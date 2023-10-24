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

// Addenda02 is a Addendumer addenda which provides business transaction information for Addenda Type
// Code 02 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
// It is used for following StandardEntryClassCode: MTE, POS, and SHR.
type Addenda02 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// TypeCode Addenda02 type code '02'
	TypeCode string `json:"typeCode"`
	// ReferenceInformationOne may be used for additional reference numbers, identification numbers,
	// or codes that the merchant needs to identify the particular transaction or customer.
	ReferenceInformationOne string `json:"referenceInformationOne,omitempty"`
	// ReferenceInformationTwo  may be used for additional reference numbers, identification numbers,
	// or codes that the merchant needs to identify the particular transaction or customer.
	ReferenceInformationTwo string `json:"referenceInformationTwo,omitempty"`
	// TerminalIdentificationCode identifies an Electronic terminal with a unique code that allows
	// a terminal owner and/or switching network to identify the terminal at which an Entry originated.
	TerminalIdentificationCode string `json:"terminalIdentificationCode"`
	// TransactionSerialNumber is assigned by the terminal at the time the transaction is originated.  The
	// number, with the Terminal Identification Code, serves as an audit trail for the transaction and is
	// usually assigned in ascending sequence.
	TransactionSerialNumber string `json:"transactionSerialNumber"`
	// TransactionDate expressed MMDD identifies the date on which the transaction occurred.
	TransactionDate string `json:"transactionDate"`
	// AuthorizationCodeOrExpireDate indicates the code that a card authorization center has
	// furnished to the merchant.
	AuthorizationCodeOrExpireDate string `json:"authorizationCodeOrExpireDate,omitempty"`
	// Terminal Location identifies the specific location of a terminal (i.e., street names of an
	// intersection, address, etc.) in accordance with the requirements of Regulation E.
	TerminalLocation string `json:"terminalLocation"`
	// TerminalCity Identifies the city in which the electronic terminal is located.
	TerminalCity string `json:"terminalCity"`
	// TerminalState Identifies the state in which the electronic terminal is located
	TerminalState string `json:"terminalState"`
	// TraceNumber Standard Entry Detail Trace Number
	//
	// Use TraceNumberField for a properly formatted string representation.
	TraceNumber string `json:"traceNumber,omitempty"`
	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters
}

// NewAddenda02 returns a new Addenda02 with default values for none exported fields
func NewAddenda02() *Addenda02 {
	addenda02 := new(Addenda02)
	addenda02.TypeCode = "02"
	return addenda02
}

// Parse takes the input record string and parses the Addenda02 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda02 *Addenda02) Parse(record string) {
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
			// 2-3 Always 02
			addenda02.TypeCode = reset()
		case 10:
			// 4-10 Based on the information entered (04-10) 7 alphanumeric
			addenda02.ReferenceInformationOne = strings.TrimSpace(reset())
		case 13:
			// 11-13 Based on the information entered (11-13) 3 alphanumeric
			addenda02.ReferenceInformationTwo = strings.TrimSpace(reset())
		case 19:
			// 14-19
			addenda02.TerminalIdentificationCode = strings.TrimSpace(reset())
		case 25:
			// 20-25
			addenda02.TransactionSerialNumber = strings.TrimSpace(reset())
		case 29:
			// 26-29
			addenda02.TransactionDate = strings.TrimSpace(reset())
		case 35:
			// 30-35
			addenda02.AuthorizationCodeOrExpireDate = strings.TrimSpace(reset())
		case 62:
			// 36-62
			addenda02.TerminalLocation = strings.TrimSpace(reset())
		case 77:
			// 63-77
			addenda02.TerminalCity = strings.TrimSpace(reset())
		case 79:
			// 78-79
			addenda02.TerminalState = strings.TrimSpace(reset())
		case 94:
			// 80-94
			addenda02.TraceNumber = strings.TrimSpace(reset())
		}
	}
}

