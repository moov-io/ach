package ach

import (
	"strings"
	"testing"
)

func mockAddendaNOC() *AddendaNOC {
	aNOC := NewAddendaNOC()
	aNOC.ChangeCode = "C01"
	aNOC.OriginalTrace = 12345
	aNOC.OriginalDFI = 9101298
	aNOC.CorrectedData = "1918171614"
	aNOC.TraceNumber = 91012980000088

	return aNOC
}

func TestAddendaNOCParse(t *testing.T) {
	aNOC := NewAddendaNOC()
	line := "798C01099912340000015      091012981918171614                                  091012980000088"
	aNOC.Parse(line)
	// walk the AddendaNOC struct
	if aNOC.recordType != "7" {
		t.Errorf("expected %v got %v", "7", aNOC.recordType)
	}
	if aNOC.typeCode != "98" {
		t.Errorf("expected %v got %v", "98", aNOC.typeCode)
	}
	if aNOC.ChangeCode != "C01" {
		t.Errorf("expected %v got %v", "C01", aNOC.ChangeCode)
	}
	if aNOC.OriginalTrace != 99912340000015 {
		t.Errorf("expected %v got %v", 99912340000015, aNOC.OriginalTrace)
	}
	if aNOC.OriginalDFI != 9101298 {
		t.Errorf("expected %v got %v", 9101298, aNOC.OriginalDFI)
	}
	if aNOC.CorrectedData != "1918171614" {
		t.Errorf("expected %v got %v", "1918171614", aNOC.CorrectedData)
	}
	if aNOC.TraceNumber != 91012980000088 {
		t.Errorf("expected %v got %v", 91012980000088, aNOC.TraceNumber)
	}
}

func TestAddendaNOCString(t *testing.T) {
	aNOC := NewAddendaNOC()
	line := "798C01099912340000015      091012981918171614                                  091012980000088"
	aNOC.Parse(line)

	if aNOC.String() != line {
		t.Errorf("\n expected: %v\n got     : %v", line, aNOC.String())
	}
}

func TestAddendaNOCValidRecordType(t *testing.T) {
	aNOC := mockAddendaNOC()
	aNOC.recordType = "63"
	if err := aNOC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestAddendaNOCValidTypeCode(t *testing.T) {
	aNOC := mockAddendaNOC()
	aNOC.typeCode = "05"
	if err := aNOC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestAddendaNOCValidCorrectedData(t *testing.T) {
	aNOC := mockAddendaNOC()
	aNOC.CorrectedData = ""
	if err := aNOC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CorrectedData" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestAddendaNOCValidateTrue(t *testing.T) {
	aNOC := mockAddendaNOC()
	aNOC.ChangeCode = "C11"
	if err := aNOC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ChangeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}
func TestAddendaNOCValidateChangeCodeFalse(t *testing.T) {
	aNOC := mockAddendaNOC()
	aNOC.ChangeCode = "C63"
	if err := aNOC.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ChangeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestAddendaNOCOriginalTraceField(t *testing.T) {
	aNOC := mockAddendaNOC()
	exp := "000000000012345"
	if aNOC.OriginalTraceField() != exp {
		t.Errorf("expected %v received %v", exp, aNOC.OriginalTraceField())
	}
}

func TestAddendaNOCOriginalDFIField(t *testing.T) {
	aNOC := mockAddendaNOC()
	exp := "09101298"
	if aNOC.OriginalDFIField() != exp {
		t.Errorf("expected %v received %v", exp, aNOC.OriginalDFIField())
	}
}

func TestAddendaNOCCorrectedDataField(t *testing.T) {
	aNOC := mockAddendaNOC()
	exp := "1918171614                   " // 29 char
	if aNOC.CorrectedDataField() != exp {
		t.Errorf("expected %v received %v", exp, aNOC.CorrectedDataField())
	}
}

func TestAddendaNOCTraceNumberField(t *testing.T) {
	aNOC := mockAddendaNOC()
	exp := "091012980000088"
	if aNOC.TraceNumberField() != exp {
		t.Errorf("expected %v received %v", exp, aNOC.TraceNumberField())
	}
}

func TestAddendaNOCNewAddendaParam(t *testing.T) {
	aParam := AddendaParam{
		TypeCode:      "98",
		ChangeCode:    "C01",
		OriginalTrace: "12345",
		OriginalDFI:   "9101298",
		CorrectedData: "1918171614",
		TraceNumber:   "91012980000088",
	}

	a, err := NewAddenda(aParam)
	if err != nil {
		t.Errorf("AddendaNOC from NewAddenda: %v", err)
	}
	aNOC, ok := a.(*AddendaNOC)
	if !ok {
		t.Errorf("expecting *AddendaNOC received %T ", a)
	}
	if aNOC.TypeCode() != aParam.TypeCode {
		t.Errorf("expected %v got %v", aParam.TypeCode, aNOC.TypeCode())
	}
	if aNOC.ChangeCode != aParam.ChangeCode {
		t.Errorf("expected %v got %v", aParam.ChangeCode, aNOC.ChangeCode)
	}
	if !strings.Contains(aNOC.OriginalTraceField(), aParam.OriginalTrace) {
		t.Errorf("expected %v got %v", aParam.OriginalTrace, aNOC.OriginalTrace)
	}
	if !strings.Contains(aNOC.OriginalDFIField(), aParam.OriginalDFI) {
		t.Errorf("expected %v got %v", aParam.OriginalDFI, aNOC.OriginalDFI)
	}
	if aNOC.CorrectedData != aParam.CorrectedData {
		t.Errorf("expected %v got %v", aParam.CorrectedData, aNOC.CorrectedData)
	}
	if !strings.Contains(aNOC.TraceNumberField(), aParam.TraceNumber) {
		t.Errorf("expected %v got %v", aParam.TraceNumber, aNOC.TraceNumber)
	}
}
