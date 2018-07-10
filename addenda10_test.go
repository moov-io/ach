// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"
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

// testAddenda10ValidRecordType validates Addenda10 recordType
func testAddenda10ValidRecordType(t testing.TB) {
	addenda10 := mockAddenda10()
	addenda10.recordType = "63"
	if err := addenda10.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddenda10ValidRecordType tests validating Addenda10 recordType
func TestAddenda10ValidRecordType(t *testing.T) {
	testAddenda10ValidRecordType(t)
}

// BenchmarkAddenda10ValidRecordType benchmarks validating Addenda10 recordType
func BenchmarkAddenda10ValidRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda10ValidRecordType(b)
	}
}

// testAddenda10ValidTypeCode validates Addenda10 TypeCode
func testAddenda10ValidTypeCode(t testing.TB) {
	addenda10 := mockAddenda10()
	addenda10.typeCode = "65"
	if err := addenda10.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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

// testAddenda10TypeCode10 TypeCode is 10 if typeCode is a valid TypeCode
func testAddenda10TypeCode10(t testing.TB) {
	addenda10 := mockAddenda10()
	addenda10.typeCode = "05"
	if err := addenda10.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestAddenda10TypeCode10 tests TypeCode is 10 if typeCode is a valid TypeCode
func TestAddenda10TypeCode10(t *testing.T) {
	testAddenda10TypeCode10(t)
}

// BenchmarkAddenda10TypeCode10 benchmarks TypeCode is 10 if typeCode is a valid TypeCode
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
	if err := addenda10.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TransactionTypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
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
	if err := addenda10.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ForeignTraceNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
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
	if err := addenda10.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "Name" {
				t.Errorf("%T: %s", err, err)
			}
		}
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

// testAddenda10FieldInclusionRecordType validates recordType fieldInclusion
func testAddenda10FieldInclusionRecordType(t testing.TB) {
	addenda10 := mockAddenda10()
	addenda10.recordType = ""
	if err := addenda10.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestAddenda10FieldInclusionRecordType tests validating recordType fieldInclusion
func TestAddenda10FieldInclusionRecordType(t *testing.T) {
	testAddenda10FieldInclusionRecordType(t)
}

// BenchmarkAddenda10FieldInclusionRecordType benchmarks validating recordType fieldInclusion
func BenchmarkAddenda10FieldInclusionRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda10FieldInclusionRecordType(b)
	}
}

// testAddenda10FieldInclusionTypeCode validates TypeCode fieldInclusion
func testAddenda10FieldInclusionTypeCode(t testing.TB) {
	addenda10 := mockAddenda10()
	addenda10.typeCode = ""
	if err := addenda10.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
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
	if err := addenda10.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldRequired {
				t.Errorf("%T: %s", err, err)
			}
		}
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
	if err := addenda10.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldRequired {
				t.Errorf("%T: %s", err, err)
			}
		}
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
	if err := addenda10.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
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
	if err := addenda10.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
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

// ToDo  Add Parse test for individual fields

// TestAddenda10String validates that a known parsed Addenda10 record can be return to a string of the same value
func testAddenda10String(t testing.TB) {
	addenda10 := NewAddenda10()
	var line = "710ANN000000000000100000928383-23938           BEK Enterprises                         0000001"
	addenda10.Parse(line)

	if addenda10.String() != line {
		t.Errorf("Strings do not match")
	}
	if addenda10.TypeCode() != "10" {
		t.Errorf("TypeCode Expected 10 got: %v", addenda10.TypeCode())
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
