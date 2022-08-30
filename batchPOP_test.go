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

// mockBatchPOPHeader creates a BatchPOP BatchHeader
func mockBatchPOPHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = POP
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "Point of Purchase"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockPOPEntryDetail creates a BatchPOP EntryDetail
func mockPOPEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingDebit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetPOPCheckSerialNumber("123456789")
	entry.SetPOPTerminalCity("PHIL")
	entry.SetPOPTerminalState("PA")
	entry.SetReceivingCompany("ABC Company")
	entry.SetTraceNumber(mockBatchPOPHeader().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

// mockBatchPOP creates a BatchPOP
func mockBatchPOP() *BatchPOP {
	mockBatch := NewBatchPOP(mockBatchPOPHeader())
	mockBatch.AddEntry(mockPOPEntryDetail())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// mockBatchPOPHeaderCredit creates a BatchPOP BatchHeader
func mockBatchPOPHeaderCredit() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = POP
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = POP
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockPOPEntryDetailCredit creates a POP EntryDetail with a credit entry
func mockPOPEntryDetailCredit() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetPOPCheckSerialNumber("123456789")
	entry.SetPOPTerminalCity("PHIL")
	entry.SetPOPTerminalState("PA")
	entry.SetReceivingCompany("ABC Company")
	entry.SetTraceNumber(mockBatchPOPHeader().ODFIIdentification, 123)
	entry.Category = CategoryForward
	return entry
}

// mockBatchPOPCredit creates a BatchPOP with a Credit entry
func mockBatchPOPCredit() *BatchPOP {
	mockBatch := NewBatchPOP(mockBatchPOPHeaderCredit())
	mockBatch.AddEntry(mockPOPEntryDetailCredit())
	return mockBatch
}

// testBatchPOPHeader creates a BatchPOP BatchHeader
func testBatchPOPHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchPOPHeader())
	err, ok := batch.(*BatchPOP)
	if !ok {
		t.Errorf("Expecting BatchPOP got %T", err)
	}
}

// TestBatchPOPHeader tests validating BatchPOP BatchHeader
func TestBatchPOPHeader(t *testing.T) {
	testBatchPOPHeader(t)
}

// BenchmarkBatchPOPHeader benchmarks validating BatchPOP BatchHeader
func BenchmarkBatchPOPHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPHeader(b)
	}
}

// testBatchPOPCreate validates BatchPOP create
func testBatchPOPCreate(t testing.TB) {
	mockBatch := mockBatchPOP()
	if err := mockBatch.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchPOPCreate tests validating BatchPOP create
func TestBatchPOPCreate(t *testing.T) {
	testBatchPOPCreate(t)
}

// BenchmarkBatchPOPCreate benchmarks validating BatchPOP create
func BenchmarkBatchPOPCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPCreate(b)
	}
}

// testBatchPOPStandardEntryClassCode validates BatchPOP create for an invalid StandardEntryClassCode
func testBatchPOPStandardEntryClassCode(t testing.TB) {
	mockBatch := mockBatchPOP()
	mockBatch.Header.StandardEntryClassCode = WEB
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchSECType) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchPOPStandardEntryClassCode tests validating BatchPOP create for an invalid StandardEntryClassCode
func TestBatchPOPStandardEntryClassCode(t *testing.T) {
	testBatchPOPStandardEntryClassCode(t)
}

// BenchmarkBatchPOPStandardEntryClassCode benchmarks validating BatchPOP create for an invalid StandardEntryClassCode
func BenchmarkBatchPOPStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPStandardEntryClassCode(b)
	}
}

// testBatchPOPServiceClassCodeEquality validates service class code equality
func testBatchPOPServiceClassCodeEquality(t testing.TB) {
	mockBatch := mockBatchPOP()
	mockBatch.GetControl().ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(220, MixedDebitsAndCredits)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchPOPServiceClassCodeEquality tests validating service class code equality
func TestBatchPOPServiceClassCodeEquality(t *testing.T) {
	testBatchPOPServiceClassCodeEquality(t)
}

// BenchmarkBatchPOPServiceClassCodeEquality benchmarks validating service class code equality
func BenchmarkBatchPOPServiceClassCodeEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPServiceClassCodeEquality(b)
	}
}

