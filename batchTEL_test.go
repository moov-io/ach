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

// mockBatchTELHeader creates a TEL batch header
func mockBatchTELHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = TEL
	bh.CompanyName = "Your Company, inc"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "Vndr Pay"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockTELEntryDetail creates a TEL entry detail
func mockTELEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingDebit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 5000000
	entry.IdentificationNumber = "Phone 333-2222"
	entry.IndividualName = "Wade Arnold"
	entry.SetTraceNumber(mockBatchTELHeader().ODFIIdentification, 1)
	entry.SetPaymentType("S")
	return entry
}

// mockBatchTEL creates a TEL batch
func mockBatchTEL() *BatchTEL {
	mockBatch := NewBatchTEL(mockBatchTELHeader())
	mockBatch.AddEntry(mockTELEntryDetail())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// testBatchTELHeader creates a TEL batch header
func testBatchTELHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchTELHeader())
	err, ok := batch.(*BatchTEL)
	if !ok {
		t.Errorf("Expecting BatchTEL got %T", err)
	}
}

// TestBatchTELHeader tests creating a TEL batch header
func TestBatchTELHeader(t *testing.T) {
	testBatchTELHeader(t)
}

// BenchmarkBatchTELHeader benchmarks creating a TEL batch header
func BenchmarkBatchTELHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTELHeader(b)
	}
}

// testBatchTELCreate validates batch create for an invalid service code
func testBatchTELCreate(t testing.TB) {
	mockBatch := mockBatchTEL()
	// Batch Header information is required to Create a batch.
	mockBatch.GetHeader().ServiceClassCode = 0
	err := mockBatch.Create()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTELCreate tests validating batch create for an invalid service code
func TestBatchTELCreate(t *testing.T) {
	testBatchTELCreate(t)
}

// BenchmarkBatchTELCreate benchmarks validating  batch create for an invalid service code
func BenchmarkBatchTELCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTELCreate(b)
	}
}

// testBatchTELAddendaCount validates addenda count for batch TEL
func testBatchTELAddendaCount(t testing.TB) {
	mockBatch := mockBatchTEL()
	// TEL can not have an addenda02
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchCalculatedControlEquality(2, 1)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTELAddendaCount tests validating addenda count for batch TEL
func TestBatchTELAddendaCount(t *testing.T) {
	testBatchTELAddendaCount(t)
}

// BenchmarkBatchTELAddendaCount benchmarks validating addenda count for batch TEL
func BenchmarkBatchTELAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTELAddendaCount(b)
	}
}

// testBatchTELSEC validates SEC code for batch TEL
func testBatchTELSEC(t testing.TB) {
	mockBatch := mockBatchTEL()
	mockBatch.Header.StandardEntryClassCode = RCK
	err := mockBatch.Validate()
	if !base.Match(err, ErrBatchSECType) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTELSEC tests validating SEC code for batch TEL
func TestBatchTELSEC(t *testing.T) {
	testBatchTELSEC(t)
}

// BenchmarkBatchTELSEC benchmarks validating SEC code for batch TEL
func BenchmarkBatchTELSEC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTELSEC(b)
	}
}

// testBatchTELDebit validates Transaction code for TEL entry detail
func testBatchTELDebit(t testing.TB) {
	mockBatch := mockBatchTEL()
	mockBatch.GetEntries()[0].TransactionCode = CheckingCredit
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchDebitOnly) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTELDebit tests validating Transaction code for TEL entry detail
func TestBatchTELDebit(t *testing.T) {
	testBatchTELDebit(t)
}

// BenchmarkBatchTELDebit benchmarks validating Transaction code for TEL entry detail
func BenchmarkBatchTELDebit(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTELDebit(b)
	}
}

// testBatchTELPaymentType validates that the entry detail
// payment type / discretionary data is either single or reoccurring
func testBatchTELPaymentType(t testing.TB) {
	mockBatch := mockBatchTEL()
	mockBatch.GetEntries()[0].DiscretionaryData = "AA"
	err := mockBatch.Validate()
	// TODO: are we expecting there to be no errors here?
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTELPaymentType tests validating that the entry detail
// payment type / discretionary data is either single or reoccurring
func TestBatchTELPaymentType(t *testing.T) {
	testBatchTELPaymentType(t)
}

// BenchmarkBatchTELPaymentTyp benchmarks validating that the entry detail
// payment type / discretionary data is either single or reoccurring
func BenchmarkBatchTELPaymentType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTELPaymentType(b)
	}
}

// TestBatchTELAddendum98 validates Addenda98 returns an error
func TestBatchTELAddendum98(t *testing.T) {
	mockBatch := NewBatchTEL(mockBatchTELHeader())
	mockBatch.AddEntry(mockTELEntryDetail())
	mockAddenda98 := mockAddenda98()
	mockAddenda98.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTELAddendum99 validates Addenda99 returns an error
func TestBatchTELAddendum99(t *testing.T) {
	mockBatch := NewBatchTEL(mockBatchTELHeader())
	mockBatch.AddEntry(mockTELEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockAddenda99.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryReturn
	mockBatch.GetEntries()[0].Addenda99 = mockAddenda99
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchTELValidTranCodeForServiceClassCode validates a transactionCode based on ServiceClassCode
func TestBatchTELValidTranCodeForServiceClassCode(t *testing.T) {
	mockBatch := mockBatchTEL()
	mockBatch.GetHeader().ServiceClassCode = CreditsOnly
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchServiceClassTranCode(CreditsOnly, 27)) {
		t.Errorf("%T: %s", err, err)
	}
}
