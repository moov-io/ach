// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strconv"
	"strings"
	"time"

	"github.com/moov-io/base"
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

// formatSimpleDate takes a github.com/moov-io/base Time and returns a string of YYMMDD
func (c *converters) formatSimpleDate(t base.Time) string {
	return t.Format("060102")
}

// parseSimpleDate returns a github.com/moov-io/base Time when passed time as YYMMDD
func (c *converters) parseSimpleDate(s string) base.Time {
	t, _ := time.Parse("060102", s)
	return base.NewTime(t)
}

// formatSimpleTime returns a string of HHMM when passed a github.com/moov-io/base Time
func (c *converters) formatSimpleTime(t base.Time) string {
	return t.Format("1504")
}

// parseSimpleTime returns a github.com/moov-io/base Time when passed a string of HHMM
func (c *converters) parseSimpleTime(s string) base.Time {
	t, _ := time.Parse("1504", s)
	return base.NewTime(t)
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
