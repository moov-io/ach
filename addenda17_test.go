// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"

	"github.com/moov-io/base"
)

// mockAddenda17 creates a mock Addenda17 record
func mockAddenda17() *Addenda17 {
	addenda17 := NewAddenda17()
	addenda17.PaymentRelatedInformation = "This is an international payment"
	addenda17.SequenceNumber = 1
	addenda17.EntryDetailSequenceNumber = 0000001
	return addenda17
}

func mockAddenda17B() *Addenda17 {
	addenda17 := NewAddenda17()
	addenda17.PaymentRelatedInformation = "Transfer of money from one country to another"
	addenda17.SequenceNumber = 2
	addenda17.EntryDetailSequenceNumber = 0000001

	return addenda17
}

// testAddenda17Parse parses Addenda17 record
func testAddenda17Parse(t testing.TB) {
	Addenda17 := NewAddenda17()
	line := "717This is an international payment                                                00010000001"
	Addenda17.Parse(line)
	// walk the Addenda17 struct
	if Addenda17.recordType != "7" {
		t.Errorf("expected %v got %v", "7", Addenda17.recordType)
	}
	if Addenda17.TypeCode != "17" {
		t.Errorf("expected %v got %v", "17", Addenda17.TypeCode)
	}
	if Addenda17.PaymentRelatedInformation != "This is an international payment" {
		t.Errorf("expected %v got %v", "This is an international payment", Addenda17.PaymentRelatedInformation)
	}
	if Addenda17.SequenceNumber != 1 {
		t.Errorf("expected: %v got: %v", 1, Addenda17.SequenceNumber)
	}
	if Addenda17.EntryDetailSequenceNumber != 0000001 {
		t.Errorf("expected: %v got: %v", 0000001, Addenda17.EntryDetailSequenceNumber)
	}
}

// TestAddenda17Parse tests parsing Addenda17 record
func TestAddenda17Parse(t *testing.T) {
	testAddenda17Parse(t)
}

// BenchmarkAddenda17Parse benchmarks parsing Addenda17 record
func BenchmarkAddenda17Parse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda17Parse(b)
	}
}

// TestMockAddenda17 validates mockAddenda17
func TestMockAddenda17(t *testing.T) {
	addenda17 := mockAddenda17()
	if err := addenda17.Validate(); err != nil {
		t.Error("mockAddenda17 does not validate and will break other tests")
	}
	if addenda17.EntryDetailSequenceNumber != 0000001 {
		t.Error("EntryDetailSequenceNumber dependent default value has changed")
	}
}

// testAddenda17String validates that a known parsed file can be return to a string of the same value
func testAddenda17String(t testing.TB) {
	addenda17 := NewAddenda17()
	var line = "717IAT                                        DIEGO MAY                            00010000001"
	addenda17.Parse(line)

	if addenda17.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestAddenda17 String tests validating that a known parsed file can be return to a string of the same value
func TestAddenda17String(t *testing.T) {
	testAddenda17String(t)
}

// BenchmarkAddenda17 String benchmarks validating that a known parsed file can be return to a string of the same value
func BenchmarkAddenda17String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda17String(b)
	}
}

func TestValidateAddenda17RecordType(t *testing.T) {
	addenda17 := mockAddenda17()
	addenda17.recordType = "63"
	err := addenda17.Validate()
	if !base.Match(err, NewErrRecordType(7)) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda17FieldInclusionTypeCode(t *testing.T) {
	addenda17 := mockAddenda17()
	addenda17.TypeCode = ""
	err := addenda17.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda17FieldInclusion(t *testing.T) {
	addenda17 := mockAddenda17()
	addenda17.EntryDetailSequenceNumber = 0
	err := addenda17.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda17FieldInclusionSequenceNumber(t *testing.T) {
	addenda17 := mockAddenda17()
	addenda17.SequenceNumber = 0
	err := addenda17.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda17FieldInclusionRecordType(t *testing.T) {
	addenda17 := mockAddenda17()
	addenda17.recordType = ""
	err := addenda17.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

//testAddenda17PaymentRelatedInformationAlphaNumeric validates PaymentRelatedInformation is alphanumeric
func testAddenda17PaymentRelatedInformationAlphaNumeric(t testing.TB) {
	addenda17 := mockAddenda17()
	addenda17.PaymentRelatedInformation = "®©"
	err := addenda17.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda17PaymentRelatedInformationAlphaNumeric tests validating PaymentRelatedInformation is alphanumeric
func TestAddenda17PaymentRelatedInformationAlphaNumeric(t *testing.T) {
	testAddenda17PaymentRelatedInformationAlphaNumeric(t)

}

// BenchmarkAddenda17PaymentRelatedInformationAlphaNumeric benchmarks PaymentRelatedInformation is alphanumeric
func BenchmarkAddenda17PaymentRelatedInformationAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda17PaymentRelatedInformationAlphaNumeric(b)
	}
}

// testAddenda17ValidTypeCode validates Addenda17 TypeCode
func testAddenda17ValidTypeCode(t testing.TB) {
	addenda17 := mockAddenda17()
	addenda17.TypeCode = "65"
	err := addenda17.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda17ValidTypeCode tests validating Addenda17 TypeCode
func TestAddenda17ValidTypeCode(t *testing.T) {
	testAddenda17ValidTypeCode(t)
}

// BenchmarkAddenda17ValidTypeCode benchmarks validating Addenda17 TypeCode
func BenchmarkAddenda17ValidTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda17ValidTypeCode(b)
	}
}

// testAddenda17TypeCode17 TypeCode is 17 if TypeCode is a valid TypeCode
func testAddenda17TypeCode17(t testing.TB) {
	addenda17 := mockAddenda17()
	addenda17.TypeCode = "05"
	err := addenda17.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda17TypeCode17 tests TypeCode is 17 if TypeCode is a valid TypeCode
func TestAddenda17TypeCode17(t *testing.T) {
	testAddenda17TypeCode17(t)
}

// BenchmarkAddenda17TypeCode17 benchmarks TypeCode is 17 if TypeCode is a valid TypeCode
func BenchmarkAddenda17TypeCode17(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda17TypeCode17(b)
	}
}
