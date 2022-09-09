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

func Example_atxWrite() {
	fh := mockFileHeader()

	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.CreditsOnly
	bh.CompanyName = "Name on Account"
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.ATX
	bh.CompanyEntryDescription = "Vndr Pay"
	bh.EffectiveEntryDate = "190816" // need EffectiveEntryDate to be fixed so it can match output
	bh.ODFIIdentification = "23138010"

	entry := ach.NewEntryDetail()
	entry.TransactionCode = ach.CheckingZeroDollarRemittanceCredit
	entry.SetRDFI("031300012")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 0
	entry.SetOriginalTraceNumber("031300010000001")
	entry.SetCATXAddendaRecords(2)
	entry.SetCATXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(bh.ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.AddendaRecordIndicator = 1

	entryOne := ach.NewEntryDetail() // Fee Entry
	entryOne.TransactionCode = ach.CheckingZeroDollarRemittanceCredit
	entryOne.SetRDFI("031300012")
	entryOne.DFIAccountNumber = "744-5678-99"
	entryOne.Amount = 0
	entryOne.SetOriginalTraceNumber("031300010000002")
	entryOne.SetCATXAddendaRecords(2)
	entryOne.SetCATXReceivingCompany("Receiver Company")
	entryOne.SetTraceNumber(bh.ODFIIdentification, 2)
	entryOne.DiscretionaryData = "01"
	entryOne.AddendaRecordIndicator = 1

	entryAd1 := ach.NewAddenda05()
	entryAd1.PaymentRelatedInformation = "Credit account 1 for service"
	entryAd1.SequenceNumber = 1
	entryAd1.EntryDetailSequenceNumber = 0000001

	entryAd2 := ach.NewAddenda05()
	entryAd2.PaymentRelatedInformation = "Credit account 2 for service"
	entryAd2.SequenceNumber = 2
	entryAd2.EntryDetailSequenceNumber = 0000001

	entryOneAd1 := ach.NewAddenda05()
	entryOneAd1.PaymentRelatedInformation = "Credit account 1 for leadership"
	entryOneAd1.SequenceNumber = 1
	entryOneAd1.EntryDetailSequenceNumber = 0000002

	entryOneAd2 := ach.NewAddenda05()
	entryOneAd2.PaymentRelatedInformation = "Credit account 2 for leadership"
	entryOneAd2.SequenceNumber = 2
	entryOneAd2.EntryDetailSequenceNumber = 0000002

	// build the batch
	batch := ach.NewBatchATX(bh)
	batch.AddEntry(entry)
	batch.GetEntries()[0].AddAddenda05(entryAd1)
	batch.GetEntries()[0].AddAddenda05(entryAd2)
	batch.AddEntry(entryOne)
	batch.GetEntries()[1].AddAddenda05(entryOneAd1)
	batch.GetEntries()[1].AddAddenda05(entryOneAd2)
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
	fmt.Println(file.Batches[0].GetEntries()[0].Addenda05[0].String())
	fmt.Println(file.Batches[0].GetEntries()[1].String())
	fmt.Println(file.Batches[0].GetEntries()[1].Addenda05[0].String())
	fmt.Println(file.Batches[0].GetControl().String())
	fmt.Println(file.Control.String())

	// Output:
	// 101 031300012 2313801041908161055A094101Federal Reserve Bank   My Bank Name           12345678
	// 5220Name on Account                     231380104 ATXVndr Pay        190816   1231380100000001
	// 624031300012744-5678-99      00000000000313000100000010002Receiver Company  011231380100000001
	// 705Credit account 1 for service                                                    00010000001
	// 624031300012744-5678-99      00000000000313000100000020002Receiver Company  011231380100000002
	// 705Credit account 1 for leadership                                                 00010000002
	// 82200000060006260002000000000000000000000000231380104                          231380100000001
	// 9000001000001000000060006260002000000000000000000000000
}
