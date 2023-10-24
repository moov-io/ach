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
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

// EntryDetail contains the actual transaction data for an individual entry.
// Fields include those designating the entry as a deposit (credit) or
// withdrawal (debit), the transit routing number for the entry recipient's financial
// institution, the account number (left justify,no zero fill), name, and dollar amount.
type EntryDetail struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// TransactionCode if the receivers account is checking, savings, general ledger (GL) or loan.
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
	// IdentificationNumber an internal identification (alphanumeric) that
	// you use to uniquely identify this Entry Detail Record
	IdentificationNumber string `json:"identificationNumber,omitempty"`
	// IndividualName The name of the receiver, usually the name on the bank account
	IndividualName string `json:"individualName"`
	// DiscretionaryData allows ODFIs to include codes, of significance only to them,
	// to enable specialized handling of the entry. There will be no
	// standardized interpretation for the value of this field. It can either
	// be a single two-character code, or two distinct one-character codes,
	// according to the needs of the ODFI and/or Originator involved. This
	// field must be returned intact for any returned entry.
	//
	// WEB and TEL batches use the Discretionary Data Field as the Payment Type Code
	DiscretionaryData string `json:"discretionaryData,omitempty"`
	// AddendaRecordIndicator indicates the existence of an Addenda Record.
	// A value of "1" indicates that one ore more addenda records follow,
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
	// Use TraceNumberField for a properly formatted string representation.
	TraceNumber string `json:"traceNumber,omitempty"`
	// Addenda02 for use with StandardEntryClassCode MTE, POS, and SHR
	Addenda02 *Addenda02 `json:"addenda02,omitempty"`
	// Addenda05 for use with StandardEntryClassCode: ACK, ATX, CCD, CIE, CTX, DNE, ENR, WEB, PPD, TRX.
	Addenda05 []*Addenda05 `json:"addenda05,omitempty"`
	// Addenda98 for user with Notification of Change
	Addenda98 *Addenda98 `json:"addenda98,omitempty"`
	// Addenda98 for user with Refused Notification of Change
	Addenda98Refused *Addenda98Refused `json:"addenda98Refused,omitempty"`
	// Addenda99 for use with Returns
	Addenda99 *Addenda99 `json:"addenda99,omitempty"`
	// Addenda99Contested for use with Contested Dishonored Returns
	Addenda99Contested *Addenda99Contested `json:"addenda99Contested,omitempty"`
	// Addenda99Dishonored for use with Dishonored Returns
	Addenda99Dishonored *Addenda99Dishonored `json:"addenda99Dishonored,omitempty"`
	// Category defines if the entry is a Forward, Return, or NOC
	Category string `json:"category,omitempty"`
	// validator is composed for data validation
	validator
	// converters is composed for ACH to golang Converters
	converters

	validateOpts *ValidateOpts
}

