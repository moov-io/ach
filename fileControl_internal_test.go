// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
)

func mockFileControl() FileControl {
	fc := NewFileControl()
	fc.BatchCount = 1
	fc.BlockCount = 1
	fc.EntryAddendaCount = 1
	fc.EntryHash = 5320001
	return fc
}

// TestParseFileControl parses a known File Control Record string.
func TestParseFileControl(t *testing.T) {
	var line = "9000001000001000000010005320001000000010500000000000000                                       "
	r := NewReader(strings.NewReader(line))
	r.line = line
	err := r.parseFileControl()
	if err != nil {
		t.Errorf("unknown error: %v", err)
	}
	record := r.File.Control

	if record.recordType != "9" {
		t.Errorf("RecordType Expected '9' got: %v", record.recordType)
	}
	if record.BatchCountField() != "000001" {
		t.Errorf("BatchCount Expected '000001' got: %v", record.BatchCountField())
	}
	if record.BlockCountField() != "000001" {
		t.Errorf("BlockCount Expected '000001' got: %v", record.BlockCountField())
	}
	if record.EntryAddendaCountField() != "00000001" {
		t.Errorf("EntryAddendaCount Expected '00000001' got: %v", record.EntryAddendaCountField())
	}
	if record.EntryHashField() != "0005320001" {
		t.Errorf("EntryHash Expected '0005320001' got: %v", record.EntryHashField())
	}
	if record.TotalDebitEntryDollarAmountInFileField() != "000000010500" {
		t.Errorf("TotalDebitEntryDollarAmountInFile Expected '0005000000010500' got: %v", record.TotalDebitEntryDollarAmountInFileField())
	}
	if record.TotalCreditEntryDollarAmountInFileField() != "000000000000" {
		t.Errorf("TotalCreditEntryDollarAmountInFile Expected '000000000000' got: %v", record.TotalCreditEntryDollarAmountInFileField())
	}
	if record.reserved != "                                       " {
		t.Errorf("Reserved Expected '                                       ' got: %v", record.reserved)
	}
}

// TestFCString validats that a known parsed file can be return to a string of the same value
func TestFCString(t *testing.T) {
	var line = "9000001000001000000010005320001000000010500000000000000                                       "
	r := NewReader(strings.NewReader(line))
	r.line = line
	err := r.parseFileControl()
	if err != nil {
		t.Errorf("unknown error: %v", err)
	}
	record := r.File.Control
	if record.String() != line {
		t.Errorf("\nStrings do not match %s\n %s", line, record.String())
	}
}

// TestValidateFCRecordType ensure error if recordType is not 9
func TestValidateFCRecordType(t *testing.T) {
	fc := mockFileControl()
	fc.recordType = "2"

	if err := fc.Validate(); err != nil {
		if !strings.Contains(err.Error(), ErrRecordType.Error()) {
			t.Errorf("Expected RecordType Error got: %v", err)
		}
	}
}

func TestFCFieldInclusion(t *testing.T) {
	fc := mockFileControl()
	// works properly
	if err := fc.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// create error is mismatch
	fc.BatchCount = 0
	if err := fc.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestFCFieldInclusionRecordType(t *testing.T) {
	fc := mockFileControl()
	fc.recordType = ""
	if err := fc.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestFCFieldInclusionBlockCount(t *testing.T) {
	fc := mockFileControl()
	fc.BlockCount = 0
	if err := fc.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestFCFieldInclusionEntryAddendaCount(t *testing.T) {
	fc := mockFileControl()
	fc.EntryAddendaCount = 0
	if err := fc.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestFCFieldInclusionEntryHash(t *testing.T) {
	fc := mockFileControl()
	fc.EntryHash = 0
	if err := fc.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}
