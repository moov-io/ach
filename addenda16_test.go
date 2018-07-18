// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"
)

// mockAddenda16 creates a mock Addenda16 record
func mockAddenda16() *Addenda16 {
	addenda16 := NewAddenda16()
	addenda16.ReceiverCityStateProvince = "LetterTown*AB\\"
	addenda16.ReceiverCountryPostalCode = "CA*80014\\"
	addenda16.EntryDetailSequenceNumber = 00000001
	return addenda16
}

// TestMockAddenda16 validates mockAddenda16
func TestMockAddenda16(t *testing.T) {
	addenda16 := mockAddenda16()
	if err := addenda16.Validate(); err != nil {
		t.Error("mockAddenda16 does not validate and will break other tests")
	}
}

// testAddenda16Parse parses Addenda16 record
func testAddenda16Parse(t testing.TB) {
	Addenda16 := NewAddenda16()
	line := "716LetterTown*AB\\                     CA*80014\\                                        0000001"
	Addenda16.Parse(line)
	// walk the Addenda16 struct
	if Addenda16.recordType != "7" {
		t.Errorf("expected %v got %v", "7", Addenda16.recordType)
	}
	if Addenda16.typeCode != "16" {
		t.Errorf("expected %v got %v", "16", Addenda16.typeCode)
	}
	if Addenda16.ReceiverCityStateProvince != "LetterTown*AB\\" {
		t.Errorf("expected %v got %v", "LetterTown*AB\\", Addenda16.ReceiverCityStateProvince)
	}
	if Addenda16.ReceiverCountryPostalCode != "CA*80014\\" {
		t.Errorf("expected: %v got: %v", "CA*80014\\", Addenda16.ReceiverCountryPostalCode)
	}
	if Addenda16.reserved != "              " {
		t.Errorf("expected: %v got: %v", "              ", Addenda16.reserved)
	}
	if Addenda16.EntryDetailSequenceNumber != 0000001 {
		t.Errorf("expected: %v got: %v", 0000001, Addenda16.EntryDetailSequenceNumber)
	}
}

// TestAddenda16Parse tests parsing Addenda16 record
func TestAddenda16Parse(t *testing.T) {
	testAddenda16Parse(t)
}

// BenchmarkAddenda16Parse benchmarks parsing Addenda16 record
func BenchmarkAddenda16Parse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda16Parse(b)
	}
}

// testAddenda16ValidRecordType validates Addenda16 recordType
func testAddenda16ValidRecordType(t testing.TB) {
	addenda16 := mockAddenda16()
	addenda16.recordType = "63"
	if err := addenda16.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddenda16ValidRecordType tests validating Addenda16 recordType
func TestAddenda16ValidRecordType(t *testing.T) {
	testAddenda16ValidRecordType(t)
}

// BenchmarkAddenda16ValidRecordType benchmarks validating Addenda16 recordType
func BenchmarkAddenda16ValidRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda16ValidRecordType(b)
	}
}

// testAddenda16ValidTypeCode validates Addenda16 TypeCode
func testAddenda16ValidTypeCode(t testing.TB) {
	addenda16 := mockAddenda16()
	addenda16.typeCode = "65"
	if err := addenda16.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddenda16ValidTypeCode tests validating Addenda16 TypeCode
func TestAddenda16ValidTypeCode(t *testing.T) {
	testAddenda16ValidTypeCode(t)
}

// BenchmarkAddenda16ValidTypeCode benchmarks validating Addenda16 TypeCode
func BenchmarkAddenda16ValidTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda16ValidTypeCode(b)
	}
}

// testAddenda16TypeCode16 TypeCode is 16 if typeCode is a valid TypeCode
func testAddenda16TypeCode16(t testing.TB) {
	addenda16 := mockAddenda16()
	addenda16.typeCode = "05"
	if err := addenda16.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddenda16TypeCode16 tests TypeCode is 16 if typeCode is a valid TypeCode
func TestAddenda16TypeCode16(t *testing.T) {
	testAddenda16TypeCode16(t)
}

// BenchmarkAddenda16TypeCode16 benchmarks TypeCode is 16 if typeCode is a valid TypeCode
func BenchmarkAddenda16TypeCode16(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda16TypeCode16(b)
	}
}

// testReceiverCityStateProvinceAlphaNumeric validates ReceiverCityStateProvince is alphanumeric
func testReceiverCityStateProvinceAlphaNumeric(t testing.TB) {
	addenda16 := mockAddenda16()
	addenda16.ReceiverCityStateProvince = "Jacobs®Town*PA"
	if err := addenda16.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ReceiverCityStateProvince" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestReceiverCityStateProvinceAlphaNumeric tests validating ReceiverCityStateProvince is alphanumeric
func TestReceiverCityStateProvinceAlphaNumeric(t *testing.T) {
	testReceiverCityStateProvinceAlphaNumeric(t)
}

// BenchmarkReceiverCityStateProvinceAlphaNumeric benchmarks validating ReceiverCityStateProvince is alphanumeric
func BenchmarkReceiverCityStateProvinceAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testReceiverCityStateProvinceAlphaNumeric(b)
	}
}

// testReceiverCountryPostalCodeAlphaNumeric validates ReceiverCountryPostalCode is alphanumeric
func testReceiverCountryPostalCodeAlphaNumeric(t testing.TB) {
	addenda16 := mockAddenda16()
	addenda16.ReceiverCountryPostalCode = "US19®305"
	if err := addenda16.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ReceiverCountryPostalCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestReceiverCountryPostalCodeAlphaNumeric tests validating ReceiverCountryPostalCode is alphanumeric
func TestReceiverCountryPostalCodeAlphaNumeric(t *testing.T) {
	testReceiverCountryPostalCodeAlphaNumeric(t)
}

// BenchmarkReceiverCountryPostalCodeAlphaNumeric benchmarks validating ReceiverCountryPostalCode is alphanumeric
func BenchmarkReceiverCountryPostalCodeAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testReceiverCountryPostalCodeAlphaNumeric(b)
	}
}

