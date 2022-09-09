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

func TestAddenda98_LookupChangecode(t *testing.T) {
	if code := LookupChangeCode(""); code != nil {
		t.Error("expected nil ChangeCode")
	}
	if code := LookupChangeCode("C05"); code == nil {
		t.Error("expected ChangeCode")
	} else {
		if code.Code != "C05" {
			t.Errorf("code.Code=%s", code.Code)
		}
		if code.Reason != "Incorrect payment code" {
			t.Errorf("code.Reason=%s", code.Reason)
		}
	}
	if code := LookupChangeCode("C64"); code == nil {
		t.Error("expected ChangeCode")
	}
	if code := LookupChangeCode("C99"); code != nil {
		t.Errorf("expected nil: %#v", code)
	}
}

func testAddenda98Parse(t testing.TB) {
	addenda98 := NewAddenda98()
	line := "798C01099912340000015      091012981918171614                                  091012980000088"
	addenda98.Parse(line)
	// walk the Addenda98 struct
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
	addenda98.ChangeCode = "C55"
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

func TestAddenda98__ChangeCodeField(t *testing.T) {
	addenda98 := mockAddenda98()
	if addenda98.ChangeCode != "C01" {
		t.Errorf("addenda98.ChangeCode=%s", addenda98.ChangeCode)
	}
	if code := addenda98.ChangeCodeField(); code == nil {
		t.Fatal("nil Addenda98.ChangeCodeField")
	} else {
		if code.Code != "C01" {
			t.Errorf("code.Code=%s", code.Code)
		}
		if code.Reason != "Incorrect bank account number" {
			t.Errorf("code.Reason=%s", code.Reason)
		}
	}

	// verify another change code
	addenda98.ChangeCode = "C07"
	if code := addenda98.ChangeCodeField(); code == nil {
		t.Fatal("nil Addenda98.ChangeCodeField")
	} else {
		if code.Code != "C07" {
			t.Errorf("code.Code=%s", code.Code)
		}
	}

	// invalid change code
	addenda98.ChangeCode = "C99"
	if code := addenda98.ChangeCodeField(); code != nil {
		t.Errorf("unexpected change code: %v", code)
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

func TestCorrectedData__first(t *testing.T) {
	if v := first(2, ""); v != "" {
		t.Errorf("got='%s'", v)
	}
	if v := first(2, "    "); v != "" {
		t.Errorf("got='%s'", v)
	}
	if v := first(3, "22"); v != "22" {
		t.Errorf("got='%s'", v)
	}
	if v := first(17, "   123   "); v != "123" {
		t.Errorf("got='%s'", v)
	}
	if v := first(17, "123456789   "); v != "123456789" {
		t.Errorf("got='%s'", v)
	}
}

func TestCorrectedData__ParseCorrectedData(t *testing.T) {
	run := func(code, data string) *CorrectedData {
		add := NewAddenda98()
		add.ChangeCode = code
		add.CorrectedData = data
		return add.ParseCorrectedData()
	}

	if v := run("C01", "123456789       "); v.AccountNumber != "123456789" {
		t.Errorf("%#v", v)
	}
	if v := run("C02", "987654320  "); v.RoutingNumber != "987654320" {
		t.Errorf("%#v", v)
	}
	if v := run("C03", "987654320   123456"); v.AccountNumber != "123456" || v.RoutingNumber != "987654320" {
		t.Errorf("%#v", v)
	}
	if v := run("C04", "Jane Doe"); v.Name != "Jane Doe" {
		t.Errorf("%#v", v)
	}
	if v := run("C05", "22  other"); v.TransactionCode != 22 {
		t.Errorf("%#v", v)
	}
	if v := run("C06", "123456789                22"); v.AccountNumber != "123456789" || v.TransactionCode != 22 {
		t.Errorf("%#v", v)
	}
	if v := run("C07", "987654320  12345  22"); v.RoutingNumber != "987654320" || v.AccountNumber != "12345" || v.TransactionCode != 22 {
		t.Errorf("%#v", v)
	}
	if v := run("C07", "9876543201242415    22"); v.RoutingNumber != "987654320" || v.AccountNumber != "1242415" || v.TransactionCode != 22 {
		t.Errorf("%#v", v)
	}
	if v := run("C07", "1234"); v != nil {
		t.Errorf("expected nil: %v", v)
	}
	if v := run("C07", "987654320 1234 1234 1234"); v != nil {
		t.Errorf("expected nil: %v", v)
	}
	if v := run("C09", "21345678    "); v.Identification != "21345678" {
		t.Errorf("%#v", v)
	}
	if v := run("C99", "    "); v != nil {
		t.Error("expected nil CorrectedData")
	}
}

func TestCorrectedData__WriteCorrectionData(t *testing.T) {
	data := &CorrectedData{AccountNumber: "12345"}
	if v := WriteCorrectionData("C01", data); v != "12345                 " {
		t.Errorf("C01 got %q (length=%d)", v, len(v))
	}
	data = &CorrectedData{RoutingNumber: "987654320"}
	if v := WriteCorrectionData("C02", data); v != "987654320             " {
		t.Errorf("C02 got %q (length=%d)", v, len(v))
	}
	data = &CorrectedData{AccountNumber: "123", RoutingNumber: "987654320"}
	if v := WriteCorrectionData("C03", data); v != "987654320          123" {
		t.Errorf("C03 got %q (length=%d)", v, len(v))
	}
	data = &CorrectedData{Name: "Jane Doe"}
	if v := WriteCorrectionData("C04", data); v != "Jane Doe              " {
		t.Errorf("C04 got %q (length=%d)", v, len(v))
	}
	data = &CorrectedData{TransactionCode: 22}
	if v := WriteCorrectionData("C05", data); v != "22                    " {
		t.Errorf("C05 got %q (length=%d)", v, len(v))
	}
	data = &CorrectedData{AccountNumber: "5421", TransactionCode: 27}
	if v := WriteCorrectionData("C06", data); v != "5421                27" {
		t.Errorf("C06 got %q (length=%d)", v, len(v))
	}
	data = &CorrectedData{RoutingNumber: "987654320", AccountNumber: "5421", TransactionCode: 32}
	if v := WriteCorrectionData("C07", data); v != "9876543205421       32" {
		t.Errorf("C07 got %q (length=%d)", v, len(v))
	}
	data = &CorrectedData{Identification: "FooBar"}
	if v := WriteCorrectionData("C09", data); v != "FooBar                " {
		t.Errorf("C09 got %q (length=%d)", v, len(v))
	}
}
