// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"

	"github.com/moov-io/base"
)

// mockAddenda12 creates a mock Addenda12 record
func mockAddenda12() *Addenda12 {
	addenda12 := NewAddenda12()
	addenda12.OriginatorCityStateProvince = "JacobsTown*PA\\"
	addenda12.OriginatorCountryPostalCode = "US*19305\\"
	addenda12.EntryDetailSequenceNumber = 00000001
	return addenda12
}

// TestMockAddenda12 validates mockAddenda12
func TestMockAddenda12(t *testing.T) {
	addenda12 := mockAddenda12()
	if err := addenda12.Validate(); err != nil {
		t.Error("mockAddenda12 does not validate and will break other tests")
	}
}

// testAddenda12Parse parses Addenda12 record
func testAddenda12Parse(t testing.TB) {
	Addenda12 := NewAddenda12()
	line := "712" + "JacobsTown*PA\\                     " + "US*19305\\                                        " + "0000001"
	Addenda12.Parse(line)
	// walk the Addenda12 struct
	if Addenda12.recordType != "7" {
		t.Errorf("expected %v got %v", "7", Addenda12.recordType)
	}
	if Addenda12.TypeCode != "12" {
		t.Errorf("expected %v got %v", "12", Addenda12.TypeCode)
	}
	if Addenda12.OriginatorCityStateProvince != "JacobsTown*PA\\" {
		t.Errorf("expected %v got %v", "JacobsTown*PA\\", Addenda12.OriginatorCityStateProvince)
	}
	if Addenda12.OriginatorCountryPostalCode != "US*19305\\" {
		t.Errorf("expected: %v got: %v", "US*19305\\", Addenda12.OriginatorCountryPostalCode)
	}
	if Addenda12.reserved != "              " {
		t.Errorf("expected: %v got: %v", "              ", Addenda12.reserved)
	}
	if Addenda12.EntryDetailSequenceNumber != 0000001 {
		t.Errorf("expected: %v got: %v", 0000001, Addenda12.EntryDetailSequenceNumber)
	}
}

// TestAddenda12Parse tests parsing Addenda12 record
func TestAddenda12Parse(t *testing.T) {
	testAddenda12Parse(t)
}

// BenchmarkAddenda12Parse benchmarks parsing Addenda12 record
func BenchmarkAddenda12Parse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda12Parse(b)
	}
}

// testAddenda12ValidRecordType validates Addenda12 recordType
func testAddenda12ValidRecordType(t testing.TB) {
	addenda12 := mockAddenda12()
	addenda12.recordType = "63"
	if err := addenda12.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddenda12ValidRecordType tests validating Addenda12 recordType
func TestAddenda12ValidRecordType(t *testing.T) {
	testAddenda12ValidRecordType(t)
}

// BenchmarkAddenda12ValidRecordType benchmarks validating Addenda12 recordType
func BenchmarkAddenda12ValidRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda12ValidRecordType(b)
	}
}

// testAddenda12ValidTypeCode validates Addenda12 TypeCode
func testAddenda12ValidTypeCode(t testing.TB) {
	addenda12 := mockAddenda12()
	addenda12.TypeCode = "65"
	if err := addenda12.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddenda12ValidTypeCode tests validating Addenda12 TypeCode
func TestAddenda12ValidTypeCode(t *testing.T) {
	testAddenda12ValidTypeCode(t)
}

// BenchmarkAddenda12ValidTypeCode benchmarks validating Addenda12 TypeCode
func BenchmarkAddenda12ValidTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda12ValidTypeCode(b)
	}
}

// testAddenda12TypeCode12 TypeCode is 12 if TypeCode is a valid TypeCode
func testAddenda12TypeCode12(t testing.TB) {
	addenda12 := mockAddenda12()
	addenda12.TypeCode = "05"
	if err := addenda12.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddenda12TypeCode12 tests TypeCode is 12 if TypeCode is a valid TypeCode
func TestAddenda12TypeCode12(t *testing.T) {
	testAddenda12TypeCode12(t)
}

// BenchmarkAddenda12TypeCode12 benchmarks TypeCode is 12 if TypeCode is a valid TypeCode
func BenchmarkAddenda12TypeCode12(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda12TypeCode12(b)
	}
}

// testOriginatorCityStateProvinceAlphaNumeric validates OriginatorCityStateProvince is alphanumeric
func testOriginatorCityStateProvinceAlphaNumeric(t testing.TB) {
	addenda12 := mockAddenda12()
	addenda12.OriginatorCityStateProvince = "Jacobs®Town*PA"
	if err := addenda12.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "OriginatorCityStateProvince" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestOriginatorCityStateProvinceAlphaNumeric tests validating OriginatorCityStateProvince is alphanumeric
func TestOriginatorCityStateProvinceAlphaNumeric(t *testing.T) {
	testOriginatorCityStateProvinceAlphaNumeric(t)
}

// BenchmarkOriginatorCityStateProvinceAlphaNumeric benchmarks validating OriginatorCityStateProvince is alphanumeric
func BenchmarkOriginatorCityStateProvinceAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testOriginatorCityStateProvinceAlphaNumeric(b)
	}
}

