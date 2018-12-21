// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
)

// BatchCIE holds the BatchHeader and BatchControl and all EntryDetail for CIE Entries.
//
// Customer-Initiated Entry (or CIE entry) is a credit entry initiated on behalf of,
// and upon the instruction of, a consumer to transfer funds to a non-consumer Receiver.
// CIE entries are usually transmitted to a company for payment of funds that the consumer
// owes to that company and are initiated by the consumer through some type of online
// banking product or bill payment service provider. With CIEs, funds owed by the consumer
// are “pushed” to the biller in the form of an ACH credit, as opposed to the biller’s use of
// a debit application (e.g., PPD, WEB) to “pull” the funds from a customer’s account.
type BatchCIE struct {
	Batch
}

// NewBatchCIE returns a *BatchCIE
func NewBatchCIE(bh *BatchHeader) *BatchCIE {
	batch := new(BatchCIE)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
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
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, CCD)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	// CIE detail entries can only be a credit, ServiceClassCode must allow credit
	switch batch.Header.ServiceClassCode {
	case MixedDebitsAndCredits, DebitsOnly:
		msg := fmt.Sprintf(msgBatchServiceClassCode, batch.Header.ServiceClassCode, CCD)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "ServiceClassCode", Msg: msg}
	}

	for _, entry := range batch.Entries {
		// CIE detail entries must be a debit
		if entry.CreditOrDebit() != "C" {
			msg := fmt.Sprintf(msgBatchTransactionCodeCredit, entry.TransactionCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TransactionCode", Msg: msg}
		}
		// CIE must have one Addenda05 record
		if len(entry.Addenda05) != 1 {
			msg := fmt.Sprintf(msgBatchRequiredAddendaCount, len(entry.Addenda05), 1, batch.Header.StandardEntryClassCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "AddendaCount", Msg: msg}
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
// call the Batch's Validate() function at the end of their execution.
func (batch *BatchCIE) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...
	return batch.Validate()
}