// testBatchPOPMixedCreditsAndDebits validates BatchPOP create for an invalid MixedCreditsAndDebits
func testBatchPOPMixedCreditsAndDebits(t testing.TB) {
	mockBatch := mockBatchPOP()
	mockBatch.Header.ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(MixedDebitsAndCredits, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchPOPMixedCreditsAndDebits tests validating BatchPOP create for an invalid MixedCreditsAndDebits
func TestBatchPOPMixedCreditsAndDebits(t *testing.T) {
	testBatchPOPMixedCreditsAndDebits(t)
}

// BenchmarkBatchPOPMixedCreditsAndDebits benchmarks validating BatchPOP create for an invalid MixedCreditsAndDebits
func BenchmarkBatchPOPMixedCreditsAndDebits(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPMixedCreditsAndDebits(b)
	}
}

// testBatchPOPCreditsOnly validates BatchPOP create for an invalid CreditsOnly
func testBatchPOPCreditsOnly(t testing.TB) {
	mockBatch := mockBatchPOP()
	mockBatch.Header.ServiceClassCode = CreditsOnly
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(CreditsOnly, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchPOPCreditsOnly tests validating BatchPOP create for an invalid CreditsOnly
func TestBatchPOPCreditsOnly(t *testing.T) {
	testBatchPOPCreditsOnly(t)
}

// BenchmarkBatchPOPCreditsOnly benchmarks validating BatchPOP create for an invalid CreditsOnly
func BenchmarkBatchPOPCreditsOnly(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPCreditsOnly(b)
	}
}

// testBatchPOPAutomatedAccountingAdvices validates BatchPOP create for an invalid AutomatedAccountingAdvices
func testBatchPOPAutomatedAccountingAdvices(t testing.TB) {
	mockBatch := mockBatchPOP()
	mockBatch.Header.ServiceClassCode = AutomatedAccountingAdvices
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(MixedDebitsAndCredits, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchPOPAutomatedAccountingAdvices tests validating BatchPOP create for an invalid AutomatedAccountingAdvices
func TestBatchPOPAutomatedAccountingAdvices(t *testing.T) {
	testBatchPOPAutomatedAccountingAdvices(t)
}

// BenchmarkBatchPOPAutomatedAccountingAdvices benchmarks validating BatchPOP create for an invalid AutomatedAccountingAdvices
func BenchmarkBatchPOPAutomatedAccountingAdvices(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPAutomatedAccountingAdvices(b)
	}
}

// testBatchPOPAmount validates BatchPOP create for an invalid Amount
func testBatchPOPAmount(t testing.TB) {
	mockBatch := mockBatchPOP()
	mockBatch.Entries[0].Amount = 2600000
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchAmount(2600000, 2500000)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchPOPAmount validates BatchPOP create for an invalid Amount
func TestBatchPOPAmount(t *testing.T) {
	testBatchPOPAmount(t)
}

// BenchmarkBatchPOPAmount validates BatchPOP create for an invalid Amount
func BenchmarkBatchPOPAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPAmount(b)
	}
}

// testBatchPOPCheckSerialNumber validates BatchPOP CheckSerialNumber / IdentificationNumber is a mandatory field
func testBatchPOPCheckSerialNumber(t testing.TB) {
	mockBatch := mockBatchPOP()
	// modify CheckSerialNumber / IdentificationNumber to nothing
	mockBatch.GetEntries()[0].SetCheckSerialNumber("")
	err := mockBatch.Validate()
	if !base.Match(err, ErrBatchCheckSerialNumber) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchPOPCheckSerialNumber tests validating BatchPOP
// CheckSerialNumber / IdentificationNumber is a mandatory field
func TestBatchPOPCheckSerialNumber(t *testing.T) {
	testBatchPOPCheckSerialNumber(t)
}

// BenchmarkBatchPOPCheckSerialNumber benchmarks validating BatchPOP
// CheckSerialNumber / IdentificationNumber is a mandatory field
func BenchmarkBatchPOPCheckSerialNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPCheckSerialNumber(b)
	}
}

// testBatchPOPCheckSerialNumberField validates POPCheckSerialNumberField characters 1-9 of underlying BatchPOP
// CheckSerialNumber / IdentificationNumber
func testBatchPOPCheckSerialNumberField(t testing.TB) {
	mockBatch := mockBatchPOP()
	tc := mockBatch.Entries[0].POPCheckSerialNumberField()
	if tc != "123456789" {
		t.Error("CheckSerialNumber is invalid")
	}
}

// TestBatchPPOPCheckSerialNumberField tests validating POPCheckSerialNumberField characters 1-9 of underlying BatchPOP
// CheckSerialNumber / IdentificationNumber
func TestBatchPOPCheckSerialNumberField(t *testing.T) {
	testBatchPOPCheckSerialNumberField(t)
}

// BenchmarkBatchPOPCheckSerialNumberField benchmarks validating POPCheckSerialNumberField characters 1-9 of underlying
// BatchPOP CheckSerialNumber / IdentificationNumber
func BenchmarkBatchPOPCheckSerialNumberField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPTerminalCityField(b)
	}
}

// testBatchPOPTerminalCityField validates POPTerminalCity characters 10-13 of underlying BatchPOP
// CheckSerialNumber / IdentificationNumber
func testBatchPOPTerminalCityField(t testing.TB) {
	mockBatch := mockBatchPOP()
	tc := mockBatch.Entries[0].POPTerminalCityField()
	if tc != "PHIL" {
		t.Error("TerminalCity is invalid")
	}
}

