// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strconv"
)

// EntryDetail contains the actual transaction data for an individual entry.
// Fields include those designating the entry as a deposit (credit) or
// withdrawal (debit), the transit routing number for the entry recipient’s financial
// institution, the account number (left justify,no zero fill), name, and dollar amount.
type EntryDetail struct {
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
	TransactionCode int

	// rdfiIdentification is the RDFI's routing number without the last digit.
	// Receiving Depository Financial Institution
	RDFIIdentification int

	// CheckDigit the last digit of the RDFI's routing number
	CheckDigit int

	// DFIAccountNumber is the receiver's bank account number you are crediting/debiting.
	// It important to note that this is an alphanumeric field, so its space padded, no zero padded
	DFIAccountNumber string

	// Amount Number of cents you are debiting/crediting this account
	Amount int

	// IdentificationNumber n internal identification (alphanumeric) that
	// you use to uniquely identify this Entry Detail Record
	IdentificationNumber string

	// IndividualName The name of the receiver, usually the name on the bank account
	IndividualName string

	// DiscretionaryData allows ODFIs to include codes, of significance only to them,
	// to enable specialized handling of the entry. There will be no
	// standardized interpretation for the value of this field. It can either
	// be a single two-character code, or two distinct one-character codes,
	// according to the needs of the ODFI and/or Originator involved. This
	// field must be returned intact for any returned entry.
	//
	// WEB uses the Discretionary Data Field as the Payment Type Code
	DiscretionaryData string

	// AddendaRecordIndicator indicates the existence of an Addenda Record.
	// A value of "1" indicates that one ore more addenda records follow,
	// and "0" means no such record is present.
	AddendaRecordIndicator int

	// TraceNumber assigned by the ODFI in ascending sequence, is included in each
	// Entry Detail Record, Corporate Entry Detail Record, and addenda Record.
	// Trace Numbers uniquely identify each entry within a batch in an ACH input file.
	// In association with the Batch Number, transmission (File Creation) Date,
	// and File ID Modifier, the Trace Number uniquely identifies an entry within a given file.
	// For addenda Records, the Trace Number will be identical to the Trace Number
	// in the associated Entry Detail Record, since the Trace Number is associated
	// with an entry or item rather than a physical record.
	TraceNumber int

	// Addendum a list of Addenda for the Entry
	// keeping separarte lists for different types of addenda...
	Addendum       []Addenda
	ReturnAddendum []ReturnAddenda
	// validator is composed for data validation
	validator
	// converters is composed for ACH to golang Converters
	converters
}

// EntryParam is the minimal fields required to make a ach entry
type EntryParam struct {
	ReceivingDFI      string `json:"receiving_dfi"`
	RDFIAccount       string `json:"rdfi_account"`
	Amount            string `json:"amount"`
	IDNumber          string `json:"id_number,omitempty"`
	IndividualName    string `json:"individual_name,omitempty"`
	ReceivingCompany  string `json:"receiving_company,omitempty"`
	DiscretionaryData string `json:"discretionary_data,omitempty"`
	TransactionCode   string `json:"transaction_code"`
}

// NewEntryDetail returns a new EntryDetail with default values for none exported fields
func NewEntryDetail(params ...EntryParam) *EntryDetail {
	entry := &EntryDetail{
		recordType: "6",
	}
	if len(params) > 0 {
		entry.SetRDFI(entry.parseNumField(params[0].ReceivingDFI))
		entry.DFIAccountNumber = params[0].RDFIAccount
		entry.Amount = entry.parseNumField(params[0].Amount)
		entry.IdentificationNumber = params[0].IDNumber
		if params[0].IndividualName != "" {
			entry.IndividualName = params[0].IndividualName
		} else {
			entry.IndividualName = params[0].ReceivingCompany
		}
		entry.DiscretionaryData = params[0].DiscretionaryData
		entry.TransactionCode = entry.parseNumField(params[0].TransactionCode)

		entry.setTraceNumber(entry.RDFIIdentification, 1)
		return entry
	}
	return entry
}

