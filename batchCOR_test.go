package ach

import (
	"testing"
)

// TODO make all the mock values cor fields

func mockBatchCORHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 220
	bh.StandardEntryClassCode = "COR"
	bh.CompanyName = "Your Company, inc"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "Vndr Pay"
	bh.ODFIIdentification = "121042882"
	return bh
}

func mockCOREntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 27
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 0
	entry.IdentificationNumber = "location #23"
	entry.SetReceivingCompany("Best Co. #23")
	entry.SetTraceNumber(mockBatchCORHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "S"
	return entry
}

func mockBatchCOR() *BatchCOR {
	mockBatch := NewBatchCOR(mockBatchCORHeader())
	mockBatch.AddEntry(mockCOREntryDetail())
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda98())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// testBatchCORHeader creates a COR batch header
func testBatchCORHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchCORHeader())

	_, ok := batch.(*BatchCOR)
	if !ok {
		t.Error("Expecting BachCOR")
	}
}

// TestBatchCORHeader tests creating a COR batch header
func TestBatchCORHeader(t *testing.T) {
	testBatchCORHeader(t)
}

// BenchmarkBatchCORHeader benchmarks creating a COR batch header
func BenchmarkBatchCORHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCORHeader(b)
	}
}

// testBatchCORSEC verifies COR SEC code
func testBatchCORSEC(t testing.TB) {
	mockBatch := mockBatchCOR()
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

// TestBatchCORSEC tests verifying COR SEC code
func TestBatchCORSEC(t *testing.T) {
	testBatchCORSEC(t)
}

// BenchmarkBatchCORSEC benchmarks verifying COR SEC code
func BenchmarkBatchCORSEC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCORSEC(b)
	}
}

//  testBatchCORAddendumCountTwo verifies addendum count of 2
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

// TestBatchCORAddendumCountTwo tests verifying addendum count of 2
func TestBatchCORAddendumCountTwo(t *testing.T) {
	testBatchCORAddendumCountTwo(t)
}

// BenchmarkBatchCORAddendumCountTwo benchmarks verifying addendum count of 2
func BenchmarkBatchCORAddendumCountTwo(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCORAddendumCountTwo(b)
	}
}

// testBatchCORAddendaCountZero verifies addendum count of 0
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

// TestBatchCORAddendaCountZero tests verifying addendum count of 0
func TestBatchCORAddendaCountZero(t *testing.T) {
	testBatchCORAddendaCountZero(t)
}

// BenchmarkBatchCORAddendaCountZero benchmarks verifying addendum count of 0
func BenchmarkBatchCORAddendaCountZero(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCORAddendaCountZero(b)
	}
}

// testBatchCORAddendaType verifies that Addendum is of type Addenda98
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

// TestBatchCORAddendaType tests verifying that Addendum is of type Addenda98
func TestBatchCORAddendaType(t *testing.T) {
	testBatchCORAddendaType(t)
}

// BenchmarkBatchCORAddendaType benchmarks verifying that Addendum is of type Addenda98
func BenchmarkBatchCORAddendaType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCORAddendaType(b)
	}
}

// testBatchCORAddendaTypeCode verifies Type Code
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

// TestBatchCORAddendaTypeCode tests verifying Type Code
func TestBatchCORAddendaTypeCode(t *testing.T) {
	testBatchCORAddendaTypeCode(t)
}

// BenchmarkBatchCORAddendaTypeCode benchmarks verifying Type Code
func BenchmarkBatchCORAddendaTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCORAddendaTypeCode(b)
	}
}

// testBatchCORAmount verifies batch COR amount
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

// TestBatchCORAmount tests verifying batch COR amount
func TestBatchCORAmount(t *testing.T) {
	testBatchCORAmount(t)
}

// BenchmarkBatchCORAmount benchmarks verifying batch COR amount
func BenchmarkBatchCORAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCORAmount(b)
	}
}

// testBatchCORCreate verifies creates batch COR
func testBatchCORCreate(t testing.TB) {
	mockBatch := mockBatchCOR()
	// Must have valid batch header to create a batch
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

// TestBatchCORCreate tests creating batch COR
func TestBatchCORCreate(t *testing.T) {
	testBatchCORCreate(t)
}

// BenchmarkBatchCORCreate benchmarks creating batch COR
func BenchmarkBatchCORCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCORCreate(b)
	}
}
