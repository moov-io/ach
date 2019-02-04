// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package base

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
)

// UnwrappableError is an interface for errors that wrap another error with some extra context
// The interface allows these errors to get automatically unwrapped by the Match function
type UnwrappableError interface {
	Error() string
	Unwrap() error
}

// ParseError is returned for parsing reader errors.
// The first line is 1.
type ParseError struct {
	Line   int    // Line number where the error occurred
	Record string // Name of the record type being parsed
	Err    error  // The actual error
}

func (e ParseError) Error() string {
	if e.Record == "" {
		return fmt.Sprintf("line:%d %T %s", e.Line, e.Err, e.Err)
	}
	return fmt.Sprintf("line:%d record:%s %T %s", e.Line, e.Record, e.Err, e.Err)
}

// Unwrap implements the UnwrappableError interface for ParseError
func (e ParseError) Unwrap() error {
	return e.Err
}

// ErrorList represents an array of errors which is also an error itself.
type ErrorList []error

// Add appends err onto the ErrorList. Errors are kept in append order.
func (e *ErrorList) Add(err error) {
	*e = append(*e, err)
}

// Err returns the first error (or nil).
func (e ErrorList) Err() error {
	if e == nil || len(e) == 0 {
		return nil
	}
	return e[0]
}

// Error implements the error interface
func (e ErrorList) Error() string {
	if len(e) == 0 {
		return "<nil>"
	}
	var buf bytes.Buffer
	e.Print(&buf)
	return buf.String()
}

// Print formats the ErrorList into a string written to w.
// If ErrorList contains multiple errors those after the first
// are indented.
func (e ErrorList) Print(w io.Writer) {
	if w == nil || len(e) == 0 {
		fmt.Fprintf(w, "<nil>")
		return
	}

	fmt.Fprintf(w, "%s", e[0])
	if len(e) > 1 {
		fmt.Fprintf(w, "\n")
	}

	for i := 1; i < len(e); i++ {
		fmt.Fprintf(w, "  %s", e[i])
		if i < len(e)-1 { // don't add \n to last error
			fmt.Fprintf(w, "\n")
		}
	}
}

// Empty no errors to return
func (e ErrorList) Empty() bool {
	return e == nil || len(e) == 0
}

// MarshalJSON marshals error list
func (e ErrorList) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Error())
}

// Match takes in two errors and compares them, returning true if they match and false if they don't
// The matching is done by basic equality for simple errors (i.e. defined by errors.New) and by type
// for other errors. If errA is wrapped with an error supporting the UnwrappableError interface it
// will also unwrap it and then recursively compare the unwrapped error with errB.
func Match(errA, errB error) bool {
	if errA == nil {
		return errB == nil
	}

	// typed errors can be compared by type
	if reflect.TypeOf(errA) == reflect.TypeOf(errB) {
		simpleError := errors.New("simple error")
		if reflect.TypeOf(errB) == reflect.TypeOf(simpleError) {
			// simple errors all have the same type, so we need to compare them directly
			return errA == errB
		}
		return true
	}

	// match wrapped errors
	uwErr, ok := errA.(UnwrappableError)
	if ok {
		return Match(uwErr.Unwrap(), errB)
	}

	return false
}

// Has takes in a (potential) list of errors, and an error to check for. If any of the errors
// in the list have the same type as the error to check, it returns true. If the "list" isn't
// actually a list (typically because it is nil), or no errors in the list match the other error
// it returns false. So it can be used as an easy way to check for a particular kind of error.
func Has(list error, err error) bool {
	el, ok := list.(ErrorList)
	if !ok {
		return false
	}
	for i := 0; i < len(el); i++ {
		if Match(el[i], err) {
			return true
		}
	}
	return false
}
