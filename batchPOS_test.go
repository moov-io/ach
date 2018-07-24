// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "testing"

// mockBatchPOSHeader creates a BatchPOS BatchHeader
func mockBatchPOSHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 225
	bh.StandardEntryClassCode = "POS"
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "ACH POS"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockPOSEntryDetail creates a BatchPOS EntryDetail
func mockPOSEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 27
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.IdentificationNumber = "45689033"
	entry.IndividualName = "Wade Arnold"
	//entry.SetReceivingCompany("ABC Company")
	entry.SetTraceNumber(mockBatchPOSHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.Category = CategoryForward
	return entry
}

// mockBatchPOS creates a BatchPOS
func mockBatchPOS() *BatchPOS {
	mockBatch := NewBatchPOS(mockBatchPOSHeader())
	mockBatch.AddEntry(mockPOSEntryDetail())
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda02())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// testBatchPOSHeader creates a BatchPOS BatchHeader
func testBatchPOSHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchPOSHeader())
	err, ok := batch.(*BatchPOS)
	if !ok {
		t.Errorf("Expecting BatchPOS got %T", err)
	}
}

// TestBatchPOSHeader tests validating BatchPOS BatchHeader
func TestBatchPOSHeader(t *testing.T) {
	testBatchPOSHeader(t)
}

// BenchmarkBatchPOSHeader benchmarks validating BatchPOS BatchHeader
func BenchmarkBatchPOSHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOSHeader(b)
	}
}

// testBatchPOSCreate validates BatchPOS create
func testBatchPOSCreate(t testing.TB) {
	mockBatch := mockBatchPOS()
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchPOSCreate tests validating BatchPOS create
func TestBatchPOSCreate(t *testing.T) {
	testBatchPOSCreate(t)
}

// BenchmarkBatchPOSCreate benchmarks validating BatchPOS create
func BenchmarkBatchPOSCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOSCreate(b)
	}
}

// testBatchPOSStandardEntryClassCode validates BatchPOS create for an invalid StandardEntryClassCode
func testBatchPOSStandardEntryClassCode(t testing.TB) {
	mockBatch := mockBatchPOS()
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

// TestBatchPOSStandardEntryClassCode tests validating BatchPOS create for an invalid StandardEntryClassCode
func TestBatchPOSStandardEntryClassCode(t *testing.T) {
	testBatchPOSStandardEntryClassCode(t)
}

// BenchmarkBatchPOSStandardEntryClassCode benchmarks validating BatchPOS create for an invalid StandardEntryClassCode
func BenchmarkBatchPOSStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOSStandardEntryClassCode(b)
	}
}

// testBatchPOSServiceClassCodeEquality validates service class code equality
func testBatchPOSServiceClassCodeEquality(t testing.TB) {
	mockBatch := mockBatchPOS()
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

// TestBatchPOSServiceClassCodeEquality tests validating service class code equality
func TestBatchPOSServiceClassCodeEquality(t *testing.T) {
	testBatchPOSServiceClassCodeEquality(t)
}

// BenchmarkBatchPOSServiceClassCodeEquality benchmarks validating service class code equality
func BenchmarkBatchPOSServiceClassCodeEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOSServiceClassCodeEquality(b)
	}
}

// testBatchPOSServiceClass200 validates BatchPOS create for an invalid ServiceClassCode 200
func testBatchPOSServiceClass200(t testing.TB) {
	mockBatch := mockBatchPOS()
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

// TestBatchPOSServiceClass200 tests validating BatchPOS create for an invalid ServiceClassCode 200
func TestBatchPOSServiceClass200(t *testing.T) {
	testBatchPOSServiceClass200(t)
}

// BenchmarkBatchPOSServiceClass200 benchmarks validating BatchPOS create for an invalid ServiceClassCode 200
func BenchmarkBatchPOSServiceClass200(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOSServiceClass200(b)
	}
}

// testBatchPOSServiceClass220 validates BatchPOS create for an invalid ServiceClassCode 220
func testBatchPOSServiceClass220(t testing.TB) {
	mockBatch := mockBatchPOS()
	mockBatch.Header.ServiceClassCode = 220
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

// TestBatchPOSServiceClass220 tests validating BatchPOS create for an invalid ServiceClassCode 220
func TestBatchPOSServiceClass220(t *testing.T) {
	testBatchPOSServiceClass220(t)
}

// BenchmarkBatchPOSServiceClass220 benchmarks validating BatchPOS create for an invalid ServiceClassCode 220
func BenchmarkBatchPOSServiceClass220(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOSServiceClass220(b)
	}
}

// testBatchPOSServiceClass280 validates BatchPOS create for an invalid ServiceClassCode 280
func testBatchPOSServiceClass280(t testing.TB) {
	mockBatch := mockBatchPOS()
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

// TestBatchPOSServiceClass280 tests validating BatchPOS create for an invalid ServiceClassCode 280
func TestBatchPOSServiceClass280(t *testing.T) {
	testBatchPOSServiceClass280(t)
}

// BenchmarkBatchPOSServiceClass280 benchmarks validating BatchPOS create for an invalid ServiceClassCode 280
func BenchmarkBatchPOSServiceClass280(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOSServiceClass280(b)
	}
}

