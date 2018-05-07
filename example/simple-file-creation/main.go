package main

import (
	"fmt"
	"os"

	"strconv"
	"time"

	"github.com/moov-io/ach"
)

func main() {
	// To create a file
	fh := ach.NewFileHeader()
	fh.ImmediateDestination = 9876543210
	fh.ImmediateOrigin = 1234567890
	fh.FileCreationDate = time.Now()
	fh.ImmediateDestinationName = "Federal Reserve Bank"
	fh.ImmediateOriginName = "My Bank Name"
	file := ach.NewFile()
	file.SetHeader(fh)

	// To create a batch.
	// Errors only if payment type is not supported.
	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = 200
	bh.CompanyName = "Your Company"
	bh.CompanyIdentification = strconv.Itoa(file.Header.ImmediateOrigin)
	bh.StandardEntryClassCode = "PPD"
	bh.CompanyEntryDescription = "Trans. Description"
	bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1)
	bh.ODFIIdentification = 123456789

	batch, _ := ach.NewBatch(bh)

	// To create an entry
	entry := ach.NewEntryDetail()
	entry.TransactionCode = 22
	entry.SetRDFI(9101298)
	entry.DFIAccountNumber = "123456789"
	entry.Amount = 100000000
	entry.IndividualName = "Wade Arnold"
	entry.SetTraceNumber(bh.ODFIIdentification, 1)
	entry.IdentificationNumber = "ABC##jvkdjfuiwn"
	entry.Category = ach.CategoryForward

	// To add one or more optional addenda records for an entry
	addenda, _ := ach.NewAddenda(ach.AddendaParam{
		PaymentRelatedInfo: "bonus pay for amazing work on #OSS"})
	entry.AddAddenda(addenda)

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
	bh2.ServiceClassCode = 220
	bh2.CompanyName = "Your Company"
	bh2.CompanyIdentification = strconv.Itoa(file.Header.ImmediateOrigin)
	bh2.StandardEntryClassCode = "WEB"
	bh2.CompanyEntryDescription = "Subscr"
	bh2.EffectiveEntryDate = time.Now().AddDate(0, 0, 1)
	bh2.ODFIIdentification = 123456789

	batch2, _ := ach.NewBatch(bh2)

	// Add an entry and define if it is a single or recurring payment
	// The following is a recurring payment for $7.99

	entry2 := ach.NewEntryDetail()
	entry2.TransactionCode = 22
	entry2.SetRDFI(102001017)
	entry2.DFIAccountNumber = "5343121"
	entry2.Amount = 799
	entry2.IndividualName = "Wade Arnold"
	entry2.SetTraceNumber(bh2.ODFIIdentification, 2)
	entry2.IdentificationNumber = "#123456"
	entry2.DiscretionaryData = "R"
	entry2.Category = ach.CategoryForward

	addenda2, _ := ach.NewAddenda(ach.AddendaParam{
		PaymentRelatedInfo: "Monthly Membership Subscription"})

	// add the entry to the batch
	entry2.AddAddenda(addenda2)

	// Create and add the second batch

	batch2.AddEntry(entry2)
	if err := batch2.Create(); err != nil {
		fmt.Printf("%T: %s", err, err)
	}
	file.AddBatch(batch2)

	// Once we added all our batches we must create the file

	if err := file.Create(); err != nil {
		fmt.Printf("%T: %s", err, err)
	}

	// Finally we wnt to write the file to an io.Writer
	w := ach.NewWriter(os.Stdout)
	if err := w.Write(file); err != nil {
		fmt.Printf("%T: %s", err, err)
	}
	w.Flush()
}
