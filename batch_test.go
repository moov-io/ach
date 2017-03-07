// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"
	"time"
)

// TestBatchEntryCountMismatc check for control out-of-balance error.
func TestBatchEntryCountMismatch(t *testing.T) {
	a := NewAddenda()
	e := NewEntryDetail()
	e.AddendaRecordIndicator = 1
	e.addAddenda(a)
	mockBatch := Batch{}
	mockBatch.addEntryDetail(e)
	mockBatch.Control.EntryAddendaCount = 2
	// works properly
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// create error is mismatch
	mockBatch.Control.EntryAddendaCount = 1
	if err := mockBatch.Validate(); err != nil {
		if err != ErrBatchEntryCountMismatch {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBatchServiceClassMismatch(t *testing.T) {
	mockBatch := Batch{}
	mockBatch.Header.ServiceClassCode = 220
	mockBatch.Control.ServiceClassCode = 220
	// works properly
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// create error is mismatch
	mockBatch.Control.ServiceClassCode = 225
	if err := mockBatch.Validate(); err != nil {
		if err != ErrBatchServiceClassMismatch {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBatchCompanyIdentification(t *testing.T) {
	mockBatch := Batch{}
	mockBatch.Header.CompanyIdentification = "ABC Corp"
	mockBatch.Control.CompanyIdentification = "ABC Corp"
	// works properly
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// create error is mismatch
	mockBatch.Control.CompanyIdentification = "XYZ Inc"
	if err := mockBatch.Validate(); err != nil {
		if err != ErrBatchCompanyID {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBatchODFIIDMismatch(t *testing.T) {
	mockBatch := Batch{}
	mockBatch.Header.ODFIIdentification = 123456789
	mockBatch.Control.ODFIIdentification = 123456789
	// works properly
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// create error is mismatch
	mockBatch.Control.ODFIIdentification = 987654321
	if err := mockBatch.Validate(); err != nil {
		if err != ErrBatchODFIIDMismatch {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBatchNumberMismatch(t *testing.T) {
	mockBatch := Batch{}
	mockBatch.Header.BatchNumber = 1
	mockBatch.Control.BatchNumber = 1
	// works properly
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// create error is mismatch
	mockBatch.Control.BatchNumber = 2
	if err := mockBatch.Validate(); err != nil {
		if err != ErrBatchNumberMismatch {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBatchIsSequenceAscending(t *testing.T) {
	e1 := NewEntryDetail()
	e2 := NewEntryDetail()
	e2.TraceNumber = e1.TraceNumber + 1
	mockBatch := Batch{}
	mockBatch.addEntryDetail(e1)
	mockBatch.addEntryDetail(e2)
	mockBatch.Control.EntryAddendaCount = 2
	// works properly
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// create error
	e3 := NewEntryDetail()
	e3.TraceNumber = e2.TraceNumber - 1
	mockBatch.addEntryDetail(e3)
	mockBatch.Control.EntryAddendaCount = 3
	if err := mockBatch.Validate(); err != nil {
		if err != ErrBatchAscendingTraceNumber {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

// isBatchAmountMismatch
func TestCreditBatchIsBatchAmountMismatch(t *testing.T) {
	e1 := NewEntryDetail()
	e1.TransactionCode = 22
	e1.Amount = 100
	e2 := NewEntryDetail()
	e2.TransactionCode = 22
	e2.Amount = 100
	e2.TraceNumber = e1.TraceNumber + 1
	mockBatch := Batch{}
	mockBatch.addEntryDetail(e1)
	mockBatch.addEntryDetail(e2)
	mockBatch.Control.EntryAddendaCount = 2
	mockBatch.Control.TotalCreditEntryDollarAmount = 200
	// works properly
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// create error
	mockBatch.Control.TotalCreditEntryDollarAmount = 1
	if err := mockBatch.Validate(); err != nil {
		if err != ErrBatchAmountMismatch {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestSavingsBatchIsBatchAmountMismatch(t *testing.T) {
	e1 := NewEntryDetail()
	e1.TransactionCode = 32
	e1.Amount = 100
	e2 := NewEntryDetail()
	e2.TransactionCode = 37
	e2.Amount = 100
	e2.TraceNumber = e1.TraceNumber + 1
	mockBatch := Batch{}
	mockBatch.addEntryDetail(e1)
	mockBatch.addEntryDetail(e2)
	mockBatch.Control.EntryAddendaCount = 2
	mockBatch.Control.TotalCreditEntryDollarAmount = 100
	mockBatch.Control.TotalDebitEntryDollarAmount = 100

	// works properly
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// create error
	mockBatch.Control.TotalDebitEntryDollarAmount = 1
	if err := mockBatch.Validate(); err != nil {
		if err != ErrBatchAmountMismatch {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBatchIsEntryHashMismatch(t *testing.T) {
	e1 := NewEntryDetail()
	e1.RDFIIdentification = 111111111
	e2 := NewEntryDetail()
	e2.RDFIIdentification = 111111111
	e2.TraceNumber = e1.TraceNumber + 1
	mockBatch := Batch{}
	mockBatch.addEntryDetail(e1)
	mockBatch.addEntryDetail(e2)
	mockBatch.Control.EntryAddendaCount = 2
	mockBatch.Control.validate = 222222222
	// works properly
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// create error
	mockBatch.Control.validate = 1
	if err := mockBatch.Validate(); err != nil {
		if err != ErrValidEntryHash {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

// isOriginatorDNEMismatch
func TestBatchIsOriginatorDNEMismatch(t *testing.T) {
	e1 := NewEntryDetail()
	e1.TransactionCode = 23
	e2 := NewEntryDetail()
	e2.TraceNumber = e1.TraceNumber + 1
	mockBatch := Batch{}
	mockBatch.addEntryDetail(e1)
	mockBatch.addEntryDetail(e2)
	mockBatch.Header.OriginatorStatusCode = 2
	mockBatch.Control.EntryAddendaCount = 2
	// works properly
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// Make it fail
	mockBatch.Header.OriginatorStatusCode = 1

	if err := mockBatch.Validate(); err != nil {
		if err != ErrBatchOriginatorDNE {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

// ErrBatchTraceNumberNotODFI
func TestBatchTraceNumberNotODFI(t *testing.T) {
	odfi := 123456789
	e1 := NewEntryDetail()
	e1.TraceNumber = 1234567890000001
	mockBatch := Batch{}
	mockBatch.addEntryDetail(e1)
	mockBatch.Header.ODFIIdentification = odfi
	mockBatch.Control.ODFIIdentification = odfi
	mockBatch.Control.EntryAddendaCount = 1
	// works properly
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// Make it fail
	mockBatch.Entries[0].TraceNumber = 1
	if err := mockBatch.Validate(); err != nil {
		if err != ErrBatchTraceNumberNotODFI {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

// ErrBatchAddendaIndicator
func TestBatchAddendaIndicator(t *testing.T) {
	a := NewAddenda()
	e1 := NewEntryDetail()
	e1.AddendaRecordIndicator = 1
	e1.addAddenda(a)

	mockBatch := Batch{}
	mockBatch.addEntryDetail(e1)
	mockBatch.Control.EntryAddendaCount = 2
	// works properly
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// Make it fail
	mockBatch.Entries[0].AddendaRecordIndicator = 0
	if err := mockBatch.Validate(); err != nil {
		if err != ErrBatchAddendaIndicator {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

// isAddendaSequence
func TestBatchAddendaSequence(t *testing.T) {
	a1 := NewAddenda()
	a1.SequenceNumber = 1
	a2 := NewAddenda()
	a2.SequenceNumber = 2
	e1 := NewEntryDetail()
	e1.AddendaRecordIndicator = 1
	e1.addAddenda(a1)
	e1.addAddenda(a2)
	mockBatch := Batch{}
	mockBatch.addEntryDetail(e1)
	mockBatch.Control.EntryAddendaCount = 3
	// works properly
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// Make it fail
	mockBatch.Entries[0].Addendums[0].SequenceNumber = 10
	if err := mockBatch.Validate(); err != nil {
		if err != ErrBatchAddendaSequence {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

// ErrBatchAddendaTraceNumber
func TestBatchAddendaTraceNumber(t *testing.T) {
	a1 := NewAddenda()
	a1.SequenceNumber = 1
	a1.EntryDetailSequenceNumber = 0000001
	a2 := NewAddenda()
	a2.SequenceNumber = 2
	a2.EntryDetailSequenceNumber = 0000001
	e1 := NewEntryDetail()
	e1.AddendaRecordIndicator = 1
	e1.TraceNumber = 1234567890000001
	e1.addAddenda(a1)
	e1.addAddenda(a2)
	mockBatch := Batch{}
	mockBatch.addEntryDetail(e1)
	mockBatch.Control.EntryAddendaCount = 3
	mockBatch.Header.ODFIIdentification = 123456789
	mockBatch.Control.ODFIIdentification = 123456789
	// works properly
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// Make it fail
	mockBatch.Entries[0].Addendums[0].EntryDetailSequenceNumber = 99
	if err := mockBatch.Validate(); err != nil {
		if err != ErrBatchAddendaTraceNumber {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBatchBuild(t *testing.T) {
	mockBatch := Batch{}
	header := NewBatchHeader()
	header.ServiceClassCode = 200
	header.CompanyName = "MY BEST COMP."
	header.CompanyDiscretionaryData = "INCLUDES OVERTIME"
	header.CompanyIdentification = "1419871234"
	header.StandardEntryClassCode = "PPD"
	header.CompanyEntryDescription = "PAYROLL"
	header.EffectiveEntryDate = time.Now()
	header.ODFIIdentification = 109991234
	mockBatch.Header = header

	entry := NewEntryDetail()
	entry.TransactionCode = 22                            // ACH Credit
	entry.setRDFI(81086674)                               // scottrade bank routing number
	entry.dfiAccountNumber = "62292250"                   // scottrade account number
	entry.Amount = 1000000                                // 1k dollars
	entry.IndividualIdentificationNumber = "658-888-2468" // Unique ID for payment
	entry.IndividualName = "Wade Arnold"
	entry.setTraceNumber(header.ODFIIdentification, 1)
	a1 := NewAddenda()
	a2 := NewAddenda()
	entry.addAddenda(a1)
	entry.addAddenda(a2)
	mockBatch.addEntryDetail(entry)
	if err := mockBatch.Build(); err != nil {
		t.Errorf("Unexpected Batch.Build error: %v", err.Error())
	}
	//fmt.Printf("Batch: %+v \n", mockBatch)

	if err := mockBatch.ValidateAll(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
}
