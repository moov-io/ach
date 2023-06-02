// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package ach

import (
	"encoding/json"
	"strconv"
)

// IATBatch holds the Batch Header and Batch Control and all Entry Records for an IAT batch
//
// An IAT entry is a credit or debit ACH entry that is part of a payment transaction involving
// a financial agency's office (i.e., depository financial institution or business issuing money
// orders) that is not located in the territorial jurisdiction of the United States. IAT entries
// can be made to or from a corporate or consumer account and must be accompanied by seven (7)
// mandatory addenda records identifying the name and physical address of the Originator, name
// and physical address of the Receiver, Receiver's account number, Receiver's bank identity and
// reason for the payment.
type IATBatch struct {
	// ID is a client defined string used as a reference to this record.
	ID      string            `json:"id"`
	Header  *IATBatchHeader   `json:"IATBatchHeader"`
	Entries []*IATEntryDetail `json:"IATEntryDetails"`
	Control *BatchControl     `json:"batchControl"`

	// category defines if the entry is a Forward, Return, or NOC
	category string
	// Converters is composed for ACH to GoLang Converters
	converters

	validateOpts *ValidateOpts
}

// NewIATBatch takes a BatchHeader and returns a matching SEC code batch type that is a batcher. Returns an error if the SEC code is not supported.
func NewIATBatch(bh *IATBatchHeader) IATBatch {
	if bh == nil {
		bh = NewIATBatchHeader()
	}

	iatBatch := IATBatch{}
	iatBatch.SetControl(NewBatchControl())
	iatBatch.SetHeader(bh)
	iatBatch.ID = bh.ID
	return iatBatch
}

// UnmarshalJSON un-marshals JSON IATBatch
func (iatBatch *IATBatch) UnmarshalJSON(p []byte) error {
	if iatBatch == nil {
		b := NewIATBatch(nil)
		iatBatch = &b
	} else {
		iatBatch.Header = NewIATBatchHeader()
		iatBatch.Control = NewBatchControl()
	}

	type Alias IATBatch
	aux := struct {
		*Alias
	}{
		(*Alias)(iatBatch),
	}
	if err := json.Unmarshal(p, &aux); err != nil {
		return err
	}
	return nil
}

// verify checks basic valid NACHA batch rules. Assumes properly parsed records. This does not mean it is a valid batch as validity is tied to each batch type
func (iatBatch *IATBatch) verify() error {
	// No entries in batch
	if len(iatBatch.Entries) <= 0 {
		return iatBatch.Error("entries", ErrBatchNoEntries)
	}
	// verify field inclusion in all the records of the iatBatch.
	if err := iatBatch.isFieldInclusion(); err != nil {
		// wrap the field error in to a batch error for a consistent api
		return iatBatch.Error("FieldError", err)
	}

	// validate batch header and control codes are the same
	if (iatBatch.validateOpts == nil || !iatBatch.validateOpts.UnequalServiceClassCode) &&
		iatBatch.Header.ServiceClassCode != iatBatch.Control.ServiceClassCode {
		return iatBatch.Error("ServiceClassCode",
			NewErrBatchHeaderControlEquality(iatBatch.Header.ServiceClassCode, iatBatch.Control.ServiceClassCode))
	}

	// Control ODFIIdentification must be the same as batch header
	if iatBatch.Header.ODFIIdentification != iatBatch.Control.ODFIIdentification {
		return iatBatch.Error("ODFIIdentification",
			NewErrBatchHeaderControlEquality(iatBatch.Header.ODFIIdentification, iatBatch.Control.ODFIIdentification))
	}
	// batch number header and control must match
	if iatBatch.Header.BatchNumber != iatBatch.Control.BatchNumber {
		return iatBatch.Error("BatchNumber",
			NewErrBatchHeaderControlEquality(iatBatch.Header.BatchNumber, iatBatch.Control.BatchNumber))
	}
	if err := iatBatch.isBatchEntryCount(); err != nil {
		return err
	}
	if iatBatch.validateOpts == nil || !iatBatch.validateOpts.CustomTraceNumbers {
		if err := iatBatch.isSequenceAscending(); err != nil {
			return err
		}
	}
	if err := iatBatch.isBatchAmount(); err != nil {
		return err
	}
	if err := iatBatch.isEntryHash(); err != nil {
		return err
	}
	if iatBatch.validateOpts == nil || !iatBatch.validateOpts.CustomTraceNumbers {
		if err := iatBatch.isTraceNumberODFI(); err != nil {
			return err
		}
		if err := iatBatch.isAddendaSequence(); err != nil {
			return err
		}
	}
	if err := iatBatch.isCategory(); err != nil {
		return err
	}
	return nil
}

