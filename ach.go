// Copyright 2016 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// Package ach reads and writes (ACH) Automated Clearing House files. ACH is the
// primary method of electronic money movemenet through the United States.
// https://en.wikipedia.org/wiki/Automated_Clearing_House
// Their are several kinds of ACH files PPD, PPD+, CCD, CCD+, and CTX; this
// package currently supports the (PPD) Prearranged Payment and Deposit Entry
// type transactions described by NACHA Operating Rules
// https://www.nacha.org//
// An ACH PPD format is a file with multiple lines of ACII text, each line is 94
// characters in length. A line is called a "record" in ACH paralance.
// There are five main record types in an ACH file:
//
// • File Header Record
// • Company/Batch Header Record
// • Entry Detail Record
// • Company/Batch Control Record
// • File Trailer Record
package ach

// First position of all Record Types. These codes are uniquily assigned to
// the first byte of each row in a file.
const (
	headerPos       = "1"
	batchPos        = "5"
	entryDetailPos  = "6"
	entryAgendaPos  = "7"
	batchControlPos = "8"
	fileControlPos  = "9"
)

// ACH contains the structures of a parsed ACH File.
type ACH struct {
	FileHeader   FileHeaderRecord
	BatchHeader  BatchHeaderRecord
	EntryDetail  EntryDetailRecord
	BatchControl BatchControlRecord
}

// FileHeaderRecord designate physical file characteristics and identify
// the origin (sending point) and destination (receiving point) of the entries
// contained in the file. The file header also includes creation date and time
// fields which can be used to uniquely identify a file.
type FileHeaderRecord struct {
	// RecordType defines the type of record in the block. headerPos
	RecordType string

	// PriorityCode conists of the numerals 01
	PriorityCode string

	// ImmediateDestination contains the Routing Number of the ACH Operator or receiving
	// point to which the file is being sent. The 10 character field begins with
	// a blank in the first position, followed by the four digit Federal Reserve
	// Routing Symbol, the four digit ABA Institution Identifier, and the Check
	// Digit (bTTTTAAAAC).
	ImmediateDestination string

	// ImmediateOrigin contains the Routing Number of the ACH Operator or sending
	// point that is sending the file. The 10 character field begins with
	// a blank in the first position, followed by the four digit Federal Reserve
	// Routing Symbol, the four digit ABA Insitution Identifier, and the Check
	// Digit (bTTTTAAAAC).
	ImmediateOrigin string

	// FileCreationDate is expressed in a "YYMMDD" format. The File Creation
	// Date is the date on which the file is prepared by an ODFI (ACH input files)
	// or the date (exchange date) on which a file is transmitted from ACH Operator
	// to ACH Operator, or from ACH Operator to RDFIs (ACH output files).
	FileCreationDate string

	// FileCreationTime is expressed ina n "HHMM" (24 hour clock) format.
	// The system time when the ACH file was created
	FileCreationTime string

	// This field should start at zero and increment by 1 (up to 9) and then go to
	// letters starting at A through Z for each subsequent file that is created for
	// a single system date. (34-34) 1 numeric 0-9 or uppercase alpha A-Z.
	// I have yet to see this ID not A
	FileIdModifier string

	// RecordSize indicates the number of characters contained in each
	// record. At this time, the value "094" must be used.
	RecordSize string

	// BlockingFactor defines the number of physical records within a block
	// (a block is 940 characters). For all files moving between a DFI and an ACH
	// Operator (either way), the value "10" must be used. If the number of records
	// within the file is not a multiple of ten, the remainder of the block must
	// be nine-filled.
	BlockingFactor string

	// FormatCode a code to allow for future format variations. As
	// currently defined, this field will contain a value of "1".
	FormatCode string

	// ImmediateDestinationName us the name of the ACH or receiving point for which that
	// file is destined. Name corresponding to the ImmediateDestination
	ImmediateDestinationName string

	// ImmidiateOriginName is the name of the ACH operator or sending point that is
	// sending the file. Name corresponding to the ImmediateOrigin
	ImmidiateOriginName string

	// ReferenceCode is reserved for information pertinent to the Originator.
	ReferenceCode string
}

