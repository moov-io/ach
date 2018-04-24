// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
)

func mockBatchHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = 220
	bh.StandardEntryClassCode = "PPD"
	bh.CompanyName = "ACME Corporation"
	bh.CompanyIdentification = "123456789"
	bh.CompanyEntryDescription = "PAYROLL"
	bh.ODFIIdentification = "6200001"
	return bh
}

func TestMockBatchHeader(t *testing.T) {
	bh := mockBatchHeader()
	if err := bh.Validate(); err != nil {
		t.Error("mockBatchHeader does not validate and will break other tests")
	}
	if bh.ServiceClassCode != 220 {
		t.Error("ServiceClassCode dependent default value has changed")
	}
	if bh.StandardEntryClassCode != "PPD" {
		t.Error("StandardEntryClassCode dependent default value has changed")
	}
	if bh.CompanyName != "ACME Corporation" {
		t.Error("CompanyName dependent default value has changed")
	}
	if bh.CompanyIdentification != "123456789" {
		t.Error("CompanyIdentification dependent default value has changed")
	}
	if bh.CompanyEntryDescription != "PAYROLL" {
		t.Error("CompanyEntryDescription dependent default value has changed")
	}
	if bh.ODFIIdentification != "6200001" {
		t.Error("ODFIIdentification dependent default value has changed")
	}
}

// TestParseBatchHeader parses a known Batch Header Record string.
func TestParseBatchHeader(t *testing.T) {
	var line = "5225companyname                         origid    PPDCHECKPAYMT000002080730   1076401250000001"
	r := NewReader(strings.NewReader(line))
	r.line = line
	if err := r.parseBatchHeader(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentBatch.GetHeader()

	if record.recordType != "5" {
		t.Errorf("RecordType Expected '5' got: %v", record.recordType)
	}
	if record.ServiceClassCode != 225 {
		t.Errorf("ServiceClassCode Expected '225' got: %v", record.ServiceClassCode)
	}
	if record.CompanyNameField() != "companyname     " {
		t.Errorf("CompanyName Expected 'companyname    ' got: '%v'", record.CompanyNameField())
	}
	if record.CompanyDiscretionaryDataField() != "                    " {
		t.Errorf("CompanyDiscretionaryData Expected '                    ' got: %v", record.CompanyDiscretionaryDataField())
	}
	if record.CompanyIdentificationField() != "origid    " {
		t.Errorf("CompanyIdentification Expected 'origid    ' got: %v", record.CompanyIdentificationField())
	}
	if record.StandardEntryClassCode != "PPD" {
		t.Errorf("StandardEntryClassCode Expected 'PPD' got: %v", record.StandardEntryClassCode)
	}
	if record.CompanyEntryDescriptionField() != "CHECKPAYMT" {
		t.Errorf("CompanyEntryDescription Expected 'CHECKPAYMT' got: %v", record.CompanyEntryDescriptionField())
	}
	if record.CompanyDescriptiveDate != "000002" {
		t.Errorf("CompanyDescriptiveDate Expected '000002' got: %v", record.CompanyDescriptiveDate)
	}
	if record.EffectiveEntryDateField() != "080730" {
		t.Errorf("EffectiveEntryDate Expected '080730' got: %v", record.EffectiveEntryDateField())
	}
	if record.settlementDate != "   " {
		t.Errorf("SettlementDate Expected '   ' got: %v", record.settlementDate)
	}
	if record.OriginatorStatusCode != 1 {
		t.Errorf("OriginatorStatusCode Expected 1 got: %v", record.OriginatorStatusCode)
	}
	if record.ODFIIdentificationField() != "07640125" {
		t.Errorf("OdfiIdentification Expected '07640125' got: %v", record.ODFIIdentificationField())
	}
	if record.BatchNumberField() != "0000001" {
		t.Errorf("BatchNumber Expected '0000001' got: %v", record.BatchNumberField())
	}
}

// TestBHString validats that a known parsed file can be return to a string of the same value
func TestBHString(t *testing.T) {
	var line = "5225companyname                         origid    PPDCHECKPAYMT000002080730   1076401250000001"
	r := NewReader(strings.NewReader(line))
	r.line = line
	if err := r.parseBatchHeader(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentBatch.GetHeader()

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestValidateBHRecordType ensure error if recordType is not 5
func TestValidateBHRecordType(t *testing.T) {
	bh := mockBatchHeader()
	bh.recordType = "2"
	if err := bh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestInvalidServiceCode ensure error if service class is not valid
func TestInvalidServiceCode(t *testing.T) {
	bh := mockBatchHeader()
	bh.ServiceClassCode = 123
	if err := bh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestValidateInvalidServiceCode ensure error if service class is not valid
func TestInvalidSECCode(t *testing.T) {
	bh := mockBatchHeader()
	bh.StandardEntryClassCode = "123"
	if err := bh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "StandardEntryClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestInvalidOrigStatusCode ensure error if originator status code is not valid
func TestInvalidOrigStatusCode(t *testing.T) {
	bh := mockBatchHeader()
	bh.OriginatorStatusCode = 3
	if err := bh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "OriginatorStatusCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestBatchHeaderFieldInclusion(t *testing.T) {
	bh := mockBatchHeader()
	bh.BatchNumber = 0
	if err := bh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "BatchNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestBatchHeaderCompanyNameAlphaNumeric(t *testing.T) {
	bh := mockBatchHeader()
	bh.CompanyName = "AT&T®"
	if err := bh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CompanyName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestBatchCompanyDiscretionaryDataAlphaNumeric(t *testing.T) {
	bh := mockBatchHeader()
	bh.CompanyDiscretionaryData = "®"
	if err := bh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CompanyDiscretionaryData" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestBatchCompanyIdentificationAlphaNumeric(t *testing.T) {
	bh := mockBatchHeader()
	bh.CompanyIdentification = "®"
	if err := bh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CompanyIdentification" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestBatchCompanyEntryDescriptionAlphaNumeric(t *testing.T) {
	bh := mockBatchHeader()
	bh.CompanyEntryDescription = "P®YROLL"
	if err := bh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CompanyEntryDescription" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestBHFieldInclusionRecordType(t *testing.T) {
	bh := mockBatchHeader()
	bh.recordType = ""
	if err := bh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestBHFieldInclusionCompanyName(t *testing.T) {
	bh := mockBatchHeader()
	bh.CompanyName = ""
	if err := bh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CompanyName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestBHFieldInclusionCompanyIdentification(t *testing.T) {
	bh := mockBatchHeader()
	bh.CompanyIdentification = ""
	if err := bh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CompanyIdentification" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestBHFieldInclusionStandardEntryClassCode(t *testing.T) {
	bh := mockBatchHeader()
	bh.StandardEntryClassCode = ""
	if err := bh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "StandardEntryClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestBHFieldInclusionCompanyEntryDescription(t *testing.T) {
	bh := mockBatchHeader()
	bh.CompanyEntryDescription = ""
	if err := bh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CompanyEntryDescription" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

func TestBHFieldInclusionOriginatorStatusCode(t *testing.T) {
	bh := mockBatchHeader()
	bh.OriginatorStatusCode = 0
	if err := bh.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "OriginatorStatusCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}
