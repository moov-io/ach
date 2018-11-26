// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
)

// BatchACK is a batch file that handles SEC payment type ACK and ACK+.
// Acknowledgement of a Corporate credit by the Receiving Depository Financial Institution (RDFI).
// For commercial accounts only.
type BatchACK struct {
	Batch
}

// NewBatchACK returns a *BatchACK
func NewBatchACK(bh *BatchHeader) *BatchACK {
	batch := new(BatchACK)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate ensures the batch meets NACHA rules specific to this batch type.
func (batch *BatchACK) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration and type specific validation.
	if batch.Header.StandardEntryClassCode != "ACK" {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, "ACK")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}
	// Range through Entries
	for _, entry := range batch.Entries {
		// Amount must be zero for Acknowledgement Entries
		if entry.Amount > 0 {
			msg := fmt.Sprintf(msgBatchAmountZero, entry.Amount, "ACK")
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Amount", Msg: msg}
		}
		if len(entry.Addenda05) > 1 {
			msg := fmt.Sprintf(msgBatchAddendaCount, len(entry.Addenda05), 1, batch.Header.StandardEntryClassCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "AddendaCount", Msg: msg}
		}
		// TransactionCode must be either 24 or 34 for Acknowledgement Entries
		switch entry.TransactionCode {
		case 24, 34:
		default:
			msg := fmt.Sprintf(msgBatchTransactionCode, entry.TransactionCode, "ACK")
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TransactionCode", Msg: msg}
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

// Create builds the batch sequence numbers and batch control. Additional creation
func (batch *BatchACK) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	return batch.Validate()
}
