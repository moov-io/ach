package main

import (
	"log"
	"os"
	"time"

	"github.com/moov-io/ach"
)

func main() {
	// Example transfer to write an ACH ENR file acknowledging a credit
	// Important: All financial institutions are different and will require registration and exact field values.

	// Set originator bank ODFI and destination Operator for the financial institution
	// this is the funding/receiving source of the transfer
	fh := ach.NewFileHeader()
	fh.ImmediateDestination = "031300012" // Routing Number of the ACH Operator or receiving point to which the file is being sent
	fh.ImmediateOrigin = "231380104"      // Routing Number of the ACH Operator or sending point that is sending the file
	fh.ImmediateDestinationName = "Federal Reserve Bank"
	fh.ImmediateOriginName = "My Bank Name"
	fh.FileCreationDate = time.Now().Format("060102")

	// BatchHeader identifies the originating entity and the type of transactions contained in the batch
	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.DebitsOnly
	bh.CompanyName = "Name on Account" // The name of the company/person that has relationship with receiver
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.ENR
	bh.CompanyEntryDescription = "AUTOENROLL"
	bh.ODFIIdentification = "23138010" // Originating Routing Number

	// Identifies the receivers account information
	// can be multiple entry's per batch
	entry := ach.NewEntryDetail()
	// Identifies the entry as a debit and credit entry AND to what type of account (Savings, DDA, Loan, GL)
	entry.TransactionCode = ach.CheckingDebit
	entry.SetRDFI("031300012")             // Receivers bank transit routing number
	entry.DFIAccountNumber = "744-5678-99" // Receivers bank account number
	entry.Amount = 0                       // Amount of transaction with no decimal. One dollar and eleven cents = 111
	entry.SetOriginalTraceNumber("031300010000001")
	entry.SetReceivingCompany("Best. #1")
	entry.SetTraceNumber(bh.ODFIIdentification, 1)

	addenda05 := ach.NewAddenda05()
	addenda05.PaymentRelatedInformation = `22*12200004*3*123987654321*777777777*DOE*JOHN*1\` // From NACHA 2013 Official Rules
	entry.AddAddenda05(addenda05)
	entry.AddendaRecordIndicator = 1

	// build the batch
	batch := ach.NewBatchENR(bh)
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
