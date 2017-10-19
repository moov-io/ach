// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"
)

func mockReturnAddenda() ReturnAddenda {
	rAddenda := NewReturnAddenda()

	return rAddenda
}

func TestMockReturnAddenda(t *testing.T) {
	// TODO: build a mock addenda
}

// This is not an exported function but utilized for validation
func TestReturnAddendaMakeReturnCodeDict(t *testing.T) {
	codes := makeReturnCodeDict()
	// check if known code is present
	_, prs := codes["R01"]
	if !prs {
		t.Error("Return Code R01 was not found in the ReturnCodeDict")
	}
	// check if invalid code is present
	_, prs = codes["ABC"]
	if prs {
		t.Error("Valid return for an invalid return code key")
	}
}

func TestReturnAddendaValidate(t *testing.T) {
	rAddenda := mockReturnAddenda()
	if err := rAddenda.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ReturnCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// Notification of Change COR
//"798C01099912340000015      091012981918171614 091012980000088"
