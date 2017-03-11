// Copyright 2016 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Errors specific to a File Header Record
var (
	ErrRecordType       = errors.New("Wrong Record type")
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
	ImmediateDestination int

	// ImmediateOrigin contains the Routing Number of the ACH Operator or sending
	// point that is sending the file. The 10 character field begins with
	// a blank in the first position, followed by the four digit Federal Reserve
	// Routing Symbol, the four digit ABA Institution Identifier, and the Check
	// Digit (bTTTTAAAAC).
	ImmediateOrigin int

	// FileCreationDate is expressed in a "YYMMDD" format. The File Creation
	// Date is the date on which the file is prepared by an ODFI (ACH input files)
	// or the date (exchange date) on which a file is transmitted from ACH Operator
	// to ACH Operator, or from ACH Operator to RDFIs (ACH output files).
	FileCreationDate time.Time

	// FileCreationTime is expressed ina n "HHMM" (24 hour clock) format.
	// The system time when the ACH file was created
	FileCreationTime time.Time

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

	// ImmediateOriginName is the name of the ACH operator or sending point that is
	// sending the file. Name corresponding to the ImmediateOrigin
	ImmediateOriginName string

	// ReferenceCode is reserved for information pertinent to the Originator.
	ReferenceCode string
	// validator is composed for data validation
	validator
	// converters is composed for ACH to golang Converters
	converters
}

