// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strings"
)

// Addenda18 is an addenda which provides business transaction information for Addenda Type
// Code 18 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
//
// Addenda18 is optional for IAT entries
//
// The Addenda18 record identifies information on each Foreign Correspondent Bank involved in the
// processing of the IAT entry. If no Foreign Correspondent Bank is involved,t he record should not be
// included. A maximum of five of these Addenda Records may be included with each IAT entry.
type Addenda18 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record in the block. entryAddenda18 Pos 7
	recordType string
	// TypeCode Addenda18 types code '18'
	TypeCode string `json:"typeCode"`
	// ForeignCorrespondentBankName contains the name of the Foreign Correspondent Bank
	ForeignCorrespondentBankName string `json:"foreignCorrespondentBankName"`
	// Foreign Correspondent Bank Identification Number Qualifier contains a 2-digit code that
	// identifies the numbering scheme used in the Foreign Correspondent Bank Identification Number
	// field. Code values for this field are:
	// “01” = National Clearing System
	// “02” = BIC Code
	// “03” = IBAN Code
	ForeignCorrespondentBankIDNumberQualifier string `json:"foreignCorrespondentBankIDNumberQualifier"`
	// Foreign Correspondent Bank Identification Number contains the bank ID number of the Foreign
	// Correspondent Bank
	ForeignCorrespondentBankIDNumber string `json:"foreignCorrespondentBankIDNumber"`
	// Foreign Correspondent Bank Branch Country Code contains the two-character code, as approved by
	// the International Organization for Standardization (ISO), to identify the country in which the
	// branch of the Foreign Correspondent Bank is located. Values can be found on the International
	// Organization for Standardization website: www.iso.org
	ForeignCorrespondentBankBranchCountryCode string `json:"foreignCorrespondentBankBranchCountryCode"`
	// reserved - Leave blank
	reserved string
	// SequenceNumber is consecutively assigned to each Addenda18 Record following
	// an Entry Detail Record. The first addenda18 sequence number must always
	// be a "1".
	SequenceNumber int `json:"sequenceNumber,omitempty"`
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

// NewAddenda18 returns a new Addenda18 with default values for none exported fields
func NewAddenda18() *Addenda18 {
	addenda18 := new(Addenda18)
	addenda18.recordType = "7"
	addenda18.TypeCode = "18"
	return addenda18
}

// Parse takes the input record string and parses the Addenda18 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm successful parsing and data validity.
func (addenda18 *Addenda18) Parse(record string) {
	// 1-1 Always "7"
	addenda18.recordType = "7"
	// 2-3 Always 18
	addenda18.TypeCode = record[1:3]
	// 4-83 Based on the information entered (04-38) 35 alphanumeric
	addenda18.ForeignCorrespondentBankName = strings.TrimSpace(record[3:38])
	// 39-40  Based on the information entered (39-40) 2 alphanumeric
	// “01” = National Clearing System
	// “02” = BIC Code
	// “03” = IBAN Code
	addenda18.ForeignCorrespondentBankIDNumberQualifier = record[38:40]
	// 41-74 Based on the information entered (41-74) 34 alphanumeric
	addenda18.ForeignCorrespondentBankIDNumber = strings.TrimSpace(record[40:74])
	// 75-77 Based on the information entered (75-77) 3 alphanumeric
	addenda18.ForeignCorrespondentBankBranchCountryCode = strings.TrimSpace(record[74:77])
	// 78-83 - Blank space
	addenda18.reserved = "      "
	// 84-87 SequenceNumber is consecutively assigned to each Addenda18 Record following
	// an Entry Detail Record
	addenda18.SequenceNumber = addenda18.parseNumField(record[83:87])
	// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
	addenda18.EntryDetailSequenceNumber = addenda18.parseNumField(record[87:94])
}

