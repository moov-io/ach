// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strconv"
)

var (
	msgIATBatchAddendaRequired  = "is required for an IAT detail entry"
	msgIATBatchAddendaIndicator = "is invalid for addenda record(s) found"
)

// IATBatch holds the Batch Header and Batch Control and all Entry Records for an IAT batch
//
// An IAT entry is a credit or debit ACH entry that is part of a payment transaction involving
// a financial agency’s office (i.e., depository financial institution or business issuing money
// orders) that is not located in the territorial jurisdiction of the United States. IAT entries
// can be made to or from a corporate or consumer account and must be accompanied by seven (7)
// mandatory addenda records identifying the name and physical address of the Originator, name
// and physical address of the Receiver, Receiver’s account number, Receiver’s bank identity and
// reason for the payment.
type IATBatch struct {
	// ID is a client defined string used as a reference to this record.
	ID      string            `json:"id"`
	Header  *IATBatchHeader   `json:"IATBatchHeader,omitempty"`
	Entries []*IATEntryDetail `json:"IATEntryDetails,omitempty"`
	Control *BatchControl     `json:"batchControl,omitempty"`

	// category defines if the entry is a Forward, Return, or NOC
	category string
	// Converters is composed for ACH to GoLang Converters
	converters
}

// IATNewBatch takes a BatchHeader and returns a matching SEC code batch type that is a batcher. Returns an error if the SEC code is not supported.
func IATNewBatch(bh *IATBatchHeader) IATBatch {
	iatBatch := IATBatch{}
	iatBatch.SetControl(NewBatchControl())
	iatBatch.SetHeader(bh)
	return iatBatch
}

// verify checks basic valid NACHA batch rules. Assumes properly parsed records. This does not mean it is a valid batch as validity is tied to each batch type
func (batch *IATBatch) verify() error {
	batchNumber := batch.Header.BatchNumber

	// No entries in batch
	if len(batch.Entries) <= 0 {
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "entries", Msg: msgBatchEntries}
	}
	// verify field inclusion in all the records of the batch.
	if err := batch.isFieldInclusion(); err != nil {
		// convert the field error in to a batch error for a consistent api
		if e, ok := err.(*FieldError); ok {
			return &BatchError{BatchNumber: batchNumber, FieldName: e.FieldName, Msg: e.Msg}
		}
		return &BatchError{BatchNumber: batchNumber, FieldName: "FieldError", Msg: err.Error()}
	}
	// validate batch header and control codes are the same
	if batch.Header.ServiceClassCode != batch.Control.ServiceClassCode {
		msg := fmt.Sprintf(msgBatchHeaderControlEquality, batch.Header.ServiceClassCode, batch.Control.ServiceClassCode)
		return &BatchError{BatchNumber: batchNumber, FieldName: "ServiceClassCode", Msg: msg}
	}
	// Control ODFIIdentification must be the same as batch header
	if batch.Header.ODFIIdentification != batch.Control.ODFIIdentification {
		msg := fmt.Sprintf(msgBatchHeaderControlEquality, batch.Header.ODFIIdentification, batch.Control.ODFIIdentification)
		return &BatchError{BatchNumber: batchNumber, FieldName: "ODFIIdentification", Msg: msg}
	}
	// batch number header and control must match
	if batch.Header.BatchNumber != batch.Control.BatchNumber {
		msg := fmt.Sprintf(msgBatchHeaderControlEquality, batch.Header.BatchNumber, batch.Control.BatchNumber)
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
	if err := batch.isTraceNumberODFI(); err != nil {
		return err
	}
	// TODO this is specific to batch SEC types and should be called by that validator
	if err := batch.isAddendaSequence(); err != nil {
		return err
	}
	if err := batch.isCategory(); err != nil {
		return err
	}
	return nil
}

