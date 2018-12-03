// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"
	"time"
)

// batch should never be used directly.
func mockBatch() *Batch {
	mockBatch := &Batch{}
	mockBatch.SetHeader(mockBatchHeader())
	mockBatch.AddEntry(mockEntryDetail())
	if err := mockBatch.build(); err != nil {
		panic(err)
	}
	return mockBatch
}

// Batch with mismatched TraceNumber ODFI
func mockBatchInvalidTraceNumberODFI() *Batch {
	mockBatch := &Batch{}
	mockBatch.SetHeader(mockBatchHeader())
	mockBatch.AddEntry(mockEntryDetailInvalidTraceNumberODFI())
	return mockBatch
}

// EntryDetail with mismatched TraceNumber ODFI
func mockEntryDetailInvalidTraceNumberODFI() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingCredit
	entry.SetRDFI("121042882")
	entry.DFIAccountNumber = "123456789"
	entry.Amount = 100000000
	entry.IndividualName = "Wade Arnold"
	entry.SetTraceNumber("9928272", 1)
	entry.IdentificationNumber = "ABC##jvkdjfuiwn"
	entry.Category = CategoryForward
	return entry
}

// Batch with no entries
func mockBatchNoEntry() *Batch {
	mockBatch := &Batch{}
	mockBatch.SetHeader(mockBatchHeader())
	return mockBatch
}

// Invalid SEC CODE BatchHeader
func mockBatchInvalidSECHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 220
	bh.StandardEntryClassCode = "NIL"
	bh.CompanyName = "ACME Corporation"
	bh.CompanyIdentification = "123456789"
	bh.CompanyEntryDescription = "PAYROLL"
	bh.EffectiveEntryDate = time.Now()
	bh.ODFIIdentification = "123456789"
	return bh
}

// Test cases that apply to all batch types
// testBatchNumberMismatch validates BatchNumber mismatch
func testBatchNumberMismatch(t testing.TB) {
	mockBatch := mockBatch()
	mockBatch.GetControl().BatchNumber = 2
	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "BatchNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchNumberMismatch tests validating BatchNumber mismatch
func TestBatchNumberMismatch(t *testing.T) {
	testBatchNumberMismatch(t)
}

// BenchmarkBatchNumberMismatch benchmarks validating BatchNumber mismatch
func BenchmarkBatchNumberMismatch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchNumberMismatch(b)
	}
}

// testCreditBatchIsBatchAmount validates Batch TotalCreditEntryDollarAmount
func testCreditBatchIsBatchAmount(t testing.TB) {
	mockBatch := mockBatch()
	mockBatch.SetHeader(mockBatchHeader())
	e1 := mockBatch.GetEntries()[0]
	e1.TransactionCode = CheckingCredit
	e1.Amount = 100
	e2 := mockEntryDetail()
	e2.TransactionCode = CheckingCredit
	e2.Amount = 100
	// replace last 2 of TraceNumber
	e2.TraceNumber = e1.TraceNumber[:13] + "10"
	mockBatch.AddEntry(e2)
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	mockBatch.GetControl().TotalCreditEntryDollarAmount = 1
	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TotalCreditEntryDollarAmount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestCreditBatchIsBatchAmount test validating Batch TotalCreditEntryDollarAmount
func TestCreditBatchIsBatchAmount(t *testing.T) {
	testCreditBatchIsBatchAmount(t)
}

// BenchmarkCreditBatchIsBatchAmount benchmarks Batch TotalCreditEntryDollarAmount
func BenchmarkCreditBatchIsBatchAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCreditBatchIsBatchAmount(b)
	}

}

