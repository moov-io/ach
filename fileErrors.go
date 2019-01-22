// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"reflect"

	"github.com/moov-io/base"
)

// FileTooLongErr is the error given when a file exceeds the maximum possible length
type FileTooLongErr string

func (e FileTooLongErr) Error() string {
	return string(e)
}

// RecordWrongLengthErr is the error given when a record is the wrong length
type RecordWrongLengthErr struct {
	Message string
	Length  int
}

// NewRecordWrongLengthErr creates a new error of the RecordWrongLengthErr type
func NewRecordWrongLengthErr(length int) RecordWrongLengthErr {
	return RecordWrongLengthErr{
		Message: fmt.Sprintf("must be 94 characters and found %d", length),
		Length:  length,
	}
}

func (e RecordWrongLengthErr) Error() string {
	return e.Message
}

// Has takes in a (potential) list of errors, and an error to check for. If any of the errors
// in the list have the same type as the error to check, it returns true. If the "list" isn't
// actually a list (typically because it is nil), or no errors in the list match the other error
// it returns false. So it can be used as an easy way to check for a particular kind of error.
func Has(list error, err error) bool {
	el, ok := list.(base.ErrorList)
	if !ok {
		return false
	}
	for i := 0; i < len(el); i++ {
		if reflect.TypeOf(el[i]) == reflect.TypeOf(err) {
			return true
		}
	}
	return false
}
