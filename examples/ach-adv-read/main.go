// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/moov-io/ach"
)

func main() {
	f, err := os.Open(filepath.Join("test", "ach-adv-read", "adv-read.ach"))
	if err != nil {
		log.Fatal(err)
	}
	r := ach.NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		fmt.Printf("Issue reading file: %+v \n", err)
	}
	// ensure we have a validated file structure
	if achFile.Validate(); err != nil {
		fmt.Printf("Could not validate entire read file: %v", err)
	}
	// If you trust the file but it's formatting is off building will probably resolve the malformed file.
	if err := achFile.Create(); err != nil {
		fmt.Printf("Could not create file with read properties: %v", err)
	}

	fmt.Printf("Credit Total Amount: %v \n", achFile.ADVControl.TotalCreditEntryDollarAmountInFile)
	fmt.Printf("Debit Total Amount: %v \n", achFile.ADVControl.TotalDebitEntryDollarAmountInFile)
	fmt.Printf("OriginatorStatusCode: %v \n", achFile.Batches[0].GetHeader().OriginatorStatusCode)
	fmt.Printf("Batch Credit Total Amount: %v \n", achFile.Batches[0].GetADVControl().TotalCreditEntryDollarAmount)
	fmt.Printf("Batch Debit Total Amount: %v \n", achFile.Batches[0].GetADVControl().TotalDebitEntryDollarAmount)
	fmt.Printf("SEC Code: %v \n", achFile.Batches[0].GetHeader().StandardEntryClassCode)
	fmt.Printf("Entry Amount: %v \n", achFile.Batches[0].GetADVEntries()[0].Amount)
	fmt.Printf("Sequence Number: %v \n", achFile.Batches[0].GetADVEntries()[0].SequenceNumber)
	fmt.Printf("EntryOne Amount: %v \n", achFile.Batches[0].GetADVEntries()[1].Amount)
	fmt.Printf("EntryOne Sequence Number: %v \n", achFile.Batches[0].GetADVEntries()[1].SequenceNumber)
}
