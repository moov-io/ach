// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "testing"

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

// testAddenda10FieldInclusionRecordType validates TypeCode fieldInclusion
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

// TestAddenda10FieldInclusionRecordType tests validating TypeCode fieldInclusion
func TestAddenda10FieldInclusionTypeCode(t *testing.T) {
	testAddenda10FieldInclusionTypeCode(t)
}

// BenchmarkAddenda10FieldInclusionRecordType benchmarks validating TypeCode fieldInclusion
func BenchmarkAddenda10FieldInclusionTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda10FieldInclusionTypeCode(b)
	}
}
