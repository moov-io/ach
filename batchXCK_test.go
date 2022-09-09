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

// mockBatchXCKHeader creates a BatchXCK BatchHeader
func mockBatchXCKHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = XCK
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = XCK
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockXCKEntryDetail creates a BatchXCK EntryDetail
func mockXCKEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingDebit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetCheckSerialNumber("123456789")
	entry.SetProcessControlField("CHECK1")
	entry.SetItemResearchNumber("182726")
	entry.DiscretionaryData = ""
	entry.SetTraceNumber(mockBatchXCKHeader().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

// mockBatchXCK creates a BatchXCK
func mockBatchXCK() *BatchXCK {
	mockBatch := NewBatchXCK(mockBatchXCKHeader())
	mockBatch.AddEntry(mockXCKEntryDetail())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// mockBatchXCKHeaderCredit creates a BatchXCK BatchHeader
func mockBatchXCKHeaderCredit() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = XCK
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = XCK
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockXCKEntryDetailCredit creates a XCK EntryDetail with a credit entry
func mockXCKEntryDetailCredit() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetCheckSerialNumber("123456789")
	entry.SetProcessControlField("CHECK1")
	entry.SetItemResearchNumber("182726")
	entry.SetTraceNumber(mockBatchXCKHeader().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

// mockBatchXCKCredit creates a BatchXCK with a Credit entry
func mockBatchXCKCredit() *BatchXCK {
	mockBatch := NewBatchXCK(mockBatchXCKHeaderCredit())
	mockBatch.AddEntry(mockXCKEntryDetailCredit())
	return mockBatch
}

// testBatchXCKHeader creates a BatchXCK BatchHeader
func testBatchXCKHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchXCKHeader())
	err, ok := batch.(*BatchXCK)
	if !ok {
		t.Errorf("Expecting BatchXCK got %T", err)
	}
}

// TestBatchXCKHeader tests validating BatchXCK BatchHeader
func TestBatchXCKHeader(t *testing.T) {
	testBatchXCKHeader(t)
}

// BenchmarkBatchXCKHeader benchmarks validating BatchXCK BatchHeader
func BenchmarkBatchXCKHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchXCKHeader(b)
	}
}

// testBatchXCKCreate validates BatchXCK create
func testBatchXCKCreate(t testing.TB) {
	mockBatch := mockBatchXCK()
	if err := mockBatch.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchXCKCreate tests validating BatchXCK create
func TestBatchXCKCreate(t *testing.T) {
	testBatchXCKCreate(t)
}

// BenchmarkBatchXCKCreate benchmarks validating BatchXCK create
func BenchmarkBatchXCKCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchXCKCreate(b)
	}
}

// testBatchXCKStandardEntryClassCode validates BatchXCK create for an invalid StandardEntryClassCode
func testBatchXCKStandardEntryClassCode(t testing.TB) {
	mockBatch := mockBatchXCK()
	mockBatch.Header.StandardEntryClassCode = WEB
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchSECType) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchXCKStandardEntryClassCode tests validating BatchXCK create for an invalid StandardEntryClassCode
func TestBatchXCKStandardEntryClassCode(t *testing.T) {
	testBatchXCKStandardEntryClassCode(t)
}

// BenchmarkBatchXCKStandardEntryClassCode benchmarks validating BatchXCK create for an invalid StandardEntryClassCode
func BenchmarkBatchXCKStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchXCKStandardEntryClassCode(b)
	}
}

// testBatchXCKServiceClassCodeEquality validates service class code equality
func testBatchXCKServiceClassCodeEquality(t testing.TB) {
	mockBatch := mockBatchXCK()
	mockBatch.GetControl().ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(220, MixedDebitsAndCredits)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchXCKServiceClassCodeEquality tests validating service class code equality
func TestBatchXCKServiceClassCodeEquality(t *testing.T) {
	testBatchXCKServiceClassCodeEquality(t)
}

// BenchmarkBatchXCKServiceClassCodeEquality benchmarks validating service class code equality
func BenchmarkBatchXCKServiceClassCodeEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchXCKServiceClassCodeEquality(b)
	}
}

