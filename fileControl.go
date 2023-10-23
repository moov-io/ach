// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package ach

import (
	"strings"
	"unicode/utf8"
)

// FileControl record contains entry counts, dollar totals and hash
// totals accumulated from each batch control record in the file.
type FileControl struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// BatchCount total number of batches (i.e., '5' records) in the file
	BatchCount int `json:"batchCount"`
	// BlockCount total number of records in the file (include all headers and trailer) divided
	// by 10 (This number must be evenly divisible by 10. If not, additional records consisting of all 9's are added to the file after the initial '9' record to fill out the block 10.)
	BlockCount int `json:"blockCount"`
	// EntryAddendaCount is a tally of each Entry Detail Record and each Addenda
	// Record processed, within either the batch or file as appropriate.
	EntryAddendaCount int `json:"entryAddendaCount"`
	// EntryHash calculated in the same manner as the batch has total but includes total from entire file
	EntryHash int `json:"entryHash"`
	// TotalDebitEntryDollarAmountInFile contains accumulated Batch debit totals within the file.
	TotalDebitEntryDollarAmountInFile int `json:"totalDebit"`
	// TotalCreditEntryDollarAmountInFile contains accumulated Batch credit totals within the file.
	TotalCreditEntryDollarAmountInFile int `json:"totalCredit"`
	// validator is composed for data validation
	validator
	// converters is composed for ACH to golang Converters
	converters
}

// Parse takes the input record string and parses the FileControl values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (fc *FileControl) Parse(record string) {
	runeCount := utf8.RuneCountInString(record)
	if runeCount != 94 {
		return
	}

	buf := getBuffer()
	defer saveBuffer(buf)

	reset := func() string {
		out := buf.String()
		buf.Reset()
		return out
	}

	var idx int
	for _, r := range record {
		idx++

		// Append rune to buffer
		buf.WriteRune(r)

		// At each field cutoff save the buffer and reset
		switch idx {
		case 1:
			// 1-1 Always "9"
			reset()
		case 7:
			// 2-7 The total number of Batch Header Record in the file. For example: "000003
			fc.BatchCount = fc.parseNumField(reset())
		case 13:
			// 8-13 e total number of blocks on the file, including the File Header and File Control records. One block is 10 lines, so it's effectively the number of lines in the file divided by 10.
			fc.BlockCount = fc.parseNumField(reset())
		case 21:
			// 14-21 Total number of Entry Detail Record in the file
			fc.EntryAddendaCount = fc.parseNumField(reset())
		case 31:
			// 22-31 Total of all positions 4-11 on each Entry Detail Record in the file. This is essentially the sum of all the RDFI routing numbers in the file.
			// If the sum exceeds 10 digits (because you have lots of Entry Detail Records), lop off the most significant digits of the sum until there are only 10
			fc.EntryHash = fc.parseNumField(reset())
		case 43:
			// 32-43 Number of cents of debit entries within the file
			fc.TotalDebitEntryDollarAmountInFile = fc.parseNumField(reset())
		case 55:
			// 44-55 Number of cents of credit entries within the file
			fc.TotalCreditEntryDollarAmountInFile = fc.parseNumField(reset())
		case 94:
			// 56-94 Reserved Always blank (just fill with spaces)
			reset()
		}
	}
}

// NewFileControl returns a new FileControl with default values for none exported fields
func NewFileControl() FileControl {
	return FileControl{}
}

// String writes the FileControl struct to a 94 character string.
func (fc *FileControl) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(fileControlPos)
	buf.WriteString(fc.BatchCountField())
	buf.WriteString(fc.BlockCountField())
	buf.WriteString(fc.EntryAddendaCountField())
	buf.WriteString(fc.EntryHashField())
	buf.WriteString(fc.TotalDebitEntryDollarAmountInFileField())
	buf.WriteString(fc.TotalCreditEntryDollarAmountInFileField())
	buf.WriteString("                                       ")
	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fc *FileControl) Validate() error {
	if err := fc.fieldInclusion(); err != nil {
		return err
	}
	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (fc *FileControl) fieldInclusion() error {
	if fc.BlockCount == 0 {
		return fieldError("BlockCount", ErrConstructor, fc.BlockCountField())
	}
	if fc.TotalCreditEntryDollarAmountInFile != 0 || fc.TotalDebitEntryDollarAmountInFile != 0 {
		if fc.BatchCount == 0 {
			return fieldError("BatchCount", ErrConstructor, fc.BatchCountField())
		}
		if fc.EntryAddendaCount == 0 {
			return fieldError("EntryAddendaCount", ErrConstructor, fc.EntryAddendaCountField())
		}
		if fc.EntryHash == 0 {
			return fieldError("EntryHash", ErrConstructor, fc.EntryAddendaCountField())
		}
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
