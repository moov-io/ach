// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"log"
	"testing"
)

// mockIATBatch
func mockIATBatch() IATBatch {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATBatchHeaderFF())
	mockBatch.AddEntry(mockIATEntryDetail())
	mockBatch.Entries[0].Addenda10 = mockAddenda10()
	mockBatch.Entries[0].Addenda11 = mockAddenda11()
	mockBatch.Entries[0].Addenda12 = mockAddenda12()
	mockBatch.Entries[0].Addenda13 = mockAddenda13()
	mockBatch.Entries[0].Addenda14 = mockAddenda14()
	mockBatch.Entries[0].Addenda15 = mockAddenda15()
	mockBatch.Entries[0].Addenda16 = mockAddenda16()
	if err := mockBatch.build(); err != nil {
		log.Fatal(err)
	}
	return mockBatch
}

// mockIATBatchManyEntries
func mockIATBatchManyEntries() IATBatch {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATBatchHeaderFF())

	mockBatch.AddEntry(mockIATEntryDetail())

	mockBatch.Entries[0].Addenda10 = mockAddenda10()
	mockBatch.Entries[0].Addenda11 = mockAddenda11()
	mockBatch.Entries[0].Addenda12 = mockAddenda12()
	mockBatch.Entries[0].Addenda13 = mockAddenda13()
	mockBatch.Entries[0].Addenda14 = mockAddenda14()
	mockBatch.Entries[0].Addenda15 = mockAddenda15()
	mockBatch.Entries[0].Addenda16 = mockAddenda16()
	mockBatch.Entries[0].AddIATAddenda(mockAddenda17())
	mockBatch.Entries[0].AddIATAddenda(mockAddenda17B())
	mockBatch.Entries[0].AddIATAddenda(mockAddenda18())
	mockBatch.Entries[0].AddIATAddenda(mockAddenda18B())
	mockBatch.Entries[0].AddIATAddenda(mockAddenda18C())
	mockBatch.Entries[0].AddIATAddenda(mockAddenda18D())
	mockBatch.Entries[0].AddIATAddenda(mockAddenda18E())

	mockBatch.AddEntry(mockIATEntryDetail2())

	mockBatch.Entries[1].Addenda10 = mockAddenda10()
	mockBatch.Entries[1].Addenda11 = mockAddenda11()
	mockBatch.Entries[1].Addenda12 = mockAddenda12()
	mockBatch.Entries[1].Addenda13 = mockAddenda13()
	mockBatch.Entries[1].Addenda14 = mockAddenda14()
	mockBatch.Entries[1].Addenda15 = mockAddenda15()
	mockBatch.Entries[1].Addenda16 = mockAddenda16()
	mockBatch.Entries[1].AddIATAddenda(mockAddenda17())
	mockBatch.Entries[1].AddIATAddenda(mockAddenda17B())
	mockBatch.Entries[1].AddIATAddenda(mockAddenda18())
	mockBatch.Entries[1].AddIATAddenda(mockAddenda18B())
	mockBatch.Entries[1].AddIATAddenda(mockAddenda18C())
	mockBatch.Entries[1].AddIATAddenda(mockAddenda18D())
	mockBatch.Entries[1].AddIATAddenda(mockAddenda18E())

	if err := mockBatch.build(); err != nil {
		log.Fatal(err)
	}
	return mockBatch
}

// mockIATBatch
func mockInvalidIATBatch() IATBatch {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATBatchHeaderFF())
	mockBatch.AddEntry(mockIATEntryDetail())
	mockBatch.Entries[0].Addenda10 = mockAddenda10()
	mockBatch.Entries[0].Addenda11 = mockAddenda11()
	mockBatch.Entries[0].Addenda12 = mockAddenda12()
	mockBatch.Entries[0].Addenda13 = mockAddenda13()
	mockBatch.Entries[0].Addenda14 = mockAddenda14()
	mockBatch.Entries[0].Addenda15 = mockAddenda15()
	mockBatch.Entries[0].Addenda16 = mockAddenda16()
	mockBatch.Entries[0].AddIATAddenda(mockInvalidAddenda17())
	if err := mockBatch.build(); err != nil {
		log.Fatal(err)
	}
	return mockBatch
}