// Parse takes the input record string and parses the EntryDetail values
func (ed *EntryDetail) Parse(record string) {
	// 1-1 Always "6"
	ed.recordType = "6"
	// 2-3 is checking credit 22 debit 27 savings credit 32 debit 37
	ed.TransactionCode = ed.parseNumField(record[1:3])
	// 4-11 the RDFI's routing number without the last digit.
	ed.RDFIIdentification = ed.parseNumField(record[3:11])
	// 12-12 The last digit of the RDFI's routing number
	ed.CheckDigit = ed.parseNumField(record[11:12])
	// 13-29 The receiver's bank account number you are crediting/debiting
	ed.DFIAccountNumber = record[12:29]
	// 30-39 Number of cents you are debiting/crediting this account
	ed.Amount = ed.parseNumField(record[29:39])
	// 40-54 An internal identification (alphanumeric) that you use to uniquely identify this Entry Detail Record
	ed.IdentificationNumber = record[39:54]
	// 55-76 The name of the receiver, usually the name on the bank account
	ed.IndividualName = record[54:76]
	// 77-78 allows ODFIs to include codes of significance only to them
	// normally blank
	ed.DiscretionaryData = record[76:78]
	// 79-79 1 if addenda exists 0 if it does not
	ed.AddendaRecordIndicator = ed.parseNumField(record[78:79])
	// 80-84 An internal identification (alphanumeric) that you use to uniquely identify
	// this Entry Detail Recor This number should be unique to the transaction and will help identify the transaction in case of an inquiry
	ed.TraceNumber = ed.parseNumField(record[79:94])
}