const (
	// CategoryForward defines the entry as being sent to the receiving institution
	CategoryForward = "Forward"
	// CategoryReturn defines the entry as being a return of a forward entry back to the originating institution
	CategoryReturn = "Return"
	// CategoryNOC defines the entry as being a notification of change of a forward entry to the originating institution
	CategoryNOC = "NOC"
	// CategoryDishonoredReturn defines the entry as being a dishonored return initiated by the ODFI to the RDFI that
	// submitted the return entry
	CategoryDishonoredReturn = "DishonoredReturn"
	// CategoryDishonoredReturnContested defines the entry as a contested dishonored return initiated by the RDFI to
	// the ODFI that submitted the dishonored return
	CategoryDishonoredReturnContested = "DishonoredReturnContested"

	// TransactionCode Values

	// CheckingCredit is a credit to the receiver's checking account
	CheckingCredit = 22
	// CheckingReturnNOCCredit is a return that credits the receiver's checking account
	CheckingReturnNOCCredit = 21
	// CheckingPrenoteCredit is a pre-notification of a credit to the receiver's checking account
	CheckingPrenoteCredit = 23
	// CheckingZeroDollarRemittanceCredit is a zero dollar remittance data credit to a checking account for CCD, CTX,
	// ACK, and ATX entries
	CheckingZeroDollarRemittanceCredit = 24
	// CheckingDebit is a debit to the receivers checking account
	CheckingDebit = 27
	// CheckingReturnNOCDebit is a return that debits the receiver's checking account
	CheckingReturnNOCDebit = 26
	// CheckingPrenoteDebit is a pre-notification of a debit to the receiver's checking account
	CheckingPrenoteDebit = 28
	// CheckingZeroDollarRemittanceDebit is a zero dollar remittance data debit to a checking account for CCD, CTX,
	// ACK, and ATX entries
	CheckingZeroDollarRemittanceDebit = 29

	// SavingsCredit is a credit to the receiver's savings account
	SavingsCredit = 32
	// SavingsReturnNOCCredit is a return that credits the receiver's savings account
	SavingsReturnNOCCredit = 31
	// SavingsPrenoteCredit is a pre-notification of a credit to the receiver's savings account
	SavingsPrenoteCredit = 33
	// SavingsZeroDollarRemittanceCredit is a zero dollar remittance data credit to a savings account for CCD
	// and CTX entries
	SavingsZeroDollarRemittanceCredit = 34
	// SavingsDebit is a debit to the receivers savings account
	SavingsDebit = 37
	// SavingsReturnNOCDebit is a return that debits the receiver's savings account
	SavingsReturnNOCDebit = 36
	// SavingsPrenoteDebit is a pre-notification of a debit to the receiver's savings account
	SavingsPrenoteDebit = 38
	// SavingsZeroDollarRemittanceDebit is a zero dollar remittance data debit to a savings account for CCD
	// and CTX entries
	SavingsZeroDollarRemittanceDebit = 39

	// GLCredit is a credit to the receiver's general ledger (GL) account
	GLCredit = 42
	// GLReturnNOCCredit is a return that credits the receiver's general ledger (GL) account
	GLReturnNOCCredit = 41
	// GLPrenoteCredit is a pre-notification of a credit to the receiver's general ledger (GL) account
	GLPrenoteCredit = 43
	// GLZeroDollarRemittanceCredit is a zero dollar remittance data credit to the receiver's general ledger (GL) account
	GLZeroDollarRemittanceCredit = 44
	// GLDebit is a debit to the receiver's general ledger (GL) account
	GLDebit = 47
	// GLReturnNOCDebit is a return that debits the receiver's general ledger (GL) account
	GLReturnNOCDebit = 46
	// GLPrenoteDebit is a pre-notification of a debit to the receiver's general ledger (GL) account
	GLPrenoteDebit = 48
	// GLZeroDollarRemittanceDebit is a zero dollar remittance data debit to the receiver's general ledger (GL) account
	GLZeroDollarRemittanceDebit = 49

	// LoanCredit is a credit to the receiver's loan account
	LoanCredit = 52
	// LoanReturnNOCCredit is a return that credits the receiver's loan account
	LoanReturnNOCCredit = 51
	// LoanPrenoteCredit is a pre-notification of a credit to the receiver's loan account
	LoanPrenoteCredit = 53
	// LoanZeroDollarRemittanceCredit is a zero dollar remittance data credit to the receiver's loan account
	LoanZeroDollarRemittanceCredit = 54
	// LoanDebit is a debit (Reversal's Only) to the receiver's loan account
	LoanDebit = 55
	// LoanReturnNOCDebit is a return that debits the receiver's loan account
	LoanReturnNOCDebit = 56
	// LoanPrenoteDebit is N/A
	// LoanZeroDollarRemittanceDebit is N/A

	// End of TransactionCode Values
)

// NewEntryDetail returns a new EntryDetail with default values for non exported fields
func NewEntryDetail() *EntryDetail {
	var entry EntryDetail
	entry.Category = CategoryForward
	return &entry
}

// Parse takes the input record string and parses the EntryDetail values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (ed *EntryDetail) Parse(record string) {
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
			ed.TransactionCode = ed.parseNumField(reset())

		case 11:
			// 4-11 the RDFI's routing number without the last digit.
			ed.RDFIIdentification = reset()

		case 12:
			// 12-12 The last digit of the RDFI's routing number
			ed.CheckDigit = reset()

		case 29:
			// 13-29 The receiver's bank account number you are crediting/debiting
			ed.DFIAccountNumber = ed.parseStringFieldWithOpts(reset(), ed.validateOpts)

		case 39:
			// 30-39 Number of cents you are debiting/crediting this account
			ed.Amount = ed.parseNumField(reset())

		case 54:
			// 40-54 An internal identification (alphanumeric) that you use to uniquely identify this Entry Detail Record
			ed.IdentificationNumber = reset()

		case 76:
			// 55-76 The name of the receiver, usually the name on the bank account
			ed.IndividualName = reset()

		case 78:
			// 77-78 allows ODFIs to include codes of significance only to them, normally blank
			// For WEB and TEL batches this field is the PaymentType which is either R(reoccurring) or S(single)
			ed.DiscretionaryData = reset()

		case 79:
			// 79-79 1 if addenda exists 0 if it does not
			ed.AddendaRecordIndicator = ed.parseNumField(reset())

		case 94:
			// 80-94 An internal identification (numeric) that you use to uniquely identify
			// this Entry Detail Record This number should be unique to the transaction and will help identify the transaction in case of an inquiry
			ed.TraceNumber = reset() // capture end of record
		}
	}
}

