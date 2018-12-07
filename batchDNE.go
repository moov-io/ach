// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strings"
)

// BatchDNE is a batch file that handles SEC code Death Notification Entry (DNE)
// United States Federal agencies (e.g. Social Security) use this to notify depository
// financial institutions that the recipient of government benefit paymetns has died.
//
// Notes:
//  - Date of death always in positions 18-23
//  - SSN (positions 38-46) are zero if no SSN
//  - Beneficiary payment starts at position 55
type BatchDNE struct {
	Batch
}

// NewBatchDNE returns a *BatchDNE
func NewBatchDNE(bh *BatchHeader) *BatchDNE {
	batch := new(BatchDNE)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate ensures the batch meets NACHA rules specific to this batch type.
func (batch *BatchDNE) Validate() error {
	if err := batch.verify(); err != nil {
		return err
	}

	// SEC code
	if batch.Header.StandardEntryClassCode != DNE {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, DNE)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	// Range over Entries
	for _, entry := range batch.Entries {
		if entry.Amount != 0 {
			msg := fmt.Sprintf(msgBatchAmountZero, entry.Amount, DNE)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Amount", Msg: msg}
		}

		switch entry.TransactionCode {
		case CheckingReturnNOCCredit, CheckingPrenoteCredit, SavingsReturnNOCCredit, SavingsPrenoteCredit:
		default:
			msg := fmt.Sprintf(msgBatchTransactionCode, entry.TransactionCode, DNE)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TransactionCode", Msg: msg}
		}

		// DNE must have one Addenda05
		if len(entry.Addenda05) != 1 {
			msg := fmt.Sprintf(msgBatchAddendaCount, len(entry.Addenda05), 1, batch.Header.StandardEntryClassCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "AddendaCount", Msg: msg}
		}
		// Verify the TransactionCode is valid for a ServiceClassCode
		if err := batch.ValidTranCodeForServiceClassCode(entry); err != nil {
			return err
		}
		// Verify Addenda* FieldInclusion based on entry.Category and batchHeader.StandardEntryClassCode
		if err := batch.addendaFieldInclusion(entry); err != nil {
			return err
		}
	}
	return nil
}

// Create builds the batch sequence numbers and batch control. Additional creation
func (batch *BatchDNE) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	return batch.Validate()
}

// details returns the Date of Death (YYMMDD), Customer SSN (9 digits), and Amount ($$$$.cc)
// from the Addenda05 record. This method assumes the addenda05 PaymentRelatedInformation is valid.
func (batch *BatchDNE) details() (string, string, string) {
	addendas := batch.Entries[0].Addenda05
	if len(addendas) != 1 {
		return "", "", ""
	}
	line := addendas[0].PaymentRelatedInformation
	return line[18:24], line[37:46], strings.TrimSuffix(line[54:], `\`)
}

// DateOfDeath returns the YYMMDD string from Addenda05's PaymentRelatedInformation
func (batch *BatchDNE) DateOfDeath() string {
	date, _, _ := batch.details()
	return date
}

// CustomerSSN returns the SSN string from Addenda05's PaymentRelatedInformation
func (batch *BatchDNE) CustomerSSN() string {
	_, ssn, _ := batch.details()
	return ssn
}

// Amount returns the amount to be dispursed to the named beneficiary from Addenda05's PaymentRelatedInformation.
func (batch *BatchDNE) Amount() string {
	_, _, amount := batch.details()
	return amount
}
