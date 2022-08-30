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

// mockBatchARCHeader creates a BatchARC BatchHeader
func mockBatchARCHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = ARC
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = ARC
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockARCEntryDetail creates a BatchARC EntryDetail
func mockARCEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingDebit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetCheckSerialNumber("123456789")
	entry.SetReceivingCompany("ABC Company")
	entry.SetTraceNumber(mockBatchARCHeader().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

// mockBatchARC creates a BatchARC
func mockBatchARC() *BatchARC {
	mockBatch := NewBatchARC(mockBatchARCHeader())
	mockBatch.AddEntry(mockARCEntryDetail())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// mockBatchARCHeaderCredit creates a BatchARC BatchHeader
func mockBatchARCHeaderCredit() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = ARC
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = ARC
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockARCEntryDetailCredit creates a ARC EntryDetail with a credit entry
func mockARCEntryDetailCredit() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetCheckSerialNumber("123456789")
	entry.SetReceivingCompany("ABC Company")
	entry.SetTraceNumber(mockBatchARCHeader().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

// mockBatchARCCredit creates a BatchARC with a Credit entry
func mockBatchARCCredit() *BatchARC {
	mockBatch := NewBatchARC(mockBatchARCHeaderCredit())
	mockBatch.AddEntry(mockARCEntryDetailCredit())
	return mockBatch
}

// testBatchARCHeader creates a BatchARC BatchHeader
func testBatchARCHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchARCHeader())
	err, ok := batch.(*BatchARC)
	if !ok {
		t.Errorf("Expecting BatchARC got %T", err)
	}
}

// TestBatchARCHeader tests validating BatchARC BatchHeader
func TestBatchARCHeader(t *testing.T) {
	testBatchARCHeader(t)
}

// BenchmarkBatchARCHeader benchmarks validating BatchARC BatchHeader
func BenchmarkBatchARCHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCHeader(b)
	}
}

// testBatchARCCreate validates BatchARC create
func testBatchARCCreate(t testing.TB) {
	mockBatch := mockBatchARC()
	if err := mockBatch.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchARCCreate tests validating BatchARC create
func TestBatchARCCreate(t *testing.T) {
	testBatchARCCreate(t)
}

// BenchmarkBatchARCCreate benchmarks validating BatchARC create
func BenchmarkBatchARCCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCCreate(b)
	}
}

// testBatchARCStandardEntryClassCode validates BatchARC create for an invalid StandardEntryClassCode
func testBatchARCStandardEntryClassCode(t testing.TB) {
	mockBatch := mockBatchARC()
	mockBatch.Header.StandardEntryClassCode = WEB
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchSECType) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchARCStandardEntryClassCode tests validating BatchARC create for an invalid StandardEntryClassCode
func TestBatchARCStandardEntryClassCode(t *testing.T) {
	testBatchARCStandardEntryClassCode(t)
}

// BenchmarkBatchARCStandardEntryClassCode benchmarks validating BatchARC create for an invalid StandardEntryClassCode
func BenchmarkBatchARCStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCStandardEntryClassCode(b)
	}
}

// testBatchARCServiceClassCodeEquality validates service class code equality
func testBatchARCServiceClassCodeEquality(t testing.TB) {
	mockBatch := mockBatchARC()
	mockBatch.GetControl().ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(220, MixedDebitsAndCredits)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchARCServiceClassCodeEquality tests validating service class code equality
func TestBatchARCServiceClassCodeEquality(t *testing.T) {
	testBatchARCServiceClassCodeEquality(t)
}

// BenchmarkBatchARCServiceClassCodeEquality benchmarks validating service class code equality
func BenchmarkBatchARCServiceClassCodeEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCServiceClassCodeEquality(b)
	}
}

