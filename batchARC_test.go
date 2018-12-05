// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "testing"

// mockBatchARCHeader creates a BatchARC BatchHeader
func mockBatchARCHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = "ARC"
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "ARC"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockARCEntryDetail creates a BatchARC EntryDetail
func mockARCEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingDebit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetCheckSerialNumber("123456789")
	entry.SetReceivingCompany("ABC Company")
	entry.SetTraceNumber(mockBatchARCHeader().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

// mockBatchARC creates a BatchARC
func mockBatchARC() *BatchARC {
	mockBatch := NewBatchARC(mockBatchARCHeader())
	mockBatch.AddEntry(mockARCEntryDetail())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// mockBatchARCHeaderCredit creates a BatchARC BatchHeader
func mockBatchARCHeaderCredit() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = DebitsOnly
	bh.StandardEntryClassCode = "ARC"
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "ARC"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockARCEntryDetailCredit creates a ARC EntryDetail with a credit entry
func mockARCEntryDetailCredit() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetCheckSerialNumber("123456789")
	entry.SetReceivingCompany("ABC Company")
	entry.SetTraceNumber(mockBatchARCHeader().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

// mockBatchARCCredit creates a BatchARC with a Credit entry
func mockBatchARCCredit() *BatchARC {
	mockBatch := NewBatchARC(mockBatchARCHeaderCredit())
	mockBatch.AddEntry(mockARCEntryDetailCredit())
	return mockBatch
}

// testBatchARCHeader creates a BatchARC BatchHeader
func testBatchARCHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchARCHeader())
	err, ok := batch.(*BatchARC)
	if !ok {
		t.Errorf("Expecting BatchARC got %T", err)
	}
}

// TestBatchARCHeader tests validating BatchARC BatchHeader
func TestBatchARCHeader(t *testing.T) {
	testBatchARCHeader(t)
}

// BenchmarkBatchARCHeader benchmarks validating BatchARC BatchHeader
func BenchmarkBatchARCHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCHeader(b)
	}
}

// testBatchARCCreate validates BatchARC create
func testBatchARCCreate(t testing.TB) {
	mockBatch := mockBatchARC()
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchARCCreate tests validating BatchARC create
func TestBatchARCCreate(t *testing.T) {
	testBatchARCCreate(t)
}

// BenchmarkBatchARCCreate benchmarks validating BatchARC create
func BenchmarkBatchARCCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCCreate(b)
	}
}

// testBatchARCStandardEntryClassCode validates BatchARC create for an invalid StandardEntryClassCode
func testBatchARCStandardEntryClassCode(t testing.TB) {
	mockBatch := mockBatchARC()
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

// TestBatchARCStandardEntryClassCode tests validating BatchARC create for an invalid StandardEntryClassCode
func TestBatchARCStandardEntryClassCode(t *testing.T) {
	testBatchARCStandardEntryClassCode(t)
}

// BenchmarkBatchARCStandardEntryClassCode benchmarks validating BatchARC create for an invalid StandardEntryClassCode
func BenchmarkBatchARCStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCStandardEntryClassCode(b)
	}
}

// testBatchARCServiceClassCodeEquality validates service class code equality
func testBatchARCServiceClassCodeEquality(t testing.TB) {
	mockBatch := mockBatchARC()
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

// TestBatchARCServiceClassCodeEquality tests validating service class code equality
func TestBatchARCServiceClassCodeEquality(t *testing.T) {
	testBatchARCServiceClassCodeEquality(t)
}

// BenchmarkBatchARCServiceClassCodeEquality benchmarks validating service class code equality
func BenchmarkBatchARCServiceClassCodeEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCServiceClassCodeEquality(b)
	}
}

// testBatchARCServiceClass200 validates BatchARC create for an invalid ServiceClassCode 200
func testBatchARCServiceClass200(t testing.TB) {
	mockBatch := mockBatchARC()
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

// TestBatchARCServiceClass200 tests validating BatchARC create for an invalid ServiceClassCode 200
func TestBatchARCServiceClass200(t *testing.T) {
	testBatchARCServiceClass200(t)
}

// BenchmarkBatchARCServiceClass200 benchmarks validating BatchARC create for an invalid ServiceClassCode 200
func BenchmarkBatchARCServiceClass200(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCServiceClass200(b)
	}
}

// testBatchARCServiceClass220 validates BatchARC create for an invalid ServiceClassCode 220
func testBatchARCServiceClass220(t testing.TB) {
	mockBatch := mockBatchARC()
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

// TestBatchARCServiceClass220 tests validating BatchARC create for an invalid ServiceClassCode 220
func TestBatchARCServiceClass220(t *testing.T) {
	testBatchARCServiceClass220(t)
}

// BenchmarkBatchARCServiceClass220 benchmarks validating BatchARC create for an invalid ServiceClassCode 220
func BenchmarkBatchARCServiceClass220(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCServiceClass220(b)
	}
}

