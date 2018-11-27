// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "fmt"

// BatchBOC holds the BatchHeader and BatchControl and all EntryDetail for BOC Entries.
//
// Back Office Conversion (BOC) A single entry debit initiated at the point of purchase
// or at a manned bill payment location to transfer funds through conversion to an
// ACH debit entry during back office processing.
//
// BOC allows retailers/billers, and ODFIs acting as Originators,
// to electronically convert checks received at the point-of-purchase as well as at a
// manned bill payment location into a single-entry ACH debit. The authorization to
// convert the check will be obtained through a notice at the checkout or manned bill
// payment location (e.g., loan payment at financial institution’s teller window) and the
// receipt of the Receiver’s check. The decision to process the check item as an ACH debit
// will be made in the “back office” instead of at the point-of-purchase. The customer’s
// check will solely be used as a source document to obtain the routing number, account
// number and check serial number.
//
// Unlike ARC entries, BOC conversions require the customer to be present and a notice that
// checks may be converted to BOC ACH entries be posted.
type BatchBOC struct {
	Batch
}

// NewBatchBOC returns a *BatchBOC
func NewBatchBOC(bh *BatchHeader) *BatchBOC {
	batch := new(BatchBOC)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate checks valid NACHA batch rules. Assumes properly parsed records.
func (batch *BatchBOC) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration and type specific validation for this type.

	if batch.Header.StandardEntryClassCode != "BOC" {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, "BOC")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	// BOC detail entries can only be a debit, ServiceClassCode must allow debits
	switch batch.Header.ServiceClassCode {
	case 200, 220:
		msg := fmt.Sprintf(msgBatchServiceClassCode, batch.Header.ServiceClassCode, "RCK")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "ServiceClassCode", Msg: msg}
	}

	for _, entry := range batch.Entries {
		// BOC detail entries must be a debit
		if entry.CreditOrDebit() != "D" {
			msg := fmt.Sprintf(msgBatchTransactionCodeCredit, entry.TransactionCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TransactionCode", Msg: msg}
		}

		// Amount must be 25,000 or less
		if entry.Amount > 2500000 {
			msg := fmt.Sprintf(msgBatchAmount, "25,000", "BOC")
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Amount", Msg: msg}
		}

		// CheckSerialNumber underlying IdentificationNumber, must be defined
		if entry.IdentificationNumber == "" {
			msg := fmt.Sprintf(msgBatchCheckSerialNumber, "BOC")
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "CheckSerialNumber", Msg: msg}
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

// Create takes Batch Header and Entries and builds a valid batch
func (batch *BatchBOC) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...

	return batch.Validate()
}
