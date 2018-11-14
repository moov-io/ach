// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

// ParseError is returned for parsing reader errors.
// The first line is 1.
type ParseError struct {
	Line   int    // Line number where the error occurred
	Record string // Name of the record type being parsed
	Err    error  // The actual error
}

func (e ParseError) Error() string {
	if e.Record == "" {
		return fmt.Sprintf("line:%d %T %s", e.Line, e.Err, e.Err)
	}
	return fmt.Sprintf("line:%d record:%s %T %s", e.Line, e.Record, e.Err, e.Err)
}

// ErrorList represents an array of errors which is also an error itself.
type ErrorList []error

// Add appends err onto the ErrorList. Errors are kept in order.
func (e *ErrorList) Add(err error) {
	*e = append(*e, err)
}

// Err returns the first error (or nil).
func (e ErrorList) Err() error {
	if e == nil || len(e) == 0 {
		return nil
	}
	return e[0]
}

// Error implements the error interface
func (e ErrorList) Error() string {
	if len(e) == 0 {
		return "<nil>"
	}
	var buf bytes.Buffer
	e.Print(&buf)
	return buf.String()
}

// Print formats the ErrorList into a string written to w.
// If ErrorList contains multiple errors those after the first
// are indented.
func (e ErrorList) Print(w io.Writer) {
	if w == nil || len(e) == 0 {
		fmt.Fprintf(w, "<nil>")
		return
	}

	fmt.Fprintf(w, "%s", e[0])
	if len(e) > 1 {
		fmt.Fprintf(w, "\n")
	}

	for i := 1; i < len(e); i++ {
		fmt.Fprintf(w, "  %s", e[i])
		if i < len(e)-1 { // don't add \n to last error
			fmt.Fprintf(w, "\n")
		}
	}
}

// Empty no errors to return
func (e ErrorList) Empty() bool {
	return e == nil || len(e) == 0
}

// MarshalJSON marshals error list
func (e ErrorList) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Error())
}

// Reader reads records from a ACH-encoded file.
type Reader struct {
	// file is ach.file model being built as r is parsed.
	File File

	// IATCurrentBatch is the current IATBatch entries being parsed
	IATCurrentBatch IATBatch

	// r handles the IO.Reader sent to be parser.
	scanner *bufio.Scanner

	// line is the current line being parsed from the input r
	line string

	// currentBatch is the current Batch entries being parsed
	currentBatch Batcher

	// line number of the file being parsed
	lineNum int

	// recordName holds the current record name being parsed.
	recordName string

	// errors holds each error encountered when attempting to parse the file
	errors ErrorList
}

// error returns a new ParseError based on err
func (r *Reader) parseError(err error) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(*ParseError); ok {
		return err
	}
	return &ParseError{
		Line:   r.lineNum,
		Record: r.recordName,
		Err:    err,
	}
}

// addCurrentBatch creates the current batch type for the file being read. A successful
// current batch will be added to r.File once parsed.
func (r *Reader) addCurrentBatch(batch Batcher) {
	r.currentBatch = batch
}

// addCurrentBatch creates the current batch type for the file being read. A successful
// current batch will be added to r.File once parsed.
func (r *Reader) addIATCurrentBatch(iatBatch IATBatch) {
	r.IATCurrentBatch = iatBatch
}

// NewReader returns a new ACH Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		scanner: bufio.NewScanner(r),
	}
}

// Read reads each line of the ACH file and defines which parser to use based
// on the first character of each line. It also enforces ACH formatting rules and returns
// the appropriate error if issues are found.
//
// A parsed file may not be valid and callers should confirm with Validate(). Invalid files may
// be rejected by other Financial Institutions or ACH tools.
func (r *Reader) Read() (File, error) {
	r.lineNum = 0
	// read through the entire file
	for r.scanner.Scan() {
		line := r.scanner.Text()
		r.lineNum++
		lineLength := len(line)

		switch {
		case r.lineNum == 1 && lineLength > RecordLength && lineLength%RecordLength == 0:
			if err := r.processFixedWidthFile(&line); err != nil {
				r.errors.Add(err)
			}
		case lineLength != RecordLength:
			msg := fmt.Sprintf(msgRecordLength, lineLength)
			err := &FileError{FieldName: "RecordLength", Value: strconv.Itoa(lineLength), Msg: msg}
			r.errors.Add(r.parseError(err))
		default:
			r.line = line
			if err := r.parseLine(); err != nil {
				r.errors.Add(err)
			}
		}
	}
	if (FileHeader{}) == r.File.Header {
		// There must be at least one File Header
		r.recordName = "FileHeader"
		r.errors.Add(r.parseError(&FileError{Msg: msgFileHeader}))
	}

	if !r.File.IsADV() {
		if (FileControl{}) == r.File.Control {
			// There must be at least one File Control
			r.recordName = "FileControl"
			r.errors.Add(r.parseError(&FileError{Msg: msgFileControl}))
		}
	} else {
		if (ADVFileControl{}) == r.File.ADVControl {
			// There must be at least one File Control
			r.recordName = "FileControl"
			r.errors.Add(r.parseError(&FileError{Msg: msgFileControl}))
		}
	}
	if r.errors.Empty() {
		return r.File, nil
	}
	return r.File, r.errors
}

