// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"
)

// TODO make all the mock values cor fields

// mockBatchCORHeader creates a COR BatchHeader
func mockBatchCORHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 220
	bh.StandardEntryClassCode = "COR"
	bh.CompanyName = "Your Company, inc"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "Vendor Pay"
	bh.ODFIIdentification = "121042882"
	return bh
}

// mockCOREntryDetail creates a COR EntryDetail
func mockCOREntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 21
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
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda98())
	if err := mockBatch.Create(); err != nil {
		panic(err)
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
	mockBatch.Header.StandardEntryClassCode = "COR"
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
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda98())
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
			if e.FieldName != "Addendum" {
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
	mockBatch.GetEntries()[0].Addendum[0].(*Addenda98).typeCode = "07"
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
	mockBatch.GetEntries()[0].TransactionCode = 27
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

// testBatchCORTransactionCode21 validates BatchCOR TransactionCode 21
func testBatchCORTransactionCode21(t testing.TB) {
	mockBatch := mockBatchCOR()
	mockBatch.GetEntries()[0].TransactionCode = 21
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
