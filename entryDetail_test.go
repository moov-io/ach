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
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/moov-io/base"
)

// mockEntryDetail creates an entry detail
func mockEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("121042882")
	entry.DFIAccountNumber = "123456789"
	entry.Amount = 100000000
	entry.IndividualName = "Wade Arnold"
	entry.SetTraceNumber(mockBatchHeader().ODFIIdentification, 1)
	entry.IdentificationNumber = "ABC##jvkdjfuiwn"
	entry.Category = CategoryForward
	return entry
}

// testMockEntryDetail validates an entry detail record
func testMockEntryDetail(t testing.TB) {
	entry := mockEntryDetail()
	if err := entry.Validate(); err != nil {
		t.Error("mockEntryDetail does not validate and will break other tests")
	}
	if entry.TransactionCode != CheckingCredit {
		t.Error("TransactionCode dependent default value has changed")
	}
	if entry.DFIAccountNumber != "123456789" {
		t.Error("DFIAccountNumber dependent default value has changed")
	}
	if entry.Amount != 100000000 {
		t.Error("Amount dependent default value has changed")
	}
	if entry.IndividualName != "Wade Arnold" {
		t.Error("IndividualName dependent default value has changed")
	}
	if entry.TraceNumber != "121042880000001" {
		t.Errorf("TraceNumber dependent default value has changed %v", entry.TraceNumber)
	}
}

// TestMockEntryDetail tests validating an entry detail record
func TestMockEntryDetail(t *testing.T) {
	testMockEntryDetail(t)
}

// BenchmarkMockEntryDetail benchmarks validating an entry detail record
func BenchmarkMockEntryDetail(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockEntryDetail(b)
	}
}

// testParseEntryDetail parses a known entry detail record string.
func testParseEntryDetail(t testing.TB) {
	var line = "62705320001912345            0000010500c-1            Arnold Wade           DD0076401255655291"
	r := NewReader(strings.NewReader(line))
	r.addCurrentBatch(NewBatchPPD(mockBatchPPDHeader()))
	r.currentBatch.SetHeader(mockBatchHeader())
	r.line = line
	if err := r.parseEntryDetail(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentBatch.GetEntries()[0]

	if record.TransactionCode != CheckingDebit {
		t.Errorf("TransactionCode Expected '27' got: %v", record.TransactionCode)
	}
	if record.RDFIIdentificationField() != "05320001" {
		t.Errorf("RDFIIdentification Expected '05320001' got: '%v'", record.RDFIIdentificationField())
	}
	if record.CheckDigit != "9" {
		t.Errorf("CheckDigit Expected '9' got: %v", record.CheckDigit)
	}
	if record.DFIAccountNumberField() != "12345            " {
		t.Errorf("DfiAccountNumber Expected '12345            ' got: %v", record.DFIAccountNumberField())
	}
	if record.AmountField() != "0000010500" {
		t.Errorf("Amount Expected '0000010500' got: %v", record.AmountField())
	}

	if record.IdentificationNumber != "c-1            " {
		t.Errorf("IdentificationNumber Expected 'c-1            ' got: %v", record.IdentificationNumber)
	}
	if record.IndividualName != "Arnold Wade           " {
		t.Errorf("IndividualName Expected 'Arnold Wade           ' got: %v", record.IndividualName)
	}
	if record.DiscretionaryData != "DD" {
		t.Errorf("DiscretionaryData Expected 'DD' got: %v", record.DiscretionaryData)
	}
	if record.AddendaRecordIndicator != 0 {
		t.Errorf("AddendaRecordIndicator Expected '0' got: %v", record.AddendaRecordIndicator)
	}
	if record.TraceNumberField() != "076401255655291" {
		t.Errorf("TraceNumber Expected '076401255655291' got: %v", record.TraceNumberField())
	}
}

// TestParseEntryDetail tests parsing a known entry detail record string.
func TestParseEntryDetail(t *testing.T) {
	testParseEntryDetail(t)
}

// BenchmarkParseEntryDetail benchmarks parsing a known entry detail record string.
func BenchmarkParseEntryDetail(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testParseEntryDetail(b)
	}
}

