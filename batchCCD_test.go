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

// mockBatchCCDHeader creates a CCD batch header
func mockBatchCCDHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = CCD
	bh.CompanyName = "Your Company, inc"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "Vndr Pay"
	bh.ODFIIdentification = "121042882"
	return bh
}

// mockCCDEntryDetail creates a CCD entry detail
func mockCCDEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingDebit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 5000000
	entry.IdentificationNumber = "location #23"
	entry.SetReceivingCompany("Best Co. #23")
	entry.SetTraceNumber(mockBatchCCDHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "S"
	return entry
}

// mockBatchCCD creates a CCD batch
func mockBatchCCD() *BatchCCD {
	mockBatch := NewBatchCCD(mockBatchCCDHeader())
	mockBatch.AddEntry(mockCCDEntryDetail())
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// testBatchCCDHeader creates a CCD batch header
func testBatchCCDHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchCCDHeader())
	_, ok := batch.(*BatchCCD)
	if !ok {
		t.Error("Expecting BatchCCD")
	}
}

// TestBatchCCDHeader tests creating a CCD batch header
func TestBatchCCDHeader(t *testing.T) {
	testBatchCCDHeader(t)
}

// BenchmarkBatchCCDHeader benchmark creating a CCD batch header
func BenchmarkBatchCCDHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCCDHeader(b)
	}
}

// testBatchCCDAddendumCount batch control CCD can only have one addendum per entry detail
func testBatchCCDAddendumCount(t testing.TB) {
	mockBatch := mockBatchCCD()
	// Adding a second addenda to the mock entry
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchCalculatedControlEquality(3, 2)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCCDAddendumCount tests batch control CCD can only have one addendum per entry detail
func TestBatchCCDAddendumCount(t *testing.T) {
	testBatchCCDAddendumCount(t)
}

// BenchmarkBatchCCDAddendumCount benchmarks batch control CCD can only have one addendum per entry detail
func BenchmarkBatchCCDAddendumCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCCDAddendumCount(b)
	}
}

// TestBatchCCDAddendum98 validates Addenda98 returns an error
func TestBatchCCDAddendum98(t *testing.T) {
	mockBatch := NewBatchCCD(mockBatchCCDHeader())
	mockBatch.AddEntry(mockCCDEntryDetail())
	mockAddenda98 := mockAddenda98()
	mockAddenda98.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCCDAddendum99 validates Addenda99 returns an error
func TestBatchCCDAddendum99(t *testing.T) {
	mockBatch := NewBatchCCD(mockBatchCCDHeader())
	mockBatch.AddEntry(mockCCDEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockAddenda99.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryReturn
	mockBatch.GetEntries()[0].Addenda99 = mockAddenda99
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// testBatchCCDReceivingCompanyName validates Receiving company / Individual name is a mandatory field
func testBatchCCDReceivingCompanyName(t testing.TB) {
	mockBatch := mockBatchCCD()
	// modify the Individual name / receiving company to nothing
	mockBatch.GetEntries()[0].SetReceivingCompany("")
	err := mockBatch.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCCDReceivingCompanyName tests validating receiving company / Individual name is a mandatory field
func TestBatchCCDReceivingCompanyName(t *testing.T) {
	testBatchCCDReceivingCompanyName(t)
}

// BenchmarkBatchCCDReceivingCompanyName benchmarks validating receiving company / Individual name is a mandatory field
func BenchmarkBatchCCDReceivingCompanyName(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCCDReceivingCompanyName(b)
	}
}

// testBatchCCDAddendaTypeCode validates addenda type code is 05
func testBatchCCDAddendaTypeCode(t testing.TB) {
	mockBatch := mockBatchCCD()
	mockBatch.GetEntries()[0].Addenda05[0].TypeCode = "07"
	err := mockBatch.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCCDAddendaTypeCode tests validating addenda type code is 05
func TestBatchCCDAddendaTypeCode(t *testing.T) {
	testBatchCCDAddendaTypeCode(t)
}

// BenchmarkBatchCCDAddendaTypeCod benchmarks validating addenda type code is 05
func BenchmarkBatchCCDAddendaTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCCDAddendaTypeCode(b)
	}
}

// testBatchCCDSEC validates that the standard entry class code is CCD for batchCCD
func testBatchCCDSEC(t testing.TB) {
	mockBatch := mockBatchCCD()
	mockBatch.Header.StandardEntryClassCode = RCK
	err := mockBatch.Validate()
	if !base.Match(err, ErrBatchSECType) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCCDSEC tests validating that the standard entry class code is CCD for batchCCD
func TestBatchCCDSEC(t *testing.T) {
	testBatchCCDSEC(t)
}

// BenchmarkBatchCCDSEC benchmarks validating that the standard entry class code is CCD for batch CCD
func BenchmarkBatchCCDSEC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCCDSEC(b)
	}
}

// testBatchCCDAddendaCount validates batch CCD addenda count
func testBatchCCDAddendaCount(t testing.TB) {
	mockBatch := mockBatchCCD()
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchAddendaCount(0, 1)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCCDAddendaCount tests validating batch CCD addenda count
func TestBatchCCDAddendaCount(t *testing.T) {
	testBatchCCDAddendaCount(t)
}

// BenchmarkBatchCCDAddendaCount benchmarks validating batch CCD addenda count
func BenchmarkBatchCCDAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCCDAddendaCount(b)
	}
}

// testBatchCCDCreate creates a batch CCD
func testBatchCCDCreate(t testing.TB) {
	mockBatch := mockBatchCCD()
	// Batch Header information is required to Create a batch.
	mockBatch.GetHeader().ServiceClassCode = 0
	err := mockBatch.Create()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCCDCreate Test creating a batch CCD
func TestBatchCCDCreate(t *testing.T) {
	testBatchCCDCreate(t)
}

// BenchmarkBatchCCDCreate benchmark creating a batch CCD
func BenchmarkBatchCCDCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCCDCreate(b)
	}
}

// testBatchCCDReceivingCompanyField validates CCDReceivingCompanyField
// underlying IndividualName
func testBatchCCDReceivingCompanyField(t testing.TB) {
	mockBatch := mockBatchCCD()
	ts := mockBatch.Entries[0].ReceivingCompanyField()
	if ts != "Best Co. #23          " {
		t.Error("Receiving Company Field is invalid")
	}
}

// TestBatchCCDReceivingCompanyField tests validating CCDReceivingCompanyField
// underlying IndividualName
func TestBatchCCDReceivingCompanyFieldField(t *testing.T) {
	testBatchCCDReceivingCompanyField(t)
}

// BenchmarkBatchCCDReceivingCompanyField benchmarks validating CCDReceivingCompanyField
// underlying IndividualName
func BenchmarkBatchCCDReceivingCompanyField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCCDReceivingCompanyField(b)
	}
}

// TestBatchCCDValidTranCodeForServiceClassCode validates a transactionCode based on ServiceClassCode
func TestBatchCCDValidTranCodeForServiceClassCode(t *testing.T) {
	mockBatch := mockBatchCCD()
	mockBatch.GetHeader().ServiceClassCode = CreditsOnly
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchServiceClassTranCode(CreditsOnly, 27)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCCDAddenda02 validates BatchCCD cannot have Addenda02
func TestBatchCCDAddenda02(t *testing.T) {
	mockBatch := mockBatchCCD()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02()
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAddendaCategory) {
		t.Errorf("%T: %s", err, err)
	}
}
