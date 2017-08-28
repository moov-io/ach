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

func mockBatchWEB() *BatchWEB {
	mockBatch := NewBatchWEB()
	mockBatch.SetHeader(mockBatchWEBHeader())
	mockBatch.AddEntry(mockEntryDetail())
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

	//mockBatch := mockBatchWEB()
}
