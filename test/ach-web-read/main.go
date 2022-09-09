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
	f, err := os.Open("web-credit.ach")
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

	fmt.Printf("Total Amount Credit: %d\n", achFile.Control.TotalCreditEntryDollarAmountInFile)
	fmt.Printf("SEC Code: %s\n", achFile.Batches[0].GetHeader().StandardEntryClassCode)
	fmt.Printf("Entry One : %s\n", achFile.Batches[0].GetEntries()[0])
	fmt.Printf("Entry One Addenda Record Indicator: %d\n", achFile.Batches[0].GetEntries()[0].AddendaRecordIndicator)
	fmt.Printf("Entry One Addenda: %s\n", achFile.Batches[0].GetEntries()[0].Addenda05[0])
	eadOne := new(ach.Addenda05)
	EntryOneAddenda := achFile.Batches[0].GetEntries()[0].Addenda05[0].String()
	eadOne.Parse(EntryOneAddenda)
	fmt.Printf("Entry One Addenda Type Code: %s\n", string(eadOne.TypeCode))
	fmt.Printf("Entry One Addenda Payment Related Information: %s\n", eadOne.PaymentRelatedInformation)
	fmt.Printf("Entry Two: %s\n", achFile.Batches[0].GetEntries()[1])
	fmt.Printf("Entry Two Addenda Record Indicator: %d\n", achFile.Batches[0].GetEntries()[1].AddendaRecordIndicator)
	fmt.Printf("Entry Two Addenda: %s\n", achFile.Batches[0].GetEntries()[1].Addenda05[0])
	eadTwo := new(ach.Addenda05)
	EntryTwoAddenda := achFile.Batches[0].GetEntries()[1].Addenda05[0].String()
	eadTwo.Parse(EntryTwoAddenda)
	fmt.Printf("Entry Two Addenda Payment Related Information: %s\n", eadTwo.PaymentRelatedInformation)
	fmt.Printf("Entry One Addenda Type Code: %s\n", string(eadTwo.TypeCode))
}
