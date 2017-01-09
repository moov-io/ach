// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Converters handles golang to ACH type Converters
type Converters struct{}

func (v *Converters) parseNumField(r string) (s int) {
	s, err := strconv.Atoi(strings.TrimSpace(r))
	if err != nil {
		// TODO: This is horrible
		fmt.Printf("%v", err)
		return
	}
	return s
}

// formatSimpleDate takes a time.Time and returns a string of YYMMDD
func (v *Converters) formatSimpleDate(t time.Time) string {
	return t.Format("060102")
}

// parseSimpleDate returns a time.Time when passed time as YYMMDD
func (v *Converters) parseSimpleDate(s string) time.Time {
	t, _ := time.Parse("060102", s)
	return t
}

// formatSimpleTime returns a string of HHMM when  passed a time.Time
func (v *Converters) formatSimpleTime(t time.Time) string {
	return t.Format("1504")
}

// parseSimpleTime returns a time.Time when passed a string of HHMM
func (v *Converters) parseSimpleTime(s string) time.Time {
	t, _ := time.Parse("1504", s)
	return t
}

// rightPad takes a string and padds the left side of s to overallLen with padStr.
func (v *Converters) rightPad(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + (overallLen - len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

// leftPad takes a string and padds the right side of s to overallLen with padStr.
func (v *Converters) leftPad(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + (overallLen - len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}

// @TODO remove decimel space from amount int
