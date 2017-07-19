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
	return fmt.Sprintf("LineNum:%d, RecordName:%s : %s \n", e.Line, e.Record, e.Err)
}

// These are the errors that can be returned in Parse.Error
// additional errors can occur in the Record
var (
	ErrFileRead           = errors.New("File could not be read")
	ErrRecordLen          = errors.New("Wrong number of characters in record expected 94")
	ErrBatchHeader        = errors.New("None or More than one Batch Headers exist in Batch.")
	ErrBatchControl       = errors.New("No terminating batch control record found in file for previous batch")
	ErrUnknownRecordType  = errors.New("Unhandled Record Type")
	ErrFileHeader         = errors.New("None or more than one File Headers exists")
	ErrFileControl        = errors.New("None or more than one File Control exists")
	ErrEntryOutside       = errors.New("Entry Detail record outside of a batch")
	ErrAddendaOutside     = errors.New("Entry Addenda without a preceding Entry Detail")
	ErrAddendaNoIndicator = errors.New("Addenda without Entry Detail Addenda Inicator")
)

// currently support SEC codes
const (
	ppd = "PPD"
)

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

// NewReader returns a new Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		scanner: bufio.NewScanner(r),
	}
}

// Read reads each line of the ACH file and defines which parser to use based
// on the first character of each line. It also enforces ACH formating rules and returns
// the appropriate error if issues are found.
func (r *Reader) Read() (File, error) {
	r.lineNum = 0
	// read through the entire file
	for r.scanner.Scan() {
		line := r.scanner.Text()
		r.lineNum++
		lineLength := len(line)

		// handle the situation where there are no line breaks
		if r.lineNum == 1 && lineLength > RecordLength && lineLength%RecordLength == 0 {
			if err := r.processFixedWidthFile(&line); err != nil {
				return r.File, r.error(err)
			}
			break
		}

		// Only 94 ASCII characters to a line
		if lineLength != RecordLength {
			return r.File, r.error(ErrRecordLen)
		}

		r.line = line

		if err := r.parseLine(); err != nil {
			return r.File, r.error(err)
		}
	}

	if err := r.scanner.Err(); err != nil {
		return r.File, r.error(ErrFileRead)
	}

	if (FileHeader{}) == r.File.Header {
		// Their must be at least one File Header
		return r.File, r.error(ErrFileHeader)
	}

	if (FileControl{}) == r.File.Control {
		// Their must be at least one File Control
		return r.File, r.error(ErrFileControl)
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
	case headerPos:
		if err := r.parseFileHeader(); err != nil {
			return err
		}
	case batchPos:
		if err := r.parseBatchHeader(); err != nil {
			return err
		}
	case entryDetailPos:
		if err := r.parseEntryDetail(); err != nil {
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
			return err
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
		if err := r.File.Validate(); err != nil {
			return err
		}
	default:
		return r.error(ErrUnknownRecordType)
	}

	return nil
}

// parseFileHeader takes the input record string and parses the FileHeaderRecord values
func (r *Reader) parseFileHeader() error {
	r.recordName = "FileHeader"
	if (FileHeader{}) != r.File.Header {
		// Their can only be one File Header per File exit
		return ErrFileHeader
	}
	r.File.Header.Parse(r.line)

	if err := r.File.Header.Validate(); err != nil {
		return err
	}
	return nil
}

// parseBatchHeader takes the input record string and parses the FileHeaderRecord values
func (r *Reader) parseBatchHeader() error {
	r.recordName = "BatchHeader"
	if r.currentBatch != nil {
		return ErrBatchHeader
	}
	bh := NewBatchHeader()
	bh.Parse(r.line)
	switch sec := bh.StandardEntryClassCode; sec {
	case ppd:
		r.addCurrentBatch(NewBatchPPD())
	// @TODO add additional batch types to creation
	default:
		return errors.New("Support for Batch's of SEC(standard entry class): " +
			sec + ", has not been implemented")
	}
	r.currentBatch.SetHeader(bh)

	if err := r.currentBatch.GetHeader().Validate(); err != nil {
		return err
	}
	return nil
}

// parseEntryDetail takes the input record string and parses the EntryDetailRecord values
func (r *Reader) parseEntryDetail() error {
	r.recordName = "EntryDetail"
	if r.currentBatch == nil {
		return ErrEntryOutside
	}
	// @TODO change to EntryDetailer once interface is implmented if EntryDetails change
	ed := new(EntryDetail)

	switch sec := r.currentBatch.GetHeader().StandardEntryClassCode; sec {
	case ppd:
		ed.Parse(r.line)
		if err := ed.Validate(); err != nil {
			return err
		}
	//case "CCD":
	default:
		return errors.New("Support for EntryDetail of SEC(standard entry class): " +
			sec + ", has not been implemented")
	}

	r.currentBatch.AddEntry(ed)
	return nil
}

// parseAddendaRecord takes the input record string and parses the AddendaRecord values
func (r *Reader) parseAddenda() error {
	r.recordName = "Addenda"

	if r.currentBatch == nil {
		return ErrAddendaOutside
	}
	if len(r.currentBatch.GetEntries()) == 0 {
		return ErrAddendaOutside
	}
	entryIndex := len(r.currentBatch.GetEntries()) - 1
	entry := r.currentBatch.GetEntries()[entryIndex]

	switch sec := r.currentBatch.GetHeader().StandardEntryClassCode; sec {
	case ppd:
		if entry.AddendaRecordIndicator == 1 {
			addenda := Addenda{}
			addenda.Parse(r.line)
			if err := addenda.Validate(); err != nil {
				return err
			}
			r.currentBatch.GetEntries()[entryIndex].AddAddenda(addenda)
		} else {
			return ErrAddendaNoIndicator
		}
	// case "CCD":
	default:
		return errors.New("Support for addenda of SEC(standard entry class): " +
			sec + ", has not been implemented")
	}

	return nil
}

// parseBatchControl takes the input record string and parses the BatchControlRecord values
func (r *Reader) parseBatchControl() error {
	r.recordName = "BatchControl"
	if r.currentBatch == nil {
		return ErrBatchHeader
	}

	r.currentBatch.GetControl().Parse(r.line)
	if err := r.currentBatch.GetControl().Validate(); err != nil {
		return err
	}
	return nil
}

// parseFileControl takes the input record string and parses the FileControlRecord values
func (r *Reader) parseFileControl() error {
	r.recordName = "FileControl"
	if (FileControl{}) != r.File.Control {
		// Their can only be one File Control per File exit
		return ErrFileControl
	}
	r.File.Control.Parse(r.line)
	if err := r.File.Control.Validate(); err != nil {
		return err
	}
	return nil
}