func mockInvalidAddenda17() *Addenda17 {
	addenda17 := NewAddenda17()
	addenda17.PaymentRelatedInformation = "Transfer of money from one country to another"
	addenda17.typeCode = "02"
	addenda17.SequenceNumber = 2
	addenda17.EntryDetailSequenceNumber = 0000002

	return addenda17
}

// TestMockIATBatch validates mockIATBatch
func TestMockIATBatch(t *testing.T) {
	iatBatch := mockIATBatch()
	if err := iatBatch.verify(); err != nil {
		t.Error("mockIATBatch does not validate and will break other tests")
	}
}

// testIATBatchAddenda10Error validates IATBatch returns an error if Addenda10 is not included
func testIATBatchAddenda10Error(t testing.TB) {
	iatBatch := mockIATBatch()
	iatBatch.GetEntries()[0].Addenda10 = nil
	if err := iatBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "FieldError" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIATBatchAddenda10Error tests validating IATBatch returns an error
// if Addenda10 is not included
func TestIATBatchAddenda10Error(t *testing.T) {
	testIATBatchAddenda10Error(t)
}

// BenchmarkIATBatchAddenda10Error benchmarks validating IATBatch returns an error
// if Addenda10 is not included
func BenchmarkIATBatchAddenda10Error(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda10Error(b)
	}
}

// testIATBatchAddenda11Error validates IATBatch returns an error if Addenda11 is not included
func testIATBatchAddenda11Error(t testing.TB) {
	iatBatch := mockIATBatch()
	iatBatch.GetEntries()[0].Addenda11 = nil
	if err := iatBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "FieldError" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIATBatchAddenda11Error tests validating IATBatch returns an error
// if Addenda11 is not included
func TestIATBatchAddenda11Error(t *testing.T) {
	testIATBatchAddenda11Error(t)
}

// BenchmarkIATBatchAddenda11Error benchmarks validating IATBatch returns an error
// if Addenda11 is not included
func BenchmarkIATBatchAddenda11Error(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda11Error(b)
	}
}

// testIATBatchAddenda12Error validates IATBatch returns an error if Addenda12 is not included
func testIATBatchAddenda12Error(t testing.TB) {
	iatBatch := mockIATBatch()
	iatBatch.GetEntries()[0].Addenda12 = nil
	if err := iatBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "FieldError" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIATBatchAddenda12Error tests validating IATBatch returns an error
// if Addenda12 is not included
func TestIATBatchAddenda12Error(t *testing.T) {
	testIATBatchAddenda12Error(t)
}

// BenchmarkIATBatchAddenda12Error benchmarks validating IATBatch returns an error
// if Addenda12 is not included
func BenchmarkIATBatchAddenda12Error(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda12Error(b)
	}
}

// testIATBatchAddenda13Error validates IATBatch returns an error if Addenda13 is not included
func testIATBatchAddenda13Error(t testing.TB) {
	iatBatch := mockIATBatch()
	iatBatch.GetEntries()[0].Addenda13 = nil
	if err := iatBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "FieldError" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIATBatchAddenda13Error tests validating IATBatch returns an error
// if Addenda13 is not included
func TestIATBatchAddenda13Error(t *testing.T) {
	testIATBatchAddenda13Error(t)
}

// BenchmarkIATBatchAddenda13Error benchmarks validating IATBatch returns an error
// if Addenda13 is not included
func BenchmarkIATBatchAddenda13Error(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda13Error(b)
	}
}

// testIATBatchAddenda14Error validates IATBatch returns an error if Addenda14 is not included
func testIATBatchAddenda14Error(t testing.TB) {
	iatBatch := mockIATBatch()
	iatBatch.GetEntries()[0].Addenda14 = nil
	if err := iatBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "FieldError" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIATBatchAddenda14Error tests validating IATBatch returns an error
// if Addenda14 is not included
func TestIATBatchAddenda14Error(t *testing.T) {
	testIATBatchAddenda14Error(t)
}

// BenchmarkIATBatchAddenda14Error benchmarks validating IATBatch returns an error
// if Addenda14 is not included
func BenchmarkIATBatchAddenda14Error(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda14Error(b)
	}
}

