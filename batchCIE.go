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

// BatchCIE holds the BatchHeader and BatchControl and all EntryDetail for CIE Entries.
//
// Customer-Initiated Entry (or CIE entry) is a credit entry initiated on behalf of,
// and upon the instruction of, a consumer to transfer funds to a non-consumer Receiver.
// CIE entries are usually transmitted to a company for payment of funds that the consumer
// owes to that company and are initiated by the consumer through some type of online
// banking product or bill payment service provider. With CIEs, funds owed by the consumer
// are “pushed” to the biller in the form of an ACH credit, as opposed to the biller's use of
// a debit application (e.g., PPD, WEB) to “pull” the funds from a customer's account.
type BatchCIE struct {
	Batch
}

// NewBatchCIE returns a *BatchCIE
func NewBatchCIE(bh *BatchHeader) *BatchCIE {
	batch := new(BatchCIE)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	batch.SetID(bh.ID)
	return batch
}

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *BatchCIE) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration based validation for this type.

	// Add type specific validation.

	if batch.Header.StandardEntryClassCode != CIE {
		return batch.Error("StandardEntryClassCode", ErrBatchSECType, CIE)
	}

	// CIE detail entries can only be a credit, ServiceClassCode must allow credit
	switch batch.Header.ServiceClassCode {
	case DebitsOnly:
		return batch.Error("ServiceClassCode", ErrBatchServiceClassCode, batch.Header.ServiceClassCode)
	}

	for _, entry := range batch.Entries {
		// CIE detail entries must be a debit
		if entry.CreditOrDebit() != "C" {
			return batch.Error("TransactionCode", ErrBatchDebitOnly, entry.TransactionCode)
		}
		// CIE must have a maximum of one Addenda05 record
		if len(entry.Addenda05) > 1 {
			return batch.Error("AddendaCount", NewErrBatchAddendaCount(len(entry.Addenda05), 1))
		}
		// Verify the Amount is valid for SEC code and TransactionCode
		if err := batch.ValidAmountForCodes(entry); err != nil {
			return err
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
func (batch *BatchCIE) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...
	return batch.Validate()
}
