// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE File.

package ach

import (
	"strings"
	"testing"
	"time"
)

// mockFileHeader build a validate File Header for tests
func mockFileHeader() FileHeader {
	fh := NewFileHeader()
	fh.ImmediateDestination = "231380104"
	fh.ImmediateOrigin = "121042882"
	fh.FileCreationDate = time.Now().AddDate(0, 0, 1).Format("060102") // YYMMDD
	fh.ImmediateDestinationName = "Federal Reserve Bank"
	fh.ImmediateOriginName = "My Bank Name"
	return fh
}

// testMockFileHeader validates a file header
func testMockFileHeader(t testing.TB) {
	fh := mockFileHeader()
	if err := fh.Validate(); err != nil {
		t.Error("mockFileHeader does not validate and will break other tests")
	}
	if fh.ImmediateDestination != "231380104" {
		t.Error("ImmediateDestination dependent default value has changed")
	}
	if fh.ImmediateOrigin != "121042882" {
		t.Error("ImmediateOrigin dependent default value has changed")
	}
	if fh.ImmediateDestinationName != "Federal Reserve Bank" {
		t.Error("ImmediateDestinationName dependent default value has changed")
	}
	if fh.ImmediateOriginName != "My Bank Name" {
		t.Error("ImmediateOriginName dependent default value has changed")
	}
}

// TestMockFileHeader tests validating a file header
func TestMockFileHeader(t *testing.T) {
	testMockFileHeader(t)
}

// BenchmarkMockFileHeader benchmarks validating a file header
func BenchmarkMockFileHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockFileHeader(b)
	}
}

// parseFileHeader validates parsing a file header
func parseFileHeader(t testing.TB) {
	var line = "101 076401251 0764012511807291511A094101achdestname            companyname                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	if err := r.parseFileHeader(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.File.Header

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

	if record.FileCreationDateField() != "180729" {
		t.Errorf("FileCreationDate Expected '180729' got:'%v'", record.FileCreationDateField())
	}

	if record.FileCreationTimeField() != "1511" {
		t.Errorf("FileCreationTime Expected '1900' got:'%v'", record.FileCreationTimeField())
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
		t.Errorf("ImmediateOriginName Expected 'companyname          ' got: '%v'", record.ImmediateOriginNameField())
	}
	if record.ReferenceCodeField() != "        " {
		t.Errorf("ReferenceCode Expected '        ' got:'%v'", record.ReferenceCodeField())
	}
}

// TestParseFileHeader test validates parsing a file header
func TestParseFileHeader(t *testing.T) {
	parseFileHeader(t)
}

// BenchmarkParseFileHeader benchmark validates parsing a file header
func BenchmarkParseFileHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parseFileHeader(b)
	}
}

// testFHString validates that a known parsed file can return to a string of the same value
func testFHString(t testing.TB) {
	var line = "101 076401251 0764012511807291511A094101achdestname            companyname                    "
	r := NewReader(strings.NewReader(line))
	r.line = line
	if err := r.parseFileHeader(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFHString tests validating that a known parsed file can return to a string of the same value
func TestFHString(t *testing.T) {
	testFHString(t)
}

// BenchmarkFHString benchmarks validating that a known parsed file
// can return to a string of the same value
func BenchmarkFHString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFHString(b)
	}
}

// testValidateFHRecordType validates error if record type is not 1
func testValidateFHRecordType(t testing.TB) {
	fh := mockFileHeader()
	fh.recordType = "2"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestValidateFHRecordType tests validating error if record type is not 1
func TestValidateFHRecordType(t *testing.T) {
	testValidateFHRecordType(t)
}

// BenchmarkValidateFHRecordType benchmarks validating error if record type is not 1
func BenchmarkValidateFHRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateFHRecordType(b)
	}
}

// testValidateIDModifier validates ID modifier is upper alphanumeric
func testValidateIDModifier(t testing.TB) {
	fh := mockFileHeader()
	fh.FileIDModifier = "®"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "FileIDModifier" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestValidateIDModifier tests validating ID modifier is upper alphanumeric
func TestValidateIDModifier(t *testing.T) {
	testValidateIDModifier(t)
}

// BenchmarkValidateIDModifier benchmarks validating ID modifier is upper alphanumeric
func BenchmarkValidateIDModifier(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateIDModifier(b)
	}
}

