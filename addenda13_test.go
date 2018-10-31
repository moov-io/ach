// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
)

// mockAddenda13 creates a mock Addenda13 record
func mockAddenda13() *Addenda13 {
	addenda13 := NewAddenda13()
	addenda13.ODFIName = "Wells Fargo"
	addenda13.ODFIIDNumberQualifier = "01"
	addenda13.ODFIIdentification = "121042882"
	addenda13.ODFIBranchCountryCode = "US"
	addenda13.EntryDetailSequenceNumber = 00000001
	return addenda13
}

// testAddenda13Parse parses Addenda13 record
func testAddenda13Parse(t testing.TB) {
	Addenda13 := NewAddenda13()
	line := "713Wells Fargo                        01121042882                         US           0000001"
	Addenda13.Parse(line)
	// walk the Addenda13 struct
	if Addenda13.recordType != "7" {
		t.Errorf("expected %v got %v", "7", Addenda13.recordType)
	}
	if Addenda13.typeCode() != "13" {
		t.Errorf("expected %v got %v", "13", Addenda13.TypeCode)
	}
	if Addenda13.ODFIName != "Wells Fargo" {
		t.Errorf("expected %v got %v", "Wells Fargo", Addenda13.ODFIName)
	}
	if Addenda13.ODFIIDNumberQualifier != "01" {
		t.Errorf("expected: %v got: %v", "01", Addenda13.ODFIIDNumberQualifier)
	}
	if Addenda13.ODFIIdentification != "121042882" {
		t.Errorf("expected: %v got: %v", "121042882", Addenda13.ODFIIdentification)
	}
	if Addenda13.ODFIBranchCountryCode != "US" {
		t.Errorf("expected: %s got: %s", "US", Addenda13.ODFIBranchCountryCode)
	}
	if Addenda13.reserved != "          " {
		t.Errorf("expected: %v got: %v", "          ", Addenda13.reserved)
	}
	if Addenda13.EntryDetailSequenceNumber != 0000001 {
		t.Errorf("expected: %v got: %v", 0000001, Addenda13.EntryDetailSequenceNumber)
	}
}

// TestAddenda13Parse tests parsing Addenda13 record
func TestAddenda13Parse(t *testing.T) {
	testAddenda13Parse(t)
}

// BenchmarkAddenda13Parse benchmarks parsing Addenda13 record
func BenchmarkAddenda13Parse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda13Parse(b)
	}
}

// TestMockAddenda13 validates mockAddenda13
func TestMockAddenda13(t *testing.T) {
	addenda13 := mockAddenda13()
	if err := addenda13.Validate(); err != nil {
		t.Error("mockAddenda13 does not validate and will break other tests")
	}
}

// testAddenda13ValidRecordType validates Addenda13 recordType
func testAddenda13ValidRecordType(t testing.TB) {
	addenda13 := mockAddenda13()
	addenda13.recordType = "63"
	if err := addenda13.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddenda13ValidRecordType tests validating Addenda13 recordType
func TestAddenda13ValidRecordType(t *testing.T) {
	testAddenda13ValidRecordType(t)
}

// BenchmarkAddenda13ValidRecordType benchmarks validating Addenda13 recordType
func BenchmarkAddenda13ValidRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda13ValidRecordType(b)
	}
}

// testAddenda13ValidTypeCode validates Addenda13 TypeCode
func testAddenda13ValidTypeCode(t testing.TB) {
	addenda13 := mockAddenda13()
	addenda13.TypeCode = "65"
	if err := addenda13.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddenda13ValidTypeCode tests validating Addenda13 TypeCode
func TestAddenda13ValidTypeCode(t *testing.T) {
	testAddenda13ValidTypeCode(t)
}

// BenchmarkAddenda13ValidTypeCode benchmarks validating Addenda13 TypeCode
func BenchmarkAddenda13ValidTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda13ValidTypeCode(b)
	}
}

// testAddenda13TypeCode13 TypeCode is 13 if TypeCode is a valid TypeCode
func testAddenda13TypeCode13(t testing.TB) {
	addenda13 := mockAddenda13()
	addenda13.TypeCode = "05"
	if err := addenda13.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddenda13TypeCode13 tests TypeCode is 13 if TypeCode is a valid TypeCode
func TestAddenda13TypeCode13(t *testing.T) {
	testAddenda13TypeCode13(t)
}

// BenchmarkAddenda13TypeCode13 benchmarks TypeCode is 13 if TypeCode is a valid TypeCode
func BenchmarkAddenda13TypeCode13(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda13TypeCode13(b)
	}
}