func (r *Reader) processFixedWidthFile(line *string) error {
	// it should be safe to parse this byte by byte since ACH files are ascii only
	record := ""
	for i, c := range *line {
		record = record + string(c)
		if i > 0 && (i+1)%RecordLength == 0 {
			r.line = record
			if err := r.parseLine(); err != nil {
				return err
			}
			record = ""
		}
	}
	return nil
}

func (r *Reader) parseLine() error {
	switch r.line[:1] {
	case fileHeaderPos:
		if err := r.parseFileHeader(); err != nil {
			return err
		}
	case batchHeaderPos:
		if err := r.parseBH(); err != nil {
			return err
		}
	case entryDetailPos:
		if err := r.parseED(); err != nil {
			return err
		}
	case entryAddendaPos:
		if err := r.parseEDAddenda(); err != nil {
			return err
		}
	case batchControlPos:
		if err := r.parseBatchControl(); err != nil {
			return err
		}
		if r.currentBatch != nil {
			if err := r.currentBatch.Validate(); err != nil {
				r.recordName = "Batches"
				return r.parseError(err)
			}
			r.File.AddBatch(r.currentBatch)
			r.currentBatch = nil
		} else {
			if err := r.IATCurrentBatch.Validate(); err != nil {
				r.recordName = "Batches"
				return r.parseError(err)
			}
			r.File.AddIATBatch(r.IATCurrentBatch)
			r.IATCurrentBatch = IATBatch{}
		}
	case fileControlPos:
		if r.line[:2] == "99" {
			// final blocking padding
			break
		}
		if err := r.parseFileControl(); err != nil {
			return err
		}
	default:
		msg := fmt.Sprintf(msgUnknownRecordType, r.line[:1])
		return r.parseError(&FileError{
			FieldName: "recordType",
			Value:     r.line[:1],
			Msg:       msg,
		})
	}
	return nil
}

// parseBH parses determines whether to parse an IATBatchHeader or BatchHeader
func (r *Reader) parseBH() error {
	if r.line[50:53] == "IAT" {
		if err := r.parseIATBatchHeader(); err != nil {
			return err
		}
	} else {
		if err := r.parseBatchHeader(); err != nil {
			return err
		}
	}
	return nil
}

// parseEd parses determines whether to parse an IATEntryDetail or EntryDetail
func (r *Reader) parseED() error {
	// IAT Indicator field
	if r.line[16:29] == "             " {
		if err := r.parseIATEntryDetail(); err != nil {
			return err
		}
	} else {
		if err := r.parseEntryDetail(); err != nil {
			return err
		}
	}
	return nil
}

// parseEd parses determines whether to parse an IATEntryDetail Addenda or EntryDetail Addenda
func (r *Reader) parseEDAddenda() error {
	if r.currentBatch != nil {
		if err := r.parseAddenda(); err != nil {
			return err
		}
	} else {
		if err := r.parseIATAddenda(); err != nil {
			return err
		}
	}
	return nil
}

// parseFileHeader takes the input record string and parses the FileHeaderRecord values
func (r *Reader) parseFileHeader() error {
	r.recordName = "FileHeader"
	if (FileHeader{}) != r.File.Header {
		// There can only be one File Header per File exit
		return r.parseError(&FileError{Msg: msgFileHeader})
	}
	r.File.Header.Parse(r.line)

	if err := r.File.Header.Validate(); err != nil {
		return r.parseError(err)
	}
	return nil
}

