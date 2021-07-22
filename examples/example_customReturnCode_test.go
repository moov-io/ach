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
	"strings"

	"github.com/moov-io/ach"
)

// Example_customReturnCode writes an ACH file with a non-standard Return Code
func Example_customReturnCode() {
	fh := mockFileHeader()

	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.CreditsOnly
	bh.CompanyName = "My Company"
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.PPD
	bh.CompanyEntryDescription = "Cash Back"
	bh.EffectiveEntryDate = "190816" // need EffectiveEntryDate to be fixed so it can match output
	bh.ODFIIdentification = "987654320"

	entry := ach.NewEntryDetail()
	entry.TransactionCode = 22 // example of a custom code
	entry.SetRDFI("123456780")
	entry.DFIAccountNumber = "12567"
	entry.Amount = 100000000
	entry.SetTraceNumber(bh.ODFIIdentification, 2)
	entry.IndividualName = "Jane Smith"
	addenda99 := ach.NewAddenda99()
	addenda99.ReturnCode = "abc"
	entry.Addenda99 = addenda99
	entry.Category = ach.CategoryReturn
	entry.AddendaRecordIndicator = 1
	entry.Addenda99.SetValidation(&ach.ValidateOpts{
		CustomReturnCodes: true,
	})

	// build the batch
	batch := ach.NewBatchPPD(bh)
	batch.AddEntry(entry)
	if err := batch.Create(); err != nil {
		log.Fatalf("Unexpected error building batch: %s\n", err)
	}

	// build the file
	file := ach.NewFile()
	file.SetHeader(fh)
	file.AddBatch(batch)
	if err := file.Create(); err != nil {
		log.Fatalf("Unexpected error building file: %s\n", err)
	}

	fmt.Printf("TransactionCode=%d\n", file.Batches[0].GetEntries()[0].TransactionCode)
	fmt.Println(file.Batches[0].GetEntries()[0].String())
	fmt.Printf("ReturnCode=%s\n", strings.TrimSpace(file.Batches[0].GetEntries()[0].Addenda99.ReturnCode))

	// Output:
	// TransactionCode=22
	// 62212345678012567            0100000000               Jane Smith              1987654320000002
	// ReturnCode=abc
}
