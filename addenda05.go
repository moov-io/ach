// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strings"
)

// Addenda05 is a Addendumer addenda which provides business transaction information for Addenda Type
// Code 05 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
// It is used for the following StandardEntryClassCode: ACK, ATX, CCD, CIE, CTX, DNE, ENR, WEB, PPD, TRX.
type Addenda05 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record in the block. entryAddenda05 Pos 7
	recordType string
	// TypeCode Addenda05 types code '05'
	TypeCode string `json:"typeCode"`
	// PaymentRelatedInformation
	PaymentRelatedInformation string `json:"paymentRelatedInformation"`
	// SequenceNumber is consecutively assigned to each Addenda05 Record following
	// an Entry Detail Record. The first addenda05 sequence number must always
	// be a "1".
	SequenceNumber int `json:"sequenceNumber,omitempty"`
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

// NewAddenda05 returns a new Addenda05 with default values for none exported fields
func NewAddenda05() *Addenda05 {
	addenda05 := new(Addenda05)
	addenda05.recordType = "7"
	addenda05.TypeCode = "05"
	return addenda05
}

// Parse takes the input record string and parses the Addenda05 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm successful parsing and data validity.
func (addenda05 *Addenda05) Parse(record string) {
	// 1-1 Always "7"
	addenda05.recordType = "7"
	// 2-3 Always 05
	addenda05.TypeCode = record[1:3]
	// 4-83 Based on the information entered (04-83) 80 alphanumeric
	addenda05.PaymentRelatedInformation = strings.TrimSpace(record[3:83])
	// 84-87 SequenceNumber is consecutively assigned to each Addenda05 Record following
	// an Entry Detail Record
	addenda05.SequenceNumber = addenda05.parseNumField(record[83:87])
	// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
	addenda05.EntryDetailSequenceNumber = addenda05.parseNumField(record[87:94])
}

// String writes the Addenda05 struct to a 94 character string.
func (addenda05 *Addenda05) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(addenda05.recordType)
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
	if addenda05.recordType != "7" {
		msg := fmt.Sprintf(msgRecordType, 7)
		return &FieldError{FieldName: "recordType", Value: addenda05.recordType, Msg: msg}
	}
	if err := addenda05.isTypeCode(addenda05.TypeCode); err != nil {
		return &FieldError{FieldName: "TypeCode", Value: addenda05.TypeCode, Msg: err.Error()}
	}
	// Type Code must be 05
	if addenda05.TypeCode != "05" {
		return &FieldError{FieldName: "TypeCode", Value: addenda05.TypeCode, Msg: msgAddendaTypeCode}
	}
	if err := addenda05.isAlphanumeric(addenda05.PaymentRelatedInformation); err != nil {
		return &FieldError{FieldName: "PaymentRelatedInformation", Value: addenda05.PaymentRelatedInformation, Msg: err.Error()}
	}

	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda05 *Addenda05) fieldInclusion() error {
	if addenda05.recordType == "" {
		return &FieldError{
			FieldName: "recordType",
			Value:     addenda05.recordType,
			Msg:       msgFieldInclusion + ", did you use NewAddenda05()?",
		}
	}
	if addenda05.TypeCode == "" {
		return &FieldError{
			FieldName: "TypeCode",
			Value:     addenda05.TypeCode,
			Msg:       msgFieldInclusion + ", did you use NewAddenda05()?",
		}
	}
	if addenda05.SequenceNumber == 0 {
		return &FieldError{
			FieldName: "SequenceNumber",
			Value:     addenda05.SequenceNumberField(),
			Msg:       msgFieldInclusion + ", did you use NewAddenda05()?",
		}
	}
	if addenda05.EntryDetailSequenceNumber == 0 {
		return &FieldError{
			FieldName: "EntryDetailSequenceNumber",
			Value:     addenda05.EntryDetailSequenceNumberField(),
			Msg:       msgFieldInclusion + ", did you use NewAddenda05()?",
		}
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
