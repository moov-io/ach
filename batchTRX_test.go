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
	"log"
	"testing"

	"github.com/moov-io/base"
)

// mockBatchTRXHeader creates a BatchTRX BatchHeader
func mockBatchTRXHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = TRX
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "ACH TRX"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockTRXEntryDetail creates a BatchTRX EntryDetail
func mockTRXEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingDebit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.IdentificationNumber = "45689033"
	entry.SetCATXAddendaRecords(1)
	entry.SetCATXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(mockBatchTRXHeader().ODFIIdentification, 1)
	entry.SetItemTypeIndicator("01")
	entry.Category = CategoryForward
	return entry
}

// mockBatchTRX creates a BatchTRX
func mockBatchTRX() *BatchTRX {
	mockBatch := NewBatchTRX(mockBatchTRXHeader())
	mockBatch.AddEntry(mockTRXEntryDetail())
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	if err := mockBatch.Create(); err != nil {
		log.Fatal(err)
	}
	return mockBatch
}

// mockBatchTRXHeaderCredit creates a BatchTRX BatchHeader
func mockBatchTRXHeaderCredit() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = TRX
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = TRX
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockTRXEntryDetailCredit creates a TRX EntryDetail with a credit entry
func mockTRXEntryDetailCredit() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetCheckSerialNumber("123456789")
	entry.SetProcessControlField("CHECK1")
	entry.SetItemResearchNumber("182726")
	entry.SetTraceNumber(mockBatchTRXHeader().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

// mockBatchTRXCredit creates a BatchTRX with a Credit entry
func mockBatchTRXCredit() *BatchTRX {
	mockBatch := NewBatchTRX(mockBatchTRXHeaderCredit())
	mockBatch.AddEntry(mockTRXEntryDetailCredit())
	return mockBatch
}

// testBatchTRXHeader creates a BatchTRX BatchHeader
func testBatchTRXHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchTRXHeader())
	err, ok := batch.(*BatchTRX)
	if !ok {
		t.Errorf("Expecting BatchTRX got %T", err)
	}
}

// TestBatchTRXHeader tests validating BatchTRX BatchHeader
func TestBatchTRXHeader(t *testing.T) {
	testBatchTRXHeader(t)
}

// BenchmarkBatchTRXHeader benchmarks validating BatchTRX BatchHeader
func BenchmarkBatchTRXHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRXHeader(b)
	}
}

// testBatchTRXCreate validates BatchTRX create
func testBatchTRXCreate(t testing.TB) {
	mockBatch := mockBatchTRX()
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTRXCreate tests validating BatchTRX create
func TestBatchTRXCreate(t *testing.T) {
	testBatchTRXCreate(t)
}

// BenchmarkBatchTRXCreate benchmarks validating BatchTRX create
func BenchmarkBatchTRXCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRXCreate(b)
	}
}

// testBatchTRXStandardEntryClassCode validates BatchTRX create for an invalid StandardEntryClassCode
func testBatchTRXStandardEntryClassCode(t testing.TB) {
	mockBatch := mockBatchTRX()
	mockBatch.Header.StandardEntryClassCode = WEB
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchSECType) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTRXStandardEntryClassCode tests validating BatchTRX create for an invalid StandardEntryClassCode
func TestBatchTRXStandardEntryClassCode(t *testing.T) {
	testBatchTRXStandardEntryClassCode(t)
}

// BenchmarkBatchTRXStandardEntryClassCode benchmarks validating BatchTRX create for an invalid StandardEntryClassCode
func BenchmarkBatchTRXStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRXStandardEntryClassCode(b)
	}
}

// testBatchTRXServiceClassCodeEquality validates service class code equality
func testBatchTRXServiceClassCodeEquality(t testing.TB) {
	mockBatch := mockBatchTRX()
	mockBatch.GetControl().ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(220, MixedDebitsAndCredits)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTRXServiceClassCodeEquality tests validating service class code equality
func TestBatchTRXServiceClassCodeEquality(t *testing.T) {
	testBatchTRXServiceClassCodeEquality(t)
}

// BenchmarkBatchTRXServiceClassCodeEquality benchmarks validating service class code equality
func BenchmarkBatchTRXServiceClassCodeEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRXServiceClassCodeEquality(b)
	}
}