// testODFINameAlphaNumeric validates ODFIName is alphanumeric
func testODFINameAlphaNumeric(t testing.TB) {
	addenda13 := mockAddenda13()
	addenda13.ODFIName = "Wells®Fargo"
	if err := addenda13.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ODFIName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestODFINameAlphaNumeric tests validating ODFIName is alphanumeric
func TestODFINameAlphaNumeric(t *testing.T) {
	testODFINameAlphaNumeric(t)
}

// BenchmarkODFINameAlphaNumeric benchmarks validating ODFIName is alphanumeric
func BenchmarkODFINameAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testODFINameAlphaNumeric(b)
	}
}

// testODFIIDNumberQualifierValid validates ODFIIDNumberQualifier is valid
func testODFIIDNumberQualifierValid(t testing.TB) {
	addenda13 := mockAddenda13()
	addenda13.ODFIIDNumberQualifier = "®1"
	if err := addenda13.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ODFIIDNumberQualifier" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestODFIIDNumberQualifierValid tests validating ODFIIDNumberQualifier is valid
func TestODFIIDNumberQualifierValid(t *testing.T) {
	testODFIIDNumberQualifierValid(t)
}

// BenchmarkODFIIDNumberQualifierValid benchmarks validating ODFIIDNumberQualifier is valid
func BenchmarkODFIIDNumberQualifierValid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testODFIIDNumberQualifierValid(b)
	}
}

// testODFIIdentificationAlphaNumeric validates ODFIIdentification is alphanumeric
func testODFIIdentificationAlphaNumeric(t testing.TB) {
	addenda13 := mockAddenda13()
	addenda13.ODFIIdentification = "®121042882"
	if err := addenda13.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ODFIIdentification" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestODFIIdentificationAlphaNumeric tests validating ODFIIdentification is alphanumeric
func TestODFIIdentificationAlphaNumeric(t *testing.T) {
	testODFIIdentificationAlphaNumeric(t)
}

// BenchmarkODFIIdentificationAlphaNumeric benchmarks validating ODFIIdentification is alphanumeric
func BenchmarkODFIIdentificationAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testODFIIdentificationAlphaNumeric(b)
	}
}

// testODFIBranchCountryCodeAlphaNumeric validates ODFIBranchCountryCode is alphanumeric
func testODFIBranchCountryCodeAlphaNumeric(t testing.TB) {
	addenda13 := mockAddenda13()
	addenda13.ODFIBranchCountryCode = "U®"
	if err := addenda13.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ODFIBranchCountryCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestODFIBranchCountryCodeAlphaNumeric tests validating ODFIBranchCountryCode is alphanumeric
func TestODFIBranchCountryCodeAlphaNumeric(t *testing.T) {
	testODFIBranchCountryCodeAlphaNumeric(t)
}

// BenchmarkODFIBranchCountryCodeAlphaNumeric benchmarks validating ODFIBranchCountryCode is alphanumeric
func BenchmarkODFIBranchCountryCodeAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testODFIBranchCountryCodeAlphaNumeric(b)
	}
}

// testAddenda13FieldInclusionRecordType validates recordType fieldInclusion
func testAddenda13FieldInclusionRecordType(t testing.TB) {
	addenda13 := mockAddenda13()
	addenda13.recordType = ""
	if err := addenda13.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestAddenda13FieldInclusionRecordType tests validating recordType fieldInclusion
func TestAddenda13FieldInclusionRecordType(t *testing.T) {
	testAddenda13FieldInclusionRecordType(t)
}

// BenchmarkAddenda13FieldInclusionRecordType benchmarks validating recordType fieldInclusion
func BenchmarkAddenda13FieldInclusionRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda13FieldInclusionRecordType(b)
	}
}

// testAddenda13FieldInclusionTypeCode validates TypeCode fieldInclusion
func testAddenda13FieldInclusionTypeCode(t testing.TB) {
	addenda13 := mockAddenda13()
	addenda13.TypeCode = ""
	if err := addenda13.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestAddenda13FieldInclusionTypeCode tests validating TypeCode fieldInclusion
func TestAddenda13FieldInclusionTypeCode(t *testing.T) {
	testAddenda13FieldInclusionTypeCode(t)
}

// BenchmarkAddenda13FieldInclusionTypeCode benchmarks validating TypeCode fieldInclusion
func BenchmarkAddenda13FieldInclusionTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda13FieldInclusionTypeCode(b)
	}
}

// testAddenda13FieldInclusionODFIName validates ODFIName fieldInclusion
func testAddenda13FieldInclusionODFIName(t testing.TB) {
	addenda13 := mockAddenda13()
	addenda13.ODFIName = ""
	if err := addenda13.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestAddenda13FieldInclusionODFIName tests validating ODFIName fieldInclusion
func TestAddenda13FieldInclusionODFIName(t *testing.T) {
	testAddenda13FieldInclusionODFIName(t)
}

// BenchmarkAddenda13FieldInclusionODFIName benchmarks validating ODFIName fieldInclusion
func BenchmarkAddenda13FieldInclusionODFIName(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda13FieldInclusionODFIName(b)
	}
}

