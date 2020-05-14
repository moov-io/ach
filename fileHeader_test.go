// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package ach

import (
	"strings"
	"testing"
	"time"

	"github.com/moov-io/base"
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

func TestFileHeader__ImmediateOrigin(t *testing.T) {
	// From https://github.com/moov-io/ach/issues/510
	// We should allow a blank space or '1' followed by a 9 digit routing number and 9 digits
	header := NewFileHeader()
	header.ImmediateOrigin = " 123456789" // ' ' + routing number
	if v := header.ImmediateOriginField(); v != " 123456789" {
		t.Errorf("got %q", v)
	}
	header.ImmediateOrigin = "123456789" // 9 digit routing number
	if v := header.ImmediateOriginField(); v != " 123456789" {
		t.Errorf("got %q", v)
	}
	header.ImmediateOrigin = trimImmediateOriginLeadingZero("0123456789") // 0 + routing number
	if v := header.ImmediateOriginField(); v != " 123456789" {
		t.Errorf("got %q", v)
	}
	header.ImmediateOrigin = trimImmediateOriginLeadingZero("1123456789") // 1 + routing number
	if v := header.ImmediateOriginField(); v != " 112345678" {
		t.Errorf("got %q", v)
	}

	// Test with BypassOriginValidation
	header.SetValidation(&ValidateOpts{BypassOriginValidation: true})
	header.ImmediateOrigin = "1234567899"
	if v := header.ImmediateOriginField(); v != header.ImmediateOrigin {
		t.Errorf("got %q", v)
	}

	// make sure our trim works that we hook into FileHeader.Parse(..)
	if v := trimImmediateOriginLeadingZero("0123456789"); v != "123456789" {
		t.Errorf("got %q", v)
	}
	if v := trimImmediateOriginLeadingZero("0012345678"); v != "012345678" {
		t.Errorf("got %q", v)
	}
	zeros := strings.Repeat("0", 10)
	if v := trimImmediateOriginLeadingZero(zeros); v != zeros {
		t.Errorf("got %q", v)
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
		t.Errorf("ImmediateOrigin Expected ' 076401251' got: %v", record.ImmediateOriginField())
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
	err := fh.Validate()
	if !base.Match(err, NewErrRecordType(1)) {
		t.Errorf("%T: %s", err, err)
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
	err := fh.Validate()
	if !base.Match(err, ErrUpperAlpha) {
		t.Errorf("%T: %s", err, err)
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
	err := fh.Validate()
	if !base.Match(err, ErrRecordSize) {
		t.Errorf("%T: %s", err, err)
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
	err := fh.Validate()
	if !base.Match(err, ErrBlockingFactor) {
		t.Errorf("%T: %s", err, err)
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
	err := fh.Validate()
	if !base.Match(err, ErrFormatCode) {
		t.Errorf("%T: %s", err, err)
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
	err := fh.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
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
	err := fh.Validate()
	if !base.Match(err, ErrUpperAlpha) {
		t.Errorf("%T: %s", err, err)
	}

	fh.FileIDModifier = "AA"
	err = fh.Validate()
	if !base.Match(err, NewErrValidFieldLength(1)) {
		t.Errorf("%T: %s", err, err)
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
	err := fh.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
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
	err := fh.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
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
	err := fh.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
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
	err := fh.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
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
	err := fh.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
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
	err := fh.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
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
	err := fh.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
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
	err := fh.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
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
	err := fh.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
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

func TestFHImmediateDestinationInvalidLength(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateDestination = "198387"
	err := fh.Validate()
	if !strings.Contains(err.Error(), "invalid routing number length") {
		t.Errorf("%T: %s", err, err)
	}
}

func TestFHImmediateDestinationInvalidCheckSum(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateDestination = "121042880"
	err := fh.Validate()
	if !strings.Contains(err.Error(), "routing number checksum mismatch") {
		t.Errorf("%T: %s", err, err)
	}
}

func TestFHImmediateOriginValidate(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateOrigin = "0000000000"
	if err := fh.Validate(); err == nil {
		t.Error("expected error")
	}

	// use an alphanumeric code (NACHA rules allow this with specific
	// agreements between the ODFI and originator)
	fh.ImmediateOrigin = "ABC124"
	if err := fh.Validate(); err != nil {
		t.Error(err)
	}
}

func TestFHFieldInclusionFileCreationDate(t *testing.T) {
	fh := mockFileHeader()
	fh.FileCreationDate = ""
	if err := fh.Validate(); !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// testFileHeaderCreationDate validates creation date field inclusion
func testFileHeaderCreationDate(t testing.TB) {
	fh := mockFileHeader()
	fh.FileCreationDate = time.Now().AddDate(0, 0, 1).Format("060102") // YYMMDD
	if err := fh.Validate(); !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}

	fh.FileCreationDate = time.Now().Format(base.ISO8601Format)
	if err := fh.Validate(); !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
	yymmdd := time.Now().Format("060102")
	if v := fh.FileCreationDateField(); v != yymmdd {
		t.Errorf("got %q", v)
	}

	fh.FileCreationDate = "      "
	if err := fh.Validate(); !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}

	fh.FileCreationDate = ""
	if v := fh.FileCreationDateField(); len(v) != 6 {
		t.Errorf("got %q", v)
	}

	fh.FileCreationDate = "05/01/2019" // non ISO 8601 date
	if err := fh.Validate(); !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
	if v := fh.FileCreationDateField(); v != "" {
		t.Errorf("got %q", v)
	}
}

// TestFileHeaderCreationDate tests validating creation date field inclusion
func TestFileHeaderCreationDate(t *testing.T) {
	testFileHeaderCreationDate(t)
}

// BenchmarkFileHeaderCreationDate benchmarks validating creation date field inclusion
func BenchmarkFileHeaderCreationDate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileHeaderCreationDate(b)
	}
}

// testFileHeaderCreationTime validates creation date field inclusion
func testFileHeaderCreationTime(t testing.TB) {
	fh := mockFileHeader()
	fh.FileCreationTime = time.Now().AddDate(0, 0, 1).Format("1504") // HHmm
	if err := fh.Validate(); !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}

	fh.FileCreationTime = time.Now().Format(base.ISO8601Format)
	if err := fh.Validate(); !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
	if v := fh.FileCreationTimeField(); len(v) != 4 {
		t.Errorf("got %q", v)
	}

	fh.FileCreationTime = "    "
	if err := fh.Validate(); !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}

	fh.FileCreationTime = ""
	if v := fh.FileCreationTimeField(); len(v) != 4 {
		t.Errorf("got %q", v)
	}

	fh.FileCreationTime = "05/01/2019" // non ISO 8601 date
	if err := fh.Validate(); !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
	if v := fh.FileCreationTimeField(); v != "" {
		t.Errorf("got %q", v)
	}
}

// TestFileHeaderCreationTime tests validating creation date field inclusion
func TestFileHeaderCreationTime(t *testing.T) {
	testFileHeaderCreationTime(t)
}

// BenchmarkFileHeaderCreationTime benchmarks validating creation date field inclusion
func BenchmarkFileHeaderCreationTime(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileHeaderCreationTime(b)
	}
}

func TestFileHeader__ValidateOrigin(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateOrigin = "0000000000"

	err := fh.ValidateWith(&ValidateOpts{
		RequireABAOrigin: true,
	})
	if err != nil {
		if !strings.Contains(err.Error(), ErrConstructor.Error()) {
			t.Errorf("unexpected error: %v", err)
		}
	}

	err = fh.ValidateWith(&ValidateOpts{})
	if err == nil {
		t.Error("expected error")
	}

	err = fh.ValidateWith(&ValidateOpts{
		BypassOriginValidation: true,
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestFileHeader__SetValidation(t *testing.T) {
	fh := mockFileHeader()
	fh.SetValidation(nil)
	fh.SetValidation(&ValidateOpts{})
}

func TestFileHeader__ValidateDestination(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateDestination = "123456789" // invalid routing number

	if err := fh.ValidateWith(&ValidateOpts{}); err != nil {
		if !strings.Contains(err.Error(), "routing number checksum mismatch") {
			t.Error(err)
		}
	}

	// skip ImmediateDestination validation
	if err := fh.ValidateWith(&ValidateOpts{BypassDestinationValidation: true}); err != nil {
		t.Error(err)
	}
}