// String writes the Addenda18 struct to a 94 character string.
func (addenda18 *Addenda18) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(addenda18.recordType)
	buf.WriteString(addenda18.TypeCode)
	buf.WriteString(addenda18.ForeignCorrespondentBankNameField())
	buf.WriteString(addenda18.ForeignCorrespondentBankIDNumberQualifierField())
	buf.WriteString(addenda18.ForeignCorrespondentBankIDNumberField())
	buf.WriteString(addenda18.ForeignCorrespondentBankBranchCountryCodeField())
	buf.WriteString(addenda18.reservedField())
	buf.WriteString(addenda18.SequenceNumberField())
	buf.WriteString(addenda18.EntryDetailSequenceNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda18 *Addenda18) Validate() error {
	if err := addenda18.fieldInclusion(); err != nil {
		return err
	}
	if addenda18.recordType != "7" {
		msg := fmt.Sprintf(msgRecordType, 7)
		return &FieldError{FieldName: "recordType", Value: addenda18.recordType, Msg: msg}
	}
	if err := addenda18.isTypeCode(addenda18.TypeCode); err != nil {
		return &FieldError{FieldName: "TypeCode", Value: addenda18.TypeCode, Msg: err.Error()}
	}
	// Type Code must be 18
	if addenda18.TypeCode != "18" {
		return &FieldError{FieldName: "TypeCode", Value: addenda18.TypeCode, Msg: msgAddendaTypeCode}
	}
	if err := addenda18.isAlphanumeric(addenda18.ForeignCorrespondentBankName); err != nil {
		return &FieldError{FieldName: "ForeignCorrespondentBankName", Value: addenda18.ForeignCorrespondentBankName, Msg: err.Error()}
	}
	if err := addenda18.isAlphanumeric(addenda18.ForeignCorrespondentBankIDNumberQualifier); err != nil {
		return &FieldError{FieldName: "ForeignCorrespondentBankIDNumberQualifier", Value: addenda18.ForeignCorrespondentBankIDNumberQualifier, Msg: err.Error()}
	}
	if err := addenda18.isAlphanumeric(addenda18.ForeignCorrespondentBankIDNumber); err != nil {
		return &FieldError{FieldName: "ForeignCorrespondentBankIDNumber", Value: addenda18.ForeignCorrespondentBankIDNumber, Msg: err.Error()}
	}
	if err := addenda18.isAlphanumeric(addenda18.ForeignCorrespondentBankBranchCountryCode); err != nil {
		return &FieldError{FieldName: "ForeignCorrespondentBankBranchCountryCode", Value: addenda18.ForeignCorrespondentBankBranchCountryCode, Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda18 *Addenda18) fieldInclusion() error {
	if addenda18.recordType == "" {
		return &FieldError{
			FieldName: "recordType",
			Value:     addenda18.recordType,
			Msg:       msgFieldInclusion + ", did you use NewAddenda18()?",
		}
	}
	if addenda18.TypeCode == "" {
		return &FieldError{
			FieldName: "TypeCode",
			Value:     addenda18.TypeCode,
			Msg:       msgFieldInclusion + ", did you use NewAddenda18()?",
		}
	}
	if addenda18.ForeignCorrespondentBankName == "" {
		return &FieldError{
			FieldName: "ForeignCorrespondentBankName",
			Value:     addenda18.ForeignCorrespondentBankName,
			Msg:       msgFieldInclusion + ", did you use NewAddenda18()?",
		}
	}
	if addenda18.ForeignCorrespondentBankIDNumberQualifier == "" {
		return &FieldError{
			FieldName: "ForeignCorrespondentBankIDNumberQualifier",
			Value:     addenda18.ForeignCorrespondentBankIDNumberQualifier,
			Msg:       msgFieldInclusion + ", did you use NewAddenda18()?",
		}
	}
	if addenda18.ForeignCorrespondentBankIDNumber == "" {
		return &FieldError{
			FieldName: "ForeignCorrespondentBankIDNumber",
			Value:     addenda18.ForeignCorrespondentBankIDNumber,
			Msg:       msgFieldInclusion + ", did you use NewAddenda18()?",
		}
	}
	if addenda18.ForeignCorrespondentBankBranchCountryCode == "" {
		return &FieldError{
			FieldName: "ForeignCorrespondentBankBranchCountryCode",
			Value:     addenda18.ForeignCorrespondentBankBranchCountryCode,
			Msg:       msgFieldInclusion + ", did you use NewAddenda18()?",
		}
	}
	if addenda18.SequenceNumber == 0 {
		return &FieldError{
			FieldName: "SequenceNumber",
			Value:     addenda18.SequenceNumberField(),
			Msg:       msgFieldInclusion + ", did you use NewAddenda18()?",
		}
	}
	if addenda18.EntryDetailSequenceNumber == 0 {
		return &FieldError{
			FieldName: "EntryDetailSequenceNumber",
			Value:     addenda18.EntryDetailSequenceNumberField(),
			Msg:       msgFieldInclusion + ", did you use NewAddenda18()?",
		}
	}
	return nil
}

// ForeignCorrespondentBankNameField returns a zero padded ForeignCorrespondentBankName string
func (addenda18 *Addenda18) ForeignCorrespondentBankNameField() string {
	return addenda18.alphaField(addenda18.ForeignCorrespondentBankName, 35)
}

// ForeignCorrespondentBankIDNumberQualifierField returns a zero padded ForeignCorrespondentBankIDNumberQualifier string
func (addenda18 *Addenda18) ForeignCorrespondentBankIDNumberQualifierField() string {
	return addenda18.alphaField(addenda18.ForeignCorrespondentBankIDNumberQualifier, 2)
}

// ForeignCorrespondentBankIDNumberField returns a zero padded ForeignCorrespondentBankIDNumber string
func (addenda18 *Addenda18) ForeignCorrespondentBankIDNumberField() string {
	return addenda18.alphaField(addenda18.ForeignCorrespondentBankIDNumber, 34)
}

// ForeignCorrespondentBankBranchCountryCodeField returns a zero padded ForeignCorrespondentBankBranchCountryCode string
func (addenda18 *Addenda18) ForeignCorrespondentBankBranchCountryCodeField() string {
	return addenda18.alphaField(addenda18.ForeignCorrespondentBankBranchCountryCode, 3)
}

// SequenceNumberField returns a zero padded SequenceNumber string
func (addenda18 *Addenda18) SequenceNumberField() string {
	return addenda18.numericField(addenda18.SequenceNumber, 4)
}

// reservedField gets reserved - blank space
func (addenda18 *Addenda18) reservedField() string {
	return addenda18.alphaField(addenda18.reserved, 6)
}

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda18 *Addenda18) EntryDetailSequenceNumberField() string {
	return addenda18.numericField(addenda18.EntryDetailSequenceNumber, 7)
}
