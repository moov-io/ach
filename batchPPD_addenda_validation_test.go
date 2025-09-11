package ach

import (
	"strings"
	"testing"
)

func TestBatchPPD_Addenda05_InvalidFormat_IsCaughtByBatchValidate(t *testing.T) {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 200
	bh.StandardEntryClassCode = "PPD"
	bh.CompanyName = "COMPANY"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "PAYROLL"
	bh.ODFIIdentification = "12104288"

	batch := NewBatchPPD(bh)

	ed := NewEntryDetail()
	ed.TransactionCode = 22
	ed.RDFIIdentification = "23138010"
	ed.CheckDigit = "4"
	ed.DFIAccountNumber = "123456789"
	ed.Amount = 100
	ed.IdentificationNumber = "ID0001"
	ed.IndividualName = "Jane Doe"
	ed.TraceNumber = "121042880000001"

	add := NewAddenda05()
	// Intentionally invalid: >80 chars
	add.PaymentRelatedInformation = strings.Repeat("X", 81)

	ed.AddAddenda05(add)
	ed.AddendaRecordIndicator = 1
	batch.AddEntry(ed)

	// Create may validate internally; accept error here OR on Validate() below.
	if err := batch.Create(); err != nil {
		// Ensure it's the expected validation error (and not something unrelated)
		if !strings.Contains(err.Error(), "PaymentRelatedInformation exceeds 80 characters") {
			t.Fatalf("unexpected error from batch.Create(): %v", err)
		}
		return // test passes: invalid addenda was caught
	}

	// If Create() succeeds, Validate() must catch it.
	if err := batch.Validate(); err == nil {
		t.Fatalf("expected validation error due to invalid Addenda05, got nil")
	}
}
