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
	"strconv"
	"strings"
	"testing"

	"github.com/moov-io/base"
)

// mockIATEntryDetail creates an IAT EntryDetail
func mockIATEntryDetail() *IATEntryDetail {
	entry := NewIATEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("121042882")
	entry.AddendaRecords = 007
	entry.DFIAccountNumber = "123456789"
	entry.Amount = 100000 // 1000.00
	entry.SetTraceNumber(mockIATBatchHeaderFF().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

// mockIATEntryDetail2 creates an EntryDetail
func mockIATEntryDetail2() *IATEntryDetail {
	entry := NewIATEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("121042882")
	entry.AddendaRecords = 007
	entry.DFIAccountNumber = "123456789"
	entry.Amount = 200000 // 2000.00
	entry.SetTraceNumber(mockIATBatchHeaderFF().ODFIIdentification, 2)
	entry.Category = CategoryForward
	return entry
}

// testMockIATEntryDetail validates an IATEntryDetail record
func testMockIATEntryDetail(t testing.TB) {
	entry := mockIATEntryDetail()
	if err := entry.Validate(); err != nil {
		t.Error("mockEntryDetail does not validate and will break other tests")
	}
	if entry.TransactionCode != CheckingCredit {
		t.Error("TransactionCode dependent default value has changed")
	}
	if entry.RDFIIdentification != "12104288" {
		t.Error("RDFIIdentification dependent default value has changed")
	}
	if entry.AddendaRecords != 7 {
		t.Error("AddendaRecords default dependent value has changed")
	}
	if entry.DFIAccountNumber != "123456789" {
		t.Error("DFIAccountNumber dependent default value has changed")
	}
	if entry.Amount != 100000 {
		t.Error("Amount dependent default value has changed")
	}
	if entry.TraceNumber != "231380100000001" {
		t.Errorf("TraceNumber dependent default value has changed %v", entry.TraceNumber)
	}
}

// TestMockIATEntryDetail tests validating an IATEntryDetail record
func TestMockIATEntryDetail(t *testing.T) {
	testMockIATEntryDetail(t)
}

// BenchmarkMockIATEntryDetail benchmarks validating an IATEntryDetail record
func BenchmarkIATMockEntryDetail(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockIATEntryDetail(b)
	}
}

// testParseIATEntryDetail parses a known IATEntryDetail record string.
func testParseIATEntryDetail(t testing.TB) {
	var line = "6221210428820007             000010000012345678901234567890123456789012345    1231380100000001"
	r := NewReader(strings.NewReader(line))
	r.addIATCurrentBatch(NewIATBatch(mockIATBatchHeaderFF()))
	r.IATCurrentBatch.SetHeader(mockIATBatchHeaderFF())
	r.line = line
	if err := r.parseIATEntryDetail(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.IATCurrentBatch.GetEntries()[0]

	if record.TransactionCode != CheckingCredit {
		t.Errorf("TransactionCode Expected '22' got: %v", record.TransactionCode)
	}
	if record.RDFIIdentificationField() != "12104288" {
		t.Errorf("RDFIIdentification Expected '12104288' got: '%v'", record.RDFIIdentificationField())
	}

	if record.AddendaRecordsField() != "0007" {
		t.Errorf("addendaRecords Expected '0007' got: %v", record.AddendaRecords)
	}
	if record.CheckDigit != "2" {
		t.Errorf("CheckDigit Expected '2' got: %v", record.CheckDigit)
	}
	if record.AmountField() != "0000100000" {
		t.Errorf("Amount Expected '0000100000' got: %v", record.AmountField())
	}
	if record.DFIAccountNumberField() != "12345678901234567890123456789012345" {
		t.Errorf("DfiAccountNumber Expected '12345678901234567890123456789012345' got: %v", record.DFIAccountNumberField())
	}
	if record.AddendaRecordIndicator != 1 {
		t.Errorf("AddendaRecordIndicator Expected '0' got: %v", record.AddendaRecordIndicator)
	}
	if record.TraceNumberField() != "231380100000001" {
		t.Errorf("TraceNumber Expected '231380100000001' got: %v", record.TraceNumberField())
	}
}

// TestParseIATEntryDetail tests parsing a known IATEntryDetail record string.
func TestParseIATEntryDetail(t *testing.T) {
	testParseIATEntryDetail(t)
}

// BenchmarkParseIATEntryDetail benchmarks parsing a known IATEntryDetail record string.
func BenchmarkParseIATEntryDetail(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testParseIATEntryDetail(b)
	}
}

// testIATEDString validates that a known parsed entry
// detail can be returned to a string of the same value
func testIATEDString(t testing.TB) {
	var line = "6221210428820007             000010000012345678901234567890123456789012345    1231380100000001"
	r := NewReader(strings.NewReader(line))
	r.addIATCurrentBatch(NewIATBatch(mockIATBatchHeaderFF()))
	r.IATCurrentBatch.SetHeader(mockIATBatchHeaderFF())
	r.line = line
	if err := r.parseIATEntryDetail(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.IATCurrentBatch.GetEntries()[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestIATEDString tests validating that a known parsed entry
// detail can be returned to a string of the same value
func TestIATEDString(t *testing.T) {
	testIATEDString(t)
}

// BenchmarkIATEDString benchmarks validating that a known parsed entry
// detail can be returned to a string of the same value
func BenchmarkIATEDString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATEDString(b)
	}
}

// testIATEDInvalidTransactionCode validates error for IATEntryDetail invalid TransactionCode
func testIATEDInvalidTransactionCode(t testing.TB) {
	iatEd := mockIATEntryDetail()
	iatEd.TransactionCode = 77
	err := iatEd.Validate()
	if !base.Match(err, ErrTransactionCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATEDInvalidTransactionCode  tests validating error for IATEntryDetail invalid TransactionCode
func TestIATEDInvalidTransactionCode(t *testing.T) {
	testIATEDInvalidTransactionCode(t)
}

// BenchmarkIATEDTransactionCode benchmarks validating error for IATEntryDetail invalid TransactionCode
func BenchmarkIATEDInvalidTransactionCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATEDInvalidTransactionCode(b)
	}
}

// testEDIATDFIAccountNumberAlphaNumeric validates company identification is alphanumeric
func testEDIATDFIAccountNumberAlphaNumeric(t testing.TB) {
	ed := mockIATEntryDetail()
	ed.DFIAccountNumber = "®"
	err := ed.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestEDIATDFIAccountNumberAlphaNumeric tests validating company identification is alphanumeric
func TestEDIATDFIAccountNumberAlphaNumeric(t *testing.T) {
	testEDIATDFIAccountNumberAlphaNumeric(t)
}

// BenchmarkEDIATDFIAccountNumberAlphaNumeric benchmarks validating company identification is alphanumeric
func BenchmarkEDIATDFIAccountNumberAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDIATDFIAccountNumberAlphaNumeric(b)
	}
}

// testEDIATisCheckDigit validates check digit
func testEDIATisCheckDigit(t testing.TB) {
	ed := mockIATEntryDetail()
	ed.CheckDigit = "1"
	err := ed.Validate()
	if !base.Match(err, NewErrValidCheckDigit(7)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestEDIATisCheckDigit tests validating check digit
func TestEDIATisCheckDigit(t *testing.T) {
	testEDIATisCheckDigit(t)
}

// BenchmarkEDIATisCheckDigit benchmarks validating check digit
func BenchmarkEDIATisCheckDigit(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDIATisCheckDigit(b)
	}
}

// testEDIATSetRDFII validates setting RDFI
func testEDIATSetRDFI(t testing.TB) {
	ed := NewIATEntryDetail()
	ed.SetRDFI("810866774")
	if ed.RDFIIdentification != "81086677" {
		t.Error("RDFI identification")
	}
	if ed.CheckDigit != "4" {
		t.Error("Unexpected check digit")
	}
}

// TestEDIATSetRDFI  tests validating setting RDFI
func TestEDIATSetRDFI(t *testing.T) {
	testEDIATSetRDFI(t)
}

// BenchmarkEDIATSetRDFI benchmarks validating setting RDFI
func BenchmarkEDIATSetRDFI(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDIATSetRDFI(b)
	}
}

// testValidateEDIATCheckDigit validates CheckDigit error
func testValidateEDIATCheckDigit(t testing.TB) {
	ed := mockIATEntryDetail()
	ed.CheckDigit = "XYZ"
	err := ed.Validate()
	if !base.Match(err, &strconv.NumError{}) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestValidateEDIATCheckDigit tests validating validates CheckDigit error
func TestValidateEDIATCheckDigit(t *testing.T) {
	testValidateEDIATCheckDigit(t)
}

// BenchmarkValidateEDIATCheckDigit benchmarks validating CheckDigit error
func BenchmarkValidateEDIATCheckDigit(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateEDIATCheckDigit(b)
	}
}

// testIATEDTransactionCode validates IATEntryDetail TransactionCode fieldInclusion
func testIATEDTransactionCode(t testing.TB) {
	iatEd := mockIATEntryDetail()
	iatEd.TransactionCode = 0
	err := iatEd.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATEDTransactionCode tests validating IATEntryDetail TransactionCode fieldInclusion
func TestIATEDTransactionCode(t *testing.T) {
	testIATEDTransactionCode(t)
}

// BenchmarkIATEDTransactionCode benchmarks validating IATEntryDetail TransactionCode fieldInclusion
func BenchmarkIATEDTransactionCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATEDTransactionCode(b)
	}
}

