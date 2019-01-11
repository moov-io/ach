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
	bh.ServiceClassCode = ach.CreditsOnly
	bh.CompanyName = "Name on Account" // The name of the company/person that has relationship with receiver
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.PPD
	bh.CompanyEntryDescription = "REG.SALARY" // will be on receiving accounts statement
	// Same Day ACH
	bh.EffectiveEntryDate = time.Now().Format("060102")
	// CompanyDescriptiveDate ODFIs at their discretion may require their Originators to further show intent for
	// same-day settlement using an optional, yet standardized, same-day indicator in the Company Descriptive Date
	// field. The Company Descriptive Date field (5 record, field 8) is an optional field with 6 positions available
	// (positions 64-69).
	bh.CompanyDescriptiveDate = "SD1300"
	bh.ODFIIdentification = "121042882" // Originating Routing Number

	// Identifies the receivers account information
	// can be multiple entry's per batch
	entry := ach.NewEntryDetail()
	// Identifies the entry as a debit and credit entry AND to what type of account (Savings, DDA, Loan, GL)
	entry.TransactionCode = ach.CheckingCredit
	entry.SetRDFI("231380104")          // Receivers bank transit routing number
	entry.DFIAccountNumber = "12345678" // Receivers bank account number
	entry.Amount = 100000000            // Amount of transaction with no decimal. One dollar and eleven cents = 111
	entry.SetTraceNumber(bh.ODFIIdentification, 1)
	entry.IndividualName = "Receiver Account Name" // Identifies the receiver of the transaction

	// build the batch
	batch := ach.NewBatchPPD(bh)
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
