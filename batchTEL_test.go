package ach

import (
	"testing"
)

func mockBatchTELHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 225
	bh.StandardEntryClassCode = "TEL"
	bh.CompanyName = "Your Company, inc"
	bh.CompanyIdentification = "123456789"
	bh.CompanyEntryDescription = "Vndr Pay"
	bh.ODFIIdentification = 6200001
	return bh
}

func mockTELEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 27
	entry.SetRDFI(9101298)
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 5000000
	entry.IdentificationNumber = "Phone 333-2222"
	entry.IndividualName = "Wade Arnold"
	entry.setTraceNumber(6200001, 123)
	entry.SetPaymentType("S")
	return entry
}

func mockBatchTEL() *BatchTEL {
	mockBatch := NewBatchTEL(mockBatchTELHeader())
	mockBatch.AddEntry(mockTELEntryDetail())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

func TestBatchTELHeader(t *testing.T) {
	batch, _ := NewBatch(mockBatchTELHeader())
	err, ok := batch.(*BatchTEL)
	if !ok {
		t.Errorf("Expecting BatchTEL got %T", err)
	}
}

func TestBatchTELCreate(t *testing.T) {
	mockBatch := mockBatchTEL()
	// Batch Header information is required to Create a batch.
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

func TestBatchTELAddendaCount(t *testing.T) {
	mockBatch := mockBatchTEL()
	// TEL can not have an addendum
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda())
	mockBatch.Create()
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "AddendaCount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchTELSEC(t *testing.T) {
	mockBatch := mockBatchTEL()
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

func TestBatchTELDebit(t *testing.T) {
	mockBatch := mockBatchTEL()
	mockBatch.GetEntries()[0].TransactionCode = 22
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "TransactionCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// verify that the entry detail payment type / discretionary data is either single or reoccurring for the
func TestBatchTELPaymentType(t *testing.T) {
	mockBatch := mockBatchTEL()
	mockBatch.GetEntries()[0].DiscretionaryData = "AA"
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			println(e.Error())
			if e.FieldName != "PaymentType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}