// testBatchPOSTransactionCode validates BatchPOS TransactionCode is not a credit
func testBatchPOSTransactionCode(t testing.TB) {
	mockBatch := mockBatchPOS()
	mockBatch.GetEntries()[0].TransactionCode = 22
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

// TestBatchPOSTransactionCode tests validating BatchPOS TransactionCode is not a credit
func TestBatchPOSTransactionCode(t *testing.T) {
	testBatchPOSTransactionCode(t)
}

// BenchmarkBatchPOSTransactionCode benchmarks validating BatchPOS TransactionCode is not a credit
func BenchmarkBatchPOSTransactionCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOSTransactionCode(b)
	}
}

// testBatchPOSAddendaCount validates BatchPOS Addendum count of 2
func testBatchPOSAddendaCount(t testing.TB) {
	mockBatch := mockBatchPOS()
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda02())
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addendum" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchPOSAddendaCount tests validating BatchPOS Addendum count of 2
func TestBatchPOSAddendaCount(t *testing.T) {
	testBatchPOSAddendaCount(t)
}

// BenchmarkBatchPOSAddendaCount benchmarks validating BatchPOS Addendum count of 2
func BenchmarkBatchPOSAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOSAddendaCount(b)
	}
}

// testBatchPOSAddendaCountZero validates Addendum count of 0
func testBatchPOSAddendaCountZero(t testing.TB) {
	mockBatch := NewBatchPOS(mockBatchPOSHeader())
	mockBatch.AddEntry(mockPOSEntryDetail())
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addendum" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchPOSAddendaCountZero tests validating Addendum count of 0
func TestBatchPOSAddendaCountZero(t *testing.T) {
	testBatchPOSAddendaCountZero(t)
}

// BenchmarkBatchPOSAddendaCountZero benchmarks validating Addendum count of 0
func BenchmarkBatchPOSAddendaCountZero(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOSAddendaCountZero(b)
	}
}

// testBatchPOSInvalidAddendum validates Addendum must be Addenda02
func testBatchPOSInvalidAddendum(t testing.TB) {
	mockBatch := NewBatchPOS(mockBatchPOSHeader())
	mockBatch.AddEntry(mockPOSEntryDetail())
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda05())
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addendum" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchPOSInvalidAddendum tests validating Addendum must be Addenda02
func TestBatchPOSInvalidAddendum(t *testing.T) {
	testBatchPOSInvalidAddendum(t)
}

// BenchmarkBatchPOSInvalidAddendum benchmarks validating Addendum must be Addenda02
func BenchmarkBatchPOSInvalidAddendum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOSInvalidAddendum(b)
	}
}

// testBatchPOSInvalidAddenda validates Addendum must be Addenda02
func testBatchPOSInvalidAddenda(t testing.TB) {
	mockBatch := NewBatchPOS(mockBatchPOSHeader())
	mockBatch.AddEntry(mockPOSEntryDetail())
	addenda02 := mockAddenda02()
	addenda02.recordType = "63"
	mockBatch.GetEntries()[0].AddAddenda(addenda02)
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

// TestBatchPOSInvalidAddenda tests validating Addendum must be Addenda02
func TestBatchPOSInvalidAddenda(t *testing.T) {
	testBatchPOSInvalidAddenda(t)
}

// BenchmarkBatchPOSInvalidAddenda benchmarks validating Addendum must be Addenda02
func BenchmarkBatchPOSInvalidAddenda(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOSInvalidAddenda(b)
	}
}

// testBatchPOSInvalidBuild validates an invalid batch build
func testBatchPOSInvalidBuild(t testing.TB) {
	mockBatch := mockBatchPOS()
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

// TestBatchPOSInvalidBuild tests validating an invalid batch build
func TestBatchPOSInvalidBuild(t *testing.T) {
	testBatchPOSInvalidBuild(t)
}

// BenchmarkBatchPOSInvalidBuild benchmarks validating an invalid batch build
func BenchmarkBatchPOSInvalidBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOSInvalidBuild(b)
	}
}

// testBatchPOSCardTransactionType validates BatchPOS create for an invalid CardTransactionType
func testBatchPOSCardTransactionType(t testing.TB) {
	mockBatch := mockBatchPOS()
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

// TestBatchPOSCardTransactionType tests validating BatchPOS create for an invalid CardTransactionType
func TestBatchPOSCardTransactionType(t *testing.T) {
	testBatchPOSCardTransactionType(t)
}

// BenchmarkBatchPOSCardTransactionType benchmarks validating BatchPOS create for an invalid CardTransactionType
func BenchmarkBatchPOSCardTransactionType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOSCardTransactionType(b)
	}
}
