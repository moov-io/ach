// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "testing"

// mockBatchBOCHeader creates a BatchBOC BatchHeader
func mockBatchBOCHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = "BOC"
	bh.CompanyName = "Company Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "BOC"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockBOCEntryDetail creates a BatchBOC EntryDetail
func mockBOCEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingDebit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetCheckSerialNumber("123456789")
	entry.SetReceivingCompany("ABC Company")
	entry.SetTraceNumber(mockBatchBOCHeader().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

// mockBatchBOC creates a BatchBOC
func mockBatchBOC() *BatchBOC {
	mockBatch := NewBatchBOC(mockBatchBOCHeader())
	mockBatch.AddEntry(mockBOCEntryDetail())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// mockBatchBOCHeaderCredit creates a BatchBOC BatchHeader
func mockBatchBOCHeaderCredit() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = "BOC"
	bh.CompanyName = "Company Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "REDEPCHECK"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockBOCEntryDetailCredit creates a BatchBOC EntryDetail with a credit
func mockBOCEntryDetailCredit() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetCheckSerialNumber("123456789")
	entry.SetReceivingCompany("ABC Company")
	entry.SetTraceNumber(mockBatchBOCHeader().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

// mockBatchBOCCredit creates a BatchBOC with a Credit entry
func mockBatchBOCCredit() *BatchBOC {
	mockBatch := NewBatchBOC(mockBatchBOCHeaderCredit())
	mockBatch.AddEntry(mockBOCEntryDetailCredit())
	return mockBatch
}

// testBatchBOCHeader creates a BatchBOC BatchHeader
func testBatchBOCHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchBOCHeader())
	err, ok := batch.(*BatchBOC)
	if !ok {
		t.Errorf("Expecting BatchBOC got %T", err)
	}
}

// TestBatchBOCHeader tests validating BatchBOC BatchHeader
func TestBatchBOCHeader(t *testing.T) {
	testBatchBOCHeader(t)
}

// BenchmarkBatchBOCHeader benchmarks validating BatchBOC BatchHeader
func BenchmarkBatchBOCHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCHeader(b)
	}
}

// testBatchBOCCreate validates BatchBOC create
func testBatchBOCCreate(t testing.TB) {
	mockBatch := mockBatchBOC()
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchBOCCreate tests validating BatchBOC create
func TestBatchBOCCreate(t *testing.T) {
	testBatchBOCCreate(t)
}

// BenchmarkBatchBOCCreate benchmarks validating BatchBOC create
func BenchmarkBatchBOCCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCCreate(b)
	}
}

// testBatchBOCStandardEntryClassCode validates BatchBOC create for an invalid StandardEntryClassCode
func testBatchBOCStandardEntryClassCode(t testing.TB) {
	mockBatch := mockBatchBOC()
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

// TestBatchBOCStandardEntryClassCode tests validating BatchBOC create for an invalid StandardEntryClassCode
func TestBatchBOCStandardEntryClassCode(t *testing.T) {
	testBatchBOCStandardEntryClassCode(t)
}

// BenchmarkBatchBOCStandardEntryClassCode benchmarks validating BatchBOC create for an invalid StandardEntryClassCode
func BenchmarkBatchBOCStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCStandardEntryClassCode(b)
	}
}

// testBatchBOCServiceClassCodeEquality validates service class code equality
func testBatchBOCServiceClassCodeEquality(t testing.TB) {
	mockBatch := mockBatchBOC()
	mockBatch.GetControl().ServiceClassCode = 200
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

// TestBatchBOCServiceClassCodeEquality tests validating service class code equality
func TestBatchBOCServiceClassCodeEquality(t *testing.T) {
	testBatchBOCServiceClassCodeEquality(t)
}

// BenchmarkBatchBOCServiceClassCodeEquality benchmarks validating service class code equality
func BenchmarkBatchBOCServiceClassCodeEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCServiceClassCodeEquality(b)
	}
}

// testBatchBOCServiceClass200 validates BatchBOC create for an invalid ServiceClassCode 200
func testBatchBOCServiceClass200(t testing.TB) {
	mockBatch := mockBatchBOC()
	mockBatch.Header.ServiceClassCode = 200
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

// TestBatchBOCServiceClass200 tests validating BatchBOC create for an invalid ServiceClassCode 200
func TestBatchBOCServiceClass200(t *testing.T) {
	testBatchBOCServiceClass200(t)
}

// BenchmarkBatchBOCServiceClass200 benchmarks validating BatchBOC create for an invalid ServiceClassCode 200
func BenchmarkBatchBOCServiceClass200(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCServiceClass200(b)
	}
}

