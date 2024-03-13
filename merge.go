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
	"fmt"

	"github.com/igrmk/treemap/v2"
)

const NACHAFileLineLimit = 10000

// MergeFiles is a helper function for consolidating an array of ACH Files into as few files as possible.
// This is useful for optimizing cost and network utilization.
//
// This operation will override batch numbers in each file to ensure they do not collide.
// The ascending batch numbers will start at 1.
//
// Entries with TraceNumbers are allowed in the same file, but must be in separate batches
// and are automatically separated.
//
// ADV and IAT Batches and Entries are currently not merged together.
//
// Old rules limit files to 10,000 lines (when rendered in their ASCII encoding), which
// is the default for this function. Use MergeFilesWith for a higher limit.
//
// File Batches can only be merged if they are unique and routed to and from the same ABA routing numbers.
func MergeFiles(files []*File) ([]*File, error) {
	return MergeFilesWith(files, Conditions{
		MaxLines: NACHAFileLineLimit,
	})
}

// NewMerger returns a Merge which can have custom ValidateOpts
func NewMerger(opts *ValidateOpts) Merger {
	return &merger{opts: opts}
}

// Merge can merge ACH files with custom ValidateOpts
type Merger interface {
	MergeWith(files []*File, conditions Conditions) ([]*File, error)
}

type merger struct {
	opts *ValidateOpts
}

func (m *merger) MergeWith(files []*File, conditions Conditions) ([]*File, error) {
	if m.opts != nil {
		for i := range files {
			files[i].SetValidation(m.opts)
		}
	}
	return MergeFilesWith(files, conditions)
}

type Conditions struct {
	// MaxLines will limit each merged files line count.
	MaxLines int `json:"maxLines"`

	// MaxDollarAmount will limit each merged file's total dollar amount.
	MaxDollarAmount int64 `json:"maxDollarAmount"`
}

