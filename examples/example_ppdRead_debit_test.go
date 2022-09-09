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

func Example_ppdReadDebit() {
	// Open a file for reading, any io.Reader can be used
	f, err := os.Open(filepath.Join("testdata", "ppd-debit.ach"))
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

	fmt.Printf("SEC Code: %s\n", achFile.Batches[0].GetHeader().StandardEntryClassCode)
	fmt.Printf("Total File Debit Amount: %d\n", achFile.Control.TotalDebitEntryDollarAmountInFile)
	fmt.Printf("Total File Credit Amount: %d\n", achFile.Control.TotalCreditEntryDollarAmountInFile)

	// Output:
	// SEC Code: PPD
	// Total File Debit Amount: 200000000
	// Total File Credit Amount: 0
}
