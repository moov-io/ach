package ach

import (
	"testing"
)

func mockBatchCCDHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 220
	bh.StandardEntryClassCode = "CCD"
	bh.CompanyName = "Your Company, inc"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "Vndr Pay"
	bh.ODFIIdentification = "121042882"
	return bh
}

func mockCCDEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 27
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 5000000
	entry.IdentificationNumber = "location #23"
	entry.SetReceivingCompany("Best Co. #23")
	entry.SetTraceNumber(mockBatchCCDHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "S"
	return entry
}

func mockBatchCCD() *BatchCCD {
	mockBatch := NewBatchCCD(mockBatchCCDHeader())
	mockBatch.AddEntry(mockCCDEntryDetail())
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda05())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// Batch CCD Header
func testBatchCCDHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchCCDHeader())
	_, ok := batch.(*BatchCCD)
	if !ok {
		t.Error("Expecting BatchCCD")
	}
}

// Test Batch CCD Header
func TestBatchCCDHeader(t *testing.T) {
	testBatchCCDHeader(t)
}

// Benchmark Batch CCD Header
func BenchmarkBatchCCDHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCCDHeader(b)
	}
}

// Batch control CCD can only have one addendum per entry detail
func testBatchCCDAddendumCount(t testing.TB) {
	mockBatch := mockBatchCCD()
	// Adding a second addenda to the mock entry
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda05())
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "EntryAddendaCount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// Test batch control CCD can only have one addendum per entry detail
func TestBatchCCDAddendumCount(t *testing.T) {
	testBatchCCDAddendumCount(t)
}

// Benchmark batch control CCD can only have one addendum per entry detail
func BenchmarkBatchCCDAddendumCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCCDAddendumCount(b)
	}
}

// Receiving company / Individual name is a mandatory field
func testBatchCCDReceivingCompanyName(t testing.TB) {
	mockBatch := mockBatchCCD()
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

// Test receiving company / Individual name is a mandatory field
func TestBatchCCDReceivingCompanyName(t *testing.T) {
	testBatchCCDReceivingCompanyName(t)
}

// Benchmark receiving company / Individual name is a mandatory field
func BenchmarkBatchCCDReceivingCompanyName(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCCDReceivingCompanyName(b)
	}
}

// Verify addenda type code is 05
func testBatchCCDAddendaTypeCode(t testing.TB) {
	mockBatch := mockBatchCCD()
	mockBatch.GetEntries()[0].Addendum[0].(*Addenda05).typeCode = "07"
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

// Test verify addenda type code is 05
func TestBatchCCDAddendaTypeCode(t *testing.T) {
	testBatchCCDAddendaTypeCode(t)
}

// Benchmark verify addenda type code is 05
func BenchmarkBatchCCDAddendaTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCCDAddendaTypeCode(b)
	}
}

// Verify that the standard entry class code is CCD for batchCCD
func testBatchCCDSEC(t testing.TB) {
	mockBatch := mockBatchCCD()
	mockBatch.header.StandardEntryClassCode = "RCK"
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

// Test verify that the standard entry class code is CCD for batchCCD
func TestBatchCCDSEC(t *testing.T) {
	testBatchCCDSEC(t)
}

// Benchmark verify that the standard entry class code is CCD for batchCCD
func BenchmarkBatchCCDSEC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCCDSEC(b)
	}
}

// Verify batch CCD addenda count
func testBatchCCDAddendaCount(t testing.TB) {
	mockBatch := mockBatchCCD()
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda05())
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

// Test verify batch CCD addenda count
func TestBatchCCDAddendaCount(t *testing.T) {
	testBatchCCDAddendaCount(t)
}

// Benchmark verify batch CCD addenda count
func BenchmarkBatchCCDAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCCDAddendaCount(b)
	}
}

// Batch CCD create
func testBatchCCDCreate(t testing.TB) {
	mockBatch := mockBatchCCD()
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

// Test batch CCD create
func TestBatchCCDCreate(t *testing.T) {
	testBatchCCDCreate(t)
}

// Benchmark batch CCD create
func BenchmarkBatchCCDCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCCDCreate(b)
	}
}
