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
	bh.ODFIIdentification = 6200001
	bh.BatchNumber = 1
	return bh
}

// TestParseBatchHeader parses a known Batch Header Record string.
func TestParseBatchHeader(t *testing.T) {
	var line = "5225companyname                         origid    PPDCHECKPAYMT000002080730   1076401250000001"
	r := NewReader(strings.NewReader(line))
	r.line = line
	err := r.parseBatchHeader()
	//fmt.Printf("holler %+v \n", r.currentBatch.Header)

	if err != nil {
		t.Errorf("unknown error: %v", err)
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
	err := r.parseBatchHeader()
	if err != nil {
		t.Errorf("unknown error: %v", err)
	}
	record := r.currentBatch.GetHeader()

	if record.String() != line {
		t.Errorf("Strings do not match")

		//fmt.Printf("%s \n %s", line, r.currentBatch.Header.String())
	}
}

// TestValidateBHRecordType ensure error if recordType is not 5
func TestValidateBHRecordType(t *testing.T) {
	bh := mockBatchHeader()
	bh.recordType = "2"
	if err := bh.Validate(); err != nil {
		if !strings.Contains(err.Error(), ErrRecordType.Error()) {
			t.Errorf("Expected RecordType Error got: %v", err)
		}
	}

}

// TestInvalidServiceCode ensure error if service class is not valid
func TestInvalidServiceCode(t *testing.T) {
	bh := mockBatchHeader()
	bh.ServiceClassCode = 123
	if err := bh.Validate(); err != nil {
		if !strings.Contains(err.Error(), ErrServiceClass.Error()) {
			t.Errorf("Expected Service Class Error got: %v", err)
		}
	}
}

// TestValidateInvalidServiceCode ensure error if service class is not valid
func TestInvalidSECCode(t *testing.T) {
	bh := mockBatchHeader()
	bh.StandardEntryClassCode = "ABC"
	if err := bh.Validate(); err != nil {
		if !strings.Contains(err.Error(), ErrSECCode.Error()) {
			t.Errorf("Expected SEC CodeError got: %v", err)
		}
	}
}

// TestInvalidOrigStatusCode ensure error if originator status code is not valid
func TestInvalidOrigStatusCode(t *testing.T) {
	bh := mockBatchHeader()
	bh.OriginatorStatusCode = 3
	if err := bh.Validate(); err != nil {
		if !strings.Contains(err.Error(), ErrOrigStatusCode.Error()) {
			t.Errorf("Expected Originator Status CodeError got: %v", err)
		}
	}
}

func TestBatchHeaderFieldInclusion(t *testing.T) {
	bh := mockBatchHeader()
	// works properly
	if err := bh.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// create error is mismatch
	bh.BatchNumber = 0
	if err := bh.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBatchHeaderCompanyNameAlphaNumeric(t *testing.T) {
	bh := mockBatchHeader()
	// works properly
	if err := bh.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// create error is mismatch
	bh.CompanyName = "AT&T"
	if err := bh.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBatchCompanyDiscretionaryDataAlphaNumeric(t *testing.T) {
	bh := mockBatchHeader()
	// works properly
	if err := bh.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// create error is mismatch
	bh.CompanyDiscretionaryData = "Invoice: #12345"
	if err := bh.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBatchCompanyIdentificationAlphaNumeric(t *testing.T) {
	bh := mockBatchHeader()
	// works properly
	if err := bh.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// create error is mismatch
	bh.CompanyIdentification = "EIN:12345"
	if err := bh.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBatchCompanyEntryDescriptionAlphaNumeric(t *testing.T) {
	bh := mockBatchHeader()
	// works properly
	if err := bh.Validate(); err != nil {
		t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
	}
	// create error is mismatch
	bh.CompanyEntryDescription = "P@YROLL"
	if err := bh.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBHFieldInclusionRecordType(t *testing.T) {
	bh := mockBatchHeader()
	bh.recordType = ""
	if err := bh.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBHFieldInclusionCompanyName(t *testing.T) {
	bh := mockBatchHeader()
	bh.CompanyName = ""
	if err := bh.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBHFieldInclusionCompanyIdentification(t *testing.T) {
	bh := mockBatchHeader()
	bh.CompanyIdentification = ""
	if err := bh.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBHFieldInclusionStandardEntryClassCode(t *testing.T) {
	bh := mockBatchHeader()
	bh.StandardEntryClassCode = ""
	if err := bh.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBHFieldInclusionCompanyEntryDescription(t *testing.T) {
	bh := mockBatchHeader()
	bh.CompanyEntryDescription = ""
	if err := bh.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}

func TestBHFieldInclusionOriginatorStatusCode(t *testing.T) {
	bh := mockBatchHeader()
	bh.OriginatorStatusCode = 0
	if err := bh.Validate(); err != nil {
		_, ok := err.(*ValidateError)
		if !ok {
			t.Errorf("Unexpected Batch.Validation error: %v", err.Error())
		}
	}
}
