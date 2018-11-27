// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"
)

// mockBatchSHRHeader creates a BatchSHR BatchHeader
func mockBatchSHRHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 225
	bh.StandardEntryClassCode = "SHR"
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "ACH SHR"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockSHREntryDetail creates a BatchSHR EntryDetail
func mockSHREntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 27
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetSHRCardExpirationDate("0722")
	entry.SetSHRDocumentReferenceNumber("12345678910")
	entry.SetSHRIndividualCardAccountNumber("1234567891123456789")
	entry.SetTraceNumber(mockBatchSHRHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.Category = CategoryForward
	return entry
}

// mockBatchSHR creates a BatchSHR
func mockBatchSHR() *BatchSHR {
	mockBatch := NewBatchSHR(mockBatchSHRHeader())
	mockBatch.AddEntry(mockSHREntryDetail())
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// testBatchSHRHeader creates a BatchSHR BatchHeader
func testBatchSHRHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchSHRHeader())
	err, ok := batch.(*BatchSHR)
	if !ok {
		t.Errorf("Expecting BatchSHR got %T", err)
	}
}

// TestBatchSHRHeader tests validating BatchSHR BatchHeader
func TestBatchSHRHeader(t *testing.T) {
	testBatchSHRHeader(t)
}

// BenchmarkBatchSHRHeader benchmarks validating BatchSHR BatchHeader
func BenchmarkBatchSHRHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRHeader(b)
	}
}

// testBatchSHRCreate validates BatchSHR create
func testBatchSHRCreate(t testing.TB) {
	mockBatch := mockBatchSHR()
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchSHRCreate tests validating BatchSHR create
func TestBatchSHRCreate(t *testing.T) {
	testBatchSHRCreate(t)
}

// BenchmarkBatchSHRCreate benchmarks validating BatchSHR create
func BenchmarkBatchSHRCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRCreate(b)
	}
}

// testBatchSHRStandardEntryClassCode validates BatchSHR create for an invalid StandardEntryClassCode
func testBatchSHRStandardEntryClassCode(t testing.TB) {
	mockBatch := mockBatchSHR()
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

// TestBatchSHRStandardEntryClassCode tests validating BatchSHR create for an invalid StandardEntryClassCode
func TestBatchSHRStandardEntryClassCode(t *testing.T) {
	testBatchSHRStandardEntryClassCode(t)
}

// BenchmarkBatchSHRStandardEntryClassCode benchmarks validating BatchSHR create for an invalid StandardEntryClassCode
func BenchmarkBatchSHRStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRStandardEntryClassCode(b)
	}
}

// testBatchSHRServiceClassCodeEquality validates service class code equality
func testBatchSHRServiceClassCodeEquality(t testing.TB) {
	mockBatch := mockBatchSHR()
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

// TestBatchSHRServiceClassCodeEquality tests validating service class code equality
func TestBatchSHRServiceClassCodeEquality(t *testing.T) {
	testBatchSHRServiceClassCodeEquality(t)
}

// BenchmarkBatchSHRServiceClassCodeEquality benchmarks validating service class code equality
func BenchmarkBatchSHRServiceClassCodeEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRServiceClassCodeEquality(b)
	}
}

// testBatchSHRServiceClass200 validates BatchSHR create for an invalid ServiceClassCode 200
func testBatchSHRServiceClass200(t testing.TB) {
	mockBatch := mockBatchSHR()
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

// TestBatchSHRServiceClass200 tests validating BatchSHR create for an invalid ServiceClassCode 200
func TestBatchSHRServiceClass200(t *testing.T) {
	testBatchSHRServiceClass200(t)
}

// BenchmarkBatchSHRServiceClass200 benchmarks validating BatchSHR create for an invalid ServiceClassCode 200
func BenchmarkBatchSHRServiceClass200(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRServiceClass200(b)
	}
}

