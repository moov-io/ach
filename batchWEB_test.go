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

// mockBatchWEBHeader creates a WEB batch header
func mockBatchWEBHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.StandardEntryClassCode = WEB
	bh.CompanyName = "Your Company, inc"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "Online Order"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockWEBEntryDetail creates a WEB entry detail
func mockWEBEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "123456789"
	entry.Amount = 100000000
	entry.IndividualName = "Wade Arnold"
	entry.SetTraceNumber(mockBatchWEBHeader().ODFIIdentification, 1)
	entry.SetPaymentType("S")
	return entry
}

// mockBatchWEB creates a WEB batch
func mockBatchWEB() *BatchWEB {
	mockBatch := NewBatchWEB(mockBatchWEBHeader())
	mockBatch.AddEntry(mockWEBEntryDetail())
	mockBatch.GetEntries()[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// testBatchWebAddenda validates No more than 1 batch per entry detail record can exist
// and no more than 1 addenda record per entry detail record can exist
func testBatchWebAddenda(t testing.TB) {
	mockBatch := mockBatchWEB()
	// mock batch already has one addenda. Creating two addenda should error
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchAddendaCount(2, 1)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchWebAddenda tests validating no more than 1 batch per entry detail
// record can exist and no more than 1 addenda record per entry detail record can exist
func TestBatchWebAddenda(t *testing.T) {
	testBatchWebAddenda(t)
}

// BenchmarkBatchWebAddenda benchmarks validating no more than 1 batch per entry detail
// record can exist and no more than 1 addenda record per entry detail record can exist
func BenchmarkBatchWebAddenda(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchWebAddenda(b)
	}
}

// testBatchWebIndividualNameRequired validates Individual name is a mandatory field
func testBatchWebIndividualNameRequired(t testing.TB) {
	mockBatch := mockBatchWEB()
	// mock batch already has one addenda. Creating two addenda should error
	mockBatch.GetEntries()[0].IndividualName = ""
	err := mockBatch.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchWebIndividualNameRequired tests validating Individual name is a mandatory field
func TestBatchWebIndividualNameRequired(t *testing.T) {
	testBatchWebIndividualNameRequired(t)
}

// BenchmarkBatchWebIndividualNameRequired benchmarks validating Individual name is a mandatory field
func BenchmarkBatchWebIndividualNameRequired(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchWebIndividualNameRequired(b)
	}
}

// TestBatchWEBAddendum98 validates Addenda98 returns an error
func TestBatchWEBAddendum98(t *testing.T) {
	mockBatch := NewBatchWEB(mockBatchWEBHeader())
	mockBatch.AddEntry(mockWEBEntryDetail())
	mockAddenda98 := mockAddenda98()
	mockAddenda98.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchWEBAddendum99 validates Addenda99 returns an error
func TestBatchWEBAddendum99(t *testing.T) {
	mockBatch := NewBatchWEB(mockBatchWEBHeader())
	mockBatch.AddEntry(mockWEBEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockAddenda99.TypeCode = "05"
	mockBatch.GetEntries()[0].Category = CategoryReturn
	mockBatch.GetEntries()[0].Addenda99 = mockAddenda99
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// testBatchWEBAddendaTypeCode validates addenda type code is valid
func testBatchWEBAddendaTypeCode(t testing.TB) {
	mockBatch := mockBatchWEB()
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	mockBatch.GetEntries()[0].Addenda05[0].TypeCode = "02"
	err := mockBatch.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchWEBAddendaTypeCode tests validating addenda type code is valid
func TestBatchWEBAddendaTypeCode(t *testing.T) {
	testBatchWEBAddendaTypeCode(t)
}

// BenchmarkBatchWEBAddendaTypeCode benchmarks validating addenda type code is valid
func BenchmarkBatchWEBAddendaTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchWEBAddendaTypeCode(b)
	}
}

// testBatchWebSEC validates that the standard entry class code is WEB for batch Web
func testBatchWebSEC(t testing.TB) {
	mockBatch := mockBatchWEB()
	mockBatch.Header.StandardEntryClassCode = RCK
	err := mockBatch.Validate()
	if !base.Match(err, ErrBatchSECType) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchWebSEC tests validating that the standard entry class code is WEB for batch WEB
func TestBatchWebSEC(t *testing.T) {
	testBatchWebSEC(t)
}

// BenchmarkBatchWebSEC benchmarks validating that the standard entry class code is WEB for batch WEB
func BenchmarkBatchWebSEC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchWebSEC(b)
	}
}

// testBatchWebPaymentType validates that the entry detail
// payment type / discretionary data is either single or reoccurring
func testBatchWebPaymentType(t testing.TB) {
	mockBatch := mockBatchWEB()
	mockBatch.GetEntries()[0].DiscretionaryData = "AA"
	err := mockBatch.Validate()
	// TODO: are we expecting there to be no errors here?
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchWebPaymentType tests validating that the entry detail
// payment type / discretionary data is either single or reoccurring
func TestBatchWebPaymentType(t *testing.T) {
	testBatchWebPaymentType(t)
}

// BenchmarkBatchWebPaymentType benchmarks validating that the entry detail
// payment type / discretionary data is either single or reoccurring
func BenchmarkBatchWebPaymentType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchWebPaymentType(b)
	}
}

// testBatchWebCreate creates a WEB batch
func testBatchWebCreate(t testing.TB) {
	mockBatch := mockBatchWEB()
	// Must have valid batch header to create a batch
	mockBatch.GetHeader().ServiceClassCode = 63
	err := mockBatch.Create()
	if !base.Match(err, ErrServiceClass) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchWebCreate tests creating a WEB batch
func TestBatchWebCreate(t *testing.T) {
	testBatchWebCreate(t)
}

// BenchmarkBatchWebCreate benchmarks creating a WEB batch
func BenchmarkBatchWebCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchWebCreate(b)
	}
}

// testBatchWebPaymentTypeR validates that the entry detail
// payment type / discretionary data is reoccurring
func testBatchWebPaymentTypeR(t testing.TB) {
	mockBatch := mockBatchWEB()
	mockBatch.GetEntries()[0].DiscretionaryData = "R"
	err := mockBatch.Validate()
	// TODO: are we expecting there to be no errors here?
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
	if mockBatch.GetEntries()[0].PaymentTypeField() != "R" {
		t.Errorf("PaymentTypeField %v was expecting R", mockBatch.GetEntries()[0].PaymentTypeField())
	}
}

// TestBatchWebPaymentTypeR tests validating that the entry detail
// payment type / discretionary data is reoccurring
func TestBatchWebPaymentTypeR(t *testing.T) {
	testBatchWebPaymentTypeR(t)
}

// BenchmarkBatchWebPaymentTypeR benchmarks validating that the entry detail
// payment type / discretionary data is reoccurring
func BenchmarkBatchWebPaymentTypeR(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchWebPaymentTypeR(b)
	}
}

