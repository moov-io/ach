// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strings"
)

// Addenda15 is an addenda which provides business transaction information for Addenda Type
// Code 15 in a machine readable format. It is usually formatted according to ANSI, ASC, X12 Standard.
//
// Addenda15 is mandatory for IAT entries
//
// The Addenda15 record identifies key information related to the Receiver.
type Addenda15 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record in the block.
	recordType string
	// TypeCode Addenda15 types code '15'
	TypeCode string `json:"typeCode"`
	// Receiver Identification Number contains the accounting number by which the Originator is known to
	// the Receiver for descriptive purposes. NACHA Rules recommend but do not require the RDFI to print
	// the contents of this field on the receiver's statement.
	ReceiverIDNumber string `json:"receiverIDNumber,omitempty"`
	// Receiver Street Address contains the Receiver‟s physical address
	ReceiverStreetAddress string `json:"receiverStreetAddress"`
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

// NewAddenda15 returns a new Addenda15 with default values for none exported fields
func NewAddenda15() *Addenda15 {
	addenda15 := new(Addenda15)
	addenda15.recordType = "7"
	addenda15.TypeCode = "15"
	return addenda15
}

// Parse takes the input record string and parses the Addenda15 values
func (addenda15 *Addenda15) Parse(record string) {
	// 1-1 Always "7"
	addenda15.recordType = "7"
	// 2-3 Always 15
	addenda15.TypeCode = record[1:3]
	// 4-18
	addenda15.ReceiverIDNumber = addenda15.parseStringField(record[3:18])
	// 19-53
	addenda15.ReceiverStreetAddress = strings.TrimSpace(record[18:53])
	// 54-87 reserved - Leave blank
	addenda15.reserved = "                                  "
	// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
	addenda15.EntryDetailSequenceNumber = addenda15.parseNumField(record[87:94])
}

// String writes the Addenda15 struct to a 94 character string.
func (addenda15 *Addenda15) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(addenda15.recordType)
	buf.WriteString(addenda15.TypeCode)
	buf.WriteString(addenda15.ReceiverIDNumberField())
	buf.WriteString(addenda15.ReceiverStreetAddressField())
	buf.WriteString(addenda15.reservedField())
	buf.WriteString(addenda15.EntryDetailSequenceNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda15 *Addenda15) Validate() error {
	if err := addenda15.fieldInclusion(); err != nil {
		return err
	}
	if addenda15.recordType != "7" {
		msg := fmt.Sprintf(msgRecordType, 7)
		return &FieldError{FieldName: "recordType", Value: addenda15.recordType, Msg: msg}
	}
	if err := addenda15.isTypeCode(addenda15.TypeCode); err != nil {
		return &FieldError{FieldName: "TypeCode", Value: addenda15.TypeCode, Msg: err.Error()}
	}
	// Type Code must be 15
	if addenda15.TypeCode != "15" {
		return &FieldError{FieldName: "TypeCode", Value: addenda15.TypeCode, Msg: msgAddendaTypeCode}
	}
	if err := addenda15.isAlphanumeric(addenda15.ReceiverIDNumber); err != nil {
		return &FieldError{FieldName: "ReceiverIDNumber", Value: addenda15.ReceiverIDNumber, Msg: err.Error()}
	}
	if err := addenda15.isAlphanumeric(addenda15.ReceiverStreetAddress); err != nil {
		return &FieldError{FieldName: "ReceiverStreetAddress", Value: addenda15.ReceiverStreetAddress, Msg: err.Error()}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda15 *Addenda15) fieldInclusion() error {
	if addenda15.recordType == "" {
		return &FieldError{
			FieldName: "recordType",
			Value:     addenda15.recordType,
			Msg:       msgFieldInclusion + ", did you use NewAddenda15()?",
		}
	}
	if addenda15.TypeCode == "" {
		return &FieldError{
			FieldName: "TypeCode",
			Value:     addenda15.TypeCode,
			Msg:       msgFieldInclusion + ", did you use NewAddenda15()?",
		}
	}
	if addenda15.ReceiverStreetAddress == "" {
		return &FieldError{
			FieldName: "ReceiverStreetAddress",
			Value:     addenda15.ReceiverStreetAddress,
			Msg:       msgFieldInclusion + ", did you use NewAddenda15()?",
		}
	}
	if addenda15.EntryDetailSequenceNumber == 0 {
		return &FieldError{
			FieldName: "EntryDetailSequenceNumber",
			Value:     addenda15.EntryDetailSequenceNumberField(),
			Msg:       msgFieldInclusion + ", did you use NewAddenda15()?",
		}
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

// reservedField gets reserved - blank space
func (addenda15 *Addenda15) reservedField() string {
	return addenda15.alphaField(addenda15.reserved, 34)
}

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda15 *Addenda15) EntryDetailSequenceNumberField() string {
	return addenda15.numericField(addenda15.EntryDetailSequenceNumber, 7)
}
