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
	"strings"
	"time"
	"unicode/utf8"

	"github.com/moov-io/base"
)

// FileHeader is a Record designating physical file characteristics and identify
// the origin (sending point) and destination (receiving point) of the entries
// contained in the file. The file header also includes creation date and time
// fields which can be used to uniquely identify a file.
type FileHeader struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record in the block. headerPos
	recordType string
	// PriorityCode consists of the numerals 01
	priorityCode string

	// ImmediateDestination contains the Routing Number of the ACH Operator or receiving
	// point to which the file is being sent. The ach file format specifies a 10 character
	// field begins with a blank space in the first position, followed by the four digit
	// Federal Reserve Routing Symbol, the four digit ABA Institution Identifier, and the Check
	// Digit (bTTTTAAAAC). ImmediateDestinationField() will append the blank space to the
	// routing number.
	ImmediateDestination string `json:"immediateDestination"`

	// ImmediateOrigin contains the Routing Number of the ACH Operator or sending
	// point that is sending the file. The ach file format specifies a 10 character
	// field begins with a blank space in the first position, followed by the four digit
	// Federal Reserve Routing Symbol, the four digit ABA Institution Identifier, and the Check
	// Digit (bTTTTAAAAC). ImmediateOriginField() will append the blank space to the
	// routing number.
	ImmediateOrigin string `json:"immediateOrigin"`

	// FileCreationDate is the date on which the file is prepared by an ODFI (ACH input files)
	// or the date (exchange date) on which a file is transmitted from ACH Operator
	// to ACH Operator, or from ACH Operator to RDFIs (ACH output files).
	//
	// The format is: YYMMDD. Y=Year, M=Month, D=Day
	FileCreationDate string `json:"fileCreationDate"`

	// FileCreationTime is the system time when the ACH file was created.
	//
	// The format is: HHmm. H=Hour, m=Minute
	FileCreationTime string `json:"fileCreationTime"`

	// This field should start at zero and increment by 1 (up to 9) and then go to
	// letters starting at A through Z for each subsequent file that is created for
	// a single system date. (34-34) 1 numeric 0-9 or uppercase alpha A-Z.
	// I have yet to see this ID not A
	FileIDModifier string `json:"fileIDModifier,omitempty"`

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
	ImmediateDestinationName string `json:"immediateDestinationName"`

	// ImmediateOriginName is the name of the ACH operator or sending point that is
	// sending the file. Name corresponding to the ImmediateOrigin
	ImmediateOriginName string `json:"immediateOriginName"`

	// ReferenceCode is reserved for information pertinent to the Originator.
	ReferenceCode string `json:"referenceCode,omitempty"`
	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters

	validateOpts *ValidateOpts
}

// NewFileHeader returns a new FileHeader with default values for none exported fields
func NewFileHeader() FileHeader {
	fh := FileHeader{
		recordType:     "1",
		priorityCode:   "01",
		FileIDModifier: "A",
		recordSize:     "094",
		blockingFactor: "10",
		formatCode:     "1",
	}
	return fh
}

// Parse takes the input record string and parses the FileHeader values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm successful parsing and data validity.
func (fh *FileHeader) Parse(record string) {
	if utf8.RuneCountInString(record) != 94 {
		return
	}

	// (character position 1-1) Always "1"
	fh.recordType = "1"
	// (2-3) Always "01"
	fh.priorityCode = "01"
	// (4-13) A blank space followed by your ODFI's routing number. For example: " 121140399"
	fh.ImmediateDestination = fh.parseStringField(record[3:13])
	// (14-23) A 10-digit number assigned to you by the ODFI once they approve you to originate ACH files through them
	fh.ImmediateOrigin = trimImmediateOriginLeadingZero(fh.parseStringField(record[13:23]))
	// 24-29 Today's date in YYMMDD format
	// must be after today's date.
	fh.FileCreationDate = fh.validateSimpleDate(record[23:29])
	// 30-33 The current time in HHmm format
	fh.FileCreationTime = fh.validateSimpleTime(record[29:33])
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

func trimImmediateOriginLeadingZero(s string) string {
	if utf8.RuneCountInString(s) == 10 && s[0] == '0' && s != "0000000000" {
		// trim off a leading 0 as ImmediateOriginField() will pad it back
		return strings.TrimSpace(s[1:])
	}
	return strings.TrimSpace(s)
}

// String writes the FileHeader struct to a 94 character string.
func (fh *FileHeader) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(fh.recordType)
	buf.WriteString(fh.priorityCode)
	buf.WriteString(fh.ImmediateDestinationField())
	buf.WriteString(fh.ImmediateOriginField())
	buf.WriteString(fh.FileCreationDateField())
	buf.WriteString(fh.FileCreationTimeField())
	buf.WriteString(fh.FileIDModifier)
	buf.WriteString(fh.recordSize)
	buf.WriteString(fh.blockingFactor)
	buf.WriteString(fh.formatCode)
	buf.WriteString(fh.ImmediateDestinationNameField())
	buf.WriteString(fh.ImmediateOriginNameField())
	buf.WriteString(fh.ReferenceCodeField())
	return buf.String()
}

