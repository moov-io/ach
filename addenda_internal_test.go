// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
)

func mockAddenda() *Addenda {
	a, _ := NewAddenda()
	addenda := a.(*Addenda)
	addenda.EntryDetailSequenceNumber = 1234567
	return addenda
}

func TestMockAddenda(t *testing.T) {
	addenda := mockAddenda()
	if err := addenda.Validate(); err != nil {
		t.Error("mockAddenda does not validate and will break other tests")
	}
	if addenda.EntryDetailSequenceNumber != 1234567 {
		t.Error("EntryDetailSequenceNumber dependent default value has changed")
	}
}

func TestParseAddenda(t *testing.T) {
	var line = "705WEB                                        DIEGO MAY                            00010000001"

	r := NewReader(strings.NewReader(line))
	r.addCurrentBatch(NewBatchPPD())
	r.currentBatch.GetHeader().StandardEntryClassCode = "PPD"
	r.currentBatch.AddEntry(&EntryDetail{TransactionCode: 22, AddendaRecordIndicator: 1})
	r.line = line
	err := r.parseAddenda()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentBatch.GetEntries()[0].Addendum[0].(*Addenda)

	if record.recordType != "7" {
		t.Errorf("RecordType Expected '7' got: %v", record.recordType)
	}
	if record.TypeCode() != "05" {
		t.Errorf("TypeCode Expected 10 got: %v", record.TypeCode())
	}
	if record.PaymentRelatedInformationField() != "WEB                                        DIEGO MAY                            " {
		t.Errorf("PaymentRelatedInformation Expected 'WEB                                        DIEGO MAY                            ' got: %v", record.PaymentRelatedInformationField())
	}
	if record.SequenceNumberField() != "0001" {
		t.Errorf("SequenceNumber Expected '0001' got: %v", record.SequenceNumberField())
	}
	if record.EntryDetailSequenceNumberField() != "0000001" {
		t.Errorf("EntryDetailSequenceNumber Expected '0000001' got: %v", record.EntryDetailSequenceNumberField())
	}
}

// TestAddendaString validats that a known parsed file can be return to a string of the same value
func TestAddendaString(t *testing.T) {
	var line = "705WEB                                        DIEGO MAY                            00010000001"
	r := NewReader(strings.NewReader(line))
	r.addCurrentBatch(NewBatchPPD())
	r.currentBatch.GetHeader().StandardEntryClassCode = "PPD"
	r.currentBatch.AddEntry(&EntryDetail{AddendaRecordIndicator: 1})
	r.line = line
	err := r.parseAddenda()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentBatch.GetEntries()[0].Addendum[0]
	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

func TestValidateAddendaRecordType(t *testing.T) {
	addenda := mockAddenda()
	addenda.recordType = "2"
	if err := addenda.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestValidateAddendaTypeCode(t *testing.T) {
	addenda := mockAddenda()
	addenda.typeCode = "23"
	if err := addenda.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestAddendaFieldInclusion(t *testing.T) {
	addenda := mockAddenda()
	addenda.EntryDetailSequenceNumber = 0
	if err := addenda.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EntryDetailSequenceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestAddendaFieldInclusionRecordType(t *testing.T) {
	addenda := mockAddenda()
	addenda.recordType = ""
	if err := addenda.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestAddendaPaymentRelatedInformationAlphaNumeric(t *testing.T) {
	addenda := mockAddenda()
	addenda.PaymentRelatedInformation = "®©"
	if err := addenda.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PaymentRelatedInformation" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestAddendaTyeCodeNil(t *testing.T) {
	addenda := mockAddenda()
	addenda.typeCode = ""
	if err := addenda.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}
