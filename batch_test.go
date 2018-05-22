// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"
	"time"
)

// batch should never be used directly.
func mockBatch() *batch {
	mockBatch := &batch{}
	mockBatch.SetHeader(mockBatchHeader())
	mockBatch.AddEntry(mockEntryDetail())
	if err := mockBatch.build(); err != nil {
		panic(err)
	}
	return mockBatch
}

// Batch with mismatched Trace Number ODFI
func mockBatchInvalidTraceNumberODFI() *batch {
	mockBatch := &batch{}
	mockBatch.SetHeader(mockBatchHeader())
	mockBatch.AddEntry(mockEntryDetailInvalidTraceNumberODFI())
	return mockBatch
}

// Entry Detail with mismatched Trace Number ODFI
func mockEntryDetailInvalidTraceNumberODFI() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 22
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
func mockBatchNoEntry() *batch {
	mockBatch := &batch{}
	mockBatch.SetHeader(mockBatchHeader())
	return mockBatch
}

// Invalid SEC CODE Batch Header
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

func TestBatchNumberMismatch(t *testing.T) {
	testBatchNumberMismatch(t)
}
func BenchmarkBatchNumberMismatch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchNumberMismatch(b)
	}
}

func testCreditBatchisBatchAmount(t testing.TB) {
	mockBatch := mockBatch()
	mockBatch.SetHeader(mockBatchHeader())
	e1 := mockBatch.GetEntries()[0]
	e1.TransactionCode = 22
	e1.Amount = 100
	e2 := mockEntryDetail()
	e2.TransactionCode = 22
	e2.Amount = 100
	e2.TraceNumber = e1.TraceNumber + 10
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

func TestCreditBatchisBatchAmount(t *testing.T) {
	testCreditBatchisBatchAmount(t)
}

func BenchmarkCreditBatchisBatchAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testCreditBatchisBatchAmount(b)
	}

}
func testSavingsBatchisBatchAmount(t testing.TB) {
	mockBatch := mockBatch()
	mockBatch.SetHeader(mockBatchHeader())
	e1 := mockBatch.GetEntries()[0]
	e1.TransactionCode = 32
	e1.Amount = 100
	e2 := mockEntryDetail()
	e2.TransactionCode = 37
	e2.Amount = 100
	e2.TraceNumber = e1.TraceNumber + 10

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

func TestSavingsBatchisBatchAmount(t *testing.T) {
	testSavingsBatchisBatchAmount(t)
}

func BenchmarkSavingsBatchisBatchAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testSavingsBatchisBatchAmount(b)
	}
}

func testBatchisEntryHash(t testing.TB) {
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

func TestBatchisEntryHash(t *testing.T) {
	testBatchisEntryHash(t)
}

func BenchmarkBatchisEntryHash(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchisEntryHash(b)
	}
}

func testBatchDNEMismatch(t testing.TB) {
	mockBatch := mockBatch()
	mockBatch.SetHeader(mockBatchHeader())
	ed := mockBatch.GetEntries()[0]
	ed.AddAddenda(mockAddenda05())
	ed.AddAddenda(mockAddenda05())
	mockBatch.build()

	mockBatch.GetHeader().OriginatorStatusCode = 1
	mockBatch.GetEntries()[0].TransactionCode = 23
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
	e.AddAddenda(a)
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
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda05())
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
	ed.AddAddenda(mockAddenda05())
	ed.AddAddenda(mockAddenda05())
	mockBatch.build()

	mockBatch.GetEntries()[0].Addendum[0].(*Addenda05).SequenceNumber = 2
	mockBatch.GetEntries()[0].Addendum[1].(*Addenda05).SequenceNumber = 1
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
	e3.TraceNumber = 1
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
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda05())
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	mockBatch.GetEntries()[0].Addendum[0].(*Addenda05).EntryDetailSequenceNumber = 99
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

func TestNewBatchDefault(t *testing.T) {
	testNewBatchDefault(t)
}

func BenchmarkNewBatchDefault(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testNewBatchDefault(b)
	}
}

func testBatchCategory(t testing.TB) {
	mockBatch := mockBatch()
	// Add a Addenda Return to the mock batch
	entry := mockEntryDetail()
	entry.AddAddenda(mockAddenda99())
	mockBatch.AddEntry(entry)

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	if mockBatch.Category() != CategoryReturn {
		t.Errorf("Addenda99 added to batch and category is %s", mockBatch.Category())
	}
}

func TestBatchCategory(t *testing.T) {
	testBatchCategory(t)
}

func BenchmarkBatchCategory(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCategory(b)
	}
}

func testBatchCategoryForwardReturn(t testing.TB) {
	mockBatch := mockBatch()
	// Add a Addenda Return to the mock batch
	entry := mockEntryDetail()
	entry.AddAddenda(mockAddenda99())
	entry.TraceNumber = entry.TraceNumber + 10
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

func TestBatchCategoryForwardReturn(t *testing.T) {
	testBatchCategoryForwardReturn(t)
}
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
	mockBatch.header.ODFIIdentification = ""
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

func testBatchInvalidTraceNumberODFI(t testing.TB) {
	mockBatch := mockBatchInvalidTraceNumberODFI()
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

func TestBatchInvalidTraceNumberODFI(t *testing.T) {
	testBatchInvalidTraceNumberODFI(t)
}

func BenchmarkBatchInvalidTraceNumberODFI(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchInvalidTraceNumberODFI(b)
	}
}

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
}

func TestBatchNoEntry(t *testing.T) {
	testBatchNoEntry(t)
}

func BenchmarkBatchNoEntry(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchNoEntry(b)
	}
}

func testBatchControl(t testing.TB) {
	mockBatch := mockBatch()
	mockBatch.control.ODFIIdentification = ""
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

func TestBatchControl(t *testing.T) {
	testBatchControl(t)
}

func BenchmarkBatchControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchControl(b)
	}
}
