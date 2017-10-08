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
	if err := mockBatch.Create(); err != nil {
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

func TestBatchServiceClassCodeEquality(t *testing.T) {
	mockBatch := mockBatchPPD()
	mockBatch.GetControl().ServiceClassCode = 225
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchPPDCreate(t *testing.T) {
	mockBatch := mockBatchPPD()
	// can not have default values in Batch Header to build batch
	mockBatch.GetHeader().ServiceClassCode = 0
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchPPDTypeCode(t *testing.T) {
	mockBatch := mockBatchPPD()
	// change an addendum to an invalid type code
	a := mockAddenda()
	a.TypeCode = "63"
	mockBatch.GetEntries()[0].AddAddenda(a)
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchCompanyIdentification(t *testing.T) {
	mockBatch := mockBatchPPD()
	mockBatch.GetControl().CompanyIdentification = "XYZ Inc"
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "CompanyIdentification" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchODFIIDMismatch(t *testing.T) {
	mockBatch := mockBatchPPD()
	mockBatch.GetControl().ODFIIdentification = 987654321
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ODFIIdentification" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchNumberMismatch(t *testing.T) {
	mockBatch := mockBatchPPD()
	mockBatch.GetControl().BatchNumber = 2
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "BatchNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

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
	if err := mockBatch.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	mockBatch.GetControl().TotalCreditEntryDollarAmount = 1
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TotalCreditEntryDollarAmount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
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
	if err := mockBatch.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	mockBatch.GetControl().TotalDebitEntryDollarAmount = 1
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TotalDebitEntryDollarAmount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchisEntryHash(t *testing.T) {
	mockBatch := mockBatchPPD()
	mockBatch.GetControl().EntryHash = 1
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "EntryHash" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
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
	mockBatch.Create()

	mockBatch.GetHeader().OriginatorStatusCode = 1
	mockBatch.GetEntries()[0].TransactionCode = 23
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "OriginatorStatusCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchTraceNumberNotODFI(t *testing.T) {
	mockBatch := mockBatchPPD()
	mockBatch.GetEntries()[0].setTraceNumber(12345678, 1)
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "ODFIIdentificationField" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
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
	if err := mockBatch.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	mockBatch.GetControl().EntryAddendaCount = 1
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "EntryAddendaCount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchAddendaIndicator(t *testing.T) {
	mockBatch := mockBatchPPD()
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda())
	mockBatch.GetEntries()[0].AddendaRecordIndicator = 0
	mockBatch.GetControl().EntryAddendaCount = 2
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "AddendaRecordIndicator" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
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
	mockBatch.Create()

	mockBatch.GetEntries()[0].Addendum[0].SequenceNumber = 2
	mockBatch.GetEntries()[0].Addendum[1].SequenceNumber = 1
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "SequenceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchIsSequenceAscending(t *testing.T) {
	mockBatch := mockBatchPPD()
	e3 := mockEntryDetail()
	e3.TraceNumber = 1
	mockBatch.AddEntry(e3)
	mockBatch.GetControl().EntryAddendaCount = 2
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TraceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchAddendaTraceNumber(t *testing.T) {
	mockBatch := mockBatchPPD()
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda())
	if err := mockBatch.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	mockBatch.GetEntries()[0].Addendum[0].EntryDetailSequenceNumber = 99
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TraceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

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
	entry.TransactionCode = 22                  // ACH Credit
	entry.SetRDFI(81086674)                     // scottrade bank routing number
	entry.DFIAccountNumber = "62292250"         // scottrade account number
	entry.Amount = 1000000                      // 1k dollars
	entry.IdentificationNumber = "658-888-2468" // Unique ID for payment
	entry.IndividualName = "Wade Arnold"
	entry.setTraceNumber(header.ODFIIdentification, 1)
	a1 := NewAddenda()
	entry.AddAddenda(a1)
	mockBatch.AddEntry(entry)
	if err := mockBatch.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}
