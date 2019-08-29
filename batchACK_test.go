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
	"log"
	"testing"

	"github.com/moov-io/base"
)

// mockBatchACKHeader creates a ACK batch header
func mockBatchACKHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.StandardEntryClassCode = ACK
	bh.CompanyName = "Your Company, inc"
	bh.CompanyIdentification = "231380104"
	bh.CompanyEntryDescription = "Vndr Pay"
	bh.ODFIIdentification = "23138010"
	return bh
}

// mockACKEntryDetail creates a ACK entry detail
func mockACKEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingZeroDollarRemittanceCredit
	entry.SetRDFI("121042882")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 0
	entry.SetOriginalTraceNumber("121042880000001")
	entry.SetReceivingCompany("Best Co. #23")
	entry.SetTraceNumber(mockBatchACKHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "S"
	entry.AddendaRecordIndicator = 1
	entry.AddAddenda05(mockAddenda05())
	return entry
}

// mockBatchACK creates a ACK batch
func mockBatchACK() *BatchACK {
	mockBatch := NewBatchACK(mockBatchACKHeader())
	mockBatch.AddEntry(mockACKEntryDetail())
	if err := mockBatch.Create(); err != nil {
		log.Fatal(err)
	}
	return mockBatch
}

// testBatchACKHeader creates a ACK batch header
func testBatchACKHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchACKHeader())
	_, ok := batch.(*BatchACK)
	if !ok {
		t.Error("Expecting BatchACK")
	}
}

// TestBatchACKHeader tests creating a ACK batch header
func TestBatchACKHeader(t *testing.T) {
	testBatchACKHeader(t)
}

// BenchmarkBatchACKHeader benchmark creating a ACK batch header
func BenchmarkBatchACKHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchACKHeader(b)
	}
}

// testBatchACKAddendumCount batch control ACK can only have one addendum per entry detail
func testBatchACKAddendumCount(t testing.TB) {
	mockBatch := mockBatchACK()
	// Adding a second addenda to the mock entry
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchAddendaCount(2, 1)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchACKAddendumCount tests batch control ACK can only have one addendum per entry detail
func TestBatchACKAddendumCount(t *testing.T) {
	testBatchACKAddendumCount(t)
}

// BenchmarkBatchACKAddendumCount benchmarks batch control ACK can only have one addendum per entry detail
func BenchmarkBatchACKAddendumCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchACKAddendumCount(b)
	}
}

// TestBatchACKAddendum98 validates Addenda98 returns an error
func TestBatchACKAddendum98(t *testing.T) {
	mockBatch := NewBatchACK(mockBatchACKHeader())
	mockBatch.AddEntry(mockACKEntryDetail())
	mockAddenda98 := mockAddenda98()
	mockAddenda98.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchACKAddendum99 validates Addenda99 returns an error
func TestBatchACKAddendum99(t *testing.T) {
	mockBatch := NewBatchACK(mockBatchACKHeader())
	mockBatch.AddEntry(mockACKEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockAddenda99.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryReturn
	mockBatch.GetEntries()[0].Addenda99 = mockAddenda99
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// testBatchACKReceivingCompanyName validates Receiving company / Individual name is a mandatory field
func testBatchACKReceivingCompanyName(t testing.TB) {
	mockBatch := mockBatchACK()
	// modify the Individual name / receiving company to nothing
	mockBatch.GetEntries()[0].SetReceivingCompany("")
	err := mockBatch.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchACKReceivingCompanyName tests validating receiving company / Individual name is a mandatory field
func TestBatchACKReceivingCompanyName(t *testing.T) {
	testBatchACKReceivingCompanyName(t)
}

// BenchmarkBatchACKReceivingCompanyName benchmarks validating receiving company / Individual name is a mandatory field
func BenchmarkBatchACKReceivingCompanyName(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchACKReceivingCompanyName(b)
	}
}

// testBatchACKAddendaTypeCode validates addenda type code is 05
func testBatchACKAddendaTypeCode(t testing.TB) {
	mockBatch := mockBatchACK()
	mockBatch.GetEntries()[0].Addenda05[0].TypeCode = "07"
	err := mockBatch.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchACKAddendaTypeCode tests validating addenda type code is 05
func TestBatchACKAddendaTypeCode(t *testing.T) {
	testBatchACKAddendaTypeCode(t)
}

// BenchmarkBatchACKAddendaTypeCod benchmarks validating addenda type code is 05
func BenchmarkBatchACKAddendaTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchACKAddendaTypeCode(b)
	}
}

// testBatchACKSEC validates that the standard entry class code is ACK for batchACK
func testBatchACKSEC(t testing.TB) {
	mockBatch := mockBatchACK()
	mockBatch.Header.StandardEntryClassCode = RCK
	err := mockBatch.Validate()
	if !base.Match(err, ErrBatchSECType) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchACKSEC tests validating that the standard entry class code is ACK for batchACK
func TestBatchACKSEC(t *testing.T) {
	testBatchACKSEC(t)
}

// BenchmarkBatchACKSEC benchmarks validating that the standard entry class code is ACK for batch ACK
func BenchmarkBatchACKSEC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchACKSEC(b)
	}
}

// testBatchACKAddendaCount validates batch ACK addenda count
func testBatchACKAddendaCount(t testing.TB) {
	mockBatch := mockBatchACK()
	addenda05 := mockAddenda05()
	mockBatch.GetEntries()[0].AddAddenda05(addenda05)
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchAddendaCount(2, 1)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchACKAddendaCount tests validating batch ACK addenda count
func TestBatchACKAddendaCount(t *testing.T) {
	testBatchACKAddendaCount(t)
}

// BenchmarkBatchACKAddendaCount benchmarks validating batch ACK addenda count
func BenchmarkBatchACKAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchACKAddendaCount(b)
	}
}

