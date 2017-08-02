// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
)

// BatchCCD holds the Batch Header and Batch Control and all Entry Records for PPD Entries
type BatchCCD struct {
	header  *BatchHeader
	entries []*EntryDetail
	control *BatchControl
	// Converters is composed for ACH to golang Converters
	converters
}

// NewBatchCCD returns a *BatchCCD
func NewBatchCCD() *BatchCCD {
	batch := new(BatchCCD)
	bh := NewBatchHeader()
	bh.StandardEntryClassCode = ccd
	batch.SetHeader(bh)
	batch.SetControl(NewBatchControl())
	batch.GetHeader().StandardEntryClassCode = ccd
	return batch
}

// Validate checks valid NACHA batch rules. Assumes properly parsed records.
func (batch *BatchCCD) Validate() error {
	batchNumber := batch.header.BatchNumber
	// validate batch header and control codes are the same
	if batch.header.ServiceClassCode != batch.control.ServiceClassCode {
		msg := fmt.Sprintf(msgBatchHeaderControlEquality, batch.header.ServiceClassCode, batch.control.ServiceClassCode)
		return &BatchError{BatchNumber: batchNumber, FieldName: "ServiceClassCode", Msg: msg}
	}
	// Company Identification must match the Company ID from the batch header record
	if batch.header.CompanyIdentification != batch.control.CompanyIdentification {
		msg := fmt.Sprintf(msgBatchHeaderControlEquality, batch.header.CompanyIdentification, batch.control.CompanyIdentification)
		return &BatchError{BatchNumber: batchNumber, FieldName: "CompanyIdentification", Msg: msg}
	}
	// Control ODFI Identification must be the same as batch header
	if batch.header.ODFIIdentification != batch.control.ODFIIdentification {
		msg := fmt.Sprintf(msgBatchHeaderControlEquality, batch.header.ODFIIdentification, batch.control.ODFIIdentification)
		return &BatchError{BatchNumber: batchNumber, FieldName: "ODFIIdentification", Msg: msg}
	}
	// batch number header and control must match
	if batch.header.BatchNumber != batch.control.BatchNumber {
		msg := fmt.Sprintf(msgBatchHeaderControlEquality, batch.header.ODFIIdentification, batch.control.ODFIIdentification)
		return &BatchError{BatchNumber: batchNumber, FieldName: "BatchNumber", Msg: msg}
	}

	if err := batch.isBatchEntryCount(); err != nil {
		return err
	}

	if err := batch.isSequenceAscending(); err != nil {
		return err
	}

	if err := batch.isBatchAmount(); err != nil {
		return err
	}

	if err := batch.isEntryHash(); err != nil {
		return err
	}

	if err := batch.isOriginatorDNE(); err != nil {
		return err
	}

	if err := batch.isTraceNumberODFI(); err != nil {
		return err
	}

	if err := batch.isAddendaSequence(); err != nil {
		return err
	}

	return nil
}

// ValidateAll is a deep validation of all recods within the batch
func (batch *BatchCCD) ValidateAll() error {
	if err := batch.header.Validate(); err != nil {
		return err
	}
	for _, entry := range batch.entries {
		if err := entry.Validate(); err != nil {
			return err
		}
		for _, addenda := range entry.Addendums {
			if err := addenda.Validate(); err != nil {
				return err
			}
		}
	}
	if err := batch.control.Validate(); err != nil {
		return err
	}
	// Validate the Batch wrapper.
	if err := batch.Validate(); err != nil {
		return err
	}
	return nil
}

