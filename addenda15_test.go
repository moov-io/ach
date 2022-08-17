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

// mockAddenda15 creates a mock Addenda15 record
func mockAddenda15() *Addenda15 {
	addenda15 := NewAddenda15()
	addenda15.ReceiverIDNumber = "987465493213987"
	addenda15.ReceiverStreetAddress = "2121 Front Street"
	addenda15.EntryDetailSequenceNumber = 00000001
	return addenda15
}

// TestMockAddenda15 validates mockAddenda15
func TestMockAddenda15(t *testing.T) {
	addenda15 := mockAddenda15()
	if err := addenda15.Validate(); err != nil {
		t.Error("mockAddenda15 does not validate and will break other tests")
	}
}

// testAddenda15Parse parses Addenda15 record
func testAddenda15Parse(t testing.TB) {
	Addenda15 := NewAddenda15()
	line := "7159874654932139872121 Front Street                                                    0000001"
	Addenda15.Parse(line)
	// walk the Addenda15 struct
	if Addenda15.TypeCode != "15" {
		t.Errorf("expected %v got %v", "15", Addenda15.TypeCode)
	}
	if Addenda15.ReceiverIDNumber != "987465493213987" {
		t.Errorf("expected %v got %v", "987465493213987", Addenda15.ReceiverIDNumber)
	}
	if Addenda15.ReceiverStreetAddress != "2121 Front Street" {
		t.Errorf("expected: %v got: %v", "2121 Front Street", Addenda15.ReceiverStreetAddress)
	}
	if Addenda15.EntryDetailSequenceNumber != 0000001 {
		t.Errorf("expected: %v got: %v", 0000001, Addenda15.EntryDetailSequenceNumber)
	}
}

// TestAddenda15Parse tests parsing Addenda15 record
func TestAddenda15Parse(t *testing.T) {
	testAddenda15Parse(t)
}

// BenchmarkAddenda15Parse benchmarks parsing Addenda15 record
func BenchmarkAddenda15Parse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda15Parse(b)
	}
}

// testAddenda15ValidTypeCode validates Addenda15 TypeCode
func testAddenda15ValidTypeCode(t testing.TB) {
	addenda15 := mockAddenda15()
	addenda15.TypeCode = "65"
	err := addenda15.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda15ValidTypeCode tests validating Addenda15 TypeCode
func TestAddenda15ValidTypeCode(t *testing.T) {
	testAddenda15ValidTypeCode(t)
}

// BenchmarkAddenda15ValidTypeCode benchmarks validating Addenda15 TypeCode
func BenchmarkAddenda15ValidTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda15ValidTypeCode(b)
	}
}

// testAddenda15TypeCode15 TypeCode is 15 if TypeCode is a valid TypeCode
func testAddenda15TypeCode15(t testing.TB) {
	addenda15 := mockAddenda15()
	addenda15.TypeCode = "05"
	err := addenda15.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda15TypeCode15 tests TypeCode is 15 if TypeCode is a valid TypeCode
func TestAddenda15TypeCode15(t *testing.T) {
	testAddenda15TypeCode15(t)
}

// BenchmarkAddenda15TypeCode15 benchmarks TypeCode is 15 if TypeCode is a valid TypeCode
func BenchmarkAddenda15TypeCode15(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda15TypeCode15(b)
	}
}

// testReceiverIDNumberAlphaNumeric validates ReceiverIDNumber is alphanumeric
func testReceiverIDNumberAlphaNumeric(t testing.TB) {
	addenda15 := mockAddenda15()
	addenda15.ReceiverIDNumber = "9874654932®1398"
	err := addenda15.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestReceiverIDNumberAlphaNumeric tests validating ReceiverIDNumber is alphanumeric
func TestReceiverIDNumberAlphaNumeric(t *testing.T) {
	testReceiverIDNumberAlphaNumeric(t)
}

// BenchmarkReceiverIDNumberAlphaNumeric benchmarks validating ReceiverIDNumber is alphanumeric
func BenchmarkReceiverIDNumberAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testReceiverIDNumberAlphaNumeric(b)
	}
}

