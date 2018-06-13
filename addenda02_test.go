// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"
	"strings"
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

func testParseAddenda02(t testing.TB) {
	addendaPOS := NewAddenda02()
	var line = "702REFONEAREFTERM021000490612123456Target Store 0049          PHILADELPHIA   PA091012980000088"
	addendaPOS.Parse(line)

	r := NewReader(strings.NewReader(line))

	//Add a new BatchPOS
	r.addCurrentBatch(NewBatchPOS(mockBatchPOSHeader()))

	//Add a POS EntryDetail
	entryDetail := mockPOSEntryDetail()

	//Add an addenda to the POS EntryDetail
	entryDetail.AddAddenda(addendaPOS)

	// add the POSentry detail to the batch
	r.currentBatch.AddEntry(entryDetail)

	record := r.currentBatch.GetEntries()[0].Addendum[0].(*Addenda02)

	if record.recordType != "7" {
		t.Errorf("RecordType Expected '7' got: %v", record.recordType)
	}
	if record.TypeCode() != "02" {
		t.Errorf("TypeCode Expected 02 got: %v", record.TypeCode())
	}
	if record.ReferenceInformationOne != "REFONEA" {
		t.Errorf("ReferenceInformationOne Expected 'REFONEA' got: %v", record.ReferenceInformationOneField())
	}
	if record.ReferenceInformationTwo != "REF" {
		t.Errorf("ReferenceInformationTwo Expected 'REF' got: %v", record.ReferenceInformationTwoField())
	}
	if record.TerminalIdentificationCode != "TERM02" {
		t.Errorf("TerminalIdentificationCode Expected 'TERM02' got: %v", record.TerminalIdentificationCodeField())
	}
	if record.TransactionSerialNumber != "100049" {
		t.Errorf("TransactionSerialNumber Expected '100049' got: %v", record.TransactionSerialNumberField())
	}
	if record.TransactionDate != "0612" {
		t.Errorf("TransactionDate Expected '0612' got: %v", record.TransactionDateField())
	}
	if record.AuthorizationCodeOrExpireDate != "123456" {
		t.Errorf("AuthorizationCodeOrExpireDate Expected '123456' got: %v", record.AuthorizationCodeOrExpireDateField())
	}
	if record.TerminalLocation != "Target Store 0049" {
		t.Errorf("TerminalLocation Expected 'Target Store 0049' got: %v", record.TerminalLocationField())
	}
	if record.TerminalCity != "PHILADELPHIA" {
		t.Errorf("TerminalCity Expected '123456' got: %v", record.TerminalCityField())
	}
	if record.TerminalState != "PA" {
		t.Errorf("TerminalState Expected '123456' got: %v", record.TerminalStateField())
	}
	if record.TraceNumber != 91012980000088 {
		t.Errorf("TraceNumber Expected '91012980000088' got: %v", record.TraceNumberField())
	}
}

func TestParseAddenda02(t *testing.T) {
	testParseAddenda02(t)
}

func BenchmarkParseAddenda02(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testParseAddenda02(b)
	}
}