// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "testing"

// mockBatchCIEHeader creates a BatchCIE BatchHeader
func mockBatchCIEHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.StandardEntryClassCode = "CIE"
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
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
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
	mockBatch.Header.StandardEntryClassCode = "WEB"
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "StandardEntryClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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
	mockBatch.GetControl().ServiceClassCode = 200
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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

// testBatchCIEServiceClass200 validates BatchCIE create for an invalid ServiceClassCode 200
func testBatchCIEServiceClass200(t testing.TB) {
	mockBatch := mockBatchCIE()
	mockBatch.Header.ServiceClassCode = 200
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCIEServiceClass200 tests validating BatchCIE create for an invalid ServiceClassCode 200
func TestBatchCIEServiceClass200(t *testing.T) {
	testBatchCIEServiceClass200(t)
}

// BenchmarkBatchCIEServiceClass200 benchmarks validating BatchCIE create for an invalid ServiceClassCode 200
func BenchmarkBatchCIEServiceClass200(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCIEServiceClass200(b)
	}
}

// testBatchCIEServiceClass225 validates BatchCIE create for an invalid ServiceClassCode 225
func testBatchCIEServiceClass225(t testing.TB) {
	mockBatch := mockBatchCIE()
	mockBatch.Header.ServiceClassCode = DebitsOnly
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCIEServiceClass225 tests validating BatchCIE create for an invalid ServiceClassCode 225
func TestBatchCIEServiceClass225(t *testing.T) {
	testBatchCIEServiceClass225(t)
}

// BenchmarkBatchCIEServiceClass225 benchmarks validating BatchCIE create for an invalid ServiceClassCode 225
func BenchmarkBatchCIEServiceClass225(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCIEServiceClass225(b)
	}
}

// testBatchCIEServiceClass280 validates BatchCIE create for an invalid ServiceClassCode 280
func testBatchCIEServiceClass280(t testing.TB) {
	mockBatch := mockBatchCIE()
	mockBatch.Header.ServiceClassCode = 280
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCIEServiceClass280 tests validating BatchCIE create for an invalid ServiceClassCode 280
func TestBatchCIEServiceClass280(t *testing.T) {
	testBatchCIEServiceClass280(t)
}

// BenchmarkBatchCIEServiceClass280 benchmarks validating BatchCIE create for an invalid ServiceClassCode 280
func BenchmarkBatchCIEServiceClass280(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCIEServiceClass280(b)
	}
}

// testBatchCIETransactionCode validates BatchCIE TransactionCode is not a debit
func testBatchCIETransactionCode(t testing.TB) {
	mockBatch := mockBatchCIE()
	mockBatch.GetEntries()[0].TransactionCode = CheckingDebit
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TransactionCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "AddendaCount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "AddendaCount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "AddendaCount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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

// testBatchCIEInvalidAddenda validates Addendum must be Addenda05 with record type 7
func testBatchCIEInvalidAddenda(t testing.TB) {
	mockBatch := NewBatchCIE(mockBatchCIEHeader())
	mockBatch.AddEntry(mockCIEEntryDetail())
	addenda05 := mockAddenda05()
	addenda05.recordType = "63"
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCIEInvalidAddenda tests validating Addendum must be Addenda05 with record type 7
func TestBatchCIEInvalidAddenda(t *testing.T) {
	testBatchCIEInvalidAddenda(t)
}

// BenchmarkBatchCIEInvalidAddenda benchmarks validating Addendum must be Addenda05 with record type 7
func BenchmarkBatchCIEInvalidAddenda(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCIEInvalidAddenda(b)
	}
}

// testBatchCIEInvalidBuild validates an invalid batch build
func testBatchCIEInvalidBuild(t testing.TB) {
	mockBatch := mockBatchCIE()
	mockBatch.GetHeader().recordType = "3"
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "CardTransactionType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCIEAddenda validates no more than 1 addenda record per entry detail record can exist
func TestBatchCIEAddenda(t *testing.T) {
	mockBatch := mockBatchCIE()
	// mock batch already has one addenda. Creating two addenda should error
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "AddendaCount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCIEAddenda02 validates BatchCIE cannot have Addenda02
func TestBatchCIEAddenda02(t *testing.T) {
	mockBatch := mockBatchCIE()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02()
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addenda02" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}
