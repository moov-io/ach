// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"bufio"
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

func (e *ParseError) Error() string {
	if e.Record == "" {
		return fmt.Sprintf("line:%d %T %s", e.Line, e.Err, e.Err)
	}
	return fmt.Sprintf("line:%d record:%s %T %s", e.Line, e.Record, e.Err, e.Err)
}

// Reader reads records from a ACH-encoded file.
type Reader struct {
	// r handles the IO.Reader sent to be parser.
	scanner *bufio.Scanner
	// file is ach.file model being built as r is parsed.
	File File
	// line is the current line being parsed from the input r
	line string
	// currentBatch is the current Batch entries being parsed
	currentBatch Batcher
	// IATCurrentBatch is the current IATBatch entries being parsed
	IATCurrentBatch IATBatcher
	// line number of the file being parsed
	lineNum int
	// recordName holds the current record name being parsed.
	recordName string
}

// error creates a new ParseError based on err.
func (r *Reader) error(err error) error {
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
func (r *Reader) addIATCurrentBatch(iatBatch IATBatcher) {
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
				return r.File, err
			}
		case lineLength != RecordLength:
			msg := fmt.Sprintf(msgRecordLength, lineLength)
			err := &FileError{FieldName: "RecordLength", Value: strconv.Itoa(lineLength), Msg: msg}
			return r.File, r.error(err)
		default:
			r.line = line
			if err := r.parseLine(); err != nil {
				return r.File, err
			}
		}
	}
	if (FileHeader{}) == r.File.Header {
		// Their must be at least one File Header
		r.recordName = "FileHeader"
		return r.File, r.error(&FileError{Msg: msgFileHeader})
	}
	if (FileControl{}) == r.File.Control {
		// Their must be at least one File Control
		r.recordName = "FileControl"
		return r.File, r.error(&FileError{Msg: msgFileControl})
	}

	return r.File, nil
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
		if err := r.parseAddenda(); err != nil {
			return err
		}
	case batchControlPos:
		if err := r.parseBatchControl(); err != nil {
			return err
		}
		if err := r.currentBatch.Validate(); err != nil {
			r.recordName = "Batches"
			return r.error(err)
		}
		r.File.AddBatch(r.currentBatch)
		r.currentBatch = nil
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
		return r.error(&FileError{FieldName: "recordType", Value: r.line[:1], Msg: msg})
	}
	return nil
}

// parseFileHeader takes the input record string and parses the FileHeaderRecord values
func (r *Reader) parseFileHeader() error {
	r.recordName = "FileHeader"
	if (FileHeader{}) != r.File.Header {
		// Their can only be one File Header per File exit
		r.error(&FileError{Msg: msgFileHeader})
	}
	r.File.Header.Parse(r.line)

	if err := r.File.Header.Validate(); err != nil {
		return r.error(err)
	}
	return nil
}

// parseBatchHeader takes the input record string and parses the FileHeaderRecord values
func (r *Reader) parseBatchHeader() error {
	r.recordName = "BatchHeader"
	if r.currentBatch != nil {
		// batch header inside of current batch
		return r.error(&FileError{Msg: msgFileBatchInside})
	}

	// Ensure we have a valid batch header before building a batch.
	bh := NewBatchHeader()
	bh.Parse(r.line)
	if err := bh.Validate(); err != nil {
		return r.error(err)
	}

	// Passing BatchHeader into NewBatch creates a Batcher of SEC code type.
	batch, err := NewBatch(bh)
	if err != nil {
		return r.error(err)
	}

	r.addCurrentBatch(batch)
	return nil
}

// parseEntryDetail takes the input record string and parses the EntryDetailRecord values
func (r *Reader) parseEntryDetail() error {
	r.recordName = "EntryDetail"

	if r.currentBatch == nil {
		return r.error(&FileError{Msg: msgFileBatchOutside})
	}
	ed := new(EntryDetail)
	ed.Parse(r.line)
	if err := ed.Validate(); err != nil {
		return r.error(err)
	}
	r.currentBatch.AddEntry(ed)
	return nil
}

