// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "testing"

// mockBatchPOPHeader creates a BatchPOP BatchHeader
func mockBatchPOPHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 225
	bh.StandardEntryClassCode = "POP"
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "Point of Purchase"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockPOPEntryDetail creates a BatchPOP EntryDetail
func mockPOPEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 27
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetPOPCheckSerialNumber("123456789")
	entry.SetPOPTerminalCity("PHIL")
	entry.SetPOPTerminalState("PA")
	entry.SetReceivingCompany("ABC Company")
	entry.SetTraceNumber(mockBatchPOPHeader().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

// mockBatchPOP creates a BatchPOP
func mockBatchPOP() *BatchPOP {
	mockBatch := NewBatchPOP(mockBatchPOPHeader())
	mockBatch.AddEntry(mockPOPEntryDetail())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// mockBatchPOPHeaderCredit creates a BatchPOP BatchHeader
func mockBatchPOPHeaderCredit() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 225
	bh.StandardEntryClassCode = "POP"
	bh.CompanyName = "Payee Name"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "POP"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockPOPEntryDetailCredit creates a POP EntryDetail with a credit entry
func mockPOPEntryDetailCredit() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 22
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.SetPOPCheckSerialNumber("123456789")
	entry.SetPOPTerminalCity("PHIL")
	entry.SetPOPTerminalState("PA")
	entry.SetReceivingCompany("ABC Company")
	entry.SetTraceNumber(mockBatchPOPHeader().ODFIIdentification, 123)
	entry.Category = CategoryForward
	return entry
}

// mockBatchPOPCredit creates a BatchPOP with a Credit entry
func mockBatchPOPCredit() *BatchPOP {
	mockBatch := NewBatchPOP(mockBatchPOPHeaderCredit())
	mockBatch.AddEntry(mockPOPEntryDetailCredit())
	return mockBatch
}

// testBatchPOPHeader creates a BatchPOP BatchHeader
func testBatchPOPHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchPOPHeader())
	err, ok := batch.(*BatchPOP)
	if !ok {
		t.Errorf("Expecting BatchPOP got %T", err)
	}
}

// TestBatchPOPHeader tests validating BatchPOP BatchHeader
func TestBatchPOPHeader(t *testing.T) {
	testBatchPOPHeader(t)
}

// BenchmarkBatchPOPHeader benchmarks validating BatchPOP BatchHeader
func BenchmarkBatchPOPHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPHeader(b)
	}
}

// testBatchPOPCreate validates BatchPOP create
func testBatchPOPCreate(t testing.TB) {
	mockBatch := mockBatchPOP()
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchPOPCreate tests validating BatchPOP create
func TestBatchPOPCreate(t *testing.T) {
	testBatchPOPCreate(t)
}

// BenchmarkBatchPOPCreate benchmarks validating BatchPOP create
func BenchmarkBatchPOPCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPCreate(b)
	}
}

// testBatchPOPStandardEntryClassCode validates BatchPOP create for an invalid StandardEntryClassCode
func testBatchPOPStandardEntryClassCode(t testing.TB) {
	mockBatch := mockBatchPOP()
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

// TestBatchPOPStandardEntryClassCode tests validating BatchPOP create for an invalid StandardEntryClassCode
func TestBatchPOPStandardEntryClassCode(t *testing.T) {
	testBatchPOPStandardEntryClassCode(t)
}

// BenchmarkBatchPOPStandardEntryClassCode benchmarks validating BatchPOP create for an invalid StandardEntryClassCode
func BenchmarkBatchPOPStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPStandardEntryClassCode(b)
	}
}

// testBatchPOPServiceClassCodeEquality validates service class code equality
func testBatchPOPServiceClassCodeEquality(t testing.TB) {
	mockBatch := mockBatchPOP()
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

// TestBatchPOPServiceClassCodeEquality tests validating service class code equality
func TestBatchPOPServiceClassCodeEquality(t *testing.T) {
	testBatchPOPServiceClassCodeEquality(t)
}

// BenchmarkBatchPOPServiceClassCodeEquality benchmarks validating service class code equality
func BenchmarkBatchPOPServiceClassCodeEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPServiceClassCodeEquality(b)
	}
}

