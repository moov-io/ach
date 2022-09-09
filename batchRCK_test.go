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

// mockBatchRCKHeader creates a BatchRCK BatchHeader
func mockBatchRCKHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = RCK
	bh.CompanyName = "Company Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "REDEPCHECK"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockRCKEntryDetail creates a BatchRCK EntryDetail
func mockRCKEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingDebit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 2400
	entry.SetCheckSerialNumber("123456789")
	entry.IndividualName = "Wade Arnold"
	entry.SetTraceNumber(mockBatchRCKHeader().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

// mockBatchRCK creates a BatchRCK
func mockBatchRCK() *BatchRCK {
	mockBatch := NewBatchRCK(mockBatchRCKHeader())
	mockBatch.AddEntry(mockRCKEntryDetail())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// mockBatchRCKHeaderCredit creates a BatchRCK BatchHeader
func mockBatchRCKHeaderCredit() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = RCK
	bh.CompanyName = "Company Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "REDEPCHECK"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockRCKEntryDetailCredit creates a BatchRCK EntryDetail with a credit entry
func mockRCKEntryDetailCredit() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 2400
	entry.SetCheckSerialNumber("123456789")
	entry.IndividualName = "Wade Arnold"
	entry.SetTraceNumber(mockBatchRCKHeader().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

// mockBatchRCKCredit creates a BatchRCK with a credit entry
func mockBatchRCKCredit() *BatchRCK {
	mockBatch := NewBatchRCK(mockBatchRCKHeaderCredit())
	mockBatch.AddEntry(mockRCKEntryDetailCredit())
	return mockBatch
}

// testBatchRCKHeader creates a BatchRCK BatchHeader
func testBatchRCKHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchRCKHeader())
	err, ok := batch.(*BatchRCK)
	if !ok {
		t.Errorf("Expecting BatchRCK got %T", err)
	}
}

// TestBatchRCKHeader tests validating BatchRCK BatchHeader
func TestBatchRCKHeader(t *testing.T) {
	testBatchRCKHeader(t)
}

// BenchmarkBatchRCKHeader benchmarks validating BatchRCK BatchHeader
func BenchmarkBatchRCKHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchRCKHeader(b)
	}
}

// testBatchRCKCreate validates BatchRCK create
func testBatchRCKCreate(t testing.TB) {
	mockBatch := mockBatchRCK()
	if err := mockBatch.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchRCKCreate tests validating BatchRCK create
func TestBatchRCKCreate(t *testing.T) {
	testBatchRCKCreate(t)
}

// BenchmarkBatchRCKCreate benchmarks validating BatchRCK create
func BenchmarkBatchRCKCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchRCKCreate(b)
	}
}

// testBatchRCKStandardEntryClassCode validates BatchRCK create for an invalid StandardEntryClassCode
func testBatchRCKStandardEntryClassCode(t testing.TB) {
	mockBatch := mockBatchRCK()
	mockBatch.Header.StandardEntryClassCode = WEB
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchSECType) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchRCKStandardEntryClassCode tests validating BatchRCK create for an invalid StandardEntryClassCode
func TestBatchRCKStandardEntryClassCode(t *testing.T) {
	testBatchRCKStandardEntryClassCode(t)
}

// BenchmarkBatchRCKStandardEntryClassCode benchmarks validating BatchRCK create for an invalid StandardEntryClassCode
func BenchmarkBatchRCKStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchRCKStandardEntryClassCode(b)
	}
}

// testBatchRCKServiceClassCodeEquality validates service class code equality
func testBatchRCKServiceClassCodeEquality(t testing.TB) {
	mockBatch := mockBatchRCK()
	mockBatch.GetControl().ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(220, MixedDebitsAndCredits)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchRCKServiceClassCodeEquality tests validating service class code equality
func TestBatchRCKServiceClassCodeEquality(t *testing.T) {
	testBatchRCKServiceClassCodeEquality(t)
}

// BenchmarkBatchRCKServiceClassCodeEquality benchmarks validating service class code equality
func BenchmarkBatchRCKServiceClassCodeEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchRCKServiceClassCodeEquality(b)
	}
}

