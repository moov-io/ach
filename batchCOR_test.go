package ach

import (
	"testing"
)

// TODO make all the mock values cor fields

func mockBatchCORHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 220
	bh.StandardEntryClassCode = "COR"
	bh.CompanyName = "Your Company, inc"
	bh.CompanyIdentification = "123456789"
	bh.CompanyEntryDescription = "Vndr Pay"
	bh.ODFIIdentification = "6200001"
	return bh
}

func mockCOREntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 27
	entry.SetRDFI(9101298)
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 0
	entry.IdentificationNumber = "location #23"
	entry.SetReceivingCompany("Best Co. #23")
	entry.SetTraceNumber(mockBatchCORHeader().ODFIIdentification, 1)
	entry.DiscretionaryData = "S"
	return entry
}

func mockBatchCOR() *BatchCOR {
	mockBatch := NewBatchCOR(mockBatchCORHeader())
	mockBatch.AddEntry(mockCOREntryDetail())
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda98())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// Test Batch Header
func TestBatchCORHeader(t *testing.T) {
	batch, _ := NewBatch(mockBatchCORHeader())

	_, ok := batch.(*BatchCOR)
	if !ok {
		t.Error("Expecting BachCOR")
	}
}

func TestBatchCORSEC(t *testing.T) {
	mockBatch := mockBatchCOR()
	mockBatch.header.StandardEntryClassCode = "RCK"
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "StandardEntryClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchCORAddendumCountTwo(t *testing.T) {
	mockBatch := mockBatchCOR()
	// Adding a second addenda to the mock entry
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda98())

	if err := mockBatch.Create(); err != nil {
		//	fmt.Printf("err: %v \n", err)
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addendum" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchCORAddendaCountZero(t *testing.T) {
	mockBatch := NewBatchCOR(mockBatchCORHeader())
	mockBatch.AddEntry(mockCOREntryDetail())
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addendum" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// check that Addendum is of type Addenda98
func TestBatchCORAddendaType(t *testing.T) {
	mockBatch := NewBatchCOR(mockBatchCORHeader())
	mockBatch.AddEntry(mockCOREntryDetail())
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda05())
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Addendum" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchCORAddendaTypeCode(t *testing.T) {
	mockBatch := mockBatchCOR()
	mockBatch.GetEntries()[0].Addendum[0].(*Addenda98).typeCode = "07"
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

func TestBatchCORAmount(t *testing.T) {
	mockBatch := mockBatchCOR()
	mockBatch.GetEntries()[0].Amount = 9999
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "Amount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchCORCreate(t *testing.T) {
	mockBatch := mockBatchCOR()
	// Must have valid batch header to create a batch
	mockBatch.GetHeader().ServiceClassCode = 63
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}
