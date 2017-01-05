// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
)

// TestParseBatchControl parses a known Batch ControlRecord string.
func TestParseBatchControl(t *testing.T) {
	var line = "82250000010005320001000000010500000000000000origid                             076401250000001"
	r := NewReader(strings.NewReader(line))
	r.record = line
	record := r.parseBatchControl()

	if record.RecordType != "8" {
		t.Errorf("RecordType Expected '8' got: %v", record.RecordType)
	}
	if record.ServiceClassCode != "225" {
		t.Errorf("ServiceClassCode Expected '225' got: %v", record.ServiceClassCode)
	}
	if record.EntryAddendaCount != "000001" {
		t.Errorf("EntryAddendaCount Expected '000001' got: %v", record.EntryAddendaCount)
	}
	if record.EntryHash != "0005320001" {
		t.Errorf("EntryHash Expected '0005320001' got: %v", record.EntryHash)
	}
	if record.TotalDebitEntryDollarAmount != "000000010500" {
		t.Errorf("TotalDebitEntryDollarAmount Expected '000000010500' got: %v", record.TotalDebitEntryDollarAmount)
	}
	if record.TotalCreditEntryDollarAmount != "000000000000" {
		t.Errorf("TotalCreditEntryDollarAmount Expected '000000000000' got: %v", record.TotalCreditEntryDollarAmount)
	}
	if record.CompanyIdentification != "origid    " {
		t.Errorf("CompanyIdentification Expected 'origid    ' got: %v", record.CompanyIdentification)
	}
	if record.MessageAuthenticationCode != "                   " {
		t.Errorf("MessageAuthenticationCode Expected '                   ' got: %v", record.MessageAuthenticationCode)
	}
	if record.Reserved != "      " {
		t.Errorf("Reserved Expected '      ' got: %v", record.Reserved)
	}
	if record.OdfiIdentification != "07640125" {
		t.Errorf("OdfiIdentification Expected '07640125' got: %v", record.OdfiIdentification)
	}
	if record.BatchNumber != "0000001" {
		t.Errorf("BatchNumber Expected '0000001' got: %v", record.BatchNumber)
	}
}
