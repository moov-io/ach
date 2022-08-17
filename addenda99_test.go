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
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/base"
)

func mockAddenda99() *Addenda99 {
	addenda99 := NewAddenda99()
	addenda99.ReturnCode = "R07"
	addenda99.OriginalTrace = "99912340000015"
	addenda99.AddendaInformation = "Authorization Revoked"
	addenda99.OriginalDFI = "9101298"

	return addenda99
}

func testAddenda99Parse(t testing.TB) {
	addenda99 := NewAddenda99()
	line := "799R07099912340000015      09101298Authorization revoked                       091012980000066"
	addenda99.Parse(line)
	// walk the Addenda99 struct
	if addenda99.TypeCode != "99" {
		t.Errorf("expected %v got %v", "99", addenda99.TypeCode)
	}
	if addenda99.ReturnCode != "R07" {
		t.Errorf("expected %v got %v", "R07", addenda99.ReturnCode)
	}
	if addenda99.OriginalTrace != "099912340000015" {
		t.Errorf("expected: %v got: %v", "099912340000015", addenda99.OriginalTrace)
	}
	if addenda99.DateOfDeathField() == "" {
		t.Errorf("got: %v", addenda99.DateOfDeath)
	}
	if addenda99.OriginalDFI != "09101298" {
		t.Errorf("expected: %s got: %s", "09101298", addenda99.OriginalDFI)
	}
	if addenda99.AddendaInformation != "Authorization revoked" {
		t.Errorf("expected: %v got: %v", "Authorization revoked", addenda99.AddendaInformation)
	}
	if addenda99.TraceNumber != "091012980000066" {
		t.Errorf("expected: %v got: %v", "091012980000066", addenda99.TraceNumber)
	}
	if code := addenda99.ReturnCodeField(); code != nil {
		if code.Code != "R07" || code.Reason != "Authorization Revoked by Customer" {
			t.Errorf("code.Code=%q code.Reason=%q", code.Code, code.Reason)
		}
	} else {
		t.Errorf("got nil ReturnCode")
	}
}

func TestAddenda99__LookupReturnCode(t *testing.T) {
	if code := LookupReturnCode(""); code != nil {
		t.Error("expected nil ReturnCode")
	}
	if code := LookupReturnCode("R02"); code == nil {
		t.Error("expected ReturnCode")
	} else {
		if code.Code != "R02" {
			t.Errorf("code.Code=%s", code.Code)
		}
		if code.Reason != "Account Closed" {
			t.Errorf("code.Reason=%s", code.Reason)
		}
	}
	if code := LookupReturnCode("R99"); code != nil {
		t.Errorf("expected nil: %#v", code)
	}
}

func TestAddenda99Parse(t *testing.T) {
	testAddenda99Parse(t)
}

func BenchmarkAddenda99Parse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda99Parse(b)
	}
}

func testAddenda99String(t testing.TB) {
	addenda99 := NewAddenda99()
	line := "799R07099912340000015      09101298Authorization revoked                       091012980000066"
	addenda99.Parse(line)

	if addenda99.String() != line {
		t.Errorf("\n expected: %v\n got     : %v", line, addenda99.String())
	}
}

func TestAddenda99String(t *testing.T) {
	testAddenda99String(t)
}

func BenchmarkAddenda99String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda99String(b)
	}
}

// This is not an exported function but utilized for validation
func testAddenda99MakeReturnCodeDict(t testing.TB) {
	codes := makeReturnCodeDict()
	// check if known code is present
	_, prs := codes["R01"]
	if !prs {
		t.Error("Return Code R01 was not found in the ReturnCodeDict")
	}
	// check if invalid code is present
	_, prs = codes["ABC"]
	if prs {
		t.Error("Valid return for an invalid return code key")
	}
}

func TestAddenda99MakeReturnCodeDict(t *testing.T) {
	testAddenda99MakeReturnCodeDict(t)
}

