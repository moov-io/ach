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

func Example_advWrite() {
	fh := mockFileHeader()

	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.AutomatedAccountingAdvices
	bh.CompanyName = "Company Name, Inc"
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.ADV
	bh.CompanyEntryDescription = "Accounting"
	bh.EffectiveEntryDate = "190816" // need EffectiveEntryDate to be fixed so it can match output
	bh.ODFIIdentification = "121042882"
	bh.OriginatorStatusCode = 0

	entry := ach.NewADVEntryDetail()
	entry.TransactionCode = ach.CreditForDebitsOriginated
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 50000
	entry.AdviceRoutingNumber = "121042882"
	entry.FileIdentification = "11131"
	entry.ACHOperatorData = ""
	entry.IndividualName = "Name"
	entry.DiscretionaryData = ""
	entry.AddendaRecordIndicator = 0
	entry.ACHOperatorRoutingNumber = "01100001"
	entry.JulianDay = 50
	entry.SequenceNumber = 1

	entryOne := ach.NewADVEntryDetail()
	entryOne.TransactionCode = ach.DebitForCreditsOriginated
	entryOne.SetRDFI("231380104")
	entryOne.DFIAccountNumber = "744-5678-99"
	entryOne.Amount = 250000
	entryOne.AdviceRoutingNumber = "121042882"
	entryOne.FileIdentification = "11139"
	entryOne.ACHOperatorData = ""
	entryOne.IndividualName = "Name"
	entryOne.DiscretionaryData = ""
	entryOne.AddendaRecordIndicator = 0
	entryOne.ACHOperatorRoutingNumber = "01100001"
	entryOne.JulianDay = 50
	entryOne.SequenceNumber = 2

	// build the batch
	batch := ach.NewBatchADV(bh)
	batch.AddADVEntry(entry)
	batch.AddADVEntry(entryOne)
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
	fmt.Println(file.Batches[0].GetADVEntries()[0].String())
	fmt.Println(file.Batches[0].GetADVEntries()[1].String())
	fmt.Println(file.Batches[0].GetADVControl().String())
	fmt.Println(file.Control.String())

	// Output:
	// 101 031300012 2313801041908161055A094101Federal Reserve Bank   My Bank Name           12345678
	// 5280Company Name, In                    231380104 ADVAccounting      190816   0121042880000001
	// 681231380104744-5678-99    00000005000012104288211131 Name                    0011000010500001
	// 682231380104744-5678-99    00000025000012104288211139 Name                    0011000010500002
	// 828000000200462760200000000000000025000000000000000000050000Company Name, Inc  121042880000001
	// 9000000000000000000000000000000000000000000000000000000
}
