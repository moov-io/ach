// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"os"
	"strings"
	"testing"
)

// TestDecode is a complete file decoding test.
func TestPPDDebitRead(t *testing.T) {
	f, err := os.Open("./testdata/ppd-debit.ach")
	if err != nil {
		t.Errorf("%s: ", err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()
	if err != nil {
		t.Errorf("Can not ach.read file: %v", err)
	}
	err = r.File.ValidateAll()
	if err != nil {
		t.Errorf("Could not validate entire read file: %v", err)
	}
}

// TestDecode is a complete file decoding test.
func TestPPDDebitFiexedLengthRead(t *testing.T) {
	f, err := os.Open("./testdata/ppd-debit-fixedLength.ach")
	if err != nil {
		t.Errorf("%s: ", err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()
	if err != nil {
		t.Errorf("Can not ach.read file: %v", err)
	}
}

/*
func TestMultiBatchFile(t *testing.T) {
	f, err := os.Open("./testdata/20110805A.ach")
	if err != nil {
		t.Errorf("%s: ", err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()
	if err != nil {
		t.Errorf("Can not ach.read file: %v", err)
	}
}
*/
func TestRecordTypeUnknown(t *testing.T) {
	var line = "301 076401251 0764012510807291511A094101achdestname            companyname                    "
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if !strings.Contains(err.Error(), ErrUnknownRecordType.Error()) {
		t.Errorf("Expected RecordType Error got: %v", err)
	}
}

func TestTwoFileHeaders(t *testing.T) {
	var line = "101 076401251 0764012510807291511A094101achdestname            companyname                    "
	var twoHeaders = line + "\n" + line
	r := NewReader(strings.NewReader(twoHeaders))
	_, err := r.Read()

	if !strings.Contains(err.Error(), ErrFileHeader.Error()) {
		t.Errorf("Expected File Header Error got: %v", err)
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

	if !strings.Contains(err.Error(), ErrFileControl.Error()) {
		t.Errorf("Expected File Control Error got: %v", err)
	}
}

func TestFileLineShort(t *testing.T) {
	var line = "1 line is only 90 characters                                               !"
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if !strings.Contains(err.Error(), ErrRecordLen.Error()) {
		t.Errorf("Unexpected read.Read() error: %v", err)
	}
}

func TestFileLineLong(t *testing.T) {
	var line = "1 line is only 100 characters                                                        !"
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if !strings.Contains(err.Error(), ErrRecordLen.Error()) {
		t.Errorf("Unexpected read.Read() error: %v", err)
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
	if !strings.Contains(err.Error(), ErrValidFieldInclusion.Error()) {
		t.Errorf("Unexpected read.Read() error: %v", err)
	}
}

// TestFileBatchHeaderErr ensure a parse validation error flows back from the parser.
func TestFileBatchHeaderErr(t *testing.T) {
	bh := mockBatchHeader()
	bh.ODFIIdentification = 0
	r := NewReader(strings.NewReader(bh.String()))
	_, err := r.Read()
	if !strings.Contains(err.Error(), ErrValidFieldInclusion.Error()) {
		t.Errorf("Unexpected read.Read() error: %v", err)
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
	if !strings.Contains(err.Error(), "BatchHeader") {
		t.Errorf("Unexpected read.Read() error: %v", err)
	}
}

// TestFileEntryDetailOutsideBatch ensure EntryDetail does not exist outside of Batch
func TestFileEntryDetailOutsideBatch(t *testing.T) {
	ed := mockEntryDetail()
	r := NewReader(strings.NewReader(ed.String()))
	_, err := r.Read()
	if !strings.Contains(err.Error(), ErrEntryOutside.Error()) {
		t.Errorf("Unexpected read.Read() error: %v", err)
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
	if !strings.Contains(err.Error(), ErrValidFieldInclusion.Error()) {
		t.Errorf("Unexpected read.Read() error: %v", err)
	}
}

// TestFileEntryDetailNotPPD validation error populates through the reader
func TestFileEntryDetailNotPPD(t *testing.T) {
	ed := mockEntryDetail()
	ed.CheckDigit = 0
	line := ed.String()
	r := NewReader(strings.NewReader(line))
	r.addCurrentBatch(NewBatchPPD())
	r.currentBatch.SetHeader(mockBatchHeader())
	r.currentBatch.GetHeader().StandardEntryClassCode = "ABCXYZ"
	_, err := r.Read()
	if !strings.Contains(err.Error(), "ABCXYZ") {
		t.Errorf("Unexpected read.Read() error: %v", err)
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
	if !strings.Contains(err.Error(), "Addenda") {
		t.Errorf("Unexpected read.Read() error: %v", err)
	}
}

// TestFileAddendaOutsideBatch validation error populates through the reader
func TestFileAddendaOutsideBatch(t *testing.T) {
	addenda := mockAddenda()
	r := NewReader(strings.NewReader(addenda.String()))
	_, err := r.Read()
	if !strings.Contains(err.Error(), ErrAddendaOutside.Error()) {
		t.Errorf("Unexpected read.Read() error: %v", err)
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
	if !strings.Contains(err.Error(), ErrAddendaNoIndicator.Error()) {
		t.Errorf("Unexpected read.Read() error: %v", err)
	}
}

func TestFileFileControlErr(t *testing.T) {

	fc := mockFileControl()
	fc.BatchCount = 0
	r := NewReader(strings.NewReader(fc.String()))
	_, err := r.Read()
	if !strings.Contains(err.Error(), ErrValidFieldInclusion.Error()) {
		t.Errorf("Unexpected read.Read() error: %v", err)
	}
}

// TestFileLongErr Batch Header Service Class is 000 which does not validate
func TestFileLongErr(t *testing.T) {
	line := "101 076401251 0764012510807291511A094101achdestname            companyname                    5000companyname                         origid    PPDCHECKPAYMT000002080730   1076401250000001"
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if !strings.Contains(err.Error(), ErrServiceClass.Error()) {
		t.Errorf("Unexpected read.Read() error: %v", err)
	}
}
