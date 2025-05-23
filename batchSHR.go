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

// BatchSHR holds the BatchHeader and BatchControl and all EntryDetail for SHR Entries.
//
// Shared Network Entry (SHR) is a debit Entry initiated at an “electronic terminal,”
// as that term is defined in Regulation E, to a Consumer Account of the Receiver to pay
// an obligation incurred in a point-of-sale transaction, or to effect a point-of-sale
// terminal cash withdrawal. Also an adjusting or other credit Entry related to such debit
// Entry, transfer of funds, or obligation. SHR Entries are initiated in a shared network
// where the ODFI and RDFI have an agreement in addition to these Rules to process such
// Entries.
type BatchSHR struct {
	Batch
}

// NewBatchSHR returns a *BatchSHR
func NewBatchSHR(bh *BatchHeader) *BatchSHR {
	batch := new(BatchSHR)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	batch.SetID(bh.ID)
	return batch
}

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *BatchSHR) Validate() error {
	if batch.validateOpts != nil && batch.validateOpts.SkipAll {
		return nil
	}

	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}

	// Add configuration and type specific validation for this type.
	if batch.Header.StandardEntryClassCode != SHR {
		return batch.Error("StandardEntryClassCode", ErrBatchSECType, SHR)
	}

	// SHR entries can be debit, credit or mixed ServiceClassCode
	switch batch.Header.ServiceClassCode {
	case MixedDebitsAndCredits, CreditsOnly, DebitsOnly:
		// do nothing
	default:
		return batch.Error("ServiceClassCode", ErrBatchServiceClassCode, batch.Header.ServiceClassCode)
	}

	invalidEntries := batch.InvalidEntries()
	if len(invalidEntries) > 0 {
		return invalidEntries[0].Error // return the first invalid entry's error
	}

	return nil
}

// InvalidEntries returns entries with validation errors in the batch
func (batch *BatchSHR) InvalidEntries() []InvalidEntry {
	var out []InvalidEntry

	for _, entry := range batch.Entries {
		// SHR detail entries can be debit or credit
		switch entry.CreditOrDebit() {
		case "C", "D":
			// do nothing
		default:
			out = append(out, InvalidEntry{
				Entry: entry,
				Error: batch.Error("TransactionCode", ErrBatchTransactionCode, entry.TransactionCode),
			})
		}
		if err := entry.isCardTransactionType(entry.DiscretionaryData); err != nil {
			out = append(out, InvalidEntry{
				Entry: entry,
				Error: batch.Error("CardTransactionType", ErrBatchInvalidCardTransactionType, entry.DiscretionaryData),
			})
		}

		// CardExpirationDate BatchSHR ACH File format is MMYY.  Validate MM is 01-12.
		month := entry.parseStringField(entry.SHRCardExpirationDateField()[0:2])
		year := entry.parseStringField(entry.SHRCardExpirationDateField()[2:4])
		if err := entry.isMonth(month); err != nil {
			out = append(out, InvalidEntry{
				Entry: entry,
				Error: fieldError("CardExpirationDate", ErrValidMonth, month),
			})
		}
		if err := entry.isCreditCardYear(year); err != nil {
			out = append(out, InvalidEntry{
				Entry: entry,
				Error: fieldError("CardExpirationDate", ErrValidYear, year),
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
		if entry.Category == CategoryForward {
			if entry.Addenda02 != nil {
				if !usabbrev.Valid(entry.Addenda02.TerminalState) {
					out = append(out, InvalidEntry{
						Entry: entry,
						Error: batch.Error("TerminalState", ErrValidState, entry.Addenda02.TerminalState),
					})
				}
			}
		}
	}

	return out
}

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate function at the end of their execution.
func (batch *BatchSHR) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...
	return batch.Validate()
}