// Build creates valid batch by building sequence numbers and batch batch control. An error is returned if
// the batch being built has invalid records.
func (batch *IATBatch) build() error {
	// Requires a valid BatchHeader
	if err := batch.Header.Validate(); err != nil {
		return err
	}
	if len(batch.Entries) <= 0 {
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "entries", Msg: msgBatchEntries}
	}
	// Create record sequence numbers
	entryCount := 0
	seq := 1
	for i, entry := range batch.Entries {
		entryCount = entryCount + 1 + 7 + len(entry.Addendum)
		//ToDo: Add  Addenda17 and Addenda18  maximum of 2 addenda17 and 5 addenda18

		// Verifies the required addenda* properties for an IAT entry detail are defined
		if err := batch.addendaFieldInclusion(entry); err != nil {
			return err
		}

		currentTraceNumberODFI, err := strconv.Atoi(entry.TraceNumberField()[:8])
		if err != nil {
			return err
		}

		batchHeaderODFI, err := strconv.Atoi(batch.Header.ODFIIdentificationField()[:8])
		if err != nil {
			return err
		}

		// Add a sequenced TraceNumber if one is not already set. Have to keep original trance number Return and NOC entries
		if currentTraceNumberODFI != batchHeaderODFI {
			batch.Entries[i].SetTraceNumber(batch.Header.ODFIIdentification, seq)
		}
		seq++

		entry.Addenda10.EntryDetailSequenceNumber = batch.parseNumField(batch.Entries[i].TraceNumberField()[8:])
		entry.Addenda11.EntryDetailSequenceNumber = batch.parseNumField(batch.Entries[i].TraceNumberField()[8:])
		entry.Addenda12.EntryDetailSequenceNumber = batch.parseNumField(batch.Entries[i].TraceNumberField()[8:])
		entry.Addenda13.EntryDetailSequenceNumber = batch.parseNumField(batch.Entries[i].TraceNumberField()[8:])
		entry.Addenda14.EntryDetailSequenceNumber = batch.parseNumField(batch.Entries[i].TraceNumberField()[8:])
		entry.Addenda15.EntryDetailSequenceNumber = batch.parseNumField(batch.Entries[i].TraceNumberField()[8:])
		entry.Addenda16.EntryDetailSequenceNumber = batch.parseNumField(batch.Entries[i].TraceNumberField()[8:])

		// Addenda17 and Addenda18 SequenceNumber and EntryDetailSequenceNumber
		seq++
		addendaSeq := 1
		for x := range entry.Addendum {
			if a, ok := batch.Entries[i].Addendum[x].(*Addenda17); ok {
				a.SequenceNumber = addendaSeq
				a.EntryDetailSequenceNumber = batch.parseNumField(batch.Entries[i].TraceNumberField()[8:])
			}
			addendaSeq++
		}

	}

	// build a BatchControl record
	bc := NewBatchControl()
	bc.ServiceClassCode = batch.Header.ServiceClassCode
	bc.ODFIIdentification = batch.Header.ODFIIdentification
	bc.BatchNumber = batch.Header.BatchNumber
	bc.EntryAddendaCount = entryCount
	bc.EntryHash = batch.parseNumField(batch.calculateEntryHash())
	bc.TotalCreditEntryDollarAmount, bc.TotalDebitEntryDollarAmount = batch.calculateBatchAmounts()
	batch.Control = bc

	return nil
}

// SetHeader appends an BatchHeader to the Batch
func (batch *IATBatch) SetHeader(batchHeader *IATBatchHeader) {
	batch.Header = batchHeader
}

// GetHeader returns the current Batch header
func (batch *IATBatch) GetHeader() *IATBatchHeader {
	return batch.Header
}

// SetControl appends an BatchControl to the Batch
func (batch *IATBatch) SetControl(batchControl *BatchControl) {
	batch.Control = batchControl
}

// GetControl returns the current Batch Control
func (batch *IATBatch) GetControl() *BatchControl {
	return batch.Control
}

// GetEntries returns a slice of entry details for the batch
func (batch *IATBatch) GetEntries() []*IATEntryDetail {
	return batch.Entries
}

// AddEntry appends an EntryDetail to the Batch
func (batch *IATBatch) AddEntry(entry *IATEntryDetail) {
	batch.category = entry.Category
	batch.Entries = append(batch.Entries, entry)
}

// Category returns IATBatch Category
// ToDo: Verify this process is the same as a non IAT Batch
func (batch *IATBatch) Category() string {
	return batch.category
}

// isFieldInclusion iterates through all the records in the batch and verifies against default fields
func (batch *IATBatch) isFieldInclusion() error {
	if err := batch.Header.Validate(); err != nil {
		return err
	}
	for _, entry := range batch.Entries {
		if err := entry.Validate(); err != nil {
			return err
		}
		// Verifies the required Addenda* properties for an IAT entry detail are included
		if err := batch.addendaFieldInclusion(entry); err != nil {
			return err
		}

		// Verifies each Addenda* record is valid
		if err := entry.Addenda10.Validate(); err != nil {
			return err
		}
		if err := entry.Addenda11.Validate(); err != nil {
			return err
		}
		if err := entry.Addenda12.Validate(); err != nil {
			return err
		}
		if err := entry.Addenda13.Validate(); err != nil {
			return err
		}
		if err := entry.Addenda14.Validate(); err != nil {
			return err
		}
		if err := entry.Addenda15.Validate(); err != nil {
			return err
		}
		if err := entry.Addenda16.Validate(); err != nil {
			return err
		}
	}
	return batch.Control.Validate()
}