// BatchHeaderRecord identifies the originating entity and the type of transactions
// contained in the batch (i.e., the standard entry class, PPD for consumer, CCD
// or CTX for corporate). This record also contains the effective date, or desired
// settlement date, for all entries contained in this batch. The settlement date
// field is not entered as it is determined by the ACH operator.
type BatchHeaderRecord struct {
	// RecordType defines the type of record in the block. 5
	RecordType string
	// ServiceClassCode ACH Mixed Debits and Credits ‘200’
	// ACH Credits Only ‘220’
	// ACH Debits Only ‘225'
	ServiceClassCode string
	// CompanyName the company originating the entries in the batch
	CompanyName string

	// CompanyDiscretionaryData allows Originators and/or ODFIs to include codes (one or more),
	// of significance only to them, to enable specialized handling of all
	// subsequent entries in that batch. There will be no standardized
	// interpretation for the value of the field. This field must be returned
	// intact on any return entry.
	CompanyDiscretionaryData string
	// CompanyIdentification The 9 digit FEIN number (proceeded by a predetermined
	// alpha or numeric character) of the entity in the company name field
	CompanyIdentification string

	// StandardEntryClassCode PPD’ for consumer transactions, ‘CCD’ or ‘CTX’ for corporate
	StandardEntryClassCode string

	// CompanyEntryDescription A description of the entries contained in the batch
	//
	//The Originator establishes the value of this field to provide a
	// description of the purpose of the entry to be displayed back to
	// the receiver. For example, "GAS BILL," "REG. SALARY," "INS. PREM,"
	// "SOC. SEC.," "DTC," "TRADE PAY," "PURCHASE," etc.
	//
	// This field must contain the word "REVERSAL" (left justified) when the
	// batch contains reversing entries.
	//
	// This field must contain the word "RECLAIM" (left justified) when the
	// batch contains reclamation entries.
	//
	// This field must contain the word "NONSETTLED" (left justified) when the
	// batch contains entries which could not settle.
	CompanyEntryDescription string

	// CompanyDescriptiveDate except as otherwise noted below, the Originator establishes this field
	// as the date it would like to see displayed to the receiver for
	// descriptive purposes. This field is never used to control timing of any
	// computer or manual operation. It is solely for descriptive purposes.
	// The RDFI should not assume any specific format. Examples of possible
	// entries in this field are "011392,", "01 92," "JAN 13," "JAN 92," etc.
	CompanyDescriptiveDate string

	// EffectiveEntryDate the date on which the entries are to settle
	EffectiveEntryDate string

	// SettlementDate Leave blank, this field is inserted by the ACH operator
	SettlementDate string

	// OriginatorStatusCode '1'
	OriginatorStatusCode string

	//OdfiIdentification First 8 digits of the originating DFI transit routing number
	OdfiIdentification string

	// BatchNumber Sequential batch number, zero fill
	//
	// This number is assigned in ascending sequence to each batch by the ODFI
	// or its Sending Point in a given file of entries. Since the batch number
	// in the Batch Header Record and the Batch Control Record is the same,
	// the ascending sequence number should be assigned by batch and not by
	// record.
	BatchNumber string
}

// EntryDetailRecord contains the actual transaction data for an individual entry.
// Fields include those designating the entry as a deposit (credit) or
// withdrawal (debit), the transit routing number for the entry recipient’s financial
// institution, the account number (left justify,no zero fill), name, and dollar amount.
type EntryDetailRecord struct {
	// RecordType defines the type of record in the block. 6
	RecordType string
	// TransactionCode if the recievers account is:
	// Credit (deposit) to checking account ‘22’
	// Prenote for credit to checking account ‘23’
	// Debit (withdrawal) to checking account ‘27’
	// Prenote for debit to checking account ‘28’
	// Credit to savings account ‘32’
	// Prenote for credit to savings account ‘33’
	// Debit to savings account ‘37’
	// Prenote for debit to savings account ‘38’
	TransactionCode string

	// RoutingNumber is the RDFI's routing number without the last digit.
	RdfiIdentification string

	// CheckDigit the last digit of the RDFI's routing number
	CheckDigit string

	// The receiver's bank account number you are crediting/debiting.
	// It important to note that this is an alphanumeric field, so its space padded, no zero padded
	DfiAccountNumber string

	// Amount Number of cents you are debiting/crediting this account
	Amount string

	// IndividualIdentificationNumber n internal identification (alphanumeric) that
	// you use to uniquely identify this Entry Detail Record
	IndividualIdentificationNumber string

	// IndividualName The name of the receiver, usually the name on the bank account
	IndividualName string

	// DiscretionaryData allows ODFIs to include codes, of significance only to them,
	// to enable specialized handling of the entry. There will be no
	// standardized interpretation for the value of this field. It can either
	// be a single two-character code, or two distince one-character codes,
	// according to the needs of the ODFI and/or Originator involved. This
	// field must be returned intact for any returned entry.
	DiscretionaryData string

	// AddendaRecordIndicator indicates the existence of an Addenda Record.
	// A value of "1" indicates that one ore more addenda records follow,
	// and "0" means no such record is present.
	AddendaRecordIndicator string

	// TraceNumber assigned by the ODFI in ascending sequence, is included in each
	// Entry Detail Record, Corporate Entry Detail Record, and addenda Record.
	// Trace Numbers uniquely identify each entry within a batch in an ACH input file.
	// In association with the Batch Number, transmission (File Creation) Date,
	// and File ID Modifier, the Trace Number uniquely identifies an entry within a given file.
	// For addenda Records, the Trace Number will be identical to the Trace Number
	// in the associated Entry Detail Record, since the Trace Number is associated
	// with an entry or item rather than a physical record.
	TraceNumber string

	Addenda AddendaRecord
}

// AddendaRecord provides business transaction information in a machine
// readable format. It is usually formatted according to ANSI, ASC, X12 Standard
type AddendaRecord struct {
	// TODO implement structure
	RecordType                string
	TypeCode                  string
	PaymentRelatedInformation string
	SequenceNumber            string
	EntryDetailSequenceNumber string
}

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

// FileControlRecord record contains entry counts, dollar totals and hash
// totals accumulated from each batch control record in the file.
type FileControlRecord struct {
	// RecordType defines the type of record in the block. fileControlPos
	RecordType string

	// BatchCount total number of batches (i.e., ‘5’ records) in the file
	BatchCount string

	// BlockCount total number of records in the file (include all headers and trailer) divided
	// by 10 (This number must be evenly divisible by 10. If not, additional records consisting of all 9’s are added to the file after the initial ‘9’ record to fill out the block 10.)
	BlockCount string

	// EntryAddendaCount total detail and addenda records in the file
	EntryAddendaCount string

	// EntryHash calculated in the same manner as the batch has total but includes total from entire file
	EntryHash string

	// TotalDebitEntryDollarAmountInFile contains accumulated Batch debit totals within the file.
	TotalDebitEntryDollarAmountInFile string

	// TotalCreditEntryDollarAmountInFile contains accumulated Batch credit totals within the file.
	TotalCreditEntryDollarAmountInFile string

	// Reserved should be blank.
	Reserved string
}
