package main

import (
	"github.com/moov-io/ach"
	"log"
	"os"
	"time"
)

func main() {
	// Example transfer to write an ACH RCK file to debit a external institutions account
	// Important: All financial institutions are different and will require registration and exact field values.

	fh := ach.NewFileHeader()
	fh.ImmediateDestination = "231380104" // Routing Number of the ACH Operator or receiving point to which the file is being sent
	fh.ImmediateOrigin = "121042882"      // Routing Number of the ACH Operator or sending point that is sending the file
	fh.FileCreationDate = time.Now()      // Today's Date
	fh.ImmediateDestinationName = "Federal Reserve Bank"
	fh.ImmediateOriginName = "My Bank Name"

	// BatchHeader identifies the originating entity and the type of transactions contained in the batch
	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = 225          // ACH credit pushes money out, 225 debits/pulls money in.
	bh.CompanyName = "Name on Account" // The name of the company/person that has relationship with receiver
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = "RCK"         // Consumer destination vs Company CCD
	bh.CompanyEntryDescription = "REDEPCHECK" // will be on receiving accounts statement
	bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1)
	bh.ODFIIdentification = "121042882" // Originating Routing Number

	// Identifies the receivers account information
	// can be multiple entry's per batch
	entry := ach.NewEntryDetail()
	// Identifies the entry as a debit and credit entry AND to what type of account (Savings, DDA, Loan, GL)
	entry.TransactionCode = ach.CheckingDebit // Code 27: Debit (withdrawal) from checking account
	entry.SetRDFI("231380104")                // Receivers bank transit routing number
	entry.DFIAccountNumber = "12345678"       // Receivers bank account number
	entry.Amount = 2400                       // Amount of transaction with no decimal. One dollar and eleven cents = 111
	entry.SetTraceNumber(bh.ODFIIdentification, 1)
	entry.IndividualName = "Wade Arnold" // Identifies the receiver of the transaction
	entry.SetCheckSerialNumber("123123123")

	// build the batch
	batch := ach.NewBatchRCK(bh)
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

	// write the file to std out. Anything io.Writer
	w := ach.NewWriter(os.Stdout)
	if err := w.Write(file); err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}
	w.Flush()
}
