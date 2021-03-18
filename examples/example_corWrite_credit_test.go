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

func Example_corWriteCredit() {
	fh := mockFileHeader()

	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.CreditsOnly
	bh.StandardEntryClassCode = ach.COR
	bh.CompanyName = "Your Company, inc"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "Vendor Pay"
	bh.ODFIIdentification = "121042882" // Originating Routing Number

	entry := ach.NewEntryDetail()
	entry.TransactionCode = ach.CheckingReturnNOCCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 0
	entry.IdentificationNumber = "location #23"
	entry.SetReceivingCompany("Best Co. #23")
	entry.SetTraceNumber(bh.ODFIIdentification, 1)
	entry.AddendaRecordIndicator = 1

	addenda98 := ach.NewAddenda98()
	addenda98.ChangeCode = "C01"
	addenda98.OriginalTrace = "121042880000001"
	addenda98.OriginalDFI = "121042882"
	addenda98.CorrectedData = "1918171614"
	addenda98.TraceNumber = "91012980000088"

	entry.Addenda98 = addenda98
	entry.Category = ach.CategoryNOC

	// build the batch
	batch := ach.NewBatchCOR(bh)
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

	fmt.Println(file.Header.String())
	fmt.Println(file.Batches[0].GetHeader().String())
	fmt.Println(file.Batches[0].GetEntries()[0].String())
	fmt.Println(file.Batches[0].GetEntries()[0].Addenda98.String())
	fmt.Println(file.Batches[0].GetControl().String())
	fmt.Println(file.Control.String())

	// Output:
	// 101 031300012 2313801041908161055A094101Federal Reserve Bank   My Bank Name           12345678
	// 5220Your Company, in                    121042882 CORVendor Pay      000000   1121042880000001
	// 621231380104744-5678-99      0000000000location #23   Best Co. #23            1121042880000001
	// 798C01121042880000001      121042881918171614                                  091012980000088
	// 82200000020023138010000000000000000000000000121042882                          121042880000001
	// 9000001000001000000020023138010000000000000000000000000
}