// Build takes Batch Header and Entries and builds a valid batch
func (batch *BatchCCD) Build() error {
	// Requires a valid BatchHeader
	if err := batch.header.Validate(); err != nil {
		return err
	}
	// Specific to Batch type.
	if len(batch.entries) <= 0 {
		return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "entries", Msg: msgBatchEntries}
	}
	// build controls and sequence numbers
	entryCount := 0
	seq := 1
	for i, entry := range batch.entries {
		entryCount = entryCount + 1 + len(entry.Addendums)
		batch.entries[i].setTraceNumber(batch.header.ODFIIdentification, seq)
		seq++
		addendaSeq := 1
		for x := range entry.Addendums {
			batch.entries[i].Addendums[x].SequenceNumber = addendaSeq
			batch.entries[i].Addendums[x].EntryDetailSequenceNumber = batch.parseNumField(batch.entries[i].TraceNumberField()[8:])
			addendaSeq++
		}
	}

	// build a BatchControl record
	bc := NewBatchControl()
	bc.ServiceClassCode = batch.header.ServiceClassCode
	bc.CompanyIdentification = batch.header.CompanyIdentification
	bc.ODFIIdentification = batch.header.ODFIIdentification
	bc.BatchNumber = batch.header.BatchNumber
	bc.EntryAddendaCount = entryCount
	bc.EntryHash = batch.parseNumField(batch.calculateEntryHash())
	bc.TotalCreditEntryDollarAmount, bc.TotalDebitEntryDollarAmount = batch.calculateBatchAmounts()
	batch.control = bc

	// validate checks that the above build covered all validation checks
	if err := batch.ValidateAll(); err != nil {
		return err // only errors if source code of build is inconcisstent with validate
	}
	return nil
}

// SetHeader appends an BatchHeader to the Batch
func (batch *BatchCCD) SetHeader(batchHeader *BatchHeader) {
	batch.header = batchHeader
}

// GetHeader returns the curent Batch header
func (batch *BatchCCD) GetHeader() *BatchHeader {
	return batch.header
}

// SetControl appends an BatchControl to the Batch
func (batch *BatchCCD) SetControl(batchControl *BatchControl) {
	batch.control = batchControl
}

// GetControl returns the curent Batch Control
func (batch *BatchCCD) GetControl() *BatchControl {
	return batch.control
}

// GetEntries returns a slice of entry details for the batch
func (batch *BatchCCD) GetEntries() []*EntryDetail {
	return batch.entries
}

// AddEntry appends an EntryDetail to the Batch
func (batch *BatchCCD) AddEntry(entry *EntryDetail) Batcher {
	batch.entries = append(batch.entries, entry)
	//	return batch.entries
	return Batcher(batch)
}

// isBatchEntryCount validate Entry count is accurate
// The Entry/Addenda Count Field is a tally of each Entry Detail and Addenda
// Record processed within the batch
func (batch *BatchCCD) isBatchEntryCount() error {
	// entryCount := 0
	// for _, entry := range batch.entries {
	// 	entryCount = entryCount + 1 + len(entry.Addendums)
	// }
	// if entryCount != batch.control.EntryAddendaCount {
	// 	msg := fmt.Sprintf(msgBatchCalculatedControlEquality, entryCount, batch.control.EntryAddendaCount)
	// 	return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "EntryAddendaCount", Msg: msg}
	// }
	return nil
}

// isBatchAmount validate Amount is the same as what is in the Entries
// The Total Debit and Credit Entry Dollar Amount fields contain accumulated
// Entry Detail debit and credit totals within a given batch
func (batch *BatchCCD) isBatchAmount() error {
	credit, debit := batch.calculateBatchAmounts()
	if debit != batch.control.TotalDebitEntryDollarAmount {
		msg := fmt.Sprintf(msgBatchCalculatedControlEquality, debit, batch.control.TotalDebitEntryDollarAmount)
		return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "TotalDebitEntryDollarAmount", Msg: msg}
	}

	if credit != batch.control.TotalCreditEntryDollarAmount {
		msg := fmt.Sprintf(msgBatchCalculatedControlEquality, credit, batch.control.TotalCreditEntryDollarAmount)
		return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "TotalCreditEntryDollarAmount", Msg: msg}
	}
	return nil
}

