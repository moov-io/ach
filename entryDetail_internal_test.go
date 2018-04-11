// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
)

func mockEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 22
	entry.SetRDFI(9101298)
	entry.DFIAccountNumber = "123456789"
	entry.Amount = 100000000
	entry.IndividualName = "Wade Arnold"
	entry.SetTraceNumber(mockBatchHeader().ODFIIdentification, 1)
	entry.IdentificationNumber = "ABC##jvkdjfuiwn"
	entry.Category = CategoryForward
	return entry
}

func mockEntryDemandDebit() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 27
	entry.SetRDFI(102001017)
	entry.DFIAccountNumber = "5343121"
	entry.Amount = 17500
	entry.IndividualName = "Robert Smith"
	entry.SetTraceNumber(mockBatchHeader().ODFIIdentification, 1)
	entry.IdentificationNumber = "ABC##jvkdjfuiwn"
	entry.Category = CategoryForward
	return entry
}

func TestMockEntryDetail(t *testing.T) {
	entry := mockEntryDetail()
	if err := entry.Validate(); err != nil {
		t.Error("mockEntryDetail does not validate and will break other tests")
	}
	if entry.TransactionCode != 22 {
		t.Error("TransactionCode dependent default value has changed")
	}
	if entry.DFIAccountNumber != "123456789" {
		t.Error("DFIAccountNumber dependent default value has changed")
	}
	if entry.Amount != 100000000 {
		t.Error("Amount dependent default value has changed")
	}
	if entry.IndividualName != "Wade Arnold" {
		t.Error("IndividualName dependent default value has changed")
	}
	if entry.TraceNumber != 62000010000001 {
		t.Errorf("TraceNumber dependent default value has changed %v", entry.TraceNumber)
	}
}

