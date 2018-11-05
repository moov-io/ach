// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strconv"
)

// BatchATX holds the BatchHeader and BatchControl and all EntryDetail for ATX (Acknowledgment)
// Entries.
//
// The ATX entry is an acknowledgement by the Receiving Depository Financial Institution (RDFI) that a
// Corporate Credit (CTX) has been received.
type BatchATX struct {
	batch
}

var (
	msgBatchATXAddendaCount = "%v entry detail addenda records not equal to addendum %v"
)

// NewBatchATX returns a *BatchATX
func NewBatchATX(bh *BatchHeader) *BatchATX {
	batch := new(BatchATX)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate checks valid NACHA batch rules. Assumes properly parsed records.
func (batch *BatchATX) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}

	// Add configuration and type specific validation for this type.
	if batch.Header.StandardEntryClassCode != "ATX" {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, "ATX")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	for _, entry := range batch.Entries {
		// Amount must be zero for Acknowledgement Entries
		if entry.Amount > 0 {
			msg := fmt.Sprintf(msgBatchAmountZero, entry.Amount, "ATX")
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Amount", Msg: msg}
		}

		// TransactionCode must be either 24 or 34 for Acknowledgement Entries
		switch entry.TransactionCode {
		// Prenote credit  23, 33, 43, 53
		// Prenote debit 28, 38, 48
		case 24, 34:
		default:
			msg := fmt.Sprintf(msgBatchTransactionCode, entry.TransactionCode, "ATX")
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TransactionCode", Msg: msg}
		}

		// Trapping this error, as entry.ATXAddendaRecordsField() can not be greater than 9999
		if len(entry.Addenda05) > 9999 {
			msg := fmt.Sprintf(msgBatchAddendaCount, len(entry.Addenda05), 9999, batch.Header.StandardEntryClassCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "AddendaCount", Msg: msg}
		}

		// validate ATXAddendaRecord Field is equal to the actual number of Addenda records
		// use 0 value if there is no Addenda records
		addendaRecords, _ := strconv.Atoi(entry.CATXAddendaRecordsField())
		if len(entry.Addenda05) != addendaRecords {
			msg := fmt.Sprintf(msgBatchATXAddendaCount, addendaRecords, len(entry.Addenda05))
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "AddendaCount", Msg: msg}
		}

		// Verify Addenda* FieldInclusion based on entry.Category and batchHeader.StandardEntryClassCode
		if err := batch.addendaFieldInclusion(entry); err != nil {
			return err
		}
	}
	return nil
}

// Create takes Batch Header and Entries and builds a valid batch
func (batch *BatchATX) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...
	return batch.Validate()
}
