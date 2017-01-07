// Copyright 2016 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// Errors specific to a File Header Record
var (
	ErrRecordType       = errors.New("Wrong Record type")
	ErrIDModifier       = errors.New("File ID Modifier is not uppercase A-Z or 0-9")
	ErrRecordSize       = errors.New("Record size is not 094")
	ErrBlockingFactor   = errors.New("Blocking Factor is not 10")
	ErrFormatCode       = errors.New("Format Code is not 1.")
	ErrFileCreationDate = errors.New("File was created before today")
)

// FileHeader is a Record designating physical file characteristics and identify
// the origin (sending point) and destination (receiving point) of the entries
// contained in the file. The file header also includes creation date and time
// fields which can be used to uniquely identify a file.
type FileHeader struct {
	// RecordType defines the type of record in the block. headerPos
	recordType string

	// PriorityCode conists of the numerals 01
	priorityCode string

	// ImmediateDestination contains the Routing Number of the ACH Operator or receiving
	// point to which the file is being sent. The 10 character field begins with
	// a blank in the first position, followed by the four digit Federal Reserve
	// Routing Symbol, the four digit ABA Institution Identifier, and the Check
	// Digit (bTTTTAAAAC).
	immediateDestination int

	// ImmediateOrigin contains the Routing Number of the ACH Operator or sending
	// point that is sending the file. The 10 character field begins with
	// a blank in the first position, followed by the four digit Federal Reserve
	// Routing Symbol, the four digit ABA Insitution Identifier, and the Check
	// Digit (bTTTTAAAAC).
	immediateOrigin int

	// FileCreationDate is expressed in a "YYMMDD" format. The File Creation
	// Date is the date on which the file is prepared by an ODFI (ACH input files)
	// or the date (exchange date) on which a file is transmitted from ACH Operator
	// to ACH Operator, or from ACH Operator to RDFIs (ACH output files).
	fileCreationDate time.Time

	// FileCreationTime is expressed ina n "HHMM" (24 hour clock) format.
	// The system time when the ACH file was created
	fileCreationTime time.Time

	// This field should start at zero and increment by 1 (up to 9) and then go to
	// letters starting at A through Z for each subsequent file that is created for
	// a single system date. (34-34) 1 numeric 0-9 or uppercase alpha A-Z.
	// I have yet to see this ID not A
	FileIDModifier string

	// RecordSize indicates the number of characters contained in each
	// record. At this time, the value "094" must be used.
	recordSize string

	// BlockingFactor defines the number of physical records within a block
	// (a block is 940 characters). For all files moving between a DFI and an ACH
	// Operator (either way), the value "10" must be used. If the number of records
	// within the file is not a multiple of ten, the remainder of the block must
	// be nine-filled.
	blockingFactor string

	// FormatCode a code to allow for future format variations. As
	// currently defined, this field will contain a value of "1".
	formatCode string

	// ImmediateDestinationName us the name of the ACH or receiving point for which that
	// file is destined. Name corresponding to the ImmediateDestination
	ImmediateDestinationName string

	// ImmidiateOriginName is the name of the ACH operator or sending point that is
	// sending the file. Name corresponding to the ImmediateOrigin
	ImmidiateOriginName string

	// ReferenceCode is reserved for information pertinent to the Originatofh.
	ReferenceCode string
	// Validator is composed for data conversion and validation
	Validator
}

// NewFileHeader returns a new FileHeader with default values for none exported fields
func NewFileHeader() *FileHeader {
	return &FileHeader{
		recordType:     "1",
		priorityCode:   "01",
		FileIDModifier: "A",
		recordSize:     "094",
		blockingFactor: "10",
		formatCode:     "1",
	}
}

// Parse takes the input record string and parses the FileHeader values
func (fh *FileHeader) Parse(record string) {
	// (character position 1-1) Always "1"
	fh.recordType = record[:1]
	// (2-3) Always "01"
	fh.priorityCode = record[1:3]
	// (4-13) A blank space followed by your ODFI's routing numbefh. For example: " 121140399"
	fh.immediateDestination = fh.parseNumField(record[3:13])
	// (14-23) A 10-digit number assigned to you by the ODFI once they approve you to originate ACH files through them
	fh.immediateOrigin = fh.parseNumField(record[13:23])
	// 24-29 Today's date in YYMMDD format
	// must be after todays date.
	fh.fileCreationDate = fh.parseSimpleDate(record[23:29])
	// 30-33 The current time in HHMM format
	fh.fileCreationTime = fh.parseSimpleTime(record[29:33])
	// 35-37 Always "A"
	fh.FileIDModifier = record[33:34]
	// 35-37 always "094"
	fh.recordSize = record[34:37]
	//38-39 always "10"
	fh.blockingFactor = record[37:39]
	//40 always "1"
	fh.formatCode = record[39:40]
	//41-63 The name of the ODFI. example "SILICON VALLEY BANK    "
	fh.ImmediateDestinationName = record[40:63]
	//64-86 ACH operator or sending point that is sending the file
	fh.ImmidiateOriginName = record[63:86]
	//97-94 Optional field that may be used to describe the ACH file for internal accounting purposes
	fh.ReferenceCode = record[86:94]
}

// String writes the FileHeader struct to a 94 character string.
func (fh *FileHeader) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v",
		fh.recordType,
		fh.priorityCode,
		fh.ImmediateDestination(),
		fh.ImmediateOrigin(),
		fh.FileCreationDate(),
		fh.FileCreationTime(),
		fh.FileIDModifier,
		fh.recordSize,
		fh.blockingFactor,
		fh.formatCode,
		fh.rightPad(fh.ImmediateDestinationName, " ", 23),
		fh.rightPad(fh.ImmidiateOriginName, " ", 23),
		fh.rightPad(fh.ReferenceCode, " ", 8),
	)

}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fh *FileHeader) Validate() (bool, error) {

	if fh.recordType != "1" {
		return false, ErrRecordType
	}
	if !fh.isUpperAlphanumeric(fh.FileIDModifier) {
		return false, ErrIDModifier
	}
	if fh.recordSize != "094" {
		return false, ErrRecordSize
	}
	if fh.blockingFactor != "10" {
		return false, ErrBlockingFactor
	}
	if fh.formatCode != "1" {
		return false, ErrFormatCode
	}
	// todo: handle test cases for before date
	/*
		if fh.fileCreationDate.Before(time.Now()) {
			return false, ErrFileCreationDate
		}
	*/
	return true, nil
}

// FileCreationDate gets the file cereation date in YYMMDD format
func (fh *FileHeader) FileCreationDate() string {
	return fh.formatSimpleDate(fh.fileCreationDate)
}

// FileCreationTime gets the file creation time in HHMM format
func (fh *FileHeader) FileCreationTime() string {
	return fh.formatSimpleTime(fh.fileCreationTime)
}

// ImmediateDestination gets the immidiate destination number with zero padding
func (fh *FileHeader) ImmediateDestination() string {
	return " " + fh.leftPad(strconv.Itoa(fh.immediateDestination), "0", 9)
}

// ImmediateOrigin gets the immidiate origen number with 0 padding
func (fh *FileHeader) ImmediateOrigin() string {
	return " " + fh.leftPad(strconv.Itoa(fh.immediateOrigin), "0", 9)
}
