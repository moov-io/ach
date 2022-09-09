// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package ach

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/moov-io/base"
)

var (
	// maxLines is the maximum number of lines a file can have. It is limited by the
	// EntryAddendaCount field which has 8 digits, and the BatchCount field which has
	// 6 digits in the File Control Record. So we can have at most the 2 file records,
	// 2 records for each of 10^6 batches, 10^8 entry and addenda records, and 8 lines
	// of 9's to round up to the nearest multiple of 10.
	maxLines = 2 + 2000000 + 100000000 + 8
)

// Reader reads records from an ACH-encoded file.
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
	errors base.ErrorList
}

// error returns a new ParseError based on err
func (r *Reader) parseError(err error) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(*base.ParseError); ok {
		return err
	}
	return &base.ParseError{
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

// SetValidation stores ValidateOpts on the Reader's underlying File which are to be used
// to override the default NACHA validation rules.
func (r *Reader) SetValidation(opts *ValidateOpts) {
	if r == nil || opts == nil {
		return
	}
	r.File.SetValidation(opts)
}

// ReadFile attempts to open a file at path and read the contents before closing
// and returning the parsed ACH File.
func ReadFile(path string) (*File, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("problem reading %s: %v", path, err)
	}
	defer fd.Close()

	file, err := NewReader(fd).Read()
	return &file, err
}

// ReadFiles attempts to open files at the given paths and read the contents
// of each before closing and returning the parsed ACH Files.
func ReadFiles(paths []string) ([]*File, error) {
	var out []*File
	for i := range paths {
		file, err := ReadFile(paths[i])
		if err != nil {
			return nil, err
		}
		if file != nil {
			out = append(out, file)
		}
	}
	return out, nil
}

// NewReader returns a new ACH Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		scanner: bufio.NewScanner(r),
	}
}

// Read reads each line in the underlying io.Reader and returns a File and any errors encountered.
//
// Read enforces ACH formatting rules and the first character of each line determines which parser is used.
//
// The returned File may not be valid. Callers should tabulate the File with File.Create followed by
// File.Validate to ensure it is Nacha compliant.
//
// Invalid files may be rejected by other financial institutions or ACH tools.
func (r *Reader) Read() (File, error) {
	r.lineNum = 0
	// read through the entire file
	for r.scanner.Scan() {
		line := r.scanner.Text()
		r.lineNum++
		if r.lineNum > maxLines {
			r.errors.Add(ErrFileTooLong)
			return r.File, r.errors
		}

		lineLength := utf8.RuneCountInString(line)

		switch {
		case r.lineNum == 1 && lineLength > RecordLength:
			extraChars := lineLength % RecordLength
			if extraChars != 0 {
				err := fmt.Errorf(
					"%d extra character(s) in ACH file: must be %d but found %d",
					extraChars,
					lineLength-extraChars,
					lineLength,
				)
				r.errors.Add(r.parseError(err))
			} else if err := r.processFixedWidthFile(&line); err != nil {
				r.errors.Add(err)
			}
		case lineLength != RecordLength:
			if lineLength > RecordLength {
				line = trimSpacesFromLongLine(line)
			}
			// right-pad the line with spaces
			line, err := rightPadShortLine(line)
			if err != nil {
				r.errors.Add(r.parseError(err))
			}
			r.line = line
			// parse the line
			if err := r.parseLine(); err != nil {
				r.errors.Add(r.parseError(NewRecordWrongLengthErr(lineLength)))
				r.errors.Add(err)
			}

		default:
			r.line = line
			if err := r.parseLine(); err != nil {
				r.errors.Add(err)
			}
		}
	}

	// Add a lingering Batch to the file if there was no BatchControl record.
	// This is common when files just contain a BatchHeader and EntryDetail records.
	if r.currentBatch != nil {
		r.File.AddBatch(r.currentBatch)
		r.currentBatch = nil
	}

	// Carry through any ValidateOpts for this comparison
	if (FileHeader{validateOpts: r.File.validateOpts}) == r.File.Header {
		// There must be at least one File Header
		r.recordName = "FileHeader"
		r.errors.Add(ErrFileHeader)
	}

	if !r.File.IsADV() {
		if (FileControl{}) == r.File.Control {
			// There must be at least one File Control
			r.recordName = "FileControl"
			r.errors.Add(ErrFileControl)
		}
	} else {
		if (ADVFileControl{}) == r.File.ADVControl {
			// There must be at least one File Control
			r.recordName = "FileControl"
			r.errors.Add(ErrFileControl)
		}
	}
	if r.errors.Empty() {
		return r.File, nil
	}
	return r.File, r.errors
}

func trimSpacesFromLongLine(s string) string {
	return strings.TrimSuffix(s[:94], " ")
}

func rightPadShortLine(s string) (string, error) {
	if n := len(s); n > RecordLength {
		return s, NewRecordWrongLengthErr(n)
	}
	return s + strings.Repeat(" ", 94-len(s)), nil
}

func (r *Reader) processFixedWidthFile(line *string) error {
	// It should be safe to parse this byte by byte since ACH files are ASCII only.
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
			batch := r.currentBatch
			r.currentBatch = nil
			batch.SetValidation(r.File.validateOpts)
			r.File.AddBatch(batch)
			if err := batch.Validate(); err != nil {
				r.recordName = "Batches"
				return r.parseError(err)
			}
		} else {
			batch := r.IATCurrentBatch
			r.IATCurrentBatch = IATBatch{}
			r.File.AddIATBatch(batch)
			if err := batch.Validate(); err != nil {
				r.recordName = "Batches"
				return r.parseError(err)
			}
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
		return NewErrUnknownRecordType(r.line[:1])
	}
	return nil
}

