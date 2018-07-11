// Copyright 2018 The ACH Authors
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
	entry.AddAddenda(mockAddenda05())
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
	entry.AddAddenda(mockAddenda05())
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
	iatBatch.Create()
	file.AddIATBatch(iatBatch)

	/*	iatBatch2 := IATBatch{}
		iatBatch2.SetHeader(mockIATBatchHeaderFF())
		iatBatch2.AddEntry(mockIATEntryDetail())
		iatBatch2.Entries[0].Addenda10 = mockAddenda10()
		iatBatch2.Entries[0].Addenda11 = mockAddenda11()
		iatBatch2.Entries[0].Addenda12 = mockAddenda12()
		iatBatch2.Entries[0].Addenda13 = mockAddenda13()
		iatBatch2.Entries[0].Addenda14 = mockAddenda14()
		iatBatch2.Entries[0].Addenda15 = mockAddenda15()
		iatBatch2.Entries[0].Addenda16 = mockAddenda16()
		iatBatch2.Create()
		file.AddIATBatch(iatBatch2)*/

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