// testBatchRCKMixedCreditsAndDebits validates BatchRCK create for an invalid MixedCreditsAndDebits
func testBatchRCKMixedCreditsAndDebits(t testing.TB) {
	mockBatch := mockBatchRCK()
	mockBatch.Header.ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(MixedDebitsAndCredits, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchRCKMixedCreditsAndDebits tests validating BatchRCK create for an invalid MixedCreditsAndDebits
func TestBatchRCKMixedCreditsAndDebits(t *testing.T) {
	testBatchRCKMixedCreditsAndDebits(t)
}

// BenchmarkBatchRCKMixedCreditsAndDebits benchmarks validating BatchRCK create for an invalid MixedCreditsAndDebits
func BenchmarkBatchRCKMixedCreditsAndDebits(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchRCKMixedCreditsAndDebits(b)
	}
}

// testBatchRCKCreditsOnly validates BatchRCK create for an invalid CreditsOnly
func testBatchRCKCreditsOnly(t testing.TB) {
	mockBatch := mockBatchRCK()
	mockBatch.Header.ServiceClassCode = CreditsOnly
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(CreditsOnly, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchRCKCreditsOnly tests validating BatchRCK create for an invalid CreditsOnly
func TestBatchRCKCreditsOnly(t *testing.T) {
	testBatchRCKCreditsOnly(t)
}

// BenchmarkBatchRCKCreditsOnly benchmarks validating BatchRCK create for an invalid CreditsOnly
func BenchmarkBatchRCKCreditsOnly(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchRCKCreditsOnly(b)
	}
}

// testBatchRCKAutomatedAccountingAdvices validates BatchRCK create for an invalid AutomatedAccountingAdvices
func testBatchRCKAutomatedAccountingAdvices(t testing.TB) {
	mockBatch := mockBatchRCK()
	mockBatch.Header.ServiceClassCode = AutomatedAccountingAdvices
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(AutomatedAccountingAdvices, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchRCKAutomatedAccountingAdvices tests validating BatchRCK create for an invalid AutomatedAccountingAdvices
func TestBatchRCKAutomatedAccountingAdvices(t *testing.T) {
	testBatchRCKAutomatedAccountingAdvices(t)
}

// BenchmarkBatchRCKAutomatedAccountingAdvices benchmarks validating BatchRCK create for an invalid AutomatedAccountingAdvices
func BenchmarkBatchRCKAutomatedAccountingAdvices(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchRCKAutomatedAccountingAdvices(b)
	}
}

// testBatchRCKCompanyEntryDescription validates BatchRCK create for an invalid CompanyEntryDescription
func testBatchRCKCompanyEntryDescription(t testing.TB) {
	mockBatch := mockBatchRCK()
	mockBatch.Header.CompanyEntryDescription = "XYZ975"
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchCompanyEntryDescriptionREDEPCHECK) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchRCKCompanyEntryDescription validates BatchRCK create for an invalid CompanyEntryDescription
func TestBatchRCKCompanyEntryDescription(t *testing.T) {
	testBatchRCKCompanyEntryDescription(t)
}

// BenchmarkBatchRCKCompanyEntryDescription validates BatchRCK create for an invalid CompanyEntryDescription
func BenchmarkBatchRCKCompanyEntryDescription(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchRCKCompanyEntryDescription(b)
	}
}

// testBatchRCKAmount validates BatchRCK create for an invalid Amount
func testBatchRCKAmount(t testing.TB) {
	mockBatch := mockBatchRCK()
	mockBatch.Entries[0].Amount = 250001
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchAmount(250001, 250000)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchRCKAmount validates BatchRCK create for an invalid Amount
func TestBatchRCKAmount(t *testing.T) {
	testBatchRCKAmount(t)
}

// BenchmarkBatchRCKAmount validates BatchRCK create for an invalid Amount
func BenchmarkBatchRCKAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchRCKAmount(b)
	}
}

// testBatchRCKCheckSerialNumber validates BatchRCK CheckSerialNumber / IdentificationNumber is a mandatory field
func testBatchRCKCheckSerialNumber(t testing.TB) {
	mockBatch := mockBatchRCK()
	// modify CheckSerialNumber / IdentificationNumber to empty string
	mockBatch.GetEntries()[0].SetCheckSerialNumber("")
	err := mockBatch.Validate()
	if !base.Match(err, ErrBatchCheckSerialNumber) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchRCKCheckSerialNumber tests validating BatchRCK
// CheckSerialNumber / IdentificationNumber is a mandatory field
func TestBatchRCKCheckSerialNumber(t *testing.T) {
	testBatchRCKCheckSerialNumber(t)
}

// BenchmarkBatchRCKCheckSerialNumber benchmarks validating BatchRCK
// CheckSerialNumber / IdentificationNumber is a mandatory field
func BenchmarkBatchRCKCheckSerialNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchRCKCheckSerialNumber(b)
	}
}

