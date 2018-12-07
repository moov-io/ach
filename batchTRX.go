// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strconv"
)

// BatchTRX holds the BatchHeader and BatchControl and all EntryDetail for TRX Entries.
//
// Check Truncation Entries Exchange is used to identify a debit entry of a truncated checks (multiple).
type BatchTRX struct {
	Batch
}

var (
	msgBatchTRXAddendaCount = "%v entry detail addenda records not equal to addendum %v"
)

// NewBatchTRX returns a *BatchTRX
func NewBatchTRX(bh *BatchHeader) *BatchTRX {
	batch := new(BatchTRX)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate checks valid NACHA batch rules. Assumes properly parsed records.
func (batch *BatchTRX) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration and type specific validation for this type.

	if batch.Header.StandardEntryClassCode != TRX {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, TRX)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	// TRX detail entries can only be a debit, ServiceClassCode must allow debits
	switch batch.Header.ServiceClassCode {
	case MixedDebitsAndCredits, CreditsOnly:
		msg := fmt.Sprintf(msgBatchServiceClassCode, batch.Header.ServiceClassCode, TRX)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "ServiceClassCode", Msg: msg}
	}

	for _, entry := range batch.Entries {
		// TRX detail entries must be a debit
		if entry.CreditOrDebit() != "D" {
			msg := fmt.Sprintf(msgBatchTransactionCodeCredit, entry.TransactionCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TransactionCode", Msg: msg}
		}
		// Trapping this error, as entry.CTXAddendaRecordsField() can not be greater than 9999
		if len(entry.Addenda05) > 9999 {
			msg := fmt.Sprintf(msgBatchAddendaCount, len(entry.Addenda05), 9999, batch.Header.StandardEntryClassCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "AddendaCount", Msg: msg}
		}
		// validate CTXAddendaRecord Field is equal to the actual number of Addenda records
		// use 0 value if there is no Addenda records
		addendaRecords, _ := strconv.Atoi(entry.CATXAddendaRecordsField())
		if len(entry.Addenda05) != addendaRecords {
			msg := fmt.Sprintf(msgBatchTRXAddendaCount, addendaRecords, len(entry.Addenda05))
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addendum", Msg: msg}
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
func (batch *BatchTRX) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...

	return batch.Validate()
}
