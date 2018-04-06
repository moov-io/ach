// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strconv"
	"strings"
)

// BatchControl contains entry counts, dollar total and has totals for all
// entries contained in the preceding batch
type BatchControl struct {
	// RecordType defines the type of record in the block. batchControlPos 8
	recordType string
	// ServiceClassCode ACH Mixed Debits and Credits ‘200’
	// ACH Credits Only ‘220’
	// ACH Debits Only ‘225'
	// Same as 'ServiceClassCode' in BatchHeaderRecord
	ServiceClassCode int
	// EntryAddendaCount is a tally of each Entry Detail Record and each Addenda
	// Record processed, within either the batch or file as appropriate.
	EntryAddendaCount int
	// validate the Receiving DFI Identification in each Entry Detail Record is hashed
	// to provide a check against inadvertent alteration of data contents due
	// to hardware failure or program erro
	//
	// In this context the Entry Hash is the sum of the corresponding fields in the
	// Entry Detail Records on the file.
	EntryHash int
	// TotalDebitEntryDollarAmount Contains accumulated Entry debit totals within the batch.
	TotalDebitEntryDollarAmount int
	// TotalCreditEntryDollarAmount Contains accumulated Entry credit totals within the batch.
	TotalCreditEntryDollarAmount int
	// CompanyIdentification is an alphameric code used to identify an Originato
	// The Company Identification Field must be included on all
	// prenotification records and on each entry initiated puruant to such
	// prenotification. The Company ID may begin with the ANSI one-digit
	// Identification Code Designators (ICD), followed by the identification
	// numbe The ANSI Identification Numbers and related Identification Code
	// Designators (ICD) are:
	//
	// IRS Employer Identification Number (EIN) "1"
	// Data Universal Numbering Systems (DUNS) "3"
	// User Assigned Number "9"
	CompanyIdentification string
	// MessageAuthenticationCode the MAC is an eight character code derived from a special key used in
	// conjunction with the DES algorithm. The purpose of the MAC is to
	// validate the authenticity of ACH entries. The DES algorithm and key
	// message standards must be in accordance with standards adopted by the
	// American National Standards Institute. The remaining eleven characters
	// of this field are blank.
	MessageAuthenticationCode string
	// Reserved for the future - Blank, 6 characters long
	reserved string
	// OdfiIdentification the routing number is used to identify the DFI originating entries within a given branch.
	ODFIIdentification int
	// BatchNumber this number is assigned in ascending sequence to each batch by the ODFI
	// or its Sending Point in a given file of entries. Since the batch number
	// in the Batch Header Record and the Batch Control Record is the same,
	// the ascending sequence number should be assigned by batch and not by record.
	BatchNumber int
	// validator is composed for data validation
	validator
	// converters is composed for ACH to golang Converters
	converters
}

// Parse takes the input record string and parses the EntryDetail values
func (bc *BatchControl) Parse(record string) {
	// 1-1 Always "8"
	bc.recordType = "8"
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
	bc.CompanyIdentification = strings.TrimSpace(record[44:54])
	// 55-73 Seems to always be blank
	bc.MessageAuthenticationCode = strings.TrimSpace(record[54:73])
	// 74-79 Always blank (just fill with spaces)
	bc.reserved = "      "
	// 80-87 This is the same as the "ODFI identification" field in previous Batch Header Record
	bc.ODFIIdentification = bc.parseNumField(record[79:87])
	// 88-94 This is the same as the "Batch number" field in previous Batch Header Record
	bc.BatchNumber = bc.parseNumField(record[87:94])
}

// NewBatchControl returns a new BatchControl with default values for none exported fields
func NewBatchControl() *BatchControl {
	return &BatchControl{
		recordType:       "8",
		ServiceClassCode: 200,
		EntryHash:        1,
		BatchNumber:      1,
	}
}

// String writes the BatchControl struct to a 94 character string.
func (bc *BatchControl) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v",
		bc.recordType,
		bc.ServiceClassCode,
		bc.EntryAddendaCountField(),
		bc.EntryHashField(),
		bc.TotalDebitEntryDollarAmountField(),
		bc.TotalCreditEntryDollarAmountField(),
		bc.CompanyIdentificationField(),
		bc.MessageAuthenticationCodeField(),
		"      ",
		bc.ODFIIdentificationField(),
		bc.BatchNumberField(),
	)
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (bc *BatchControl) Validate() error {
	if err := bc.fieldInclusion(); err != nil {
		return err
	}
	if bc.recordType != "8" {
		msg := fmt.Sprintf(msgRecordType, 7)
		return &FieldError{FieldName: "recordType", Value: bc.recordType, Msg: msg}
	}
	if err := bc.isServiceClass(bc.ServiceClassCode); err != nil {
		return &FieldError{FieldName: "ServiceClassCode", Value: strconv.Itoa(bc.ServiceClassCode), Msg: err.Error()}
	}

	if err := bc.isAlphanumeric(bc.CompanyIdentification); err != nil {
		return &FieldError{FieldName: "CompanyIdentification", Value: bc.CompanyIdentification, Msg: err.Error()}
	}

	if err := bc.isAlphanumeric(bc.MessageAuthenticationCode); err != nil {
		return &FieldError{FieldName: "MessageAuthenticationCode", Value: bc.MessageAuthenticationCode, Msg: err.Error()}
	}

	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (bc *BatchControl) fieldInclusion() error {
	if bc.recordType == "" {
		return &FieldError{FieldName: "recordType", Value: bc.recordType, Msg: msgFieldInclusion}
	}
	if bc.ServiceClassCode == 0 {
		return &FieldError{FieldName: "ServiceClassCode", Value: strconv.Itoa(bc.ServiceClassCode), Msg: msgFieldInclusion}
	}
	if bc.ODFIIdentification == 0 {
		return &FieldError{FieldName: "ODFIIdentification", Value: bc.ODFIIdentificationField(), Msg: msgFieldInclusion}
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

//TotalDebitEntryDollarAmountField get a zero padded Debity Entry Amount
func (bc *BatchControl) TotalDebitEntryDollarAmountField() string {
	return bc.numericField(bc.TotalDebitEntryDollarAmount, 12)
}

// TotalCreditEntryDollarAmountField get a zero padded Credit Entry Amount
func (bc *BatchControl) TotalCreditEntryDollarAmountField() string {
	return bc.numericField(bc.TotalCreditEntryDollarAmount, 12)
}

// CompanyIdentificationField get the CompanyIdentification righ padded
func (bc *BatchControl) CompanyIdentificationField() string {
	return bc.alphaField(bc.CompanyIdentification, 10)
}

// MessageAuthenticationCodeField get the MessageAuthenticationCode right padded
func (bc *BatchControl) MessageAuthenticationCodeField() string {
	return bc.alphaField(bc.MessageAuthenticationCode, 19)
}

// ODFIIdentificationField get the odfi number zero padded
func (bc *BatchControl) ODFIIdentificationField() string {
	return bc.numericField(bc.ODFIIdentification, 8)
}

// BatchNumberField gets a string of the batch number zero padded
func (bc *BatchControl) BatchNumberField() string {
	return bc.numericField(bc.BatchNumber, 7)
}