// testIATBatchAddenda15Error validates IATBatch returns an error if Addenda15 is not included
func testIATBatchAddenda15Error(t testing.TB) {
	iatBatch := mockIATBatch()
	iatBatch.GetEntries()[0].Addenda15 = nil
	if err := iatBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "FieldError" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIATBatchAddenda15Error tests validating IATBatch returns an error
// if Addenda15 is not included
func TestIATBatchAddenda15Error(t *testing.T) {
	testIATBatchAddenda15Error(t)
}

// BenchmarkIATBatchAddenda15Error benchmarks validating IATBatch returns an error
// if Addenda15 is not included
func BenchmarkIATBatchAddenda15Error(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda15Error(b)
	}
}

// testIATBatchAddenda16Error validates IATBatch returns an error if Addenda16 is not included
func testIATBatchAddenda16Error(t testing.TB) {
	iatBatch := mockIATBatch()
	iatBatch.GetEntries()[0].Addenda16 = nil
	if err := iatBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "FieldError" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIATBatchAddenda16Error tests validating IATBatch returns an error
// if Addenda16 is not included
func TestIATBatchAddenda16Error(t *testing.T) {
	testIATBatchAddenda16Error(t)
}

// BenchmarkIATBatchAddenda16Error benchmarks validating IATBatch returns an error
// if Addenda16 is not included
func BenchmarkIATBatchAddenda16Error(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda16Error(b)
	}
}

// testAddenda10EntryDetailSequenceNumber validates IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func testAddenda10EntryDetailSequenceNumber(t testing.TB) {
	iatBatch := mockIATBatch()
	iatBatch.GetEntries()[0].Addenda10.EntryDetailSequenceNumber = 00000005
	if err := iatBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TraceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddenda10EntryDetailSequenceNumber tests validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func TestAddenda10EntryDetailSequenceNumber(t *testing.T) {
	testAddenda10EntryDetailSequenceNumber(t)
}

// BenchmarkAddenda10EntryDetailSequenceNumber benchmarks validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func BenchmarkAddenda10EntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda10EntryDetailSequenceNumber(b)
	}
}

// testAddenda11EntryDetailSequenceNumber validates IATBatch returns an error if EntryDetailSequenceNumber
// is not valid
func testAddenda11EntryDetailSequenceNumber(t testing.TB) {
	iatBatch := mockIATBatch()
	iatBatch.GetEntries()[0].Addenda11.EntryDetailSequenceNumber = 00000005
	if err := iatBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TraceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddenda11EntryDetailSequenceNumber tests validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func TestAddenda11EntryDetailSequenceNumber(t *testing.T) {
	testAddenda11EntryDetailSequenceNumber(t)
}

// BenchmarkAddenda11EntryDetailSequenceNumber benchmarks validating IATBatch returns an error
// if EntryDetailSequenceNumber is not valid
func BenchmarkAddenda11EntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda11EntryDetailSequenceNumber(b)
	}
}

// testAddenda12EntryDetailSequenceNumber validates IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func testAddenda12EntryDetailSequenceNumber(t testing.TB) {
	iatBatch := mockIATBatch()
	iatBatch.GetEntries()[0].Addenda12.EntryDetailSequenceNumber = 00000005
	if err := iatBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TraceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddenda12EntryDetailSequenceNumber tests validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func TestAddenda12EntryDetailSequenceNumber(t *testing.T) {
	testAddenda12EntryDetailSequenceNumber(t)
}

// BenchmarkAddenda12EntryDetailSequenceNumber benchmarks validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func BenchmarkAddenda12EntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda12EntryDetailSequenceNumber(b)
	}
}

// testAddenda13EntryDetailSequenceNumber validates IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func testAddenda13EntryDetailSequenceNumber(t testing.TB) {
	iatBatch := mockIATBatch()
	iatBatch.GetEntries()[0].Addenda13.EntryDetailSequenceNumber = 00000005
	if err := iatBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TraceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddenda13EntryDetailSequenceNumber tests validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func TestAddenda13EntryDetailSequenceNumber(t *testing.T) {
	testAddenda13EntryDetailSequenceNumber(t)
}

// BenchmarkAddenda13EntryDetailSequenceNumber benchmarks validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func BenchmarkAddenda13EntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda13EntryDetailSequenceNumber(b)
	}
}

