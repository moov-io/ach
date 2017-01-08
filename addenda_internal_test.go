// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
)

// TestParseAddendaRecord parses a known Addenda Record string.
func TestParseAddenda(t *testing.T) {
	var line = "710WEB                                        DIEGO MAY                                0000001"

	r := NewReader(strings.NewReader(line))
	r.currentBatch.Header.StandardEntryClassCode = "PPD"
	r.currentBatch.addEntryDetail(EntryDetail{AddendaRecordIndicator: 1})
	r.line = line
	err := r.parseAddenda()
	if err != nil {
		t.Errorf("unknown error: %v", err)
	}
	record := r.currentBatch.Entries[0].Addendums[0]

	if record.recordType != "7" {
		t.Errorf("RecordType Expected '7' got: %v", record.recordType)
	}
	if record.TypeCode != "10" {
		t.Errorf("TypeCode Expected 10 got: %v", record.TypeCode)
	}
	if record.PaymentRelatedInformation != "WEB                                        DIEGO MAY                            " {
		t.Errorf("PaymentRelatedInformation Expected 'WEB                                        DIEGO MAY                            ' got: %v", record.PaymentRelatedInformation)
	}
	if record.SequenceNumber != "    " {
		t.Errorf("SequenceNumber Expected '    ' got: %v", record.SequenceNumber)
	}
	if record.EntryDetailSequenceNumber != "0000001" {
		t.Errorf("EntryDetailSequenceNumber Expected '0000001' got: %v", record.EntryDetailSequenceNumber)
	}
}

// TestAddendaString validats that a known parsed file can be return to a string of the same value
func TestAddendaString(t *testing.T) {
	var line = "710WEB                                        DIEGO MAY                                0000001"
	r := NewReader(strings.NewReader(line))
	r.currentBatch.Header.StandardEntryClassCode = "PPD"
	r.currentBatch.addEntryDetail(EntryDetail{AddendaRecordIndicator: 1})
	r.line = line
	err := r.parseAddenda()
	if err != nil {
		t.Errorf("unknown error: %v", err)
	}
	record := r.currentBatch.Entries[0].Addendums[0]
	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestValidateAddendaRecordType ensure error if recordType is not 7
func TestValidateAddendaRecordType(t *testing.T) {
	addenda := NewAddenda()
	addenda.recordType = "2"
	_, err := addenda.Validate()
	if !strings.Contains(err.Error(), ErrRecordType.Error()) {
		t.Errorf("Expected RecordType Error got: %v", err)
	}
}

// TestValidateAddendaRecordType ensure error if recordType is not 7
func TestValidateAddendaTypeCode(t *testing.T) {
	addenda := NewAddenda()
	addenda.TypeCode = "23"
	_, err := addenda.Validate()
	if !strings.Contains(err.Error(), ErrAddendaTypeCode.Error()) {
		t.Errorf("Expected Type Code Error got: %v", err)
	}
}
