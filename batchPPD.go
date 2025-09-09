// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
package ach

import (
	"fmt"
)

// BatchPPD holds the Batch Header and Batch Control and all Entry Records for PPD Entries
type BatchPPD struct {
	Batch
}

// NewBatchPPD returns a *BatchPPD
func NewBatchPPD(bh *BatchHeader) *BatchPPD {
	batch := new(BatchPPD)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	batch.SetID(bh.ID)
	return batch
}

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *BatchPPD) Validate() error {
	if batch.validateOpts != nil && (batch.validateOpts.SkipAll || batch.validateOpts.BypassBatchValidation) {
		return nil
	}

	// Basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}

	// PPD-specific header check.
	if batch.Header.StandardEntryClassCode != PPD {
		return batch.Error("StandardEntryClassCode", ErrBatchSECType, PPD)
	}

	// Surface any previously-detected invalid entries.
	if invalid := batch.InvalidEntries(); len(invalid) > 0 {
		return invalid[0].Error
	}

	// --- Addenda05 validation (explicit) ---
	// Enforce: PaymentRelatedInformation must be <= 80 characters per NACHA.
	for _, ed := range batch.Entries { // or batch.GetEntries() if your type uses an accessor
		if ed == nil || ed.AddendaRecordIndicator != 1 || len(ed.Addenda05) == 0 {
			continue
		}
		for _, a := range ed.Addenda05 {
			if a == nil {
				continue
			}
			// hard limit to ensure we catch invalid data even if Create() renumbers fields, etc.
			if len(a.PaymentRelatedInformation) > 80 {
				return fmt.Errorf("Addenda05: PaymentRelatedInformation exceeds 80 characters")
			}
			// still run the type's own validation in case it flags anything else
			if err := a.Validate(); err != nil {
				return err
			}
		}
	}
	// --------------------------------------

	return nil
}

// InvalidEntries returns entries with validation errors in the batch
func (batch *BatchPPD) InvalidEntries() []InvalidEntry {
	var out []InvalidEntry

	for _, entry := range batch.Entries {
		// PPD can have up to one Addenda05 record
		if len(entry.Addenda05) > 1 {
			out = append(out, InvalidEntry{
				Entry: entry,
				Error: batch.Error("AddendaCount", NewErrBatchAddendaCount(len(entry.Addenda05), 1)),
			})
		}
		// Verify the Amount is valid for SEC code and TransactionCode
		if err := batch.ValidAmountForCodes(entry); err != nil {
			out = append(out, InvalidEntry{
				Entry: entry,
				Error: err,
			})
		}
		// Verify the TransactionCode is valid for a ServiceClassCode
		if err := batch.ValidTranCodeForServiceClassCode(entry); err != nil {
			out = append(out, InvalidEntry{
				Entry: entry,
				Error: err,
			})
		}
		// Verify Addenda* FieldInclusion based on entry.Category and batchHeader.StandardEntryClassCode
		if err := batch.addendaFieldInclusion(entry); err != nil {
			out = append(out, InvalidEntry{
				Entry: entry,
				Error: err,
			})
		}
	}

	return out
}

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate function at the end of their execution.
func (batch *BatchPPD) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...

	return batch.Validate()
}
