// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"errors"
	"fmt"
	"strconv"
)

// Batch holds the Batch Header and Batch Control and all Entry Records
type batch struct {
	// id is a client defined string used as a reference to this record. accessed via ID/SetID
	id      string
	Header  *BatchHeader   `json:"batchHeader,omitempty"`
	Entries []*EntryDetail `json:"entryDetails,omitempty"`
	Control *BatchControl  `json:"batchControl,omitempty"`

	// category defines if the entry is a Forward, Return, or NOC
	category string
	// Converters is composed for ACH to GoLang Converters
	converters
}

// NewBatch takes a BatchHeader and returns a matching SEC code batch type that is a batcher. Returns an error if the SEC code is not supported.
func NewBatch(bh *BatchHeader) (Batcher, error) {
	switch bh.StandardEntryClassCode {
	case "ACK":
		return NewBatchACK(bh), nil
	case "ARC":
		return NewBatchARC(bh), nil
	case "ATX":
		return NewBatchATX(bh), nil
	case "BOC":
		return NewBatchBOC(bh), nil
	case "CCD":
		return NewBatchCCD(bh), nil
	case "CIE":
		return NewBatchCIE(bh), nil
	case "COR":
		return NewBatchCOR(bh), nil
	case "CTX":
		return NewBatchCTX(bh), nil
	case "DNE":
		return NewBatchDNE(bh), nil
	case "ENR":
		return NewBatchENR(bh), nil
	case "IAT":
		msg := fmt.Sprintf(msgFileIATSEC, bh.StandardEntryClassCode)
		return nil, &FileError{FieldName: "StandardEntryClassCode", Value: bh.StandardEntryClassCode, Msg: msg}
	case "POP":
		return NewBatchPOP(bh), nil
	case "POS":
		return NewBatchPOS(bh), nil
	case "PPD":
		return NewBatchPPD(bh), nil
	case "RCK":
		return NewBatchRCK(bh), nil
	case "SHR":
		return NewBatchSHR(bh), nil
	case "TEL":
		return NewBatchTEL(bh), nil
	case "WEB":
		return NewBatchWEB(bh), nil
	default:
	}
	msg := fmt.Sprintf(msgFileNoneSEC, bh.StandardEntryClassCode)
	return nil, &FileError{FieldName: "StandardEntryClassCode", Value: bh.StandardEntryClassCode, Msg: msg}
}

func (batch *batch) Create() error {
	return errors.New("use an implementation of batch or NewBatch")
}

func (batch *batch) Validate() error {
	return errors.New("use an implementation of batch or NewBatch")
}