// testEDString validates that a known parsed entry
// detail can be returned to a string of the same value
func testEDString(t testing.TB) {
	var line = "62705320001912345            0000010500c-1            Arnold Wade           DD0076401255655291"
	r := NewReader(strings.NewReader(line))
	r.addCurrentBatch(NewBatchPPD(mockBatchPPDHeader()))
	r.currentBatch.SetHeader(mockBatchHeader())
	r.line = line
	if err := r.parseEntryDetail(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentBatch.GetEntries()[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestEDString tests validating that a known parsed entry
// detail can be returned to a string of the same value
func TestEDString(t *testing.T) {
	testEDString(t)
}

// BenchmarkEDString benchmarks validating that a known parsed entry
// detail can be returned to a string of the same value
func BenchmarkEDString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDString(b)
	}
}

// testValidateEDTransactionCode validates error if transaction code is not valid
func testValidateEDTransactionCode(t testing.TB) {
	ed := mockEntryDetail()
	ed.TransactionCode = 63
	err := ed.Validate()
	if !base.Match(err, ErrTransactionCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestValidateEDTransactionCode tests validating error if transaction code is not valid
func TestValidateEDTransactionCode(t *testing.T) {
	testValidateEDTransactionCode(t)
}

// BenchmarkValidateEDTransactionCode benchmarks validating error if transaction code is not valid
func BenchmarkValidateEDTransactionCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateEDTransactionCode(b)
	}
}

func TestEntryDetailAmountNegative(t *testing.T) {
	ed := mockEntryDetail()
	ed.Amount = -100
	if err := ed.Validate(); !base.Match(err, ErrNegativeAmount) {
		t.Errorf("%T: %s", err, err)
	}
}

// testEDFieldInclusion validates entry detail field inclusion
func testEDFieldInclusion(t testing.TB) {
	ed := mockEntryDetail()
	ed.Amount = 0
	err := ed.Validate()
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestEDFieldInclusion tests validating entry detail field inclusion
func TestEDFieldInclusion(t *testing.T) {
	testEDFieldInclusion(t)
}

// BenchmarkEDFieldInclusion benchmarks validating entry detail field inclusion
func BenchmarkEDFieldInclusion(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDFieldInclusion(b)
	}
}

// testEDdfiAccountNumberAlphaNumeric validates DFI account number is alpha numeric
func testEDdfiAccountNumberAlphaNumeric(t testing.TB) {
	ed := mockEntryDetail()
	ed.DFIAccountNumber = "®"
	err := ed.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestEDdfiAccountNumberAlphaNumeric tests validating DFI account number is alpha numeric
func TestEDdfiAccountNumberAlphaNumeric(t *testing.T) {
	testEDdfiAccountNumberAlphaNumeric(t)
}

// BenchmarkEDdfiAccountNumberAlphaNumeric benchmarks validating DFI account number is alpha numeric
func BenchmarkEDdfiAccountNumberAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDdfiAccountNumberAlphaNumeric(b)
	}
}

// testEDIdentificationNumberAlphaNumeric validates identification number is alpha numeric
func testEDIdentificationNumberAlphaNumeric(t testing.TB) {
	ed := mockEntryDetail()
	ed.IdentificationNumber = "®"
	err := ed.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestEDIdentificationNumberAlphaNumeric tests validating identification number is alpha numeric
func TestEDIdentificationNumberAlphaNumeric(t *testing.T) {
	testEDIdentificationNumberAlphaNumeric(t)
}

// BenchmarkEDIdentificationNumberAlphaNumeric benchmarks validating identification number is alpha numeric
func BenchmarkEDIdentificationNumberAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDIdentificationNumberAlphaNumeric(b)
	}
}

// testEDIndividualNameAlphaNumeric validates individual name is alpha numeric
func testEDIndividualNameAlphaNumeric(t testing.TB) {
	ed := mockEntryDetail()
	ed.IndividualName = "W®DE"
	err := ed.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestEDIndividualNameAlphaNumeric tests validating individual name is alpha numeric
func TestEDIndividualNameAlphaNumeric(t *testing.T) {
	testEDIndividualNameAlphaNumeric(t)
}

// BenchmarkEDIndividualNameAlphaNumeric benchmarks validating individual name is alpha numeric
func BenchmarkEDIndividualNameAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDIndividualNameAlphaNumeric(b)
	}
}

