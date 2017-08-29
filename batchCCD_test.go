package ach

import "testing"

func mockBatchCCDHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 220
	bh.StandardEntryClassCode = "CCD"
	bh.CompanyName = "Your Company, inc"
	bh.CompanyIdentification = "123456789"
	bh.CompanyEntryDescription = "Vndr Pay"
	bh.ODFIIdentification = 6200001
	return bh
}

func mockCCDEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 27
	entry.SetRDFI(9101298)
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 5000000
	entry.IndividualName = "Wade Arnold"
	entry.TraceNumber = 123456789
	entry.DiscretionaryData = "S"
	return entry
}

func mockBatchCCD() *BatchCCD {
	mockBatch := NewBatchCCD()
	mockBatch.SetHeader(mockBatchCCDHeader())
	mockBatch.AddEntry(mockCCDEntryDetail())
	mockBatch.GetEntries()[0].AddAddenda(mockAddenda())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// A Batch CCD can only have one addendum per entry detail
func TestBatchCCDAddendumCount(t *testing.T) {
	mockBatch := mockBatchCCD()
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

// Individual name is a mandatory field
func TestBatchCCDReceivingCompanyName(t *testing.T) {
	mockBatch := mockBatchCCD()
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

func TestBatchCCDIdentificationNumber(t *testing.T) {
	mockBatch := mockBatchCCD()
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
func TestBatchCCDAddendaTypeCode(t *testing.T) {
	mockBatch := mockBatchCCD()
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

// verify that the standard entry class code is CCD for batchCCD
func TestBatchCCDSEC(t *testing.T) {
	mockBatch := mockBatchCCD()
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
