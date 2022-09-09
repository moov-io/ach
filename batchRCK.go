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

// BatchRCK holds the BatchHeader and BatchControl and all EntryDetail for RCK Entries.
//
// Represented Check Entries (RCK). A physical check that was presented but returned because of
// insufficient funds may be represented as an ACH entry.
type BatchRCK struct {
	Batch
}

// NewBatchRCK returns a *BatchRCK
func NewBatchRCK(bh *BatchHeader) *BatchRCK {
	batch := new(BatchRCK)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	batch.SetID(bh.ID)
	return batch
}

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *BatchRCK) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}

	// Add configuration and type specific validation for this type.
	if batch.Header.StandardEntryClassCode != RCK {
		return batch.Error("StandardEntryClassCode", ErrBatchSECType, RCK)
	}

	// RCK detail entries can only be a debit, ServiceClassCode must allow debits
	switch batch.Header.ServiceClassCode {
	case CreditsOnly:
		return batch.Error("ServiceClassCode", ErrBatchServiceClassCode, batch.Header.ServiceClassCode)
	}

	// CompanyEntryDescription is required to be REDEPCHECK
	if batch.Header.CompanyEntryDescription != "REDEPCHECK" {
		return batch.Error("CompanyEntryDescription", ErrBatchCompanyEntryDescriptionREDEPCHECK, batch.Header.CompanyEntryDescription)
	}

	for _, entry := range batch.Entries {
		// RCK detail entries must be a debit
		if entry.CreditOrDebit() != "D" {
			return batch.Error("TransactionCode", ErrBatchDebitOnly, entry.TransactionCode)
		}
		// // Amount must be 2,500 or less
		if entry.Amount > 250000 {
			return batch.Error("Amount", NewErrBatchAmount(entry.Amount, 250000))
		}
		// CheckSerialNumber underlying IdentificationNumber, must be defined
		if entry.IdentificationNumber == "" {
			return batch.Error("CheckSerialNumber", ErrBatchCheckSerialNumber)
		}
		// Verify the TransactionCode is valid for a ServiceClassCode
		if err := batch.ValidTranCodeForServiceClassCode(entry); err != nil {
			return err
		}
		// Verify Addenda* FieldInclusion based on entry.Category and batchHeader.StandardEntryClassCode
		if err := batch.addendaFieldInclusion(entry); err != nil {
			return err
		}
	}
	return nil
}

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate function at the end of their execution.
func (batch *BatchRCK) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...

	return batch.Validate()
}
