// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
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
	return batch
}

// Validate checks valid NACHA batch rules. Assumes properly parsed records.
func (batch *BatchSHR) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}

	// Add configuration and type specific validation for this type.
	if batch.Header.StandardEntryClassCode != "SHR" {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, "SHR")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	// SHR detail entries can only be a debit, ServiceClassCode must allow debits
	switch batch.Header.ServiceClassCode {
	case 200, 220:
		msg := fmt.Sprintf(msgBatchServiceClassCode, batch.Header.ServiceClassCode, "SHR")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "ServiceClassCode", Msg: msg}
	}

	for _, entry := range batch.Entries {
		// SHR detail entries must be a debit
		if entry.CreditOrDebit() != "D" {
			msg := fmt.Sprintf(msgBatchTransactionCodeCredit, entry.TransactionCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TransactionCode", Msg: msg}
		}
		if err := entry.isCardTransactionType(entry.DiscretionaryData); err != nil {
			msg := fmt.Sprintf(msgBatchCardTransactionType, entry.DiscretionaryData)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "CardTransactionType", Msg: msg}
		}

		// CardExpirationDate BatchSHR ACH File format is MMYY.  Validate MM is 01-12.
		month := entry.parseStringField(entry.SHRCardExpirationDateField()[0:2])
		year := entry.parseStringField(entry.SHRCardExpirationDateField()[2:4])
		if err := entry.isMonth(month); err != nil {
			return &FieldError{FieldName: "CardExpirationDate", Value: month, Msg: msgValidMonth}
		}
		if err := entry.isYear(year); err != nil {
			return &FieldError{FieldName: "CardExpirationDate", Value: year, Msg: msgValidYear}
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
				msg := fmt.Sprintf("%q is not a valid US state or territory", entry.Addenda02.TerminalState)
				return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TerminalState", Msg: msg}
			}
		}
	}
	return nil
}

// Create takes Batch Header and Entries and builds a valid batch
func (batch *BatchSHR) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...
	return batch.Validate()
}
