package ach

import (
	"fmt"
	"strconv"
)

// FileControl record contains entry counts, dollar totals and hash
// totals accumulated from each batch control record in the file.
type FileControl struct {
	// RecordType defines the type of record in the block. fileControlPos 9
	recordType string

	// BatchCount total number of batches (i.e., ‘5’ records) in the file
	batchCount int

	// BlockCount total number of records in the file (include all headers and trailer) divided
	// by 10 (This number must be evenly divisible by 10. If not, additional records consisting of all 9’s are added to the file after the initial ‘9’ record to fill out the block 10.)
	blockCount int

	// EntryAddendaCount total detail and addenda records in the file
	entryAddendaCount int

	// EntryHash calculated in the same manner as the batch has total but includes total from entire file
	entryHash int

	// TotalDebitEntryDollarAmountInFile contains accumulated Batch debit totals within the file.
	totalDebitEntryDollarAmountInFile int

	// TotalCreditEntryDollarAmountInFile contains accumulated Batch credit totals within the file.
	totalCreditEntryDollarAmountInFile int
	// Reserved should be blank.
	reserved string
	// Validator is composed for data validation
	Validator
	// Converters is composed for ACH to golang Converters
	Converters
}

// Parse takes the input record string and parses the FileControl values
func (fc *FileControl) Parse(record string) {
	// 1-1 Always "9"
	fc.recordType = record[:1]
	// 2-7 The total number of Batch Header Record in the file. For example: "000003
	fc.batchCount = fc.parseNumField(record[1:7])
	// 8-13 e total number of blocks on the file, including the File Header and File Control records. One block is 10 lines, so it's effectively the number of lines in the file divided by 10.
	fc.blockCount = fc.parseNumField(record[7:13])
	// 14-21 Total number of Entry Detail Record in the file
	fc.entryAddendaCount = fc.parseNumField(record[13:21])
	// 22-31 Total of all positions 4-11 on each Entry Detail Record in the file. This is essentially the sum of all the RDFI routing numbers in the file.
	// If the sum exceeds 10 digits (because you have lots of Entry Detail Records), lop off the most significant digits of the sum until there are only 10
	fc.entryHash = fc.parseNumField(record[21:31])
	// 32-43 Number of cents of debit entries within the file
	fc.totalDebitEntryDollarAmountInFile = fc.parseNumField(record[31:43])
	// 44-55 Number of cents of credit entries within the file
	fc.totalCreditEntryDollarAmountInFile = fc.parseNumField(record[43:55])
	// 56-94 Reserved Always blank (just fill with spaces)
	fc.reserved = record[55:94]
}

// NewFileControl returns a new FileControl with default values for none exported fields
func NewFileControl() *FileControl {
	return &FileControl{
		recordType: "9",
		reserved:   "                                       ",
	}
}

// String writes the FileControl struct to a 94 character string.
func (fc *FileControl) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v%v",
		fc.recordType,
		fc.BatchCount(),
		fc.BlockCount(),
		fc.EntryAddendaCount(),
		fc.EntryHash(),
		fc.TotalDebitEntryDollarAmountInFile(),
		fc.TotalCreditEntryDollarAmountInFile(),
		fc.reserved,
	)
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fc *FileControl) Validate() (bool, error) {
	if fc.recordType != "9" {
		return false, ErrRecordType
	}
	return true, nil
}

// BatchCount gets a string of the batch count zero padded
func (fc *FileControl) BatchCount() string {
	return fc.leftPad(strconv.Itoa(fc.batchCount), "0", 6)
}

// BlockCount gets a string of the block count zero padded
func (fc *FileControl) BlockCount() string {
	return fc.leftPad(strconv.Itoa(fc.blockCount), "0", 6)
}

// EntryAddendaCount gets a string of entry addenda batch count zero padded
func (fc *FileControl) EntryAddendaCount() string {
	return fc.leftPad(strconv.Itoa(fc.entryAddendaCount), "0", 8)
}

// EntryHash gets a string of entry hash zero padded
func (fc *FileControl) EntryHash() string {
	return fc.leftPad(strconv.Itoa(fc.entryHash), "0", 10)
}

// TotalDebitEntryDollarAmountInFile get a zero padded Total debit Entry Amount
func (fc *FileControl) TotalDebitEntryDollarAmountInFile() string {
	return fc.leftPad(strconv.Itoa(fc.totalDebitEntryDollarAmountInFile), "0", 12)
}

// TotalCreditEntryDollarAmountInFile get a zero padded Total credit Entry Amount
func (fc *FileControl) TotalCreditEntryDollarAmountInFile() string {
	return fc.leftPad(strconv.Itoa(fc.totalCreditEntryDollarAmountInFile), "0", 12)
}
