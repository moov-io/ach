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

// BatchControl contains entry counts, dollar total and has totals for all
// entries contained in the preceding batch
type BatchControl struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// ServiceClassCode ACH Mixed Debits and Credits '200'
	// ACH Credits Only '220'
	// ACH Debits Only '225'
	// Constants: MixedCreditsAnDebits (220), CReditsOnly 9220), DebitsOnly (225)
	// Same as 'ServiceClassCode' in BatchHeaderRecord
	ServiceClassCode int `json:"serviceClassCode"`
	// EntryAddendaCount is a tally of each Entry Detail Record and each Addenda
	// Record processed, within either the batch or file as appropriate.
	EntryAddendaCount int `json:"entryAddendaCount"`
	// validate the Receiving DFI Identification in each Entry Detail Record is hashed
	// to provide a check against inadvertent alteration of data contents due
	// to hardware failure or program error
	//
	// In this context the Entry Hash is the sum of the corresponding fields in the
	// Entry Detail Records on the file.
	EntryHash int `json:"entryHash"`
	// TotalDebitEntryDollarAmount Contains accumulated Entry debit totals within the batch.
	TotalDebitEntryDollarAmount int `json:"totalDebit"`
	// TotalCreditEntryDollarAmount Contains accumulated Entry credit totals within the batch.
	TotalCreditEntryDollarAmount int `json:"totalCredit"`
	// CompanyIdentification is an alphanumeric code used to identify an Originator
	// The Company Identification Field must be included on all
	// prenotification records and on each entry initiated pursuant to such
	// prenotification. The Company ID may begin with the ANSI one-digit
	// Identification Code Designator (ICD), followed by the identification
	// number The ANSI Identification Numbers and related Identification Code
	// Designator (ICD) are:
	//
	// IRS Employer Identification Number (EIN) "1"
	// Data Universal Numbering Systems (DUNS) "3"
	// User Assigned Number "9"
	CompanyIdentification string `json:"companyIdentification"`
	// MessageAuthenticationCode the MAC is an eight character code derived from a special key used in
	// conjunction with the DES algorithm. The purpose of the MAC is to
	// validate the authenticity of ACH entries. The DES algorithm and key
	// message standards must be in accordance with standards adopted by the
	// American National Standards Institute. The remaining eleven characters
	// of this field are blank.
	MessageAuthenticationCode string `json:"messageAuthentication,omitempty"`
	// ODFIIdentification the routing number is used to identify the DFI originating entries within a given branch.
	ODFIIdentification string `json:"ODFIIdentification"`
	// BatchNumber this number is assigned in ascending sequence to each batch by the ODFI
	// or its Sending Point in a given file of entries. Since the batch number
	// in the Batch Header Record and the Batch Control Record is the same,
	// the ascending sequence number should be assigned by batch and not by record.
	BatchNumber int `json:"batchNumber"`
	// validator is composed for data validation
	validator
	// converters is composed for ACH to golang Converters
	converters

	validateOpts *ValidateOpts
}

func (bc *BatchControl) SetValidation(opts *ValidateOpts) {
	if bc == nil {
		return
	}
	bc.validateOpts = opts
}

// Parse takes the input record string and parses the EntryDetail values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (bc *BatchControl) Parse(record string) {
	if utf8.RuneCountInString(record) != 94 {
		return
	}

	// 1-1 Always "8"
	// 2-4 This is the same as the "Service code" field in previous Batch Header Record
	bc.ServiceClassCode = bc.parseNumField(record[1:4])
	// 5-10 Total number of Entry Detail Record in the batch
	bc.EntryAddendaCount = bc.parseNumField(record[4:10])
	// 11-20 Total of all positions 4-11 on each Entry Detail Record in the batch. This is essentially the sum of all the RDFI routing numbers in the batch.
	// If the sum exceeds 10 digits (because you have lots of Entry Detail Records), lop off the most significant digits of the sum until there are only 10
	bc.EntryHash = bc.parseNumField(record[10:20])
	// 21-32 Number of cents of debit entries within the batch
	bc.TotalDebitEntryDollarAmount = bc.parseNumField(record[20:32])
	// 33-44 Number of cents of credit entries within the batch
	bc.TotalCreditEntryDollarAmount = bc.parseNumField(record[32:44])
	// 45-54 This is the same as the "Company identification" field in previous Batch Header Record
	bc.CompanyIdentification = bc.parseStringFieldWithOpts(record[44:54], bc.validateOpts)
	// 55-73 Seems to always be blank
	bc.MessageAuthenticationCode = bc.parseStringFieldWithOpts(record[54:73], bc.validateOpts)
	// 74-79 Always blank (just fill with spaces)
	// 80-87 This is the same as the "ODFI identification" field in previous Batch Header Record
	bc.ODFIIdentification = bc.parseStringFieldWithOpts(record[79:87], bc.validateOpts)
	// 88-94 This is the same as the "Batch number" field in previous Batch Header Record
	bc.BatchNumber = bc.parseNumField(record[87:94])
}

