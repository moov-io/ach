package ach

import (
	"bytes"
	"testing"
)

func TestFileRecord(t *testing.T) {
	f := NewFile()
	f.SetHeader(mockFileHeader())
	if err := f.Header.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	if f.Header.ImmediateOriginName != "My Bank Name" {
		t.Errorf("FileParam value was not copied to file.Header")
	}
}

func TestBatchRecord(t *testing.T) {
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

func TestEntryDetail(t *testing.T) {
	entry := mockEntryDetail()
	//override mockEntryDetail
	entry.TransactionCode = 27

	if err := entry.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

func TestEntryDetailPaymentType(t *testing.T) {

	entry := mockEntryDetail()
	//override mockEntryDetail
	entry.TransactionCode = 27
	entry.DiscretionaryData = "R"

	if err := entry.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

func TestEntryDetailReceivingCompany(t *testing.T) {
	entry := mockEntryDetail()
	//override mockEntryDetail
	entry.TransactionCode = 27
	entry.IdentificationNumber = "location #23"
	entry.IndividualName = "Best Co. #23"

	if err := entry.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddendaRecord(t *testing.T) {
	addenda05 := NewAddenda05()
	addenda05.PaymentRelatedInformation = "Currently string needs ASC X12 Interchange Control Structures"
	addenda05.SequenceNumber = 1
	addenda05.EntryDetailSequenceNumber = 1234567

	if err := addenda05.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

func TestBuildFile(t *testing.T) {
	// To create a file
	file := NewFile()
	file.SetHeader(mockFileHeader())

	// To create a batch. Errors only if payment type is not supported.
	batch, _ := NewBatch(mockBatchHeader())

	// To create an entry
	entry := mockPPDEntryDetail()

	// To add one or more optional addenda records for an entry
	addendaPPD := NewAddenda05()
	addendaPPD.PaymentRelatedInformation = "Currently string needs ASC X12 Interchange Control Structures"

	// Add the addenda record to the detail entry
	entry.AddAddenda(addendaPPD)

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

	// Add an entry and define if it is a single or recurring payment
	// The following is a reoccuring payment for $7.99

	entry = mockWEBEntryDetail()

	addendaWEB := NewAddenda05()
	addendaWEB.PaymentRelatedInformation = "Monthly Membership Subscription"

	// Add the addenda record to the detail entry
	entry.AddAddenda(addendaWEB)

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
