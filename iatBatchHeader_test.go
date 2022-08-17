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

// mockIATBatchHeaderFF creates a IAT BatchHeader that is Fixed-Fixed
func mockIATBatchHeaderFF() *IATBatchHeader {
	bh := NewIATBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.ForeignExchangeIndicator = "FF"
	bh.ForeignExchangeReferenceIndicator = 3
	bh.ISODestinationCountryCode = "US"
	bh.OriginatorIdentification = "123456789"
	bh.StandardEntryClassCode = IAT
	bh.CompanyEntryDescription = "TRADEPAYMT"
	bh.ISOOriginatingCurrencyCode = "CAD"
	bh.ISODestinationCurrencyCode = "USD"
	bh.ODFIIdentification = "23138010"
	return bh
}

// mockIATBatchReturnHeaderFF creates a IAT Return BatchHeader that is Fixed-Fixed
func mockIATReturnBatchHeaderFF() *IATBatchHeader {
	bh := NewIATBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.ForeignExchangeIndicator = "FF"
	bh.ForeignExchangeReferenceIndicator = 3
	bh.ISODestinationCountryCode = "US"
	bh.OriginatorIdentification = "123456789"
	bh.StandardEntryClassCode = IAT
	bh.CompanyEntryDescription = "TRADEPAYMT"
	bh.ISOOriginatingCurrencyCode = "CAD"
	bh.ISODestinationCurrencyCode = "USD"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockIATNOCBatchHeaderFF creates a IAT Return BatchHeader that is Fixed-Fixed
func mockIATNOCBatchHeaderFF() *IATBatchHeader {
	bh := NewIATBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.IATIndicator = "IATCOR"
	bh.ForeignExchangeIndicator = "FF"
	bh.ForeignExchangeReferenceIndicator = 3
	bh.ISODestinationCountryCode = "US"
	bh.OriginatorIdentification = "123456789"
	bh.StandardEntryClassCode = COR
	bh.CompanyEntryDescription = "TRADEPAYMT"
	bh.ISOOriginatingCurrencyCode = "CAD"
	bh.ISODestinationCurrencyCode = "USD"
	bh.ODFIIdentification = "12104288"
	return bh
}

// testMockIATBatchHeaderFF creates a IAT BatchHeader Fixed-Fixed
func testMockIATBatchHeaderFF(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	if err := bh.Validate(); err != nil {
		t.Error("mockIATBatchHeaderFF does not validate and will break other tests: ", err)
	}
	if bh.ServiceClassCode != CreditsOnly {
		t.Error("ServiceClassCode dependent default value has changed")
	}
	if bh.ForeignExchangeIndicator != "FF" {
		t.Error("ForeignExchangeIndicator does not validate and will break other tests")
	}
	if bh.ForeignExchangeReferenceIndicator != 3 {
		t.Error("ForeignExchangeReferenceIndicator does not validate and will break other tests")
	}
	if bh.ISODestinationCountryCode != "US" {
		t.Error("DestinationCountryCode dependent default value has changed")
	}
	if bh.OriginatorIdentification != "123456789" {
		t.Error("OriginatorIdentification dependent default value has changed")
	}
	if bh.StandardEntryClassCode != IAT {
		t.Error("StandardEntryClassCode dependent default value has changed")
	}
	if bh.CompanyEntryDescription != "TRADEPAYMT" {
		t.Error("CompanyEntryDescription dependent default value has changed")
	}
	if bh.ISOOriginatingCurrencyCode != "CAD" {
		t.Error("ISOOriginatingCurrencyCode dependent default value has changed")
	}
	if bh.ISODestinationCurrencyCode != "USD" {
		t.Error("ISODestinationCurrencyCode dependent default value has changed")
	}
	if bh.ODFIIdentification != "23138010" {
		t.Error("ODFIIdentification dependent default value has changed")
	}
}

// TestMockIATBatchHeaderFF tests creating a IAT BatchHeader Fixed-Fixed
func TestMockIATBatchHeaderFF(t *testing.T) {
	testMockIATBatchHeaderFF(t)
}

// BenchmarkMockIATBatchHeaderFF benchmarks creating a IAT BatchHeader Fixed-Fixed
func BenchmarkMockIATBatchHeaderFF(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockIATBatchHeaderFF(b)
	}
}

