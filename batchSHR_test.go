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

// mockBatchSHRHeader creates a BatchSHR BatchHeader
func mockBatchSHRHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = SHR
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "ACH SHR"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockSHREntryDetail creates a BatchSHR EntryDetail
func mockSHREntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingDebit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetSHRCardExpirationDate("0722")
	entry.SetSHRDocumentReferenceNumber("12345678910")
	entry.SetSHRIndividualCardAccountNumber("1234567891123456789")
	entry.SetTraceNumber(mockBatchSHRHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.Category = CategoryForward
	return entry
}

// mockBatchSHR creates a BatchSHR
func mockBatchSHR() *BatchSHR {
	mockBatch := NewBatchSHR(mockBatchSHRHeader())
	mockBatch.AddEntry(mockSHREntryDetail())
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// testBatchSHRHeader creates a BatchSHR BatchHeader
func testBatchSHRHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchSHRHeader())
	err, ok := batch.(*BatchSHR)
	if !ok {
		t.Errorf("Expecting BatchSHR got %T", err)
	}
}

// TestBatchSHRHeader tests validating BatchSHR BatchHeader
func TestBatchSHRHeader(t *testing.T) {
	testBatchSHRHeader(t)
}

// BenchmarkBatchSHRHeader benchmarks validating BatchSHR BatchHeader
func BenchmarkBatchSHRHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRHeader(b)
	}
}

// testBatchSHRCreate validates BatchSHR create
func testBatchSHRCreate(t testing.TB) {
	mockBatch := mockBatchSHR()
	if err := mockBatch.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchSHRCreate tests validating BatchSHR create
func TestBatchSHRCreate(t *testing.T) {
	testBatchSHRCreate(t)
}

// BenchmarkBatchSHRCreate benchmarks validating BatchSHR create
func BenchmarkBatchSHRCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRCreate(b)
	}
}

// testBatchSHRStandardEntryClassCode validates BatchSHR create for an invalid StandardEntryClassCode
func testBatchSHRStandardEntryClassCode(t testing.TB) {
	mockBatch := mockBatchSHR()
	mockBatch.Header.StandardEntryClassCode = WEB
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchSECType) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchSHRStandardEntryClassCode tests validating BatchSHR create for an invalid StandardEntryClassCode
func TestBatchSHRStandardEntryClassCode(t *testing.T) {
	testBatchSHRStandardEntryClassCode(t)
}

// BenchmarkBatchSHRStandardEntryClassCode benchmarks validating BatchSHR create for an invalid StandardEntryClassCode
func BenchmarkBatchSHRStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRStandardEntryClassCode(b)
	}
}

// testBatchSHRServiceClassCodeEquality validates service class code equality
func testBatchSHRServiceClassCodeEquality(t testing.TB) {
	mockBatch := mockBatchSHR()
	mockBatch.GetControl().ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(220, MixedDebitsAndCredits)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchSHRServiceClassCodeEquality tests validating service class code equality
func TestBatchSHRServiceClassCodeEquality(t *testing.T) {
	testBatchSHRServiceClassCodeEquality(t)
}

// BenchmarkBatchSHRServiceClassCodeEquality benchmarks validating service class code equality
func BenchmarkBatchSHRServiceClassCodeEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRServiceClassCodeEquality(b)
	}
}