// testEDDiscretionaryDataAlphaNumeric validates discretionary data is alpha numeric
func testEDDiscretionaryDataAlphaNumeric(t testing.TB) {
	ed := mockEntryDetail()
	ed.DiscretionaryData = "®!"
	err := ed.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestEDDiscretionaryDataAlphaNumeric tests validating discretionary data is alpha numeric
func TestEDDiscretionaryDataAlphaNumeric(t *testing.T) {
	testEDDiscretionaryDataAlphaNumeric(t)
}

// BenchmarkEDDiscretionaryDataAlphaNumeric benchmarks validating discretionary data is alpha numeric
func BenchmarkEDDiscretionaryDataAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDDiscretionaryDataAlphaNumeric(b)
	}
}

// testEDisCheckDigit validates check digit
func testEDisCheckDigit(t testing.TB) {
	ed := mockEntryDetail()
	ed.CheckDigit = "1"
	err := ed.Validate()
	if !base.Match(err, NewErrValidCheckDigit(7)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestEDisCheckDigit tests validating check digit
func TestEDisCheckDigit(t *testing.T) {
	testEDisCheckDigit(t)
}

// BenchmarkEDSetRDFI benchmarks validating check digit
func BenchmarkEDisCheckDigit(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDisCheckDigit(b)
	}
}

// testEDSetRDFI validates setting RDFI
func testEDSetRDFI(t testing.TB) {
	ed := NewEntryDetail()
	ed.SetRDFI("810866774")
	if ed.RDFIIdentification != "81086677" {
		t.Error("RDFI identification")
	}
	if ed.CheckDigit != "4" {
		t.Error("Unexpected check digit")
	}
}

// TestEDSetRDFI tests validating setting RDFI
func TestEDSetRDFI(t *testing.T) {
	testEDSetRDFI(t)
}

// BenchmarkEDSetRDFI benchmarks validating setting RDFI
func BenchmarkEDSetRDFI(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDSetRDFI(b)
	}
}

// testEDFieldInclusionTransactionCode validates transaction code field inclusion
func testEDFieldInclusionTransactionCode(t testing.TB) {
	entry := mockEntryDetail()
	entry.TransactionCode = 0
	err := entry.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestEDFieldInclusionTransactionCode tests validating transaction code field inclusion
func TestEDFieldInclusionTransactionCode(t *testing.T) {
	testEDFieldInclusionTransactionCode(t)
}

// BenchmarkEDFieldInclusionTransactionCode benchmarks validating transaction code field inclusion
func BenchmarkEDFieldInclusionTransactionCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDFieldInclusionTransactionCode(b)
	}
}

// testEDFieldInclusionRDFIIdentification validates RDFI identification field inclusion
func testEDFieldInclusionRDFIIdentification(t testing.TB) {
	entry := mockEntryDetail()
	entry.RDFIIdentification = ""
	err := entry.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestEDFieldInclusionRDFIIdentification tests validating RDFI identification field inclusion
func TestEDFieldInclusionRDFIIdentification(t *testing.T) {
	testEDFieldInclusionRDFIIdentification(t)
}

// BenchmarkEDFieldInclusionRDFIIdentification benchmarks validating RDFI identification field inclusion
func BenchmarkEDFieldInclusionRDFIIdentification(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDFieldInclusionRDFIIdentification(b)
	}
}

// testEDFieldInclusionDFIAccountNumber validates DFI account number field inclusion
func testEDFieldInclusionDFIAccountNumber(t testing.TB) {
	entry := mockEntryDetail()
	entry.DFIAccountNumber = ""
	err := entry.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestEDFieldInclusionDFIAccountNumber tests validating DFI account number field inclusion
func TestEDFieldInclusionDFIAccountNumber(t *testing.T) {
	testEDFieldInclusionDFIAccountNumber(t)
}

// BenchmarkEDFieldInclusionDFIAccountNumber benchmarks validating DFI account number field inclusion
func BenchmarkEDFieldInclusionDFIAccountNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDFieldInclusionDFIAccountNumber(b)
	}
}

// testEDFieldInclusionIndividualName validates individual name field inclusion
func testEDFieldInclusionIndividualName(t testing.TB) {
	entry := mockEntryDetail()
	entry.IndividualName = ""
	err := entry.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestEDFieldInclusionIndividualName tests validating individual name field inclusion
func TestEDFieldInclusionIndividualName(t *testing.T) {
	testEDFieldInclusionIndividualName(t)
}

// BenchmarkEDFieldInclusionIndividualName benchmarks validating individual name field inclusion
func BenchmarkEDFieldInclusionIndividualName(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDFieldInclusionIndividualName(b)
	}
}

