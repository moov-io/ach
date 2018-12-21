// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
)

// BatchADV holds the Batch Header and Batch Control and all Entry Records for ADV Entries
//
// The ADV entry identifies a Non-Monetary Entry that is used by an ACH Operator to provide accounting information
// regarding an entry to participating DFI's.  It's an optional service provided by ACH operators and must be requested
// by a DFI wanting the service.
type BatchADV struct {
	Batch
}

// NewBatchADV returns a *BatchADV
func NewBatchADV(bh *BatchHeader) *BatchADV {
	batch := new(BatchADV)
	batch.SetADVControl(NewADVBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *BatchADV) Validate() error {

	if batch.Header.StandardEntryClassCode != ADV {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, ADV)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}
	if batch.Header.ServiceClassCode != AutomatedAccountingAdvices {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.ServiceClassCode, ADV)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "ServiceClassCode", Msg: msg}
	}
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration and type specific validation for this type.
	for _, entry := range batch.ADVEntries {

		if entry.Category == CategoryForward {
			switch entry.TransactionCode {
			case CreditForDebitsOriginated, CreditForCreditsReceived, CreditForCreditsRejected, CreditSummary,
				DebitForCreditsOriginated, DebitForDebitsReceived, DebitForDebitsRejectedBatches, DebitSummary:
			default:
				msg := fmt.Sprintf(msgBatchTransactionCode, entry.TransactionCode, ADV)
				return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TransactionCode", Msg: msg}
			}
		}
	}
	return nil
}

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate() function at the end of their execution.
func (batch *BatchADV) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...
	return batch.Validate()
}
