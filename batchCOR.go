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
	Batch
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
	if batch.Header.StandardEntryClassCode != COR {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, COR)
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
		*/
		switch entry.TransactionCode {
		case
			CheckingCredit, CheckingDebit, CheckingPrenoteCredit, CheckingPrenoteDebit,
			CheckingZeroDollarRemittanceCredit, CheckingZeroDollarRemittanceDebit,
			SavingsCredit, SavingsDebit, SavingsPrenoteCredit, SavingsPrenoteDebit,
			SavingsZeroDollarRemittanceCredit, SavingsZeroDollarRemittanceDebit,
			GLCredit, GLDebit, GLPrenoteCredit, GLPrenoteDebit, GLZeroDollarRemittanceCredit,
			GLZeroDollarRemittanceDebit, LoanCredit, LoanDebit, LoanPrenoteCredit,
			LoanZeroDollarRemittanceCredit:
			msg := fmt.Sprintf(msgBatchTransactionCode, entry.TransactionCode, COR)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TransactionCode", Msg: msg}
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

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate() function at the end of their execution.
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
		if entry.Addenda98 == nil {
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda98", Msg: msgBatchCORAddenda}
		}
	}
	return nil
}
