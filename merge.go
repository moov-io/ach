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
				}
			}
			if !batchExistsInMerged {
				outf.AddBatch(fs.infiles[i].Batches[j])
				if err := outf.Create(); err != nil {
					return nil, err
				}
				n, err := lineCount(outf)
				if n == 0 || err != nil {
					return nil, fmt.Errorf("problem getting line count of File (header: %#v): %v", outf.Header, err)
				}
				if n > NACHAFileLineLimit {
					outf.RemoveBatch(fs.infiles[i].Batches[j])
					if err := outf.Create(); err != nil { // rebalance ACH file after removing the Batch
						return nil, err
					}
					f := *outf
					fs.locMaxed = append(fs.locMaxed, &f)

					outf = fs.create(outf) // replace output file with the one we just created

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
		}
	}

	return out, nil
}

type mergableFiles struct {
	infiles  []*File
	outfiles []*File
	locMaxed []*File
}

// create returns the index of a newly created file in fs.outfiles given the details from f.Header
func (fs *mergableFiles) create(f *File) *File { // returns the outfiles index of the created file
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
	out.Create()
	fs.outfiles = append(fs.outfiles, out) // add the new outfile

	return out
}

// findOutfile optionally returns a File from fs.files if the FileHeaders match.
// This is done because we append batches into files to minimize the count of output files.
//
// findOutfile will return the existing file (stored in outfiles) if no matching file exists.
func (fs *mergableFiles) findOutfile(f *File) *File {
	inTraceNumbers := getTraceNumbers(f)

	var lookup func(int) *File
	lookup = func(start int) *File {
		// To allow recursive lookups we need to memorize the current index so deeper calls
		// will bypass files with conflicting trace numbers.
		for i := start; i < len(fs.outfiles); i++ {
			if fs.outfiles[i].Header.ImmediateDestination == f.Header.ImmediateDestination &&
				fs.outfiles[i].Header.ImmediateOrigin == f.Header.ImmediateOrigin {

				// found a matching file, so verify the TraceNumber isn't alreay inside
				outTraceNumbers := getTraceNumbers(fs.outfiles[i])
				for trace := range inTraceNumbers {
					// If any of our incoming trace numbers match the existing merged file
					// return the entire file as separate. This keeps partially overlapping
					// batches self-contained.
					if outTraceNumbers.contains(trace) {
						return lookup(i + 1)
					}
				}

				// No conflicting TraceNumber was found, so return current merge file
				return fs.outfiles[i]
			}
		}
		// Record a newly mergable file we can use in future merge attempts
		fs.outfiles = append(fs.outfiles, f)
		return f
	}
	return lookup(0)
}

func lineCount(f *File) (int, error) {
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
			if entries[j].Addenda99 != nil {
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
	return lines, nil
}
