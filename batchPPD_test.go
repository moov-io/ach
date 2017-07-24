// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"
	"time"
)

func mockBatchPPD() *BatchPPD {
	mockBatch := NewBatchPPD()
	mockBatch.SetHeader(mockBatchHeader())
	mockBatch.AddEntry(mockEntryDetail())
	if err := mockBatch.Build(); err != nil {
		panic(err)
	}
	return mockBatch
}

func TestBatchError(t *testing.T) {
	err := &BatchError{BatchNumber: 1, FieldName: "mock", Msg: "test message"}
	if err.Error() != "BatchNumber 1 mock test message" {
		t.Error("BatchError Error has changed formatting")
	}
}
func TestBatchServiceClassCodeEsquality(t *testing.T) {
	mockBatch := mockBatchPPD()
	mockBatch.GetControl().ServiceClassCode = 225
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Error(err)
			}
		}
	}
}

func TestBatchCompanyIdentification(t *testing.T) {
	mockBatch := mockBatchPPD()
	mockBatch.GetControl().CompanyIdentification = "XYZ Inc"
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "CompanyIdentification" {
				t.Error(err)
			}
		}
	}
}

func TestBatchODFIIDMismatch(t *testing.T) {
	mockBatch := mockBatchPPD()
	mockBatch.GetControl().ODFIIdentification = 987654321
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ODFIIdentification" {
				t.Error(err)
			}
		}
	}
}

func TestBatchNumberMismatch(t *testing.T) {
	mockBatch := mockBatchPPD()
	mockBatch.GetControl().BatchNumber = 2
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "BatchNumber" {
				t.Error(err)
			}
		}
	}
}

// isBatchAmount
func TestCreditBatchisBatchAmount(t *testing.T) {
	mockBatch := NewBatchPPD()
	mockBatch.SetHeader(mockBatchHeader())
	e1 := mockEntryDetail()
	e1.TransactionCode = 22
	e1.Amount = 100
	e2 := mockEntryDetail()
	e2.TransactionCode = 22
	e2.Amount = 100
	mockBatch.AddEntry(e1)
	mockBatch.AddEntry(e2)
	// works properly
	if err := mockBatch.Build(); err != nil {
		t.Error(err)
	}
	// create error
	mockBatch.GetControl().TotalCreditEntryDollarAmount = 1
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TotalCreditEntryDollarAmount" {
				t.Error(err)
			}
		}
	}
}

func TestSavingsBatchisBatchAmount(t *testing.T) {
	mockBatch := NewBatchPPD()
	mockBatch.SetHeader(mockBatchHeader())
	e1 := mockEntryDetail()
	e1.TransactionCode = 32
	e1.Amount = 100
	e2 := mockEntryDetail()
	e2.TransactionCode = 37
	e2.Amount = 100
	mockBatch.AddEntry(e1)
	mockBatch.AddEntry(e2)
	// works properly
	if err := mockBatch.Build(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// create error
	mockBatch.GetControl().TotalDebitEntryDollarAmount = 1
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TotalDebitEntryDollarAmount" {
				t.Error(err)
			}
		}
	}
}

func TestBatchisEntryHash(t *testing.T) {
	mockBatch := mockBatchPPD()
	mockBatch.GetControl().EntryHash = 1
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "EntryHash" {
				t.Error(err)
			}
		}
	}
}

// isOriginatorDNE
func TestBatchisOriginatorDNE(t *testing.T) {
	mockBatch := mockBatchPPD()
	// Make it fail
	mockBatch.GetHeader().OriginatorStatusCode = 1

	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "EntryHash" {
				t.Error(err)
			}
		}
	}
}

func TestBatchDNEMismatch(t *testing.T) {
	mockBatch := NewBatchPPD()
	mockBatch.SetHeader(mockBatchHeader())
	ed := mockEntryDetail()
	ed.AddAddenda(mockAddenda())
	ed.AddAddenda(mockAddenda())
	mockBatch.AddEntry(ed)
	mockBatch.Build()

	mockBatch.GetHeader().OriginatorStatusCode = 1
	mockBatch.GetEntries()[0].TransactionCode = 23
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "OriginatorStatusCode" {
				t.Error(err)
			}
		}
	}
}

// ErrBatchTraceNumberNotODFI
func TestBatchTraceNumberNotODFI(t *testing.T) {
	mockBatch := mockBatchPPD()
	mockBatch.GetEntries()[0].setTraceNumber(12345678, 1)
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ODFIIdentificationField" {
				t.Error(err)
			}
		}
	}
}

func TestBatchEntryCountEquality(t *testing.T) {
	mockBatch := NewBatchPPD()
	mockBatch.SetHeader(mockBatchHeader())
	e := mockEntryDetail()
	a := mockAddenda()
	e.AddAddenda(a)
	mockBatch.AddEntry(e)

	// Build a valid batch
	if err := mockBatch.Build(); err != nil {
		t.Error(err)
	}
	// create error batch error
	mockBatch.GetControl().EntryAddendaCount = 1
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "EntryAddendaCount" {
				t.Error(err)
			}
		}
	}
}

// ErrBatchAddendaIndicator
func TestBatchAddendaIndicator(t *testing.T) {
	mockBatch := mockBatchPPD()
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda())
	mockBatch.GetEntries()[0].AddendaRecordIndicator = 0
	mockBatch.GetControl().EntryAddendaCount = 2
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "AddendaRecordIndicator" {
				t.Error(err)
			}
		}
	}
}

