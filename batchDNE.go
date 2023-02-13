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
	"strings"
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
	if err := batch.verify(); err != nil {
		return err
	}

	// SEC code
	if batch.Header.StandardEntryClassCode != DNE {
		return batch.Error("StandardEntryClassCode", ErrBatchSECType, DNE)
	}

	// Range over Entries
	for _, entry := range batch.Entries {
		if entry.Amount != 0 {
			return batch.Error("Amount", ErrBatchAmountNonZero, entry.Amount)
		}

		switch entry.TransactionCode {
		case CheckingReturnNOCCredit, CheckingPrenoteCredit, SavingsReturnNOCCredit, SavingsPrenoteCredit:
		default:
			return batch.Error("TransactionCode", ErrBatchTransactionCode, entry.TransactionCode)
		}

		// DNE must have one Addenda05
		if len(entry.Addenda05) != 1 {
			return batch.Error("AddendaCount", NewErrBatchAddendaCount(len(entry.Addenda05), 1))
		}
		// // Verify the Amount is valid for SEC code and TransactionCode
		// if err := batch.ValidAmountForCodes(entry); err != nil { // TODO(adam): https://github.com/moov-io/ach/issues/1171
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
func (batch *BatchDNE) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	return batch.Validate()
}

// details returns the Date of Death (YYMMDD), Customer SSN (9 digits), and Amount ($$$$.cc)
// from the Addenda05 record. This method assumes the addenda05 PaymentRelatedInformation is valid.
func (batch *BatchDNE) details() (string, string, string) {
	if batch == nil || len(batch.Entries) == 0 {
		return "", "", ""
	}

	addendas := batch.Entries[0].Addenda05
	if len(addendas) != 1 {
		return "", "", ""
	}

	line := addendas[0].PaymentRelatedInformation
	return line[18:24], line[37:46], strings.TrimSuffix(line[54:], `\`)
}

// DateOfDeath returns the YYMMDD string from Addenda05's PaymentRelatedInformation
func (batch *BatchDNE) DateOfDeath() string {
	date, _, _ := batch.details()
	return date
}

// CustomerSSN returns the SSN string from Addenda05's PaymentRelatedInformation
func (batch *BatchDNE) CustomerSSN() string {
	_, ssn, _ := batch.details()
	return ssn
}

// Amount returns the amount to be dispursed to the named beneficiary from Addenda05's PaymentRelatedInformation.
func (batch *BatchDNE) Amount() string {
	_, _, amount := batch.details()
	return amount
}
