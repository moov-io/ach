// Copyright 2016 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

const (
	// RecordLength character count of each line representing a letter in a file
	RecordLength = 94
)

// ParseError is returned for parsing reader errors.
// The first line is 1.
type ParseError struct {
	Line   int    // Line number where the error accurd
	Record string // Name of the record type being parsed
	Err    error  // The actual error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("line %d, Record  %s: %s", e.Line, e.Record, e.Err)
}

// These are the errors that can be returned in Parse.Error
// additional errors can occur in the Record
var (
	ErrRecordLen         = errors.New("Wrong number of fields in record expect 94")
	ErrBatchControl      = errors.New("No terminating batch control record found in file")
	ErrUnknownRecordType = errors.New("Unhandled Record Type")
	ErrFileHeader        = errors.New("None or more than one File Headers exists")
	ErrFileControl       = errors.New("None or more than one Batch Control exists")
)

// Reader reads records from a ACH-encoded file.
type Reader struct {
	// r handles the IO.Reader sent to be parser.
	r *bufio.Reader
	// line number of the file being parsed
	line int
	// record holds the current record type being parser.
	record string
	// field number of the record currently being parsed
	//field string
}

// error creates a new ParseError based on err.
func (r *Reader) error(err error) error {
	return &ParseError{
		Line:   r.line,
		Record: r.record,
		Err:    err,
	}
}

// NewReader returns a new Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		r: bufio.NewReader(r),
	}
}

// Read reads each line of the ACH file and defines which parser to use based
// on the first character of each line. It also enforces ACH formating rules and returns
// the appropriate error if issues are found.
func (r *Reader) Read() (file File, err error) {
	// read through the entire file
	for {
		line, _, err := r.r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("%s when reading file", err)
				// TODO: return nil, error
			}
		}
		// Only 94 ASCII characters to a line

		if len(line) != RecordLength {
			return file, r.error(ErrRecordLen)
		}
		// TODO: Check that all characters are accepter.

		r.line++

		r.record = string(line)
		switch r.record[:1] {
		case headerPos:
			if (FileHeader{}) != file.FileHeader {
				// Their can only be one File Header per File exit
				return file, r.error(ErrFileHeader)
			}
			file.FileHeader = r.parseFileHeader()
			v, err := file.FileHeader.Validate()
			if !v {
				r.record = "FileHeader"
				return file, r.error(err)
			}
		case batchPos:
			file.BatchHeader = r.parseBatchHeader()
			v, err := file.BatchHeader.Validate()
			if !v {
				r.record = "BatchHeader"
				return file, r.error(err)
			}
		case entryDetailPos:
			file.EntryDetail = r.parseEntryDetail()
			v, err := file.EntryDetail.Validate()
			if !v {
				r.record = "EntryDetail"
				return file, r.error(err)
			}
		case entryAddendaPos:
			file.Addenda = r.parseAddenda()
			v, err := file.Addenda.Validate()
			if !v {
				r.record = "Addenda"
				return file, r.error(err)
			}
		case batchControlPos:
			file.BatchControl = r.parseBatchControl()
			v, err := file.BatchControl.Validate()
			if !v {
				r.record = "BatchControl"
				return file, r.error(err)
			}
		case fileControlPos:
			if (FileControl{}) != file.FileControl {
				// Their can only be one File Control per File exit
				return file, r.error(ErrFileControl)
			}
			file.FileControl = r.parseFileControl()
			v, err := file.FileControl.Validate()
			if !v {
				r.record = "FileControl"
				return file, r.error(err)
			}
		default:
			return file, r.error(ErrUnknownRecordType)
		}

	}

	if (FileHeader{}) == file.FileHeader {
		// Their must be at least one file header
		return file, r.error(ErrFileHeader)
	}
	if (FileControl{}) == file.FileControl {
		// Their must be at least one file control
		return file, r.error(ErrFileControl)
	}

	// TODO: number of lines in file must be divisable by 10 the blocking factor
	//fmt.Printf("Number of lines in file: %v \n", r.line)
	return file, nil
}

// parseFileHeader takes the input record string and parses the FileHeaderRecord values
func (r *Reader) parseFileHeader() (fh FileHeader) {
	fh.Parse(r.record)
	return fh
}

// parseBatchHeader takes the input record string and parses the FileHeaderRecord values
func (r *Reader) parseBatchHeader() (bh BatchHeader) {
	bh.Parse(r.record)
	return bh
}

// parseEntryDetail takes the input record string and parses the EntryDetailRecord values
func (r *Reader) parseEntryDetail() (ed EntryDetail) {
	ed.Parse(r.record)
	return ed
}

// parseAddendaRecord takes the input record string and parses the AddendaRecord values
func (r *Reader) parseAddenda() (addenda Addenda) {
	addenda.Parse(r.record)
	return addenda
}

// parseBatchControl takes the input record string and parses the BatchControlRecord values
func (r *Reader) parseBatchControl() (bc BatchControl) {
	bc.Parse(r.record)
	return bc
}

// parseFileControl takes the input record string and parses the FileControlRecord values
func (r *Reader) parseFileControl() (fc FileControl) {
	fc.Parse(r.record)
	return fc
}
