// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"log"
	"testing"
)

// mockBatchCTXHeader creates a BatchCTX BatchHeader
func mockBatchCTXHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 220
	bh.StandardEntryClassCode = "CTX"
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "ACH CTX"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockCTXEntryDetail creates a BatchCTX EntryDetail
func mockCTXEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 22
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.IdentificationNumber = "45689033"
	entry.SetCTXAddendaRecords(1)
	entry.SetCTXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(mockBatchCTXHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.Category = CategoryForward
	return entry
}

// mockBatchCTX creates a BatchCTX
func mockBatchCTX() *BatchCTX {
	mockBatch := NewBatchCTX(mockBatchCTXHeader())
	mockBatch.AddEntry(mockCTXEntryDetail())
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda05())
	if err := mockBatch.Create(); err != nil {
		log.Fatal(err)
	}
	return mockBatch
}

// testBatchCTXHeader creates a BatchCTX BatchHeader
func testBatchCTXHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchCTXHeader())
	err, ok := batch.(*BatchCTX)
	if !ok {
		t.Errorf("Expecting BatchCTX got %T", err)
	}
}

// TestBatchCTXHeader tests validating BatchCTX BatchHeader
func TestBatchCTXHeader(t *testing.T) {
	testBatchCTXHeader(t)
}

// BenchmarkBatchCTXHeader benchmarks validating BatchCTX BatchHeader
func BenchmarkBatchCTXHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXHeader(b)
	}
}

// testBatchCTXCreate validates BatchCTX create
func testBatchCTXCreate(t testing.TB) {
	mockBatch := mockBatchCTX()
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCTXCreate tests validating BatchCTX create
func TestBatchCTXCreate(t *testing.T) {
	testBatchCTXCreate(t)
}

// BenchmarkBatchCTXCreate benchmarks validating BatchCTX create
func BenchmarkBatchCTXCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXCreate(b)
	}
}

// testBatchCTXStandardEntryClassCode validates BatchCTX create for an invalid StandardEntryClassCode
func testBatchCTXStandardEntryClassCode(t testing.TB) {
	mockBatch := mockBatchCTX()
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

// TestBatchCTXStandardEntryClassCode tests validating BatchCTX create for an invalid StandardEntryClassCode
func TestBatchCTXStandardEntryClassCode(t *testing.T) {
	testBatchCTXStandardEntryClassCode(t)
}

// BenchmarkBatchCTXStandardEntryClassCode benchmarks validating BatchCTX create for an invalid StandardEntryClassCode
func BenchmarkBatchCTXStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXStandardEntryClassCode(b)
	}
}

// testBatchCTXServiceClassCodeEquality validates service class code equality
func testBatchCTXServiceClassCodeEquality(t testing.TB) {
	mockBatch := mockBatchCTX()
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

// TestBatchCTXServiceClassCodeEquality tests validating service class code equality
func TestBatchCTXServiceClassCodeEquality(t *testing.T) {
	testBatchCTXServiceClassCodeEquality(t)
}

// BenchmarkBatchCTXServiceClassCodeEquality benchmarks validating service class code equality
func BenchmarkBatchCTXServiceClassCodeEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXServiceClassCodeEquality(b)
	}
}

// testBatchCTXAddendaCount validates BatchCTX Addendum count of 2
func testBatchCTXAddendaCount(t testing.TB) {
	mockBatch := mockBatchCTX()
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda05())
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

// TestBatchCTXAddendaCount tests validating BatchCTX Addendum count of 2
func TestBatchCTXAddendaCount(t *testing.T) {
	testBatchCTXAddendaCount(t)
}

// BenchmarkBatchCTXAddendaCount benchmarks validating BatchCTX Addendum count of 2
func BenchmarkBatchCTXAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXAddendaCount(b)
	}
}

// testBatchCTXAddendaCountZero validates Addendum count of 0
func testBatchCTXAddendaCountZero(t testing.TB) {
	mockBatch := NewBatchCTX(mockBatchCTXHeader())
	mockBatch.AddEntry(mockCTXEntryDetail())
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

// TestBatchCTXAddendaCountZero tests validating Addendum count of 0
func TestBatchCTXAddendaCountZero(t *testing.T) {
	testBatchCTXAddendaCountZero(t)
}

// BenchmarkBatchCTXAddendaCountZero benchmarks validating Addendum count of 0
func BenchmarkBatchCTXAddendaCountZero(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXAddendaCountZero(b)
	}
}

// testBatchCTXInvalidAddendum validates Addendum must be Addenda05
func testBatchCTXInvalidAddendum(t testing.TB) {
	mockBatch := NewBatchCTX(mockBatchCTXHeader())
	mockBatch.AddEntry(mockCTXEntryDetail())
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda02())
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

// TestBatchCTXInvalidAddendum tests validating Addendum must be Addenda05
func TestBatchCTXInvalidAddendum(t *testing.T) {
	testBatchCTXInvalidAddendum(t)
}

// BenchmarkBatchCTXInvalidAddendum benchmarks validating Addendum must be Addenda05
func BenchmarkBatchCTXInvalidAddendum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXInvalidAddendum(b)
	}
}