// Build creates valid batch by building sequence numbers and batch control. An error is returned if
// the batch being built has invalid records.
func (iatBatch *IATBatch) build() error {
	// Requires a valid BatchHeader
	if err := iatBatch.Header.Validate(); err != nil {
		return err
	}
	if len(iatBatch.Entries) <= 0 {
		return iatBatch.Error("entries", ErrBatchNoEntries)
	}
	// Create record sequence numbers
	entryCount := 0
	seq := 1
	for i, entry := range iatBatch.Entries {
		entryCount = entryCount + 1 + 7 + len(entry.Addenda17) + len(entry.Addenda18)

		if entry.Addenda98 != nil {
			entryCount = entryCount + 1
		}
		if entry.Addenda99 != nil {
			entryCount = entryCount + 1
		}

		// Verifies the required addenda* properties for an IAT entry detail are defined
		if err := iatBatch.addendaFieldInclusion(entry); err != nil {
			return err
		}

		currentTraceNumberODFI, err := strconv.Atoi(entry.TraceNumberField()[:8])
		if err != nil {
			return err
		}

		batchHeaderODFI, err := strconv.Atoi(iatBatch.Header.ODFIIdentificationField()[:8])
		if err != nil {
			return err
		}

		// Add a sequenced TraceNumber if one is not already set.
		if currentTraceNumberODFI != batchHeaderODFI {
			if opts := iatBatch.validateOpts; opts == nil {
				iatBatch.Entries[i].SetTraceNumber(iatBatch.Header.ODFIIdentification, seq)
			} else {
				// Automatically set the TraceNumber if we are validating Origin and don't have custom trace numbers
				if !opts.BypassOriginValidation && !opts.CustomTraceNumbers {
					iatBatch.Entries[i].SetTraceNumber(iatBatch.Header.ODFIIdentification, seq)
				}
			}
		}

		if entry.Category != CategoryNOC {
			// Set TraceNumber for IATEntryDetail Addenda10-16 Record Properties
			entry.Addenda10.EntryDetailSequenceNumber = iatBatch.parseNumField(iatBatch.Entries[i].TraceNumberField()[8:])
			entry.Addenda11.EntryDetailSequenceNumber = iatBatch.parseNumField(iatBatch.Entries[i].TraceNumberField()[8:])
			entry.Addenda12.EntryDetailSequenceNumber = iatBatch.parseNumField(iatBatch.Entries[i].TraceNumberField()[8:])
			entry.Addenda13.EntryDetailSequenceNumber = iatBatch.parseNumField(iatBatch.Entries[i].TraceNumberField()[8:])
			entry.Addenda14.EntryDetailSequenceNumber = iatBatch.parseNumField(iatBatch.Entries[i].TraceNumberField()[8:])
			entry.Addenda15.EntryDetailSequenceNumber = iatBatch.parseNumField(iatBatch.Entries[i].TraceNumberField()[8:])
			entry.Addenda16.EntryDetailSequenceNumber = iatBatch.parseNumField(iatBatch.Entries[i].TraceNumberField()[8:])
		}
		// Set TraceNumber for Addendumer Addenda17 and Addenda18 SequenceNumber and EntryDetailSequenceNumber
		seq++
		addenda17Seq := 1
		addenda18Seq := 1

		for _, addenda17 := range entry.Addenda17 {
			addenda17.SequenceNumber = addenda17Seq
			addenda17.EntryDetailSequenceNumber = iatBatch.parseNumField(iatBatch.Entries[i].TraceNumberField()[8:])
			addenda17Seq++
		}

		for _, addenda18 := range entry.Addenda18 {
			addenda18.SequenceNumber = addenda18Seq
			addenda18.EntryDetailSequenceNumber = iatBatch.parseNumField(iatBatch.Entries[i].TraceNumberField()[8:])
			addenda18Seq++
		}
	}

	// build a BatchControl record
	bc := NewBatchControl()
	bc.ServiceClassCode = iatBatch.Header.ServiceClassCode
	bc.ODFIIdentification = iatBatch.Header.ODFIIdentification
	bc.BatchNumber = iatBatch.Header.BatchNumber
	bc.EntryAddendaCount = entryCount
	bc.EntryHash = iatBatch.calculateEntryHash()
	bc.TotalCreditEntryDollarAmount, bc.TotalDebitEntryDollarAmount = iatBatch.calculateBatchAmounts()
	iatBatch.Control = bc

	return nil
}

