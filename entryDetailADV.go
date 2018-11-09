// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// EntryDetailADV contains the actual transaction data for an individual entry.
// Fields include those designating the entry as a deposit (credit) or
// withdrawal (debit), the transit routing number for the entry recipient’s financial
// institution, the account number (left justify,no zero fill), name, and dollar amount.
type EntryDetailADV struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record in the block. 6
	recordType string
	// TransactionCode representing Accounting Entries
	// Credit for ACH debits originated - 81
	// Debit for ACH credits originated - 82
	// Credit for ACH credits received 83
	// Debit for ACH debits received 84
	// Credit for ACH credits in rejected batches 85
	// Debit for ACH debits in rejected batches - 86
	// Summary credit for respondent ACH activity - 87
	// Summary debit for respondent ACH activity - 88
	TransactionCode int `json:"transactionCode"`
	// RDFIIdentification is the RDFI's routing number without the last digit.
	// Receiving Depository Financial Institution
	RDFIIdentification string `json:"RDFIIdentification"`
	// CheckDigit the last digit of the RDFI's routing number
	CheckDigit string `json:"checkDigit"`
	// DFIAccountNumber is the receiver's bank account number you are crediting/debiting.
	// It important to note that this is an alphanumeric field, so its space padded, no zero padded
	DFIAccountNumber string `json:"DFIAccountNumber"`
	// Amount Number of cents you are debiting/crediting this account
	Amount int `json:"amount"`
	// AdviceRoutingNumber
	AdviceRoutingNumber string `json:"adviceRoutingNumber"`
	// FileIdentification
	FileIdentification string `json:"fileIdentification,omitempty"`
	// ACHOperatorData
	ACHOperatorData string `json:"achOperatorData,omitempty"`
	// IndividualName The name of the receiver, usually the name on the bank account
	IndividualName string `json:"individualName"`
	// DiscretionaryData allows ODFIs to include codes, of significance only to them,
	// to enable specialized handling of the entry. There will be no
	// standardized interpretation for the value of this field. It can either
	// be a single two-character code, or two distinct one-character codes,
	// according to the needs of the ODFI and/or Originator involved. This
	// field must be returned intact for any returned entry.
	DiscretionaryData string `json:"discretionaryData,omitempty"`
	// AddendaRecordIndicator indicates the existence of an Addenda Record.
	// A value of "1" indicates that one ore more addenda records follow,
	// and "0" means no such record is present.
	AddendaRecordIndicator int `json:"addendaRecordIndicator,omitempty"`
	// ACHOperatorRoutingNumber
	ACHOperatorRoutingNumber string `json:"achOperatorRoutingNumber"`
	// JulianDateDay
	JulianDateDay int `json:"julianDateDay"`
	// SequenceNumber
	SequenceNumber int `json:"sequenceNumber,omitEmpty"`
	// Addenda98 for user with NOC
	Addenda98 *Addenda98 `json:"addenda98,omitempty"`
	// Addenda99 for use with Returns
	Addenda99 *Addenda99 `json:"addenda99,omitempty"`
	// Category defines if the entry is a Forward, Return, or NOC
	Category string `json:"category,omitempty"`
	// validator is composed for data validation
	validator
	// converters is composed for ACH to golang Converters
	converters
}

// NewEntryDetailADV returns a new EntryDetailADV with default values for non exported fields
func NewEntryDetailADV() *EntryDetailADV {
	entry := &EntryDetailADV{
		recordType: "6",
		Category:   CategoryForward,
	}
	return entry
}

//ToDo: ADV Specific Properties

// Parse takes the input record string and parses the EntryDetailADV values
func (ed *EntryDetailADV) Parse(record string) {
	if utf8.RuneCountInString(record) != 94 {
		return
	}

	// 1-1 Always "6"
	ed.recordType = "6"
	// 2-3 is checking credit 22 debit 27 savings credit 32 debit 37
	ed.TransactionCode = ed.parseNumField(record[1:3])
	// 4-11 the RDFI's routing number without the last digit.
	ed.RDFIIdentification = ed.parseStringField(record[3:11])
	// 12-12 The last digit of the RDFI's routing number
	ed.CheckDigit = ed.parseStringField(record[11:12])
	// 13-27 The receiver's bank account number you are crediting/debiting
	ed.DFIAccountNumber = record[12:27]
	// 28-39 Number of cents you are debiting/crediting this account
	ed.Amount = ed.parseNumField(record[27:39])
	// 40-48 Advice Routing Number
	ed.AdviceRoutingNumber = ed.parseStringField(record[39:48])
	// 49-53 File Identification
	ed.FileIdentification = ed.parseStringField(record[48:53])
	// 54-54 ACH Operator Data
	ed.ACHOperatorData = ed.parseStringField(record[53:54])
	// 55-76 Individual Name
	ed.IndividualName = record[54:76]
	// 77-78 allows ODFIs to include codes of significance only to them, normally blank
	ed.DiscretionaryData = record[76:78]
	// 79-79 1 if addenda exists 0 if it does not
	ed.AddendaRecordIndicator = ed.parseNumField(record[78:79])
	// 80-87
	ed.ACHOperatorRoutingNumber = ed.parseStringField(record[79:87])
	// 88-90
	ed.JulianDateDay = ed.parseNumField(record[87:90])
	// 91-94
	ed.SequenceNumber = ed.parseNumField(record[90:94])
}