// parseBH parses determines whether to parse an IATBatchHeader or BatchHeader
func (r *Reader) parseBH() error {
	if r.line[50:53] == IAT || strings.TrimSpace(r.line[04:20]) == IATCOR {
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
	if r.IATCurrentBatch.Header != nil {
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
	if r.currentBatch != nil && r.currentBatch.GetHeader().CompanyName != IATCOR {
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
	// Pass through any ValidateOpts from the Reader for this comparison
	// as we need to compare the other struct fields (e.g. origin, destination)
	if (FileHeader{validateOpts: r.File.validateOpts}) != r.File.Header {
		// There can only be one File Header per File exit
		return ErrFileHeader
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
		return ErrFileBatchHeaderInsideBatch
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
		return ErrFileEntryOutsideBatch
	}
	if r.currentBatch.GetHeader().StandardEntryClassCode != ADV {
		ed := NewEntryDetail()
		ed.Parse(r.line)
		if err := ed.Validate(); err != nil {
			return r.parseError(err)
		}
		r.currentBatch.AddEntry(ed)
	} else {
		ed := NewADVEntryDetail()
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
		return ErrFileAddendaOutsideBatch
	}

	if r.currentBatch.GetHeader().StandardEntryClassCode != ADV {
		if len(r.currentBatch.GetEntries()) == 0 {
			return ErrFileAddendaOutsideEntry
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
				r.currentBatch.GetEntries()[entryIndex].Category = CategoryNOC
				r.currentBatch.GetEntries()[entryIndex].Addenda98 = addenda98
			case "99":
				// Addenda99, Addenda99Dishonored, Addenda99Contested records both have their code
				// in the same spot, so we need to determine which to parse by the value.
				switch {
				case IsDishonoredReturnCode(r.line[3:6]):
					addenda99Dishonored := NewAddenda99Dishonored()
					addenda99Dishonored.Parse(r.line)
					addenda99Dishonored.SetValidation(r.File.validateOpts)
					if err := addenda99Dishonored.Validate(); err != nil {
						return r.parseError(err)
					}
					r.currentBatch.GetEntries()[entryIndex].Addenda99Dishonored = addenda99Dishonored
					r.currentBatch.GetEntries()[entryIndex].Category = CategoryDishonoredReturn

				case IsContestedReturnCode(r.line[3:6]):
					addenda99Contested := NewAddenda99Contested()
					addenda99Contested.Parse(r.line)
					addenda99Contested.SetValidation(r.File.validateOpts)
					if err := addenda99Contested.Validate(); err != nil {
						return r.parseError(err)
					}
					r.currentBatch.GetEntries()[entryIndex].Addenda99Contested = addenda99Contested
					r.currentBatch.GetEntries()[entryIndex].Category = CategoryDishonoredReturnContested

				default:
					addenda99 := NewAddenda99()
					addenda99.Parse(r.line)
					addenda99.SetValidation(r.File.validateOpts)
					if err := addenda99.Validate(); err != nil {
						return r.parseError(err)
					}
					r.currentBatch.GetEntries()[entryIndex].Addenda99 = addenda99
					r.currentBatch.GetEntries()[entryIndex].Category = CategoryReturn
				}
			}
		} else {
			return r.parseError(r.currentBatch.Error("AddendaRecordIndicator", ErrBatchAddendaIndicator))
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
		return ErrFileAddendaOutsideEntry
	}
	entryIndex := len(r.currentBatch.GetADVEntries()) - 1
	entry := r.currentBatch.GetADVEntries()[entryIndex]

	if entry.AddendaRecordIndicator != 1 {
		return r.parseError(r.currentBatch.Error("AddendaRecordIndicator", ErrBatchAddendaIndicator))
	}
	addenda99 := NewAddenda99()
	addenda99.Parse(r.line)
	if err := addenda99.Validate(); err != nil {
		return r.parseError(err)
	}
	r.currentBatch.GetADVEntries()[entryIndex].Category = CategoryReturn
	r.currentBatch.GetADVEntries()[entryIndex].Addenda99 = addenda99
	return nil
}

// parseBatchControl takes the input record string and parses the BatchControlRecord values
func (r *Reader) parseBatchControl() error {
	r.recordName = "BatchControl"
	if r.currentBatch == nil && r.IATCurrentBatch.GetEntries() == nil {
		// batch Control without a current batch
		return ErrFileBatchControlOutsideBatch
	}
	if r.currentBatch != nil {
		if r.currentBatch.GetHeader().StandardEntryClassCode == ADV {
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
			return ErrFileControl
		}
		r.File.Control.Parse(r.line)
		if err := r.File.Control.Validate(); err != nil {
			return r.parseError(err)
		}
	} else {
		if (ADVFileControl{}) != r.File.ADVControl {
			// Can be only one file control per file
			return ErrFileControl
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
		return ErrFileBatchHeaderInsideBatch
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
		return ErrFileEntryOutsideBatch
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
		return ErrFileAddendaOutsideEntry
	}
	entryIndex := len(r.IATCurrentBatch.GetEntries()) - 1
	entry := r.IATCurrentBatch.GetEntries()[entryIndex]

	if entry.AddendaRecordIndicator == 1 {
		err := r.switchIATAddenda(entryIndex)
		if err != nil {
			return r.parseError(err)
		}
	} else {
		return r.parseError(fieldError("AddendaRecordIndicator", ErrIATBatchAddendaIndicator))
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