// testValidateRecordSize validates record size is "094"
func testValidateRecordSize(t testing.TB) {
	fh := mockFileHeader()
	fh.recordSize = "666"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordSize" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestValidateRecordSize tests validating record size is "094"
func TestValidateRecordSize(t *testing.T) {
	testValidateRecordSize(t)
}

// BenchmarkValidateRecordSize benchmarks validating record size is "094"
func BenchmarkValidateRecordSize(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateRecordSize(b)
	}
}

// testBlockingFactor validates blocking factor  is "10"
func testBlockingFactor(t testing.TB) {
	fh := mockFileHeader()
	fh.blockingFactor = "99"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "blockingFactor" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestBlockingFactor tests validating blocking factor  is "10"
func TestBlockingFactor(t *testing.T) {
	testBlockingFactor(t)
}

// BenchmarkBlockingFactor benchmarks validating blocking factor  is "10"
func BenchmarkBlockingFactor(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBlockingFactor(b)
	}
}

// testFormatCode validates format code is "1"
func testFormatCode(t testing.TB) {
	fh := mockFileHeader()
	fh.formatCode = "2"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "formatCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFormatCode tests validating format code is "1"
func TestFormatCode(t *testing.T) {
	testFormatCode(t)
}

// BenchmarkFormatCode benchmarks validating format code is "1"
func BenchmarkFormatCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFormatCode(b)
	}
}

// testFHFieldInclusion validates file header field inclusion
func testFHFieldInclusion(t testing.TB) {
	fh := mockFileHeader()
	fh.ImmediateOrigin = ""
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFHFieldInclusion tests validating file header field inclusion
func TestFHFieldInclusion(t *testing.T) {
	testFHFieldInclusion(t)
}

// BenchmarkFHFieldInclusion benchmarks validating file header field inclusion
func BenchmarkFHFieldInclusion(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFHFieldInclusion(b)
	}
}

// testUpperLengthFileID validates file ID
func testUpperLengthFileID(t testing.TB) {
	fh := mockFileHeader()
	fh.FileIDModifier = "a"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "FileIDModifier" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}

	fh.FileIDModifier = "AA"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "FileIDModifier" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestUpperLengthFileID  tests validating file ID
func TestUpperLengthFileID(t *testing.T) {
	testUpperLengthFileID(t)
}

// BenchmarkUpperLengthFileID  benchmarks validating file ID
func BenchmarkUpperLengthFileID(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testUpperLengthFileID(b)
	}
}

// testImmediateDestinationNameAlphaNumeric validates immediate destination name is alphanumeric
func testImmediateDestinationNameAlphaNumeric(t testing.TB) {
	fh := mockFileHeader()
	fh.ImmediateDestinationName = "Super Big Bank"
	fh.ImmediateDestinationName = "Big ®$$ Bank"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ImmediateDestinationName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestImmediateDestinationNameAlphaNumeric tests validating
// immediate destination name is alphanumeric
func TestImmediateDestinationNameAlphaNumeric(t *testing.T) {
	testImmediateDestinationNameAlphaNumeric(t)
}

// BenchmarkImmediateDestinationNameAlphaNumeric benchmarks validating
// immediate destination name is alphanumeric
func BenchmarkImmediateDestinationNameAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testImmediateDestinationNameAlphaNumeric(b)
	}
}

// testImmediateOriginNameAlphaNumeric validates immediate origin name is alphanumeric
func testImmediateOriginNameAlphaNumeric(t testing.TB) {
	fh := mockFileHeader()
	fh.ImmediateOriginName = "Super Big Bank"
	fh.ImmediateOriginName = "Bigger ®$$ Bank"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ImmediateOriginName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestImmediateOriginNameAlphaNumeric tests validating immediate origin name is alphanumeric
func TestImmediateOriginNameAlphaNumeric(t *testing.T) {
	testImmediateOriginNameAlphaNumeric(t)
}

// BenchmarkImmediateOriginNameAlphaNumeric benchmarks validating
// immediate origin name is alphanumeric
func BenchmarkImmediateOriginNameAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testImmediateOriginNameAlphaNumeric(b)
	}
}

// testImmediateReferenceCodeAlphaNumeric validates immediate reference is alphanumeric
func testImmediateReferenceCodeAlphaNumeric(t testing.TB) {
	fh := mockFileHeader()
	fh.ReferenceCode = " "
	fh.ReferenceCode = "®"
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ReferenceCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestImmediateReferenceCodeAlphaNumeric tests validating immediate reference is alphanumeric
func TestImmediateReferenceCodeAlphaNumeric(t *testing.T) {
	testImmediateReferenceCodeAlphaNumeric(t)
}

// BenchmarkImmediateReferenceCodeAlphaNumeric benchmarks validating immediate reference is alphanumeric
func BenchmarkImmediateReferenceCodeAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testImmediateReferenceCodeAlphaNumeric(b)
	}
}

