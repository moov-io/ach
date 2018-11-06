// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"log"
	"testing"
)

// mockBatchADVHeader creates a ADV batch header
func mockBatchADVHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 220
	bh.StandardEntryClassCode = "ADV"
	bh.CompanyName = "Your Company, inc"
	bh.CompanyIdentification = "231380104"
	bh.CompanyEntryDescription = "Vndr Pay"
	bh.ODFIIdentification = "23138010"
	return bh
}

// mockADVEntryDetail creates a ADV entry detail
func mockADVEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 24
	entry.SetRDFI("121042882")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 0
	entry.SetOriginalTraceNumber("121042880000001")
	entry.SetReceivingCompany("Best Co. #23")
	entry.SetTraceNumber(mockBatchADVHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "S"
	entry.AddendaRecordIndicator = 1
	entry.AddAddenda05(mockAddenda05())
	return entry
}

// mockBatchADV creates a ADV batch
func mockBatchADV() *BatchADV {
	mockBatch := NewBatchADV(mockBatchADVHeader())
	mockBatch.AddEntry(mockADVEntryDetail())
	if err := mockBatch.Create(); err != nil {
		log.Fatal(err)
	}
	return mockBatch
}

// testBatchADVHeader creates a ADV batch header
func testBatchADVHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchADVHeader())
	_, ok := batch.(*BatchADV)
	if !ok {
		t.Error("Expecting BatchADV")
	}
}

// TestBatchADVHeader tests creating a ADV batch header
func TestBatchADVHeader(t *testing.T) {
	testBatchADVHeader(t)
}

// BenchmarkBatchADVHeader benchmark creating a ADV batch header
func BenchmarkBatchADVHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchADVHeader(b)
	}
}

// testBatchADVAddendumCount batch control ADV can only have one addendum per entry detail
func testBatchADVAddendumCount(t testing.TB) {
	mockBatch := mockBatchADV()
	// Adding a second addenda to the mock entry
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
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

// TestBatchADVAddendumCount tests batch control ADV can only have one addendum per entry detail
func TestBatchADVAddendumCount(t *testing.T) {
	testBatchADVAddendumCount(t)
}

// BenchmarkBatchADVAddendumCount benchmarks batch control ADV can only have one addendum per entry detail
func BenchmarkBatchADVAddendumCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchADVAddendumCount(b)
	}
}

// TestBatchADVAddendum98 validates Addenda98 returns an error
func TestBatchADVAddendum98(t *testing.T) {
	mockBatch := NewBatchADV(mockBatchADVHeader())
	mockBatch.AddEntry(mockADVEntryDetail())
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

// TestBatchADVAddendum99 validates Addenda99 returns an error
func TestBatchADVAddendum99(t *testing.T) {
	mockBatch := NewBatchADV(mockBatchADVHeader())
	mockBatch.AddEntry(mockADVEntryDetail())
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

// testBatchADVReceivingCompanyName validates Receiving company / Individual name is a mandatory field
func testBatchADVReceivingCompanyName(t testing.TB) {
	mockBatch := mockBatchADV()
	// modify the Individual name / receiving company to nothing
	mockBatch.GetEntries()[0].SetReceivingCompany("")
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "IndividualName" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchADVReceivingCompanyName tests validating receiving company / Individual name is a mandatory field
func TestBatchADVReceivingCompanyName(t *testing.T) {
	testBatchADVReceivingCompanyName(t)
}

// BenchmarkBatchADVReceivingCompanyName benchmarks validating receiving company / Individual name is a mandatory field
func BenchmarkBatchADVReceivingCompanyName(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchADVReceivingCompanyName(b)
	}
}

// testBatchADVAddendaTypeCode validates addenda type code is 05
func testBatchADVAddendaTypeCode(t testing.TB) {
	mockBatch := mockBatchADV()
	mockBatch.GetEntries()[0].Addenda05[0].TypeCode = "07"
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

// TestBatchADVAddendaTypeCode tests validating addenda type code is 05
func TestBatchADVAddendaTypeCode(t *testing.T) {
	testBatchADVAddendaTypeCode(t)
}

// BenchmarkBatchADVAddendaTypeCod benchmarks validating addenda type code is 05
func BenchmarkBatchADVAddendaTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchADVAddendaTypeCode(b)
	}
}

// testBatchADVSEC validates that the standard entry class code is ADV for batchADV
func testBatchADVSEC(t testing.TB) {
	mockBatch := mockBatchADV()
	mockBatch.Header.StandardEntryClassCode = "RCK"
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

// TestBatchADVSEC tests validating that the standard entry class code is ADV for batchADV
func TestBatchADVSEC(t *testing.T) {
	testBatchADVSEC(t)
}

// BenchmarkBatchADVSEC benchmarks validating that the standard entry class code is ADV for batch ADV
func BenchmarkBatchADVSEC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchADVSEC(b)
	}
}

