// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strings"
)

// Addenda13 is an addenda which provides business transaction information for Addenda Type
// Code 13 in a machine readable format. It is usually formatted according to ANSI, ASC, X13 Standard.
//
// Addenda13 is mandatory for IAT entries
//
// The Addenda13 contains information related to the financial institution originating the entry.
// For inbound IAT entries, the Fourth Addenda Record must contain information to identify the
// foreign financial institution that is providing the funding and payment instruction for the IAT entry.
type Addenda13 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record in the block.
	recordType string
	// TypeCode Addenda13 types code '13'
	TypeCode string `json:"typeCode"`
	// Originating DFI Name
	// For Outbound IAT Entries, this field must contain the name of the U.S. ODFI.
	// For Inbound IATs: Name of the foreign bank providing funding for the payment transaction
	ODFIName string `json:"ODFIName"`
	// Originating DFI Identification Number Qualifier
	// For Inbound IATs: The 2-digit code that identifies the numbering scheme used in the
	// Foreign DFI Identification Number field:
	// 01 = National Clearing System
	// 02 = BIC Code
	// 03 = IBAN Code
	ODFIIDNumberQualifier string `json:"ODFIIDNumberQualifier"`
	// Originating DFI Identification
	// This field contains the routing number that identifies the U.S. ODFI initiating the entry.
	// For Inbound IATs: This field contains the bank ID number of the Foreign Bank providing funding
	// for the payment transaction.
	ODFIIdentification string `json:"ODFIIdentification"`
	// Originating DFI Branch Country Code
	// USb” = United States
	//(“b” indicates a blank space)
	// For Inbound IATs: This 3 position field contains a 2-character code as approved by the
	// International Organization for Standardization (ISO) used to identify the country in which
	// the branch of the bank that originated the entry is located. Values for other countries can
	// be found on the International Organization for Standardization website: www.iso.org.
	ODFIBranchCountryCode string `json:"ODFIBranchCountryCode"`
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

// NewAddenda13 returns a new Addenda13 with default values for none exported fields
func NewAddenda13() *Addenda13 {
	addenda13 := new(Addenda13)
	addenda13.recordType = "7"
	addenda13.TypeCode = "13"
	return addenda13
}

// Parse takes the input record string and parses the Addenda13 values
func (addenda13 *Addenda13) Parse(record string) {
	// 1-1 Always "7"
	addenda13.recordType = "7"
	// 2-3 Always 13
	addenda13.TypeCode = record[1:3]
	// 4-38 ODFIName
	addenda13.ODFIName = strings.TrimSpace(record[3:38])
	// 39-40 ODFIIDNumberQualifier
	addenda13.ODFIIDNumberQualifier = record[38:40]
	// 41-74 ODFIIdentification
	addenda13.ODFIIdentification = addenda13.parseStringField(record[40:74])
	// 75-77
	addenda13.ODFIBranchCountryCode = strings.TrimSpace(record[74:77])
	// 78-87 reserved - Leave blank
	addenda13.reserved = "          "
	// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
	addenda13.EntryDetailSequenceNumber = addenda13.parseNumField(record[87:94])
}

