package main

import (
	"github.com/moov-io/ach"
	"log"
	"os"
	"time"
)

func main() {
	// Example transfer to write an ACH SHR file to send/credit a external institutions account
	// Important: All financial institutions are different and will require registration and exact field values.

	// Set originator bank ODFI and destination Operator for the financial institution
	// this is the funding/receiving source of the transfer
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
	bh.StandardEntryClassCode = "SHR"      // Consumer destination vs Company CCD
	bh.CompanyEntryDescription = "Payment" // will be on receiving accounts statement
	bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1)
	bh.ODFIIdentification = "121042882" // Originating Routing Number

	// Identifies the receivers account information
	// can be multiple entry's per batch
	entry := ach.NewEntryDetail()
	// Identifies the entry as a debit and credit entry AND to what type of account (Savings, DDA, Loan, GL)
	entry.TransactionCode = 27          // Code 22: Debit (withdrawal) from checking account
	entry.SetRDFI("231380104")          // Receivers bank transit routing number
	entry.DFIAccountNumber = "12345678" // Receivers bank account number
	entry.Amount = 100000000            // Amount of transaction with no decimal. One dollar and eleven cents = 111
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

	// write the file to std out. Anything io.Writer
	w := ach.NewWriter(os.Stdout)
	if err := w.Write(file); err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}
	w.Flush()
}
