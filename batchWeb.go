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

var (
	msgBatchWebPaymentType = "%v is not a valid payment type S (single entry) or R (recurring)"
)

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
				if addenda.typeCode() != "05" {
					msg := fmt.Sprintf(msgBatchTypeCode, addenda.typeCode(), "05", entry.Category, batch.Header.StandardEntryClassCode)
					return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TypeCode", Msg: msg}
				}
				if len(entry.Addendum) > 1 {
					msg := fmt.Sprintf(msgBatchAddendaCount, len(entry.Addendum), 1, batch.Header.StandardEntryClassCode)
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
func (batch *BatchWEB) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}

	return batch.Validate()
}