// verify checks basic valid NACHA batch rules. Assumes properly parsed records. This does not mean it is a valid batch as validity is tied to each batch type
func (batch *batch) verify() error {
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
	// Company Identification must match the Company ID from the batch header record
	if batch.Header.CompanyIdentification != batch.Control.CompanyIdentification {
		msg := fmt.Sprintf(msgBatchHeaderControlEquality, batch.Header.CompanyIdentification, batch.Control.CompanyIdentification)
		return &BatchError{BatchNumber: batchNumber, FieldName: "CompanyIdentification", Msg: msg}
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
	if err := batch.isOriginatorDNE(); err != nil {
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
func (batch *batch) build() error {
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
		entryCount = entryCount + 1

		// Add in Addenda Count
		if entry.Addenda02 != nil {
			entryCount = entryCount + 1
		}
		entryCount = entryCount + len(entry.Addenda05)
		if entry.Addenda98 != nil {
			entryCount = entryCount + 1
		}

		if entry.Addenda99 != nil {
			entryCount = entryCount + 1
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
		addendaSeq := 1
		for _, a := range entry.Addenda05 {
			// sequences don't exist in NOC or Return addenda
			a.SequenceNumber = addendaSeq
			a.EntryDetailSequenceNumber = batch.parseNumField(batch.Entries[i].TraceNumberField()[8:])
			addendaSeq++
		}

	}

	// build a BatchControl record
	bc := NewBatchControl()
	bc.ServiceClassCode = batch.Header.ServiceClassCode
	bc.CompanyIdentification = batch.Header.CompanyIdentification
	bc.ODFIIdentification = batch.Header.ODFIIdentification
	bc.BatchNumber = batch.Header.BatchNumber
	bc.EntryAddendaCount = entryCount
	bc.EntryHash = batch.parseNumField(batch.calculateEntryHash())
	bc.TotalCreditEntryDollarAmount, bc.TotalDebitEntryDollarAmount = batch.calculateBatchAmounts()
	batch.Control = bc

	return nil
}

// SetHeader appends an BatchHeader to the Batch
func (batch *batch) SetHeader(batchHeader *BatchHeader) {
	batch.Header = batchHeader
}

// GetHeader returns the current Batch header
func (batch *batch) GetHeader() *BatchHeader {
	return batch.Header
}

// SetControl appends an BatchControl to the Batch
func (batch *batch) SetControl(batchControl *BatchControl) {
	batch.Control = batchControl
}

// GetControl returns the current Batch Control
func (batch *batch) GetControl() *BatchControl {
	return batch.Control
}

// GetEntries returns a slice of entry details for the batch
func (batch *batch) GetEntries() []*EntryDetail {
	return batch.Entries
}

// AddEntry appends an EntryDetail to the Batch
func (batch *batch) AddEntry(entry *EntryDetail) {
	batch.category = entry.Category
	batch.Entries = append(batch.Entries, entry)
}

// IsReturn is true if the batch contains an Entry Return
func (batch *batch) Category() string {
	return batch.category
}

// ID returns the id of the batch
func (batch *batch) ID() string {
	return batch.id
}

// SetID sets the batch id
func (batch *batch) SetID(id string) {
	batch.id = id
}

// isFieldInclusion iterates through all the records in the batch and verifies against default fields
func (batch *batch) isFieldInclusion() error {
	if err := batch.Header.Validate(); err != nil {
		return err
	}
	for _, entry := range batch.Entries {
		if err := entry.Validate(); err != nil {
			return err
		}

		if entry.Addenda02 != nil {
			if err := entry.Addenda02.Validate(); err != nil {
				return err
			}
		}
		for _, addenda05 := range entry.Addenda05 {
			if err := addenda05.Validate(); err != nil {
				return err
			}
		}
		if entry.Addenda98 != nil {
			if err := entry.Addenda98.Validate(); err != nil {
				return err
			}
		}
		if entry.Addenda99 != nil {
			if err := entry.Addenda99.Validate(); err != nil {
				return err
			}
		}
	}
	return batch.Control.Validate()
}

// isBatchEntryCount validate Entry count is accurate
// The Entry/Addenda Count Field is a tally of each Entry Detail and Addenda
// Record processed within the batch
func (batch *batch) isBatchEntryCount() error {
	entryCount := 0
	for _, entry := range batch.Entries {
		entryCount = entryCount + 1

		// Add in Addenda Count
		if entry.Addenda02 != nil {
			entryCount = entryCount + 1
		}
		entryCount = entryCount + len(entry.Addenda05)
		if entry.Addenda98 != nil {
			entryCount = entryCount + 1
		}

		if entry.Addenda99 != nil {
			entryCount = entryCount + 1
		}
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
func (batch *batch) isBatchAmount() error {
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

func (batch *batch) calculateBatchAmounts() (credit int, debit int) {
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
func (batch *batch) isSequenceAscending() error {
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
func (batch *batch) isEntryHash() error {
	hashField := batch.calculateEntryHash()
	if hashField != batch.Control.EntryHashField() {
		msg := fmt.Sprintf(msgBatchCalculatedControlEquality, hashField, batch.Control.EntryHashField())
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "EntryHash", Msg: msg}
	}
	return nil
}

// calculateEntryHash This field is prepared by hashing the 8-digit Routing Number in each entry.
// The Entry Hash provides a check against inadvertent alteration of data
func (batch *batch) calculateEntryHash() string {
	hash := 0
	for _, entry := range batch.Entries {

		entryRDFI, _ := strconv.Atoi(entry.RDFIIdentification)

		hash = hash + entryRDFI
	}
	return batch.numericField(hash, 10)
}

// The Originator Status Code is not equal to “2” for DNE if the Transaction Code is 23 or 33
func (batch *batch) isOriginatorDNE() error {
	if batch.Header.OriginatorStatusCode != 2 {
		for _, entry := range batch.Entries {
			if entry.TransactionCode == 23 || entry.TransactionCode == 33 {
				msg := fmt.Sprintf(msgBatchOriginatorDNE, batch.Header.OriginatorStatusCode)
				return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "OriginatorStatusCode", Msg: msg}
			}
		}
	}
	return nil
}

// isTraceNumberODFI checks if the first 8 positions of the entry detail trace number
// match the batch header ODFI
func (batch *batch) isTraceNumberODFI() error {
	for _, entry := range batch.Entries {
		if batch.Header.ODFIIdentificationField() != entry.TraceNumberField()[:8] {
			msg := fmt.Sprintf(msgBatchTraceNumberNotODFI, batch.Header.ODFIIdentificationField(), entry.TraceNumberField()[:8])
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "ODFIIdentificationField", Msg: msg}
		}
	}
	return nil
}

// isAddendaSequence check multiple errors on addenda records in the batch entries
func (batch *batch) isAddendaSequence() error {
	for _, entry := range batch.Entries {

		if entry.Addenda02 != nil {
			if entry.AddendaRecordIndicator != 1 {
				return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "AddendaRecordIndicator", Msg: msgBatchAddendaIndicator}
			}
		}

		if len(entry.Addenda05) > 0 {
			// addenda without indicator flag of 1
			if entry.AddendaRecordIndicator != 1 {
				return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "AddendaRecordIndicator", Msg: msgBatchAddendaIndicator}
			}
			lastSeq := -1
			// check if sequence is ascending
			for _, a := range entry.Addenda05 {
				// sequences don't exist in NOC or Return addenda

				if a.SequenceNumber < lastSeq {
					msg := fmt.Sprintf(msgBatchAscending, a.SequenceNumber, lastSeq)
					return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "SequenceNumber", Msg: msg}
				}
				lastSeq = a.SequenceNumber
				// check that we are in the correct Entry Detail
				if !(a.EntryDetailSequenceNumberField() == entry.TraceNumberField()[8:]) {
					msg := fmt.Sprintf(msgBatchAddendaTraceNumber, a.EntryDetailSequenceNumberField(), entry.TraceNumberField()[8:])
					return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "TraceNumber", Msg: msg}
				}
			}
		}
		if entry.Addenda98 != nil {
			if entry.AddendaRecordIndicator != 1 {
				return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "AddendaRecordIndicator", Msg: msgBatchAddendaIndicator}
			}
		}
		if entry.Addenda99 != nil {
			if entry.AddendaRecordIndicator != 1 {
				return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "AddendaRecordIndicator", Msg: msgBatchAddendaIndicator}
			}
		}
	}
	return nil
}

