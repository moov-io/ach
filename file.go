// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

// FileError is an error describing issues validating a file
type FileError struct {
	FieldName string
	Value     string
	Msg       string
}

func (e FileError) Error() string {
	return fmt.Sprintf("%s %s", e.FieldName, e.Msg)
}

// File contains the structures of a parsed ACH File.
type File struct {
	ID         string         `json:"id"`
	Header     FileHeader     `json:"fileHeader"`
	Batches    []Batcher      `json:"batches"`
	IATBatches []IATBatch     `json:"IATBatches"`
	Control    FileControl    `json:"fileControl"`
	ADVControl ADVFileControl `json:"fileADVControl"`

	// NotificationOfChange (Notification of change) is a slice of references to BatchCOR in file.Batches
	NotificationOfChange []Batcher
	// ReturnEntries is a slice of references to file.Batches that contain return entries
	ReturnEntries []Batcher
}

// NewFile constructs a file template.
func NewFile() *File {
	return &File{
		Header:  NewFileHeader(),
		Control: NewFileControl(),
	}
}

type file struct {
	ID string `json:"id"`
}

type fileHeader struct {
	Header FileHeader `json:"fileHeader"`
}

type fileControl struct {
	Control FileControl `json:"fileControl"`
}

type advFileControl struct {
	ADVControl ADVFileControl `json:"advFileControl"`
}

// FileFromJSON attempts to return a *File object assuming the input is valid JSON.
//
// Callers should always check for a nil-error before using the returned file.
//
// The File returned may not be valid and callers should confirm with Validate(). Invalid files may
// be rejected by other Financial Institutions or ACH tools.
func FileFromJSON(bs []byte) (*File, error) {
	if len(bs) == 0 {
		return nil, errors.New("no JSON data provided")
	}

	// read file root level
	var f file
	file := NewFile()
	if err := json.NewDecoder(bytes.NewReader(bs)).Decode(&f); err != nil {
		return nil, fmt.Errorf("problem reading File: %v", err)
	}
	file.ID = f.ID

	// Read FileHeader
	header := fileHeader{
		Header: file.Header,
	}
	if err := json.NewDecoder(bytes.NewReader(bs)).Decode(&header); err != nil {
		return nil, fmt.Errorf("problem reading FileHeader: %v", err)
	}
	file.Header = header.Header

	// Build resulting file
	if err := file.setBatchesFromJSON(bs); err != nil {
		return nil, err
	}

	if !file.IsADV() {
		// Read FileControl
		control := fileControl{
			Control: NewFileControl(),
		}
		if err := json.NewDecoder(bytes.NewReader(bs)).Decode(&control); err != nil {
			return nil, fmt.Errorf("problem reading FileControl: %v", err)
		}
		file.Control = control.Control
	} else {
		// Read ADVFileControl
		advControl := advFileControl{
			ADVControl: NewADVFileControl(),
		}
		if err := json.NewDecoder(bytes.NewReader(bs)).Decode(&advControl); err != nil {
			return nil, fmt.Errorf("problem reading ADVFileControl: %v", err)
		}
		file.ADVControl = advControl.ADVControl
	}

	if !file.IsADV() {
		file.Control.BatchCount = len(file.Batches)
	} else {
		file.ADVControl.BatchCount = len(file.Batches)
	}

	if err := file.Create(); err != nil {
		return file, err
	}
	return file, nil
}

// UnmarshalJSON returns error json struct tag unmarshal is deprecated, use ach.FileFromJSON instead
func (f *File) UnmarshalJSON(p []byte) error {
	return errors.New("json struct tag unmarshal is deprecated, use ach.FileFromJSON instead")
}

type batchesJSON struct {
	Batches []*Batch `json:"batches"`
}

type iatBatchesJSON struct {
	IATBatches []IATBatch `json:"iatBatches"`
}

func setEntryRecordType(e *EntryDetail) {
	e.recordType = "6"
	if e.Addenda02 != nil {
		e.Addenda02.recordType = "7"
		e.Addenda02.TypeCode = "02"
	}
	for _, a := range e.Addenda05 {
		a.recordType = "7"
		a.TypeCode = "05"
	}
	if e.Addenda98 != nil {
		e.Addenda98.recordType = "7"
		e.Addenda98.TypeCode = "98"
	}
	if e.Addenda99 != nil {
		e.Addenda99.recordType = "7"
		e.Addenda99.TypeCode = "99"
	}
}

