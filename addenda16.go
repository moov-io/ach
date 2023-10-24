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

// Addenda16 is an addenda which provides business transaction information for Addenda Type
// Code 16 in a machine readable format. It is usually formatted according to ANSI, ASC, X16 Standard.
//
// # Addenda16 is mandatory for IAT entries
//
// The Addenda16 record identifies key information related to the Receiver.
type Addenda16 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// TypeCode Addenda16 types code '16'
	TypeCode string `json:"typeCode"`
	// Receiver City & State / Province
	// Data elements City and State / Province  should be separated with an asterisk (*) as a delimiter
	// and the field should end with a backslash (\).
	// For example: San Francisco*CA\
	ReceiverCityStateProvince string `json:"receiverCityStateProvince"`
	// Receiver Country & Postal Code
	// Data elements must be separated by an asterisk (*) and must end with a backslash (\)
	// For example: US*10036\
	ReceiverCountryPostalCode string `json:"receiverCountryPostalCode"`
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

// NewAddenda16 returns a new Addenda16 with default values for none exported fields
func NewAddenda16() *Addenda16 {
	addenda16 := new(Addenda16)
	addenda16.TypeCode = "16"
	return addenda16
}

// Parse takes the input record string and parses the Addenda16 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda16 *Addenda16) Parse(record string) {
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
			// 2-3 Always 16
			addenda16.TypeCode = reset()
		case 38:
			// 4-38 ReceiverCityStateProvince
			addenda16.ReceiverCityStateProvince = strings.TrimSpace(reset())
		case 73:
			// 39-73 ReceiverCountryPostalCode
			addenda16.ReceiverCountryPostalCode = strings.TrimSpace(reset())
		case 87:
			// 74-87 reserved - Leave blank
			reset()
		case 94:
			// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
			addenda16.EntryDetailSequenceNumber = addenda16.parseNumField(reset())
		}
	}
}

// String writes the Addenda16 struct to a 94 character string.
func (addenda16 *Addenda16) String() string {
	if addenda16 == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(entryAddendaPos)
	buf.WriteString(addenda16.TypeCode)
	buf.WriteString(addenda16.ReceiverCityStateProvinceField())
	buf.WriteString(addenda16.ReceiverCountryPostalCodeField())
	buf.WriteString("              ")
	buf.WriteString(addenda16.EntryDetailSequenceNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda16 *Addenda16) Validate() error {
	if err := addenda16.fieldInclusion(); err != nil {
		return err
	}
	if err := addenda16.isTypeCode(addenda16.TypeCode); err != nil {
		return fieldError("TypeCode", err, addenda16.TypeCode)
	}
	// Type Code must be 16
	if addenda16.TypeCode != "16" {
		return fieldError("TypeCode", ErrAddendaTypeCode, addenda16.TypeCode)
	}
	if err := addenda16.isAlphanumeric(addenda16.ReceiverCityStateProvince); err != nil {
		return fieldError("ReceiverCityStateProvince", err, addenda16.ReceiverCityStateProvince)
	}
	if err := addenda16.isAlphanumeric(addenda16.ReceiverCountryPostalCode); err != nil {
		return fieldError("ReceiverCountryPostalCode", err, addenda16.ReceiverCountryPostalCode)
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda16 *Addenda16) fieldInclusion() error {
	if addenda16.TypeCode == "" {
		return fieldError("TypeCode", ErrConstructor, addenda16.TypeCode)
	}
	if addenda16.ReceiverCityStateProvince == "" {
		return fieldError("ReceiverCityStateProvince", ErrConstructor, addenda16.ReceiverCityStateProvince)
	}
	if addenda16.ReceiverCountryPostalCode == "" {
		return fieldError("ReceiverCountryPostalCode", ErrConstructor, addenda16.ReceiverCountryPostalCode)
	}
	if addenda16.EntryDetailSequenceNumber == 0 {
		return fieldError("EntryDetailSequenceNumber", ErrConstructor, addenda16.EntryDetailSequenceNumberField())
	}
	return nil
}

// ReceiverCityStateProvinceField gets the ReceiverCityStateProvinceField left padded
func (addenda16 *Addenda16) ReceiverCityStateProvinceField() string {
	return addenda16.alphaField(addenda16.ReceiverCityStateProvince, 35)
}

// ReceiverCountryPostalCodeField gets the ReceiverCountryPostalCode field left padded
func (addenda16 *Addenda16) ReceiverCountryPostalCodeField() string {
	return addenda16.alphaField(addenda16.ReceiverCountryPostalCode, 35)
}

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda16 *Addenda16) EntryDetailSequenceNumberField() string {
	return addenda16.numericField(addenda16.EntryDetailSequenceNumber, 7)
}