// String writes the EntryDetail struct to a 94 character string.
func (ed *EntryDetail) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v",
		ed.recordType,
		ed.TransactionCode,
		ed.RDFIIdentificationField(),
		ed.CheckDigit,
		ed.DFIAccountNumberField(),
		ed.AmountField(),
		ed.IdentificationNumberField(),
		ed.IndividualNameField(),
		ed.DiscretionaryDataField(),
		ed.AddendaRecordIndicator,
		ed.TraceNumberField())
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ed *EntryDetail) Validate() error {
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
	if err := ed.isAlphanumeric(ed.IdentificationNumber); err != nil {
		return &FieldError{FieldName: "IdentificationNumber", Value: ed.IdentificationNumber, Msg: err.Error()}
	}
	if err := ed.isAlphanumeric(ed.IndividualName); err != nil {
		return &FieldError{FieldName: "IndividualName", Value: ed.IndividualName, Msg: err.Error()}
	}
	if err := ed.isAlphanumeric(ed.DiscretionaryData); err != nil {
		return &FieldError{FieldName: "DiscretionaryData", Value: ed.DiscretionaryData, Msg: err.Error()}
	}
	calculated := ed.CalculateCheckDigit(ed.RDFIIdentificationField())
	if calculated != ed.CheckDigit {
		msg := fmt.Sprintf(msgValidCheckDigit, calculated)
		return &FieldError{FieldName: "RDFIIdentification", Value: strconv.Itoa(ed.CheckDigit), Msg: msg}
	}

	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (ed *EntryDetail) fieldInclusion() error {
	if ed.recordType == "" {
		return &FieldError{FieldName: "recordType", Value: ed.recordType, Msg: msgFieldInclusion}
	}
	if ed.TransactionCode == 0 {
		return &FieldError{FieldName: "TransactionCode", Value: strconv.Itoa(ed.TransactionCode), Msg: msgFieldInclusion}
	}
	if ed.RDFIIdentification == 0 {
		return &FieldError{FieldName: "RDFIIdentification", Value: ed.RDFIIdentificationField(), Msg: msgFieldInclusion}
	}
	if ed.DFIAccountNumber == "" {
		return &FieldError{FieldName: "DFIAccountNumber", Value: ed.DFIAccountNumber, Msg: msgFieldInclusion}
	}
	// amount can be 0 if it's COR, should probably be more specific...
	/*
		if ed.Amount == 0 {
			return &FieldError{FieldName: "Amount", Value: ed.AmountField(), Msg: msgFieldInclusion}
		}*/
	if ed.IndividualName == "" {
		return &FieldError{FieldName: "IndividualName", Value: ed.IndividualName, Msg: msgFieldInclusion}
	}
	if ed.TraceNumber == 0 {
		return &FieldError{FieldName: "TraceNumber", Value: ed.TraceNumberField(), Msg: msgFieldInclusion}
	}
	return nil
}

// AddAddenda appends an EntryDetail to the Addendum
func (ed *EntryDetail) AddAddenda(addenda Addenda) []Addenda {
	ed.AddendaRecordIndicator = 1
	// checks to make sure that we only have either or, not both
	if ed.ReturnAddendum != nil {
		return nil
	}
	ed.Addendum = append(ed.Addendum, addenda)
	return ed.Addendum
}

// AddReturnAddenda appends an ReturnAddendum to the entry
func (ed *EntryDetail) AddReturnAddenda(returnAddendum ReturnAddenda) []ReturnAddenda {
	ed.AddendaRecordIndicator = 1
	// checks to make sure that we only have either or, not both
	if ed.Addendum != nil {
		return nil
	}
	ed.ReturnAddendum = append(ed.ReturnAddendum, returnAddendum)
	return ed.ReturnAddendum
}

// SetRDFI takes the 9 digit RDFI account number and separates it for RDFIIdentification and CheckDigit
func (ed *EntryDetail) SetRDFI(rdfi int) *EntryDetail {
	s := ed.numericField(rdfi, 9)
	ed.RDFIIdentification = ed.parseNumField(s[:8])
	ed.CheckDigit = ed.parseNumField(s[8:9])
	return ed
}

// setTraceNumber takes first 8 digits of RDFI and concatenates a sequence number onto the TraceNumber
func (ed *EntryDetail) setTraceNumber(RDFIIdentification int, seq int) {
	trace := ed.numericField(RDFIIdentification, 8) + ed.numericField(seq, 7)
	ed.TraceNumber = ed.parseNumField(trace)
}

// RDFIIdentificationField get the rdfiIdentification with zero padding
func (ed *EntryDetail) RDFIIdentificationField() string {
	return ed.numericField(ed.RDFIIdentification, 8)
}

// DFIAccountNumberField gets the DFIAccountNumber with space padding
func (ed *EntryDetail) DFIAccountNumberField() string {
	return ed.alphaField(ed.DFIAccountNumber, 17)
}

// AmountField returns a zero padded string of amount
func (ed *EntryDetail) AmountField() string {
	return ed.numericField(ed.Amount, 10)
}

// IdentificationNumberField returns a space padded string of IdentificationNumber
func (ed *EntryDetail) IdentificationNumberField() string {
	return ed.alphaField(ed.IdentificationNumber, 15)
}

// IndividualNameField returns a space padded string of IndividualName
func (ed *EntryDetail) IndividualNameField() string {
	return ed.alphaField(ed.IndividualName, 22)
}

// ReceivingCompanyField is used in CCD files but returns the underlying IndividualName field
func (ed *EntryDetail) ReceivingCompanyField() string {
	return ed.IndividualNameField()
}

// DiscretionaryDataField returns a space padded string of DiscretionaryData
func (ed *EntryDetail) DiscretionaryDataField() string {
	return ed.alphaField(ed.DiscretionaryData, 2)
}

// PaymentType returns the discretionary data field used in web batch files
func (ed *EntryDetail) PaymentType() string {
	if ed.DiscretionaryData == "" {
		ed.DiscretionaryData = "S"
	}
	return ed.DiscretionaryDataField()
}

// TraceNumberField returns a zero padded traceNumber string
func (ed *EntryDetail) TraceNumberField() string {
	return ed.numericField(ed.TraceNumber, 15)
}

// HasReturnAddenda returns true if entry has return addenda
func (ed *EntryDetail) HasReturnAddenda() bool {
	return ed.ReturnAddendum != nil
}
