// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"log"
	"testing"
)

// mockBatchATXHeader creates a BatchATX BatchHeader
func mockBatchATXHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.StandardEntryClassCode = "ATX"
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "231380104"
	bh.CompanyEntryDescription = "ACH ATX"
	bh.ODFIIdentification = "23138010"
	return bh
}

// mockATXEntryDetail creates a BatchATX EntryDetail
func mockATXEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingZeroDollarRemittanceCredit
	entry.SetRDFI("121042882")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 0
	entry.SetOriginalTraceNumber("121042880000001")
	entry.SetCATXAddendaRecords(1)
	entry.SetCATXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(mockBatchATXHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.AddendaRecordIndicator = 1
	entry.AddAddenda05(mockAddenda05())
	entry.Category = CategoryForward
	return entry
}

// mockBatchATX creates a BatchATX
func mockBatchATX() *BatchATX {
	mockBatch := NewBatchATX(mockBatchATXHeader())
	mockBatch.AddEntry(mockATXEntryDetail())
	if err := mockBatch.Create(); err != nil {
		log.Fatal(err)
	}
	return mockBatch
}

// testBatchATXHeader creates a BatchATX BatchHeader
func testBatchATXHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchATXHeader())
	err, ok := batch.(*BatchATX)
	if !ok {
		t.Errorf("Expecting BatchATX got %T", err)
	}
}

// TestBatchATXHeader tests validating BatchATX BatchHeader
func TestBatchATXHeader(t *testing.T) {
	testBatchATXHeader(t)
}

// BenchmarkBatchATXHeader benchmarks validating BatchATX BatchHeader
func BenchmarkBatchATXHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchATXHeader(b)
	}
}

// testBatchATXCreate validates BatchATX create
func testBatchATXCreate(t testing.TB) {
	mockBatch := mockBatchATX()
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchATXCreate tests validating BatchATX create
func TestBatchATXCreate(t *testing.T) {
	testBatchATXCreate(t)
}

// BenchmarkBatchATXCreate benchmarks validating BatchATX create
func BenchmarkBatchATXCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchATXCreate(b)
	}
}

// testBatchATXStandardEntryClassCode validates BatchATX create for an invalid StandardEntryClassCode
func testBatchATXStandardEntryClassCode(t testing.TB) {
	mockBatch := mockBatchATX()
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

// TestBatchATXStandardEntryClassCode tests validating BatchATX create for an invalid StandardEntryClassCode
func TestBatchATXStandardEntryClassCode(t *testing.T) {
	testBatchATXStandardEntryClassCode(t)
}

// BenchmarkBatchATXStandardEntryClassCode benchmarks validating BatchATX create for an invalid StandardEntryClassCode
func BenchmarkBatchATXStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchATXStandardEntryClassCode(b)
	}
}

// testBatchATXServiceClassCodeEquality validates service class code equality
func testBatchATXServiceClassCodeEquality(t testing.TB) {
	mockBatch := mockBatchATX()
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

// TestBatchATXServiceClassCodeEquality tests validating service class code equality
func TestBatchATXServiceClassCodeEquality(t *testing.T) {
	testBatchATXServiceClassCodeEquality(t)
}

// BenchmarkBatchATXServiceClassCodeEquality benchmarks validating service class code equality
func BenchmarkBatchATXServiceClassCodeEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchATXServiceClassCodeEquality(b)
	}
}

// testBatchATXAddendaCount validates BatchATX Addenda05 count of 2
func testBatchATXAddendaCount(t testing.TB) {
	mockBatch := mockBatchATX()
	mockBatch.GetEntries()[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
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

// TestBatchATXAddendaCount tests validating BatchATX Addenda05 count of 2
func TestBatchATXAddendaCount(t *testing.T) {
	testBatchATXAddendaCount(t)
}

// BenchmarkBatchATXAddendaCount benchmarks validating BatchATX Addendum count of 2
func BenchmarkBatchATXAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchATXAddendaCount(b)
	}
}

// testBatchATXAddendaCountZero validates Addendum count of 0
func testBatchATXAddendaCountZero(t testing.TB) {
	mockBatch := NewBatchATX(mockBatchATXHeader())
	mockBatch.AddEntry(mockATXEntryDetail())
	//mockBatch.GetEntries()[0].Addenda05[0].
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

// TestBatchATXAddendaCountZero tests validating Addendum count of 0
func TestBatchATXAddendaCountZero(t *testing.T) {
	testBatchATXAddendaCountZero(t)
}

// BenchmarkBatchATXAddendaCountZero benchmarks validating Addendum count of 0
func BenchmarkBatchATXAddendaCountZero(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchATXAddendaCountZero(b)
	}
}

