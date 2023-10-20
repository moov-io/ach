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

// Addenda14 is an addenda which provides business transaction information for Addenda Type
// Code 14 in a machine readable format. It is usually formatted according to ANSI, ASC, X14 Standard.
//
// # Addenda14 is mandatory for IAT entries
//
// The Addenda14 identifies the Receiving financial institution holding the Receiver's account.
type Addenda14 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// TypeCode Addenda14 types code '14'
	TypeCode string `json:"typeCode"`
	// Name of the Receiver's bank
	RDFIName string `json:"RDFIName"`
	// Receiving DFI Identification Number Qualifier
	// The 2-digit code that identifies the numbering scheme used in the
	// Receiving DFI Identification Number field:
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

// NewAddenda14 returns a new Addenda14 with default values for none exported fields
func NewAddenda14() *Addenda14 {
	addenda14 := new(Addenda14)
	addenda14.TypeCode = "14"
	return addenda14
}

// Parse takes the input record string and parses the Addenda14 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda14 *Addenda14) Parse(record string) {
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
			// 2-3 Always 14
			addenda14.TypeCode = reset()
		case 38:
			// 4-38 RDFIName
			addenda14.RDFIName = strings.TrimSpace(reset())
		case 40:
			// 39-40 RDFIIDNumberQualifier
			addenda14.RDFIIDNumberQualifier = reset()
		case 74:
			// 41-74 RDFIIdentification
			addenda14.RDFIIdentification = addenda14.parseStringField(reset())
		case 77:
			// 75-77
			addenda14.RDFIBranchCountryCode = strings.TrimSpace(reset())
		case 87:
			// 78-87 reserved - Leave blank
			reset()
		case 94:
			// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
			addenda14.EntryDetailSequenceNumber = addenda14.parseNumField(reset())
		}
	}
}

// String writes the Addenda14 struct to a 94 character string.
func (addenda14 *Addenda14) String() string {
	if addenda14 == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(entryAddendaPos)
	buf.WriteString(addenda14.TypeCode)
	buf.WriteString(addenda14.RDFINameField())
	buf.WriteString(addenda14.RDFIIDNumberQualifierField())
	buf.WriteString(addenda14.RDFIIdentificationField())
	buf.WriteString(addenda14.RDFIBranchCountryCodeField())
	buf.WriteString("          ")
	buf.WriteString(addenda14.EntryDetailSequenceNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda14 *Addenda14) Validate() error {
	if err := addenda14.fieldInclusion(); err != nil {
		return err
	}
	if err := addenda14.isTypeCode(addenda14.TypeCode); err != nil {
		return fieldError("TypeCode", err, addenda14.TypeCode)
	}
	// Type Code must be 14
	if addenda14.TypeCode != "14" {
		return fieldError("TypeCode", ErrAddendaTypeCode, addenda14.TypeCode)
	}
	if err := addenda14.isAlphanumeric(addenda14.RDFIName); err != nil {
		return fieldError("RDFIName", err, addenda14.RDFIName)
	}
	// Valid RDFI Identification Number Qualifier
	if err := addenda14.isIDNumberQualifier(addenda14.RDFIIDNumberQualifier); err != nil {
		return fieldError("RDFIIDNumberQualifier", ErrIDNumberQualifier, addenda14.RDFIIDNumberQualifier)
	}
	if err := addenda14.isAlphanumeric(addenda14.RDFIIdentification); err != nil {
		return fieldError("RDFIIdentification", err, addenda14.RDFIIdentification)
	}
	if err := addenda14.isAlphanumeric(addenda14.RDFIBranchCountryCode); err != nil {
		return fieldError("RDFIBranchCountryCode", err, addenda14.RDFIBranchCountryCode)
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda14 *Addenda14) fieldInclusion() error {
	if addenda14.TypeCode == "" {
		return fieldError("TypeCode", ErrConstructor, addenda14.TypeCode)
	}
	if addenda14.RDFIName == "" {
		return fieldError("RDFIName", ErrConstructor, addenda14.RDFIName)
	}
	if addenda14.RDFIIDNumberQualifier == "" {
		return fieldError("RDFIIDNumberQualifier", ErrConstructor, addenda14.RDFIIDNumberQualifier)
	}
	if addenda14.RDFIIdentification == "" {
		return fieldError("RDFIIdentification", ErrConstructor, addenda14.RDFIIdentification)
	}
	if addenda14.RDFIBranchCountryCode == "" {
		return fieldError("RDFIBranchCountryCode", ErrConstructor, addenda14.RDFIBranchCountryCode)
	}
	if addenda14.EntryDetailSequenceNumber == 0 {
		return fieldError("EntryDetailSequenceNumber", ErrConstructor, addenda14.EntryDetailSequenceNumberField())
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

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda14 *Addenda14) EntryDetailSequenceNumberField() string {
	return addenda14.numericField(addenda14.EntryDetailSequenceNumber, 7)
}
