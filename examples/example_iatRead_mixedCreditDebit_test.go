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

func Example_iatReadMixedCreditDebit() {
	// Open a file for reading, any io.Reader can be used
	f, err := os.Open(filepath.Join("testdata", "iat-mixedCreditDebit.ach"))
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

	fmt.Printf("SEC Code: %s\n", achFile.IATBatches[0].GetHeader().StandardEntryClassCode)
	fmt.Printf("Debit Entry: %s\n", achFile.IATBatches[0].Entries[0].String())
	fmt.Printf("Addenda10: %s\n", achFile.IATBatches[0].Entries[0].Addenda10.String())
	fmt.Printf("Addenda11: %s\n", achFile.IATBatches[0].Entries[0].Addenda11.String())
	fmt.Printf("Addenda12: %s\n", achFile.IATBatches[0].Entries[0].Addenda12.String())
	fmt.Printf("Addenda13: %s\n", achFile.IATBatches[0].Entries[0].Addenda13.String())
	fmt.Printf("Addenda14: %s\n", achFile.IATBatches[0].Entries[0].Addenda14.String())
	fmt.Printf("Addenda15: %s\n", achFile.IATBatches[0].Entries[0].Addenda15.String())
	fmt.Printf("Addenda16: %s\n", achFile.IATBatches[0].Entries[0].Addenda16.String())
	fmt.Printf("Addenda17: %s\n", achFile.IATBatches[0].Entries[0].Addenda17[0].String())
	fmt.Printf("Addenda18: %s\n", achFile.IATBatches[0].Entries[0].Addenda18[0].String())
	fmt.Printf("Total File Debit Amount: %d\n", achFile.Control.TotalDebitEntryDollarAmountInFile)
	fmt.Printf("Credit Entry: %s\n", achFile.IATBatches[0].Entries[1].String())
	fmt.Printf("Addenda10: %s\n", achFile.IATBatches[0].Entries[1].Addenda10.String())
	fmt.Printf("Addenda11: %s\n", achFile.IATBatches[0].Entries[1].Addenda11.String())
	fmt.Printf("Addenda12: %s\n", achFile.IATBatches[0].Entries[1].Addenda12.String())
	fmt.Printf("Addenda13: %s\n", achFile.IATBatches[0].Entries[1].Addenda13.String())
	fmt.Printf("Addenda14: %s\n", achFile.IATBatches[0].Entries[1].Addenda14.String())
	fmt.Printf("Addenda15: %s\n", achFile.IATBatches[0].Entries[1].Addenda15.String())
	fmt.Printf("Addenda16: %s\n", achFile.IATBatches[0].Entries[1].Addenda16.String())
	fmt.Printf("Addenda17: %s\n", achFile.IATBatches[0].Entries[1].Addenda17[0].String())
	fmt.Printf("Addenda18: %s\n", achFile.IATBatches[0].Entries[1].Addenda18[0].String())
	fmt.Printf("Total File Credit Amount: %d\n", achFile.Control.TotalCreditEntryDollarAmountInFile)

	// Output:
	// SEC Code: IAT
	// Debit Entry: 6271210428820007             0000100000123456789                              1231380100000001
	// Addenda10: 710ANN000000000000100000928383-23938          BEK Enterprises                          0000001
	// Addenda11: 711BEK Solutions                      15 West Place Street                             0000001
	// Addenda12: 712JacobsTown*PA\                     US*19305\                                        0000001
	// Addenda13: 713Wells Fargo                        01231380104                         US           0000001
	// Addenda14: 714Citadel Bank                       01121042882                         CA           0000001
	// Addenda15: 7159874654932139872121 Front Street                                                    0000001
	// Addenda16: 716LetterTown*AB\                     CA*80014\                                        0000001
	// Addenda17: 717This is an international payment                                                00010000001
	// Addenda18: 718Bank of France                     01456456456987987                   FR       00010000001
	// Total File Debit Amount: 100000
	// Credit Entry: 6221210428820007             0000100000123456789                              1231380100000002
	// Addenda10: 710ANN000000000000100000928383-23938          ADCAF Enterprises                        0000002
	// Addenda11: 711ADCAF Solutions                    15 West Place Street                             0000002
	// Addenda12: 712JacobsTown*PA\                     US*19305\                                        0000002
	// Addenda13: 713Wells Fargo                        01231380104                         US           0000002
	// Addenda14: 714Citadel Bank                       01121042882                         CA           0000002
	// Addenda15: 71598746549321398718 Fifth Street                                                      0000002
	// Addenda16: 716LetterTown*AB\                     CA*80014\                                        0000002
	// Addenda17: 717This is an international payment                                                00010000002
	// Addenda18: 718Bank of France                     01456456456987987                   FR       00010000002
	// Total File Credit Amount: 100000
}