// testAddenda14EntryDetailSequenceNumber validates IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func testAddenda14EntryDetailSequenceNumber(t testing.TB) {
	iatBatch := mockIATBatch()
	iatBatch.GetEntries()[0].Addenda14.EntryDetailSequenceNumber = 00000005
	if err := iatBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TraceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddenda14EntryDetailSequenceNumber tests validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func TestAddenda14EntryDetailSequenceNumber(t *testing.T) {
	testAddenda14EntryDetailSequenceNumber(t)
}

// BenchmarkAddenda14EntryDetailSequenceNumber benchmarks validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func BenchmarkAddenda14EntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda14EntryDetailSequenceNumber(b)
	}
}

// testAddenda15EntryDetailSequenceNumber validates IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func testAddenda15EntryDetailSequenceNumber(t testing.TB) {
	iatBatch := mockIATBatch()
	iatBatch.GetEntries()[0].Addenda15.EntryDetailSequenceNumber = 00000005
	if err := iatBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TraceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddenda15EntryDetailSequenceNumber tests validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func TestAddenda15EntryDetailSequenceNumber(t *testing.T) {
	testAddenda15EntryDetailSequenceNumber(t)
}

// BenchmarkAddenda15EntryDetailSequenceNumber benchmarks validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func BenchmarkAddenda15EntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda15EntryDetailSequenceNumber(b)
	}
}

// testAddenda16EntryDetailSequenceNumber validates IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func testAddenda16EntryDetailSequenceNumber(t testing.TB) {
	iatBatch := mockIATBatch()
	iatBatch.GetEntries()[0].Addenda16.EntryDetailSequenceNumber = 00000005
	if err := iatBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TraceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddenda16EntryDetailSequenceNumber tests validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func TestAddenda16EntryDetailSequenceNumber(t *testing.T) {
	testAddenda16EntryDetailSequenceNumber(t)
}

// BenchmarkAddenda16EntryDetailSequenceNumber benchmarks validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func BenchmarkAddenda16EntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda16EntryDetailSequenceNumber(b)
	}
}

// testIATBatchNumberMismatch validates BatchNumber mismatch
func testIATBatchNumberMismatch(t testing.TB) {
	mockBatch := mockIATBatch()
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

// TestIATBatchNumberMismatch tests validating BatchNumber mismatch
func TestIATBatchNumberMismatch(t *testing.T) {
	testIATBatchNumberMismatch(t)
}

// BenchmarkIATBatchNumberMismatch benchmarks validating BatchNumber mismatch
func BenchmarkIATBatchNumberMismatch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchNumberMismatch(b)
	}
}

// testIATServiceClassCodeMismatch validates ServiceClassCode mismatch
func testIATServiceClassCodeMismatch(t testing.TB) {
	mockBatch := mockIATBatch()
	mockBatch.GetControl().ServiceClassCode = 225
	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIATServiceClassCodeMismatch tests validating ServiceClassCode mismatch
func TestServiceClassCodeMismatch(t *testing.T) {
	testIATServiceClassCodeMismatch(t)
}

// BenchmarkIATServiceClassCoderMismatch benchmarks validating ServiceClassCode mismatch
func BenchmarkIATServiceClassCodeMismatch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATServiceClassCodeMismatch(b)
	}
}

// testIATBatchCreditIsBatchAmount validates credit isBatchAmount
func testIATBatchCreditIsBatchAmount(t testing.TB) {
	mockBatch := mockIATBatch()
	e1 := mockBatch.GetEntries()[0]
	e2 := mockIATEntryDetail()
	e2.TransactionCode = 22
	e2.Amount = 5000
	e2.TraceNumber = e1.TraceNumber + 10
	mockBatch.AddEntry(e2)
	mockBatch.Entries[1].Addenda10 = mockAddenda10()
	mockBatch.Entries[1].Addenda11 = mockAddenda11()
	mockBatch.Entries[1].Addenda12 = mockAddenda12()
	mockBatch.Entries[1].Addenda13 = mockAddenda13()
	mockBatch.Entries[1].Addenda14 = mockAddenda14()
	mockBatch.Entries[1].Addenda15 = mockAddenda15()
	mockBatch.Entries[1].Addenda16 = mockAddenda16()
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	mockBatch.GetControl().TotalCreditEntryDollarAmount = 1000
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

// TestIATBatchCreditIsBatchAmount tests validating credit isBatchAmount
func TestIATBatchCreditIsBatchAmount(t *testing.T) {
	testIATBatchCreditIsBatchAmount(t)
}

// BenchmarkIATBatchCreditIsBatchAmount benchmarks validating credit isBatchAmount
func BenchmarkIATBatchCreditIsBatchAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchCreditIsBatchAmount(b)
	}

}