// testParseIATBatchHeader parses a known IAT BatchHeader record string
func testParseIATBatchHeader(t testing.TB) {
	var line = "5220                FF3               US123456789 IATTRADEPAYMTCADUSD180621   1231380100000001"
	r := NewReader(strings.NewReader(line))
	r.line = line
	if err := r.parseIATBatchHeader(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.IATCurrentBatch.GetHeader()

	if record.ServiceClassCode != CreditsOnly {
		t.Errorf("ServiceClassCode Expected '220' got: %v", record.ServiceClassCode)
	}
	if record.IATIndicator != "" {
		t.Errorf("IATIndicator Expected '' got: %v", record.IATIndicator)
	}
	if record.ForeignExchangeIndicator != "FF" {
		t.Errorf("ForeignExchangeIndicator Expected '                ' got: %v",
			record.ForeignExchangeIndicator)
	}
	if record.ForeignExchangeReferenceIndicator != 3 {
		t.Errorf("ForeignExchangeReferenceIndicator Expected '                ' got: %v",
			record.ForeignExchangeReferenceIndicator)
	}
	if record.ForeignExchangeReferenceField() != "               " {
		t.Errorf("ForeignExchangeReference Expected '                ' got: %v",
			record.ForeignExchangeReference)
	}
	if record.StandardEntryClassCode != IAT {
		t.Errorf("StandardEntryClassCode Expected 'PPD' got: %v", record.StandardEntryClassCode)
	}
	if record.CompanyEntryDescription != "TRADEPAYMT" {
		t.Errorf("CompanyEntryDescription Expected 'TRADEPAYMT' got: %v", record.CompanyEntryDescriptionField())
	}

	if record.EffectiveEntryDateField() != "180621" {
		t.Errorf("EffectiveEntryDate Expected '180621' got: %v", record.EffectiveEntryDateField())
	}
	if record.SettlementDate != "   " {
		t.Errorf("SettlementDate Expected '   ' got: %v", record.SettlementDate)
	}
	if record.OriginatorStatusCode != 1 {
		t.Errorf("OriginatorStatusCode Expected 1 got: %v", record.OriginatorStatusCode)
	}
	if record.ODFIIdentification != "23138010" {
		t.Errorf("OdfiIdentification Expected '23138010' got: %v", record.ODFIIdentificationField())
	}
	if record.BatchNumberField() != "0000001" {
		t.Errorf("BatchNumber Expected '0000001' got: %v", record.BatchNumberField())
	}
}

// TestParseIATBatchHeader tests parsing a known IAT BatchHeader record string
func TestParseIATBatchHeader(t *testing.T) {
	testParseIATBatchHeader(t)
}

// BenchmarkParseBatchHeader benchmarks parsing a known IAT BatchHeader record string
func BenchmarkParseIATBatchHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testParseIATBatchHeader(b)
	}
}

// testIATBHString validates that a known parsed IAT Batch Header
// can be return to a string of the same value
func testIATBHString(t testing.TB) {
	var line = "5220                FF3               US123456789 IATTRADEPAYMTCADUSD180621   1231380100000001"
	r := NewReader(strings.NewReader(line))
	r.line = line
	if err := r.parseIATBatchHeader(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.IATCurrentBatch.GetHeader()

	if v := record.String(); v != line {
		t.Errorf("Strings do not match:\n   v=%q\nline=%q", v, line) // vertically aligned
	}
}

// TestIATBHString tests validating that a known parsed IAT BatchHeader
// can be return to a string of the same value
func TestIATBHString(t *testing.T) {
	testIATBHString(t)
}

// BenchmarkIATBHString benchmarks validating that a known parsed IAT BatchHeader
// can be return to a string of the same value
func BenchmarkIATBHString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBHString(b)
	}
}

