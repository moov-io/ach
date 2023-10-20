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

// Addenda13 is an addenda which provides business transaction information for Addenda Type
// Code 13 in a machine readable format. It is usually formatted according to ANSI, ASC, X13 Standard.
//
// # Addenda13 is mandatory for IAT entries
//
// The Addenda13 contains information related to the financial institution originating the entry.
// For inbound IAT entries, the Fourth Addenda Record must contain information to identify the
// foreign financial institution that is providing the funding and payment instruction for the IAT entry.
type Addenda13 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
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

// NewAddenda13 returns a new Addenda13 with default values for none exported fields
func NewAddenda13() *Addenda13 {
	addenda13 := new(Addenda13)
	addenda13.TypeCode = "13"
	return addenda13
}

// Parse takes the input record string and parses the Addenda13 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda13 *Addenda13) Parse(record string) {
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
			// 2-3 Always 13
			addenda13.TypeCode = reset()
		case 38:
			// 4-38 ODFIName
			addenda13.ODFIName = strings.TrimSpace(reset())
		case 40:
			// 39-40 ODFIIDNumberQualifier
			addenda13.ODFIIDNumberQualifier = reset()
		case 74:
			// 41-74 ODFIIdentification
			addenda13.ODFIIdentification = addenda13.parseStringField(reset())
		case 77:
			// 75-77
			addenda13.ODFIBranchCountryCode = strings.TrimSpace(reset())
		case 87:
			// 78-87 reserved - Leave blank
			reset()
		case 94:
			// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
			addenda13.EntryDetailSequenceNumber = addenda13.parseNumField(reset())
		}
	}
}

// String writes the Addenda13 struct to a 94 character string.
func (addenda13 *Addenda13) String() string {
	if addenda13 == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(entryAddendaPos)
	buf.WriteString(addenda13.TypeCode)
	buf.WriteString(addenda13.ODFINameField())
	buf.WriteString(addenda13.ODFIIDNumberQualifierField())
	buf.WriteString(addenda13.ODFIIdentificationField())
	buf.WriteString(addenda13.ODFIBranchCountryCodeField())
	buf.WriteString("          ")
	buf.WriteString(addenda13.EntryDetailSequenceNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda13 *Addenda13) Validate() error {
	if err := addenda13.fieldInclusion(); err != nil {
		return err
	}
	if err := addenda13.isTypeCode(addenda13.TypeCode); err != nil {
		return fieldError("TypeCode", err, addenda13.TypeCode)
	}
	// Type Code must be 13
	if addenda13.TypeCode != "13" {
		return fieldError("TypeCode", ErrAddendaTypeCode, addenda13.TypeCode)
	}
	if err := addenda13.isAlphanumeric(addenda13.ODFIName); err != nil {
		return fieldError("ODFIName", err, addenda13.ODFIName)
	}
	// Valid ODFI Identification Number Qualifier
	if err := addenda13.isIDNumberQualifier(addenda13.ODFIIDNumberQualifier); err != nil {
		return fieldError("ODFIIDNumberQualifier", err, addenda13.ODFIIDNumberQualifier)
	}
	if err := addenda13.isAlphanumeric(addenda13.ODFIIdentification); err != nil {
		return fieldError("ODFIIdentification", err, addenda13.ODFIIdentification)
	}
	if err := addenda13.isAlphanumeric(addenda13.ODFIBranchCountryCode); err != nil {
		return fieldError("ODFIBranchCountryCode", err, addenda13.ODFIBranchCountryCode)
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda13 *Addenda13) fieldInclusion() error {
	if addenda13.TypeCode == "" {
		return fieldError("TypeCode", ErrConstructor, addenda13.TypeCode)
	}
	if addenda13.ODFIName == "" {
		return fieldError("ODFIName", ErrConstructor, addenda13.ODFIName)
	}
	if addenda13.ODFIIDNumberQualifier == "" {
		return fieldError("ODFIIDNumberQualifier", ErrConstructor, addenda13.ODFIIDNumberQualifier)
	}
	if addenda13.ODFIIdentification == "" {
		return fieldError("ODFIIdentification", ErrConstructor, addenda13.ODFIIdentification)
	}
	if addenda13.ODFIBranchCountryCode == "" {
		return fieldError("ODFIBranchCountryCode", ErrConstructor, addenda13.ODFIBranchCountryCode)
	}
	if addenda13.EntryDetailSequenceNumber == 0 {
		return fieldError("EntryDetailSequenceNumber", ErrConstructor, addenda13.EntryDetailSequenceNumberField())
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

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda13 *Addenda13) EntryDetailSequenceNumberField() string {
	return addenda13.numericField(addenda13.EntryDetailSequenceNumber, 7)
}