func TestBatchIsAddendaSeqAscending(t *testing.T) {
	mockBatch := NewBatchPPD()
	mockBatch.SetHeader(mockBatchHeader())
	ed := mockEntryDetail()
	ed.AddAddenda(mockAddenda())
	ed.AddAddenda(mockAddenda())
	mockBatch.AddEntry(ed)
	mockBatch.Build()

	mockBatch.GetEntries()[0].Addendums[0].SequenceNumber = 2
	mockBatch.GetEntries()[0].Addendums[1].SequenceNumber = 1
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "SequenceNumber" {
				t.Error(err)
			}
		}
	}
}

func TestBatchIsSequenceAscending(t *testing.T) {
	mockBatch := mockBatchPPD()
	e3 := mockEntryDetail()
	e3.TraceNumber = 0
	mockBatch.AddEntry(e3)
	mockBatch.GetControl().EntryAddendaCount = 2
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TraceNumber" {
				t.Error(err)
			}
		}
	}
}

// ErrBatchAddendaTraceNumber
func TestBatchAddendaTraceNumber(t *testing.T) {
	mockBatch := mockBatchPPD()
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda())
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda())
	// works properly
	if err := mockBatch.Build(); err != nil {
		t.Error(err)
	}
	// Make it fail
	mockBatch.GetEntries()[0].Addendums[0].EntryDetailSequenceNumber = 99
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TraceNumber" {
				t.Error(err)
			}
		}
	}
}

/** the previous tests where for validation the following tests are for building a batch **/

func TestBatchBuild(t *testing.T) {
	mockBatch := NewBatchPPD()
	header := NewBatchHeader()
	header.ServiceClassCode = 200
	header.CompanyName = "MY BEST COMP."
	header.CompanyDiscretionaryData = "INCLUDES OVERTIME"
	header.CompanyIdentification = "1419871234"
	header.StandardEntryClassCode = "PPD"
	header.CompanyEntryDescription = "PAYROLL"
	header.EffectiveEntryDate = time.Now()
	header.ODFIIdentification = 109991234
	mockBatch.SetHeader(header)

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
	mockBatch.AddEntry(entry)
	if err := mockBatch.Build(); err != nil {
		t.Error(err)
	}
}

func TestBatchValidateAllEntries(t *testing.T) {
	mockBatch := mockBatchPPD()

	// Make it fail
	mockBatch.GetEntries()[0].DFIAccountNumber = ""
	if err := mockBatch.ValidateAll(); err != nil {
		switch e := err.(type) {
		case *BatchError:
			t.Error(err)
		case *FieldError:
			if e.Msg != msgFieldInclusion {
				t.Error(err)
			}
		}
	}
}

func TestBatchValidateAllAddenda(t *testing.T) {
	mockBatch := mockBatchPPD()
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda())
	// works properly
	if err := mockBatch.Build(); err != nil {
		t.Error(err)
	}

	// Make it fail
	mockBatch.GetEntries()[0].Addendums[0].TypeCode = ""
	if err := mockBatch.ValidateAll(); err != nil {
		switch e := err.(type) {
		case *BatchError:
			t.Error(err)
		case *FieldError:
			if e.Msg != msgFieldInclusion {
				t.Error(err)
			}
		}
	}
}

func TestBatchValidateAllBatchControl(t *testing.T) {
	mockBatch := mockBatchPPD()

	// Make it fail
	mockBatch.GetControl().ODFIIdentification = 0
	if err := mockBatch.ValidateAll(); err != nil {
		switch e := err.(type) {
		case *BatchError:
			t.Error(err)
		case *FieldError:
			if e.Msg != msgFieldInclusion {
				t.Error(err)
			}
		}
	}
}

func TestBatchValidateAllHeader(t *testing.T) {
	mockBatch := mockBatchPPD()

	// Make it fail
	mockBatch.GetHeader().ODFIIdentification = 0
	if err := mockBatch.ValidateAll(); err != nil {
		switch e := err.(type) {
		case *BatchError:
			t.Error(err)
		case *FieldError:
			if e.Msg != msgFieldInclusion {
				t.Error(err)
			}
		}
	}
}

func TestBatchValidateAllBatch(t *testing.T) {
	mockBatch := mockBatchPPD()

	// Make it fail
	mockBatch.GetHeader().ODFIIdentification = 123456
	if err := mockBatch.ValidateAll(); err != nil {
		switch e := err.(type) {
		case *BatchError:
			if e.FieldName != "ODFIIdentification" {
				t.Error(err)
			}
		case *FieldError:
			t.Error(err)
		}
	}
}

func TestBatchBuildNoEntries(t *testing.T) {
	mockBatch := NewBatchPPD()
	mockBatch.SetHeader(mockBatchHeader())
	if err := mockBatch.Build(); err != nil {
		switch e := err.(type) {
		case *BatchError:
			if e.Msg != msgBatchEntries {
				t.Error(err)
			}
		case *FieldError:
			t.Error(err)
		}
	}
}

func TestBatchBuildHeader(t *testing.T) {
	mockBatch := mockBatchPPD()
	mockBatch.GetHeader().CompanyIdentification = ""
	if err := mockBatch.Build(); err != nil {
		switch e := err.(type) {
		case *BatchError:
			t.Error(err)
		case *FieldError:
			if e.Msg != msgFieldInclusion {
				t.Error(err)
			}
		}
	}
}
