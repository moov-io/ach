// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "fmt"

// BatchPPD holds the Batch Header and Batch Control and all Entry Records for PPD Entries
type BatchPPD struct {
	batch
}

// NewBatchPPD returns a *BatchPPD
func NewBatchPPD(bh *BatchHeader) *BatchPPD {
	batch := new(BatchPPD)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate checks valid NACHA batch rules. Assumes properly parsed records.
func (batch *BatchPPD) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration and type specific validation for this type.

	if batch.Header.StandardEntryClassCode != "PPD" {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, "PPD")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	for _, entry := range batch.Entries {
		// PPD can have up to one Record TypeCode = 05, or there can be a NOC (98) or Return (99)
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

// Create takes Batch Header and Entries and builds a valid batch
func (batch *BatchPPD) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...

	return batch.Validate()
}
