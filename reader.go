// Copyright 2016 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

const (
	// RecordLength character count of each line representing a letter in a file
	RecordLength = 94
)

// ParseError is returned for parsing reader errors.
// The first line is 1.
type ParseError struct {
	Line  int    // Line number where the error accurd
	Field string // Name of the field being parsed
	Err   error  // The actual error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("line %d, field  %s: %s", e.Line, e.Field, e.Err)
}

// These are the errors that can be returned in Parse.Error
var (
	ErrRecordLen    = errors.New("Wrong number of fields in record expect 94")
	ErrBatchControl = errors.New("No terminating batch control record found in file")
	ErrRecordType   = errors.New("Unhandled Record Type")
)

// Reader reads records from a ACH-encoded file.
type Reader struct {
	// r handles the IO.Reader sent to be parser.
	r *bufio.Reader
	// line number of the file being parsed
	line int
	// record holds the current line being parser.
	record string
	// field number of the record currently being parsed
	field string
}

// error creates a new ParseError based on err.
func (r *Reader) error(err error) error {
	return &ParseError{
		Line:  r.line,
		Field: r.field,
		Err:   err,
	}
}

// NewReader returns a new Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	return &Reader{
		r: bufio.NewReader(r),
	}
}

// Read reads each line of the ACH file and defines which parser to use based
// on the first byte of each line. It also enforces ACH formating rules and returns
// the appropriate error if issues are founr.
func (r *Reader) Read() (ach ACH, err error) {
	// read through the entire file
	for {
		line, _, err := r.r.ReadLine()
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
			return ach, r.error(ErrRecordLen)
		}
		// TODO: Check that all characters are accepter.

		r.line++

		r.record = string(line)
		switch r.record[:1] {
		case headerPos:
			ach.FileHeader = r.parseFileHeader()
		case batchPos:
			ach.BatchHeader = r.parseBatchHeader()
		case entryDetailPos:
			ach.EntryDetail = r.parseEntryDetail()
		case entryAddendaPos:
			// TODO: implement parseEntryAgenda
			//ach.EntryAgenda = parseEntryAgenda(record)
		case batchControlPos:
			ach.BatchControl = r.parseBatchControl()
		case fileControlPos:
			ach.FileControl = r.parseFileControl()
		default:
			//fmt.Println("Record type not detected")
			// TODO: return nil, error
			return ach, r.error(ErrRecordType)
		}

	}
	// TODO: number of lines in file must be divisable by 10 the blocking factor
	//fmt.Printf("Number of lines in file: %v \n", r.line)
	return ach, nil
}

// parseFileHeader takes the input record string and parses the FileHeaderRecord values
func (r *Reader) parseFileHeader() (fh FileHeader) {
	fh.parse(r.record)
	return fh
}

// parseBatchHeader takes the input record string and parses the FileHeaderRecord values
func (r *Reader) parseBatchHeader() (batchHeader BatchHeaderRecord) {
	// 1-1 Always "5"
	batchHeader.RecordType = r.record[:1]
	// 2-4 If the entries are credits, always "220". If the entries are debits, always "225"
	batchHeader.ServiceClassCode = r.record[1:4]
	// 5-20 Your company's name. This name may appear on the receivers’ statements prepared by the RDFI.
	batchHeader.CompanyName = r.record[4:20]
	// 21-40 Optional field you may use to describe the batch for internal accounting purposes
	batchHeader.CompanyDiscretionaryData = r.record[20:40]
	// 41-50 A 10-digit number assigned to you by the ODFI once they approve you to
	// originate ACH files through them. This is the same as the "Immediate origin" field in File Header Record
	batchHeader.CompanyIdentification = r.record[40:50]
	// 51-53 If the entries are PPD (credits/debits towards consumer account), use "PPD".
	// If the entries are CCD (credits/debits towards corporate account), use "CCD".
	// The difference between the 2 class codes are outside of the scope of this post, but generally most ACH transfers to consumer bank accounts should use "PPD"
	batchHeader.StandardEntryClassCode = r.record[50:53]
	// 54-63 Your description of the transaction. This text will appear on the receivers’ bank statement.
	// For example: "Payroll   "
	batchHeader.CompanyEntryDescription = r.record[53:63]
	// 64-69 The date you choose to identify the transactions in YYMMDD format.
	// This date may be printed on the receivers’ bank statement by the RDFI
	batchHeader.CompanyDescriptiveDate = r.record[63:69]
	// 70-75 Date transactions are to be posted to the receivers’ account.
	// You almost always want the transaction to post as soon as possible, so put tomorrow's date in YYMMDD format
	batchHeader.EffectiveEntryDate = r.record[69:75]
	// 76-79 Always blank (just fill with spaces)
	batchHeader.SettlementDate = r.record[75:78]
	// 79-79 Always 1
	batchHeader.OriginatorStatusCode = r.record[78:79]
	// 80-87 Your ODFI's routing number without the last digit. The last digit is simply a
	// checksum digit, which is why it is not necessary
	batchHeader.OdfiIdentification = r.record[79:87]
	// 88-94 Sequential number of this Batch Header Recorr.
	// For example, put "1" if this is the first Batch Header Record in the file
	batchHeader.BatchNumber = r.record[87:94]

	return batchHeader
}

