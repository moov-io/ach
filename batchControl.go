package ach

// BatchControl contains entry counts, dollar total and has totals for all
// entries contained in the preceding batch
type BatchControl struct {
	// RecordType defines the type of record in the block. batchControlPos 8
	RecordType string
	// ServiceClassCode ACH Mixed Debits and Credits ‘200’
	// ACH Credits Only ‘220’
	// ACH Debits Only ‘225'
	// Sane as 'ServiceClassCode' in BatchHeaderRecord
	ServiceClassCode string
	// EntryAddendaCount is a tally of each Entry Detail Record and each Addenda
	// Record processed, within either the batch or file as appropriate.
	EntryAddendaCount string
	// EntryHash the Receiving DFI Identification in each Entry Detail Record is hashed
	// to provide a check against inadvertent alteration of data contents due
	// to hardware failure or program erro
	//
	// In this context the Entry Hash is the sum of the corresponding fields in the
	// Entry Detail Records on the file.
	EntryHash string
	// TotalDebitEntryDollarAmount Contains accumulated Entry debit totals within the batch.
	TotalDebitEntryDollarAmount string
	// TotalCreditEntryDollarAmount Contains accumulated Entry credit totals within the batch.
	TotalCreditEntryDollarAmount string
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
	BatchNumber string
}

// Parse takes the input record string and parses the EntryDetail values
func (bc *BatchControl) Parse(record string) {
	// 1-1 Always "8"
	bc.RecordType = record[:1]
	// 2-4 This is the same as the "Service code" field in previous Batch Header Record
	bc.ServiceClassCode = record[1:4]
	// 5-10 Total number of Entry Detail Record in the batch
	bc.EntryAddendaCount = record[4:10]
	// 11-20 Total of all positions 4-11 on each Entry Detail Record in the batch. This is essentially the sum of all the RDFI routing numbers in the batch.
	// If the sum exceeds 10 digits (because you have lots of Entry Detail Records), lop off the most significant digits of the sum until there are only 10
	bc.EntryHash = record[10:20]
	// 21-32 Number of cents of debit entries within the batch
	bc.TotalDebitEntryDollarAmount = record[20:32]
	// 33-44 Number of cents of credit entries within the batch
	bc.TotalCreditEntryDollarAmount = record[32:44]
	// 45-54 This is the same as the "Company identification" field in previous Batch Header Record
	bc.CompanyIdentification = record[44:54]
	// 55-73 Seems to always be blank
	bc.MessageAuthenticationCode = record[54:73]
	// 74-79 Always blank (just fill with spaces)
	bc.Reserved = record[73:79]
	// 80-87 This is the same as the "ODFI identification" field in previous Batch Header Record
	bc.OdfiIdentification = record[79:87]
	// 88-94 This is the same as the "Batch number" field in previous Batch Header Record
	bc.BatchNumber = record[87:94]
}

// NewBatchControl returns a new BatchControl with default values for none exported fields
func NewBatchControl() *BatchControl {
	return &BatchControl{}
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (bc *BatchControl) Validate() (bool, error) {
	return true, nil
}