// TestBatchPOPTerminalCityField tests validating POPTerminalCity characters 10-13 of underlying BatchPOP
// CheckSerialNumber / IdentificationNumber
func TestBatchPOPTerminalCityField(t *testing.T) {
	testBatchPOPTerminalCityField(t)
}

// BenchmarkBatchPOPTerminalCityField benchmarks validating POPTerminalCity characters 10-13 of underlying
// BatchPOP CheckSerialNumber / IdentificationNumber
func BenchmarkBatchPOPTerminalCityField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPTerminalCityField(b)
	}
}

// testBatchPOPTerminalStateField validates POPTerminalState characters 14-15 of underlying BatchPOP
// CheckSerialNumber / IdentificationNumber
func testBatchPOPTerminalStateField(t testing.TB) {
	mockBatch := mockBatchPOP()
	ts := mockBatch.Entries[0].POPTerminalStateField()
	if ts != "PA" {
		t.Error("TerminalState is invalid")
	}
}

// TestBatchPOPTerminalStateField tests validating POPTerminalState characters 14-15 of underlying BatchPOP
// CheckSerialNumber / IdentificationNumber
func TestBatchPOPTerminalStateField(t *testing.T) {
	testBatchPOPTerminalStateField(t)
}

// BenchmarkBatchPOPTerminalStateField benchmarks validating POPTerminalState characters 14-15 of underlying
// BatchPOP CheckSerialNumber / IdentificationNumber
func BenchmarkBatchPOPTerminalStateField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPTerminalStateField(b)
	}
}

// testBatchPOPTransactionCode validates BatchPOP TransactionCode is not a credit
func testBatchPOPTransactionCode(t testing.TB) {
	mockBatch := mockBatchPOPCredit()
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchDebitOnly) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchPOPTransactionCode tests validating BatchPOP TransactionCode is not a credit
func TestBatchPOPTransactionCode(t *testing.T) {
	testBatchPOPTransactionCode(t)
}

// BenchmarkBatchPOPTransactionCode benchmarks validating BatchPOP TransactionCode is not a credit
func BenchmarkBatchPOPTransactionCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPTransactionCode(b)
	}
}

// testBatchPOPAddendaCount validates BatchPOP Addenda count
func testBatchPOPAddendaCount(t testing.TB) {
	mockBatch := mockBatchPOP()
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAddendaCategory) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchPOPAddendaCount tests validating BatchPOP Addenda count
func TestBatchPOPAddendaCount(t *testing.T) {
	testBatchPOPAddendaCount(t)
}

// BenchmarkBatchPOPAddendaCount benchmarks validating BatchPOP Addenda count
func BenchmarkBatchPOPAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCAddendaCount(b)
	}
}

// testBatchPOPInvalidBuild validates an invalid batch build
func testBatchPOPInvalidBuild(t testing.TB) {
	mockBatch := mockBatchPOP()
	mockBatch.GetHeader().ServiceClassCode = 3
	err := mockBatch.Create()
	if !base.Match(err, ErrServiceClass) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchPOPInvalidBuild tests validating an invalid batch build
func TestBatchPOPInvalidBuild(t *testing.T) {
	testBatchPOPInvalidBuild(t)
}

// BenchmarkBatchPOPInvalidBuild benchmarks validating an invalid batch build
func BenchmarkBatchPOPInvalidBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPInvalidBuild(b)
	}
}

// TestBatchPOPAddendum98 validates Addenda98 returns an error
func TestBatchPOPAddendum98(t *testing.T) {
	mockBatch := NewBatchPOP(mockBatchPOPHeader())
	mockBatch.AddEntry(mockPOPEntryDetail())
	mockAddenda98 := mockAddenda98()
	mockAddenda98.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchPOPAddendum99 validates Addenda99 returns an error
func TestBatchPOPAddendum99(t *testing.T) {
	mockBatch := NewBatchPOP(mockBatchPOPHeader())
	mockBatch.AddEntry(mockPOPEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockAddenda99.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryReturn
	mockBatch.GetEntries()[0].Addenda99 = mockAddenda99
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

func testBatchPOPMixedDebitsAndCredits(t testing.TB) {
	mockBatch := mockBatchPOP()
	mockBatch.Header.ServiceClassCode = MixedDebitsAndCredits
	mockBatch.Batch.Control.ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchPOPMixedDebitsAndCredits tests validating BatchPOP create for MixedDebitsAndCredits with debit transaction code
func TestBatchPOPMixedDebitsAndCredits(t *testing.T) {
	testBatchPOPMixedDebitsAndCredits(t)
}
