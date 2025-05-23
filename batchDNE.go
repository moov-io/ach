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
	"fmt"
	"strings"
	"time"
)

// BatchDNE is a batch file that handles SEC code Death Notification Entry (DNE)
// United States Federal agencies (e.g. Social Security) use this to notify depository
// financial institutions that the recipient of government benefit payments has died.
//
// Notes:
//   - Date of death always in positions 18-23
//   - SSN (positions 38-46) are zero if no SSN
//   - Beneficiary payment starts at position 55
type BatchDNE struct {
	Batch
}

// NewBatchDNE returns a *BatchDNE
func NewBatchDNE(bh *BatchHeader) *BatchDNE {
	batch := new(BatchDNE)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	batch.SetID(bh.ID)
	return batch
}

// Validate ensures the batch meets NACHA rules specific to this batch type.
func (batch *BatchDNE) Validate() error {
	if batch.validateOpts != nil && batch.validateOpts.SkipAll {
		return nil
	}

	if err := batch.verify(); err != nil {
		return err
	}

	// SEC code
	if batch.Header.StandardEntryClassCode != DNE {
		return batch.Error("StandardEntryClassCode", ErrBatchSECType, DNE)
	}

	invalidEntries := batch.InvalidEntries()
	if len(invalidEntries) > 0 {
		return invalidEntries[0].Error // return the first invalid entry's error
	}

	return nil
}

// InvalidEntries returns entries with validation errors in the batch
func (batch *BatchDNE) InvalidEntries() []InvalidEntry {
	var out []InvalidEntry

	// Range over Entries
	for _, entry := range batch.Entries {
		if entry.Amount != 0 {
			out = append(out, InvalidEntry{
				Entry: entry,
				Error: batch.Error("Amount", ErrBatchAmountNonZero, entry.Amount),
			})
		}

		switch entry.TransactionCode {
		case CheckingPrenoteCredit, SavingsPrenoteCredit:
		default:
			out = append(out, InvalidEntry{
				Entry: entry,
				Error: batch.Error("TransactionCode", ErrBatchTransactionCode, entry.TransactionCode),
			})
		}

		// DNE must have one Addenda05
		if len(entry.Addenda05) != 1 {
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
func (batch *BatchDNE) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	return batch.Validate()
}

type DNEPaymentInformation struct {
	DateOfDeath time.Time
	CustomerSSN string

	// Amount is a two-decimal float value formatted as a string
	// Example: 123.45
	Amount string
}

// ParseDNEPaymentInformation returns an DNEPaymentInformation for a given Addenda05 record. The information is parsed from the addenda's
// PaymentRelatedInformation field.
//
// The returned information is not validated for correctness.
func ParseDNEPaymentInformation(addenda05 *Addenda05) (*DNEPaymentInformation, error) {
	if addenda05 == nil {
		return nil, nil
	}

	fields := strings.Split(strings.TrimSuffix(addenda05.PaymentRelatedInformation, `\`), "*")
	if len(fields) != 6 {
		return nil, fmt.Errorf("unexpected %d fields", len(fields))
	}

	dateOfDeath, err := time.Parse("010206", fields[1])
	if err != nil {
		return nil, fmt.Errorf("parsing DateOfDeath: %w", err)
	}

	return &DNEPaymentInformation{
		DateOfDeath: dateOfDeath,
		CustomerSSN: fields[3],
		Amount:      fields[5],
	}, nil
}

func (info DNEPaymentInformation) String() string {
	return fmt.Sprintf(`DATE OF DEATH*%s*CUSTOMER SSN*%s*AMOUNT*%s\`,
		info.DateOfDeath.Format("010206"), info.CustomerSSN, info.Amount)
}
