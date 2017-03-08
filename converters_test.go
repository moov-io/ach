// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "testing"

//testAlphaField ensire that padding and two long of strings get properly made
func TestAlphaFieldShort(t *testing.T) {
	c := converters{}
	result := c.alphaField("ABC123", 10)
	if result != "ABC123    " {
		t.Errorf("Left justified space filled got:'%v'", result)
	}
}

// TestAlphaFieldLong ensure that string is left justified and sliced to max
func TestAlphaFieldLong(t *testing.T) {
	c := converters{}
	result := c.alphaField("abcdEFGH123", 10)
	if result != "abcdEFGH12" {
		t.Errorf("Left justified space filled got:'%v'", result)
	}
}

// TestNumericFieldShort ensures zero padding and right justified
func TestNumericFieldShort(t *testing.T) {
	c := converters{}
	result := c.numericField(12345, 10)
	if result != "0000012345" {
		t.Errorf("Right justified zero got: '%v'", result)
	}
}

// TestNumericFieldLong right justified and sliced to max length
func TestNumericFieldLong(t *testing.T) {
	c := converters{}
	result := c.numericField(123456, 5)
	if result != "23456" {
		t.Errorf("Right justified zero got: '%v'", result)
	}
}

//TestParseNumField handle zero and spaces in number conversion
func TestParseNumField(t *testing.T) {
	c := converters{}
	result := c.parseNumField(" 012345")
	if result != 12345 {
		t.Errorf("Right justified zero got: '%v'", result)
	}
}
