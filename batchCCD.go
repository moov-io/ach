// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
)

// BatchCCD is a batch file that handles SEC payment type CCD and CCD+.
// Corporate credit or debit. Identifies an Entry initiated by an Organization to transfer funds to or from an account of that Organization or another Organization.
// For commercial accounts only.
type BatchCCD struct {
	Batch
}

// NewBatchCCD returns a *BatchCCD
func NewBatchCCD(bh *BatchHeader) *BatchCCD {
	batch := new(BatchCCD)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate ensures the batch meets NACHA rules specific to this batch type.
func (batch *BatchCCD) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}

	// Add configuration and type specific validation.
	if batch.Header.StandardEntryClassCode != "CCD" {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, "CCD")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	for _, entry := range batch.Entries {
		// CCD can have up to one Addenda05 record,
		if len(entry.Addenda05) > 1 {
			msg := fmt.Sprintf(msgBatchAddendaCount, len(entry.Addenda05), 1, batch.Header.StandardEntryClassCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "AddendaCount", Msg: msg}
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
func (batch *BatchCCD) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}

	return batch.Validate()
}