// String writes the EntryDetailADV struct to a 94 character string.
func (ed *EntryDetailADV) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(ed.recordType)
	buf.WriteString(fmt.Sprintf("%v", ed.TransactionCode))
	buf.WriteString(ed.RDFIIdentificationField())
	buf.WriteString(ed.CheckDigit)
	buf.WriteString(ed.DFIAccountNumberField())
	buf.WriteString(ed.AmountField())
	buf.WriteString(ed.AdviceRoutingNumberField())
	buf.WriteString(ed.FileIdentificationField())
	buf.WriteString(ed.ACHOperatorDataField())
	buf.WriteString(ed.IndividualNameField())
	buf.WriteString(ed.DiscretionaryDataField())
	buf.WriteString(fmt.Sprintf("%v", ed.AddendaRecordIndicator))
	buf.WriteString(ed.ACHOperatorRoutingNumberField())
	buf.WriteString(ed.JulianDateDayField())
	buf.WriteString(ed.SequenceNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ed *EntryDetailADV) Validate() error {
	if err := ed.fieldInclusion(); err != nil {
		return err
	}
	if ed.recordType != "6" {
		msg := fmt.Sprintf(msgRecordType, 6)
		return &FieldError{FieldName: "recordType", Value: ed.recordType, Msg: msg}
	}
	if err := ed.isTransactionCode(ed.TransactionCode); err != nil {
		return &FieldError{FieldName: "TransactionCode", Value: strconv.Itoa(ed.TransactionCode), Msg: err.Error()}
	}
	if err := ed.isAlphanumeric(ed.DFIAccountNumber); err != nil {
		return &FieldError{FieldName: "DFIAccountNumber", Value: ed.DFIAccountNumber, Msg: err.Error()}
	}
	if err := ed.isAlphanumeric(ed.AdviceRoutingNumber); err != nil {
		return &FieldError{FieldName: "AdviceRoutingNumber", Value: ed.AdviceRoutingNumber, Msg: err.Error()}
	}
	if err := ed.isAlphanumeric(ed.IndividualName); err != nil {
		return &FieldError{FieldName: "IndividualName", Value: ed.IndividualName, Msg: err.Error()}
	}
	if err := ed.isAlphanumeric(ed.DiscretionaryData); err != nil {
		return &FieldError{FieldName: "DiscretionaryData", Value: ed.DiscretionaryData, Msg: err.Error()}
	}
	if err := ed.isAlphanumeric(ed.ACHOperatorRoutingNumber); err != nil {
		return &FieldError{FieldName: "ACHOperatorRoutingNumber", Value: ed.ACHOperatorRoutingNumber, Msg: err.Error()}
	}
	calculated := ed.CalculateCheckDigit(ed.RDFIIdentificationField())

	edCheckDigit, err := strconv.Atoi(ed.CheckDigit)
	if err != nil {
		return &FieldError{FieldName: "CheckDigit", Value: ed.CheckDigit, Msg: err.Error()}
	}

	if calculated != edCheckDigit {
		msg := fmt.Sprintf(msgValidCheckDigit, calculated)
		return &FieldError{FieldName: "RDFIIdentification", Value: ed.CheckDigit, Msg: msg}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (ed *EntryDetailADV) fieldInclusion() error {
	if ed.recordType == "" {
		return &FieldError{
			FieldName: "recordType",
			Value:     ed.recordType,
			Msg:       msgFieldInclusion + ", did you use NewEntryDetailADV()?",
		}
	}
	if ed.TransactionCode == 0 {
		return &FieldError{
			FieldName: "TransactionCode",
			Value:     strconv.Itoa(ed.TransactionCode),
			Msg:       msgFieldInclusion + ", did you use NewEntryDetailADV()?",
		}
	}
	if ed.RDFIIdentification == "" {
		return &FieldError{
			FieldName: "RDFIIdentification",
			Value:     ed.RDFIIdentificationField(),
			Msg:       msgFieldInclusion + ", did you use NewEntryDetailADV()?",
		}
	}
	if ed.DFIAccountNumber == "" {
		return &FieldError{
			FieldName: "DFIAccountNumber",
			Value:     ed.DFIAccountNumber,
			Msg:       msgFieldInclusion + ", did you use NewEntryDetailADV()?",
		}
	}
	if ed.AdviceRoutingNumber == "" {
		return &FieldError{
			FieldName: "AdviceRoutingNumber",
			Value:     ed.AdviceRoutingNumber,
			Msg:       msgFieldInclusion + ", did you use NewEntryDetailADV()?",
		}
	}
	if ed.IndividualName == "" {
		return &FieldError{
			FieldName: "IndividualName",
			Value:     ed.IndividualName,
			Msg:       msgFieldRequired + ", did you use NewEntryDetailADV()?",
		}
	}
	if ed.ACHOperatorRoutingNumber == "" {
		return &FieldError{
			FieldName: "ACHOperatorRoutingNumber",
			Value:     ed.ACHOperatorRoutingNumber,
			Msg:       msgFieldRequired + ", did you use NewEntryDetailADV()?",
		}
	}
	return nil
}

// SetRDFI takes the 9 digit RDFI account number and separates it for RDFIIdentification and CheckDigit
func (ed *EntryDetailADV) SetRDFI(rdfi string) *EntryDetailADV {
	s := ed.stringField(rdfi, 9)
	ed.RDFIIdentification = ed.parseStringField(s[:8])
	ed.CheckDigit = ed.parseStringField(s[8:9])
	return ed
}

// RDFIIdentificationField get the rdfiIdentification with zero padding
func (ed *EntryDetailADV) RDFIIdentificationField() string {
	return ed.stringField(ed.RDFIIdentification, 8)
}

// DFIAccountNumberField gets the DFIAccountNumber with space padding
func (ed *EntryDetailADV) DFIAccountNumberField() string {
	return ed.alphaField(ed.DFIAccountNumber, 17)
}

// AmountField returns a zero padded string of amount
func (ed *EntryDetailADV) AmountField() string {
	return ed.numericField(ed.Amount, 10)
}

// AdviceRoutingNumberField gets the AdviceRoutingNumber with zero padding
func (ed *EntryDetailADV) AdviceRoutingNumberField() string {
	return ed.stringField(ed.AdviceRoutingNumber, 9)
}

// FileIdentificationField returns a space padded string of FileIdentification
func (ed *EntryDetailADV) FileIdentificationField() string {
	return ed.alphaField(ed.FileIdentification, 5)
}

// ACHOperatorDataField returns a space padded string of ACHOperatorData
func (ed *EntryDetailADV) ACHOperatorDataField() string {
	return ed.alphaField(ed.ACHOperatorData, 1)
}

// IndividualNameField returns a space padded string of IndividualName
func (ed *EntryDetailADV) IndividualNameField() string {
	return ed.alphaField(ed.IndividualName, 22)
}

// DiscretionaryDataField returns a space padded string of DiscretionaryData
func (ed *EntryDetailADV) DiscretionaryDataField() string {
	return ed.alphaField(ed.DiscretionaryData, 2)
}

// ACHOperatorRoutingNumberField returns a space padded string of ACHOperatorRoutingNumber
func (ed *EntryDetailADV) ACHOperatorRoutingNumberField() string {
	return ed.alphaField(ed.ACHOperatorRoutingNumber, 8)
}

// JulianDateDayField returns a zero padded string of JulianDateDay
func (ed *EntryDetailADV) JulianDateDayField() string {
	return ed.numericField(ed.JulianDateDay, 3)
}

// SequenceNumberField returns a zero padded string of SequenceNumber
func (ed *EntryDetailADV) SequenceNumberField() string {
	return ed.numericField(ed.SequenceNumber, 4)
}

// CreditOrDebit returns a "C" for credit or "D" for debit based on the entry TransactionCode
func (ed *EntryDetailADV) CreditOrDebit() string {
	if ed.TransactionCode < 10 || ed.TransactionCode > 99 {
		return ""
	}
	tc := strconv.Itoa(ed.TransactionCode)

	// take the second number in the TransactionCode
	switch tc[1:2] {
	case "1", "3", "5", "7":
		return "C"
	case "2", "4", "6", "8":
		return "D"
	default:
	}
	return ""
}
