package ach

import (
	"bytes"
	"strings"
	"testing"
)

func TestPPDWrite(t *testing.T) {
	file := NewFile().SetHeader(mockFileHeader())
	entry := mockEntryDetail()
	entry.AddAddenda(mockAddenda())
	batch := NewBatchPPD()
	batch.SetHeader(mockBatchHeader())
	batch.AddEntry(entry)
	batch.Build()
	file.AddBatch(batch)

	if err := file.Build(); err != nil {
		t.Errorf("Could not build file: %v", err)
	}
	if err := file.ValidateAll(); err != nil {
		t.Errorf("Could not validate built file: %v", err)
	}

	b := &bytes.Buffer{}
	f := NewWriter(b)

	if err := f.WriteAll([]*File{file}); err != nil {
		t.Errorf("Unexpected error: %s\n", err)
	}
	r := NewReader(strings.NewReader(b.String()))
	_, err := r.Read()
	if err != nil {
		t.Errorf("Can not ach.Read generated file: %v", err)
	}
	if err = r.File.ValidateAll(); err != nil {
		t.Errorf("Could not validate entire generated file: %v", err)
	}
}
