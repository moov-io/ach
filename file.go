// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// Package ach reads and writes (ACH) Automated Clearing House files. ACH is the
// primary method of electronic money movemenet through the United States.
//
// https://en.wikipedia.org/wiki/Automated_Clearing_House
package ach

import (
	"errors"
)

// First position of all Record Types. These codes are uniquily assigned to
// the first byte of each row in a file.
const (
	headerPos       = "1"
	batchPos        = "5"
	entryDetailPos  = "6"
	entryAddendaPos = "7"
	batchControlPos = "8"
	fileControlPos  = "9"
)

// Errors specific to parsing a Batch container
var (
	ErrFileBatchCount   = errors.New("Total Number of Batches in file is out-of-balance with File Control")
	ErrFileEntryCount   = errors.New("Total Entries and Addenda count is out-of-balance with File Control")
	ErrFileDebitAmount  = errors.New("Total Debit amount is out-of-balance with File Control")
	ErrFileCreditAmount = errors.New("Total Credit amount is out-of-balance with File Control")
	ErrFileEntryHash    = errors.New("Calculated Batch Control Entry hash does not match File Control Entry Hash")
	ErrFileBatches      = errors.New("File must have []*Batches to be built")
)

// File contains the structures of a parsed ACH File.
type File struct {
	Header  FileHeader
	Batches []*Batch
	Control FileControl
	// converters is composed for ACH to golang Converters
	converters
}

// NewFile constucuts a file template.
func NewFile() *File {
	return &File{
		Header: NewFileHeader(),
		// Batches: []Batch, TODO need a NewBatch
		Control: NewFileControl(),
	}
}

// Build creates a valid file and requires that the FileHeader and at least one Batch
func (f *File) Build() error {
	// Requires a valid FileHeader to build FileControl
	if err := f.Header.Validate(); err != nil {
		return err
	}
	// Requires at least one Batch in the new file.
	if len(f.Batches) <= 0 {
		return ErrFileBatches
	}
	// add 2 for FileHeader/control and reset if build was called twice do to error
	totalRecordsInFile := 2
	batchSeq := 1
	fileEntryAddendaCount := 0
	fileEntryHashSum := 0
	totalDebitAmount := 0
	totalCreditAmount := 0
	for i, batch := range f.Batches {
		// create ascending batch numbers
		f.Batches[i].Header.BatchNumber = batchSeq
		f.Batches[i].Control.BatchNumber = batchSeq
		batchSeq++
		// sum file entry and addenda records. Assume batch.Build() batch properly calculated control
		fileEntryAddendaCount = fileEntryAddendaCount + batch.Control.EntryAddendaCount
		// add 2 for Batch header/control + entry added count
		totalRecordsInFile = totalRecordsInFile + 2 + batch.Control.EntryAddendaCount
		// sum hash from batch control. Assume Batch.Build properly calculated field.
		fileEntryHashSum = fileEntryHashSum + batch.Control.EntryHash
		totalDebitAmount = totalDebitAmount + batch.Control.TotalDebitEntryDollarAmount
		totalCreditAmount = totalCreditAmount + batch.Control.TotalCreditEntryDollarAmount

	}
	// create FileControl from calculated values
	fc := NewFileControl()
	fc.BatchCount = batchSeq - 1
	// blocking factor of 10 is static default value in f.Header.blockingFactor.
	if (totalRecordsInFile % 10) != 0 {
		fc.BlockCount = totalRecordsInFile/10 + 1
	} else {
		fc.BlockCount = totalRecordsInFile / 10
	}
	fc.EntryAddendaCount = fileEntryAddendaCount
	fc.EntryHash = fileEntryHashSum
	fc.TotalDebitEntryDollarAmountInFile = totalDebitAmount
	fc.TotalCreditEntryDollarAmountInFile = totalCreditAmount
	f.Control = fc
	return nil
}

// AddBatch appends a Batch to the ach.File
func (f *File) AddBatch(batch *Batch) []*Batch {
	f.Batches = append(f.Batches, batch)
	return f.Batches
}

// SetHeader allows for header to be built.
func (f *File) SetHeader(h FileHeader) *File {
	f.Header = h
	return f
}

// Validate NACHA rules on the entire batch before being added to a File
func (f *File) Validate() error {
	// The value of the Batch Count Field is equal to the number of Company/Batch/Header Records in the file.
	if f.Control.BatchCount != len(f.Batches) {
		return ErrFileBatchCount
	}

	if err := f.isEntryAddendaCount(); err != nil {
		return err
	}

	if err := f.isFileAmount(); err != nil {
		return err
	}

	if err := f.isEntryHashMismatch(); err != nil {
		return err
	}

	return nil
}

// ValidateAll walks the enture data structure and validates each record
func (f *File) ValidateAll() error {

	// validate inward out of the File Struct
	for _, batch := range f.Batches {
		if err := batch.ValidateAll(); err != nil {
			return err
		}
	}
	if err := f.Header.Validate(); err != nil {
		return err
	}
	if err := f.Control.Validate(); err != nil {
		return err
	}
	if err := f.Validate(); err != nil {
		return err
	}
	return nil
}

// This field is prepared by hashing the RDFI’s 8-digit Routing Number in each entry.
//The Entry Hash provides a check against inadvertent alteration of data
func (f *File) isEntryAddendaCount() error {
	count := 0
	// we assume that each batch block has already validated the addenda count in batch control.
	for _, batch := range f.Batches {
		count += batch.Control.EntryAddendaCount
	}
	if f.Control.EntryAddendaCount != count {
		return ErrFileEntryCount
	}
	return nil
}

// isFileAmount tThe Total Debit and Credit Entry Dollar Amounts Fields contain accumulated
// Entry Detail debit and credit totals within the file
func (f *File) isFileAmount() error {
	debit := 0
	credit := 0
	for _, batch := range f.Batches {
		debit += batch.Control.TotalDebitEntryDollarAmount
		credit += batch.Control.TotalCreditEntryDollarAmount
	}
	if f.Control.TotalDebitEntryDollarAmountInFile != debit {
		return ErrFileDebitAmount
	}
	if f.Control.TotalCreditEntryDollarAmountInFile != credit {
		return ErrFileCreditAmount
	}
	return nil
}

// isEntryHashMismatch validates the hash by recalulating the result
func (f *File) isEntryHashMismatch() error {
	hashField := f.calculateEntryHash()
	if hashField != f.Control.EntryHashField() {
		return ErrFileEntryHash
	}
	return nil
}

// calculateEntryHash This field is prepared by hashing the 8-digit Routing Number in each batch.
// The Entry Hash provides a check against inadvertent alteration of data
func (f *File) calculateEntryHash() string {
	hash := 0
	for _, batch := range f.Batches {
		hash = hash + batch.Control.EntryHash
	}
	return f.numericField(hash, 10)
}
