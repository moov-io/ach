// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "testing"

// mockBatchRCKHeader creates a BatchRCK BatchHeader
func mockBatchRCKHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 225
	bh.StandardEntryClassCode = "RCK"
	bh.CompanyName = "Company Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "REDEPCHECK"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockRCKEntryDetail creates a BatchRCK EntryDetail
func mockRCKEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 27
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
	bh.ServiceClassCode = 225
	bh.StandardEntryClassCode = "RCK"
	bh.CompanyName = "Company Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "REDEPCHECK"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockRCKEntryDetailCredit creates a BatchRCK EntryDetail with a credit entry
func mockRCKEntryDetailCredit() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 22
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
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
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

// testBatchRCKServiceClass200 validates BatchRCK create for an invalid ServiceClassCode 200
func testBatchRCKServiceClass200(t testing.TB) {
	mockBatch := mockBatchRCK()
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

// TestBatchRCKServiceClass200 tests validating BatchRCK create for an invalid ServiceClassCode 200
func TestBatchRCKServiceClass200(t *testing.T) {
	testBatchRCKServiceClass200(t)
}

// BenchmarkBatchRCKServiceClass200 benchmarks validating BatchRCK create for an invalid ServiceClassCode 200
func BenchmarkBatchRCKServiceClass200(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchRCKServiceClass200(b)
	}
}

// testBatchRCKServiceClass220 validates BatchRCK create for an invalid ServiceClassCode 220
func testBatchRCKServiceClass220(t testing.TB) {
	mockBatch := mockBatchRCK()
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

// TestBatchRCKServiceClass220 tests validating BatchRCK create for an invalid ServiceClassCode 220
func TestBatchRCKServiceClass220(t *testing.T) {
	testBatchRCKServiceClass220(t)
}

// BenchmarkBatchRCKServiceClass220 benchmarks validating BatchRCK create for an invalid ServiceClassCode 220
func BenchmarkBatchRCKServiceClass220(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchRCKServiceClass220(b)
	}
}

// testBatchRCKServiceClass280 validates BatchRCK create for an invalid ServiceClassCode 280
func testBatchRCKServiceClass280(t testing.TB) {
	mockBatch := mockBatchRCK()
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

// TestBatchRCKServiceClass280 tests validating BatchRCK create for an invalid ServiceClassCode 280
func TestBatchRCKServiceClass280(t *testing.T) {
	testBatchRCKServiceClass280(t)
}

// BenchmarkBatchRCKServiceClass280 benchmarks validating BatchRCK create for an invalid ServiceClassCode 280
func BenchmarkBatchRCKServiceClass280(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchRCKServiceClass280(b)
	}
}

// testBatchRCKCompanyEntryDescription validates BatchRCK create for an invalid CompanyEntryDescription
func testBatchRCKCompanyEntryDescription(t testing.TB) {
	mockBatch := mockBatchRCK()
	mockBatch.Header.CompanyEntryDescription = "XYZ975"
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "CompanyEntryDescription" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Amount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "CheckSerialNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda05())
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
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
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