// testBatchTRXAddendaCount validates BatchTRX Addendum count of 2
func testBatchTRXAddendaCount(t testing.TB) {
	mockBatch := mockBatchTRX()
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchExpectedAddendaCount(2, 1)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTRXAddendaCount tests validating BatchTRX Addendum count of 2
func TestBatchTRXAddendaCount(t *testing.T) {
	testBatchTRXAddendaCount(t)
}

// BenchmarkBatchTRXAddendaCount benchmarks validating BatchTRX Addendum count of 2
func BenchmarkBatchTRXAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRXAddendaCount(b)
	}
}

// testBatchTRXAddendaCountZero validates Addendum count of 0
func testBatchTRXAddendaCountZero(t testing.TB) {
	mockBatch := NewBatchTRX(mockBatchTRXHeader())
	mockBatch.AddEntry(mockTRXEntryDetail())
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchExpectedAddendaCount(0, 1)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTRXAddendaCountZero tests validating Addendum count of 0
func TestBatchTRXAddendaCountZero(t *testing.T) {
	testBatchTRXAddendaCountZero(t)
}

// BenchmarkBatchTRXAddendaCountZero benchmarks validating Addendum count of 0
func BenchmarkBatchTRXAddendaCountZero(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRXAddendaCountZero(b)
	}
}

// testBatchTRXInvalidAddendum validates Addendum must be Addenda05
func testBatchTRXInvalidAddendum(t testing.TB) {
	mockBatch := NewBatchTRX(mockBatchTRXHeader())
	mockBatch.AddEntry(mockTRXEntryDetail())
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchExpectedAddendaCount(0, 1)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTRXInvalidAddendum tests validating Addendum must be Addenda05
func TestBatchTRXInvalidAddendum(t *testing.T) {
	testBatchTRXInvalidAddendum(t)
}

// BenchmarkBatchTRXInvalidAddendum benchmarks validating Addendum must be Addenda05
func BenchmarkBatchTRXInvalidAddendum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRXInvalidAddendum(b)
	}
}

// testBatchTRXInvalidAddenda validates Addendum must be Addenda05 with type code 05
func testBatchTRXInvalidAddenda(t testing.TB) {
	mockBatch := NewBatchTRX(mockBatchTRXHeader())
	mockBatch.AddEntry(mockTRXEntryDetail())
	addenda05 := mockAddenda05()
	addenda05.TypeCode = "63"
	mockBatch.GetEntries()[0].AddAddenda05(addenda05)
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTRXInvalidAddenda tests validating Addendum must be Addenda05 with record type 7
func TestBatchTRXInvalidAddenda(t *testing.T) {
	testBatchTRXInvalidAddenda(t)
}

// BenchmarkBatchTRXInvalidAddenda benchmarks validating Addendum must be Addenda05 with record type 7
func BenchmarkBatchTRXInvalidAddenda(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRXInvalidAddenda(b)
	}
}

// testBatchTRXInvalidBuild validates an invalid batch build
func testBatchTRXInvalidBuild(t testing.TB) {
	mockBatch := mockBatchTRX()
	mockBatch.GetHeader().ServiceClassCode = 3
	err := mockBatch.Create()
	if !base.Match(err, ErrServiceClass) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTRXInvalidBuild tests validating an invalid batch build
func TestBatchTRXInvalidBuild(t *testing.T) {
	testBatchTRXInvalidBuild(t)
}

// BenchmarkBatchTRXInvalidBuild benchmarks validating an invalid batch build
func BenchmarkBatchTRXInvalidBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRXInvalidBuild(b)
	}
}

