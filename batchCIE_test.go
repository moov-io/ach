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

// mockBatchCIEHeader creates a BatchCIE BatchHeader
func mockBatchCIEHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.StandardEntryClassCode = CIE
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "ACH CIE"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockCIEEntryDetail creates a BatchCIE EntryDetail
func mockCIEEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.IdentificationNumber = "45689033"
	entry.IndividualName = "Receiver Account Name"
	entry.SetTraceNumber(mockBatchCIEHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.Category = CategoryForward
	return entry
}

// mockBatchCIE creates a BatchCIE
func mockBatchCIE() *BatchCIE {
	mockBatch := NewBatchCIE(mockBatchCIEHeader())
	mockBatch.AddEntry(mockCIEEntryDetail())
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// testBatchCIEHeader creates a BatchCIE BatchHeader
func testBatchCIEHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchCIEHeader())
	err, ok := batch.(*BatchCIE)
	if !ok {
		t.Errorf("Expecting BatchCIE got %T", err)
	}
}

// TestBatchCIEHeader tests validating BatchCIE BatchHeader
func TestBatchCIEHeader(t *testing.T) {
	testBatchCIEHeader(t)
}

// BenchmarkBatchCIEHeader benchmarks validating BatchCIE BatchHeader
func BenchmarkBatchCIEHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCIEHeader(b)
	}
}

// testBatchCIECreate validates BatchCIE create
func testBatchCIECreate(t testing.TB) {
	mockBatch := mockBatchCIE()
	if err := mockBatch.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCIECreate tests validating BatchCIE create
func TestBatchCIECreate(t *testing.T) {
	testBatchCIECreate(t)
}

// BenchmarkBatchCIECreate benchmarks validating BatchCIE create
func BenchmarkBatchCIECreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCIECreate(b)
	}
}

// testBatchCIEStandardEntryClassCode validates BatchCIE create for an invalid StandardEntryClassCode
func testBatchCIEStandardEntryClassCode(t testing.TB) {
	mockBatch := mockBatchCIE()
	mockBatch.Header.StandardEntryClassCode = WEB
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchSECType) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCIEStandardEntryClassCode tests validating BatchCIE create for an invalid StandardEntryClassCode
func TestBatchCIEStandardEntryClassCode(t *testing.T) {
	testBatchCIEStandardEntryClassCode(t)
}

// BenchmarkBatchCIEStandardEntryClassCode benchmarks validating BatchCIE create for an invalid StandardEntryClassCode
func BenchmarkBatchCIEStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCIEStandardEntryClassCode(b)
	}
}

// testBatchCIEServiceClassCodeEquality validates service class code equality
func testBatchCIEServiceClassCodeEquality(t testing.TB) {
	mockBatch := mockBatchCIE()
	mockBatch.GetControl().ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(220, MixedDebitsAndCredits)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCIEServiceClassCodeEquality tests validating service class code equality
func TestBatchCIEServiceClassCodeEquality(t *testing.T) {
	testBatchCIEServiceClassCodeEquality(t)
}

// BenchmarkBatchCIEServiceClassCodeEquality benchmarks validating service class code equality
func BenchmarkBatchCIEServiceClassCodeEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCIEServiceClassCodeEquality(b)
	}
}

// testBatchCIEMixedCreditsAndDebits validates BatchCIE create for an invalid MixedCreditsAndDebits
func testBatchCIEMixedCreditsAndDebits(t testing.TB) {
	mockBatch := mockBatchCIE()
	mockBatch.Header.ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(MixedDebitsAndCredits, 220)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCIEMixedCreditsAndDebits tests validating BatchCIE create for an invalid MixedCreditsAndDebits
func TestBatchCIEMixedCreditsAndDebits(t *testing.T) {
	testBatchCIEMixedCreditsAndDebits(t)
}

// BenchmarkBatchCIEMixedCreditsAndDebits benchmarks validating BatchCIE create for an invalid MixedCreditsAndDebits
func BenchmarkBatchCIEMixedCreditsAndDebits(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCIEMixedCreditsAndDebits(b)
	}
}

