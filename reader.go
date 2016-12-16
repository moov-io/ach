// Copyright 2016 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.
//
package ach

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
)

// A Reader decodes records from a NACHA ACH ecoded file
//
// Data Specifications Within the ACH File
// • Each line is a Record and consists of 94 characters in
// • Fields within each record type are alphabetic, numeric or alphameric.
// • All alphabetic fields must be left justified and blank padded/filled.
// • All alphabetic characters must be in upper case or "caps".
// • All numeric fields must be right justified, unsigned, and zero padded/filled.
// • All records are 94 characters in length.
// • The file's blocking factor is '10', as indicated in positions 38-39 of the
// File Header '1' record. Every 10 records are a block. If the number of
// records within the file is not a multiple of 10, the remainder of the block
// must be nine filled. The total number of records in your file must be evenly
// divisible by 10.

const (
	RecordLength = 94
)

// A ParseError is returned for parsing errors.
// The first line is 1.
type ParseError struct {
	Line    int   // line where the error occurred
	CharPos int   // The character position where the error was found
	Err     error // The actual error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("line %d, character position %d: %s", e.Line, e.CharPos, e.Err)
}

// These are the errors that can be returned in Parse.Error
var (
	ErrFieldCount = errors.New("wrong number of fields in record expect 94")
)

// A decoder reads records from a ACH-encoded file.
type decoder struct {
	line int
	//charpos int
	r *bufio.Reader
	// lineBuffer holds the unescaped content read by readRecord
	lineBuffer bytes.Buffer
	//Indexes of the fields inside of lineBuffer
	fieldIndexes []int
	header       FileHeaderRecord
}

// error creates a new ParseError based on err.
func (d *decoder) error(err error) error {
	return &ParseError{
		Line: d.line,
		//CharPos: d.charpos,
		Err: err,
	}
}

// Decode reads a ACH file from r and returns it as a ach.ACH
func Decode(r io.Reader) (ach ACH, err error) {
	var d decoder
	return d.decode(r)
}

// Decode reads each line of the ACH file and defines which parser to use based
// on the first byte of each line. It also enforces ACH formating rules and returns
// the appropriate error if issues are found.
func (d *decoder) decode(r io.Reader) (ach ACH, err error) {

	d.r = bufio.NewReader(r)

	// read through the entire file
	for {
		line, _, err := d.r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("%s when reading file", err)
				// TODO: return nil, error
			}
		}
		// Only 94 ASCII characters to a line

		if len(line) != RecordLength {
			return ach, d.error(ErrFieldCount)
		}
		// TODO: Check that all characters are accepted.

		d.line++

		record := string(line)
		switch record[:1] {
		case headerPos:
			ach.FileHeader = parseFileHeader(record)
		case batchPos:
			ach.BatchHeader = parseBatchHeader(record)
		case entryDetailPos:
			ach.EntryDetail = parseEntryDetail(record)
		case entryAddendaPos:
			// TODO: implement parseEntryAgenda
			//ach.EntryAgenda = parseEntryAgenda(record)
		case batchControlPos:
			ach.BatchControl = parseBatchControl(record)
		case fileControlPos:
			ach.FileControl = parseFileControl(record)
		default:
			//fmt.Println("Record type not detected")
			// TODO: return nil, error
		}

	}
	// TODO: number of lines in file must be divisable by 10 the blocking factor
	//fmt.Printf("Number of lines in file: %v \n", d.line)
	return ach, nil
}

// parseFileHeader takes the input record string and parses the FileHeaderRecord values
func parseFileHeader(record string) (fileHeader FileHeaderRecord) {
	// (character position 1-1) Always "1"
	fileHeader.RecordType = record[:1]
	// (2-3) Always "01"
	fileHeader.PriorityCode = record[1:3]
	// (4-13) A blank space followed by your ODFI's routing number. For example: " 121140399"
	fileHeader.ImmediateDestination = record[3:13]
	// (14-23) A 10-digit number assigned to you by the ODFI once they approve you to originate ACH files through them
	fileHeader.ImmediateOrigin = record[13:23]
	// 24-29 Today's date in YYMMDD format
	fileHeader.FileCreationDate = record[23:29]
	// 30-33 The current time in HHMM format
	fileHeader.FileCreationTime = record[29:33]
	// 35-37 Always "A"
	fileHeader.FileIdModifier = record[33:34]
	// 35-37 always "094"
	fileHeader.RecordSize = record[34:37]
	//38-39 always "10"
	fileHeader.BlockingFactor = record[37:39]
	//40 always "1"
	fileHeader.FormatCode = record[39:40]
	//41-63 The name of the ODFI. example "SILICON VALLEY BANK    "
	fileHeader.ImmediateDestinationName = record[40:63]
	//64-86 ACH operator or sending point that is sending the file
	fileHeader.ImmidiateOriginName = record[63:86]
	//97-94 Optional field that may be used to describe the ACH file for internal accounting purposes
	fileHeader.ReferenceCode = record[86:94]

	return fileHeader
}