// testIATBHFVString validates that a known parsed IAT Batch Header
// can be return to a string of the same value
func testIATBHFVString(t testing.TB) {
	var line = "5220                FV2123456789012345US123456789 IATTRADEPAYMTCADUSD180621   1231380100000001"
	r := NewReader(strings.NewReader(line))
	r.line = line
	if err := r.parseIATBatchHeader(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.IATCurrentBatch.GetHeader()

	if v := record.String(); v != line {
		t.Errorf("Strings do not match:\n   v=%q\nline=%q", v, line)
	}
}

// TestIATBHFVString tests validating that a known parsed IAT BatchHeader
// can be return to a string of the same value
func TestIATBHFVString(t *testing.T) {
	testIATBHFVString(t)
}

// BenchmarkIATBHFVString benchmarks validating that a known parsed IAT BatchHeader
// can be return to a string of the same value
func BenchmarkIATBHFVString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBHFVString(b)
	}
}

// testValidateIATBHServiceClassCode validates error if IATBatchHeader
// ServiceClassCode is invalid
func testValidateIATBHServiceClassCode(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.ServiceClassCode = 999
	err := bh.Validate()
	if !base.Match(err, ErrServiceClass) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestValidateIATBHServiceClassCode tests validating error if IATBatchHeader
// ServiceClassCode is invalid
func TestValidateIATBHServiceClassCode(t *testing.T) {
	testValidateIATBHServiceClassCode(t)
}

// BenchmarkValidateIATBHServiceClassCode benchmarks validating error if IATBatchHeader
// ServiceClassCode is invalid
func BenchmarkValidateIATBHServiceClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateIATBHServiceClassCode(b)
	}
}

// testValidateIATBHForeignExchangeIndicator validates error if IATBatchHeader
// ForeignExchangeIndicator is invalid
func testValidateIATBHForeignExchangeIndicator(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.ForeignExchangeIndicator = "XY"
	err := bh.Validate()
	if !base.Match(err, ErrForeignExchangeIndicator) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestValidateIATBHForeignExchangeIndicator tests validating error if IATBatchHeader
// ForeignExchangeIndicator is invalid
func TestValidateIATBHForeignExchangeIndicator(t *testing.T) {
	testValidateIATBHForeignExchangeIndicator(t)
}

// BenchmarkValidateIATBHForeignExchangeIndicator benchmarks validating error if IATBatchHeader
// ForeignExchangeIndicator is invalid
func BenchmarkValidateIATBHForeignExchangeIndicator(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateIATBHForeignExchangeIndicator(b)
	}
}

// testValidateIATBHForeignExchangeReferenceIndicator validates error if IATBatchHeader
// ForeignExchangeReferenceIndicator is invalid
func testValidateIATBHForeignExchangeReferenceIndicator(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.ForeignExchangeReferenceIndicator = 5
	err := bh.Validate()
	if !base.Match(err, ErrForeignExchangeReferenceIndicator) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestValidateIATBHForeignExchangeReferenceIndicator tests validating error if IATBatchHeader
// ForeignExchangeReferenceIndicator is invalid
func TestValidateIATBHForeignExchangeReferenceIndicator(t *testing.T) {
	testValidateIATBHForeignExchangeReferenceIndicator(t)
}

// BenchmarkValidateIATBHForeignExchangeReferenceIndicator benchmarks validating error if IATBatchHeader
// ForeignExchangeReferenceIndicator is invalid
func BenchmarkValidateIATBHForeignExchangeReferenceIndicator(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateIATBHForeignExchangeReferenceIndicator(b)
	}
}

// testValidateIATBHISODestinationCountryCode validates error if IATBatchHeader
// ISODestinationCountryCode is invalid
func testValidateIATBHISODestinationCountryCode(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.ISODestinationCountryCode = "®"
	err := bh.Validate()
	if !base.Match(err, ErrValidISO3166) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestValidateIATBHISODestinationCountryCode tests validating error if IATBatchHeader
// ISODestinationCountryCode is invalid
func TestValidateIATBHISODestinationCountryCode(t *testing.T) {
	testValidateIATBHISODestinationCountryCode(t)
}

// BenchmarkValidateIATBHISODestinationCountryCode benchmarks validating error if IATBatchHeader
// ISODestinationCountryCode is invalid
func BenchmarkValidateIATBHISODestinationCountryCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateIATBHISODestinationCountryCode(b)
	}
}

