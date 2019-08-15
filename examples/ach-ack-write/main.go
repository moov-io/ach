// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"github.com/moov-io/ach"
	"log"
	"os"
	"time"
)

func main() {
	// Example transfer to write an ACH ACK file acknowledging a credit

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
	bh.StandardEntryClassCode = ach.ACK
	bh.CompanyEntryDescription = "Vndr Pay"                              // will be on receiving accounts statement
	bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1).Format("060102") // YYMMDD
	bh.ODFIIdentification = "23138010"                                   // Originating Routing Number

	// Identifies the receivers account information
	// can be multiple entry's per batch
	entry := ach.NewEntryDetail()
	// Identifies the entry as a debit and credit entry AND to what type of account (Savings, DDA, Loan, GL)
	entry.TransactionCode = ach.CheckingZeroDollarRemittanceCredit
	entry.SetRDFI("031300012")             // Receivers bank transit routing number
	entry.DFIAccountNumber = "744-5678-99" // Receivers bank account number
	entry.Amount = 0                       // Amount of transaction with no decimal. Zero Dollar Amount
	entry.SetOriginalTraceNumber("031300010000001")
	entry.SetReceivingCompany("Best. #1")
	entry.SetTraceNumber(bh.ODFIIdentification, 1)

	entryOne := ach.NewEntryDetail() // Fee Entry
	entryOne.TransactionCode = ach.CheckingZeroDollarRemittanceCredit
	entryOne.SetRDFI("031300012")             // Receivers bank transit routing number
	entryOne.DFIAccountNumber = "744-5678-99" // Receivers bank account number
	entryOne.Amount = 0                       // Amount of transaction with no decimal.  Zero Dollar Amount
	entryOne.SetOriginalTraceNumber("031300010000002")
	entryOne.SetReceivingCompany("Best. #1")
	entryOne.SetTraceNumber(bh.ODFIIdentification, 2)

	// build the batch
	batch := ach.NewBatchACK(bh)
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

	// write the file to std out. Anything io.Writer
	w := ach.NewWriter(os.Stdout)
	if err := w.Write(file); err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}
	w.Flush()
}
