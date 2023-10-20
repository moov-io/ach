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
	"time"
	"unicode/utf8"
)

// BatchHeader identifies the originating entity and the type of transactions
// contained in the batch (i.e., the standard entry class, PPD for consumer, CCD
// or CTX for corporate). This record also contains the effective date, or desired
// settlement date, for all entries contained in this batch. The settlement date
// field is not entered as it is determined by the ACH operator
type BatchHeader struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`

	// ServiceClassCode ACH Mixed Debits and Credits '200'
	// ACH Credits Only '220'
	// ACH Debits Only '225'
	ServiceClassCode int `json:"serviceClassCode"`

	// CompanyName the company originating the entries in the batch
	CompanyName string `json:"companyName"`

	// CompanyDiscretionaryData allows Originators and/or ODFIs to include codes (one or more),
	// of significance only to them, to enable specialized handling of all
	// subsequent entries in that batch. There will be no standardized
	// interpretation for the value of the field. This field must be returned
	// intact on any return entry.
	CompanyDiscretionaryData string `json:"companyDiscretionaryData,omitempty"`

	// CompanyIdentification The 9 digit FEIN number (proceeded by a predetermined
	// alpha or numeric character) of the entity in the company name field
	CompanyIdentification string `json:"companyIdentification"`

	// StandardEntryClassCode
	// Identifies the payment type (product) found within an ACH batch-using a 3-character code.
	// The SEC Code pertains to all items within batch.
	// Determines format of the detail records.
	// Determines addenda records (required or optional PLUS one or up to 9,999 records).
	// Determines rules to follow (return time frames).
	// Some SEC codes require specific data in predetermined fields within the ACH record
	StandardEntryClassCode string `json:"standardEntryClassCode"`

	// CompanyEntryDescription A description of the entries contained in the batch
	//
	//The Originator establishes the value of this field to provide a
	// description of the purpose of the entry to be displayed back to
	// the receive For example, "GAS BILL," "REG. SALARY," "INS. PREM,"
	// "SOC. SEC.," "DTC," "TRADE PAY," "PURCHASE," etc.
	//
	// This field must contain the word "REVERSAL" (left justified) when the
	// batch contains reversing entries.
	//
	// This field must contain the word "RECLAIM" (left justified) when the
	// batch contains reclamation entries.
	//
	// This field must contain the word "NONSETTLED" (left justified) when the
	// batch contains entries which could not settle.
	CompanyEntryDescription string `json:"companyEntryDescription,omitempty"`

	// CompanyDescriptiveDate currently, the Rules provide that the “Originator establishes this field as the date it
	// would like to see displayed to the Receiver for descriptive purposes.” NACHA recommends that, as desired,
	// the content of this field be formatted using the convention “SDHHMM”, where the “SD” in positions 64- 65 denotes
	// the intent for same-day settlement, and the hours and minutes in positions 66-69 denote the desired settlement
	// time using a 24-hour clock. When electing to use this convention, the ODFI would validate that the field
	// contains either.
	//
	// ODFIs at their discretion may require their Originators to further show intent for
	// same-day settlement using an optional, yet standardized, same-day indicator in the Company Descriptive Date
	// field. The Company Descriptive Date field (5 record, field 8) is an optional field with 6 positions available
	// (positions 64-69).
	CompanyDescriptiveDate string `json:"companyDescriptiveDate,omitempty"`

	// EffectiveEntryDate the date on which the entries are to settle. Format: YYMMDD (Y=Year, M=Month, D=Day)
	EffectiveEntryDate string `json:"effectiveEntryDate,omitempty"`

	// SettlementDate Leave blank, this field is inserted by the ACH operator
	SettlementDate string `json:"settlementDate,omitempty"`

	// OriginatorStatusCode refers to the ODFI initiating the Entry.
	// 0 ADV File prepared by an ACH Operator.
	// 1 This code identifies the Originator as a depository financial institution.
	// 2 This code identifies the Originator as a Federal Government entity or agency.
	OriginatorStatusCode int `json:"originatorStatusCode,omitempty"`

	//ODFIIdentification First 8 digits of the originating DFI transit routing number
	ODFIIdentification string `json:"ODFIIdentification"`

	// BatchNumber is assigned in ascending sequence to each batch by the ODFI
	// or its Sending Point in a given file of entries. Since the batch number
	// in the Batch Header Record and the Batch Control Record is the same,
	// the ascending sequence number should be assigned by batch and not by
	// record.
	BatchNumber int `json:"batchNumber"`

	// validator is composed for data validation
	validator

	// converters is composed for ACH to golang Converters
	converters

	validateOpts *ValidateOpts
}

const (
	// BatchHeader.ServiceClassCode and BatchControl.ServiceClassCode

	// MixedDebitsAndCredits indicates a batch can have debit and credit ACH entries
	MixedDebitsAndCredits = 200
	// CreditsOnly indicates a batch can only have credit ACH entries
	CreditsOnly = 220
	// DebitsOnly indicates a batch can only have debit ACH entries
	DebitsOnly = 225
	// AutomatedAccountingAdvices indicates a batch can only have Automated Accounting Advices (debit and credit)
	AutomatedAccountingAdvices = 280
)

// NewBatchHeader returns a new BatchHeader with default values for non exported fields
func NewBatchHeader() *BatchHeader {
	bh := &BatchHeader{
		OriginatorStatusCode: 1, // Prepared by a financial institution
		BatchNumber:          1,
	}
	return bh
}

// Parse takes the input record string and parses the BatchHeader values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (bh *BatchHeader) Parse(record string) {
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
		case 1:
			reset()
		case 4:
			// 2-4 MixedCreditsAnDebits (200), CreditsOnly (220), DebitsOnly (225)
			bh.ServiceClassCode = bh.parseNumField(reset())
		case 20:
			// 5-20 Your company's name. This name may appear on the receivers' statements prepared by the RDFI.
			bh.CompanyName = bh.parseStringFieldWithOpts(reset(), bh.validateOpts)
		case 40:
			// 21-40 Optional field you may use to describe the batch for internal accounting purposes
			bh.CompanyDiscretionaryData = bh.parseStringFieldWithOpts(reset(), bh.validateOpts)
		case 50:
			// 41-50 A 10-digit number assigned to you by the ODFI once they approve you to
			// originate ACH files through them. This is the same as the "Immediate origin" field in File Header Record
			bh.CompanyIdentification = bh.parseStringFieldWithOpts(reset(), bh.validateOpts)
		case 53:
			// 51-53 If the entries are PPD (credits/debits towards consumer account), use PPD.
			// If the entries are CCD (credits/debits towards corporate account), use CCD.
			// The difference between the 2 SEC codes are outside of the scope of this post.
			bh.StandardEntryClassCode = reset()
		case 63:
			// 54-63 Your description of the transaction. This text will appear on the receivers' bank statement.
			// For example: "Payroll   "
			bh.CompanyEntryDescription = bh.parseStringFieldWithOpts(reset(), bh.validateOpts)
		case 69:
			// 64-69 The date you choose to identify the transactions in YYMMDD format.
			// This date may be printed on the receivers' bank statement by the RDFI
			bh.CompanyDescriptiveDate = bh.parseStringFieldWithOpts(reset(), bh.validateOpts)
		case 75:
			// 70-75 Date transactions are to be posted to the receivers' account.
			// You almost always want the transaction to post as soon as possible, so put tomorrow's date in YYMMDD format
			bh.EffectiveEntryDate = bh.validateSimpleDate(reset())
		case 78:
			// 76-78 Always blank if creating batches (just fill with spaces).
			// Set to file value when parsing. Julian day format.
			bh.SettlementDate = bh.validateSettlementDate(reset())
		case 79:
			// 79-79 Always 1
			bh.OriginatorStatusCode = bh.parseNumField(reset())
		case 87:
			// 80-87 Your ODFI's routing number without the last digit. The last digit is simply a
			// checksum digit, which is why it is not necessary
			bh.ODFIIdentification = bh.parseStringFieldWithOpts(reset(), bh.validateOpts)
		case 94:
			// 88-94 Sequential number of this Batch Header Record
			// For example, put "1" if this is the first Batch Header Record in the file
			bh.BatchNumber = bh.parseNumField(reset())
		}
	}
}

// String writes the BatchHeader struct to a 94 character string.
func (bh *BatchHeader) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(batchHeaderPos)
	buf.WriteString(fmt.Sprintf("%v", bh.ServiceClassCode))
	buf.WriteString(bh.CompanyNameField())
	buf.WriteString(bh.CompanyDiscretionaryDataField())
	buf.WriteString(bh.CompanyIdentificationField())
	buf.WriteString(bh.StandardEntryClassCode)
	buf.WriteString(bh.CompanyEntryDescriptionField())
	buf.WriteString(bh.CompanyDescriptiveDateField())
	buf.WriteString(bh.EffectiveEntryDateField())
	buf.WriteString(bh.SettlementDateField())
	buf.WriteString(fmt.Sprintf("%v", bh.OriginatorStatusCode))
	buf.WriteString(bh.ODFIIdentificationField())
	buf.WriteString(bh.BatchNumberField())
	return buf.String()
}

// Equal returns true only if two BatchHeaders are equal.
// Equality is determined by the Nacha defined fields of each record.
func (bh *BatchHeader) Equal(other *BatchHeader) bool {
	if bh == nil || other == nil {
		return false
	}

	if bh.ServiceClassCode != other.ServiceClassCode {
		return false
	}
	if !strings.EqualFold(bh.CompanyName, other.CompanyName) {
		return false
	}
	if bh.CompanyIdentification != other.CompanyIdentification {
		return false
	}
	if bh.StandardEntryClassCode != other.StandardEntryClassCode {
		return false
	}
	if bh.CompanyEntryDescription != other.CompanyEntryDescription {
		return false
	}
	if bh.EffectiveEntryDate != other.EffectiveEntryDate {
		return false
	}
	if bh.ODFIIdentification != other.ODFIIdentification {
		return false
	}
	return true
}

// SetValidation stores ValidateOpts on the BatchHeader which are to be used to override
// the default NACHA validation rules.
func (bh *BatchHeader) SetValidation(opts *ValidateOpts) {
	if bh == nil {
		return
	}
	bh.validateOpts = opts
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (bh *BatchHeader) Validate() error {
	if err := bh.fieldInclusion(); err != nil {
		return err
	}
	if bh.validateOpts == nil || bh.validateOpts.CheckTransactionCode == nil {
		// Ensure the ServiceClassCode follows NACHA standards if we have no TransactionCode
		// validation overrides. Custom TransactionCode's don't allow for standard validation.
		if err := bh.isServiceClass(bh.ServiceClassCode); err != nil {
			return fieldError("ServiceClassCode", err, bh.ServiceClassCode)
		}
	}
	if err := bh.isSECCode(bh.StandardEntryClassCode); err != nil {
		return fieldError("StandardEntryClassCode", err, bh.StandardEntryClassCode)
	}
	if err := bh.isOriginatorStatusCode(bh.OriginatorStatusCode); err != nil {
		return fieldError("OriginatorStatusCode", err, bh.OriginatorStatusCode)
	}

	// Originator status code 0 is used for ADV batches only
	if bh.StandardEntryClassCode != ADV && bh.OriginatorStatusCode == 0 {
		return fieldError("OriginatorStatusCode", ErrOrigStatusCode, bh.OriginatorStatusCode)
	}

	if err := bh.isAlphanumeric(bh.CompanyName); err != nil {
		return fieldError("CompanyName", err, bh.CompanyName)
	}
	if err := bh.isAlphanumeric(bh.CompanyDiscretionaryData); err != nil {
		return fieldError("CompanyDiscretionaryData", err, bh.CompanyDiscretionaryData)
	}
	if err := bh.isAlphanumeric(bh.CompanyIdentification); err != nil {
		return fieldError("CompanyIdentification", err, bh.CompanyIdentification)
	}
	if err := bh.isAlphanumeric(bh.CompanyEntryDescription); err != nil {
		return fieldError("CompanyEntryDescription", err, bh.CompanyEntryDescription)
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (bh *BatchHeader) fieldInclusion() error {
	if bh.ServiceClassCode == 0 {
		return fieldError("ServiceClassCode", ErrConstructor, strconv.Itoa(bh.ServiceClassCode))
	}
	if bh.CompanyName == "" {
		return fieldError("CompanyName", ErrConstructor, bh.CompanyName)
	}
	if bh.CompanyIdentification == "" {
		return fieldError("CompanyIdentification", ErrConstructor, bh.CompanyIdentification)
	}
	if bh.StandardEntryClassCode == "" {
		return fieldError("StandardEntryClassCode", ErrConstructor, bh.StandardEntryClassCode)
	}
	if bh.CompanyEntryDescription == "" {
		return fieldError("CompanyEntryDescription", ErrConstructor, bh.CompanyEntryDescription)
	}
	if bh.ODFIIdentification == "" {
		return fieldError("ODFIIdentification", ErrConstructor, bh.ODFIIdentificationField())
	}
	return nil
}

// CompanyNameField get the CompanyName left padded
func (bh *BatchHeader) CompanyNameField() string {
	return bh.alphaField(bh.CompanyName, 16)
}

// CompanyDiscretionaryDataField get the CompanyDiscretionaryData left padded
func (bh *BatchHeader) CompanyDiscretionaryDataField() string {
	return bh.alphaField(bh.CompanyDiscretionaryData, 20)
}

// CompanyIdentificationField get the CompanyIdentification left padded
func (bh *BatchHeader) CompanyIdentificationField() string {
	return bh.alphaField(bh.CompanyIdentification, 10)
}

// CompanyEntryDescriptionField get the CompanyEntryDescription left padded
func (bh *BatchHeader) CompanyEntryDescriptionField() string {
	return bh.alphaField(bh.CompanyEntryDescription, 10)
}

// CompanyDescriptiveDateField get the CompanyDescriptiveDate left padded
func (bh *BatchHeader) CompanyDescriptiveDateField() string {
	return bh.alphaField(bh.CompanyDescriptiveDate, 6)
}

// EffectiveEntryDateField get the EffectiveEntryDate in YYMMDD format
func (bh *BatchHeader) EffectiveEntryDateField() string {
	// ENR records require EffectiveEntryDate to be space filled. NACHA Page OR108
	if bh.CompanyEntryDescription == "AUTOENROLL" {
		return bh.alphaField("", 6)
	}
	return bh.stringField(bh.EffectiveEntryDate, 6) // YYMMDD
}

// ODFIIdentificationField get the odfi number zero padded
func (bh *BatchHeader) ODFIIdentificationField() string {
	return bh.stringField(bh.ODFIIdentification, 8)
}

// BatchNumberField get the batch number zero padded
func (bh *BatchHeader) BatchNumberField() string {
	return bh.numericField(bh.BatchNumber, 7)
}

func (bh *BatchHeader) SettlementDateField() string {
	return bh.alphaField(bh.SettlementDate, 3)
}

func (bh *BatchHeader) LiftEffectiveEntryDate() (time.Time, error) {
	return time.Parse("060102", bh.EffectiveEntryDate) // YYMMDD
}
