// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
)

// Addenda12 is an addenda which provides business transaction information for Addenda Type
// Code 12 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
//
// Addenda12 is mandatory for IAT entries
//
// The Addenda12 record identifies key information related to the Originator of
// the entry.
type Addenda12 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record in the block.
	recordType string
	// TypeCode Addenda12 types code '12'
	TypeCode string `json:"typeCode"`
	// Originator City & State / Province
	// Data elements City and State / Province  should be separated with an asterisk (*) as a delimiter
	// and the field should end with a backslash (\).
	// For example: San FranciscoCA.
	OriginatorCityStateProvince string `json:"originatorCityStateProvince"`
	// Originator Country & Postal Code
	// Data elements must be separated by an asterisk (*) and must end with a backslash (\)
	// For example: US10036\
	OriginatorCountryPostalCode string `json:"originatorCountryPostalCode"`
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

// NewAddenda12 returns a new Addenda12 with default values for none exported fields
func NewAddenda12() *Addenda12 {
	addenda12 := new(Addenda12)
	addenda12.recordType = "7"
	addenda12.TypeCode = "12"
	return addenda12
}

// Parse takes the input record string and parses the Addenda12 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm successful parsing and data validity.
func (addenda12 *Addenda12) Parse(record string) {
	// 1-1 Always "7"
	addenda12.recordType = "7"
	// 2-3 Always 12
	addenda12.TypeCode = record[1:3]
	// 4-38
	addenda12.OriginatorCityStateProvince = strings.TrimSpace(record[3:38])
	// 39-73
	addenda12.OriginatorCountryPostalCode = strings.TrimSpace(record[38:73])
	// 74-87 reserved - Leave blank
	addenda12.reserved = "              "
	// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
	addenda12.EntryDetailSequenceNumber = addenda12.parseNumField(record[87:94])
}

// String writes the Addenda12 struct to a 94 character string.
func (addenda12 *Addenda12) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(addenda12.recordType)
	buf.WriteString(addenda12.TypeCode)
	buf.WriteString(addenda12.OriginatorCityStateProvinceField())
	// ToDo Validator for backslash
	buf.WriteString(addenda12.OriginatorCountryPostalCodeField())
	buf.WriteString(addenda12.reservedField())
	buf.WriteString(addenda12.EntryDetailSequenceNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda12 *Addenda12) Validate() error {
	if err := addenda12.fieldInclusion(); err != nil {
		return err
	}
	if addenda12.recordType != "7" {
		return fieldError("recordType", NewErrRecordType(7), addenda12.recordType)
	}
	if err := addenda12.isTypeCode(addenda12.TypeCode); err != nil {
		return fieldError("TypeCode", err, addenda12.TypeCode)
	}
	// Type Code must be 12
	if addenda12.TypeCode != "12" {
		return fieldError("TypeCode", ErrAddendaTypeCode, addenda12.TypeCode)
	}
	if err := addenda12.isAlphanumeric(addenda12.OriginatorCityStateProvince); err != nil {
		return fieldError("OriginatorCityStateProvince", err, addenda12.OriginatorCityStateProvince)
	}
	if err := addenda12.isAlphanumeric(addenda12.OriginatorCountryPostalCode); err != nil {
		return fieldError("OriginatorCountryPostalCode", err, addenda12.OriginatorCountryPostalCode)
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda12 *Addenda12) fieldInclusion() error {
	if addenda12.recordType == "" {
		return fieldError("recordType", ErrConstructor, addenda12.recordType)
	}
	if addenda12.TypeCode == "" {
		return fieldError("TypeCode", ErrConstructor, addenda12.TypeCode)
	}
	if addenda12.OriginatorCityStateProvince == "" {
		return fieldError("OriginatorCityStateProvince", ErrConstructor, addenda12.OriginatorCityStateProvince)
	}
	if addenda12.OriginatorCountryPostalCode == "" {
		return fieldError("OriginatorCountryPostalCode", ErrConstructor, addenda12.OriginatorCountryPostalCode)
	}
	if addenda12.EntryDetailSequenceNumber == 0 {
		return fieldError("EntryDetailSequenceNumber", ErrConstructor, addenda12.EntryDetailSequenceNumberField())
	}
	return nil
}

// OriginatorCityStateProvinceField gets the OriginatorCityStateProvinceField left padded
func (addenda12 *Addenda12) OriginatorCityStateProvinceField() string {
	return addenda12.alphaField(addenda12.OriginatorCityStateProvince, 35)
}

// OriginatorCountryPostalCodeField gets the OriginatorCountryPostalCode field left padded
func (addenda12 *Addenda12) OriginatorCountryPostalCodeField() string {
	return addenda12.alphaField(addenda12.OriginatorCountryPostalCode, 35)
}

// reservedField gets reserved - blank space
func (addenda12 *Addenda12) reservedField() string {
	return addenda12.alphaField(addenda12.reserved, 14)
}

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda12 *Addenda12) EntryDetailSequenceNumberField() string {
	return addenda12.numericField(addenda12.EntryDetailSequenceNumber, 7)
}
