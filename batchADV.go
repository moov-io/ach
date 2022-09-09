// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package ach

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
	batch.SetID(bh.ID)
	return batch
}

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *BatchADV) Validate() error {
	if batch.Header.StandardEntryClassCode != ADV {
		return batch.Error("StandardEntryClassCode", ErrBatchSECType, ADV)
	}
	if batch.Header.ServiceClassCode != AutomatedAccountingAdvices {
		return batch.Error("ServiceClassCode", ErrBatchServiceClassCode, batch.Header.ServiceClassCode)
	}
	if batch.Header.OriginatorStatusCode != 0 {
		return batch.Error("OriginatorStatusCode", ErrOrigStatusCode, batch.Header.OriginatorStatusCode)
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
				return batch.Error("TransactionCode", ErrBatchTransactionCode, entry.TransactionCode)
			}
			if entry.Addenda99 != nil {
				return batch.Error("Addenda99", ErrBatchAddendaCategory, entry.Category)
			}
		}
	}
	return nil
}

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate function at the end of their execution.
func (batch *BatchADV) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...
	return batch.Validate()
}
