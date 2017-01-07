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
	err := r.parseBatchHeader()
	if err != nil {
		t.Errorf("unknown error: %v", err)
	}
	record := r.currentBatch.Header

	if record.recordType != "5" {
		t.Errorf("RecordType Expected '5' got: %v", record.recordType)
	}
	if record.ServiceClassCode != 225 {
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
	if record.EffectiveEntryDate() != "080730" {
		t.Errorf("EffectiveEntryDate Expected '080730' got: %v", record.EffectiveEntryDate())
	}
	if record.settlementDate != "   " {
		t.Errorf("SettlementDate Expected '   ' got: %v", record.settlementDate)
	}
	if record.OriginatorStatusCode != 1 {
		t.Errorf("OriginatorStatusCode Expected 1 got: %v", record.OriginatorStatusCode)
	}
	if record.ODFIIdentification() != "07640125" {
		t.Errorf("OdfiIdentification Expected '07640125' got: %v", record.ODFIIdentification())
	}
	if record.BatchNumber() != "0000001" {
		t.Errorf("BatchNumber Expected '0000001' got: %v", record.BatchNumber())
	}
}

// TestBHString validats that a known parsed file can be return to a string of the same value
func TestBHString(t *testing.T) {
	var line = "5225companyname                         origid    PPDCHECKPAYMT000002080730   1076401250000001"
	r := NewReader(strings.NewReader(line))
	r.record = line
	err := r.parseBatchHeader()
	if err != nil {
		t.Errorf("unknown error: %v", err)
	}
	record := r.currentBatch.Header

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestValidateBHRecordType ensure error if recordType is not 5
func TestValidateBHRecordType(t *testing.T) {
	bh := NewBatchHeader()
	bh.recordType = "2"
	_, err := bh.Validate()
	if !strings.Contains(err.Error(), ErrRecordType.Error()) {
		t.Errorf("Expected RecordType Error got: %v", err)
	}
}

// TestInvalidServiceCode ensure error if service class is not valid
func TestInvalidServiceCode(t *testing.T) {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 123
	_, err := bh.Validate()
	if !strings.Contains(err.Error(), ErrServiceClass.Error()) {
		t.Errorf("Expected Service Class Error got: %v", err)
	}
}

// TestValidateInvalidServiceCode ensure error if service class is not valid
func TestInvalidSECCode(t *testing.T) {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 200
	bh.StandardEntryClassCode = "ABC"
	_, err := bh.Validate()
	if !strings.Contains(err.Error(), ErrSECCode.Error()) {
		t.Errorf("Expected SEC CodeError got: %v", err)
	}
}

// TestInvalidOrigStatusCode ensure error if originator status code is not valid
func TestInvalidOrigStatusCode(t *testing.T) {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 200
	bh.StandardEntryClassCode = "PPD"
	bh.OriginatorStatusCode = 3
	_, err := bh.Validate()
	if !strings.Contains(err.Error(), ErrOrigStatusCode.Error()) {
		t.Errorf("Expected Originator Status CodeError got: %v", err)
	}
}