// parseBatchHeader takes the input record string and parses the FileHeaderRecord values
func (r *Reader) parseBatchHeader() error {
	r.recordName = "BatchHeader"
	if r.currentBatch != nil {
		// batch header inside of current batch
		return r.parseError(&FileError{Msg: msgFileBatchInside})
	}

	// Ensure we have a valid batch header before building a batch.
	bh := NewBatchHeader()
	bh.Parse(r.line)
	if err := bh.Validate(); err != nil {
		return r.parseError(err)
	}

	// Passing BatchHeader into NewBatch creates a Batcher of SEC code type.
	batch, err := NewBatch(bh)
	if err != nil {
		return r.parseError(err)
	}

	r.addCurrentBatch(batch)
	return nil
}

// parseEntryDetail takes the input record string and parses the EntryDetailRecord values
func (r *Reader) parseEntryDetail() error {
	r.recordName = "EntryDetail"

	if r.currentBatch == nil {
		return r.parseError(&FileError{Msg: msgFileBatchOutside})
	}
	if r.currentBatch.GetHeader().StandardEntryClassCode != "ADV" {
		ed := new(EntryDetail)
		ed.Parse(r.line)
		if err := ed.Validate(); err != nil {
			return r.parseError(err)
		}
		r.currentBatch.AddEntry(ed)
	} else {
		ed := new(ADVEntryDetail)
		ed.Parse(r.line)
		if err := ed.Validate(); err != nil {
			return r.parseError(err)
		}
		r.currentBatch.AddADVEntry(ed)
	}
	return nil
}

// parseAddendaRecord takes the input record string and create an Addenda Type appended to the last EntryDetail
func (r *Reader) parseAddenda() error {
	r.recordName = "Addenda"

	if r.currentBatch == nil {
		msg := fmt.Sprint(msgFileBatchOutside)
		return r.parseError(&FileError{FieldName: "Addenda", Msg: msg})
	}

	if r.currentBatch.GetHeader().StandardEntryClassCode != "ADV" {
		if len(r.currentBatch.GetEntries()) == 0 {
			return r.parseError(&FileError{FieldName: "Addenda", Msg: msgFileBatchOutside})
		}
		entryIndex := len(r.currentBatch.GetEntries()) - 1
		entry := r.currentBatch.GetEntries()[entryIndex]

		if entry.AddendaRecordIndicator == 1 {
			switch r.line[1:3] {
			case "02":
				addenda02 := NewAddenda02()
				addenda02.Parse(r.line)
				if err := addenda02.Validate(); err != nil {
					return r.parseError(err)
				}
				r.currentBatch.GetEntries()[entryIndex].Addenda02 = addenda02
			case "05":
				addenda05 := NewAddenda05()
				addenda05.Parse(r.line)
				if err := addenda05.Validate(); err != nil {
					return r.parseError(err)
				}
				r.currentBatch.GetEntries()[entryIndex].AddAddenda05(addenda05)
			case "98":
				addenda98 := NewAddenda98()
				addenda98.Parse(r.line)
				if err := addenda98.Validate(); err != nil {
					return r.parseError(err)
				}
				r.currentBatch.GetEntries()[entryIndex].Addenda98 = addenda98
			case "99":
				addenda99 := NewAddenda99()
				addenda99.Parse(r.line)
				if err := addenda99.Validate(); err != nil {
					return r.parseError(err)
				}
				r.currentBatch.GetEntries()[entryIndex].Addenda99 = addenda99
			}
		} else {
			return r.parseError(&FileError{
				FieldName: "AddendaRecordIndicator",
				Msg:       fmt.Sprint(msgBatchAddendaIndicator),
			})
		}
	} else {
		if err := r.parseADVAddenda(); err != nil {
			return err
		}
	}
	return nil
}

// parseADVAddenda takes the input record string and create an Addenda99 appended to the last ADVEntryDetail
func (r *Reader) parseADVAddenda() error {
	if len(r.currentBatch.GetADVEntries()) == 0 {
		return r.parseError(&FileError{FieldName: "Addenda", Msg: msgFileBatchOutside})
	}
	entryIndex := len(r.currentBatch.GetADVEntries()) - 1
	entry := r.currentBatch.GetADVEntries()[entryIndex]

	if entry.AddendaRecordIndicator != 1 {
		return r.parseError(&FileError{
			FieldName: "AddendaRecordIndicator",
			Msg:       fmt.Sprint(msgBatchAddendaIndicator),
		})
	}
	addenda99 := NewAddenda99()
	addenda99.Parse(r.line)
	if err := addenda99.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentBatch.GetADVEntries()[entryIndex].Addenda99 = addenda99
	return nil
}

