package ach

import "testing"

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

func TestNOCAddendaValidateTrue(t *testing.T) {
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
