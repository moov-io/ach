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

// Addenda10 is an addenda which provides business transaction information for Addenda Type
// Code 10 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
//
// # Addenda10 is mandatory for IAT entries
//
// The Addenda10 Record identifies the Receiver of the transaction and the dollar amount of
// the payment.
type Addenda10 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// TypeCode Addenda10 types code '10'
	TypeCode string `json:"typeCode"`
	// Transaction Type Code Describes the type of payment:
	// ANN = Annuity, BUS = Business/Commercial, DEP = Deposit, LOA = Loan, MIS = Miscellaneous, MOR = Mortgage
	// PEN = Pension, RLS = Rent/Lease, REM = Remittance2, SAL = Salary/Payroll, TAX = Tax, TEL = Telephone-Initiated Transaction
	// WEB = Internet-Initiated Transaction, ARC = Accounts Receivable Entry, BOC = Back Office Conversion Entry,
	// POP = Point of Purchase Entry, RCK = Re-presented Check Entry
	TransactionTypeCode string `json:"transactionTypeCode"`
	// Foreign Payment Amount $$$$$$$$$$$$$$$$¢¢
	// For inbound IAT payments this field should contain the USD amount or may be blank.
	ForeignPaymentAmount int `json:"foreignPaymentAmount"`
	// Foreign Trace Number
	ForeignTraceNumber string `json:"foreignTraceNumber,omitempty"`
	// Receiving Company Name/Individual Name
	Name string `json:"name"`
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

// NewAddenda10 returns a new Addenda10 with default values for none exported fields
func NewAddenda10() *Addenda10 {
	addenda10 := new(Addenda10)
	addenda10.TypeCode = "10"
	return addenda10
}

// Parse takes the input record string and parses the Addenda10 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda10 *Addenda10) Parse(record string) {
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
			// 2-3 Always 10
			addenda10.TypeCode = reset()
		case 6:
			// 04-06 Describes the type of payment
			addenda10.TransactionTypeCode = reset()
		case 24:
			// 07-24 Payment Amount	For inbound IAT payments this field should contain the USD amount or may be blank.
			addenda10.ForeignPaymentAmount = addenda10.parseNumField(reset())
		case 46:
			//  25-46 Insert blanks or zeros
			addenda10.ForeignTraceNumber = strings.TrimSpace(reset())
		case 81:
			// 47-81 Receiving Company Name/Individual Name
			addenda10.Name = strings.TrimSpace(reset())
		case 87:
			// 82-87 reserved - Leave blank
			reset()
		case 94:
			// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
			addenda10.EntryDetailSequenceNumber = addenda10.parseNumField(reset())
		}
	}
}

// String writes the Addenda10 struct to a 94 character string.
func (addenda10 *Addenda10) String() string {
	if addenda10 == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(entryAddendaPos)
	buf.WriteString(addenda10.TypeCode)
	// TransactionTypeCode Validator
	buf.WriteString(addenda10.TransactionTypeCode)
	buf.WriteString(addenda10.ForeignPaymentAmountField())
	buf.WriteString(addenda10.ForeignTraceNumberField())
	buf.WriteString(addenda10.NameField())
	buf.WriteString("      ")
	buf.WriteString(addenda10.EntryDetailSequenceNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda10 *Addenda10) Validate() error {
	if err := addenda10.fieldInclusion(); err != nil {
		return err
	}
	if err := addenda10.isTypeCode(addenda10.TypeCode); err != nil {
		return fieldError("TypeCode", err, addenda10.TypeCode)
	}
	// Type Code must be 10
	if addenda10.TypeCode != "10" {
		return fieldError("TypeCode", ErrAddendaTypeCode, addenda10.TypeCode)
	}
	if err := addenda10.isTransactionTypeCode(addenda10.TransactionTypeCode); err != nil {
		return fieldError("TransactionTypeCode", err, addenda10.TransactionTypeCode)
	}
	// ToDo: Foreign Payment Amount blank ?
	if err := addenda10.isAlphanumeric(addenda10.ForeignTraceNumber); err != nil {
		return fieldError("ForeignTraceNumber", err, addenda10.ForeignTraceNumber)
	}
	if err := addenda10.isAlphanumeric(addenda10.Name); err != nil {
		return fieldError("Name", err, addenda10.Name)
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda10 *Addenda10) fieldInclusion() error {
	if addenda10.TypeCode == "" {
		return fieldError("TypeCode", ErrConstructor, addenda10.TypeCode)
	}
	if addenda10.TransactionTypeCode == "" {
		return fieldError("TransactionTypeCode", ErrFieldRequired, addenda10.TransactionTypeCode)
	}
	// ToDo:  Commented because it appears this value can be all 000 (maybe blank?)
	/*	if addenda10.ForeignPaymentAmount == 0 {
		return fieldError( "ForeignPaymentAmount", ErrFieldRequired,  strconv.Itoa(addenda10.ForeignPaymentAmount))
	}*/
	if addenda10.Name == "" {
		return fieldError("Name", ErrConstructor, addenda10.Name)
	}
	if addenda10.EntryDetailSequenceNumber == 0 {
		return fieldError("EntryDetailSequenceNumber", ErrConstructor, addenda10.EntryDetailSequenceNumberField())
	}
	return nil
}

// ForeignPaymentAmountField returns ForeignPaymentAmount zero padded
// ToDo: Review/Add logic for blank ?
func (addenda10 *Addenda10) ForeignPaymentAmountField() string {
	return addenda10.numericField(addenda10.ForeignPaymentAmount, 18)
}

// ForeignTraceNumberField gets the Foreign TraceNumber left padded
func (addenda10 *Addenda10) ForeignTraceNumberField() string {
	return addenda10.alphaField(addenda10.ForeignTraceNumber, 22)
}

// NameField gets the name field - Receiving Company Name/Individual Name left padded
func (addenda10 *Addenda10) NameField() string {
	return addenda10.alphaField(addenda10.Name, 35)
}

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda10 *Addenda10) EntryDetailSequenceNumberField() string {
	return addenda10.numericField(addenda10.EntryDetailSequenceNumber, 7)
}
