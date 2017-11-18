package ach

import (
	"testing"
)

// batch should never be used directly.
func mockBatch() *batch {
	mockBatch := &batch{}
	mockBatch.SetHeader(mockBatchHeader())
	mockBatch.AddEntry(mockEntryDetail())
	if err := mockBatch.build(); err != nil {
		panic(err)
	}
	return mockBatch
}

// Test cases that apply to all batch types

func TestBatchNumberMismatch(t *testing.T) {
	mockBatch := mockBatch()
	mockBatch.GetControl().BatchNumber = 2
	if err := mockBatch.verify(); err != nil {
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
	mockBatch := mockBatch()
	mockBatch.SetHeader(mockBatchHeader())
	e1 := mockEntryDetail()
	e1.TransactionCode = 22
	e1.Amount = 100
	e2 := mockEntryDetail()
	e2.TransactionCode = 22
	e2.Amount = 100
	mockBatch.AddEntry(e1)
	mockBatch.AddEntry(e2)
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	mockBatch.GetControl().TotalCreditEntryDollarAmount = 1
	if err := mockBatch.verify(); err != nil {
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
	mockBatch := mockBatch()
	mockBatch.SetHeader(mockBatchHeader())
	e1 := mockEntryDetail()
	e1.TransactionCode = 32
	e1.Amount = 100
	e2 := mockEntryDetail()
	e2.TransactionCode = 37
	e2.Amount = 100
	mockBatch.AddEntry(e1)
	mockBatch.AddEntry(e2)
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	mockBatch.GetControl().TotalDebitEntryDollarAmount = 1
	if err := mockBatch.verify(); err != nil {
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
	mockBatch := mockBatch()
	mockBatch.GetControl().EntryHash = 1
	if err := mockBatch.verify(); err != nil {
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
	mockBatch := mockBatch()
	mockBatch.SetHeader(mockBatchHeader())
	ed := mockEntryDetail()
	ed.AddAddenda(mockAddenda())
	ed.AddAddenda(mockAddenda())
	mockBatch.AddEntry(ed)
	mockBatch.build()

	mockBatch.GetHeader().OriginatorStatusCode = 1
	mockBatch.GetEntries()[0].TransactionCode = 23
	if err := mockBatch.verify(); err != nil {
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
	mockBatch := mockBatch()
	mockBatch.GetEntries()[0].setTraceNumber(12345678, 1)
	if err := mockBatch.verify(); err != nil {
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
	mockBatch := mockBatch()
	mockBatch.SetHeader(mockBatchHeader())
	e := mockEntryDetail()
	a := mockAddenda()
	e.AddAddenda(a)
	mockBatch.AddEntry(e)
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	mockBatch.GetControl().EntryAddendaCount = 1
	if err := mockBatch.verify(); err != nil {
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
	mockBatch := mockBatch()
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda())
	mockBatch.GetEntries()[0].AddendaRecordIndicator = 0
	mockBatch.GetControl().EntryAddendaCount = 2
	if err := mockBatch.verify(); err != nil {
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
	mockBatch := mockBatch()
	mockBatch.SetHeader(mockBatchHeader())
	ed := mockEntryDetail()
	ed.AddAddenda(mockAddenda())
	ed.AddAddenda(mockAddenda())
	mockBatch.AddEntry(ed)
	mockBatch.build()

	mockBatch.GetEntries()[1].Addendum[0].(*Addenda).SequenceNumber = 2
	mockBatch.GetEntries()[1].Addendum[1].(*Addenda).SequenceNumber = 1
	if err := mockBatch.verify(); err != nil {
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
	mockBatch := mockBatch()
	e3 := mockEntryDetail()
	e3.TraceNumber = 1
	mockBatch.AddEntry(e3)
	mockBatch.GetControl().EntryAddendaCount = 2
	if err := mockBatch.verify(); err != nil {
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
	mockBatch := mockBatch()
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda())
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	mockBatch.GetEntries()[0].Addendum[0].(*Addenda).EntryDetailSequenceNumber = 99
	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TraceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestNewBatchDefault(t *testing.T) {
	_, err := NewBatch(BatchParam{
		StandardEntryClass: "NIL"})

	if e, ok := err.(*FileError); ok {
		if e.FieldName != "StandardEntryClassCode" {
			t.Errorf("%T: %s", err, err)
		}
	} else {
		t.Errorf("%T: %s", err, err)
	}
}

func TestBatchCategory(t *testing.T) {
	mockBatch := mockBatch()
	// Add a Addenda Return to the mock batch
	mockBatch.GetEntries()[0].AddAddenda(mockAddendaReturn())
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	if mockBatch.Category() != CategoryReturn {
		t.Errorf("AddendaReturn added to batch and category is %s", mockBatch.Category())
	}
}

func TestBatchCategoryForwardReturn(t *testing.T) {
	mockBatch := mockBatch()
	// Add a Addenda Return to the mock batch
	entry := mockEntryDetail()
	entry.AddAddenda(mockAddendaReturn())
	mockBatch.AddEntry(entry)
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := mockBatch.verify(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Category" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}
