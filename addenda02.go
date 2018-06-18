// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strings"
)

// Addenda02 is a Addendumer addenda which provides business transaction information for Addenda Type
// Code 02 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
type Addenda02 struct {
	//ToDo: Verify which fields should be omitempty
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record in the block. entryAddenda02 Pos 7
	recordType string
	// TypeCode Addenda02 type code '02'
	typeCode string
	// ReferenceInformationOne may be used for additional reference numbers, identification numbers,
	// or codes that the merchant needs to identify the particular transaction or customer.
	ReferenceInformationOne string `json:"referenceInformationOne, omitempty"`
	// ReferenceInformationTwo  may be used for additional reference numbers, identification numbers,
	// or codes that the merchant needs to identify the particular transaction or customer.
	ReferenceInformationTwo string `json:"referenceInformationTwo, omitempty"`
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
	AuthorizationCodeOrExpireDate string `json:"authorizationCodeOrExpireDate, omitempty"`
	// Terminal Location identifies the specific location of a terminal (i.e., street names of an
	// intersection, address, etc.) in accordance with the requirements of Regulation E.
	TerminalLocation string `json:"terminalLocation"`
	// TerminalCity Identifies the city in which the electronic terminal is located.
	TerminalCity string `json:"terminalCity"`
	// TerminalState Identifies the state in which the electronic terminal is located
	TerminalState string `json:"terminalState"`
	// TraceNumber Standard Entry Detail Trace Number
	TraceNumber int `json:"traceNumber,omitempty"`
	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters
}

// NewAddenda02 returns a new Addenda02 with default values for none exported fields
func NewAddenda02() *Addenda02 {
	addenda02 := new(Addenda02)
	addenda02.recordType = "7"
	addenda02.typeCode = "02"
	return addenda02
}

// Parse takes the input record string and parses the Addenda02 values
func (addenda02 *Addenda02) Parse(record string) {
	// 1-1 Always "7"
	addenda02.recordType = "7"
	// 2-3 Always 02
	addenda02.typeCode = record[1:3]
	// 4-10 Based on the information entered (04-10) 7 alphanumeric
	addenda02.ReferenceInformationOne = strings.TrimSpace(record[3:10])
	// 11-13 Based on the information entered (11-13) 3 alphanumeric
	addenda02.ReferenceInformationTwo = strings.TrimSpace(record[10:13])
	// 14-19
	addenda02.TerminalIdentificationCode = strings.TrimSpace(record[13:19])
	// 20-25
	addenda02.TransactionSerialNumber = strings.TrimSpace(record[19:25])
	// 26-29
	addenda02.TransactionDate = strings.TrimSpace(record[25:29])
	// 30-35
	addenda02.AuthorizationCodeOrExpireDate = strings.TrimSpace(record[29:35])
	// 36-62
	addenda02.TerminalLocation = strings.TrimSpace(record[35:62])
	// 63-77
	addenda02.TerminalCity = strings.TrimSpace(record[62:77])
	// 78-79
	addenda02.TerminalState = strings.TrimSpace(record[77:79])
	// 80-94
	addenda02.TraceNumber = addenda02.parseNumField(record[79:94])
}

