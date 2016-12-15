// Copyright 2016 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.
//
package ach

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
)

// A Reader decodes records from a NACHA ACH ecoded file
//
// Data Specifications Within the ACH File
// • Each line is a Record and consists of 94 characters in
// • Fields within each record type are alphabetic, numeric or alphameric.
// • All alphabetic fields must be left justified and blank padded/filled.
// • All alphabetic characters must be in upper case or "caps".
// • All numeric fields must be right justified, unsigned, and zero padded/filled.
// • All records are 94 characters in length.
// • The file's blocking factor is '10', as indicated in positions 38-39 of the
// File Header '1' record. Every 10 records are a block. If the number of
// records within the file is not a multiple of 10, the remainder of the block
// must be nine filled. The total number of records in your file must be evenly
// divisible by 10.

const (
	RecordLength = 94
)

// A ParseError is returned for parsing errors.
// The first line is 1.
type ParseError struct {
	Line    int   // line where the error occurred
	CharPos int   // The character position where the error was found
	Err     error // The actual error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("line %d, character position %d: %s", e.Line, e.CharPos, e.Err)
}

// These are the errors that can be returned in Parse.Error
var (
	ErrFieldCount = errors.New("wrong number of fields in record expect 94")
)

// A decoder reads records from a ACH-encoded file.
type decoder struct {
	line    int
	charpos int
	r       *bufio.Reader
	// lineBuffer holds the unescaped content read by readRecord
	lineBuffer bytes.Buffer
	//Indexes of the fields inside of lineBuffer
	fieldIndexes []int
	header       FileHeaderRecord
}

// error creates a new ParseError based on err.
func (d *decoder) error(err error) error {
	return &ParseError{
		Line:    d.line,
		CharPos: d.charpos,
		Err:     err,
	}
}

// Decode reads a ACH file from r and returns it as a ach.ACH
func Decode(r io.Reader) (ach ACH, err error) {
	var d decoder
	return d.decode(r)
}

// Decode reads each line of the ACH file and defines which parser to use based
// on the first byte of each line. It also enforces ACH formating rules and returns
// the appropriate error if issues are found.
func (d *decoder) decode(r io.Reader) (ach ACH, err error) {

	d.r = bufio.NewReader(r)

	// read through the entire file
	for {
		line, _, err := d.r.ReadLine()
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
			fmt.Println("character length is not 94: ")
			break
			// TODO: return nil, err
		}
		// TODO: Check that all characters are accepted.

		d.line++

		record := string(line)
		switch record[:1] {
		case FILE_HEADER:
			ach.FileHeader, err = parseFileHeader(record)
			return ach, err
			//fmt.Println("FileHeader")
		case BATCH_HEADER:
			//fmt.Println("BatchHeader")
		case ENTRY_DETAIL:
			//fmt.Println("EntryDetail")
		case BATCH_CONTROL:
			//fmt.Println("BatchControl")
		case FILE_CONTROL:
			//fmt.Println("FileControl")
		default:
			//fmt.Println("Record type not detected")
			// TODO: return nil, error
		}

	}
	// TODO: number of lines in file must be divisable by 10 the blocking factor
	//fmt.Printf("Number of lines in file: %v \n", d.line)
	return ach, nil
}

// parseFileHeader takes the input record string and parses the FileHeaderRecord values
func parseFileHeader(record string) (fileHeader FileHeaderRecord, err error) {
	// (character position 1-1) Always "1"
	fileHeader.RecordType = record[:1]
	// (2-3) Always "01"
	fileHeader.PriorityCode = record[1:3]
	// (4-13) A blank space followed by your ODFI's routing number. For example: " 121140399"
	fileHeader.ImmediateDestination = record[3:13]
	// (14-23) A 10-digit number assigned to you by the ODFI once they approve you to originate ACH files through them
	fileHeader.ImmediateOrigin = record[13:23]
	// 24-29 Today's date in YYMMDD format
	fileHeader.FileCreationDate = record[23:29]
	// 30-33 The current time in HHMM format
	fileHeader.FileCreationTime = record[29:33]
	// 35-37 Always "A"
	fileHeader.FileIdModifier = record[33:34]
	// 35-37 always "094"
	fileHeader.RecordSize = record[34:37]
	//38-39 always "10"
	fileHeader.BlockingFactor = record[37:39]
	//40 always "1"
	fileHeader.FormatCode = record[39:40]
	//41-63 The name of the ODFI. example "SILICON VALLEY BANK    "
	fileHeader.ImmediateDestinationName = record[40:63]
	//64-86 ACH operator or sending point that is sending the file
	fileHeader.ImmidiateOriginName = record[63:86]
	//97-94 Optional field that may be used to describe the ACH file for internal accounting purposes
	fileHeader.ReferenceCode = record[86:94]

	return fileHeader, nil
}
