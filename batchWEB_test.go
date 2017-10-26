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
	entry.TransactionCode = 22
	entry.SetRDFI(9101298)
	entry.DFIAccountNumber = "123456789"
	entry.Amount = 100000000
	entry.IndividualName = "Wade Arnold"
	entry.TraceNumber = 123456789
	entry.SetPaymentType("S")
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

// No more than 1 batch per entry detail record can exist
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
func TestBatchWEBAddendaTypeCode(t *testing.T) {
	mockBatch := mockBatchWEB()
	mockBatch.GetEntries()[0].Addendum[0].(*Addenda).typeCode = "07"
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

// verify that the standard entry class code is WEB for batchWeb
func TestBatchWebSEC(t *testing.T) {
	mockBatch := mockBatchWEB()
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

// verify that the entry detail payment type / discretionary data is either single or reoccurring for the
func TestBatchWebPaymentType(t *testing.T) {
	mockBatch := mockBatchWEB()
	mockBatch.GetEntries()[0].DiscretionaryData = "AA"
	if err := mockBatch.Validate(); err != nil {
		if e, ok := err.(*BatchError); ok {
			if e.FieldName != "PaymentType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestBatchWebCreate(t *testing.T) {
	mockBatch := mockBatchWEB()
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
