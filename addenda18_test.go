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

// mockAddenda18 creates a mock Addenda18 record
func mockAddenda18() *Addenda18 {
	addenda18 := NewAddenda18()
	addenda18.ForeignCorrespondentBankName = "Bank of Germany"
	addenda18.ForeignCorrespondentBankIDNumberQualifier = "01"
	addenda18.ForeignCorrespondentBankIDNumber = "987987987654654"
	addenda18.ForeignCorrespondentBankBranchCountryCode = "DE"
	addenda18.SequenceNumber = 1
	addenda18.EntryDetailSequenceNumber = 0000001
	return addenda18
}

func mockAddenda18B() *Addenda18 {
	addenda18 := NewAddenda18()
	addenda18.ForeignCorrespondentBankName = "Bank of Spain"
	addenda18.ForeignCorrespondentBankIDNumberQualifier = "01"
	addenda18.ForeignCorrespondentBankIDNumber = "987987987123123"
	addenda18.ForeignCorrespondentBankBranchCountryCode = "ES"
	addenda18.SequenceNumber = 2
	addenda18.EntryDetailSequenceNumber = 0000001
	return addenda18
}

func mockAddenda18C() *Addenda18 {
	addenda18 := NewAddenda18()
	addenda18.ForeignCorrespondentBankName = "Bank of France"
	addenda18.ForeignCorrespondentBankIDNumberQualifier = "01"
	addenda18.ForeignCorrespondentBankIDNumber = "456456456987987"
	addenda18.ForeignCorrespondentBankBranchCountryCode = "FR"
	addenda18.SequenceNumber = 3
	addenda18.EntryDetailSequenceNumber = 0000001
	return addenda18
}

func mockAddenda18D() *Addenda18 {
	addenda18 := NewAddenda18()
	addenda18.ForeignCorrespondentBankName = "Bank of Turkey"
	addenda18.ForeignCorrespondentBankIDNumberQualifier = "01"
	addenda18.ForeignCorrespondentBankIDNumber = "12312345678910"
	addenda18.ForeignCorrespondentBankBranchCountryCode = "TR"
	addenda18.SequenceNumber = 4
	addenda18.EntryDetailSequenceNumber = 0000001
	return addenda18
}

func mockAddenda18E() *Addenda18 {
	addenda18 := NewAddenda18()
	addenda18.ForeignCorrespondentBankName = "Bank of United Kingdom"
	addenda18.ForeignCorrespondentBankIDNumberQualifier = "01"
	addenda18.ForeignCorrespondentBankIDNumber = "1234567890123456789012345678901234"
	addenda18.ForeignCorrespondentBankBranchCountryCode = "GB"
	addenda18.SequenceNumber = 5
	addenda18.EntryDetailSequenceNumber = 0000001
	return addenda18
}

func mockAddenda18F() *Addenda18 {
	addenda18 := NewAddenda18()
	addenda18.ForeignCorrespondentBankName = "Bank of Antarctica"
	addenda18.ForeignCorrespondentBankIDNumberQualifier = "01"
	addenda18.ForeignCorrespondentBankIDNumber = "123456789012345678901"
	addenda18.ForeignCorrespondentBankBranchCountryCode = "AQ"
	addenda18.SequenceNumber = 6
	addenda18.EntryDetailSequenceNumber = 0000001
	return addenda18
}

// TestMockAddenda18 validates mockAddenda18
func TestMockAddenda18(t *testing.T) {
	addenda18 := mockAddenda18()
	if err := addenda18.Validate(); err != nil {
		t.Error("mockAddenda18 does not validate and will break other tests")
	}
	if addenda18.ForeignCorrespondentBankName != "Bank of Germany" {
		t.Error("ForeignCorrespondentBankName dependent default value has changed")
	}
	if addenda18.ForeignCorrespondentBankIDNumberQualifier != "01" {
		t.Error("ForeignCorrespondentBankIDNumberQualifier dependent default value has changed")
	}
	if addenda18.ForeignCorrespondentBankIDNumber != "987987987654654" {
		t.Error("ForeignCorrespondentBankIDNumber dependent default value has changed")
	}
	if addenda18.ForeignCorrespondentBankBranchCountryCode != "DE" {
		t.Error("ForeignCorrespondentBankBranchCountryCode dependent default value has changed")
	}
	if addenda18.EntryDetailSequenceNumber != 0000001 {
		t.Error("EntryDetailSequenceNumber dependent default value has changed")
	}
}