// isCategory verifies that a Forward and Return Category are not in the same batch
func (batch *batch) isCategory() error {
	category := batch.GetEntries()[0].Category
	if len(batch.Entries) > 1 {
		for i := 0; i < len(batch.Entries); i++ {
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

// addendaFieldInclusion verifies Addenda* Field Inclusion based on entry.Category and
// batchHeader.StandardEntryClassCode
// Forward Entries:
// MTE, POS, and SHR can only have Addenda02
// ACK, ATX, CCD, CIE, CTX, DNE, ENR, WEB, PPD, TRX can only have Addenda05
// ARC, BOC, POP, RCK, TEL, TRC, XCK cannot have Addenda02 or Addenda05
// Notification of Change:
// COR and Addenda98
// Return:
// Addenda99
//
// ToDo: Implement Batch for MTE, TRX, TRC, XCK
//
func (batch *batch) addendaFieldInclusion(entry *EntryDetail) error {
	switch entry.Category {
	case CategoryForward:
		if err := batch.addendaFieldInclusionForward(entry); err != nil {
			return err
		}
	case CategoryNOC:
		if err := batch.addendaFieldInclusionNOC(entry); err != nil {
			return err
		}
	case CategoryReturn:
		if err := batch.addendaFieldInclusionReturn(entry); err != nil {
			return err
		}
	}
	return nil
}

// addendaFieldInclusionForward verifies Addenda* Field Inclusion for entry.Category Forward
func (batch *batch) addendaFieldInclusionForward(entry *EntryDetail) error {
	switch batch.Header.StandardEntryClassCode {
	case "MTE", "POS", "SHR":
		if entry.Addenda05 != nil {
			msg := fmt.Sprintf(msgBatchAddenda, "Addenda05", entry.Category, batch.Header.StandardEntryClassCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda05", Msg: msg}
		}
		// ACK, ATX, CCD, CIE, CTX, DNE, ENR WEB, PPD, TRX can only have Addenda05
	case "ACK", "ATX", "CCD", "CIE", "CTX", "DNE", "ENR", "WEB", "PPD", "TRX":
		if entry.Addenda02 != nil {
			msg := fmt.Sprintf(msgBatchAddenda, "Addenda02", entry.Category, batch.Header.StandardEntryClassCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda02", Msg: msg}
		}
	case "ARC", "BOC", "COR", "POP", "RCK", "TEL", "TRC", "XCK":
		if entry.Addenda02 != nil {
			msg := fmt.Sprintf(msgBatchAddenda, "Addenda02", entry.Category, batch.Header.StandardEntryClassCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda02", Msg: msg}
		}
		if entry.Addenda05 != nil {
			msg := fmt.Sprintf(msgBatchAddenda, "Addenda05", entry.Category, batch.Header.StandardEntryClassCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda05", Msg: msg}
		}
	}
	if batch.Header.StandardEntryClassCode != "COR" {
		if entry.Addenda98 != nil {
			msg := fmt.Sprintf(msgBatchAddenda, "Addenda98", entry.Category, batch.Header.StandardEntryClassCode)
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda98", Msg: msg}
		}
	}
	if entry.Addenda99 != nil {
		msg := fmt.Sprintf(msgBatchAddenda, "Addenda99", entry.Category, batch.Header.StandardEntryClassCode)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda99", Msg: msg}
	}
	return nil
}

// addendaFieldInclusionNOC verifies Addenda* Field Inclusion for entry.Category NOC
func (batch *batch) addendaFieldInclusionNOC(entry *EntryDetail) error {
	if entry.Addenda02 != nil {
		msg := fmt.Sprintf(msgBatchAddenda, "Addenda02", entry.Category, batch.Header.StandardEntryClassCode)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda02", Msg: msg}
	}
	if entry.Addenda05 != nil {
		msg := fmt.Sprintf(msgBatchAddenda, "Addenda05", entry.Category, batch.Header.StandardEntryClassCode)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda05", Msg: msg}
	}
	if batch.Header.StandardEntryClassCode == "COR" {
		if entry.Addenda98 == nil {
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda98", Msg: msgFieldInclusion}
		}
	} else {
		if entry.Addenda98 != nil {
			return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda98", Msg: msgFieldInclusion}
		}
	}
	if entry.Addenda99 != nil {
		msg := fmt.Sprintf(msgBatchAddenda, "Addenda99", entry.Category, batch.Header.StandardEntryClassCode)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda99", Msg: msg}
	}
	return nil
}

// addendaFieldInclusionReturn verifies Addenda* Field Inclusion for entry.Category Return
func (batch *batch) addendaFieldInclusionReturn(entry *EntryDetail) error {
	if entry.Addenda02 != nil {
		msg := fmt.Sprintf(msgBatchAddenda, "Addenda02", entry.Category, batch.Header.StandardEntryClassCode)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda02", Msg: msg}
	}
	if entry.Addenda05 != nil {
		msg := fmt.Sprintf(msgBatchAddenda, "Addenda05", entry.Category, batch.Header.StandardEntryClassCode)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda05", Msg: msg}
	}
	if entry.Addenda98 != nil {
		msg := fmt.Sprintf(msgBatchAddenda, "Addenda98", entry.Category, batch.Header.StandardEntryClassCode)
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda98", Msg: msg}
	}
	if entry.Addenda99 == nil {
		return &BatchError{BatchNumber: batch.Header.BatchNumber, FieldName: "Addenda99", Msg: msgFieldInclusion}
	}
	return nil
}
