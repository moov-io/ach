// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strings"
)

// Addenda11 is an addenda which provides business transaction information for Addenda Type
// Code 11 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
//
// Addenda11 is mandatory for IAT entries
//
// The Addenda11 record identifies key information related to the Originator of
// the entry.
type Addenda11 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record in the block.
	recordType string
	// TypeCode Addenda11 types code '11'
	TypeCode string `json:"typeCode"`
	// Originator Name contains the originators name (your company name / name)
	OriginatorName string `json:"originatorName"`
	// Originator Street Address Contains the originators street address (your company's address / your address)
	OriginatorStreetAddress string `json:"originatorStreetAddress"`
	// reserved - Leave blank
	reserved string
	// EntryDetailSequenceNumber contains the ascending sequence number section of the Entry
	// Detail or Corporate Entry Detail Record's trace number This number is
	// the same as the last seven digits of the trace number of the related
	// Entry Detail Record or Corporate Entry Detail Record.
	EntryDetailSequenceNumber int `json:"entryDetailSequenceNumber,omitempty"`
	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters
}

// NewAddenda11 returns a new Addenda11 with default values for none exported fields
func NewAddenda11() *Addenda11 {
	addenda11 := new(Addenda11)
	addenda11.recordType = "7"
	addenda11.TypeCode = "11"
	return addenda11
}

// Parse takes the input record string and parses the Addenda11 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm successful parsing and data validity.
func (addenda11 *Addenda11) Parse(record string) {
	// 1-1 Always "7"
	addenda11.recordType = "7"
	// 2-3 Always 11
	addenda11.TypeCode = record[1:3]
	// 4-38
	addenda11.OriginatorName = strings.TrimSpace(record[3:38])
	// 39-73
	addenda11.OriginatorStreetAddress = strings.TrimSpace(record[38:73])
	// 74-87 reserved - Leave blank
	addenda11.reserved = "              "
	// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
	addenda11.EntryDetailSequenceNumber = addenda11.parseNumField(record[87:94])
}

// String writes the Addenda11 struct to a 94 character string.
func (addenda11 *Addenda11) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(addenda11.recordType)
	buf.WriteString(addenda11.TypeCode)
	buf.WriteString(addenda11.OriginatorNameField())
	buf.WriteString(addenda11.OriginatorStreetAddressField())
	buf.WriteString(addenda11.reservedField())
	buf.WriteString(addenda11.EntryDetailSequenceNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda11 *Addenda11) Validate() error {
	if err := addenda11.fieldInclusion(); err != nil {
		return err
	}
	if addenda11.recordType != "7" {
		msg := fmt.Sprintf(msgRecordType, 7)
		return &FieldError{FieldName: "recordType", Value: addenda11.recordType, Msg: msg}
	}
	if err := addenda11.isTypeCode(addenda11.TypeCode); err != nil {
		return &FieldError{FieldName: "TypeCode", Value: addenda11.TypeCode, Msg: err.Error()}
	}
	// Type Code must be 11
	if addenda11.TypeCode != "11" {
		return &FieldError{FieldName: "TypeCode", Value: addenda11.TypeCode, Msg: msgAddendaTypeCode}
	}
	if err := addenda11.isAlphanumeric(addenda11.OriginatorName); err != nil {
		return &FieldError{FieldName: "OriginatorName", Value: addenda11.OriginatorName, Msg: err.Error()}
	}
	if err := addenda11.isAlphanumeric(addenda11.OriginatorStreetAddress); err != nil {
		return &FieldError{FieldName: "OriginatorStreetAddress", Value: addenda11.OriginatorStreetAddress, Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda11 *Addenda11) fieldInclusion() error {
	if addenda11.recordType == "" {
		return &FieldError{
			FieldName: "recordType",
			Value:     addenda11.recordType,
			Msg:       msgFieldInclusion + ", did you use NewAddenda11()?",
		}
	}
	if addenda11.TypeCode == "" {
		return &FieldError{
			FieldName: "TypeCode",
			Value:     addenda11.TypeCode,
			Msg:       msgFieldInclusion + ", did you use NewAddenda11()?",
		}
	}
	if addenda11.OriginatorName == "" {
		return &FieldError{
			FieldName: "OriginatorName",
			Value:     addenda11.OriginatorName,
			Msg:       msgFieldInclusion + ", did you use NewAddenda11()?",
		}
	}
	if addenda11.OriginatorStreetAddress == "" {
		return &FieldError{
			FieldName: "OriginatorStreetAddress",
			Value:     addenda11.OriginatorStreetAddress,
			Msg:       msgFieldInclusion + ", did you use NewAddenda11()?",
		}
	}
	if addenda11.EntryDetailSequenceNumber == 0 {
		return &FieldError{
			FieldName: "EntryDetailSequenceNumber",
			Value:     addenda11.EntryDetailSequenceNumberField(),
			Msg:       msgFieldInclusion + ", did you use NewAddenda11()?",
		}
	}
	return nil
}

// OriginatorNameField gets the OriginatorName field - Originator Company Name/Individual Name left padded
func (addenda11 *Addenda11) OriginatorNameField() string {
	return addenda11.alphaField(addenda11.OriginatorName, 35)
}

// OriginatorStreetAddressField gets the OriginatorStreetAddress field - Originator Street Address left padded
func (addenda11 *Addenda11) OriginatorStreetAddressField() string {
	return addenda11.alphaField(addenda11.OriginatorStreetAddress, 35)
}

// reservedField gets reserved - blank space
func (addenda11 *Addenda11) reservedField() string {
	return addenda11.alphaField(addenda11.reserved, 14)
}

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda11 *Addenda11) EntryDetailSequenceNumberField() string {
	return addenda11.numericField(addenda11.EntryDetailSequenceNumber, 7)
}
