// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
)

func mockAddenda() Addenda {
	addenda := NewAddenda()
	addenda.TypeCode = "08"
	addenda.SequenceNumber = 1
	addenda.EntryDetailSequenceNumber = 1234567
	return addenda
}

// TestParseAddendaRecord parses a known Addenda Record string.
func TestParseAddenda(t *testing.T) {
	var line = "710WEB                                        DIEGO MAY                            00010000001"

	r := NewReader(strings.NewReader(line))
	r.addCurrentBatch(NewBatchPPD())
	r.currentBatch.GetHeader().StandardEntryClassCode = "PPD"
	r.currentBatch.AddEntry(&EntryDetail{TransactionCode: 22, AddendaRecordIndicator: 1})
	r.line = line
	err := r.parseAddenda()
	if err != nil {
		t.Errorf("unknown error: %v", err)
	}
	record := r.currentBatch.GetEntries()[0].Addendums[0]

	if record.recordType != "7" {
		t.Errorf("RecordType Expected '7' got: %v", record.recordType)
	}
	if record.TypeCode != "10" {
		t.Errorf("TypeCode Expected 10 got: %v", record.TypeCode)
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
	var line = "710WEB                                        DIEGO MAY                            00010000001"
	r := NewReader(strings.NewReader(line))
	r.addCurrentBatch(NewBatchPPD())
	r.currentBatch.GetHeader().StandardEntryClassCode = "PPD"
	r.currentBatch.AddEntry(&EntryDetail{AddendaRecordIndicator: 1})
	r.line = line
	err := r.parseAddenda()
	if err != nil {
		t.Errorf("unknown error: %v", err)
	}
	record := r.currentBatch.GetEntries()[0].Addendums[0]
	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestValidateAddendaRecordType ensure error if recordType is not 7
func TestValidateAddendaRecordType(t *testing.T) {
	addenda := mockAddenda()
	addenda.recordType = "2"
	if err := addenda.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Error(err)
			}
		}
	}
}

// TestValidateAddendaRecordType ensure error if recordType is not 7
func TestValidateAddendaTypeCode(t *testing.T) {
	addenda := mockAddenda()
	addenda.TypeCode = "23"
	if err := addenda.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Error(err)
			}
		}
	}
}

func TestAddendaFieldInclusion(t *testing.T) {
	addenda := mockAddenda()
	// works properly
	if err := addenda.Validate(); err != nil {
		t.Errorf("Unexpected Addenda error: %v", err.Error())
	}
	// create error is mismatch
	addenda.EntryDetailSequenceNumber = 0
	if err := addenda.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "EntryDetailSequenceNumber" {
				t.Error(err)
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
				t.Error(err)
			}
		}
	}
}

func TestAddendaPaymentRelatedInformationAlphaNumeric(t *testing.T) {
	addenda := mockAddenda()
	// works properly
	if err := addenda.Validate(); err != nil {
		t.Error(err)
	}
	// create error is mismatch
	addenda.PaymentRelatedInformation = "@!"
	if err := addenda.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "PaymentRelatedInformation" {
				t.Error(err)
			}
		}
	}
}