// testBatchRCKTransactionCode validates BatchRCK TransactionCode is not a credit
func testBatchRCKTransactionCode(t testing.TB) {
	mockBatch := mockBatchRCKCredit()
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchDebitOnly) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchRCKTransactionCode tests validating BatchRCK TransactionCode is not a credit
func TestBatchRCKTransactionCode(t *testing.T) {
	testBatchRCKTransactionCode(t)
}

// BenchmarkBatchRCKTransactionCode benchmarks validating BatchRCK TransactionCode is not a credit
func BenchmarkBatchRCKTransactionCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchRCKTransactionCode(b)
	}
}

// testBatchRCKAddendaCount validates BatchRCK addenda count
func testBatchRCKAddendaCount(t testing.TB) {
	mockBatch := mockBatchRCK()
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAddendaCategory) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchRCKAddendaCount tests validating BatchRCK addenda count
func TestBatchRCKAddendaCount(t *testing.T) {
	testBatchRCKAddendaCount(t)
}

// BenchmarkBatchRCKAddendaCount benchmarks validating BatchRCK addenda count
func BenchmarkBatchRCKAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchRCKAddendaCount(b)
	}
}

// testBatchRCKParseCheckSerialNumber validates BatchRCK create
func testBatchRCKParseCheckSerialNumber(t testing.TB) {
	mockBatch := mockBatchRCK()
	if err := mockBatch.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	checkSerialNumber := "123456789      "
	if checkSerialNumber != mockBatch.GetEntries()[0].CheckSerialNumberField() {
		t.Errorf("RecordType Expected '123456789' got: %v", mockBatch.GetEntries()[0].CheckSerialNumberField())
	}
}

// TestBatchRCKParseCheckSerialNumber tests validating BatchRCK create
func TestBatchRCKParseCheckSerialNumber(t *testing.T) {
	testBatchRCKParseCheckSerialNumber(t)
}

// BenchmarkBatchRCKParseCheckSerialNumber benchmarks validating BatchRCK create
func BenchmarkBatchRCKParseCheckSerialNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchRCKParseCheckSerialNumber(b)
	}
}

// testBatchRCKInvalidBuild validates an invalid batch build
func testBatchRCKInvalidBuild(t testing.TB) {
	mockBatch := mockBatchRCK()
	mockBatch.GetHeader().ServiceClassCode = 3
	err := mockBatch.Create()
	if !base.Match(err, ErrServiceClass) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchRCKInvalidBuild tests validating an invalid batch build
func TestBatchRCKInvalidBuild(t *testing.T) {
	testBatchRCKInvalidBuild(t)
}

// BenchmarkBatchRCKInvalidBuild benchmarks validating an invalid batch build
func BenchmarkRCKBatchInvalidBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchRCKInvalidBuild(b)
	}
}

// TestBatchRCKAddendum98 validates Addenda98 returns an error
func TestBatchRCKAddendum98(t *testing.T) {
	mockBatch := NewBatchRCK(mockBatchRCKHeader())
	mockBatch.AddEntry(mockRCKEntryDetail())
	mockAddenda98 := mockAddenda98()
	mockAddenda98.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchRCKAddendum99 validates Addenda99 returns an error
func TestBatchRCKAddendum99(t *testing.T) {
	mockBatch := NewBatchRCK(mockBatchRCKHeader())
	mockBatch.AddEntry(mockRCKEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockAddenda99.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryReturn
	mockBatch.GetEntries()[0].Addenda99 = mockAddenda99
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// testBatchRCKServiceClassCodeEquality validates MixedDebitsAndCredits service class code
func testBatchRCKMixedDebitsAndCreditsServiceClassCode(t testing.TB) {
	mockBatch := mockBatchRCK()
	mockBatch.GetControl().ServiceClassCode = MixedDebitsAndCredits
	mockBatch.Header.ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchRCKServiceClassCodeEquality tests validating MixedDebitsAndCredits service class code
func TestBatchRCKMixedDebitsAndCreditsServiceClassCode(t *testing.T) {
	testBatchRCKMixedDebitsAndCreditsServiceClassCode(t)
}
