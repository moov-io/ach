package ach

import (
	"testing"
)

func mockAddenda98() *Addenda98 {
	addenda98 := NewAddenda98()
	addenda98.ChangeCode = "C01"
	addenda98.OriginalTrace = 12345
	addenda98.OriginalDFI = "9101298"
	addenda98.CorrectedData = "1918171614"
	addenda98.TraceNumber = 91012980000088

	return addenda98
}

func testAddenda98Parse(t testing.TB) {
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
	if addenda98.OriginalDFI != "09101298" {
		t.Errorf("expected %s got %s", "09101298", addenda98.OriginalDFI)
	}
	if addenda98.CorrectedData != "1918171614" {
		t.Errorf("expected %v got %v", "1918171614", addenda98.CorrectedData)
	}
	if addenda98.TraceNumber != 91012980000088 {
		t.Errorf("expected %v got %v", 91012980000088, addenda98.TraceNumber)
	}
}

func TestAddenda98Parse(t *testing.T) {
	testAddenda98Parse(t)
}

func BenchmarkAddenda98Parse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98Parse(b)
	}
}

func testAddenda98String(t testing.TB) {
	addenda98 := NewAddenda98()
	line := "798C01099912340000015      091012981918171614                                  091012980000088"
	addenda98.Parse(line)

	if addenda98.String() != line {
		t.Errorf("\n expected: %v\n got     : %v", line, addenda98.String())
	}
}

func TestAddenda98String(t *testing.T) {
	testAddenda98String(t)
}

func BenchmarkAddenda98String(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98String(b)
	}
}

func testAddenda98ValidRecordType(t testing.TB) {
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
func TestAddenda98ValidRecordType(t *testing.T) {
	testAddenda98ValidRecordType(t)
}

func BenchmarkAddenda98ValidRecordType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98ValidRecordType(b)
	}
}

func testAddenda98ValidTypeCode(t testing.TB) {
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

func TestAddenda98ValidTypeCode(t *testing.T) {
	testAddenda98ValidTypeCode(t)
}

func BenchmarkAddenda98ValidTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98ValidTypeCode(b)
	}
}

func testAddenda98ValidCorrectedData(t testing.TB) {
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

func TestAddenda98ValidCorrectedData(t *testing.T) {
	testAddenda98ValidCorrectedData(t)
}

func BenchmarkAddenda98ValidCorrectedData(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98ValidCorrectedData(b)
	}
}

func testAddenda98ValidateTrue(t testing.TB) {
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

func TestAddenda98ValidateTrue(t *testing.T) {
	testAddenda98ValidateTrue(t)
}

func BenchmarkAddenda98ValidateTrue(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98ValidateTrue(b)
	}
}

func testAddenda98ValidateChangeCodeFalse(t testing.TB) {
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

func TestAddenda98ValidateChangeCodeFalse(t *testing.T) {
	testAddenda98ValidateChangeCodeFalse(t)
}

func BenchmarkAddenda98ValidateChangeCodeFalse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98ValidateChangeCodeFalse(b)
	}
}

func testAddenda98OriginalTraceField(t testing.TB) {
	addenda98 := mockAddenda98()
	exp := "000000000012345"
	if addenda98.OriginalTraceField() != exp {
		t.Errorf("expected %v received %v", exp, addenda98.OriginalTraceField())
	}
}

func TestAddenda98OriginalTraceField(t *testing.T) {
	testAddenda98OriginalTraceField(t)
}

func BenchmarkAddenda98OriginalTraceField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98OriginalTraceField(b)
	}
}

func testAddenda98OriginalDFIField(t testing.TB) {
	addenda98 := mockAddenda98()
	exp := "09101298"
	if addenda98.OriginalDFIField() != exp {
		t.Errorf("expected %v received %v", exp, addenda98.OriginalDFIField())
	}
}

func TestAddenda98OriginalDFIField(t *testing.T) {
	testAddenda98OriginalDFIField(t)
}

func BenchmarkAddenda98OriginalDFIField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98OriginalDFIField(b)
	}
}

func testAddenda98CorrectedDataField(t testing.TB) {
	addenda98 := mockAddenda98()
	exp := "1918171614                   " // 29 char
	if addenda98.CorrectedDataField() != exp {
		t.Errorf("expected %v received %v", exp, addenda98.CorrectedDataField())
	}
}

func TestAddenda98CorrectedDataField(t *testing.T) {
	testAddenda98CorrectedDataField(t)
}

func BenchmarkAddenda98CorrectedDataField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98CorrectedDataField(b)
	}
}

func testAddenda98TraceNumberField(t testing.TB) {
	addenda98 := mockAddenda98()
	exp := "091012980000088"
	if addenda98.TraceNumberField() != exp {
		t.Errorf("expected %v received %v", exp, addenda98.TraceNumberField())
	}
}

func TestAddenda98TraceNumberField(t *testing.T) {
	testAddenda98TraceNumberField(t)
}

func BenchmarkAddenda98TraceNumberField(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda98TraceNumberField(b)
	}
}
