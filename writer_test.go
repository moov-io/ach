// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"bytes"
	"strings"
	"testing"
)

// testPPDWrite writes a PPD ACH file
func testPPDWrite(t testing.TB) {
	file := NewFile().SetHeader(mockFileHeader())
	entry := mockEntryDetail()
	entry.AddendaRecordIndicator = 1
	entry.AddAddenda05(mockAddenda05())
	batch := NewBatchPPD(mockBatchPPDHeader())
	batch.SetHeader(mockBatchHeader())
	batch.AddEntry(entry)
	batch.Create()
	file.AddBatch(batch)

	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	b := &bytes.Buffer{}
	f := NewWriter(b)

	if err := f.Write(file); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	r := NewReader(strings.NewReader(b.String()))
	_, err := r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err = r.File.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestPPDWrite tests writing a PPD ACH file
func TestPPDWrite(t *testing.T) {
	testPPDWrite(t)
}

// BenchmarkPPDWrite benchmarks validating writing a PPD ACH file
func BenchmarkPPDWrite(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testPPDWrite(b)
	}
}

// testFileWriteErr validates error for file write
func testFileWriteErr(t testing.TB) {
	file := NewFile().SetHeader(mockFileHeader())
	entry := mockEntryDetail()
	entry.AddendaRecordIndicator = 1
	entry.AddAddenda05(mockAddenda05())
	batch := NewBatchPPD(mockBatchPPDHeader())
	batch.SetHeader(mockBatchHeader())
	batch.AddEntry(entry)
	batch.Create()
	if err := batch.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	file.AddBatch(batch)

	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	file.Batches[0].GetControl().EntryAddendaCount = 10

	b := &bytes.Buffer{}
	f := NewWriter(b)

	if err := f.Write(file); err != nil {
		if e, ok := err.(*FileError); ok {
			if e.FieldName != "EntryAddendaCount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFileWriteErr tests validating error for file write
func TestFileWriteErr(t *testing.T) {
	testFileWriteErr(t)
}

// BenchmarkFileWriteErr benchmarks error for file write
func BenchmarkFileWriteErr(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileWriteErr(b)
	}
}

// testIATWrite writes a IAT ACH file
func testIATWrite(t testing.TB) {
	file := NewFile().SetHeader(mockFileHeader())
	iatBatch := IATBatch{}
	iatBatch.SetHeader(mockIATBatchHeaderFF())
	iatBatch.AddEntry(mockIATEntryDetail())
	iatBatch.Entries[0].Addenda10 = mockAddenda10()
	iatBatch.Entries[0].Addenda11 = mockAddenda11()
	iatBatch.Entries[0].Addenda12 = mockAddenda12()
	iatBatch.Entries[0].Addenda13 = mockAddenda13()
	iatBatch.Entries[0].Addenda14 = mockAddenda14()
	iatBatch.Entries[0].Addenda15 = mockAddenda15()
	iatBatch.Entries[0].Addenda16 = mockAddenda16()
	iatBatch.Entries[0].AddAddenda17(mockAddenda17())
	iatBatch.Entries[0].AddAddenda17(mockAddenda17B())
	iatBatch.Entries[0].AddAddenda18(mockAddenda18())
	iatBatch.Entries[0].AddAddenda18(mockAddenda18B())
	iatBatch.Entries[0].AddAddenda18(mockAddenda18C())
	iatBatch.Entries[0].AddAddenda18(mockAddenda18D())
	iatBatch.Entries[0].AddAddenda18(mockAddenda18E())
	iatBatch.Create()
	file.AddIATBatch(iatBatch)

	iatBatch2 := IATBatch{}
	iatBatch2.SetHeader(mockIATBatchHeaderFF())
	iatBatch2.AddEntry(mockIATEntryDetail())
	iatBatch2.GetEntries()[0].TransactionCode = 27
	iatBatch2.GetEntries()[0].Amount = 2000
	iatBatch2.Entries[0].Addenda10 = mockAddenda10()
	iatBatch2.Entries[0].Addenda11 = mockAddenda11()
	iatBatch2.Entries[0].Addenda12 = mockAddenda12()
	iatBatch2.Entries[0].Addenda13 = mockAddenda13()
	iatBatch2.Entries[0].Addenda14 = mockAddenda14()
	iatBatch2.Entries[0].Addenda15 = mockAddenda15()
	iatBatch2.Entries[0].Addenda16 = mockAddenda16()
	iatBatch2.Entries[0].AddAddenda17(mockAddenda17())
	iatBatch2.Entries[0].AddAddenda17(mockAddenda17B())
	iatBatch2.Entries[0].AddAddenda18(mockAddenda18())
	iatBatch2.Entries[0].AddAddenda18(mockAddenda18B())
	iatBatch2.Entries[0].AddAddenda18(mockAddenda18C())
	iatBatch2.Entries[0].AddAddenda18(mockAddenda18D())
	iatBatch2.Entries[0].AddAddenda18(mockAddenda18E())
	iatBatch2.Create()
	file.AddIATBatch(iatBatch2)

	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	b := &bytes.Buffer{}
	f := NewWriter(b)

	if err := f.Write(file); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	r := NewReader(strings.NewReader(b.String()))
	_, err := r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err = r.File.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	/*	// Write IAT records to standard output. Anything io.Writer
		w := NewWriter(os.Stdout)
		if err := w.Write(file); err != nil {
			log.Fatalf("Unexpected error: %s\n", err)
		}
		w.Flush()*/
}

// TestIATWrite tests writing a IAT ACH file
func TestIATWrite(t *testing.T) {
	testIATWrite(t)
}

// BenchmarkIATWrite benchmarks validating writing a IAT ACH file
func BenchmarkIATWrite(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATWrite(b)
	}
}

// testPPDIATWrite writes an ACH file which writing an ACH file which contains PPD and IAT entries
func testPPDIATWrite(t testing.TB) {
	file := NewFile().SetHeader(mockFileHeader())

	entry := mockEntryDetail()
	entry.AddendaRecordIndicator = 1
	entry.AddAddenda05(mockAddenda05())
	batch := NewBatchPPD(mockBatchPPDHeader())
	batch.SetHeader(mockBatchHeader())
	batch.AddEntry(entry)
	batch.Create()
	file.AddBatch(batch)

	iatBatch := IATBatch{}
	iatBatch.SetHeader(mockIATBatchHeaderFF())
	iatBatch.AddEntry(mockIATEntryDetail())
	iatBatch.Entries[0].Addenda10 = mockAddenda10()
	iatBatch.Entries[0].Addenda11 = mockAddenda11()
	iatBatch.Entries[0].Addenda12 = mockAddenda12()
	iatBatch.Entries[0].Addenda13 = mockAddenda13()
	iatBatch.Entries[0].Addenda14 = mockAddenda14()
	iatBatch.Entries[0].Addenda15 = mockAddenda15()
	iatBatch.Entries[0].Addenda16 = mockAddenda16()
	iatBatch.Entries[0].AddAddenda17(mockAddenda17())
	iatBatch.Entries[0].AddAddenda17(mockAddenda17B())
	iatBatch.Entries[0].AddAddenda18(mockAddenda18())
	iatBatch.Entries[0].AddAddenda18(mockAddenda18B())
	iatBatch.Entries[0].AddAddenda18(mockAddenda18C())
	iatBatch.Entries[0].AddAddenda18(mockAddenda18D())
	iatBatch.Entries[0].AddAddenda18(mockAddenda18E())
	iatBatch.Create()
	file.AddIATBatch(iatBatch)

	iatBatch2 := IATBatch{}
	iatBatch2.SetHeader(mockIATBatchHeaderFF())
	iatBatch2.AddEntry(mockIATEntryDetail())
	iatBatch2.GetEntries()[0].TransactionCode = 27
	iatBatch2.GetEntries()[0].Amount = 2000
	iatBatch2.Entries[0].Addenda10 = mockAddenda10()
	iatBatch2.Entries[0].Addenda11 = mockAddenda11()
	iatBatch2.Entries[0].Addenda12 = mockAddenda12()
	iatBatch2.Entries[0].Addenda13 = mockAddenda13()
	iatBatch2.Entries[0].Addenda14 = mockAddenda14()
	iatBatch2.Entries[0].Addenda15 = mockAddenda15()
	iatBatch2.Entries[0].Addenda16 = mockAddenda16()
	iatBatch2.Entries[0].AddAddenda17(mockAddenda17())
	iatBatch2.Entries[0].AddAddenda17(mockAddenda17B())
	iatBatch2.Entries[0].AddAddenda18(mockAddenda18())
	iatBatch2.Entries[0].AddAddenda18(mockAddenda18B())
	iatBatch2.Entries[0].AddAddenda18(mockAddenda18C())
	iatBatch2.Entries[0].AddAddenda18(mockAddenda18D())
	iatBatch2.Entries[0].AddAddenda18(mockAddenda18E())
	iatBatch2.Create()
	file.AddIATBatch(iatBatch2)

	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	b := &bytes.Buffer{}
	f := NewWriter(b)

	if err := f.Write(file); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	r := NewReader(strings.NewReader(b.String()))
	_, err := r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err = r.File.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	/*	// Write records to standard output. Anything io.Writer
		w := NewWriter(os.Stdout)
		if err := w.Write(file); err != nil {
			log.Fatalf("Unexpected error: %s\n", err)
		}
		w.Flush()*/
}

// TestPPDIATWrite tests writing a IAT ACH file
func TestPPDIATWrite(t *testing.T) {
	testPPDIATWrite(t)
}

// BenchmarkPPDIATWrite benchmarks validating writing an ACH file which contain PPD and IAT entries
func BenchmarkPPDIATWrite(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testPPDIATWrite(b)
	}
}

