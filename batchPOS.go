// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"github.com/moov-io/ach/internal/usabbrev"
)

// BatchPOS holds the BatchHeader and BatchControl and all EntryDetail for POS Entries.
//
// A POS Entry is a debit Entry initiated at an “electronic terminal” to a consumer
// account of the receiver to pay an obligation incurred in a point- of-sale
// transaction, or to effect a point-of-sale terminal cash withdrawal.
//
// Point-of-Sale Entries (POS) are ACH debit entries typically initiated by the use
// of a merchant-issued plastic card to pay an obligation at the point-of-sale. Much
// like a financial institution issued debit card, the merchant- issued debit card is
// swiped at the point-of-sale and approved for use; however, the authorization only
// verifies the card is open, active and within the card’s limits—it does not verify
// the Receiver’s account balance or debit the account at the time of the purchase.
// Settlement of the transaction moves from the card network to the ACH Network through
// the creation of a POS entry by the card issuer to debit the Receiver’s account.
type BatchPOS struct {
	Batch
}

// NewBatchPOS returns a *BatchPOS
func NewBatchPOS(bh *BatchHeader) *BatchPOS {
	batch := new(BatchPOS)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *BatchPOS) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}

	// Add configuration and type specific validation for this type.

	if batch.Header.StandardEntryClassCode != POS {
		return batch.Error("StandardEntryClassCode", ErrBatchSECType, POS)
	}

	// POS detail entries can only be a debit, ServiceClassCode must allow debits
	switch batch.Header.ServiceClassCode {
	case MixedDebitsAndCredits, CreditsOnly:
		return batch.Error("ServiceClassCode", ErrBatchServiceClassCode, batch.Header.ServiceClassCode)
	}

	for _, entry := range batch.Entries {
		// POS detail entries must be a debit
		if entry.CreditOrDebit() != "D" {
			return batch.Error("TransactionCode", ErrBatchDebitOnly, entry.TransactionCode)
		}
		if err := entry.isCardTransactionType(entry.DiscretionaryData); err != nil {
			return batch.Error("CardTransactionType", ErrBatchInvalidCardTransactionType, entry.DiscretionaryData)
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
	}
	return nil
}

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate() function at the end of their execution.
func (batch *BatchPOS) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...
	return batch.Validate()
}
