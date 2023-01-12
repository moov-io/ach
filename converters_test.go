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
	"unicode/utf8"

	"github.com/stretchr/testify/require"
)

// testAlphaField ensures that padding and two long of strings get properly made
func testAlphaFieldShort(t testing.TB) {
	c := converters{}
	result := c.alphaField("ABC123", 10)
	if result != "ABC123    " {
		t.Errorf("Left justified space filled got:'%v'", result)
	}

	answer := c.alphaField("Returned per ODFI’s Request", 44)
	expected := "Returned per ODFI’s Request                 "
	require.Equal(t, 44, utf8.RuneCountInString(expected))
	require.Equal(t, expected, answer)
}

// TestAlphaFieldShort test ensures that padding and two long of strings get properly made
func TestAlphaFieldShort(t *testing.T) {
	testAlphaFieldShort(t)
}

// BenchmarkAlphaFieldShort benchmark ensures that padding and two long of strings get properly made
func BenchmarkAlphaFieldShort(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAlphaFieldShort(b)
	}
}

// testAlphaFieldLong ensures that string is left justified and sliced to max
func testAlphaFieldLong(t testing.TB) {
	c := converters{}
	result := c.alphaField("abcdEFGH123", 10)
	if result != "abcdEFGH12" {
		t.Errorf("Left justified space filled got:'%v'", result)
	}
}

// TestAlphaFieldLong test ensures that string is left justified and sliced to max
func TestAlphaFieldLong(t *testing.T) {
	testAlphaFieldLong(t)
}

// BenchmarkAlphaFieldLong benchmark ensures that string is left justified and sliced to max
func BenchmarkAlphaFieldLong(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAlphaFieldLong(b)
	}
}

// testNumericFieldShort ensures zero padding and right justified
func testNumericFieldShort(t testing.TB) {
	c := converters{}
	result := c.numericField(12345, 10)
	if result != "0000012345" {
		t.Errorf("Right justified zero got: '%v'", result)
	}
}

// TestNumericFieldShort test ensures zero padding and right justified
func TestNumericFieldShort(t *testing.T) {
	testNumericFieldShort(t)
}

// BenchmarkNumericFieldShort benchmark ensures zero padding and right justified
func BenchmarkNumericFieldShort(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testNumericFieldShort(b)
	}
}

// testNumericFieldLong ensures right justified and sliced to max length
func testNumericFieldLong(t testing.TB) {
	c := converters{}
	result := c.numericField(123456, 5)
	if result != "23456" {
		t.Errorf("Right justified zero got: '%v'", result)
	}
}

// TestNumericFieldLong test ensures right justified and sliced to max length
func TestNumericFieldLong(t *testing.T) {
	testNumericFieldLong(t)
}

// BenchmarkNumericFieldLong benchmark ensures right justified and sliced to max length
func BenchmarkNumericFieldLong(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testNumericFieldLong(b)
	}
}

// testParseNumField handles zero and spaces in number conversion
func testParseNumField(t testing.TB) {
	c := converters{}
	result := c.parseNumField(" 012345")
	if result != 12345 {
		t.Errorf("Right justified zero got: '%v'", result)
	}
}

// TestParseNumField test handles zero and spaces in number conversion
func TestParseNumField(t *testing.T) {
	testParseNumField(t)
}

// BenchmarkParseNumField benchmark handles zero and spaces in number conversion
func BenchmarkParseNumField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testParseNumField(b)
	}
}

// testParseStringField handles spaces in string conversion
func testParseStringField(t testing.TB) {
	c := converters{}
	result := c.parseStringField(" 012345")
	if result != "012345" {
		t.Errorf("Trim spaces: '%v'", result)
	}
}

// TestParseStringField test handles spaces in string conversion
func TestParseStringField(t *testing.T) {
	testParseStringField(t)
}

// BenchmarkParseStringField benchmark handles spaces in string conversion
func BenchmarkParseStringField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testParseStringField(b)
	}
}

func TestParseStringFieldWithOpts__PreservesSpacesWithOptTrue(t *testing.T) {
	c := converters{}
	opts := ValidateOpts{
		PreserveSpaces: true,
	}

	expected := " 012345"

	result := c.parseStringFieldWithOpts(" 012345", &opts)
	if result != expected {
		t.Errorf("Expected: '%v', Actual: '%v'", expected, result)
	}
}

func TestParseStringFieldWithOpts__DoesntPreserveSpacesWithOptFalse(t *testing.T) {
	c := converters{}
	opts := ValidateOpts{
		PreserveSpaces: false,
	}

	expected := "012345"

	result := c.parseStringFieldWithOpts(" 012345", &opts)
	if result != expected {
		t.Errorf("Expected: '%v', Actual: '%v'", expected, result)
	}
}

// testRTNFieldShort ensures zero padding and right justified
func testRTNFieldShort(t testing.TB) {
	c := converters{}
	result := c.stringField("123456", 8)
	if result != "00123456" {
		t.Errorf("Zero padding 8 character string : '%v'", result)
	}
}

// TestRTNFieldShort test ensures zero padding and right justified
func TestRTNFieldShort(t *testing.T) {
	testRTNFieldShort(t)
}

// BenchmarkRTNFieldShort benchmark ensures zero padding and right justified
func BenchmarkRTNFieldShort(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testRTNFieldShort(b)
	}
}

// testRTNFieldLong ensures sliced to max length
func testRTNFieldLong(t testing.TB) {
	c := converters{}
	result := c.stringField("1234567899", 8)
	if result != "12345678" {
		t.Errorf("first 8 character string: '%v'", result)
	}
}

// TestRTNFieldLong test ensures  sliced to max length
func TestRTNFieldLong(t *testing.T) {
	testRTNFieldLong(t)
}

// BenchmarkRTNFieldLong benchmark ensures sliced to max length
func BenchmarkRTNFieldLong(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testRTNFieldLong(b)
	}
}

// testRTNFieldExact ensures exact match
func testRTNFieldExact(t testing.TB) {
	c := converters{}
	result := c.stringField("123456789", 9)
	if result != "123456789" {
		t.Errorf("first 9 character string: '%v'", result)
	}
}

// TestRTNFieldExact test ensures exact match
func TestRTNFieldExact(t *testing.T) {
	testRTNFieldExact(t)
}

// BenchmarkRTNFieldExact benchmark ensures exact match
func BenchmarkRTNFieldExact(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testRTNFieldExact(b)
	}
}

func TestLeastSignificantDigits(t *testing.T) {
	tests := []struct {
		input int
		max   uint
		want  int
	}{
		{
			input: 123,
			max:   2,
			want:  23,
		},
		{
			input: 123,
			max:   3,
			want:  123,
		},
		{
			input: 123,
			max:   5,
			want:  123,
		},
		{
			input: 12345678912,
			max:   10,
			want:  2345678912,
		},
		{
			input: 99,
			max:   0,
			want:  0,
		},
	}

	c := converters{}
	for _, tt := range tests {
		if got := c.leastSignificantDigits(tt.input, tt.max); got != tt.want {
			t.Errorf("rightmost digits: want %d, got %d", tt.want, got)
		}
	}
}
