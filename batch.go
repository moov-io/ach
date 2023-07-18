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
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Batch holds the Batch Header and Batch Control and all Entry Records
type Batch struct {
	// id is a client defined string used as a reference to this record. accessed via ID/SetID
	id         string
	Header     *BatchHeader      `json:"batchHeader"`
	Entries    []*EntryDetail    `json:"entryDetails"`
	Control    *BatchControl     `json:"batchControl"`
	ADVEntries []*ADVEntryDetail `json:"advEntryDetails,omitempty"`
	ADVControl *ADVBatchControl  `json:"advBatchControl,omitempty"`

	// offset holds the information to build an EntryDetail record which
	// balances the batch by debiting or crediting the sum of amounts in the batch.
	offset *Offset

	// category defines if the entry is a Forward, Return, or NOC
	category string
	// Converters is composed for ACH to GoLang Converters
	converters

	validateOpts *ValidateOpts
}

const (
	// ACK ACH Payment Acknowledgment - A code that indicates acknowledgment of receipt of a corporate credit payment
	// (CCD).
	ACK = "ACK"
	// ADV Automated Accounting Advice – A code that provide accounting information regarding an Entry. It is an
	// optional service.
	ADV = "ADV"
	// ARC Accounts Receivable Entry – A code that indicates a consumer check converted to a one-time ACH debit.
	// The Accounts Receivable (ARC) Entry provides initiates a single-entry ACH debit to customer accounts by
	// converting checks at the point of receipt through the U.S. mail, at a drop box location or in-person for
	// payment of a bill at a manned location.
	ARC = "ARC"
	// ATX Financial EDI Acknowledgment -  A code that indicates acknowledgement by the Receiving Depository Financial
	// Institution (RDFI) that a Corporate Credit Exchange (CTX) has been received.
	ATX = "ATX"
	// BOC Back Office Conversion Entry - A code that indicates single entry debit initiated at the point of purchase
	// or at a manned bill payment location to transfer funds through conversion to an ACH debit entry during back
	// office processing.
	BOC = "BOC"
	// CCD Corporate Credit or Debit Entry - A code that indicates an entry initiated by an Organization to transfer
	// funds to or from an account of that Organization or another Organization. For commercial accounts only.
	CCD = "CCD"
	// CIE Customer Initiated Entry - A code that indicates a credit entry initiated on behalf of, and upon the
	// instruction of, a consumer to transfer funds to a non-consumer Receiver.
	CIE = "CIE"
	// COR Notification of Change or Refused Notification of Change - A code used by an RDFI or ODFI when originating a
	// Notification of Change or Refused Notification of Change in automated format.
	COR = "COR"
	// CTX Corporate Trade Exchange - A code that indicates the ability to collect and disburse funds and information
	// between companies. Generally it is used by businesses paying one another for goods or services.
	CTX = "CTX"
	// DNE Death Notification Entry - A code that United States Federal agencies (e.g. Social Security) use to notify
	// depository financial institutions that the recipient of government benefit payments has died.
	DNE = "DNE"
	// ENR Automated Enrollment Entry - A code indicating enrollment of a person with an agency of the US government
	// for a depository financial institution.
	ENR = "ENR"
	// IAT International ACH Transaction - A code IAT indicating a  credit or debit ACH entry that is part of a payment
	// transaction involving a financial agency's office (i.e., depository financial institution or business issuing
	// money orders) that is not located in the territorial jurisdiction of the United States. IAT entries can be made
	// to or from a corporate or consumer account and must be accompanied by seven (7) mandatory addenda records
	// identifying the name and physical address of the Originator, name and physical address of the Receiver,
	// Receiver's account number, Receiver's bank identity and reason for the payment.
	IAT = "IAT"
	// MTE Machine Transfer Entry - A code that indicates when a consumer uses their debit card at an Automated Teller
	// Machine (ATM) to withdraw cash.  MTE transactions cannot be aggregated together under a single Entry.
	MTE = "MTE"
	// POP Point of Purchase Entry - A code that indicates a check presented in-person to a merchant for purchase
	// is presented as an ACH entry instead of a physical check.
	POP = "POP"
	// POS Point of Sale Entry - A code that indicates a debit entry initiated at an “electronic terminal” to a
	// consumer account of the receiver to pay an obligation incurred in a point-of-sale transaction, or to effect a
	// point-of-sale terminal cash withdrawal.
	POS = "POS"
	// PPD Prearranged Payment and Deposit Entry - A code tha indicates a an entry initiated by an organization based
	// on a standing or a single entry authorization to transfer funds.
	PPD = "PPD"
	// RCK Re-presented Check Entry - A code that indicates a physical check that was presented but returned because of
	//// insufficient funds may be represented as an ACH entry.
	RCK = "RCK"
	// SHR Shared Network Transaction - A code that indicates a debit Entry initiated at an “electronic terminal,” as
	// that term is defined in Regulation E, to a Consumer Account of the Receiver to pay an obligation incurred in a
	// point-of-sale transaction, or to effect a point-of-sale terminal cash withdrawal. Also an adjusting or other
	// credit Entry related to such debit Entry, transfer of funds, or obligation. SHR Entries are initiated in a
	// shared network where the ODFI and RDFI have an agreement in addition to these Rules to process such Entries.
	SHR = "SHR"
	// TEL Telephone Initiated Entry - A code indicating a Telephone-Initiated consumer debit transactions. The NACHA
	// Operating Rules permit TEL entries when the originator obtains the Receiver's authorization for the debit entry
	// orally via the telephone.  An entry based upon a Receiver's oral authorization must utilize the TEL
	// Standard Entry Class (SEC) Code.
	TEL = "TEL"
	// TRC Check Truncation Entry - is a code used to identify a debit entry of a truncated check.
	TRC = "TRC"
	// TRX Check Truncation Entries Exchange - used to identify a debit entry exchange of a truncated checks (multiple).
	TRX = "TRX"
	// WEB Internet-Initiated/Mobile Entry - A code indicating an entry submitted pursuant to an authorization obtained
	// solely via the Internet or a mobile network. For consumer accounts only.
	WEB = "WEB"
	// XCK Destroyed Check Entry - A code indicating a debit entry initiated for a destroyed check eligible items
	XCK = "XCK"
)