// parseBatchHeader takes the input record string and parses the FileHeaderRecord values
func parseBatchHeader(record string) (batchHeader BatchHeaderRecord) {
	// 1-1 Always "5"
	batchHeader.RecordType = record[:1]
	// 2-4 If the entries are credits, always "220". If the entries are debits, always "225"
	batchHeader.ServiceClassCode = record[1:4]
	// 5-20 Your company's name. This name may appear on the receivers’ statements prepared by the RDFI.
	batchHeader.CompanyName = record[4:20]
	// 21-40 Optional field you may use to describe the batch for internal accounting purposes
	batchHeader.CompanyDiscretionaryData = record[20:40]
	// 41-50 A 10-digit number assigned to you by the ODFI once they approve you to
	// originate ACH files through them. This is the same as the "Immediate origin" field in File Header Record
	batchHeader.CompanyIdentification = record[40:50]
	// 51-53 If the entries are PPD (credits/debits towards consumer account), use "PPD".
	// If the entries are CCD (credits/debits towards corporate account), use "CCD".
	// The difference between the 2 class codes are outside of the scope of this post, but generally most ACH transfers to consumer bank accounts should use "PPD"
	batchHeader.StandardEntryClassCode = record[50:53]
	// 54-63 Your description of the transaction. This text will appear on the receivers’ bank statement.
	// For example: "Payroll   "
	batchHeader.CompanyEntryDescription = record[53:63]
	// 64-69 The date you choose to identify the transactions in YYMMDD format.
	// This date may be printed on the receivers’ bank statement by the RDFI
	batchHeader.CompanyDescriptiveDate = record[63:69]
	// 70-75 Date transactions are to be posted to the receivers’ account.
	// You almost always want the transaction to post as soon as possible, so put tomorrow's date in YYMMDD format
	batchHeader.EffectiveEntryDate = record[69:75]
	// 76-79 Always blank (just fill with spaces)
	batchHeader.SettlementDate = record[75:78]
	// 79-79 Always 1
	batchHeader.OriginatorStatusCode = record[78:79]
	// 80-87 Your ODFI's routing number without the last digit. The last digit is simply a
	// checksum digit, which is why it is not necessary
	batchHeader.OdfiIdentification = record[79:87]
	// 88-94 Sequential number of this Batch Header Record.
	// For example, put "1" if this is the first Batch Header Record in the file
	batchHeader.BatchNumber = record[87:94]

	return batchHeader
}

// parseEntryDetail takes the input record string and parses the EntryDetailRecord values
func parseEntryDetail(record string) (entryDetail EntryDetailRecord) {
	// 1-1 Always "6"
	entryDetail.RecordType = record[:1]
	// 2-3 is checking credit 22 debit 27 savings credit 32 debit 37
	entryDetail.TransactionCode = record[1:3]
	// 4-11 the RDFI's routing number without the last digit.
	entryDetail.RdfiIdentification = record[3:11]
	// 12-12 The last digit of the RDFI's routing number
	entryDetail.CheckDigit = record[11:12]
	// 13-29 The receiver's bank account number you are crediting/debiting
	entryDetail.DfiAccountNumber = record[12:29]
	// 30-39 Number of cents you are debiting/crediting this account
	entryDetail.Amount = record[29:39]
	// 40-54 An internal identification (alphanumeric) that you use to uniquely identify this Entry Detail Record
	entryDetail.IndividualIdentificationNumber = record[39:54]
	// 55-76 The name of the receiver, usually the name on the bank account
	entryDetail.IndividualName = record[54:76]
	// 77-78 allows ODFIs to include codes of significance only to them
	// normally blank
	entryDetail.DiscretionaryData = record[76:78]
	// 79-79 1 if addenda exists 0 if it does not
	entryDetail.AddendaRecordIndicator = record[78:79]
	// 80-84 An internal identification (alphanumeric) that you use to uniquely identify
	// this Entry Detail Record. This number should be unique to the transaction and will help identify the transaction in case of an inquiry
	entryDetail.TraceNumber = record[79:94]

	return entryDetail
}

