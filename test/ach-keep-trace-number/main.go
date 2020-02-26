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
	"log"
	"os"
	"time"

	"github.com/moov-io/ach"
)

func main() {
	// To create a file
	fh := ach.NewFileHeader()
	fh.ImmediateDestination = "231380104"
	fh.ImmediateOrigin = "121042882"
	fh.FileCreationDate = time.Now().Format("060102")
	fh.ImmediateDestinationName = "Federal Reserve Bank"
	fh.ImmediateOriginName = "My Bank Name"
	file := ach.NewFile()
	file.SetHeader(fh)

	// To create a batch.
	// Errors only if payment type is not supported.
	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.MixedDebitsAndCredits
	bh.CompanyName = "Your Company"
	bh.CompanyIdentification = file.Header.ImmediateOrigin
	bh.StandardEntryClassCode = ach.PPD
	bh.CompanyEntryDescription = "Trans. Description"
	bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1).Format("060102") // YYMMDD
	bh.ODFIIdentification = "121042882"

	batch, _ := ach.NewBatch(bh)

	// To create an entry
	entry := ach.NewEntryDetail()
	entry.TransactionCode = ach.CheckingCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "81967038518"
	entry.Amount = 1000000
	entry.IndividualName = "Wade Arnold"
	entry.SetTraceNumber("12345678", 1)
	entry.IdentificationNumber = "ABC##jvkdjfuiwn"
	entry.Category = ach.CategoryForward
	entry.AddendaRecordIndicator = 1

	// To add one or more optional addenda records for an entry

	addenda := ach.NewAddenda05()
	addenda.PaymentRelatedInformation = "bonus pay for amazing work on #OSS"
	entry.AddAddenda05(addenda)

	// Entries are added to batches like so:

	batch.AddEntry(entry)

	// When all of the Entries are added to the batch we must create it.

	if err := batch.Create(); err != nil {
		fmt.Printf("%T: %s", err, err)
	}

	// And batches are added to files much the same way:

	file.AddBatch(batch)

	// Now add a new batch for accepting payments on the web

	bh2 := ach.NewBatchHeader()
	bh2.ServiceClassCode = ach.CreditsOnly
	bh2.CompanyName = "Your Company"
	bh2.CompanyIdentification = file.Header.ImmediateOrigin
	bh2.StandardEntryClassCode = ach.WEB
	bh2.CompanyEntryDescription = "Subscribe"
	bh2.EffectiveEntryDate = time.Now().AddDate(0, 0, 1).Format("060102") // YYMMDD
	bh2.ODFIIdentification = "121042882"

	batch2, _ := ach.NewBatch(bh2)

	// Add an entry and define if it is a single or recurring payment
	// The following is a recurring payment for $7.99

	entry2 := ach.NewEntryDetail()
	entry2.TransactionCode = ach.CheckingCredit
	entry2.SetRDFI("231380104")
	entry2.DFIAccountNumber = "81967038518"
	entry2.Amount = 799
	entry2.IndividualName = "Wade Arnold"
	entry2.SetTraceNumber("87654321", 2)
	entry2.IdentificationNumber = "#123456"
	entry2.DiscretionaryData = "R"
	entry2.Category = ach.CategoryForward
	entry2.AddendaRecordIndicator = 1

	// To add one or more optional addenda records for an entry
	addenda2 := ach.NewAddenda05()
	addenda2.PaymentRelatedInformation = "Monthly Membership Subscription"
	entry2.AddAddenda05(addenda2)

	// add the entry to the batch
	batch2.AddEntry(entry2)

	// Create and add the second batch
	if err := batch2.Create(); err != nil {
		fmt.Printf("%T: %s", err, err)
	}
	file.AddBatch(batch2)

	// Once we added all our batches we must create the file

	if err := file.Create(); err != nil {
		fmt.Printf("%T: %s", err, err)
	}
	// Check if the trace number was kept
	batch1Entries := file.Batches[0].GetEntries()

	if batch1Entries[0].TraceNumber != "123456780000001" {
		log.Fatal("TraceNumber was not kept " + batch1Entries[0].TraceNumber)
	}

	batch2Entries := file.Batches[1].GetEntries()

	if batch2Entries[0].TraceNumber != "876543210000002" {
		log.Fatal("TraceNumber was not kept " + batch2Entries[0].TraceNumber)
	}

	// Finally we wnt to write the file to an io.Writer
	w := ach.NewWriter(os.Stdout)
	if err := w.Write(file); err != nil {
		fmt.Printf("%T: %s", err, err)
	}
	w.Flush()
}
