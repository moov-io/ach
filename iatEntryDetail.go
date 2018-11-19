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

// IATEntryDetail contains the actual transaction data for an individual entry.
// Fields include those designating the entry as a deposit (credit) or
// withdrawal (debit), the transit routing number for the entry recipient’s financial
// institution, the account number (left justify,no zero fill), name, and dollar amount.
type IATEntryDetail struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record in the block. 6
	recordType string
	// TransactionCode if the receivers account is:
	// Credit (deposit) to checking account ‘22’
	// Prenote for credit to checking account ‘23’
	// Debit (withdrawal) to checking account ‘27’
	// Prenote for debit to checking account ‘28’
	// Credit to savings account ‘32’
	// Prenote for credit to savings account ‘33’
	// Debit to savings account ‘37’
	// Prenote for debit to savings account ‘38’
	TransactionCode int `json:"transactionCode"`
	// RDFIIdentification is the RDFI's routing number without the last digit.
	// Receiving Depository Financial Institution
	RDFIIdentification string `json:"RDFIIdentification"`
	// CheckDigit the last digit of the RDFI's routing number
	CheckDigit string `json:"checkDigit"`
	// AddendaRecords is the number of Addenda Records
	AddendaRecords int `json:"AddendaRecords"`
	// reserved - Leave blank
	reserved string
	// Amount Number of cents you are debiting/crediting this account
	Amount int `json:"amount"`
	// DFIAccountNumber is the receiver's bank account number you are crediting/debiting.
	// It important to note that this is an alphanumeric field, so its space padded, no zero padded
	DFIAccountNumber string `json:"DFIAccountNumber"`
	// reservedTwo - Leave blank
	reservedTwo string
	// OFACSreeningIndicator - Leave blank
	OFACSreeningIndicator string `json:"OFACSreeningIndicator"`
	// SecondaryOFACSreeningIndicator - Leave blank
	SecondaryOFACSreeningIndicator string `json:"SecondaryOFACSreeningIndicator"`
	// AddendaRecordIndicator indicates the existence of an Addenda Record.
	// A value of "1" indicates that one or more addenda records follow,
	// and "0" means no such record is present.
	AddendaRecordIndicator int `json:"addendaRecordIndicator,omitempty"`
	// TraceNumber assigned by the ODFI in ascending sequence, is included in each
	// Entry Detail Record, Corporate Entry Detail Record, and addenda Record.
	// Trace Numbers uniquely identify each entry within a batch in an ACH input file.
	// In association with the Batch Number, transmission (File Creation) Date,
	// and File ID Modifier, the Trace Number uniquely identifies an entry within a given file.
	// For addenda Records, the Trace Number will be identical to the Trace Number
	// in the associated Entry Detail Record, since the Trace Number is associated
	// with an entry or item rather than a physical record.
	//
	// Use TraceNumberField() for a properly formatted string representation.
	TraceNumber string `json:"traceNumber,omitempty"`
	// Addenda10 is mandatory for IAT entries
	//
	// The Addenda10 Record identifies the Receiver of the transaction and the dollar amount of
	// the payment.
	Addenda10 *Addenda10 `json:"addenda10,omitempty"`
	// Addenda11 is mandatory for IAT entries
	//
	// The Addenda11 record identifies key information related to the Originator of
	// the entry.
	Addenda11 *Addenda11 `json:"addenda11,omitempty"`
	// Addenda12 is mandatory for IAT entries
	//
	// The Addenda12 record identifies key information related to the Originator of
	// the entry.
	Addenda12 *Addenda12 `json:"addenda12,omitempty"`
	// Addenda13 is mandatory for IAT entries
	//
	// The Addenda13 contains information related to the financial institution originating the entry.
	// For inbound IAT entries, the Fourth Addenda Record must contain information to identify the
	// foreign financial institution that is providing the funding and payment instruction for
	// the IAT entry.
	Addenda13 *Addenda13 `json:"addenda13,omitempty"`
	// Addenda14 is mandatory for IAT entries
	//
	// The Addenda14 identifies the Receiving financial institution holding the Receiver's account.
	Addenda14 *Addenda14 `json:"addenda14,omitempty"`
	// Addenda15 is mandatory for IAT entries
	//
	// The Addenda15 record identifies key information related to the Receiver.
	Addenda15 *Addenda15 `json:"addenda15,omitempty"`
	// Addenda16 is mandatory for IAt entries
	//
	// Addenda16 record identifies additional key information related to the Receiver.
	Addenda16 *Addenda16 `json:"addenda16,omitempty"`
	// Addenda17 is optional for IAT entries
	//
	// This is an optional Addenda Record used to provide payment-related data. There i a maximum of up to two of these
	// Addenda Records with each IAT entry.
	Addenda17 []*Addenda17 `json:"addenda17,omitempty"`
	// Addenda18 is optional for IAT entries
	//
	// This optional addenda record is used to provide information on each Foreign Correspondent Bank involved in the
	// processing of the IAT entry. If no Foreign Correspondent Bank is involved,the record should not be included.
	// A maximum of five Addenda18 records may be included with each IAT entry.
	Addenda18 []*Addenda18 `json:"addenda18,omitempty"`
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

