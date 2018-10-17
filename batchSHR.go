// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"unicode/utf8"
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
	batch
}

var msgBatchSHRAddenda = "found and 1 Addenda02 is required for SEC code SHR"
var msgBatchSHRAddendaType = "%T found where Addenda02 is required for SEC code SHR"

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
	// Add configuration based validation for this type.

	// Add type specific validation.

	if batch.Header.StandardEntryClassCode != "SHR" {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, "SHR")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	// SHR detail entries can only be a debit, ServiceClassCode must allow debits
	switch batch.Header.ServiceClassCode {
	case 200, 220, 280:
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
		if err := entry.isMonth(entry.parseStringField(entry.SHRCardExpirationDateField()[0:2])); err != nil {
			return &FieldError{FieldName: "CardExpirationDate", Value: entry.parseStringField(entry.SHRCardExpirationDateField()[0:2]), Msg: msgValidMonth}
		}

		if v := entry.SHRCardExpirationDateField(); utf8.RuneCountInString(v) == 4 {
			exp := entry.parseStringField(v[2:4])
			if err := entry.isYear(exp); err != nil {
				return &FieldError{
					FieldName: "CardExpirationDate",
					Value: v,
					Msg: msgValidYear,
				}
			}
		} else {
			return &FieldError{
				FieldName: "CardExpirationDate",
				Value: v,
				Msg: msgFieldInclusion,
			}
		}

		// Addenda validations - SHR Addenda must be Addenda02

		// Addendum must be equal to 1
		if len(entry.Addendum) != 1 {
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addendum", Msg: msgBatchSHRAddenda}
		}

		// Addenda type assertion must be Addenda02
		addenda02, ok := entry.Addendum[0].(*Addenda02)
		if !ok {
			msg := fmt.Sprintf(msgBatchSHRAddendaType, entry.Addendum[0])
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addendum", Msg: msg}
		}

		// Addenda02 must be Validated
		if err := addenda02.Validate(); err != nil {
			// convert the field error in to a batch error for a consistent api
			if e, ok := err.(*FieldError); ok {
				return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: e.FieldName, Msg: e.Msg}
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