func BenchmarkAddenda99MakeReturnCodeDict(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda99MakeReturnCodeDict(b)
	}
}

func testAddenda99ValidateTrue(t testing.TB) {
	addenda99 := mockAddenda99()
	addenda99.ReturnCode = "R13"
	err := addenda99.Validate()
	// no error expected
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda99ValidateTrue(t *testing.T) {
	testAddenda99ValidateTrue(t)
}

func BenchmarkAddenda99ValidateTrue(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda99ValidateTrue(b)
	}
}

func testAddenda99ValidateReturnCodeFalse(t testing.TB) {
	addenda99 := mockAddenda99()
	addenda99.ReturnCode = ""
	err := addenda99.Validate()
	if !base.Match(err, ErrAddenda99ReturnCode) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda99ValidateReturnCodeFalse(t *testing.T) {
	testAddenda99ValidateReturnCodeFalse(t)
}

func BenchmarkAddenda99ValidateReturnCodeFalse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda99ValidateReturnCodeFalse(b)
	}
}

func testAddenda99OriginalTraceField(t testing.TB) {
	addenda99 := mockAddenda99()
	addenda99.OriginalTrace = "12345"
	if addenda99.OriginalTraceField() != "000000000012345" {
		t.Errorf("expected %v received %v", "000000000012345", addenda99.OriginalTraceField())
	}
}

func TestAddenda99OriginalTraceField(t *testing.T) {
	testAddenda99OriginalTraceField(t)
}

func BenchmarkAddenda99OriginalTraceField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda99OriginalTraceField(b)
	}
}

func testAddenda99DateOfDeathField(t testing.TB) {
	addenda99 := mockAddenda99()
	// Check for all zeros
	if addenda99.DateOfDeathField() != "      " {
		t.Errorf("expected %v received %v", "      ", addenda99.DateOfDeathField())
	}
	// Year: 1978 Month: October Day: 23
	addenda99.DateOfDeath = "781023"
	if addenda99.DateOfDeathField() != "781023" {
		t.Errorf("expected %v received %v", "781023", addenda99.DateOfDeathField())
	}
}

func TestAddenda99DateOfDeathField(t *testing.T) {
	testAddenda99DateOfDeathField(t)
}
func BenchmarkAddenda99DateOfDeathField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda99DateOfDeathField(b)
	}
}

func testAddenda99OriginalDFIField(t testing.TB) {
	addenda99 := mockAddenda99()
	exp := "09101298"
	if addenda99.OriginalDFIField() != exp {
		t.Errorf("expected %v received %v", exp, addenda99.OriginalDFIField())
	}
}

func TestAddenda99OriginalDFIField(t *testing.T) {
	testAddenda99OriginalDFIField(t)
}

func BenchmarkAddenda99OriginalDFIField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda99OriginalDFIField(b)
	}
}

func testAddenda99AddendaInformationField(t testing.TB) {
	addenda99 := mockAddenda99()
	exp := "Authorization Revoked                       "
	if addenda99.AddendaInformationField() != exp {
		t.Errorf("expected %v received %v", exp, addenda99.AddendaInformationField())
	}
}

func TestAddenda99AddendaInformationField(t *testing.T) {
	testAddenda99AddendaInformationField(t)
}

func BenchmarkAddenda99AddendaInformationField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda99AddendaInformationField(b)
	}
}

func testAddenda99TraceNumberField(t testing.TB) {
	addenda99 := mockAddenda99()
	addenda99.TraceNumber = "91012980000066"
	exp := "091012980000066"
	if addenda99.TraceNumberField() != exp {
		t.Errorf("expected %v received %v", exp, addenda99.TraceNumberField())
	}
}

func TestAddenda99TraceNumberField(t *testing.T) {
	testAddenda99TraceNumberField(t)
}

func BenchmarkAddenda99TraceNumberField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda99TraceNumberField(b)
	}
}

