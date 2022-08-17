// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package ach

import (
	"strings"
	"testing"

	"github.com/moov-io/base"
)

// mockBatchHeader creates a batch header
func mockBatchHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.StandardEntryClassCode = PPD
	bh.CompanyName = "ACME Corporation"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "PAYROLL"
	bh.ODFIIdentification = "12104288"
	return bh
}

// testMockBatchHeader creates a batch header
func testMockBatchHeader(t testing.TB) {
	bh := mockBatchHeader()
	if err := bh.Validate(); err != nil {
		t.Error("mockBatchHeader does not validate and will break other tests")
	}
	if bh.ServiceClassCode != CreditsOnly {
		t.Error("ServiceClassCode dependent default value has changed")
	}
	if bh.StandardEntryClassCode != PPD {
		t.Error("StandardEntryClassCode dependent default value has changed")
	}
	if bh.CompanyName != "ACME Corporation" {
		t.Error("CompanyName dependent default value has changed")
	}
	if bh.CompanyIdentification != "121042882" {
		t.Error("CompanyIdentification dependent default value has changed")
	}
	if bh.CompanyEntryDescription != "PAYROLL" {
		t.Error("CompanyEntryDescription dependent default value has changed")
	}
	if bh.ODFIIdentification != "12104288" {
		t.Error("ODFIIdentification dependent default value has changed")
	}
}

// TestMockBatchHeader tests creating a batch header
func TestMockBatchHeader(t *testing.T) {
	testMockBatchHeader(t)
}

// BenchmarkMockBatchHeader benchmarks creating a batch header
func BenchmarkMockBatchHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockBatchHeader(b)
	}
}

// testParseBatchHeader parses a known batch header record string
func testParseBatchHeader(t testing.TB) {
	var line = "5225companyname                         origid    PPDCHECKPAYMT000002190730   1076401250000001"
	r := NewReader(strings.NewReader(line))
	r.line = line
	if err := r.parseBatchHeader(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentBatch.GetHeader()

	if record.ServiceClassCode != DebitsOnly {
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
	if record.StandardEntryClassCode != PPD {
		t.Errorf("StandardEntryClassCode Expected 'PPD' got: %v", record.StandardEntryClassCode)
	}
	if record.CompanyEntryDescriptionField() != "CHECKPAYMT" {
		t.Errorf("CompanyEntryDescription Expected 'CHECKPAYMT' got: %v", record.CompanyEntryDescriptionField())
	}
	if record.CompanyDescriptiveDate != "000002" {
		t.Errorf("CompanyDescriptiveDate Expected '000002' got: %v", record.CompanyDescriptiveDate)
	}
	if record.EffectiveEntryDateField() != "190730" {
		t.Errorf("EffectiveEntryDate Expected '190730' got: %v", record.EffectiveEntryDateField())
	}
	if record.SettlementDateField() != "   " {
		t.Errorf("SettlementDate Expected '   ' got: %v", record.SettlementDateField())
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

// TestParseBatchHeader tests parsing a known batch header record string
func TestParseBatchHeader(t *testing.T) {
	testParseBatchHeader(t)
}

// BenchmarkParseBatchHeader benchmarks parsing a known batch header record string
func BenchmarkParseBatchHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testParseBatchHeader(b)
	}
}

// testBHString validates that a known parsed file can be return to a string of the same value
func testBHString(t testing.TB) {
	var line = "5225companyname                         origid    PPDCHECKPAYMT000002180730   1076401250000001"
	r := NewReader(strings.NewReader(line))
	r.line = line
	if err := r.parseBatchHeader(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentBatch.GetHeader()

	if v := record.String(); v != line {
		t.Errorf("Strings do not match:\n   v=%q\nline=%q", v, line) // these are aligned with spaces
	}
}

// TestBHString tests validating that a known parsed file can be return to a string of the same value
func TestBHString(t *testing.T) {
	testBHString(t)
}

// BenchmarkBHString benchmarks validating that a known parsed file can be return to a string of the same value
func BenchmarkBHString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBHString(b)
	}
}

// testInvalidServiceCode validates error if service code is not valid
func testInvalidServiceCode(t testing.TB) {
	bh := mockBatchHeader()
	bh.ServiceClassCode = 123
	err := bh.Validate()
	if !base.Match(err, ErrServiceClass) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestInvalidServiceCode tests validating error if service code is not valid
func TestInvalidServiceCode(t *testing.T) {
	testInvalidServiceCode(t)
}

// BenchmarkInvalidServiceCode benchmarks validating error if service code is not valid
func BenchmarkInvalidServiceCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testInvalidServiceCode(b)
	}
}

// testValidateInvalidSECCode validates error if service class is not valid
func testInvalidSECCode(t testing.TB) {
	bh := mockBatchHeader()
	bh.StandardEntryClassCode = "123"
	err := bh.Validate()
	if !base.Match(err, ErrSECCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestInvalidSECCode tests validating error if service class is not valid
func TestInvalidSECCode(t *testing.T) {
	testInvalidSECCode(t)
}

// BenchmarkInvalidSECCode benchmarks validating error if service class is not valid
func BenchmarkInvalidSECCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testInvalidSECCode(b)
	}
}

