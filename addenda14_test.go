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

// mockAddenda14 creates a mock Addenda14 record
func mockAddenda14() *Addenda14 {
	addenda14 := NewAddenda14()
	addenda14.RDFIName = "Citadel Bank"
	addenda14.RDFIIDNumberQualifier = "01"
	addenda14.RDFIIdentification = "231380104"
	addenda14.RDFIBranchCountryCode = "US"
	addenda14.EntryDetailSequenceNumber = 00000001
	return addenda14
}

// TestMockAddenda14 validates mockAddenda14
func TestMockAddenda14(t *testing.T) {
	addenda14 := mockAddenda14()
	if err := addenda14.Validate(); err != nil {
		t.Error("mockAddenda14 does not validate and will break other tests")
	}
}

// testAddenda14Parse parses Addenda14 record
func testAddenda14Parse(t testing.TB) {
	Addenda14 := NewAddenda14()
	line := "714Citadel Bank                       01231380104                         US           0000001"
	Addenda14.Parse(line)
	// walk the Addenda14 struct
	if Addenda14.TypeCode != "14" {
		t.Errorf("expected %v got %v", "14", Addenda14.TypeCode)
	}
	if Addenda14.RDFIName != "Citadel Bank" {
		t.Errorf("expected %v got %v", "Citadel Bank", Addenda14.RDFIName)
	}
	if Addenda14.RDFIIDNumberQualifier != "01" {
		t.Errorf("expected: %v got: %v", "01", Addenda14.RDFIIDNumberQualifier)
	}
	if Addenda14.RDFIIdentification != "231380104" {
		t.Errorf("expected: %v got: %v", "928383-23938", Addenda14.RDFIIdentification)
	}
	if Addenda14.RDFIBranchCountryCode != "US" {
		t.Errorf("expected: %s got: %s", "US", Addenda14.RDFIBranchCountryCode)
	}
	if Addenda14.EntryDetailSequenceNumber != 0000001 {
		t.Errorf("expected: %v got: %v", 0000001, Addenda14.EntryDetailSequenceNumber)
	}
}

// TestAddenda14Parse tests parsing Addenda14 record
func TestAddenda14Parse(t *testing.T) {
	testAddenda14Parse(t)
}

// BenchmarkAddenda14Parse benchmarks parsing Addenda14 record
func BenchmarkAddenda14Parse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda14Parse(b)
	}
}

// testAddenda14ValidTypeCode validates Addenda14 TypeCode
func testAddenda14ValidTypeCode(t testing.TB) {
	addenda14 := mockAddenda14()
	addenda14.TypeCode = "65"
	err := addenda14.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda14ValidTypeCode tests validating Addenda14 TypeCode
func TestAddenda14ValidTypeCode(t *testing.T) {
	testAddenda14ValidTypeCode(t)
}

// BenchmarkAddenda14ValidTypeCode benchmarks validating Addenda14 TypeCode
func BenchmarkAddenda14ValidTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda14ValidTypeCode(b)
	}
}

// testAddenda14TypeCode14 TypeCode is 14 if TypeCode is a valid TypeCode
func testAddenda14TypeCode14(t testing.TB) {
	addenda14 := mockAddenda14()
	addenda14.TypeCode = "05"
	err := addenda14.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda14TypeCode14 tests TypeCode is 14 if TypeCode is a valid TypeCode
func TestAddenda14TypeCode14(t *testing.T) {
	testAddenda14TypeCode14(t)
}

// BenchmarkAddenda14TypeCode14 benchmarks TypeCode is 14 if TypeCode is a valid TypeCode
func BenchmarkAddenda14TypeCode14(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda14TypeCode14(b)
	}
}

// testRDFINameAlphaNumeric validates RDFIName is alphanumeric
func testRDFINameAlphaNumeric(t testing.TB) {
	addenda14 := mockAddenda14()
	addenda14.RDFIName = "Wells®Fargo"
	err := addenda14.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestRDFINameAlphaNumeric tests validating RDFIName is alphanumeric
func TestRDFINameAlphaNumeric(t *testing.T) {
	testRDFINameAlphaNumeric(t)
}

// BenchmarkRDFINameAlphaNumeric benchmarks validating RDFIName is alphanumeric
func BenchmarkRDFINameAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testRDFINameAlphaNumeric(b)
	}
}