// TestBatchATXInvalidAddenda02 validates Addenda must be Addenda05
func TestBatchATXInvalidAddend02(t *testing.T) {
	mockBatch := NewBatchATX(mockBatchATXHeader())
	entry := mockATXEntryDetail()
	entry.Addenda02 = mockAddenda02()
	mockBatch.AddEntry(entry)
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

// testBatchATXInvalidAddenda validates Addendum must be Addenda05 with record type 7
func testBatchATXInvalidAddenda(t testing.TB) {
	mockBatch := NewBatchATX(mockBatchATXHeader())
	mockBatch.AddEntry(mockATXEntryDetail())
	addenda05 := mockAddenda05()
	addenda05.recordType = "63"
	mockBatch.GetEntries()[0].AddAddenda05(addenda05)
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

// TestBatchATXInvalidAddenda tests validating Addendum must be Addenda05 with record type 7
func TestBatchATXInvalidAddenda(t *testing.T) {
	testBatchATXInvalidAddenda(t)
}

// BenchmarkBatchATXInvalidAddenda benchmarks validating Addendum must be Addenda05 with record type 7
func BenchmarkBatchATXInvalidAddenda(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchATXInvalidAddenda(b)
	}
}

// testBatchATXInvalidBuild validates an invalid batch build
func testBatchATXInvalidBuild(t testing.TB) {
	mockBatch := mockBatchATX()
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

// TestBatchATXInvalidBuild tests validating an invalid batch build
func TestBatchATXInvalidBuild(t *testing.T) {
	testBatchATXInvalidBuild(t)
}

// BenchmarkBatchATXInvalidBuild benchmarks validating an invalid batch build
func BenchmarkBatchATXInvalidBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchATXInvalidBuild(b)
	}
}

// testBatchATXAddenda10000 validates error for 10000 Addenda
func testBatchATXAddenda10000(t testing.TB) {

	bh := NewBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.StandardEntryClassCode = "ATX"
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "231380104"
	bh.CompanyEntryDescription = "ACH ATX"
	bh.ODFIIdentification = "23138010"

	entry := NewEntryDetail()
	entry.TransactionCode = CheckingZeroDollarRemittanceCredit
	entry.SetRDFI("121042882")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 0
	entry.SetOriginalTraceNumber("121042880000001")
	entry.SetCATXAddendaRecords(9999)
	entry.SetCATXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(mockBatchATXHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.Category = CategoryForward

	mockBatch := NewBatchATX(bh)
	mockBatch.AddEntry(entry)
	mockBatch.GetEntries()[0].AddendaRecordIndicator = 1

	for i := 0; i < 10000; i++ {
		mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())

	}

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

// TestBatchATXAddenda10000 tests validating error for 10000 Addenda
func TestBatchATXAddenda10000(t *testing.T) {
	testBatchATXAddenda10000(t)
}

// BenchmarkBatchATXAddenda10000 benchmarks validating error for 10000 Addenda
func BenchmarkBatchATXAddenda10000(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchATXAddenda10000(b)
	}
}

// testBatchATXAddendaRecords validates error for AddendaRecords not equal to addendum
func testBatchATXAddendaRecords(t testing.TB) {
	bh := NewBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.StandardEntryClassCode = "ATX"
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "231380104"
	bh.CompanyEntryDescription = "ACH ATX"
	bh.ODFIIdentification = "23138010"

	entry := NewEntryDetail()
	entry.TransactionCode = CheckingZeroDollarRemittanceCredit
	entry.SetRDFI("121042882")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 0
	entry.SetOriginalTraceNumber("121042880000001")
	entry.SetCATXAddendaRecords(565)
	entry.SetCATXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(mockBatchATXHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.Category = CategoryForward
	entry.AddendaRecordIndicator = 1

	mockBatch := NewBatchATX(bh)
	mockBatch.AddEntry(entry)

	for i := 0; i < 565; i++ {
		mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	}

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

// TestBatchATXAddendaRecords tests validating error for AddendaRecords not equal to addendum
func TestBatchATXAddendaRecords(t *testing.T) {
	testBatchATXAddendaRecords(t)
}

// BenchmarkBatchAddendaRecords benchmarks validating error for AddendaRecords not equal to addendum
func BenchmarkBatchATXAddendaRecords(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchATXAddendaRecords(b)
	}
}

// testBatchATXReceivingCompany validates ATXReceivingCompany
func testBatchATXReceivingCompany(t testing.TB) {
	mockBatch := mockBatchATX()
	//mockBatch.GetEntries()[0].SetCATXReceivingCompany("Receiver")

	if mockBatch.GetEntries()[0].CATXReceivingCompanyField() != "Receiver Company" {
		t.Errorf("expected %v got %v", "Receiver Company", mockBatch.GetEntries()[0].CATXReceivingCompanyField())
	}
}

// TestBatchATXReceivingCompany tests validating ATXReceivingCompany
func TestBatchATXReceivingCompany(t *testing.T) {
	testBatchATXReceivingCompany(t)
}

// BenchmarkBatchATXReceivingCompany benchmarks validating ATXReceivingCompany
func BenchmarkBatchATXReceivingCompany(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchATXReceivingCompany(b)
	}
}

