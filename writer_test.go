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

	if err := f.WriteAll([]*File{file}); err != nil {
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

// BenchmarkPPDWrite benchmarks writing a PPD ACH file
func BenchmarkPPDWrite(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testPPDWrite(b)
	}
}
