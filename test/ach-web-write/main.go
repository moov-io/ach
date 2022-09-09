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
	// Example transfer to write an ACH WEB file to send/credit a external institution's account
	// Important: All financial institutions are different and will require registration and exact field values.

	// Set originator bank ODFI and destination Operator for the financial institution
	// this is the funding/receiving source of the transfer
	fh := ach.NewFileHeader()
	fh.ImmediateDestination = "231380104"             // Routing Number of the ACH Operator or receiving point to which the file is being sent
	fh.ImmediateOrigin = "121042882"                  // Routing Number of the ACH Operator or sending point that is sending the file
	fh.FileCreationDate = time.Now().Format("060102") // Today's Date
	fh.ImmediateDestinationName = "Federal Reserve Bank"
	fh.ImmediateOriginName = "My Bank Name"

	// BatchHeader identifies the originating entity and the type of transactions contained in the batch
	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.CreditsOnly
	bh.CompanyName = "Name on Account" // The name of the company/person that has relationship with receiver
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.WEB
	bh.CompanyEntryDescription = "Subscribe"                             // will be on receiving account's statement
	bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1).Format("060102") // YYMMDD
	bh.ODFIIdentification = "121042882"                                  // Originating Routing Number

	entry := ach.NewEntryDetail()
	entry.TransactionCode = ach.CheckingCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "12345678"
	entry.Amount = 10000 //  100.00
	entry.IndividualName = "Wade Arnold"
	entry.SetTraceNumber(bh.ODFIIdentification, 1)
	entry.IdentificationNumber = "#789654"
	entry.DiscretionaryData = "S"
	entry.Category = ach.CategoryForward
	entry.AddendaRecordIndicator = 1

	// To add one or more optional addenda records for an entry
	addenda1 := ach.NewAddenda05()
	addenda1.PaymentRelatedInformation = "PAY-GATE payment"
	entry.AddAddenda05(addenda1)

	entry2 := ach.NewEntryDetail()
	entry2.TransactionCode = ach.CheckingCredit
	entry2.SetRDFI("231380104")
	entry2.DFIAccountNumber = "81967038518"
	entry2.Amount = 799 // 7.99
	entry2.IndividualName = "Wade Arnold"
	entry2.SetTraceNumber(bh.ODFIIdentification, 2)
	entry2.IdentificationNumber = "#123456"
	entry2.DiscretionaryData = "R"
	entry2.Category = ach.CategoryForward
	entry2.AddendaRecordIndicator = 1

	// To add one or more optional addenda records for an entry
	addenda2 := ach.NewAddenda05()
	addenda2.PaymentRelatedInformation = "Monthly Membership Subscription"
	entry2.AddAddenda05(addenda2)

	// build the batch
	batch := ach.NewBatchWEB(bh)
	batch.AddEntry(entry)
	batch.AddEntry(entry2)
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
