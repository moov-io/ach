// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"
	"time"
)

func mockAddenda99() *Addenda99 {
	addenda99 := NewAddenda99()
	addenda99.ReturnCode = "R07"
	addenda99.OriginalTrace = 99912340000015
	addenda99.AddendaInformation = "Authorization Revoked"
	addenda99.OriginalDFI = "9101298"

	return addenda99
}

func TestAddenda99Parse(t *testing.T) {
	addenda99 := NewAddenda99()
	line := "799R07099912340000015      09101298Authorization revoked                       091012980000066"
	addenda99.Parse(line)
	// walk the Addenda99 struct
	if addenda99.recordType != "7" {
		t.Errorf("expected %v got %v", "7", addenda99.recordType)
	}
	if addenda99.typeCode != "99" {
		t.Errorf("expected %v got %v", "99", addenda99.typeCode)
	}
	if addenda99.ReturnCode != "R07" {
		t.Errorf("expected %v got %v", "R07", addenda99.ReturnCode)
	}
	if addenda99.OriginalTrace != 99912340000015 {
		t.Errorf("expected: %v got: %v", 99912340000015, addenda99.OriginalTrace)
	}
	if addenda99.DateOfDeath.IsZero() != true {
		t.Errorf("expected: %v got: %v", time.Time{}, addenda99.DateOfDeath)
	}
	if addenda99.OriginalDFI != "09101298" {
		t.Errorf("expected: %s got: %s", "09101298", addenda99.OriginalDFI)
	}
	if addenda99.AddendaInformation != "Authorization revoked" {
		t.Errorf("expected: %v got: %v", "Authorization revoked", addenda99.AddendaInformation)
	}
	if addenda99.TraceNumber != 91012980000066 {
		t.Errorf("expected: %v got: %v", 91012980000066, addenda99.TraceNumber)
	}
}

func TestAddenda99String(t *testing.T) {
	addenda99 := NewAddenda99()
	line := "799R07099912340000015      09101298Authorization revoked                       091012980000066"
	addenda99.Parse(line)

	if addenda99.String() != line {
		t.Errorf("\n expected: %v\n got     : %v", line, addenda99.String())
	}
}

// This is not an exported function but utilized for validation
func TestAddenda99MakeReturnCodeDict(t *testing.T) {
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

func TestAddenda99ValidateTrue(t *testing.T) {
	addenda99 := mockAddenda99()
	addenda99.ReturnCode = "R13"
	if err := addenda99.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ReturnCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestAddenda99ValidateReturnCodeFalse(t *testing.T) {
	addenda99 := mockAddenda99()
	addenda99.ReturnCode = ""
	if err := addenda99.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ReturnCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestAddenda99OriginalTraceField(t *testing.T) {
	addenda99 := mockAddenda99()
	addenda99.OriginalTrace = 12345
	if addenda99.OriginalTraceField() != "000000000012345" {
		t.Errorf("expected %v received %v", "000000000012345", addenda99.OriginalTraceField())
	}
}

func TestAddenda99DateOfDeathField(t *testing.T) {
	addenda99 := mockAddenda99()
	// Check for all zeros
	if addenda99.DateOfDeathField() != "      " {
		t.Errorf("expected %v received %v", "      ", addenda99.DateOfDeathField())
	}
	// Year: 1978 Month: October Day: 23
	addenda99.DateOfDeath = time.Date(1978, time.October, 23, 0, 0, 0, 0, time.UTC)
	if addenda99.DateOfDeathField() != "781023" {
		t.Errorf("expected %v received %v", "781023", addenda99.DateOfDeathField())
	}
}

func TestAddenda99OriginalDFIField(t *testing.T) {
	addenda99 := mockAddenda99()
	exp := "09101298"
	if addenda99.OriginalDFIField() != exp {
		t.Errorf("expected %v received %v", exp, addenda99.OriginalDFIField())
	}
}

func TestAddenda99AddendaInformationField(t *testing.T) {
	addenda99 := mockAddenda99()
	exp := "Authorization Revoked                       "
	if addenda99.AddendaInformationField() != exp {
		t.Errorf("expected %v received %v", exp, addenda99.AddendaInformationField())
	}
}

func TestAddenda99TraceNumberField(t *testing.T) {
	addenda99 := mockAddenda99()
	addenda99.TraceNumber = 91012980000066
	exp := "091012980000066"
	if addenda99.TraceNumberField() != exp {
		t.Errorf("expected %v received %v", exp, addenda99.TraceNumberField())
	}
}