// testBatchCTXInvalidAddenda validates Addendum must be Addenda05 with record type 7
func testBatchCTXInvalidAddenda(t testing.TB) {
	mockBatch := NewBatchCTX(mockBatchCTXHeader())
	mockBatch.AddEntry(mockCTXEntryDetail())
	addenda05 := mockAddenda05()
	addenda05.recordType = "63"
	mockBatch.GetEntries()[0].AddAddenda(addenda05)
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

// TestBatchCTXInvalidAddenda tests validating Addendum must be Addenda05 with record type 7
func TestBatchCTXInvalidAddenda(t *testing.T) {
	testBatchCTXInvalidAddenda(t)
}

// BenchmarkBatchCTXInvalidAddenda benchmarks validating Addendum must be Addenda05 with record type 7
func BenchmarkBatchCTXInvalidAddenda(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXInvalidAddenda(b)
	}
}

// testBatchCTXInvalidBuild validates an invalid batch build
func testBatchCTXInvalidBuild(t testing.TB) {
	mockBatch := mockBatchCTX()
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

// TestBatchCTXInvalidBuild tests validating an invalid batch build
func TestBatchCTXInvalidBuild(t *testing.T) {
	testBatchCTXInvalidBuild(t)
}

// BenchmarkBatchCTXInvalidBuild benchmarks validating an invalid batch build
func BenchmarkBatchCTXInvalidBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXInvalidBuild(b)
	}
}

// testBatchCTXAddenda10000 validates error for 10000 Addenda
func testBatchCTXAddenda10000(t testing.TB) {

	bh := NewBatchHeader()
	bh.ServiceClassCode = 220
	bh.StandardEntryClassCode = "CTX"
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "ACH CTX"
	bh.ODFIIdentification = "12104288"

	entry := NewEntryDetail()
	entry.TransactionCode = 22
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.IdentificationNumber = "45689033"
	entry.SetCTXAddendaRecords(9999)
	entry.SetCTXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(mockBatchCTXHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.Category = CategoryForward

	mockBatch := NewBatchCTX(bh)
	mockBatch.AddEntry(entry)

	for i := 0; i < 10000; i++ {
		mockBatch.GetEntries()[0].AddAddenda(mockAddenda05())
	}

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

// TestBatchCTXAddenda10000 tests validating error for 10000 Addenda
func TestBatchCTXAddenda10000(t *testing.T) {
	testBatchCTXAddenda10000(t)
}

// BenchmarkBatchCTXAddenda10000 benchmarks validating error for 10000 Addenda
func BenchmarkBatchCTXAddenda10000(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXAddenda10000(b)
	}
}

// testBatchCTXAddendaRecords validates error for AddendaRecords not equal to addendum
func testBatchCTXAddendaRecords(t testing.TB) {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 220
	bh.StandardEntryClassCode = "CTX"
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "ACH CTX"
	bh.ODFIIdentification = "12104288"

	entry := NewEntryDetail()
	entry.TransactionCode = 22
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.IdentificationNumber = "45689033"
	entry.SetCTXAddendaRecords(500)
	entry.SetCTXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(mockBatchCTXHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.Category = CategoryForward

	mockBatch := NewBatchCTX(bh)
	mockBatch.AddEntry(entry)

	for i := 0; i < 565; i++ {
		mockBatch.GetEntries()[0].AddAddenda(mockAddenda05())
	}

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

// TestBatchCTXAddendaRecords tests validating error for AddendaRecords not equal to addendum
func TestBatchCTXAddendaRecords(t *testing.T) {
	testBatchCTXAddendaRecords(t)
}

// BenchmarkBatchAddendaRecords benchmarks validating error for AddendaRecords not equal to addendum
func BenchmarkBatchCTXAddendaRecords(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXAddendaRecords(b)
	}
}

// testBatchCTXReceivingCompany validates CTXReceivingCompany
func testBatchCTXReceivingCompany(t testing.TB) {
	mockBatch := mockBatchCTX()
	//mockBatch.GetEntries()[0].SetCTXReceivingCompany("Receiver")

	if mockBatch.GetEntries()[0].CTXReceivingCompanyField() != "Receiver Company" {
		t.Errorf("expected %v got %v", "Receiver Company", mockBatch.GetEntries()[0].CTXReceivingCompanyField())
	}
}

// TestBatchCTXReceivingCompany tests validating CTXReceivingCompany
func TestBatchCTXReceivingCompany(t *testing.T) {
	testBatchCTXReceivingCompany(t)
}

// BenchmarkBatchCTXReceivingCompany benchmarks validating CTXReceivingCompany
func BenchmarkBatchCTXReceivingCompany(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXReceivingCompany(b)
	}
}

// testBatchCTXReserved validates CTXReservedField
func testBatchCTXReserved(t testing.TB) {
	mockBatch := mockBatchCTX()

	if mockBatch.GetEntries()[0].CTXReservedField() != "  " {
		t.Errorf("expected %v got %v", "  ", mockBatch.GetEntries()[0].CTXReservedField())
	}
}

// TestBatchCTXReserved tests validating CTXReservedField
func TestBatchCTXReserved(t *testing.T) {
	testBatchCTXReserved(t)
}

// BenchmarkBatchCTXReserved benchmarks validating CTXReservedField
func BenchmarkBatchCTXReserved(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCTXReserved(b)
	}
}