// testFHFieldInclusionRecordType validates field inclusion
func testFHFieldInclusionRecordType(t testing.TB) {
	fh := mockFileHeader()
	fh.recordType = ""
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFHFieldInclusionRecordType tests validating field inclusion
func TestFHFieldInclusionRecordType(t *testing.T) {
	testFHFieldInclusionRecordType(t)
}

// BenchmarkFHFieldInclusionRecordType benchmarks validating field inclusion
func BenchmarkFHFieldInclusionRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFHFieldInclusionRecordType(b)
	}
}

// testFHFieldInclusionImmediateDestination validates immediate destination field inclusion
func testFHFieldInclusionImmediateDestination(t testing.TB) {
	fh := mockFileHeader()
	fh.ImmediateDestination = ""
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFHFieldInclusionImmediateDestination tests validates immediate destination field inclusion
func TestFHFieldInclusionImmediateDestination(t *testing.T) {
	testFHFieldInclusionImmediateDestination(t)
}

// BenchmarkFHFieldInclusionImmediateDestination benchmarks validates immediate destination field inclusion
func BenchmarkFHFieldInclusionImmediateDestination(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFHFieldInclusionImmediateDestination(b)
	}
}

// testFHFieldInclusionFileIDModifier validates file ID modifier field inclusion
func testFHFieldInclusionFileIDModifier(t testing.TB) {
	fh := mockFileHeader()
	fh.FileIDModifier = ""
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFHFieldInclusionFileIDModifier tests validating file ID modifier field inclusion
func TestFHFieldInclusionFileIDModifier(t *testing.T) {
	testFHFieldInclusionFileIDModifier(t)
}

// BenchmarkFHFieldInclusionFileIDModifier benchmarks validating file ID modifier field inclusion
func BenchmarkFHFieldInclusionFileIDModifier(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFHFieldInclusionFileIDModifier(b)
	}
}

// testFHFieldInclusionRecordSize validates record size field inclusion
func testFHFieldInclusionRecordSize(t testing.TB) {
	fh := mockFileHeader()
	fh.recordSize = ""
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFHFieldInclusionRecordSize tests validating record size field inclusion
func TestFHFieldInclusionRecordSize(t *testing.T) {
	testFHFieldInclusionRecordSize(t)
}

// BenchmarkFHFieldInclusionRecordSize benchmarks validating record size field inclusion
func BenchmarkFHFieldInclusionRecordSize(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFHFieldInclusionRecordSize(b)
	}
}

// testFHFieldInclusionBlockingFactor validates blocking factor field inclusion
func testFHFieldInclusionBlockingFactor(t testing.TB) {
	fh := mockFileHeader()
	fh.blockingFactor = ""
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFHFieldInclusionBlockingFactor tests validating blocking factor field inclusion
func TestFHFieldInclusionBlockingFactor(t *testing.T) {
	testFHFieldInclusionBlockingFactor(t)
}

// BenchmarkFHFieldInclusionBlockingFactor benchmarks
// validating blocking factor field inclusion
func BenchmarkFHFieldInclusionBlockingFactor(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFHFieldInclusionBlockingFactor(b)
	}
}

// testFHFieldInclusionFormatCode validates format code field inclusion
func testFHFieldInclusionFormatCode(t testing.TB) {
	fh := mockFileHeader()
	fh.formatCode = ""
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFHFieldInclusionFormatCode tests validating format code field inclusion
func TestFHFieldInclusionFormatCode(t *testing.T) {
	testFHFieldInclusionFormatCode(t)
}

// BenchmarkFHFieldInclusionFormatCode benchmarks validating format code field inclusion
func BenchmarkFHFieldInclusionFormatCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFHFieldInclusionFormatCode(b)
	}
}

// testFHFieldInclusionCreationDate validates creation date field inclusion
func testFHFieldInclusionCreationDate(t testing.TB) {
	fh := mockFileHeader()
	fh.FileCreationDate = time.Now().AddDate(0, 0, 1).Format("060102") // YYMMDD
	if err := fh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestFHFieldInclusionCreationDate tests validating creation date field inclusion
func TestFHFieldInclusionCreationDate(t *testing.T) {
	testFHFieldInclusionCreationDate(t)
}

// BenchmarkFHFieldInclusionCreationDate benchmarks validating creation date field inclusion
func BenchmarkFHFieldInclusionCreationDate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFHFieldInclusionCreationDate(b)
	}
}
