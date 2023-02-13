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

// BatchACK is a batch file that handles SEC payment type ACK and ACK+.
// Acknowledgement of a Corporate credit by the Receiving Depository Financial Institution (RDFI).
// For commercial accounts only.
type BatchACK struct {
	Batch
}

// NewBatchACK returns a *BatchACK
func NewBatchACK(bh *BatchHeader) *BatchACK {
	batch := new(BatchACK)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	batch.SetID(bh.ID)
	return batch
}

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *BatchACK) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration and type specific validation.
	if batch.Header.StandardEntryClassCode != ACK {
		return batch.Error("StandardEntryClassCode", ErrBatchSECType, ACK)
	}
	// Range through Entries
	for _, entry := range batch.Entries {
		// Amount must be zero for Acknowledgement Entries
		if entry.Amount > 0 {
			return batch.Error("Amount", ErrBatchAmountNonZero, entry.Amount)
		}
		if len(entry.Addenda05) > 1 {
			return batch.Error("AddendaCount", NewErrBatchAddendaCount(len(entry.Addenda05), 1))
		}
		switch entry.TransactionCode {
		case CheckingZeroDollarRemittanceCredit, SavingsZeroDollarRemittanceCredit:
		default:
			return batch.Error("TransactionCode", ErrBatchTransactionCode, entry.TransactionCode)
		}
		// // Verify the Amount is valid for SEC code and TransactionCode
		// if err := batch.ValidAmountForCodes(entry); err != nil { // TODO(adam): https://github.com/moov-io/ach/issues/1170
		// 	return err
		// }
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
func (batch *BatchACK) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	return batch.Validate()
}