// NewIATEntryDetail returns a new IATEntryDetail with default values for non exported fields
func NewIATEntryDetail() *IATEntryDetail {
	iatEd := &IATEntryDetail{
		recordType:             "6",
		Category:               CategoryForward,
		AddendaRecordIndicator: 1,
	}
	return iatEd
}

// Parse takes the input record string and parses the EntryDetail values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm successful parsing and data validity.
func (ed *IATEntryDetail) Parse(record string) {
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
	// 13-16 Number of addenda records
	ed.AddendaRecords = ed.parseNumField(record[12:16])
	// 17-29 reserved - Leave blank
	ed.reserved = "             "
	// 30-39 Number of cents you are debiting/crediting this account
	ed.Amount = ed.parseNumField(record[29:39])
	// 40-74 The foreign receiver's account number you are crediting/debiting
	ed.DFIAccountNumber = record[39:74]
	// 75-76 reserved2 Leave blank
	ed.reservedTwo = "  "
	// 77 OFACScreeningIndicator
	ed.OFACSreeningIndicator = " "
	// 78-78 Secondary SecondaryOFACSreeningIndicator
	ed.SecondaryOFACSreeningIndicator = " "
	// 79-79 1 if addenda exists 0 if it does not
	//ed.AddendaRecordIndicator = 1
	ed.AddendaRecordIndicator = ed.parseNumField(record[78:79])
	// 80-94 An internal identification (alphanumeric) that you use to uniquely identify
	// this Entry Detail Record This number should be unique to the transaction and will help identify the transaction in case of an inquiry
	ed.TraceNumber = strings.TrimSpace(record[79:94])
}

