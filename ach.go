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
//
//	Entry Detail Record - The Entry Detail Records contain the information
//	necessary to route the entry to the Receiver (i.e., the Receiver's financial
//	institution, account number, account type, receiving name, and the debit or
//	credit amount.)
//
// 	Addenda Record (CCD+ not supported)- This record provides business
//	transaction information in a machine-readable format. It is usually
//	formatted according to ANSI, ASC, X12 Standard
//
//	Company/Batch Control Record - This record contains the counts, hash totals,
//	and total dollar controls for the preceding detail entries within the
//	indicated batch
//
//	File Control Record - This record contains dollar, entry, and hash total
//	accumulations from all Company/Batch Control Records in the file. This
//	record also contains counts of the number of records and the number of
//	batches within the file

package ach

// First position of all Record Types. These codes are uniquily assigned to
// the first byte of each row in a file.
const (
	FILE_HEADER   = "1"
	BATCH_HEADER  = "5"
	ENTRY_DETAIL  = "6"
	ENTRY_ADDENDA = "7"
	BATCH_CONTROL = "8"
	FILE_CONTROL  = "9"
)

// ACH contains the structures of a parsed ACH File.
type ACH struct {
	FileHeader FileHeaderRecord
}

// A FileHeader is eactly one in each ACH file, and its always the first record
// in the file. The File Header Record contains high-level information about the
// contents of the ACH file. This record designates the physical file characteristics
//	and identifies the immediate origin and destination of the entries contained
//	within the file. In addition, this record includes date, time, and file
//	identification fields used to identify the file uniquely.
type FileHeaderRecord struct {
	RecordType               string // (character position 1-1) Always "1"
	PriorityCode             string // (2-3) Always "01"
	ImmediateDestination     string // (4-13) A blank space followed by your ODFI's routing number. For example: " 121140399"
	ImmediateOrigin          string // (14-23) A 10-digit number assigned to you by the ODFI once they approve you to originate ACH files through them
	FileCreationDate         string // 24-29 Today's date in YYMMDD format
	FileCreationTime         string // 30-33 The current time in HHMM format
	FileIdModifier           string // 35-37 Always "A"
	RecordSize               string // 35-37 always "094"
	BlockingFactor           string //38-39 always "10"
	FormatCode               string //40 always "1"
	ImmediateDestinationName string //41-63 The name of the ODFI. example "SILICON VALLEY BANK    "
	ImmidiateOriginName      string //64-86 ACH operator or sending point that is sending the file
	ReferenceCode            string //97-94 Optional field that may be used to describe the ACH file for internal accounting purposes
}
