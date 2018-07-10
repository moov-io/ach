// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

// ToDo: Deprecate

// BatchIAT holds the Batch Header and Batch Control and all Entry Records for IAT Entries
type BatchIAT struct {
	IATBatch
}

// NewBatchIAT returns a *BatchIAT
func NewBatchIAT(bh *IATBatchHeader) *BatchIAT {
	iatBatch := new(BatchIAT)
	iatBatch.SetControl(NewBatchControl())
	iatBatch.SetHeader(bh)
	return iatBatch
}

// Validate checks valid NACHA batch rules. Assumes properly parsed records.
func (batch *BatchIAT) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration based validation for this type.

	// Batch can have one addenda per entry record
	/*	if err := batch.isAddendaCount(1); err != nil {
			return err
		}
		if err := batch.isTypeCode("05"); err != nil {
			return err
		}*/

	// Add type specific validation.
	// ...
	return nil
}

// Create takes Batch Header and Entries and builds a valid batch
func (batch *BatchIAT) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...

	return batch.Validate()
}