// testBatchSHRMixedCreditsAndDebits validates BatchSHR create for an invalid MixedCreditsAndDebits
func testBatchSHRMixedCreditsAndDebits(t testing.TB) {
	mockBatch := mockBatchSHR()
	mockBatch.Header.ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(MixedDebitsAndCredits, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchSHRMixedCreditsAndDebits tests validating BatchSHR create for an invalid MixedCreditsAndDebits
func TestBatchSHRMixedCreditsAndDebits(t *testing.T) {
	testBatchSHRMixedCreditsAndDebits(t)
}

// BenchmarkBatchSHRMixedCreditsAndDebits benchmarks validating BatchSHR create for an invalid MixedCreditsAndDebits
func BenchmarkBatchSHRMixedCreditsAndDebits(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRMixedCreditsAndDebits(b)
	}
}

// testBatchSHRCreditsOnly validates BatchSHR create for an invalid CreditsOnly
func testBatchSHRCreditsOnly(t testing.TB) {
	mockBatch := mockBatchSHR()
	mockBatch.Header.ServiceClassCode = CreditsOnly
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(CreditsOnly, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchSHRCreditsOnly tests validating BatchSHR create for an invalid CreditsOnly
func TestBatchSHRCreditsOnly(t *testing.T) {
	testBatchSHRCreditsOnly(t)
}

// BenchmarkBatchSHRCreditsOnly benchmarks validating BatchSHR create for an invalid CreditsOnly
func BenchmarkBatchSHRCreditsOnly(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRCreditsOnly(b)
	}
}

// testBatchSHRAutomatedAccountingAdvices validates BatchSHR create for an invalid AutomatedAccountingAdvices
func testBatchSHRAutomatedAccountingAdvices(t testing.TB) {
	mockBatch := mockBatchSHR()
	mockBatch.Header.ServiceClassCode = AutomatedAccountingAdvices
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(AutomatedAccountingAdvices, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchSHRAutomatedAccountingAdvices tests validating BatchSHR create for an invalid AutomatedAccountingAdvices
func TestBatchSHRAutomatedAccountingAdvices(t *testing.T) {
	testBatchSHRAutomatedAccountingAdvices(t)
}

// BenchmarkBatchSHRAutomatedAccountingAdvices benchmarks validating BatchSHR create for an invalid AutomatedAccountingAdvices
func BenchmarkBatchSHRAutomatedAccountingAdvices(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRAutomatedAccountingAdvices(b)
	}
}

// testBatchSHRTransactionCode validates BatchSHR TransactionCode is not a credit
func testBatchSHRTransactionCode(t testing.TB) {
	mockBatch := mockBatchSHR()
	mockBatch.GetEntries()[0].TransactionCode = CheckingCredit
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchDebitOnly) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchSHRTransactionCode tests validating BatchSHR TransactionCode is not a credit
func TestBatchSHRTransactionCode(t *testing.T) {
	testBatchSHRTransactionCode(t)
}

// BenchmarkBatchSHRTransactionCode benchmarks validating BatchSHR TransactionCode is not a credit
func BenchmarkBatchSHRTransactionCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRTransactionCode(b)
	}
}

// testBatchSHRAddendaCount validates BatchSHR Addendum count of 2
func testBatchSHRAddendaCount(t testing.TB) {
	mockBatch := mockBatchSHR()
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02()
	err := mockBatch.Create()
	// TODO: are we not expecting any errors here?
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchSHRAddendaCount tests validating BatchSHR Addendum count of 2
func TestBatchSHRAddendaCount(t *testing.T) {
	testBatchSHRAddendaCount(t)
}

// BenchmarkBatchSHRAddendaCount benchmarks validating BatchSHR Addendum count of 2
func BenchmarkBatchSHRAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRAddendaCount(b)
	}
}

// testBatchSHRAddendaCountZero validates Addendum count of 0
func testBatchSHRAddendaCountZero(t testing.TB) {
	mockBatch := NewBatchSHR(mockBatchSHRHeader())
	mockBatch.AddEntry(mockSHREntryDetail())
	mockAddenda02 := mockAddenda02()
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality("225", "200")) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchSHRAddendaCountZero tests validating Addendum count of 0
func TestBatchSHRAddendaCountZero(t *testing.T) {
	testBatchSHRAddendaCountZero(t)
}

// BenchmarkBatchSHRAddendaCountZero benchmarks validating Addendum count of 0
func BenchmarkBatchSHRAddendaCountZero(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRAddendaCountZero(b)
	}
}

// TestBatchSHRAddendum98 validates Addenda98 returns an error
func TestBatchSHRAddendum98(t *testing.T) {
	mockBatch := NewBatchSHR(mockBatchSHRHeader())
	mockBatch.AddEntry(mockSHREntryDetail())
	mockAddenda98 := mockAddenda98()
	mockAddenda98.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchSHRAddendum99 validates Addenda99 returns an error
func TestBatchSHRAddendum99(t *testing.T) {
	mockBatch := NewBatchSHR(mockBatchSHRHeader())
	mockBatch.AddEntry(mockSHREntryDetail())
	mockAddenda99 := mockAddenda99()
	mockAddenda99.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryReturn
	mockBatch.GetEntries()[0].Addenda99 = mockAddenda99
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// testBatchSHRInvalidAddendum validates Addendum must be Addenda02
func testBatchSHRInvalidAddendum(t testing.TB) {
	mockBatch := NewBatchSHR(mockBatchSHRHeader())
	mockBatch.AddEntry(mockSHREntryDetail())
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02()
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAddendaCategory) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchSHRInvalidAddendum tests validating Addendum must be Addenda02
func TestBatchSHRInvalidAddendum(t *testing.T) {
	testBatchSHRInvalidAddendum(t)
}

// BenchmarkBatchSHRInvalidAddendum benchmarks validating Addendum must be Addenda02
func BenchmarkBatchSHRInvalidAddendum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRInvalidAddendum(b)
	}
}