// testIATEDRDFIIdentification validates IATEntryDetail RDFIIdentification fieldInclusion
func testIATEDRDFIIdentification(t testing.TB) {
	iatEd := mockIATEntryDetail()
	iatEd.RDFIIdentification = ""
	err := iatEd.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATEDRDFIIdentification tests validating IATEntryDetail RDFIIdentification fieldInclusion
func TestIATEDRDFIIdentification(t *testing.T) {
	testIATEDRDFIIdentification(t)
}

// BenchmarkIATEDRDFIIdentification benchmarks validating IATEntryDetail RDFIIdentification fieldInclusion
func BenchmarkIATEDRDFIIdentification(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATEDRDFIIdentification(b)
	}
}

// testIATEDAddendaRecords validates IATEntryDetail AddendaRecords fieldInclusion
func testIATEDAddendaRecords(t testing.TB) {
	iatEd := mockIATEntryDetail()
	iatEd.AddendaRecords = 0
	err := iatEd.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATEDAddendaRecords tests validating IATEntryDetail AddendaRecords fieldInclusion
func TestIATEDAddendaRecords(t *testing.T) {
	testIATEDAddendaRecords(t)
}

// BenchmarkIATEDAddendaRecords benchmarks validating IATEntryDetail AddendaRecords fieldInclusion
func BenchmarkIATEDAddendaRecords(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATEDAddendaRecords(b)
	}
}

