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

import (
	"github.com/moov-io/ach/internal/usabbrev"
)

// BatchPOS holds the BatchHeader and BatchControl and all EntryDetail for POS Entries.
//
// A POS Entry is a debit Entry initiated at an “electronic terminal” to a consumer
// account of the receiver to pay an obligation incurred in a point- of-sale
// transaction, or to effect a point-of-sale terminal cash withdrawal.
//
// Point-of-Sale Entries (POS) are ACH debit entries typically initiated by the use
// of a merchant-issued plastic card to pay an obligation at the point-of-sale. Much
// like a financial institution issued debit card, the merchant- issued debit card is
// swiped at the point-of-sale and approved for use; however, the authorization only
// verifies the card is open, active and within the card's limits—it does not verify
// the Receiver's account balance or debit the account at the time of the purchase.
// Settlement of the transaction moves from the card network to the ACH Network through
// the creation of a POS entry by the card issuer to debit the Receiver's account.
type BatchPOS struct {
	Batch
}

// NewBatchPOS returns a *BatchPOS
func NewBatchPOS(bh *BatchHeader) *BatchPOS {
	batch := new(BatchPOS)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	batch.SetID(bh.ID)
	return batch
}

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *BatchPOS) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}

	// Add configuration and type specific validation for this type.

	if batch.Header.StandardEntryClassCode != POS {
		return batch.Error("StandardEntryClassCode", ErrBatchSECType, POS)
	}

	for _, entry := range batch.Entries {
		if err := entry.isCardTransactionType(entry.DiscretionaryData); err != nil {
			return batch.Error("CardTransactionType", ErrBatchInvalidCardTransactionType, entry.DiscretionaryData)
		}
		// Verify the TransactionCode is valid for a ServiceClassCode
		if err := batch.ValidTranCodeForServiceClassCode(entry); err != nil {
			return err
		}
		// Verify Addenda* FieldInclusion based on entry.Category and batchHeader.StandardEntryClassCode
		if err := batch.addendaFieldInclusion(entry); err != nil {
			return err
		}
		if entry.Category == CategoryForward {
			if !usabbrev.Valid(entry.Addenda02.TerminalState) {
				return batch.Error("TerminalState", ErrValidState, entry.Addenda02.TerminalState)
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
func (batch *BatchPOS) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...
	return batch.Validate()
}
