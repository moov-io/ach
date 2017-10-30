package ach

import (
	"testing"
)

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
	entry.IdentificationNumber = "location #23"
	entry.SetReceivingCompany("Best Co. #23")
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

// receiving company / Individual name is a mandatory field
func TestBatchCCDReceivingCompanyName(t *testing.T) {
	mockBatch := mockBatchCCD()
	// modify the Individual name / receiving company to nothing
	mockBatch.GetEntries()[0].SetReceivingCompany("")
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

func TestBatchCCDAddendaCount(t *testing.T) {
	mockBatch := mockBatchCCD()
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

func TestBatchCCDCreate(t *testing.T) {
	mockBatch := mockBatchCCD()
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

func TestBatchCCDParam(t *testing.T) {

	batch, _ := NewBatch(BatchParam{
		ServiceClassCode:        "220",
		CompanyName:             "Your Company, inc",
		StandardEntryClass:      "CCD",
		CompanyIdentification:   "123456789",
		CompanyEntryDescription: "Vndr Pay",
		CompanyDescriptiveDate:  "Oct 23",
		ODFIIdentification:      "123456789"})

	_, ok := batch.(*BatchCCD)
	if !ok {
		t.Error("Expecting BachCCD")
	}

}
