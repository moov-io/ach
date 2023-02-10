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

// BatchBOC holds the BatchHeader and BatchControl and all EntryDetail for BOC Entries.
//
// Back Office Conversion (BOC) A single entry debit initiated at the point of purchase
// or at a manned bill payment location to transfer funds through conversion to an
// ACH debit entry during back office processing.
//
// BOC allows retailers/billers, and ODFIs acting as Originators,
// to electronically convert checks received at the point-of-purchase as well as at a
// manned bill payment location into a single-entry ACH debit. The authorization to
// convert the check will be obtained through a notice at the checkout or manned bill
// payment location (e.g., loan payment at financial institution's teller window) and the
// receipt of the Receiver's check. The decision to process the check item as an ACH debit
// will be made in the “back office” instead of at the point-of-purchase. The customer's
// check will solely be used as a source document to obtain the routing number, account
// number and check serial number.
//
// Unlike ARC entries, BOC conversions require the customer to be present and a notice that
// checks may be converted to BOC ACH entries be posted.
type BatchBOC struct {
	Batch
}

// NewBatchBOC returns a *BatchBOC
func NewBatchBOC(bh *BatchHeader) *BatchBOC {
	batch := new(BatchBOC)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	batch.SetID(bh.ID)
	return batch
}

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *BatchBOC) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration and type specific validation for this type.

	if batch.Header.StandardEntryClassCode != BOC {
		return batch.Error("StandardEntryClassCode", ErrBatchSECType, BOC)
	}

	// BOC detail entries can only be a debit, ServiceClassCode must allow debits
	switch batch.Header.ServiceClassCode {
	case CreditsOnly:
		return batch.Error("ServiceClassCode", ErrBatchServiceClassCode, batch.Header.ServiceClassCode)
	}

	for _, entry := range batch.Entries {
		// BOC detail entries must be a debit
		if entry.CreditOrDebit() != "D" {
			return batch.Error("TransactionCode", ErrBatchDebitOnly, entry.TransactionCode)
		}

		// Amount must be 25,000 or less
		if entry.Amount > 2500000 {
			return batch.Error("Amount", NewErrBatchAmount(entry.Amount, 2500000))
		}

		// CheckSerialNumber underlying IdentificationNumber, must be defined
		if entry.IdentificationNumber == "" {
			return batch.Error("CheckSerialNumber", ErrBatchCheckSerialNumber)
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
func (batch *BatchBOC) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...

	return batch.Validate()
}