// SetHeader appends an BatchHeader to the Batch
func (iatBatch *IATBatch) SetHeader(batchHeader *IATBatchHeader) {
	iatBatch.Header = batchHeader
}

// GetHeader returns the current Batch header
func (iatBatch *IATBatch) GetHeader() *IATBatchHeader {
	return iatBatch.Header
}

// SetControl appends an BatchControl to the Batch
func (iatBatch *IATBatch) SetControl(batchControl *BatchControl) {
	iatBatch.Control = batchControl
}

// GetControl returns the current Batch Control
func (iatBatch *IATBatch) GetControl() *BatchControl {
	return iatBatch.Control
}

// GetEntries returns a slice of entry details for the batch
func (iatBatch *IATBatch) GetEntries() []*IATEntryDetail {
	return iatBatch.Entries
}

// AddEntry appends an EntryDetail to the Batch
func (iatBatch *IATBatch) AddEntry(entry *IATEntryDetail) {
	iatBatch.category = entry.Category
	iatBatch.Entries = append(iatBatch.Entries, entry)
}

// Category returns IATBatch Category
func (iatBatch *IATBatch) Category() string {
	return iatBatch.category
}

// isFieldInclusion iterates through all the records in the batch and verifies against default fields
func (iatBatch *IATBatch) isFieldInclusion() error {
	if err := iatBatch.Header.Validate(); err != nil {
		return err
	}
	for _, entry := range iatBatch.Entries {
		if err := entry.Validate(); err != nil {
			return err
		}
		// Verifies the required Addenda* properties for an IAT entry detail are included
		if err := iatBatch.addendaFieldInclusion(entry); err != nil {
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
	return iatBatch.Control.Validate()
}

// isBatchEntryCount validate Entry count is accurate
// The Entry/Addenda Count Field is a tally of each Entry Detail and Addenda
// Record processed within the batch
func (iatBatch *IATBatch) isBatchEntryCount() error {
	entryCount := 0
	for _, entry := range iatBatch.Entries {
		entryCount = entryCount + 1 + 7 + len(entry.Addenda17) + len(entry.Addenda18)

		if entry.Addenda98 != nil {
			entryCount = entryCount + 1
		}
		if entry.Addenda99 != nil {
			entryCount = entryCount + 1
		}
	}
	if entryCount != iatBatch.Control.EntryAddendaCount {
		if iatBatch.validateOpts != nil && iatBatch.validateOpts.UnequalAddendaCounts {
			return nil
		}
		return iatBatch.Error("EntryAddendaCount",
			NewErrBatchCalculatedControlEquality(entryCount, iatBatch.Control.EntryAddendaCount))
	}
	return nil
}

// isBatchAmount validate Amount is the same as what is in the Entries
// The Total Debit and Credit Entry Dollar Amount fields contain accumulated
// Entry Detail debit and credit totals within a given batch
func (iatBatch *IATBatch) isBatchAmount() error {
	credit, debit := iatBatch.calculateBatchAmounts()
	if debit != iatBatch.Control.TotalDebitEntryDollarAmount {
		return iatBatch.Error("TotalDebitEntryDollarAmount",
			NewErrBatchCalculatedControlEquality(debit, iatBatch.Control.TotalDebitEntryDollarAmount))
	}

	if credit != iatBatch.Control.TotalCreditEntryDollarAmount {
		return iatBatch.Error("TotalCreditEntryDollarAmount",
			NewErrBatchCalculatedControlEquality(credit, iatBatch.Control.TotalCreditEntryDollarAmount))
	}
	return nil
}

func (iatBatch *IATBatch) calculateBatchAmounts() (credit int, debit int) {
	for _, entry := range iatBatch.Entries {
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
func (iatBatch *IATBatch) isSequenceAscending() error {
	lastSeq := "-1"
	for _, entry := range iatBatch.Entries {
		if iatBatch.validateOpts == nil || !iatBatch.validateOpts.CustomTraceNumbers {
			if entry.TraceNumber <= lastSeq {
				return iatBatch.Error("TraceNumber", NewErrBatchAscending(lastSeq, entry.TraceNumber))
			}
		}
		lastSeq = entry.TraceNumber
	}
	return nil
}

// isEntryHash validates the hash by recalculating the result
func (iatBatch *IATBatch) isEntryHash() error {
	hashField := iatBatch.calculateEntryHash()
	if hashField != iatBatch.Control.EntryHash {
		return iatBatch.Error("EntryHash",
			NewErrBatchCalculatedControlEquality(hashField, iatBatch.Control.EntryHash))
	}
	return nil
}

// calculateEntryHash This field is prepared by hashing the 8-digit Routing Number in each entry.
// The Entry Hash provides a check against inadvertent alteration of data
func (iatBatch *IATBatch) calculateEntryHash() int {
	hash := 0
	for _, entry := range iatBatch.Entries {
		entryRDFI, _ := strconv.Atoi(aba8(entry.RDFIIdentification))
		hash += entryRDFI
	}

	// EntryHash is essentially the sum of all the RDFI routing numbers in the batch. If the sum exceeds 10 digits
	// (because you have lots of Entry Detail Records), lop off the most significant digits of the sum until there
	// are only 10.
	return iatBatch.leastSignificantDigits(hash, 10)
}

// isTraceNumberODFI checks if the first 8 positions of the entry detail trace number
// match the batch header ODFI
func (iatBatch *IATBatch) isTraceNumberODFI() error {
	if iatBatch.validateOpts != nil && iatBatch.validateOpts.BypassOriginValidation {
		return nil
	}
	for _, entry := range iatBatch.Entries {
		if iatBatch.Header.ODFIIdentificationField() != entry.TraceNumberField()[:8] {
			return iatBatch.Error("ODFIIdentificationField",
				NewErrBatchTraceNumberNotODFI(iatBatch.Header.ODFIIdentificationField(), entry.TraceNumberField()[:8]))
		}
	}
	return nil
}

// isAddendaSequence check multiple errors on addenda records in the batch entries
func (iatBatch *IATBatch) isAddendaSequence() error {
	for _, entry := range iatBatch.Entries {
		// addenda without indicator flag of 1
		if entry.AddendaRecordIndicator != 1 {
			return iatBatch.Error("AddendaRecordIndicator", ErrIATBatchAddendaIndicator)
		}

		if entry.Category == CategoryNOC {
			return nil
		}

		// Verify Addenda* entry detail sequence numbers are valid
		entryTN := entry.TraceNumberField()[8:]
		if entry.Addenda10.EntryDetailSequenceNumberField() != entryTN {
			return iatBatch.Error("TraceNumber", NewErrBatchAddendaTraceNumber(entry.Addenda10.EntryDetailSequenceNumberField(), entryTN))
		}
		if entry.Addenda11.EntryDetailSequenceNumberField() != entryTN {
			return iatBatch.Error("TraceNumber", NewErrBatchAddendaTraceNumber(entry.Addenda11.EntryDetailSequenceNumberField(), entryTN))
		}
		if entry.Addenda12.EntryDetailSequenceNumberField() != entryTN {
			return iatBatch.Error("TraceNumber", NewErrBatchAddendaTraceNumber(entry.Addenda12.EntryDetailSequenceNumberField(), entryTN))
		}
		if entry.Addenda13.EntryDetailSequenceNumberField() != entryTN {
			return iatBatch.Error("TraceNumber", NewErrBatchAddendaTraceNumber(entry.Addenda13.EntryDetailSequenceNumberField(), entryTN))
		}
		if entry.Addenda14.EntryDetailSequenceNumberField() != entryTN {
			return iatBatch.Error("TraceNumber", NewErrBatchAddendaTraceNumber(entry.Addenda14.EntryDetailSequenceNumberField(), entryTN))
		}
		if entry.Addenda15.EntryDetailSequenceNumberField() != entryTN {
			return iatBatch.Error("TraceNumber", NewErrBatchAddendaTraceNumber(entry.Addenda15.EntryDetailSequenceNumberField(), entryTN))
		}
		if entry.Addenda16.EntryDetailSequenceNumberField() != entryTN {
			return iatBatch.Error("TraceNumber", NewErrBatchAddendaTraceNumber(entry.Addenda16.EntryDetailSequenceNumberField(), entryTN))
		}

		// check if sequence is ascending for addendumer - Addenda17 and Addenda18
		lastAddenda17Seq := -1
		lastAddenda18Seq := -1

		for _, addenda17 := range entry.Addenda17 {
			if addenda17.SequenceNumber < lastAddenda17Seq {
				return iatBatch.Error("SequenceNumber", NewErrBatchAscending(lastAddenda17Seq, addenda17.SequenceNumber))
			}
			lastAddenda17Seq = addenda17.SequenceNumber
			// check that we are in the correct Entry Detail
			if !(addenda17.EntryDetailSequenceNumberField() == entry.TraceNumberField()[8:]) {
				return iatBatch.Error("TraceNumber", NewErrBatchAddendaTraceNumber(addenda17.EntryDetailSequenceNumberField(), entryTN))
			}
		}

		for _, addenda18 := range entry.Addenda18 {
			if addenda18.SequenceNumber < lastAddenda18Seq {
				return iatBatch.Error("SequenceNumber", NewErrBatchAscending(lastAddenda18Seq, addenda18.SequenceNumber))
			}
			lastAddenda18Seq = addenda18.SequenceNumber
			// check that we are in the correct Entry Detail
			if !(addenda18.EntryDetailSequenceNumberField() == entry.TraceNumberField()[8:]) {
				return iatBatch.Error("TraceNumber", NewErrBatchAddendaTraceNumber(addenda18.EntryDetailSequenceNumberField(), entryTN))
			}
		}
	}
	return nil
}

// isCategory verifies that a Forward and Return Category are not in the same batch
func (iatBatch *IATBatch) isCategory() error {
	category := iatBatch.GetEntries()[0].Category
	if len(iatBatch.Entries) > 1 {
		for i := 1; i < len(iatBatch.Entries); i++ {
			if iatBatch.Entries[i].Category == CategoryNOC {
				continue
			}
			if iatBatch.Entries[i].Category != category {
				return iatBatch.Error("Category", NewErrBatchCategory(iatBatch.Entries[i].Category, category))
			}
		}
	}
	return nil
}

func (iatBatch *IATBatch) addendaFieldInclusion(entry *IATEntryDetail) error {
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
// call the Batch's Validate function at the end of their execution.
func (iatBatch *IATBatch) Create() error {
	// generates sequence numbers and batch control
	if err := iatBatch.build(); err != nil {
		return err
	}
	// Additional steps specific to batch type
	// ...
	return iatBatch.Validate()
}

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the iatBatch.
func (iatBatch *IATBatch) Validate() error {
	// basic verification of the batch before we validate specific rules.
	if err := iatBatch.verify(); err != nil {
		return err
	}
	// Add configuration based validation for this type.

	for _, entry := range iatBatch.Entries {
		if len(entry.Addenda17) > 2 {
			return iatBatch.Error("Addenda17", NewErrBatchAddendaCount(len(entry.Addenda17), 2))
		}

		if len(entry.Addenda18) > 5 {
			return iatBatch.Error("Addenda18", NewErrBatchAddendaCount(len(entry.Addenda18), 5))
		}
		if iatBatch.Header.ServiceClassCode == AutomatedAccountingAdvices {
			return iatBatch.Error("ServiceClassCode", ErrBatchServiceClassCode, iatBatch.Header.ServiceClassCode)
		}
		if entry.Category == CategoryNOC {
			if iatBatch.GetHeader().IATIndicator != IATCOR {
				return iatBatch.Error("IATIndicator", NewErrBatchIATNOC(iatBatch.GetHeader().IATIndicator, IATCOR))
			}
			if iatBatch.GetHeader().StandardEntryClassCode != COR {
				return iatBatch.Error("StandardEntryClassCode", NewErrBatchIATNOC(iatBatch.GetHeader().StandardEntryClassCode, COR))
			}
			switch entry.TransactionCode {
			case CheckingCredit, CheckingDebit, CheckingPrenoteCredit, CheckingPrenoteDebit,
				CheckingZeroDollarRemittanceCredit, CheckingZeroDollarRemittanceDebit,
				SavingsCredit, SavingsDebit, SavingsPrenoteCredit, SavingsPrenoteDebit,
				SavingsZeroDollarRemittanceCredit, SavingsZeroDollarRemittanceDebit,
				GLCredit, GLDebit, GLPrenoteCredit, GLPrenoteDebit,
				GLZeroDollarRemittanceCredit, GLZeroDollarRemittanceDebit,
				LoanCredit, LoanDebit, LoanPrenoteCredit, LoanZeroDollarRemittanceCredit:
				return iatBatch.Error("TransactionCode", ErrBatchTransactionCode, entry.TransactionCode)
			}
		}

	}
	return nil
}

// SetValidation stores ValidateOpts on the Batch which are to be used to override
// the default NACHA validation rules.
func (iatBatch *IATBatch) SetValidation(opts *ValidateOpts) {
	if iatBatch == nil {
		return
	}
	iatBatch.validateOpts = opts
}
