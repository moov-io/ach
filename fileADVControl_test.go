// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
)

// mockFileADVControl create a file control
func mockFileADVControl() FileADVControl {
	fc := NewFileADVControl()
	fc.BatchCount = 1
	fc.BlockCount = 1
	fc.EntryAddendaCount = 1
	fc.EntryHash = 5320001
	return fc
}

// testMockFileADVControl validates a file control record
func testMockFileADVControl(t testing.TB) {
	fc := mockFileADVControl()
	if err := fc.Validate(); err != nil {
		t.Error("mockFileADVControl does not validate and will break other tests")
	}
	if fc.BatchCount != 1 {
		t.Error("BatchCount dependent default value has changed")
	}
	if fc.BlockCount != 1 {
		t.Error("BlockCount dependent default value has changed")
	}
	if fc.EntryAddendaCount != 1 {
		t.Error("EntryAddendaCount dependent default value has changed")
	}
	if fc.EntryHash != 5320001 {
		t.Error("EntryHash dependent default value has changed")
	}
}

// TestMockFileADVControl tests validating a file control record
func TestMockFileADVControl(t *testing.T) {
	testMockFileADVControl(t)
}

// BenchmarkMockFileADVControl benchmarks validating a file control record
func BenchmarkMockFileADVControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockFileADVControl(b)
	}
}

// testParseFileADVControl parses a known file control record string
func testParseFileADVControl(t testing.TB) {
	var line = "90000010000010000000100053200010000000000000001050000000000000000000000                       "
	r := NewReader(strings.NewReader(line))
	r.line = line
	batchADV := mockBatchADV()
	r.File.AddBatch(batchADV)

	err := r.parseFileControl()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.File.ADVControl

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
	if record.TotalDebitEntryDollarAmountInFileField() != "00000000000000010500" {
		t.Errorf("TotalDebitEntryDollarAmountInFile Expected '0000000000000010500' got: %v", record.TotalDebitEntryDollarAmountInFileField())
	}
	if record.TotalCreditEntryDollarAmountInFileField() != "00000000000000000000" {
		t.Errorf("TotalCreditEntryDollarAmountInFile Expected '00000000000000000000' got: %v", record.TotalCreditEntryDollarAmountInFileField())
	}
	if record.reserved != "                       " {
		t.Errorf("Reserved Expected '                       ' got: %v", record.reserved)
	}
}

// TestParseFileADVControl tests parsing a known file control record string
func TestParseFileADVControl(t *testing.T) {
	testParseFileADVControl(t)
}

// BenchmarkParseFileADVControl benchmarks parsing a known file control record string
func BenchmarkParseFileADVControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testParseFileADVControl(b)
	}
}

// testFCADVString validates that a known parsed file can be return to a string of the same value
func testFCADVString(t testing.TB) {
	var line = "90000010000010000000100053200010000000000000001050000000000000000000000                       "
	r := NewReader(strings.NewReader(line))
	r.line = line
	batchADV := mockBatchADV()
	r.File.AddBatch(batchADV)

	err := r.parseFileControl()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.File.ADVControl
	if record.String() != line {
		t.Errorf("\nStrings do not match %s\n %s", line, record.String())
	}
}

// TestFCADVString tests validating that a known parsed file can be return to a string of the same value
func TestFCADVString(t *testing.T) {
	testFCADVString(t)
}

// BenchmarkFCADVString benchmarks validating that a known parsed file can be return to a string of the same value
func BenchmarkFCADVString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFCADVString(b)
	}
}

// testValidateFCADVRecordType validates error if recordType is not 9
func testValidateFCADVRecordType(t testing.TB) {
	fc := mockFileADVControl()
	fc.recordType = "2"

	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestValidateFCADVRecordType tests validating error if recordType is not 9
func TestValidateFCADVRecordType(t *testing.T) {
	testValidateFCADVRecordType(t)
}

// BenchmarkValidateFCADVRecordType benchmarks validating error if recordType is not 9
func BenchmarkValidateFCADVRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateFCADVRecordType(b)
	}
}

// testFCADVFieldInclusion validates file control field inclusion
func testFCADVFieldInclusion(t testing.TB) {
	fc := mockFileADVControl()
	fc.BatchCount = 0
	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFCADVFieldInclusion tests validating file control field inclusion
func TestFCADVFieldInclusion(t *testing.T) {
	testFCADVFieldInclusion(t)
}

// BenchmarkFCADVFieldInclusion benchmarks validating file control field inclusion
func BenchmarkFCADVFieldInclusion(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFCADVFieldInclusion(b)
	}
}

// testFCADVFieldInclusionRecordType validates file control record type field inclusion
func testFCADVFieldInclusionRecordType(t testing.TB) {
	fc := mockFileADVControl()
	fc.recordType = ""
	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFCADVFieldInclusionRecordType tests validating file control record type field inclusion
func TestFCADVFieldInclusionRecordType(t *testing.T) {
	testFCADVFieldInclusionRecordType(t)
}

// BenchmarkFCADVFieldInclusionRecordType benchmarks tests validating file control record type field inclusion
func BenchmarkFCADVFieldInclusionRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFCADVFieldInclusionRecordType(b)
	}
}

// testFCADVFieldInclusionBlockCount validates file control block count field inclusion
func testFCADVFieldInclusionBlockCount(t testing.TB) {
	fc := mockFileControl()
	fc.BlockCount = 0
	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFCADVFieldInclusionBlockCount tests validating file control block count field inclusion
func TestFCADVFieldInclusionBlockCount(t *testing.T) {
	testFCADVFieldInclusionBlockCount(t)
}

// BenchmarkFCADVFieldInclusionBlockCount benchmarks validating file control block count field inclusion
func BenchmarkFCADVFieldInclusionBlockCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFCADVFieldInclusionBlockCount(b)
	}
}

// testFCADVFieldInclusionEntryAddendaCount validates file control addenda count field inclusion
func testFCADVFieldInclusionEntryAddendaCount(t testing.TB) {
	fc := mockFileControl()
	fc.EntryAddendaCount = 0
	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFCADVFieldInclusionEntryAddendaCount tests validating file control addenda count field inclusion
func TestFCADVFieldInclusionEntryAddendaCount(t *testing.T) {
	testFCADVFieldInclusionEntryAddendaCount(t)
}

// BenchmarkFCADVFieldInclusionEntryAddendaCount benchmarks validating file control addenda count field inclusion
func BenchmarkFCADVFieldInclusionEntryAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFCADVFieldInclusionEntryAddendaCount(b)
	}
}

// testFCADVFieldInclusionEntryHash validates file control entry hash field inclusion
func testFCADVFieldInclusionEntryHash(t testing.TB) {
	fc := mockFileControl()
	fc.EntryHash = 0
	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFCADVFieldInclusionEntryHash tests validating file control entry hash field inclusion
func TestFCADVFieldInclusionEntryHash(t *testing.T) {
	testFCADVFieldInclusionEntryHash(t)
}

// BenchmarkFCADVFieldInclusionEntryHash benchmarks validating file control entry hash field inclusion
func BenchmarkFCADVFieldInclusionEntryHash(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFCADVFieldInclusionEntryHash(b)
	}
}
