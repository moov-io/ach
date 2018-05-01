package main

import (
	"fmt"
	"os"

	"github.com/moov-io/ach"
	"time"
)

func main() {
	// To create a file
	fh := ach.NewFileHeader()
	fh.ImmediateDestination = "231380104"
	fh.ImmediateOrigin = "121042882"
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
	bh.CompanyIdentification = file.Header.ImmediateOrigin
	bh.StandardEntryClassCode = "PPD"
	bh.CompanyEntryDescription = "Trans. Description"
	bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1)
	bh.ODFIIdentification = "121042882"

	batch, _ := ach.NewBatch(bh)

	// To create an entry
	entry := ach.NewEntryDetail()
	entry.TransactionCode = 22
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "81967038518"
	entry.Amount = 1000000
	entry.IndividualName = "Wade Arnold"
	entry.SetTraceNumber(bh.ODFIIdentification, 1)
	entry.IdentificationNumber = "ABC##jvkdjfuiwn"
	entry.Category = ach.CategoryForward

	// To add one or more optional addenda records for an entry

	addenda := ach.NewAddenda05()
	addenda.PaymentRelatedInformation = "bonus pay for amazing work on #OSS"
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
	bh2.CompanyIdentification = file.Header.ImmediateOrigin
	bh2.StandardEntryClassCode = "WEB"
	bh2.CompanyEntryDescription = "Subscr"
	bh2.EffectiveEntryDate = time.Now().AddDate(0, 0, 1)
	bh2.ODFIIdentification = "121042882"

	batch2, _ := ach.NewBatch(bh2)

	// Add an entry and define if it is a single or reccuring payment
	// The following is a reccuring payment for $7.99

	entry2 := ach.NewEntryDetail()
	entry2.TransactionCode = 22
	entry2.SetRDFI("231380104")
	entry2.DFIAccountNumber = "81967038518"
	entry2.Amount = 799
	entry2.IndividualName = "Wade Arnold"
	entry2.SetTraceNumber(bh2.ODFIIdentification, 2)
	entry2.IdentificationNumber = "#123456"
	entry2.DiscretionaryData = "R"
	entry2.Category = ach.CategoryForward

	// To add one or more optional addenda records for an entry
	addenda2 := ach.NewAddenda05()
	addenda2.PaymentRelatedInformation = "Monthly Membership Subscription"
	entry2.AddAddenda(addenda2)

	// add the entry to the batch
	batch2.AddEntry(entry2)

	// Create and add the second batch
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