func setIATEntryRecordType(e *IATEntryDetail) {
	// these values need to be inferred from the json field names
	e.recordType = "6"
	if e.Addenda10 != nil {
		e.Addenda10.recordType = "7"
		e.Addenda10.TypeCode = "10"
	}
	if e.Addenda11 != nil {
		e.Addenda11.recordType = "7"
		e.Addenda11.TypeCode = "11"
	}
	if e.Addenda12 != nil {
		e.Addenda12.recordType = "7"
		e.Addenda12.TypeCode = "12"
	}
	if e.Addenda13 != nil {
		e.Addenda13.recordType = "7"
		e.Addenda13.TypeCode = "13"
	}
	if e.Addenda14 != nil {
		e.Addenda14.recordType = "7"
		e.Addenda14.TypeCode = "14"
	}
	if e.Addenda15 != nil {
		e.Addenda15.recordType = "7"
		e.Addenda15.TypeCode = "15"
	}
	if e.Addenda16 != nil {
		e.Addenda16.recordType = "7"
		e.Addenda16.TypeCode = "16"
	}
	for _, a := range e.Addenda17 {
		a.recordType = "7"
		a.TypeCode = "17"
	}
	for _, a := range e.Addenda18 {
		a.recordType = "7"
		a.TypeCode = "18"
	}
	if e.Addenda98 != nil {
		e.Addenda98.recordType = "7"
		e.Addenda98.TypeCode = "98"
	}
	if e.Addenda99 != nil {
		e.Addenda99.recordType = "7"
		e.Addenda99.TypeCode = "99"
	}
}

// setBatchesFromJson takes bs as JSON and attempts to read out all the Batches within.
//
// We have to break this out as Batcher is an interface (and can't be read by Go's
// json struct tag decoding).
func (f *File) setBatchesFromJSON(bs []byte) error {
	var batches batchesJSON
	var iatBatches iatBatchesJSON

	if err := json.Unmarshal(bs, &batches); err != nil {
		return err
	}
	// Clear out any nil batches
	for i := range f.Batches {
		if f.Batches[i] == nil {
			f.Batches = append(f.Batches[:i], f.Batches[i+1:]...)
		}
	}
	// Add new batches to file
	for i := range batches.Batches {
		if batches.Batches[i] == nil {
			continue
		}
		batch := *batches.Batches[i]
		batch.Header.recordType = batchHeaderPos

		for _, e := range batch.Entries {
			// these values need to be inferred from the json field names
			setEntryRecordType(e)
		}

		if err := batch.build(); err != nil {
			return batch.Error("Invalid Batch", err, batch.Header.ID)
		}

		// Attach a batch with the correct type
		f.Batches = append(f.Batches, ConvertBatchType(batch))
	}

	if err := json.Unmarshal(bs, &iatBatches); err != nil {
		return err
	}

	// Add new iatBatches to file
	for i := range iatBatches.IATBatches {
		if len(iatBatches.IATBatches) == 0 {
			continue
		}

		iatBatch := iatBatches.IATBatches[i]
		iatBatch.Header.recordType = "5"
		for _, e := range iatBatch.Entries {
			setIATEntryRecordType(e)
		}

		if err := iatBatch.build(); err != nil {
			return iatBatch.Error("from JSON", err)
		}
		f.IATBatches = append(f.IATBatches, iatBatch)
	}

	return nil
}

// Create will tabulate and assemble an ACH file into a valid state. This includes
// setting any posting dates, sequence numbers, counts, and sums.
//
// Create requires that the FileHeader and at least one Batch be added to the File.
//
// Create implementations are free to modify computable fields in a file and should
// call the Batch's Validate() function at the end of their execution.
func (f *File) Create() error {
	// Requires a valid FileHeader to build FileControl
	if err := f.Header.Validate(); err != nil {
		return err
	}

	// Requires at least one Batch in the new file.
	if len(f.Batches) <= 0 && len(f.IATBatches) <= 0 {
		return ErrFileNoBatches
	}

	if !f.IsADV() {
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
		for i, iatBatch := range f.IATBatches {
			// create ascending batch numbers
			f.IATBatches[i].GetHeader().BatchNumber = batchSeq
			f.IATBatches[i].GetControl().BatchNumber = batchSeq
			batchSeq++
			// sum file entry and addenda records. Assume batch.Create() batch properly calculated control
			fileEntryAddendaCount = fileEntryAddendaCount + iatBatch.GetControl().EntryAddendaCount
			// add 2 for Batch header/control + entry added count
			totalRecordsInFile = totalRecordsInFile + 2 + iatBatch.GetControl().EntryAddendaCount
			// sum hash from batch control. Assume Batch.Build properly calculated field.
			fileEntryHashSum = fileEntryHashSum + iatBatch.GetControl().EntryHash
			totalDebitAmount = totalDebitAmount + iatBatch.GetControl().TotalDebitEntryDollarAmount
			totalCreditAmount = totalCreditAmount + iatBatch.GetControl().TotalCreditEntryDollarAmount
		}

		// create FileControl from calculated values
		fc := NewFileControl()
		fc.ID = f.ID
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
	} else {
		if err := f.createFileADV(); err != nil {
			return err
		}
	}
	return nil
}

