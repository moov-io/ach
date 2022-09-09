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

// Example_customTransactionCode writes an ACH file with a non-standard NACHA TransactionCodes
func Example_customTransactionCode() {
	validationOpts := &ach.ValidateOpts{
		// CheckTransactionCode lets us override the standard set of NACHA
		// codes for applications which use custom TransactionCode values.
		CheckTransactionCode: func(code int) error {
			// Is it a custom TransactionCode of ours?
			if code == 78 {
				return nil
			}
			// Is it a NACHA standard code?
			return ach.StandardTransactionCode(code)
		},
	}

	fh := mockFileHeader()

	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.CreditsOnly
	bh.CompanyName = "My Company"
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.PPD
	bh.CompanyEntryDescription = "Cash Back"
	bh.EffectiveEntryDate = "190816" // need EffectiveEntryDate to be fixed so it can match output
	bh.ODFIIdentification = "987654320"
	bh.SetValidation(validationOpts)

	entry := ach.NewEntryDetail()
	entry.TransactionCode = 78 // example of a custom code
	entry.SetRDFI("123456780")
	entry.DFIAccountNumber = "12567"
	entry.Amount = 100000000
	entry.SetTraceNumber(bh.ODFIIdentification, 2)
	entry.IndividualName = "Jane Smith"
	entry.SetValidation(validationOpts)

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

	// Output:
	// TransactionCode=78
	// 67812345678012567            0100000000               Jane Smith              0987654320000002
}