// testAddenda18Parse parses Addenda18 record
func testAddenda18Parse(t testing.TB) {
	Addenda18 := NewAddenda18()
	line := "718Bank of Germany                    01987987987654654                   DE       00010000001"
	Addenda18.Parse(line)
	// walk the Addenda18 struct
	if Addenda18.TypeCode != "18" {
		t.Errorf("expected %v got %v", "18", Addenda18.TypeCode)
	}
	if Addenda18.ForeignCorrespondentBankName != "Bank of Germany" {
		t.Errorf("expected %v got %v", "Bank of Germany", Addenda18.ForeignCorrespondentBankName)
	}
	if Addenda18.ForeignCorrespondentBankIDNumberQualifier != "01" {
		t.Errorf("expected: %v got: %v", "01", Addenda18.ForeignCorrespondentBankIDNumberQualifier)
	}
	if Addenda18.ForeignCorrespondentBankIDNumber != "987987987654654" {
		t.Errorf("expected: %v got: %v", "987987987654654", Addenda18.ForeignCorrespondentBankIDNumber)
	}
	if Addenda18.ForeignCorrespondentBankBranchCountryCode != "DE" {
		t.Errorf("expected: %s got: %s", "DE", Addenda18.ForeignCorrespondentBankBranchCountryCode)
	}
	if Addenda18.EntryDetailSequenceNumber != 0000001 {
		t.Errorf("expected: %v got: %v", 0000001, Addenda18.EntryDetailSequenceNumber)
	}
}

// TestAddenda18Parse tests parsing Addenda18 record
func TestAddenda18Parse(t *testing.T) {
	testAddenda18Parse(t)
}

// BenchmarkAddenda18Parse benchmarks parsing Addenda18 record
func BenchmarkAddenda18Parse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda18Parse(b)
	}
}

// testAddenda18String validates that a known parsed file can be return to a string of the same value
func testAddenda18String(t testing.TB) {
	addenda18 := NewAddenda18()
	var line = "718Bank of United Kingdom             011234567890123456789012345678901234GB       00010000001"
	addenda18.Parse(line)

	if addenda18.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestAddenda18 String tests validating that a known parsed file can be return to a string of the same value
func TestAddenda18String(t *testing.T) {
	testAddenda18String(t)
}

// BenchmarkAddenda18 String benchmarks validating that a known parsed file can be return to a string of the same value
func BenchmarkAddenda18String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda18String(b)
	}
}

