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
	entry.TraceNumber = 123456789
	return entry
}

// TestParseEntryDetail Header parses a known Entry Detail Record string.
func TestParseEntryDetail(t *testing.T) {
	var line = "62705320001912345            0000010500c-1            Arnold Wade           DD0076401255655291"
	r := NewReader(strings.NewReader(line))
	r.addCurrentBatch(NewBatchPPD())
	r.currentBatch.SetHeader(mockBatchHeader())
	r.line = line
	err := r.parseEntryDetail()
	if err != nil {
		t.Errorf("unknown error: %v", err)
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

	if record.IndividualIdentificationNumber != "c-1            " {
		t.Errorf("IndividualIdentificationNumber Expected 'c-1            ' got: %v", record.IndividualIdentificationNumber)
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
	r.addCurrentBatch(NewBatchPPD())
	r.currentBatch.SetHeader(mockBatchHeader())
	r.line = line
	err := r.parseEntryDetail()
	if err != nil {
		t.Errorf("unknown error: %v", err)
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
		if !strings.Contains(err.Error(), ErrRecordType.Error()) {
			t.Errorf("Expected RecordType Error got: %v", err)
		}
	}
}

// TestValidateEDTransactionCode ensure error if TransactionCode is not valid
func TestValidateEDTransactionCode(t *testing.T) {
	ed := mockEntryDetail()
	ed.TransactionCode = 63
	if err := ed.Validate(); err != nil {
		if !strings.Contains(err.Error(), ErrTransactionCode.Error()) {
			t.Errorf("Expected Transaction Code Error got: %v", err)
		}
	}
}

func TestEDFieldInclusion(t *testing.T) {
	ed := mockEntryDetail()
	// works properly
	if err := ed.Validate(); err != nil {
		t.Errorf("Unexpected Entry Detail error: %v", err.Error())
	}
	// create error is mismatch
	ed.Amount = 0
	if err := ed.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Entry Detail error: %v", err.Error())
		}
	}
}

func TestEDdfiAccountNumberAlphaNumeric(t *testing.T) {
	ed := mockEntryDetail()
	// works properly
	if err := ed.Validate(); err != nil {
		t.Errorf("Unexpected Entry Detail error: %v", err.Error())
	}
	// create error is mismatch
	ed.DFIAccountNumber = "74647#999!"
	if err := ed.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Entry Detail error: %v", err.Error())
		}
	}
}

func TestEDIndividualIdentificationNumberAlphaNumeric(t *testing.T) {
	ed := mockEntryDetail()
	// works properly
	if err := ed.Validate(); err != nil {
		t.Errorf("Unexpected Entry Detail error: %v", err.Error())
	}
	// create error is mismatch
	ed.IndividualIdentificationNumber = "#12345!"
	if err := ed.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Entry Detail error: %v", err.Error())
		}
	}
}

func TestEDIndividualNameAlphaNumeric(t *testing.T) {
	ed := mockEntryDetail()
	// works properly
	if err := ed.Validate(); err != nil {
		t.Errorf("Unexpected Entry Detail error: %v", err.Error())
	}
	// create error is mismatch
	ed.IndividualName = "W@DE"
	if err := ed.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Entry Detail error: %v", err.Error())
		}
	}
}

func TestEDDiscretionaryDataAlphaNumeric(t *testing.T) {
	ed := mockEntryDetail()
	// works properly
	if err := ed.Validate(); err != nil {
		t.Errorf("Unexpected Entry Detail error: %v", err.Error())
	}
	// create error is mismatch
	ed.DiscretionaryData = "@!"
	if err := ed.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Entry Detail error: %v", err.Error())
		}
	}
}

func TestEDisCheckDigit(t *testing.T) {
	ed := mockEntryDetail()
	// works properly
	if err := ed.Validate(); err != nil {
		t.Errorf("Unexpected Entry Detail error: %v", err.Error())
	}
	// create error is mismatch
	ed.CheckDigit = 1
	if err := ed.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Entry Detail error: %v", err.Error())
		}
	}
}

func TestEDSetRDFI(t *testing.T) {
	ed := NewEntryDetail()
	ed.SetRDFI(81086674)
	if ed.RDFIIdentification != 8108667 {
		t.Errorf("RDFI identification")
	}
	if ed.CheckDigit != 4 {
		t.Errorf("Unexpected check digit")
	}
}

func TestEDFieldInclusionRecordType(t *testing.T) {
	entry := mockEntryDetail()
	entry.recordType = ""
	if err := entry.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestEDFieldInclusionTransactionCode(t *testing.T) {
	entry := mockEntryDetail()
	entry.TransactionCode = 0
	if err := entry.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestEDFieldInclusionRDFIIdentification(t *testing.T) {
	entry := mockEntryDetail()
	entry.RDFIIdentification = 0
	if err := entry.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestEDFieldInclusionCheckDigit(t *testing.T) {
	entry := mockEntryDetail()
	entry.CheckDigit = 0
	if err := entry.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestEDFieldInclusionDFIAccountNumber(t *testing.T) {
	entry := mockEntryDetail()
	entry.DFIAccountNumber = ""
	if err := entry.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestEDFieldInclusionIndividualName(t *testing.T) {
	entry := mockEntryDetail()
	entry.IndividualName = ""
	if err := entry.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestEDFieldInclusionTraceNumber(t *testing.T) {
	entry := mockEntryDetail()
	entry.TraceNumber = 0
	if err := entry.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}
