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
	ErrBatchAscendingTraceNumber = errors.New("Individual Entry Detail Records within individual batches must be in ascending Trace Number order")
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
	v, err := batch.isSequenceAscending()
	if !v {
		return false, err
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