// testValidateIATBHOriginatorIdentification validates error if IATBatchHeader
// OriginatorIdentification is invalid
func testValidateIATBHOriginatorIdentification(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.OriginatorIdentification = "®"
	err := bh.Validate()
	// TODO: are we expecting there to be no errors here?
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestValidateIATBHOriginatorIdentification tests validating error if IATBatchHeader
// OriginatorIdentification is invalid
func TestValidateIATBHOriginatorIdentification(t *testing.T) {
	testValidateIATBHOriginatorIdentification(t)
}

// BenchmarkValidateIATBHOriginatorIdentification benchmarks validating error if IATBatchHeader
// OriginatorIdentification is invalid
func BenchmarkValidateIATBHOriginatorIdentification(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateIATBHOriginatorIdentification(b)
	}
}

// testValidateIATBHStandardEntryClassCode validates error if IATBatchHeader
// StandardEntryClassCode is invalid
func testValidateIATBHStandardEntryClassCode(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.StandardEntryClassCode = "ABC"
	err := bh.Validate()
	if !base.Match(err, ErrSECCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestValidateIATBHStandardEntryClassCode tests validating error if IATBatchHeader
// StandardEntryClassCode is invalid
func TestValidateIATBHStandardEntryClassCode(t *testing.T) {
	testValidateIATBHStandardEntryClassCode(t)
}

// BenchmarkValidateIATBHStandardEntryClassCode benchmarks validating error if IATBatchHeader
// StandardEntryClassCode is invalid
func BenchmarkValidateIATBHStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateIATBHStandardEntryClassCode(b)
	}
}

// testValidateIATBHCompanyEntryDescription validates error if IATBatchHeader
// CompanyEntryDescription is invalid
func testValidateIATBHCompanyEntryDescription(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.CompanyEntryDescription = "®"
	err := bh.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestValidateIATBHCompanyEntryDescription tests validating error if IATBatchHeader
// CompanyEntryDescription is invalid
func TestValidateIATBHCompanyEntryDescription(t *testing.T) {
	testValidateIATBHCompanyEntryDescription(t)
}

// BenchmarkValidateIATBHCompanyEntryDescription benchmarks validating error if IATBatchHeader
// CompanyEntryDescription is invalid
func BenchmarkValidateIATBHCompanyEntryDescription(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateIATBHCompanyEntryDescription(b)
	}
}

// testValidateIATBHISOOriginatingCurrencyCode validates error if IATBatchHeader
// ISOOriginatingCurrencyCode is invalid
func testValidateIATBHISOOriginatingCurrencyCode(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.ISOOriginatingCurrencyCode = "®"
	err := bh.Validate()
	if !base.Match(err, ErrValidISO4217) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestValidateIATBHISOOriginatingCurrencyCode tests validating error if IATBatchHeader
// ISOOriginatingCurrencyCode is invalid
func TestValidateIATBHISOOriginatingCurrencyCode(t *testing.T) {
	testValidateIATBHISOOriginatingCurrencyCode(t)
}

// BenchmarkValidateIATBHISOOriginatingCurrencyCode benchmarks validating error if IATBatchHeader
// ISOOriginatingCurrencyCode is invalid
func BenchmarkValidateIATBHISOOriginatingCurrencyCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateIATBHISOOriginatingCurrencyCode(b)
	}
}

// testValidateIATBHISODestinationCurrencyCode validates error if IATBatchHeader
// ISODestinationCurrencyCode is invalid
func testValidateIATBHISODestinationCurrencyCode(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.ISODestinationCurrencyCode = "®"
	err := bh.Validate()
	if !base.Match(err, ErrValidISO4217) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestValidateIATBHISODestinationCurrencyCode tests validating error if IATBatchHeader
// ISODestinationCurrencyCode is invalid
func TestValidateIATBHISODestinationCurrencyCode(t *testing.T) {
	testValidateIATBHISODestinationCurrencyCode(t)
}

// BenchmarkValidateIATBHISODestinationCurrencyCode benchmarks validating error if IATBatchHeader
// ISODestinationCurrencyCode is invalid
func BenchmarkValidateIATBHISODestinationCurrencyCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateIATBHISODestinationCurrencyCode(b)
	}
}