// testInvalidOrigStatusCode validates error if originator status code is not valid
func testInvalidOrigStatusCode(t testing.TB) {
	bh := mockBatchHeader()
	bh.OriginatorStatusCode = 3
	err := bh.Validate()
	if !base.Match(err, ErrOrigStatusCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestInvalidOrigStatusCode tests validating error if originator status code is not valid
func TestInvalidOrigStatusCode(t *testing.T) {
	testInvalidOrigStatusCode(t)
}

// BenchmarkInvalidOrigStatusCode benchmarks  validating error if originator status code is not valid
func BenchmarkInvalidOrigStatusCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testInvalidOrigStatusCode(b)
	}
}

// testBatchHeaderFieldInclusion validates batch header field inclusion
func testBatchHeaderFieldInclusion(t testing.TB) {
	bh := mockBatchHeader()
	bh.BatchNumber = 0
	err := bh.Validate()
	// TODO: are we expecting there to be no errors here?
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchHeaderFieldInclusion tests validating batch header field inclusion
func TestBatchHeaderFieldInclusion(t *testing.T) {
	testBatchHeaderFieldInclusion(t)
}

// BenchmarkBatchHeaderFieldInclusion benchmarks validating batch header field inclusion
func BenchmarkBatchHeaderFieldInclusion(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchHeaderFieldInclusion(b)
	}
}

// testBatchHeaderCompanyNameAlphaNumeric validates batch header company name is alphanumeric
func testBatchHeaderCompanyNameAlphaNumeric(t testing.TB) {
	bh := mockBatchHeader()
	bh.CompanyName = "AT&T®"
	err := bh.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchHeaderCompanyNameAlphaNumeric tests validating batch header company name is alphanumeric
func TestBatchHeaderCompanyNameAlphaNumeric(t *testing.T) {
	testBatchHeaderCompanyNameAlphaNumeric(t)
}

// BenchmarkBatchHeaderCompanyNameAlphaNumeric benchmarks validating batch header company name is alphanumeric
func BenchmarkBatchHeaderCompanyNameAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchHeaderCompanyNameAlphaNumeric(b)
	}
}

// testBatchCompanyDiscretionaryDataAlphaNumeric validates company discretionary data is alphanumeric
func testBatchCompanyDiscretionaryDataAlphaNumeric(t testing.TB) {
	bh := mockBatchHeader()
	bh.CompanyDiscretionaryData = "®"
	err := bh.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCompanyDiscretionaryDataAlphaNumeric tests validating company discretionary data is alphanumeric
func TestBatchCompanyDiscretionaryDataAlphaNumeric(t *testing.T) {
	testBatchCompanyDiscretionaryDataAlphaNumeric(t)
}

// BenchmarkBatchCompanyDiscretionaryDataAlphaNumeric benchmarks validating company discretionary data is alphanumeric
func BenchmarkBatchCompanyDiscretionaryDataAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCompanyDiscretionaryDataAlphaNumeric(b)
	}
}

// testBatchCompanyIdentificationAlphaNumeric validates company identification is alphanumeric
func testBatchCompanyIdentificationAlphaNumeric(t testing.TB) {
	bh := mockBatchHeader()
	bh.CompanyIdentification = "®"
	err := bh.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCompanyIdentificationAlphaNumeric tests validating company identification is alphanumeric
func TestBatchCompanyIdentificationAlphaNumeric(t *testing.T) {
	testBatchCompanyIdentificationAlphaNumeric(t)
}

// BenchmarkBatchCompanyIdentificationAlphaNumeric benchmarks validating company identification is alphanumeric
func BenchmarkBatchCompanyIdentificationAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCompanyIdentificationAlphaNumeric(b)
	}
}

// testBatchCompanyEntryDescriptionAlphaNumeric validates company entry description is alphanumeric
func testBatchCompanyEntryDescriptionAlphaNumeric(t testing.TB) {
	bh := mockBatchHeader()
	bh.CompanyEntryDescription = "P®YROLL"
	err := bh.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchCompanyEntryDescriptionAlphaNumeric tests validating company entry description is alphanumeric
func TestBatchCompanyEntryDescriptionAlphaNumeric(t *testing.T) {
	testBatchCompanyEntryDescriptionAlphaNumeric(t)
}

// BenchmarkBatchCompanyEntryDescriptionAlphaNumeric benchmarks validating company entry description is alphanumeric
func BenchmarkBatchCompanyEntryDescriptionAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchCompanyEntryDescriptionAlphaNumeric(b)
	}
}

// testBHFieldInclusionCompanyName validates company name field inclusion
func testBHFieldInclusionCompanyName(t testing.TB) {
	bh := mockBatchHeader()
	bh.CompanyName = ""
	err := bh.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBHFieldInclusionCompanyName tests validating company name field inclusion
func TestBHFieldInclusionCompanyName(t *testing.T) {
	testBHFieldInclusionCompanyName(t)
}

// BenchmarkBHFieldInclusionCompanyName benchmarks validating company name field inclusion
func BenchmarkBHFieldInclusionCompanyName(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBHFieldInclusionCompanyName(b)
	}
}

