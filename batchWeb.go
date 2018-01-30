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
func NewBatchWEB(params ...BatchParam) *BatchWEB {
	batch := new(BatchWEB)
	batch.SetControl(NewBatchControl())

	if len(params) > 0 {
		bh := NewBatchHeader(params[0])
		bh.StandardEntryClassCode = "WEB"
		batch.SetHeader(bh)
		return batch
	}
	bh := NewBatchHeader()
	bh.StandardEntryClassCode = "WEB"
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
	if err := batch.isTypeCode("05"); err != nil {
		return err
	}

	// Add type specific validation.
	if batch.header.StandardEntryClassCode != "WEB" {
		msg := fmt.Sprintf(msgBatchSECType, batch.header.StandardEntryClassCode, "WEB")
		return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	if err := batch.isPaymentTypeCode(); err != nil {
		return err
	}

	return nil
}

// Create builds the batch sequence numbers and batch control. Additional creation
func (batch *BatchWEB) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}

	if err := batch.Validate(); err != nil {
		return err
	}
	return nil
}
