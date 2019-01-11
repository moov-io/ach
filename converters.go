// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strconv"
	"strings"
)

// converters handles golang to ACH type Converters
type converters struct{}

func (c *converters) parseNumField(r string) (s int) {
	s, _ = strconv.Atoi(strings.TrimSpace(r))
	return s
}

func (c *converters) parseStringField(r string) (s string) {
	s = strings.TrimSpace(r)
	return s
}

// formatSimpleDate takes a YYMMDD date and formats it for the fixed-width ACH file format
func (c *converters) formatSimpleDate(s string) string {
	if s == "" {
		return c.stringField(s, 6)
	}
	return s
}

// formatSimpleTime takes a HHMM time and formats it for the fixed-width ACH file format
func (c *converters) formatSimpleTime(s string) string {
	if s == "" {
		return c.stringField(s, 4)
	}
	return s
}

// alphaField Alphanumeric and Alphabetic fields are left-justified and space filled.
func (c *converters) alphaField(s string, max uint) string {
	ln := uint(len(s))
	if ln > max {
		return s[:max]
	}
	s += strings.Repeat(" ", int(max-ln))
	return s
}

// numericField right-justified, unsigned, and zero filled
func (c *converters) numericField(n int, max uint) string {
	s := strconv.Itoa(n)
	ln := uint(len(s))
	if ln > max {
		return s[ln-max:]
	}
	s = strings.Repeat("0", int(max-ln)) + s
	return s
}

// stringField slices to max length and zero filled
func (c *converters) stringField(s string, max uint) string {
	ln := uint(len(s))
	if ln > max {
		return s[:max]
	}
	s = strings.Repeat("0", int(max-ln)) + s
	return s
}
