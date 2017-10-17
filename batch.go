package ach

import "fmt"

// Batch holds the Batch Header and Batch Control and all Entry Records for PPD Entries
type batch struct {
	header  *BatchHeader
	entries []*EntryDetail
	control *BatchControl
	// Converters is composed for ACH to GoLang Converters
	converters
}

// NewBatch takes a BatchParm and returns a matching SEC code batch type that is a batcher. Returns and error if the SEC code is not supported.
func NewBatch(bp BatchParam) (Batcher, error) {
	switch sec := bp.StandardEntryClass; sec {
	case "PPD":
		return NewBatchPPD(bp), nil
	case "WEB":
		return NewBatchWEB(bp), nil
	case "CCD":
		return NewBatchCCD(bp), nil
	case "COR":
		return NewBatchCOR(bp), nil
	default:
		msg := fmt.Sprintf(msgFileNoneSEC, sec)
		return nil, &FileError{FieldName: "StandardEntryClassCode", Msg: msg}
	}
}

// verify checks basic valid NACHA batch rules. Assumes properly parsed records. This does not mean it is a valid batch as validity is tied to each batch type
func (batch *batch) verify() error {
	batchNumber := batch.header.BatchNumber

	// verify field inclusion in all the records of the batch.
	if err := batch.isFieldInclusion(); err != nil {
		// convert the field error in to a batch error for a consistent api
		if e, ok := err.(*FieldError); ok {
			return &BatchError{BatchNumber: batchNumber, FieldName: e.FieldName, Msg: e.Msg}
		}
		return &BatchError{BatchNumber: batchNumber, FieldName: "FieldError", Msg: err.Error()}
	}
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

// Build creates valid batch by building sequence numbers and batch batch control. An error is returned if
// the batch being built has invalid records.
func (batch *batch) build() error {
	// Requires a valid BatchHeader
	if err := batch.header.Validate(); err != nil {
		return err
	}
	if len(batch.entries) <= 0 {
		return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "entries", Msg: msgBatchEntries}
	}
	// Create record sequence numbers
	entryCount := 0
	seq := 1
	for i, entry := range batch.entries {
		entryCount = entryCount + 1 + len(entry.Addendum)
		batch.entries[i].setTraceNumber(batch.header.ODFIIdentification, seq)
		seq++
		addendaSeq := 1
		for x := range entry.Addendum {
			batch.entries[i].Addendum[x].SequenceNumber = addendaSeq
			batch.entries[i].Addendum[x].EntryDetailSequenceNumber = batch.parseNumField(batch.entries[i].TraceNumberField()[8:])
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

	return nil
}

// SetHeader appends an BatchHeader to the Batch
func (batch *batch) SetHeader(batchHeader *BatchHeader) {
	batch.header = batchHeader
}

// GetHeader returns the current Batch header
func (batch *batch) GetHeader() *BatchHeader {
	return batch.header
}

// SetControl appends an BatchControl to the Batch
func (batch *batch) SetControl(batchControl *BatchControl) {
	batch.control = batchControl
}

// GetControl returns the current Batch Control
func (batch *batch) GetControl() *BatchControl {
	return batch.control
}

// GetEntries returns a slice of entry details for the batch
func (batch *batch) GetEntries() []*EntryDetail {
	return batch.entries
}

// AddEntry appends an EntryDetail to the Batch
func (batch *batch) AddEntry(entry *EntryDetail) {
	batch.entries = append(batch.entries, entry)
}

// isFieldInclusion iterates through all the records in the batch and verifies against default fields
func (batch *batch) isFieldInclusion() error {
	if err := batch.header.Validate(); err != nil {
		return err
	}
	for _, entry := range batch.entries {
		if err := entry.Validate(); err != nil {
			return err
		}
		for _, addenda := range entry.Addendum {
			if err := addenda.Validate(); err != nil {
				return nil
			}
		}
	}
	if err := batch.control.Validate(); err != nil {
		return err
	}
	return nil
}

// isBatchEntryCount validate Entry count is accurate
// The Entry/Addenda Count Field is a tally of each Entry Detail and Addenda
// Record processed within the batch
func (batch *batch) isBatchEntryCount() error {
	entryCount := 0
	for _, entry := range batch.entries {
		entryCount = entryCount + 1 + len(entry.Addendum) + len(entry.ReturnAddendum)
	}
	if entryCount != batch.control.EntryAddendaCount {
		msg := fmt.Sprintf(msgBatchCalculatedControlEquality, entryCount, batch.control.EntryAddendaCount)
		return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "EntryAddendaCount", Msg: msg}
	}
	return nil
}

// isBatchAmount validate Amount is the same as what is in the Entries
// The Total Debit and Credit Entry Dollar Amount fields contain accumulated
// Entry Detail debit and credit totals within a given batch
func (batch *batch) isBatchAmount() error {
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

func (batch *batch) calculateBatchAmounts() (credit int, debit int) {
	for _, entry := range batch.entries {
		if entry.TransactionCode == 21 || entry.TransactionCode == 22 || entry.TransactionCode == 23 || entry.TransactionCode == 32 || entry.TransactionCode == 33 {
			credit = credit + entry.Amount
		}
		if entry.TransactionCode == 26 || entry.TransactionCode == 27 || entry.TransactionCode == 28 || entry.TransactionCode == 37 || entry.TransactionCode == 38 {
			debit = debit + entry.Amount
		}
	}
	return credit, debit
}

// isSequenceAscending Individual Entry Detail Records within individual batches must
// be in ascending Trace Number order (although Trace Numbers need not necessarily be consecutive).
func (batch *batch) isSequenceAscending() error {
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

// isEntryHash validates the hash by recalculating the result
func (batch *batch) isEntryHash() error {
	hashField := batch.calculateEntryHash()
	if hashField != batch.control.EntryHashField() {
		msg := fmt.Sprintf(msgBatchCalculatedControlEquality, hashField, batch.control.EntryHashField())
		return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "EntryHash", Msg: msg}
	}
	return nil
}

// calculateEntryHash This field is prepared by hashing the 8-digit Routing Number in each entry.
// The Entry Hash provides a check against inadvertent alteration of data
func (batch *batch) calculateEntryHash() string {
	hash := 0
	for _, entry := range batch.entries {
		hash = hash + entry.RDFIIdentification
	}
	return batch.numericField(hash, 10)
}

// The Originator Status Code is not equal to “2” for DNE if the Transaction Code is 23 or 33
func (batch *batch) isOriginatorDNE() error {
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
func (batch *batch) isTraceNumberODFI() error {
	for _, entry := range batch.entries {
		if batch.header.ODFIIdentificationField() != entry.TraceNumberField()[:8] {
			msg := fmt.Sprintf(msgBatchTraceNumberNotODFI, batch.header.ODFIIdentificationField(), entry.TraceNumberField()[:8])
			return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "ODFIIdentificationField", Msg: msg}
		}
	}

	return nil
}

// isAddendaSequence check multiple errors on addenda records in the batch entries
func (batch *batch) isAddendaSequence() error {
	for _, entry := range batch.entries {
		if len(entry.Addendum) > 0 {
			// addenda without indicator flag of 1
			if entry.AddendaRecordIndicator != 1 {
				return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "AddendaRecordIndicator", Msg: msgBatchAddendaIndicator}
			}
			lastSeq := -1
			// check if sequence is assending
			for _, addenda := range entry.Addendum {
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

// isAddendaCount iterates through each entry detail and checks the number of addendum is greater than the count paramater otherwise it returns an error.
// Following SEC codes allow for none or one Addendum
// "PPD", "WEB", "CCD", "CIE", "DNE", "MTE", "POS", "SHR"
func (batch *batch) isAddendaCount(count int) error {
	for _, entry := range batch.entries {
		if !entry.HasReturnAddenda() {
			if len(entry.Addendum) > count {
				msg := fmt.Sprintf(msgBatchAddendaCount, len(entry.Addendum), count, batch.header.StandardEntryClassCode)
				return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "AddendaCount", Msg: msg}
			}
		} else {
			if len(entry.ReturnAddendum) > count {
				msg := fmt.Sprintf(msgBatchAddendaCount, len(entry.ReturnAddendum), count, batch.header.StandardEntryClassCode)
				return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "ReturnAddendaCount", Msg: msg}
			}
		}
	}
	return nil
}

// isTypeCode takes a typecode string and verifies addenda records match
func (batch *batch) isTypeCode(typeCode string) error {
	for _, entry := range batch.entries {
		for _, addenda := range entry.Addendum {
			if addenda.TypeCode != typeCode {
				msg := fmt.Sprintf(msgBatchTypeCode, addenda.TypeCode, typeCode, batch.header.StandardEntryClassCode)
				return &BatchError{BatchNumber: batch.header.BatchNumber, FieldName: "TypeCode", Msg: msg}
			}
		}
	}
	return nil
}