// testOriginatorCountryPostalCodeAlphaNumeric validates OriginatorCountryPostalCode is alphanumeric
func testOriginatorCountryPostalCodeAlphaNumeric(t testing.TB) {
	addenda12 := mockAddenda12()
	addenda12.OriginatorCountryPostalCode = "US19®305"
	if err := addenda12.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "OriginatorCountryPostalCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestOriginatorCountryPostalCodeAlphaNumeric tests validating OriginatorCountryPostalCode is alphanumeric
func TestOriginatorCountryPostalCodeAlphaNumeric(t *testing.T) {
	testOriginatorCountryPostalCodeAlphaNumeric(t)
}

// BenchmarkOriginatorCountryPostalCodeAlphaNumeric benchmarks validating OriginatorCountryPostalCode is alphanumeric
func BenchmarkOriginatorCountryPostalCodeAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testOriginatorCountryPostalCodeAlphaNumeric(b)
	}
}

// testAddenda12FieldInclusionRecordType validates recordType fieldInclusion
func testAddenda12FieldInclusionRecordType(t testing.TB) {
	addenda12 := mockAddenda12()
	addenda12.recordType = ""
	err := addenda12.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda12FieldInclusionRecordType tests validating recordType fieldInclusion
func TestAddenda12FieldInclusionRecordType(t *testing.T) {
	testAddenda12FieldInclusionRecordType(t)
}

// BenchmarkAddenda12FieldInclusionRecordType benchmarks validating recordType fieldInclusion
func BenchmarkAddenda12FieldInclusionRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda12FieldInclusionRecordType(b)
	}
}

// testAddenda12FieldInclusionTypeCode validates TypeCode fieldInclusion
func testAddenda12FieldInclusionTypeCode(t testing.TB) {
	addenda12 := mockAddenda12()
	addenda12.TypeCode = ""
	err := addenda12.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda12FieldInclusionTypeCode tests validating TypeCode fieldInclusion
func TestAddenda12FieldInclusionTypeCode(t *testing.T) {
	testAddenda12FieldInclusionTypeCode(t)
}

// BenchmarkAddenda12FieldInclusionTypeCode benchmarks validating TypeCode fieldInclusion
func BenchmarkAddenda12FieldInclusionTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda12FieldInclusionTypeCode(b)
	}
}

// testAddenda12FieldInclusionOriginatorCityStateProvince validates OriginatorCityStateProvince fieldInclusion
func testAddenda12FieldInclusionOriginatorCityStateProvince(t testing.TB) {
	addenda12 := mockAddenda12()
	addenda12.OriginatorCityStateProvince = ""
	err := addenda12.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda12FieldInclusionOriginatorCityStateProvince tests validating OriginatorCityStateProvince fieldInclusion
func TestAddenda12FieldInclusionOriginatorCityStateProvince(t *testing.T) {
	testAddenda12FieldInclusionOriginatorCityStateProvince(t)
}

// BenchmarkAddenda12FieldInclusionOriginatorCityStateProvince benchmarks validating OriginatorCityStateProvince fieldInclusion
func BenchmarkAddenda12FieldInclusionOriginatorCityStateProvince(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda12FieldInclusionOriginatorCityStateProvince(b)
	}
}

// testAddenda12FieldInclusionOriginatorCountryPostalCode validates OriginatorCountryPostalCode fieldInclusion
func testAddenda12FieldInclusionOriginatorCountryPostalCode(t testing.TB) {
	addenda12 := mockAddenda12()
	addenda12.OriginatorCountryPostalCode = ""
	err := addenda12.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda12FieldInclusionOriginatorCountryPostalCode tests validating OriginatorCountryPostalCode fieldInclusion
func TestAddenda12FieldInclusionOriginatorCountryPostalCode(t *testing.T) {
	testAddenda12FieldInclusionOriginatorCountryPostalCode(t)
}

// BenchmarkAddenda12FieldInclusionOriginatorCountryPostalCode benchmarks validating OriginatorCountryPostalCode fieldInclusion
func BenchmarkAddenda12FieldInclusionOriginatorCountryPostalCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda12FieldInclusionOriginatorCountryPostalCode(b)
	}
}

// testAddenda12FieldInclusionEntryDetailSequenceNumber validates EntryDetailSequenceNumber fieldInclusion
func testAddenda12FieldInclusionEntryDetailSequenceNumber(t testing.TB) {
	addenda12 := mockAddenda12()
	addenda12.EntryDetailSequenceNumber = 0
	err := addenda12.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda12FieldInclusionEntryDetailSequenceNumber tests validating
// EntryDetailSequenceNumber fieldInclusion
func TestAddenda12FieldInclusionEntryDetailSequenceNumber(t *testing.T) {
	testAddenda12FieldInclusionEntryDetailSequenceNumber(t)
}

// BenchmarkAddenda12FieldInclusionEntryDetailSequenceNumber benchmarks validating
// EntryDetailSequenceNumber fieldInclusion
func BenchmarkAddenda12FieldInclusionEntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda12FieldInclusionEntryDetailSequenceNumber(b)
	}
}

// TestAddenda12String validates that a known parsed Addenda12 record can be return to a string of the same value
func testAddenda12String(t testing.TB) {
	addenda12 := NewAddenda12()
	// Backslash logic
	var line = "712" +
		"JacobsTown*PA\\                     " +
		"US*19305\\                          " +
		"              " +
		"0000001"

	addenda12.Parse(line)

	if addenda12.String() != line {
		t.Errorf("Strings do not match")
	}
	if addenda12.TypeCode != "12" {
		t.Errorf("TypeCode Expected 12 got: %v", addenda12.TypeCode)
	}
}

// TestAddenda12String tests validating that a known parsed Addenda12 record can be return to a string of the same value
func TestAddenda12String(t *testing.T) {
	testAddenda12String(t)
}

// BenchmarkAddenda12String benchmarks validating that a known parsed Addenda12 record can be return to a string of the same value
func BenchmarkAddenda12String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda12String(b)
	}
}
