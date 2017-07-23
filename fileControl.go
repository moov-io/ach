// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "fmt"

// FileControl record contains entry counts, dollar totals and hash
// totals accumulated from each batch control record in the file.
type FileControl struct {
	// RecordType defines the type of record in the block. fileControlPos 9
	recordType string

	// BatchCount total number of batches (i.e., ‘5’ records) in the file
	BatchCount int

	// BlockCount total number of records in the file (include all headers and trailer) divided
	// by 10 (This number must be evenly divisible by 10. If not, additional records consisting of all 9’s are added to the file after the initial ‘9’ record to fill out the block 10.)
	BlockCount int

	// EntryAddendaCount total detail and addenda records in the file
	EntryAddendaCount int

	// EntryHash calculated in the same manner as the batch has total but includes total from entire file
	EntryHash int

	// TotalDebitEntryDollarAmountInFile contains accumulated Batch debit totals within the file.
	TotalDebitEntryDollarAmountInFile int

	// TotalCreditEntryDollarAmountInFile contains accumulated Batch credit totals within the file.
	TotalCreditEntryDollarAmountInFile int
	// Reserved should be blank.
	reserved string
	// validator is composed for data validation
	validator
	// converters is composed for ACH to golang Converters
	converters
}

// Parse takes the input record string and parses the FileControl values
func (fc *FileControl) Parse(record string) {
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
	// 32-43 Number of cents of debit entries within the file
	fc.TotalDebitEntryDollarAmountInFile = fc.parseNumField(record[31:43])
	// 44-55 Number of cents of credit entries within the file
	fc.TotalCreditEntryDollarAmountInFile = fc.parseNumField(record[43:55])
	// 56-94 Reserved Always blank (just fill with spaces)
	fc.reserved = "                                       "
}

// NewFileControl returns a new FileControl with default values for none exported fields
func NewFileControl() FileControl {
	return FileControl{
		recordType: "9",
		reserved:   "                                       ",
	}
}

// String writes the FileControl struct to a 94 character string.
func (fc *FileControl) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v%v",
		fc.recordType,
		fc.BatchCountField(),
		fc.BlockCountField(),
		fc.EntryAddendaCountField(),
		fc.EntryHashField(),
		fc.TotalDebitEntryDollarAmountInFileField(),
		fc.TotalCreditEntryDollarAmountInFileField(),
		fc.reserved,
	)
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fc *FileControl) Validate() error {
	if err := fc.fieldInclusion(); err != nil {
		return err
	}
	if fc.recordType != "9" {
		msg := fmt.Sprintf(msgRecordType, 9)
		return &FieldError{FieldName: "recordType", Value: fc.recordType, Msg: msg}
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (fc *FileControl) fieldInclusion() error {
	if fc.recordType == "" {
		return &FieldError{FieldName: "recordType", Value: fc.recordType, Msg: msgFieldInclusion}
	}
	if fc.BatchCount == 0 {
		return &FieldError{FieldName: "BatchCount", Value: fc.BatchCountField(), Msg: msgFieldInclusion}
	}
	if fc.BlockCount == 0 {
		return &FieldError{FieldName: "BlockCount", Value: fc.BlockCountField(), Msg: msgFieldInclusion}
	}
	if fc.EntryAddendaCount == 0 {
		return &FieldError{FieldName: "EntryAddendaCount", Value: fc.EntryAddendaCountField(), Msg: msgFieldInclusion}
	}
	if fc.EntryHash == 0 {
		return &FieldError{FieldName: "EntryAddendaCount", Value: fc.EntryAddendaCountField(), Msg: msgFieldInclusion}
	}
	return nil
}

// BatchCountField gets a string of the batch count zero padded
func (fc *FileControl) BatchCountField() string {
	return fc.numericField(fc.BatchCount, 6)
}

// BlockCountField gets a string of the block count zero padded
func (fc *FileControl) BlockCountField() string {
	return fc.numericField(fc.BlockCount, 6)
}

// EntryAddendaCountField gets a string of entry addenda batch count zero padded
func (fc *FileControl) EntryAddendaCountField() string {
	return fc.numericField(fc.EntryAddendaCount, 8)
}

// EntryHashField gets a string of entry hash zero padded
func (fc *FileControl) EntryHashField() string {
	return fc.numericField(fc.EntryHash, 10)
}

// TotalDebitEntryDollarAmountInFileField get a zero padded Total debit Entry Amount
func (fc *FileControl) TotalDebitEntryDollarAmountInFileField() string {
	return fc.numericField(fc.TotalDebitEntryDollarAmountInFile, 12)
}

// TotalCreditEntryDollarAmountInFileField get a zero padded Total credit Entry Amount
func (fc *FileControl) TotalCreditEntryDollarAmountInFileField() string {
	return fc.numericField(fc.TotalCreditEntryDollarAmountInFile, 12)
}