// TestBatchWEBAddendum99Category validates Addenda99 returns an error
func TestBatchWEBAddendum99Category(t *testing.T) {
	mockBatch := NewBatchWEB(mockBatchWEBHeader())
	mockBatch.AddEntry(mockWEBEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	mockBatch.Entries[0].Category = CategoryForward
	mockBatch.Entries[0].Addenda99 = mockAddenda99
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAddendaCategory) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchWEBCategoryReturn validates CategoryReturn returns an error
func TestBatchWEBCategoryReturn(t *testing.T) {
	mockBatch := NewBatchWEB(mockBatchWEBHeader())
	mockBatch.AddEntry(mockWEBEntryDetail())
	mockAddenda05 := mockAddenda05()
	mockBatch.GetEntries()[0].Category = CategoryReturn
	mockBatch.GetEntries()[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05)
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAddendaCategory) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchWEBCategoryReturnAddenda98 validates CategoryReturn returns an error
func TestBatchWEBCategoryReturnAddenda98(t *testing.T) {
	mockBatch := NewBatchWEB(mockBatchWEBHeader())
	mockBatch.AddEntry(mockWEBEntryDetail())
	mockAddenda98 := mockAddenda98()
	mockBatch.GetEntries()[0].Category = CategoryReturn
	mockBatch.GetEntries()[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].Addenda98 = mockAddenda98
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAddendaCategory) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchWEBValidTranCodeForServiceClassCode validates a transactionCode based on ServiceClassCode
func TestBatchWEBValidTranCodeForServiceClassCode(t *testing.T) {
	mockBatch := mockBatchWEB()
	mockBatch.GetHeader().ServiceClassCode = DebitsOnly
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchServiceClassTranCode(DebitsOnly, 22)) {
		t.Errorf("%T: %s", err, err)
	}
}
