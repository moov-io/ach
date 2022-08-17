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
	"testing"

	"github.com/moov-io/base"
)

// mockAddenda10 creates a mock Addenda10 record
func mockAddenda10() *Addenda10 {
	addenda10 := NewAddenda10()
	addenda10.TransactionTypeCode = "ANN"
	addenda10.ForeignPaymentAmount = 100000
	addenda10.ForeignTraceNumber = "928383-23938"
	addenda10.Name = "BEK Enterprises"
	addenda10.EntryDetailSequenceNumber = 00000001
	return addenda10
}

// TestMockAddenda10 validates mockAddenda10
func TestMockAddenda10(t *testing.T) {
	addenda10 := mockAddenda10()
	if err := addenda10.Validate(); err != nil {
		t.Error("mockAddenda10 does not validate and will break other tests")
	}
}

// testAddenda10Parse parses Addenda10 record
func testAddenda10Parse(t testing.TB) {
	addenda10 := NewAddenda10()
	line := "710ANN000000000000100000928383-23938          BEK Enterprises                          0000001"
	addenda10.Parse(line)
	// walk the Addenda10 struct
	if addenda10.TypeCode != "10" {
		t.Errorf("expected %v got %v", "10", addenda10.TypeCode)
	}
	if addenda10.TransactionTypeCode != "ANN" {
		t.Errorf("expected %v got %v", "ANN", addenda10.TransactionTypeCode)
	}
	if addenda10.ForeignPaymentAmount != 100000 {
		t.Errorf("expected: %v got: %v", 100000, addenda10.ForeignPaymentAmount)
	}
	if addenda10.ForeignTraceNumber != "928383-23938" {
		t.Errorf("expected: %v got: %v", "928383-23938", addenda10.ForeignTraceNumber)
	}
	if addenda10.Name != "BEK Enterprises" {
		t.Errorf("expected: %s got: %s", "BEK Enterprises", addenda10.Name)
	}
	if addenda10.EntryDetailSequenceNumber != 0000001 {
		t.Errorf("expected: %v got: %v", 0000001, addenda10.EntryDetailSequenceNumber)
	}
}

// TestAddenda10Parse tests parsing Addenda10 record
func TestAddenda10Parse(t *testing.T) {
	testAddenda10Parse(t)
}

// BenchmarkAddenda10Parse benchmarks parsing Addenda10 record
func BenchmarkAddenda10Parse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda10Parse(b)
	}
}

// testAddenda10ValidTypeCode validates Addenda10 TypeCode
func testAddenda10ValidTypeCode(t testing.TB) {
	addenda10 := mockAddenda10()
	addenda10.TypeCode = "65"
	err := addenda10.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda10ValidTypeCode tests validating Addenda10 TypeCode
func TestAddenda10ValidTypeCode(t *testing.T) {
	testAddenda10ValidTypeCode(t)
}

// BenchmarkAddenda10ValidTypeCode benchmarks validating Addenda10 TypeCode
func BenchmarkAddenda10ValidTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda10ValidTypeCode(b)
	}
}

// testAddenda10TypeCode10 TypeCode is 10 if TypeCode is a valid TypeCode
func testAddenda10TypeCode10(t testing.TB) {
	addenda10 := mockAddenda10()
	addenda10.TypeCode = "05"
	err := addenda10.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda10TypeCode10 tests TypeCode is 10 if TypeCode is a valid TypeCode
func TestAddenda10TypeCode10(t *testing.T) {
	testAddenda10TypeCode10(t)
}

// BenchmarkAddenda10TypeCode10 benchmarks TypeCode is 10 if TypeCode is a valid TypeCode
func BenchmarkAddenda10TypeCode10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda10TypeCode10(b)
	}
}

// testAddenda10TransactionTypeCode validates TransactionTypeCode
func testAddenda10TransactionTypeCode(t testing.TB) {
	addenda10 := mockAddenda10()
	addenda10.TransactionTypeCode = "ABC"
	err := addenda10.Validate()
	if !base.Match(err, ErrTransactionTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda10TransactionTypeCode tests validating TransactionTypeCode
func TestAddenda10TransactionTypeCode(t *testing.T) {
	testAddenda10TransactionTypeCode(t)
}

// BenchmarkAddenda10TransactionTypeCode benchmarks validating TransactionTypeCode
func BenchmarkAddenda10TransactionTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda10TransactionTypeCode(b)
	}
}

// testForeignTraceNumberAlphaNumeric validates ForeignTraceNumber is alphanumeric
func testForeignTraceNumberAlphaNumeric(t testing.TB) {
	addenda10 := mockAddenda10()
	addenda10.ForeignTraceNumber = "®"
	err := addenda10.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestForeignTraceNumberAlphaNumeric tests validating ForeignTraceNumber is alphanumeric
func TestForeignTraceNumberAlphaNumeric(t *testing.T) {
	testForeignTraceNumberAlphaNumeric(t)
}

// BenchmarkForeignTraceNumberAlphaNumeric benchmarks validating ForeignTraceNumber is alphanumeric
func BenchmarkForeignTraceNumberAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testForeignTraceNumberAlphaNumeric(b)
	}
}

// testNameAlphaNumeric validates Name is alphanumeric
func testNameAlphaNumeric(t testing.TB) {
	addenda10 := mockAddenda10()
	addenda10.Name = "Jas®n"
	err := addenda10.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestNameAlphaNumeric tests validating Name is alphanumeric
func TestNameAlphaNumeric(t *testing.T) {
	testNameAlphaNumeric(t)
}

// BenchmarkNameAlphaNumeric benchmarks validating Name is alphanumeric
func BenchmarkNameAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testNameAlphaNumeric(b)
	}
}

