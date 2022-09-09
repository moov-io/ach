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

	"github.com/moov-io/ach"
)

func main() {
	// Open a file for reading, any io.Reader can be used
	f, err := os.Open("iat-credit.ach")
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

	fmt.Printf("Total File Amount Credit: %d\n", achFile.Control.TotalCreditEntryDollarAmountInFile)
	fmt.Printf("Total Batch Amount Credit: %d\n", achFile.IATBatches[0].Control.TotalCreditEntryDollarAmount)
	fmt.Printf("SEC Code: %s\n", achFile.IATBatches[0].GetHeader().StandardEntryClassCode)
	fmt.Printf("Entry: %s\n", achFile.IATBatches[0].GetEntries()[0])
	fmt.Printf("Entry Amount: %d\n", achFile.IATBatches[0].GetEntries()[0].Amount)
	fmt.Printf("Addenda Record Indicator: %d\n", achFile.IATBatches[0].GetEntries()[0].AddendaRecordIndicator)
	fmt.Printf("Addenda10: %s\n", achFile.IATBatches[0].GetEntries()[0].Addenda10)
	fmt.Printf("Addenda11: %s\n", achFile.IATBatches[0].GetEntries()[0].Addenda11)
	fmt.Printf("Addenda12: %s\n", achFile.IATBatches[0].GetEntries()[0].Addenda12)
	fmt.Printf("Addenda13: %s\n", achFile.IATBatches[0].GetEntries()[0].Addenda13)
	fmt.Printf("Addenda14: %s\n", achFile.IATBatches[0].GetEntries()[0].Addenda14)
	fmt.Printf("Addenda15: %s\n", achFile.IATBatches[0].GetEntries()[0].Addenda15)
	fmt.Printf("Addenda16: %s\n", achFile.IATBatches[0].GetEntries()[0].Addenda16)
	fmt.Printf("Addenda17: %s\n", achFile.IATBatches[0].GetEntries()[0].Addenda17[0].String())
	fmt.Printf("Addenda18: %s\n", achFile.IATBatches[0].GetEntries()[0].Addenda18[0].String())
}