// testRDFIIDNumberQualifierValid validates RDFIIDNumberQualifier is valid
func testRDFIIDNumberQualifierValid(t testing.TB) {
	addenda14 := mockAddenda14()
	addenda14.RDFIIDNumberQualifier = "®1"
	err := addenda14.Validate()
	if !base.Match(err, ErrIDNumberQualifier) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestRDFIIDNumberQualifierValid tests validating RDFIIDNumberQualifier is valid
func TestRDFIIDNumberQualifierValid(t *testing.T) {
	testRDFIIDNumberQualifierValid(t)
}

// BenchmarkRDFIIDNumberQualifierValid benchmarks validating RDFIIDNumberQualifier is valid
func BenchmarkRDFIIDNumberQualifierValid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testRDFIIDNumberQualifierValid(b)
	}
}

// testRDFIIdentificationAlphaNumeric validates RDFIIdentification is alphanumeric
func testRDFIIdentificationAlphaNumeric(t testing.TB) {
	addenda14 := mockAddenda14()
	addenda14.RDFIIdentification = "®121042882"
	err := addenda14.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestRDFIIdentificationAlphaNumeric tests validating RDFIIdentification is alphanumeric
func TestRDFIIdentificationAlphaNumeric(t *testing.T) {
	testRDFIIdentificationAlphaNumeric(t)
}

// BenchmarkRDFIIdentificationAlphaNumeric benchmarks validating RDFIIdentification is alphanumeric
func BenchmarkRDFIIdentificationAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testRDFIIdentificationAlphaNumeric(b)
	}
}

// testRDFIBranchCountryCodeAlphaNumeric validates RDFIBranchCountryCode is alphanumeric
func testRDFIBranchCountryCodeAlphaNumeric(t testing.TB) {
	addenda14 := mockAddenda14()
	addenda14.RDFIBranchCountryCode = "U®"
	err := addenda14.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestRDFIBranchCountryCodeAlphaNumeric tests validating RDFIBranchCountryCode is alphanumeric
func TestRDFIBranchCountryCodeAlphaNumeric(t *testing.T) {
	testRDFIBranchCountryCodeAlphaNumeric(t)
}

// BenchmarkRDFIBranchCountryCodeAlphaNumeric benchmarks validating RDFIBranchCountryCode is alphanumeric
func BenchmarkRDFIBranchCountryCodeAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testRDFIBranchCountryCodeAlphaNumeric(b)
	}
}

// testAddenda14FieldInclusionTypeCode validates TypeCode fieldInclusion
func testAddenda14FieldInclusionTypeCode(t testing.TB) {
	addenda14 := mockAddenda14()
	addenda14.TypeCode = ""
	err := addenda14.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda14FieldInclusionTypeCode tests validating TypeCode fieldInclusion
func TestAddenda14FieldInclusionTypeCode(t *testing.T) {
	testAddenda14FieldInclusionTypeCode(t)
}

// BenchmarkAddenda14FieldInclusionTypeCode benchmarks validating TypeCode fieldInclusion
func BenchmarkAddenda14FieldInclusionTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda14FieldInclusionTypeCode(b)
	}
}

// testAddenda14FieldInclusionRDFIName validates RDFIName fieldInclusion
func testAddenda14FieldInclusionRDFIName(t testing.TB) {
	addenda14 := mockAddenda14()
	addenda14.RDFIName = ""
	err := addenda14.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda14FieldInclusionRDFIName tests validating RDFIName fieldInclusion
func TestAddenda14FieldInclusionRDFIName(t *testing.T) {
	testAddenda14FieldInclusionRDFIName(t)
}

// BenchmarkAddenda14FieldInclusionRDFIName benchmarks validating RDFIName fieldInclusion
func BenchmarkAddenda14FieldInclusionRDFIName(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda14FieldInclusionRDFIName(b)
	}
}

// testAddenda14FieldInclusionRDFIIDNumberQualifier validates RDFIIDNumberQualifier fieldInclusion
func testAddenda14FieldInclusionRDFIIDNumberQualifier(t testing.TB) {
	addenda14 := mockAddenda14()
	addenda14.RDFIIDNumberQualifier = ""
	err := addenda14.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda14FieldInclusionRDFIIdNumberQualifier tests validating RDFIIdNumberQualifier fieldInclusion
func TestAddenda14FieldInclusionRDFIIdNumberQualifier(t *testing.T) {
	testAddenda14FieldInclusionRDFIIDNumberQualifier(t)
}

// BenchmarkAddenda14FieldInclusionRDFIIdNumberQualifier benchmarks validating RDFIIdNumberQualifier fieldInclusion
func BenchmarkAddenda14FieldInclusionRDFIIdNumberQualifier(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda14FieldInclusionRDFIIDNumberQualifier(b)
	}
}

