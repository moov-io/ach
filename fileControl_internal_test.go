// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
)

// TestParseFileControl parses a known File Control Record string.
func TestParseFileControl(t *testing.T) {
	var line = "9000001000001000000010005320001000000010500000000000000                                       "
	r := NewReader(strings.NewReader(line))
	r.record = line
	record := r.parseFileControl()

	if record.RecordType != "9" {
		t.Errorf("RecordType Expected '9' got: %v", record.RecordType)
	}
	if record.BatchCount != "000001" {
		t.Errorf("BatchCount Expected '000001' got: %v", record.BatchCount)
	}
	if record.BlockCount != "000001" {
		t.Errorf("BlockCount Expected '000001' got: %v", record.BlockCount)
	}
	if record.EntryAddendaCount != "00000001" {
		t.Errorf("EntryAddendaCount Expected '00000001' got: %v", record.EntryAddendaCount)
	}
	if record.EntryHash != "0005320001" {
		t.Errorf("EntryHash Expected '0005320001' got: %v", record.EntryHash)
	}
	if record.TotalDebitEntryDollarAmountInFile != "000000010500" {
		t.Errorf("TotalDebitEntryDollarAmountInFile Expected '0005000000010500' got: %v", record.TotalDebitEntryDollarAmountInFile)
	}
	if record.TotalCreditEntryDollarAmountInFile != "000000000000" {
		t.Errorf("TotalCreditEntryDollarAmountInFile Expected '000000000000' got: %v", record.TotalCreditEntryDollarAmountInFile)
	}
	if record.Reserved != "                                       " {
		t.Errorf("Reserved Expected '                                       ' got: %v", record.Reserved)
	}

}