// testIATBatchDebitIsBatchAmount validates debit isBatchAmount
func testIATBatchDebitIsBatchAmount(t testing.TB) {
	mockBatch := mockIATBatch()
	e1 := mockBatch.GetEntries()[0]
	e1.TransactionCode = 27
	e2 := mockIATEntryDetail()
	e2.TransactionCode = 27
	e2.Amount = 5000
	e2.TraceNumber = e1.TraceNumber + 10
	mockBatch.AddEntry(e2)
	mockBatch.Entries[1].Addenda10 = mockAddenda10()
	mockBatch.Entries[1].Addenda11 = mockAddenda11()
	mockBatch.Entries[1].Addenda12 = mockAddenda12()
	mockBatch.Entries[1].Addenda13 = mockAddenda13()
	mockBatch.Entries[1].Addenda14 = mockAddenda14()
	mockBatch.Entries[1].Addenda15 = mockAddenda15()
	mockBatch.Entries[1].Addenda16 = mockAddenda16()
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	mockBatch.GetControl().TotalDebitEntryDollarAmount = 1000
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

// TestIATBatchDebitIsBatchAmount tests validating debit isBatchAmount
func TestIATBatchDebitIsBatchAmount(t *testing.T) {
	testIATBatchDebitIsBatchAmount(t)
}

// BenchmarkIATBatchDebitIsBatchAmount benchmarks validating debit isBatchAmount
func BenchmarkIATBatchDebitIsBatchAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchDebitIsBatchAmount(b)
	}

}

// testIATBatchFieldInclusion validates IATBatch FieldInclusion
func testIATBatchFieldInclusion(t testing.TB) {
	mockBatch := mockIATBatch()
	mockBatch2 := mockIATBatch()
	mockBatch2.Header.recordType = "4"

	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
	if err := mockBatch2.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
	if err := mockBatch2.build(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIATBatchFieldInclusion tests validating IATBatch FieldInclusion
func TestIATBatchFieldInclusion(t *testing.T) {
	testIATBatchFieldInclusion(t)
}

// BenchmarkIATBatchFieldInclusion benchmarks validating IATBatch FieldInclusion
func BenchmarkIATBatchFieldInclusion(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchFieldInclusion(b)
	}

}

// testIATBatchBuildError validates IATBatch build error
func testIATBatchBuild(t testing.TB) {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATBatchHeaderFF())

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

// TestIATBatchBuild tests validating IATBatch build error
func TestIATBatchBuild(t *testing.T) {
	testIATBatchBuild(t)
}

// BenchmarkIATBatchBuild benchmarks validating IATBatch build error
func BenchmarkIATBatchBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchBuild(b)
	}

}

// testIATODFIIdentificationMismatch validates ODFIIdentification mismatch
func testIATODFIIdentificationMismatch(t testing.TB) {
	mockBatch := mockIATBatch()
	mockBatch.GetControl().ODFIIdentification = "53158020"
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

// TestIATODFIIdentificationMismatch tests validating ODFIIdentification mismatch
func TestODFIIdentificationMismatch(t *testing.T) {
	testIATODFIIdentificationMismatch(t)
}

// BenchmarkIATODFIIdentificationMismatch benchmarks validating ODFIIdentification mismatch
func BenchmarkIATODFIIdentificationMismatch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATODFIIdentificationMismatch(b)
	}
}

// testIATBatchAddendaRecordIndicator validates IATEntryDetail AddendaRecordIndicator
func testIATBatchAddendaRecordIndicator(t testing.TB) {
	mockBatch := mockIATBatch()
	mockBatch.GetEntries()[0].AddendaRecordIndicator = 2
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

// TestIATBatchAddendaRecordIndicator tests validating IATEntryDetail AddendaRecordIndicator
func TestIATBatchAddendaRecordIndicator(t *testing.T) {
	testIATBatchAddendaRecordIndicator(t)
}

// BenchmarkIATBatchAddendaRecordIndicator benchmarks IATEntryDetail AddendaRecordIndicator
func BenchmarkIATBatchAddendaRecordIndicator(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddendaRecordIndicator(b)
	}
}

