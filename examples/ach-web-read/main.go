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
	"github.com/moov-io/ach"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("test", "ach-web-read", "web-credit.ach"))
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

	fmt.Printf("Total Amount Credit: %v \n", achFile.Control.TotalCreditEntryDollarAmountInFile)
	fmt.Printf("SEC Code: %v \n", achFile.Batches[0].GetHeader().StandardEntryClassCode)
	fmt.Printf("Entry One : %v \n", achFile.Batches[0].GetEntries()[0])
	fmt.Printf("Entry One Addenda Record Indicator: %v \n", achFile.Batches[0].GetEntries()[0].AddendaRecordIndicator)
	fmt.Printf("Entry One Addenda: %v \n", achFile.Batches[0].GetEntries()[0].Addenda05[0])
	eadOne := new(ach.Addenda05)
	EntryOneAddenda := achFile.Batches[0].GetEntries()[0].Addenda05[0].String()
	eadOne.Parse(EntryOneAddenda)
	fmt.Printf("Entry One Addenda Type Code: %v \n", string(eadOne.TypeCode))
	fmt.Printf("Entry One Addenda Payment Related Information: %v \n", eadOne.PaymentRelatedInformation)
	fmt.Printf("Entry Two: %v \n", achFile.Batches[0].GetEntries()[1])
	fmt.Printf("Entry Two Addenda Record Indicator: %v \n", achFile.Batches[0].GetEntries()[1].AddendaRecordIndicator)
	fmt.Printf("Entry Two Addenda: %v \n", achFile.Batches[0].GetEntries()[1].Addenda05[0])
	eadTwo := new(ach.Addenda05)
	EntryTwoAddenda := achFile.Batches[0].GetEntries()[1].Addenda05[0].String()
	eadTwo.Parse(EntryTwoAddenda)
	fmt.Printf("Entry Two Addenda Payment Related Information: %v \n", eadTwo.PaymentRelatedInformation)
	fmt.Printf("Entry One Addenda Type Code: %v \n", string(eadTwo.TypeCode))
}