// testBatchBOCServiceClass220 validates BatchBOC create for an invalid ServiceClassCode 220
func testBatchBOCServiceClass220(t testing.TB) {
	mockBatch := mockBatchBOC()
	mockBatch.Header.ServiceClassCode = CreditsOnly
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

// TestBatchBOCServiceClass220 tests validating BatchBOC create for an invalid ServiceClassCode 220
func TestBatchBOCServiceClass220(t *testing.T) {
	testBatchBOCServiceClass220(t)
}

// BenchmarkBatchBOCServiceClass220 benchmarks validating BatchBOC create for an invalid ServiceClassCode 220
func BenchmarkBatchBOCServiceClass220(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCServiceClass220(b)
	}
}

// testBatchBOCServiceClass280 validates BatchBOC create for an invalid ServiceClassCode 280
func testBatchBOCServiceClass280(t testing.TB) {
	mockBatch := mockBatchBOC()
	mockBatch.Header.ServiceClassCode = 280
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

// TestBatchBOCServiceClass280 tests validating BatchBOC create for an invalid ServiceClassCode 280
func TestBatchBOCServiceClass280(t *testing.T) {
	testBatchBOCServiceClass280(t)
}

// BenchmarkBatchBOCServiceClass280 benchmarks validating BatchBOC create for an invalid ServiceClassCode 280
func BenchmarkBatchBOCServiceClass280(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCServiceClass280(b)
	}
}

// testBatchBOCAmount validates BatchBOC create for an invalid Amount
func testBatchBOCAmount(t testing.TB) {
	mockBatch := mockBatchBOC()
	mockBatch.Entries[0].Amount = 2500001
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

// TestBatchBOCAmount validates BatchBOC create for an invalid Amount
func TestBatchBOCAmount(t *testing.T) {
	testBatchBOCAmount(t)
}

// BenchmarkBatchBOCAmount validates BatchBOC create for an invalid Amount
func BenchmarkBatchBOCAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCAmount(b)
	}
}

// testBatchBOCCheckSerialNumber validates BatchBOC CheckSerialNumber / IdentificationNumber is a mandatory field
func testBatchBOCCheckSerialNumber(t testing.TB) {
	mockBatch := mockBatchBOC()
	// modify CheckSerialNumber / IdentificationNumber to empty string
	mockBatch.GetEntries()[0].SetCheckSerialNumber("")
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "CheckSerialNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchBOCCheckSerialNumber  tests validating BatchBOC CheckSerialNumber / IdentificationNumber is a mandatory field
func TestBatchBOCCheckSerialNumber(t *testing.T) {
	testBatchBOCCheckSerialNumber(t)
}

// BenchmarkBatchBOCCheckSerialNumber benchmarks validating BatchBOC
// CheckSerialNumber / IdentificationNumber is a mandatory field
func BenchmarkBatchBOCCheckSerialNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCCheckSerialNumber(b)
	}
}

// testBatchBOCTransactionCode validates BatchBOC TransactionCode is not a credit
func testBatchBOCTransactionCode(t testing.TB) {
	mockBatch := mockBatchBOCCredit()
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

// TestBatchBOCTransactionCode tests validating BatchBOC TransactionCode is not a credit
func TestBatchBOCTransactionCode(t *testing.T) {
	testBatchBOCTransactionCode(t)
}

// BenchmarkBatchBOCTransactionCode benchmarks validating BatchBOC TransactionCode is not a credit
func BenchmarkBatchBOCTransactionCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCTransactionCode(b)
	}
}

// testBatchBOCAddenda05 validates BatchBOC Addenda count
func testBatchBOCAddenda05(t testing.TB) {
	mockBatch := mockBatchBOC()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addenda05" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchBOCAddenda05 tests validating BatchBOC Addenda count
func TestBatchBOCAddenda05(t *testing.T) {
	testBatchBOCAddenda05(t)
}

// BenchmarkBatchBOCAddenda05 benchmarks validating BatchBOC Addenda count
func BenchmarkBatchBOCAddenda05(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCAddenda05(b)
	}
}

// testBatchBOCInvalidBuild validates an invalid batch build
func testBatchBOCInvalidBuild(t testing.TB) {
	mockBatch := mockBatchBOC()
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

// TestBatchBOCInvalidBuild tests validating an invalid batch build
func TestBatchBOCInvalidBuild(t *testing.T) {
	testBatchBOCInvalidBuild(t)
}

// BenchmarkBatchBOCInvalidBuild benchmarks validating an invalid batch build
func BenchmarkBatchBOCInvalidBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchBOCInvalidBuild(b)
	}
}

// TestBatchBOCAddendum98 validates Addenda98 returns an error
func TestBatchBOCAddendum98(t *testing.T) {
	mockBatch := NewBatchBOC(mockBatchBOCHeader())
	mockBatch.AddEntry(mockBOCEntryDetail())
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

// TestBatchBOCAddendum99 validates Addenda99 returns an error
func TestBatchBOCAddendum99(t *testing.T) {
	mockBatch := NewBatchBOC(mockBatchBOCHeader())
	mockBatch.AddEntry(mockBOCEntryDetail())
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
