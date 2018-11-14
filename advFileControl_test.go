// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
)

// mockADVFileControl create a file control
func mockADVFileControl() ADVFileControl {
	fc := NewADVFileControl()
	fc.BatchCount = 1
	fc.BlockCount = 1
	fc.EntryAddendaCount = 1
	fc.EntryHash = 5320001
	return fc
}

// testMockADVFileControl validates a file control record
func testMockADVFileControl(t testing.TB) {
	fc := mockADVFileControl()
	if err := fc.Validate(); err != nil {
		t.Error("mockADVFileControl does not validate and will break other tests")
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

// TestMockADVFileControl tests validating a file control record
func TestMockADVFileControl(t *testing.T) {
	testMockADVFileControl(t)
}

// BenchmarkMockADVFileControl benchmarks validating a file control record
func BenchmarkMockADVFileControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockADVFileControl(b)
	}
}

// testParseADVFileControl parses a known file control record string
func testParseADVFileControl(t testing.TB) {
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

// TestParseADVFileControl tests parsing a known file control record string
func TestParseADVFileControl(t *testing.T) {
	testParseADVFileControl(t)
}

// BenchmarkParseADVFileControl benchmarks parsing a known file control record string
func BenchmarkParseADVFileControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testParseADVFileControl(b)
	}
}

// testADVFCString validates that a known parsed file can be return to a string of the same value
func testADVFCString(t testing.TB) {
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

// TestADVFCString tests validating that a known parsed file can be return to a string of the same value
func TestADVFCString(t *testing.T) {
	testADVFCString(t)
}

// BenchmarkADVFCString benchmarks validating that a known parsed file can be return to a string of the same value
func BenchmarkADVFCString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testADVFCString(b)
	}
}

// testValidateADVFCRecordType validates error if recordType is not 9
func testValidateADVFCRecordType(t testing.TB) {
	fc := mockADVFileControl()
	fc.recordType = "2"

	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestValidateADVFCRecordType tests validating error if recordType is not 9
func TestValidateADVFCRecordType(t *testing.T) {
	testValidateADVFCRecordType(t)
}

// BenchmarkValidateADVFCRecordType benchmarks validating error if recordType is not 9
func BenchmarkValidateADVFCRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateADVFCRecordType(b)
	}
}

// testADVFCFieldInclusion validates file control field inclusion
func testADVFCFieldInclusion(t testing.TB) {
	fc := mockADVFileControl()
	fc.BatchCount = 0
	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVFCFieldInclusion tests validating file control field inclusion
func TestADVFCFieldInclusion(t *testing.T) {
	testADVFCFieldInclusion(t)
}

// BenchmarkADVFCFieldInclusion benchmarks validating file control field inclusion
func BenchmarkADVFCFieldInclusion(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testADVFCFieldInclusion(b)
	}
}

// testADVFCFieldInclusionRecordType validates file control record type field inclusion
func testADVFCFieldInclusionRecordType(t testing.TB) {
	fc := mockADVFileControl()
	fc.recordType = ""
	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVFCFieldInclusionRecordType tests validating file control record type field inclusion
func TestADVFCFieldInclusionRecordType(t *testing.T) {
	testADVFCFieldInclusionRecordType(t)
}

// BenchmarkADVFCFieldInclusionRecordType benchmarks tests validating file control record type field inclusion
func BenchmarkADVFCFieldInclusionRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testADVFCFieldInclusionRecordType(b)
	}
}

// testADVFCFieldInclusionBlockCount validates file control block count field inclusion
func testADVFCFieldInclusionBlockCount(t testing.TB) {
	fc := mockADVFileControl()
	fc.BlockCount = 0
	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVFCFieldInclusionBlockCount tests validating file control block count field inclusion
func TestADVFCFieldInclusionBlockCount(t *testing.T) {
	testADVFCFieldInclusionBlockCount(t)
}

// BenchmarkADVFCFieldInclusionBlockCount benchmarks validating file control block count field inclusion
func BenchmarkADVFCFieldInclusionBlockCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testADVFCFieldInclusionBlockCount(b)
	}
}

// testADVFCFieldInclusionEntryAddendaCount validates file control addenda count field inclusion
func testADVFCFieldInclusionEntryAddendaCount(t testing.TB) {
	fc := mockADVFileControl()
	fc.EntryAddendaCount = 0
	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVFCFieldInclusionEntryAddendaCount tests validating file control addenda count field inclusion
func TestADVFCFieldInclusionEntryAddendaCount(t *testing.T) {
	testADVFCFieldInclusionEntryAddendaCount(t)
}

// BenchmarkADVFCFieldInclusionEntryAddendaCount benchmarks validating file control addenda count field inclusion
func BenchmarkADVFCFieldInclusionEntryAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testADVFCFieldInclusionEntryAddendaCount(b)
	}
}

// testADVFCFieldInclusionEntryHash validates file control entry hash field inclusion
func testADVFCFieldInclusionEntryHash(t testing.TB) {
	fc := mockADVFileControl()
	fc.EntryHash = 0
	if err := fc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVFCFieldInclusionEntryHash tests validating file control entry hash field inclusion
func TestADVFCFieldInclusionEntryHash(t *testing.T) {
	testADVFCFieldInclusionEntryHash(t)
}

// BenchmarkADVFCFieldInclusionEntryHash benchmarks validating file control entry hash field inclusion
func BenchmarkADVFCFieldInclusionEntryHash(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testADVFCFieldInclusionEntryHash(b)
	}
}

// TestInvalidADVFCParse returns an error when parsing an ADV File Control
func TestInvalidADVFCParse(t *testing.T) {
	var line = "9000001000001000000010005320001000000000000000105"
	r := NewReader(strings.NewReader(line))
	r.line = line
	batchADV := mockBatchADV()
	r.File.AddBatch(batchADV)

	if err := r.parseFileControl(); err != nil {
		if p, ok := err.(*ParseError); ok {
			if p.Record != "FileControl" {
				t.Errorf("%T: %s", p, p)
			}
		} else {
			t.Errorf("%T: %s", p.Err, p.Err)
		}
	}
}
