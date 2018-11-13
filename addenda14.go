// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strings"
)

// Addenda14 is an addenda which provides business transaction information for Addenda Type
// Code 14 in a machine readable format. It is usually formatted according to ANSI, ASC, X14 Standard.
//
// Addenda14 is mandatory for IAT entries
//
// The Addenda14 identifies the Receiving financial institution holding the Receiver's account.
type Addenda14 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record in the block.
	recordType string
	// TypeCode Addenda14 types code '14'
	TypeCode string `json:"typeCode"`
	// Receiving DFI Name
	// Name of the Receiver's bank
	RDFIName string `json:"RDFIName"`
	// Receiving DFI Identification Number Qualifier
	// The 2-digit code that identifies the numbering scheme used in the
	// Receiving DFI Identification Number field::
	// 01 = National Clearing System
	// 02 = BIC Code
	// 03 = IBAN Code
	RDFIIDNumberQualifier string `json:"RDFIIDNumberQualifier"`
	// Receiving DFI Identification
	// This field contains the bank identification number of the DFI at which the
	// Receiver maintains his account.
	RDFIIdentification string `json:"RDFIIdentification"`
	// Receiving DFI Branch Country Code
	// USb” = United States
	//(“b” indicates a blank space)
	// This 3 position field contains a 2-character code as approved by the International
	// Organization for Standardization (ISO) used to identify the country in which the
	// branch of the bank that receives the entry is located. Values for other countries can
	// be found on the International Organization for Standardization website: www.iso.org
	RDFIBranchCountryCode string `json:"RDFIBranchCountryCode"`
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

// NewAddenda14 returns a new Addenda14 with default values for none exported fields
func NewAddenda14() *Addenda14 {
	addenda14 := new(Addenda14)
	addenda14.recordType = "7"
	addenda14.TypeCode = "14"
	return addenda14
}

// Parse takes the input record string and parses the Addenda14 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm successful parsing and data validity.
func (addenda14 *Addenda14) Parse(record string) {
	// 1-1 Always "7"
	addenda14.recordType = "7"
	// 2-3 Always 14
	addenda14.TypeCode = record[1:3]
	// 4-38 RDFIName
	addenda14.RDFIName = strings.TrimSpace(record[3:38])
	// 39-40 RDFIIDNumberQualifier
	addenda14.RDFIIDNumberQualifier = record[38:40]
	// 41-74 RDFIIdentification
	addenda14.RDFIIdentification = addenda14.parseStringField(record[40:74])
	// 75-77
	addenda14.RDFIBranchCountryCode = strings.TrimSpace(record[74:77])
	// 78-87 reserved - Leave blank
	addenda14.reserved = "          "
	// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
	addenda14.EntryDetailSequenceNumber = addenda14.parseNumField(record[87:94])
}