// testSavingsBatchIsBatchAmount validates Batch TotalDebitEntryDollarAmount
func testSavingsBatchIsBatchAmount(t testing.TB) {
	mockBatch := mockBatch()
	mockBatch.SetHeader(mockBatchHeader())
	e1 := mockBatch.GetEntries()[0]
	e1.TransactionCode = 32
	e1.Amount = 100
	e2 := mockEntryDetail()
	e2.TransactionCode = 37
	e2.Amount = 100
	// replace last 2 of TraceNumber
	e2.TraceNumber = e1.TraceNumber[:13] + "10"

	mockBatch.AddEntry(e2)
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	mockBatch.GetControl().TotalDebitEntryDollarAmount = 1
	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TotalDebitEntryDollarAmount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestSavingsBatchIsBatchAmount tests validating Batch TotalDebitEntryDollarAmount
func TestSavingsBatchIsBatchAmount(t *testing.T) {
	testSavingsBatchIsBatchAmount(t)
}

// BenchmarkSavingsBatchIsBatchAmount benchmarks validating Batch TotalDebitEntryDollarAmount
func BenchmarkSavingsBatchIsBatchAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testSavingsBatchIsBatchAmount(b)
	}
}

func testBatchIsEntryHash(t testing.TB) {
	mockBatch := mockBatch()
	mockBatch.GetControl().EntryHash = 1
	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "EntryHash" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchIsEntryHash(t *testing.T) {
	testBatchIsEntryHash(t)
}

func BenchmarkBatchIsEntryHash(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchIsEntryHash(b)
	}
}

func testBatchDNEMismatch(t testing.TB) {
	mockBatch := mockBatch()
	mockBatch.SetHeader(mockBatchHeader())
	ed := mockBatch.GetEntries()[0]
	ed.AddAddenda05(mockAddenda05())
	ed.AddAddenda05(mockAddenda05())
	mockBatch.build()

	mockBatch.GetHeader().OriginatorStatusCode = 1
	mockBatch.GetEntries()[0].TransactionCode = CheckingPrenoteCredit
	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "OriginatorStatusCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchDNEMismatch(t *testing.T) {
	testBatchDNEMismatch(t)
}

func BenchmarkBatchDNEMismatch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchDNEMismatch(b)
	}
}

func testBatchTraceNumberNotODFI(t testing.TB) {
	mockBatch := mockBatch()
	mockBatch.GetEntries()[0].SetTraceNumber("12345678", 1)
	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ODFIIdentificationField" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchTraceNumberNotODFI(t *testing.T) {
	testBatchTraceNumberNotODFI(t)
}

func BenchmarkBatchTraceNumberNotODFI(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTraceNumberNotODFI(b)
	}
}

func testBatchEntryCountEquality(t testing.TB) {
	mockBatch := mockBatch()
	mockBatch.SetHeader(mockBatchHeader())
	e := mockEntryDetail()
	a := mockAddenda05()
	e.AddAddenda05(a)
	mockBatch.AddEntry(e)
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	mockBatch.GetControl().EntryAddendaCount = 1
	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "EntryAddendaCount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchEntryCountEquality(t *testing.T) {
	testBatchEntryCountEquality(t)
}

func BenchmarkBatchEntryCountEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchEntryCountEquality(b)
	}
}

func testBatchAddendaIndicator(t testing.TB) {
	mockBatch := mockBatch()
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	mockBatch.GetEntries()[0].AddendaRecordIndicator = 0
	mockBatch.GetControl().EntryAddendaCount = 2
	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "AddendaRecordIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchAddendaIndicator(t *testing.T) {
	testBatchAddendaIndicator(t)
}

func BenchmarkBatchAddendaIndicator(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchAddendaIndicator(b)
	}
}

func testBatchIsAddendaSeqAscending(t testing.TB) {
	mockBatch := mockBatch()
	ed := mockBatch.GetEntries()[0]
	ed.AddAddenda05(mockAddenda05())
	ed.AddAddenda05(mockAddenda05())
	mockBatch.build()
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].Addenda05[0].SequenceNumber = 2
	mockBatch.GetEntries()[0].Addenda05[1].SequenceNumber = 1
	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "SequenceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchIsAddendaSeqAscending(t *testing.T) {
	testBatchIsAddendaSeqAscending(t)
}
func BenchmarkBatchIsAddendaSeqAscending(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchIsAddendaSeqAscending(b)
	}
}

func testBatchIsSequenceAscending(t testing.TB) {
	mockBatch := mockBatch()
	e3 := mockEntryDetail()
	e3.TraceNumber = "1"
	mockBatch.AddEntry(e3)
	mockBatch.GetControl().EntryAddendaCount = 2
	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TraceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchIsSequenceAscending(t *testing.T) {
	testBatchIsSequenceAscending(t)
}

func BenchmarkBatchIsSequenceAscending(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchIsSequenceAscending(b)
	}
}

