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

// mockBatchBOCHeader creates a BatchBOC BatchHeader
func mockBatchBOCHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = BOC
	bh.CompanyName = "Company Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = BOC
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockBOCEntryDetail creates a BatchBOC EntryDetail
func mockBOCEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingDebit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetCheckSerialNumber("123456789")
	entry.SetReceivingCompany("ABC Company")
	entry.SetTraceNumber(mockBatchBOCHeader().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

// mockBatchBOC creates a BatchBOC
func mockBatchBOC() *BatchBOC {
	mockBatch := NewBatchBOC(mockBatchBOCHeader())
	mockBatch.AddEntry(mockBOCEntryDetail())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// mockBatchBOCHeaderCredit creates a BatchBOC BatchHeader
func mockBatchBOCHeaderCredit() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = BOC
	bh.CompanyName = "Company Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "REDEPCHECK"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockBOCEntryDetailCredit creates a BatchBOC EntryDetail with a credit
func mockBOCEntryDetailCredit() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetCheckSerialNumber("123456789")
	entry.SetReceivingCompany("ABC Company")
	entry.SetTraceNumber(mockBatchBOCHeader().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

// mockBatchBOCCredit creates a BatchBOC with a Credit entry
func mockBatchBOCCredit() *BatchBOC {
	mockBatch := NewBatchBOC(mockBatchBOCHeaderCredit())
	mockBatch.AddEntry(mockBOCEntryDetailCredit())
	return mockBatch
}

// testBatchBOCHeader creates a BatchBOC BatchHeader
func testBatchBOCHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchBOCHeader())
	err, ok := batch.(*BatchBOC)
	if !ok {
		t.Errorf("Expecting BatchBOC got %T", err)
	}
}

// TestBatchBOCHeader tests validating BatchBOC BatchHeader
func TestBatchBOCHeader(t *testing.T) {
	testBatchBOCHeader(t)
}

// BenchmarkBatchBOCHeader benchmarks validating BatchBOC BatchHeader
func BenchmarkBatchBOCHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCHeader(b)
	}
}

// testBatchBOCCreate validates BatchBOC create
func testBatchBOCCreate(t testing.TB) {
	mockBatch := mockBatchBOC()
	if err := mockBatch.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchBOCCreate tests validating BatchBOC create
func TestBatchBOCCreate(t *testing.T) {
	testBatchBOCCreate(t)
}

// BenchmarkBatchBOCCreate benchmarks validating BatchBOC create
func BenchmarkBatchBOCCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCCreate(b)
	}
}

// testBatchBOCStandardEntryClassCode validates BatchBOC create for an invalid StandardEntryClassCode
func testBatchBOCStandardEntryClassCode(t testing.TB) {
	mockBatch := mockBatchBOC()
	mockBatch.Header.StandardEntryClassCode = WEB
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchSECType) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchBOCStandardEntryClassCode tests validating BatchBOC create for an invalid StandardEntryClassCode
func TestBatchBOCStandardEntryClassCode(t *testing.T) {
	testBatchBOCStandardEntryClassCode(t)
}

// BenchmarkBatchBOCStandardEntryClassCode benchmarks validating BatchBOC create for an invalid StandardEntryClassCode
func BenchmarkBatchBOCStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCStandardEntryClassCode(b)
	}
}

// testBatchBOCServiceClassCodeEquality validates service class code equality
func testBatchBOCServiceClassCodeEquality(t testing.TB) {
	mockBatch := mockBatchBOC()
	mockBatch.GetControl().ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(220, MixedDebitsAndCredits)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchBOCServiceClassCodeEquality tests validating service class code equality
func TestBatchBOCServiceClassCodeEquality(t *testing.T) {
	testBatchBOCServiceClassCodeEquality(t)
}

// BenchmarkBatchBOCServiceClassCodeEquality benchmarks validating service class code equality
func BenchmarkBatchBOCServiceClassCodeEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCServiceClassCodeEquality(b)
	}
}