// String writes the EntryDetail struct to a 94 character string.
func (ed *IATEntryDetail) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(ed.recordType)
	buf.WriteString(fmt.Sprintf("%v", ed.TransactionCode))
	buf.WriteString(ed.RDFIIdentificationField())
	buf.WriteString(ed.CheckDigit)
	buf.WriteString(ed.AddendaRecordsField())
	buf.WriteString(ed.reservedField())
	buf.WriteString(ed.AmountField())
	buf.WriteString(ed.DFIAccountNumberField())
	buf.WriteString(ed.reservedTwoField())
	buf.WriteString(ed.OFACSreeningIndicatorField())
	buf.WriteString(ed.SecondaryOFACSreeningIndicatorField())
	buf.WriteString(fmt.Sprintf("%v", ed.AddendaRecordIndicator))
	buf.WriteString(ed.TraceNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ed *IATEntryDetail) Validate() error {
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
	// CheckDigit calculations
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
func (ed *IATEntryDetail) fieldInclusion() error {
	if ed.recordType == "" {
		return &FieldError{
			FieldName: "recordType",
			Value:     ed.recordType,
			Msg:       msgFieldInclusion + ", did you use NewIATEntryDetail()?",
		}
	}
	if ed.TransactionCode == 0 {
		return &FieldError{
			FieldName: "TransactionCode",
			Value:     strconv.Itoa(ed.TransactionCode),
			Msg:       msgFieldInclusion + ", did you use NewIATEntryDetail()?",
		}
	}
	if ed.RDFIIdentification == "" {
		return &FieldError{
			FieldName: "RDFIIdentification",
			Value:     ed.RDFIIdentificationField(),
			Msg:       msgFieldInclusion + ", did you use NewIATEntryDetail()?",
		}
	}
	if ed.AddendaRecords == 0 {
		return &FieldError{
			FieldName: "AddendaRecords",
			Value:     strconv.Itoa(ed.AddendaRecords),
			Msg:       msgFieldInclusion + ", did you use NewIATEntryDetail()?",
		}
	}
	if ed.DFIAccountNumber == "" {
		return &FieldError{
			FieldName: "DFIAccountNumber",
			Value:     ed.DFIAccountNumber,
			Msg:       msgFieldInclusion + ", did you use NewIATEntryDetail()?",
		}
	}
	if ed.AddendaRecordIndicator == 0 {
		return &FieldError{
			FieldName: "AddendaRecordIndicator",
			Value:     strconv.Itoa(ed.AddendaRecordIndicator),
			Msg:       msgFieldInclusion + ", did you use NewIATEntryDetail()?",
		}
	}
	if ed.TraceNumber == "" {
		return &FieldError{
			FieldName: "TraceNumber",
			Value:     ed.TraceNumberField(),
			Msg:       msgFieldInclusion + ", did you use NewIATEntryDetail()?",
		}
	}
	return nil
}

// SetRDFI takes the 9 digit RDFI account number and separates it for RDFIIdentification and CheckDigit
func (ed *IATEntryDetail) SetRDFI(rdfi string) *IATEntryDetail {
	s := ed.stringField(rdfi, 9)
	ed.RDFIIdentification = ed.parseStringField(s[:8])
	ed.CheckDigit = ed.parseStringField(s[8:9])
	return ed
}

// SetTraceNumber takes first 8 digits of ODFI and concatenates a sequence number onto the TraceNumber
func (ed *IATEntryDetail) SetTraceNumber(ODFIIdentification string, seq int) {
	ed.TraceNumber = ed.stringField(ODFIIdentification, 8) + ed.numericField(seq, 7)
}

// RDFIIdentificationField get the rdfiIdentification with zero padding
func (ed *IATEntryDetail) RDFIIdentificationField() string {
	return ed.stringField(ed.RDFIIdentification, 8)
}

// AddendaRecordsField returns a zero padded AddendaRecords string
func (ed *IATEntryDetail) AddendaRecordsField() string {
	return ed.numericField(ed.AddendaRecords, 4)
}

func (ed *IATEntryDetail) reservedField() string {
	return ed.alphaField(ed.reserved, 13)
}

// AmountField returns a zero padded string of amount
func (ed *IATEntryDetail) AmountField() string {
	return ed.numericField(ed.Amount, 10)
}

// DFIAccountNumberField gets the DFIAccountNumber with space padding
func (ed *IATEntryDetail) DFIAccountNumberField() string {
	return ed.alphaField(ed.DFIAccountNumber, 35)
}

// reservedTwoField gets the reservedTwo
func (ed *IATEntryDetail) reservedTwoField() string {
	return ed.alphaField(ed.reservedTwo, 2)
}

// OFACSreeningIndicatorField gets the OFACSreeningIndicator
func (ed *IATEntryDetail) OFACSreeningIndicatorField() string {
	return ed.alphaField(ed.OFACSreeningIndicator, 1)
}

// SecondaryOFACSreeningIndicatorField gets the SecondaryOFACSreeningIndicator
func (ed *IATEntryDetail) SecondaryOFACSreeningIndicatorField() string {
	return ed.alphaField(ed.SecondaryOFACSreeningIndicator, 1)
}

// TraceNumberField returns a zero padded TraceNumber string
func (ed *IATEntryDetail) TraceNumberField() string {
	return ed.stringField(ed.TraceNumber, 15)
}

// AddAddenda17 appends an Addenda17 to the IATEntryDetail
func (ed *IATEntryDetail) AddAddenda17(addenda17 *Addenda17) {
	ed.Addenda17 = append(ed.Addenda17, addenda17)
}

// AddAddenda18 appends an Addenda18 to the IATEntryDetail
func (ed *IATEntryDetail) AddAddenda18(addenda18 *Addenda18) {
	ed.Addenda18 = append(ed.Addenda18, addenda18)
}