func testBatchAddendaTraceNumber(t testing.TB) {
	mockBatch := mockBatch()
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	mockBatch.GetEntries()[0].Addenda05[0].EntryDetailSequenceNumber = 99
	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TraceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchAddendaTraceNumber(t *testing.T) {
	testBatchAddendaTraceNumber(t)
}

func BenchmarkBatchAddendaTraceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchAddendaTraceNumber(b)
	}
}

// testNewBatchDefault validates error for NewBatch if invalid SEC Code
func testNewBatchDefault(t testing.TB) {
	_, err := NewBatch(mockBatchInvalidSECHeader())

	if e, ok := err.(*FileError); ok {
		if e.FieldName != "StandardEntryClassCode" {
			t.Errorf("%T: %s", err, err)
		}
	} else {
		t.Errorf("%T: %s", err, err)
	}
}

// TestNewBatchDefault test validating error for NewBatch if invalid SEC Code
func TestNewBatchDefault(t *testing.T) {
	testNewBatchDefault(t)
}

// BenchmarkNewBatchDefault benchmarks validating error for NewBatch if
// invalid SEC Code
func BenchmarkNewBatchDefault(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testNewBatchDefault(b)
	}
}

// testBatchCategory validates Batch Category
func testBatchCategory(t testing.TB) {
	mockBatch := mockBatch()
	// Add a Addenda Return to the mock batch
	entry := mockEntryDetail()
	entry.Addenda99 = mockAddenda99()
	entry.Category = CategoryReturn
	mockBatch.AddEntry(entry)

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	if mockBatch.Category() != CategoryReturn {
		t.Errorf("Addenda99 added to batch and category is %s", mockBatch.Category())
	}
}

// TestBatchCategory tests validating Batch Category
func TestBatchCategory(t *testing.T) {
	testBatchCategory(t)
}

// BenchmarkBatchCategory benchmarks validating Batch Category
func BenchmarkBatchCategory(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCategory(b)
	}
}

//  testBatchCategoryForwardReturn validates Category based on EntryDetail
func testBatchCategoryForwardReturn(t testing.TB) {
	mockBatch := mockBatch()
	// Add a Addenda Return to the mock batch
	entry := mockEntryDetail()
	entry.Addenda99 = mockAddenda99()
	entry.Category = CategoryReturn
	// replace last 2 of TraceNumber
	entry.TraceNumber = entry.TraceNumber[:13] + "10"
	entry.AddendaRecordIndicator = 1
	mockBatch.AddEntry(entry)
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Category" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchCategoryForwardReturn tests validating Category based on EntryDetail
func TestBatchCategoryForwardReturn(t *testing.T) {
	testBatchCategoryForwardReturn(t)
}

//  BenchmarkBatchCategoryForwardReturn benchmarks validating Category based on EntryDetail
func BenchmarkBatchCategoryForwardReturn(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCategoryForwardReturn(b)
	}
}

// Don't over write a batch trace number when building if it already exists
func testBatchTraceNumberExists(t testing.TB) {
	mockBatch := mockBatch()
	entry := mockEntryDetail()
	traceBefore := entry.TraceNumberField()
	mockBatch.AddEntry(entry)
	mockBatch.build()
	traceAfter := mockBatch.GetEntries()[1].TraceNumberField()
	if traceBefore != traceAfter {
		t.Errorf("Trace number was set to %v before batch.build and is now %v\n", traceBefore, traceAfter)
	}
}

func TestBatchTraceNumberExists(t *testing.T) {
	testBatchTraceNumberExists(t)
}

func BenchmarkBatchTraceNumberExists(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchTraceNumberExists(b)
	}
}

func testBatchFieldInclusion(t testing.TB) {
	mockBatch := mockBatch()
	mockBatch.Header.ODFIIdentification = ""
	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ODFIIdentification" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchFieldInclusion(t *testing.T) {
	testBatchFieldInclusion(t)
}

func BenchmarkBatchFieldInclusion(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchFieldInclusion(b)
	}
}

