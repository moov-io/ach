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
	r.File.AddBatch(NewBatch())
	r.File.Control.BatchCount = 1
	if err := r.File.Validate(); err != nil {
		t.Errorf("Unexpected File.Validation error: %v", err.Error())
	}
	// More batches than the file control count.
	r.File.AddBatch(&Batch{})
	if err := r.File.Validate(); err != nil {
		if err != ErrFileBatchCount {
			t.Errorf("Unexpected File.Validation error: %v", err.Error())
		}
	}
}

func TestFileEntryAddendaError(t *testing.T) {
	r := NewReader(strings.NewReader(" "))
	mockBatch := NewBatch()
	mockBatch.Control.EntryAddendaCount = 1
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
	mockBatch := NewBatch()
	mockBatch.Control.EntryAddendaCount = 1
	mockBatch.Control.TotalDebitEntryDollarAmount = 10500

	r.File.AddBatch(mockBatch)
	r.File.Control.BatchCount = 1
	r.File.Control.EntryAddendaCount = 1
	r.File.Control.TotalDebitEntryDollarAmountInFile = 10500

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
	mockBatch := NewBatch()
	mockBatch.Control.EntryAddendaCount = 1
	mockBatch.Control.TotalCreditEntryDollarAmount = 10500

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
	mockBatch1 := NewBatch()
	mockBatch1.Control.EntryAddendaCount = 1
	mockBatch1.Control.TotalCreditEntryDollarAmount = 10500
	mockBatch1.Control.EntryHash = 1212121212

	mockBatch2 := NewBatch()
	mockBatch2.Control.EntryAddendaCount = 1
	mockBatch2.Control.TotalCreditEntryDollarAmount = 10500
	mockBatch2.Control.EntryHash = 2121212121

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