// NewBatchControl returns a new BatchControl with default values for none exported fields
func NewBatchControl() *BatchControl {
	return &BatchControl{
		ServiceClassCode: MixedDebitsAndCredits,
		EntryHash:        1,
		BatchNumber:      1,
	}
}

// String writes the BatchControl struct to a 94 character string.
func (bc *BatchControl) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(batchControlPos)
	buf.WriteString(fmt.Sprintf("%v", bc.ServiceClassCode))
	buf.WriteString(bc.EntryAddendaCountField())
	buf.WriteString(bc.EntryHashField())
	buf.WriteString(bc.TotalDebitEntryDollarAmountField())
	buf.WriteString(bc.TotalCreditEntryDollarAmountField())
	buf.WriteString(bc.CompanyIdentificationField())
	buf.WriteString(bc.MessageAuthenticationCodeField())
	buf.WriteString("      ")
	buf.WriteString(bc.ODFIIdentificationField())
	buf.WriteString(bc.BatchNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (bc *BatchControl) Validate() error {
	if err := bc.fieldInclusion(); err != nil {
		return err
	}
	if err := bc.isServiceClass(bc.ServiceClassCode); err != nil {
		return fieldError("ServiceClassCode", err, strconv.Itoa(bc.ServiceClassCode))
	}

	if err := bc.isAlphanumeric(bc.CompanyIdentification); err != nil {
		return fieldError("CompanyIdentification", err, bc.CompanyIdentification)
	}

	if err := bc.isAlphanumeric(bc.MessageAuthenticationCode); err != nil {
		return fieldError("MessageAuthenticationCode", err, bc.MessageAuthenticationCode)
	}

	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (bc *BatchControl) fieldInclusion() error {
	if bc.ServiceClassCode == 0 {
		return fieldError("ServiceClassCode", ErrConstructor, strconv.Itoa(bc.ServiceClassCode))
	}
	if bc.ODFIIdentification == "000000000" {
		return fieldError("ODFIIdentification", ErrConstructor, bc.ODFIIdentificationField())
	}
	return nil
}

// EntryAddendaCountField gets a string of the addenda count zero padded
func (bc *BatchControl) EntryAddendaCountField() string {
	return bc.numericField(bc.EntryAddendaCount, 6)
}

// EntryHashField get a zero padded EntryHash
func (bc *BatchControl) EntryHashField() string {
	return bc.numericField(bc.EntryHash, 10)
}

// TotalDebitEntryDollarAmountField get a zero padded Debit Entry Amount
func (bc *BatchControl) TotalDebitEntryDollarAmountField() string {
	return bc.numericField(bc.TotalDebitEntryDollarAmount, 12)
}

// TotalCreditEntryDollarAmountField get a zero padded Credit Entry Amount
func (bc *BatchControl) TotalCreditEntryDollarAmountField() string {
	return bc.numericField(bc.TotalCreditEntryDollarAmount, 12)
}

// CompanyIdentificationField get the CompanyIdentification right padded
func (bc *BatchControl) CompanyIdentificationField() string {
	return bc.alphaField(bc.CompanyIdentification, 10)
}

// MessageAuthenticationCodeField get the MessageAuthenticationCode right padded
func (bc *BatchControl) MessageAuthenticationCodeField() string {
	return bc.alphaField(bc.MessageAuthenticationCode, 19)
}

// ODFIIdentificationField get the odfi number zero padded
func (bc *BatchControl) ODFIIdentificationField() string {
	return bc.stringField(bc.ODFIIdentification, 8)
}

// BatchNumberField gets a string of the batch number zero padded
func (bc *BatchControl) BatchNumberField() string {
	return bc.numericField(bc.BatchNumber, 7)
}