// testBatchSHRServiceClass220 validates BatchSHR create for an invalid ServiceClassCode 220
func testBatchSHRServiceClass220(t testing.TB) {
	mockBatch := mockBatchSHR()
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

// TestBatchSHRServiceClass220 tests validating BatchSHR create for an invalid ServiceClassCode 220
func TestBatchSHRServiceClass220(t *testing.T) {
	testBatchSHRServiceClass220(t)
}

// BenchmarkBatchSHRServiceClass220 benchmarks validating BatchSHR create for an invalid ServiceClassCode 220
func BenchmarkBatchSHRServiceClass220(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRServiceClass220(b)
	}
}

// testBatchSHRServiceClass280 validates BatchSHR create for an invalid ServiceClassCode 280
func testBatchSHRServiceClass280(t testing.TB) {
	mockBatch := mockBatchSHR()
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

// TestBatchSHRServiceClass280 tests validating BatchSHR create for an invalid ServiceClassCode 280
func TestBatchSHRServiceClass280(t *testing.T) {
	testBatchSHRServiceClass280(t)
}

// BenchmarkBatchSHRServiceClass280 benchmarks validating BatchSHR create for an invalid ServiceClassCode 280
func BenchmarkBatchSHRServiceClass280(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRServiceClass280(b)
	}
}

// testBatchSHRTransactionCode validates BatchSHR TransactionCode is not a credit
func testBatchSHRTransactionCode(t testing.TB) {
	mockBatch := mockBatchSHR()
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

// TestBatchSHRTransactionCode tests validating BatchSHR TransactionCode is not a credit
func TestBatchSHRTransactionCode(t *testing.T) {
	testBatchSHRTransactionCode(t)
}

// BenchmarkBatchSHRTransactionCode benchmarks validating BatchSHR TransactionCode is not a credit
func BenchmarkBatchSHRTransactionCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRTransactionCode(b)
	}
}

// testBatchSHRAddendaCount validates BatchSHR Addendum count of 2
func testBatchSHRAddendaCount(t testing.TB) {
	mockBatch := mockBatchSHR()
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02()
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

// TestBatchSHRAddendaCount tests validating BatchSHR Addendum count of 2
func TestBatchSHRAddendaCount(t *testing.T) {
	testBatchSHRAddendaCount(t)
}

// BenchmarkBatchSHRAddendaCount benchmarks validating BatchSHR Addendum count of 2
func BenchmarkBatchSHRAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRAddendaCount(b)
	}
}

// testBatchSHRAddendaCountZero validates Addendum count of 0
func testBatchSHRAddendaCountZero(t testing.TB) {
	mockBatch := NewBatchSHR(mockBatchSHRHeader())
	mockBatch.AddEntry(mockSHREntryDetail())
	mockAddenda02 := mockAddenda02()
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02
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

// TestBatchSHRAddendaCountZero tests validating Addendum count of 0
func TestBatchSHRAddendaCountZero(t *testing.T) {
	testBatchSHRAddendaCountZero(t)
}

// BenchmarkBatchSHRAddendaCountZero benchmarks validating Addendum count of 0
func BenchmarkBatchSHRAddendaCountZero(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRAddendaCountZero(b)
	}
}

