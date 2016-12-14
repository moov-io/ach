// Copyright 2016 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.
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

package ach

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
)

const (
	RecordLength = 94
)

// A ParseError is returned for parsing errors.
// The first line is 1.
type ParseError struct {
	Line    int   // line where the error occurred
	CharPos int   // The character position where the error was found
	Err     error // The actual errror
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("line %d, character position %d: %s", e.Line, e.CharPos, e.Err)
}

// These are the errors that can be returned in Parse.Error
var (
	ErrFieldCount = errors.New("wrong number of fields in record expect 94")
)

// A Reader reads records from a ACH-encoded file.
//
// As returned by NewReader, a Readers expects iunput conforming to NACHA ACH
// file format specification.
//
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
// Thy type of ach returned depends on its contents.
func Decode(r io.Reader) (ach ACH, err error) {
	var d decoder
	return d.decode(r)
}

// Decode reads one record (94 characters of a line) from r.
// If the record has an unexpected formatiing Read returns the record an error
// If there is no data left to be read, Read returns nil, io.EOF.
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
	// TODO: number of lines in file must be divisable by 5
	//fmt.Printf("Number of lines in file: %v \n", d.line)
	return ach, nil
}

func parseFileHeader(record string) (fileHeader FileHeaderRecord, err error) {
	fileHeader.RecordType = record[:1]
	fileHeader.PriorityCode = record[1:3]
	/**
	 * This field contains the Routing Number of the ACH Operator or receiving
	 * point to which the file is being sent. The 10 character field begins with
	 * a blank in the first position, followed by the four digit Federal Reserve
	 * Routing Symbol, the four digit ABA Institution Identifier, and the Check
	 * Digit (bTTTTAAAAC).
	 */
	fileHeader.ImmediateDestination = record[3:13]

	/**
	 * This field contains the Routing Number of the ACH Operator or sending
	 * point that is sending the file. The 10 character field begins with
	 * a blank in the first position, followed by the four digit Federal Reserve
	 * Routing Symbol, the four digit ABA Institution Identifier, and the Check
	 * Digit (bTTTTAAAAC).
	 */
	fileHeader.ImmediateOrigin = record[13:23]

	/**
	 * The File Creation Date is expressed in a "YYMMDD" format. The File Creation
	 * Date is the date on which the file is prepared by an ODFI (ACH input files)
	 * or the date (exchange date) on which a file is transmitted from ACH Operator
	 * to ACH Operator, or from ACH Operator to RDFIs (ACH output files).
	 */
	fileHeader.FileCreationDate = record[23:29]

	/**
	 * The File Creation Time is expressed ina n "HHMM" (24 hour clock) format.
	 */
	fileHeader.FileCreationTime = record[29:33]

	fileHeader.FileIdModifier = record[33:34]

	/**
	 * The Record Size Field indicates the number of characters contained in each
	 * record. At this time, the value "094" must be used.
	 */
	fileHeader.RecordSize = record[34:37]

	/**
	 * The Blocking Factor defines the number of physical records within a block
	 * (a block is 940 characters). For all files moving between a DFI and an ACH
	 * Operator (either way), the value "10" must be used. If the number of records
	 * within the file is not a multiple of ten, the remainder of the block must
	 * be nine-filled.
	 */
	fileHeader.BlockingFactor = record[37:39]

	/**
	 * This field identifies a code to allow for future format variations. As
	 * currently defined, this field will contain a value of "1".
	 */
	fileHeader.FormatCode = record[39:40]

	/**
	 * This field contains the name of the ACH or receiving point for which that
	 * file is destined.
	 */
	fileHeader.ImmediateDestinationName = record[40:63]

	/**
	 * This field contains the name of the ACH operator or sending point that is sending the file.
	 */
	fileHeader.ImmidiateOriginName = record[63:86]

	/**
	 * This field is reserved for information pertinent to the Originator.
	 */
	fileHeader.ReferenceCode = record[86:94]

	return fileHeader, nil
}