// parseBatchControl takes the input record string and parses the BatchControlRecord values
func (r *Reader) parseBatchControl() error {
	r.recordName = "BatchControl"
	if r.currentBatch == nil && r.IATCurrentBatch.GetEntries() == nil {
		// batch Control without a current batch
		return r.parseError(&FileError{Msg: msgFileBatchOutside})
	}
	if r.currentBatch != nil {
		if r.currentBatch.GetHeader().StandardEntryClassCode == "ADV" {
			r.currentBatch.GetADVControl().Parse(r.line)
			if err := r.currentBatch.GetADVControl().Validate(); err != nil {
				return r.parseError(err)
			}
		} else {
			r.currentBatch.GetControl().Parse(r.line)
			if err := r.currentBatch.GetControl().Validate(); err != nil {
				return r.parseError(err)
			}
		}
	} else {
		r.IATCurrentBatch.GetControl().Parse(r.line)
		if err := r.IATCurrentBatch.GetControl().Validate(); err != nil {
			return r.parseError(err)
		}

	}
	return nil
}

// parseFileControl takes the input record string and parses the FileControlRecord values
func (r *Reader) parseFileControl() error {
	r.recordName = "FileControl"

	if !r.File.IsADV() {
		if (FileControl{}) != r.File.Control {
			// Can be only one file control per file
			return r.parseError(&FileError{Msg: msgFileControl})
		}
		r.File.Control.Parse(r.line)
		if err := r.File.Control.Validate(); err != nil {
			return r.parseError(err)
		}
	} else {
		if (ADVFileControl{}) != r.File.ADVControl {
			// Can be only one file control per file
			return r.parseError(&FileError{Msg: msgFileControl})
		}
		r.File.ADVControl.Parse(r.line)
		if err := r.File.ADVControl.Validate(); err != nil {
			return r.parseError(err)
		}
	}
	return nil
}

// IAT specific reader functions

// parseIATBatchHeader takes the input record string and parses the FileHeaderRecord values
func (r *Reader) parseIATBatchHeader() error {
	r.recordName = "BatchHeader"
	if r.IATCurrentBatch.Header != nil {
		// batch header inside of current batch
		return r.parseError(&FileError{Msg: msgFileBatchInside})
	}

	// Ensure we have a valid IAT BatchHeader before building a batch.
	bh := NewIATBatchHeader()
	bh.Parse(r.line)
	if err := bh.Validate(); err != nil {
		return r.parseError(err)
	}

	// Passing BatchHeader into NewBatchIAT creates a Batcher of IAT SEC code type.
	iatBatch := NewIATBatch(bh)
	r.addIATCurrentBatch(iatBatch)

	return nil
}

// parseIATEntryDetail takes the input record string and parses the EntryDetailRecord values
func (r *Reader) parseIATEntryDetail() error {
	r.recordName = "EntryDetail"

	if r.IATCurrentBatch.Header == nil {
		return r.parseError(&FileError{Msg: msgFileBatchOutside})
	}

	ed := new(IATEntryDetail)
	ed.Parse(r.line)
	if err := ed.Validate(); err != nil {
		return r.parseError(err)
	}
	r.IATCurrentBatch.AddEntry(ed)
	return nil
}

// parseIATAddenda takes the input record string and create an Addenda Type appended to the last EntryDetail
func (r *Reader) parseIATAddenda() error {
	r.recordName = "Addenda"

	if r.IATCurrentBatch.GetEntries() == nil {
		msg := fmt.Sprint(msgFileBatchOutside)
		return r.parseError(&FileError{FieldName: "Addenda", Msg: msg})
	}
	if len(r.IATCurrentBatch.GetEntries()) == 0 {
		return r.parseError(&FileError{FieldName: "Addenda", Msg: msgFileBatchOutside})
	}
	entryIndex := len(r.IATCurrentBatch.GetEntries()) - 1
	entry := r.IATCurrentBatch.GetEntries()[entryIndex]

	if entry.AddendaRecordIndicator == 1 {
		err := r.switchIATAddenda(entryIndex)
		if err != nil {
			return r.parseError(err)
		}
	} else {
		msg := fmt.Sprint(msgIATBatchAddendaIndicator)
		return r.parseError(&FileError{FieldName: "AddendaRecordIndicator", Msg: msg})
	}
	return nil
}

