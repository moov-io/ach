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
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// IATEntryDetail contains the actual transaction data for an individual entry.
// Fields include those designating the entry as a deposit (credit) or
// withdrawal (debit), the transit routing number for the entry recipient's financial
// institution, the account number (left justify,no zero fill), name, and dollar amount.
type IATEntryDetail struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// TransactionCode if the receivers account is:
	// Credit (deposit) to checking account '22'
	// Prenote for credit to checking account '23'
	// Debit (withdrawal) to checking account '27'
	// Prenote for debit to checking account '28'
	// Credit to savings account '32'
	// Prenote for credit to savings account '33'
	// Debit to savings account '37'
	// Prenote for debit to savings account '38'
	TransactionCode int `json:"transactionCode"`
	// RDFIIdentification is the RDFI's routing number without the last digit.
	// Receiving Depository Financial Institution
	RDFIIdentification string `json:"RDFIIdentification"`
	// CheckDigit the last digit of the RDFI's routing number
	CheckDigit string `json:"checkDigit"`
	// AddendaRecords is the number of Addenda Records
	AddendaRecords int `json:"addendaRecords"`
	// Amount Number of cents you are debiting/crediting this account
	Amount int `json:"amount"`
	// DFIAccountNumber is the receiver's bank account number you are crediting/debiting.
	// It important to note that this is an alphanumeric field, so its space padded, no zero padded
	DFIAccountNumber string `json:"DFIAccountNumber"`
	// OFACScreeningIndicator - Leave blank
	OFACScreeningIndicator string `json:"OFACScreeningIndicator"`
	// SecondaryOFACScreeningIndicator - Leave blank
	SecondaryOFACScreeningIndicator string `json:"secondaryOFACScreeningIndicator"`
	// AddendaRecordIndicator indicates the existence of an Addenda Record.
	// A value of "1" indicates that one or more addenda records follow,
	// and "0" means no such record is present.
	AddendaRecordIndicator int `json:"addendaRecordIndicator"`
	// TraceNumber assigned by the ODFI in ascending sequence, is included in each
	// Entry Detail Record, Corporate Entry Detail Record, and addenda Record.
	// Trace Numbers uniquely identify each entry within a batch in an ACH input file.
	// In association with the Batch Number, transmission (File Creation) Date,
	// and File ID Modifier, the Trace Number uniquely identifies an entry within a given file.
	// For addenda Records, the Trace Number will be identical to the Trace Number
	// in the associated Entry Detail Record, since the Trace Number is associated
	// with an entry or item rather than a physical record.
	//
	// Use TraceNumberField for a properly formatted string representation.
	TraceNumber string `json:"traceNumber,omitempty"`
	// Addenda10 is mandatory for IAT entries
	//
	// The Addenda10 Record identifies the Receiver of the transaction and the dollar amount of
	// the payment.
	Addenda10 *Addenda10 `json:"addenda10"`
	// Addenda11 is mandatory for IAT entries
	//
	// The Addenda11 record identifies key information related to the Originator of
	// the entry.
	Addenda11 *Addenda11 `json:"addenda11"`
	// Addenda12 is mandatory for IAT entries
	//
	// The Addenda12 record identifies key information related to the Originator of
	// the entry.
	Addenda12 *Addenda12 `json:"addenda12"`
	// Addenda13 is mandatory for IAT entries
	//
	// The Addenda13 contains information related to the financial institution originating the entry.
	// For inbound IAT entries, the Fourth Addenda Record must contain information to identify the
	// foreign financial institution that is providing the funding and payment instruction for
	// the IAT entry.
	Addenda13 *Addenda13 `json:"addenda13"`
	// Addenda14 is mandatory for IAT entries
	//
	// The Addenda14 identifies the Receiving financial institution holding the Receiver's account.
	Addenda14 *Addenda14 `json:"addenda14"`
	// Addenda15 is mandatory for IAT entries
	//
	// The Addenda15 record identifies key information related to the Receiver.
	Addenda15 *Addenda15 `json:"addenda15"`
	// Addenda16 is mandatory for IAt entries
	//
	// Addenda16 record identifies additional key information related to the Receiver.
	Addenda16 *Addenda16 `json:"addenda16"`
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

	validateOpts *ValidateOpts
}

