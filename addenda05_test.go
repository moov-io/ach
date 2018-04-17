// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"
	"strings"
)

func mockAddenda05() *Addenda05 {
	addenda05 := NewAddenda05()
	addenda05.SequenceNumber =            1
	addenda05.EntryDetailSequenceNumber = 1234567

	return addenda05
}

func TestMockAddenda05(t *testing.T) {
	addenda05 := mockAddenda05()
	if err := addenda05.Validate(); err != nil {
		t.Error("mockAddenda05 does not validate and will break other tests")
	}
	if addenda05.EntryDetailSequenceNumber != 1234567 {
		t.Error("EntryDetailSequenceNumber dependent default value has changed")
	}
}

func TestParseAddenda05(t *testing.T) {
	addendaPPD := NewAddenda05()
	//var line = "705WEB                                        DIEGO MAY                            00010000001"
	var line = "705PPD                                        DIEGO MAY                            00010000001"
	addendaPPD.Parse(line)

	r := NewReader(strings.NewReader(line))

	//Add a new BatchWEB
	r.addCurrentBatch(NewBatchPPD(mockBatchPPDHeader()))

	//Add a WEB EntryDetail
	entryDetail := mockPPDEntryDetail()

	//Add an addenda to the WEB EntryDetail
	entryDetail.AddAddenda(addendaPPD)

	// add the WEB entry detail to the batch
	r.currentBatch.AddEntry(entryDetail)

	record := r.currentBatch.GetEntries()[0].Addendum[0].(*Addenda05)

	if record.recordType != "7" {
		t.Errorf("RecordType Expected '7' got: %v", record.recordType)
	}
	if record.TypeCode() != "05" {
		t.Errorf("TypeCode Expected 10 got: %v", record.TypeCode())
	}
	if record.PaymentRelatedInformationField() != "PPD                                        DIEGO MAY                            " {
		t.Errorf("PaymentRelatedInformation Expected 'PPD                                        DIEGO MAY                            ' got: %v", record.PaymentRelatedInformationField())
	}
	if record.SequenceNumberField() != "0001" {
		t.Errorf("SequenceNumber Expected '0001' got: %v", record.SequenceNumberField())
	}
	if record.EntryDetailSequenceNumberField() != "0000001" {
		t.Errorf("EntryDetailSequenceNumber Expected '0000001' got: %v", record.EntryDetailSequenceNumberField())
	}
}

// TestAddenda05 String validates that a known parsed file can be return to a string of the same value
func TestAddenda05String(t *testing.T) {
	addenda05 := NewAddenda05()
	var line = "705WEB                                        DIEGO MAY                            00010000001"
	addenda05.Parse(line)

	if addenda05.String() != line {
		t.Errorf("Strings do not match")
	}
}

func TestValidateAddenda05RecordType(t *testing.T) {
	addenda05 := mockAddenda05()
	addenda05.recordType = "63"
	if err := addenda05.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestValidateAddenda05TypeCode(t *testing.T) {
	addenda05 := mockAddenda05()
	addenda05.typeCode = "23"
	if err := addenda05.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestAddenda05FieldInclusion(t *testing.T) {
	addenda05 := mockAddenda05()
	addenda05.EntryDetailSequenceNumber = 0
	if err := addenda05.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EntryDetailSequenceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestAddenda05FieldInclusionRecordType(t *testing.T) {
	addenda05 := mockAddenda05()
	addenda05.recordType = ""
	if err := addenda05.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestAddenda05PaymentRelatedInformationAlphaNumeric(t *testing.T) {
	addenda05 := mockAddenda05()
	addenda05.PaymentRelatedInformation = "®©"
	if err := addenda05.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PaymentRelatedInformation" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestAddenda05TyeCodeNil(t *testing.T) {
	addenda05 := mockAddenda05()
	addenda05.typeCode = ""
	if err := addenda05.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

