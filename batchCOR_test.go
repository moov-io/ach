// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"log"
	"testing"
)

// mockBatchCORHeader creates a COR BatchHeader
func mockBatchCORHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.StandardEntryClassCode = COR
	bh.CompanyName = "Your Company, inc"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "Vendor Pay"
	bh.ODFIIdentification = "121042882"
	return bh
}

// mockCOREntryDetail creates a COR EntryDetail
func mockCOREntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingReturnNOCCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 0
	entry.IdentificationNumber = "location #23"
	entry.SetReceivingCompany("Best Co. #23")
	entry.SetTraceNumber(mockBatchCORHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "S"
	return entry
}

//  mockBatchCOR creates a BatchCOR
func mockBatchCOR() *BatchCOR {
	mockBatch := NewBatchCOR(mockBatchCORHeader())
	mockBatch.AddEntry(mockCOREntryDetail())
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	if err := mockBatch.Create(); err != nil {
		log.Fatal(err)
	}
	return mockBatch
}

// testBatchCORHeader creates a COR BatchHeader
func testBatchCORHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchCORHeader())

	_, ok := batch.(*BatchCOR)
	if !ok {
		t.Error("Expecting BachCOR")
	}
}

// TestBatchCORHeader tests creating a COR BatchHeader
func TestBatchCORHeader(t *testing.T) {
	testBatchCORHeader(t)
}

// BenchmarkBatchCORHeader benchmarks creating a COR BatchHeader
func BenchmarkBatchCORHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCORHeader(b)
	}
}

// testBatchCORSEC validates BatchCOR SEC code
func testBatchCORSEC(t testing.TB) {
	mockBatch := mockBatchCOR()
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.Header.StandardEntryClassCode = WEB
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

// TestBatchCORSEC tests validating BatchCOR SEC code
func TestBatchCORSEC(t *testing.T) {
	testBatchCORSEC(t)
}

// BenchmarkBatchCORSEC benchmarks validating BatchCOR SEC code
func BenchmarkBatchCORSEC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCORSEC(b)
	}
}

//  testBatchCORAddendumCountTwo validates Addendum count of 2
func testBatchCORAddendumCountTwo(t testing.TB) {
	mockBatch := mockBatchCOR()
	// Adding a second addenda to the mock entry
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98()
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

// TestBatchCORAddendumCountTwo tests validating Addendum count of 2
func TestBatchCORAddendumCountTwo(t *testing.T) {
	testBatchCORAddendumCountTwo(t)
}

// BenchmarkBatchCORAddendumCountTwo benchmarks validating Addendum count of 2
func BenchmarkBatchCORAddendumCountTwo(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCORAddendumCountTwo(b)
	}
}

// testBatchCORAddendaCountZero validates Addendum count of 0
func testBatchCORAddendaCountZero(t testing.TB) {
	mockBatch := NewBatchCOR(mockBatchCORHeader())
	mockBatch.AddEntry(mockCOREntryDetail())
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addenda98" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCORAddendaCountZero tests validating Addendum count of 0
func TestBatchCORAddendaCountZero(t *testing.T) {
	testBatchCORAddendaCountZero(t)
}

// BenchmarkBatchCORAddendaCountZero benchmarks validating Addendum count of 0
func BenchmarkBatchCORAddendaCountZero(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCORAddendaCountZero(b)
	}
}

// testBatchCORAddendaType validates that Addendum is of type Addenda98
func testBatchCORAddendaType(t testing.TB) {
	mockBatch := NewBatchCOR(mockBatchCORHeader())
	mockBatch.AddEntry(mockCOREntryDetail())
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addenda98" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCORAddendaType tests validating that Addendum is of type Addenda98
func TestBatchCORAddendaType(t *testing.T) {
	testBatchCORAddendaType(t)
}

// BenchmarkBatchCORAddendaType benchmarks validating that Addendum is of type Addenda98
func BenchmarkBatchCORAddendaType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCORAddendaType(b)
	}
}

// testBatchCORAddendaTypeCode validates TypeCode
func testBatchCORAddendaTypeCode(t testing.TB) {
	mockBatch := mockBatchCOR()
	mockBatch.GetEntries()[0].Addenda98.TypeCode = "07"
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCORAddendaTypeCode tests validating TypeCode
func TestBatchCORAddendaTypeCode(t *testing.T) {
	testBatchCORAddendaTypeCode(t)
}

// BenchmarkBatchCORAddendaTypeCode benchmarks validating TypeCode
func BenchmarkBatchCORAddendaTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCORAddendaTypeCode(b)
	}
}

// testBatchCORAmount validates BatchCOR Amount
func testBatchCORAmount(t testing.TB) {
	mockBatch := mockBatchCOR()
	mockBatch.GetEntries()[0].Amount = 9999
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Amount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCORAmount tests validating BatchCOR Amount
func TestBatchCORAmount(t *testing.T) {
	testBatchCORAmount(t)
}

// BenchmarkBatchCORAmount benchmarks validating BatchCOR Amount
func BenchmarkBatchCORAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCORAmount(b)
	}
}

