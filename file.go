// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// Package ach reads and writes (ACH) Automated Clearing House files. ACH is the
// primary method of electronic money movemenet through the United States.
//
// https://en.wikipedia.org/wiki/Automated_Clearing_House
package ach

import (
	"errors"
	"strconv"
)

// First position of all Record Types. These codes are uniquily assigned to
// the first byte of each row in a file.
const (
	headerPos       = "1"
	batchPos        = "5"
	entryDetailPos  = "6"
	entryAddendaPos = "7"
	batchControlPos = "8"
	fileControlPos  = "9"
)

// File contains the structures of a parsed ACH File.
type File struct {
	Header  FileHeader
	Batches []Batch
	Control FileControl

	// TODO: remove
	Addenda
}

// addEntryDetail appends an EntryDetail to the Batch
func (f *File) addBatch(batch Batch) []Batch {
	f.Batches = append(f.Batches, batch)
	return f.Batches
}

// Errors specific to parsing a Batch container
var (
	ErrBatchServiceClassMismatch = errors.New("Service Class Code is not the same in Header and Control")
	ErrBatchEntryCountMismatch   = errors.New("Batch Entry Count is out-of-balance with number of Entries")
	ErrBatchAmountMismatch       = errors.New("Batch Control debit and credit amounts are not the same as sum of Entries")

	ErrBatchNumberMismatch       = errors.New("Batch Number is not the same in Header as Control")
	ErrBatchAscendingTraceNumber = errors.New("Trace Numbers on the File are not in ascending sequence within a batch")
	ErrValidEntryHash            = errors.New("Entry Hash is not equal to the sum of Entry Detail RDFI Identification")
)

// Batch holds the Batch Header and Batch Control and all Entry Records
type Batch struct {
	Header  BatchHeader
	Entries []EntryDetail
	Control BatchControl
	// Converters is composed for ACH to golang Converters
	Converters
}

// addEntryDetail appends an EntryDetail to the Batch
func (batch *Batch) addEntryDetail(entry EntryDetail) []EntryDetail {
	batch.Entries = append(batch.Entries, entry)
	return batch.Entries
}

// Validate NACHA rules on the entire batch before being added to a File
func (batch *Batch) Validate() error {
	if err := batch.isServiceClassMismatch(); err != nil {
		return err
	}

	if err := batch.isBatchEntryCountMismatch(); err != nil {
		return err
	}

	if err := batch.isBatchNumberMismatch(); err != nil {
		return err
	}
	if err := batch.isSequenceAscending(); err != nil {
		return err
	}

	if err := batch.isBatchAmountMismatch(); err != nil {
		return err
	}

	if err := batch.isEntryHashMismatch(); err != nil {
		return err
	}
	return nil
}

// isServiceClassMismatch validate batch header and control codes are the same
func (batch *Batch) isServiceClassMismatch() error {
	if batch.Header.ServiceClassCode != batch.Control.ServiceClassCode {
		return ErrBatchServiceClassMismatch
	}
	return nil
}

// isBatchEntryCountMismatch validate Entry count is accurate
func (batch *Batch) isBatchEntryCountMismatch() error {
	if len(batch.Entries) != batch.Control.EntryAddendaCount {
		return ErrBatchEntryCountMismatch
	}
	return nil
}

// isBatchAmountMismatch validate Amount is the same as what is in the Entries
func (batch *Batch) isBatchAmountMismatch() error {
	debit := 0
	credit := 0
	savingsCredit := 0
	savingsDebit := 0
	for _, seq := range batch.Entries {
		if seq.TransactionCode == 22 || seq.TransactionCode == 23 {
			credit = credit + seq.Amount
		}
		if seq.TransactionCode == 27 || seq.TransactionCode == 28 {
			debit = debit + seq.Amount
		}
		if seq.TransactionCode == 32 || seq.TransactionCode == 33 {
			savingsCredit = savingsCredit + seq.Amount
		}
		if seq.TransactionCode == 37 || seq.TransactionCode == 38 {
			savingsDebit = savingsDebit + seq.Amount
		}

		// TODO: current unsure of what to do with savings credits and debits.
		if debit != batch.Control.TotalDebitEntryDollarAmount {
			return ErrBatchAmountMismatch
		}
		if credit != batch.Control.TotalCreditEntryDollarAmount {
			return ErrBatchAmountMismatch
		}
	}

	return nil
}

// isBatchNumberMismatch validate batch header and control numbers are the same
func (batch *Batch) isBatchNumberMismatch() error {
	if batch.Header.BatchNumber != batch.Control.BatchNumber {
		return ErrBatchNumberMismatch
	}
	return nil
}

// isSequenceAscending Individual Entry Detail Records within individual batches must
// be in ascending Trace Number order (although Trace Numbers need not necessarily be consecutive).
func (batch *Batch) isSequenceAscending() error {
	lastSeq := 0
	for _, seq := range batch.Entries {
		if seq.TraceNumber < lastSeq {
			return ErrBatchAscendingTraceNumber
		}
		lastSeq = seq.TraceNumber
	}
	return nil
}

func (batch *Batch) isEntryHashMismatch() error {
	hash := 0
	for _, seq := range batch.Entries {
		hash = hash + seq.RDFIIdentification
	}
	// need to keep just the first 10 digits
	// TODO: Need test cases on this adding up more than ten digits
	hashField := batch.leftPad(strconv.Itoa(hash), "0", 10)
	if hashField != batch.Control.EntryHashField() {
		return ErrValidEntryHash
	}
	return nil
}