// testAddenda99TypeCode99 TypeCode is 99
func testAddenda99TypeCode99(t testing.TB) {
	addenda99 := mockAddenda99()
	addenda99.TypeCode = "05"
	err := addenda99.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda99TypeCode99 tests TypeCode is 99
func TestAddenda99TypeCode99(t *testing.T) {
	testAddenda99TypeCode99(t)
}

// BenchmarkAddenda99TypeCode99 benchmarks TypeCode is 99
func BenchmarkAddenda99TypeCode99(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda99TypeCode99(b)
	}
}

// testAddenda99TypeCodeNil validates TypeCode is ""
func testAddenda99TypeCodeNil(t testing.TB) {
	addenda99 := mockAddenda99()
	addenda99.TypeCode = ""
	err := addenda99.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda99TypeCodeES tests TypeCode is ""
func TestAddenda99TypeCodeNil(t *testing.T) {
	testAddenda99TypeCodeNil(t)
}

// BenchmarkAddenda99TypeCodeNil benchmarks TypeCode is ""
func BenchmarkAddenda99TypeCodeNil(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda99TypeCodeNil(b)
	}
}

// TestAddenda99RuneCountInString validates RuneCountInString
func TestAddenda99RuneCountInString(t *testing.T) {
	addenda99 := NewAddenda99()
	var line = "799"
	addenda99.Parse(line)

	if addenda99.DateOfDeath != "" {
		t.Error("Parsed with an invalid RuneCountInString not equal to 94")
	}
}

func TestAddenda99__MissingFileHeaderControl(t *testing.T) {
	// We have a usecase where returned files don't contain the FileHeader
	// or FileControl records. Our parser can handle this, but returns
	// errors are they're expected per the NACHA spec.
	//
	// This test just checks we can parse the file and get the expected errors.
	file, err := ReadFile(filepath.Join("test", "testdata", "return-no-file-header-control.ach"))
	if err == nil {
		t.Error("expected an error")
	}
	if !strings.Contains(err.Error(), ErrFileHeader.Error()) {
		t.Errorf("unexpected error: %v", err)
	}
	if !strings.Contains(err.Error(), ErrFileControl.Error()) {
		t.Errorf("unexpected error: %v", err)
	}

	if len(file.Batches) != 1 {
		t.Errorf("got %d batches", len(file.Batches))
	}
	if entries := file.Batches[0].GetEntries(); len(entries) != 1 {
		t.Errorf("got %d entries", len(entries))
	}
}

func TestAddenda99__SetValidation(t *testing.T) {
	addenda99 := NewAddenda99()
	addenda99.SetValidation(&ValidateOpts{
		CustomReturnCodes: true,
	})

	addenda99.ReturnCode = "@#$" // can safely say this will never be a real ReasonCode
	if err := addenda99.Validate(); err != nil {
		t.Fatal(err)
	}

	addenda99.SetValidation(&ValidateOpts{
		CustomReturnCodes: false,
	})
	if err := addenda99.Validate(); err == nil {
		t.Fatal("Did not flag invalid reason code")
	}
	addenda99 = nil
	addenda99.SetValidation(&ValidateOpts{})
}

func TestAddenda99__CustomReturnCode(t *testing.T) {
	mockBatch := mockBatch()
	// Add a Addenda Return to the mock batch
	if len(mockBatch.Entries) != 1 {
		t.Fatal("Expected 1 batch entry")
	}
	addenda99 := NewAddenda99()
	addenda99.SetValidation(&ValidateOpts{
		CustomReturnCodes: true,
	})

	addenda99.ReturnCode = "@#$" // can safely say this will never be a real ReasonCode
	mockBatch.Entries[0].Addenda99 = addenda99
	mockBatch.Entries[0].Category = CategoryReturn
	mockBatch.Entries[0].AddendaRecordIndicator = 1

	// replace last 2 of TraceNumber
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := mockBatch.verify(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}
