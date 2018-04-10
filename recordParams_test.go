package ach

import (
	"testing"
	"bytes"
)

func TestFileParam(t *testing.T) {
	f := NewFile()
	f.SetHeader(mockFileHeader())
	if err := f.Header.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	if f.Header.ImmediateOriginName != "My Bank Name" {
		t.Errorf("FileParam value was not copied to file.Header")
	}
}


func TestBatchParam(t *testing.T) {
	companyName := "ACME Corporation"
	batch, _ := NewBatch(mockBatchPPDHeader())

	bh := batch.GetHeader()
	if err := bh.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if bh.CompanyName != companyName {
		t.Errorf("BatchParam value was not copied to batch.header.CompanyName")
	}
}

func TestEntryParam(t *testing.T) {
	entry := NewEntryDetail(EntryParam{
		ReceivingDFI:      "102001017",
		RDFIAccount:       "5343121",
		Amount:            "17500",
		TransactionCode:   "27",
		IDNumber:          "ABC##jvkdjfuiwn",
		IndividualName:    "Bob Smith",
		DiscretionaryData: "B1"})

	if err := entry.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

func TestEntryParamPaymentType(t *testing.T) {
	entry := NewEntryDetail(EntryParam{
		ReceivingDFI:    "102001017",
		RDFIAccount:     "5343121",
		Amount:          "17500",
		TransactionCode: "27",
		IDNumber:        "ABC##jvkdjfuiwn",
		IndividualName:  "Bob Smith",
		PaymentType:     "R"})

	if err := entry.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

func TestEntryParamReceivingCompany(t *testing.T) {
	entry := NewEntryDetail(EntryParam{
		ReceivingDFI:     "102001017",
		RDFIAccount:      "5343121",
		Amount:           "17500",
		TransactionCode:  "27",
		IDNumber:         "location #23",
		ReceivingCompany: "Best Co. #23"})

	if err := entry.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddendaParam(t *testing.T) {
	addenda, _ := NewAddenda(AddendaParam{
		PaymentRelatedInfo: "Currently string needs ASC X12 Interchange Control Structures",
	})
	if err := addenda.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

func TestBuildFileParam(t *testing.T) {
	// To create a file
	file := NewFile()
	file.SetHeader(mockFileHeader())

	// To create a batch. Errors only if payment type is not supported.
	batch, _ := NewBatch(mockBatchHeader())

	// To create an entry
	entry := NewEntryDetail(EntryParam{
		ReceivingDFI:      "102001017",
		RDFIAccount:       "5343121",
		Amount:            "17500",
		TransactionCode:   "27",
		IDNumber:          "ABC##jvkdjfuiwn",
		IndividualName:    "Robert Smith",
		DiscretionaryData: "B1"})

	// To add one or more optional addenda records for an entry

	addenda, _ := NewAddenda(AddendaParam{
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

	// Now add a new batch for accepting payments on the web

	batch, _ = NewBatch(mockBatchWEBHeader())

	// Add an entry and define if it is a single or reoccuring payment
	// The following is a reoccuring payment for $7.99

	entry = NewEntryDetail(EntryParam{
		ReceivingDFI:    "102001017",
		RDFIAccount:     "5343121",
		Amount:          "799",
		TransactionCode: "22",
		IDNumber:        "#123456",
		IndividualName:  "Wade Arnold",
		PaymentType:     "R"})

	addenda, _ = NewAddenda(AddendaParam{
		PaymentRelatedInfo: "Monthly Membership Subscription"})

	// add the entry to the batch
	entry.AddAddenda(addenda)

	// add the second batch to the file

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
