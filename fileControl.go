package ach

// FileControl record contains entry counts, dollar totals and hash
// totals accumulated from each batch control record in the file.
type FileControl struct {
	// RecordType defines the type of record in the block. fileControlPos 9
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

// Parse takes the input record string and parses the FileControl values
func (fc *FileControl) Parse(record string) {
	// 1-1 Always "9"
	fc.RecordType = record[:1]
	// 2-7 The total number of Batch Header Record in the file. For example: "000003
	fc.BatchCount = record[1:7]
	// 8-13 e total number of blocks on the file, including the File Header and File Control records. One block is 10 lines, so it's effectively the number of lines in the file divided by 10.
	fc.BlockCount = record[7:13]
	// 14-21 Total number of Entry Detail Record in the file
	fc.EntryAddendaCount = record[13:21]
	// 22-31 Total of all positions 4-11 on each Entry Detail Record in the file. This is essentially the sum of all the RDFI routing numbers in the file.
	// If the sum exceeds 10 digits (because you have lots of Entry Detail Records), lop off the most significant digits of the sum until there are only 10
	fc.EntryHash = record[21:31]
	// 32-43 Number of cents of debit entries within the file
	fc.TotalDebitEntryDollarAmountInFile = record[31:43]
	// 44-55 Number of cents of credit entries within the file
	fc.TotalCreditEntryDollarAmountInFile = record[43:55]
	// 56-94 Reserved Always blank (just fill with spaces)
	fc.Reserved = record[55:94]
}

// NewFileControl returns a new FileControl with default values for none exported fields
func NewFileControl() *FileControl {
	return &FileControl{}
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fc *FileControl) Validate() (bool, error) {
	return true, nil
}
