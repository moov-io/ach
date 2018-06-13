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
	entry.SetReceivingCompany("ABC Company")
	entry.SetTraceNumber(mockBatchPOSHeader().ODFIIdentification, 123)
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

