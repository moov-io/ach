// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"
)

func mockAddenda02() *Addenda02 {
	addenda02 := NewAddenda02()
	addenda02.ReferenceInformationOne = "REFONEA"
	addenda02.ReferenceInformationTwo = "REF"
	addenda02.TerminalIdentificationCode = "TERM02"
	addenda02.TransactionSerialNumber = "100049"
	addenda02.TransactionDate = "0612"
	addenda02.AuthorizationCodeOrExpireDate = "123456"
	addenda02.TerminalLocation = "Target Store 0049"
	addenda02.TerminalCity = "PHILADELPHIA"
	addenda02.TerminalState = "PA"
	addenda02.TraceNumber = 91012980000088
	return addenda02
}

func TestMockAddenda02(t *testing.T) {
	addenda02 := mockAddenda02()
	if err := addenda02.Validate(); err != nil {
		t.Error("mockAddenda02 does not validate and will break other tests")
	}
}

func testAddenda02ValidRecordType(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.recordType = "63"
	if err := addenda02.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}
func TestAddenda02ValidRecordType(t *testing.T) {
	testAddenda02ValidRecordType(t)
}

func BenchmarkAddenda02ValidRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02ValidRecordType(b)
	}
}

func testAddenda02ValidTypeCode(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.typeCode = "65"
	if err := addenda02.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}
func TestAddenda02ValidTypeCode(t *testing.T) {
	testAddenda02ValidTypeCode(t)
}

func BenchmarkAddenda02ValidTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02ValidTypeCode(b)
	}
}

func testAddenda02TypeCode02(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.typeCode = "05"
	if err := addenda02.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}
func TestAddenda02TypeCode02(t *testing.T) {
	testAddenda02TypeCode02(t)
}

func BenchmarkAddenda02TypeCode02(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TypeCode02(b)
	}
}

func testAddenda02RecordType(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.recordType = ""
	if err := addenda02.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestAddenda02RecordType(t *testing.T) {
	testAddenda02RecordType(t)
}

func BenchmarkAddenda02FieldInclusionRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02RecordType(b)
	}
}

func testAddenda02TypeCode(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.typeCode = ""
	if err := addenda02.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestAddenda02TypeCode(t *testing.T) {
	testAddenda02TypeCode(t)
}

func BenchmarkAddenda02TypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TypeCode(b)
	}
}

func testAddenda02TerminalIdentificationCode(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TerminalIdentificationCode = ""
	if err := addenda02.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestAddenda02TerminalIdentificationCode(t *testing.T) {
	testAddenda02TerminalIdentificationCode(t)
}

func BenchmarkAddenda02TerminalIdentificationCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TerminalIdentificationCode(b)
	}
}

func testAddenda02TransactionSerialNumber(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TransactionSerialNumber = ""
	if err := addenda02.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestAddenda02TransactionSerialNumber(t *testing.T) {
	testAddenda02TransactionSerialNumber(t)
}

func BenchmarkAddenda02TransactionSerialNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TransactionSerialNumber(b)
	}
}

func testAddenda02TransactionDate(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TransactionDate = ""
	if err := addenda02.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestAddenda02TransactionDate(t *testing.T) {
	testAddenda02TransactionDate(t)
}

func BenchmarkAddenda02TransactionDate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TransactionDate(b)
	}
}

func testAddenda02TerminalLocation(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TerminalLocation = ""
	if err := addenda02.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestAddenda02TerminalLocation(t *testing.T) {
	testAddenda02TerminalLocation(t)
}

func BenchmarkAddenda02TerminalLocation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TerminalLocation(b)
	}
}

func testAddenda02TerminalCity(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TerminalCity = ""
	if err := addenda02.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestAddenda02TerminalCity(t *testing.T) {
	testAddenda02TerminalCity(t)
}

func BenchmarkAddenda02TerminalCity(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TerminalCity(b)
	}
}

func testAddenda02TerminalState(t testing.TB) {
	addenda02 := mockAddenda02()
	addenda02.TerminalState = ""
	if err := addenda02.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestAddenda02TerminalState(t *testing.T) {
	testAddenda02TerminalState(t)
}

func BenchmarkAddenda02TerminalState(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02TerminalState(b)
	}
}

// TestAddenda02 String validates that a known parsed file can be return to a string of the same value
func testAddenda02String(t testing.TB) {
	addenda02 := NewAddenda02()
	var line = "702REFONEAREFTERM021000490612123456Target Store 0049          PHILADELPHIA   PA091012980000088"
	addenda02.Parse(line)
	if addenda02.String() != line {
		t.Errorf("Strings do not match")
	}
	if addenda02.TypeCode() != "02" {
		t.Errorf("TypeCode Expected 02 got: %v", addenda02.TypeCode())
	}
}

func TestAddenda02String(t *testing.T) {
	testAddenda02String(t)
}

func BenchmarkAddenda02String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda02String(b)
	}
}
