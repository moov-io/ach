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

// ADVBatchControl contains entry counts, dollar total and has totals for all
// entries contained in the preceding batch
type ADVBatchControl struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// This should be the same as BatchHeader ServiceClassCode for ADV: AutomatedAccountingAdvices.
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
	// ACHOperatorData is an alphanumeric code used to identify an ACH Operator
	ACHOperatorData string `json:"achOperatorData"`
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
}

// Parse takes the input record string and parses the EntryDetail values
func (bc *ADVBatchControl) Parse(record string) {
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
			// 1-1 Always "8"
			reset()
		case 4:
			// 2-4 This is the same as the "Service code" field in previous Batch Header Record
			bc.ServiceClassCode = bc.parseNumField(reset())
		case 10:
			// 5-10 Total number of Entry Detail Record in the batch
			bc.EntryAddendaCount = bc.parseNumField(reset())
		case 20:
			// 11-20 Total of all positions 4-11 on each Entry Detail Record in the batch. This is essentially the sum of all the RDFI routing numbers in the batch.
			// If the sum exceeds 10 digits (because you have lots of Entry Detail Records), lop off the most significant digits of the sum until there are only 10
			bc.EntryHash = bc.parseNumField(reset())
		case 40:
			// 21-32 Number of cents of debit entries within the batch
			bc.TotalDebitEntryDollarAmount = bc.parseNumField(reset())
		case 60:
			// 33-44 Number of cents of credit entries within the batch
			bc.TotalCreditEntryDollarAmount = bc.parseNumField(reset())
		case 79:
			// 45-54 ACH Operator Data
			bc.ACHOperatorData = strings.TrimSpace(reset())
		case 87:
			// 80-87 This is the same as the "ODFI identification" field in previous Batch Header Record
			bc.ODFIIdentification = bc.parseStringField(reset())
		case 94:
			// 88-94 This is the same as the "Batch number" field in previous Batch Header Record
			bc.BatchNumber = bc.parseNumField(reset())
		}
	}
}

// NewADVBatchControl returns a new ADVBatchControl with default values for none exported fields
func NewADVBatchControl() *ADVBatchControl {
	return &ADVBatchControl{
		ServiceClassCode: AutomatedAccountingAdvices,
		EntryHash:        1,
		BatchNumber:      1,
	}
}

// String writes the ADVBatchControl struct to a 94 character string.
func (bc *ADVBatchControl) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(batchControlPos)
	buf.WriteString(fmt.Sprintf("%v", bc.ServiceClassCode))
	buf.WriteString(bc.EntryAddendaCountField())
	buf.WriteString(bc.EntryHashField())
	buf.WriteString(bc.TotalDebitEntryDollarAmountField())
	buf.WriteString(bc.TotalCreditEntryDollarAmountField())
	buf.WriteString(bc.ACHOperatorDataField())
	buf.WriteString(bc.ODFIIdentificationField())
	buf.WriteString(bc.BatchNumberField())
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (bc *ADVBatchControl) Validate() error {
	if err := bc.fieldInclusion(); err != nil {
		return err
	}
	if err := bc.isServiceClass(bc.ServiceClassCode); err != nil {
		return fieldError("ServiceClassCode", err, strconv.Itoa(bc.ServiceClassCode))
	}

	if err := bc.isAlphanumeric(bc.ACHOperatorData); err != nil {
		return fieldError("ACHOperatorData", err, bc.ACHOperatorData)
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (bc *ADVBatchControl) fieldInclusion() error {
	if bc.ServiceClassCode == 0 {
		return fieldError("ServiceClassCode", ErrConstructor, strconv.Itoa(bc.ServiceClassCode))
	}
	if bc.ODFIIdentification == "000000000" || bc.ODFIIdentification == "" {
		return fieldError("ODFIIdentification", ErrConstructor, bc.ODFIIdentificationField())
	}
	return nil
}

// EntryAddendaCountField gets a string of the addenda count zero padded
func (bc *ADVBatchControl) EntryAddendaCountField() string {
	return bc.numericField(bc.EntryAddendaCount, 6)
}

// EntryHashField get a zero padded EntryHash
func (bc *ADVBatchControl) EntryHashField() string {
	return bc.numericField(bc.EntryHash, 10)
}

// TotalDebitEntryDollarAmountField get a zero padded Debit Entry Amount
func (bc *ADVBatchControl) TotalDebitEntryDollarAmountField() string {
	return bc.numericField(bc.TotalDebitEntryDollarAmount, 20)
}

// TotalCreditEntryDollarAmountField get a zero padded Credit Entry Amount
func (bc *ADVBatchControl) TotalCreditEntryDollarAmountField() string {
	return bc.numericField(bc.TotalCreditEntryDollarAmount, 20)
}

// ACHOperatorDataField get the ACHOperatorData right padded
func (bc *ADVBatchControl) ACHOperatorDataField() string {
	return bc.alphaField(bc.ACHOperatorData, 19)
}

// ODFIIdentificationField get the odfi number zero padded
func (bc *ADVBatchControl) ODFIIdentificationField() string {
	return bc.stringField(bc.ODFIIdentification, 8)
}

// BatchNumberField gets a string of the batch number zero padded
func (bc *ADVBatchControl) BatchNumberField() string {
	return bc.numericField(bc.BatchNumber, 7)
}
