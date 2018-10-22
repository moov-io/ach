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
	// Add configuration based validation for this type.
	// Web can have up to one addenda per entry record
	if err := batch.isAddendaCount(1); err != nil {
		return err
	}
	for _, entry := range batch.Entries {
		for _, addenda := range entry.Addendum {
			if (addenda.typeCode() != "05") && (addenda.typeCode() != "99") {
				msg := fmt.Sprintf(msgBatchTypeCode, addenda.typeCode(), addenda.typeCode(), batch.Header.StandardEntryClassCode)
				return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TypeCode", Msg: msg}
			}
		}
	}

	// Add type specific validation.
	if batch.Header.StandardEntryClassCode != "WEB" {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, "WEB")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	return batch.isPaymentTypeCode()
}

// Create builds the batch sequence numbers and batch control. Additional creation
func (batch *BatchWEB) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}

	return batch.Validate()
}