// testBatchARCMixedCreditsAndDebits validates BatchARC create for an invalid MixedCreditsAndDebits
func testBatchARCMixedCreditsAndDebits(t testing.TB) {
	mockBatch := mockBatchARC()
	mockBatch.Header.ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(MixedDebitsAndCredits, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchARCMixedCreditsAndDebits tests validating BatchARC create for an invalid MixedCreditsAndDebits
func TestBatchARCMixedCreditsAndDebits(t *testing.T) {
	testBatchARCMixedCreditsAndDebits(t)
}

// BenchmarkBatchARCMixedCreditsAndDebits benchmarks validating BatchARC create for an invalid MixedCreditsAndDebits
func BenchmarkBatchARCMixedCreditsAndDebits(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCMixedCreditsAndDebits(b)
	}
}

// testBatchARCCreditsOnly validates BatchARC create for an invalid CreditsOnly
func testBatchARCCreditsOnly(t testing.TB) {
	mockBatch := mockBatchARC()
	mockBatch.Header.ServiceClassCode = CreditsOnly
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(CreditsOnly, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchARCCreditsOnly tests validating BatchARC create for an invalid CreditsOnly
func TestBatchARCCreditsOnly(t *testing.T) {
	testBatchARCCreditsOnly(t)
}

// BenchmarkBatchARCCreditsOnly benchmarks validating BatchARC create for an invalid CreditsOnly
func BenchmarkBatchARCCreditsOnly(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCCreditsOnly(b)
	}
}

// testBatchARCAutomatedAccountingAdvices validates BatchARC create for an invalid AutomatedAccountingAdvices
func testBatchARCAutomatedAccountingAdvices(t testing.TB) {
	mockBatch := mockBatchARC()
	mockBatch.Header.ServiceClassCode = AutomatedAccountingAdvices
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(AutomatedAccountingAdvices, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchARCAutomatedAccountingAdvices tests validating BatchARC create for an invalid AutomatedAccountingAdvices
func TestBatchARCAutomatedAccountingAdvices(t *testing.T) {
	testBatchARCAutomatedAccountingAdvices(t)
}

// BenchmarkBatchARCAutomatedAccountingAdvices benchmarks validating BatchARC create for an invalid AutomatedAccountingAdvices
func BenchmarkBatchARCAutomatedAccountingAdvices(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCAutomatedAccountingAdvices(b)
	}
}

// testBatchARCAmount validates BatchARC create for an invalid Amount
func testBatchARCAmount(t testing.TB) {
	mockBatch := mockBatchARC()
	mockBatch.Entries[0].Amount = 2600000
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchAmount(2600000, 2500000)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchARCAmount validates BatchARC create for an invalid Amount
func TestBatchARCAmount(t *testing.T) {
	testBatchARCAmount(t)
}

// BenchmarkBatchARCAmount validates BatchARC create for an invalid Amount
func BenchmarkBatchARCAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCAmount(b)
	}
}

// testBatchARCCheckSerialNumber validates BatchARC CheckSerialNumber / IdentificationNumber is a mandatory field
func testBatchARCCheckSerialNumber(t testing.TB) {
	mockBatch := mockBatchARC()
	// modify CheckSerialNumber / IdentificationNumber to nothing
	mockBatch.GetEntries()[0].SetCheckSerialNumber("")
	err := mockBatch.Validate()
	if !base.Match(err, ErrBatchCheckSerialNumber) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchARCCheckSerialNumber  tests validating BatchARC
// CheckSerialNumber / IdentificationNumber is a mandatory field
func TestBatchARCCheckSerialNumber(t *testing.T) {
	testBatchARCCheckSerialNumber(t)
}

// BenchmarkBatchARCCheckSerialNumber benchmarks validating BatchARC
// CheckSerialNumber / IdentificationNumber is a mandatory field
func BenchmarkBatchARCCheckSerialNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCCheckSerialNumber(b)
	}
}

// testBatchARCTransactionCode validates BatchARC TransactionCode is not a credit
func testBatchARCTransactionCode(t testing.TB) {
	mockBatch := mockBatchARCCredit()
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchDebitOnly) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchARCTransactionCode tests validating BatchARC TransactionCode is not a credit
func TestBatchARCTransactionCode(t *testing.T) {
	testBatchARCTransactionCode(t)
}

// BenchmarkBatchARCTransactionCode benchmarks validating BatchARC TransactionCode is not a credit
func BenchmarkBatchARCTransactionCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCTransactionCode(b)
	}
}

// testBatchARCAddendaCount validates BatchARC Addenda count
func testBatchARCAddendaCount(t testing.TB) {
	mockBatch := mockBatchARC()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAddendaCategory) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchARCAddendaCount tests validating BatchARC Addenda count
func TestBatchARCAddendaCount(t *testing.T) {
	testBatchARCAddendaCount(t)
}

// BenchmarkBatchARCAddendaCount benchmarks validating BatchARC Addenda count
func BenchmarkBatchARCAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCAddendaCount(b)
	}
}

// testBatchARCInvalidBuild validates an invalid batch build
func testBatchARCInvalidBuild(t testing.TB) {
	mockBatch := mockBatchARC()
	mockBatch.GetHeader().ServiceClassCode = 3
	err := mockBatch.Create()
	if !base.Match(err, ErrServiceClass) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchARCInvalidBuild tests validating an invalid batch build
func TestBatchARCInvalidBuild(t *testing.T) {
	testBatchARCInvalidBuild(t)
}

// BenchmarkBatchARCInvalidBuild benchmarks validating an invalid batch build
func BenchmarkBatchARCInvalidBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCInvalidBuild(b)
	}
}

// TestBatchARCAddendum98 validates Addenda98 returns an error
func TestBatchARCAddendum98(t *testing.T) {
	mockBatch := NewBatchARC(mockBatchARCHeader())
	mockBatch.AddEntry(mockARCEntryDetail())
	mockAddenda98 := mockAddenda98()
	mockAddenda98.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchARCAddendum99 validates Addenda99 returns an error
func TestBatchARCAddendum99(t *testing.T) {
	mockBatch := NewBatchARC(mockBatchARCHeader())
	mockBatch.AddEntry(mockARCEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockAddenda99.TypeCode = "05"
	mockBatch.GetEntries()[0].Addenda99 = mockAddenda99
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchARCAddendum99Category validates Addenda99 returns an error
func TestBatchARCAddendum99Category(t *testing.T) {
	mockBatch := NewBatchARC(mockBatchARCHeader())
	mockBatch.AddEntry(mockARCEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].Category = CategoryForward
	mockBatch.GetEntries()[0].Addenda99 = mockAddenda99
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAddendaCategory) {
		t.Errorf("%T: %s", err, err)
	}
}

// testBatchARCMixedCreditsAndDebits validates BatchARC create for valid MixedCreditsAndDebits
func testBatchARCMixedCreditsAndDebitsBatchControlMixedDebitsAndCredits(t testing.TB) {
	mockBatch := mockBatchARC()
	mockBatch.Header.ServiceClassCode = MixedDebitsAndCredits
	mockBatch.Batch.Control.ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchARCMixedCreditsAndDebitsBatchControlMixedDebitsAndCredits tests validating BatchARC create for valid MixedCreditsAndDebits
func TestBatchARCMixedCreditsAndDebitsBatchControlMixedDebitsAndCredits(t *testing.T) {
	testBatchARCMixedCreditsAndDebitsBatchControlMixedDebitsAndCredits(t)
}
