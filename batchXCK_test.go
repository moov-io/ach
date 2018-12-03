// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "testing"

// mockBatchXCKHeader creates a BatchXCK BatchHeader
func mockBatchXCKHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 225
	bh.StandardEntryClassCode = "XCK"
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "XCK"
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
	bh.ServiceClassCode = 225
	bh.StandardEntryClassCode = "XCK"
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "XCK"
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
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
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

// testBatchXCKServiceClass200 validates BatchXCK create for an invalid ServiceClassCode 200
func testBatchXCKServiceClass200(t testing.TB) {
	mockBatch := mockBatchXCK()
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

// TestBatchXCKServiceClass200 tests validating BatchXCK create for an invalid ServiceClassCode 200
func TestBatchXCKServiceClass200(t *testing.T) {
	testBatchXCKServiceClass200(t)
}

// BenchmarkBatchXCKServiceClass200 benchmarks validating BatchXCK create for an invalid ServiceClassCode 200
func BenchmarkBatchXCKServiceClass200(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchXCKServiceClass200(b)
	}
}

// testBatchXCKServiceClass220 validates BatchXCK create for an invalid ServiceClassCode 220
func testBatchXCKServiceClass220(t testing.TB) {
	mockBatch := mockBatchXCK()
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

// TestBatchXCKServiceClass220 tests validating BatchXCK create for an invalid ServiceClassCode 220
func TestBatchXCKServiceClass220(t *testing.T) {
	testBatchXCKServiceClass220(t)
}

// BenchmarkBatchXCKServiceClass220 benchmarks validating BatchXCK create for an invalid ServiceClassCode 220
func BenchmarkBatchXCKServiceClass220(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchXCKServiceClass220(b)
	}
}

// testBatchXCKServiceClass280 validates BatchXCK create for an invalid ServiceClassCode 280
func testBatchXCKServiceClass280(t testing.TB) {
	mockBatch := mockBatchXCK()
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

// TestBatchXCKServiceClass280 tests validating BatchXCK create for an invalid ServiceClassCode 280
func TestBatchXCKServiceClass280(t *testing.T) {
	testBatchXCKServiceClass280(t)
}

// BenchmarkBatchXCKServiceClass280 benchmarks validating BatchXCK create for an invalid ServiceClassCode 280
func BenchmarkBatchXCKServiceClass280(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchXCKServiceClass280(b)
	}
}

// testBatchXCKCheckSerialNumber validates BatchXCK CheckSerialNumber is not mandatory
func testBatchXCKCheckSerialNumber(t testing.TB) {
	mockBatch := mockBatchXCK()
	// modify CheckSerialNumber / IdentificationNumber to nothing
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
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addenda05" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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

// TestBatchXCKAddendum99 validates Addenda99 returns an error
func TestBatchXCKAddendum99(t *testing.T) {
	mockBatch := NewBatchXCK(mockBatchXCKHeader())
	mockBatch.AddEntry(mockXCKEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockAddenda99.TypeCode = "05"
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

// TestBatchXCKAddendum99Category validates Addenda99 returns an error
func TestBatchXCKAddendum99Category(t *testing.T) {
	mockBatch := NewBatchXCK(mockBatchXCKHeader())
	mockBatch.AddEntry(mockXCKEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].Category = CategoryForward
	mockBatch.GetEntries()[0].Addenda99 = mockAddenda99
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addenda99" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchXCKProcessControlField returns an error if ProcessControlField is not defined.
func TestBatchXCKProcessControlField(t *testing.T) {
	mockBatch := NewBatchXCK(mockBatchXCKHeader())
	mockBatch.AddEntry(mockXCKEntryDetail())
	mockBatch.GetEntries()[0].SetProcessControlField("")
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ProcessControlField" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchXCKItemResearchNumber returns an error if ItemResearchNumber is not defined.
func TestBatchXCKItemResearchNumber(t *testing.T) {
	mockBatch := NewBatchXCK(mockBatchXCKHeader())
	mockBatch.AddEntry(mockXCKEntryDetail())
	mockBatch.GetEntries()[0].IndividualName = ""
	mockBatch.GetEntries()[0].SetProcessControlField("CHECK1")
	mockBatch.GetEntries()[0].SetItemResearchNumber("")
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ItemResearchNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchXCKAmount validates BatchXCK create for an invalid Amount
func TestBatchXCKAmount(t *testing.T) {
	mockBatch := mockBatchXCK()
	mockBatch.Entries[0].Amount = 260000
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
