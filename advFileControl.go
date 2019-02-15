// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"unicode/utf8"
)

// ADVFileControl record contains entry counts, dollar totals and hash
// totals accumulated from each batchADV control record in the file.
type ADVFileControl struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record in the block. fileControlPos 9
	recordType string

	// BatchCount total number of batches (i.e., ‘5’ records) in the file
	BatchCount int `json:"batchCount"`

	// BlockCount total number of records in the file (include all headers and trailer) divided
	// by 10 (This number must be evenly divisible by 10. If not, additional records consisting of all 9’s are added to the file after the initial ‘9’ record to fill out the block 10.)
	BlockCount int `json:"blockCount,omitempty"`

	// EntryAddendaCount total detail and addenda records in the file
	EntryAddendaCount int `json:"entryAddendaCount"`

	// EntryHash calculated in the same manner as the batch has total but includes total from entire file
	EntryHash int `json:"entryHash"`

	// TotalDebitEntryDollarAmountInFile contains accumulated Batch debit totals within the file.
	TotalDebitEntryDollarAmountInFile int `json:"totalDebit"`

	// TotalCreditEntryDollarAmountInFile contains accumulated Batch credit totals within the file.
	TotalCreditEntryDollarAmountInFile int `json:"totalCredit"`
	// Reserved should be blank.
	reserved string
	// validator is composed for data validation
	validator
	// converters is composed for ACH to golang Converters
	converters
}

// Parse takes the input record string and parses the FileControl values
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fc *ADVFileControl) Parse(record string) {
	if utf8.RuneCountInString(record) < 71 {
		return
	}

	// 1-1 Always "9"
	fc.recordType = "9"
	// 2-7 The total number of Batch Header Record in the file. For example: "000003
	fc.BatchCount = fc.parseNumField(record[1:7])
	// 8-13 e total number of blocks on the file, including the File Header and File Control records. One block is 10 lines, so it's effectively the number of lines in the file divided by 10.
	fc.BlockCount = fc.parseNumField(record[7:13])
	// 14-21 Total number of Entry Detail Record in the file
	fc.EntryAddendaCount = fc.parseNumField(record[13:21])
	// 22-31 Total of all positions 4-11 on each Entry Detail Record in the file. This is essentially the sum of all the RDFI routing numbers in the file.
	// If the sum exceeds 10 digits (because you have lots of Entry Detail Records), lop off the most significant digits of the sum until there are only 10
	fc.EntryHash = fc.parseNumField(record[21:31])
	// 32-51 Number of cents of debit entries within the file
	fc.TotalDebitEntryDollarAmountInFile = fc.parseNumField(record[31:51])
	// 52-71 Number of cents of credit entries within the file
	fc.TotalCreditEntryDollarAmountInFile = fc.parseNumField(record[51:71])
	// 72-94 Reserved Always blank (just fill with spaces)
	fc.reserved = "                       "
}

// NewADVFileControl returns a new ADVFileControl with default values for none exported fields
func NewADVFileControl() ADVFileControl {
	return ADVFileControl{
		recordType: "9",
		reserved:   "                       ",
	}
}

// String writes the ADVFileControl struct to a 94 character string.
func (fc *ADVFileControl) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(fc.recordType)
	buf.WriteString(fc.BatchCountField())
	buf.WriteString(fc.BlockCountField())
	buf.WriteString(fc.EntryAddendaCountField())
	buf.WriteString(fc.EntryHashField())
	buf.WriteString(fc.TotalDebitEntryDollarAmountInFileField())
	buf.WriteString(fc.TotalCreditEntryDollarAmountInFileField())
	buf.WriteString(fc.reserved)
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fc *ADVFileControl) Validate() error {
	if err := fc.fieldInclusion(); err != nil {
		return err
	}
	if fc.recordType != "9" {
		return fieldError("recordType", NewErrRecordType(9), fc.recordType)
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (fc *ADVFileControl) fieldInclusion() error {
	if fc.recordType == "" {
		return fieldError("recordType", ErrConstructor, fc.recordType)
	}
	if fc.BatchCount == 0 {
		return fieldError("BatchCount", ErrConstructor, fc.BatchCountField())
	}
	if fc.BlockCount == 0 {
		return fieldError("BlockCount", ErrConstructor, fc.BlockCountField())
	}
	if fc.EntryAddendaCount == 0 {
		return fieldError("EntryAddendaCount", ErrConstructor, fc.EntryAddendaCountField())
	}
	if fc.EntryHash == 0 {
		return fieldError("EntryHash", ErrConstructor, fc.EntryAddendaCountField())
	}
	return nil
}

// BatchCountField gets a string of the batch count zero padded
func (fc *ADVFileControl) BatchCountField() string {
	return fc.numericField(fc.BatchCount, 6)
}

// BlockCountField gets a string of the block count zero padded
func (fc *ADVFileControl) BlockCountField() string {
	return fc.numericField(fc.BlockCount, 6)
}

// EntryAddendaCountField gets a string of entry addenda batch count zero padded
func (fc *ADVFileControl) EntryAddendaCountField() string {
	return fc.numericField(fc.EntryAddendaCount, 8)
}

// EntryHashField gets a string of entry hash zero padded
func (fc *ADVFileControl) EntryHashField() string {
	return fc.numericField(fc.EntryHash, 10)
}

// TotalDebitEntryDollarAmountInFileField get a zero padded Total debit Entry Amount
func (fc *ADVFileControl) TotalDebitEntryDollarAmountInFileField() string {
	return fc.numericField(fc.TotalDebitEntryDollarAmountInFile, 20)
}

// TotalCreditEntryDollarAmountInFileField get a zero padded Total credit Entry Amount
func (fc *ADVFileControl) TotalCreditEntryDollarAmountInFileField() string {
	return fc.numericField(fc.TotalCreditEntryDollarAmountInFile, 20)
}