// testReceiverStreetAddressAlphaNumeric validates ReceiverStreetAddress is alphanumeric
func testReceiverStreetAddressAlphaNumeric(t testing.TB) {
	addenda15 := mockAddenda15()
	addenda15.ReceiverStreetAddress = "2121 Fr®nt Street"
	err := addenda15.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestReceiverStreetAddressAlphaNumeric tests validating ReceiverStreetAddress is alphanumeric
func TestReceiverStreetAddressAlphaNumeric(t *testing.T) {
	testReceiverStreetAddressAlphaNumeric(t)
}

// BenchmarkReceiverStreetAddressAlphaNumeric benchmarks validating ReceiverStreetAddress is alphanumeric
func BenchmarkReceiverStreetAddressAlphaNumeric(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testReceiverStreetAddressAlphaNumeric(b)
	}
}

// testAddenda15FieldInclusionTypeCode validates TypeCode fieldInclusion
func testAddenda15FieldInclusionTypeCode(t testing.TB) {
	addenda15 := mockAddenda15()
	addenda15.TypeCode = ""
	err := addenda15.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda15FieldInclusionTypeCode tests validating TypeCode fieldInclusion
func TestAddenda15FieldInclusionTypeCode(t *testing.T) {
	testAddenda15FieldInclusionTypeCode(t)
}

// BenchmarkAddenda15FieldInclusionTypeCode benchmarks validating TypeCode fieldInclusion
func BenchmarkAddenda15FieldInclusionTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda15FieldInclusionTypeCode(b)
	}
}

// testAddenda15FieldInclusionReceiverStreetAddress validates ReceiverStreetAddress fieldInclusion
func testAddenda15FieldInclusionReceiverStreetAddress(t testing.TB) {
	addenda15 := mockAddenda15()
	addenda15.ReceiverStreetAddress = ""
	err := addenda15.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda15FieldInclusionReceiverStreetAddress tests validating ReceiverStreetAddress fieldInclusion
func TestAddenda15FieldInclusionReceiverStreetAddress(t *testing.T) {
	testAddenda15FieldInclusionReceiverStreetAddress(t)
}

// BenchmarkAddenda15FieldInclusionReceiverStreetAddress benchmarks validating ReceiverStreetAddress fieldInclusion
func BenchmarkAddenda15FieldInclusionReceiverStreetAddress(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda15FieldInclusionReceiverStreetAddress(b)
	}
}

// testAddenda15FieldInclusionEntryDetailSequenceNumber validates EntryDetailSequenceNumber fieldInclusion
func testAddenda15FieldInclusionEntryDetailSequenceNumber(t testing.TB) {
	addenda15 := mockAddenda15()
	addenda15.EntryDetailSequenceNumber = 0
	err := addenda15.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda15FieldInclusionEntryDetailSequenceNumber tests validating
// EntryDetailSequenceNumber fieldInclusion
func TestAddenda15FieldInclusionEntryDetailSequenceNumber(t *testing.T) {
	testAddenda15FieldInclusionEntryDetailSequenceNumber(t)
}

// BenchmarkAddenda15FieldInclusionEntryDetailSequenceNumber benchmarks validating
// EntryDetailSequenceNumber fieldInclusion
func BenchmarkAddenda15FieldInclusionEntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda15FieldInclusionEntryDetailSequenceNumber(b)
	}
}

// TestAddenda15String validates that a known parsed Addenda15 record can be return to a string of the same value
func testAddenda15String(t testing.TB) {
	addenda15 := NewAddenda15()
	var line = "7159874654932139872121 Front Street                                                    0000001"
	addenda15.Parse(line)

	if addenda15.String() != line {
		t.Errorf("Strings do not match")
	}
	if addenda15.TypeCode != "15" {
		t.Errorf("TypeCode Expected 15 got: %v", addenda15.TypeCode)
	}
}

// TestAddenda15String tests validating that a known parsed Addenda15 record can be return to a string of the same value
func TestAddenda15String(t *testing.T) {
	testAddenda15String(t)
}

// BenchmarkAddenda15String benchmarks validating that a known parsed Addenda15 record can be return to a string of the same value
func BenchmarkAddenda15String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda15String(b)
	}
}

// TestAddenda15RuneCountInString validates RuneCountInString
func TestAddenda15RuneCountInString(t *testing.T) {
	addenda15 := NewAddenda15()
	var line = "7159874654932139872121 Front Street"
	addenda15.Parse(line)

	if addenda15.ReceiverIDNumber != "" {
		t.Error("Parsed with an invalid RuneCountInString not equal to 94")
	}
}
