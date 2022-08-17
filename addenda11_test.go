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

// mockAddenda11 creates a mock Addenda11 record
func mockAddenda11() *Addenda11 {
	addenda11 := NewAddenda11()
	addenda11.OriginatorName = "BEK Solutions"
	addenda11.OriginatorStreetAddress = "15 West Place Street"
	addenda11.EntryDetailSequenceNumber = 00000001
	return addenda11
}

// TestMockAddenda11 validates mockAddenda11
func TestMockAddenda11(t *testing.T) {
	addenda11 := mockAddenda11()
	if err := addenda11.Validate(); err != nil {
		t.Error("mockAddenda11 does not validate and will break other tests")
	}
}

// testAddenda11Parse parses Addenda11 record
func testAddenda11Parse(t testing.TB) {
	Addenda11 := NewAddenda11()
	line := "711BEK Solutions                      15 West Place Street                             0000001"
	Addenda11.Parse(line)
	// walk the Addenda11 struct
	if Addenda11.TypeCode != "11" {
		t.Errorf("expected %v got %v", "11", Addenda11.TypeCode)
	}
	if Addenda11.OriginatorName != "BEK Solutions" {
		t.Errorf("expected %v got %v", "BEK Solutions", Addenda11.OriginatorName)
	}
	if Addenda11.OriginatorStreetAddress != "15 West Place Street" {
		t.Errorf("expected: %v got: %v", "15 West Place Street", Addenda11.OriginatorStreetAddress)
	}
	if Addenda11.EntryDetailSequenceNumber != 0000001 {
		t.Errorf("expected: %v got: %v", 0000001, Addenda11.EntryDetailSequenceNumber)
	}
}

// TestAddenda11Parse tests parsing Addenda11 record
func TestAddenda11Parse(t *testing.T) {
	testAddenda11Parse(t)
}

// BenchmarkAddenda11Parse benchmarks parsing Addenda11 record
func BenchmarkAddenda11Parse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda11Parse(b)
	}
}

// testAddenda11ValidTypeCode validates Addenda11 TypeCode
func testAddenda11ValidTypeCode(t testing.TB) {
	addenda11 := mockAddenda11()
	addenda11.TypeCode = "65"
	err := addenda11.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda11ValidTypeCode tests validating Addenda11 TypeCode
func TestAddenda11ValidTypeCode(t *testing.T) {
	testAddenda11ValidTypeCode(t)
}

// BenchmarkAddenda11ValidTypeCode benchmarks validating Addenda11 TypeCode
func BenchmarkAddenda11ValidTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda11ValidTypeCode(b)
	}
}

// testAddenda11TypeCode11 TypeCode is 11 if TypeCode is a valid TypeCode
func testAddenda11TypeCode11(t testing.TB) {
	addenda11 := mockAddenda11()
	addenda11.TypeCode = "05"
	err := addenda11.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda11TypeCode11 tests TypeCode is 11 if TypeCode is a valid TypeCode
func TestAddenda11TypeCode11(t *testing.T) {
	testAddenda11TypeCode11(t)
}

// BenchmarkAddenda11TypeCode11 benchmarks TypeCode is 11 if TypeCode is a valid TypeCode
func BenchmarkAddenda11TypeCode11(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda11TypeCode11(b)
	}
}

// testOriginatorNameAlphaNumeric validates OriginatorName is alphanumeric
func testOriginatorNameAlphaNumeric(t testing.TB) {
	addenda11 := mockAddenda11()
	addenda11.OriginatorName = "BEK S®lutions"
	err := addenda11.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestOriginatorNameAlphaNumeric tests validating OriginatorName is alphanumeric
func TestOriginatorNameAlphaNumeric(t *testing.T) {
	testOriginatorNameAlphaNumeric(t)
}

// BenchmarkOriginatorNameAlphaNumeric benchmarks validating OriginatorName is alphanumeric
func BenchmarkOriginatorNameAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testOriginatorNameAlphaNumeric(b)
	}
}

// testOriginatorStreetAddressAlphaNumeric validates OriginatorStreetAddress is alphanumeric
func testOriginatorStreetAddressAlphaNumeric(t testing.TB) {
	addenda11 := mockAddenda11()
	addenda11.OriginatorStreetAddress = "15 W®st"
	err := addenda11.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestOriginatorStreetAddressAlphaNumeric tests validating OriginatorStreetAddress is alphanumeric
func TestOriginatorStreetAddressAlphaNumeric(t *testing.T) {
	testOriginatorStreetAddressAlphaNumeric(t)
}

// BenchmarkOriginatorStreetAddressAlphaNumeric benchmarks validating OriginatorStreetAddress is alphanumeric
func BenchmarkOriginatorStreetAddressAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testOriginatorStreetAddressAlphaNumeric(b)
	}
}