// testIATEDDFIAccountNumber validates IATEntryDetail DFIAccountNumber fieldInclusion
func testIATEDDFIAccountNumber(t testing.TB) {
	iatEd := mockIATEntryDetail()
	iatEd.DFIAccountNumber = ""
	err := iatEd.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATEDDFIAccountNumber tests validating IATEntryDetail DFIAccountNumber fieldInclusion
func TestIATEDDFIAccountNumber(t *testing.T) {
	testIATEDDFIAccountNumber(t)
}

// BenchmarkIATEDDFIAccountNumber benchmarks validating IATEntryDetail DFIAccountNumber fieldInclusion
func BenchmarkIATEDDFIAccountNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATEDDFIAccountNumber(b)
	}
}

// testIATEDTraceNumber validates IATEntryDetail TraceNumber fieldInclusion
func testIATEDTraceNumber(t testing.TB) {
	iatEd := mockIATEntryDetail()
	iatEd.TraceNumber = "0"
	err := iatEd.Validate()
	// TODO: are we expecting there to be no errors here?
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATEDTraceNumber tests validating IATEntryDetail TraceNumber fieldInclusion
func TestIATEDTraceNumber(t *testing.T) {
	testIATEDTraceNumber(t)
}

// BenchmarkIATEDTraceNumber benchmarks validating IATEntryDetail TraceNumber fieldInclusion
func BenchmarkIATEDTraceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATEDTraceNumber(b)
	}
}

// testIATEDAddendaRecordIndicator validates IATEntryDetail AddendaIndicator fieldInclusion
func testIATEDAddendaRecordIndicator(t testing.TB) {
	iatEd := mockIATEntryDetail()
	iatEd.AddendaRecordIndicator = 0
	err := iatEd.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATEDAddendaRecordIndicator tests validating IATEntryDetail AddendaRecordIndicator fieldInclusion
func TestIATEDAddendaRecordIndicator(t *testing.T) {
	testIATEDAddendaRecordIndicator(t)
}

// BenchmarkIATEDAddendaRecordIndicator benchmarks validating IATEntryDetail AddendaRecordIndicator fieldInclusion
func BenchmarkIATEDAddendaRecordIndicator(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATEDAddendaRecordIndicator(b)
	}
}
