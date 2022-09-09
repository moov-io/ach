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

package examples

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/moov-io/ach"
)

func Example_advRead() {
	// Open a file for reading, any io.Reader can be used
	f, err := os.Open(filepath.Join("testdata", "adv-read.ach"))
	if err != nil {
		log.Fatalln(err)
	}
	r := ach.NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		log.Fatalf("reading file: %v\n", err)
	}
	// If you trust the file but its formatting is off, building will probably resolve the malformed file
	if err := achFile.Create(); err != nil {
		log.Fatalf("creating file: %v\n", err)
	}
	// Validate the ACH file
	if err := achFile.Validate(); err != nil {
		log.Fatalf("validating file: %v\n", err)
	}

	fmt.Printf("Credit Total Amount: %d\n", achFile.ADVControl.TotalCreditEntryDollarAmountInFile)
	fmt.Printf("Debit Total Amount: %d\n", achFile.ADVControl.TotalDebitEntryDollarAmountInFile)
	fmt.Printf("OriginatorStatusCode: %d\n", achFile.Batches[0].GetHeader().OriginatorStatusCode)
	fmt.Printf("Batch Credit Total Amount: %d\n", achFile.Batches[0].GetADVControl().TotalCreditEntryDollarAmount)
	fmt.Printf("Batch Debit Total Amount: %d\n", achFile.Batches[0].GetADVControl().TotalDebitEntryDollarAmount)
	fmt.Printf("SEC Code: %s\n", achFile.Batches[0].GetHeader().StandardEntryClassCode)
	fmt.Printf("Entry Amount: %d\n", achFile.Batches[0].GetADVEntries()[0].Amount)
	fmt.Printf("Sequence Number: %d\n", achFile.Batches[0].GetADVEntries()[0].SequenceNumber)
	fmt.Printf("EntryOne Amount: %d\n", achFile.Batches[0].GetADVEntries()[1].Amount)
	fmt.Printf("EntryOne Sequence Number: %d\n", achFile.Batches[0].GetADVEntries()[1].SequenceNumber)

	// Output:
	// Credit Total Amount: 50000
	// Debit Total Amount: 250000
	// OriginatorStatusCode: 0
	// Batch Credit Total Amount: 50000
	// Batch Debit Total Amount: 250000
	// SEC Code: ADV
	// Entry Amount: 50000
	// Sequence Number: 1
	// EntryOne Amount: 250000
	// EntryOne Sequence Number: 2
}
