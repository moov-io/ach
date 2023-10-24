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

// Addenda11 is an addenda which provides business transaction information for Addenda Type
// Code 11 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
//
// # Addenda11 is mandatory for IAT entries
//
// The Addenda11 record identifies key information related to the Originator of
// the entry.
type Addenda11 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// TypeCode Addenda11 types code '11'
	TypeCode string `json:"typeCode"`
	// Originator Name contains the originators name (your company name / name)
	OriginatorName string `json:"originatorName"`
	// Originator Street Address Contains the originators street address (your company's address / your address)
	OriginatorStreetAddress string `json:"originatorStreetAddress"`
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

// NewAddenda11 returns a new Addenda11 with default values for none exported fields
func NewAddenda11() *Addenda11 {
	addenda11 := new(Addenda11)
	addenda11.TypeCode = "11"
	return addenda11
}

// Parse takes the input record string and parses the Addenda11 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda11 *Addenda11) Parse(record string) {
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
			// 2-3 Always 11
			addenda11.TypeCode = reset()
		case 38:
			// 4-38
			addenda11.OriginatorName = strings.TrimSpace(reset())
		case 73:
			// 39-73
			addenda11.OriginatorStreetAddress = strings.TrimSpace(reset())
		case 87:
			// 74-87 reserved - Leave blank
			reset()
		case 94:
			// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
			addenda11.EntryDetailSequenceNumber = addenda11.parseNumField(reset())
		}
	}
}

// String writes the Addenda11 struct to a 94 character string.
func (addenda11 *Addenda11) String() string {
	if addenda11 == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(entryAddendaPos)
	buf.WriteString(addenda11.TypeCode)
	buf.WriteString(addenda11.OriginatorNameField())
	buf.WriteString(addenda11.OriginatorStreetAddressField())
	buf.WriteString("              ")
	buf.WriteString(addenda11.EntryDetailSequenceNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda11 *Addenda11) Validate() error {
	if err := addenda11.fieldInclusion(); err != nil {
		return err
	}
	if err := addenda11.isTypeCode(addenda11.TypeCode); err != nil {
		return fieldError("TypeCode", err, addenda11.TypeCode)
	}
	// Type Code must be 11
	if addenda11.TypeCode != "11" {
		return fieldError("TypeCode", ErrAddendaTypeCode, addenda11.TypeCode)
	}
	if err := addenda11.isAlphanumeric(addenda11.OriginatorName); err != nil {
		return fieldError("OriginatorName", err, addenda11.OriginatorName)
	}
	if err := addenda11.isAlphanumeric(addenda11.OriginatorStreetAddress); err != nil {
		return fieldError("OriginatorStreetAddress", err, addenda11.OriginatorStreetAddress)
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda11 *Addenda11) fieldInclusion() error {
	if addenda11.TypeCode == "" {
		return fieldError("TypeCode", ErrConstructor, addenda11.TypeCode)
	}
	if addenda11.OriginatorName == "" {
		return fieldError("OriginatorName", ErrConstructor, addenda11.OriginatorName)
	}
	if addenda11.OriginatorStreetAddress == "" {
		return fieldError("OriginatorStreetAddress", ErrConstructor, addenda11.OriginatorStreetAddress)
	}
	if addenda11.EntryDetailSequenceNumber == 0 {
		return fieldError("EntryDetailSequenceNumber", ErrConstructor, addenda11.EntryDetailSequenceNumberField())
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

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda11 *Addenda11) EntryDetailSequenceNumberField() string {
	return addenda11.numericField(addenda11.EntryDetailSequenceNumber, 7)
}
