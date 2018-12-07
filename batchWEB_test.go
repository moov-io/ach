package ach

import (
	"testing"
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

// TestBatchWEBAddendum99 validates Addenda99 returns an error
func TestBatchWEBAddendum99(t *testing.T) {
	mockBatch := NewBatchWEB(mockBatchWEBHeader())
	mockBatch.AddEntry(mockWEBEntryDetail())
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

// testBatchWEBAddendaTypeCode validates addenda type code is valid
func testBatchWEBAddendaTypeCode(t testing.TB) {
	mockBatch := mockBatchWEB()
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	mockBatch.GetEntries()[0].Addenda05[0].TypeCode = "02"
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
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "PaymentType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "PaymentType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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

// TestBatchWEBCategoryReturn validates CategoryReturn returns an error
func TestBatchWEBCategoryReturn(t *testing.T) {
	mockBatch := NewBatchWEB(mockBatchWEBHeader())
	mockBatch.AddEntry(mockWEBEntryDetail())
	mockAddenda05 := mockAddenda05()
	mockBatch.GetEntries()[0].Category = CategoryReturn
	mockBatch.GetEntries()[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05)
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addenda05" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addenda98" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchWEBValidTranCodeForServiceClassCode validates a transactionCode based on ServiceClassCode
func TestBatchWEBValidTranCodeForServiceClassCode(t *testing.T) {
	mockBatch := mockBatchWEB()
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
