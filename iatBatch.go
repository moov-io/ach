// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"encoding/json"
	"strconv"
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

// NewIATBatch takes a BatchHeader and returns a matching SEC code batch type that is a batcher. Returns an error if the SEC code is not supported.
func NewIATBatch(bh *IATBatchHeader) IATBatch {
	if bh == nil {
		bh = NewIATBatchHeader()
	}

	iatBatch := IATBatch{}
	iatBatch.SetControl(NewBatchControl())
	iatBatch.SetHeader(bh)
	return iatBatch
}

// UnmarshalJSON un-marshals JSON IATBatch
func (batch *IATBatch) UnmarshalJSON(p []byte) error {
	if batch == nil {
		b := NewIATBatch(nil)
		batch = &b
	} else {
		batch.Header = NewIATBatchHeader()
		batch.Control = NewBatchControl()
	}

	type Alias IATBatch
	aux := struct {
		*Alias
	}{
		(*Alias)(batch),
	}
	if err := json.Unmarshal(p, &aux); err != nil {
		return err
	}
	return nil
}

// verify checks basic valid NACHA batch rules. Assumes properly parsed records. This does not mean it is a valid batch as validity is tied to each batch type
func (batch *IATBatch) verify() error {
	// No entries in batch
	if len(batch.Entries) <= 0 {
		return batch.Error("entries", ErrBatchNoEntries)
	}
	// verify field inclusion in all the records of the batch.
	if err := batch.isFieldInclusion(); err != nil {
		// wrap the field error in to a batch error for a consistent api
		return batch.Error("FieldError", err)
	}
	// validate batch header and control codes are the same
	if batch.Header.ServiceClassCode != batch.Control.ServiceClassCode {
		return batch.Error("ServiceClassCode",
			NewErrBatchHeaderControlEquality(batch.Header.ServiceClassCode, batch.Control.ServiceClassCode))
	}
	// Control ODFIIdentification must be the same as batch header
	if batch.Header.ODFIIdentification != batch.Control.ODFIIdentification {
		return batch.Error("ODFIIdentification",
			NewErrBatchHeaderControlEquality(batch.Header.ODFIIdentification, batch.Control.ODFIIdentification))
	}
	// batch number header and control must match
	if batch.Header.BatchNumber != batch.Control.BatchNumber {
		return batch.Error("BatchNumber",
			NewErrBatchHeaderControlEquality(batch.Header.BatchNumber, batch.Control.BatchNumber))
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
		return batch.Error("entries", ErrBatchNoEntries)
	}
	// Create record sequence numbers
	entryCount := 0
	seq := 1
	for i, entry := range batch.Entries {
		entryCount = entryCount + 1 + 7 + len(entry.Addenda17) + len(entry.Addenda18)

		if entry.Addenda98 != nil {
			entryCount = entryCount + 1
		}
		if entry.Addenda99 != nil {
			entryCount = entryCount + 1
		}

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

		// Add a sequenced TraceNumber if one is not already set.
		if currentTraceNumberODFI != batchHeaderODFI {
			batch.Entries[i].SetTraceNumber(batch.Header.ODFIIdentification, seq)
		}

		if entry.Category != CategoryNOC {
			// Set TraceNumber for IATEntryDetail Addenda10-16 Record Properties
			entry.Addenda10.EntryDetailSequenceNumber = batch.parseNumField(batch.Entries[i].TraceNumberField()[8:])
			entry.Addenda11.EntryDetailSequenceNumber = batch.parseNumField(batch.Entries[i].TraceNumberField()[8:])
			entry.Addenda12.EntryDetailSequenceNumber = batch.parseNumField(batch.Entries[i].TraceNumberField()[8:])
			entry.Addenda13.EntryDetailSequenceNumber = batch.parseNumField(batch.Entries[i].TraceNumberField()[8:])
			entry.Addenda14.EntryDetailSequenceNumber = batch.parseNumField(batch.Entries[i].TraceNumberField()[8:])
			entry.Addenda15.EntryDetailSequenceNumber = batch.parseNumField(batch.Entries[i].TraceNumberField()[8:])
			entry.Addenda16.EntryDetailSequenceNumber = batch.parseNumField(batch.Entries[i].TraceNumberField()[8:])
		}
		// Set TraceNumber for Addendumer Addenda17 and Addenda18 SequenceNumber and EntryDetailSequenceNumber
		seq++
		addenda17Seq := 1
		addenda18Seq := 1

		for _, addenda17 := range entry.Addenda17 {
			addenda17.SequenceNumber = addenda17Seq
			addenda17.EntryDetailSequenceNumber = batch.parseNumField(batch.Entries[i].TraceNumberField()[8:])
			addenda17Seq++
		}

		for _, addenda18 := range entry.Addenda18 {
			addenda18.SequenceNumber = addenda18Seq
			addenda18.EntryDetailSequenceNumber = batch.parseNumField(batch.Entries[i].TraceNumberField()[8:])
			addenda18Seq++
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
		if entry.Category != CategoryNOC {
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
			for _, Addenda17 := range entry.Addenda17 {
				if err := Addenda17.Validate(); err != nil {
					return err
				}
			}
			for _, Addenda18 := range entry.Addenda18 {
				if err := Addenda18.Validate(); err != nil {
					return err
				}
			}
		}
		if entry.Category == CategoryNOC {
			if entry.Addenda98 == nil {
				return fieldError("Addenda98", ErrFieldInclusion)
			}
			if err := entry.Addenda98.Validate(); err != nil {
				return err
			}
		}
		if entry.Category == CategoryReturn {
			if entry.Addenda99 == nil {
				return fieldError("Addenda99", ErrFieldInclusion)
			}

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
func (batch *IATBatch) isBatchEntryCount() error {
	entryCount := 0
	for _, entry := range batch.Entries {
		entryCount = entryCount + 1 + 7 + len(entry.Addenda17) + len(entry.Addenda18)

		if entry.Addenda98 != nil {
			entryCount = entryCount + 1
		}
		if entry.Addenda99 != nil {
			entryCount = entryCount + 1
		}
	}
	if entryCount != batch.Control.EntryAddendaCount {
		return batch.Error("EntryAddendaCount",
			NewErrBatchCalculatedControlEquality(entryCount, batch.Control.EntryAddendaCount))
	}
	return nil
}

// isBatchAmount validate Amount is the same as what is in the Entries
// The Total Debit and Credit Entry Dollar Amount fields contain accumulated
// Entry Detail debit and credit totals within a given batch
func (batch *IATBatch) isBatchAmount() error {
	credit, debit := batch.calculateBatchAmounts()
	if debit != batch.Control.TotalDebitEntryDollarAmount {
		return batch.Error("TotalDebitEntryDollarAmount",
			NewErrBatchCalculatedControlEquality(debit, batch.Control.TotalDebitEntryDollarAmount))
	}

	if credit != batch.Control.TotalCreditEntryDollarAmount {
		return batch.Error("TotalCreditEntryDollarAmount",
			NewErrBatchCalculatedControlEquality(credit, batch.Control.TotalCreditEntryDollarAmount))
	}
	return nil
}

func (batch *IATBatch) calculateBatchAmounts() (credit int, debit int) {
	for _, entry := range batch.Entries {
		switch entry.TransactionCode {
		case CheckingCredit, CheckingReturnNOCCredit, CheckingPrenoteCredit, CheckingZeroDollarRemittanceCredit,
			SavingsCredit, SavingsReturnNOCCredit, SavingsPrenoteCredit, SavingsZeroDollarRemittanceCredit, GLCredit,
			GLReturnNOCCredit, GLPrenoteCredit, GLZeroDollarRemittanceCredit, LoanCredit, LoanReturnNOCCredit,
			LoanPrenoteCredit, LoanZeroDollarRemittanceCredit:
			credit = credit + entry.Amount
		case CheckingDebit, CheckingReturnNOCDebit, CheckingPrenoteDebit, CheckingZeroDollarRemittanceDebit,
			SavingsDebit, SavingsReturnNOCDebit, SavingsPrenoteDebit, SavingsZeroDollarRemittanceDebit, GLDebit,
			GLReturnNOCDebit, GLPrenoteDebit, GLZeroDollarRemittanceDebit, LoanDebit, LoanReturnNOCDebit:
			debit = debit + entry.Amount
		}
	}
	return credit, debit
}

// isSequenceAscending Individual Entry Detail Records within individual batches must
// be in ascending Trace Number order (although Trace Numbers need not necessarily be consecutive).
func (batch *IATBatch) isSequenceAscending() error {
	lastSeq := "-1"
	for _, entry := range batch.Entries {
		if entry.TraceNumber <= lastSeq {
			return batch.Error("TraceNumber", NewErrBatchAscending(lastSeq, entry.TraceNumber))
		}
		lastSeq = entry.TraceNumber
	}
	return nil
}

// isEntryHash validates the hash by recalculating the result
func (batch *IATBatch) isEntryHash() error {
	hashField := batch.calculateEntryHash()
	if hashField != batch.Control.EntryHashField() {
		return batch.Error("EntryHash",
			NewErrBatchCalculatedControlEquality(hashField, batch.Control.EntryHash))
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
			return batch.Error("ODFIIdentificationField",
				NewErrBatchTraceNumberNotODFI(batch.Header.ODFIIdentificationField(), entry.TraceNumberField()[:8]))
		}
	}

	return nil
}

// isAddendaSequence check multiple errors on addenda records in the batch entries
func (batch *IATBatch) isAddendaSequence() error {
	for _, entry := range batch.Entries {
		// addenda without indicator flag of 1
		if entry.AddendaRecordIndicator != 1 {
			return batch.Error("AddendaRecordIndicator", ErrIATBatchAddendaIndicator)
		}

		if entry.Category == CategoryNOC {
			return nil
		}

		// Verify Addenda* entry detail sequence numbers are valid
		entryTN := entry.TraceNumberField()[8:]
		if entry.Addenda10.EntryDetailSequenceNumberField() != entryTN {
			return batch.Error("TraceNumber", NewErrBatchAddendaTraceNumber(entry.Addenda10.EntryDetailSequenceNumberField(), entryTN))
		}
		if entry.Addenda11.EntryDetailSequenceNumberField() != entryTN {
			return batch.Error("TraceNumber", NewErrBatchAddendaTraceNumber(entry.Addenda11.EntryDetailSequenceNumberField(), entryTN))
		}
		if entry.Addenda12.EntryDetailSequenceNumberField() != entryTN {
			return batch.Error("TraceNumber", NewErrBatchAddendaTraceNumber(entry.Addenda12.EntryDetailSequenceNumberField(), entryTN))
		}
		if entry.Addenda13.EntryDetailSequenceNumberField() != entryTN {
			return batch.Error("TraceNumber", NewErrBatchAddendaTraceNumber(entry.Addenda13.EntryDetailSequenceNumberField(), entryTN))
		}
		if entry.Addenda14.EntryDetailSequenceNumberField() != entryTN {
			return batch.Error("TraceNumber", NewErrBatchAddendaTraceNumber(entry.Addenda14.EntryDetailSequenceNumberField(), entryTN))
		}
		if entry.Addenda15.EntryDetailSequenceNumberField() != entryTN {
			return batch.Error("TraceNumber", NewErrBatchAddendaTraceNumber(entry.Addenda15.EntryDetailSequenceNumberField(), entryTN))
		}
		if entry.Addenda16.EntryDetailSequenceNumberField() != entryTN {
			return batch.Error("TraceNumber", NewErrBatchAddendaTraceNumber(entry.Addenda16.EntryDetailSequenceNumberField(), entryTN))
		}

		// check if sequence is ascending for addendumer - Addenda17 and Addenda18
		lastAddenda17Seq := -1
		lastAddenda18Seq := -1

		for _, addenda17 := range entry.Addenda17 {
			if addenda17.SequenceNumber < lastAddenda17Seq {
				return batch.Error("SequenceNumber", NewErrBatchAscending(lastAddenda17Seq, addenda17.SequenceNumber))
			}
			lastAddenda17Seq = addenda17.SequenceNumber
			// check that we are in the correct Entry Detail
			if !(addenda17.EntryDetailSequenceNumberField() == entry.TraceNumberField()[8:]) {
				return batch.Error("TraceNumber", NewErrBatchAddendaTraceNumber(addenda17.EntryDetailSequenceNumberField(), entryTN))
			}
		}

		for _, addenda18 := range entry.Addenda18 {
			if addenda18.SequenceNumber < lastAddenda18Seq {
				return batch.Error("SequenceNumber", NewErrBatchAscending(lastAddenda18Seq, addenda18.SequenceNumber))
			}
			lastAddenda18Seq = addenda18.SequenceNumber
			// check that we are in the correct Entry Detail
			if !(addenda18.EntryDetailSequenceNumberField() == entry.TraceNumberField()[8:]) {
				return batch.Error("TraceNumber", NewErrBatchAddendaTraceNumber(addenda18.EntryDetailSequenceNumberField(), entryTN))
			}
		}
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
				return batch.Error("Category", NewErrBatchCategory(batch.Entries[i].Category, category))
			}
		}
	}
	return nil
}

func (batch *IATBatch) addendaFieldInclusion(entry *IATEntryDetail) error {
	if entry.Category == CategoryNOC {
		return nil
	}
	if entry.Addenda10 == nil {
		return fieldError("Addenda10", ErrFieldInclusion)
	}
	if entry.Addenda11 == nil {
		return fieldError("Addenda11", ErrFieldInclusion)
	}
	if entry.Addenda12 == nil {
		return fieldError("Addenda12", ErrFieldInclusion)
	}
	if entry.Addenda13 == nil {
		return fieldError("Addenda13", ErrFieldInclusion)
	}
	if entry.Addenda14 == nil {
		return fieldError("Addenda14", ErrFieldInclusion)
	}
	if entry.Addenda15 == nil {
		return fieldError("Addenda15", ErrFieldInclusion)
	}
	if entry.Addenda16 == nil {
		return fieldError("Addenda16", ErrFieldInclusion)
	}
	return nil
}

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate() function at the end of their execution.
func (batch *IATBatch) Create() error {
	// generates sequence numbers and batch control
	if err := batch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...
	return batch.Validate()
}

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *IATBatch) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := batch.verify(); err != nil {
		return err
	}
	// Add configuration based validation for this type.

	for _, entry := range batch.Entries {
		if len(entry.Addenda17) > 2 {
			return batch.Error("Addenda17", NewErrBatchAddendaCount(len(entry.Addenda17), 2))
		}

		if len(entry.Addenda18) > 5 {
			return batch.Error("Addenda18", NewErrBatchAddendaCount(len(entry.Addenda18), 5))
		}
		if batch.Header.ServiceClassCode == AutomatedAccountingAdvices {
			return batch.Error("ServiceClassCode", ErrBatchServiceClassCode, batch.Header.ServiceClassCode)
		}
		if entry.Category == CategoryNOC {
			if batch.GetHeader().IATIndicator != IATCOR {
				return batch.Error("IATIndicator", NewErrBatchIATNOC(batch.GetHeader().IATIndicator, IATCOR))
			}
			if batch.GetHeader().StandardEntryClassCode != COR {
				return batch.Error("StandardEntryClassCode", NewErrBatchIATNOC(batch.GetHeader().StandardEntryClassCode, COR))
			}
			switch entry.TransactionCode {
			case CheckingCredit, CheckingDebit, CheckingPrenoteCredit, CheckingPrenoteDebit,
				CheckingZeroDollarRemittanceCredit, CheckingZeroDollarRemittanceDebit,
				SavingsCredit, SavingsDebit, SavingsPrenoteCredit, SavingsPrenoteDebit,
				SavingsZeroDollarRemittanceCredit, SavingsZeroDollarRemittanceDebit,
				GLCredit, GLDebit, GLPrenoteCredit, GLPrenoteDebit,
				GLZeroDollarRemittanceCredit, GLZeroDollarRemittanceDebit,
				LoanCredit, LoanDebit, LoanPrenoteCredit, LoanZeroDollarRemittanceCredit:
				return batch.Error("TransactionCode", ErrBatchTransactionCode, entry.TransactionCode)
			}
		}

	}
	return nil
}