// parseAddendaRecord takes the input record string and parses the AddendaRecord values
func parseAddendaRecord(record string) (addenda AddendaRecord) {
	// 1-1 Always "7"
	addenda.RecordType = record[:1]
	// 2-3 Defines the specific explanation and format for the addenda information contained in the same record
	addenda.TypeCode = record[1:3]
	// 4-83 Based on the information entered. (04-83) 80 alphanumeric
	addenda.PaymentRelatedInformation = record[3:83]
	// 84-87 SequenceNumber is consecutively assigned to each Addenda Record following
	// an Entry Detail Record
	addenda.SequenceNumber = record[83:87]
	// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
	addenda.EntryDetailSequenceNumber = record[87:94]

	return addenda
}

// parseBatchControl takes the input record string and parses the BatchControlRecord values
func parseBatchControl(record string) (batchControl BatchControlRecord) {
	// 1-1 Always "8"
	batchControl.RecordType = record[:1]
	// 2-4 This is the same as the "Service code" field in previous Batch Header Record
	batchControl.ServiceClassCode = record[1:4]
	// 5-10 Total number of Entry Detail Record in the batch
	batchControl.EntryAddendaCount = record[4:10]
	// 11-20 Total of all positions 4-11 on each Entry Detail Record in the batch. This is essentially the sum of all the RDFI routing numbers in the batch.
	// If the sum exceeds 10 digits (because you have lots of Entry Detail Records), lop off the most significant digits of the sum until there are only 10
	batchControl.EntryHash = record[10:20]
	// 21-32 Number of cents of debit entries within the batch
	batchControl.TotalDebitEntryDollarAmount = record[20:32]
	// 33-44 Number of cents of credit entries within the batch
	batchControl.TotalCreditEntryDollarAmount = record[32:44]
	// 45-54 This is the same as the "Company identification" field in previous Batch Header Record
	batchControl.CompanyIdentification = record[44:54]
	// 55-73 Seems to always be blank
	batchControl.MessageAuthenticationCode = record[54:73]
	// 74-79 Always blank (just fill with spaces)
	batchControl.Reserved = record[73:79]
	// 80-87 This is the same as the "ODFI identification" field in previous Batch Header Record
	batchControl.OdfiIdentification = record[79:87]
	// 88-94 This is the same as the "Batch number" field in previous Batch Header Record
	batchControl.BatchNumber = record[87:94]

	return batchControl
}

// parseFileControl takes the input record string and parses the FileControlRecord values
func parseFileControl(record string) (fileControl FileControlRecord) {
	// 1-1 Always "9"
	fileControl.RecordType = record[:1]
	// 2-7 The total number of Batch Header Record in the file. For example: "000003
	fileControl.BatchCount = record[1:7]
	// 8-13 e total number of blocks on the file, including the File Header and File Control records. One block is 10 lines, so it's effectively the number of lines in the file divided by 10.
	fileControl.BlockCount = record[7:13]
	// 14-21 Total number of Entry Detail Record in the file
	fileControl.EntryAddendaCount = record[13:21]
	// 22-31 Total of all positions 4-11 on each Entry Detail Record in the file. This is essentially the sum of all the RDFI routing numbers in the file.
	// If the sum exceeds 10 digits (because you have lots of Entry Detail Records), lop off the most significant digits of the sum until there are only 10
	fileControl.EntryHash = record[21:31]
	// 32-43 Number of cents of debit entries within the file
	fileControl.TotalDebitEntryDollarAmountInFile = record[31:43]
	// 44-55 Number of cents of credit entries within the file
	fileControl.TotalCreditEntryDollarAmountInFile = record[43:55]
	// 56-94 Reserved Always blank (just fill with spaces)
	fileControl.Reserved = record[55:94]

	return fileControl
}