// testBatchInvalidTraceNumberODFI validates TraceNumberODFI
func testBatchInvalidTraceNumberODFI(t testing.TB) {
	mockBatch := mockBatchInvalidTraceNumberODFI()
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchInvalidTraceNumberODFI tests validating TraceNumberODFI
func TestBatchInvalidTraceNumberODFI(t *testing.T) {
	testBatchInvalidTraceNumberODFI(t)
}

// BenchmarkBatchInvalidTraceNumberODFI benchmarks validating TraceNumberODFI
func BenchmarkBatchInvalidTraceNumberODFI(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchInvalidTraceNumberODFI(b)
	}
}

// testBatchNoEntry validates error for a batch with no entries
func testBatchNoEntry(t testing.TB) {
	mockBatch := mockBatchNoEntry()
	if err := mockBatch.build(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "entries" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}

	// test verify
	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "entries" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}

}

// TestBatchNoEntry tests validating error for a batch with no entries
func TestBatchNoEntry(t *testing.T) {
	testBatchNoEntry(t)
}

// BenchmarkBatchNoEntry benchmarks validating error for a batch with no entries
func BenchmarkBatchNoEntry(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchNoEntry(b)
	}
}

// testBatchControl validates BatchControl ODFIIdentification
func testBatchControl(t testing.TB) {
	mockBatch := mockBatch()
	mockBatch.Control.ODFIIdentification = ""
	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ODFIIdentification" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchControl tests validating BatchControl ODFIIdentification
func TestBatchControl(t *testing.T) {
	testBatchControl(t)
}

// BenchmarkBatchControl benchmarks validating BatchControl ODFIIdentification
func BenchmarkBatchControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchControl(b)
	}
}

// testIATBatch validates an IAT batch returns an error for batch
func testIATBatch(t testing.TB) {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 220
	bh.StandardEntryClassCode = "IAT"
	bh.CompanyName = "ACME Corporation"
	bh.CompanyIdentification = "123456789"
	bh.CompanyEntryDescription = "PAYROLL"
	bh.EffectiveEntryDate = time.Now()
	bh.ODFIIdentification = "123456789"

	_, err := NewBatch(bh)

	if e, ok := err.(*FileError); ok {
		if e.FieldName != "StandardEntryClassCode" {
			t.Errorf("%T: %s", err, err)
		}
	} else {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatch tests validating an IAT batch returns an error for batch
func TestIATBatch(t *testing.T) {
	testIATBatch(t)
}

// BenchmarkIATBatch benchmarks validating an IAT batch returns an error for batch
func BenchmarkIATBatch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatch(b)
	}
}

