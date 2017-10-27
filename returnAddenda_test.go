// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"
	"time"
)

func mockReturnAddenda() *ReturnAddenda {
	rAddenda := NewReturnAddenda()
	rAddenda.typeCode = "99"
	rAddenda.ReturnCode = "R07"
	rAddenda.AddendaInformation = "Authorization Revoked"
	rAddenda.OriginalDFI = 9101298

	return rAddenda
}

func TestMockReturnAddenda(t *testing.T) {
	// TODO: build a mock addenda
}

func TestReturnAddendaParse(t *testing.T) {
	rAddenda := NewReturnAddenda()
	line := "799R07099912340000015      09101298Authorization revoked                       091012980000066"
	rAddenda.Parse(line)
	// walk the returnAddenda struct
	if rAddenda.recordType != "7" {
		t.Errorf("expected %v got %v", "7", rAddenda.recordType)
	}
	if rAddenda.typeCode != "99" {
		t.Errorf("expected %v got %v", "99", rAddenda.typeCode)
	}
	if rAddenda.ReturnCode != "R07" {
		t.Errorf("expected %v got %v", "R07", rAddenda.ReturnCode)
	}
	if rAddenda.OriginalTrace != 99912340000015 {
		t.Errorf("expected: %v got: %v", 99912340000015, rAddenda.OriginalTrace)
	}
	if rAddenda.DateOfDeath.IsZero() != true {
		t.Errorf("expected: %v got: %v", time.Time{}, rAddenda.DateOfDeath)
	}
	if rAddenda.OriginalDFI != 9101298 {
		t.Errorf("expected: %v got: %v", 9101298, rAddenda.OriginalDFI)
	}
	if rAddenda.AddendaInformation != "Authorization revoked" {
		t.Errorf("expected: %v got: %v", "Authorization revoked", rAddenda.AddendaInformation)
	}
	if rAddenda.TraceNumber != 91012980000066 {
		t.Errorf("expected: %v got: %v", 91012980000066, rAddenda.TraceNumber)
	}
}

func TestReturnAddendaString(t *testing.T) {
	rAddenda := NewReturnAddenda()
	line := "799R07099912340000015      09101298Authorization revoked                       091012980000066"
	rAddenda.Parse(line)

	if rAddenda.String() != line {
		t.Errorf("\n expected: %v\n got     : %v", line, rAddenda.String())
	}
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

func TestReturnAddendaValidateTrue(t *testing.T) {
	rAddenda := mockReturnAddenda()
	rAddenda.ReturnCode = "R13"
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

func TestReturnAddendaValidateReturnCodeFalse(t *testing.T) {
	rAddenda := mockReturnAddenda()
	rAddenda.ReturnCode = ""
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

func TestReturnAddendaOriginalTraceField(t *testing.T) {
	rAddenda := mockReturnAddenda()
	rAddenda.OriginalTrace = 12345
	if rAddenda.OriginalTraceField() != "000000000012345" {
		t.Errorf("expected %v received %v", "000000000012345", rAddenda.OriginalTraceField())
	}
}

func TestReturnAddendaDateOfDeathField(t *testing.T) {
	rAddenda := mockReturnAddenda()
	// Check for all zeros
	if rAddenda.DateOfDeathField() != "      " {
		t.Errorf("expected %v received %v", "      ", rAddenda.DateOfDeathField())
	}
	// Year: 1978 Month: October Day: 23
	rAddenda.DateOfDeath = time.Date(1978, time.October, 23, 0, 0, 0, 0, time.UTC)
	if rAddenda.DateOfDeathField() != "781023" {
		t.Errorf("expected %v received %v", "781023", rAddenda.DateOfDeathField())
	}
}

func TestReturnAddendaOriginalDFIField(t *testing.T) {
	rAddenda := mockReturnAddenda()
	exp := "09101298"
	if rAddenda.OriginalDFIField() != exp {
		t.Errorf("expected %v received %v", exp, rAddenda.OriginalDFIField())
	}
}

func TestReturnAddendaAddendaInformationField(t *testing.T) {
	rAddenda := mockReturnAddenda()
	exp := "Authorization Revoked                       "
	if rAddenda.AddendaInformationField() != exp {
		t.Errorf("expected %v received %v", exp, rAddenda.AddendaInformationField())
	}
}

func TestReturnAddendaTraceNumberField(t *testing.T) {
	rAddenda := mockReturnAddenda()
	rAddenda.TraceNumber = 91012980000066
	exp := "091012980000066"
	if rAddenda.TraceNumberField() != exp {
		t.Errorf("expected %v received %v", exp, rAddenda.TraceNumberField())
	}
}