// testBatchATXReserved validates ATXReservedField
func testBatchATXReserved(t testing.TB) {
	mockBatch := mockBatchATX()

	if mockBatch.GetEntries()[0].CATXReservedField() != "  " {
		t.Errorf("expected %v got %v", "  ", mockBatch.GetEntries()[0].CATXReservedField())
	}
}

// TestBatchATXReserved tests validating ATXReservedField
func TestBatchATXReserved(t *testing.T) {
	testBatchATXReserved(t)
}

// BenchmarkBatchATXReserved benchmarks validating ATXReservedField
func BenchmarkBatchATXReserved(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchATXReserved(b)
	}
}

// testBatchATXZeroAddendaRecords validates zero addenda records
func testBatchATXZeroAddendaRecords(t testing.TB) {
	bh := NewBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.StandardEntryClassCode = "ATX"
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "231380104"
	bh.CompanyEntryDescription = "ACH ATX"
	bh.ODFIIdentification = "23138010"

	entry := NewEntryDetail()
	entry.TransactionCode = CheckingZeroDollarRemittanceCredit
	entry.SetRDFI("121042882")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 0
	entry.SetOriginalTraceNumber("121042880000001")
	entry.SetCATXAddendaRecords(1)
	entry.SetCATXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(mockBatchATXHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.AddAddenda05(mockAddenda05())
	entry.AddendaRecordIndicator = 1
	entry.Category = CategoryForward

	mockBatch := NewBatchATX(bh)
	mockBatch.AddEntry(entry)

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

// TestBatchATXZeroAddendaRecords tests validating zero addenda records
func TestBatchATXZeroAddendaRecords(t *testing.T) {
	testBatchATXZeroAddendaRecords(t)
}

// BenchmarkBatchZeroAddendaRecords benchmarks validating zero addenda records
func BenchmarkBatchATXZeroAddendaRecords(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchATXZeroAddendaRecords(b)
	}
}

// testBatchATXTransactionCode validates TransactionCode
func testBatchATXTransactionCode(t testing.TB) {
	bh := NewBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.StandardEntryClassCode = "ATX"
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "231380104"
	bh.CompanyEntryDescription = "ACH ATX"
	bh.ODFIIdentification = "23138010"
	bh.OriginatorStatusCode = 2

	entry := NewEntryDetail()
	entry.TransactionCode = CheckingPrenoteCredit
	entry.SetRDFI("121042882")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 0
	entry.SetOriginalTraceNumber("121042880000001")
	entry.SetCATXAddendaRecords(1)
	entry.SetCATXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(mockBatchATXHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.Category = CategoryForward

	mockBatch := NewBatchATX(bh)
	mockBatch.AddEntry(entry)
	mockBatch.GetEntries()[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())

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

// TestBatchATXTransactionCode tests validating prenote addenda records
func TestBatchATXTransactionCode(t *testing.T) {
	testBatchATXTransactionCode(t)
}

// BenchmarkBatchATXTransactionCode benchmarks validating prenote addenda records
func BenchmarkBatchATXTransactionCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchATXTransactionCode(b)
	}
}

// TestBatchATXAmount validates Amount
func TestBatchATXAmount(t *testing.T) {
	bh := NewBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.StandardEntryClassCode = "ATX"
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "231380104"
	bh.CompanyEntryDescription = "ACH ATX"
	bh.ODFIIdentification = "23138010"
	bh.OriginatorStatusCode = 2

	entry := NewEntryDetail()
	entry.TransactionCode = CheckingPrenoteCredit
	entry.SetRDFI("121042882")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetOriginalTraceNumber("121042880000001")
	entry.SetCATXAddendaRecords(1)
	entry.SetCATXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(mockBatchATXHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.Category = CategoryForward

	mockBatch := NewBatchATX(bh)
	mockBatch.AddEntry(entry)
	mockBatch.GetEntries()[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())

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

// TestBatchATXAddendum98 validates Addenda98 returns an error
func TestBatchATXAddendum98(t *testing.T) {
	mockBatch := NewBatchATX(mockBatchATXHeader())
	mockBatch.AddEntry(mockATXEntryDetail())
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

// TestBatchATXAddendum99 validates Addenda99 returns an error
func TestBatchATXAddendum99(t *testing.T) {
	mockBatch := NewBatchATX(mockBatchATXHeader())
	mockBatch.AddEntry(mockATXEntryDetail())
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

// TestBatchATXValidTranCodeForServiceClassCode validates a transactionCode based on ServiceClassCode
func TestBatchATXValidTranCodeForServiceClassCode(t *testing.T) {
	mockBatch := mockBatchATX()
	mockBatch.GetHeader().ServiceClassCode = DebitsOnly
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
