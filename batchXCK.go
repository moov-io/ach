// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "fmt"

// BatchXCK holds the BatchHeader and BatchControl and all EntryDetail for XCK Entries.
//
// Destroyed Check Entry identifies a debit entry initiated for a XCK eligible items.
type BatchXCK struct {
	Batch
}

// NewBatchXCK returns a *BatchXCK
func NewBatchXCK(bh *BatchHeader) *BatchXCK {
	batch := new(BatchXCK)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate checks valid NACHA batch rules. Assumes properly parsed records.
func (batch *BatchXCK) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration and type specific validation for this type.

	if batch.Header.StandardEntryClassCode != XCK {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, XCK)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	// XCK detail entries can only be a debit, ServiceClassCode must allow debits
	switch batch.Header.ServiceClassCode {
	case MixedDebitsAndCredits, CreditsOnly:
		msg := fmt.Sprintf(msgBatchServiceClassCode, batch.Header.ServiceClassCode, XCK)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "ServiceClassCode", Msg: msg}
	}

	for _, entry := range batch.Entries {
		// XCK detail entries must be a debit
		if entry.CreditOrDebit() != "D" {
			msg := fmt.Sprintf(msgBatchTransactionCodeCredit, entry.TransactionCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TransactionCode", Msg: msg}
		}
		// Amount must be 2,500 or less
		if entry.Amount > 250000 {
			msg := fmt.Sprintf(msgBatchAmount, "2,500", XCK)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Amount", Msg: msg}
		}
		// ProcessControlField underlying IdentificationNumber, must be defined
		if entry.ProcessControlField() == "" {
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "ProcessControlField", Msg: msgFieldRequired}
		}
		// ItemResearchNumber underlying IdentificationNumber, must be defined
		if entry.ItemResearchNumber() == "" {
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "ItemResearchNumber", Msg: msgFieldRequired}
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
func (batch *BatchXCK) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...

	return batch.Validate()
}
