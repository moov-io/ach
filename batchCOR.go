// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
)

// BatchCOR COR - Automated Notification of Change (NOC) or Refused Notification of Change
// This Standard Entry Class Code is used by an RDFI or ODFI when originating a Notification of Change or Refused Notification of Change in automated format.
// A Notification of Change may be created by an RDFI to notify the ODFI that a posted Entry or Prenotification Entry contains invalid or erroneous information and should be changed.
type BatchCOR struct {
	batch
}

var msgBatchCORAmount = "debit:%v credit:%v entry detail amount fields must be zero for SEC type COR"
var msgBatchCORAddenda = "found and 1 Addenda98 is required for SEC Type COR"

// NewBatchCOR returns a *BatchCOR
func NewBatchCOR(bh *BatchHeader) *BatchCOR {
	batch := new(BatchCOR)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate ensures the batch meets NACHA rules specific to this batch type.
func (batch *BatchCOR) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration based validation for this type.
	// COR Addenda must be Addenda98
	if err := batch.isAddenda98(); err != nil {
		return err
	}

	// Add type specific validation.
	if batch.Header.StandardEntryClassCode != "COR" {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, "COR")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	// The Amount field must be zero
	// batch.verify calls batch.isBatchAmount which ensures the batch.Control values are accurate.
	if batch.Control.TotalCreditEntryDollarAmount != 0 || batch.Control.TotalDebitEntryDollarAmount != 0 {
		msg := fmt.Sprintf(msgBatchCORAmount, batch.Control.TotalCreditEntryDollarAmount, batch.Control.TotalDebitEntryDollarAmount)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Amount", Msg: msg}
	}

	for _, entry := range batch.Entries {
		/* COR TransactionCode must be a Return or NOC transaction Code
			   Return/NOC
			   Credit:  21, 31, 41, 51
			   Debit: 26, 36, 46, 56

			   Automated payment/deposit
			   Credit: 22, 32, 42, 52
			   Debit: 27, 37, 47, 55 (reversal)

			   Prenote
			   Credit:  23, 33, 43, 53
			   Debit: 28, 38, 48

			   Zero dollar amount with remittance data
			   Credit: 24, 34, 44, 54
		 	   Debit: 29, 39, 49
		*/
		switch entry.TransactionCode {
		case 22, 27, 32, 37, 42, 47, 52, 55,
			23, 28, 33, 38, 43, 48, 53,
			24, 29, 34, 39, 44, 49, 54:
			msg := fmt.Sprintf(msgBatchTransactionCode, entry.TransactionCode, "COR")
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TransactionCode", Msg: msg}
		}

		// Verify Addenda* FieldInclusion based on entry.Category and batchHeader.StandardEntryClassCode
		if err := batch.addendaFieldInclusion(entry); err != nil {
			return err
		}
	}

	return nil
}

// Create builds the batch sequence numbers and batch control. Additional creation
func (batch *BatchCOR) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}

	return batch.Validate()
}

// isAddenda98 verifies that a Addenda98 exists for each EntryDetail and is Validated
func (batch *BatchCOR) isAddenda98() error {
	for _, entry := range batch.Entries {
		// ToDo: May be able to get rid of the first check
		if entry.Addenda98 == nil {
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addendum", Msg: msgBatchCORAddenda}
		}
		// Addenda98 must be Validated
		if err := entry.Addenda98.Validate(); err != nil {
			// convert the field error in to a batch error for a consistent api
			if e, ok := err.(*FieldError); ok {
				return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: e.FieldName, Msg: e.Msg}
			}
		}
	}
	return nil
}