// String writes the Addenda02 struct to a 94 character string.
func (addenda02 *Addenda02) String() string {
	if addenda02 == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(entryAddendaPos)
	buf.WriteString(addenda02.TypeCode)
	buf.WriteString(addenda02.ReferenceInformationOneField())
	buf.WriteString(addenda02.ReferenceInformationTwoField())
	buf.WriteString(addenda02.TerminalIdentificationCodeField())
	buf.WriteString(addenda02.TransactionSerialNumberField())
	buf.WriteString(addenda02.TransactionDateField())
	buf.WriteString(addenda02.AuthorizationCodeOrExpireDateField())
	buf.WriteString(addenda02.TerminalLocationField())
	buf.WriteString(addenda02.TerminalCityField())
	buf.WriteString(addenda02.TerminalStateField())
	buf.WriteString(addenda02.TraceNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda02 *Addenda02) Validate() error {
	if err := addenda02.fieldInclusion(); err != nil {
		return err
	}
	if err := addenda02.isTypeCode(addenda02.TypeCode); err != nil {
		return fieldError("TypeCode", err, addenda02.TypeCode)
	}
	// Type Code must be 02
	if addenda02.TypeCode != "02" {
		return fieldError("TypeCode", ErrAddendaTypeCode, addenda02.TypeCode)
	}
	if err := addenda02.isAlphanumeric(addenda02.ReferenceInformationOne); err != nil {
		return fieldError("ReferenceInformationOne", err, addenda02.ReferenceInformationOne)
	}
	if err := addenda02.isAlphanumeric(addenda02.ReferenceInformationTwo); err != nil {
		return fieldError("ReferenceInformationTwo", err, addenda02.ReferenceInformationTwo)
	}
	if err := addenda02.isAlphanumeric(addenda02.TerminalIdentificationCode); err != nil {
		return fieldError("TerminalIdentificationCode", err, addenda02.TerminalIdentificationCode)
	}
	if err := addenda02.isAlphanumeric(addenda02.TransactionSerialNumber); err != nil {
		return fieldError("TransactionSerialNumber", err, addenda02.TransactionSerialNumber)
	}

	// TransactionDate Addenda02 ACH File format is MMDD. Validate MM is 01-12 and day for the
	// month 01-31 depending on month.
	mm := addenda02.parseStringField(addenda02.TransactionDateField()[0:2])
	dd := addenda02.parseStringField(addenda02.TransactionDateField()[2:4])
	if err := addenda02.isMonth(mm); err != nil {
		return fieldError("TransactionDate", ErrValidMonth, mm)
	}
	if err := addenda02.isDay(mm, dd); err != nil {
		return fieldError("TransactionDate", ErrValidDay, mm)
	}

	if err := addenda02.isAlphanumeric(addenda02.AuthorizationCodeOrExpireDate); err != nil {
		return fieldError("AuthorizationCodeOrExpireDate", err, addenda02.AuthorizationCodeOrExpireDate)
	}
	if err := addenda02.isAlphanumeric(addenda02.TerminalLocation); err != nil {
		return fieldError("TerminalLocation", err, addenda02.TerminalLocation)
	}
	if err := addenda02.isAlphanumeric(addenda02.TerminalCity); err != nil {
		return fieldError("TerminalCity", err, addenda02.TerminalCity)
	}
	if err := addenda02.isAlphanumeric(addenda02.TerminalState); err != nil {
		return fieldError("TerminalState", err, addenda02.TerminalState)
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values  and required fields are defined. If fields are
// invalid the ACH transfer will be returned.

func (addenda02 *Addenda02) fieldInclusion() error {
	if addenda02.TypeCode == "" {
		return fieldError("TypeCode", ErrConstructor, addenda02.TypeCode)
	}
	// Required Fields
	if addenda02.TransactionSerialNumber == "" {
		return fieldError("TransactionSerialNumber", ErrFieldRequired, addenda02.TransactionSerialNumber)
	}
	if addenda02.TransactionDate == "" {
		return fieldError("TransactionDate", ErrFieldRequired, addenda02.TransactionDate)
	}
	if addenda02.TerminalLocation == "" {
		return fieldError("TerminalLocation", ErrFieldRequired, addenda02.TerminalLocation)
	}
	if addenda02.TerminalCity == "" {
		return fieldError("TerminalCity", ErrFieldRequired, addenda02.TerminalCity)
	}
	if addenda02.TerminalState == "" {
		return fieldError("TerminalState", ErrFieldRequired, addenda02.TerminalState)
	}
	return nil
}

// ReferenceInformationOneField returns a space padded ReferenceInformationOne string
func (addenda02 *Addenda02) ReferenceInformationOneField() string {
	return addenda02.alphaField(addenda02.ReferenceInformationOne, 7)
}

// ReferenceInformationTwoField returns a space padded ReferenceInformationTwo string
func (addenda02 *Addenda02) ReferenceInformationTwoField() string {
	return addenda02.alphaField(addenda02.ReferenceInformationOne, 3)
}

// TerminalIdentificationCodeField returns a space padded TerminalIdentificationCode string
func (addenda02 *Addenda02) TerminalIdentificationCodeField() string {
	return addenda02.alphaField(addenda02.TerminalIdentificationCode, 6)
}

// TransactionSerialNumberField returns a zero padded TransactionSerialNumber string
func (addenda02 *Addenda02) TransactionSerialNumberField() string {
	return addenda02.alphaField(addenda02.TransactionSerialNumber, 6)
}

// TransactionDateField returns TransactionDate MMDD string
func (addenda02 *Addenda02) TransactionDateField() string {
	return addenda02.alphaField(addenda02.TransactionDate, 4)
}

// AuthorizationCodeOrExpireDateField returns a space padded AuthorizationCodeOrExpireDate string
func (addenda02 *Addenda02) AuthorizationCodeOrExpireDateField() string {
	return addenda02.alphaField(addenda02.AuthorizationCodeOrExpireDate, 6)
}

// TerminalLocationField returns a space padded TerminalLocation string
func (addenda02 *Addenda02) TerminalLocationField() string {
	return addenda02.alphaField(addenda02.TerminalLocation, 27)
}

// TerminalCityField returns a space padded TerminalCity string
func (addenda02 *Addenda02) TerminalCityField() string {
	return addenda02.alphaField(addenda02.TerminalCity, 15)
}

// TerminalStateField returns a space padded TerminalState string
func (addenda02 *Addenda02) TerminalStateField() string {
	return addenda02.alphaField(addenda02.TerminalState, 2)
}

// TraceNumberField returns a space padded TraceNumber string
func (addenda02 *Addenda02) TraceNumberField() string {
	return addenda02.stringField(addenda02.TraceNumber, 15)
}