func TestAddenda18FieldInclusionTypeCode(t *testing.T) {
	addenda18 := mockAddenda18()
	addenda18.TypeCode = ""
	err := addenda18.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda18FieldInclusion(t *testing.T) {
	addenda18 := mockAddenda18()
	addenda18.EntryDetailSequenceNumber = 0
	err := addenda18.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda18FieldInclusionSequenceNumber(t *testing.T) {
	addenda18 := mockAddenda18()
	addenda18.SequenceNumber = 0
	err := addenda18.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda18FieldInclusionFCBankName(t *testing.T) {
	addenda18 := mockAddenda18()
	addenda18.ForeignCorrespondentBankName = ""
	err := addenda18.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda18FieldInclusionFCBankIDNumberQualifier(t *testing.T) {
	addenda18 := mockAddenda18()
	addenda18.ForeignCorrespondentBankIDNumberQualifier = ""
	err := addenda18.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda18FieldInclusionFCBankIDNumber(t *testing.T) {
	addenda18 := mockAddenda18()
	addenda18.ForeignCorrespondentBankIDNumber = ""
	err := addenda18.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda18FieldInclusionFCBankBranchCountryCode(t *testing.T) {
	addenda18 := mockAddenda18()
	addenda18.ForeignCorrespondentBankBranchCountryCode = ""
	err := addenda18.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// testAddenda18ForeignCorrespondentBankNameAlphaNumeric validates ForeignCorrespondentBankName is alphanumeric
func testAddenda18ForeignCorrespondentBankNameAlphaNumeric(t testing.TB) {
	addenda18 := mockAddenda18()
	addenda18.ForeignCorrespondentBankName = "®©"
	err := addenda18.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda18ForeignCorrespondentBankNameAlphaNumeric tests validating ForeignCorrespondentBankName is alphanumeric
func TestAddenda18ForeignCorrespondentBankNameAlphaNumeric(t *testing.T) {
	testAddenda18ForeignCorrespondentBankNameAlphaNumeric(t)

}

// BenchmarkAddenda18ForeignCorrespondentBankNameAlphaNumeric benchmarks ForeignCorrespondentBankName is alphanumeric
func BenchmarkAddenda18ForeignCorrespondentBankNameAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda18ForeignCorrespondentBankNameAlphaNumeric(b)
	}
}

// testAddenda18ForeignCorrespondentBankIDQualifierAlphaNumeric validates ForeignCorrespondentBankIDNumberQualifier is alphanumeric
func testAddenda18ForeignCorrespondentBankIDQualifierAlphaNumeric(t testing.TB) {
	addenda18 := mockAddenda18()
	addenda18.ForeignCorrespondentBankIDNumberQualifier = "®©"
	err := addenda18.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda18ForeignCorrespondentBankIDQualifierAlphaNumeric tests validating ForeignCorrespondentBankIDNumberQualifier is alphanumeric
func TestAddenda18ForeignCorrespondentBankIDQualifierAlphaNumeric(t *testing.T) {
	testAddenda18ForeignCorrespondentBankIDQualifierAlphaNumeric(t)
}

// BenchmarkAddenda18ForeignCorrespondentBankIDQualifierAlphaNumeric benchmarks ForeignCorrespondentBankIDNumberQualifier is alphanumeric
func BenchmarkAddenda18ForeignCorrespondentBankIDQualifierAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda18ForeignCorrespondentBankIDQualifierAlphaNumeric(b)
	}
}

// testAddenda18ForeignCorrespondentBankBranchCountryCodeAlphaNumeric validates ForeignCorrespondentBankBranchCountryCode is alphanumeric
func testAddenda18ForeignCorrespondentBankBranchCountryCodeAlphaNumeric(t testing.TB) {
	addenda18 := mockAddenda18()
	addenda18.ForeignCorrespondentBankBranchCountryCode = "®©"
	err := addenda18.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda18ForeignCorrespondentBankBranchCountryCodeNumeric tests validating ForeignCorrespondentBankBranchCountryCode is alphanumeric
func TestAddenda18ForeignCorrespondentBankBranchCountryCodeAlphaNumeric(t *testing.T) {
	testAddenda18ForeignCorrespondentBankBranchCountryCodeAlphaNumeric(t)
}

// BenchmarkAddenda18ForeignCorrespondentBankBranchCountryCodeAlphaNumeric benchmarks ForeignCorrespondentBankBranchCountryCode is alphanumeric
func BenchmarkAddenda18ForeignCorrespondentBankBranchCountryCodeAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda18ForeignCorrespondentBankBranchCountryCodeAlphaNumeric(b)
	}
}

// testAddenda18ForeignCorrespondentBankIDNumberAlphaNumeric validates ForeignCorrespondentBankIDNumber is alphanumeric
func testAddenda18ForeignCorrespondentBankIDNumberAlphaNumeric(t testing.TB) {
	addenda18 := mockAddenda18()
	addenda18.ForeignCorrespondentBankIDNumber = "®©"
	err := addenda18.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda18ForeignCorrespondentBankIDNumberAlphaNumeric tests validating ForeignCorrespondentBankIDNumber is alphanumeric
func TestAddenda18ForeignCorrespondentBankIDNumberAlphaNumeric(t *testing.T) {
	testAddenda18ForeignCorrespondentBankIDNumberAlphaNumeric(t)
}

// BenchmarkAddenda18ForeignCorrespondentBankIDNumberAlphaNumeric benchmarks ForeignCorrespondentBankIDNumber is alphanumeric
func BenchmarkAddendaForeignCorrespondentBankIDNumberAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda18ForeignCorrespondentBankIDNumberAlphaNumeric(b)
	}
}

// testAddenda18ValidTypeCode validates Addenda18 TypeCode
func testAddenda18ValidTypeCode(t testing.TB) {
	addenda18 := mockAddenda18()
	addenda18.TypeCode = "65"
	err := addenda18.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda18ValidTypeCode tests validating Addenda18 TypeCode
func TestAddenda18ValidTypeCode(t *testing.T) {
	testAddenda18ValidTypeCode(t)
}

// BenchmarkAddenda18ValidTypeCode benchmarks validating Addenda18 TypeCode
func BenchmarkAddenda18ValidTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda18ValidTypeCode(b)
	}
}

// testAddenda18TypeCode18 TypeCode is 18 if TypeCode is a valid TypeCode
func testAddenda18TypeCode18(t testing.TB) {
	addenda18 := mockAddenda18()
	addenda18.TypeCode = "05"
	err := addenda18.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda18TypeCode18 tests TypeCode is 18 if TypeCode is a valid TypeCode
func TestAddenda18TypeCode18(t *testing.T) {
	testAddenda18TypeCode18(t)
}

// BenchmarkAddenda18TypeCode18 benchmarks TypeCode is 18 if TypeCode is a valid TypeCode
func BenchmarkAddenda18TypeCode18(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda18TypeCode18(b)
	}
}

// TestAddenda18RuneCountInString validates RuneCountInString
func TestAddenda18RuneCountInString(t *testing.T) {
	addenda18 := NewAddenda18()
	var line = "718Bank of United Kingdom           "
	addenda18.Parse(line)

	if addenda18.ForeignCorrespondentBankBranchCountryCode != "" {
		t.Error("Parsed with an invalid RuneCountInString not equal to 94")
	}
}
