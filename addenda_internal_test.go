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
	r.record = line
	record := r.parseAddenda()

	if record.RecordType != "7" {
		t.Errorf("RecordType Expected '7' got: %v", record.RecordType)
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
