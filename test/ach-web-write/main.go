package main

import (
	"log"
	"os"
	"time"

	"github.com/moov-io/ach"
)

func main() {
	// Example transfer to write an ACH WEB file to send/credit a external institutions account
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
	bh.ServiceClassCode = 220          //
	bh.CompanyName = "Name on Account" // The name of the company/person that has relationship with receiver
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = "WEB"        // Consumer destination vs Company CCD
	bh.CompanyEntryDescription = "Subscribe" // will be on receiving accounts statement
	bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1)
	bh.ODFIIdentification = "121042882" // Originating Routing Number

	entry := ach.NewEntryDetail()
	entry.TransactionCode = 22
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "12345678"
	entry.Amount = 10000 //  100.00
	entry.IndividualName = "Wade Arnold"
	entry.SetTraceNumber(bh.ODFIIdentification, 1)
	entry.IdentificationNumber = "#789654"
	entry.DiscretionaryData = "S"
	entry.Category = ach.CategoryForward

	// To add one or more optional addenda records for an entry
	addenda1 := ach.NewAddenda05()
	addenda1.PaymentRelatedInformation = "PAY-GATE payment"
	entry.AddAddenda(addenda1)

	entry2 := ach.NewEntryDetail()
	entry2.TransactionCode = 22
	entry2.SetRDFI("231380104")
	entry2.DFIAccountNumber = "81967038518"
	entry2.Amount = 799 // 7.99
	entry2.IndividualName = "Wade Arnold"
	entry2.SetTraceNumber(bh.ODFIIdentification, 2)
	entry2.IdentificationNumber = "#123456"
	entry2.DiscretionaryData = "R"
	entry2.Category = ach.CategoryForward

	// To add one or more optional addenda records for an entry
	addenda2 := ach.NewAddenda05()
	addenda2.PaymentRelatedInformation = "Monthly Membership Subscription"
	entry2.AddAddenda(addenda2)

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

	// write the file to std out. Anything io.Writer
	w := ach.NewWriter(os.Stdout)
	if err := w.Write(file); err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}
	w.Flush()
}
