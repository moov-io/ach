// Copyright 2018 The ACH Authors
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

var msgBatchPOSAddenda = "found and 1 Addenda02 is required for SEC code POS"
var msgBatchPOSAddendaType = "%T found where Addenda02 is required for SEC code POS"

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
	// Add configuration based validation for this type.

	// POS Addenda must be Addenda02
	if err := batch.isAddenda02(); err != nil {
		return err
	}
	// Add type specific validation.
	if batch.Header.StandardEntryClassCode != "POS" {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, "POS")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	//ToDo;  Determine if this is needed
	/*	// POS detail entries can only be a debit, ServiceClassCode must allow debits
		switch batch.Header.ServiceClassCode {
		case 200, 220, 280:
			msg := fmt.Sprintf(msgBatchServiceClassCode, batch.Header.ServiceClassCode, "POS")
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "ServiceClassCode", Msg: msg}
		}*/

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

		//ToDo;  Additional validations -  move isAddenda02 logic so we only range through a batch of entries once
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

// isAddenda02 verifies that a Addenda02 exists for each EntryDetail and is Validated
func (batch *BatchPOS) isAddenda02() error {
	for _, entry := range batch.Entries {
		// Addenda type must be equal to 1
		if len(entry.Addendum) != 1 {
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addendum", Msg: msgBatchPOSAddenda}
		}
		// Addenda type assertion must be Addenda02
		addenda02, ok := entry.Addendum[0].(*Addenda02)
		if !ok {
			msg := fmt.Sprintf(msgBatchPOSAddendaType, entry.Addendum[0])
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addendum", Msg: msg}
		}
		// Addenda02 must be Validated
		if err := addenda02.Validate(); err != nil {
			// convert the field error in to a batch error for a consistent api
			if e, ok := err.(*FieldError); ok {
				return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: e.FieldName, Msg: e.Msg}
			}
		}
	}
	return nil
}