// testBatchSHRInvalidAddenda validates Addendum must be Addenda02
func testBatchSHRInvalidAddenda(t testing.TB) {
	mockBatch := NewBatchSHR(mockBatchSHRHeader())
	mockBatch.AddEntry(mockSHREntryDetail())
	addenda02 := mockAddenda02()
	addenda02.TypeCode = "63"
	mockBatch.GetEntries()[0].Addenda02 = addenda02
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchSHRInvalidAddenda tests validating Addendum must be Addenda02
func TestBatchSHRInvalidAddenda(t *testing.T) {
	testBatchSHRInvalidAddenda(t)
}

// BenchmarkBatchSHRInvalidAddenda benchmarks validating Addendum must be Addenda02
func BenchmarkBatchSHRInvalidAddenda(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRInvalidAddenda(b)
	}
}

// testBatchSHRInvalidBuild validates an invalid batch build
func testBatchSHRInvalidBuild(t testing.TB) {
	mockBatch := mockBatchSHR()
	mockBatch.GetHeader().ServiceClassCode = 3
	err := mockBatch.Create()
	if !base.Match(err, ErrServiceClass) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchSHRInvalidBuild tests validating an invalid batch build
func TestBatchSHRInvalidBuild(t *testing.T) {
	testBatchSHRInvalidBuild(t)
}

// BenchmarkBatchSHRInvalidBuild benchmarks validating an invalid batch build
func BenchmarkBatchSHRInvalidBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRInvalidBuild(b)
	}
}

// testBatchSHRCardTransactionType validates BatchSHR create for an invalid CardTransactionType
func testBatchSHRCardTransactionType(t testing.TB) {
	mockBatch := mockBatchSHR()
	mockBatch.GetEntries()[0].DiscretionaryData = "555"
	err := mockBatch.Validate()
	if !base.Match(err, ErrBatchInvalidCardTransactionType) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchSHRCardTransactionType tests validating BatchSHR create for an invalid CardTransactionType
func TestBatchSHRCardTransactionType(t *testing.T) {
	testBatchSHRCardTransactionType(t)
}

// BenchmarkBatchSHRCardTransactionType benchmarks validating BatchSHR create for an invalid CardTransactionType
func BenchmarkBatchSHRCardTransactionType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRCardTransactionType(b)
	}
}

// testBatchSHRCardExpirationDateField validates SHRCardExpirationDate
// characters 0-4 of underlying IdentificationNumber
func testBatchSHRCardExpirationDateField(t testing.TB) {
	mockBatch := mockBatchSHR()
	ts := mockBatch.Entries[0].SHRCardExpirationDateField()
	if ts != "0722" {
		t.Error("Card Expiration Date is invalid")
	}
}

// TestBatchSHRCardExpirationDateField tests validatingSHRCardExpirationDate
// characters 0-4 of underlying IdentificationNumber
func TestBatchSHRCardExpirationDateField(t *testing.T) {
	testBatchSHRCardExpirationDateField(t)
}

// BenchmarkBatchSHRCardExpirationDateField benchmarks validating SHRCardExpirationDate
// characters 0-4 of underlying IdentificationNumber
func BenchmarkBatchSHRCardExpirationDateField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRCardExpirationDateField(b)
	}
}

// testBatchSHRDocumentReferenceNumberField validates SHRDocumentReferenceNumberField
// characters 5-15 of underlying IdentificationNumber
func testBatchSHRDocumentReferenceNumberField(t testing.TB) {
	mockBatch := mockBatchSHR()
	ts := mockBatch.Entries[0].SHRDocumentReferenceNumberField()
	if ts != "12345678910" {
		t.Error("Document Reference Number is invalid")
	}
}

