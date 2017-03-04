// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

// validator is common validation and formating of golang types to ach type strings
type validator struct{}

// ValidateError is returned for errors at a field level in a record
type ValidateError struct {
	FieldName string // field name where error happend
	Value     string // value that cause error
	Err       error  // the actual error
}

func (e *ValidateError) Error() string {
	return fmt.Sprintf("FieldName:%s Value:%s msg: %s ", e.FieldName, e.Value, e.Err)
}

// Errors specific to validation
var (
	ErrValidAlphanumeric   = errors.New("Field has non alphanumeric characters")
	ErrValidAlpha          = errors.New("Field has non alpha characters")
	ErrUpperAlpha          = errors.New("Field is not uppercase A-Z or 0-9")
	ErrValidFieldInclusion = errors.New("A mandatory field has a default value")
	ErrValidFieldLength    = errors.New("Field length is invalid")
	ErrServiceClass        = errors.New("Invalid Service Class Code")
	ErrSECCode             = errors.New("Invalid Standard Entry Class Code")
	ErrOrigStatusCode      = errors.New("Invalid Originator Status Code")
	ErrAddendaTypeCode     = errors.New("Invalid Addenda Type Code")
	ErrTransactionCode     = errors.New("Invalid Transaction Code")
	ErrValidCheckDigit     = errors.New("Check Digit does not match calculated check digit")
)

// iServiceClass returns true if a valid service class code
func (v *validator) isServiceClass(code int) error {
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
		return nil
	}
	return ErrServiceClass
}

// isSECCode returns true if a SEC Code is found
func (v *validator) isSECCode(code string) error {
	switch code {
	case
		"ACK", "ADV", "ARC", "ATX", "BOC", "CCD", "CIE", "COR", "CTX", "DNE", "ENR",
		"IAT", "MTE", "POS", "PPD", "POP", "RCK", "SHR", "TEL", "TRC", "TRX", "WEB", "XCK":
		return nil
	}
	return ErrSECCode
}

// isTypeCode returns true if a valid type code is found
func (v *validator) isTypeCode(code string) error {
	switch code {
	case
		// For POS, SHR or MTE Entries
		"02",
		// Addenda Record (ACK, ATX, CCD, CIE, CTX, DNE, ENR, PPD, TRX and WEB Entries
		"08",
		// Notification of Change and Refused Notification of Change Entry
		"98",
		// Return, Dishonored Return and Contested Dishonored Return Entries
		"99",
		//  IAT forward Entries and IAT Returns
		"10", "11", "12", "13", "14", "15", "16", "17",
		// CCD Addenda Record
		"05":
		return nil
	}
	return ErrAddendaTypeCode
}

// isTransactionCode ensures TransactionCode code is valid
func (v *validator) isTransactionCode(code int) error {
	switch code {
	// TransactionCode if the receivers account is:
	case
		// Credit (deposit) to checking account ‘22’
		22,
		// Prenote for credit to checking account ‘23’
		23,
		// Debit (withdrawal) to checking account ‘27’
		27,
		// Prenote for debit to checking account ‘28’
		28,
		// Credit to savings account ‘32’
		32,
		// Prenote for credit to savings account ‘33’
		33,
		// Debit to savings account ‘37’
		37,
		// Prenote for debit to savings account ‘38’
		38:
		return nil
	}
	return ErrTransactionCode
}

// isOriginatorStatusCode ensures status code is valid
func (v *validator) isOriginatorStatusCode(code int) error {
	switch code {
	case
		// ADV file - prepared by an ACH Operator
		0,
		//Originator is a financial institution
		1,
		// Originator is a Government Agency or other agency not subject to ACH Rules
		2:
		return nil
	}
	return ErrOrigStatusCode
}

// isUpperAlphanumeric checks if string only contains ASCII alphanumeric upper case characters
func (v *validator) isUpperAlphanumeric(s string) error {
	if regexp.MustCompile(`[^A-Z0-9]+`).MatchString(s) {
		return ErrUpperAlpha
	}
	return nil
}

// isAlphanumeric checks if a string only contains ASCII alphanumeric characters
func (v *validator) isAlphanumeric(s string) error {
	if regexp.MustCompile(`[^ a-zA-Z0-9_*-\/]+`).MatchString(s) {
		// ^[ A-Za-z0-9_@./#&+-]*$/
		return ErrValidAlphanumeric
	}
	return nil
}

// isOriginatorStatusCode ensures status code is valid
func (v *validator) isCheckDigit(routingNumber string, checkDigit int) error {
	if v.CalculateCheckDigit(routingNumber) != checkDigit {
		return ErrValidCheckDigit
	}
	return nil
}

// CalculateCheckDigit returns a check digit for a rounting number
// Multiply each digit in the Routing number by a weighting factor. The weighting factors for each digit are:
// Position: 1 2 3 4 5 6 7 8
// Weights : 3 7 1 3 7 1 3 7
// Add the results of the eight multiplications
// Subtract the sum from the next highest multiple of 10.
// The result is the Check Digit
func (v *validator) CalculateCheckDigit(routingNumber string) int {
	var routeIndex [8]string
	for i := 0; i < 8; i++ {
		routeIndex[i] = string(routingNumber[i])
	}
	n, _ := strconv.Atoi(routeIndex[0])
	sum := (n * 3)
	n, _ = strconv.Atoi(routeIndex[1])
	sum = sum + (n * 7)
	n, _ = strconv.Atoi(routeIndex[2])
	sum = sum + n // multiply by 1
	n, _ = strconv.Atoi(routeIndex[3])
	sum = sum + (n * 3)
	n, _ = strconv.Atoi(routeIndex[4])
	sum = sum + (n * 7)
	n, _ = strconv.Atoi(routeIndex[5])
	sum = sum + n // multiply by 1
	n, _ = strconv.Atoi(routeIndex[6])
	sum = sum + (n * 3)
	n, _ = strconv.Atoi(routeIndex[7])
	sum = sum + (n * 7)

	return v.roundUp10(sum) - sum
}

// roundUp10 round number up to the next ten spot.
func (v *validator) roundUp10(n int) int {
	return int(math.Ceil(float64(n)/10.0)) * 10
}
