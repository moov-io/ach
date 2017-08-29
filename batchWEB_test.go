package ach

import "testing"

func mockBatchWEBHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 220
	bh.StandardEntryClassCode = "WEB"
	bh.CompanyName = "Your Company, inc"
	bh.CompanyIdentification = "123456789"
	bh.CompanyEntryDescription = "Online Order"
	bh.ODFIIdentification = 6200001
	return bh
}

func mockWEBEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 27
	entry.SetRDFI(9101298)
	entry.DFIAccountNumber = "123456789"
	entry.Amount = 100000000
	entry.IndividualName = "Wade Arnold"
	entry.TraceNumber = 123456789
	return entry
}

func mockBatchWEB() *BatchWEB {
	mockBatch := NewBatchWEB()
	mockBatch.SetHeader(mockBatchWEBHeader())
	mockBatch.AddEntry(mockWEBEntryDetail())
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// A Batch web can only have one addendum per entry detail
func TestBatchWEBAddendumCount(t *testing.T) {
	mockBatch := mockBatchWEB()
	// Adding a second addenda to the mock entry
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda())

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

func TestBatchWEBPaymentTypeSingleEntry(t *testing.T) {
	//mockBatch := mockBatchWEB()
}

func TestBatchWEBDebitOnly(t *testing.T) {
	mockBatch := mockBatchWEB()
	// make the entry a credit
	mockBatch.GetEntries()[0].TransactionCode = 22
	// re-create batch to rebuild control and create runs validate
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

// Nor more than 1 batch per entry detail record can exist
func TestBatchWebAddenda(t *testing.T) {
	mockBatch := mockBatchWEB()
	// mock batch already has one addenda. Creating two addenda should error
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda())
	if err := mockBatch.Create(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "AddendaCount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// Individual name is a mandatory field
func TestBatchWebIndividualNameRequired(t *testing.T) {
	mockBatch := mockBatchWEB()
	// mock batch already has one addenda. Creating two addenda should error
	mockBatch.GetEntries()[0].IndividualName = ""
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "IndividualName" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// verify addenda type code is 05
func TestAddendaTypeCode(t *testing.T) {
	mockBatch := mockBatchWEB()
	mockBatch.GetEntries()[0].Addendum[0].TypeCode = "07"
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