// testIATBatchInvalidTraceNumberODFI validates TraceNumberODFI
func testIATBatchInvalidTraceNumberODFI(t testing.TB) {
	mockBatch := mockIATBatch()
	mockBatch.GetEntries()[0].SetTraceNumber("9928272", 1)
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

// TestIATBatchInvalidTraceNumberODFI tests validating TraceNumberODFI
func TestIATBatchInvalidTraceNumberODFI(t *testing.T) {
	testIATBatchInvalidTraceNumberODFI(t)
}

// BenchmarkIATBatchInvalidTraceNumberODFI benchmarks validating TraceNumberODFI
func BenchmarkIATBatchInvalidTraceNumberODFI(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchInvalidTraceNumberODFI(b)
	}
}

// testIATBatchControl validates BatchControl ODFIIdentification
func testIATBatchControl(t testing.TB) {
	mockBatch := mockIATBatch()
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

// TestIATBatchControl tests validating BatchControl ODFIIdentification
func TestIATBatchControl(t *testing.T) {
	testIATBatchControl(t)
}

// BenchmarkIATBatchControl benchmarks validating BatchControl ODFIIdentification
func BenchmarkIATBatchControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchControl(b)
	}
}

// testIATBatchEntryCountEquality validates IATBatch EntryAddendaCount
func testIATBatchEntryCountEquality(t testing.TB) {
	mockBatch := mockIATBatch()
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

// TestIATBatchEntryCountEquality tests validating IATBatch EntryAddendaCount
func TestIATBatchEntryCountEquality(t *testing.T) {
	testIATBatchEntryCountEquality(t)
}

// BenchmarkIATBatchEntryCountEquality benchmarks validating IATBatch EntryAddendaCount
func BenchmarkIATBatchEntryCountEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchEntryCountEquality(b)
	}
}

// testIATBatchisEntryHash validates IATBatch EntryHash
func testIATBatchisEntryHash(t testing.TB) {
	mockBatch := mockIATBatch()
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

//TestIATBatchisEntryHash tests validating IATBatch EntryHash
func TestIATBatchisEntryHash(t *testing.T) {
	testIATBatchisEntryHash(t)
}

//BenchmarkIATBatchisEntryHash benchmarks validating IATBatch EntryHash
func BenchmarkIATBatchisEntryHash(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchisEntryHash(b)
	}
}

// testIATBatchIsSequenceAscending validates sequence ascending
func testIATBatchIsSequenceAscending(t testing.TB) {
	mockBatch := mockIATBatch()
	e2 := mockIATEntryDetail()
	e2.TraceNumber = 1
	mockBatch.AddEntry(e2)
	mockBatch.Entries[1].Addenda10 = mockAddenda10()
	mockBatch.Entries[1].Addenda11 = mockAddenda11()
	mockBatch.Entries[1].Addenda12 = mockAddenda12()
	mockBatch.Entries[1].Addenda13 = mockAddenda13()
	mockBatch.Entries[1].Addenda14 = mockAddenda14()
	mockBatch.Entries[1].Addenda15 = mockAddenda15()
	mockBatch.Entries[1].Addenda16 = mockAddenda16()
	mockBatch.GetControl().EntryAddendaCount = 16
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

// TestIATBatchIsSequenceAscending tests validating sequence ascending
func TestIATBatchIsSequenceAscending(t *testing.T) {
	testIATBatchIsSequenceAscending(t)
}

// BenchmarkIATBatchIsSequenceAscending tests validating sequence ascending
func BenchmarkIATBatchIsSequenceAscending(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchIsSequenceAscending(b)
	}
}

// testIATBatchIsCategory validates category
func testIATBatchIsCategory(t testing.TB) {
	mockBatch := mockIATBatchManyEntries()
	mockBatch.GetEntries()[1].Category = CategoryReturn

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

// TestIATBatchIsCategory tests validating category
func TestIATBatchIsCategory(t *testing.T) {
	testIATBatchIsCategory(t)
}

// BenchmarkIATBatchIsCategory tests validating category
func BenchmarkIATBatchIsCategory(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchIsCategory(b)
	}
}