// testBatchXCKMixedCreditsAndDebits validates BatchXCK create for an invalid MixedCreditsAndDebits
func testBatchXCKMixedCreditsAndDebits(t testing.TB) {
	mockBatch := mockBatchXCK()
	mockBatch.Header.ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(MixedDebitsAndCredits, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchXCKMixedCreditsAndDebits tests validating BatchXCK create for an invalid MixedCreditsAndDebits
func TestBatchXCKMixedCreditsAndDebits(t *testing.T) {
	testBatchXCKMixedCreditsAndDebits(t)
}

// BenchmarkBatchXCKMixedCreditsAndDebits benchmarks validating BatchXCK create for an invalid MixedCreditsAndDebits
func BenchmarkBatchXCKMixedCreditsAndDebits(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchXCKMixedCreditsAndDebits(b)
	}
}

// testBatchXCKCreditsOnly validates BatchXCK create for an invalid CreditsOnly
func testBatchXCKCreditsOnly(t testing.TB) {
	mockBatch := mockBatchXCK()
	mockBatch.Header.ServiceClassCode = CreditsOnly
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(CreditsOnly, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchXCKCreditsOnly tests validating BatchXCK create for an invalid CreditsOnly
func TestBatchXCKCreditsOnly(t *testing.T) {
	testBatchXCKCreditsOnly(t)
}

// BenchmarkBatchXCKCreditsOnly benchmarks validating BatchXCK create for an invalid CreditsOnly
func BenchmarkBatchXCKCreditsOnly(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchXCKCreditsOnly(b)
	}
}

// testBatchXCKAutomatedAccountingAdvices validates BatchXCK create for an invalid AutomatedAccountingAdvices
func testBatchXCKAutomatedAccountingAdvices(t testing.TB) {
	mockBatch := mockBatchXCK()
	mockBatch.Header.ServiceClassCode = AutomatedAccountingAdvices
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(AutomatedAccountingAdvices, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchXCKAutomatedAccountingAdvices tests validating BatchXCK create for an invalid AutomatedAccountingAdvices
func TestBatchXCKAutomatedAccountingAdvices(t *testing.T) {
	testBatchXCKAutomatedAccountingAdvices(t)
}

// BenchmarkBatchXCKAutomatedAccountingAdvices benchmarks validating BatchXCK create for an invalid AutomatedAccountingAdvices
func BenchmarkBatchXCKAutomatedAccountingAdvices(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchXCKAutomatedAccountingAdvices(b)
	}
}

// testBatchXCKCheckSerialNumber validates BatchXCK CheckSerialNumber is not mandatory
func testBatchXCKCheckSerialNumber(t testing.TB) {
	mockBatch := mockBatchXCK()
	// modify CheckSerialNumber / IdentificationNumber to nothing
	mockBatch.GetEntries()[0].SetCheckSerialNumber("")
	err := mockBatch.Validate()
	// no error expected
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchXCKCheckSerialNumber  tests validating BatchXCK
// CheckSerialNumber / IdentificationNumber is a mandatory field
func TestBatchXCKCheckSerialNumber(t *testing.T) {
	testBatchXCKCheckSerialNumber(t)
}

// BenchmarkBatchXCKCheckSerialNumber benchmarks validating BatchXCK
// CheckSerialNumber / IdentificationNumber is a mandatory field
func BenchmarkBatchXCKCheckSerialNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchXCKCheckSerialNumber(b)
	}
}

// testBatchXCKTransactionCode validates BatchXCK TransactionCode is not a credit
func testBatchXCKTransactionCode(t testing.TB) {
	mockBatch := mockBatchXCKCredit()
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchDebitOnly) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchXCKTransactionCode tests validating BatchXCK TransactionCode is not a credit
func TestBatchXCKTransactionCode(t *testing.T) {
	testBatchXCKTransactionCode(t)
}

// BenchmarkBatchXCKTransactionCode benchmarks validating BatchXCK TransactionCode is not a credit
func BenchmarkBatchXCKTransactionCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchXCKTransactionCode(b)
	}
}

// testBatchXCKAddendaCount validates BatchXCK Addenda count
func testBatchXCKAddendaCount(t testing.TB) {
	mockBatch := mockBatchXCK()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAddendaCategory) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchXCKAddendaCount tests validating BatchXCK Addenda count
