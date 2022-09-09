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
	// Example transfer to write an ACH SHR file to send/credit a external institution's account
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
	bh.ServiceClassCode = ach.DebitsOnly
	bh.CompanyName = "Name on Account" // The name of the company/person that has relationship with receiver
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.SHR
	bh.CompanyEntryDescription = "Payment"                               // will be on receiving account's statement
	bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1).Format("060102") // YYMMDD
	bh.ODFIIdentification = "121042882"                                  // Originating Routing Number

	// Identifies the receiver's account information
	// can be multiple entries per batch
	entry := ach.NewEntryDetail()
	// Identifies the entry as a debit and credit entry AND to what type of account (Savings, DDA, Loan, GL)
	entry.TransactionCode = ach.CheckingDebit // Code 22: Debit (withdrawal) from checking account
	entry.SetRDFI("231380104")                // Receiver's bank transit routing number
	entry.DFIAccountNumber = "12345678"       // Receiver's bank account number
	entry.Amount = 100000000                  // Amount of transaction with no decimal. One dollar and eleven cents = 111
	entry.SetTraceNumber(bh.ODFIIdentification, 1)
	entry.SetSHRCardExpirationDate("0722")
	entry.SetSHRDocumentReferenceNumber("12345678910")
	entry.SetSHRIndividualCardAccountNumber("1234567891123456789")
	entry.DiscretionaryData = "01"
	entry.AddendaRecordIndicator = 1

	addenda02 := ach.NewAddenda02()
	addenda02.ReferenceInformationOne = "REFONEA"
	addenda02.ReferenceInformationTwo = "REF"
	addenda02.TerminalIdentificationCode = "TERM02"
	addenda02.TransactionSerialNumber = "100049"
	addenda02.TransactionDate = "0614"
	addenda02.AuthorizationCodeOrExpireDate = "123456"
	addenda02.TerminalLocation = "Target Store 0049"
	addenda02.TerminalCity = "PHILADELPHIA"
	addenda02.TerminalState = "PA"
	addenda02.TraceNumber = "121042880000001"

	// build the batch
	batch := ach.NewBatchSHR(bh)
	batch.AddEntry(entry)
	batch.GetEntries()[0].Addenda02 = addenda02
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
