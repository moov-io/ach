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
	batch
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
		// TransactionCode must be either 24 or 34 for Acknowledgement Entries
		switch entry.TransactionCode {
		case 24, 34:
		default:
			msg := fmt.Sprintf(msgBatchTransactionCode, entry.TransactionCode, "ACK")
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TransactionCode", Msg: msg}
		}

		// ACK can have up to one Record TypeCode = 05, or there can be a NOC (98) or Return (99)
		for _, addenda := range entry.Addendum {
			switch entry.Category {
			case CategoryForward:
				if err := batch.categoryForwardAddenda05(entry, addenda); err != nil {
					return err
				}
				if len(entry.Addendum) > 1 {
					msg := fmt.Sprintf(msgBatchAddendaCount, len(entry.Addendum), 1, batch.Header.StandardEntryClassCode)
					return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "AddendaCount", Msg: msg}
				}
			case CategoryNOC:
				if err := batch.categoryNOCAddenda98(entry, addenda); err != nil {
					return err
				}
			case CategoryReturn:
				if err := batch.categoryReturnAddenda99(entry, addenda); err != nil {
					return err
				}
			}
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