// isBatchEntryCount validate Entry count is accurate
// The Entry/Addenda Count Field is a tally of each Entry Detail and Addenda
// Record processed within the batch
func (batch *IATBatch) isBatchEntryCount() error {
	entryCount := 0
	for _, entry := range batch.Entries {
		entryCount = entryCount + 1 + 7 + len(entry.Addendum)
	}
	if entryCount != batch.Control.EntryAddendaCount {
		msg := fmt.Sprintf(msgBatchCalculatedControlEquality, entryCount, batch.Control.EntryAddendaCount)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "EntryAddendaCount", Msg: msg}
	}
	return nil
}

// isBatchAmount validate Amount is the same as what is in the Entries
// The Total Debit and Credit Entry Dollar Amount fields contain accumulated
// Entry Detail debit and credit totals within a given batch
func (batch *IATBatch) isBatchAmount() error {
	credit, debit := batch.calculateBatchAmounts()
	if debit != batch.Control.TotalDebitEntryDollarAmount {
		msg := fmt.Sprintf(msgBatchCalculatedControlEquality, debit, batch.Control.TotalDebitEntryDollarAmount)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TotalDebitEntryDollarAmount", Msg: msg}
	}

	if credit != batch.Control.TotalCreditEntryDollarAmount {
		msg := fmt.Sprintf(msgBatchCalculatedControlEquality, credit, batch.Control.TotalCreditEntryDollarAmount)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TotalCreditEntryDollarAmount", Msg: msg}
	}
	return nil
}

func (batch *IATBatch) calculateBatchAmounts() (credit int, debit int) {
	for _, entry := range batch.Entries {
		if entry.TransactionCode == 21 || entry.TransactionCode == 22 || entry.TransactionCode == 23 || entry.TransactionCode == 32 || entry.TransactionCode == 33 {
			credit = credit + entry.Amount
		}
		if entry.TransactionCode == 26 || entry.TransactionCode == 27 || entry.TransactionCode == 28 || entry.TransactionCode == 36 || entry.TransactionCode == 37 || entry.TransactionCode == 38 {
			debit = debit + entry.Amount
		}
	}
	return credit, debit
}

// isSequenceAscending Individual Entry Detail Records within individual batches must
// be in ascending Trace Number order (although Trace Numbers need not necessarily be consecutive).
func (batch *IATBatch) isSequenceAscending() error {
	lastSeq := -1
	for _, entry := range batch.Entries {
		if entry.TraceNumber <= lastSeq {
			msg := fmt.Sprintf(msgBatchAscending, entry.TraceNumber, lastSeq)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TraceNumber", Msg: msg}
		}
		lastSeq = entry.TraceNumber
	}
	return nil
}

// isEntryHash validates the hash by recalculating the result
func (batch *IATBatch) isEntryHash() error {
	hashField := batch.calculateEntryHash()
	if hashField != batch.Control.EntryHashField() {
		msg := fmt.Sprintf(msgBatchCalculatedControlEquality, hashField, batch.Control.EntryHashField())
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "EntryHash", Msg: msg}
	}
	return nil
}

// calculateEntryHash This field is prepared by hashing the 8-digit Routing Number in each entry.
// The Entry Hash provides a check against inadvertent alteration of data
func (batch *IATBatch) calculateEntryHash() string {
	hash := 0
	for _, entry := range batch.Entries {

		entryRDFI, _ := strconv.Atoi(entry.RDFIIdentification)

		hash = hash + entryRDFI
	}
	return batch.numericField(hash, 10)
}

// isTraceNumberODFI checks if the first 8 positions of the entry detail trace number
// match the batch header ODFI
func (batch *IATBatch) isTraceNumberODFI() error {
	for _, entry := range batch.Entries {
		if batch.Header.ODFIIdentificationField() != entry.TraceNumberField()[:8] {
			msg := fmt.Sprintf(msgBatchTraceNumberNotODFI, batch.Header.ODFIIdentificationField(), entry.TraceNumberField()[:8])
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "ODFIIdentificationField", Msg: msg}
		}
	}

	return nil
}

