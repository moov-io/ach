// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "testing"

//testAlphaField ensire that padding and two long of strings get properly made
func TestAlphaFieldShort(t *testing.T) {
	c := Converters{}
	result := c.alphaField("ABC123", 10)
	if result != "ABC123    " {
		t.Errorf("Left justified space filled got:'%v'", result)
	}
}

func TestAlphaFieldLong(t *testing.T) {
	c := Converters{}
	result := c.alphaField("abcdEFGH123", 10)
	if result != "abcdEFGH12" {
		t.Errorf("Left justified space filled got:'%v'", result)
	}
}

func TestNumericFieldShort(t *testing.T) {
	c := Converters{}
	result := c.numericField(12345, 10)
	if result != "0000012345" {
		t.Errorf("Right justified zero got: '%v'", result)
	}
}

func TestNumericFieldLong(t *testing.T) {
	c := Converters{}
	result := c.numericField(123456, 5)
	if result != "23456" {
		t.Errorf("Right justified zero got: '%v'", result)
	}
}