// testEDFieldInclusionTraceNumber validates trace number field inclusion
func testEDFieldInclusionTraceNumber(t testing.TB) {
	entry := mockEntryDetail()
	entry.TraceNumber = "0"
	err := entry.Validate()
	// TODO: are we expecting to see no error here?
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestEDFieldInclusionTraceNumber tests validating trace number field inclusion
func TestEDFieldInclusionTraceNumber(t *testing.T) {
	testEDFieldInclusionTraceNumber(t)
}

// BenchmarkEDFieldInclusionTraceNumber benchmarks validating trace number field inclusion
func BenchmarkEDFieldInclusionTraceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDFieldInclusionTraceNumber(b)
	}
}

// testEDAddAddenda99 validates adding Addenda99 to an entry detail
func testEDAddAddenda99(t testing.TB) {
	entry := mockEntryDetail()
	entry.Addenda99 = mockAddenda99()
	entry.Category = CategoryReturn
	entry.AddendaRecordIndicator = 1
	if entry.Category != CategoryReturn {
		t.Error("Addenda99 added and isReturn is false")
	}
	if entry.AddendaRecordIndicator != 1 {
		t.Error("Addenda99 added and record indicator is not 1")
	}

}

// TestEDAddAddenda99 tests validating adding Addenda99 to an entry detail
func TestEDAddAddenda99(t *testing.T) {
	testEDAddAddenda99(t)
}

// BenchmarkEDAddAddenda99 benchmarks validating adding Addenda99 to an entry detail
func BenchmarkEDAddAddenda99(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDAddAddenda99(b)
	}
}

// testEDAddAddenda99Twice validates only one Addenda99 is added to an entry detail
func testEDAddAddenda99Twice(t testing.TB) {
	entry := mockEntryDetail()
	entry.Addenda99 = mockAddenda99()
	entry.Addenda99 = mockAddenda99()
	entry.Category = CategoryReturn
	if entry.Category != CategoryReturn {
		t.Error("Addenda99 added and Category is not CategoryReturn")
	}
}

// TestEDAddAddenda99Twice tests validating only one Addenda99 is added to an entry detail
func TestEDAddAddenda99Twice(t *testing.T) {
	testEDAddAddenda99Twice(t)
}

// BenchmarkEDAddAddenda99Twice benchmarks validating only one Addenda99 is added to an entry detail
func BenchmarkEDAddAddenda99Twice(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDAddAddenda99Twice(b)
	}
}

// testEDCreditOrDebit validates debit and credit transaction code
func testEDCreditOrDebit(t testing.TB) {
	entry := mockEntryDetail()
	if entry.CreditOrDebit() != "C" { // our mock's default
		t.Errorf("TransactionCode %v expected a Credit(C) got %v", entry.TransactionCode, entry.CreditOrDebit())
	}

	// TransactionCode -> C or D
	var cases = map[int]string{
		// invalid
		-1:  "",
		00:  "", // invalid
		1:   "",
		108: "",
		// valid
		22: "C",
		23: "C",
		27: "D",
		28: "D",
		32: "C",
		33: "C",
		37: "D",
		38: "D",
	}
	for code, expected := range cases {
		entry.TransactionCode = code
		if v := entry.CreditOrDebit(); v != expected {
			t.Errorf("TransactionCode %d expected %s, got %s", code, expected, v)
		}
	}
}

// TestEDCreditOrDebit tests validating debit and credit transaction code
func TestEDCreditOrDebit(t *testing.T) {
	testEDCreditOrDebit(t)
}

// BenchmarkEDCreditOrDebit benchmarks validating debit and credit transaction code
func BenchmarkEDCreditOrDebit(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDCreditOrDebit(b)
	}
}

// testValidateEDCheckDigit validates CheckDigit error
func testValidateEDCheckDigit(t testing.TB) {
	ed := mockEntryDetail()
	ed.CheckDigit = "XYZ"
	err := ed.Validate()
	if !base.Match(err, &strconv.NumError{}) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestValidateEDCheckDigit tests validating validates CheckDigit error
func TestValidateEDCheckDigit(t *testing.T) {
	testValidateEDCheckDigit(t)
}

// BenchmarkValidateEDCheckDigit benchmarks validating CheckDigit error
func BenchmarkValidateEDCheckDigit(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateEDCheckDigit(b)
	}
}

