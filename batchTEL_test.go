package ach

import (
	"testing"
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
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addenda02" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			println(e.Error())
			if e.FieldName != "PaymentType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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

// TestBatchTELAddendum99 validates Addenda99 returns an error
func TestBatchTELAddendum99(t *testing.T) {
	mockBatch := NewBatchTEL(mockBatchTELHeader())
	mockBatch.AddEntry(mockTELEntryDetail())
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

// TestBatchTELValidTranCodeForServiceClassCode validates a transactionCode based on ServiceClassCode
func TestBatchTELValidTranCodeForServiceClassCode(t *testing.T) {
	mockBatch := mockBatchTEL()
	mockBatch.GetHeader().ServiceClassCode = CreditsOnly
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