// testBatchPOPServiceClass200 validates BatchPOP create for an invalid ServiceClassCode 200
func testBatchPOPServiceClass200(t testing.TB) {
	mockBatch := mockBatchPOP()
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

// TestBatchPOPServiceClass200 tests validating BatchPOP create for an invalid ServiceClassCode 200
func TestBatchPOPServiceClass200(t *testing.T) {
	testBatchPOPServiceClass200(t)
}

// BenchmarkBatchPOPServiceClass200 benchmarks validating BatchPOP create for an invalid ServiceClassCode 200
func BenchmarkBatchPOPServiceClass200(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPServiceClass200(b)
	}
}

// testBatchPOPServiceClass220 validates BatchPOP create for an invalid ServiceClassCode 220
func testBatchPOPServiceClass220(t testing.TB) {
	mockBatch := mockBatchPOP()
	mockBatch.Header.ServiceClassCode = 220
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

// TestBatchPOPServiceClass220 tests validating BatchPOP create for an invalid ServiceClassCode 220
func TestBatchPOPServiceClass220(t *testing.T) {
	testBatchPOPServiceClass220(t)
}

// BenchmarkBatchPOPServiceClass220 benchmarks validating BatchPOP create for an invalid ServiceClassCode 220
func BenchmarkBatchPOPServiceClass220(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPServiceClass220(b)
	}
}

// testBatchPOPServiceClass280 validates BatchPOP create for an invalid ServiceClassCode 280
func testBatchPOPServiceClass280(t testing.TB) {
	mockBatch := mockBatchPOP()
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

// TestBatchPOPServiceClass280 tests validating BatchPOP create for an invalid ServiceClassCode 280
func TestBatchPOPServiceClass280(t *testing.T) {
	testBatchPOPServiceClass280(t)
}

// BenchmarkBatchPOPServiceClass280 benchmarks validating BatchPOP create for an invalid ServiceClassCode 280
func BenchmarkBatchPOPServiceClass280(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPServiceClass280(b)
	}
}

// testBatchPOPAmount validates BatchPOP create for an invalid Amount
func testBatchPOPAmount(t testing.TB) {
	mockBatch := mockBatchPOP()
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

// TestBatchPOPAmount validates BatchPOP create for an invalid Amount
func TestBatchPOPAmount(t *testing.T) {
	testBatchPOPAmount(t)
}

// BenchmarkBatchPOPAmount validates BatchPOP create for an invalid Amount
func BenchmarkBatchPOPAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPAmount(b)
	}
}

// testBatchPOPCheckSerialNumber validates BatchPOP CheckSerialNumber / IdentificationNumber is a mandatory field
func testBatchPOPCheckSerialNumber(t testing.TB) {
	mockBatch := mockBatchPOP()
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

// TestBatchPOPCheckSerialNumber tests validating BatchPOP
// CheckSerialNumber / IdentificationNumber is a mandatory field
func TestBatchPOPCheckSerialNumber(t *testing.T) {
	testBatchPOPCheckSerialNumber(t)
}

// BenchmarkBatchPOPCheckSerialNumber benchmarks validating BatchPOP
// CheckSerialNumber / IdentificationNumber is a mandatory field
func BenchmarkBatchPOPCheckSerialNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPCheckSerialNumber(b)
	}
}

// testBatchPOPCheckSerialNumberField validates POPCheckSerialNumberField characters 1-9 of underlying BatchPOP
// CheckSerialNumber / IdentificationNumber
func testBatchPOPCheckSerialNumberField(t testing.TB) {
	mockBatch := mockBatchPOP()
	tc := mockBatch.Entries[0].POPCheckSerialNumberField()
	if tc != "123456789" {
		t.Error("CheckSerialNumber is invalid")
	}
}

// TestBatchPPOPCheckSerialNumberField tests validating POPCheckSerialNumberField characters 1-9 of underlying BatchPOP
// CheckSerialNumber / IdentificationNumber
func TestBatchPOPCheckSerialNumberField(t *testing.T) {
	testBatchPOPCheckSerialNumberField(t)
}

