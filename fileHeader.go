// Copyright 2016 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strings"
	"time"
)

// Errors specific to a File Header Record
var (
	msgRecordType       = "received expecting %d"
	msgRecordSize       = "is not 094"
	msgBlockingFactor   = "is not 10"
	msgFormatCode       = "is not 1"
	msgFileCreationDate = "was created before " + time.Now().String()
)

// FileHeader is a Record designating physical file characteristics and identify
// the origin (sending point) and destination (receiving point) of the entries
// contained in the file. The file header also includes creation date and time
// fields which can be used to uniquely identify a file.
type FileHeader struct {
	// RecordType defines the type of record in the block. headerPos
	recordType string

	// PriorityCode consists of the numerals 01
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
	// converters is composed for ACH to GoLang Converters
	converters
}

// NewFileHeader returns a new FileHeader with default values for none exported fields
func NewFileHeader(params ...FileParam) FileHeader {
	fh := FileHeader{
		recordType:     "1",
		priorityCode:   "01",
		FileIDModifier: "A",
		recordSize:     "094",
		blockingFactor: "10",
		formatCode:     "1",
	}
	if len(params) > 0 {
		fh.ImmediateDestination = fh.parseNumField(params[0].ImmediateDestination)
		fh.ImmediateOrigin = fh.parseNumField(params[0].ImmediateOrigin)
		fh.ImmediateDestinationName = params[0].ImmediateDestinationName
		fh.ImmediateOriginName = params[0].ImmediateOriginName
		fh.ReferenceCode = params[0].ReferenceCode
		fh.FileCreationDate = time.Now()
		return fh
	}
	return fh
}

// Parse takes the input record string and parses the FileHeader values
func (fh *FileHeader) Parse(record string) {
	// (character position 1-1) Always "1"
	fh.recordType = "1"
	// (2-3) Always "01"
	fh.priorityCode = "01"
	// (4-13) A blank space followed by your ODFI's routing number. For example: " 121140399"
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
		msg := fmt.Sprintf(msgRecordType, 1)
		return &FieldError{FieldName: "recordType", Value: fh.recordType, Msg: msg}
	}
	if err := fh.isUpperAlphanumeric(fh.FileIDModifier); err != nil {
		return &FieldError{FieldName: "FileIDModifier", Value: fh.FileIDModifier, Msg: err.Error()}
	}
	if len(fh.FileIDModifier) != 1 {
		msg := fmt.Sprintf(msgValidFieldLength, 1)
		return &FieldError{FieldName: "FileIDModifier", Value: fh.FileIDModifier, Msg: msg}
	}
	if fh.recordSize != "094" {
		return &FieldError{FieldName: "recordSize", Value: fh.recordSize, Msg: msgRecordSize}
	}
	if fh.blockingFactor != "10" {
		return &FieldError{FieldName: "blockingFactor", Value: fh.blockingFactor, Msg: msgBlockingFactor}
	}
	if fh.formatCode != "1" {
		return &FieldError{FieldName: "formatCode", Value: fh.formatCode, Msg: msgFormatCode}
	}
	if err := fh.isAlphanumeric(fh.ImmediateDestinationName); err != nil {
		return &FieldError{FieldName: "ImmediateDestinationName", Value: fh.ImmediateDestinationName, Msg: err.Error()}
	}
	if err := fh.isAlphanumeric(fh.ImmediateOriginName); err != nil {
		return &FieldError{FieldName: "ImmediateOriginName", Value: fh.ImmediateOriginName, Msg: err.Error()}
	}
	if err := fh.isAlphanumeric(fh.ReferenceCode); err != nil {
		return &FieldError{FieldName: "ReferenceCode", Value: fh.ReferenceCode, Msg: err.Error()}
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
		return &FieldError{FieldName: "recordType", Value: fh.recordType, Msg: msgFieldInclusion}
	}
	if fh.ImmediateDestination == 0 {
		return &FieldError{FieldName: "ImmediateDestination", Value: fh.ImmediateDestinationField(), Msg: msgFieldInclusion}
	}
	if fh.ImmediateOrigin == 0 {
		return &FieldError{FieldName: "ImmediateOrigin", Value: fh.ImmediateOriginField(), Msg: msgFieldInclusion}
	}
	if fh.FileCreationDate.IsZero() {
		return &FieldError{FieldName: "FileCreationDate", Value: fh.FileCreationDate.String(), Msg: msgFieldInclusion}
	}
	if fh.FileIDModifier == "" {
		return &FieldError{FieldName: "FileIDModifier", Value: fh.FileIDModifier, Msg: msgFieldInclusion}
	}
	if fh.recordSize == "" {
		return &FieldError{FieldName: "recordSize", Value: fh.recordSize, Msg: msgFieldInclusion}
	}
	if fh.blockingFactor == "" {
		return &FieldError{FieldName: "blockingFactor", Value: fh.blockingFactor, Msg: msgFieldInclusion}
	}
	if fh.formatCode == "" {
		return &FieldError{FieldName: "formatCode", Value: fh.formatCode, Msg: msgFieldInclusion}
	}
	return nil
}

// ImmediateDestinationField gets the immediate destination number with zero padding
func (fh *FileHeader) ImmediateDestinationField() string {
	return " " + fh.numericField(fh.ImmediateDestination, 9)
}

// ImmediateOriginField gets the immediate origin number with 0 padding
func (fh *FileHeader) ImmediateOriginField() string {
	return " " + fh.numericField(fh.ImmediateOrigin, 9)
}

// FileCreationDateField gets the file creation date in YYMMDD format
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
