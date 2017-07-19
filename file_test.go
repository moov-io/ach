// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
)

// TestBatchCountError tests for to many batch counts
func TestBatchCountError(t *testing.T) {
	r := NewReader(strings.NewReader(" "))
	r.File.AddBatch(NewBatchPPD())
	r.File.Control.BatchCount = 1
	if err := r.File.Validate(); err != nil {
		t.Errorf("Unexpected File.Validation error: %v", err.Error())
	}
	// More batches than the file control count.
	r.File.AddBatch(NewBatchPPD())
	if err := r.File.Validate(); err != nil {
		if err != ErrFileBatchCount {
			t.Errorf("Unexpected File.Validation error: %v", err.Error())
		}
	}
}

func TestFileEntryAddendaError(t *testing.T) {
	r := NewReader(strings.NewReader(" "))
	mockBatch := NewBatchPPD()
	mockBatch.GetControl().EntryAddendaCount = 1
	r.File.AddBatch(mockBatch)
	r.File.Control.BatchCount = 1
	r.File.Control.EntryAddendaCount = 1
	if err := r.File.Validate(); err != nil {
		t.Errorf("Unexpected File.Validation error: %v", err.Error())
	}

	// more entries than the file control
	r.File.Control.EntryAddendaCount = 5
	if err := r.File.Validate(); err != nil {
		if err != ErrFileEntryCount {
			t.Errorf("Unexpected File.Validation error: %v", err.Error())
		}
	}
}

func TestFileDebitAmount(t *testing.T) {

	r := NewReader(strings.NewReader(" "))
	mockBatch := NewBatchPPD()
	bc := BatchControl{
		EntryAddendaCount:           1,
		TotalDebitEntryDollarAmount: 105000,
	}
	mockBatch.SetControl(&bc)

	r.File.AddBatch(mockBatch)
	r.File.Control.BatchCount = 1
	r.File.Control.EntryAddendaCount = 1
	r.File.Control.TotalDebitEntryDollarAmountInFile = 105000

	if err := r.File.Validate(); err != nil {
		t.Errorf("Unexpected File.Validation error: %v", err.Error())
	}

	r.File.Control.TotalDebitEntryDollarAmountInFile = 0
	if err := r.File.Validate(); err != nil {
		if err != ErrFileDebitAmount {
			t.Errorf("Unexpected File.Validation error: %v", err.Error())
		}
	}
}

func TestFileCreditAmount(t *testing.T) {
	r := NewReader(strings.NewReader(" "))
	mockBatch := NewBatchPPD()
	bc := BatchControl{
		EntryAddendaCount:            1,
		TotalCreditEntryDollarAmount: 10500,
	}
	mockBatch.SetControl(&bc)

	r.File.AddBatch(mockBatch)
	r.File.Control.BatchCount = 1
	r.File.Control.EntryAddendaCount = 1
	r.File.Control.TotalCreditEntryDollarAmountInFile = 10500

	if err := r.File.Validate(); err != nil {
		t.Errorf("Unexpected File.Validation error: %v", err.Error())
	}

	r.File.Control.TotalCreditEntryDollarAmountInFile = 0
	if err := r.File.Validate(); err != nil {
		if err != ErrFileCreditAmount {
			t.Errorf("Unexpected File.Validation error: %v", err.Error())
		}
	}
}

func TestFileEntryHash(t *testing.T) {
	r := NewReader(strings.NewReader(" "))
	mockBatch1 := NewBatchPPD()
	bc := BatchControl{
		EntryAddendaCount:            1,
		TotalCreditEntryDollarAmount: 10500,
		EntryHash:                    1212121212,
	}
	mockBatch1.SetControl(&bc)

	mockBatch2 := NewBatchPPD()
	bc2 := BatchControl{
		EntryAddendaCount:            1,
		TotalCreditEntryDollarAmount: 10500,
		EntryHash:                    2121212121,
	}
	mockBatch2.SetControl(&bc2)

	r.File.AddBatch(mockBatch1)
	r.File.AddBatch(mockBatch2)

	r.File.Control.BatchCount = 2
	r.File.Control.EntryAddendaCount = 2
	r.File.Control.TotalCreditEntryDollarAmountInFile = 21000
	r.File.Control.EntryHash = 3333333333
	if err := r.File.Validate(); err != nil {
		t.Errorf("Unexpected File.Validation error: %v", err.Error())
	}

	r.File.Control.EntryHash = 0
	if err := r.File.Validate(); err != nil {
		if err != ErrFileEntryHash {
			t.Errorf("Unexpected File.Validation error: %v", err.Error())
		}
	}
}

func TestFileBlockCount10(t *testing.T) {
	file := NewFile().SetHeader(mockFileHeader())
	batch := NewBatchPPD()
	batch.SetHeader(mockBatchHeader())
	batch.AddEntry(mockEntryDetail())
	batch.AddEntry(mockEntryDetail())
	batch.AddEntry(mockEntryDetail())
	batch.AddEntry(mockEntryDetail())
	batch.AddEntry(mockEntryDetail())
	batch.AddEntry(mockEntryDetail())
	batch.Build()
	file.AddBatch(batch)

	// ensure with 10 records in file we don't get 2 for a block count
	if err := file.Build(); err != nil {
		t.Errorf("Unexpected File.Validation error: %v", err.Error())
	}
	if file.Control.BlockCount != 1 {
		t.Errorf("Unexpected block count in file expect 1 got: %v", file.Control.BlockCount)
	}
}

func TestFileBuildBadFileHeader(t *testing.T) {
	file := NewFile().SetHeader(FileHeader{})
	if err := file.Build(); err != nil {
		if !strings.Contains(err.Error(), ErrRecordType.Error()) {
			t.Errorf("Unexpected File.Build error: %v", err.Error())
		}
	}
}

func TestFileBuildNoBatch(t *testing.T) {
	file := NewFile().SetHeader(mockFileHeader())
	if err := file.Build(); err != nil {
		if !strings.Contains(err.Error(), ErrFileBatches.Error()) {
			t.Errorf("Unexpected File.Build error: %v", err.Error())
		}
	}
}

func TestFileValidateAllBatch(t *testing.T) {
	file := NewFile().SetHeader(mockFileHeader())
	batch := NewBatchPPD()
	batch.SetHeader(mockBatchHeader())
	batch.AddEntry(mockEntryDetail())
	batch.Build()
	file.AddBatch(batch)
	// break the file header
	file.Batches[0].GetHeader().ODFIIdentification = 0
	if err := file.Build(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected File.ValidationAll error: %v", err.Error())
		}
	}
}

func TestFileValidateAllFileHeader(t *testing.T) {
	file := NewFile().SetHeader(mockFileHeader())
	batch := NewBatchPPD()
	batch.SetHeader(mockBatchHeader())
	batch.AddEntry(mockEntryDetail())
	batch.Build()
	file.AddBatch(batch)
	// break the file header
	file.Header.ImmediateOrigin = 0
	if err := file.Build(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected File.ValidationAll error: %v", err.Error())
		}
	}
}

func TestFileValidateAllFileControl(t *testing.T) {
	file := NewFile().SetHeader(mockFileHeader())
	batch := NewBatchPPD()
	batch.SetHeader(mockBatchHeader())
	batch.AddEntry(mockEntryDetail())
	batch.Build()
	file.AddBatch(batch)
	// break the file header
	file.Control.BatchCount = 0
	if err := file.Build(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected File.ValidationAll error: %v", err.Error())
		}
	}
}
