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
	"time"
)

const NACHAFileLineLimit = 10000

// MergeFiles is a helper function for consolidating an array of ACH Files into as few files
// as possible. This is useful for optimizing cost and network efficiency.
//
// This operation will override batch numbers in each file to ensure they do not collide.
// The ascending batch numbers will start at 1.
//
// Duplicate TraceNumbers will not be allowed in the same file. Multiple files will be created.
//
// Per NACHA rules files must remain under 10,000 lines (when rendered in their ASCII encoding)
//
// File Batches can only be merged if they are unique and routed to and from the same ABA routing numbers.
func MergeFiles(files []*File) ([]*File, error) {
	return mergeFilesHelper(files, Conditions{
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

func MergeFilesWith(files []*File, conditions Conditions) ([]*File, error) {
	return mergeFilesHelper(files, conditions)
}

func mergeFilesHelper(files []*File, conditions Conditions) ([]*File, error) {
	fs := &mergableFiles{infiles: files}
	for i := range fs.infiles {
		if fs.infiles[i] == nil {
			continue // skip nil Files
		}
		outf := fs.findOutfile(fs.infiles[i])
		for j := range fs.infiles[i].Batches {
			batchExistsInMerged := false
			for k := range outf.Batches {
				if fs.infiles[i].Batches[j].Equal(outf.Batches[k]) {
					batchExistsInMerged = true

					batch, err := mergeEntries(outf.Batches[k], fs.infiles[i].Batches[j])
					if err != nil {
						return nil, err
					}
					outf.Batches[k] = batch
					break
				}
			}
			if !batchExistsInMerged {
				outf.AddBatch(fs.infiles[i].Batches[j])
				if err := outf.Create(); err != nil {
					return nil, err
				}

				fileTooLong := (conditions.MaxLines > 0 && (lineCount(outf) > conditions.MaxLines))
				fileTooExpensive := (conditions.MaxDollarAmount > 0 && (dollarAmount(outf) > conditions.MaxDollarAmount))

				// Split into a new file if needed
				if fileTooLong || fileTooExpensive {
					// In the event of a file with one batch and one entry just keep it
					if len(outf.Batches) > 1 || len(outf.Batches[0].GetEntries()) > 1 {
						outf.RemoveBatch(fs.infiles[i].Batches[j])
						if err := outf.Create(); err != nil { // rebalance ACH file after removing the Batch
							return nil, err
						}
						f := *outf
						fs.locMaxed = append(fs.locMaxed, &f)
					}

					outf = fs.swapLocMaxedFile(outf) // replace output file with the one we just created

					outf.AddBatch(fs.infiles[i].Batches[j])
					if err := outf.Create(); err != nil {
						return nil, err
					}
				}
			}
		}
	}

	// Return LOC-maxed files and merged files
	out := append(fs.locMaxed, fs.outfiles...)

	// Override batch numbers to ensure they don't collide
	for _, f := range out {
		for i, b := range f.Batches {
			b.GetHeader().BatchNumber = i + 1
			b.GetControl().BatchNumber = i + 1
			// Tabulate each Batch after combining them
			if err := b.Create(); err != nil {
				return out, err
			}
		}
		// Tabulate the files before returning them
		if err := f.Create(); err != nil {
			return out, err
		}
	}
	return out, nil
}

type mergableFiles struct {
	infiles  []*File
	outfiles []*File
	locMaxed []*File
}

// swapLocMaxedFile replaces an ACH file that is over the Nacha line limit with an empty file containing
// a matching FileHeader record. This allows future iterations inside of MergeFiles to append
func (fs *mergableFiles) swapLocMaxedFile(f *File) *File {
	now := time.Now()

	// remove the current file from outfiles
	for i := range fs.outfiles {
		if fs.outfiles[i].Header.ImmediateDestination == f.Header.ImmediateDestination &&
			fs.outfiles[i].Header.ImmediateOrigin == f.Header.ImmediateOrigin {
			// found a matching file, so remove it from fs.outfiles
			fs.outfiles = append(fs.outfiles[:i], fs.outfiles[i+1:]...)
			goto next
		}
	}
next:
	out := NewFile()
	out.Header = f.Header
	out.Header.FileCreationDate = now.Format("060102") // YYMMDD
	out.Header.FileCreationTime = now.Format("1504")   // HHmm
	out.SetValidation(f.validateOpts)
	out.Create()
	fs.outfiles = append(fs.outfiles, out) // add the new outfile

	return out
}

// findOutfile optionally returns a File from fs.files if the FileHeaders match.
// This is done because we append batches into files to minimize the count of output files.
//
// findOutfile will return the existing file (stored in outfiles) if no matching file exists.
func (fs *mergableFiles) findOutfile(f *File) *File {
	var lookup func(int) *File
	lookup = func(start int) *File {
		// To allow recursive lookups we need to memorize the current index so deeper calls
		// will bypass files with conflicting trace numbers.
		for i := start; i < len(fs.outfiles); i++ {
			if fs.outfiles[i].Header.ImmediateDestination == f.Header.ImmediateDestination &&
				fs.outfiles[i].Header.ImmediateOrigin == f.Header.ImmediateOrigin {

				// found a matching file, so verify the TraceNumber isn't alreay inside
				for inB := range f.Batches {
					inEntries := f.Batches[inB].GetEntries()
					for inE := range inEntries {
						// Compare against outfiles
						for outB := range fs.outfiles[i].Batches {
							outEntries := fs.outfiles[i].Batches[outB].GetEntries()
							for outE := range outEntries {
								// If any of our incoming trace numbers match the existing merged file
								// return the entire file as separate. This keeps partially overlapping
								// batches self-contained.
								if inEntries[inE].TraceNumber == outEntries[outE].TraceNumber {
									return lookup(i + 1)
								}
							}
						}
					}
				}

				// No conflicting TraceNumber was found, so return current merge file
				return fs.outfiles[i]
			}
		}
		// Record a newly mergable File/FileHeader we can use in future merge attempts
		outf := NewFile()
		outf.Header = f.Header
		outf.SetValidation(f.validateOpts)
		outf.Control = f.Control
		fs.outfiles = append(fs.outfiles, outf)
		return outf
	}
	return lookup(0)
}

func mergeEntries(b1, b2 Batcher) (Batcher, error) {
	b, _ := NewBatch(b1.GetHeader())
	entries := sortEntriesByTraceNumber(append(b1.GetEntries(), b2.GetEntries()...))
	for i := range entries {
		b.AddEntry(entries[i])
	}
	b.SetControl(b1.GetControl())
	if err := b.Create(); err != nil {
		return nil, err
	}
	return b, nil
}

func lineCount(f *File) int {
	lines := 2 // FileHeader, FileControl
	for i := range f.Batches {
		lines += 2 // BatchHeader, BatchControl
		entries := f.Batches[i].GetEntries()
		for j := range entries {
			lines++
			if entries[j].Addenda02 != nil {
				lines++
			}
			lines += len(entries[j].Addenda05)
			if entries[j].Addenda98 != nil {
				lines++
			}
			if entries[j].Addenda98Refused != nil {
				lines++
			}
			if entries[j].Addenda99 != nil {
				lines++
			}
			if entries[j].Addenda99Dishonored != nil {
				lines++
			}
			if entries[j].Addenda99Contested != nil {
				lines++
			}
		}
	}
	for i := range f.IATBatches {
		lines += 2 // IATBatchHeader, BatchControl
		for j := range f.IATBatches[i].Entries {
			lines++
			if f.IATBatches[i].Entries[j].Addenda10 != nil {
				lines++
			}
			if f.IATBatches[i].Entries[j].Addenda11 != nil {
				lines++
			}
			if f.IATBatches[i].Entries[j].Addenda12 != nil {
				lines++
			}
			if f.IATBatches[i].Entries[j].Addenda13 != nil {
				lines++
			}
			if f.IATBatches[i].Entries[j].Addenda14 != nil {
				lines++
			}
			if f.IATBatches[i].Entries[j].Addenda15 != nil {
				lines++
			}
			if f.IATBatches[i].Entries[j].Addenda16 != nil {
				lines++
			}

			lines += len(f.IATBatches[i].Entries[j].Addenda17)
			lines += len(f.IATBatches[i].Entries[j].Addenda18)

			if f.IATBatches[i].Entries[j].Addenda98 != nil {
				lines++
			}
			if f.IATBatches[i].Entries[j].Addenda99 != nil {
				lines++
			}
		}
	}
	return lines
}

func dollarAmount(f *File) int64 {
	var total int64
	for i := range f.Batches {
		entries := f.Batches[i].GetEntries()
		for j := range entries {
			total += int64(entries[j].Amount)
		}
	}
	return total
}