// testBatchTRXAddenda10000 validates error for 10000 Addenda
func testBatchTRXAddenda10000(t testing.TB) {

	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = TRX
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "ACH TRX"
	bh.ODFIIdentification = "12104288"

	entry := NewEntryDetail()
	entry.TransactionCode = CheckingDebit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.IdentificationNumber = "45689033"
	entry.SetCATXAddendaRecords(9999)
	entry.SetCATXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(mockBatchTRXHeader().ODFIIdentification, 1)
	entry.SetItemTypeIndicator("01")
	entry.Category = CategoryForward

	mockBatch := NewBatchTRX(bh)
	mockBatch.AddEntry(entry)

	for i := 0; i < 10000; i++ {
		mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	}
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchAddendaCount(10000, 9999)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTRXAddenda10000 tests validating error for 10000 Addenda
func TestBatchTRXAddenda10000(t *testing.T) {
	testBatchTRXAddenda10000(t)
}

// BenchmarkBatchTRXAddenda10000 benchmarks validating error for 10000 Addenda
func BenchmarkBatchTRXAddenda10000(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRXAddenda10000(b)
	}
}

// testBatchTRXAddendaRecords validates error for AddendaRecords not equal to addendum
func testBatchTRXAddendaRecords(t testing.TB) {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = TRX
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "ACH TRX"
	bh.ODFIIdentification = "12104288"

	entry := NewEntryDetail()
	entry.TransactionCode = CheckingDebit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.IdentificationNumber = "45689033"
	entry.SetCATXAddendaRecords(500)
	entry.SetCATXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(mockBatchTRXHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.Category = CategoryForward

	mockBatch := NewBatchTRX(bh)
	mockBatch.AddEntry(entry)

	for i := 0; i < 565; i++ {
		mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	}
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchExpectedAddendaCount(565, 500)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTRXAddendaRecords tests validating error for AddendaRecords not equal to addendum
func TestBatchTRXAddendaRecords(t *testing.T) {
	testBatchTRXAddendaRecords(t)
}

// BenchmarkBatchAddendaRecords benchmarks validating error for AddendaRecords not equal to addendum
func BenchmarkBatchTRXAddendaRecords(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRXAddendaRecords(b)
	}
}

// testBatchTRXReceivingCompany validates TRXReceivingCompany
func testBatchTRXReceivingCompany(t testing.TB) {
	mockBatch := mockBatchTRX()
	//mockBatch.GetEntries()[0].SetCATXReceivingCompany("Receiver")

	if mockBatch.GetEntries()[0].CATXReceivingCompanyField() != "Receiver Company" {
		t.Errorf("expected %v got %v", "Receiver Company", mockBatch.GetEntries()[0].CATXReceivingCompanyField())
	}
}

// TestBatchTRXReceivingCompany tests validating TRXReceivingCompany
func TestBatchTRXReceivingCompany(t *testing.T) {
	testBatchTRXReceivingCompany(t)
}

// BenchmarkBatchTRXReceivingCompany benchmarks validating TRXReceivingCompany
func BenchmarkBatchTRXReceivingCompany(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRXReceivingCompany(b)
	}
}

// testBatchTRXReserved validates TRXReservedField
func testBatchTRXReserved(t testing.TB) {
	mockBatch := mockBatchTRX()

	if mockBatch.GetEntries()[0].CATXReservedField() != "  " {
		t.Errorf("expected %v got %v", "  ", mockBatch.GetEntries()[0].CATXReservedField())
	}
}

// TestBatchTRXReserved tests validating TRXReservedField
func TestBatchTRXReserved(t *testing.T) {
	testBatchTRXReserved(t)
}

// BenchmarkBatchTRXReserved benchmarks validating TRXReservedField
func BenchmarkBatchTRXReserved(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRXReserved(b)
	}
}

// testBatchTRXZeroAddendaRecords validates zero addenda records
func testBatchTRXZeroAddendaRecords(t testing.TB) {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = TRX
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "ACH TRX"
	bh.ODFIIdentification = "12104288"

	entry := NewEntryDetail()
	entry.TransactionCode = CheckingDebit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.IdentificationNumber = "45689033"
	entry.SetCATXAddendaRecords(1)
	entry.SetCATXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(mockBatchTRXHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.Category = CategoryForward

	mockBatch := NewBatchTRX(bh)
	mockBatch.AddEntry(entry)

	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchExpectedAddendaCount(0, 1)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTRXZeroAddendaRecords tests validating zero addenda records
func TestBatchTRXZeroAddendaRecords(t *testing.T) {
	testBatchTRXZeroAddendaRecords(t)
}

// BenchmarkBatchZeroAddendaRecords benchmarks validating zero addenda records
func BenchmarkBatchTRXZeroAddendaRecords(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRXZeroAddendaRecords(b)
	}
}

// testBatchTRXTransactionCode validates BatchTRX TransactionCode is not a credit
func testBatchTRXTransactionCode(t testing.TB) {
	mockBatch := mockBatchTRXCredit()
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchDebitOnly) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTRXTransactionCode tests validating BatchTRX TransactionCode is not a credit
func TestBatchTRXTransactionCode(t *testing.T) {
	testBatchTRXTransactionCode(t)
}

// BenchmarkBatchTRXTransactionCode benchmarks validating BatchTRX TransactionCode is not a credit
func BenchmarkBatchTRXTransactionCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRXTransactionCode(b)
	}
}

