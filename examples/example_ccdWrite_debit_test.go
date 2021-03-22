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

func Example_ccdWriteDebit() {
	fh := mockFileHeader()

	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.DebitsOnly
	bh.CompanyName = "Name on Account"
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.CCD
	bh.CompanyEntryDescription = "Vndr Pay"
	bh.EffectiveEntryDate = "190816" // need EffectiveEntryDate to be fixed so it can match output
	bh.ODFIIdentification = "031300012"

	entry := ach.NewEntryDetail()
	entry.TransactionCode = ach.CheckingDebit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 500000
	entry.IdentificationNumber = "location #1234"
	entry.SetReceivingCompany("Best Co. #1")
	entry.SetTraceNumber(bh.ODFIIdentification, 1)
	entry.DiscretionaryData = "S"

	entryOne := ach.NewEntryDetail()
	entryOne.TransactionCode = ach.CheckingDebit
	entryOne.SetRDFI("231380104")
	entryOne.DFIAccountNumber = "744-5678-99"
	entryOne.Amount = 125
	entryOne.IdentificationNumber = "Fee #1"
	entryOne.SetReceivingCompany("Best Co. #1")
	entryOne.SetTraceNumber(bh.ODFIIdentification, 2)
	entryOne.DiscretionaryData = "S"

	// build the batch
	batch := ach.NewBatchCCD(bh)
	batch.AddEntry(entry)
	batch.AddEntry(entryOne)
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

	fmt.Println(file.Header.String())
	fmt.Println(file.Batches[0].GetHeader().String())
	fmt.Println(file.Batches[0].GetEntries()[0].String())
	fmt.Println(file.Batches[0].GetEntries()[1].String())
	fmt.Println(file.Batches[0].GetControl().String())
	fmt.Println(file.Control.String())

	// Output:
	// 101 031300012 2313801041908161055A094101Federal Reserve Bank   My Bank Name           12345678
	// 5225Name on Account                     231380104 CCDVndr Pay        190816   1031300010000001
	// 627231380104744-5678-99      0000500000location #1234 Best Co. #1           S 0031300010000001
	// 627231380104744-5678-99      0000000125Fee #1         Best Co. #1           S 0031300010000002
	// 82250000020046276020000000500125000000000000231380104                          031300010000001
	// 9000001000001000000020046276020000000500125000000000000
}
