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
	return fmt.Sprintf("LineNum: %d, RecordName:  %s, Error: %s ", e.Line, e.Record, e.Err)
}

// These are the errors that can be returned in Parse.Error
// additional errors can occur in the Record
var (
	ErrRecordLen          = errors.New("Wrong number of fields in record expect 94")
	ErrBatchControl       = errors.New("No terminating batch control record found in file for previous batch")
	ErrUnknownRecordType  = errors.New("Unhandled Record Type")
	ErrFileHeader         = errors.New("None or more than one File Headers exists")
	ErrFileControl        = errors.New("None or more than one File Control exists")
	ErrEntryOutside       = errors.New("Entry Detail record outside of a batch")
	ErrAddendaOutside     = errors.New("Entry Addenda without a preceeding Entry Detail")
	ErrAddendaNoIndicator = errors.New("Addenda without Entry Detail Addenda Inicator")
)

// currently support SEC codes
const (
	ppd = "PPD"
)

// Reader reads records from a ACH-encoded file.
type Reader struct {
	// r handles the IO.Reader sent to be parser.
	r *bufio.Reader
	// file is ach.file model being built as r is parsed.
	file File
	// line is the current line being parsed from the input r
	line string
	// currentBatch is the current Batch entries being parsed
	currentBatch Batch
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

// NewReader returns a new Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		r: bufio.NewReader(r),
	}
}

// Read reads each line of the ACH file and defines which parser to use based
// on the first character of each line. It also enforces ACH formating rules and returns
// the appropriate error if issues are found.
func (r *Reader) Read() (File, error) {
	// read through the entire file
	for {
		i, _, err := r.r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("%s when reading file", err)
				// TODO: return nil, error
				// "File cannot be read"
			}
		}
		r.line = string(i)
		r.lineNum++
		// Only 94 ASCII characters to a line
		if len(r.line) != RecordLength {
			return r.file, r.error(ErrRecordLen)
		}

		switch r.line[:1] {
		case headerPos:
			err = r.parseFileHeader()
			if err != nil {
				return r.file, err
			}
		case batchPos:
			err = r.parseBatchHeader()
			if err != nil {
				return r.file, err
			}
		case entryDetailPos:
			err = r.parseEntryDetail()
			if err != nil {
				return r.file, err
			}
		case entryAddendaPos:
			err = r.parseAddenda()
			if err != nil {
				return r.file, err
			}
		case batchControlPos:
			err = r.parseBatchControl()
			if err != nil {
				return r.file, err
			}
			if err = r.currentBatch.Validate(); err != nil {
				r.recordName = "Batches"
				return r.file, r.error(err)
			}
			r.file.addBatch(r.currentBatch)
			r.currentBatch = Batch{}
		case fileControlPos:
			err = r.parseFileControl()
			if err != nil {
				return r.file, err
			}
		default:
			return r.file, r.error(ErrUnknownRecordType)
		}
	}

	if (FileHeader{}) == r.file.Header {
		// Their must be at least one file header
		return r.file, r.error(ErrFileHeader)
	}

	if (FileControl{}) == r.file.Control {
		// Their must be at least one file control
		return r.file, r.error(ErrFileControl)
	}

	// TODO: Validate cross Record type values

	// TODO: number of lines in file must be divisable by 10 the blocking factor
	return r.file, nil
}

// parseFileHeader takes the input record string and parses the FileHeaderRecord values
func (r *Reader) parseFileHeader() error {
	r.recordName = "FileHeader"
	if (FileHeader{}) != r.file.Header {
		// Their can only be one File Header per File exit
		return r.error(ErrFileHeader)
	}
	r.file.Header.Parse(r.line)

	if err := r.file.Header.Validate(); err != nil {
		return r.error(err)
	}
	return nil
}

// parseBatchHeader takes the input record string and parses the FileHeaderRecord values
func (r *Reader) parseBatchHeader() error {
	r.recordName = "BatchHeader"
	if (BatchHeader{}) != r.currentBatch.Header {
		// Ensure we have an empty Batch
		return r.error(ErrBatchControl)
	}
	r.currentBatch.Header.Parse(r.line)
	if err := r.currentBatch.Header.Validate(); err != nil {
		return r.error(err)
	}
	return nil
}

// parseEntryDetail takes the input record string and parses the EntryDetailRecord values
func (r *Reader) parseEntryDetail() error {
	r.recordName = "EntryDetail"
	if (BatchHeader{}) == r.currentBatch.Header {
		return r.error(ErrEntryOutside)
	}
	if r.currentBatch.Header.StandardEntryClassCode == ppd {
		ed := EntryDetail{}
		ed.Parse(r.line)
		if err := ed.Validate(); err != nil {
			return r.error(err)
		}
		r.currentBatch.addEntryDetail(ed)
	} else {
		return r.error(errors.New("Support for detail entries of standard entry class " +
			r.currentBatch.Header.StandardEntryClassCode + " has not been implemented"))
	}
	return nil
}

// parseAddendaRecord takes the input record string and parses the AddendaRecord values
func (r *Reader) parseAddenda() error {
	r.recordName = "Addenda"
	if len(r.currentBatch.Entries) == 0 {
		return r.error(ErrAddendaOutside)
	}
	entryIndex := len(r.currentBatch.Entries) - 1
	entry := r.currentBatch.Entries[entryIndex]
	if r.currentBatch.Header.StandardEntryClassCode == ppd {
		if entry.AddendaRecordIndicator == 1 {
			addenda := Addenda{}
			addenda.Parse(r.line)
			if err := addenda.Validate(); err != nil {
				return r.error(err)
			}
			r.currentBatch.Entries[entryIndex].addAddenda(addenda)
		} else {
			return r.error(ErrAddendaNoIndicator)
		}
	} else {
		return r.error(errors.New("Support for Addenda records for standard entry class " +
			r.currentBatch.Header.StandardEntryClassCode + " has not been implemented"))
	}

	return nil
}

// parseBatchControl takes the input record string and parses the BatchControlRecord values
func (r *Reader) parseBatchControl() error {
	r.recordName = "BatchControl"
	r.currentBatch.Control.Parse(r.line)
	if err := r.currentBatch.Control.Validate(); err != nil {
		return r.error(err)
	}
	return nil
}

// parseFileControl takes the input record string and parses the FileControlRecord values
func (r *Reader) parseFileControl() error {
	r.recordName = "FileControl"
	if (FileControl{}) != r.file.Control {
		// Their can only be one File Control per File exit
		return r.error(ErrFileControl)
	}
	r.file.Control.Parse(r.line)
	if err := r.file.Control.Validate(); err != nil {
		return r.error(err)
	}
	return nil
}
