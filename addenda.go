package ach

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Errors specific to a Batach Addenda Record
var (
	ErrAddendaTypeCode = errors.New("Invalid Addenda Type Code")
)

// Addenda provides business transaction information in a machine
// readable format. It is usually formatted according to ANSI, ASC, X12 Standard
type Addenda struct {
	// RecordType defines the type of record in the block. entryAddendaPos 7
	recordType string
	// TypeCode Addenda types code '05'
	TypeCode string
	// PaymentRelatedInformation
	PaymentRelatedInformation string
	// SequenceNumber is consecutively assigned to each Addenda Record following
	// an Entry Detail Record. The first addenda sequence number must always
	// be a "1".
	SequenceNumber int
	// EntryDetailSequenceNumber contains the ascending sequence number section of the Entry
	// Detail or Corporate Entry Detail Record's trace numbe This number is
	// the same as the last seven digits of the trace number of the related
	// Entry Detail Record or Corporate Entry Detail Record.
	EntryDetailSequenceNumber int
	// Validator is composed for data validation
	Validator
	// Converters is composed for ACH to golang Converters
	Converters
}

// NewAddenda returns a new Addenda with default values for none exported fields
func NewAddenda() *Addenda {
	return &Addenda{
		recordType: "7",
	}
}

// Parse takes the input record string and parses the Addenda values
func (addenda *Addenda) Parse(record string) {
	// 1-1 Always "7"
	addenda.recordType = "7"
	// 2-3 Defines the specific explanation and format for the addenda information contained in the same record
	addenda.TypeCode = record[1:3]
	// 4-83 Based on the information entere (04-83) 80 alphanumeric
	addenda.PaymentRelatedInformation = strings.TrimSpace(record[3:83])
	// 84-87 SequenceNumber is consecutively assigned to each Addenda Record following
	// an Entry Detail Record
	addenda.SequenceNumber = addenda.parseNumField(record[83:87])
	// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
	addenda.EntryDetailSequenceNumber = addenda.parseNumField(record[87:94])

}

// String writes the Addenda struct to a 94 character string.
func (addenda *Addenda) String() string {
	return fmt.Sprintf("%v%v%v%v%v",
		addenda.recordType,
		addenda.TypeCode,
		addenda.PaymentRelatedInformationField(),
		addenda.SequenceNumberField(),
		addenda.EntryDetailSequenceNumberField())
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addenda *Addenda) Validate() (bool, error) {
	v, err := addenda.fieldInclusion()
	if !v {
		return false, error(err)
	}

	if addenda.recordType != "7" {
		return false, ErrRecordType
	}
	if !addenda.isTypeCode(addenda.TypeCode) {
		return false, ErrAddendaTypeCode
	}

	if !addenda.isAlphanumeric(addenda.PaymentRelatedInformation) {
		return false, ErrValidAlphanumeric
	}

	return true, nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda *Addenda) fieldInclusion() (bool, error) {
	if addenda.recordType == "" &&
		addenda.TypeCode == "" &&
		addenda.SequenceNumber == 0 &&
		addenda.EntryDetailSequenceNumber == 0 {
		return false, ErrValidFieldInclusion
	}
	return true, nil
}

// PaymentRelatedInformationField returns a zero padded PaymentRelatedInformation string
func (addenda *Addenda) PaymentRelatedInformationField() string {
	return addenda.rightPad(addenda.PaymentRelatedInformation, " ", 80)
}

// SequenceNumberField returns a zero padded SequenceNumber string
func (addenda *Addenda) SequenceNumberField() string {
	return addenda.leftPad(strconv.Itoa(addenda.SequenceNumber), "0", 4)
}

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda *Addenda) EntryDetailSequenceNumberField() string {
	return addenda.leftPad(strconv.Itoa(addenda.EntryDetailSequenceNumber), "0", 7)
}
