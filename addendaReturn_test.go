// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
	"time"
)

func mockAddendaReturn() *AddendaReturn {
	addendaReturn := NewAddendaReturn()
	addendaReturn.typeCode = "99"
	addendaReturn.ReturnCode = "R07"
	addendaReturn.OriginalTrace = 99912340000015
	addendaReturn.AddendaInformation = "Authorization Revoked"
	addendaReturn.OriginalDFI = 9101298

	return addendaReturn
}

func TestMockAddendaReturn(t *testing.T) {
	// TODO: build a mock addenda
}

func TestAddendaReturnParse(t *testing.T) {
	addendaReturn := NewAddendaReturn()
	line := "799R07099912340000015      09101298Authorization revoked                       091012980000066"
	addendaReturn.Parse(line)
	// walk the addendaReturn struct
	if addendaReturn.recordType != "7" {
		t.Errorf("expected %v got %v", "7", addendaReturn.recordType)
	}
	if addendaReturn.typeCode != "99" {
		t.Errorf("expected %v got %v", "99", addendaReturn.typeCode)
	}
	if addendaReturn.ReturnCode != "R07" {
		t.Errorf("expected %v got %v", "R07", addendaReturn.ReturnCode)
	}
	if addendaReturn.OriginalTrace != 99912340000015 {
		t.Errorf("expected: %v got: %v", 99912340000015, addendaReturn.OriginalTrace)
	}
	if addendaReturn.DateOfDeath.IsZero() != true {
		t.Errorf("expected: %v got: %v", time.Time{}, addendaReturn.DateOfDeath)
	}
	if addendaReturn.OriginalDFI != 9101298 {
		t.Errorf("expected: %v got: %v", 9101298, addendaReturn.OriginalDFI)
	}
	if addendaReturn.AddendaInformation != "Authorization revoked" {
		t.Errorf("expected: %v got: %v", "Authorization revoked", addendaReturn.AddendaInformation)
	}
	if addendaReturn.TraceNumber != 91012980000066 {
		t.Errorf("expected: %v got: %v", 91012980000066, addendaReturn.TraceNumber)
	}
}

func TestAddendaReturnString(t *testing.T) {
	addendaReturn := NewAddendaReturn()
	line := "799R07099912340000015      09101298Authorization revoked                       091012980000066"
	addendaReturn.Parse(line)

	if addendaReturn.String() != line {
		t.Errorf("\n expected: %v\n got     : %v", line, addendaReturn.String())
	}
}

// This is not an exported function but utilized for validation
func TestAddendaReturnMakeReturnCodeDict(t *testing.T) {
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

func TestAddendaReturnValidateTrue(t *testing.T) {
	addendaReturn := mockAddendaReturn()
	addendaReturn.ReturnCode = "R13"
	if err := addendaReturn.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ReturnCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestAddendaReturnValidateReturnCodeFalse(t *testing.T) {
	addendaReturn := mockAddendaReturn()
	addendaReturn.ReturnCode = ""
	if err := addendaReturn.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ReturnCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestAddendaReturnOriginalTraceField(t *testing.T) {
	addendaReturn := mockAddendaReturn()
	addendaReturn.OriginalTrace = 12345
	if addendaReturn.OriginalTraceField() != "000000000012345" {
		t.Errorf("expected %v received %v", "000000000012345", addendaReturn.OriginalTraceField())
	}
}

func TestAddendaReturnDateOfDeathField(t *testing.T) {
	addendaReturn := mockAddendaReturn()
	// Check for all zeros
	if addendaReturn.DateOfDeathField() != "      " {
		t.Errorf("expected %v received %v", "      ", addendaReturn.DateOfDeathField())
	}
	// Year: 1978 Month: October Day: 23
	addendaReturn.DateOfDeath = time.Date(1978, time.October, 23, 0, 0, 0, 0, time.UTC)
	if addendaReturn.DateOfDeathField() != "781023" {
		t.Errorf("expected %v received %v", "781023", addendaReturn.DateOfDeathField())
	}
}

func TestAddendaReturnOriginalDFIField(t *testing.T) {
	addendaReturn := mockAddendaReturn()
	exp := "09101298"
	if addendaReturn.OriginalDFIField() != exp {
		t.Errorf("expected %v received %v", exp, addendaReturn.OriginalDFIField())
	}
}

func TestAddendaReturnAddendaInformationField(t *testing.T) {
	addendaReturn := mockAddendaReturn()
	exp := "Authorization Revoked                       "
	if addendaReturn.AddendaInformationField() != exp {
		t.Errorf("expected %v received %v", exp, addendaReturn.AddendaInformationField())
	}
}

func TestAddendaReturnTraceNumberField(t *testing.T) {
	addendaReturn := mockAddendaReturn()
	addendaReturn.TraceNumber = 91012980000066
	exp := "091012980000066"
	if addendaReturn.TraceNumberField() != exp {
		t.Errorf("expected %v received %v", exp, addendaReturn.TraceNumberField())
	}
}

func TestAddendaReturnNewAddendaParam(t *testing.T) {
	aParam := AddendaParam{
		TypeCode:      "99",
		ReturnCode:    "R07",
		OriginalTrace: "99912340000015",
		OriginalDFI:   "09101298",
		AddendaInfo:   "Authorization Revoked",
		TraceNumber:   "091012980000066",
	}

	a, err := NewAddenda(aParam)
	if err != nil {
		t.Errorf("addendaReturn from NewAddeda: %v", err)
	}
	addendaReturn, ok := a.(*AddendaReturn)
	if !ok {
		t.Errorf("expecting *AddendaReturn received %T ", a)
	}
	if addendaReturn.TypeCode() != aParam.TypeCode {
		t.Errorf("expected %v got %v", aParam.TypeCode, addendaReturn.TypeCode())
	}
	if addendaReturn.ReturnCode != aParam.ReturnCode {
		t.Errorf("expected %v got %v", aParam.ReturnCode, addendaReturn.ReturnCode)
	}
	if !strings.Contains(addendaReturn.OriginalTraceField(), aParam.OriginalTrace) {
		t.Errorf("expected %v got %v", aParam.OriginalTrace, addendaReturn.OriginalTrace)
	}
	if !strings.Contains(addendaReturn.OriginalDFIField(), aParam.OriginalDFI) {
		t.Errorf("expected %v got %v", aParam.OriginalDFI, addendaReturn.OriginalDFI)
	}
	if addendaReturn.AddendaInformation != aParam.AddendaInfo {
		t.Errorf("expected %v got %v", aParam.AddendaInfo, addendaReturn.AddendaInformation)
	}
	if !strings.Contains(addendaReturn.TraceNumberField(), aParam.TraceNumber) {
		t.Errorf("expected %v got %v", aParam.TraceNumber, addendaReturn.TraceNumber)
	}
}
