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

func Example_ppdReadSegmentFile() {
	// Open a file for reading, any io.Reader can be used
	fCredit, err := os.Open(filepath.Join("testdata", "segmentFile-ppd-credit.ach"))
	if err != nil {
		log.Fatalln(err)
	}
	rCredit := ach.NewReader(fCredit)
	achFileCredit, err := rCredit.Read()
	if err != nil {
		log.Fatalf("reading file: %v\n", err)
	}
	// If you trust the file but its formatting is off, building will probably resolve the malformed file
	if err := achFileCredit.Create(); err != nil {
		log.Fatalf("creating file: %v\n", err)
	}
	// Validate the ACH file
	if err := achFileCredit.Validate(); err != nil {
		log.Fatalf("validating file: %v\n", err)
	}

	// Open a file for reading, any io.Reader can be used
	fDebit, err := os.Open(filepath.Join("testdata", "segmentFile-ppd-debit.ach"))
	if err != nil {
		log.Fatalln(err)
	}
	rDebit := ach.NewReader(fDebit)
	achFileDebit, err := rDebit.Read()
	if err != nil {
		log.Fatalf("reading file: %v\n", err)
	}
	// If you trust the file but its formatting is off, building will probably resolve the malformed file
	if err := achFileDebit.Create(); err != nil {
		log.Fatalf("creating file: %v\n", err)
	}
	// Validate the ACH file
	if err := achFileDebit.Validate(); err != nil {
		log.Fatalf("validating file: %v\n", err)
	}

	fmt.Printf("Total Credit Amount: %d\n", achFileCredit.Control.TotalCreditEntryDollarAmountInFile)
	fmt.Printf("SEC Code: %s\n", achFileCredit.Batches[0].GetHeader().StandardEntryClassCode)
	fmt.Printf("Total Debit Amount: %d\n", achFileDebit.Control.TotalDebitEntryDollarAmountInFile)
	fmt.Printf("SEC Code: %s\n", achFileDebit.Batches[0].GetHeader().StandardEntryClassCode)

	// Output:
	// Total Credit Amount: 200000000
	// SEC Code: PPD
	// Total Debit Amount: 200000000
	// SEC Code: PPD
}