// testBatchACKServiceClassCode validates ServiceClassCode
func testBatchACKServiceClassCode(t testing.TB) {
	mockBatch := mockBatchACK()
	// Batch Header information is required to Create a batch.
	mockBatch.GetHeader().ServiceClassCode = 0
	err := mockBatch.Create()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchACKServiceClassCode tests validating ServiceClassCode
func TestBatchACKServiceClassCode(t *testing.T) {
	testBatchACKServiceClassCode(t)
}

// BenchmarkBatchACKServiceClassCode benchmarks validating ServiceClassCode
func BenchmarkBatchACKServiceClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchACKServiceClassCode(b)
	}
}

// testBatchACKReceivingCompanyField validates ACKReceivingCompanyField
// underlying IndividualName
func testBatchACKReceivingCompanyField(t testing.TB) {
	mockBatch := mockBatchACK()
	ts := mockBatch.Entries[0].ReceivingCompanyField()
	if ts != "Best Co. #23          " {
		t.Error("Receiving Company Field is invalid")
	}
}

// TestBatchACKReceivingCompanyField tests validating ACKReceivingCompanyField
// underlying IndividualName
func TestBatchACKReceivingCompanyFieldField(t *testing.T) {
	testBatchACKReceivingCompanyField(t)
}

// BenchmarkBatchACKReceivingCompanyField benchmarks validating ACKReceivingCompanyField
// underlying IndividualName
func BenchmarkBatchACKReceivingCompanyField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchACKReceivingCompanyField(b)
	}
}

// TestBatchACKAmount validates Amount
func TestBatchACKAmount(t *testing.T) {
	mockBatch := mockBatchACK()
	// Batch Header information is required to Create a batch.
	mockBatch.GetEntries()[0].Amount = 25000
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAmountNonZero) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchACKTransactionCode validates TransactionCode
func TestBatchACKTransactionCode(t *testing.T) {
	mockBatch := mockBatchACK()
	// Batch Header information is required to Create a batch.
	mockBatch.GetEntries()[0].TransactionCode = CheckingCredit
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchTransactionCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchACKAddendum99Category validates Addenda99 returns an error
func TestBatchACKAddendum99Category(t *testing.T) {
	mockBatch := NewBatchACK(mockBatchACKHeader())
	mockBatch.AddEntry(mockACKEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockBatch.GetEntries()[0].Category = CategoryForward
	mockBatch.GetEntries()[0].Addenda99 = mockAddenda99
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAddendaCategory) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchACKValidTranCodeForServiceClassCode validates a transactionCode based on ServiceClassCode
func TestBatchACKValidTranCodeForServiceClassCode(t *testing.T) {
	mockBatch := mockBatchACK()
	mockBatch.GetHeader().ServiceClassCode = DebitsOnly
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchServiceClassTranCode(DebitsOnly, 24)) {
		t.Errorf("%T: %s", err, err)
	}
}
