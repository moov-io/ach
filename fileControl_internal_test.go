// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
)

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
	fc := NewFileControl()
	fc.recordType = "2"

	if err := fc.Validate(); err != nil {
		if !strings.Contains(err.Error(), ErrRecordType.Error()) {
			t.Errorf("Expected RecordType Error got: %v", err)
		}

	}
}