func (batch *Batch) MarshalJSON() ([]byte, error) {
	type Alias Batch
	aux := struct {
		*Alias
		Offset *Offset `json:"offset"`
	}{
		(*Alias)(batch),
		batch.offset,
	}
	return json.Marshal(aux)
}

func (batch *Batch) UnmarshalJSON(p []byte) error {
	if batch == nil {
		batch = &Batch{}
	}
	// blank out the fields of our Batch before reading
	batch.Header = NewBatchHeader()
	batch.Control = NewBatchControl()
	batch.ADVControl = NewADVBatchControl()

	type Alias Batch
	aux := struct {
		*Alias
		Offset *Offset `json:"offset"`
	}{
		(*Alias)(batch),
		nil,
	}
	if err := json.Unmarshal(p, &aux); err != nil {
		if e, ok := err.(*json.UnmarshalTypeError); ok {
			return fmt.Errorf("%s: %v", e.Field, err)
		}
		return err
	}
	if aux.Offset != nil {
		batch.offset = aux.Offset
	}
	return nil
}

// NewBatch takes a BatchHeader and returns a matching SEC code batch type that is a batcher. Returns an error if the SEC code is not supported.
func NewBatch(bh *BatchHeader) (Batcher, error) {
	if bh == nil {
		return nil, errors.New("nil BatchHeader provided")
	}

	switch bh.StandardEntryClassCode {
	case ACK:
		return NewBatchACK(bh), nil
	case ADV:
		return NewBatchADV(bh), nil
	case ARC:
		return NewBatchARC(bh), nil
	case ATX:
		return NewBatchATX(bh), nil
	case BOC:
		return NewBatchBOC(bh), nil
	case CCD:
		return NewBatchCCD(bh), nil
	case CIE:
		return NewBatchCIE(bh), nil
	case COR:
		return NewBatchCOR(bh), nil
	case CTX:
		return NewBatchCTX(bh), nil
	case DNE:
		return NewBatchDNE(bh), nil
	case ENR:
		return NewBatchENR(bh), nil
	case IAT:
		return nil, ErrFileIATSEC
	case MTE:
		return NewBatchMTE(bh), nil
	case POP:
		return NewBatchPOP(bh), nil
	case POS:
		return NewBatchPOS(bh), nil
	case PPD:
		return NewBatchPPD(bh), nil
	case RCK:
		return NewBatchRCK(bh), nil
	case SHR:
		return NewBatchSHR(bh), nil
	case TEL:
		return NewBatchTEL(bh), nil
	case TRC:
		return NewBatchTRC(bh), nil
	case TRX:
		return NewBatchTRX(bh), nil
	case WEB:
		return NewBatchWEB(bh), nil
	case XCK:
		return NewBatchXCK(bh), nil
	default:
	}
	return nil, NewErrFileUnknownSEC(bh.StandardEntryClassCode)
}

// ConvertBatchType will take a batch object and convert it into one of the correct batch type
func ConvertBatchType(b Batch) Batcher {
	switch b.Header.StandardEntryClassCode {
	case ACK:
		return &BatchACK{b}
	case ADV:
		return &BatchADV{b}
	case ARC:
		return &BatchARC{b}
	case ATX:
		return &BatchATX{b}
	case BOC:
		return &BatchBOC{b}
	case CCD:
		return &BatchCCD{b}
	case CIE:
		return &BatchCIE{b}
	case COR:
		return &BatchCOR{b}
	case CTX:
		return &BatchCTX{b}
	case DNE:
		return &BatchDNE{b}
	case ENR:
		return &BatchENR{b}
	case MTE:
		return &BatchMTE{b}
	case POP:
		return &BatchPOP{b}
	case POS:
		return &BatchPOS{b}
	case PPD:
		return &BatchPPD{b}
	case RCK:
		return &BatchRCK{b}
	case SHR:
		return &BatchSHR{b}
	case TEL:
		return &BatchTEL{b}
	case TRC:
		return &BatchTRC{b}
	case TRX:
		return &BatchTRX{b}
	case WEB:
		return &BatchWEB{b}
	case XCK:
		return &BatchXCK{b}
	default:
		return &b
	}
}

// Create will tabulate and assemble an ACH batch into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate function at the end of their execution.
func (batch *Batch) Create() error {
	return errors.New("use an implementation of batch or NewBatch")
}

// Validate checks properties of the ACH batch to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the batch.
func (batch *Batch) Validate() error {
	return errors.New("use an implementation of batch or NewBatch")
}

// SetValidation stores ValidateOpts on the Batch which are to be used to override
// the default NACHA validation rules.
func (batch *Batch) SetValidation(opts *ValidateOpts) {
	if batch == nil {
		return
	}
	batch.validateOpts = opts
}