// isAddendaSequence check multiple errors on addenda records in the batch entries
func (batch *IATBatch) isAddendaSequence() error {
	for _, entry := range batch.Entries {
		// addenda without indicator flag of 1
		if entry.AddendaRecordIndicator != 1 {
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "AddendaRecordIndicator", Msg: msgIATBatchAddendaIndicator}
		}
		// Verify Addenda* entry detail sequence numbers are valid
		entryTN := entry.TraceNumberField()[8:]

		if entry.Addenda10.EntryDetailSequenceNumberField() != entryTN {
			msg := fmt.Sprintf(msgBatchAddendaTraceNumber, entry.Addenda10.EntryDetailSequenceNumberField(), entry.TraceNumberField()[8:])
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TraceNumber", Msg: msg}
		}
		if entry.Addenda11.EntryDetailSequenceNumberField() != entryTN {
			msg := fmt.Sprintf(msgBatchAddendaTraceNumber, entry.Addenda11.EntryDetailSequenceNumberField(), entry.TraceNumberField()[8:])
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TraceNumber", Msg: msg}
		}
		if entry.Addenda12.EntryDetailSequenceNumberField() != entryTN {
			msg := fmt.Sprintf(msgBatchAddendaTraceNumber, entry.Addenda12.EntryDetailSequenceNumberField(), entry.TraceNumberField()[8:])
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TraceNumber", Msg: msg}
		}
		if entry.Addenda13.EntryDetailSequenceNumberField() != entryTN {
			msg := fmt.Sprintf(msgBatchAddendaTraceNumber, entry.Addenda13.EntryDetailSequenceNumberField(), entry.TraceNumberField()[8:])
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TraceNumber", Msg: msg}
		}
		if entry.Addenda14.EntryDetailSequenceNumberField() != entryTN {
			msg := fmt.Sprintf(msgBatchAddendaTraceNumber, entry.Addenda14.EntryDetailSequenceNumberField(), entry.TraceNumberField()[8:])
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TraceNumber", Msg: msg}
		}
		if entry.Addenda15.EntryDetailSequenceNumberField() != entryTN {
			msg := fmt.Sprintf(msgBatchAddendaTraceNumber, entry.Addenda15.EntryDetailSequenceNumberField(), entry.TraceNumberField()[8:])
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TraceNumber", Msg: msg}
		}
		if entry.Addenda16.EntryDetailSequenceNumberField() != entryTN {
			msg := fmt.Sprintf(msgBatchAddendaTraceNumber, entry.Addenda16.EntryDetailSequenceNumberField(), entry.TraceNumberField()[8:])
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TraceNumber", Msg: msg}
		}
		//ToDo:  Add Addenda17 and Addenda 18 logic for SequenceNUmber and EntryDetailSequenceNumber
	}
	return nil
}

// isCategory verifies that a Forward and Return Category are not in the same batch
func (batch *IATBatch) isCategory() error {
	category := batch.GetEntries()[0].Category
	if len(batch.Entries) > 1 {
		for i := 1; i < len(batch.Entries); i++ {
			if batch.Entries[i].Category == CategoryNOC {
				continue
			}
			if batch.Entries[i].Category != category {
				return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Category", Msg: msgBatchForwardReturn}
			}
		}
	}
	return nil
}

func (batch *IATBatch) addendaFieldInclusion(entry *IATEntryDetail) error {
	if entry.Addenda10 == nil {
		msg := fmt.Sprint(msgIATBatchAddendaRequired)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda10", Msg: msg}
	}
	if entry.Addenda11 == nil {
		msg := fmt.Sprint(msgIATBatchAddendaRequired)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda11", Msg: msg}
	}
	if entry.Addenda12 == nil {
		msg := fmt.Sprint(msgIATBatchAddendaRequired)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda12", Msg: msg}
	}
	if entry.Addenda13 == nil {
		msg := fmt.Sprint(msgIATBatchAddendaRequired)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda13", Msg: msg}
	}
	if entry.Addenda14 == nil {
		msg := fmt.Sprint(msgIATBatchAddendaRequired)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda14", Msg: msg}
	}
	if entry.Addenda15 == nil {
		msg := fmt.Sprint(msgIATBatchAddendaRequired)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda15", Msg: msg}
	}
	if entry.Addenda16 == nil {
		msg := fmt.Sprint(msgIATBatchAddendaRequired)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda16", Msg: msg}
	}
	return nil
}

// Create takes Batch Header and Entries and builds a valid batch
func (batch *IATBatch) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...
	return batch.Validate()
}

// Validate checks valid NACHA batch rules. Assumes properly parsed records.
func (batch *IATBatch) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration based validation for this type.

	// IATBatch must have the following mandatory addenda per entry detail:
	// Addenda10,Addenda11,Addenda12,Addenda13,Addenda14,Addenda15,Addenda16

	// ToDo:  IATBatch can have a maximum of 2 optional Addenda17 records
	// ToDo:  IAtBatch can have a maximum of 5 optional Addenda18 records

	// Add type specific validation.
	// ...
	return nil
}
