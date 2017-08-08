package main

import (
	"fmt"
	"os"

	"github.com/moov-io/ach"
)

func main() {
	// To create a file
	file := ach.NewFile(ach.FileParam{
		ImmediateDestination:     "0210000890",
		ImmediateOrigin:          "123456789",
		ImmediateDestinationName: "Your Bank",
		ImmediateOriginName:      "Your Company",
		ReferenceCode:            "#00000A1"})

	// To create a batch
	batch := ach.NewBatchPPD(ach.BatchParam{
		ServiceClassCode:        "220",
		CompanyName:             "Your Company",
		StandardEntryClass:      "PPD",
		CompanyIdentification:   "123456789",
		CompanyEntryDescription: "Trans. Description",
		CompanyDescriptiveDate:  "Oct 23",
		ODFIIdentification:      "123456789"})

	// To create an entry
	entry := ach.NewEntryDetail(ach.EntryParam{
		ReceivingDFI:      "102001017",
		RDFIAccount:       "5343121",
		Amount:            "17500",
		TransactionCode:   "22",
		IDNumber:          "ABC##jvkdjfuiwn",
		IndividualName:    "Bob Smith",
		DiscretionaryData: "B1"})

	// To add one or more optional addenda records for an entry
	addenda := ach.NewAddenda(ach.AddendaParam{
		PaymentRelatedInfo: "bonus pay for amazing work on #OSS"})
	entry.AddAddenda(addenda)

	// Entries are added to batches like so:

	batch.AddEntry(entry)

	// When all of the Entries are added to the batch we must build it.

	if err := batch.Build(); err != nil {
		fmt.Printf("%T: %s", err, err)
	}

	// And batches are added to files much the same way:

	file.AddBatch(batch)

	// Once we added all our batches we must build the file

	if err := file.Build(); err != nil {
		fmt.Printf("%T: %s", err, err)
	}

	// Finally we wnt to write the file to an io.Writer
	w := ach.NewWriter(os.Stdout)
	if err := w.Write(file); err != nil {
		fmt.Printf("%T: %s", err, err)
	}
	w.Flush()
}
