package ach

import (
	"fmt"
)

type BatchCOR struct {
	batch
}

var (
	msgBatchWebPaymentType = "%v is not a valid payment type S (single entry) or R (recurring)"
)

// NewBatchCOR returns a *BatchWEB
func NewBatchCOR(params ...BatchParam) *BatchCOR {
	batch := new(BatchCOR)
	batch.SetControl(NewBatchControl())

	if len(params) > 0 {
		bh := NewBatchHeader(params[0])
		bh.StandardEntryClassCode = "COR"
		batch.SetHeader(bh)
		return batch
	}
	bh := NewBatchHeader()
	bh.StandardEntryClassCode = "COR"
	batch.SetHeader(bh)
	return batch
}

// Validate ensures the batch meets NACHA rules specific to this batch type.
func (batch *BatchCOR) Validate() error {
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
	if batch.header.StandardEntryClassCode != "COR" {
		msg := fmt.Sprintf(msgBatchSECType, batch.header.StandardEntryClassCode, "COR")
		return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	return nil
}

// Create builds the batch sequence numbers and batch control. Additional creation
func (batch *BatchCOR) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}

	if err := batch.Validate(); err != nil {
		return err
	}
	return nil
}