// TestBatchSHRAddendum98 validates Addenda98 returns an error
func TestBatchSHRAddendum98(t *testing.T) {
	mockBatch := NewBatchSHR(mockBatchSHRHeader())
	mockBatch.AddEntry(mockSHREntryDetail())
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

// TestBatchSHRAddendum99 validates Addenda99 returns an error
func TestBatchSHRAddendum99(t *testing.T) {
	mockBatch := NewBatchSHR(mockBatchSHRHeader())
	mockBatch.AddEntry(mockSHREntryDetail())
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

// testBatchSHRInvalidAddendum validates Addendum must be Addenda02
func testBatchSHRInvalidAddendum(t testing.TB) {
	mockBatch := NewBatchSHR(mockBatchSHRHeader())
	mockBatch.AddEntry(mockSHREntryDetail())
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addenda05" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchSHRInvalidAddendum tests validating Addendum must be Addenda02
func TestBatchSHRInvalidAddendum(t *testing.T) {
	testBatchSHRInvalidAddendum(t)
}

// BenchmarkBatchSHRInvalidAddendum benchmarks validating Addendum must be Addenda02
func BenchmarkBatchSHRInvalidAddendum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRInvalidAddendum(b)
	}
}

// testBatchSHRInvalidAddenda validates Addendum must be Addenda02
func testBatchSHRInvalidAddenda(t testing.TB) {
	mockBatch := NewBatchSHR(mockBatchSHRHeader())
	mockBatch.AddEntry(mockSHREntryDetail())
	addenda02 := mockAddenda02()
	addenda02.recordType = "63"
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02()
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

// TestBatchSHRInvalidAddenda tests validating Addendum must be Addenda02
func TestBatchSHRInvalidAddenda(t *testing.T) {
	testBatchSHRInvalidAddenda(t)
}

// BenchmarkBatchSHRInvalidAddenda benchmarks validating Addendum must be Addenda02
func BenchmarkBatchSHRInvalidAddenda(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRInvalidAddenda(b)
	}
}

// testBatchSHRInvalidBuild validates an invalid batch build
func testBatchSHRInvalidBuild(t testing.TB) {
	mockBatch := mockBatchSHR()
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

// TestBatchSHRInvalidBuild tests validating an invalid batch build
func TestBatchSHRInvalidBuild(t *testing.T) {
	testBatchSHRInvalidBuild(t)
}

// BenchmarkBatchSHRInvalidBuild benchmarks validating an invalid batch build
func BenchmarkBatchSHRInvalidBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRInvalidBuild(b)
	}
}

// testBatchSHRCardTransactionType validates BatchSHR create for an invalid CardTransactionType
func testBatchSHRCardTransactionType(t testing.TB) {
	mockBatch := mockBatchSHR()
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

// TestBatchSHRCardTransactionType tests validating BatchSHR create for an invalid CardTransactionType
func TestBatchSHRCardTransactionType(t *testing.T) {
	testBatchSHRCardTransactionType(t)
}

// BenchmarkBatchSHRCardTransactionType benchmarks validating BatchSHR create for an invalid CardTransactionType
func BenchmarkBatchSHRCardTransactionType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRCardTransactionType(b)
	}
}

// testBatchSHRCardExpirationDateField validates SHRCardExpirationDate
// characters 0-4 of underlying IdentificationNumber
func testBatchSHRCardExpirationDateField(t testing.TB) {
	mockBatch := mockBatchSHR()
	ts := mockBatch.Entries[0].SHRCardExpirationDateField()
	if ts != "0722" {
		t.Error("Card Expiration Date is invalid")
	}
}

// TestBatchSHRCardExpirationDateField tests validatingSHRCardExpirationDate
// characters 0-4 of underlying IdentificationNumber
func TestBatchSHRCardExpirationDateField(t *testing.T) {
	testBatchSHRCardExpirationDateField(t)
}

// BenchmarkBatchSHRCardExpirationDateField benchmarks validating SHRCardExpirationDate
// characters 0-4 of underlying IdentificationNumber
func BenchmarkBatchSHRCardExpirationDateField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRCardExpirationDateField(b)
	}
}

// testBatchSHRDocumentReferenceNumberField validates SHRDocumentReferenceNumberField
// characters 5-15 of underlying IdentificationNumber
func testBatchSHRDocumentReferenceNumberField(t testing.TB) {
	mockBatch := mockBatchSHR()
	ts := mockBatch.Entries[0].SHRDocumentReferenceNumberField()
	if ts != "12345678910" {
		t.Error("Document Reference Number is invalid")
	}
}

