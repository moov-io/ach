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

	"github.com/moov-io/ach"
)

// Example_customTraceNumber writes an ACH file with a non-standard NACHA TraceNumber
func Example_customTraceNumber() {
	fh := mockFileHeader()

	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.CreditsOnly
	bh.CompanyName = "My Company"
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.PPD
	bh.CompanyEntryDescription = "Cash Back"
	// fix EffectiveEntryDate for consistent output
	bh.EffectiveEntryDate = "190816"
	bh.ODFIIdentification = "987654320"

	entry := ach.NewEntryDetail()
	entry.TransactionCode = ach.CheckingCredit
	entry.SetRDFI("123456780")
	entry.DFIAccountNumber = "12567"
	entry.Amount = 100000000
	entry.TraceNumber = "321888"
	entry.IndividualName = "Jane Smith"

	// build the batch
	batch := ach.NewBatchPPD(bh)
	batch.SetValidation(&ach.ValidateOpts{
		CustomTraceNumbers: true,
	})
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
	if err := file.Validate(); err != nil {
		log.Fatalf("Unexpected validation error: %v", err)
	}

	if trace := file.Batches[0].GetEntries()[0].TraceNumber; trace != "321888" {
		log.Fatalf("Unexpected trace number: %s", trace)
	}

	fmt.Printf("TransactionCode=%d\n", file.Batches[0].GetEntries()[0].TransactionCode)
	fmt.Printf("%s\n", file.Batches[0].GetEntries()[0].String())

	// Output:
	// TransactionCode=22
	// 62212345678012567            0100000000               Jane Smith              0000000000321888
}