// testAddenda10FieldInclusionTypeCode validates TypeCode fieldInclusion
func testAddenda10FieldInclusionTypeCode(t testing.TB) {
	addenda10 := mockAddenda10()
	addenda10.TypeCode = ""
	err := addenda10.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda10FieldInclusionTypeCode tests validating TypeCode fieldInclusion
func TestAddenda10FieldInclusionTypeCode(t *testing.T) {
	testAddenda10FieldInclusionTypeCode(t)
}

// BenchmarkAddenda10FieldInclusionTypeCode benchmarks validating TypeCode fieldInclusion
func BenchmarkAddenda10FieldInclusionTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda10FieldInclusionTypeCode(b)
	}
}

// testAddenda10FieldInclusionTransactionTypeCode validates TransactionTypeCode fieldInclusion
func testAddenda10FieldInclusionTransactionTypeCode(t testing.TB) {
	addenda10 := mockAddenda10()
	addenda10.TransactionTypeCode = ""
	err := addenda10.Validate()
	if !base.Match(err, ErrFieldRequired) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda10FieldInclusionTransactionTypeCode tests validating
// TransactionTypeCode fieldInclusion
func TestAddenda10FieldInclusionTransactionTypeCode(t *testing.T) {
	testAddenda10FieldInclusionTransactionTypeCode(t)
}

// BenchmarkAddenda10FieldInclusionTransactionTypeCode benchmarks validating
// TransactionTypeCode fieldInclusion
func BenchmarkAddenda10FieldInclusionTransactionTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda10FieldInclusionTransactionTypeCode(b)
	}
}

// testAddenda10FieldInclusionForeignPaymentAmount validates ForeignPaymentAmount fieldInclusion
func testAddenda10FieldInclusionForeignPaymentAmount(t testing.TB) {
	addenda10 := mockAddenda10()
	addenda10.ForeignPaymentAmount = 0
	err := addenda10.Validate()
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda10FieldInclusionForeignPaymentAmount tests validating ForeignPaymentAmount fieldInclusion
func TestAddenda10FieldInclusionForeignPaymentAmount(t *testing.T) {
	testAddenda10FieldInclusionForeignPaymentAmount(t)
}

// BenchmarkAddenda10FieldInclusionForeignPaymentAmount benchmarks validating ForeignPaymentAmount fieldInclusion
func BenchmarkAddenda10FieldInclusionForeignPaymentAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda10FieldInclusionForeignPaymentAmount(b)
	}
}

// testAddenda10FieldInclusionName validates Name fieldInclusion
func testAddenda10FieldInclusionName(t testing.TB) {
	addenda10 := mockAddenda10()
	addenda10.Name = ""
	err := addenda10.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda10FieldInclusionName tests validating Name fieldInclusion
func TestAddenda10FieldInclusionName(t *testing.T) {
	testAddenda10FieldInclusionName(t)
}

// BenchmarkAddenda10FieldInclusionName benchmarks validating Name fieldInclusion
func BenchmarkAddenda10FieldInclusionName(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda10FieldInclusionName(b)
	}
}

// testAddenda10FieldInclusionEntryDetailSequenceNumber validates EntryDetailSequenceNumber fieldInclusion
func testAddenda10FieldInclusionEntryDetailSequenceNumber(t testing.TB) {
	addenda10 := mockAddenda10()
	addenda10.EntryDetailSequenceNumber = 0
	err := addenda10.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda10FieldInclusionEntryDetailSequenceNumber tests validating
// EntryDetailSequenceNumber fieldInclusion
func TestAddenda10FieldInclusionEntryDetailSequenceNumber(t *testing.T) {
	testAddenda10FieldInclusionEntryDetailSequenceNumber(t)
}

// BenchmarkAddenda10FieldInclusionEntryDetailSequenceNumber benchmarks validating
// EntryDetailSequenceNumber fieldInclusion
func BenchmarkAddenda10FieldInclusionEntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda10FieldInclusionEntryDetailSequenceNumber(b)
	}
}

// TestAddenda10String validates that a known parsed Addenda10 record can be return to a string of the same value
func testAddenda10String(t testing.TB) {
	addenda10 := NewAddenda10()
	var line = "710ANN000000000000100000928383-23938          BEK Enterprises                          0000001"
	addenda10.Parse(line)

	if addenda10.String() != line {
		t.Errorf("Strings do not match")
	}
	if addenda10.TypeCode != "10" {
		t.Errorf("TypeCode Expected 10 got: %v", addenda10.TypeCode)
	}
}

// TestAddenda10String tests validating that a known parsed Addenda10 record can be return to a string of the same value
func TestAddenda10String(t *testing.T) {
	testAddenda10String(t)
}

// BenchmarkAddenda10String benchmarks validating that a known parsed Addenda10 record can be return to a string of the same value
func BenchmarkAddenda10String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda10String(b)
	}
}

// TestAddenda10RuneCountInString validates RuneCountInString
func TestAddenda10RuneCountInString(t *testing.T) {
	addenda10 := NewAddenda10()
	var line = "710ANN000000000000100000928383-23938          BEK Enterprises"
	addenda10.Parse(line)

	if addenda10.TransactionTypeCode != "" {
		t.Error("Parsed with an invalid RuneCountInString not equal to 94")
	}
}