// testBatchBOCMixedCreditsAndDebits validates BatchBOC create for an invalid MixedCreditsAndDebits
func testBatchBOCMixedCreditsAndDebits(t testing.TB) {
	mockBatch := mockBatchBOC()
	mockBatch.Header.ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(MixedDebitsAndCredits, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchBOCMixedCreditsAndDebits tests validating BatchBOC create for an invalid MixedCreditsAndDebits
func TestBatchBOCMixedCreditsAndDebits(t *testing.T) {
	testBatchBOCMixedCreditsAndDebits(t)
}

// BenchmarkBatchBOCMixedCreditsAndDebits benchmarks validating BatchBOC create for an invalid MixedCreditsAndDebits
func BenchmarkBatchBOCMixedCreditsAndDebits(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCMixedCreditsAndDebits(b)
	}
}

// testBatchBOCCreditsOnly validates BatchBOC create for an invalid CreditsOnly
func testBatchBOCCreditsOnly(t testing.TB) {
	mockBatch := mockBatchBOC()
	mockBatch.Header.ServiceClassCode = CreditsOnly

	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(CreditsOnly, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchBOCCreditsOnly tests validating BatchBOC create for an invalid CreditsOnly
func TestBatchBOCCreditsOnly(t *testing.T) {
	testBatchBOCCreditsOnly(t)
}

// BenchmarkBatchBOCCreditsOnly benchmarks validating BatchBOC create for an invalid CreditsOnly
func BenchmarkBatchBOCCreditsOnly(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCCreditsOnly(b)
	}
}

// testBatchBOCAutomatedAccountingAdvices validates BatchBOC create for an invalid AutomatedAccountingAdvices
func testBatchBOCAutomatedAccountingAdvices(t testing.TB) {
	mockBatch := mockBatchBOC()
	mockBatch.Header.ServiceClassCode = AutomatedAccountingAdvices
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(AutomatedAccountingAdvices, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchBOCAutomatedAccountingAdvices tests validating BatchBOC create for an invalid AutomatedAccountingAdvices
func TestBatchBOCAutomatedAccountingAdvices(t *testing.T) {
	testBatchBOCAutomatedAccountingAdvices(t)
}

// BenchmarkBatchBOCAutomatedAccountingAdvices benchmarks validating BatchBOC create for an invalid AutomatedAccountingAdvices
func BenchmarkBatchBOCAutomatedAccountingAdvices(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCAutomatedAccountingAdvices(b)
	}
}

// testBatchBOCAmount validates BatchBOC create for an invalid Amount
func testBatchBOCAmount(t testing.TB) {
	mockBatch := mockBatchBOC()
	mockBatch.Entries[0].Amount = 2500001
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchAmount(2500001, 2500000)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchBOCAmount validates BatchBOC create for an invalid Amount
func TestBatchBOCAmount(t *testing.T) {
	testBatchBOCAmount(t)
}

// BenchmarkBatchBOCAmount validates BatchBOC create for an invalid Amount
func BenchmarkBatchBOCAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCAmount(b)
	}
}

// testBatchBOCCheckSerialNumber validates BatchBOC CheckSerialNumber / IdentificationNumber is a mandatory field
func testBatchBOCCheckSerialNumber(t testing.TB) {
	mockBatch := mockBatchBOC()
	// modify CheckSerialNumber / IdentificationNumber to empty string
	mockBatch.GetEntries()[0].SetCheckSerialNumber("")
	err := mockBatch.Validate()
	if !base.Match(err, ErrBatchCheckSerialNumber) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchBOCCheckSerialNumber  tests validating BatchBOC CheckSerialNumber / IdentificationNumber is a mandatory field
func TestBatchBOCCheckSerialNumber(t *testing.T) {
	testBatchBOCCheckSerialNumber(t)
}

// BenchmarkBatchBOCCheckSerialNumber benchmarks validating BatchBOC
// CheckSerialNumber / IdentificationNumber is a mandatory field
func BenchmarkBatchBOCCheckSerialNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCCheckSerialNumber(b)
	}
}

// testBatchBOCTransactionCode validates BatchBOC TransactionCode is not a credit
func testBatchBOCTransactionCode(t testing.TB) {
	mockBatch := mockBatchBOCCredit()
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchDebitOnly) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchBOCTransactionCode tests validating BatchBOC TransactionCode is not a credit
func TestBatchBOCTransactionCode(t *testing.T) {
	testBatchBOCTransactionCode(t)
}

// BenchmarkBatchBOCTransactionCode benchmarks validating BatchBOC TransactionCode is not a credit
func BenchmarkBatchBOCTransactionCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCTransactionCode(b)
	}
}

// testBatchBOCAddenda05 validates BatchBOC Addenda count
func testBatchBOCAddenda05(t testing.TB) {
	mockBatch := mockBatchBOC()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAddendaCategory) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchBOCAddenda05 tests validating BatchBOC Addenda count
func TestBatchBOCAddenda05(t *testing.T) {
	testBatchBOCAddenda05(t)
}

// BenchmarkBatchBOCAddenda05 benchmarks validating BatchBOC Addenda count
func BenchmarkBatchBOCAddenda05(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCAddenda05(b)
	}
}

// testBatchBOCInvalidBuild validates an invalid batch build
func testBatchBOCInvalidBuild(t testing.TB) {
	mockBatch := mockBatchBOC()
	mockBatch.GetHeader().ServiceClassCode = 3
	err := mockBatch.Create()
	if !base.Match(err, ErrServiceClass) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchBOCInvalidBuild tests validating an invalid batch build
func TestBatchBOCInvalidBuild(t *testing.T) {
	testBatchBOCInvalidBuild(t)
}

// BenchmarkBatchBOCInvalidBuild benchmarks validating an invalid batch build
func BenchmarkBatchBOCInvalidBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCInvalidBuild(b)
	}
}

// TestBatchBOCAddendum98 validates Addenda98 returns an error
func TestBatchBOCAddendum98(t *testing.T) {
	mockBatch := NewBatchBOC(mockBatchBOCHeader())
	mockBatch.AddEntry(mockBOCEntryDetail())
	mockAddenda98 := mockAddenda98()
	mockAddenda98.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchBOCAddendum99 validates Addenda99 returns an error
func TestBatchBOCAddendum99(t *testing.T) {
	mockBatch := NewBatchBOC(mockBatchBOCHeader())
	mockBatch.AddEntry(mockBOCEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockAddenda99.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryReturn
	mockBatch.GetEntries()[0].Addenda99 = mockAddenda99
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// testBatchBOCCreditsOnly validates BatchBOC create for Valid SCC MixedDebitsAndCredits with transCode Debit
func testBatchBOCMixedDebitsAndCreditsWithCreditTransCode(t testing.TB) {
	mockBatch := mockBatchBOC()
	mockBatch.Header.ServiceClassCode = MixedDebitsAndCredits
	mockBatch.Batch.Control.ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchBOCCreditsOnly tests validating BatchBOC create for Valid SCC MixedDebitsAndCredits with transCode Debit
func TestMixedDebitsAndCreditsWithCreditTransCode(t *testing.T) {
	testBatchBOCMixedDebitsAndCreditsWithCreditTransCode(t)
}