// BenchmarkBatchPOPCheckSerialNumberField benchmarks validating POPCheckSerialNumberField characters 1-9 of underlying
// BatchPOP CheckSerialNumber / IdentificationNumber
func BenchmarkBatchPOPCheckSerialNumberField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPTerminalCityField(b)
	}
}

// testBatchPOPTerminalCityField validates POPTerminalCity characters 10-13 of underlying BatchPOP
// CheckSerialNumber / IdentificationNumber
func testBatchPOPTerminalCityField(t testing.TB) {
	mockBatch := mockBatchPOP()
	tc := mockBatch.Entries[0].POPTerminalCityField()
	if tc != "PHIL" {
		t.Error("TerminalCity is invalid")
	}
}

// TestBatchPOPTerminalCityField tests validating POPTerminalCity characters 10-13 of underlying BatchPOP
// CheckSerialNumber / IdentificationNumber
func TestBatchPOPTerminalCityField(t *testing.T) {
	testBatchPOPTerminalCityField(t)
}

// BenchmarkBatchPOPTerminalCityField benchmarks validating POPTerminalCity characters 10-13 of underlying
// BatchPOP CheckSerialNumber / IdentificationNumber
func BenchmarkBatchPOPTerminalCityField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPTerminalCityField(b)
	}
}

// testBatchPOPTerminalStateField validates POPTerminalState characters 14-15 of underlying BatchPOP
// CheckSerialNumber / IdentificationNumber
func testBatchPOPTerminalStateField(t testing.TB) {
	mockBatch := mockBatchPOP()
	ts := mockBatch.Entries[0].POPTerminalStateField()
	if ts != "PA" {
		t.Error("TerminalState is invalid")
	}
}

// TestBatchPOPTerminalStateField tests validating POPTerminalState characters 14-15 of underlying BatchPOP
// CheckSerialNumber / IdentificationNumber
func TestBatchPOPTerminalStateField(t *testing.T) {
	testBatchPOPTerminalStateField(t)
}

// BenchmarkBatchPOPTerminalStateField benchmarks validating POPTerminalState characters 14-15 of underlying
// BatchPOP CheckSerialNumber / IdentificationNumber
func BenchmarkBatchPOPTerminalStateField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPTerminalStateField(b)
	}
}

// testBatchPOPTransactionCode validates BatchPOP TransactionCode is not a credit
func testBatchPOPTransactionCode(t testing.TB) {
	mockBatch := mockBatchPOPCredit()
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

// TestBatchPOPTransactionCode tests validating BatchPOP TransactionCode is not a credit
func TestBatchPOPTransactionCode(t *testing.T) {
	testBatchPOPTransactionCode(t)
}

// BenchmarkBatchPOPTransactionCode benchmarks validating BatchPOP TransactionCode is not a credit
func BenchmarkBatchPOPTransactionCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPTransactionCode(b)
	}
}

// testBatchPOPAddendaCount validates BatchPOP Addenda count
func testBatchPOPAddendaCount(t testing.TB) {
	mockBatch := mockBatchPOP()
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

// TestBatchPOPAddendaCount tests validating BatchPOP Addenda count
func TestBatchPOPAddendaCount(t *testing.T) {
	testBatchPOPAddendaCount(t)
}

// BenchmarkBatchPOPAddendaCount benchmarks validating BatchPOP Addenda count
func BenchmarkBatchPOPAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchARCAddendaCount(b)
	}
}

// testBatchPOPInvalidBuild validates an invalid batch build
func testBatchPOPInvalidBuild(t testing.TB) {
	mockBatch := mockBatchPOP()
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

// TestBatchPOPInvalidBuild tests validating an invalid batch build
func TestBatchPOPInvalidBuild(t *testing.T) {
	testBatchPOPInvalidBuild(t)
}

// BenchmarkBatchPOPInvalidBuild benchmarks validating an invalid batch build
func BenchmarkBatchPOPInvalidBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchPOPInvalidBuild(b)
	}
}
