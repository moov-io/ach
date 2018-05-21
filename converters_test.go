// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "testing"

//testAlphaField ensures that padding and two long of strings get properly made
func testAlphaFieldShort(t testing.TB) {
	c := converters{}
	result := c.alphaField("ABC123", 10)
	if result != "ABC123    " {
		t.Errorf("Left justified space filled got:'%v'", result)
	}
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

// testRTNFieldShort ensures zero padding and right justified
func testRTNFieldShort(t testing.TB) {
	c := converters{}
	result := c.stringRTNField("123456", 8)
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
	result := c.stringRTNField("1234567899", 8)
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
	result := c.stringRTNField("123456789", 9)
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
