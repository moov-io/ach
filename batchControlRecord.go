package ach

// BatchControlRecord contains entry counts, dollar total and has totals for all
// entries contained in the preceding batch
type BatchControlRecord struct {
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
	// to hardware failure or program error.
	//
	// In this context the Entry Hash is the sum of the corresponding fields in the
	// Entry Detail Records on the file.
	EntryHash string
	// TotalDebitEntryDollarAmount Contains accumulated Entry debit totals within the batch.
	TotalDebitEntryDollarAmount string
	// TotalCreditEntryDollarAmount Contains accumulated Entry credit totals within the batch.
	TotalCreditEntryDollarAmount string
	// CompanyIdentification is an alphameric code used to identify an Originator.
	// The Company Identification Field must be included on all
	// prenotification records and on each entry initiated puruant to such
	// prenotification. The Company ID may begin with the ANSI one-digit
	// Identification Code Designators (ICD), followed by the identification
	// number. The ANSI Identification Numbers and related Identification Code
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
