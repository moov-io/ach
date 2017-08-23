package ach

import (
	"bytes"
	"testing"
)

func TestFileParam(t *testing.T) {
	f := NewFile(
		FileParam{
			ImmediateDestination:     "0210000890",
			ImmediateOrigin:          "123456789",
			ImmediateDestinationName: "Your Bank",
			ImmediateOriginName:      "Your Company",
			ReferenceCode:            "#00000A1"})
	if err := f.Header.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	if f.Header.ImmediateOriginName != "Your Company" {
		t.Errorf("FileParam value was not copied to file.Header")
	}
}

func TestBatchParam(t *testing.T) {
	companyName := "Your Company"
	batch := NewBatchPPD(BatchParam{
		ServiceClassCode:        "220",
		CompanyName:             companyName,
		StandardEntryClass:      "PPD",
		CompanyIdentification:   "123456789",
		CompanyEntryDescription: "Trans. Description",
		CompanyDescriptiveDate:  "Oct 23",
		ODFIIdentification:      "123456789"})

	if err := batch.header.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if batch.header.CompanyName != companyName {
		t.Errorf("BatchParam value was not copied to batch.header.CompanyName")
	}
}
func TestEntryParam(t *testing.T) {
	entry := NewEntryDetail(EntryParam{
		ReceivingDFI:      "102001017",
		RDFIAccount:       "5343121",
		Amount:            "17500",
		TransactionCode:   "22",
		IDNumber:          "ABC##jvkdjfuiwn",
		IndividualName:    "Bob Smith",
		DiscretionaryData: "B1"})

	if err := entry.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddendaParam(t *testing.T) {
	addenda := NewAddenda(AddendaParam{
		PaymentRelatedInfo: "Currently string needs ASC X12 Interchange Control Structures",
	})
	if err := addenda.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

func TestBuildFileParam(t *testing.T) {
	// To create a file
	file := NewFile(FileParam{
		ImmediateDestination:     "0210000890",
		ImmediateOrigin:          "123456789",
		ImmediateDestinationName: "Your Bank",
		ImmediateOriginName:      "Your Company",
		ReferenceCode:            "#00000A1"})

	// To create a batch. Errors only if payment type is not supported.
	batch, _ := NewBatch(BatchParam{
		ServiceClassCode:        "220",
		CompanyName:             "Your Company",
		StandardEntryClass:      "PPD",
		CompanyIdentification:   "123456789",
		CompanyEntryDescription: "Trans. Description",
		CompanyDescriptiveDate:  "Oct 23",
		ODFIIdentification:      "123456789"})

	// To create an entry
	entry := NewEntryDetail(EntryParam{
		ReceivingDFI:      "102001017",
		RDFIAccount:       "5343121",
		Amount:            "17500",
		TransactionCode:   "22",
		IDNumber:          "ABC##jvkdjfuiwn",
		IndividualName:    "Robert Smith",
		DiscretionaryData: "B1"})

	// To add one or more optional addenda records for an entry

	addenda := NewAddenda(AddendaParam{
		PaymentRelatedInfo: "bonus pay for amazing work on #OSS"})
	entry.AddAddenda(addenda)

	// Entries are added to batches like so:

	batch.AddEntry(entry)

	// When all of the Entries are added to the batch we must build it.

	if err := batch.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	// And batches are added to files much the same way:

	file.AddBatch(batch)

	// Once we added all our batches we must build the file

	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	// Finally we write the file to an io.Writer
	var b bytes.Buffer
	w := NewWriter(&b)
	if err := w.WriteAll([]*File{file}); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	w.Flush()
}