// testValidateIATBHOriginatorStatusCode validates error if IATBatchHeader
// OriginatorStatusCode is invalid
func testValidateIATBHOriginatorStatusCode(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.OriginatorStatusCode = 7
	err := bh.Validate()
	if !base.Match(err, ErrOrigStatusCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestValidateIATBHOriginatorStatusCode tests validating error if IATBatchHeader
// OriginatorStatusCode is invalid
func TestValidateIATBHOriginatorStatusCode(t *testing.T) {
	testValidateIATBHOriginatorStatusCode(t)
}

// BenchmarkValidateIATBHOriginatorStatusCode benchmarks validating error if IATBatchHeader
// OriginatorStatusCode is invalid
func BenchmarkValidateIATBHOriginatorStatusCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testValidateIATBHOriginatorStatusCode(b)
	}
}

//FieldInclusion

// testIATBHServiceClassCode validates IATBatchHeader ServiceClassCode fieldInclusion
func testIATBHServiceClassCode(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.ServiceClassCode = 0
	err := bh.Validate()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBHServiceClassCode tests validating IATBatchHeader ServiceClassCode fieldInclusion
func TestIATBHServiceClassCode(t *testing.T) {
	testIATBHServiceClassCode(t)
}

// BenchmarkIATBHServiceClassCode benchmarks validating IATBatchHeader ServiceClassCode fieldInclusion
func BenchmarkIATBHServiceClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBHServiceClassCode(b)
	}
}

// testIATBHForeignExchangeIndicator validates IATBatchHeader ForeignExchangeIndicator fieldInclusion
func testIATBHForeignExchangeIndicator(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.ForeignExchangeIndicator = ""
	err := bh.Validate()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBHForeignExchangeIndicator tests validating IATBatchHeader ForeignExchangeIndicator fieldInclusion
func TestIATBHForeignExchangeIndicator(t *testing.T) {
	testIATBHForeignExchangeIndicator(t)
}

// BenchmarkIATBHForeignExchangeIndicator benchmarks validating IATBatchHeader ForeignExchangeIndicator fieldInclusion
func BenchmarkIATBHForeignExchangeIndicator(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBHForeignExchangeIndicator(b)
	}
}

// testIATBHForeignExchangeReferenceIndicator validates IATBatchHeader ForeignExchangeReferenceIndicator fieldInclusion
func testIATBHForeignExchangeReferenceIndicator(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.ForeignExchangeReferenceIndicator = 0
	err := bh.Validate()
	if !base.Match(err, ErrFieldRequired) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBHForeignExchangeReferenceIndicator tests validating IATBatchHeader ForeignExchangeReferenceIndicator fieldInclusion
func TestIATBHForeignExchangeReferenceIndicator(t *testing.T) {
	testIATBHForeignExchangeReferenceIndicator(t)
}

// BenchmarkIATBHForeignExchangeReferenceIndicator benchmarks validating IATBatchHeader ForeignExchangeReferenceIndicator fieldInclusion
func BenchmarkIATBHForeignExchangeReferenceIndicator(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBHForeignExchangeReferenceIndicator(b)
	}
}

// testIATBHISODestinationCountryCode validates IATBatchHeader ISODestinationCountryCode fieldInclusion
func testIATBHISODestinationCountryCode(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.ISODestinationCountryCode = ""
	err := bh.Validate()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBHISODestinationCountryCode tests validating IATBatchHeader ISODestinationCountryCode fieldInclusion
func TestIATBHISODestinationCountryCode(t *testing.T) {
	testIATBHISODestinationCountryCode(t)
}

// BenchmarkIATBHISODestinationCountryCode benchmarks validating IATBatchHeader ISODestinationCountryCode fieldInclusion
func BenchmarkIATBHISODestinationCountryCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBHISODestinationCountryCode(b)
	}
}

