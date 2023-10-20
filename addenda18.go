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
	"unicode/utf8"
)

// Addenda18 is an addenda which provides business transaction information for Addenda Type
// Code 18 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
//
// # Addenda18 is optional for IAT entries
//
// The Addenda18 record identifies information on each Foreign Correspondent Bank involved in the
// processing of the IAT entry. If no Foreign Correspondent Bank is involved,t he record should not be
// included. A maximum of five of these Addenda Records may be included with each IAT entry.
type Addenda18 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
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
	// SequenceNumber is consecutively assigned to each Addenda18 Record following
	// an Entry Detail Record. The first addenda18 sequence number must always
	// be a "1".
	SequenceNumber int `json:"sequenceNumber"`
	// EntryDetailSequenceNumber contains the ascending sequence number section of the Entry
	// Detail or Corporate Entry Detail Record's trace number This number is
	// the same as the last seven digits of the trace number of the related
	// Entry Detail Record or Corporate Entry Detail Record.
	EntryDetailSequenceNumber int `json:"entryDetailSequenceNumber"`
	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters
}

// NewAddenda18 returns a new Addenda18 with default values for none exported fields
func NewAddenda18() *Addenda18 {
	addenda18 := new(Addenda18)
	addenda18.TypeCode = "18"
	return addenda18
}

// Parse takes the input record string and parses the Addenda18 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda18 *Addenda18) Parse(record string) {
	runeCount := utf8.RuneCountInString(record)
	if runeCount != 94 {
		return
	}

	buf := getBuffer()
	defer saveBuffer(buf)

	reset := func() string {
		out := buf.String()
		buf.Reset()
		return out
	}

	// We're going to process the record rune-by-rune and at each field cutoff save the value.
	var idx int
	for _, r := range record {
		idx++

		// Append rune to buffer
		buf.WriteRune(r)

		// At each cutoff save the buffer and reset
		switch idx {
		case 0, 1:
			// 1-1 Always 7
			reset()
		case 3:
			// 2-3 Always 18
			addenda18.TypeCode = reset()
		case 38:
			// 4-38 Based on the information entered (04-38) 35 alphanumeric
			addenda18.ForeignCorrespondentBankName = strings.TrimSpace(reset())
		case 40:
			// 39-40  Based on the information entered (39-40) 2 alphanumeric
			// “01” = National Clearing System
			// “02” = BIC Code
			// “03” = IBAN Code
			addenda18.ForeignCorrespondentBankIDNumberQualifier = reset()
		case 74:
			// 41-74 Based on the information entered (41-74) 34 alphanumeric
			addenda18.ForeignCorrespondentBankIDNumber = strings.TrimSpace(reset())
		case 77:
			// 75-77 Based on the information entered (75-77) 3 alphanumeric
			addenda18.ForeignCorrespondentBankBranchCountryCode = strings.TrimSpace(reset())
		case 83:
			// 78-83 - Blank space
			reset()
		case 87:
			// 84-87 SequenceNumber is consecutively assigned to each Addenda18 Record following an Entry Detail Record
			addenda18.SequenceNumber = addenda18.parseNumField(reset())
		case 94:
			// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
			addenda18.EntryDetailSequenceNumber = addenda18.parseNumField(reset())
		}
	}
}

// String writes the Addenda18 struct to a 94 character string.
func (addenda18 *Addenda18) String() string {
	if addenda18 == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(entryAddendaPos)
	buf.WriteString(addenda18.TypeCode)
	buf.WriteString(addenda18.ForeignCorrespondentBankNameField())
	buf.WriteString(addenda18.ForeignCorrespondentBankIDNumberQualifierField())
	buf.WriteString(addenda18.ForeignCorrespondentBankIDNumberField())
	buf.WriteString(addenda18.ForeignCorrespondentBankBranchCountryCodeField())
	buf.WriteString("      ")
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
	if err := addenda18.isTypeCode(addenda18.TypeCode); err != nil {
		return fieldError("TypeCode", err, addenda18.TypeCode)
	}
	// Type Code must be 18
	if addenda18.TypeCode != "18" {
		return fieldError("TypeCode", ErrAddendaTypeCode, addenda18.TypeCode)
	}
	if err := addenda18.isAlphanumeric(addenda18.ForeignCorrespondentBankName); err != nil {
		return fieldError("ForeignCorrespondentBankName", err, addenda18.ForeignCorrespondentBankName)
	}
	if err := addenda18.isAlphanumeric(addenda18.ForeignCorrespondentBankIDNumberQualifier); err != nil {
		return fieldError("ForeignCorrespondentBankIDNumberQualifier", err, addenda18.ForeignCorrespondentBankIDNumberQualifier)
	}
	if err := addenda18.isAlphanumeric(addenda18.ForeignCorrespondentBankIDNumber); err != nil {
		return fieldError("ForeignCorrespondentBankIDNumber", err, addenda18.ForeignCorrespondentBankIDNumber)
	}
	if err := addenda18.isAlphanumeric(addenda18.ForeignCorrespondentBankBranchCountryCode); err != nil {
		return fieldError("ForeignCorrespondentBankBranchCountryCode", err, addenda18.ForeignCorrespondentBankBranchCountryCode)
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda18 *Addenda18) fieldInclusion() error {
	if addenda18.TypeCode == "" {
		return fieldError("TypeCode", ErrConstructor, addenda18.TypeCode)
	}
	if addenda18.ForeignCorrespondentBankName == "" {
		return fieldError("ForeignCorrespondentBankName", ErrConstructor, addenda18.ForeignCorrespondentBankName)
	}
	if addenda18.ForeignCorrespondentBankIDNumberQualifier == "" {
		return fieldError("ForeignCorrespondentBankIDNumberQualifier", ErrConstructor, addenda18.ForeignCorrespondentBankIDNumberQualifier)
	}
	if addenda18.ForeignCorrespondentBankIDNumber == "" {
		return fieldError("ForeignCorrespondentBankIDNumber", ErrConstructor, addenda18.ForeignCorrespondentBankIDNumber)
	}
	if addenda18.ForeignCorrespondentBankBranchCountryCode == "" {
		return fieldError("ForeignCorrespondentBankBranchCountryCode", ErrConstructor, addenda18.ForeignCorrespondentBankBranchCountryCode)
	}
	if addenda18.SequenceNumber == 0 {
		return fieldError("SequenceNumber", ErrConstructor, addenda18.SequenceNumberField())
	}
	if addenda18.EntryDetailSequenceNumber == 0 {
		return fieldError("EntryDetailSequenceNumber", ErrConstructor, addenda18.EntryDetailSequenceNumberField())
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

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda18 *Addenda18) EntryDetailSequenceNumberField() string {
	return addenda18.numericField(addenda18.EntryDetailSequenceNumber, 7)
}
