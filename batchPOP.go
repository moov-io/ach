// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "fmt"

// BatchPOP holds the BatchHeader and BatchControl and all EntryDetail for POP Entries.
//
// Point-of-Purchase. A check presented in-person to a merchant for purchase is presented
// as an ACH entry instead of a physical check.
//
// This ACH debit application is used by originators as a method of payment for the
// in-person purchase of goods or services by consumers. These Single Entry debit
// entries are initiated by the originator based on a written authorization and
// account information drawn from the source document (a check) obtained from the
// consumer at the point-of-purchase. The source document, which is voided by the
// merchant and returned to the consumer at the point-of-purchase, is used to
// collect the consumer’s routing number, account number and check serial number that
// will be used to generate the debit entry to the consumer’s account.
//
// The difference between POP and ARC is that ARC can result from a check mailed in whereas POP is in-person.
type BatchPOP struct {
	Batch
}

// NewBatchPOP returns a *BatchPOP
func NewBatchPOP(bh *BatchHeader) *BatchPOP {
	batch := new(BatchPOP)
	batch.SetControl(NewBatchControl())
	batch.SetHeader(bh)
	return batch
}

// Validate checks valid NACHA batch rules. Assumes properly parsed records.
func (batch *BatchPOP) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}

	// Add configuration and type specific validation for this type.
	if batch.Header.StandardEntryClassCode != "POP" {
		msg := fmt.Sprintf(msgBatchSECType, batch.Header.StandardEntryClassCode, "POP")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "StandardEntryClassCode", Msg: msg}
	}

	// POP detail entries can only be a debit, ServiceClassCode must allow debits
	switch batch.Header.ServiceClassCode {
	case 200, 220:
		msg := fmt.Sprintf(msgBatchServiceClassCode, batch.Header.ServiceClassCode, "POP")
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "ServiceClassCode", Msg: msg}
	}

	for _, entry := range batch.Entries {
		// POP detail entries must be a debit
		if entry.CreditOrDebit() != "D" {
			msg := fmt.Sprintf(msgBatchTransactionCodeCredit, entry.TransactionCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TransactionCode", Msg: msg}
		}
		// Amount must be 25,000 or less
		if entry.Amount > 2500000 {
			msg := fmt.Sprintf(msgBatchAmount, "25,000", "POP")
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Amount", Msg: msg}
		}
		// CheckSerialNumber, Terminal City, Terminal State underlying IdentificationNumber, must be defined
		if entry.IdentificationNumber == "" {
			msg := fmt.Sprintf(msgBatchCheckSerialNumber, "POP")
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "CheckSerialNumber", Msg: msg}
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

// Create takes Batch Header and Entries and builds a valid batch
func (batch *BatchPOP) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...

	return batch.Validate()
}