// String writes the Addenda14 struct to a 94 character string.
func (addenda14 *Addenda14) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(addenda14.recordType)
	buf.WriteString(addenda14.TypeCode)
	buf.WriteString(addenda14.RDFINameField())
	buf.WriteString(addenda14.RDFIIDNumberQualifierField())
	buf.WriteString(addenda14.RDFIIdentificationField())
	buf.WriteString(addenda14.RDFIBranchCountryCodeField())
	buf.WriteString(addenda14.reservedField())
	buf.WriteString(addenda14.EntryDetailSequenceNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda14 *Addenda14) Validate() error {
	if err := addenda14.fieldInclusion(); err != nil {
		return err
	}
	if addenda14.recordType != "7" {
		msg := fmt.Sprintf(msgRecordType, 7)
		return &FieldError{FieldName: "recordType", Value: addenda14.recordType, Msg: msg}
	}
	if err := addenda14.isTypeCode(addenda14.TypeCode); err != nil {
		return &FieldError{FieldName: "TypeCode", Value: addenda14.TypeCode, Msg: err.Error()}
	}
	// Type Code must be 14
	if addenda14.TypeCode != "14" {
		return &FieldError{FieldName: "TypeCode", Value: addenda14.TypeCode, Msg: msgAddendaTypeCode}
	}
	if err := addenda14.isAlphanumeric(addenda14.RDFIName); err != nil {
		return &FieldError{FieldName: "RDFIName",
			Value: addenda14.RDFIName, Msg: err.Error()}
	}
	// Valid RDFI Identification Number Qualifier
	if err := addenda14.isIDNumberQualifier(addenda14.RDFIIDNumberQualifier); err != nil {
		return &FieldError{FieldName: "RDFIIDNumberQualifier",
			Value: addenda14.RDFIIDNumberQualifier, Msg: msgIDNumberQualifier}
	}
	if err := addenda14.isAlphanumeric(addenda14.RDFIIdentification); err != nil {
		return &FieldError{FieldName: "RDFIIdentification",
			Value: addenda14.RDFIIdentification, Msg: err.Error()}
	}
	if err := addenda14.isAlphanumeric(addenda14.RDFIBranchCountryCode); err != nil {
		return &FieldError{FieldName: "RDFIBranchCountryCode",
			Value: addenda14.RDFIBranchCountryCode, Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda14 *Addenda14) fieldInclusion() error {
	if addenda14.recordType == "" {
		return &FieldError{
			FieldName: "recordType",
			Value:     addenda14.recordType,
			Msg:       msgFieldInclusion + ", did you use NewAddenda14()?",
		}
	}
	if addenda14.TypeCode == "" {
		return &FieldError{
			FieldName: "TypeCode",
			Value:     addenda14.TypeCode,
			Msg:       msgFieldInclusion + ", did you use NewAddenda14()?",
		}
	}
	if addenda14.RDFIName == "" {
		return &FieldError{
			FieldName: "RDFIName",
			Value:     addenda14.RDFIName,
			Msg:       msgFieldInclusion + ", did you use NewAddenda14()?",
		}
	}
	if addenda14.RDFIIDNumberQualifier == "" {
		return &FieldError{
			FieldName: "RDFIIDNumberQualifier",
			Value:     addenda14.RDFIIDNumberQualifier,
			Msg:       msgFieldInclusion + ", did you use NewAddenda14()?",
		}
	}
	if addenda14.RDFIIdentification == "" {
		return &FieldError{
			FieldName: "RDFIIdentification",
			Value:     addenda14.RDFIIdentification,
			Msg:       msgFieldInclusion + ", did you use NewAddenda14()?",
		}
	}
	if addenda14.RDFIBranchCountryCode == "" {
		return &FieldError{
			FieldName: "RDFIBranchCountryCode",
			Value:     addenda14.RDFIBranchCountryCode,
			Msg:       msgFieldInclusion + ", did you use NewAddenda14()?",
		}
	}
	if addenda14.EntryDetailSequenceNumber == 0 {
		return &FieldError{
			FieldName: "EntryDetailSequenceNumber",
			Value:     addenda14.EntryDetailSequenceNumberField(),
			Msg:       msgFieldInclusion + ", did you use NewAddenda14()?",
		}
	}
	return nil
}

// RDFINameField gets the RDFIName field left padded
func (addenda14 *Addenda14) RDFINameField() string {
	return addenda14.alphaField(addenda14.RDFIName, 35)
}

// RDFIIDNumberQualifierField gets the RDFIIDNumberQualifier field left padded
func (addenda14 *Addenda14) RDFIIDNumberQualifierField() string {
	return addenda14.alphaField(addenda14.RDFIIDNumberQualifier, 2)
}

// RDFIIdentificationField gets the RDFIIdentificationCode field left padded
func (addenda14 *Addenda14) RDFIIdentificationField() string {
	return addenda14.alphaField(addenda14.RDFIIdentification, 34)
}

// RDFIBranchCountryCodeField gets the RDFIBranchCountryCode field left padded
func (addenda14 *Addenda14) RDFIBranchCountryCodeField() string {
	return addenda14.alphaField(addenda14.RDFIBranchCountryCode, 3)
}

// reservedField gets reserved - blank space
func (addenda14 *Addenda14) reservedField() string {
	return addenda14.alphaField(addenda14.reserved, 10)
}

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda14 *Addenda14) EntryDetailSequenceNumberField() string {
	return addenda14.numericField(addenda14.EntryDetailSequenceNumber, 7)
}
