// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
)

// BatchPOS holds the BatchHeader and BatchControl and all EntryDetail for POS Entries.
//
// A POS Entry is a debit Entry initiated at an “electronic terminal” to a Consumer
// Account of the Receiver to pay an obligation incurred in a point- of-sale
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
	batch
}

// NewBatchPOS returns a *BatchPOS
func NewBatchPOS(bh *BatchHeader) *BatchPOS {
	batch := new(BatchPOS)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate checks valid NACHA batch rules. Assumes properly parsed records.
func (batch *BatchPOS) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}

	// Add configuration and type specific validation for this type.

	if batch.Header.StandardEntryClassCode != "POS" {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, "POS")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	// POS detail entries can only be a debit, ServiceClassCode must allow debits
	switch batch.Header.ServiceClassCode {
	case 200, 220, 280:
		msg := fmt.Sprintf(msgBatchServiceClassCode, batch.Header.ServiceClassCode, "POS")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "ServiceClassCode", Msg: msg}
	}

	for _, entry := range batch.Entries {
		// POS detail entries must be a debit
		if entry.CreditOrDebit() != "D" {
			msg := fmt.Sprintf(msgBatchTransactionCodeCredit, entry.TransactionCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TransactionCode", Msg: msg}
		}
		if err := entry.isCardTransactionType(entry.DiscretionaryData); err != nil {
			msg := fmt.Sprintf(msgBatchCardTransactionType, entry.DiscretionaryData)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "CardTransactionType", Msg: msg}
		}

		// Verify Addenda* FieldInclusion based on entry.Category and batchHeader.StandardEntryClassCode
		if err := batch.addendaFieldInclusion(entry); err != nil {
			return err
		}
	}
	return nil
}

// Create takes Batch Header and Entries and builds a valid batch
func (batch *BatchPOS) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...
	return batch.Validate()
}
