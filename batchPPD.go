// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"errors"
)

// Errors specific to parsing a Batch container
var (
	ErrBatchServiceClassMismatch = errors.New("Service Class Code is not the same in Header and Control")
	ErrBatchEntryCountMismatch   = errors.New("Batch Entry Count is out-of-balance with number of Entries")
	ErrBatchAmountMismatch       = errors.New("Batch Control debit and credit amounts are not the same as sum of Entries")
	ErrBatchNumberMismatch       = errors.New("Batch Number is not the same in Header as Control")
	ErrBatchAscendingTraceNumber = errors.New("Trace Numbers on the File are not in ascending order within a batch")
	ErrBatchAddendaSequence      = errors.New("Addenda Sequence numbers are not in ascending order")
	ErrValidEntryHash            = errors.New("Entry Hash is not equal to the sum of Entry Detail RDFI Identification")
	ErrBatchOriginatorDNE        = errors.New("Originator Status Code is not equal to “2” for DNE if the Transaction Code is 23 or 33")
	ErrBatchCompanyID            = errors.New("Company Identification must match the Company ID from the batch header record")
	ErrBatchODFIIDMismatch       = errors.New("Batch Control ODFI Identification must be the same as batch header")
	ErrBatchTraceNumberNotODFI   = errors.New("Trace Number in an Entry Detail Record are not the same as the ODFI Routing Number")
	ErrBatchAddendaIndicator     = errors.New("Addenda found with no Addenda Indicator in proceeding Entry Detail")
	ErrBatchAddendaTraceNumber   = errors.New("Addenda Entry Detail Sequence number does not match proceeding Entry Detail Trace Number")
	ErrBatchEntries              = errors.New("Batch must have Entrie Record(s) to be built")
)

// BatchPPD holds the Batch Header and Batch Control and all Entry Records for PPD Entries
type BatchPPD struct {
	header  *BatchHeader
	entries []*EntryDetail
	control *BatchControl
	// Converters is composed for ACH to golang Converters
	converters
}

// NewBatchPPD returns a *BatchPPD
func NewBatchPPD() *BatchPPD {
	batch := new(BatchPPD)
	bh := NewBatchHeader()
	bh.StandardEntryClassCode = ppd
	batch.SetHeader(bh)
	batch.SetControl(NewBatchControl())
	batch.GetHeader().StandardEntryClassCode = ppd
	return batch
}

