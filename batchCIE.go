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
	batch
}

// NewBatchCIE returns a *BatchCIE
func NewBatchCIE(bh *BatchHeader) *BatchCIE {
	batch := new(BatchCIE)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate checks valid NACHA batch rules. Assumes properly parsed records.
func (batch *BatchCIE) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration based validation for this type.

	// Add type specific validation.

	if batch.Header.StandardEntryClassCode != "CIE" {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, "CIE")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	// CIE detail entries can only be a debit, ServiceClassCode must allow debits
	switch batch.Header.ServiceClassCode {
	case 200, 225, 280:
		msg := fmt.Sprintf(msgBatchServiceClassCode, batch.Header.ServiceClassCode, "CIE")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "ServiceClassCode", Msg: msg}
	}

	for _, entry := range batch.Entries {
		// CIE detail entries must be a debit
		if entry.CreditOrDebit() != "C" {
			msg := fmt.Sprintf(msgBatchTransactionCodeCredit, entry.TransactionCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TransactionCode", Msg: msg}
		}

		// CIE must have one Addenda05 record
		if len(entry.Addendum) != 1 {
			msg := fmt.Sprintf(msgBatchRequiredAddendaCount, len(entry.Addendum), 1, batch.Header.StandardEntryClassCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "AddendaCount", Msg: msg}
		}

		// CIE can have up to one Record TypeCode = 05, or there can be a NOC (98) or Return (99)
		for _, addenda := range entry.Addendum {
			switch entry.Category {
			case CategoryForward:
				if addenda.typeCode() != "05" {
					msg := fmt.Sprintf(msgBatchTypeCode, addenda.typeCode(), "05", entry.Category, batch.Header.StandardEntryClassCode)
					return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TypeCode", Msg: msg}
				}
				if len(entry.Addendum) > 1 {
					msg := fmt.Sprintf(msgBatchAddendaCount, len(entry.Addendum), 0, batch.Header.StandardEntryClassCode)
					return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "AddendaCount", Msg: msg}
				}
			case CategoryNOC:
				if addenda.typeCode() != "98" {
					msg := fmt.Sprintf(msgBatchTypeCode, addenda.typeCode(), "98", entry.Category, batch.Header.StandardEntryClassCode)
					return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TypeCode", Msg: msg}
				}
				// Do not need a length check on entry.Addendum as addAddenda.EntryDetail only allows one Addenda98
			case CategoryReturn:
				if addenda.typeCode() != "99" {
					msg := fmt.Sprintf(msgBatchTypeCode, addenda.typeCode(), "99", entry.Category, batch.Header.StandardEntryClassCode)
					return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TypeCode", Msg: msg}
				}
				// Do not need a length check on entry.Addendum as addAddenda.EntryDetail only allows one Addenda99
			}
		}
	}
	return nil
}

// Create takes Batch Header and Entries and builds a valid batch
func (batch *BatchCIE) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...
	return batch.Validate()
}