// testIATBHOriginatorIdentification validates IATBatchHeader OriginatorIdentification fieldInclusion
func testIATBHOriginatorIdentification(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.OriginatorIdentification = ""
	err := bh.Validate()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBHOriginatorIdentification tests validating IATBatchHeader OriginatorIdentification fieldInclusion
func TestIATBHOriginatorIdentification(t *testing.T) {
	testIATBHOriginatorIdentification(t)
}

// BenchmarkIATBHOriginatorIdentification benchmarks validating IATBatchHeader OriginatorIdentification fieldInclusion
func BenchmarkIATBHOriginatorIdentification(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBHOriginatorIdentification(b)
	}
}

// testIATBHStandardEntryClassCode validates IATBatchHeader StandardEntryClassCode fieldInclusion
func testIATBHStandardEntryClassCode(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.StandardEntryClassCode = ""
	err := bh.Validate()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBHStandardEntryClassCode tests validating IATBatchHeader StandardEntryClassCode fieldInclusion
func TestIATBHStandardEntryClassCode(t *testing.T) {
	testIATBHStandardEntryClassCode(t)
}

// BenchmarkIATBHStandardEntryClassCode benchmarks validating IATBatchHeader StandardEntryClassCode fieldInclusion
func BenchmarkIATBHStandardEntryClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBHStandardEntryClassCode(b)
	}
}

// testIATBHCompanyEntryDescription validates IATBatchHeader CompanyEntryDescription fieldInclusion
func testIATBHCompanyEntryDescription(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.CompanyEntryDescription = ""
	err := bh.Validate()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBHCompanyEntryDescription tests validating IATBatchHeader CompanyEntryDescription fieldInclusion
func TestIATBHCompanyEntryDescription(t *testing.T) {
	testIATBHCompanyEntryDescription(t)
}

// BenchmarkIATBHCompanyEntryDescription benchmarks validating IATBatchHeader CompanyEntryDescription fieldInclusion
func BenchmarkIATBHCompanyEntryDescription(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBHCompanyEntryDescription(b)
	}
}

// testIATBHISOOriginatingCurrencyCode validates IATBatchHeader ISOOriginatingCurrencyCode fieldInclusion
func testIATBHISOOriginatingCurrencyCode(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.ISOOriginatingCurrencyCode = ""
	err := bh.Validate()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBHISOOriginatingCurrencyCode tests validating IATBatchHeader ISOOriginatingCurrencyCode fieldInclusion
func TestIATBHISOOriginatingCurrencyCode(t *testing.T) {
	testIATBHISOOriginatingCurrencyCode(t)
}

// BenchmarkIATBHISOOriginatingCurrencyCode benchmarks validating IATBatchHeader ISOOriginatingCurrencyCode fieldInclusion
func BenchmarkIATBHISOOriginatingCurrencyCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBHISOOriginatingCurrencyCode(b)
	}
}

// testIATBHISODestinationCurrencyCode validates IATBatchHeader ISODestinationCurrencyCode fieldInclusion
func testIATBHISODestinationCurrencyCode(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.ISODestinationCurrencyCode = ""
	err := bh.Validate()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBHISODestinationCurrencyCode tests validating IATBatchHeader ISODestinationCurrencyCode fieldInclusion
func TestIATBHISODestinationCurrencyCode(t *testing.T) {
	testIATBHISODestinationCurrencyCode(t)
}

// BenchmarkIATBHISODestinationCurrencyCode benchmarks validating IATBatchHeader ISODestinationCurrencyCode fieldInclusion
func BenchmarkIATBHISODestinationCurrencyCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBHISODestinationCurrencyCode(b)
	}
}

// testIATBHODFIIdentification validates IATBatchHeader ODFIIdentification fieldInclusion
func testIATBHODFIIdentification(t testing.TB) {
	bh := mockIATBatchHeaderFF()
	bh.ODFIIdentification = ""
	err := bh.Validate()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBHODFIIdentification tests validating IATBatchHeader ODFIIdentification fieldInclusion
func TestIATBHODFIIdentification(t *testing.T) {
	testIATBHODFIIdentification(t)
}

// BenchmarkIATBHODFIIdentification benchmarks validating IATBatchHeader ODFIIdentification fieldInclusion
func BenchmarkIATBHODFIIdentification(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBHODFIIdentification(b)
	}
}
