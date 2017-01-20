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
	"fmt"
	"strconv"
)

// Errors specific to parsing a Batch container
var (
	ErrBatchServiceClassMismatch = errors.New("Service Class Code is not the same in Header and Control")
	ErrBatchEntryCountMismatch   = errors.New("Batch Entry Count is out-of-balance with number of Entries")
	ErrBatchAmountMismatch       = errors.New("Batch Control debit and credit amounts are not the same as sum of Entries")
	ErrBatchNumberMismatch       = errors.New("Batch Number is not the same in Header as Control")
	ErrBatchAscendingTraceNumber = errors.New("Trace Numbers on the File are not in ascending entryuence within a batch")
	ErrValidEntryHash            = errors.New("Entry Hash is not equal to the sum of Entry Detail RDFI Identification")
	ErrBatchOriginatorDNE        = errors.New("Originator Status Code is not equal to “2” for DNE if the Transaction Code is 23 or 33")
	ErrBatchCompanyID            = errors.New("Company Identification must match the Company ID from the batch header record")
	ErrBatchODFIIDMismatch       = errors.New("Batch Control ODFI Identification must be the same as batch header")
	ErrBatchTraceNumberNotODFI   = errors.New("Trace Number in an Entry Detail Record are not the same as the ODFI Routing Number")
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
	// validate batch header and control codes are the same
	if batch.Header.ServiceClassCode != batch.Control.ServiceClassCode {
		return ErrBatchServiceClassMismatch
	}
	// Company Identification must match the Company ID from the batch header record
	if batch.Header.CompanyIdentification != batch.Control.CompanyIdentification {
		return ErrBatchCompanyID
	}
	// Control ODFI Identification must be the same as batch header
	if batch.Header.ODFIIdentification != batch.Control.ODFIIdentification {
		return ErrBatchODFIIDMismatch
	}

	// batch number header and control must match
	if batch.Header.BatchNumber != batch.Control.BatchNumber {
		return ErrBatchNumberMismatch
	}

	if err := batch.isBatchEntryCountMismatch(); err != nil {
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

	if err := batch.isOriginatorDNEMismatch(); err != nil {
		return err
	}

	if err := batch.isTraceNumberODFI(); err != nil {
		return err
	}

	return nil
}

// isBatchEntryCountMismatch validate Entry count is accurate
// The Entry/Addenda Count Field is a tally of each Entry Detail and Addenda
// Record processed within the batch
func (batch *Batch) isBatchEntryCountMismatch() error {
	entryCount := 0
	for _, entry := range batch.Entries {
		entryCount = entryCount + 1 + len(entry.Addendums)
	}
	if entryCount != batch.Control.EntryAddendaCount {
		return ErrBatchEntryCountMismatch
	}
	return nil
}

// isBatchAmountMismatch validate Amount is the same as what is in the Entries
// The Total Debit and Credit Entry Dollar Amount fields contain accumulated
// Entry Detail debit and credit totals within a given batch
func (batch *Batch) isBatchAmountMismatch() error {
	debit := 0
	credit := 0
	savingsCredit := 0
	savingsDebit := 0
	for _, entry := range batch.Entries {
		if entry.TransactionCode == 22 || entry.TransactionCode == 23 {
			credit = credit + entry.Amount
		}
		if entry.TransactionCode == 27 || entry.TransactionCode == 28 {
			debit = debit + entry.Amount
		}
		if entry.TransactionCode == 32 || entry.TransactionCode == 33 {
			savingsCredit = savingsCredit + entry.Amount
		}
		if entry.TransactionCode == 37 || entry.TransactionCode == 38 {
			savingsDebit = savingsDebit + entry.Amount
		}
	}

	// TODO: current unsure of what to do with savings credits and debits.
	if debit != batch.Control.TotalDebitEntryDollarAmount {
		return ErrBatchAmountMismatch
	}
	if credit != batch.Control.TotalCreditEntryDollarAmount {
		return ErrBatchAmountMismatch
	}

	return nil
}

// isSequenceAscending Individual Entry Detail Records within individual batches must
// be in ascending Trace Number order (although Trace Numbers need not necessarily be consecutive).
func (batch *Batch) isSequenceAscending() error {
	lastSeq := 0
	for _, entry := range batch.Entries {
		if entry.TraceNumber < lastSeq {
			return ErrBatchAscendingTraceNumber
		}
		lastSeq = entry.TraceNumber
	}
	return nil
}

// isEntryHashMismatch validates the hash by recalulating the result
// This field is prepared by hashing the 8-digit Routing Number in each entry.
// The Entry Hash provides a check against inadvertent alteration of data
func (batch *Batch) isEntryHashMismatch() error {
	hash := 0
	for _, entry := range batch.Entries {
		hash = hash + entry.RDFIIdentification
	}
	// need to keep just the first 10 digits
	// TODO: Need test cases on this adding up more than ten digits
	hashField := batch.leftPad(strconv.Itoa(hash), "0", 10)
	if hashField != batch.Control.EntryHashField() {
		return ErrValidEntryHash
	}
	return nil
}

// The Originator Status Code is not equal to “2” for DNE if the Transaction Code is 23 or 33
func (batch *Batch) isOriginatorDNEMismatch() error {
	if batch.Header.OriginatorStatusCode != 2 {
		for _, entry := range batch.Entries {
			if entry.TransactionCode == 23 || entry.TransactionCode == 33 {
				return ErrBatchOriginatorDNE
			}
		}
	}
	return nil
}

// isTraceNumberODFI checks if the first 8 positions of the entry detail trace number
// match the batch header odfi
func (batch *Batch) isTraceNumberODFI() error {
	for _, entry := range batch.Entries {
		if batch.Header.ODFIIdentificationField() != entry.TraceNumberField()[:8] {
			fmt.Printf("header odfi: %s trace slice: %s", batch.Header.ODFIIdentificationField(), entry.TraceNumberField()[:8])
			return ErrBatchTraceNumberNotODFI
		}
	}

	return nil
}