// testBatchCIEDebitsOnly validates BatchCIE create for an invalid DebitsOnly
func testBatchCIEDebitsOnly(t testing.TB) {
	mockBatch := mockBatchCIE()
	mockBatch.Header.ServiceClassCode = DebitsOnly
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(DebitsOnly, 220)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCIEDebitsOnly tests validating BatchCIE create for an invalid DebitsOnly
func TestBatchCIEDebitsOnly(t *testing.T) {
	testBatchCIEDebitsOnly(t)
}

// BenchmarkBatchCIEDebitsOnly benchmarks validating BatchCIE create for an invalid DebitsOnly
func BenchmarkBatchCIEDebitsOnly(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCIEDebitsOnly(b)
	}
}

// testBatchCIEAutomatedAccountingAdvices validates BatchCIE create for an invalid AutomatedAccountingAdvices
func testBatchCIEAutomatedAccountingAdvices(t testing.TB) {
	mockBatch := mockBatchCIE()
	mockBatch.Header.ServiceClassCode = AutomatedAccountingAdvices
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchHeaderControlEquality(AutomatedAccountingAdvices, 220)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCIEAutomatedAccountingAdvices tests validating BatchCIE create for an invalid AutomatedAccountingAdvices
func TestBatchCIEAutomatedAccountingAdvices(t *testing.T) {
	testBatchCIEAutomatedAccountingAdvices(t)
}

// BenchmarkBatchCIEAutomatedAccountingAdvices benchmarks validating BatchCIE create for an invalid AutomatedAccountingAdvices
func BenchmarkBatchCIEAutomatedAccountingAdvices(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCIEAutomatedAccountingAdvices(b)
	}
}

// testBatchCIETransactionCode validates BatchCIE TransactionCode is not a debit
func testBatchCIETransactionCode(t testing.TB) {
	mockBatch := mockBatchCIE()
	mockBatch.GetEntries()[0].TransactionCode = CheckingDebit
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchDebitOnly) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCIETransactionCode tests validating BatchCIE TransactionCode is not a credit
func TestBatchCIETransactionCode(t *testing.T) {
	testBatchCIETransactionCode(t)
}

// BenchmarkBatchCIETransactionCode benchmarks validating BatchCIE TransactionCode is not a credit
func BenchmarkBatchCIETransactionCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCIETransactionCode(b)
	}
}

// testBatchCIEAddendaCount validates BatchCIE Addendum count of 2
func testBatchCIEAddendaCount(t testing.TB) {
	mockBatch := mockBatchCIE()
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchAddendaCount(2, 1)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCIEAddendaCount tests validating BatchCIE Addendum count of 2
func TestBatchCIEAddendaCount(t *testing.T) {
	testBatchCIEAddendaCount(t)
}

// BenchmarkBatchCIEAddendaCount benchmarks validating BatchCIE Addendum count of 2
func BenchmarkBatchCIEAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCIEAddendaCount(b)
	}
}

// testBatchCIEAddendaCountZero validates Addendum count of 0
func testBatchCIEAddendaCountZero(t testing.TB) {
	mockBatch := NewBatchCIE(mockBatchCIEHeader())
	mockBatch.AddEntry(mockCIEEntryDetail())
	err := mockBatch.Create()
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCIEAddendaCountZero tests validating Addendum count of 0
func TestBatchCIEAddendaCountZero(t *testing.T) {
	testBatchCIEAddendaCountZero(t)
}

// BenchmarkBatchCIEAddendaCountZero benchmarks validating Addendum count of 0
func BenchmarkBatchCIEAddendaCountZero(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCIEAddendaCountZero(b)
	}
}

// testBatchCIEInvalidAddendum validates Addendum must be Addenda05
func testBatchCIEInvalidAddendum(t testing.TB) {
	mockBatch := NewBatchCIE(mockBatchCIEHeader())
	mockBatch.AddEntry(mockCIEEntryDetail())
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAddendaCategory) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCIEInvalidAddendum tests validating Addendum must be Addenda05
func TestBatchCIEInvalidAddendum(t *testing.T) {
	testBatchCIEInvalidAddendum(t)
}

// BenchmarkBatchCIEInvalidAddendum benchmarks validating Addendum must be Addenda05
func BenchmarkBatchCIEInvalidAddendum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCIEInvalidAddendum(b)
	}
}

