// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

// BatchACK is a batch file that handles SEC payment type ACK and ACK+.
// Acknowledgement of a Corporate credit by the Receiving Depository Financial Institution (RDFI).
// For commercial accounts only.
type BatchACK struct {
	Batch
}

// NewBatchACK returns a *BatchACK
func NewBatchACK(bh *BatchHeader) *BatchACK {
	batch := new(BatchACK)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *BatchACK) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration and type specific validation.
	if batch.Header.StandardEntryClassCode != ACK {
		return batch.Error("StandardEntryClassCode", ErrBatchSECType, ACK)
	}
	// Range through Entries
	for _, entry := range batch.Entries {
		// Amount must be zero for Acknowledgement Entries
		if entry.Amount > 0 {
			return batch.Error("Amount", ErrBatchAmountNonZero, entry.Amount)
		}
		if len(entry.Addenda05) > 1 {
			return batch.Error("AddendaCount", NewErrBatchAddendaCount(len(entry.Addenda05), 1))
		}
		switch entry.TransactionCode {
		case CheckingZeroDollarRemittanceCredit, SavingsZeroDollarRemittanceCredit:
		default:
			return batch.Error("TransactionCode", ErrBatchTransactionCode, entry.TransactionCode)
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
func (batch *BatchACK) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	return batch.Validate()
}
