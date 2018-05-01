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
	bc.CompanyIdentification = "121042882"
	bc.ODFIIdentification = "12104288"
	return bc
}

func TestMockBatchControl(t *testing.T) {
	bc := mockBatchControl()
	if err := bc.Validate(); err != nil {
		t.Error("mockBatchControl does not validate and will break other tests")
	}
	if bc.ServiceClassCode != 220 {
		t.Error("ServiceClassCode depedendent default value has changed")
	}
	if bc.CompanyIdentification != "121042882" {
		t.Error("CompanyIdentification depedendent default value has changed")
	}
	if bc.ODFIIdentification != "12104288" {
		t.Error("ODFIIdentification depedendent default value has changed")
	}
}

// TestParseBatchControl parses a known Batch ControlRecord string.
func TestParseBatchControl(t *testing.T) {
	var line = "82250000010005320001000000010500000000000000origid                             076401250000001"
	r := NewReader(strings.NewReader(line))
	r.line = line
	bh := BatchHeader{BatchNumber: 1,
		ServiceClassCode:      225,
		CompanyIdentification: "origid",
		ODFIIdentification:    "7640125"}
	r.addCurrentBatch(NewBatchPPD(&bh))

	r.currentBatch.AddEntry(&EntryDetail{TransactionCode: 27, Amount: 10500, RDFIIdentification: "5320001", TraceNumber: 76401255655291})
	if err := r.parseBatchControl(); err != nil {
		t.Errorf("%T: %s", err, err)
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
	bh := BatchHeader{BatchNumber: 1,
		ServiceClassCode:      225,
		CompanyIdentification: "origid",
		ODFIIdentification:    "7640125"}
	r.addCurrentBatch(NewBatchPPD(&bh))

	r.currentBatch.AddEntry(&EntryDetail{TransactionCode: 27, Amount: 10500, RDFIIdentification: "5320001", TraceNumber: 76401255655291})
	if err := r.parseBatchControl(); err != nil {
		t.Errorf("%T: %s", err, err)
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
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestBCisServiceClassErr(t *testing.T) {
	bc := mockBatchControl()
	bc.ServiceClassCode = 123
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestBCBatchNumber(t *testing.T) {
	bc := mockBatchControl()
	bc.BatchNumber = 0
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "BatchNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestBCCompanyIdentificationAlphaNumeric(t *testing.T) {
	bc := mockBatchControl()
	bc.CompanyIdentification = "®"
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CompanyIdentification" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestBCMessageAuthenticationCodeAlphaNumeric(t *testing.T) {
	bc := mockBatchControl()
	bc.MessageAuthenticationCode = "®"
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "MessageAuthenticationCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestBCFieldInclusionRecordType(t *testing.T) {
	bc := mockBatchControl()
	bc.recordType = ""
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestBCFieldInclusionServiceClassCode(t *testing.T) {
	bc := mockBatchControl()
	bc.ServiceClassCode = 0
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestBCFieldInclusionODFIIdentification(t *testing.T) {
	bc := mockBatchControl()
	bc.ODFIIdentification = "000000000"
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestBatchControlLength(t *testing.T) {
	bc := NewBatchControl()
	recordLength := len(bc.String())
	if recordLength != 94 {
		t.Errorf("Instantiated length of Batch Control string is not 94 but %v", recordLength)
	}
}