// String writes the Addenda13 struct to a 94 character string.
func (addenda13 *Addenda13) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(addenda13.recordType)
	buf.WriteString(addenda13.TypeCode)
	buf.WriteString(addenda13.ODFINameField())
	buf.WriteString(addenda13.ODFIIDNumberQualifierField())
	buf.WriteString(addenda13.ODFIIdentificationField())
	buf.WriteString(addenda13.ODFIBranchCountryCodeField())
	buf.WriteString(addenda13.reservedField())
	buf.WriteString(addenda13.EntryDetailSequenceNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda13 *Addenda13) Validate() error {
	if err := addenda13.fieldInclusion(); err != nil {
		return err
	}
	if addenda13.recordType != "7" {
		msg := fmt.Sprintf(msgRecordType, 7)
		return &FieldError{FieldName: "recordType", Value: addenda13.recordType, Msg: msg}
	}
	if err := addenda13.isTypeCode(addenda13.TypeCode); err != nil {
		return &FieldError{FieldName: "TypeCode", Value: addenda13.TypeCode, Msg: err.Error()}
	}
	// Type Code must be 13
	if addenda13.TypeCode != "13" {
		return &FieldError{FieldName: "TypeCode", Value: addenda13.TypeCode, Msg: msgAddendaTypeCode}
	}
	if err := addenda13.isAlphanumeric(addenda13.ODFIName); err != nil {
		return &FieldError{FieldName: "ODFIName",
			Value: addenda13.ODFIName, Msg: err.Error()}
	}
	// Valid ODFI Identification Number Qualifier
	if err := addenda13.isIDNumberQualifier(addenda13.ODFIIDNumberQualifier); err != nil {
		return &FieldError{FieldName: "ODFIIDNumberQualifier",
			Value: addenda13.ODFIIDNumberQualifier, Msg: err.Error()}
	}
	if err := addenda13.isAlphanumeric(addenda13.ODFIIdentification); err != nil {
		return &FieldError{FieldName: "ODFIIdentification",
			Value: addenda13.ODFIIdentification, Msg: err.Error()}
	}
	if err := addenda13.isAlphanumeric(addenda13.ODFIBranchCountryCode); err != nil {
		return &FieldError{FieldName: "ODFIBranchCountryCode",
			Value: addenda13.ODFIBranchCountryCode, Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda13 *Addenda13) fieldInclusion() error {
	if addenda13.recordType == "" {
		return &FieldError{
			FieldName: "recordType",
			Value:     addenda13.recordType,
			Msg:       msgFieldInclusion + ", did you use NewAddenda13()?",
		}
	}
	if addenda13.TypeCode == "" {
		return &FieldError{
			FieldName: "TypeCode",
			Value:     addenda13.TypeCode,
			Msg:       msgFieldInclusion + ", did you use NewAddenda13()?",
		}
	}
	if addenda13.ODFIName == "" {
		return &FieldError{
			FieldName: "ODFIName",
			Value:     addenda13.ODFIName,
			Msg:       msgFieldInclusion + ", did you use NewAddenda13()?",
		}
	}
	if addenda13.ODFIIDNumberQualifier == "" {
		return &FieldError{
			FieldName: "ODFIIDNumberQualifier",
			Value:     addenda13.ODFIIDNumberQualifier,
			Msg:       msgFieldInclusion + ", did you use NewAddenda13()?",
		}
	}
	if addenda13.ODFIIdentification == "" {
		return &FieldError{
			FieldName: "ODFIIdentification",
			Value:     addenda13.ODFIIdentification,
			Msg:       msgFieldInclusion + ", did you use NewAddenda13()?",
		}
	}
	if addenda13.ODFIBranchCountryCode == "" {
		return &FieldError{
			FieldName: "ODFIBranchCountryCode",
			Value:     addenda13.ODFIBranchCountryCode,
			Msg:       msgFieldInclusion + ", did you use NewAddenda13()?",
		}
	}
	if addenda13.EntryDetailSequenceNumber == 0 {
		return &FieldError{
			FieldName: "EntryDetailSequenceNumber",
			Value:     addenda13.EntryDetailSequenceNumberField(),
			Msg:       msgFieldInclusion + ", did you use NewAddenda13()?",
		}
	}
	return nil
}

// ODFINameField gets the ODFIName field left padded
func (addenda13 *Addenda13) ODFINameField() string {
	return addenda13.alphaField(addenda13.ODFIName, 35)
}

// ODFIIDNumberQualifierField gets the ODFIIDNumberQualifier field left padded
func (addenda13 *Addenda13) ODFIIDNumberQualifierField() string {
	return addenda13.alphaField(addenda13.ODFIIDNumberQualifier, 2)
}

// ODFIIdentificationField gets the ODFIIdentificationCode field left padded
func (addenda13 *Addenda13) ODFIIdentificationField() string {
	return addenda13.alphaField(addenda13.ODFIIdentification, 34)
}

// ODFIBranchCountryCodeField gets the ODFIBranchCountryCode field left padded
func (addenda13 *Addenda13) ODFIBranchCountryCodeField() string {
	return addenda13.alphaField(addenda13.ODFIBranchCountryCode, 3)
}

// reservedField gets reserved - blank space
func (addenda13 *Addenda13) reservedField() string {
	return addenda13.alphaField(addenda13.reserved, 10)
}

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda13 *Addenda13) EntryDetailSequenceNumberField() string {
	return addenda13.numericField(addenda13.EntryDetailSequenceNumber, 7)
}
