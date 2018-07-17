// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"
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
	addenda18.EntryDetailSequenceNumber = 0000002
	return addenda18
}

func mockAddenda18C() *Addenda18 {
	addenda18 := NewAddenda18()
	addenda18.ForeignCorrespondentBankName = "Bank of France"
	addenda18.ForeignCorrespondentBankIDNumberQualifier = "01"
	addenda18.ForeignCorrespondentBankIDNumber = "456456456987987"
	addenda18.ForeignCorrespondentBankBranchCountryCode = "FR"
	addenda18.SequenceNumber = 2
	addenda18.EntryDetailSequenceNumber = 0000003
	return addenda18
}

func mockAddenda18D() *Addenda18 {
	addenda18 := NewAddenda18()
	addenda18.ForeignCorrespondentBankName = "Bank of Turkey"
	addenda18.ForeignCorrespondentBankIDNumberQualifier = "01"
	addenda18.ForeignCorrespondentBankIDNumber = "12312345678910"
	addenda18.ForeignCorrespondentBankBranchCountryCode = "TR"
	addenda18.SequenceNumber = 2
	addenda18.EntryDetailSequenceNumber = 0000004
	return addenda18
}

func mockAddenda18E() *Addenda18 {
	addenda18 := NewAddenda18()
	addenda18.ForeignCorrespondentBankName = "Bank of United Kingdom"
	addenda18.ForeignCorrespondentBankIDNumberQualifier = "01"
	addenda18.ForeignCorrespondentBankIDNumber = "1234567890123456789012345678901234"
	addenda18.ForeignCorrespondentBankBranchCountryCode = "GB"
	addenda18.SequenceNumber = 2
	addenda18.EntryDetailSequenceNumber = 0000005
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

// ToDo: Add parse logic

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

func TestValidateAddenda18RecordType(t *testing.T) {
	addenda18 := mockAddenda18()
	addenda18.recordType = "63"
	if err := addenda18.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestAddenda18TypeCodeFieldInclusion(t *testing.T) {
	addenda18 := mockAddenda18()
	addenda18.typeCode = ""
	if err := addenda18.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestAddenda18FieldInclusion(t *testing.T) {
	addenda18 := mockAddenda18()
	addenda18.EntryDetailSequenceNumber = 0
	if err := addenda18.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EntryDetailSequenceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestAddenda18FieldInclusionRecordType(t *testing.T) {
	addenda18 := mockAddenda18()
	addenda18.recordType = ""
	if err := addenda18.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

//testAddenda18ForeignCorrespondentBankNameAlphaNumeric validates ForeignCorrespondentBankName is alphanumeric
func testAddenda18ForeignCorrespondentBankNameAlphaNumeric(t testing.TB) {
	addenda18 := mockAddenda18()
	addenda18.ForeignCorrespondentBankName = "®©"
	if err := addenda18.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ForeignCorrespondentBankName" {
				t.Errorf("%T: %s", err, err)
			}
		}
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

//testAddenda18ForeignCorrespondentBankIDQualifierAlphaNumeric validates ForeignCorrespondentBankIDNumberQualifier is alphanumeric
func testAddenda18ForeignCorrespondentBankIDQualifierAlphaNumeric(t testing.TB) {
	addenda18 := mockAddenda18()
	addenda18.ForeignCorrespondentBankIDNumberQualifier = "®©"
	if err := addenda18.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ForeignCorrespondentBankIDNumberQualifier" {
				t.Errorf("%T: %s", err, err)
			}
		}
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

//testAddenda18ForeignCorrespondentBankBranchCountryCodeAlphaNumeric validates ForeignCorrespondentBankBranchCountryCode is alphanumeric
func testAddenda18ForeignCorrespondentBankBranchCountryCodeAlphaNumeric(t testing.TB) {
	addenda18 := mockAddenda18()
	addenda18.ForeignCorrespondentBankBranchCountryCode = "®©"
	if err := addenda18.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ForeignCorrespondentBankBranchCountryCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
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


//testAddenda18ForeignCorrespondentBankIDNumberAlphaNumeric validates ForeignCorrespondentBankIDNumber is alphanumeric
func testAddenda18ForeignCorrespondentBankIDNumberAlphaNumeric(t testing.TB) {
	addenda18 := mockAddenda18()
	addenda18.ForeignCorrespondentBankIDNumber = "®©"
	if err := addenda18.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ForeignCorrespondentBankIDNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
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
	addenda18.typeCode = "65"
	if err := addenda18.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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

// testAddenda18TypeCode18 TypeCode is 18 if typeCode is a valid TypeCode
func testAddenda18TypeCode18(t testing.TB) {
	addenda18 := mockAddenda18()
	addenda18.typeCode = "05"
	if err := addenda18.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddenda18TypeCode18 tests TypeCode is 18 if typeCode is a valid TypeCode
func TestAddenda18TypeCode18(t *testing.T) {
	testAddenda18TypeCode18(t)
}

// BenchmarkAddenda18TypeCode18 benchmarks TypeCode is 18 if typeCode is a valid TypeCode
func BenchmarkAddenda18TypeCode18(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda18TypeCode18(b)
	}
}
