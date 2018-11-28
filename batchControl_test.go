// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
)

func mockBatchControl() *BatchControl {
	bc := NewBatchControl()
	bc.ServiceClassCode = 220
	bc.CompanyIdentification = "121042882"
	bc.ODFIIdentification = "12104288"
	return bc
}

// testMockBatchControl tests mock batch control
func testMockBatchControl(t testing.TB) {
	bc := mockBatchControl()
	if err := bc.Validate(); err != nil {
		t.Error("mockBatchControl does not validate and will break other tests")
	}
	if bc.ServiceClassCode != 220 {
		t.Error("ServiceClassCode depedendent default value has changed")
	}
	if bc.CompanyIdentification != "121042882" {
		t.Error("CompanyIdentification depedendent default value has changed")
	}
	if bc.ODFIIdentification != "12104288" {
		t.Error("ODFIIdentification depedendent default value has changed")
	}
}

// TestMockBatchControl test mock batch control
func TestMockBatchControl(t *testing.T) {
	testMockBatchControl(t)
}

// BenchmarkMockBatchControl benchmarks mock batch control
func BenchmarkMockBatchControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockBatchControl(b)
	}
}

// TestParseBatchControl parses a known Batch ControlRecord string.
func testParseBatchControl(t testing.TB) {
	var line = "82250000010005320001000000010500000000000000origid                             076401250000001"
	r := NewReader(strings.NewReader(line))
	r.line = line
	bh := BatchHeader{BatchNumber: 1,
		ServiceClassCode:      225,
		CompanyIdentification: "origid",
		ODFIIdentification:    "7640125"}
	r.addCurrentBatch(NewBatchPPD(&bh))

	r.currentBatch.AddEntry(&EntryDetail{TransactionCode: 27, Amount: 10500, RDFIIdentification: "5320001", TraceNumber: "76401255655291"})
	if err := r.parseBatchControl(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentBatch.GetControl()

	if record.recordType != "8" {
		t.Errorf("RecordType Expected '8' got: %v", record.recordType)
	}
	if record.ServiceClassCode != 225 {
		t.Errorf("ServiceClassCode Expected '225' got: %v", record.ServiceClassCode)
	}
	if record.EntryAddendaCountField() != "000001" {
		t.Errorf("EntryAddendaCount Expected '000001' got: %v", record.EntryAddendaCountField())
	}
	if record.EntryHashField() != "0005320001" {
		t.Errorf("EntryHash Expected '0005320001' got: %v", record.EntryHashField())
	}
	if record.TotalDebitEntryDollarAmountField() != "000000010500" {
		t.Errorf("TotalDebitEntryDollarAmount Expected '000000010500' got: %v", record.TotalDebitEntryDollarAmountField())
	}
	if record.TotalCreditEntryDollarAmountField() != "000000000000" {
		t.Errorf("TotalCreditEntryDollarAmount Expected '000000000000' got: %v", record.TotalCreditEntryDollarAmountField())
	}
	if record.CompanyIdentificationField() != "origid    " {
		t.Errorf("CompanyIdentification Expected 'origid    ' got: %v", record.CompanyIdentificationField())
	}
	if record.MessageAuthenticationCodeField() != "                   " {
		t.Errorf("MessageAuthenticationCode Expected '                   ' got: %v", record.MessageAuthenticationCodeField())
	}
	if record.reserved != "      " {
		t.Errorf("Reserved Expected '      ' got: %v", record.reserved)
	}
	if record.ODFIIdentificationField() != "07640125" {
		t.Errorf("OdfiIdentification Expected '07640125' got: %v", record.ODFIIdentificationField())
	}
	if record.BatchNumberField() != "0000001" {
		t.Errorf("BatchNumber Expected '0000001' got: %v", record.BatchNumberField())
	}
}

// TestParseBatchControl tests parsing a known Batch ControlRecord string.
func TestParseBatchControl(t *testing.T) {
	testParseBatchControl(t)
}

// BenchmarkParseBatchControl benchmarks parsing a known Batch ControlRecord string.
func BenchmarkParseBatchControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testParseBatchControl(b)
	}
}

// testBCString validates that a known parsed file can be return to a string of the same value
func testBCString(t testing.TB) {
	var line = "82250000010005320001000000010500000000000000origid                             076401250000001"
	r := NewReader(strings.NewReader(line))
	r.line = line
	bh := BatchHeader{BatchNumber: 1,
		ServiceClassCode:      225,
		CompanyIdentification: "origid",
		ODFIIdentification:    "7640125"}
	r.addCurrentBatch(NewBatchPPD(&bh))

	r.currentBatch.AddEntry(&EntryDetail{TransactionCode: 27, Amount: 10500, RDFIIdentification: "5320001", TraceNumber: "76401255655291"})
	if err := r.parseBatchControl(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentBatch.GetControl()

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestBCString tests validating that a known parsed file can be return to a string of the same value
func TestBCString(t *testing.T) {
	testBCString(t)
}

// BenchmarkBCString benchmarks validating that a known parsed file can be return to a string of the same value
func BenchmarkBCString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBCString(b)
	}
}

