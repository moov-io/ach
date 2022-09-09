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
	"log"
	"os"
	"time"

	"github.com/moov-io/ach"
)

func main() {
	// Example transfer to write an ACH AXT file acknowledging a CTX credit
	// Important: All financial institutions are different and will require registration and exact field values.

	// Set originator bank ODFI and destination Operator for the financial institution
	// this is the funding/receiving source of the transfer
	fh := ach.NewFileHeader()
	fh.ImmediateDestination = "031300012"             // Routing Number of the ACH Operator or receiving point to which the file is being sent
	fh.ImmediateOrigin = "231380104"                  // Routing Number of the ACH Operator or sending point that is sending the file
	fh.FileCreationDate = time.Now().Format("060102") // Today's Date
	fh.ImmediateDestinationName = "Federal Reserve Bank"
	fh.ImmediateOriginName = "My Bank Name"

	// BatchHeader identifies the originating entity and the type of transactions contained in the batch
	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.CreditsOnly
	bh.CompanyName = "Name on Account" // The name of the company/person that has relationship with receiver
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.ATX
	bh.CompanyEntryDescription = "Vndr Pay"                              // will be on receiving account's statement
	bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1).Format("060102") // YYMMDD
	bh.ODFIIdentification = "23138010"                                   // Originating Routing Number

	// Identifies the receiver's account information
	// can be multiple entries per batch
	entry := ach.NewEntryDetail()
	// Identifies the entry as a debit and credit entry AND to what type of account (Savings, DDA, Loan, GL)
	entry.TransactionCode = ach.CheckingZeroDollarRemittanceCredit // Code 22: Demand Debit(deposit) to checking account
	entry.SetRDFI("031300012")                                     // Receiver's bank transit routing number
	entry.DFIAccountNumber = "744-5678-99"                         // Receiver's bank account number
	entry.Amount = 0                                               // Amount of transaction with no decimal. One dollar and eleven cents = 111
	entry.SetOriginalTraceNumber("031300010000001")
	entry.SetCATXAddendaRecords(2)
	entry.SetCATXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(bh.ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.AddendaRecordIndicator = 1

	entryOne := ach.NewEntryDetail() // Fee Entry
	entryOne.TransactionCode = ach.CheckingZeroDollarRemittanceCredit
	entryOne.SetRDFI("031300012")             // Receiver's bank transit routing number
	entryOne.DFIAccountNumber = "744-5678-99" // Receiver's bank account number
	entryOne.Amount = 0                       // Amount of transaction with no decimal. One dollar and eleven cents = 111
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

	// Write the file to stdout, any io.Writer can be used
	w := ach.NewWriter(os.Stdout)
	if err := w.Write(file); err != nil {
		log.Fatalf("Unexpected error writing file: %s\n", err)
	}
	w.Flush()
}
