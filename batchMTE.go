// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"github.com/moov-io/ach/internal/usabbrev"
)

// BatchMTE holds the BatchHeader, BatchControl, and EntryDetail for Machine Transfer Entry (MTE) entries.
//
// A MTE transaction is created when a consumer uses their debit card at an Automated Teller Machine (ATM) to withdraw cash.
// MTE transactions cannot be aggregated together under a single Entry.
type BatchMTE struct {
	Batch
}

// NewBatchMTE returns a *BatchMTE
func NewBatchMTE(bh *BatchHeader) *BatchMTE {
	batch := new(BatchMTE)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *BatchMTE) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}

	if batch.Header.StandardEntryClassCode != MTE {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, MTE)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	for _, entry := range batch.Entries {
		if entry.Amount <= 0 {
			msg := fmt.Sprintf(msgBatchAmountNonZero, entry.Amount, MTE)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Amount", Msg: msg}
		}
		// Verify the TransactionCode is valid for a ServiceClassCode
		if err := batch.ValidTranCodeForServiceClassCode(entry); err != nil {
			return err
		}
		// Verify Addenda* FieldInclusion based on entry.Category and batchHeader.StandardEntryClassCode
		if err := batch.addendaFieldInclusion(entry); err != nil {
			return err
		}
		if entry.Category == CategoryForward {
			if !usabbrev.Valid(entry.Addenda02.TerminalState) {
				msg := fmt.Sprintf("%q is not a valid US state or territory", entry.Addenda02.TerminalState)
				return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TerminalState", Msg: msg}
			}
		}
	}
	return nil
}

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate() function at the end of their execution.
func (batch *BatchMTE) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	return batch.Validate()
}
