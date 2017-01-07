// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
)

// TestParseEntryDetail Header parses a known Entry Detail Record string.
func TestParseEntryDetail(t *testing.T) {
	var line = "62705320001912345            0000010500c-1            Arnold Wade           DD0076401255655291"
	r := NewReader(strings.NewReader(line))
	r.currentBatch.Header.StandardEntryClassCode = "PPD"
	r.line = line
	err := r.parseEntryDetail()
	if err != nil {
		t.Errorf("unknown error: %v", err)
	}
	record := r.currentBatch.Entries[0]

	if record.recordType != "6" {
		t.Errorf("RecordType Expected '6' got: %v", record.recordType)
	}
	if record.TransactionCode != 27 {
		t.Errorf("TransactionCode Expected '27' got: %v", record.TransactionCode)
	}
	if record.RDFIIdentification() != "05320001" {
		t.Errorf("RDFIIdentification Expected '05320001' got: '%v'", record.RDFIIdentification())
	}
	if record.CheckDigit != 9 {
		t.Errorf("CheckDigit Expected '9' got: %v", record.CheckDigit)
	}
	if record.DFIAccountNumber() != "12345            " {
		t.Errorf("DfiAccountNumber Expected '12345            ' got: %v", record.DFIAccountNumber())
	}
	if record.Amount() != "0000010500" {
		t.Errorf("Amount Expected '0000010500' got: %v", record.Amount())
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
	if record.TraceNumber() != "076401255655291" {
		t.Errorf("TraceNumber Expected '076401255655291' got: %v", record.TraceNumber())
	}
}

// TestEDString validats that a known parsed file can be return to a string of the same value
func TestEDString(t *testing.T) {
	var line = "62705320001912345            0000010500c-1            Arnold Wade           DD0076401255655291"
	r := NewReader(strings.NewReader(line))
	r.currentBatch.Header.StandardEntryClassCode = "PPD"
	r.line = line
	err := r.parseEntryDetail()
	if err != nil {
		t.Errorf("unknown error: %v", err)
	}
	record := r.currentBatch.Entries[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestValidateEDRecordType ensure error if recordType is not 6
func TestValidateEDRecordType(t *testing.T) {
	ed := NewEntryDetail()
	ed.recordType = "2"
	_, err := ed.Validate()
	if !strings.Contains(err.Error(), ErrRecordType.Error()) {
		t.Errorf("Expected RecordType Error got: %v", err)
	}
}

// TestValidateEDTransactionCode ensure error if TransactionCode is not valid
func TestValidateEDTransactionCode(t *testing.T) {
	ed := NewEntryDetail()
	ed.TransactionCode = 63
	_, err := ed.Validate()
	if !strings.Contains(err.Error(), ErrTransactionCode.Error()) {
		t.Errorf("Expected Transaction Code Error got: %v", err)
	}
}