// NewIATEntryDetail returns a new IATEntryDetail with default values for non exported fields
func NewIATEntryDetail() *IATEntryDetail {
	iatEd := &IATEntryDetail{
		Category:               CategoryForward,
		AddendaRecordIndicator: 1,
	}
	return iatEd
}

// Parse takes the input record string and parses the EntryDetail values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (iatEd *IATEntryDetail) Parse(record string) {
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
			// do nothing, ignore "6" record type
			reset()

		case 3:
			// 2-3 is checking credit 22 debit 27 savings credit 32 debit 37
			iatEd.TransactionCode = iatEd.parseNumField(reset())
		case 11:
			// 4-11 the RDFI's routing number without the last digit.
			iatEd.RDFIIdentification = iatEd.parseStringField(reset())
		case 12:
			// 12-12 The last digit of the RDFI's routing number
			iatEd.CheckDigit = iatEd.parseStringField(reset())
		case 16:
			// 13-16 Number of addenda records
			iatEd.AddendaRecords = iatEd.parseNumField(reset())
		case 29:
			// 17-29 reserved - Leave blank
			reset()
		case 39:
			// 30-39 Number of cents you are debiting/crediting this account
			iatEd.Amount = iatEd.parseNumField(reset())
		case 74:
			// 40-74 The foreign receiver's account number you are crediting/debiting
			iatEd.DFIAccountNumber = string(reset())
		case 76:
			// 75-76 reserved Leave blank
			reset()
		case 77:
			// 77 OFACScreeningIndicator
			reset()
			iatEd.OFACScreeningIndicator = " "
		case 78:
			// 78-78 Secondary SecondaryOFACScreeningIndicator
			reset()
			iatEd.SecondaryOFACScreeningIndicator = " "
		case 79:
			// 79-79 1 if addenda exists 0 if it does not
			iatEd.AddendaRecordIndicator = iatEd.parseNumField(reset())
		case 94:
			// 80-94 An internal identification (alphanumeric) that you use to uniquely identify
			// this Entry Detail Record This number should be unique to the transaction and will
			// help identify the transaction in case of an inquiry
			iatEd.TraceNumber = strings.TrimSpace(reset())
		}
	}
}

// String writes the EntryDetail struct to a 94 character string.
func (iatEd *IATEntryDetail) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(entryDetailPos)
	buf.WriteString(fmt.Sprintf("%v", iatEd.TransactionCode))
	buf.WriteString(iatEd.RDFIIdentificationField())
	buf.WriteString(iatEd.CheckDigit)
	buf.WriteString(iatEd.AddendaRecordsField())
	buf.WriteString("             ")
	buf.WriteString(iatEd.AmountField())
	buf.WriteString(iatEd.DFIAccountNumberField())
	buf.WriteString("  ")
	buf.WriteString(iatEd.OFACScreeningIndicatorField())
	buf.WriteString(iatEd.SecondaryOFACScreeningIndicatorField())
	buf.WriteString(fmt.Sprintf("%v", iatEd.AddendaRecordIndicator))
	buf.WriteString(iatEd.TraceNumberField())
	return buf.String()
}

