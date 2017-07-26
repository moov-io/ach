// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"os"
	"strings"
	"testing"
)

func TestParseError(t *testing.T) {
	e := &FieldError{FieldName: "testField", Value: "nil", Msg: "could not parse"}
	err := &ParseError{Line: 63, Err: e}
	if err.Error() != "line:63 *ach.FieldError testField nil could not parse" {
		t.Error("ParseError error string formating has changed")
	}
	err.Record = "TestRecord"
	if err.Error() != "line:63 record:TestRecord *ach.FieldError testField nil could not parse" {
		t.Error("ParseError error string formating has changed")
	}
}

// TestDecode is a complete file decoding test. A canary test
func TestPPDDebitRead(t *testing.T) {
	f, err := os.Open("./testdata/ppd-debit.ach")
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err = r.File.ValidateAll(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestDecode is a complete file decoding test. A canary test
func TestPPDDebitFiexedLengthRead(t *testing.T) {
	f, err := os.Open("./testdata/ppd-debit-fixedLength.ach")
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

func TestRecordTypeUnknown(t *testing.T) {
	var line = "301 076401251 0764012510807291511A094101achdestname            companyname                    "
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FileError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", e, e)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestTwoFileHeaders(t *testing.T) {
	var line = "101 076401251 0764012510807291511A094101achdestname            companyname                    "
	var twoHeaders = line + "\n" + line
	r := NewReader(strings.NewReader(twoHeaders))
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FileError); ok {
			if e.Msg != msgFileControl {
				t.Errorf("%T: %s", e, e)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestTwoFileControls(t *testing.T) {
	var line = "9000001000001000000010005320001000000010500000000000000                                       "
	var twoControls = line + "\n" + line
	r := NewReader(strings.NewReader(twoControls))
	r.addCurrentBatch(NewBatchPPD())
	bc := BatchControl{EntryAddendaCount: 1,
		TotalDebitEntryDollarAmount: 10500,
		EntryHash:                   5320001}
	r.currentBatch.SetControl(&bc)

	r.File.AddBatch(r.currentBatch)
	r.File.Control.EntryHash = 5320001
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FileError); ok {
			if e.Msg != msgFileControl {
				t.Errorf("%T: %s", e, e)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFileLineShort(t *testing.T) {
	var line = "1 line is only 70 characters ........................................!"
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FileError); ok {
			if e.FieldName != "RecordLength" {
				t.Errorf("%T: %s", e, e)
			}
		} else {
			t.Errorf("%T: %s", e, e)
		}
	}
}

func TestFileLineLong(t *testing.T) {
	var line = "1 line is 100 characters ..........................................................................!"
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FileError); ok {
			if e.FieldName != "RecordLength" {
				t.Errorf("%T: %s", e, e)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFileFileHeaderErr ensure a parse validation error flows back from the parser.
func TestFileFileHeaderErr(t *testing.T) {
	fh := mockFileHeader()
	fh.ImmediateOrigin = 0
	r := NewReader(strings.NewReader(fh.String()))
	// necessary to have a file control not nil
	r.File.Control = mockFileControl()
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", e, e)
			}
		}
	} else {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileBatchHeaderErr ensure a parse validation error flows back from the parser.
func TestFileBatchHeaderErr(t *testing.T) {
	bh := mockBatchHeader()
	bh.ODFIIdentification = 0
	r := NewReader(strings.NewReader(bh.String()))
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", e, e)
			}
		}
	} else {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileBatchHeaderErr Error when two batch headers exists in a current batch
func TestFileBatchHeaderDuplicate(t *testing.T) {
	// create a new Batch header string
	bh := mockBatchHeader()
	r := NewReader(strings.NewReader(bh.String()))
	// instantitate a batch header in the reader
	r.addCurrentBatch(NewBatchPPD())
	// read should fail because it is parsing a second batch header and there can only be one.
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", e, e)
			}
		}
	} else {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileEntryDetailOutsideBatch ensure EntryDetail does not exist outside of Batch
func TestFileEntryDetailOutsideBatch(t *testing.T) {
	ed := mockEntryDetail()
	r := NewReader(strings.NewReader(ed.String()))
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FileError); ok {
			if e.Msg != msgFileBatchOutside {
				t.Errorf("%T: %s", e, e)
			}
		}
	} else {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileEntryDetail validation error populates through the reader
func TestFileEntryDetail(t *testing.T) {
	ed := mockEntryDetail()
	ed.CheckDigit = 0
	line := ed.String()
	r := NewReader(strings.NewReader(line))
	r.addCurrentBatch(NewBatchPPD())
	r.currentBatch.SetHeader(mockBatchHeader())
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", e, e)
			}
		}
	} else {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileEntryDetailSECNoSupport validation error populates through the reader
func TestFileEntryDetailSECNoSupport(t *testing.T) {
	ed := mockEntryDetail()
	ed.CheckDigit = 0
	line := ed.String()
	r := NewReader(strings.NewReader(line))
	r.addCurrentBatch(NewBatchPPD())
	r.currentBatch.SetHeader(mockBatchHeader())
	r.currentBatch.GetHeader().StandardEntryClassCode = "ABCXYZ"
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FieldError); ok {
			if e.FieldName != "StandardEntryClassCode" {
				t.Errorf("%T: %s", e, e)
			}
		}
	} else {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileAddenda validation error populates through the reader
func TestFileAddenda(t *testing.T) {
	bh := mockBatchHeader()
	ed := mockEntryDetail()
	addenda := mockAddenda()
	addenda.SequenceNumber = 0
	ed.AddAddenda(addenda)
	line := bh.String() + "\n" + ed.String() + "\n" + ed.Addendums[0].String()
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if err != nil {
		if p, ok := err.(*ParseError); ok {
			if e, ok := p.Err.(*FieldError); ok {
				if e.Msg != msgFieldInclusion {
					t.Errorf("%T: %s", e, e)
				}
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFileAddendaOutsideBatch validation error populates through the reader
func TestFileAddendaOutsideBatch(t *testing.T) {
	addenda := mockAddenda()
	r := NewReader(strings.NewReader(addenda.String()))
	_, err := r.Read()
	if err != nil {
		if p, ok := err.(*ParseError); ok {
			if e, ok := p.Err.(*FileError); ok {
				if e.Msg != msgFileBatchOutside {
					t.Errorf("%T: %s", e, e)
				}
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFileAddendaNoIndicator
func TestFileAddendaNoIndicator(t *testing.T) {
	bh := mockBatchHeader()
	ed := mockEntryDetail()
	addenda := mockAddenda()
	line := bh.String() + "\n" + ed.String() + "\n" + addenda.String()
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FileError); ok {
			if e.FieldName != "AddendaRecordIndicator" {
				t.Errorf("%T: %s", e, e)
			}
		}
	} else {
		t.Errorf("%T: %s", err, err)
	}
}

func TestFileFileControlErr(t *testing.T) {
	fc := mockFileControl()
	fc.BatchCount = 0
	r := NewReader(strings.NewReader(fc.String()))
	_, err := r.Read()
	if p, ok := err.(*ParseError); ok {
		if e, ok := p.Err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", e, e)
			}
		}
	} else {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileLongErr Batch Header Service Class is 000 which does not validate
func TestFileLongErr(t *testing.T) {
	line := "101 076401251 0764012510807291511A094101achdestname            companyname                    5000companyname                         origid    PPDCHECKPAYMT000002080730   1076401250000001"
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if e, ok := err.(*ParseError); ok {
		if e, ok := e.Err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}