// verify checks basic valid NACHA batch rules. Assumes properly parsed records. This does not mean it is a valid batch as validity is tied to each batch type
func (batch *Batch) verify() error {
	// No entries in batch
	if len(batch.Entries) <= 0 && len(batch.ADVEntries) <= 0 {
		return batch.Error("entries", ErrBatchNoEntries)
	}
	// verify field inclusion in all the records of the batch.
	if err := batch.isFieldInclusion(); err != nil {
		// convert the field error in to a batch error for a consistent api
		return batch.Error("FieldError", err)
	}

	if !batch.IsADV() {
		// validate batch header and control codes are the same
		if (batch.validateOpts == nil || !batch.validateOpts.UnequalServiceClassCode) &&
			batch.Header.ServiceClassCode != batch.Control.ServiceClassCode {
			return batch.Error("ServiceClassCode",
				NewErrBatchHeaderControlEquality(batch.Header.ServiceClassCode, batch.Control.ServiceClassCode))
		}
		// Company Identification in the batch header and control must match if bypassCompanyIdentificationMatch is not enabled.
		if batch.Header.CompanyIdentification != batch.Control.CompanyIdentification &&
			!(batch.validateOpts != nil && batch.validateOpts.BypassCompanyIdentificationMatch) {
			return batch.Error("CompanyIdentification",
				NewErrBatchHeaderControlEquality(batch.Header.CompanyIdentification, batch.Control.CompanyIdentification))
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
	} else {
		if (batch.validateOpts == nil || !batch.validateOpts.UnequalServiceClassCode) &&
			batch.Header.ServiceClassCode != batch.ADVControl.ServiceClassCode {
			return batch.Error("ServiceClassCode",
				NewErrBatchHeaderControlEquality(batch.Header.ServiceClassCode, batch.ADVControl.ServiceClassCode))
		}
		// Control ODFIIdentification must be the same as batch header
		if batch.Header.ODFIIdentification != batch.ADVControl.ODFIIdentification {
			return batch.Error("ODFIIdentification",
				NewErrBatchHeaderControlEquality(batch.Header.ODFIIdentification, batch.ADVControl.ODFIIdentification))
		}
		// batch number header and control must match
		if batch.Header.BatchNumber != batch.ADVControl.BatchNumber {
			return batch.Error("BatchNumber",
				NewErrBatchHeaderControlEquality(batch.Header.BatchNumber, batch.ADVControl.BatchNumber))
		}
	}

	if err := batch.isBatchEntryCount(); err != nil {
		return err
	}
	if batch.validateOpts == nil || !batch.validateOpts.CustomTraceNumbers {
		if err := batch.isSequenceAscending(); err != nil {
			return err
		}
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
	if batch.validateOpts == nil || !batch.validateOpts.CustomTraceNumbers {
		if err := batch.isTraceNumberODFI(); err != nil {
			return err
		}
		if err := batch.isAddendaSequence(); err != nil {
			return err
		}
	}
	if err := batch.isCategory(); err != nil {
		return err
	}
	return nil
}

// Build creates valid batch by building sequence numbers and batch control. An error is returned if
// the batch being built has invalid records.
func (batch *Batch) build() error {
	// Requires a valid BatchHeader
	if err := batch.Header.Validate(); err != nil {
		return err
	}
	if len(batch.Entries) <= 0 && len(batch.ADVEntries) <= 0 {
		return batch.Error("entries", ErrBatchNoEntries)
	}
	// Create record sequence numbers
	entryCount := 0
	seq := 1

	if !batch.IsADV() {
		for i, entry := range batch.Entries {
			entryCount += 1 + entry.addendaCount()

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
				if opts := batch.validateOpts; opts == nil {
					entry.SetTraceNumber(batch.Header.ODFIIdentification, seq)
				} else {
					// Automatically set the TraceNumber if we are validating Origin and don't have custom trace numbers
					if !opts.BypassOriginValidation && !opts.CustomTraceNumbers {
						entry.SetTraceNumber(batch.Header.ODFIIdentification, seq)
					}
				}
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
		bc.EntryHash = batch.calculateEntryHash()
		bc.TotalCreditEntryDollarAmount, bc.TotalDebitEntryDollarAmount = batch.calculateBatchAmounts()
		batch.Control = bc
	} else {
		for i, entry := range batch.ADVEntries {
			entryCount++

			if entry.Addenda99 != nil {
				entryCount++
			}
			// Set Sequence Number
			batch.ADVEntries[i].SequenceNumber = seq

			seq++

			if seq > 9999 {
				return batch.Error("SequenceNumber", ErrBatchADVCount)
			}
		}
		// build a BatchADVControl record
		bcADV := NewADVBatchControl()
		bcADV.ServiceClassCode = batch.Header.ServiceClassCode
		bcADV.ACHOperatorData = batch.Header.CompanyName
		bcADV.ODFIIdentification = batch.Header.ODFIIdentification
		bcADV.BatchNumber = batch.Header.BatchNumber
		bcADV.EntryAddendaCount = entryCount
		bcADV.EntryHash = batch.calculateEntryHash()
		bcADV.TotalCreditEntryDollarAmount, bcADV.TotalDebitEntryDollarAmount = batch.calculateADVBatchAmounts()
		batch.ADVControl = bcADV
	}
	return batch.upsertOffsets()
}

// SetHeader appends an BatchHeader to the Batch
func (batch *Batch) SetHeader(batchHeader *BatchHeader) {
	batch.Header = batchHeader
}

// GetHeader returns the current Batch header
func (batch *Batch) GetHeader() *BatchHeader {
	return batch.Header
}

// SetControl appends an BatchControl to the Batch
func (batch *Batch) SetControl(batchControl *BatchControl) {
	batch.Control = batchControl
}

// GetControl returns the current Batch Control
func (batch *Batch) GetControl() *BatchControl {
	return batch.Control
}

// SetADVControl appends an BatchADVControl to the Batch
func (batch *Batch) SetADVControl(batchADVControl *ADVBatchControl) {
	batch.ADVControl = batchADVControl
}

// GetADVControl returns the current Batch ADVControl
func (batch *Batch) GetADVControl() *ADVBatchControl {
	return batch.ADVControl
}

// GetEntries returns a slice of entry details for the batch
func (batch *Batch) GetEntries() []*EntryDetail {
	return batch.Entries
}

// AddEntry appends an EntryDetail to the Batch
func (batch *Batch) AddEntry(entry *EntryDetail) {
	if entry == nil {
		return
	}

	batch.category = entry.Category
	batch.Entries = append(batch.Entries, entry)
}

// AddADVEntry appends an ADV EntryDetail to the Batch
func (batch *Batch) AddADVEntry(entry *ADVEntryDetail) {
	batch.category = entry.Category
	batch.ADVEntries = append(batch.ADVEntries, entry)
}

// GetADVEntries returns a slice of entry details for the batch
func (batch *Batch) GetADVEntries() []*ADVEntryDetail {
	return batch.ADVEntries
}

// Category returns batch category
func (batch *Batch) Category() string {
	if len(batch.Entries) == 0 && batch.category != "" {
		return batch.category
	}
	// If an Entry has NOC or Return that's the Batch's category
	for i := range batch.Entries {
		switch batch.Entries[i].Category {
		case CategoryReturn, CategoryNOC:
			return batch.Entries[i].Category
		}
	}
	for i := range batch.ADVEntries {
		switch batch.ADVEntries[i].Category {
		case CategoryReturn, CategoryNOC:
			return batch.ADVEntries[i].Category
		}
	}
	return CategoryForward
}

// ID returns the id of the batch
func (batch *Batch) ID() string {
	return batch.id
}

// SetID sets the batch id
func (batch *Batch) SetID(id string) {
	batch.id = id
}

// isFieldInclusion iterates through all the records in the batch and verifies against default fields
func (batch *Batch) isFieldInclusion() error {
	if err := batch.Header.Validate(); err != nil {
		return err
	}

	if !batch.IsADV() {
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
			if entry.Addenda98Refused != nil {
				if err := entry.Addenda98Refused.Validate(); err != nil {
					return err
				}
			}
			if entry.Addenda99 != nil {
				if err := entry.Addenda99.Validate(); err != nil {
					return err
				}
			}
			if entry.Addenda99Dishonored != nil {
				if err := entry.Addenda99Dishonored.Validate(); err != nil {
					return err
				}
			}
			if entry.Addenda99Contested != nil {
				if err := entry.Addenda99Contested.Validate(); err != nil {
					return err
				}
			}

		}
		return batch.Control.Validate()
	}
	// ADV File/Batch
	for _, entry := range batch.ADVEntries {
		if err := entry.Validate(); err != nil {
			return err
		}
		if entry.Addenda99 != nil {
			if err := entry.Addenda99.Validate(); err != nil {
				return err
			}
		}
	}
	return batch.ADVControl.Validate()
}

// isBatchEntryCount validate Entry count is accurate
// The Entry/Addenda Count Field is a tally of each Entry Detail and Addenda
// Record processed within the batch
func (batch *Batch) isBatchEntryCount() error {
	entryCount := 0

	if !batch.IsADV() {
		for _, entry := range batch.Entries {
			entryCount += 1 + entry.addendaCount()
		}
		if entryCount != batch.Control.EntryAddendaCount {
			if batch.validateOpts != nil && batch.validateOpts.UnequalAddendaCounts {
				return nil
			}
			return batch.Error("EntryAddendaCount",
				NewErrBatchCalculatedControlEquality(entryCount, batch.Control.EntryAddendaCount))
		}
	} else {
		for _, entry := range batch.ADVEntries {
			entryCount++
			if entry.Addenda99 != nil {
				entryCount++
			}
		}
		if entryCount != batch.ADVControl.EntryAddendaCount {
			if batch.validateOpts != nil && batch.validateOpts.UnequalAddendaCounts {
				return nil
			}
			return batch.Error("EntryAddendaCount",
				NewErrBatchCalculatedControlEquality(entryCount, batch.ADVControl.EntryAddendaCount))
		}
	}
	return nil
}

// isBatchAmount validate Amount is the same as what is in the Entries
// The Total Debit and Credit Entry Dollar Amount fields contain accumulated
// Entry Detail debit and credit totals within a given batch
func (batch *Batch) isBatchAmount() error {
	var credit, debit int

	// ToDo: Consider going back to one function for calculating BatchAmounts, but I'm not sure I want to have
	// calculateBatchAmounts with ADV TransactionCodes.  In addition the smaller functions help keep the -over for
	// gocyclo lower, although since we are currently at 25 (originally it was 18 or 19) it probably won't matter now
	// in this case.  Based on what I see in other github go code, I'm not sure 25 is a high enough number either.
	// Balancing easy to understand functions without having to create functions just for the purpose of meeting the
	// -over number convinces me that it should be higher than 25.

	if !batch.IsADV() {
		credit, debit = batch.calculateBatchAmounts()
		if debit != batch.Control.TotalDebitEntryDollarAmount {
			return batch.Error("TotalDebitEntryDollarAmount",
				NewErrBatchCalculatedControlEquality(debit, batch.Control.TotalDebitEntryDollarAmount))
		}
		if credit != batch.Control.TotalCreditEntryDollarAmount {
			return batch.Error("TotalCreditEntryDollarAmount",
				NewErrBatchCalculatedControlEquality(credit, batch.Control.TotalCreditEntryDollarAmount))
		}
	} else {
		credit, debit = batch.calculateADVBatchAmounts()
		if debit != batch.ADVControl.TotalDebitEntryDollarAmount {
			return batch.Error("TotalDebitEntryDollarAmount",
				NewErrBatchCalculatedControlEquality(debit, batch.ADVControl.TotalDebitEntryDollarAmount))
		}
		if credit != batch.ADVControl.TotalCreditEntryDollarAmount {
			return batch.Error("TotalCreditEntryDollarAmount",
				NewErrBatchCalculatedControlEquality(credit, batch.ADVControl.TotalCreditEntryDollarAmount))
		}
	}
	return nil
}

func (batch *Batch) calculateBatchAmounts() (credit int, debit int) {
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

func (batch *Batch) calculateADVBatchAmounts() (credit int, debit int) {
	for _, entry := range batch.ADVEntries {
		if entry.TransactionCode == CreditForDebitsOriginated ||
			entry.TransactionCode == CreditForCreditsReceived ||
			entry.TransactionCode == CreditForCreditsRejected ||
			entry.TransactionCode == CreditSummary {
			credit = credit + entry.Amount
		}
		if entry.TransactionCode == DebitForCreditsOriginated ||
			entry.TransactionCode == DebitForDebitsReceived ||
			entry.TransactionCode == DebitForDebitsRejectedBatches ||
			entry.TransactionCode == DebitSummary {
			debit = debit + entry.Amount
		}
	}
	return credit, debit
}

// isSequenceAscending Individual Entry Detail Records within individual batches must
// be in ascending Trace Number order (although Trace Numbers need not necessarily be consecutive).
func (batch *Batch) isSequenceAscending() error {
	if !batch.IsADV() {
		lastSeq := "0"
		for _, entry := range batch.Entries {
			if batch.validateOpts == nil || !batch.validateOpts.CustomTraceNumbers {
				if entry.TraceNumber <= lastSeq {
					return batch.Error("TraceNumber", NewErrBatchAscending(lastSeq, entry.TraceNumber))
				}
			}
			lastSeq = entry.TraceNumber
		}
	}
	return nil
}

// isEntryHash validates the hash by recalculating the result
func (batch *Batch) isEntryHash() error {

	hashField := batch.calculateEntryHash()
	if !batch.IsADV() {
		if hashField != batch.Control.EntryHash {
			return batch.Error("EntryHash",
				NewErrBatchCalculatedControlEquality(hashField, batch.Control.EntryHash))
		}
	} else {
		if hashField != batch.ADVControl.EntryHash {
			return batch.Error("EntryHash",
				NewErrBatchCalculatedControlEquality(hashField, batch.ADVControl.EntryHash))
		}
	}
	return nil
}

// calculateEntryHash This field is prepared by hashing the 8-digit Routing Number in each entry.
// The Entry Hash provides a check against inadvertent alteration of data
func (batch *Batch) calculateEntryHash() int {
	hash := 0

	if !batch.IsADV() {
		for _, entry := range batch.Entries {
			entryRDFI, _ := strconv.Atoi(aba8(entry.RDFIIdentification))
			hash += entryRDFI
		}
	} else {
		for _, entry := range batch.ADVEntries {
			entryRDFI, _ := strconv.Atoi(aba8(entry.RDFIIdentification))
			hash += entryRDFI
		}
	}

	// EntryHash is essentially the sum of all the RDFI routing numbers in the batch. If the sum exceeds 10 digits
	// (because you have lots of Entry Detail Records), lop off the most significant digits of the sum until there
	// are only 10.
	return batch.leastSignificantDigits(hash, 10)
}

// "Only an agency of the United States Government may originate a DNE entry" - NACHA Operating Rules
// Origination code '2' is for government agencies. Codes 21, 23, 31, and 33 are the only transaction codes
// allowed for DNEs. Tranaction codes 21 and 31 are just for returns or NOCs of the 23 and 33 codes.
// So we check that the Originator Status Code is not equal to “2” for DNE if the Transaction Code is 23 or 33
func (batch *Batch) isOriginatorDNE() error {
	if batch.Header.OriginatorStatusCode != 2 && batch.Header.StandardEntryClassCode == DNE {
		for _, entry := range batch.Entries {
			if entry.TransactionCode == CheckingPrenoteCredit || entry.TransactionCode == SavingsPrenoteCredit {
				return batch.Error("OriginatorStatusCode", ErrBatchOriginatorDNE, batch.Header.OriginatorStatusCode)
			}
		}
	}
	return nil
}

// isTraceNumberODFI checks if the first 8 positions of the entry detail trace number
// match the batch header ODFI
func (batch *Batch) isTraceNumberODFI() error {
	if batch.validateOpts != nil && batch.validateOpts.BypassOriginValidation {
		return nil
	}
	for _, entry := range batch.Entries {
		if batch.Header.ODFIIdentificationField() != entry.TraceNumberField()[:8] {
			return batch.Error("ODFIIdentificationField",
				NewErrBatchTraceNumberNotODFI(batch.Header.ODFIIdentificationField(), entry.TraceNumberField()[:8]))
		}
	}
	return nil
}

// isAddendaSequence check multiple errors on addenda records in the batch entries
func (batch *Batch) isAddendaSequence() error {
	for _, entry := range batch.Entries {

		if entry.Addenda02 != nil {
			if entry.AddendaRecordIndicator != 1 {
				return batch.Error("AddendaRecordIndicator", ErrBatchAddendaIndicator)
			}
		}
		if len(entry.Addenda05) > 0 {
			// addenda without indicator flag of 1
			if entry.AddendaRecordIndicator != 1 {
				return batch.Error("AddendaRecordIndicator", ErrBatchAddendaIndicator)
			}
			lastSeq := -1
			// check if sequence is ascending
			for _, a := range entry.Addenda05 {
				// sequences don't exist in NOC or Return addenda

				if a.SequenceNumber < lastSeq {
					return batch.Error("SequenceNumber", NewErrBatchAscending(lastSeq, a.SequenceNumber))
				}
				lastSeq = a.SequenceNumber
				// check that we are in the correct Entry Detail
				if !(a.EntryDetailSequenceNumberField() == entry.TraceNumberField()[8:]) {
					return batch.Error("TraceNumber", NewErrBatchAscending(lastSeq, a.SequenceNumber))
				}
			}
		}
		if entry.Addenda98 != nil {
			if entry.AddendaRecordIndicator != 1 {
				return batch.Error("AddendaRecordIndicator", ErrBatchAddendaIndicator)
			}
		}
		if entry.Addenda98Refused != nil {
			if entry.AddendaRecordIndicator != 1 {
				return batch.Error("AddendaRecordIndicator", ErrBatchAddendaIndicator)
			}
		}
		if entry.Addenda99 != nil {
			if entry.AddendaRecordIndicator != 1 {
				return batch.Error("AddendaRecordIndicator", ErrBatchAddendaIndicator)
			}
		}
		if entry.Addenda99Dishonored != nil {
			if entry.AddendaRecordIndicator != 1 {
				return batch.Error("AddendaRecordIndicator", ErrBatchAddendaIndicator)
			}
		}
		if entry.Addenda99Contested != nil {
			if entry.AddendaRecordIndicator != 1 {
				return batch.Error("AddendaRecordIndicator", ErrBatchAddendaIndicator)
			}
		}
	}
	return nil
}

// isCategory verifies that a Forward and Return Category are not in the same batch
func (batch *Batch) isCategory() error {
	if !batch.IsADV() {
		category := batch.GetEntries()[0].Category
		if len(batch.Entries) > 1 {
			for i := 0; i < len(batch.Entries); i++ {
				if batch.Entries[i].Category == CategoryNOC {
					continue
				}
				if batch.Entries[i].Category != category {
					return batch.Error("Category", NewErrBatchCategory(batch.Entries[i].Category, category))
				}
			}
		}
	} else {
		category := batch.GetADVEntries()[0].Category
		if len(batch.ADVEntries) > 1 {
			for i := 0; i < len(batch.ADVEntries); i++ {
				if batch.ADVEntries[i].Category != category {
					return batch.Error("Category", NewErrBatchCategory(batch.ADVEntries[i].Category, category))
				}
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
// Addenda99, Addenda99Dishonored, Addenda99Contested
func (batch *Batch) addendaFieldInclusion(entry *EntryDetail) error {
	switch entry.Category {
	case CategoryForward:
		if err := batch.addendaFieldInclusionForward(entry); err != nil {
			return err
		}
	case CategoryNOC:
		if err := batch.addendaFieldInclusionNOC(entry); err != nil {
			return err
		}
	case CategoryReturn, CategoryDishonoredReturn, CategoryDishonoredReturnContested:
		if err := batch.addendaFieldInclusionReturn(entry); err != nil {
			return err
		}
	}
	return nil
}

// addendaFieldInclusionForward verifies Addenda* Field Inclusion for entry.Category Forward
func (batch *Batch) addendaFieldInclusionForward(entry *EntryDetail) error {
	switch batch.Header.StandardEntryClassCode {
	case MTE, POS, SHR:
		if entry.Addenda02 == nil {
			return batch.Error("Addenda02", ErrFieldInclusion)
		}
		if entry.Addenda05 != nil {
			return batch.Error("Addenda05", ErrBatchAddendaCategory, entry.Category)
		}
	// ACK, ATX, CCD, CIE, CTX, DNE, ENR WEB, PPD, TRX can only have Addenda05
	case ACK, ATX, CCD, CIE, CTX, DNE, ENR, WEB, PPD, TRX:
		if entry.Addenda02 != nil {
			return batch.Error("Addenda02", ErrBatchAddendaCategory, entry.Category)
		}
	case ARC, BOC, COR, POP, RCK, TEL, TRC, XCK:
		if entry.Addenda02 != nil {
			return batch.Error("Addenda02", ErrBatchAddendaCategory, entry.Category)
		}
		if entry.Addenda05 != nil {
			return batch.Error("Addenda05", ErrBatchAddendaCategory, entry.Category)
		}
	}
	if batch.Header.StandardEntryClassCode != COR {
		if entry.Addenda98 != nil || entry.Addenda98Refused != nil {
			return batch.Error("Addenda98", ErrBatchAddendaCategory, entry.Category)
		}
	}
	if entry.Addenda99 != nil {
		return batch.Error("Addenda99", ErrBatchAddendaCategory, entry.Category)
	}
	return nil
}

// addendaFieldInclusionNOC verifies Addenda* Field Inclusion for entry.Category NOC
func (batch *Batch) addendaFieldInclusionNOC(entry *EntryDetail) error {
	if entry.Addenda02 != nil {
		return batch.Error("Addenda02", ErrBatchAddendaCategory, entry.Category)
	}
	if entry.Addenda05 != nil {
		return batch.Error("Addenda05", ErrBatchAddendaCategory, entry.Category)
	}
	if batch.Header.StandardEntryClassCode != COR {
		if entry.Addenda98 != nil || entry.Addenda98Refused != nil {
			return batch.Error("Addenda98", ErrFieldInclusion)
		}
	}
	if entry.Addenda99 != nil {
		return batch.Error("Addenda99", ErrBatchAddendaCategory, entry.Category)
	}
	return nil
}

// addendaFieldInclusionReturn verifies Addenda* Field Inclusion for entry.Category Return
func (batch *Batch) addendaFieldInclusionReturn(entry *EntryDetail) error {
	if entry.Addenda02 != nil {
		return batch.Error("Addenda02", ErrBatchAddendaCategory, entry.Category)
	}
	if entry.Addenda05 != nil {
		switch batch.Header.StandardEntryClassCode {
		case CTX:
			// do nothing, CTX allows Addneda05 records
		default:
			return batch.Error("Addenda05", ErrBatchAddendaCategory, entry.Category)
		}
	}
	if entry.Addenda98 != nil || entry.Addenda98Refused != nil {
		return batch.Error("Addenda98", ErrBatchAddendaCategory, entry.Category)
	}
	if entry.Addenda99 == nil && entry.Addenda99Dishonored == nil && entry.Addenda99Contested == nil {
		// Offset entries within a Return batch will not have an Addenda99 record as they might be
		// used to zero accounting entries.
		//
		// See: https://github.com/moov-io/ach/issues/1010
		if entry.IndividualName == offsetIndividualName {
			return nil
		}
		return batch.Error("Addenda99", ErrFieldInclusion)
	}
	return nil
}

// IsADV determines if a batch is batch type ADV - BatchADV
func (batch *Batch) IsADV() bool {
	ok := batch.GetHeader().StandardEntryClassCode == ADV
	return ok
}

func (batch *Batch) ValidAmountForCodes(entry *EntryDetail) error {
	if batch.validateOpts != nil && batch.validateOpts.AllowInvalidAmounts {
		return nil
	}
	if entry != nil && (entry.Addenda98 != nil || entry.Addenda98Refused != nil) {
		// NOC entries will have a zero'd amount value
		if entry.Amount != 0 {
			return ErrBatchAmountNonZero
		}
		return nil
	}

	var isPrenoteDesc bool
	if batch != nil && batch.Header != nil {
		isPrenoteDesc = strings.EqualFold(batch.Header.CompanyEntryDescription, "PRENOTE")
	}
	isPrenoteTxCode := entry.isPrenote(entry.TransactionCode)
	if isPrenoteDesc || isPrenoteTxCode {
		if entry.Amount == 0 {
			return nil
		}
		return fieldError("Amount", ErrBatchAmountNonZero, entry.Amount)
	} else {
		if entry.Amount == 0 {
			return fieldError("Amount", ErrBatchAmountZero, entry.Amount)
		}
	}
	return nil
}

// ValidTranCodeForServiceClassCode validates a TransactionCode is valid for a ServiceClassCode
func (batch *Batch) ValidTranCodeForServiceClassCode(entry *EntryDetail) error {
	// ADV should use ADVEntryDetail
	switch entry.TransactionCode {
	case CreditForDebitsOriginated, CreditForCreditsReceived, CreditForCreditsRejected, CreditSummary,
		DebitForCreditsOriginated, DebitForDebitsReceived, DebitForDebitsRejectedBatches, DebitSummary:
		return batch.Error("TransactionCode", ErrBatchTransactionCode, entry.TransactionCode)
	}

	if entry.validateOpts != nil && entry.validateOpts.CheckTransactionCode != nil {
		// We're unable to validate the ServiceClassCode with custom TransactionCode validation.
		return nil
	}

	switch batch.Header.ServiceClassCode {
	case AutomatedAccountingAdvices:
		return batch.Error("ServiceClassCode", ErrBatchServiceClassCode, batch.Header.ServiceClassCode)

	case MixedDebitsAndCredits:
		return nil

	case CreditsOnly:
		if entry.CreditOrDebit() != "C" {
			return batch.Error("TransactionCode", NewErrBatchServiceClassTranCode(batch.Header.ServiceClassCode, entry.TransactionCode))
		}
	case DebitsOnly:
		if entry.CreditOrDebit() != "D" {
			return batch.Error("TransactionCode", NewErrBatchServiceClassTranCode(batch.Header.ServiceClassCode, entry.TransactionCode))
		}
	}
	return nil
}

// Equal returns true only if two Batch (or any Batcher) objects are equal. Equality is determined by
// many of the ACH Batch and EntryDetail properties.
func (batch *Batch) Equal(other Batcher) bool {
	// Some fields are intentionally not compared as they could vary between batches that would otherwise be the same.
	if batch == nil || other == nil || batch.Header == nil || other.GetHeader() == nil {
		return false
	}
	if !batch.Header.Equal(other.GetHeader()) {
		return false
	}
	oentries := other.GetEntries()
	if len(batch.Entries) != len(oentries) {
		return false
	}
	equalEntries := 0
	for i := range batch.Entries {
		for j := range oentries {
			if batch.Entries[i].TransactionCode != oentries[j].TransactionCode {
				continue // skip to next EntryDetail
			}
			if batch.Entries[i].RDFIIdentification != oentries[j].RDFIIdentification {
				continue // skip to next EntryDetail
			}
			if batch.Entries[i].CheckDigit != oentries[j].CheckDigit {
				continue // skip to next EntryDetail
			}
			if batch.Entries[i].DFIAccountNumber != oentries[j].DFIAccountNumber {
				continue // skip to next EntryDetail
			}
			if batch.Entries[i].Amount != oentries[j].Amount {
				continue // skip to next EntryDetail
			}
			if batch.Entries[i].IdentificationNumber != oentries[j].IdentificationNumber {
				continue // skip to next EntryDetail
			}
			if batch.Entries[i].IndividualName != oentries[j].IndividualName {
				continue // skip to next EntryDetail
			}
			if batch.Entries[i].DiscretionaryData != oentries[j].DiscretionaryData {
				continue // skip to next EntryDetail
			}
			equalEntries++
		}
	}
	return len(batch.Entries) == equalEntries && equalEntries != 0
}

// WithOffset sets the Offset information onto a Batch so that during Create a balanced offset record(s) at the end of each batch.
//
// If there are debits, there is a credit offset matching the sum of the debits. If there are credits, there is a debit offset matching
// the sum of the credits. They are mutually exclusive.
func (b *Batch) WithOffset(off *Offset) {
	b.offset = off
}

const offsetIndividualName = "OFFSET"

func (b *Batch) upsertOffsets() error {
	if b == nil || b.offset == nil {
		return nil
	}
	if err := CheckRoutingNumber(b.offset.RoutingNumber); err != nil {
		return fmt.Errorf("offset: invalid routing number %s: %v", b.offset.RoutingNumber, err)
	}

	// remove any Offset records already on the batch
	for i := 0; i < len(b.Entries); i++ {
		// TODO(adam): Should we remove this based on checking the last element is
		// debit/credit and sums to all the other elements (which are mutually exclusive to
		// the last record being debit or credit)?
		// See: https://github.com/moov-io/ach/issues/540
		if strings.EqualFold(b.Entries[i].IndividualName, offsetIndividualName) {
			// fixup BatchControl records for our conditional after this for loop
			if b.Entries[i].TransactionCode == CheckingCredit || b.Entries[i].TransactionCode == SavingsCredit {
				b.Control.TotalCreditEntryDollarAmount -= b.Entries[i].Amount
			} else {
				b.Control.TotalDebitEntryDollarAmount -= b.Entries[i].Amount
			}
			// remove the EntryDetail
			b.Control.EntryAddendaCount -= 1
			b.Entries = append(b.Entries[:i], b.Entries[i+i:]...)
			i--
		}
	}

	// Make sure the offset account type is valid
	if err := b.offset.AccountType.validate(); err != nil {
		return err
	}

	offsetCount := 1

	// Create our debit offset EntryDetail
	debitED := createOffsetEntryDetail(b.offset, b)
	debitED.TraceNumber = fmt.Sprintf("%15.15d", lastTraceNumber(b.Entries)+offsetCount)
	debitED.Amount = b.Control.TotalCreditEntryDollarAmount
	switch b.offset.AccountType {
	case OffsetChecking:
		debitED.TransactionCode = CheckingDebit
	case OffsetSavings:
		debitED.TransactionCode = SavingsDebit
	}
	if debitED.Amount == 0 {
		debitED = nil // zero out so we don't add an empty OFFSET EntryDetail
	} else {
		offsetCount += 1
	}

	// Create our credit offset EntryDetail
	creditED := createOffsetEntryDetail(b.offset, b)
	creditED.TraceNumber = fmt.Sprintf("%15.15d", lastTraceNumber(b.Entries)+offsetCount)
	creditED.Amount = b.Control.TotalDebitEntryDollarAmount
	switch b.offset.AccountType {
	case OffsetChecking:
		creditED.TransactionCode = CheckingCredit
	case OffsetSavings:
		creditED.TransactionCode = SavingsCredit
	}
	if creditED.Amount == 0 {
		creditED = nil // zero out so we don't add an empty OFFSET EntryDetail
	}

	// Add both EntryDetails to our Batch and recalculate some fields
	if debitED != nil {
		b.AddEntry(debitED)
		b.Control.EntryAddendaCount += 1
		b.Control.TotalDebitEntryDollarAmount += debitED.Amount
	}
	if creditED != nil {
		b.AddEntry(creditED)
		b.Control.EntryAddendaCount += 1
		b.Control.TotalCreditEntryDollarAmount += creditED.Amount
	}
	b.Header.ServiceClassCode = MixedDebitsAndCredits

	b.Control.ServiceClassCode = MixedDebitsAndCredits
	b.Control.EntryHash = b.calculateEntryHash()

	return nil
}

func createOffsetEntryDetail(off *Offset, batch *Batch) *EntryDetail {
	ed := NewEntryDetail()
	ed.RDFIIdentification = batch.offset.RoutingNumber[:8]
	ed.CheckDigit = batch.offset.RoutingNumber[8:9]
	ed.DFIAccountNumber = batch.offset.AccountNumber
	ed.IdentificationNumber = "" // left empty
	ed.IndividualName = offsetIndividualName
	ed.DiscretionaryData = batch.offset.Description
	if len(batch.Entries) > 0 {
		ed.Category = batch.Entries[0].Category
	}
	return ed
}

// aba8 returns the first 8 digits of an ABA routing number.
// If the input is invalid then an empty string is returned.
func aba8(rtn string) string {
	n := utf8.RuneCountInString(rtn)
	switch {
	case n > 10:
		return ""
	case n == 10:
		if rtn[0] == '0' || rtn[0] == '1' {
			return rtn[1:9] // ACH server will prefix with space, 0, or 1
		}
		return ""
	case n != 8 && n != 9:
		return ""
	default:
		return rtn[:8]
	}
}

func lastTraceNumber(entries []*EntryDetail) int {
	if len(entries) == 0 {
		return 0
	}
	n, err := strconv.Atoi(entries[len(entries)-1].TraceNumber)
	if err != nil {
		return 0
	}
	return n
}