// testAddenda16FieldInclusionRecordType validates recordType fieldInclusion
func testAddenda16FieldInclusionRecordType(t testing.TB) {
	addenda16 := mockAddenda16()
	addenda16.recordType = ""
	if err := addenda16.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestAddenda16FieldInclusionRecordType tests validating recordType fieldInclusion
func TestAddenda16FieldInclusionRecordType(t *testing.T) {
	testAddenda16FieldInclusionRecordType(t)
}

// BenchmarkAddenda16FieldInclusionRecordType benchmarks validating recordType fieldInclusion
func BenchmarkAddenda16FieldInclusionRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda16FieldInclusionRecordType(b)
	}
}

// testAddenda16FieldInclusionTypeCode validates TypeCode fieldInclusion
func testAddenda16FieldInclusionTypeCode(t testing.TB) {
	addenda16 := mockAddenda16()
	addenda16.typeCode = ""
	if err := addenda16.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestAddenda16FieldInclusionTypeCode tests validating TypeCode fieldInclusion
func TestAddenda16FieldInclusionTypeCode(t *testing.T) {
	testAddenda16FieldInclusionTypeCode(t)
}

// BenchmarkAddenda16FieldInclusionTypeCode benchmarks validating TypeCode fieldInclusion
func BenchmarkAddenda16FieldInclusionTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda16FieldInclusionTypeCode(b)
	}
}

// testAddenda16FieldInclusionReceiverCityStateProvince validates ReceiverCityStateProvince fieldInclusion
func testAddenda16FieldInclusionReceiverCityStateProvince(t testing.TB) {
	addenda16 := mockAddenda16()
	addenda16.ReceiverCityStateProvince = ""
	if err := addenda16.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestAddenda16FieldInclusionReceiverCityStateProvince tests validating ReceiverCityStateProvince fieldInclusion
func TestAddenda16FieldInclusionReceiverCityStateProvince(t *testing.T) {
	testAddenda16FieldInclusionReceiverCityStateProvince(t)
}

// BenchmarkAddenda16FieldInclusionReceiverCityStateProvince benchmarks validating ReceiverCityStateProvince fieldInclusion
func BenchmarkAddenda16FieldInclusionReceiverCityStateProvince(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda16FieldInclusionReceiverCityStateProvince(b)
	}
}

// testAddenda16FieldInclusionReceiverCountryPostalCode validates ReceiverCountryPostalCode fieldInclusion
func testAddenda16FieldInclusionReceiverCountryPostalCode(t testing.TB) {
	addenda16 := mockAddenda16()
	addenda16.ReceiverCountryPostalCode = ""
	if err := addenda16.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestAddenda16FieldInclusionReceiverCountryPostalCode tests validating ReceiverCountryPostalCode fieldInclusion
func TestAddenda16FieldInclusionReceiverCountryPostalCode(t *testing.T) {
	testAddenda16FieldInclusionReceiverCountryPostalCode(t)
}

// BenchmarkAddenda16FieldInclusionReceiverCountryPostalCode benchmarks validating ReceiverCountryPostalCode fieldInclusion
func BenchmarkAddenda16FieldInclusionReceiverCountryPostalCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda16FieldInclusionReceiverCountryPostalCode(b)
	}
}

// testAddenda16FieldInclusionEntryDetailSequenceNumber validates EntryDetailSequenceNumber fieldInclusion
func testAddenda16FieldInclusionEntryDetailSequenceNumber(t testing.TB) {
	addenda16 := mockAddenda16()
	addenda16.EntryDetailSequenceNumber = 0
	if err := addenda16.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestAddenda16FieldInclusionEntryDetailSequenceNumber tests validating
// EntryDetailSequenceNumber fieldInclusion
func TestAddenda16FieldInclusionEntryDetailSequenceNumber(t *testing.T) {
	testAddenda16FieldInclusionEntryDetailSequenceNumber(t)
}

// BenchmarkAddenda16FieldInclusionEntryDetailSequenceNumber benchmarks validating
// EntryDetailSequenceNumber fieldInclusion
func BenchmarkAddenda16FieldInclusionEntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda16FieldInclusionEntryDetailSequenceNumber(b)
	}
}

// ToDo  Add Parse test for individual fields

// TestAddenda16String validates that a known parsed Addenda16 record can be return to a string of the same value
func testAddenda16String(t testing.TB) {
	addenda16 := NewAddenda16()
	// Backslash logic
	var line = "716" +
		"LetterTown*AB\\                     " +
		"CA*80014\\                          " +
		"              " +
		"0000001"

	addenda16.Parse(line)

	if addenda16.String() != line {
		t.Errorf("Strings do not match")
	}
	if addenda16.TypeCode() != "16" {
		t.Errorf("TypeCode Expected 16 got: %v", addenda16.TypeCode())
	}
}

// TestAddenda16String tests validating that a known parsed Addenda16 record can be return to a string of the same value
func TestAddenda16String(t *testing.T) {
	testAddenda16String(t)
}

// BenchmarkAddenda16String benchmarks validating that a known parsed Addenda16 record can be return to a string of the same value
func BenchmarkAddenda16String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda16String(b)
	}
}