// testValidateBCRecordType ensure error if recordType is not 8
func testValidateBCRecordType(t testing.TB) {
	bc := mockBatchControl()
	bc.recordType = "2"
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestValidateBCRecordType tests ensuring an error if recordType is not 8
func TestValidateBCRecordType(t *testing.T) {
	testValidateBCRecordType(t)
}

// BenchmarkValidateBCRecordType benchmarks ensuring an error if recordType is not 8
func BenchmarkValidateBCRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateBCRecordType(b)
	}
}

// testBCisServiceClassErr verifies service class code
func testBCisServiceClassErr(t testing.TB) {
	bc := mockBatchControl()
	bc.ServiceClassCode = 123
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ServiceClassCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestBCisServiceClassErr tests verifying service class code
func TestBCisServiceClassErr(t *testing.T) {
	testBCisServiceClassErr(t)
}

// BenchmarkBCisServiceClassErr benchmarks verifying service class code
func BenchmarkBCisServiceClassErr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testBCisServiceClassErr(b)
	}
}

// testBCBatchNumber verifies batch number
func testBCBatchNumber(t testing.TB) {
	bc := mockBatchControl()
	bc.BatchNumber = 0
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "BatchNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestBCBatchNumber tests verifying batch number
func TestBCBatchNumber(t *testing.T) {
	testBCBatchNumber(t)
}

// BenchmarkBCBatchNumber benchmarks verifying batch number
func BenchmarkBCBatchNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBCBatchNumber(b)
	}
}

// testBCCompanyIdentificationAlphaNumeric verifies Company Identification is AlphaNumeric
func testBCCompanyIdentificationAlphaNumeric(t testing.TB) {
	bc := mockBatchControl()
	bc.CompanyIdentification = "®"
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CompanyIdentification" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestBCCompanyIdentificationAlphaNumeric tests verifying Company Identification is AlphaNumeric
func TestBCCompanyIdentificationAlphaNumeric(t *testing.T) {
	testBCCompanyIdentificationAlphaNumeric(t)
}

// BenchmarkBCCompanyIdentificationAlphaNumeric benchmarks verifying Company Identification is AlphaNumeric
func BenchmarkBCCompanyIdentificationAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBCCompanyIdentificationAlphaNumeric(b)
	}
}

// testBCMessageAuthenticationCodeAlphaNumeric verifies AuthenticationCode is AlphaNumeric
func testBCMessageAuthenticationCodeAlphaNumeric(t testing.TB) {
	bc := mockBatchControl()
	bc.MessageAuthenticationCode = "®"
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "MessageAuthenticationCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestBCMessageAuthenticationCodeAlphaNumeric tests verifying AuthenticationCode is AlphaNumeric
func TestBCMessageAuthenticationCodeAlphaNumeric(t *testing.T) {
	testBCMessageAuthenticationCodeAlphaNumeric(t)
}

// BenchmarkBCMessageAuthenticationCodeAlphaNumeric benchmarks verifying AuthenticationCode is AlphaNumeric
func BenchmarkBCMessageAuthenticationCodeAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBCMessageAuthenticationCodeAlphaNumeric(b)
	}
}

// testBCFieldInclusionRecordType verifies Record Type is included
func testBCFieldInclusionRecordType(t testing.TB) {
	bc := mockBatchControl()
	bc.recordType = ""
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestBCFieldInclusionRecordType tests verifying Record Type is included
func TestBCFieldInclusionRecordType(t *testing.T) {
	testBCFieldInclusionRecordType(t)
}

// BenchmarkBCFieldInclusionRecordType benchmarks verifying Record Type is included
func BenchmarkBCFieldInclusionRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBCFieldInclusionRecordType(b)
	}
}

// testBCFieldInclusionServiceClassCode verifies Service Class Code is included
func testBCFieldInclusionServiceClassCode(t testing.TB) {
	bc := mockBatchControl()
	bc.ServiceClassCode = 0
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestBCFieldInclusionServiceClassCode tests verifying Service Class Code is included
func TestBCFieldInclusionServiceClassCode(t *testing.T) {
	testBCFieldInclusionServiceClassCode(t)
}

// BenchmarkBCFieldInclusionServiceClassCod benchmarks verifying Service Class Code is included
func BenchmarkBCFieldInclusionServiceClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBCFieldInclusionServiceClassCode(b)
	}
}

// testBCFieldInclusionODFIIdentification verifies batch control ODFIIdentification
func testBCFieldInclusionODFIIdentification(t testing.TB) {
	bc := mockBatchControl()
	bc.ODFIIdentification = "000000000"
	if err := bc.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestBCFieldInclusionODFIIdentification tests verifying batch control ODFIIdentification
func TestBCFieldInclusionODFIIdentification(t *testing.T) {
	testBCFieldInclusionODFIIdentification(t)
}

// BenchmarkBCFieldInclusionODFIIdentification benchmarks verifying batch control ODFIIdentification
func BenchmarkBCFieldInclusionODFIIdentification(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBCFieldInclusionODFIIdentification(b)
	}
}

// testBatchControlLength verifies batch control length
func testBatchControlLength(t testing.TB) {
	bc := NewBatchControl()
	recordLength := len(bc.String())
	if recordLength != 94 {
		t.Errorf("Instantiated length of Batch Control string is not 94 but %v", recordLength)
	}
}

// TestBatchControlLength tests verifying batch control length
func TestBatchControlLength(t *testing.T) {
	testBatchControlLength(t)
}

// BenchmarkBatchControlLength benchmarks verifying batch control length
func BenchmarkBatchControlLength(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchControlLength(b)
	}
}