func TestBatchXCKAddendaCount(t *testing.T) {
	testBatchXCKAddendaCount(t)
}

// BenchmarkBatchXCKAddendaCount benchmarks validating BatchXCK Addenda count
func BenchmarkBatchXCKAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchXCKAddendaCount(b)
	}
}

// testBatchXCKInvalidBuild validates an invalid batch build
func testBatchXCKInvalidBuild(t testing.TB) {
	mockBatch := mockBatchXCK()
	mockBatch.GetHeader().ServiceClassCode = 3
	err := mockBatch.Create()
	if !base.Match(err, ErrServiceClass) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchXCKInvalidBuild tests validating an invalid batch build
func TestBatchXCKInvalidBuild(t *testing.T) {
	testBatchXCKInvalidBuild(t)
}

// BenchmarkBatchXCKInvalidBuild benchmarks validating an invalid batch build
func BenchmarkBatchXCKInvalidBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchXCKInvalidBuild(b)
	}
}

// TestBatchXCKAddendum98 validates Addenda98 returns an error
func TestBatchXCKAddendum98(t *testing.T) {
	mockBatch := NewBatchXCK(mockBatchXCKHeader())
	mockBatch.AddEntry(mockXCKEntryDetail())
	mockAddenda98 := mockAddenda98()
	mockAddenda98.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchXCKAddendum99 validates Addenda99 returns an error
func TestBatchXCKAddendum99(t *testing.T) {
	mockBatch := NewBatchXCK(mockBatchXCKHeader())
	mockBatch.AddEntry(mockXCKEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockAddenda99.TypeCode = "05"
	mockBatch.GetEntries()[0].Addenda99 = mockAddenda99
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchXCKAddendum99Category validates Addenda99 returns an error
func TestBatchXCKAddendum99Category(t *testing.T) {
	mockBatch := NewBatchXCK(mockBatchXCKHeader())
	mockBatch.AddEntry(mockXCKEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].Category = CategoryForward
	mockBatch.GetEntries()[0].Addenda99 = mockAddenda99
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAddendaCategory) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchXCKProcessControlField returns an error if ProcessControlField is not defined.
func TestBatchXCKProcessControlField(t *testing.T) {
	mockBatch := NewBatchXCK(mockBatchXCKHeader())
	mockBatch.AddEntry(mockXCKEntryDetail())
	mockBatch.GetEntries()[0].SetProcessControlField("")
	err := mockBatch.Create()
	if !base.Match(err, ErrFieldRequired) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchXCKItemResearchNumber returns an error if ItemResearchNumber is not defined.
func TestBatchXCKItemResearchNumber(t *testing.T) {
	mockBatch := NewBatchXCK(mockBatchXCKHeader())
	mockBatch.AddEntry(mockXCKEntryDetail())
	mockBatch.GetEntries()[0].IndividualName = ""
	mockBatch.GetEntries()[0].SetProcessControlField("CHECK1")
	mockBatch.GetEntries()[0].SetItemResearchNumber("")
	err := mockBatch.Create()
	if !base.Match(err, ErrFieldRequired) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchXCKAmount validates BatchXCK create for an invalid Amount
func TestBatchXCKAmount(t *testing.T) {
	mockBatch := mockBatchXCK()
	mockBatch.Entries[0].Amount = 260000
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchAmount(260000, 250000)) {
		t.Errorf("%T: %s", err, err)
	}
}

// testBatchXCKMixedDebitsAndCreditsServiceClassCode validates MixedDebitsAndCredits service class code
func testBatchXCKMixedDebitsAndCreditsServiceClassCode(t testing.TB) {
	mockBatch := mockBatchXCK()
	mockBatch.GetControl().ServiceClassCode = MixedDebitsAndCredits
	mockBatch.Header.ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchXCKMixedDebitsAndCreditsServiceClassCode tests validates MixedDebitsAndCredits service class code
func TestBatchXCKMixedDebitsAndCreditsServiceClassCode(t *testing.T) {
	testBatchXCKMixedDebitsAndCreditsServiceClassCode(t)
}
