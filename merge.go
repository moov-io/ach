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
	"context"
	"fmt"
	"io/fs"
	"path/filepath"
	"sync"

	"github.com/igrmk/treemap/v2"
	"golang.org/x/sync/errgroup"
)

const NACHAFileLineLimit = 10000

// MergeFiles is a helper function for consolidating an array of ACH Files into as few files as possible.
// This is useful for optimizing cost and network utilization.
//
// This operation will override batch numbers in each file to ensure they do not collide.
// The ascending batch numbers will start at 1.
//
// Entries with duplicate TraceNumbers are allowed in the same file, but must be in separate batches
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
// Entries with duplicate TraceNumbers are allowed in the same file, but must be in separate batches
// and are automatically separated.
//
// ADV and IAT Batches and Entries are currently not merged together.
//
// Conditions allows for capping the maximum line length or dollar amount of merged files.
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
		err := sorted.add(incoming[i])
		if err != nil {
			return nil, err
		}
	}

	return convertToFiles(sorted, conditions)
}

// MergeDir will consolidate a directory of ACH files into as few files as possible.
// This is useful for optimizing cost and network utilization.
//
// This operation will override batch numbers in each file to ensure they do not collide.
// The ascending batch numbers will start at 1.
//
// Entries with duplicate TraceNumbers are allowed in the same file, but must be in separate batches
// and are automatically separated.
//
// ADV and IAT Batches and Entries are currently not merged together.
//
// MergeDir is typically more performant than MergeFiles as it reads files concurrently while merging occurs.
// This has a more stable cpu and memory usage trend over reading all files into memory and then calling MergeFiles.
//
// File Batches can only be merged if they are unique and routed to and from the same ABA routing numbers.
func MergeDir(dir string, conditions Conditions) ([]*File, error) {
	sorted := &outFile{}
	var setup sync.Once

	// We've observed the slowest part of MergeDir is reading files from disk and
	// parsing them into File structs. We want to have a decent buffer of *File
	// structs that are ready to merge.
	//
	// For example we have observed (on an Intel Mac w/ SSD)
	//    filepath.Walk        50-250µs
	//    queueFileForMerging  20-250ms
	//    sorted.add             1-25ms
	var g errgroup.Group
	parseWorkers := 50 // active ACH Reader's
	discoveredPaths := make(chan string, parseWorkers*2)
	mergableFiles := make(chan *File, parseWorkers*2)

	ctx, cancelFunc := context.WithCancel(context.Background())

	// We are going to scan the directory for files to parse and merge.
	g.Go(func() error {
		defer func() {
			// After we're done reading paths close the channel
			cancelFunc()
			close(discoveredPaths)
		}()
		return filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
			// For now don't delve into subdirs and bubble up errors
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			if path != "" {
				discoveredPaths <- path
			}

			return nil
		})
	})

	// Setup concurrent ACH file parsers which is typically the longest part of merging.
	var wg sync.WaitGroup
	wg.Add(parseWorkers)
	for i := 0; i < parseWorkers; i++ {
		g.Go(func() error {
			defer wg.Done()

			return queueFileForMerging(ctx, discoveredPaths, &setup, sorted, mergableFiles)
		})
	}
	g.Go(func() error {
		wg.Wait()

		// Sending a nil file is the signal to stop merging
		mergableFiles <- nil
		close(mergableFiles)

		return nil
	})

	// Merge ACH files into the final output
	g.Go(func() error {
		for {
			file := <-mergableFiles
			if file == nil {
				return nil
			}

			// accumulate the file into our merged set
			err := sorted.add(file)
			if err != nil {
				return fmt.Errorf("adding file into merged set failed: %w", err)
			}
		}
	})

	err := g.Wait()
	if err != nil {
		return nil, fmt.Errorf("merging %s failed: %w", dir, err)
	}

	return convertToFiles(sorted, conditions)
}

func queueFileForMerging(ctx context.Context, discoveredPaths chan string, setup *sync.Once, sorted *outFile, mergableFiles chan *File) error {
	for {
		select {
		case path := <-discoveredPaths:
			if path == "" {
				return nil
			}

			// Read the file
			file, err := ReadFile(path)
			if file == nil || err != nil {
				return fmt.Errorf("reading %s failed: %w", path, err)
			}

			// Save the first file's header information if it's not already
			setup.Do(func() {
				sorted.header = file.Header
				sorted.validateOpts = file.GetValidation()
			})

			// Only send non-nil files, once this channel receives a nil file we stop merging
			if file != nil {
				mergableFiles <- file
			}

		case <-ctx.Done():
			return nil
		}
	}
}

// outFile is a partial ACH file with batches and forms a linked list to additional files
type outFile struct {
	header  FileHeader
	batches []*batch

	validateOpts *ValidateOpts

	next *outFile
}

func (outf *outFile) add(incoming *File) error {
	outFile := pickOutFile(incoming.Header, outf)
	if outFile == nil {
		return fmt.Errorf("found no outfile: %w", ErrPleaseReportBug)
	}
	outFile.validateOpts = outFile.validateOpts.merge(incoming.GetValidation())

	for j := range incoming.Batches {
		bh := incoming.Batches[j].GetHeader()
		if bh == nil {
			return fmt.Errorf("batch[%d] has nil BatchHeader", j)
		}

		entries := incoming.Batches[j].GetEntries()
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

	return nil
}

func convertToFiles(sorted *outFile, conditions Conditions) ([]*File, error) {
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