// String writes the EntryDetail struct to a 94 character string.
func (ed *EntryDetail) String() string {
	buf := getBuffer()
	defer saveBuffer(buf)

	buf.WriteString(entryDetailPos)
	buf.WriteString(fmt.Sprintf("%v", ed.TransactionCode))
	buf.WriteString(ed.RDFIIdentificationField())
	buf.WriteString(ed.CheckDigit)
	buf.WriteString(ed.DFIAccountNumberField())
	buf.WriteString(ed.AmountField())
	buf.WriteString(ed.IdentificationNumberField())
	buf.WriteString(ed.IndividualNameField())
	buf.WriteString(ed.DiscretionaryDataField())
	buf.WriteString(fmt.Sprintf("%v", ed.AddendaRecordIndicator))
	buf.WriteString(ed.TraceNumberField())

	return buf.String()
}

// SetValidation stores ValidateOpts on the EntryDetail which are to be used to override
// the default NACHA validation rules.
func (ed *EntryDetail) SetValidation(opts *ValidateOpts) {
	if ed == nil {
		return
	}
	ed.validateOpts = opts
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (ed *EntryDetail) Validate() error {
	if err := ed.fieldInclusion(); err != nil {
		return err
	}
	if ed.validateOpts != nil && ed.validateOpts.CheckTransactionCode != nil {
		if err := ed.validateOpts.CheckTransactionCode(ed.TransactionCode); err != nil {
			return fieldError("TransactionCode", err, strconv.Itoa(ed.TransactionCode))
		}
	} else {
		if err := ed.isTransactionCode(ed.TransactionCode); err != nil {
			return fieldError("TransactionCode", err, strconv.Itoa(ed.TransactionCode))
		}
	}
	if err := ed.isAlphanumeric(ed.DFIAccountNumber); err != nil {
		return fieldError("DFIAccountNumber", err, ed.DFIAccountNumber)
	}
	if ed.Amount < 0 {
		return fieldError("Amount", ErrNegativeAmount, ed.Amount)
	}
	if err := ed.amountOverflowsField(); err != nil {
		return fieldError("Amount", err, ed.Amount)
	}
	if err := ed.isAlphanumeric(ed.IdentificationNumber); err != nil {
		return fieldError("IdentificationNumber", err, ed.IdentificationNumber)
	}
	if err := ed.isAlphanumeric(ed.IndividualName); err != nil {
		return fieldError("IndividualName", err, ed.IndividualName)
	}
	if err := ed.isAlphanumeric(ed.DiscretionaryData); err != nil {
		return fieldError("DiscretionaryData", err, ed.DiscretionaryData)
	}

	if ed.validateOpts == nil || !ed.validateOpts.AllowInvalidCheckDigit {
		calculated := CalculateCheckDigit(ed.RDFIIdentificationField())

		edCheckDigit, err := strconv.Atoi(ed.CheckDigit)
		if err != nil {
			return fieldError("CheckDigit", err, ed.CheckDigit)
		}

		if calculated != edCheckDigit {
			return fieldError("RDFIIdentification", NewErrValidCheckDigit(calculated), ed.CheckDigit)
		}
	}

	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (ed *EntryDetail) fieldInclusion() error {
	if ed.TransactionCode == 0 {
		return fieldError("TransactionCode", ErrConstructor, strconv.Itoa(ed.TransactionCode))
	}
	if ed.RDFIIdentification == "" {
		return fieldError("RDFIIdentification", ErrConstructor, ed.RDFIIdentificationField())
	}
	if ed.DFIAccountNumber == "" {
		return fieldError("DFIAccountNumber", ErrConstructor, ed.DFIAccountNumber)
	}
	if ed.IndividualName == "" {
		return fieldError("IndividualName", ErrConstructor, ed.IndividualName)
	}
	if ed.TraceNumber == "" {
		return fieldError("TraceNumber", ErrConstructor, ed.TraceNumberField())
	}
	return nil
}

func (ed *EntryDetail) amountOverflowsField() error {
	intstr := strconv.Itoa(ed.Amount)
	strstr := ed.AmountField()
	if intstr == "0" && strstr == "0000000000" {
		return nil // both are empty values
	}
	if len(intstr) > len(strstr) {
		return fmt.Errorf("does not match formatted value %s", strstr)
	}
	return nil
}

// SetRDFI takes the 9 digit RDFI account number and separates it for RDFIIdentification and CheckDigit
func (ed *EntryDetail) SetRDFI(rdfi string) *EntryDetail {
	s := ed.stringField(rdfi, 9)
	ed.RDFIIdentification = ed.parseStringField(s[:8])
	ed.CheckDigit = ed.parseStringField(s[8:9])
	return ed
}

// SetTraceNumber takes first 8 digits of ODFI and concatenates a sequence number onto the TraceNumber
func (ed *EntryDetail) SetTraceNumber(ODFIIdentification string, seq int) {
	traceNumber := ed.stringField(ODFIIdentification, 8) + ed.numericField(seq, 7)
	ed.TraceNumber = traceNumber

	// Populate TraceNumber of addenda records that should match the Entry's trace number
	if ed.Addenda02 != nil {
		ed.Addenda02.TraceNumber = traceNumber
	}
	if ed.Addenda98 != nil {
		ed.Addenda98.TraceNumber = traceNumber
	}
	if ed.Addenda98Refused != nil {
		ed.Addenda98Refused.TraceNumber = traceNumber
	}
	if ed.Addenda99 != nil {
		ed.Addenda99.TraceNumber = traceNumber
	}
	if ed.Addenda99Contested != nil {
		ed.Addenda99Contested.TraceNumber = traceNumber
	}
	if ed.Addenda99Dishonored != nil {
		ed.Addenda99Dishonored.TraceNumber = traceNumber
	}
}

// RDFIIdentificationField get the rdfiIdentification with zero padding
func (ed *EntryDetail) RDFIIdentificationField() string {
	return ed.stringField(ed.RDFIIdentification, 8)
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

// CheckSerialNumberField is used in RCK, ARC, BOC files but returns
// a space padded string of the underlying IdentificationNumber field
func (ed *EntryDetail) CheckSerialNumberField() string {
	return ed.alphaField(ed.IdentificationNumber, 15)
}

// SetCheckSerialNumber setter for RCK, ARC, BOC CheckSerialNumber
// which is underlying IdentificationNumber
func (ed *EntryDetail) SetCheckSerialNumber(s string) {
	ed.IdentificationNumber = s
}

// SetPOPCheckSerialNumber setter for POP CheckSerialNumber
// which is characters 1-9 of underlying CheckSerialNumber \ IdentificationNumber
func (ed *EntryDetail) SetPOPCheckSerialNumber(s string) {
	ed.IdentificationNumber = ed.alphaField(s, 9)
}

// SetPOPTerminalCity setter for POP Terminal City
// which is characters 10-13 of underlying CheckSerialNumber \ IdentificationNumber
func (ed *EntryDetail) SetPOPTerminalCity(s string) {
	ed.IdentificationNumber = ed.IdentificationNumber + ed.alphaField(s, 4)
}

// SetPOPTerminalState setter for POP Terminal State
// which is characters 14-15 of underlying CheckSerialNumber \ IdentificationNumber
func (ed *EntryDetail) SetPOPTerminalState(s string) {
	ed.IdentificationNumber = ed.IdentificationNumber + ed.alphaField(s, 2)
}

// POPCheckSerialNumberField is used in POP, characters 1-9 of underlying BatchPOP
// CheckSerialNumber / IdentificationNumber
func (ed *EntryDetail) POPCheckSerialNumberField() string {
	return ed.parseStringField(ed.IdentificationNumber[0:9])
}

// POPTerminalCityField is used in POP, characters 10-13 of underlying BatchPOP
// CheckSerialNumber / IdentificationNumber
func (ed *EntryDetail) POPTerminalCityField() string {
	return ed.parseStringField(ed.IdentificationNumber[9:13])
}

// POPTerminalStateField is used in POP, characters 14-15 of underlying BatchPOP
// CheckSerialNumber / IdentificationNumber
func (ed *EntryDetail) POPTerminalStateField() string {
	return ed.parseStringField(ed.IdentificationNumber[13:15])
}

// SetSHRCardExpirationDate format MMYY is used in SHR, characters 1-4 of underlying
// IdentificationNumber
func (ed *EntryDetail) SetSHRCardExpirationDate(s string) {
	ed.IdentificationNumber = ed.alphaField(s, 4)
}

// SetSHRDocumentReferenceNumber format int is used in SHR, characters 5-15 of underlying
// IdentificationNumber
func (ed *EntryDetail) SetSHRDocumentReferenceNumber(s string) {
	ed.IdentificationNumber = ed.IdentificationNumber + ed.stringField(s, 11)
}

// SetSHRIndividualCardAccountNumber format int is used in SHR, underlying
// IndividualName
func (ed *EntryDetail) SetSHRIndividualCardAccountNumber(s string) {
	ed.IndividualName = ed.stringField(s, 22)
}

// SHRCardExpirationDateField format MMYY is used in SHR, characters 1-4 of underlying
// IdentificationNumber
func (ed *EntryDetail) SHRCardExpirationDateField() string {
	return ed.alphaField(ed.parseStringField(ed.IdentificationNumber[0:4]), 4)
}

// SHRDocumentReferenceNumberField format int is used in SHR, characters 5-15 of underlying
// IdentificationNumber
func (ed *EntryDetail) SHRDocumentReferenceNumberField() string {
	return ed.stringField(ed.IdentificationNumber[4:15], 11)
}

// SHRIndividualCardAccountNumberField format int is used in SHR, underlying
// IndividualName
func (ed *EntryDetail) SHRIndividualCardAccountNumberField() string {
	return ed.stringField(ed.IndividualName, 22)
}

// IndividualNameField returns a space padded string of IndividualName
func (ed *EntryDetail) IndividualNameField() string {
	return ed.alphaField(ed.IndividualName, 22)
}

// ReceivingCompanyField is used in CCD files but returns the underlying IndividualName field
func (ed *EntryDetail) ReceivingCompanyField() string {
	return ed.IndividualNameField()
}

// SetReceivingCompany setter for CCD ReceivingCompany which is underlying IndividualName
func (ed *EntryDetail) SetReceivingCompany(s string) {
	ed.IndividualName = s
}

// OriginalTraceNumberField is used in ACK and ATX files but returns the underlying IdentificationNumber field
func (ed *EntryDetail) OriginalTraceNumberField() string {
	return ed.IdentificationNumberField()
}

// SetOriginalTraceNumber setter for ACK and ATX OriginalTraceNumber which is underlying IdentificationNumber
func (ed *EntryDetail) SetOriginalTraceNumber(s string) {
	ed.IdentificationNumber = s
}

// SetCATXAddendaRecords setter for CTX and ATX AddendaRecords characters 1-4 of underlying IndividualName
func (ed *EntryDetail) SetCATXAddendaRecords(i int) {
	count := ed.numericField(i, 4)
	current := ed.IndividualName
	if utf8.RuneCountInString(current) > 4 {
		ed.IndividualName = count + current[4:]
	} else {
		ed.IndividualName = count + ed.alphaField(" ", 16) + "  "
	}
}

// SetCATXReceivingCompany setter for CTX and ATX ReceivingCompany characters 5-20 underlying IndividualName
// Position 21-22 of underlying Individual Name are reserved blank space for CTX "  "
func (ed *EntryDetail) SetCATXReceivingCompany(s string) {
	current := ed.IndividualName
	if utf8.RuneCountInString(current) > 4 {
		count := current[:4]
		ed.IndividualName = count + ed.alphaField(s, 16) + "  "
	} else {
		ed.IndividualName = "0000" + ed.alphaField(s, 16) + "  "
	}
}

// CATXAddendaRecordsField is used in CTX and ATX files, characters 1-4 of underlying IndividualName field
func (ed *EntryDetail) CATXAddendaRecordsField() string {
	if utf8.RuneCountInString(ed.IndividualName) < 5 {
		return ed.IndividualName
	}
	return ed.parseStringField(ed.IndividualName[:4])
}

// CATXReceivingCompanyField is used in CTX and ATX files, characters 5-20 of underlying IndividualName field
func (ed *EntryDetail) CATXReceivingCompanyField() string {
	if utf8.RuneCountInString(ed.IndividualName) < 4 {
		return ""
	}
	return ed.IndividualName[4:]
}

// CATXReservedField is used in CTX and ATX files, characters 21-22 of underlying IndividualName field
func (ed *EntryDetail) CATXReservedField() string {
	return ed.IndividualName[20:22]
}

// DiscretionaryDataField returns a space padded string of DiscretionaryData
func (ed *EntryDetail) DiscretionaryDataField() string {
	return ed.alphaField(ed.DiscretionaryData, 2)
}

// PaymentTypeField returns the DiscretionaryData field used in WEB and TEL batch files
func (ed *EntryDetail) PaymentTypeField() string {
	// because DiscretionaryData can be changed outside of PaymentType we reset the value for safety
	ed.SetPaymentType(ed.DiscretionaryData)
	return ed.DiscretionaryData
}

// SetPaymentType as R (Recurring) all other values will result in S (single).
// This is used for WEB and TEL batch files in-place of DiscretionaryData.
func (ed *EntryDetail) SetPaymentType(t string) {
	t = strings.ToUpper(strings.TrimSpace(t))
	if t == "R" {
		ed.DiscretionaryData = "R"
	} else {
		ed.DiscretionaryData = "S"
	}
}

// SetProcessControlField setter for TRC Process Control Field characters 1-6 of underlying IndividualName
func (ed *EntryDetail) SetProcessControlField(s string) {
	ed.IndividualName = ed.alphaField(s, 6)
}

// SetItemResearchNumber setter for TRC Item Research Number characters 7-22 of underlying IndividualName
func (ed *EntryDetail) SetItemResearchNumber(s string) {
	ed.IndividualName = ed.IndividualName + ed.alphaField(s, 16)
}

// SetItemTypeIndicator setter for TRC Item Type Indicator which is underlying Discretionary Data
func (ed *EntryDetail) SetItemTypeIndicator(s string) {
	ed.DiscretionaryData = ed.alphaField(s, 2)
}

// ProcessControlField getter for TRC Process Control Field characters 1-6 of underlying IndividualName
func (ed *EntryDetail) ProcessControlField() string {
	return ed.parseStringField(ed.IndividualName[0:6])
}

// ItemResearchNumber getter for TRC Item Research Number characters 7-22 of underlying IndividualName
func (ed *EntryDetail) ItemResearchNumber() string {
	return ed.parseStringField(ed.IndividualName[6:22])
}

// ItemTypeIndicator getter for TRC Item Type Indicator which is underlying Discretionary Data
func (ed *EntryDetail) ItemTypeIndicator() string {
	return ed.DiscretionaryData
}

// TraceNumberField returns a zero padded TraceNumber string
func (ed *EntryDetail) TraceNumberField() string {
	return ed.stringField(ed.TraceNumber, 15)
}

// CreditOrDebit returns a "C" for credit or "D" for debit based on the entry TransactionCode
func (ed *EntryDetail) CreditOrDebit() string {
	if ed.TransactionCode < 10 || ed.TransactionCode > 99 {
		return ""
	}
	tc := strconv.Itoa(ed.TransactionCode)

	// take the second number in the TransactionCode
	switch tc[1:2] {
	case "1", "2", "3", "4":
		return "C"
	case "5", "6", "7", "8", "9":
		return "D"
	default:
	}
	return ""
}

// AddAddenda05 appends an Addenda05 to the EntryDetail
func (ed *EntryDetail) AddAddenda05(addenda05 *Addenda05) {
	ed.Addenda05 = append(ed.Addenda05, addenda05)
}

// addendaCount returns the count of Addenda records added onto this EntryDetail
func (ed *EntryDetail) addendaCount() (n int) {
	if ed.Addenda02 != nil {
		n += 1
	}
	for i := range ed.Addenda05 {
		if ed.Addenda05[i] != nil {
			n += 1
		}
	}
	if ed.Addenda98 != nil {
		n += 1
	}
	if ed.Addenda98Refused != nil {
		n += 1
	}
	if ed.Addenda99 != nil {
		n += 1
	}
	if ed.Addenda99Dishonored != nil {
		n += 1
	}
	if ed.Addenda99Contested != nil {
		n += 1
	}
	return n
}

func sortEntriesByTraceNumber(entries []*EntryDetail) []*EntryDetail {
	sort.Slice(entries[:], func(i, j int) bool {
		return entries[i].TraceNumber < entries[j].TraceNumber
	})
	return entries
}