// testBatchADVAddendaCount validates batch ADV addenda count
func testBatchADVAddendaCount(t testing.TB) {
	mockBatch := mockBatchADV()
	addenda05 := mockAddenda05()
	mockBatch.GetEntries()[0].AddAddenda05(addenda05)
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

// TestBatchADVAddendaCount tests validating batch ADV addenda count
func TestBatchADVAddendaCount(t *testing.T) {
	testBatchADVAddendaCount(t)
}

// BenchmarkBatchADVAddendaCount benchmarks validating batch ADV addenda count
func BenchmarkBatchADVAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchADVAddendaCount(b)
	}
}

// testBatchADVServiceClassCode validates ServiceClassCode
func testBatchADVServiceClassCode(t testing.TB) {
	mockBatch := mockBatchADV()
	// Batch Header information is required to Create a batch.
	mockBatch.GetHeader().ServiceClassCode = 0
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

// TestBatchADVServiceClassCode tests validating ServiceClassCode
func TestBatchADVServiceClassCode(t *testing.T) {
	testBatchADVServiceClassCode(t)
}

// BenchmarkBatchADVServiceClassCode benchmarks validating ServiceClassCode
func BenchmarkBatchADVServiceClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchADVServiceClassCode(b)
	}
}

// testBatchADVReceivingCompanyField validates ADVReceivingCompanyField
// underlying IndividualName
func testBatchADVReceivingCompanyField(t testing.TB) {
	mockBatch := mockBatchADV()
	ts := mockBatch.Entries[0].ReceivingCompanyField()
	if ts != "Best Co. #23          " {
		t.Error("Receiving Company Field is invalid")
	}
}

// TestBatchADVReceivingCompanyField tests validating ADVReceivingCompanyField
// underlying IndividualName
func TestBatchADVReceivingCompanyFieldField(t *testing.T) {
	testBatchADVReceivingCompanyField(t)
}

// BenchmarkBatchADVReceivingCompanyField benchmarks validating ADVReceivingCompanyField
// underlying IndividualName
func BenchmarkBatchADVReceivingCompanyField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchADVReceivingCompanyField(b)
	}
}

// TestBatchADVAmount validates Amount
func TestBatchADVAmount(t *testing.T) {
	mockBatch := mockBatchADV()
	// Batch Header information is required to Create a batch.
	mockBatch.GetEntries()[0].Amount = 25000
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

// TestBatchADVTransactionCode validates Amount
func TestBatchADVTransactionCode(t *testing.T) {
	mockBatch := mockBatchADV()
	// Batch Header information is required to Create a batch.
	mockBatch.GetEntries()[0].TransactionCode = 22
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TransactionCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchADVAddendum99Category validates Addenda99 returns an error
func TestBatchADVAddendum99Category(t *testing.T) {
	mockBatch := NewBatchADV(mockBatchADVHeader())
	mockBatch.AddEntry(mockADVEntryDetail())
	mockAddenda99 := mockAddenda99()
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