// TestParseEntryDetail Header parses a known Entry Detail Record string.
func TestParseEntryDetail(t *testing.T) {
	var line = "62705320001912345            0000010500c-1            Arnold Wade           DD0076401255655291"
	r := NewReader(strings.NewReader(line))
	r.addCurrentBatch(NewBatchPPD(mockBatchPPDHeader()))
	r.currentBatch.SetHeader(mockBatchHeader())
	r.line = line
	if err := r.parseEntryDetail(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentBatch.GetEntries()[0]

	if record.recordType != "6" {
		t.Errorf("RecordType Expected '6' got: %v", record.recordType)
	}
	if record.TransactionCode != 27 {
		t.Errorf("TransactionCode Expected '27' got: %v", record.TransactionCode)
	}
	if record.RDFIIdentificationField() != "05320001" {
		t.Errorf("RDFIIdentification Expected '05320001' got: '%v'", record.RDFIIdentificationField())
	}
	if record.CheckDigit != 9 {
		t.Errorf("CheckDigit Expected '9' got: %v", record.CheckDigit)
	}
	if record.DFIAccountNumberField() != "12345            " {
		t.Errorf("DfiAccountNumber Expected '12345            ' got: %v", record.DFIAccountNumberField())
	}
	if record.AmountField() != "0000010500" {
		t.Errorf("Amount Expected '0000010500' got: %v", record.AmountField())
	}

	if record.IdentificationNumber != "c-1            " {
		t.Errorf("IdentificationNumber Expected 'c-1            ' got: %v", record.IdentificationNumber)
	}
	if record.IndividualName != "Arnold Wade           " {
		t.Errorf("IndividualName Expected 'Arnold Wade           ' got: %v", record.IndividualName)
	}
	if record.DiscretionaryData != "DD" {
		t.Errorf("DiscretionaryData Expected 'DD' got: %v", record.DiscretionaryData)
	}
	if record.AddendaRecordIndicator != 0 {
		t.Errorf("AddendaRecordIndicator Expected '0' got: %v", record.AddendaRecordIndicator)
	}
	if record.TraceNumberField() != "076401255655291" {
		t.Errorf("TraceNumber Expected '076401255655291' got: %v", record.TraceNumberField())
	}
}

// TestEDString validats that a known parsed file can be return to a string of the same value
func TestEDString(t *testing.T) {
	var line = "62705320001912345            0000010500c-1            Arnold Wade           DD0076401255655291"
	r := NewReader(strings.NewReader(line))
	r.addCurrentBatch(NewBatchPPD(mockBatchPPDHeader()))
	r.currentBatch.SetHeader(mockBatchHeader())
	r.line = line
	if err := r.parseEntryDetail(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentBatch.GetEntries()[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestValidateEDRecordType ensure error if recordType is not 6
func TestValidateEDRecordType(t *testing.T) {
	ed := mockEntryDetail()
	ed.recordType = "2"
	if err := ed.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestValidateEDTransactionCode ensure error if TransactionCode is not valid
func TestValidateEDTransactionCode(t *testing.T) {
	ed := mockEntryDetail()
	ed.TransactionCode = 63
	if err := ed.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TransactionCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestEDFieldInclusion(t *testing.T) {
	ed := mockEntryDetail()
	ed.Amount = 0
	if err := ed.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestEDdfiAccountNumberAlphaNumeric(t *testing.T) {
	ed := mockEntryDetail()
	ed.DFIAccountNumber = "®"
	if err := ed.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "DFIAccountNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestEDIdentificationNumberAlphaNumeric(t *testing.T) {
	ed := mockEntryDetail()
	ed.IdentificationNumber = "®"
	if err := ed.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "IdentificationNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestEDIndividualNameAlphaNumeric(t *testing.T) {
	ed := mockEntryDetail()
	ed.IndividualName = "W®DE"
	if err := ed.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "IndividualName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestEDDiscretionaryDataAlphaNumeric(t *testing.T) {
	ed := mockEntryDetail()
	ed.DiscretionaryData = "®!"
	if err := ed.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "DiscretionaryData" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestEDisCheckDigit(t *testing.T) {
	ed := mockEntryDetail()
	ed.CheckDigit = 1
	if err := ed.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "RDFIIdentification" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestEDSetRDFI(t *testing.T) {
	ed := NewEntryDetail()
	ed.SetRDFI(81086674)
	if ed.RDFIIdentification != 8108667 {
		t.Error("RDFI identification")
	}
	if ed.CheckDigit != 4 {
		t.Error("Unexpected check digit")
	}
}

func TestEDFieldInclusionRecordType(t *testing.T) {
	entry := mockEntryDetail()
	entry.recordType = ""
	if err := entry.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestEDFieldInclusionTransactionCode(t *testing.T) {
	entry := mockEntryDetail()
	entry.TransactionCode = 0
	if err := entry.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestEDFieldInclusionRDFIIdentification(t *testing.T) {
	entry := mockEntryDetail()
	entry.RDFIIdentification = 0
	if err := entry.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestEDFieldInclusionDFIAccountNumber(t *testing.T) {
	entry := mockEntryDetail()
	entry.DFIAccountNumber = ""
	if err := entry.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestEDFieldInclusionIndividualName(t *testing.T) {
	entry := mockEntryDetail()
	entry.IndividualName = ""
	if err := entry.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestEDFieldInclusionTraceNumber(t *testing.T) {
	entry := mockEntryDetail()
	entry.TraceNumber = 0
	if err := entry.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestEDAddAddendaAddendaReturn(t *testing.T) {
	entry := mockEntryDetail()
	entry.AddAddenda(mockAddendaReturn())
	if entry.Category != CategoryReturn {
		t.Error("AddendaReturn added and isReturn is false")
	}
	if entry.AddendaRecordIndicator != 1 {
		t.Error("AddendaReturn added and record indicator is not 1")
	}

}

func TestEDAddAddendaAddendaReturnTwice(t *testing.T) {
	entry := mockEntryDetail()
	entry.AddAddenda(mockAddendaReturn())
	entry.AddAddenda(mockAddendaReturn())
	if entry.Category != CategoryReturn {
		t.Error("AddendaReturn added and Category is not CategoryReturn")
	}

	if len(entry.Addendum) != 1 {
		t.Error("AddendaReturn added and isReturn is false")
	}
}

func TestEDCreditOrDebit(t *testing.T) {
	// TODO add more credit and debit transaction code's to this test
	entry := mockEntryDetail()
	if entry.CreditOrDebit() != "C" {
		t.Errorf("TransactionCode %v expected a Credit(C) got %v", entry.TransactionCode, entry.CreditOrDebit())
	}
	entry.TransactionCode = 27
	if entry.CreditOrDebit() != "D" {
		t.Errorf("TransactionCode %v expected a Debit(D) got %v", entry.TransactionCode, entry.CreditOrDebit())
	}
}
