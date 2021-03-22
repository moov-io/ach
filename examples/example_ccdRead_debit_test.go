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

func Example_ccdReadDebit() {
	// Open a file for reading, any io.Reader can be used
	f, err := os.Open(filepath.Join("testdata", "ccd-debit.ach"))
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

	fmt.Printf("Total Amount Debit: %d\n", achFile.Control.TotalDebitEntryDollarAmountInFile)
	fmt.Printf("Total Amount Credit: %d\n", achFile.Control.TotalCreditEntryDollarAmountInFile)
	fmt.Printf("SEC Code: %s\n", achFile.Batches[0].GetHeader().StandardEntryClassCode)
	fmt.Printf("CCD Entry Identification Number: %s\n", achFile.Batches[0].GetEntries()[0].IdentificationNumber)
	fmt.Printf("CCD Entry Receiving Company: %s\n", achFile.Batches[0].GetEntries()[0].IndividualName)
	fmt.Printf("CCD Entry Trace Number: %s\n", achFile.Batches[0].GetEntries()[0].TraceNumberField())
	fmt.Printf("CCD Fee Identification Number: %s\n", achFile.Batches[0].GetEntries()[1].IdentificationNumber)
	fmt.Printf("CCD Fee Receiving Company: %s\n", achFile.Batches[0].GetEntries()[1].IndividualName)
	fmt.Printf("CCD Fee Trace Number: %s\n", achFile.Batches[0].GetEntries()[1].TraceNumberField())

	// Output:
	// Total Amount Debit: 500125
	// Total Amount Credit: 0
	// SEC Code: CCD
	// CCD Entry Identification Number: location1234567
	// CCD Entry Receiving Company: Best Co. #123456789012
	// CCD Entry Trace Number: 031300010000001
	// CCD Fee Identification Number: Fee123456789012
	// CCD Fee Receiving Company: Best Co. #123456789012
	// CCD Fee Trace Number: 031300010000002
}
