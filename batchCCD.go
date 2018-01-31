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
func NewBatchCCD(params ...BatchParam) *BatchCCD {
	batch := new(BatchCCD)
	batch.SetControl(NewBatchControl())

	if len(params) > 0 {
		bh := NewBatchHeader(params[0])
		bh.StandardEntryClassCode = "CCD"
		batch.SetHeader(bh)
		return batch
	}
	bh := NewBatchHeader()
	bh.StandardEntryClassCode = "CCD"
	batch.SetHeader(bh)
	return batch
}

// Validate ensures the batch meets NACHA rules specific to this batch type.
func (batch *BatchCCD) Validate() error {
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
	if batch.header.StandardEntryClassCode != "CCD" {
		msg := fmt.Sprintf(msgBatchSECType, batch.header.StandardEntryClassCode, "CCD")
		return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	return nil
}

// Create builds the batch sequence numbers and batch control. Additional creation
func (batch *BatchCCD) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}

	if err := batch.Validate(); err != nil {
		return err
	}
	return nil
}