// testAddenda14FieldInclusionRDFIIdentification validates RDFIIdentification fieldInclusion
func testAddenda14FieldInclusionRDFIIdentification(t testing.TB) {
	addenda14 := mockAddenda14()
	addenda14.RDFIIdentification = ""
	err := addenda14.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda14FieldInclusionRDFIIdentification tests validating RDFIIdentification fieldInclusion
func TestAddenda14FieldInclusionRDFIIdentification(t *testing.T) {
	testAddenda14FieldInclusionRDFIIdentification(t)
}

// BenchmarkAddenda14FieldInclusionRDFIIdentification benchmarks validating RDFIIdentification fieldInclusion
func BenchmarkAddenda14FieldInclusionRDFIIdentification(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda14FieldInclusionRDFIIdentification(b)
	}
}

// testAddenda14FieldInclusionRDFIBranchCountryCode validates RDFIBranchCountryCode fieldInclusion
func testAddenda14FieldInclusionRDFIBranchCountryCode(t testing.TB) {
	addenda14 := mockAddenda14()
	addenda14.RDFIBranchCountryCode = ""
	err := addenda14.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda14FieldInclusionRDFIBranchCountryCode tests validating RDFIBranchCountryCode fieldInclusion
func TestAddenda14FieldInclusionRDFIBranchCountryCode(t *testing.T) {
	testAddenda14FieldInclusionRDFIBranchCountryCode(t)
}

// BenchmarkAddenda14FieldInclusionRDFIBranchCountryCode benchmarks validating RDFIBranchCountryCode fieldInclusion
func BenchmarkAddenda14FieldInclusionRDFIBranchCountryCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda14FieldInclusionRDFIBranchCountryCode(b)
	}
}

// testAddenda14FieldInclusionEntryDetailSequenceNumber validates EntryDetailSequenceNumber fieldInclusion
func testAddenda14FieldInclusionEntryDetailSequenceNumber(t testing.TB) {
	addenda14 := mockAddenda14()
	addenda14.EntryDetailSequenceNumber = 0
	err := addenda14.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda14FieldInclusionEntryDetailSequenceNumber tests validating
// EntryDetailSequenceNumber fieldInclusion
func TestAddenda14FieldInclusionEntryDetailSequenceNumber(t *testing.T) {
	testAddenda14FieldInclusionEntryDetailSequenceNumber(t)
}

// BenchmarkAddenda14FieldInclusionEntryDetailSequenceNumber benchmarks validating
// EntryDetailSequenceNumber fieldInclusion
func BenchmarkAddenda14FieldInclusionEntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda14FieldInclusionEntryDetailSequenceNumber(b)
	}
}

// TestAddenda14String validates that a known parsed Addenda14 record can be return to a string of the same value
func testAddenda14String(t testing.TB) {
	addenda14 := NewAddenda14()
	var line = "714Citadel Bank                       231380104                         US             0000001"

	addenda14.Parse(line)

	if addenda14.String() != line {
		t.Errorf("Strings do not match")
	}
	if addenda14.TypeCode != "14" {
		t.Errorf("TypeCode Expected 14 got: %v", addenda14.TypeCode)
	}
}

// TestAddenda14String tests validating that a known parsed Addenda14 record can be return to a string of the same value
func TestAddenda14String(t *testing.T) {
	testAddenda14String(t)
}

// BenchmarkAddenda14String benchmarks validating that a known parsed Addenda14 record can be return to a string of the same value
func BenchmarkAddenda14String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda14String(b)
	}
}

// TestAddenda14RuneCountInString validates RuneCountInString
func TestAddenda14RuneCountInString(t *testing.T) {
	addenda14 := NewAddenda14()
	var line = "714Citadel Bank                       231380104                         US"
	addenda14.Parse(line)

	if addenda14.RDFIBranchCountryCode != "" {
		t.Error("Parsed with an invalid RuneCountInString not equal to 94")
	}
}