// AddBatch appends a Batch to the ach.File
func (f *File) AddBatch(batch Batcher) []Batcher {
	if batch.Category() == CategoryNOC {
		f.NotificationOfChange = append(f.NotificationOfChange, batch)
	}
	if batch.Category() == CategoryReturn {
		f.ReturnEntries = append(f.ReturnEntries, batch)
	}
	f.Batches = append(f.Batches, batch)
	return f.Batches
}

// AddIATBatch appends a IATBatch to the ach.File
func (f *File) AddIATBatch(iatBatch IATBatch) []IATBatch {
	f.IATBatches = append(f.IATBatches, iatBatch)
	return f.IATBatches
}

// SetHeader allows for header to be built.
func (f *File) SetHeader(h FileHeader) *File {
	f.Header = h
	return f
}

// Validate checks properties of the ACH file to ensure they match NACHA guidelines.
// This includes computing checksums, totals, and sequence orderings.
//
// Validate will never modify the file.
func (f *File) Validate() error {
	if err := f.Header.Validate(); err != nil {
		return err
	}

	if !f.IsADV() {
		// The value of the Batch Count Field is equal to the number of Company/Batch/Header Records in the file.
		if f.Control.BatchCount != (len(f.Batches) + len(f.IATBatches)) {
			return NewErrFileCalculatedControlEquality("BatchCount", len(f.Batches), f.Control.BatchCount)
		}

		for _, b := range f.Batches {
			if err := b.Validate(); err != nil {
				return err
			}
		}

		if err := f.Control.Validate(); err != nil {
			return err
		}
		if err := f.isEntryAddendaCount(false); err != nil {
			return err
		}
		if err := f.isFileAmount(false); err != nil {
			return err
		}
		return f.isEntryHash(false)
	}

	// File contains ADV batches BatchADV

	// The value of the Batch Count Field is equal to the number of Company/Batch/Header Records in the file.
	if f.ADVControl.BatchCount != len(f.Batches) {
		return NewErrFileCalculatedControlEquality("BatchCount", len(f.Batches), f.ADVControl.BatchCount)
	}
	if err := f.ADVControl.Validate(); err != nil {
		return err
	}
	if err := f.isEntryAddendaCount(true); err != nil {
		return err
	}
	if err := f.isFileAmount(true); err != nil {
		return err
	}
	return f.isEntryHash(true)
}

// isEntryAddendaCount is prepared by hashing the RDFI’s 8-digit Routing Number in each entry.
// The Entry Hash provides a check against inadvertent alteration of data
func (f *File) isEntryAddendaCount(IsADV bool) error {
	// IsADV
	// true: the file contains ADV batches
	// false: the file contains other batch types

	count := 0

	// we assume that each batch block has already validated the addenda count is accurate in batch control.

	if !IsADV {
		for _, batch := range f.Batches {
			count += batch.GetControl().EntryAddendaCount
		}
		for _, iatBatch := range f.IATBatches {
			count += iatBatch.GetControl().EntryAddendaCount
		}
		if f.Control.EntryAddendaCount != count {
			return NewErrFileCalculatedControlEquality("EntryAddendaCount", count, f.Control.EntryAddendaCount)
		}
	} else {
		for _, batch := range f.Batches {
			count += batch.GetADVControl().EntryAddendaCount
		}
		if f.ADVControl.EntryAddendaCount != count {
			return NewErrFileCalculatedControlEquality("EntryAddendaCount", count, f.ADVControl.EntryAddendaCount)
		}
	}
	return nil
}

