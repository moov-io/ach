// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

// BatchPPD holds the Batch Header and Batch Control and all Entry Records for PPD Entries
type BatchPPD struct {
	batch
}

// NewBatchPPD returns a *BatchPPD
func NewBatchPPD(params ...BatchParam) *BatchPPD {
	batch := new(BatchPPD)
	batch.SetControl(NewBatchControl())

	if len(params) > 0 {
		bh := NewBatchHeader(params[0])
		bh.StandardEntryClassCode = ppd
		batch.SetHeader(bh)
		return batch
	}
	bh := NewBatchHeader()
	bh.StandardEntryClassCode = ppd
	batch.SetHeader(bh)
	return batch
}

// Validate checks valid NACHA batch rules. Assumes properly parsed records.
func (batch *BatchPPD) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration based validation for this type.

	// Batch can have one addenda per entry record
	if err := batch.isAddendaCount(1); err != nil {
		return err
	}
	if err := batch.isTypeCode("05"); err != nil {
		return err
	}

	// Add type specific validation.
	// ...
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

	if err := batch.Validate(); err != nil {
		return err
	}
	return nil
}
