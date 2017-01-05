package ach

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Validator is common validation and formating of golang types to ach type strings
type Validator struct{}

// isValidServiceClass returns true if a valid service class code
func (v *Validator) isValidServiceClass(code int) bool {
	switch code {
	case
		// ACH Mixed Debits and Credits
		200,
		// ACH Credits Only
		220,
		// ACH Debits Only
		225,
		// ACH Automated Accounting Advices
		280:
		return true
	}
	return false
}

// isValidSECCode returns true if a SEC Code is found
func (v *Validator) isValidSECCode(code string) bool {
	switch code {
	case
		"ACK", "ADV", "ARC", "ATX", "BOC", "CCD", "CIE", "COR", "CTX", "DNE", "ENR",
		"IAT", "MTE", "POS", "PPD", "POP", "RCK", "SHR", "TEL", "TRC", "TRX", "WEB", "XCK":
		return true
	}
	return false
}

// isOriginatorStatusCode ensures status code is valid
func (v *Validator) isOriginatorStatusCode(code int) bool {
	switch code {
	case
		// ADV file - prepared by an ACH Operator
		0,
		//Originator is a financial institution
		1,
		// Originator is a Government Agency or other agency not subject to ACH Rules
		2:
		return true
	}
	return false
}

func (v *Validator) parseNumField(r string) (s int) {
	s, err := strconv.Atoi(strings.TrimSpace(r))
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	return s
}

// formatSimpleDate takes a time.Time and returns a string of YYMMDD
func (v *Validator) formatSimpleDate(t time.Time) string {
	return t.Format("060102")
}

// parseSimpleDate returns a time.Time when passed time as YYMMDD
func (v *Validator) parseSimpleDate(s string) time.Time {
	t, _ := time.Parse("060102", s)
	return t
}

// formatSimpleTime returns a string of HHMM when  passed a time.Time
func (v *Validator) formatSimpleTime(t time.Time) string {
	return t.Format("1504")
}

// parseSimpleTime returns a time.Time when passed a string of HHMM
func (v *Validator) parseSimpleTime(s string) time.Time {
	t, _ := time.Parse("1504", s)
	return t
}

// isUpperAlphanumeric checks if string only contains ASCII alphanumeric upper case characters
func (v *Validator) isUpperAlphanumeric(s string) (b bool) {
	if regexp.MustCompile(`[^A-Z0-9]+`).MatchString(s) {
		return false
	}
	return true
}

// isAlphanumeric checks if a string only contains ASCII alphanumeric characters
func (v *Validator) isAlphanumeric(s string) (b bool) {
	if regexp.MustCompile(`[^ a-zA-Z0-9_*\/]+`).MatchString(s) {
		// ^[ A-Za-z0-9_@./#&+-]*$/
		return false
	}
	return true
}

// isLetter checks if a string contains only ASCII letters
func (v *Validator) isLetter(s string) bool {
	fmt.Println(s)
	for _, r := range s {
		fmt.Printf("%v ", r)
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

// rightPad takes a string and padds the left side of s to overallLen with padStr.
func (v *Validator) rightPad(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + (overallLen - len(padStr))
	var retStr = s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}

// leftPad takes a string and padds the right side of s to overallLen with padStr.
func (v *Validator) leftPad(s string, padStr string, overallLen int) string {
	var padCountInt int
	padCountInt = 1 + (overallLen - len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + s
	return retStr[(len(retStr) - overallLen):]
}
