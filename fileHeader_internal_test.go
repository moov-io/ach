// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
)

// TestParseFileHeader parses a known File Header Record string.
func TestParseFileHeader(t *testing.T) {
	var line = "101 076401251 0764012510807291511A094101achdestname            companyname                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	if err := r.parseFileHeader(); err != nil {
		t.Errorf("unknown error: %v", err)
	}
	record := r.file.Header

	if record.recordType != "1" {
		t.Errorf("RecordType Expected 1 got: %v", record.recordType)
	}
	if record.priorityCode != "01" {
		t.Errorf("PriorityCode Expected 01 got: %v", record.priorityCode)
	}
	if record.ImmediateDestinationField() != " 076401251" {
		t.Errorf("ImmediateDestination Expected ' 076401251' got: %v", record.ImmediateDestinationField())
	}
	if record.ImmediateOriginField() != " 076401251" {
		t.Errorf("ImmediateOrigin Expected '   076401251' got: %v", record.ImmediateOriginField())
	}

	if record.FileCreationDateField() != "080729" {
		t.Errorf("FileCreationDate Expected '080729' got:'%v'", record.FileCreationDateField())
	}
	if record.FileCreationTimeField() != "1511" {
		t.Errorf("FileCreationTime Expected '1511' got:'%v'", record.FileCreationTimeField())
	}

	if record.FileIDModifier != "A" {
		t.Errorf("FileIDModifier Expected 'A' got:'%v'", record.FileIDModifier)
	}
	if record.recordSize != "094" {
		t.Errorf("RecordSize Expected '094' got:'%v'", record.recordSize)
	}
	if record.blockingFactor != "10" {
		t.Errorf("BlockingFactor Expected '10' got:'%v'", record.blockingFactor)
	}
	if record.formatCode != "1" {
		t.Errorf("FormatCode Expected '1' got:'%v'", record.formatCode)
	}
	if record.ImmediateDestinationNameField() != "achdestname            " {
		t.Errorf("ImmediateDestinationName Expected 'achdestname           ' got:'%v'", record.ImmediateDestinationNameField())
	}
	if record.ImmediateOriginNameField() != "companyname            " {
		t.Errorf("ImmidiateOriginName Expected 'companyname          ' got: '%v'", record.ImmediateOriginNameField())
	}
	if record.ReferenceCodeField() != "        " {
		t.Errorf("ReferenceCode Expected '        ' got:'%v'", record.ReferenceCodeField())
	}
}

// TestString validats that a known parsed file can be return to a string of the same value
func TestFHString(t *testing.T) {
	var line = "101 076401251 0764012510807291511A094101achdestname            companyname                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	err := r.parseFileHeader()
	if err != nil {
		t.Errorf("unknown error: %v", err)
	}
	record := r.file.Header

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestValidateFHRecordType ensure error if recordType is not 1
func TestValidateFHRecordType(t *testing.T) {
	fh := NewFileHeader()
	fh.recordType = "2"
	if err := fh.Validate(); err != nil {
		if !strings.Contains(err.Error(), ErrRecordType.Error()) {
			t.Errorf("Expected RecordType Error got: %v", err)
		}
	}
}

// TestValidateIDModifier ensure ID Modiier is upper alphanumeric
func TestValidateIDModifier(t *testing.T) {
	fh := NewFileHeader()
	fh.FileIDModifier = "A"
	if err := fh.Validate(); err != nil {
		if !strings.Contains(err.Error(), ErrValidFieldLength.Error()) {
			t.Errorf("Expected ID Modifier Error got: %v", err)
		}
	}
}

// TestValidateRecordSize ensure record size is "094"
func TestValidateRecordSize(t *testing.T) {
	fh := NewFileHeader()
	fh.recordSize = "666"
	if err := fh.Validate(); err != nil {
		if !strings.Contains(err.Error(), ErrRecordSize.Error()) {
			t.Errorf("Expected Record Size Error got: %v", err)
		}
	}
}

// TestBlockingFactor ensure blocking factor  is "10"
func TestBlockingFactor(t *testing.T) {
	fh := NewFileHeader()
	fh.blockingFactor = "99"
	if err := fh.Validate(); err != nil {
		if !strings.Contains(err.Error(), ErrBlockingFactor.Error()) {
			t.Errorf("Expected Blocking Factor Error got: %v", err)
		}
	}
}

// TestFormatCode ensure format code is "1"
func TestFormatCode(t *testing.T) {
	fh := NewFileHeader()
	fh.formatCode = "2"
	if err := fh.Validate(); err != nil {
		if !strings.Contains(err.Error(), ErrFormatCode.Error()) {
			t.Errorf("Expected Format Code Error got: %v", err)
		}
	}
}