// SetValidation stores ValidateOpts on the EntryDetail which are to be used to override
// the default NACHA validation rules.
func (iatEd *IATEntryDetail) SetValidation(opts *ValidateOpts) {
	if iatEd == nil {
		return
	}
	iatEd.validateOpts = opts
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (iatEd *IATEntryDetail) Validate() error {
	if err := iatEd.fieldInclusion(); err != nil {
		return err
	}
	if iatEd.validateOpts != nil && iatEd.validateOpts.CheckTransactionCode != nil {
		if err := iatEd.validateOpts.CheckTransactionCode(iatEd.TransactionCode); err != nil {
			return fieldError("TransactionCode", err, strconv.Itoa(iatEd.TransactionCode))
		}
	} else {
		if err := iatEd.isTransactionCode(iatEd.TransactionCode); err != nil {
			return fieldError("TransactionCode", err, strconv.Itoa(iatEd.TransactionCode))
		}
	}
	if err := iatEd.isAlphanumeric(iatEd.DFIAccountNumber); err != nil {
		return fieldError("DFIAccountNumber", err, iatEd.DFIAccountNumber)
	}
	// CheckDigit calculations
	calculated := CalculateCheckDigit(iatEd.RDFIIdentificationField())

	edCheckDigit, err := strconv.Atoi(iatEd.CheckDigit)
	if err != nil {
		return fieldError("CheckDigit", err, iatEd.CheckDigit)
	}
	if calculated != edCheckDigit {
		return fieldError("RDFIIdentification", NewErrValidCheckDigit(calculated), iatEd.CheckDigit)
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (iatEd *IATEntryDetail) fieldInclusion() error {
	if iatEd.TransactionCode == 0 {
		return fieldError("TransactionCode", ErrConstructor, strconv.Itoa(iatEd.TransactionCode))
	}
	if iatEd.RDFIIdentification == "" {
		return fieldError("RDFIIdentification", ErrConstructor, iatEd.RDFIIdentificationField())
	}
	if iatEd.AddendaRecords == 0 {
		return fieldError("AddendaRecords", ErrConstructor, strconv.Itoa(iatEd.AddendaRecords))
	}
	if iatEd.DFIAccountNumber == "" {
		return fieldError("DFIAccountNumber", ErrConstructor, iatEd.DFIAccountNumber)
	}
	if iatEd.AddendaRecordIndicator == 0 {
		return fieldError("AddendaRecordIndicator", ErrConstructor, strconv.Itoa(iatEd.AddendaRecordIndicator))
	}
	if iatEd.TraceNumber == "" {
		return fieldError("TraceNumber", ErrConstructor, iatEd.TraceNumberField())
	}
	return nil
}

// SetRDFI takes the 9 digit RDFI account number and separates it for RDFIIdentification and CheckDigit
func (iatEd *IATEntryDetail) SetRDFI(rdfi string) *IATEntryDetail {
	s := iatEd.stringField(rdfi, 9)
	iatEd.RDFIIdentification = iatEd.parseStringField(s[:8])
	iatEd.CheckDigit = iatEd.parseStringField(s[8:9])
	return iatEd
}

// SetTraceNumber takes first 8 digits of ODFI and concatenates a sequence number onto the TraceNumber
func (iatEd *IATEntryDetail) SetTraceNumber(ODFIIdentification string, seq int) {
	iatEd.TraceNumber = iatEd.stringField(ODFIIdentification, 8) + iatEd.numericField(seq, 7)
}

// RDFIIdentificationField get the rdfiIdentification with zero padding
func (iatEd *IATEntryDetail) RDFIIdentificationField() string {
	return iatEd.stringField(iatEd.RDFIIdentification, 8)
}

// AddendaRecordsField returns a zero padded AddendaRecords string
func (iatEd *IATEntryDetail) AddendaRecordsField() string {
	return iatEd.numericField(iatEd.AddendaRecords, 4)
}

// AmountField returns a zero padded string of amount
func (iatEd *IATEntryDetail) AmountField() string {
	return iatEd.numericField(iatEd.Amount, 10)
}

// DFIAccountNumberField gets the DFIAccountNumber with space padding
func (iatEd *IATEntryDetail) DFIAccountNumberField() string {
	return iatEd.alphaField(iatEd.DFIAccountNumber, 35)
}

// OFACScreeningIndicatorField gets the OFACScreeningIndicator
func (iatEd *IATEntryDetail) OFACScreeningIndicatorField() string {
	return iatEd.alphaField(iatEd.OFACScreeningIndicator, 1)
}

// SecondaryOFACScreeningIndicatorField gets the SecondaryOFACScreeningIndicator
func (iatEd *IATEntryDetail) SecondaryOFACScreeningIndicatorField() string {
	return iatEd.alphaField(iatEd.SecondaryOFACScreeningIndicator, 1)
}

// TraceNumberField returns a zero padded TraceNumber string
func (iatEd *IATEntryDetail) TraceNumberField() string {
	return iatEd.stringField(iatEd.TraceNumber, 15)
}

// AddAddenda17 appends an Addenda17 to the IATEntryDetail
func (iatEd *IATEntryDetail) AddAddenda17(addenda17 *Addenda17) {
	iatEd.Addenda17 = append(iatEd.Addenda17, addenda17)
}

// AddAddenda18 appends an Addenda18 to the IATEntryDetail
func (iatEd *IATEntryDetail) AddAddenda18(addenda18 *Addenda18) {
	iatEd.Addenda18 = append(iatEd.Addenda18, addenda18)
}
