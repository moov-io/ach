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

// BatchTEL is a batch that handles SEC payment type Telephone-Initiated Entries (TEL)
// Telephone-Initiated Entries (TEL) are consumer debit transactions. The NACHA Operating Rules permit TEL entries when
// the Originator obtains the Receiver's authorization for the debit entry orally via the telephone.
// An entry based upon a Receiver's oral authorization must utilize the TEL (Telephone-Initiated Entry)
// Standard Entry Class (SEC) Code.
type BatchTEL struct {
	Batch
}

// NewBatchTEL returns a *BatchTEL
func NewBatchTEL(bh *BatchHeader) *BatchTEL {
	batch := new(BatchTEL)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	batch.SetID(bh.ID)
	return batch
}

// Validate ensures the batch meets NACHA rules specific to the SEC type TEL
func (batch *BatchTEL) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration and type specific based validation for this type.
	if batch.Header.StandardEntryClassCode != TEL {
		return batch.Error("StandardEntryClassCode", ErrBatchSECType, TEL)
	}
	// can not have credits in TEL batches
	for _, entry := range batch.Entries {
		if entry.CreditOrDebit() != "D" {
			return batch.Error("TransactionCode", ErrBatchDebitOnly, entry.TransactionCode)
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
func (batch *BatchTEL) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}

	return batch.Validate()
}
