// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
)

var (
	msgBatchDNEAddenda     = "found and 1 Addenda05 is required for SEC Type DNE"
	msgBatchDNEAddendaType = "%T found where Addenda05 is required for SEC type DNE"
)

// BatchDNE is a batch file that handles SEC code Death Notification Entry (DNE)
// United States Federal agencies (e.g. Social Security) use this to notify depository
// financial institutions that the recipient of government benefit paymetns has died.
//
// Notes:
//  - Date of death always in positions 18-23
//  - SSN (positions 38-46) are zero if no SSN
//  - Beneficiary payment starts at position 55
type BatchDNE struct {
	batch
}

func NewBatchDNE(bh *BatchHeader) *BatchDNE {
	batch := new(BatchDNE)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate ensures the batch meets NACHA rules specific to this batch type.
func (batch *BatchDNE) Validate() error {
	if err := batch.verify(); err != nil {
		return err
	}

	// SEC code
	if batch.Header.StandardEntryClassCode != "DNE" {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, "DNE")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	// Range over Entries
	for _, entry := range batch.Entries {
		if entry.Amount != 0 {
			msg := fmt.Sprintf(msgBatchAmountZero, entry.Amount, "DNE")
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Amount", Msg: msg}
		}

		switch entry.TransactionCode {
		case 21, 23, 31, 33:
		default:
			msg := fmt.Sprintf(msgBatchTransactionCode, entry.TransactionCode, "DNE")
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TransactionCode", Msg: msg}
		}

		// DNE must have one Addenda05
		if len(entry.Addendum) != 1 {
			msg := fmt.Sprintf(msgBatchAddendaCount, len(entry.Addendum), 1, batch.Header.StandardEntryClassCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "AddendaCount", Msg: msg}
		}
	}

	// Check Addenda05
	return batch.isAddenda05()
}

// Create builds the batch sequence numbers and batch control. Additional creation
func (batch *BatchDNE) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	return batch.Validate()
}

// isAddenda05 verifies that a Addenda04 exists for each EntryDetail and is Validated
func (batch *BatchDNE) isAddenda05() error {
	if len(batch.Entries) != 1 {
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "entries", Msg: msgBatchEntries}
	}

	for _, entry := range batch.Entries {
		// Addenda type must be equal to 1
		if len(entry.Addendum) != 1 {
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addendum", Msg: msgBatchDNEAddenda}
		}
		// Addenda type assertion must be Addenda05
		addenda05, ok := entry.Addendum[0].(*Addenda05)
		if !ok {
			msg := fmt.Sprintf(msgBatchDNEAddendaType, entry.Addendum[0])
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
	return nil
}