func TestEntryDetail__CategoryJSON(t *testing.T) {
	ed := mockEntryDetail()

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(ed); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(buf.String(), `"category":"Forward"`) {
		t.Error(buf.String())
	}
	buf.Reset()

	// read our return file and marshal
	file, err := readACHFilepath(filepath.Join("test", "testdata", "return-WEB.ach"))
	if err != nil {
		t.Fatal(err)
	}
	if n := len(file.ReturnEntries); n != 2 {
		t.Errorf("got %d ReturnEntries", n)
	}
	if err := json.NewEncoder(&buf).Encode(file); err != nil {
		t.Fatal(err)
	}
	// There are two ReturnEntries and two Batches
	if n := strings.Count(buf.String(), `"category":"Return"`); n != 4 {
		// return-WEB.ach has two EntryDetail records
		t.Errorf("got %d category:Return\n%s", n, buf.String())
	}
	buf.Reset()

	// COR / Notification of Change
	file, err = readACHFilepath(filepath.Join("test", "testdata", "cor-example.ach"))
	if err != nil {
		t.Fatal(err)
	}
	if err := json.NewEncoder(&buf).Encode(file); err != nil {
		t.Fatal(err)
	}
	if n := len(file.ReturnEntries); n != 0 {
		t.Errorf("got %d ReturnEntries", n)
	}
	if n := len(file.NotificationOfChange); n != 1 {
		t.Errorf("got %d NotificationOfChange", n)
	}
	// one NOC entry in NotificationOfChange and one in Batches
	if n := strings.Count(buf.String(), `"category":"NOC"`); n != 2 {
		// return-WEB.ach has two EntryDetail records
		t.Errorf("got %d category:NOC\n%s", n, buf.String())
	}
}

func TestEntryDetail__ParseReturn(t *testing.T) {
	file, err := readACHFilepath(filepath.Join("test", "testdata", "return-WEB.ach"))
	if err != nil {
		t.Fatal(err)
	}
	if n := len(file.Batches); n != 2 {
		t.Errorf("got %d batches: %#v", n, file.Batches)
	}
	for i := range file.Batches {
		entries := file.Batches[i].GetEntries()
		if n := len(entries); n != 1 {
			t.Errorf("got %d EntryDetail records: %#v", n, entries)
		}
		for j := range entries {
			if entries[j].Category != CategoryReturn {
				t.Errorf("EntryDetail.Category=%s\n  %#v", entries[j].Category, entries[j])
			}
		}
	}
}

func TestEntryDetail__ParseNOC(t *testing.T) {
	file, err := readACHFilepath(filepath.Join("test", "testdata", "cor-example.ach"))
	if err != nil {
		t.Fatal(err)
	}
	if n := len(file.Batches); n != 1 {
		t.Errorf("got %d batches: %#v", n, file.Batches)
	}

	entries := file.Batches[0].GetEntries()
	if n := len(entries); n != 1 {
		t.Errorf("got %d EntryDetail records: %#v", n, entries)
	}
	if entries[0].Category != CategoryNOC {
		t.Errorf("EntryDetail.Category=%s\n  %#v", entries[0].Category, entries[0])
	}
}

func TestEntryDetail__SetValidation(t *testing.T) {
	ed := mockEntryDetail()
	ed.SetValidation(&ValidateOpts{
		CheckTransactionCode: func(code int) error {
			if code == 999 {
				return nil
			}
			return fmt.Errorf("unexpected TransactionCode: %d", code)
		},
	})
	ed.TransactionCode = 999
	if err := ed.Validate(); err != nil {
		t.Fatal(err)
	}

	// nil out
	ed = nil
	ed.SetValidation(&ValidateOpts{})
}

func TestEntryDetail__LargeAmountStrings(t *testing.T) {
	ed := mockEntryDetail()
	ed.Amount = math.MaxInt64

	if err := ed.Validate(); err == nil {
		t.Error("expected error")
	} else {
		if !strings.Contains(err.Error(), "Amount 9223372036854775807 does not match formatted value 6854775807") {
			t.Errorf("unexpected error: %v", err)
		}
	}
}