//testIATBatchCategory tests IATBatch Category
func testIATBatchCategory(t testing.TB) {
	mockBatch := mockIATBatch()

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	if mockBatch.Category() != CategoryForward {
		t.Errorf("No returns and Category is %s", mockBatch.Category())
	}
}

// TestIATBatchCategory tests IATBatch Category
func TestIATBatchCategory(t *testing.T) {
	testIATBatchCategory(t)
}

// BenchmarkIATBatchCategory benchmarks IATBatch Category
func BenchmarkIATBatchCategory(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchCategory(b)
	}
}

// testIATBatchValidateEntry validates EntryDetail
func testIATBatchValidateEntry(t testing.TB) {
	mockBatch := mockIATBatch()
	mockBatch.GetEntries()[0].recordType = "5"

	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIATBatchValidateEntry tests validating Entry
func TestIATBatchValidateEntry(t *testing.T) {
	testIATBatchValidateEntry(t)
}

// BenchmarkIATBatchValidateEntry tests validating Entry
func BenchmarkIATBatchValidateEntry(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidateEntry(b)
	}
}

// testIATBatchValidateAddenda10 validates Addenda10
func testIATBatchValidateAddenda10(t testing.TB) {
	mockBatch := mockIATBatchManyEntries()
	mockBatch.GetEntries()[1].Addenda10.typeCode = "02"

	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIATBatchValidateAddenda10 tests validating Addenda10
func TestIATBatchValidateAddenda10(t *testing.T) {
	testIATBatchValidateAddenda10(t)
}

// BenchmarkIATBatchValidateAddenda10 tests validating Addenda10
func BenchmarkIATBatchValidateAddenda10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidateAddenda10(b)
	}
}

// testIATBatchValidateAddenda11 validates Addenda11
func testIATBatchValidateAddenda11(t testing.TB) {
	mockBatch := mockIATBatchManyEntries()
	mockBatch.GetEntries()[1].Addenda11.typeCode = "02"

	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIATBatchValidateAddenda11 tests validating Addenda11
func TestIATBatchValidateAddenda11(t *testing.T) {
	testIATBatchValidateAddenda11(t)
}

// BenchmarkIATBatchValidateAddenda11 tests validating Addenda11
func BenchmarkIATBatchValidateAddenda11(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidateAddenda11(b)
	}
}

// testIATBatchValidateAddenda12 validates Addenda12
func testIATBatchValidateAddenda12(t testing.TB) {
	mockBatch := mockIATBatchManyEntries()
	mockBatch.GetEntries()[1].Addenda12.typeCode = "02"

	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIATBatchValidateAddenda12 tests validating Addenda12
func TestIATBatchValidateAddenda12(t *testing.T) {
	testIATBatchValidateAddenda12(t)
}

// BenchmarkIATBatchValidateAddenda12 tests validating Addenda12
func BenchmarkIATBatchValidateAddenda12(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidateAddenda12(b)
	}
}

// testIATBatchValidateAddenda13 validates Addenda13
func testIATBatchValidateAddenda13(t testing.TB) {
	mockBatch := mockIATBatchManyEntries()
	mockBatch.GetEntries()[1].Addenda13.typeCode = "02"

	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIATBatchValidateAddenda13 tests validating Addenda13
func TestIATBatchValidateAddenda13(t *testing.T) {
	testIATBatchValidateAddenda13(t)
}

// BenchmarkIATBatchValidateAddenda13 tests validating Addenda13
func BenchmarkIATBatchValidateAddenda13(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidateAddenda13(b)
	}
}

// testIATBatchValidateAddenda14 validates Addenda14
func testIATBatchValidateAddenda14(t testing.TB) {
	mockBatch := mockIATBatchManyEntries()
	mockBatch.GetEntries()[1].Addenda14.typeCode = "02"

	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIATBatchValidateAddenda14 tests validating Addenda14
func TestIATBatchValidateAddenda14(t *testing.T) {
	testIATBatchValidateAddenda14(t)
}

// BenchmarkIATBatchValidateAddenda14 tests validating Addenda14
func BenchmarkIATBatchValidateAddenda14(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidateAddenda14(b)
	}
}