// NewFileHeader returns a new FileHeader with default values for none exported fields
func NewFileHeader() FileHeader {
	return FileHeader{
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
	fh.recordType = "1"
	// (2-3) Always "01"
	fh.priorityCode = "01"
	// (4-13) A blank space followed by your ODFI's routing numbefh. For example: " 121140399"
	fh.ImmediateDestination = fh.parseNumField(record[3:13])
	// (14-23) A 10-digit number assigned to you by the ODFI once they approve you to originate ACH files through them
	fh.ImmediateOrigin = fh.parseNumField(record[13:23])
	// 24-29 Today's date in YYMMDD format
	// must be after todays date.
	fh.FileCreationDate = fh.parseSimpleDate(record[23:29])
	// 30-33 The current time in HHMM format
	fh.FileCreationTime = fh.parseSimpleTime(record[29:33])
	// 35-37 Always "A"
	fh.FileIDModifier = record[33:34]
	// 35-37 always "094"
	fh.recordSize = "094"
	//38-39 always "10"
	fh.blockingFactor = "10"
	//40 always "1"
	fh.formatCode = "1"
	//41-63 The name of the ODFI. example "SILICON VALLEY BANK    "
	fh.ImmediateDestinationName = strings.TrimSpace(record[40:63])
	//64-86 ACH operator or sending point that is sending the file
	fh.ImmediateOriginName = strings.TrimSpace(record[63:86])
	//97-94 Optional field that may be used to describe the ACH file for internal accounting purposes
	fh.ReferenceCode = strings.TrimSpace(record[86:94])
}

// String writes the FileHeader struct to a 94 character string.
func (fh *FileHeader) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v",
		fh.recordType,
		fh.priorityCode,
		fh.ImmediateDestinationField(),
		fh.ImmediateOriginField(),
		fh.FileCreationDateField(),
		fh.FileCreationTimeField(),
		fh.FileIDModifier,
		fh.recordSize,
		fh.blockingFactor,
		fh.formatCode,
		fh.ImmediateDestinationNameField(),
		fh.ImmediateOriginNameField(),
		fh.ReferenceCodeField(),
	)

}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (fh *FileHeader) Validate() error {

	if err := fh.fieldInclusion(); err != nil {
		return err
	}
	if fh.recordType != "1" {
		return &ValidateError{FieldName: "recordType", Value: fh.recordType, Err: ErrRecordType}
	}
	if err := fh.isUpperAlphanumeric(fh.FileIDModifier); err != nil {
		return &ValidateError{FieldName: "FileIDModifier", Value: fh.FileIDModifier, Err: err}
	}
	if len(fh.FileIDModifier) != 1 {
		return &ValidateError{FieldName: "FileIDModifier", Value: fh.FileIDModifier, Err: ErrValidFieldLength}
	}
	if fh.recordSize != "094" {
		return &ValidateError{FieldName: "recordSize", Value: fh.recordSize, Err: ErrRecordSize}
	}
	if fh.blockingFactor != "10" {
		return &ValidateError{FieldName: "blockingFactor", Value: fh.blockingFactor, Err: ErrBlockingFactor}
	}
	if fh.formatCode != "1" {
		return &ValidateError{FieldName: "formatCode", Value: fh.formatCode, Err: ErrFormatCode}
	}
	if err := fh.isAlphanumeric(fh.ImmediateDestinationName); err != nil {
		return &ValidateError{FieldName: "ImmediateDestinationName", Value: fh.ImmediateDestinationName, Err: err}
	}
	if err := fh.isAlphanumeric(fh.ImmediateOriginName); err != nil {
		return &ValidateError{FieldName: "ImmediateOriginName", Value: fh.ImmediateOriginName, Err: err}
	}
	if err := fh.isAlphanumeric(fh.ReferenceCode); err != nil {
		return &ValidateError{FieldName: "ReferenceCode", Value: fh.ReferenceCode, Err: err}
	}

	// todo: handle test cases for before date
	/*
		if fh.fileCreationDate.Before(time.Now()) {
			return false, ErrFileCreationDate
		}
	*/
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (fh *FileHeader) fieldInclusion() error {
	if fh.recordType == "" {
		return &ValidateError{FieldName: "recordType", Value: fh.recordType, Err: ErrRecordType}
	}
	if fh.ImmediateDestination == 0 {
		return &ValidateError{FieldName: "ImmediateDestination", Value: string(fh.ImmediateDestination), Err: ErrValidFieldInclusion}
	}
	if fh.ImmediateOrigin == 0 {
		return &ValidateError{FieldName: "ImmediateOrigin", Value: string(fh.ImmediateOrigin), Err: ErrValidFieldInclusion}
	}
	if fh.FileCreationDate.IsZero() {
		return &ValidateError{FieldName: "FileCreationDate", Value: fh.FileCreationDate.String(), Err: ErrValidFieldInclusion}
	}
	if fh.FileIDModifier == "" {
		return &ValidateError{FieldName: "FileIDModifier", Value: fh.FileIDModifier, Err: ErrRecordType}
	}
	if fh.recordSize == "" {
		return &ValidateError{FieldName: "recordSize", Value: fh.recordSize, Err: ErrRecordType}
	}
	if fh.blockingFactor == "" {
		return &ValidateError{FieldName: "blockingFactor", Value: fh.blockingFactor, Err: ErrRecordType}
	}
	if fh.formatCode == "" {
		return &ValidateError{FieldName: "formatCode", Value: fh.formatCode, Err: ErrRecordType}
	}
	return nil
}

// ImmediateDestinationField gets the immediate destination number with zero padding
func (fh *FileHeader) ImmediateDestinationField() string {
	return " " + fh.numericField(fh.ImmediateDestination, 9)
}

// ImmediateOriginField gets the immediate origen number with 0 padding
func (fh *FileHeader) ImmediateOriginField() string {
	return " " + fh.numericField(fh.ImmediateOrigin, 9)
}

// FileCreationDateField gets the file cereation date in YYMMDD format
func (fh *FileHeader) FileCreationDateField() string {
	return fh.formatSimpleDate(fh.FileCreationDate)
}

// FileCreationTimeField gets the file creation time in HHMM format
func (fh *FileHeader) FileCreationTimeField() string {
	return fh.formatSimpleTime(fh.FileCreationTime)
}

// ImmediateDestinationNameField gets the ImmediateDestinationName field padded
func (fh *FileHeader) ImmediateDestinationNameField() string {
	return fh.alphaField(fh.ImmediateDestinationName, 23)
}

// ImmediateOriginNameField gets the ImmImmediateOriginName field padded
func (fh *FileHeader) ImmediateOriginNameField() string {
	return fh.alphaField(fh.ImmediateOriginName, 23)
}

// ReferenceCodeField gets the ReferenceCode field padded
func (fh *FileHeader) ReferenceCodeField() string {
	return fh.alphaField(fh.ReferenceCode, 8)
}
