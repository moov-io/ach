package ach

import (
	"fmt"
)

// BatchWEB creates a batch file that handles SEC payment type WEB.
// Entry submitted pursuant to an authorization obtained solely via the Internet or a wireless network
// For consumer accounts only.
type BatchWEB struct {
	batch
}

// NewBatchWEB returns a *BatchWEB
func NewBatchWEB(bh *BatchHeader) *BatchWEB {
	batch := new(BatchWEB)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate ensures the batch meets NACHA rules specific to this batch type.
func (batch *BatchWEB) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration and type specific validation for this type.
	if batch.Header.StandardEntryClassCode != "WEB" {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, "WEB")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	for _, entry := range batch.Entries {
		// WEB can have up to one Addenda Record TypeCode = 05, or there can be a NOC (98) or Return (99)
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
func (batch *BatchWEB) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}

	return batch.Validate()
}