// TestBatchTRXAddendum98 validates Addenda98 returns an error
func TestBatchTRXAddendum98(t *testing.T) {
	mockBatch := NewBatchTRX(mockBatchTRXHeader())
	mockBatch.AddEntry(mockTRXEntryDetail())
	mockAddenda98 := mockAddenda98()
	mockAddenda98.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTRXAddendum99 validates Addenda99 returns an error
func TestBatchTRXAddendum99(t *testing.T) {
	mockBatch := NewBatchTRX(mockBatchTRXHeader())
	mockBatch.AddEntry(mockCTXEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockAddenda99.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryReturn
	mockBatch.GetEntries()[0].Addenda99 = mockAddenda99
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// testBatchTRXCreditsOnly validates BatchTRX create for an invalid CreditsOnly
func testBatchTRXCreditsOnly(t testing.TB) {
	mockBatch := mockBatchTRX()
	mockBatch.Header.ServiceClassCode = CreditsOnly
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(CreditsOnly, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTRXCreditsOnly tests validating BatchTRX create for an invalid CreditsOnly
func TestBatchTRXCreditsOnly(t *testing.T) {
	testBatchTRXCreditsOnly(t)
}

// BenchmarkBatchTRXCreditsOnly benchmarks validating BatchTRX create for an invalid CreditsOnly
func BenchmarkBatchTRXCreditsOnly(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRXCreditsOnly(b)
	}
}

// testBatchTRXAutomatedAccountingAdvices validates BatchTRX create for an invalid AutomatedAccountingAdvices
func testBatchTRXAutomatedAccountingAdvices(t testing.TB) {
	mockBatch := mockBatchTRX()
	mockBatch.Header.ServiceClassCode = AutomatedAccountingAdvices
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(AutomatedAccountingAdvices, 225)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTRXAutomatedAccountingAdvices tests validating BatchTRX create for an invalid AutomatedAccountingAdvices
func TestBatchTRXAutomatedAccountingAdvices(t *testing.T) {
	testBatchTRXAutomatedAccountingAdvices(t)
}

// BenchmarkBatchTRXAutomatedAccountingAdvices benchmarks validating BatchTRX create for an invalid AutomatedAccountingAdvices
func BenchmarkBatchTRXAutomatedAccountingAdvices(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTRXAutomatedAccountingAdvices(b)
	}
}

// TestBatchTRXAddenda02 validates BatchTRX create for an invalid Addenda02
func TestBatchTRXAddenda02(t *testing.T) {
	mockBatch := mockBatchTRX()
	mockBatch.Entries[0].Addenda02 = mockAddenda02()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAddendaCategory) {
		t.Errorf("%T: %s", err, err)
	}
}

// testBatchTRXMixedDebitsAndCreditsServiceClassCode validates MixedDebitsAndCredits service class code
func testBatchTRXMixedDebitsAndCreditsServiceClassCode(t testing.TB) {
	mockBatch := mockBatchTRX()
	mockBatch.GetControl().ServiceClassCode = MixedDebitsAndCredits
	mockBatch.Header.ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTRXMixedDebitsAndCreditsServiceClassCode tests validates MixedDebitsAndCredits service class code
func TestBatchTRXMixedDebitsAndCreditsServiceClassCode(t *testing.T) {
	testBatchTRXMixedDebitsAndCreditsServiceClassCode(t)
}
