// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
)

// BatchCIE holds the BatchHeader and BatchControl and all EntryDetail for CIE Entries.
//
// Customer-Initiated Entry (or CIE entry) is a credit entry initiated on behalf of,
// and upon the instruction of, a consumer to transfer funds to a non-consumer Receiver.
// CIE entries are usually transmitted to a company for payment of funds that the consumer
// owes to that company and are initiated by the consumer through some type of online
// banking product or bill payment service provider. With CIEs, funds owed by the consumer
// are “pushed” to the biller in the form of an ACH credit, as opposed to the biller’s use of
// a debit application (e.g., PPD, WEB) to “pull” the funds from a customer’s account.
type BatchCIE struct {
	batch
}

var msgBatchCIEAddenda = "found and 1 Addenda05 is the maximum for SEC code CIE"
var msgBatchCIEAddendaType = "%T found where Addenda05 is required for SEC code CIE"

// NewBatchCIE returns a *BatchCIE
func NewBatchCIE(bh *BatchHeader) *BatchCIE {
	batch := new(BatchCIE)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate checks valid NACHA batch rules. Assumes properly parsed records.
func (batch *BatchCIE) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration based validation for this type.

	// Add type specific validation.

	if batch.Header.StandardEntryClassCode != "CIE" {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, "CIE")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	// CIE detail entries can only be a debit, ServiceClassCode must allow debits
	switch batch.Header.ServiceClassCode {
	case 200, 225, 280:
		msg := fmt.Sprintf(msgBatchServiceClassCode, batch.Header.ServiceClassCode, "CIE")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "ServiceClassCode", Msg: msg}
	}

	for _, entry := range batch.Entries {
		// CIE detail entries must be a debit
		if entry.CreditOrDebit() != "C" {
			msg := fmt.Sprintf(msgBatchTransactionCodeCredit, entry.TransactionCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TransactionCode", Msg: msg}
		}

		// Addenda validations - CIE Addenda must be Addenda05

		// Addendum must be equal to 1
		if len(entry.Addendum) > 1 {
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addendum", Msg: msgBatchCIEAddenda}
		}

		if len(entry.Addendum) > 0 {
			// Addenda type assertion must be Addenda05
			addenda05, ok := entry.Addendum[0].(*Addenda05)
			if !ok {
				msg := fmt.Sprintf(msgBatchCIEAddendaType, entry.Addendum[0])
				return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addendum", Msg: msg}
			}

			// Addenda05 must be Validated
			if err := addenda05.Validate(); err != nil {
				// convert the field error in to a batch error for a consistent api
				if e, ok := err.(*FieldError); ok {
					return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: e.FieldName, Msg: e.Msg}
				}
			}
		}
	}
	return nil
}

// Create takes Batch Header and Entries and builds a valid batch
func (batch *BatchCIE) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...
	return batch.Validate()
}