// Validate NACHA rules on the entire batch before being added to a File
func (batch *BatchPPD) Validate() error {
	// validate batch header and control codes are the same
	if batch.header.ServiceClassCode != batch.control.ServiceClassCode {
		return ErrBatchServiceClassMismatch
	}
	// Company Identification must match the Company ID from the batch header record
	if batch.header.CompanyIdentification != batch.control.CompanyIdentification {
		return ErrBatchCompanyID
	}
	// Control ODFI Identification must be the same as batch header
	if batch.header.ODFIIdentification != batch.control.ODFIIdentification {
		return ErrBatchODFIIDMismatch
	}

	// batch number header and control must match
	if batch.header.BatchNumber != batch.control.BatchNumber {
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

	if err := batch.isAddendaSequence(); err != nil {
		return err
	}

	return nil
}

// ValidateAll validate all dependency records in the batch.
func (batch *BatchPPD) ValidateAll() error {
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
func (batch *BatchPPD) Build() error {
	// Requires a valid BatchHeader
	if err := batch.header.Validate(); err != nil {
		return err
	}
	if len(batch.entries) <= 0 {
		return ErrBatchEntries
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

	// Validate the built batch
	if err := batch.ValidateAll(); err != nil {
		return err
	}
	return nil
}

// SetHeader appends an BatchHeader to the Batch
func (batch *BatchPPD) SetHeader(batchHeader *BatchHeader) {
	batch.header = batchHeader
}

// GetHeader returns the curent Batch header
func (batch *BatchPPD) GetHeader() *BatchHeader {
	return batch.header
}

// SetControl appends an BatchControl to the Batch
func (batch *BatchPPD) SetControl(batchControl *BatchControl) {
	batch.control = batchControl
}

// GetControl returns the curent Batch Control
func (batch *BatchPPD) GetControl() *BatchControl {
	return batch.control
}

// GetEntries returns a slice of entry details for the batch
func (batch *BatchPPD) GetEntries() []*EntryDetail {
	return batch.entries
}

// AddEntry appends an EntryDetail to the Batch
//func (batch *Batch) AddEntryDetail(entry EntryDetail) []EntryDetail {
func (batch *BatchPPD) AddEntry(entry *EntryDetail) Batcher {
	//entry.setTraceNumber(batch.header.ODFIIdentification, 1)
	batch.entries = append(batch.entries, entry)
	//	return batch.entries
	return Batcher(batch)
}

// isBatchEntryCountMismatch validate Entry count is accurate
// The Entry/Addenda Count Field is a tally of each Entry Detail and Addenda
// Record processed within the batch
func (batch *BatchPPD) isBatchEntryCountMismatch() error {
	entryCount := 0
	for _, entry := range batch.entries {
		entryCount = entryCount + 1 + len(entry.Addendums)
	}
	if entryCount != batch.control.EntryAddendaCount {
		return ErrBatchEntryCountMismatch
	}
	return nil
}

// isBatchAmountMismatch validate Amount is the same as what is in the Entries
// The Total Debit and Credit Entry Dollar Amount fields contain accumulated
// Entry Detail debit and credit totals within a given batch
func (batch *BatchPPD) isBatchAmountMismatch() error {
	credit, debit := batch.calculateBatchAmounts()
	//fmt.Printf("debit: %v batch debit: %v \n", debit, batch.Control.TotalDebitEntryDollarAmount)

	if debit != batch.control.TotalDebitEntryDollarAmount {
		return ErrBatchAmountMismatch
	}
	//fmt.Printf("credit: %v batch credit: %v \n", credit, batch.Control.TotalCreditEntryDollarAmount)

	if credit != batch.control.TotalCreditEntryDollarAmount {
		return ErrBatchAmountMismatch
	}
	return nil
}

func (batch *BatchPPD) calculateBatchAmounts() (credit int, debit int) {
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
func (batch *BatchPPD) isSequenceAscending() error {
	lastSeq := -1
	for _, entry := range batch.entries {
		if entry.TraceNumber <= lastSeq {
			return ErrBatchAscendingTraceNumber
		}
		lastSeq = entry.TraceNumber
	}
	return nil
}

// isEntryHashMismatch validates the hash by recalulating the result
func (batch *BatchPPD) isEntryHashMismatch() error {
	hashField := batch.calculateEntryHash()
	if hashField != batch.control.EntryHashField() {
		return ErrValidEntryHash
	}
	return nil
}

// calculateEntryHash This field is prepared by hashing the 8-digit Routing Number in each entry.
// The Entry Hash provides a check against inadvertent alteration of data
func (batch *BatchPPD) calculateEntryHash() string {
	hash := 0
	for _, entry := range batch.entries {
		hash = hash + entry.RDFIIdentification
	}
	return batch.numericField(hash, 10)
}

// The Originator Status Code is not equal to “2” for DNE if the Transaction Code is 23 or 33
func (batch *BatchPPD) isOriginatorDNEMismatch() error {
	if batch.header.OriginatorStatusCode != 2 {
		for _, entry := range batch.entries {
			if entry.TransactionCode == 23 || entry.TransactionCode == 33 {
				return ErrBatchOriginatorDNE
			}
		}
	}
	return nil
}

// isTraceNumberODFI checks if the first 8 positions of the entry detail trace number
// match the batch header odfi
func (batch *BatchPPD) isTraceNumberODFI() error {
	for _, entry := range batch.entries {
		if batch.header.ODFIIdentificationField() != entry.TraceNumberField()[:8] {
			return ErrBatchTraceNumberNotODFI
		}
	}

	return nil
}

// isAddendaSequence check multiple errors on addenda records in the batch entries
func (batch *BatchPPD) isAddendaSequence() error {
	for _, entry := range batch.entries {
		if len(entry.Addendums) > 0 {
			// addenda without indicator flag of 1
			if entry.AddendaRecordIndicator != 1 {
				return ErrBatchAddendaIndicator
			}
			seq := -1
			// check if sequence is assending
			for _, addenda := range entry.Addendums {
				if !(addenda.SequenceNumber > seq) {
					return ErrBatchAddendaSequence
				}
				seq = addenda.SequenceNumber
				// check that we are in the correct Entry Detail
				if !(addenda.EntryDetailSequenceNumberField() == entry.TraceNumberField()[8:]) {
					return ErrBatchAddendaTraceNumber
				}
			}
		}
	}
	return nil
}
