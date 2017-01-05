// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
)

// TestParseBatchHeader parses a known Batch Header Record string.
func TestParseBatchHeader(t *testing.T) {
	var line = "5225companyname                         origid    PPDCHECKPAYMT000002080730   1076401250000001"
	r := NewReader(strings.NewReader(line))
	r.record = line
	record := r.parseBatchHeader()

	if record.RecordType != "5" {
		t.Errorf("RecordType Expected '5' got: %v", record.RecordType)
	}
	if record.ServiceClassCode != "225" {
		t.Errorf("ServiceClassCode Expected '225' got: %v", record.ServiceClassCode)
	}
	if record.CompanyName != "companyname     " {
		t.Errorf("CompanyName Expected 'companyname    ' got: '%v'", record.CompanyName)
	}
	if record.CompanyDiscretionaryData != "                    " {
		t.Errorf("CompanyDiscretionaryData Expected '                    ' got: %v", record.CompanyDiscretionaryData)
	}
	if record.CompanyIdentification != "origid    " {
		t.Errorf("CompanyIdentification Expected 'origid    ' got: %v", record.CompanyIdentification)
	}
	if record.StandardEntryClassCode != "PPD" {
		t.Errorf("StandardEntryClassCode Expected 'PPD' got: %v", record.StandardEntryClassCode)
	}
	if record.CompanyEntryDescription != "CHECKPAYMT" {
		t.Errorf("CompanyEntryDescription Expected 'CHECKPAYMT' got: %v", record.CompanyEntryDescription)
	}
	if record.CompanyDescriptiveDate != "000002" {
		t.Errorf("CompanyDescriptiveDate Expected '000002' got: %v", record.CompanyDescriptiveDate)
	}
	if record.EffectiveEntryDate != "080730" {
		t.Errorf("EffectiveEntryDate Expected '080730' got: %v", record.EffectiveEntryDate)
	}
	if record.SettlementDate != "   " {
		t.Errorf("SettlementDate Expected '   ' got: %v", record.SettlementDate)
	}
	if record.OriginatorStatusCode != "1" {
		t.Errorf("OriginatorStatusCode Expected 1 got: %v", record.OriginatorStatusCode)
	}
	if record.OdfiIdentification != "07640125" {
		t.Errorf("OdfiIdentification Expected '07640125' got: %v", record.OdfiIdentification)
	}
	if record.BatchNumber != "0000001" {
		t.Errorf("BatchNumber Expected '0000001' got: %v", record.BatchNumber)
	}
}