// parseAddendaRecord takes the input record string and create an Addenda Type appended to the last EntryDetail
func (r *Reader) parseAddenda() error {
	r.recordName = "Addenda"

	if r.currentBatch == nil {
		msg := fmt.Sprint(msgFileBatchOutside)
		return r.error(&FileError{FieldName: "Addenda", Msg: msg})
	}
	if len(r.currentBatch.GetEntries()) == 0 {
		return r.error(&FileError{FieldName: "Addenda", Msg: msgFileBatchOutside})
	}
	entryIndex := len(r.currentBatch.GetEntries()) - 1
	entry := r.currentBatch.GetEntries()[entryIndex]

	if entry.AddendaRecordIndicator == 1 {
		switch r.line[1:3] {
		case "02":
			addenda02 := NewAddenda02()
			addenda02.Parse(r.line)
			if err := addenda02.Validate(); err != nil {
				return r.error(err)
			}
			r.currentBatch.GetEntries()[entryIndex].AddAddenda(addenda02)
		case "05":
			addenda05 := NewAddenda05()
			addenda05.Parse(r.line)
			if err := addenda05.Validate(); err != nil {
				return r.error(err)
			}
			r.currentBatch.GetEntries()[entryIndex].AddAddenda(addenda05)
		case "98":
			addenda98 := NewAddenda98()
			addenda98.Parse(r.line)
			if err := addenda98.Validate(); err != nil {
				return r.error(err)
			}
			r.currentBatch.GetEntries()[entryIndex].AddAddenda(addenda98)
		case "99":
			addenda99 := NewAddenda99()
			addenda99.Parse(r.line)
			if err := addenda99.Validate(); err != nil {
				return r.error(err)
			}
			r.currentBatch.GetEntries()[entryIndex].AddAddenda(addenda99)
		}
	} else {
		msg := fmt.Sprint(msgBatchAddendaIndicator)
		return r.error(&FileError{FieldName: "AddendaRecordIndicator", Msg: msg})
	}

	return nil
}

// parseBatchControl takes the input record string and parses the BatchControlRecord values
func (r *Reader) parseBatchControl() error {
	r.recordName = "BatchControl"
	if r.currentBatch == nil {
		// batch Control without a current batch
		return r.error(&FileError{Msg: msgFileBatchOutside})
	}
	r.currentBatch.GetControl().Parse(r.line)
	if err := r.currentBatch.GetControl().Validate(); err != nil {
		return r.error(err)
	}
	return nil
}

// parseFileControl takes the input record string and parses the FileControlRecord values
func (r *Reader) parseFileControl() error {
	r.recordName = "FileControl"
	if (FileControl{}) != r.File.Control {
		// Can be only one file control per file
		return r.error(&FileError{Msg: msgFileControl})
	}
	r.File.Control.Parse(r.line)
	if err := r.File.Control.Validate(); err != nil {
		return r.error(err)
	}
	return nil
}

// parseIATBatchHeader takes the input record string and parses the FileHeaderRecord values
func (r *Reader) parseIATBatchHeader() error {
	r.recordName = "IATBatchHeader"
	if r.IATCurrentBatch != nil {
		// batch header inside of current batch
		return r.error(&FileError{Msg: msgFileBatchInside})
	}

	// Ensure we have a valid IAT BatchHeader before building a batch.
	bh := NewIATBatchHeader()
	bh.Parse(r.line)
	if err := bh.Validate(); err != nil {
		return r.error(err)
	}

	// Passing BatchHeader into NewBatchIAT creates a Batcher of IAT SEC code type.
	iatBatch, err := IATNewBatch(bh)
	if err != nil {
		return r.error(err)
	}

	r.addIATCurrentBatch(iatBatch)

	return nil
}

// parseIATEntryDetail takes the input record string and parses the EntryDetailRecord values
func (r *Reader) parseIATEntryDetail() error {
	r.recordName = "IATEntryDetail"

	if r.IATCurrentBatch == nil {
		return r.error(&FileError{Msg: msgFileBatchOutside})
	}
	ed := new(IATEntryDetail)
	ed.Parse(r.line)
	if err := ed.Validate(); err != nil {
		return r.error(err)
	}
	r.IATCurrentBatch.AddEntry(ed)
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
	// ToDo: Review if this can be true for domestic files.
	// IATIndicator field
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