// testBatchCORTransactionCode27 validates BatchCOR TransactionCode 27 returns an error
func testBatchCORTransactionCode27(t testing.TB) {
	mockBatch := mockBatchCOR()
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

// TestBatchCORTransactionCode27 tests validating BatchCOR TransactionCode 27 returns an error
func TestBatchCORTransactionCode27(t *testing.T) {
	testBatchCORTransactionCode27(t)
}

// BenchmarkBatchCORTransactionCode27 benchmarks validating
// BatchCOR TransactionCode 27 returns an error
func BenchmarkBatchCORTransactionCode27(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCORTransactionCode27(b)
	}
}

// testBatchCORTransactionCode21 validates BatchCOR TransactionCode 21 is a valid TransactionCode to be used for NOC
// mockBatch.Create() should not return an error for this test
func testBatchCORTransactionCode21(t testing.TB) {
	mockBatch := mockBatchCOR()
	mockBatch.GetEntries()[0].TransactionCode = CheckingReturnNOCCredit
	if err := mockBatch.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCORTransactionCode21 tests validating BatchCOR TransactionCode 21
func TestBatchCORTransactionCode21(t *testing.T) {
	testBatchCORTransactionCode21(t)
}

// BenchmarkBatchCORTransactionCode21 benchmarks validating BatchCOR TransactionCode 21
func BenchmarkBatchCORTransactionCode21(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCORTransactionCode21(b)
	}
}

// testBatchCORCreate creates BatchCOR
func testBatchCORCreate(t testing.TB) {
	mockBatch := mockBatchCOR()
	// Must have valid batch header to create a Batch
	mockBatch.GetHeader().ServiceClassCode = 63
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCORCreate tests creating BatchCOR
func TestBatchCORCreate(t *testing.T) {
	testBatchCORCreate(t)
}

// BenchmarkBatchCORCreate benchmarks creating BatchCOR
func BenchmarkBatchCORCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCORCreate(b)
	}
}

// testBatchCORServiceClassCodeEquality validates service class code equality
func testBatchCORServiceClassCodeEquality(t testing.TB) {
	mockBatch := mockBatchCOR()
	mockBatch.GetControl().ServiceClassCode = MixedDebitsAndCredits
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

// TestBatchCORServiceClassCodeEquality tests validating service class code equality
func TestBatchCORServiceClassCodeEquality(t *testing.T) {
	testBatchCORServiceClassCodeEquality(t)
}

// BenchmarkBatchCORServiceClassCodeEquality benchmarks validating service class code equality
func BenchmarkBatchCORServiceClassCodeEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCORServiceClassCodeEquality(b)
	}
}

// TestBatchCORCategoryNOCAddenda05 validates that an error is returned if valid Addenda05 is defined for CategoryNOC
func TestBatchCORCategoryNOCAddenda05(t *testing.T) {
	mockBatch := NewBatchCOR(mockBatchCORHeader())
	mockBatch.AddEntry(mockCOREntryDetail())
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98()
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

// TestBatchCORCategoryNOCAddenda02 validates that an error is returned if valid Addenda02 is defined for CategoryNOC
func TestBatchCORCategoryNOCAddenda02(t *testing.T) {
	mockBatch := NewBatchCOR(mockBatchCORHeader())
	mockBatch.AddEntry(mockCOREntryDetail())
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98()
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addenda02" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCORCategoryNOCAddenda98 validates that no error is returned if Addenda098 is defined for CategoryNOC
func TestBatchCORCategoryNOCAddenda98(t *testing.T) {
	mockBatch := NewBatchCOR(mockBatchCORHeader())
	mockBatch.AddEntry(mockCOREntryDetail())
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addenda98" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCORValidTranCodeForServiceClassCode validates a transactionCode based on ServiceClassCode
func TestBatchCORValidTranCodeForServiceClassCode(t *testing.T) {
	mockBatch := mockBatchCOR()
	mockBatch.GetHeader().ServiceClassCode = AutomatedAccountingAdvices
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCORInvalidAddenda98 validates that an error is returned if Addenda98 is invalid
func TestBatchCORTestBatchCORInvalidAddenda98(t *testing.T) {
	mockBatch := NewBatchCOR(mockBatchCORHeader())
	mockBatch.AddEntry(mockCOREntryDetail())
	mockBatch.GetEntries()[0].Category = CategoryNOC
	addenda98 := mockAddenda98()
	addenda98.recordType = "03"
	mockBatch.GetEntries()[0].Addenda98 = addenda98

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

// TestBatchCORTransactionCodeInvalid validates BatchCOR returns an error for an invalid TransactionCode
func TestBatchCORAutomatedAccountingAdvices(t *testing.T) {
	mockBatch := mockBatchCOR()
	mockBatch.GetEntries()[0].TransactionCode = 65
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