// testAddenda13FieldInclusionODFIIDNumberQualifier validates ODFIIDNumberQualifier fieldInclusion
func testAddenda13FieldInclusionODFIIDNumberQualifier(t testing.TB) {
	addenda13 := mockAddenda13()
	addenda13.ODFIIDNumberQualifier = ""
	if err := addenda13.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestAddenda13FieldInclusionODFIIDNumberQualifier tests validating ODFIIDNumberQualifier fieldInclusion
func TestAddenda13FieldInclusionODFIIDNumberQualifier(t *testing.T) {
	testAddenda13FieldInclusionODFIIDNumberQualifier(t)
}

// BenchmarkAddenda13FieldInclusionODFIIDNumberQualifier benchmarks validating ODFIIDNumberQualifier fieldInclusion
func BenchmarkAddenda13FieldInclusionODFIIDNumberQualifier(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda13FieldInclusionODFIIDNumberQualifier(b)
	}
}

// testAddenda13FieldInclusionODFIIdentification validates ODFIIdentification fieldInclusion
func testAddenda13FieldInclusionODFIIdentification(t testing.TB) {
	addenda13 := mockAddenda13()
	addenda13.ODFIIdentification = ""
	if err := addenda13.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestAddenda13FieldInclusionODFIIdentification tests validating ODFIIdentification fieldInclusion
func TestAddenda13FieldInclusionODFIIdentification(t *testing.T) {
	testAddenda13FieldInclusionODFIIdentification(t)
}

// BenchmarkAddenda13FieldInclusionODFIIdentification benchmarks validating ODFIIdentification fieldInclusion
func BenchmarkAddenda13FieldInclusionODFIIdentification(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda13FieldInclusionODFIIdentification(b)
	}
}

// testAddenda13FieldInclusionODFIBranchCountryCode validates ODFIBranchCountryCode fieldInclusion
func testAddenda13FieldInclusionODFIBranchCountryCode(t testing.TB) {
	addenda13 := mockAddenda13()
	addenda13.ODFIBranchCountryCode = ""
	if err := addenda13.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestAddenda13FieldInclusionODFIBranchCountryCode tests validating ODFIBranchCountryCode fieldInclusion
func TestAddenda13FieldInclusionODFIBranchCountryCode(t *testing.T) {
	testAddenda13FieldInclusionODFIBranchCountryCode(t)
}

// BenchmarkAddenda13FieldInclusionODFIBranchCountryCode benchmarks validating ODFIBranchCountryCode fieldInclusion
func BenchmarkAddenda13FieldInclusionODFIBranchCountryCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda13FieldInclusionODFIBranchCountryCode(b)
	}
}

// testAddenda13FieldInclusionEntryDetailSequenceNumber validates EntryDetailSequenceNumber fieldInclusion
func testAddenda13FieldInclusionEntryDetailSequenceNumber(t testing.TB) {
	addenda13 := mockAddenda13()
	addenda13.EntryDetailSequenceNumber = 0
	if err := addenda13.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestAddenda13FieldInclusionEntryDetailSequenceNumber tests validating
// EntryDetailSequenceNumber fieldInclusion
func TestAddenda13FieldInclusionEntryDetailSequenceNumber(t *testing.T) {
	testAddenda13FieldInclusionEntryDetailSequenceNumber(t)
}

// BenchmarkAddenda13FieldInclusionEntryDetailSequenceNumber benchmarks validating
// EntryDetailSequenceNumber fieldInclusion
func BenchmarkAddenda13FieldInclusionEntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda13FieldInclusionEntryDetailSequenceNumber(b)
	}
}

// TestAddenda13String validates that a known parsed Addenda13 record can be return to a string of the same value
func testAddenda13String(t testing.TB) {
	addenda13 := NewAddenda13()
	var line = "713Wells Fargo                        121042882                         US             0000001"
	addenda13.Parse(line)

	if addenda13.String() != line {
		t.Errorf("Strings do not match")
	}
	if addenda13.TypeCode != "13" {
		t.Errorf("TypeCode Expected 13 got: %v", addenda13.TypeCode)
	}
}

// TestAddenda13String tests validating that a known parsed Addenda13 record can be return to a string of the same value
func TestAddenda13String(t *testing.T) {
	testAddenda13String(t)
}

// BenchmarkAddenda13String benchmarks validating that a known parsed Addenda13 record can be return to a string of the same value
func BenchmarkAddenda13String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda13String(b)
	}
}
