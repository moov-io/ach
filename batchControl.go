package ach

import (
	"fmt"
	"strconv"
)

// BatchControl contains entry counts, dollar total and has totals for all
// entries contained in the preceding batch
type BatchControl struct {
	// RecordType defines the type of record in the block. batchControlPos 8
	recordType string
	// ServiceClassCode ACH Mixed Debits and Credits ‘200’
	// ACH Credits Only ‘220’
	// ACH Debits Only ‘225'
	// Same as 'ServiceClassCode' in BatchHeaderRecord
	ServiceClassCode int
	// EntryAddendaCount is a tally of each Entry Detail Record and each Addenda
	// Record processed, within either the batch or file as appropriate.
	entryAddendaCount int
	// EntryHash the Receiving DFI Identification in each Entry Detail Record is hashed
	// to provide a check against inadvertent alteration of data contents due
	// to hardware failure or program erro
	//
	// In this context the Entry Hash is the sum of the corresponding fields in the
	// Entry Detail Records on the file.
	entryHash int
	// TotalDebitEntryDollarAmount Contains accumulated Entry debit totals within the batch.
	totalDebitEntryDollarAmount int
	// TotalCreditEntryDollarAmount Contains accumulated Entry credit totals within the batch.
	totalCreditEntryDollarAmount int
	// CompanyIdentification is an alphameric code used to identify an Originato
	// The Company Identification Field must be included on all
	// prenotification records and on each entry initiated puruant to such
	// prenotification. The Company ID may begin with the ANSI one-digit
	// Identification Code Designators (ICD), followed by the identification
	// numbe The ANSI Identification Numbers and related Identification Code
	// Designators (ICD) are:
	//
	// IRS Employer Identification Number (EIN) "1"
	// Data Universal Numbering Systems (DUNS) "3"
	// User Assigned Number "9"
	CompanyIdentification string
	// MessageAuthenticationCode the MAC is an eight character code derived from a special key used in
	// conjunction with the DES algorithm. The purpose of the MAC is to
	// validate the authenticity of ACH entries. The DES algorithm and key
	// message standards must be in accordance with standards adopted by the
	// American National Standards Institute. The remaining eleven characters
	// of this field are blank.
	MessageAuthenticationCode string
	// Reserved for the future - Blank, 6 characters long
	Reserved string
	// OdfiIdentification the routing number is used to identify the DFI originating entries within a given branch.
	OdfiIdentification string
	// BatchNumber this number is assigned in ascending sequence to each batch by the ODFI
	// or its Sending Point in a given file of entries. Since the batch number
	// in the Batch Header Record and the Batch Control Record is the same,
	// the ascending sequence number should be assigned by batch and not by record.
	batchNumber int
	// Validator is composed for data validation
	Validator
	// Converters is composed for ACH to golang Converters
	Converters
}

// Parse takes the input record string and parses the EntryDetail values
func (bc *BatchControl) Parse(record string) {
	// 1-1 Always "8"
	bc.recordType = record[:1]
	// 2-4 This is the same as the "Service code" field in previous Batch Header Record
	bc.ServiceClassCode = bc.parseNumField(record[1:4])
	// 5-10 Total number of Entry Detail Record in the batch
	bc.entryAddendaCount = bc.parseNumField(record[4:10])
	// 11-20 Total of all positions 4-11 on each Entry Detail Record in the batch. This is essentially the sum of all the RDFI routing numbers in the batch.
	// If the sum exceeds 10 digits (because you have lots of Entry Detail Records), lop off the most significant digits of the sum until there are only 10
	bc.entryHash = bc.parseNumField(record[10:20])
	// 21-32 Number of cents of debit entries within the batch
	bc.totalDebitEntryDollarAmount = bc.parseNumField(record[20:32])
	// 33-44 Number of cents of credit entries within the batch
	bc.totalCreditEntryDollarAmount = bc.parseNumField(record[32:44])
	// 45-54 This is the same as the "Company identification" field in previous Batch Header Record
	bc.CompanyIdentification = record[44:54]
	// 55-73 Seems to always be blank
	bc.MessageAuthenticationCode = record[54:73]
	// 74-79 Always blank (just fill with spaces)
	bc.Reserved = record[73:79]
	// 80-87 This is the same as the "ODFI identification" field in previous Batch Header Record
	bc.OdfiIdentification = record[79:87]
	// 88-94 This is the same as the "Batch number" field in previous Batch Header Record
	bc.batchNumber = bc.parseNumField(record[87:94])
}

// NewBatchControl returns a new BatchControl with default values for none exported fields
func NewBatchControl() *BatchControl {
	return &BatchControl{
		recordType: "8",
	}
}

// String writes the BatchControl struct to a 94 character string.
func (bc *BatchControl) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v",
		bc.recordType,
		bc.ServiceClassCode,
		bc.EntryAddendaCount(),
		bc.EntryHash(),
		bc.TotalDebitEntryDollarAmount(),
		bc.TotalCreditEntryDollarAmount(),
		bc.CompanyIdentification,
		bc.MessageAuthenticationCode,
		bc.Reserved,
		bc.OdfiIdentification,
		bc.BatchNumber(),
	)
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (bc *BatchControl) Validate() (bool, error) {
	if bc.recordType != "8" {
		return false, ErrRecordType
	}
	if !bc.isServiceClass(bc.ServiceClassCode) {
		return false, ErrServiceClass
	}
	return true, nil
}

// EntryAddendaCount gets a string of the addenda count zero padded
func (bc *BatchControl) EntryAddendaCount() string {
	return bc.leftPad(strconv.Itoa(bc.entryAddendaCount), "0", 6)
}

// EntryHash get a zero padded EntryHash
func (bc *BatchControl) EntryHash() string {
	return bc.leftPad(strconv.Itoa(bc.entryHash), "0", 10)
}

//TotalDebitEntryDollarAmount get a zero padded Debity Entry Amount
func (bc *BatchControl) TotalDebitEntryDollarAmount() string {
	return bc.leftPad(strconv.Itoa(bc.totalDebitEntryDollarAmount), "0", 12)
}

// TotalCreditEntryDollarAmount get a zero padded Credit Entry Amount
func (bc *BatchControl) TotalCreditEntryDollarAmount() string {
	return bc.leftPad(strconv.Itoa(bc.totalCreditEntryDollarAmount), "0", 12)
}

// BatchNumber gets a string of the batch number zero padded
func (bc *BatchControl) BatchNumber() string {
	return bc.leftPad(strconv.Itoa(bc.batchNumber), "0", 7)
}
