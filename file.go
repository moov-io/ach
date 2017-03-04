// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// Package ach reads and writes (ACH) Automated Clearing House files. ACH is the
// primary method of electronic money movemenet through the United States.
//
// https://en.wikipedia.org/wiki/Automated_Clearing_House
package ach

import "errors"

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
	ErrFileDebitAmount  = errors.New("Total Debit amountis out-of-balance with File Control")
	ErrFileCreditAmount = errors.New("Total Credit amountis out-of-balance with File Control")
	ErrFileEntryHash    = errors.New("Calculated Batch Control Entry hash does not match File Control Entry Hash")
)

// File contains the structures of a parsed ACH File.
type File struct {
	Header  FileHeader
	Batches []Batch
	Control FileControl
	// Converters is composed for ACH to golang Converters
	Converters
}

// addEntryDetail appends an EntryDetail to the Batch
func (f *File) addBatch(batch Batch) []Batch {
	f.Batches = append(f.Batches, batch)
	return f.Batches
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

// ValidateAll walks the enture file data structure and validates each record
func (f *File) ValidateAll() error {

	// validate inward out of the File Struct
	for _, batch := range f.Batches {
		for _, entry := range batch.Entries {
			for _, addenda := range entry.Addendums {
				if err := addenda.Validate(); err != nil {
					return err
				}
			}
			if err := entry.Validate(); err != nil {
				return err
			}
		}
		if err := batch.Header.Validate(); err != nil {
			return err
		}
		if err := batch.Control.Validate(); err != nil {
			return err
		}
		if err := batch.Validate(); err != nil {
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

// TODO: isEntryHashMismatch
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
// This field is prepared by hashing the 8-digit Routing Number in each batch.
// The Entry Hash provides a check against inadvertent alteration of data
func (f *File) isEntryHashMismatch() error {
	hash := 0
	for _, batch := range f.Batches {
		hash = hash + batch.Control.EntryHash
	}
	hashField := f.numericField(hash, 10)
	if hashField != f.Control.EntryHashField() {
		return ErrFileEntryHash
	}
	return nil
}