// parseEntryDetail takes the input record string and parses the EntryDetailRecord values
func (r *Reader) parseEntryDetail() (entryDetail EntryDetailRecord) {
	// 1-1 Always "6"
	entryDetail.RecordType = r.record[:1]
	// 2-3 is checking credit 22 debit 27 savings credit 32 debit 37
	entryDetail.TransactionCode = r.record[1:3]
	// 4-11 the RDFI's routing number without the last digit.
	entryDetail.RdfiIdentification = r.record[3:11]
	// 12-12 The last digit of the RDFI's routing number
	entryDetail.CheckDigit = r.record[11:12]
	// 13-29 The receiver's bank account number you are crediting/debiting
	entryDetail.DfiAccountNumber = r.record[12:29]
	// 30-39 Number of cents you are debiting/crediting this account
	entryDetail.Amount = r.record[29:39]
	// 40-54 An internal identification (alphanumeric) that you use to uniquely identify this Entry Detail Record
	entryDetail.IndividualIdentificationNumber = r.record[39:54]
	// 55-76 The name of the receiver, usually the name on the bank account
	entryDetail.IndividualName = r.record[54:76]
	// 77-78 allows ODFIs to include codes of significance only to them
	// normally blank
	entryDetail.DiscretionaryData = r.record[76:78]
	// 79-79 1 if addenda exists 0 if it does not
	entryDetail.AddendaRecordIndicator = r.record[78:79]
	// 80-84 An internal identification (alphanumeric) that you use to uniquely identify
	// this Entry Detail Recorr. This number should be unique to the transaction and will help identify the transaction in case of an inquiry
	entryDetail.TraceNumber = r.record[79:94]

	return entryDetail
}

// parseAddendaRecord takes the input record string and parses the AddendaRecord values
func (r *Reader) parseAddendaRecord() (addenda AddendaRecord) {
	// 1-1 Always "7"
	addenda.RecordType = r.record[:1]
	// 2-3 Defines the specific explanation and format for the addenda information contained in the same record
	addenda.TypeCode = r.record[1:3]
	// 4-83 Based on the information enterer. (04-83) 80 alphanumeric
	addenda.PaymentRelatedInformation = r.record[3:83]
	// 84-87 SequenceNumber is consecutively assigned to each Addenda Record following
	// an Entry Detail Record
	addenda.SequenceNumber = r.record[83:87]
	// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
	addenda.EntryDetailSequenceNumber = r.record[87:94]

	return addenda
}

// parseBatchControl takes the input record string and parses the BatchControlRecord values
func (r *Reader) parseBatchControl() (batchControl BatchControlRecord) {
	// 1-1 Always "8"
	batchControl.RecordType = r.record[:1]
	// 2-4 This is the same as the "Service code" field in previous Batch Header Record
	batchControl.ServiceClassCode = r.record[1:4]
	// 5-10 Total number of Entry Detail Record in the batch
	batchControl.EntryAddendaCount = r.record[4:10]
	// 11-20 Total of all positions 4-11 on each Entry Detail Record in the batch. This is essentially the sum of all the RDFI routing numbers in the batch.
	// If the sum exceeds 10 digits (because you have lots of Entry Detail Records), lop off the most significant digits of the sum until there are only 10
	batchControl.EntryHash = r.record[10:20]
	// 21-32 Number of cents of debit entries within the batch
	batchControl.TotalDebitEntryDollarAmount = r.record[20:32]
	// 33-44 Number of cents of credit entries within the batch
	batchControl.TotalCreditEntryDollarAmount = r.record[32:44]
	// 45-54 This is the same as the "Company identification" field in previous Batch Header Record
	batchControl.CompanyIdentification = r.record[44:54]
	// 55-73 Seems to always be blank
	batchControl.MessageAuthenticationCode = r.record[54:73]
	// 74-79 Always blank (just fill with spaces)
	batchControl.Reserved = r.record[73:79]
	// 80-87 This is the same as the "ODFI identification" field in previous Batch Header Record
	batchControl.OdfiIdentification = r.record[79:87]
	// 88-94 This is the same as the "Batch number" field in previous Batch Header Record
	batchControl.BatchNumber = r.record[87:94]

	return batchControl
}

// parseFileControl takes the input record string and parses the FileControlRecord values
func (r *Reader) parseFileControl() (fileControl FileControlRecord) {
	// 1-1 Always "9"
	fileControl.RecordType = r.record[:1]
	// 2-7 The total number of Batch Header Record in the file. For example: "000003
	fileControl.BatchCount = r.record[1:7]
	// 8-13 e total number of blocks on the file, including the File Header and File Control records. One block is 10 lines, so it's effectively the number of lines in the file divided by 10.
	fileControl.BlockCount = r.record[7:13]
	// 14-21 Total number of Entry Detail Record in the file
	fileControl.EntryAddendaCount = r.record[13:21]
	// 22-31 Total of all positions 4-11 on each Entry Detail Record in the file. This is essentially the sum of all the RDFI routing numbers in the file.
	// If the sum exceeds 10 digits (because you have lots of Entry Detail Records), lop off the most significant digits of the sum until there are only 10
	fileControl.EntryHash = r.record[21:31]
	// 32-43 Number of cents of debit entries within the file
	fileControl.TotalDebitEntryDollarAmountInFile = r.record[31:43]
	// 44-55 Number of cents of credit entries within the file
	fileControl.TotalCreditEntryDollarAmountInFile = r.record[43:55]
	// 56-94 Reserved Always blank (just fill with spaces)
	fileControl.Reserved = r.record[55:94]

	return fileControl
}
