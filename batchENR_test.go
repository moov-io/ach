// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"log"
	"testing"
)

// mockBatchENRHeader creates a ENR batch header
func mockBatchENRHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.CompanyName = "Name on Account"
	bh.CompanyIdentification = "231380104"
	bh.StandardEntryClassCode = "ENR"
	bh.CompanyEntryDescription = "AUTOENROLL"
	bh.ODFIIdentification = "23138010"
	return bh
}

// mockENREntryDetail creates a ENR entry detail
func mockENREntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("031300012")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 0
	entry.SetOriginalTraceNumber("031300010000001")
	entry.SetReceivingCompany("Best. #1")
	entry.SetTraceNumber("23138010", 1)

	addenda := NewAddenda05()
	addenda.PaymentRelatedInformation = `21*12200004*3*123987654321*777777777*DOE*JOHN*1\`
	entry.AddAddenda05(addenda)
	entry.AddendaRecordIndicator = 1

	return entry
}

// mockBatchENR creates a ENR batch
func mockBatchENR() *BatchENR {
	batch := NewBatchENR(mockBatchENRHeader())
	batch.AddEntry(mockENREntryDetail())
	if err := batch.Create(); err != nil {
		log.Fatalf("Unexpected error building batch: %s\n", err)
	}
	return batch
}

// testBatchENRHeader creates a ENR batch header
func testBatchENRHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchENRHeader())
	_, ok := batch.(*BatchENR)
	if !ok {
		t.Error("Expecting BatchENR")
	}
}

// TestBatchENRHeader tests creating a ENR batch header
func TestBatchENRHeader(t *testing.T) {
	testBatchENRHeader(t)
}

// BenchmarkBatchENRHeader benchmark creating a ENR batch header
func BenchmarkBatchENRHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchENRHeader(b)
	}
}

// testBatchENRAddendumCount batch control ENR must have 1-9999 Addenda05 records
func testBatchENRAddendumCount(t testing.TB) {
	mockBatch := mockBatchENR()
	// Adding a second addenda to the mock entry
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	if err := mockBatch.Create(); err != nil {
		t.Errorf("Adding addenda is allowed: %v", err)
	}
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("Adding addendas is allowed: %v", err)
	}
}

// TestBatchENRAddendumCount tests batch control ENR can only have one addendum per entry detail
func TestBatchENRAddendumCount(t *testing.T) {
	testBatchENRAddendumCount(t)
}

// BenchmarkBatchENRAddendumCount benchmarks batch control ENR can only have one addendum per entry detail
func BenchmarkBatchENRAddendumCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchENRAddendumCount(b)
	}
}

// TestBatchENRAddendum98 validates Addenda05 returns an error
func TestBatchENRAddendum98(t *testing.T) {
	mockBatch := NewBatchENR(mockBatchENRHeader())
	mockBatch.AddEntry(mockENREntryDetail())
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

// testBatchENRCompanyEntryDescription validates CompanyEntryDescription
func testBatchENRCompanyEntryDescription(t testing.TB) {
	mockBatch := mockBatchENR()
	mockBatch.Header.CompanyEntryDescription = "bad"
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "CompanyEntryDescription" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchENRCompanyEntryDescription tests validating receiving company / Individual name is a mandatory field
func TestBatchENRCompanyEntryDescription(t *testing.T) {
	testBatchENRCompanyEntryDescription(t)
}

// BenchmarkBatchENRCompanyEntryDescription benchmarks validating receiving company / Individual name is a mandatory field
func BenchmarkBatchENRCompanyEntryDescription(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchENRCompanyEntryDescription(b)
	}
}

// testBatchENRAddendaTypeCode validates addenda type code is 05
func testBatchENRAddendaTypeCode(t testing.TB) {
	mockBatch := mockBatchENR()
	mockBatch.GetEntries()[0].Addenda05[0].TypeCode = "98"
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

// TestBatchENRAddendaTypeCode tests validating addenda type code is 05
func TestBatchENRAddendaTypeCode(t *testing.T) {
	testBatchENRAddendaTypeCode(t)
}

// BenchmarkBatchENRAddendaTypeCod benchmarks validating addenda type code is 05
func BenchmarkBatchENRAddendaTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchENRAddendaTypeCode(b)
	}
}

// testBatchENRSEC validates that the standard entry class code is ENR for batchENR
func testBatchENRSEC(t testing.TB) {
	mockBatch := mockBatchENR()
	mockBatch.Header.StandardEntryClassCode = "ACK"
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

// TestBatchENRSEC tests validating that the standard entry class code is ENR for batchENR
func TestBatchENRSEC(t *testing.T) {
	testBatchENRSEC(t)
}

// BenchmarkBatchENRSEC benchmarks validating that the standard entry class code is ENR for batch ENR
func BenchmarkBatchENRSEC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchENRSEC(b)
	}
}

// testBatchENRAddendaCount validates batch ENR addenda count
func testBatchENRAddendaCount(t testing.TB) {
	mockBatch := mockBatchENR()
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

// TestBatchENRAddendaCount tests validating batch ENR addenda count
func TestBatchENRAddendaCount(t *testing.T) {
	testBatchENRAddendaCount(t)
}

// BenchmarkBatchENRAddendaCount benchmarks validating batch ENR addenda count
func BenchmarkBatchENRAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchENRAddendaCount(b)
	}
}

// testBatchENRServiceClassCode validates ServiceClassCode
func testBatchENRServiceClassCode(t testing.TB) {
	mockBatch := mockBatchENR()
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

// TestBatchENRServiceClassCode tests validating ServiceClassCode
func TestBatchENRServiceClassCode(t *testing.T) {
	testBatchENRServiceClassCode(t)
}

// BenchmarkBatchENRServiceClassCode benchmarks validating ServiceClassCode
func BenchmarkBatchENRServiceClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchENRServiceClassCode(b)
	}
}

// TestBatchENRAmount validates Amount
func TestBatchENRAmount(t *testing.T) {
	mockBatch := mockBatchENR()
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

// TestBatchENRTransactionCode validates TransactionCode
func TestBatchENRTransactionCode(t *testing.T) {
	mockBatch := mockBatchENR()
	mockBatch.GetEntries()[0].TransactionCode = CheckingReturnNOCCredit
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

func TestBatchENR__PaymentInformation(t *testing.T) {
	batch := mockBatchENR()
	if err := batch.Validate(); err != nil {
		t.Fatal(err)
	}
	addenda05 := batch.GetEntries()[0].Addenda05[0]
	info, err := batch.ParsePaymentInformation(addenda05)
	if err != nil {
		t.Fatal(err)
	}

	if v := info.TransactionCode; v != CheckingReturnNOCCredit {
		t.Errorf("TransactionCode: %d", v)
	}
	if v := info.RDFIIdentification; v != "12200004" {
		t.Errorf("RDFIIdentification: %s", v)
	}
	if v := info.CheckDigit; v != "3" {
		t.Errorf("CheckDigit: %s", v)
	}
	if v := info.DFIAccountNumber; v != "123987654321" {
		t.Errorf("DFIAccountNumber: %s", v)
	}
	if v := info.IndividualIdentification; v != "777777777" {
		t.Errorf("IndividualIdentification: %s", v)
	}
	if v := info.IndividualName; v != "JOHN DOE" {
		t.Errorf("IndividualName: %s", v)
	}
	if v := info.EnrolleeClassificationCode; v != 1 {
		t.Errorf("EnrolleeClassificationCode: %d", v)
	}
}

// TestBatchENRValidTranCodeForServiceClassCode validates a transactionCode based on ServiceClassCode
func TestBatchENRValidTranCodeForServiceClassCode(t *testing.T) {
	mockBatch := mockBatchENR()
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

// TestBatchENRAddenda02 validates BatchENR cannot have Addenda02
func TestBatchENRAddenda02(t *testing.T) {
	mockBatch := mockBatchENR()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].Addenda02 = mockAddenda02()
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