func (r *Reader) switchIATAddenda(entryIndex int) error {
	switch r.line[1:3] {
	// IAT mandatory and optional Addenda
	case "10", "11", "12", "13", "14", "15", "16", "17", "18":
		err := r.mandatoryOptionalIATAddenda(entryIndex)
		if err != nil {
			return err
		}
	// IATNOC
	case "98":
		err := r.nocIATAddenda(entryIndex)
		if err != nil {
			return err
		}
	// IAT return Addenda
	case "99":
		err := r.returnIATAddenda(entryIndex)
		if err != nil {
			return err
		}
	}
	return nil
}

// mandatoryOptionalIATAddenda parses and validates mandatory IAT addenda records: Addenda10,
// Addenda11, Addenda12, Addenda13, Addenda14, Addenda15, Addenda16, Addenda17, Addenda18
func (r *Reader) mandatoryOptionalIATAddenda(entryIndex int) error {
	switch r.line[1:3] {
	case "10":
		addenda10 := NewAddenda10()
		addenda10.Parse(r.line)
		if err := addenda10.Validate(); err != nil {
			return err
		}
		r.IATCurrentBatch.Entries[entryIndex].Addenda10 = addenda10
	case "11":
		addenda11 := NewAddenda11()
		addenda11.Parse(r.line)
		if err := addenda11.Validate(); err != nil {
			return err
		}
		r.IATCurrentBatch.Entries[entryIndex].Addenda11 = addenda11
	case "12":
		addenda12 := NewAddenda12()
		addenda12.Parse(r.line)
		if err := addenda12.Validate(); err != nil {
			return err
		}
		r.IATCurrentBatch.Entries[entryIndex].Addenda12 = addenda12
	case "13":
		addenda13 := NewAddenda13()
		addenda13.Parse(r.line)
		if err := addenda13.Validate(); err != nil {
			return err
		}
		r.IATCurrentBatch.Entries[entryIndex].Addenda13 = addenda13
	case "14":
		addenda14 := NewAddenda14()
		addenda14.Parse(r.line)
		if err := addenda14.Validate(); err != nil {
			return err
		}
		r.IATCurrentBatch.Entries[entryIndex].Addenda14 = addenda14
	case "15":
		addenda15 := NewAddenda15()
		addenda15.Parse(r.line)
		if err := addenda15.Validate(); err != nil {
			return err
		}
		r.IATCurrentBatch.Entries[entryIndex].Addenda15 = addenda15
	case "16":
		addenda16 := NewAddenda16()
		addenda16.Parse(r.line)
		if err := addenda16.Validate(); err != nil {
			return err
		}
		r.IATCurrentBatch.Entries[entryIndex].Addenda16 = addenda16
	case "17":
		addenda17 := NewAddenda17()
		addenda17.Parse(r.line)
		if err := addenda17.Validate(); err != nil {
			return err
		}
		r.IATCurrentBatch.Entries[entryIndex].AddAddenda17(addenda17)
	case "18":
		addenda18 := NewAddenda18()
		addenda18.Parse(r.line)
		if err := addenda18.Validate(); err != nil {
			return err
		}
		r.IATCurrentBatch.Entries[entryIndex].AddAddenda18(addenda18)
	}
	return nil
}

// nocIATAddenda parses and validates IAT NOC record Addenda98
func (r *Reader) nocIATAddenda(entryIndex int) error {
	addenda98 := NewAddenda98()
	addenda98.Parse(r.line)
	if err := addenda98.Validate(); err != nil {
		return err
	}
	r.IATCurrentBatch.Entries[entryIndex].Addenda98 = addenda98
	return nil
}

// returnIATAddenda parses and validates IAT return record Addenda99
func (r *Reader) returnIATAddenda(entryIndex int) error {
	addenda99 := NewAddenda99()
	addenda99.Parse(r.line)
	if err := addenda99.Validate(); err != nil {
		return err
	}
	r.IATCurrentBatch.Entries[entryIndex].Addenda99 = addenda99
	return nil
}