// TestBatchSHRDocumentReferenceNumberField tests validating SHRDocumentReferenceNumberField
// characters 5-15 of underlying IdentificationNumber
func TestBatchSHRDocumentReferenceNumberField(t *testing.T) {
	testBatchSHRDocumentReferenceNumberField(t)
}

// BenchmarkBatchSHRDocumentReferenceNumberField benchmarks validating SHRDocumentReferenceNumberField
// characters 5-15 of underlying IdentificationNumber
func BenchmarkSHRDocumentReferenceNumberField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRDocumentReferenceNumberField(b)
	}
}

// testBatchSHRIndividualCardAccountNumberField validates SHRIndividualCardAccountNumberField
// underlying IndividualName
func testBatchSHRIndividualCardAccountNumberField(t testing.TB) {
	mockBatch := mockBatchSHR()
	ts := mockBatch.Entries[0].SHRIndividualCardAccountNumberField()
	if ts != "0001234567891123456789" {
		t.Error("Individual Card Account Number is invalid")
	}
}

// TestBatchSHRIndividualCardAccountNumberField tests validating SHRIndividualCardAccountNumberField
// underlying IndividualName
func TestBatchSHRIndividualCardAccountNumberField(t *testing.T) {
	testBatchSHRIndividualCardAccountNumberField(t)
}

// BenchmarkBatchSHRIndividualCardAccountNumberField benchmarks validating SHRIndividualCardAccountNumberField
// underlying IndividualName
func BenchmarkBatchSHRDocumentReferenceNumberField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRIndividualCardAccountNumberField(b)
	}
}

// testSHRCardExpirationDateMonth validates the month is valid for CardExpirationDate
func testSHRCardExpirationDateMonth(t testing.TB) {
	mockBatch := mockBatchSHR()
	mockBatch.GetEntries()[0].SetSHRCardExpirationDate("1306")
	err := mockBatch.Validate()
	if !base.Match(err, ErrValidMonth) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestSHRCardExpirationDateMonth tests validating the month is valid for CardExpirationDate
func TestSHRSHRCardExpirationDateMonth(t *testing.T) {
	testSHRCardExpirationDateMonth(t)
}

// BenchmarkSHRCardExpirationDateMonth test validating the month is valid for CardExpirationDate
func BenchmarkSHRCardExpirationDateMonth(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testSHRCardExpirationDateMonth(b)
	}
}

// testSHRCardExpirationDateYear validates the year is valid for CardExpirationDate
func testSHRCardExpirationDateYear(t testing.TB) {
	mockBatch := mockBatchSHR()
	mockBatch.GetEntries()[0].SetSHRCardExpirationDate("0612")
	err := mockBatch.Validate()
	if !base.Match(err, ErrValidYear) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestSHRCardExpirationDateYear tests validating the year is valid for CardExpirationDate
func TestSHRSHRCardExpirationDateYear(t *testing.T) {
	testSHRCardExpirationDateYear(t)
}

// BenchmarkSHRCardExpirationDateYear test validating the year is valid for CardExpirationDate
func BenchmarkSHRCardExpirationDateYear(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testSHRCardExpirationDateYear(b)
	}
}

// TestBatchSHRAddendum99Category validates Addenda99 returns an error
func TestBatchSHRAddendum99Category(t *testing.T) {
	mockBatch := NewBatchSHR(mockBatchSHRHeader())
	mockBatch.AddEntry(mockSHREntryDetail())
	mockAddenda99 := mockAddenda99()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	mockBatch.Entries[0].Category = CategoryNOC
	mockBatch.Entries[0].Addenda99 = mockAddenda99
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAddendaCategory) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchSHRTerminalState validates TerminalState returns an error if invalid from usabbrev
func TestBatchSHRTerminalState(t *testing.T) {
	mockBatch := NewBatchSHR(mockBatchSHRHeader())
	mockBatch.AddEntry(mockSHREntryDetail())
	mockAddenda02 := mockAddenda02()
	mockAddenda02.TerminalState = "YY"
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	err := mockBatch.Create()
	if !base.Match(err, ErrValidState) {
		t.Errorf("%T: %s", err, err)
	}
}
