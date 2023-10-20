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

// Addenda15 is an addenda which provides business transaction information for Addenda Type
// Code 15 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
//
// # Addenda15 is mandatory for IAT entries
//
// The Addenda15 record identifies key information related to the Receiver.
type Addenda15 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// TypeCode Addenda15 types code '15'
	TypeCode string `json:"typeCode"`
	// Receiver Identification Number contains the accounting number by which the Originator is known to
	// the Receiver for descriptive purposes. NACHA Rules recommend but do not require the RDFI to print
	// the contents of this field on the receiver's statement.
	ReceiverIDNumber string `json:"receiverIDNumber,omitempty"`
	// Receiver Street Address contains the Receiver's physical address
	ReceiverStreetAddress string `json:"receiverStreetAddress"`
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

// NewAddenda15 returns a new Addenda15 with default values for none exported fields
func NewAddenda15() *Addenda15 {
	addenda15 := new(Addenda15)
	addenda15.TypeCode = "15"
	return addenda15
}

// Parse takes the input record string and parses the Addenda15 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda15 *Addenda15) Parse(record string) {
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
			// 2-3 Always 15
			addenda15.TypeCode = reset()
		case 18:
			// 4-18
			addenda15.ReceiverIDNumber = addenda15.parseStringField(reset())
		case 53:
			// 19-53
			addenda15.ReceiverStreetAddress = strings.TrimSpace(reset())
		case 87:
			// 54-87 reserved - Leave blank
			reset()
		case 94:
			// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
			addenda15.EntryDetailSequenceNumber = addenda15.parseNumField(reset())
		}
	}
}

// String writes the Addenda15 struct to a 94 character string.
func (addenda15 *Addenda15) String() string {
	if addenda15 == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(entryAddendaPos)
	buf.WriteString(addenda15.TypeCode)
	buf.WriteString(addenda15.ReceiverIDNumberField())
	buf.WriteString(addenda15.ReceiverStreetAddressField())
	buf.WriteString("                                  ")
	buf.WriteString(addenda15.EntryDetailSequenceNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda15 *Addenda15) Validate() error {
	if err := addenda15.fieldInclusion(); err != nil {
		return err
	}
	if err := addenda15.isTypeCode(addenda15.TypeCode); err != nil {
		return fieldError("TypeCode", err, addenda15.TypeCode)
	}
	// Type Code must be 15
	if addenda15.TypeCode != "15" {
		return fieldError("TypeCode", ErrAddendaTypeCode, addenda15.TypeCode)
	}
	if err := addenda15.isAlphanumeric(addenda15.ReceiverIDNumber); err != nil {
		return fieldError("ReceiverIDNumber", err, addenda15.ReceiverIDNumber)
	}
	if err := addenda15.isAlphanumeric(addenda15.ReceiverStreetAddress); err != nil {
		return fieldError("ReceiverStreetAddress", err, addenda15.ReceiverStreetAddress)
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda15 *Addenda15) fieldInclusion() error {
	if addenda15.TypeCode == "" {
		return fieldError("TypeCode", ErrConstructor, addenda15.TypeCode)
	}
	if addenda15.ReceiverStreetAddress == "" {
		return fieldError("ReceiverStreetAddress", ErrConstructor, addenda15.ReceiverStreetAddress)
	}
	if addenda15.EntryDetailSequenceNumber == 0 {
		return fieldError("EntryDetailSequenceNumber", ErrConstructor, addenda15.EntryDetailSequenceNumberField())
	}
	return nil
}

// ReceiverIDNumberField gets the ReceiverIDNumber field left padded
func (addenda15 *Addenda15) ReceiverIDNumberField() string {
	return addenda15.alphaField(addenda15.ReceiverIDNumber, 15)
}

// ReceiverStreetAddressField gets the ReceiverStreetAddressField field left padded
func (addenda15 *Addenda15) ReceiverStreetAddressField() string {
	return addenda15.alphaField(addenda15.ReceiverStreetAddress, 35)
}

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda15 *Addenda15) EntryDetailSequenceNumberField() string {
	return addenda15.numericField(addenda15.EntryDetailSequenceNumber, 7)
}
