package ach

import (
	"flag"
	"fmt"
	"strings"
)

func init() {
	flag.Lookup("alsologtostderr").Value.Set("true")
}

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
	// Detail or Corporate Entry Detail Record's trace number This number is
	// the same as the last seven digits of the trace number of the related
	// Entry Detail Record or Corporate Entry Detail Record.
	EntryDetailSequenceNumber int
	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters
}

// AddendaParam is the minimal fields required to make a ach addenda
type AddendaParam struct {
	PaymentRelatedInfo string `json:"payment_related_info"`
}

// NewAddenda returns a new Addenda with default values for none exported fields
func NewAddenda(params ...AddendaParam) Addenda {
	addenda := Addenda{
		recordType:                "7",
		TypeCode:                  "05",
		SequenceNumber:            1,
		EntryDetailSequenceNumber: 1,
	}

	if len(params) > 0 {
		addenda.PaymentRelatedInformation = params[0].PaymentRelatedInfo
		return addenda
	}
	return addenda
}

// Parse takes the input record string and parses the Addenda values
func (addenda *Addenda) Parse(record string) {
	// 1-1 Always "7"
	addenda.recordType = "7"
	// 2-3 Defines the specific explanation and format for the addenda information contained in the same record
	addenda.TypeCode = record[1:3]
	// 4-83 Based on the information entered (04-83) 80 alphanumeric
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
func (addenda *Addenda) Validate() error {
	if err := addenda.fieldInclusion(); err != nil {
		return err
	}
	if addenda.recordType != "7" {
		msg := fmt.Sprintf(msgRecordType, 7)
		return &FieldError{FieldName: "recordType", Value: addenda.recordType, Msg: msg}
	}
	if err := addenda.isTypeCode(addenda.TypeCode); err != nil {
		return &FieldError{FieldName: "TypeCode", Value: addenda.TypeCode, Msg: err.Error()}
	}
	if err := addenda.isAlphanumeric(addenda.PaymentRelatedInformation); err != nil {
		return &FieldError{FieldName: "PaymentRelatedInformation", Value: addenda.PaymentRelatedInformation, Msg: err.Error()}
	}

	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addenda *Addenda) fieldInclusion() error {
	if addenda.recordType == "" {
		return &FieldError{FieldName: "recordType", Value: addenda.recordType, Msg: msgFieldInclusion}
	}
	if addenda.TypeCode == "" {
		return &FieldError{FieldName: "TypeCode", Value: addenda.TypeCode, Msg: msgFieldInclusion}
	}
	if addenda.SequenceNumber == 0 {
		return &FieldError{FieldName: "SequenceNumber", Value: addenda.SequenceNumberField(), Msg: msgFieldInclusion}
	}
	if addenda.EntryDetailSequenceNumber == 0 {
		return &FieldError{FieldName: "EntryDetailSequenceNumber", Value: addenda.EntryDetailSequenceNumberField(), Msg: msgFieldInclusion}
	}
	return nil
}

// PaymentRelatedInformationField returns a zero padded PaymentRelatedInformation string
func (addenda *Addenda) PaymentRelatedInformationField() string {
	return addenda.alphaField(addenda.PaymentRelatedInformation, 80)
}

// SequenceNumberField returns a zero padded SequenceNumber string
func (addenda *Addenda) SequenceNumberField() string {
	return addenda.numericField(addenda.SequenceNumber, 4)
}

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addenda *Addenda) EntryDetailSequenceNumberField() string {
	return addenda.numericField(addenda.EntryDetailSequenceNumber, 7)
}
