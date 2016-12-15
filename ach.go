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
	FileHeader  FileHeaderRecord
	BatchHeader BatchHeaderRecord
}

// FileHeaderRecord designate physical file characteristics and identify
// the origin (sending point) and destination (receiving point) of the entries
// contained in the file. The file header also includes creation date and time
// fields which can be used to uniquely identify a file.
type FileHeaderRecord struct {
	// RecordType defines the type of record in the block. FILE_HEADER
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
}

// EntryDetailRecord contains the actual transaction data for an individual entry.
// Fields include those designating the entry as a deposit (credit) or
// withdrawal (debit), the transit routing number for the entry recipient’s financial
// institution, the account number (left justify,no zero fill), name, and dollar amount.
type EntryDetailRecord struct{}

// AddendaRecord provides business transaction information in a machine
// readable format. It is usually formatted according to ANSI, ASC, X12 Standard
type AddendaRecord struct {
}

// BatchControlRecord contains entry counts, dollar total and has totals for all
// entries contained in the preceding batch
type BatchControlRecord struct {
}

// FileControlRecord This record contains entry counts, dollar totals and hash
// totals accumulated from each batch control record in the file.
type FileControlRecord struct {
}