func (batch *BatchCCD) calculateBatchAmounts() (credit int, debit int) {
	for _, entry := range batch.entries {
		if entry.TransactionCode == 22 || entry.TransactionCode == 23 {
			credit = credit + entry.Amount
		}
		if entry.TransactionCode == 27 || entry.TransactionCode == 28 {
			debit = debit + entry.Amount
		}
		// savings credit
		if entry.TransactionCode == 32 || entry.TransactionCode == 33 {
			credit = credit + entry.Amount
		}
		// savings debit
		if entry.TransactionCode == 37 || entry.TransactionCode == 38 {
			debit = debit + entry.Amount
		}
	}
	return credit, debit
}

// isSequenceAscending Individual Entry Detail Records within individual batches must
// be in ascending Trace Number order (although Trace Numbers need not necessarily be consecutive).
func (batch *BatchCCD) isSequenceAscending() error {
	lastSeq := -1
	for _, entry := range batch.entries {
		if entry.TraceNumber <= lastSeq {
			msg := fmt.Sprintf(msgBatchAscending, entry.TraceNumber, lastSeq)
			return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "TraceNumber", Msg: msg}
		}
		lastSeq = entry.TraceNumber
	}
	return nil
}

// isEntryHash validates the hash by recalulating the result
func (batch *BatchCCD) isEntryHash() error {
	hashField := batch.calculateEntryHash()
	if hashField != batch.control.EntryHashField() {
		msg := fmt.Sprintf(msgBatchCalculatedControlEquality, hashField, batch.control.EntryHashField())
		return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "EntryHash", Msg: msg}
	}
	return nil
}

// calculateEntryHash This field is prepared by hashing the 8-digit Routing Number in each entry.
// The Entry Hash provides a check against inadvertent alteration of data
func (batch *BatchCCD) calculateEntryHash() string {
	hash := 0
	for _, entry := range batch.entries {
		hash = hash + entry.RDFIIdentification
	}
	return batch.numericField(hash, 10)
}

// The Originator Status Code is not equal to “2” for DNE if the Transaction Code is 23 or 33
func (batch *BatchCCD) isOriginatorDNE() error {
	if batch.header.OriginatorStatusCode != 2 {
		for _, entry := range batch.entries {
			if entry.TransactionCode == 23 || entry.TransactionCode == 33 {
				msg := fmt.Sprintf(msgBatchOriginatorDNE, batch.header.OriginatorStatusCode)
				return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "OriginatorStatusCode", Msg: msg}
			}
		}
	}
	return nil
}

// isTraceNumberODFI checks if the first 8 positions of the entry detail trace number
// match the batch header odfi
func (batch *BatchCCD) isTraceNumberODFI() error {
	for _, entry := range batch.entries {
		if batch.header.ODFIIdentificationField() != entry.TraceNumberField()[:8] {
			msg := fmt.Sprintf(msgBatchTraceNumberNotODFI, batch.header.ODFIIdentificationField(), entry.TraceNumberField()[:8])
			return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "ODFIIdentificationField", Msg: msg}
		}
	}

	return nil
}

// isAddendaSequence check multiple errors on addenda records in the batch entries
func (batch *BatchCCD) isAddendaSequence() error {
	for _, entry := range batch.entries {
		if len(entry.Addendums) > 0 {
			// addenda without indicator flag of 1
			if entry.AddendaRecordIndicator != 1 {
				return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "AddendaRecordIndicator", Msg: msgBatchAddendaIndicator}
			}
			lastSeq := -1
			// check if sequence is assending
			for _, addenda := range entry.Addendums {
				if addenda.SequenceNumber < lastSeq {
					msg := fmt.Sprintf(msgBatchAscending, addenda.SequenceNumber, lastSeq)
					return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "SequenceNumber", Msg: msg}
				}
				lastSeq = addenda.SequenceNumber
				// check that we are in the correct Entry Detail
				if !(addenda.EntryDetailSequenceNumberField() == entry.TraceNumberField()[8:]) {
					msg := fmt.Sprintf(msgBatchAddendaTraceNumber, addenda.EntryDetailSequenceNumberField(), entry.TraceNumberField()[8:])
					return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "TraceNumber", Msg: msg}
				}
			}
		}
	}
	return nil
}