// testBatchCIEInvalidAddenda validates Addendum must be Addenda05 with type code 05
func testBatchCIEInvalidAddenda(t testing.TB) {
	mockBatch := NewBatchCIE(mockBatchCIEHeader())
	mockBatch.AddEntry(mockCIEEntryDetail())
	addenda05 := mockAddenda05()
	addenda05.TypeCode = "63"
	mockBatch.GetEntries()[0].AddAddenda05(addenda05)
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCIEInvalidAddenda tests validating Addendum must be Addenda05 with type code 05
func TestBatchCIEInvalidAddenda(t *testing.T) {
	testBatchCIEInvalidAddenda(t)
}

// BenchmarkBatchCIEInvalidAddenda benchmarks validating Addendum must be Addenda05 with type code 05
func BenchmarkBatchCIEInvalidAddenda(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCIEInvalidAddenda(b)
	}
}

// testBatchCIEInvalidBuild validates an invalid batch build
func testBatchCIEInvalidBuild(t testing.TB) {
	mockBatch := mockBatchCIE()
	mockBatch.GetHeader().ServiceClassCode = 3
	err := mockBatch.Create()
	if !base.Match(err, ErrServiceClass) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCIEInvalidBuild tests validating an invalid batch build
func TestBatchCIEInvalidBuild(t *testing.T) {
	testBatchCIEInvalidBuild(t)
}

// BenchmarkBatchCIEInvalidBuild benchmarks validating an invalid batch build
func BenchmarkBatchCIEInvalidBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCIEInvalidBuild(b)
	}
}

// testBatchCIECardTransactionType validates BatchCIE create for an invalid CardTransactionType
func testBatchCIECardTransactionType(t testing.TB) {
	mockBatch := mockBatchCIE()
	mockBatch.GetEntries()[0].DiscretionaryData = "555"
	err := mockBatch.Validate()
	// TODO: are we not expecting any errors here?
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCIECardTransactionType tests validating BatchCIE create for an invalid CardTransactionType
func TestBatchCIECardTransactionType(t *testing.T) {
	testBatchCIECardTransactionType(t)
}

// BenchmarkBatchCIECardTransactionType benchmarks validating BatchCIE create for an invalid CardTransactionType
func BenchmarkBatchCIECardTransactionType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCIECardTransactionType(b)
	}
}

// TestBatchCIEAddendum98 validates Addenda98 returns an error
func TestBatchCIEAddendum98(t *testing.T) {
	mockBatch := NewBatchCIE(mockBatchCIEHeader())
	mockBatch.AddEntry(mockCIEEntryDetail())
	mockAddenda98 := mockAddenda98()
	mockAddenda98.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCIEAddendum99 validates Addenda99 returns an error
func TestBatchCIEAddendum99(t *testing.T) {
	mockBatch := NewBatchCIE(mockBatchCIEHeader())
	mockBatch.AddEntry(mockCIEEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockAddenda99.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryReturn
	mockBatch.GetEntries()[0].Addenda99 = mockAddenda99
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCIEAddenda validates no more than 1 addenda record per entry detail record can exist
func TestBatchCIEAddenda(t *testing.T) {
	mockBatch := mockBatchCIE()
	// mock batch already has one addenda. Creating two addenda should error
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchAddendaCount(2, 1)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCIEAddenda02 validates BatchCIE cannot have Addenda02
func TestBatchCIEAddenda02(t *testing.T) {
	mockBatch := mockBatchCIE()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02()
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAddendaCategory) {
		t.Errorf("%T: %s", err, err)
	}
}

// testBatchCIEServiceClassCodeEquality validates service class code MixedDebitsAndCredits
func testBatchCIEMixedDebitsAndCreditsServiceClassCode(t testing.TB) {
	mockBatch := mockBatchCIE()
	mockBatch.GetControl().ServiceClassCode = MixedDebitsAndCredits
	mockBatch.Header.ServiceClassCode = MixedDebitsAndCredits
	err := mockBatch.Validate()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCIEMixedDebitsAndCreditsServiceClassCode tests validating SCC MixedDebitsAndCredits
func TestBatchCIEMixedDebitsAndCreditsServiceClassCode(t *testing.T) {
	testBatchCIEMixedDebitsAndCreditsServiceClassCode(t)
}
