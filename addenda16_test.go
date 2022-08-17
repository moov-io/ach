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
	"testing"

	"github.com/moov-io/base"
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
	if Addenda16.TypeCode != "16" {
		t.Errorf("expected %v got %v", "16", Addenda16.TypeCode)
	}
	if Addenda16.ReceiverCityStateProvince != "LetterTown*AB\\" {
		t.Errorf("expected %v got %v", "LetterTown*AB\\", Addenda16.ReceiverCityStateProvince)
	}
	if Addenda16.ReceiverCountryPostalCode != "CA*80014\\" {
		t.Errorf("expected: %v got: %v", "CA*80014\\", Addenda16.ReceiverCountryPostalCode)
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

// testAddenda16ValidTypeCode validates Addenda16 TypeCode
func testAddenda16ValidTypeCode(t testing.TB) {
	addenda16 := mockAddenda16()
	addenda16.TypeCode = "65"
	err := addenda16.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
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

// testAddenda16TypeCode16 TypeCode is 16 if TypeCode is a valid TypeCode
func testAddenda16TypeCode16(t testing.TB) {
	addenda16 := mockAddenda16()
	addenda16.TypeCode = "05"
	err := addenda16.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda16TypeCode16 tests TypeCode is 16 if TypeCode is a valid TypeCode
func TestAddenda16TypeCode16(t *testing.T) {
	testAddenda16TypeCode16(t)
}

// BenchmarkAddenda16TypeCode16 benchmarks TypeCode is 16 if TypeCode is a valid TypeCode
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
	err := addenda16.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
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
	err := addenda16.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
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

// testAddenda16FieldInclusionTypeCode validates TypeCode fieldInclusion
func testAddenda16FieldInclusionTypeCode(t testing.TB) {
	addenda16 := mockAddenda16()
	addenda16.TypeCode = ""
	err := addenda16.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
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
	err := addenda16.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
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
	err := addenda16.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
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
	err := addenda16.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
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
	if addenda16.TypeCode != "16" {
		t.Errorf("TypeCode Expected 16 got: %v", addenda16.TypeCode)
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

// TestAddenda16RuneCountInString validates RuneCountInString
func TestAddenda16RuneCountInString(t *testing.T) {
	addenda16 := NewAddenda16()
	var line = "716"
	addenda16.Parse(line)

	if addenda16.ReceiverCityStateProvince != "" {
		t.Error("Parsed with an invalid RuneCountInString not equal to 94")
	}
}