// TestBatchSHRDocumentReferenceNumberField tests validating SHRDocumentReferenceNumberField
// characters 5-15 of underlying IdentificationNumber
func TestBatchSHRDocumentReferenceNumberField(t *testing.T) {
	testBatchSHRDocumentReferenceNumberField(t)
}

// BenchmarkBatchSHRDocumentReferenceNumberField benchmarks validating SHRDocumentReferenceNumberField
// characters 5-15 of underlying IdentificationNumber
func BenchmarkSHRDocumentReferenceNumberField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRDocumentReferenceNumberField(b)
	}
}

// testBatchSHRIndividualCardAccountNumberField validates SHRIndividualCardAccountNumberField
// underlying IndividualName
func testBatchSHRIndividualCardAccountNumberField(t testing.TB) {
	mockBatch := mockBatchSHR()
	ts := mockBatch.Entries[0].SHRIndividualCardAccountNumberField()
	if ts != "0001234567891123456789" {
		t.Error("Individual Card Account Number is invalid")
	}
}

// TestBatchSHRIndividualCardAccountNumberField tests validating SHRIndividualCardAccountNumberField
// underlying IndividualName
func TestBatchSHRIndividualCardAccountNumberField(t *testing.T) {
	testBatchSHRIndividualCardAccountNumberField(t)
}

// BenchmarkBatchSHRIndividualCardAccountNumberField benchmarks validating SHRIndividualCardAccountNumberField
// underlying IndividualName
func BenchmarkBatchSHRDocumentReferenceNumberField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchSHRIndividualCardAccountNumberField(b)
	}
}

// testSHRCardExpirationDateMonth validates the month is valid for CardExpirationDate
func testSHRCardExpirationDateMonth(t testing.TB) {
	mockBatch := mockBatchSHR()
	mockBatch.GetEntries()[0].SetSHRCardExpirationDate("1306")
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CardExpirationDate" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSHRCardExpirationDateMonth tests validating the month is valid for CardExpirationDate
func TestSHRSHRCardExpirationDateMonth(t *testing.T) {
	testSHRCardExpirationDateMonth(t)
}

// BenchmarkSHRCardExpirationDateMonth test validating the month is valid for CardExpirationDate
func BenchmarkSHRCardExpirationDateMonth(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testSHRCardExpirationDateMonth(b)
	}
}

// testSHRCardExpirationDateYear validates the year is valid for CardExpirationDate
func testSHRCardExpirationDateYear(t testing.TB) {
	mockBatch := mockBatchSHR()
	mockBatch.GetEntries()[0].SetSHRCardExpirationDate("0612")
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CardExpirationDate" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSHRCardExpirationDateYear tests validating the year is valid for CardExpirationDate
func TestSHRSHRCardExpirationDateYear(t *testing.T) {
	testSHRCardExpirationDateYear(t)
}

// BenchmarkSHRCardExpirationDateYear test validating the year is valid for CardExpirationDate
func BenchmarkSHRCardExpirationDateYear(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testSHRCardExpirationDateYear(b)
	}
}

// TestBatchSHRAddendum99Category validates Addenda99 returns an error
func TestBatchSHRAddendum99Category(t *testing.T) {
	mockBatch := NewBatchSHR(mockBatchSHRHeader())
	mockBatch.AddEntry(mockSHREntryDetail())
	mockAddenda99 := mockAddenda99()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	mockBatch.Entries[0].Category = CategoryNOC
	mockBatch.Entries[0].Addenda99 = mockAddenda99
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

// TestBatchSHRTerminalState validates TerminalState returns an error if invalid from usabbrev
func TestBatchSHRTerminalState(t *testing.T) {
	mockBatch := NewBatchSHR(mockBatchSHRHeader())
	mockBatch.AddEntry(mockSHREntryDetail())
	mockAddenda02 := mockAddenda02()
	mockAddenda02.TerminalState = "YY"
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TerminalState" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}
