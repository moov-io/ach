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

// Addenda17 is an addenda which provides business transaction information for Addenda Type
// Code 17 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
//
// # Addenda17 is optional for IAT entries
//
// The Addenda17 record identifies payment-related data. A maximum of two of these Addenda Records
// may be included with each IAT entry.
type Addenda17 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// TypeCode Addenda17 types code '17'
	TypeCode string `json:"typeCode"`
	// PaymentRelatedInformation
	PaymentRelatedInformation string `json:"paymentRelatedInformation"`
	// SequenceNumber is consecutively assigned to each Addenda17 Record following
	// an Entry Detail Record. The first addenda17 sequence number must always
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

// NewAddenda17 returns a new Addenda17 with default values for none exported fields
func NewAddenda17() *Addenda17 {
	addenda17 := new(Addenda17)
	addenda17.TypeCode = "17"
	return addenda17
}

// Parse takes the input record string and parses the Addenda17 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda17 *Addenda17) Parse(record string) {
	if utf8.RuneCountInString(record) != 94 {
		return
	}

	// 1-1 Always 7
	// 2-3 Always 17
	addenda17.TypeCode = record[1:3]
	// 4-83 Based on the information entered (04-83) 80 alphanumeric
	addenda17.PaymentRelatedInformation = strings.TrimSpace(record[3:83])
	// 84-87 SequenceNumber is consecutively assigned to each Addenda17 Record following
	// an Entry Detail Record
	addenda17.SequenceNumber = addenda17.parseNumField(record[83:87])
	// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
	addenda17.EntryDetailSequenceNumber = addenda17.parseNumField(record[87:94])
}

// String writes the Addenda17 struct to a 94 character string.
func (addenda17 *Addenda17) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(entryAddendaPos)
	buf.WriteString(addenda17.TypeCode)
	buf.WriteString(addenda17.PaymentRelatedInformationField())
	buf.WriteString(addenda17.SequenceNumberField())
	buf.WriteString(addenda17.EntryDetailSequenceNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda17 *Addenda17) Validate() error {
	if err := addenda17.fieldInclusion(); err != nil {
		return err
	}
	if err := addenda17.isTypeCode(addenda17.TypeCode); err != nil {
		return fieldError("TypeCode", err, addenda17.TypeCode)
	}
	// Type Code must be 17
	if addenda17.TypeCode != "17" {
		return fieldError("TypeCode", ErrAddendaTypeCode, addenda17.TypeCode)
	}
	if err := addenda17.isAlphanumeric(addenda17.PaymentRelatedInformation); err != nil {
		return fieldError("PaymentRelatedInformation", err, addenda17.PaymentRelatedInformation)
	}

	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda17 *Addenda17) fieldInclusion() error {
	if addenda17.TypeCode == "" {
		return fieldError("TypeCode", ErrConstructor, addenda17.TypeCode)
	}
	if addenda17.SequenceNumber == 0 {
		return fieldError("SequenceNumber", ErrConstructor, addenda17.SequenceNumberField())
	}
	if addenda17.EntryDetailSequenceNumber == 0 {
		return fieldError("EntryDetailSequenceNumber", ErrConstructor, addenda17.EntryDetailSequenceNumberField())
	}
	return nil
}

// PaymentRelatedInformationField returns a zero padded PaymentRelatedInformation string
func (addenda17 *Addenda17) PaymentRelatedInformationField() string {
	return addenda17.alphaField(addenda17.PaymentRelatedInformation, 80)
}

// SequenceNumberField returns a zero padded SequenceNumber string
func (addenda17 *Addenda17) SequenceNumberField() string {
	return addenda17.numericField(addenda17.SequenceNumber, 4)
}

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda17 *Addenda17) EntryDetailSequenceNumberField() string {
	return addenda17.numericField(addenda17.EntryDetailSequenceNumber, 7)
}