// String writes the Addenda02 struct to a 94 character string.
func (addenda02 *Addenda02) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v",
		addenda02.recordType,
		addenda02.typeCode,
		addenda02.ReferenceInformationOneField(),
		addenda02.ReferenceInformationTwoField(),
		addenda02.TerminalIdentificationCodeField(),
		addenda02.TransactionSerialNumberField(),
		// ToDo: Follow up on best way to get TransactionDate - should it be treated as an alpha field
		addenda02.TransactionDateField(),
		addenda02.AuthorizationCodeOrExpireDateField(),
		addenda02.TerminalLocationField(),
		addenda02.TerminalCityField(),
		addenda02.TerminalStateField(),
		addenda02.TraceNumberField(),
	)
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda02 *Addenda02) Validate() error {
	if err := addenda02.fieldInclusion(); err != nil {
		return err
	}
	if addenda02.recordType != "7" {
		msg := fmt.Sprintf(msgRecordType, 7)
		return &FieldError{FieldName: "recordType", Value: addenda02.recordType, Msg: msg}
	}
	if err := addenda02.isTypeCode(addenda02.typeCode); err != nil {
		return &FieldError{FieldName: "TypeCode", Value: addenda02.typeCode, Msg: err.Error()}
	}
	// Type Code must be 02
	// ToDo: Evaluate if Addenda05 and Addenda99 should be modified to validate on 05 and 99
	if addenda02.typeCode != "02" {
		return &FieldError{FieldName: "TypeCode", Value: addenda02.typeCode, Msg: msgAddendaTypeCode}
	}
	// TransactionDate Addenda02 ACH File format is MMDD.  Validate MM is 01-12.
	if err := addenda02.isMonth(addenda02.parseStringField(addenda02.TransactionDate[0:2])); err != nil {
		return &FieldError{FieldName: "TransactionDate", Value: addenda02.parseStringField(addenda02.TransactionDate[0:2]), Msg: msgValidMonth}
	}
	// TransactionDate Addenda02 ACH File format is MMDD.  If the month is valid, validate the day for the
	// month 01-31 depending on month.
	if err := addenda02.isDay(addenda02.parseStringField(addenda02.TransactionDate[0:2]), addenda02.parseStringField(addenda02.TransactionDate[2:4])); err != nil {
		return &FieldError{FieldName: "TransactionDate", Value: addenda02.parseStringField(addenda02.TransactionDate[0:2]), Msg: msgValidDay}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values  and rquired fields are defined. If fields are
// invalid the ACH transfer will be returned.

// ToDo: check if we should do fieldInclusion or validate on required fields

func (addenda02 *Addenda02) fieldInclusion() error {
	if addenda02.recordType == "" {
		return &FieldError{FieldName: "recordType", Value: addenda02.recordType, Msg: msgFieldInclusion}
	}
	if addenda02.typeCode == "" {
		return &FieldError{FieldName: "TypeCode", Value: addenda02.typeCode, Msg: msgFieldInclusion}
	}
	// Required Fields
	if addenda02.TerminalIdentificationCode == "" {
		return &FieldError{FieldName: "TerminalIdentificationCode", Value: addenda02.TerminalIdentificationCode, Msg: msgFieldRequired}
	}
	if addenda02.TransactionSerialNumber == "" {
		return &FieldError{FieldName: "TransactionSerialNumber", Value: addenda02.TransactionSerialNumber, Msg: msgFieldRequired}
	}
	if addenda02.TransactionDate == "" {
		return &FieldError{FieldName: "TransactionDate", Value: addenda02.TransactionDate, Msg: msgFieldRequired}
	}
	if addenda02.TerminalLocation == "" {
		return &FieldError{FieldName: "TerminalLocation", Value: addenda02.TerminalLocation, Msg: msgFieldRequired}
	}
	if addenda02.TerminalCity == "" {
		return &FieldError{FieldName: "TerminalCity", Value: addenda02.TerminalCity, Msg: msgFieldRequired}
	}
	if addenda02.TerminalState == "" {
		return &FieldError{FieldName: "TerminalState", Value: addenda02.TerminalState, Msg: msgFieldRequired}
	}
	return nil
}

// TypeCode Defines the specific explanation and format for the addenda02 information
func (addenda02 *Addenda02) TypeCode() string {
	return addenda02.typeCode
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
	return addenda02.TransactionDate
}

// AuthorizationCodeOrExpireDateField returns a space padded AuthorizationCodeOrExpireDate string
func (addenda02 *Addenda02) AuthorizationCodeOrExpireDateField() string {
	return addenda02.alphaField(addenda02.AuthorizationCodeOrExpireDate, 6)
}

//TerminalLocationField returns a space padded TerminalLocation string
func (addenda02 *Addenda02) TerminalLocationField() string {
	return addenda02.alphaField(addenda02.TerminalLocation, 27)
}

//TerminalCityField returns a space padded TerminalCity string
func (addenda02 *Addenda02) TerminalCityField() string {
	return addenda02.alphaField(addenda02.TerminalCity, 15)
}

//TerminalStateField returns a space padded TerminalState string
func (addenda02 *Addenda02) TerminalStateField() string {
	return addenda02.alphaField(addenda02.TerminalState, 2)
}

// TraceNumberField returns a space padded traceNumber string
func (addenda02 *Addenda02) TraceNumberField() string {
	return addenda02.numericField(addenda02.TraceNumber, 15)
}
