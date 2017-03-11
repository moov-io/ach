// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
	"time"
)

// TestBatchEntryCountMismatc check for control out-of-balance error.
func TestBatchEntryCountMismatch(t *testing.T) {
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	e := mockEntryDetail()
	a := mockAddenda()
	e.AddAddenda(a)
	mockBatch.AddEntryDetail(e)
	mockBatch.Build()
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
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	mockBatch.AddEntryDetail(mockEntryDetail())
	mockBatch.Build()

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
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	mockBatch.AddEntryDetail(mockEntryDetail())
	mockBatch.Build()
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
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	mockBatch.AddEntryDetail(mockEntryDetail())
	mockBatch.Build()
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
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	mockBatch.AddEntryDetail(mockEntryDetail())
	mockBatch.Build()
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
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	mockBatch.AddEntryDetail(mockEntryDetail())
	mockBatch.AddEntryDetail(mockEntryDetail())
	mockBatch.Build()
	// works properly
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// create error
	e3 := mockEntryDetail()
	e3.TraceNumber = 0
	mockBatch.AddEntryDetail(e3)
	mockBatch.Control.EntryAddendaCount = 3
	if err := mockBatch.Validate(); err != nil {
		if err != ErrBatchAscendingTraceNumber {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

// isBatchAmountMismatch
func TestCreditBatchIsBatchAmountMismatch(t *testing.T) {
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	e1 := mockEntryDetail()
	e1.TransactionCode = 22
	e1.Amount = 100
	e2 := mockEntryDetail()
	e2.TransactionCode = 22
	e2.Amount = 100
	mockBatch.AddEntryDetail(e1)
	mockBatch.AddEntryDetail(e2)
	mockBatch.Build()
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
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	e1 := mockEntryDetail()
	e1.TransactionCode = 32
	e1.Amount = 100
	e2 := mockEntryDetail()
	e2.TransactionCode = 37
	e2.Amount = 100
	mockBatch.AddEntryDetail(e1)
	mockBatch.AddEntryDetail(e2)
	mockBatch.Build()
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
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	mockBatch.AddEntryDetail(mockEntryDetail())
	mockBatch.AddEntryDetail(mockEntryDetail())
	mockBatch.Build()
	// works properly
	if err := mockBatch.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// create error
	mockBatch.Control.EntryHash = 0
	if err := mockBatch.Validate(); err != nil {
		if err != ErrValidEntryHash {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

// isOriginatorDNEMismatch
func TestBatchIsOriginatorDNEMismatch(t *testing.T) {
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	mockBatch.AddEntryDetail(mockEntryDetail())
	mockBatch.Build()
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
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	mockBatch.AddEntryDetail(mockEntryDetail())
	mockBatch.Build()
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
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	ed := mockEntryDetail()
	ed.AddAddenda(mockAddenda())
	mockBatch.AddEntryDetail(ed)
	mockBatch.Build()

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
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	ed := mockEntryDetail()
	ed.AddAddenda(mockAddenda())
	mockBatch.AddEntryDetail(ed)
	mockBatch.Build()
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
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	ed := mockEntryDetail()
	ed.AddAddenda(mockAddenda())
	ed.AddAddenda(mockAddenda())
	mockBatch.AddEntryDetail(ed)
	mockBatch.Build()

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
	mockBatch := NewBatch()
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
	entry.SetRDFI(81086674)                               // scottrade bank routing number
	entry.DFIAccountNumber = "62292250"                   // scottrade account number
	entry.Amount = 1000000                                // 1k dollars
	entry.IndividualIdentificationNumber = "658-888-2468" // Unique ID for payment
	entry.IndividualName = "Wade Arnold"
	entry.setTraceNumber(header.ODFIIdentification, 1)
	a1 := NewAddenda()
	a2 := NewAddenda()
	entry.AddAddenda(a1)
	entry.AddAddenda(a2)
	mockBatch.AddEntryDetail(entry)
	if err := mockBatch.Build(); err != nil {
		t.Errorf("Unexpected Batch.Build error: %v", err.Error())
	}
	//fmt.Printf("Batch: %+v \n", mockBatch)

	if err := mockBatch.ValidateAll(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
}

func TestBatchValidateAllBH(t *testing.T) {
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	ed := mockEntryDetail()
	ed.AddAddenda(mockAddenda())
	ed.AddAddenda(mockAddenda())
	mockBatch.AddEntryDetail(ed)
	mockBatch.Build()

	// Make it fail
	mockBatch.Header.ODFIIdentification = 0
	if err := mockBatch.ValidateAll(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBatchValidateAllED(t *testing.T) {
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	ed := mockEntryDetail()
	ed.AddAddenda(mockAddenda())
	ed.AddAddenda(mockAddenda())
	mockBatch.AddEntryDetail(ed)
	mockBatch.Build()

	// Make it fail
	mockBatch.Entries[0].DFIAccountNumber = ""
	if err := mockBatch.ValidateAll(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBatchValidateAllAddenda(t *testing.T) {
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	ed := mockEntryDetail()
	ed.AddAddenda(mockAddenda())
	ed.AddAddenda(mockAddenda())
	mockBatch.AddEntryDetail(ed)
	mockBatch.Build()

	// Make it fail
	mockBatch.Entries[0].Addendums[0].TypeCode = ""
	if err := mockBatch.ValidateAll(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBatchValidateAllBatchControl(t *testing.T) {
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	ed := mockEntryDetail()
	ed.AddAddenda(mockAddenda())
	ed.AddAddenda(mockAddenda())
	mockBatch.AddEntryDetail(ed)
	mockBatch.Build()

	// Make it fail
	mockBatch.Control.ODFIIdentification = 0
	if err := mockBatch.ValidateAll(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.ValidationAll error: %v", err.Error())
		}
	}
}

func TestBatchBuildHeader(t *testing.T) {
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	ed := mockEntryDetail()
	ed.AddAddenda(mockAddenda())
	ed.AddAddenda(mockAddenda())
	mockBatch.AddEntryDetail(ed)

	mockBatch.Header.ODFIIdentification = 0
	if err := mockBatch.Build(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Build error: %v", err.Error())
		}
	}
}

func TestBatchBuildNoEntries(t *testing.T) {
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	if err := mockBatch.Build(); err != nil {
		if !strings.Contains(err.Error(), ErrBatchEntries.Error()) {
			t.Errorf("Unexpected Batch.Build error: %v", err.Error())
		}
	}
}

func TestBatchDNEMismatch(t *testing.T) {
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	ed := mockEntryDetail()
	ed.AddAddenda(mockAddenda())
	ed.AddAddenda(mockAddenda())
	mockBatch.AddEntryDetail(ed)
	mockBatch.Build()

	mockBatch.Header.OriginatorStatusCode = 1
	mockBatch.Entries[0].TransactionCode = 23
	if err := mockBatch.Validate(); err != nil {
		if !strings.Contains(err.Error(), ErrBatchOriginatorDNE.Error()) {
			t.Errorf("Unexpected Batch.validation error: %v", err.Error())
		}
	}
}

func TestBatchAddendaSeq(t *testing.T) {
	mockBatch := NewBatch().SetHeader(mockBatchHeader())
	ed := mockEntryDetail()
	ed.AddAddenda(mockAddenda())
	ed.AddAddenda(mockAddenda())
	mockBatch.AddEntryDetail(ed)
	mockBatch.Build()

	mockBatch.Entries[0].Addendums[0].SequenceNumber = 2
	mockBatch.Entries[0].Addendums[1].SequenceNumber = 1
	if err := mockBatch.Validate(); err != nil {
		if !strings.Contains(err.Error(), ErrBatchAddendaSequence.Error()) {
			t.Errorf("Unexpected Batch.validation error: %v", err.Error())
		}
	}
}