// testAddenda11FieldInclusionTypeCode validates TypeCode fieldInclusion
func testAddenda11FieldInclusionTypeCode(t testing.TB) {
	addenda11 := mockAddenda11()
	addenda11.TypeCode = ""
	err := addenda11.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda11FieldInclusionTypeCode tests validating TypeCode fieldInclusion
func TestAddenda11FieldInclusionTypeCode(t *testing.T) {
	testAddenda11FieldInclusionTypeCode(t)
}

// BenchmarkAddenda11FieldInclusionTypeCode benchmarks validating TypeCode fieldInclusion
func BenchmarkAddenda11FieldInclusionTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda11FieldInclusionTypeCode(b)
	}
}

// testAddenda11FieldInclusionOriginatorName validates OriginatorName fieldInclusion
func testAddenda11FieldInclusionOriginatorName(t testing.TB) {
	addenda11 := mockAddenda11()
	addenda11.OriginatorName = ""
	err := addenda11.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda11FieldInclusionOriginatorName tests validating OriginatorName fieldInclusion
func TestAddenda11FieldInclusionOriginatorName(t *testing.T) {
	testAddenda11FieldInclusionOriginatorName(t)
}

// BenchmarkAddenda11FieldInclusionOriginatorName benchmarks validating OriginatorName fieldInclusion
func BenchmarkAddenda11FieldInclusionOriginatorName(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda11FieldInclusionOriginatorName(b)
	}
}

// testAddenda11FieldInclusionOriginatorStreetAddress validates OriginatorStreetAddress fieldInclusion
func testAddenda11FieldInclusionOriginatorStreetAddress(t testing.TB) {
	addenda11 := mockAddenda11()
	addenda11.OriginatorStreetAddress = ""
	err := addenda11.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda11FieldInclusionOriginatorStreetAddress tests validating OriginatorStreetAddress fieldInclusion
func TestAddenda11FieldInclusionOriginatorStreetAddress(t *testing.T) {
	testAddenda11FieldInclusionOriginatorStreetAddress(t)
}

// BenchmarkAddenda11FieldInclusionOriginatorStreetAddress benchmarks validating OriginatorStreetAddress fieldInclusion
func BenchmarkAddenda11FieldInclusionOriginatorStreetAddress(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda11FieldInclusionOriginatorStreetAddress(b)
	}
}

// testAddenda11FieldInclusionEntryDetailSequenceNumber validates EntryDetailSequenceNumber fieldInclusion
func testAddenda11FieldInclusionEntryDetailSequenceNumber(t testing.TB) {
	addenda11 := mockAddenda11()
	addenda11.EntryDetailSequenceNumber = 0
	err := addenda11.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda11FieldInclusionEntryDetailSequenceNumber tests validating
// EntryDetailSequenceNumber fieldInclusion
func TestAddenda11FieldInclusionEntryDetailSequenceNumber(t *testing.T) {
	testAddenda11FieldInclusionEntryDetailSequenceNumber(t)
}

// BenchmarkAddenda11FieldInclusionEntryDetailSequenceNumber benchmarks validating
// EntryDetailSequenceNumber fieldInclusion
func BenchmarkAddenda11FieldInclusionEntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda11FieldInclusionEntryDetailSequenceNumber(b)
	}
}

// TestAddenda11String validates that a known parsed Addenda11 record can be return to a string of the same value
func testAddenda11String(t testing.TB) {
	addenda11 := NewAddenda11()
	var line = "711BEK Solutions                      15 West Place Street                             0000001"
	addenda11.Parse(line)

	if addenda11.String() != line {
		t.Errorf("Strings do not match")
	}
	if addenda11.TypeCode != "11" {
		t.Errorf("TypeCode Expected 11 got: %v", addenda11.TypeCode)
	}
}

// TestAddenda11String tests validating that a known parsed Addenda11 record can be return to a string of the same value
func TestAddenda11String(t *testing.T) {
	testAddenda11String(t)
}

// BenchmarkAddenda11String benchmarks validating that a known parsed Addenda11 record can be return to a string of the same value
func BenchmarkAddenda11String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda11String(b)
	}
}

// TestAddenda11RuneCountInString validates RuneCountInString
func TestAddenda11RuneCountInString(t *testing.T) {
	addenda11 := NewAddenda11()
	var line = "711BEK Solutions                      15 West Place Street"
	addenda11.Parse(line)

	if addenda11.OriginatorName != "" {
		t.Error("Parsed with an invalid RuneCountInString not equal to 94")
	}
}
