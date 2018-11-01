// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strings"
)

// Addenda10 is an addenda which provides business transaction information for Addenda Type
// Code 10 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
//
// Addenda10 is mandatory for IAT entries
//
// The Addenda10 Record identifies the Receiver of the transaction and the dollar amount of
// the payment.
type Addenda10 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record in the block.
	recordType string
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
	// reserved - Leave blank
	reserved string
	// EntryDetailSequenceNumber contains the ascending sequence number section of the Entry
	// Detail or Corporate Entry Detail Record's trace number This number is
	// the same as the last seven digits of the trace number of the related
	// Entry Detail Record or Corporate Entry Detail Record.
	EntryDetailSequenceNumber int `json:"entryDetailSequenceNumber,omitempty"`
	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters
}

// NewAddenda10 returns a new Addenda10 with default values for none exported fields
func NewAddenda10() *Addenda10 {
	addenda10 := new(Addenda10)
	addenda10.recordType = "7"
	addenda10.TypeCode = "10"
	return addenda10
}

// Parse takes the input record string and parses the Addenda10 values
func (addenda10 *Addenda10) Parse(record string) {
	// 1-1 Always "7"
	addenda10.recordType = "7"
	// 2-3 Always 10
	addenda10.TypeCode = record[1:3]
	// 04-06 Describes the type of payment
	addenda10.TransactionTypeCode = record[3:6]
	// 07-24 Payment Amount	For inbound IAT payments this field should contain the USD amount or may be blank.
	addenda10.ForeignPaymentAmount = addenda10.parseNumField(record[06:24])
	//  25-46 Insert blanks or zeros
	addenda10.ForeignTraceNumber = strings.TrimSpace(record[24:46])
	// 47-81 Receiving Company Name/Individual Name
	addenda10.Name = strings.TrimSpace(record[46:81])
	// 82-87 reserved - Leave blank
	addenda10.reserved = "      "
	// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
	addenda10.EntryDetailSequenceNumber = addenda10.parseNumField(record[87:94])
}

// String writes the Addenda10 struct to a 94 character string.
func (addenda10 *Addenda10) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(addenda10.recordType)
	buf.WriteString(addenda10.TypeCode)
	// TransactionTypeCode Validator
	buf.WriteString(addenda10.TransactionTypeCode)
	buf.WriteString(addenda10.ForeignPaymentAmountField())
	buf.WriteString(addenda10.ForeignTraceNumberField())
	buf.WriteString(addenda10.NameField())
	buf.WriteString(addenda10.reservedField())
	buf.WriteString(addenda10.EntryDetailSequenceNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda10 *Addenda10) Validate() error {
	if err := addenda10.fieldInclusion(); err != nil {
		return err
	}
	if addenda10.recordType != "7" {
		msg := fmt.Sprintf(msgRecordType, 7)
		return &FieldError{FieldName: "recordType", Value: addenda10.recordType, Msg: msg}
	}
	if err := addenda10.isTypeCode(addenda10.TypeCode); err != nil {
		return &FieldError{FieldName: "TypeCode", Value: addenda10.TypeCode, Msg: err.Error()}
	}
	// Type Code must be 10
	if addenda10.TypeCode != "10" {
		return &FieldError{FieldName: "TypeCode", Value: addenda10.TypeCode, Msg: msgAddendaTypeCode}
	}
	if err := addenda10.isTransactionTypeCode(addenda10.TransactionTypeCode); err != nil {
		return &FieldError{FieldName: "TransactionTypeCode", Value: addenda10.TransactionTypeCode, Msg: err.Error()}
	}
	// ToDo: Foreign Payment Amount blank ?
	if err := addenda10.isAlphanumeric(addenda10.ForeignTraceNumber); err != nil {
		return &FieldError{FieldName: "ForeignTraceNumber", Value: addenda10.ForeignTraceNumber, Msg: err.Error()}
	}
	if err := addenda10.isAlphanumeric(addenda10.Name); err != nil {
		return &FieldError{FieldName: "Name", Value: addenda10.Name, Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda10 *Addenda10) fieldInclusion() error {
	if addenda10.recordType == "" {
		return &FieldError{
			FieldName: "recordType",
			Value:     addenda10.recordType,
			Msg:       msgFieldInclusion + ", did you use NewAddenda10()?",
		}
	}
	if addenda10.TypeCode == "" {
		return &FieldError{
			FieldName: "TypeCode",
			Value:     addenda10.TypeCode,
			Msg:       msgFieldInclusion + ", did you use NewAddenda10()?",
		}
	}
	if addenda10.TransactionTypeCode == "" {
		return &FieldError{
			FieldName: "TransactionTypeCode",
			Value:     addenda10.TransactionTypeCode,
			Msg:       msgFieldRequired,
		}
	}
	// ToDo:  Commented because it appears this value can be all 000 (maybe blank?)
	/*	if addenda10.ForeignPaymentAmount == 0 {
		return &FieldError{FieldName: "ForeignPaymentAmount",
			Value: strconv.Itoa(addenda10.ForeignPaymentAmount), Msg: msgFieldRequired}
	}*/
	if addenda10.Name == "" {
		return &FieldError{
			FieldName: "Name",
			Value:     addenda10.Name,
			Msg:       msgFieldInclusion + ", did you use NewAddenda10()?",
		}
	}
	if addenda10.EntryDetailSequenceNumber == 0 {
		return &FieldError{
			FieldName: "EntryDetailSequenceNumber",
			Value:     addenda10.EntryDetailSequenceNumberField(),
			Msg:       msgFieldInclusion + ", did you use NewAddenda10()?",
		}
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

// reservedField gets reserved - blank space
func (addenda10 *Addenda10) reservedField() string {
	return addenda10.alphaField(addenda10.reserved, 6)
}

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda10 *Addenda10) EntryDetailSequenceNumberField() string {
	return addenda10.numericField(addenda10.EntryDetailSequenceNumber, 7)
}
