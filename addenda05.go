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

// Addenda05 is a Addendumer addenda which provides business transaction information for Addenda Type
// Code 05 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
// It is used for the following StandardEntryClassCode: ACK, ATX, CCD, CIE, CTX, DNE, ENR, WEB, PPD, TRX.
type Addenda05 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// TypeCode Addenda05 types code '05'
	TypeCode string `json:"typeCode"`
	// PaymentRelatedInformation
	PaymentRelatedInformation string `json:"paymentRelatedInformation"`
	// SequenceNumber is consecutively assigned to each Addenda05 Record following
	// an Entry Detail Record. The first addenda05 sequence number must always
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

// NewAddenda05 returns a new Addenda05 with default values for none exported fields
func NewAddenda05() *Addenda05 {
	addenda05 := new(Addenda05)
	addenda05.TypeCode = "05"
	return addenda05
}

// Parse takes the input record string and parses the Addenda05 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda05 *Addenda05) Parse(record string) {
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
			// 2-3 Always 05
			addenda05.TypeCode = reset()
		case 83:
			// 4-83 Based on the information entered (04-83) 80 alphanumeric
			addenda05.PaymentRelatedInformation = strings.TrimSpace(reset())
		case 87:
			// 84-87 SequenceNumber is consecutively assigned to each Addenda05 Record following
			// an Entry Detail Record
			addenda05.SequenceNumber = addenda05.parseNumField(reset())
		case 94:
			// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
			addenda05.EntryDetailSequenceNumber = addenda05.parseNumField(reset())
		}
	}
}

// String writes the Addenda05 struct to a 94 character string.
func (addenda05 *Addenda05) String() string {
	if addenda05 == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(entryAddendaPos)
	buf.WriteString(addenda05.TypeCode)
	buf.WriteString(addenda05.PaymentRelatedInformationField())
	buf.WriteString(addenda05.SequenceNumberField())
	buf.WriteString(addenda05.EntryDetailSequenceNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda05 *Addenda05) Validate() error {
	if err := addenda05.fieldInclusion(); err != nil {
		return err
	}

	if err := addenda05.isTypeCode(addenda05.TypeCode); err != nil {
		return fieldError("TypeCode", err, addenda05.TypeCode)
	}
	// Type Code must be 05
	if addenda05.TypeCode != "05" {
		return fieldError("TypeCode", ErrAddendaTypeCode, addenda05.TypeCode)
	}
	if err := addenda05.isAlphanumeric(addenda05.PaymentRelatedInformation); err != nil {
		return fieldError("PaymentRelatedInformation", err, addenda05.PaymentRelatedInformation)
	}

	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda05 *Addenda05) fieldInclusion() error {
	if addenda05.TypeCode == "" {
		return fieldError("TypeCode", ErrConstructor, addenda05.TypeCode)
	}
	if addenda05.SequenceNumber == 0 {
		return fieldError("SequenceNumber", ErrConstructor, addenda05.SequenceNumberField())
	}
	if addenda05.EntryDetailSequenceNumber == 0 {
		return fieldError("EntryDetailSequenceNumber", ErrConstructor, addenda05.EntryDetailSequenceNumberField())
	}
	return nil
}

// PaymentRelatedInformationField returns a zero padded PaymentRelatedInformation string
func (addenda05 *Addenda05) PaymentRelatedInformationField() string {
	return addenda05.alphaField(addenda05.PaymentRelatedInformation, 80)
}

// SequenceNumberField returns a zero padded SequenceNumber string
func (addenda05 *Addenda05) SequenceNumberField() string {
	return addenda05.numericField(addenda05.SequenceNumber, 4)
}

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda05 *Addenda05) EntryDetailSequenceNumberField() string {
	return addenda05.numericField(addenda05.EntryDetailSequenceNumber, 7)
}