// testIATBatchValidateAddenda15 validates Addenda15
func testIATBatchValidateAddenda15(t testing.TB) {
	mockBatch := mockIATBatchManyEntries()
	mockBatch.GetEntries()[1].Addenda15.typeCode = "02"

	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIATBatchValidateAddenda15 tests validating Addenda15
func TestIATBatchValidateAddenda15(t *testing.T) {
	testIATBatchValidateAddenda15(t)
}

// BenchmarkIATBatchValidateAddenda15 tests validating Addenda15
func BenchmarkIATBatchValidateAddenda15(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidateAddenda15(b)
	}
}

// testIATBatchValidateAddenda16 validates Addenda16
func testIATBatchValidateAddenda16(t testing.TB) {
	mockBatch := mockIATBatchManyEntries()
	mockBatch.GetEntries()[1].Addenda16.typeCode = "02"

	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIATBatchValidateAddenda16 tests validating Addenda16
func TestIATBatchValidateAddenda16(t *testing.T) {
	testIATBatchValidateAddenda16(t)
}

// BenchmarkIATBatchValidateAddenda16 tests validating Addenda16
func BenchmarkIATBatchValidateAddenda16(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidateAddenda16(b)
	}
}

// testIATBatchValidateAddenda17 validates Addenda17
func testIATBatchValidateAddenda17(t testing.TB) {
	mockBatch := mockInvalidIATBatch()

	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestIATBatchValidateAddenda17 tests validating Addenda17
func TestIATBatchValidateAddenda17(t *testing.T) {
	testIATBatchValidateAddenda17(t)
}

// BenchmarkIATBatchValidateAddenda17 tests validating Addenda17
func BenchmarkIATBatchValidateAddenda17(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidateAddenda17(b)
	}
}

// testIATBatchCreateError validates IATBatch create error
func testIATBatchCreate(t testing.TB) {
	file := NewFile().SetHeader(mockFileHeader())
	mockBatch := mockIATBatch()
	mockBatch.GetHeader().recordType = "7"

	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}

	file.AddIATBatch(mockBatch)
}

// TestIATBatchCreate tests validating IATBatch create error
func TestIATBatchCreate(t *testing.T) {
	testIATBatchCreate(t)
}

// BenchmarkIATBatchCreate benchmarks validating IATBatch create error
func BenchmarkIATBatchCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchCreate(b)
	}

}

// testIATBatchValidate validates IATBatch validate error
func testIATBatchValidate(t testing.TB) {
	file := NewFile().SetHeader(mockFileHeader())
	mockBatch := mockIATBatch()
	mockBatch.GetHeader().ServiceClassCode = 225

	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}

	file.AddIATBatch(mockBatch)
}

// TestIATBatchValidate tests validating IATBatch validate error
func TestIATBatchValidate(t *testing.T) {
	testIATBatchValidate(t)
}

// BenchmarkIATBatchValidate benchmarks validating IATBatch validate error
func BenchmarkIATBatchValidate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidate(b)
	}

}

// testIATBatchEntryAddendum validates IATBatch EntryAddendum error
func testIATBatchEntryAddendum(t testing.TB) {
	file := NewFile().SetHeader(mockFileHeader())
	mockBatch := mockIATBatch()
	mockBatch.Entries[0].AddIATAddenda(mockAddenda17())
	mockBatch.Entries[0].AddIATAddenda(mockAddenda17B())
	mockBatch.Entries[0].AddIATAddenda(mockAddenda18())
	mockBatch.Entries[0].AddIATAddenda(mockAddenda18B())
	mockBatch.Entries[0].AddIATAddenda(mockAddenda18C())
	mockBatch.Entries[0].AddIATAddenda(mockAddenda18D())
	mockBatch.Entries[0].AddIATAddenda(mockAddenda18E())
	mockBatch.Entries[0].AddIATAddenda(mockAddenda18F())

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addendum" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}

	file.AddIATBatch(mockBatch)
}

// TestIATBatchEntryAddendum tests validating IATBatch EntryAddendum error
func TestIATBatchEntryAddendum(t *testing.T) {
	testIATBatchEntryAddendum(t)
}

// BenchmarkIATBatchEntryAddendum benchmarks validating IATBatch EntryAddendum error
func BenchmarkIATBatchEntryAddendum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchEntryAddendum(b)
	}
}
