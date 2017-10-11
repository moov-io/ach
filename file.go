// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// Package ach reads and writes (ACH) Automated Clearing House files. ACH is the
// primary method of electronic money movement through the United States.
//
// https://en.wikipedia.org/wiki/Automated_Clearing_House
package ach

import (
	"fmt"
	"strconv"
)

// First position of all Record Types. These codes are uniquely assigned to
// the first byte of each row in a file.
const (
	fileHeaderPos   = "1"
	batchHeaderPos  = "5"
	entryDetailPos  = "6"
	entryAddendaPos = "7"
	batchControlPos = "8"
	fileControlPos  = "9"

	// RecordLength character count of each line representing a letter in a file
	RecordLength = 94
)

// currently supported SEC codes
const (
	ppd = "PPD"
	web = "WEB"
	ccd = "CCD"
	cor = "COR"
)

// Errors strings specific to parsing a Batch container
var (
	msgFileControlEquality           = "header %v is not equal to control %v"
	msgFileCalculatedControlEquality = "calculated %v is out-of-balance with control %v"
	// specific messages
	msgRecordLength      = "must be 94 characters and found %d"
	msgFileBatchOutside  = "outside of current batch"
	msgFileBatchInside   = "inside of current batch"
	msgFileControl       = "none or more than one file control exists"
	msgFileHeader        = "none or more than one file headers exists"
	msgUnknownRecordType = "%s is an unknown record type"
	msgFileNoneSEC       = "%v SEC(standard entry class) is not implemented"
)

// FileError is an error describing issues validating a file
type FileError struct {
	FieldName string
	Value     string
	Msg       string
}

func (e *FileError) Error() string {
	return fmt.Sprintf("%s %s", e.FieldName, e.Msg)
}

// File contains the structures of a parsed ACH File.
type File struct {
	Header  FileHeader
	Batches []Batcher
	Control FileControl

	converters
}

// FileParam is the minimal fields required to make a ach file header
type FileParam struct {
	// ImmediateDestination is the originating banks ABA routing number. Frequently your banks ABA routing number.
	ImmediateDestination string `json:"immediate_destination"`
	// ImmediateOrigin
	ImmediateOrigin string `json:"immediate_origin"`
	// ImmediateDestinationName is the originating banks name.
	ImmediateDestinationName string `json:"immediate_destination_name"`
	ImmediateOriginName      string `json:"immediate_origin_name"`
	ReferenceCode            string `json:"reference_code,omitempty"`
}

// NewFile constructs a file template.
func NewFile(params ...FileParam) *File {
	if len(params) > 0 {
		fh := NewFileHeader(params[0])
		return &File{
			Header:  fh,
			Control: NewFileControl(),
		}

	}
	return &File{
		Header:  NewFileHeader(),
		Control: NewFileControl(),
	}
}

// Create creates a valid file and requires that the FileHeader and at least one Batch
func (f *File) Create() error {
	// Requires a valid FileHeader to build FileControl
	if err := f.Header.Validate(); err != nil {
		return err
	}
	// Requires at least one Batch in the new file.
	if len(f.Batches) <= 0 {
		return &FileError{FieldName: "Batchs", Value: strconv.Itoa(len(f.Batches)), Msg: "must have []*Batches to be built"}
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
		f.Batches[i].GetHeader().BatchNumber = batchSeq
		f.Batches[i].GetControl().BatchNumber = batchSeq
		batchSeq++
		// sum file entry and addenda records. Assume batch.Create() batch properly calculated control
		fileEntryAddendaCount = fileEntryAddendaCount + batch.GetControl().EntryAddendaCount
		// add 2 for Batch header/control + entry added count
		totalRecordsInFile = totalRecordsInFile + 2 + batch.GetControl().EntryAddendaCount
		// sum hash from batch control. Assume Batch.Build properly calculated field.
		fileEntryHashSum = fileEntryHashSum + batch.GetControl().EntryHash
		totalDebitAmount = totalDebitAmount + batch.GetControl().TotalDebitEntryDollarAmount
		totalCreditAmount = totalCreditAmount + batch.GetControl().TotalCreditEntryDollarAmount

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
func (f *File) AddBatch(batch Batcher) []Batcher {
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
		msg := fmt.Sprintf(msgFileCalculatedControlEquality, len(f.Batches), f.Control.BatchCount)
		return &FileError{FieldName: "BatchCount", Value: strconv.Itoa(len(f.Batches)), Msg: msg}
	}

	if err := f.isEntryAddendaCount(); err != nil {
		return err
	}

	if err := f.isFileAmount(); err != nil {
		return err
	}

	if err := f.isEntryHash(); err != nil {
		return err
	}

	return nil
}

// isEntryAddenda is prepared by hashing the RDFI’s 8-digit Routing Number in each entry.
//The Entry Hash provides a check against inadvertent alteration of data
func (f *File) isEntryAddendaCount() error {
	count := 0
	// we assume that each batch block has already validated the addenda count is accurate in batch control.
	for _, batch := range f.Batches {
		count += batch.GetControl().EntryAddendaCount
	}
	if f.Control.EntryAddendaCount != count {
		msg := fmt.Sprintf(msgFileCalculatedControlEquality, count, f.Control.EntryAddendaCount)
		return &FileError{FieldName: "EntryAddendaCount", Value: f.Control.EntryAddendaCountField(), Msg: msg}
	}
	return nil
}

// isFileAmount tThe Total Debit and Credit Entry Dollar Amounts Fields contain accumulated
// Entry Detail debit and credit totals within the file
func (f *File) isFileAmount() error {
	debit := 0
	credit := 0
	for _, batch := range f.Batches {
		debit += batch.GetControl().TotalDebitEntryDollarAmount
		credit += batch.GetControl().TotalCreditEntryDollarAmount
	}
	if f.Control.TotalDebitEntryDollarAmountInFile != debit {
		msg := fmt.Sprintf(msgFileCalculatedControlEquality, debit, f.Control.TotalDebitEntryDollarAmountInFile)
		return &FileError{FieldName: "TotalDebitEntryDollarAmountInFile", Value: f.Control.TotalDebitEntryDollarAmountInFileField(), Msg: msg}
	}
	if f.Control.TotalCreditEntryDollarAmountInFile != credit {
		msg := fmt.Sprintf(msgFileCalculatedControlEquality, credit, f.Control.TotalCreditEntryDollarAmountInFile)
		return &FileError{FieldName: "TotalCreditEntryDollarAmountInFile", Value: f.Control.TotalCreditEntryDollarAmountInFileField(), Msg: msg}
	}
	return nil
}

// isEntryHash validates the hash by recalculating the result
func (f *File) isEntryHash() error {
	hashField := f.calculateEntryHash()
	if hashField != f.Control.EntryHashField() {
		msg := fmt.Sprintf(msgFileCalculatedControlEquality, hashField, f.Control.EntryHashField())
		return &FileError{FieldName: "EntryHash", Value: f.Control.EntryHashField(), Msg: msg}
	}
	return nil
}

// calculateEntryHash This field is prepared by hashing the 8-digit Routing Number in each batch.
// The Entry Hash provides a check against inadvertent alteration of data
func (f *File) calculateEntryHash() string {
	hash := 0
	for _, batch := range f.Batches {
		hash = hash + batch.GetControl().EntryHash
	}
	return f.numericField(hash, 10)
}