// MergeFilesWith is a function for consolidating an array of ACH Files into a few files as possible.
// This is useful for optimizing cost and network utilization.
//
// This operation will override batch numbers in each file to ensure they do not collide.
// The ascending batch numbers will start at 1.
//
// Entries with TraceNumbers are allowed in the same file, but must be in separate batches
// and are automatically separated.
//
// ADV and IAT Batches and Entries are currently not merged together.
//
// Old rules limit files to 10,000 lines (when rendered in their ASCII encoding), which
// is the default for this function. Use MergeFilesWith for a higher limit.
//
// File Batches can only be merged if they are unique and routed to and from the same ABA routing numbers.
func MergeFilesWith(incoming []*File, conditions Conditions) ([]*File, error) {
	if len(incoming) == 0 {
		return nil, nil
	}

	sorted := &outFile{
		header:       incoming[0].Header,
		validateOpts: incoming[0].GetValidation(),
	}

	for i := range incoming {
		outFile := pickOutFile(incoming[i].Header, sorted)
		if outFile == nil {
			return nil, fmt.Errorf("finding outfile from incoming[%d]: %w", i, ErrPleaseReportBug)
		}
		outFile.validateOpts = outFile.validateOpts.merge(incoming[i].GetValidation())

		for j := range incoming[i].Batches {
			bh := incoming[i].Batches[j].GetHeader()
			if bh == nil {
				return nil, fmt.Errorf("incoming[%d].batch[%d] has nil batchHeader", i, j)
			}

			entries := incoming[i].Batches[j].GetEntries()
			for m := range entries {
				// Find a batch where this entry can fit
				b := findOutBatch(bh, outFile.batches, entries[m])

				// No batch can hold this EntryDetail so create one
				if b == nil {
					b = &batch{
						header:  *bh,
						entries: treemap.New[string, *EntryDetail](),
					}
					outFile.batches = append(outFile.batches, b)
				}

				b.entries.Set(entries[m].TraceNumber, entries[m])
			}
		}
	}

	var batchNumber int

	var out []*File
	for {
		// Run through the linked list (sorted.next) until we terminate
		if sorted == nil {
			break
		}

		file := NewFile()
		file.Header = sorted.header

		if sorted.validateOpts != nil {
			file.SetValidation(sorted.validateOpts)
		}

		currentFileLineCount := 2 // FileHeader, FileControl
		var currentFileDollarAmount int

		for i := range sorted.batches {
			nextBatch := sorted.batches[i]

			bh := nextBatch.header
			batchNumber += 1
			bh.BatchNumber = batchNumber

			batch, err := NewBatch(&bh)
			if err != nil {
				return nil, fmt.Errorf("creating batch from sorted.batches[%d] failed: %w", i, err)
			}

			currentFileLineCount += 2 // BatchHeader, BatchControl

			// add each entry detail
			for it := nextBatch.entries.Iterator(); it.Valid(); it.Next() {
				nextEntry := it.Value()

				// Check if we're going to exceed the merge conditions before adding the entry
				entryLineCount := 1 + nextEntry.addendaCount()
				if conditions.MaxLines > 0 {
					// File will be too large, so make a new file and batch
					if currentFileLineCount+entryLineCount > conditions.MaxLines {
						goto overflow
					}
				}

				// File would exceed the dollar amount we're limited to
				if conditions.MaxDollarAmount > 0 {
					if int64(currentFileDollarAmount)+int64(nextEntry.Amount) > conditions.MaxDollarAmount {
						goto overflow
					}
				}

				// Without a condition being exceeded jump into adding the entry in the current batch
				goto merge

			overflow:
				// Close out the current batch and file since we exceeded some limit
				if len(batch.GetEntries()) > 0 {
					err = batch.Create()
					if err != nil {
						return nil, fmt.Errorf("problem creating batch for new file/batch: %w", err)
					}
					file.AddBatch(batch)
				}
				if len(file.Batches) > 0 {
					err = file.Create()
					if err != nil {
						return nil, fmt.Errorf("problem creating file for new file/batch: %w", err)
					}
					out = append(out, file)
				}

				// Reset counters
				currentFileLineCount = 4 // FileHeader, FileControl, BatchHeader, BatchControl
				currentFileDollarAmount = 0

				// Create the new file and batch
				file = NewFile()
				file.Header = sorted.header

				batch, err = NewBatch(&nextBatch.header)
				if err != nil {
					return nil, fmt.Errorf("problem creating overflow batch: %w", err)
				}
				batchNumber += 1
				batch.GetHeader().BatchNumber = batchNumber

			merge:
				// Add the entry to the current batch
				batch.AddEntry(nextEntry)

				currentFileLineCount += 1 + nextEntry.addendaCount()
				currentFileDollarAmount += nextEntry.Amount
			}

			if len(batch.GetEntries()) > 0 {
				err = batch.Create()
				if err != nil {
					return nil, fmt.Errorf("problem creating batch for outfile: %w", err)
				}
				file.AddBatch(batch)
			}
		}

		if len(file.Batches) > 0 {
			err := file.Create()
			if err != nil {
				return nil, fmt.Errorf("problem creating outfile: %w", err)
			}
			out = append(out, file)
		}

		sorted = sorted.next
	}
	return out, nil
}

// outFile is a partial ACH file with batches and forms a linked list to additional files
type outFile struct {
	header  FileHeader
	batches []*batch

	validateOpts *ValidateOpts

	next *outFile
}

// batch contains a BatcHeader and tree of entries sorted by TraceNumber, which allows for
// faster lookup and insertion into an ACH file
type batch struct {
	header  BatchHeader
	entries *treemap.TreeMap[string, *EntryDetail]
}

// pickOutFile will search for an existing outFile matching the FileHeader Origin and Destination.
// If no such file can be found it will create one. A nil file will never be returned.
func pickOutFile(fh FileHeader, file *outFile) *outFile {
	if file == nil {
		return &outFile{
			header: fh,
		}
	}
	if fh.ImmediateOrigin == file.header.ImmediateOrigin &&
		fh.ImmediateDestination == file.header.ImmediateDestination {
		return file
	}
	if file.next == nil {
		file.next = &outFile{
			header: fh,
		}
		return file.next
	}
	return pickOutFile(fh, file.next)
}

// findOutBatch searches an array of batches for one whose BatcHeader matches bh
// and doesn't contain the TraceNumber from entry.
func findOutBatch(bh *BatchHeader, batches []*batch, entry *EntryDetail) *batch {
	for i := range batches {
		if batches[i].header.Equal(bh) {
			// Make sure this batch doesn't contain the TraceNumber already
			var found bool
			if entry != nil {
				found = batches[i].entries.Contains(entry.TraceNumber)
			}
			if !found {
				return batches[i]
			}
		}
	}
	return nil
}
