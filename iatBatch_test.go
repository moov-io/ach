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

// testIATAddendaRecordIndicator validates AddendaRecordIndicator FieldInclusion
func testIATAddendaRecordIndicator(t testing.TB) {
	mockBatch := mockIATBatch()
	mockBatch.GetEntries()[0].AddendaRecordIndicator = 0
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

// TestIATAddendaRecordIndicator tests validating AddendaRecordIndicator FieldInclusion
func TestIATAddendaRecordIndicator(t *testing.T) {
	testIATAddendaRecordIndicator(t)
}

// BenchmarkIATAddendaRecordIndicator benchmarks validating AddendaRecordIndicator FieldInclusion
func BenchmarkIATAddendaRecordIndicator(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATAddendaRecordIndicator(b)
	}
}
