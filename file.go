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

// File contains the structures of a parsed ACH File.
type File struct {
	Header  FileHeader
	Batches []Batch
	Control FileControl

	// TODO: remove
	Addenda
}

// addEntryDetail appends an EntryDetail to the Batch
func (f *File) addBatch(batch Batch) []Batch {
	f.Batches = append(f.Batches, batch)
	return f.Batches
}

// Errors specific to parsing a Batch container
var (
	ErrBatchServiceClassMismatch = errors.New("Service Class Code is not the same in Header and Control")
	ErrBatchEntryCountMismatch   = errors.New("Batch Entry Count is out-of-balance with number of entries")
	ErrBatchNumberMismatch       = errors.New("Batch Number is not the same in Header as Control")
	ErrBatchAscendingTraceNumber = errors.New("Trace Numbers on the File are not in ascending sequence within a batch")
)

// Batch holds the Batch Header and Batch Control and all Entry Records
type Batch struct {
	Header  BatchHeader
	Entries []EntryDetail
	Control BatchControl
}

// addEntryDetail appends an EntryDetail to the Batch
func (batch *Batch) addEntryDetail(entry EntryDetail) []EntryDetail {
	batch.Entries = append(batch.Entries, entry)
	return batch.Entries
}

// Validate NACHA rules on the entire batch before being added to a File
func (batch *Batch) Validate() (bool, error) {
	v, err := batch.isServiceClassMismatch()
	if !v {
		return false, err
	}

	v, err = batch.isBatchEntryCountMismatch()
	if !v {
		return false, err
	}

	v, err = batch.isBatchNumberMismatch()
	if !v {
		return false, err
	}

	v, err = batch.isSequenceAscending()
	if !v {
		return false, err
	}

	return true, nil
}

// isServiceClassMismatch validate batch header and control codes are the same
func (batch *Batch) isServiceClassMismatch() (bool, error) {
	if batch.Header.ServiceClassCode != batch.Control.ServiceClassCode {
		return false, ErrBatchServiceClassMismatch
	}
	return true, nil
}

// isBatchEntryCountMismatch validate Entry count is accurate
func (batch *Batch) isBatchEntryCountMismatch() (bool, error) {
	if len(batch.Entries) != batch.Control.EntryAddendaCount {
		return false, ErrBatchEntryCountMismatch
	}
	return true, nil
}

// isBatchNumberMismatch validate batch header and control numbers are the same
func (batch *Batch) isBatchNumberMismatch() (bool, error) {
	if batch.Header.BatchNumber != batch.Control.BatchNumber {
		return false, ErrBatchNumberMismatch
	}
	return true, nil
}

// isSequenceAscending Individual Entry Detail Records within individual batches must
// be in ascending Trace Number order (although Trace Numbers need not necessarily be consecutive).
func (batch *Batch) isSequenceAscending() (bool, error) {
	lastSeq := 0
	for _, seq := range batch.Entries {
		if seq.TraceNumber < lastSeq {
			return false, ErrBatchAscendingTraceNumber
		}
		lastSeq = seq.TraceNumber
	}
	return true, nil
}
