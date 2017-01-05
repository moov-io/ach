package ach

// FileControlRecord record contains entry counts, dollar totals and hash
// totals accumulated from each batch control record in the file.
type FileControlRecord struct {
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
