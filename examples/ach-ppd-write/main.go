// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"os"
	"time"

	"github.com/moov-io/ach"
)

func main() {
	// Example transfer to write an ACH PPD file to send/credit a external institutions account
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
	bh.ServiceClassCode = ach.MixedDebitsAndCredits
	bh.CompanyName = "Name on Account" // The name of the company/person that has relationship with receiver
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.PPD
	bh.CompanyEntryDescription = "REG.SALARY"                            // will be on receiving accounts statement
	bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1).Format("060102") // YYMMDD
	bh.ODFIIdentification = "121042882"                                  // Originating Routing Number

	// Identifies the receivers account information
	// can be multiple entry's per batch
	entryOne := ach.NewEntryDetail()
	// Identifies the entry as a debit and credit entry AND to what type of account (Savings, DDA, Loan, GL)
	entryOne.TransactionCode = ach.CheckingDebit
	entryOne.SetRDFI("231380104")           // Receivers bank transit routing number
	entryOne.DFIAccountNumber = "123456789" // Receivers bank account number
	entryOne.Amount = 200000000             // Amount of transaction with no decimal. One dollar and eleven cents = 111
	entryOne.SetTraceNumber(bh.ODFIIdentification, 1)
	entryOne.IndividualName = "Debit Account" // Identifies the receiver of the transaction

	entryTwo := ach.NewEntryDetail()
	// Identifies the entry as a debit and credit entry AND to what type of account (Savings, DDA, Loan, GL)
	entryTwo.TransactionCode = ach.CheckingCredit
	entryTwo.SetRDFI("231380104")           // Receivers bank transit routing number
	entryTwo.DFIAccountNumber = "987654321" // Receivers bank account number
	entryTwo.Amount = 100000000             // Amount of transaction with no decimal. One dollar and eleven cents = 111
	entryTwo.SetTraceNumber(bh.ODFIIdentification, 2)
	entryTwo.IndividualName = "Credit Account 1" // Identifies the receiver of the transaction

	entryThree := ach.NewEntryDetail()
	// Identifies the entry as a debit and credit entry AND to what type of account (Savings, DDA, Loan, GL)
	entryThree.TransactionCode = ach.CheckingCredit
	entryThree.SetRDFI("231380104")           // Receivers bank transit routing number
	entryThree.DFIAccountNumber = "837098765" // Receivers bank account number
	entryThree.Amount = 100000000             // Amount of transaction with no decimal. One dollar and eleven cents = 111
	entryThree.SetTraceNumber(bh.ODFIIdentification, 3)
	entryThree.IndividualName = "Credit Account 2" // Identifies the receiver of the transaction

	// build the batch
	batch := ach.NewBatchPPD(bh)
	batch.AddEntry(entryOne)
	batch.AddEntry(entryTwo)
	batch.AddEntry(entryThree)
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

	// write the file to std out. Anything io.Writer
	w := ach.NewWriter(os.Stdout)
	if err := w.Write(file); err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}
	w.Flush()
}