// testBHFieldInclusionCompanyIdentification validates company identification field inclusion
func testBHFieldInclusionCompanyIdentification(t testing.TB) {
	bh := mockBatchHeader()
	bh.CompanyIdentification = ""
	err := bh.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBHFieldInclusionCompanyIdentification tests validating company identification field inclusion
func TestBHFieldInclusionCompanyIdentification(t *testing.T) {
	testBHFieldInclusionCompanyIdentification(t)
}

// BenchmarkBHFieldInclusionCompanyIdentification benchmarks validating company identification field inclusion
func BenchmarkBHFieldInclusionCompanyIdentification(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBHFieldInclusionCompanyIdentification(b)
	}
}

// testBHFieldInclusionStandardEntryClassCode validates SEC Code field inclusion
func testBHFieldInclusionStandardEntryClassCode(t testing.TB) {
	bh := mockBatchHeader()
	bh.StandardEntryClassCode = ""
	err := bh.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBHFieldInclusionStandardEntryClassCode tests validating SEC Code field inclusion
func TestBHFieldInclusionStandardEntryClassCode(t *testing.T) {
	testBHFieldInclusionStandardEntryClassCode(t)
}

// BenchmarkBHFieldInclusionStandardEntryClassCode benchmarks validating SEC Code field inclusion
func BenchmarkBHFieldInclusionStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBHFieldInclusionStandardEntryClassCode(b)
	}
}

// testBHFieldInclusionCompanyEntryDescription validates Company Entry Description field inclusion
func testBHFieldInclusionCompanyEntryDescription(t testing.TB) {
	bh := mockBatchHeader()
	bh.CompanyEntryDescription = ""
	err := bh.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBHFieldInclusionCompanyEntryDescription tests validating Company Entry Description field inclusion
func Test(t *testing.T) {
	testBHFieldInclusionCompanyEntryDescription(t)
}

// BenchmarkBHFieldInclusionCompanyEntryDescription benchmarks validating Company Entry Description field inclusion
func BenchmarkBHFieldInclusionCompanyEntryDescription(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBHFieldInclusionCompanyEntryDescription(b)
	}
}

// testBHFieldInclusionOriginatorStatusCode validates Originator Status Code field inclusion
func testBHFieldInclusionOriginatorStatusCode(t testing.TB) {
	bh := mockBatchHeader()
	bh.OriginatorStatusCode = 0
	err := bh.Validate()
	if !base.Match(err, ErrOrigStatusCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBHFieldInclusionOriginatorStatusCode tests validating Originator Status Code field inclusion
func TestBHFieldInclusionOriginatorStatusCode(t *testing.T) {
	testBHFieldInclusionOriginatorStatusCode(t)
}

// BenchmarkBHFieldInclusionOriginatorStatusCode benchmarks validating Originator Status Code field inclusion
func BenchmarkBHFieldInclusionOriginatorStatusCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBHFieldInclusionOriginatorStatusCode(b)
	}
}

// testBHFieldInclusionODFIIdentification validates ODFIIdentification field inclusion
func testBHFieldInclusionODFIIdentification(t testing.TB) {
	bh := mockBatchHeader()
	bh.ODFIIdentification = ""
	err := bh.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBHFieldInclusionODFIIdentification tests validating ODFIIdentification field inclusion
func TestBHFieldInclusionODFIIdentification(t *testing.T) {
	testBHFieldInclusionODFIIdentification(t)
}

// BenchmarkBHFieldInclusionODFIIdentification benchmarks validating ODFIIdentification field inclusion
func BenchmarkBHFieldInclusionODFIIdentification(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBHFieldInclusionODFIIdentification(b)
	}
}

func TestBatchHeaderENR__EffectiveEntryDateField(t *testing.T) {
	bh := mockBatchHeader()

	// ENR batches require EffectiveEntryDate to be space filled
	bh.CompanyEntryDescription = "AUTOENROLL"
	if v, ans := bh.EffectiveEntryDateField(), "      "; v != ans {
		t.Errorf("got %q (len=%d), expected space filled (len=6)", v, len(ans))
	}
}

func TestBatchHeader__LiftEffectiveEntryDate(t *testing.T) {
	bh := mockBatchHeader()
	bh.EffectiveEntryDate = "190730"

	if tt, err := bh.LiftEffectiveEntryDate(); err != nil {
		t.Fatal(err)
	} else {
		if tt.String() != "2019-07-30 00:00:00 +0000 UTC" {
			t.Errorf("tt=%v", tt)
		}
	}

	bh.EffectiveEntryDate = "aaaaaaaa"
	if _, err := bh.LiftEffectiveEntryDate(); err == nil {
		t.Error("expected error")
	}
}

func TestBatchHeader__SetValidation(t *testing.T) {
	bh := NewBatchHeader()
	bh.SetValidation(&ValidateOpts{})

	// nil out
	bh = nil
	bh.SetValidation(&ValidateOpts{})
}
