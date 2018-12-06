package main

import (
	"github.com/moov-io/ach"
	"log"
	"os"
	"time"
)

func main() {
	// Example transfer to write an ACH ARC file to debit an external institutions account
	// Important: All financial institutions are different and will require registration and exact field values.

	fh := ach.NewFileHeader()
	fh.ImmediateDestination = "231380104" // Routing Number of the ACH Operator or receiving point to which the file is being sent
	fh.ImmediateOrigin = "121042882"      // Routing Number of the ACH Operator or sending point that is sending the file
	fh.FileCreationDate = time.Now()      // Today's Date
	fh.ImmediateDestinationName = "Federal Reserve Bank"
	fh.ImmediateOriginName = "My Bank Name"

	// BatchHeader identifies the originating entity and the type of transactions contained in the batch
	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.AutomatedAccountingAdvices
	bh.CompanyName = "Company Name, Inc" // The name of the company/person that has relationship with receiver
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = "ADV"         // Consumer destination vs Company CCD
	bh.CompanyEntryDescription = "Accounting" // will be on receiving accounts statement
	bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1)
	bh.ODFIIdentification = "121042882" // Originating Routing Number

	// Identifies the receivers account information
	// can be multiple entry's per batch
	entry := ach.NewADVEntryDetail()
	// Credit for ACH debits originated
	entry.TransactionCode = 81 //
	entry.SetRDFI("231380104") // Receivers bank transit routing number
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 50000
	entry.AdviceRoutingNumber = "121042882"
	entry.FileIdentification = "11131"
	entry.ACHOperatorData = ""
	entry.IndividualName = "Name"
	entry.DiscretionaryData = ""
	entry.AddendaRecordIndicator = 0
	entry.ACHOperatorRoutingNumber = "01100001"
	entry.JulianDay = 50
	entry.SequenceNumber = 1

	entryOne := ach.NewADVEntryDetail()
	// Debit for ACH credits originated
	entryOne.TransactionCode = 82 //
	entryOne.SetRDFI("231380104") // Receivers bank transit routing number
	entryOne.DFIAccountNumber = "744-5678-99"
	entryOne.Amount = 250000
	entryOne.AdviceRoutingNumber = "121042882"
	entryOne.FileIdentification = "11139"
	entryOne.ACHOperatorData = ""
	entryOne.IndividualName = "Name"
	entryOne.DiscretionaryData = ""
	entryOne.AddendaRecordIndicator = 0
	entryOne.ACHOperatorRoutingNumber = "01100001"
	entryOne.JulianDay = 50
	entryOne.SequenceNumber = 2

	// build the batch
	batch := ach.NewBatchADV(bh)
	batch.AddADVEntry(entry)
	batch.AddADVEntry(entryOne)
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
