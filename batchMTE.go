// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"

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
		return batch.Error("StandardEntryClassCode", ErrBatchSECType, MTE)
	}

	for _, entry := range batch.Entries {
		if entry.Amount <= 0 {
			return batch.Error("Amount", ErrBatchAmountZero, entry.Amount)
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
				return batch.Error("TerminalState", ErrValidState, entry.Addenda02.TerminalState)
			}
		}

		// MTE entries cannot have an identification number that is all spaces or all zeros
		if strings.Trim(entry.IdentificationNumber, " 0") == "" {
			return batch.Error("IdentificationNumber", ErrIdentificationNumber, entry.IdentificationNumber)
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
