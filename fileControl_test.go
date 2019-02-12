// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"

	"github.com/moov-io/base"
)

// mockFileControl create a file control
func mockFileControl() FileControl {
	fc := NewFileControl()
	fc.BatchCount = 1
	fc.BlockCount = 1
	fc.EntryAddendaCount = 1
	fc.EntryHash = 5320001
	return fc
}

// testMockFileControl validates a file control record
func testMockFileControl(t testing.TB) {
	fc := mockFileControl()
	if err := fc.Validate(); err != nil {
		t.Error("mockFileControl does not validate and will break other tests")
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

// TestMockFileControl tests validating a file control record
func TestMockFileControl(t *testing.T) {
	testMockFileControl(t)
}

// BenchmarkMockFileControl benchmarks validating a file control record
func BenchmarkMockFileControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockFileControl(b)
	}
}

// testParseFileControl parses a known file control record string
func testParseFileControl(t testing.TB) {
	var line = "9000001000001000000010005320001000000010500000000000000                                       "
	r := NewReader(strings.NewReader(line))
	r.line = line
	err := r.parseFileControl()
	if err != nil {
		t.Errorf("%T: %s", err, err)
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

// TestParseFileControl tests parsing a known file control record string
func TestParseFileControl(t *testing.T) {
	testParseFileControl(t)
}

// BenchmarkParseFileControl benchmarks parsing a known file control record string
func BenchmarkParseFileControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testParseFileControl(b)
	}
}

// testFCString validates that a known parsed file can be return to a string of the same value
func testFCString(t testing.TB) {
	var line = "9000001000001000000010005320001000000010500000000000000                                       "
	r := NewReader(strings.NewReader(line))
	r.line = line
	err := r.parseFileControl()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.File.Control
	if record.String() != line {
		t.Errorf("\nStrings do not match %s\n %s", line, record.String())
	}
}

// TestFCString tests validating that a known parsed file can be return to a string of the same value
func TestFCString(t *testing.T) {
	testFCString(t)
}

// BenchmarkFCString benchmarks validating that a known parsed file can be return to a string of the same value
func BenchmarkFCString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFCString(b)
	}
}

// testValidateFCRecordType validates error if recordType is not 9
func testValidateFCRecordType(t testing.TB) {
	fc := mockFileControl()
	fc.recordType = "2"

	err := fc.Validate()
	// TODO: are we expecting there to be an error here?
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestValidateFCRecordType tests validating error if recordType is not 9
func TestValidateFCRecordType(t *testing.T) {
	testValidateFCRecordType(t)
}

// BenchmarkValidateFCRecordType benchmarks validating error if recordType is not 9
func BenchmarkValidateFCRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateFCRecordType(b)
	}
}

// testFCFieldInclusion validates file control field inclusion
func testFCFieldInclusion(t testing.TB) {
	fc := mockFileControl()
	fc.BatchCount = 0
	err := fc.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFCFieldInclusion tests validating file control field inclusion
func TestFCFieldInclusion(t *testing.T) {
	testFCFieldInclusion(t)
}

// BenchmarkFCFieldInclusion benchmarks validating file control field inclusion
func BenchmarkFCFieldInclusion(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFCFieldInclusion(b)
	}
}

// testFCFieldInclusionRecordType validates file control record type field inclusion
func testFCFieldInclusionRecordType(t testing.TB) {
	fc := mockFileControl()
	fc.recordType = ""
	err := fc.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFCFieldInclusionRecordType tests validating file control record type field inclusion
func TestFCFieldInclusionRecordType(t *testing.T) {
	testFCFieldInclusionRecordType(t)
}

// BenchmarkFCFieldInclusionRecordType benchmarks tests validating file control record type field inclusion
func BenchmarkFCFieldInclusionRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFCFieldInclusionRecordType(b)
	}
}

// testFCFieldInclusionBlockCount validates file control block count field inclusion
func testFCFieldInclusionBlockCount(t testing.TB) {
	fc := mockFileControl()
	fc.BlockCount = 0
	err := fc.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFCFieldInclusionBlockCount tests validating file control block count field inclusion
func TestFCFieldInclusionBlockCount(t *testing.T) {
	testFCFieldInclusionBlockCount(t)
}

// BenchmarkFCFieldInclusionBlockCount benchmarks validating file control block count field inclusion
func BenchmarkFCFieldInclusionBlockCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFCFieldInclusionBlockCount(b)
	}
}

// testFCFieldInclusionEntryAddendaCount validates file control addenda count field inclusion
func testFCFieldInclusionEntryAddendaCount(t testing.TB) {
	fc := mockFileControl()
	fc.EntryAddendaCount = 0
	err := fc.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFCFieldInclusionEntryAddendaCount tests validating file control addenda count field inclusion
func TestFCFieldInclusionEntryAddendaCount(t *testing.T) {
	testFCFieldInclusionEntryAddendaCount(t)
}

// BenchmarkFCFieldInclusionEntryAddendaCount benchmarks validating file control addenda count field inclusion
func BenchmarkFCFieldInclusionEntryAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFCFieldInclusionEntryAddendaCount(b)
	}
}

// testFCFieldInclusionEntryHash validates file control entry hash field inclusion
func testFCFieldInclusionEntryHash(t testing.TB) {
	fc := mockFileControl()
	fc.EntryHash = 0
	err := fc.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFCFieldInclusionEntryHash tests validating file control entry hash field inclusion
func TestFCFieldInclusionEntryHash(t *testing.T) {
	testFCFieldInclusionEntryHash(t)
}

// BenchmarkFCFieldInclusionEntryHash benchmarks validating file control entry hash field inclusion
func BenchmarkFCFieldInclusionEntryHash(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFCFieldInclusionEntryHash(b)
	}
}
