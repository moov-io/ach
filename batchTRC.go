// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "fmt"

// BatchTRC holds the BatchHeader and BatchControl and all EntryDetail for TRC Entries.
//
// Check Truncation Entry (Truncated Entry) is used to identify a debit entry of a truncated check.
type BatchTRC struct {
	Batch
}

// NewBatchTRC returns a *BatchTRC
func NewBatchTRC(bh *BatchHeader) *BatchTRC {
	batch := new(BatchTRC)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate checks valid NACHA batch rules. Assumes properly parsed records.
func (batch *BatchTRC) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration and type specific validation for this type.

	if batch.Header.StandardEntryClassCode != TRC {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, TRC)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	// TRC detail entries can only be a debit, ServiceClassCode must allow debits
	switch batch.Header.ServiceClassCode {
	case MixedDebitsAndCredits, CreditsOnly:
		msg := fmt.Sprintf(msgBatchServiceClassCode, batch.Header.ServiceClassCode, TRC)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "ServiceClassCode", Msg: msg}
	}

	for _, entry := range batch.Entries {
		// TRC detail entries must be a debit
		if entry.CreditOrDebit() != "D" {
			msg := fmt.Sprintf(msgBatchTransactionCodeCredit, entry.TransactionCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TransactionCode", Msg: msg}
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
func (batch *BatchTRC) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...

	return batch.Validate()
}