// testIATReturn writes a IAT ACH Return file
func testIATReturn(t testing.TB) {
	file := NewFile().SetHeader(mockFileHeader())
	iatBatch := IATBatch{}
	iatBatch.SetHeader(mockIATBatchHeaderFF())
	iatBatch.AddEntry(mockIATEntryDetail())
	iatBatch.Entries[0].Addenda10 = mockAddenda10()
	iatBatch.Entries[0].Addenda11 = mockAddenda11()
	iatBatch.Entries[0].Addenda12 = mockAddenda12()
	iatBatch.Entries[0].Addenda13 = mockAddenda13()
	iatBatch.Entries[0].Addenda14 = mockAddenda14()
	iatBatch.Entries[0].Addenda15 = mockAddenda15()
	iatBatch.Entries[0].Addenda16 = mockAddenda16()
	iatBatch.Entries[0].AddIATAddenda(mockIATAddenda99())
	iatBatch.Entries[0].Category = CategoryReturn
	iatBatch.Create()
	file.AddIATBatch(iatBatch)

	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	b := &bytes.Buffer{}
	f := NewWriter(b)

	if err := f.Write(file); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	r := NewReader(strings.NewReader(b.String()))
	_, err := r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err = r.File.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	/*		// Write IAT records to standard output. Anything io.Writer
			w := NewWriter(os.Stdout)
			if err := w.Write(file); err != nil {
				log.Fatalf("Unexpected error: %s\n", err)
			}
			w.Flush()*/
}

// TestIATReturn tests writing a IAT ACH Return file
func TestIATReturn(t *testing.T) {
	testIATReturn(t)
}

// BenchmarkIATReturn benchmarks validating writing a IAT ACH Return file
func BenchmarkIATReturn(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATReturn(b)
	}
}