// isFileAmount The Total Debit and Credit Entry Dollar Amounts Fields contain accumulated
// Entry Detail debit and credit totals within the file
func (f *File) isFileAmount(IsADV bool) error {
	// IsADV
	// true: the file contains ADV batches
	// false: the file contains other batch types

	debit := 0
	credit := 0

	if !IsADV {
		for _, batch := range f.Batches {
			debit += batch.GetControl().TotalDebitEntryDollarAmount
			credit += batch.GetControl().TotalCreditEntryDollarAmount
		}
		// IAT
		for _, iatBatch := range f.IATBatches {
			debit += iatBatch.GetControl().TotalDebitEntryDollarAmount
			credit += iatBatch.GetControl().TotalCreditEntryDollarAmount
		}

		if f.Control.TotalDebitEntryDollarAmountInFile != debit {
			return NewErrFileCalculatedControlEquality("TotalDebitEntryDollarAmountInFile", debit, f.Control.TotalDebitEntryDollarAmountInFile)
		}
		if f.Control.TotalCreditEntryDollarAmountInFile != credit {
			return NewErrFileCalculatedControlEquality("TotalCreditEntryDollarAmountInFile", credit, f.Control.TotalCreditEntryDollarAmountInFile)
		}
	} else {
		for _, batch := range f.Batches {
			debit += batch.GetADVControl().TotalDebitEntryDollarAmount
			credit += batch.GetADVControl().TotalCreditEntryDollarAmount
		}

		if f.ADVControl.TotalDebitEntryDollarAmountInFile != debit {
			return NewErrFileCalculatedControlEquality("TotalDebitEntryDollarAmountInFile", debit, f.ADVControl.TotalDebitEntryDollarAmountInFile)
		}
		if f.ADVControl.TotalCreditEntryDollarAmountInFile != credit {
			return NewErrFileCalculatedControlEquality("TotalCreditEntryDollarAmountInFile", credit, f.ADVControl.TotalCreditEntryDollarAmountInFile)

		}
	}
	return nil
}

// isEntryHash validates the hash by recalculating the result
func (f *File) isEntryHash(IsADV bool) error {
	// IsADV
	// true: the file contains ADV batches
	// false: the file contains other batch types but not ADV

	hashField := f.calculateEntryHash(IsADV)

	if !IsADV {
		if hashField != f.Control.EntryHash {
			return NewErrFileCalculatedControlEquality("EntryHash", hashField, f.Control.EntryHash)
		}
	} else {
		if hashField != f.ADVControl.EntryHash {
			return NewErrFileCalculatedControlEquality("EntryHash", hashField, f.ADVControl.EntryHash)
		}
	}
	return nil
}

// calculateEntryHash This field is prepared by hashing the 8-digit Routing Number in each batch.
// The Entry Hash provides a check against inadvertent alteration of data
func (f *File) calculateEntryHash(IsADV bool) int {
	// IsADV
	// true: the file contains ADV batches
	// false: the file contains other batch types but not ADV

	hash := 0

	if !IsADV {
		for _, batch := range f.Batches {
			hash = hash + batch.GetControl().EntryHash
		}
		// IAT
		for _, iatBatch := range f.IATBatches {
			hash = hash + iatBatch.GetControl().EntryHash
		}
	} else {
		for _, batch := range f.Batches {
			hash = hash + batch.GetADVControl().EntryHash
		}
	}
	return hash
}

// IsADV determines if the File is an File containing ADV batches
func (f *File) IsADV() bool {
	ok := false
	for _, batch := range f.Batches {
		if v := batch.GetHeader(); v == nil {
			batch.SetHeader(NewBatchHeader())
		}

		if v := batch.GetControl(); v == nil {
			batch.SetControl(NewBatchControl())
		}

		ok = batch.GetHeader().StandardEntryClassCode == ADV
		if ok {
			break
		}
	}
	return ok
}

func (f *File) createFileADV() error {

	// add 2 for FileHeader/control and reset if build was called twice do to error
	totalRecordsInFile := 2
	batchSeq := 1
	fileEntryAddendaCount := 0
	fileEntryHashSum := 0
	totalDebitAmount := 0
	totalCreditAmount := 0

	for i, batch := range f.Batches {
		// create ascending batch numbers

		if batch.GetHeader().StandardEntryClassCode != ADV {
			return ErrFileADVOnly
		}

		f.Batches[i].GetHeader().BatchNumber = batchSeq
		f.Batches[i].GetADVControl().BatchNumber = batchSeq
		batchSeq++
		// sum file entry and addenda records. Assume batch.Create() batch properly calculated control
		fileEntryAddendaCount = fileEntryAddendaCount + batch.GetADVControl().EntryAddendaCount
		// add 2 for Batch header/control + entry added count
		totalRecordsInFile = totalRecordsInFile + 2 + batch.GetADVControl().EntryAddendaCount
		// sum hash from batch control. Assume Batch.Build properly calculated field.
		fileEntryHashSum = fileEntryHashSum + batch.GetADVControl().EntryHash
		totalDebitAmount = totalDebitAmount + batch.GetADVControl().TotalDebitEntryDollarAmount
		totalCreditAmount = totalCreditAmount + batch.GetADVControl().TotalCreditEntryDollarAmount
	}

	fc := NewADVFileControl()
	fc.ID = f.ID
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
	f.ADVControl = fc

	return nil
}
