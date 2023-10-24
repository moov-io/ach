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
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
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

func (c *converters) parseStringFieldWithOpts(r string, opts *ValidateOpts) string {
	if opts != nil && opts.PreserveSpaces {
		return r
	} else {
		return c.parseStringField(r)
	}
}

// formatSimpleDate takes a YYMMDD date and formats it for the fixed-width ACH file format
func (c *converters) formatSimpleDate(s string) string {
	if s == "" {
		return c.stringField(s, 6)
	}
	return s
}

// formatSimpleTime takes a HHmm (H=hour, m=minute) time and formats it for the fixed-width ACH file format
func (c *converters) formatSimpleTime(s string) string {
	if s == "" {
		return c.stringField(s, 4)
	}
	return s
}

var (
	spaceZeros  map[int]string = populateMap(94, " ")
	stringZeros map[int]string = populateMap(94, "0")
)

// populateMap will allocate strings for padding ACH fields.
//
// In Go strings are immutable so they can be reused across objects without needing to allocate new objects.
func populateMap(max int, zero string) map[int]string {
	out := make(map[int]string, max)
	for i := 0; i < max; i++ {
		out[i] = strings.Repeat(zero, i)
	}
	return out
}

// alphaField Alphanumeric and Alphabetic fields are left-justified and space filled.
func (c *converters) alphaField(s string, max uint) string {
	ln := uint(utf8.RuneCountInString(s))
	if ln > max {
		return s[:max]
	}

	m := int(max - ln)
	pad, exists := spaceZeros[m]
	if exists {
		return s + pad
	}
	// slow path
	return s + strings.Repeat(" ", int(max-ln))
}

// numericField right-justified, unsigned, and zero filled
func (c *converters) numericField(n int, max uint) string {
	s := strconv.FormatInt(int64(n), 10)
	if l := uint(len(s)); l > max {
		return s[l-max:]
	} else {
		m := int(max - l)
		pad, exists := stringZeros[m]
		if exists {
			return pad + s
		}
		// slow path
		return strings.Repeat("0", m) + s
	}
}

// stringField slices to max length and zero filled
func (c *converters) stringField(s string, max uint) string {
	ln := uint(utf8.RuneCountInString(s))
	if ln > max {
		return s[:max]
	}

	// Pad with preallocated string
	m := int(max - ln)
	pad, exists := stringZeros[m]
	if exists {
		return pad + s
	}
	// slow path
	return strings.Repeat("0", m) + s
}

// leastSignificantDigits returns the least significant digits of v limited by maxDigits.
func (c *converters) leastSignificantDigits(v int, maxDigits uint) int {
	return v % int(math.Pow10(int(maxDigits)))
}
