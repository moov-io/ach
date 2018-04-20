package ach

import (
	"testing"
)

func mockAddenda98() *Addenda98 {
	addenda98 := NewAddenda98()
	addenda98.ChangeCode = "C01"
	addenda98.OriginalTrace = 12345
	addenda98.OriginalDFI = 9101298
	addenda98.CorrectedData = "1918171614"
	addenda98.TraceNumber = 91012980000088

	return addenda98
}

func TestAddenda98Parse(t *testing.T) {
	addenda98 := NewAddenda98()
	line := "798C01099912340000015      091012981918171614                                  091012980000088"
	addenda98.Parse(line)
	// walk the Addenda98 struct
	if addenda98.recordType != "7" {
		t.Errorf("expected %v got %v", "7", addenda98.recordType)
	}
	if addenda98.typeCode != "98" {
		t.Errorf("expected %v got %v", "98", addenda98.typeCode)
	}
	if addenda98.ChangeCode != "C01" {
		t.Errorf("expected %v got %v", "C01", addenda98.ChangeCode)
	}
	if addenda98.OriginalTrace != 99912340000015 {
		t.Errorf("expected %v got %v", 99912340000015, addenda98.OriginalTrace)
	}
	if addenda98.OriginalDFI != 9101298 {
		t.Errorf("expected %v got %v", 9101298, addenda98.OriginalDFI)
	}
	if addenda98.CorrectedData != "1918171614" {
		t.Errorf("expected %v got %v", "1918171614", addenda98.CorrectedData)
	}
	if addenda98.TraceNumber != 91012980000088 {
		t.Errorf("expected %v got %v", 91012980000088, addenda98.TraceNumber)
	}
}

func TestAddenda98String(t *testing.T) {
	addenda98 := NewAddenda98()
	line := "798C01099912340000015      091012981918171614                                  091012980000088"
	addenda98.Parse(line)

	if addenda98.String() != line {
		t.Errorf("\n expected: %v\n got     : %v", line, addenda98.String())
	}
}

func TestAddenda98ValidRecordType(t *testing.T) {
	addenda98 := mockAddenda98()
	addenda98.recordType = "63"
	if err := addenda98.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestAddenda98ValidTypeCode(t *testing.T) {
	addenda98 := mockAddenda98()
	addenda98.typeCode = "05"
	if err := addenda98.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TypeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestAddenda98ValidCorrectedData(t *testing.T) {
	addenda98 := mockAddenda98()
	addenda98.CorrectedData = ""
	if err := addenda98.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "CorrectedData" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestAddenda98ValidateTrue(t *testing.T) {
	addenda98 := mockAddenda98()
	addenda98.ChangeCode = "C11"
	if err := addenda98.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ChangeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}
func TestAddenda98ValidateChangeCodeFalse(t *testing.T) {
	addenda98 := mockAddenda98()
	addenda98.ChangeCode = "C63"
	if err := addenda98.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ChangeCode" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestAddenda98OriginalTraceField(t *testing.T) {
	addenda98 := mockAddenda98()
	exp := "000000000012345"
	if addenda98.OriginalTraceField() != exp {
		t.Errorf("expected %v received %v", exp, addenda98.OriginalTraceField())
	}
}

func TestAddenda98OriginalDFIField(t *testing.T) {
	addenda98 := mockAddenda98()
	exp := "09101298"
	if addenda98.OriginalDFIField() != exp {
		t.Errorf("expected %v received %v", exp, addenda98.OriginalDFIField())
	}
}

func TestAddenda98CorrectedDataField(t *testing.T) {
	addenda98 := mockAddenda98()
	exp := "1918171614                   " // 29 char
	if addenda98.CorrectedDataField() != exp {
		t.Errorf("expected %v received %v", exp, addenda98.CorrectedDataField())
	}
}

func TestAddenda98TraceNumberField(t *testing.T) {
	addenda98 := mockAddenda98()
	exp := "091012980000088"
	if addenda98.TraceNumberField() != exp {
		t.Errorf("expected %v received %v", exp, addenda98.TraceNumberField())
	}
}
