// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
)

// BatchCCD is a batch file that handles SEC payment type CCD amd CCD+.
// Corporate credit or debit. Identifies an Entry initiated by an Organization to transfer funds to or from an account of that Organization or another Organization.
// For commercial accounts only.
type BatchCCD struct {
	batch
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
		// CCD can have up to one Record TypeCode = 05, or there can be a NOC (98) or Return (99)
		for _, addenda := range entry.Addendum {
			switch entry.Category {
			case CategoryForward:
				if addenda.typeCode() != "05" {
					msg := fmt.Sprintf(msgBatchTypeCode, addenda.typeCode(), "05", entry.Category, batch.Header.StandardEntryClassCode)
					return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TypeCode", Msg: msg}
				}
				if len(entry.Addendum) > 1 {
					msg := fmt.Sprintf(msgBatchAddendaCount, len(entry.Addendum), 0, batch.Header.StandardEntryClassCode)
					return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "AddendaCount", Msg: msg}
				}
			case CategoryNOC:
				if addenda.typeCode() != "98" {
					msg := fmt.Sprintf(msgBatchTypeCode, addenda.typeCode(), "98", entry.Category, batch.Header.StandardEntryClassCode)
					return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TypeCode", Msg: msg}
				}
				// Do not need a length check on entry.Addendum as addAddenda.EntryDetail only allows one Addenda98
			case CategoryReturn:
				if addenda.typeCode() != "99" {
					msg := fmt.Sprintf(msgBatchTypeCode, addenda.typeCode(), "99", entry.Category, batch.Header.StandardEntryClassCode)
					return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TypeCode", Msg: msg}
				}
				// Do not need a length check on entry.Addendum as addAddenda.EntryDetail only allows one Addenda99
			}
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