// testBatchARCServiceClass280 validates BatchARC create for an invalid ServiceClassCode 280
func testBatchARCServiceClass280(t testing.TB) {
	mockBatch := mockBatchARC()
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

// TestBatchARCServiceClass280 tests validating BatchARC create for an invalid ServiceClassCode 280
func TestBatchARCServiceClass280(t *testing.T) {
	testBatchARCServiceClass280(t)
}

// BenchmarkBatchARCServiceClass280 benchmarks validating BatchARC create for an invalid ServiceClassCode 280
func BenchmarkBatchARCServiceClass280(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCServiceClass280(b)
	}
}

// testBatchARCAmount validates BatchARC create for an invalid Amount
func testBatchARCAmount(t testing.TB) {
	mockBatch := mockBatchARC()
	mockBatch.Entries[0].Amount = 2600000
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

// TestBatchARCAmount validates BatchARC create for an invalid Amount
func TestBatchARCAmount(t *testing.T) {
	testBatchARCAmount(t)
}

// BenchmarkBatchARCAmount validates BatchARC create for an invalid Amount
func BenchmarkBatchARCAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCAmount(b)
	}
}

// testBatchARCCheckSerialNumber validates BatchARC CheckSerialNumber / IdentificationNumber is a mandatory field
func testBatchARCCheckSerialNumber(t testing.TB) {
	mockBatch := mockBatchARC()
	// modify CheckSerialNumber / IdentificationNumber to nothing
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

// TestBatchARCCheckSerialNumber  tests validating BatchARC
// CheckSerialNumber / IdentificationNumber is a mandatory field
func TestBatchARCCheckSerialNumber(t *testing.T) {
	testBatchARCCheckSerialNumber(t)
}

// BenchmarkBatchARCCheckSerialNumber benchmarks validating BatchARC
// CheckSerialNumber / IdentificationNumber is a mandatory field
func BenchmarkBatchARCCheckSerialNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCCheckSerialNumber(b)
	}
}

// testBatchARCTransactionCode validates BatchARC TransactionCode is not a credit
func testBatchARCTransactionCode(t testing.TB) {
	mockBatch := mockBatchARCCredit()
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

// TestBatchARCTransactionCode tests validating BatchARC TransactionCode is not a credit
func TestBatchARCTransactionCode(t *testing.T) {
	testBatchARCTransactionCode(t)
}

// BenchmarkBatchARCTransactionCode benchmarks validating BatchARC TransactionCode is not a credit
func BenchmarkBatchARCTransactionCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCTransactionCode(b)
	}
}

// testBatchARCAddendaCount validates BatchARC Addenda count
func testBatchARCAddendaCount(t testing.TB) {
	mockBatch := mockBatchARC()
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

// TestBatchARCAddendaCount tests validating BatchARC Addenda count
func TestBatchARCAddendaCount(t *testing.T) {
	testBatchARCAddendaCount(t)
}

// BenchmarkBatchARCAddendaCount benchmarks validating BatchARC Addenda count
func BenchmarkBatchARCAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCAddendaCount(b)
	}
}

// testBatchARCInvalidBuild validates an invalid batch build
func testBatchARCInvalidBuild(t testing.TB) {
	mockBatch := mockBatchARC()
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

// TestBatchARCInvalidBuild tests validating an invalid batch build
func TestBatchARCInvalidBuild(t *testing.T) {
	testBatchARCInvalidBuild(t)
}

// BenchmarkBatchARCInvalidBuild benchmarks validating an invalid batch build
func BenchmarkBatchARCInvalidBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCInvalidBuild(b)
	}
}

// TestBatchARCAddendum98 validates Addenda98 returns an error
func TestBatchARCAddendum98(t *testing.T) {
	mockBatch := NewBatchARC(mockBatchARCHeader())
	mockBatch.AddEntry(mockARCEntryDetail())
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

// TestBatchARCAddendum99 validates Addenda99 returns an error
func TestBatchARCAddendum99(t *testing.T) {
	mockBatch := NewBatchARC(mockBatchARCHeader())
	mockBatch.AddEntry(mockARCEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockAddenda99.TypeCode = "05"
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

// TestBatchARCAddendum99Category validates Addenda99 returns an error
func TestBatchARCAddendum99Category(t *testing.T) {
	mockBatch := NewBatchARC(mockBatchARCHeader())
	mockBatch.AddEntry(mockARCEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
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
