// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"
	"time"
)

func mockBatchPPDHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 220
	bh.StandardEntryClassCode = "PPD"
	bh.CompanyName = "ACME Corporation"
	bh.CompanyIdentification = "123456789"
	bh.CompanyEntryDescription = "PAYROLL"
	bh.EffectiveEntryDate = time.Now()
	bh.ODFIIdentification = 6200001
	return bh
}

func mockPPDEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 22
	entry.SetRDFI(9101298)
	entry.DFIAccountNumber = "123456789"
	entry.Amount = 100000000
	entry.IndividualName = "Wade Arnold"
	entry.SetTraceNumber(mockBatchPPDHeader().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

func mockBatchPPDHeader2() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 200
	bh.CompanyName = "MY BEST COMP."
	bh.CompanyDiscretionaryData = "INCLUDES OVERTIME"
	bh.CompanyIdentification = "1419871234"
	bh.StandardEntryClassCode = "PPD"
	bh.CompanyEntryDescription = "PAYROLL"
	bh.EffectiveEntryDate = time.Now()
	bh.ODFIIdentification = 109991234
	return bh
}

func mockPPDEntry2() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = 22                  // ACH Credit
	entry.SetRDFI(81086674)                     // scottrade bank routing number
	entry.DFIAccountNumber = "62292250"         // scottrade account number
	entry.Amount = 1000000                      // 1k dollars
	entry.IdentificationNumber = "658-888-2468" // Unique ID for payment
	entry.IndividualName = "Wade Arnold"
	entry.SetTraceNumber(mockBatchPPDHeader2().ODFIIdentification, 1)
	entry.Category = CategoryForward
	return entry
}

func mockBatchPPD() *BatchPPD {
	mockBatch := NewBatchPPD(mockBatchPPDHeader())
	mockBatch.AddEntry(mockPPDEntryDetail())
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
	a.typeCode = "63"
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

func TestBatchBuild(t *testing.T) {
	mockBatch := NewBatchPPD(mockBatchPPDHeader2())
	entry := mockPPDEntry2()
	addenda05 := NewAddenda05()
	entry.AddAddenda(addenda05)
	mockBatch.AddEntry(entry)
	if err := mockBatch.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}