// TestBatchADVInvalidServiceClassCode validates ServiceClassCode
func TestBatchADVInvalidServiceClassCode(t *testing.T) {
	mockBatch := mockBatchADV()
	mockBatch.Create()
	mockBatch.ADVControl.ServiceClassCode = 220
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

// TestBatchADVInvalidODFIIdentification validates ODFIIdentification
func TestBatchADVInvalidODFIIdentification(t *testing.T) {
	mockBatch := mockBatchADV()
	mockBatch.Create()
	mockBatch.ADVControl.ODFIIdentification = "231380104"
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ODFIIdentification" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchADVInvalidBatchNumber validates BatchNumber
func TestBatchADVInvalidBatchNumber(t *testing.T) {
	mockBatch := mockBatchADV()
	mockBatch.Create()
	mockBatch.ADVControl.BatchNumber = 2
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "BatchNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchADVEntryAddendaCount validates EntryAddendaCount
func TestBatchADVInvalidEntryAddendaCount(t *testing.T) {
	mockBatch := mockBatchADV()
	mockBatch.Create()
	mockBatch.ADVControl.EntryAddendaCount = CheckingCredit
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

// TestBatchADVTotalDebitEntryDollarAmount validates TotalDebitEntryDollarAmount
func TestBatchADVInvalidTotalDebitEntryDollarAmount(t *testing.T) {
	mockBatch := mockBatchADV()
	mockBatch.GetADVEntries()[0].TransactionCode = DebitForCreditsOriginated
	mockBatch.Create()
	mockBatch.ADVControl.TotalDebitEntryDollarAmount = 2200
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TotalDebitEntryDollarAmount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchADVTotalCreditEntryDollarAmount validates TotalCreditEntryDollarAmount
func TestBatchADVInvalidTotalCreditEntryDollarAmount(t *testing.T) {
	mockBatch := mockBatchADV()
	mockBatch.Create()
	mockBatch.ADVControl.TotalCreditEntryDollarAmount = 2200
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TotalCreditEntryDollarAmount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchADVEntryHash validates EntryHash
func TestBatchADVInvalidEntryHash(t *testing.T) {
	mockBatch := mockBatchADV()
	mockBatch.Create()
	mockBatch.ADVControl.EntryHash = 2200233
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "EntryHash" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchAddenda98InvalidAddendaRecordIndicator validates AddendaRecordIndicator
func TestBatchAddenda98InvalidAddendaRecordIndicator(t *testing.T) {
	mockBatch := mockBatchCOR()
	mockBatch.GetEntries()[0].AddendaRecordIndicator = 0
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "AddendaRecordIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchAddenda02InvalidAddendaRecordIndicator validates AddendaRecordIndicator
func TestBatchAddenda02InvalidAddendaRecordIndicator(t *testing.T) {
	mockBatch := mockBatchPOS()
	mockBatch.GetEntries()[0].AddendaRecordIndicator = 0
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "AddendaRecordIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchADVCategory validates Category
func TestBatchADVCategory(t *testing.T) {
	mockBatch := mockBatchADV()

	entryOne := NewADVEntryDetail()
	entryOne.TransactionCode = CreditForDebitsOriginated
	entryOne.SetRDFI("231380104")
	entryOne.DFIAccountNumber = "744-5678-99"
	entryOne.Amount = 50000
	entryOne.AdviceRoutingNumber = "121042882"
	entryOne.FileIdentification = "FILE1"
	entryOne.ACHOperatorData = ""
	entryOne.IndividualName = "Name"
	entryOne.DiscretionaryData = ""
	entryOne.AddendaRecordIndicator = 0
	entryOne.ACHOperatorRoutingNumber = "01100001"
	entryOne.JulianDay = 50
	entryOne.SequenceNumber = 1
	entryOne.Category = CategoryReturn

	mockBatch.AddADVEntry(entryOne)
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Category" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestBatchDishonoredReturnsCategory validates Category for Returns
func TestBatchDishonoredReturnsCategory(t *testing.T) {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingDebit
	entry.SetRDFI("121042882")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 25000
	entry.IdentificationNumber = "45689033"
	entry.IndividualName = "Wade Arnold"
	entry.SetTraceNumber(mockBatchPOSHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "01"
	entry.AddendaRecordIndicator = 1
	entry.Category = CategoryDishonoredReturn

	addenda99 := mockAddenda99()
	addenda99.ReturnCode = "R68"
	addenda99.AddendaInformation = "Untimely Return"
	entry.Addenda99 = addenda99

	entryOne := NewEntryDetail()
	entryOne.TransactionCode = CheckingDebit
	entryOne.SetRDFI("121042882")
	entryOne.DFIAccountNumber = "744-5678-99"
	entryOne.Amount = 23000
	entryOne.IdentificationNumber = "45689033"
	entryOne.IndividualName = "Adam Decaf"
	entryOne.SetTraceNumber(mockBatchPOSHeader().ODFIIdentification, 1)
	entryOne.DiscretionaryData = "01"
	entryOne.AddendaRecordIndicator = 1
	entryOne.Category = CategoryReturn

	addenda99One := mockAddenda99()
	addenda99One.ReturnCode = "R68"
	addenda99One.AddendaInformation = "Untimely Return"
	entryOne.Addenda99 = addenda99One

	posHeader := NewBatchHeader()
	posHeader.ServiceClassCode = 225
	posHeader.StandardEntryClassCode = "POS"
	posHeader.CompanyName = "Payee Name"
	posHeader.CompanyIdentification = "231380104"
	posHeader.CompanyEntryDescription = "ACH POS"
	posHeader.ODFIIdentification = "23138010"

	batch := NewBatchPOS(posHeader)
	batch.SetHeader(posHeader)
	batch.AddEntry(entry)
	batch.AddEntry(entryOne)

	if err := batch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Category" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}