// SetValidation stores ValidateOpts on the Batch which are to be used to override
// the default NACHA validation rules.
func (fh *FileHeader) SetValidation(opts *ValidateOpts) {
	if fh == nil {
		return
	}
	fh.validateOpts = opts
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops the parsing.
func (fh *FileHeader) Validate() error {
	return fh.ValidateWith(fh.validateOpts)
}

// ValidateWith performs NACHA format rule checks on each record according to their specification
// overlayed with any custom flags.
// The first error encountered is returned and stops the parsing.
func (fh *FileHeader) ValidateWith(opts *ValidateOpts) error {
	if opts == nil {
		opts = &ValidateOpts{}
	}
	if err := fh.fieldInclusion(); err != nil {
		return err
	}
	if fh.recordType != "1" {
		return fieldError("recordType", NewErrRecordType(1), fh.recordType)
	}
	if err := fh.isUpperAlphanumeric(fh.FileIDModifier); err != nil {
		return fieldError("FileIDModifier", err, fh.FileIDModifier)
	}
	if len(fh.FileIDModifier) != 1 {
		return fieldError("FileIDModifier", NewErrValidFieldLength(1), fh.FileIDModifier)
	}
	if fh.recordSize != "094" {
		return fieldError("recordSize", ErrRecordSize, fh.recordSize)
	}
	if fh.blockingFactor != "10" {
		return fieldError("blockingFactor", ErrBlockingFactor, fh.blockingFactor)
	}
	if fh.formatCode != "1" {
		return fieldError("formatCode", ErrFormatCode, fh.formatCode)
	}
	if err := fh.isAlphanumeric(fh.ImmediateDestinationName); err != nil {
		return fieldError("ImmediateDestinationName", err, fh.ImmediateDestinationName)
	}
	if !opts.BypassOriginValidation {
		if opts.RequireABAOrigin {
			if err := CheckRoutingNumber(fh.ImmediateOrigin); err != nil {
				return fieldError("ImmediateOrigin", ErrConstructor, fh.ImmediateOrigin)
			}
		} else {
			if fh.ImmediateOrigin == "0000000000" {
				return fieldError("ImmediateOrigin", ErrConstructor, fh.ImmediateOrigin)
			}
		}
	}
	if !opts.BypassDestinationValidation {
		if fh.ImmediateDestination == "000000000" {
			return fieldError("ImmediateDestination", ErrConstructor, fh.ImmediateDestination)
		}
		if err := CheckRoutingNumber(fh.ImmediateDestination); err != nil {
			return fieldError("ImmediateDestination", err, fh.ImmediateDestination)
		}
	}
	if err := fh.isAlphanumeric(fh.ImmediateOriginName); err != nil {
		return fieldError("ImmediateOriginName", err, fh.ImmediateOriginName)
	}
	if err := fh.isAlphanumeric(fh.ReferenceCode); err != nil {
		return fieldError("ReferenceCode", err, fh.ReferenceCode)
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
		return fieldError("recordType", ErrConstructor, fh.recordType)
	}
	if fh.ImmediateDestination == "" {
		return fieldError("ImmediateDestination", ErrConstructor, fh.ImmediateDestinationField())
	}
	if fh.ImmediateOrigin == "" {
		return fieldError("ImmediateOrigin", ErrConstructor, fh.ImmediateOriginField())
	}
	if fh.FileCreationDate == "" {
		return fieldError("FileCreationDate", ErrConstructor, fh.FileCreationDate)
	}
	if fh.FileIDModifier == "" {
		return fieldError("FileIDModifier", ErrConstructor, fh.FileIDModifier)
	}
	if fh.recordSize == "" {
		return fieldError("recordSize", ErrConstructor, fh.recordSize)
	}
	if fh.blockingFactor == "" {
		return fieldError("blockingFactor", ErrConstructor, fh.blockingFactor)
	}
	if fh.formatCode == "" {
		return fieldError("formatCode", ErrConstructor, fh.formatCode)
	}
	return nil
}

// ImmediateDestinationField gets the immediate destination number with zero padding
func (fh *FileHeader) ImmediateDestinationField() string {
	return " " + fh.stringField(fh.ImmediateDestination, 9)
}

// ImmediateOriginField gets the immediate origin number with 0 padding
func (fh *FileHeader) ImmediateOriginField() string {
	if fh.ImmediateOrigin == "" {
		return strings.Repeat(" ", 10)
	}
	if fh.validateOpts != nil && fh.validateOpts.BypassOriginValidation {
		return fh.stringField(strings.TrimSpace(fh.ImmediateOrigin), 10)
	}
	return " " + fh.stringField(strings.TrimSpace(fh.ImmediateOrigin), 9)
}

// FileCreationDateField gets the file creation date in YYMMDD (year, month, day) format
// A blank string is returned when an error occurred while parsing the timestamp. ISO 8601
// is the only other format supported.
func (fh *FileHeader) FileCreationDateField() string {
	switch utf8.RuneCountInString(fh.FileCreationDate) {
	case 0:
		return time.Now().Format("060102")
	case 6:
		return fh.formatSimpleDate(fh.FileCreationDate) // YYMMDD
	}
	t, err := time.Parse(base.ISO8601Format, fh.FileCreationDate)
	if err != nil {
		return ""
	}
	return t.Format("060102") // YYMMDD
}

// FileCreationTimeField gets the file creation time in HHmm (hour, minute) format
// A blank string is returned when an error occurred while parsing the timestamp. ISO 8601
// is the only other format supported.
func (fh *FileHeader) FileCreationTimeField() string {
	switch utf8.RuneCountInString(fh.FileCreationTime) {
	case 0:
		return time.Now().Format("1504")
	case 4:
		return fh.formatSimpleTime(fh.FileCreationTime) // HHmm
	}
	t, err := time.Parse(base.ISO8601Format, fh.FileCreationTime)
	if err != nil {
		return ""
	}
	return t.Format("1504") // HHmm
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
