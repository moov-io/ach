// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strconv"
)

// BatchTRX holds the BatchHeader and BatchControl and all EntryDetail for TRX Entries.
//
// Check Truncation Entries Exchange is used to identify a debit entry of a truncated checks (multiple).
type BatchTRX struct {
	Batch
}

// NewBatchTRX returns a *BatchTRX
func NewBatchTRX(bh *BatchHeader) *BatchTRX {
	batch := new(BatchTRX)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *BatchTRX) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration and type specific validation for this type.

	if batch.Header.StandardEntryClassCode != TRX {
		return batch.Error("StandardEntryClassCode", ErrBatchSECType, TRX)
	}

	// TRX detail entries can only be a debit, ServiceClassCode must allow debits
	switch batch.Header.ServiceClassCode {
	case MixedDebitsAndCredits, CreditsOnly:
		return batch.Error("ServiceClassCode", ErrBatchServiceClassCode, batch.Header.ServiceClassCode)
	}

	for _, entry := range batch.Entries {
		// TRX detail entries must be a debit
		if entry.CreditOrDebit() != "D" {
			return batch.Error("TransactionCode", ErrBatchDebitOnly, entry.TransactionCode)
		}
		// Trapping this error, as entry.CTXAddendaRecordsField() can not be greater than 9999
		if len(entry.Addenda05) > 9999 {
			return batch.Error("AddendaCount", NewErrBatchAddendaCount(len(entry.Addenda05), 9999))
		}
		// validate CTXAddendaRecord Field is equal to the actual number of Addenda records
		// use 0 value if there is no Addenda records
		addendaRecords, _ := strconv.Atoi(entry.CATXAddendaRecordsField())
		if len(entry.Addenda05) != addendaRecords {
			return batch.Error("AddendaCount", NewErrBatchExpectedAddendaCount(len(entry.Addenda05), addendaRecords))
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
func (batch *BatchTRX) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...

	return batch.Validate()
}
