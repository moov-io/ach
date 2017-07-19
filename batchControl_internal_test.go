// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
)

func mockBatchControl() *BatchControl {
	bc := NewBatchControl()
	bc.ServiceClassCode = 220
	bc.CompanyIdentification = "123456789"
	bc.EntryHash = 1
	bc.ODFIIdentification = 6200001
	bc.BatchNumber = 1
	return bc
}

// TestParseBatchControl parses a known Batch ControlRecord string.
func TestParseBatchControl(t *testing.T) {
	var line = "82250000010005320001000000010500000000000000origid                             076401250000001"
	r := NewReader(strings.NewReader(line))
	r.line = line
	r.addCurrentBatch(NewBatchPPD())
	bh := BatchHeader{BatchNumber: 1,
		ServiceClassCode:      225,
		CompanyIdentification: "origid",
		ODFIIdentification:    7640125}
	r.currentBatch.SetHeader(&bh)

	r.currentBatch.AddEntry(&EntryDetail{TransactionCode: 27, Amount: 10500, RDFIIdentification: 5320001, TraceNumber: 76401255655291})
	//fmt.Printf("%+v \n", r.line)
	err := r.parseBatchControl()
	if err != nil {
		t.Errorf("unknown error: %v", err)
	}
	record := r.currentBatch.GetControl()

	if record.recordType != "8" {
		t.Errorf("RecordType Expected '8' got: %v", record.recordType)
	}
	if record.ServiceClassCode != 225 {
		t.Errorf("ServiceClassCode Expected '225' got: %v", record.ServiceClassCode)
	}
	if record.EntryAddendaCountField() != "000001" {
		t.Errorf("EntryAddendaCount Expected '000001' got: %v", record.EntryAddendaCountField())
	}
	if record.EntryHashField() != "0005320001" {
		t.Errorf("EntryHash Expected '0005320001' got: %v", record.EntryHashField())
	}
	if record.TotalDebitEntryDollarAmountField() != "000000010500" {
		t.Errorf("TotalDebitEntryDollarAmount Expected '000000010500' got: %v", record.TotalDebitEntryDollarAmountField())
	}
	if record.TotalCreditEntryDollarAmountField() != "000000000000" {
		t.Errorf("TotalCreditEntryDollarAmount Expected '000000000000' got: %v", record.TotalCreditEntryDollarAmountField())
	}
	if record.CompanyIdentificationField() != "origid    " {
		t.Errorf("CompanyIdentification Expected 'origid    ' got: %v", record.CompanyIdentificationField())
	}
	if record.MessageAuthenticationCodeField() != "                   " {
		t.Errorf("MessageAuthenticationCode Expected '                   ' got: %v", record.MessageAuthenticationCodeField())
	}
	if record.reserved != "      " {
		t.Errorf("Reserved Expected '      ' got: %v", record.reserved)
	}
	if record.ODFIIdentificationField() != "07640125" {
		t.Errorf("OdfiIdentification Expected '07640125' got: %v", record.ODFIIdentificationField())
	}
	if record.BatchNumberField() != "0000001" {
		t.Errorf("BatchNumber Expected '0000001' got: %v", record.BatchNumberField())
	}
}

// TestBCString validats that a known parsed file can be return to a string of the same value
func TestBCString(t *testing.T) {
	var line = "82250000010005320001000000010500000000000000origid                             076401250000001"
	r := NewReader(strings.NewReader(line))
	r.line = line
	r.addCurrentBatch(NewBatchPPD())
	bh := BatchHeader{BatchNumber: 1,
		ServiceClassCode:      225,
		CompanyIdentification: "origid",
		ODFIIdentification:    7640125}
	r.currentBatch.SetHeader(&bh)

	r.currentBatch.AddEntry(&EntryDetail{TransactionCode: 27, Amount: 10500, RDFIIdentification: 5320001, TraceNumber: 76401255655291})
	err := r.parseBatchControl()
	if err != nil {
		t.Errorf("unknown error: %v", err)
	}
	record := r.currentBatch.GetControl()

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestValidateBCRecordType ensure error if recordType is not 8
func TestValidateBCRecordType(t *testing.T) {
	bc := mockBatchControl()
	bc.recordType = "2"
	if err := bc.Validate(); err != nil {
		if !strings.Contains(err.Error(), ErrRecordType.Error()) {
			t.Errorf("Expected RecordType Error got: %v", err)
		}
	}
}

func TestBCisServiceClassErr(t *testing.T) {
	bc := mockBatchControl()
	// works properly
	if err := bc.Validate(); err != nil {
		t.Errorf("Unexpected BatchControl error: %v", err.Error())
	}
	// create error is mismatch
	bc.ServiceClassCode = 123
	if err := bc.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected BatchControl error: %v", err.Error())
		}
	}
}

func TestBCFieldInclusion(t *testing.T) {
	bc := mockBatchControl()
	// works properly
	if err := bc.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected BatchControl error: %v", err.Error())
		}
	}
	// create error is mismatch
	bc.BatchNumber = 0
	if err := bc.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected BatchControl error: %v", err.Error())
		}
	}
}

func TestBCCompanyIdentificationAlphaNumeric(t *testing.T) {
	bc := mockBatchControl()
	// works properly
	if err := bc.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected BatchControl error: %v", err.Error())
		}
	}
	// create error is mismatch
	bc.CompanyIdentification = "@!"
	if err := bc.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected BatchControl error: %v", err.Error())
		}
	}
}

func TestBCMessageAuthenticationCodeAlphaNumeric(t *testing.T) {
	bc := mockBatchControl()
	// works properly
	if err := bc.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected BatchControl error: %v", err.Error())
		}
	}
	// create error is mismatch
	bc.MessageAuthenticationCode = "@!"
	if err := bc.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected BatchControl error: %v", err.Error())
		}
	}
}

func TestBCFieldInclusionRecordType(t *testing.T) {
	bc := mockBatchControl()
	bc.recordType = ""
	if err := bc.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBCFieldInclusionServiceClassCode(t *testing.T) {
	bc := mockBatchControl()
	bc.ServiceClassCode = 0
	if err := bc.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBCFieldInclusionODFIIdentification(t *testing.T) {
	bc := mockBatchControl()
	bc.ODFIIdentification = 0
	if err := bc.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}
