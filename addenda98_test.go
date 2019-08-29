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

func mockAddenda98() *Addenda98 {
	addenda98 := NewAddenda98()
	addenda98.ChangeCode = "C01"
	addenda98.OriginalTrace = "12345"
	addenda98.OriginalDFI = "9101298"
	addenda98.CorrectedData = "1918171614"
	addenda98.TraceNumber = "91012980000088"

	return addenda98
}

func testAddenda98Parse(t testing.TB) {
	addenda98 := NewAddenda98()
	line := "798C01099912340000015      091012981918171614                                  091012980000088"
	addenda98.Parse(line)
	// walk the Addenda98 struct
	if addenda98.recordType != "7" {
		t.Errorf("expected %v got %v", "7", addenda98.recordType)
	}
	if addenda98.TypeCode != "98" {
		t.Errorf("expected %v got %v", "98", addenda98.TypeCode)
	}
	if addenda98.ChangeCode != "C01" {
		t.Errorf("expected %v got %v", "C01", addenda98.ChangeCode)
	}
	if addenda98.OriginalTrace != "099912340000015" {
		t.Errorf("expected %v got %v", "099912340000015", addenda98.OriginalTrace)
	}
	if addenda98.OriginalDFI != "09101298" {
		t.Errorf("expected %s got %s", "09101298", addenda98.OriginalDFI)
	}
	if addenda98.CorrectedData != "1918171614" {
		t.Errorf("expected %v got %v", "1918171614", addenda98.CorrectedData)
	}
	if addenda98.TraceNumber != "091012980000088" {
		t.Errorf("expected %v got %v", "091012980000088", addenda98.TraceNumber)
	}
}

func TestAddenda98Parse(t *testing.T) {
	testAddenda98Parse(t)
}

func BenchmarkAddenda98Parse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98Parse(b)
	}
}

func testAddenda98String(t testing.TB) {
	addenda98 := NewAddenda98()
	line := "798C01099912340000015      091012981918171614                                  091012980000088"
	addenda98.Parse(line)

	if addenda98.String() != line {
		t.Errorf("\n expected: %v\n got     : %v", line, addenda98.String())
	}
}

func TestAddenda98String(t *testing.T) {
	testAddenda98String(t)
}

func BenchmarkAddenda98String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98String(b)
	}
}

func testAddenda98ValidRecordType(t testing.TB) {
	addenda98 := mockAddenda98()
	addenda98.recordType = "63"
	err := addenda98.Validate()
	if !base.Match(err, NewErrRecordType(7)) {
		t.Errorf("%T: %s", err, err)
	}
}
func TestAddenda98ValidRecordType(t *testing.T) {
	testAddenda98ValidRecordType(t)
}

func BenchmarkAddenda98ValidRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98ValidRecordType(b)
	}
}

func testAddenda98ValidTypeCode(t testing.TB) {
	addenda98 := mockAddenda98()
	addenda98.TypeCode = "05"
	err := addenda98.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda98ValidTypeCode(t *testing.T) {
	testAddenda98ValidTypeCode(t)
}

func BenchmarkAddenda98ValidTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98ValidTypeCode(b)
	}
}

func testAddenda98ValidCorrectedData(t testing.TB) {
	addenda98 := mockAddenda98()
	addenda98.CorrectedData = ""
	err := addenda98.Validate()
	if !base.Match(err, ErrAddenda98CorrectedData) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda98ValidCorrectedData(t *testing.T) {
	testAddenda98ValidCorrectedData(t)
}

func BenchmarkAddenda98ValidCorrectedData(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98ValidCorrectedData(b)
	}
}

func testAddenda98ValidateTrue(t testing.TB) {
	addenda98 := mockAddenda98()
	addenda98.ChangeCode = "C11"
	err := addenda98.Validate()
	// no error expected
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda98ValidateTrue(t *testing.T) {
	testAddenda98ValidateTrue(t)
}

func BenchmarkAddenda98ValidateTrue(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98ValidateTrue(b)
	}
}

func testAddenda98ValidateChangeCodeFalse(t testing.TB) {
	addenda98 := mockAddenda98()
	addenda98.ChangeCode = "C63"
	err := addenda98.Validate()
	if !base.Match(err, ErrAddenda98ChangeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestAddenda98ValidateChangeCodeFalse(t *testing.T) {
	testAddenda98ValidateChangeCodeFalse(t)
}

func BenchmarkAddenda98ValidateChangeCodeFalse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98ValidateChangeCodeFalse(b)
	}
}

func testAddenda98OriginalTraceField(t testing.TB) {
	addenda98 := mockAddenda98()
	exp := "000000000012345"
	if addenda98.OriginalTraceField() != exp {
		t.Errorf("expected %v received %v", exp, addenda98.OriginalTraceField())
	}
}

func TestAddenda98OriginalTraceField(t *testing.T) {
	testAddenda98OriginalTraceField(t)
}

func BenchmarkAddenda98OriginalTraceField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98OriginalTraceField(b)
	}
}

func testAddenda98OriginalDFIField(t testing.TB) {
	addenda98 := mockAddenda98()
	exp := "09101298"
	if addenda98.OriginalDFIField() != exp {
		t.Errorf("expected %v received %v", exp, addenda98.OriginalDFIField())
	}
}

func TestAddenda98OriginalDFIField(t *testing.T) {
	testAddenda98OriginalDFIField(t)
}

func BenchmarkAddenda98OriginalDFIField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98OriginalDFIField(b)
	}
}

func testAddenda98CorrectedDataField(t testing.TB) {
	addenda98 := mockAddenda98()
	exp := "1918171614                   " // 29 char
	if addenda98.CorrectedDataField() != exp {
		t.Errorf("expected %v received %v", exp, addenda98.CorrectedDataField())
	}
}

func TestAddenda98CorrectedDataField(t *testing.T) {
	testAddenda98CorrectedDataField(t)
}

func BenchmarkAddenda98CorrectedDataField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98CorrectedDataField(b)
	}
}

func testAddenda98TraceNumberField(t testing.TB) {
	addenda98 := mockAddenda98()
	exp := "091012980000088"
	if addenda98.TraceNumberField() != exp {
		t.Errorf("expected %v received %v", exp, addenda98.TraceNumberField())
	}
}

func TestAddenda98TraceNumberField(t *testing.T) {
	testAddenda98TraceNumberField(t)
}

func BenchmarkAddenda98TraceNumberField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98TraceNumberField(b)
	}
}

// testAddenda98TypeCodeNil validates TypeCode is ""
func testAddenda98TypeCodeNil(t testing.TB) {
	addenda98 := mockAddenda98()
	addenda98.TypeCode = ""
	err := addenda98.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda98TypeCodeES tests TypeCode is ""
func TestAddenda98TypeCodeNil(t *testing.T) {
	testAddenda98TypeCodeNil(t)
}

// BenchmarkAddenda98TypeCodeNil benchmarks TypeCode is ""
func BenchmarkAddenda98TypeCodeNil(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98TypeCodeNil(b)
	}
}

// TestAddenda98RuneCountInString validates RuneCountInString
func TestAddenda98RuneCountInString(t *testing.T) {
	addenda98 := NewAddenda98()
	var line = "798"
	addenda98.Parse(line)

	if addenda98.CorrectedData != "" {
		t.Error("Parsed with an invalid RuneCountInString not equal to 94")
	}
}
