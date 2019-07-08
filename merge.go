// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"bufio"
	"bytes"
	"fmt"
	"time"
)

const NACHAFileLineLimit = 10000

// MergeFiles is a helper function for consolidating an array of ACH Files into as few files
// as possible. This is useful for optimizing cost and network efficiency.
//
// Per NACHA rules files must remain under 10,000 lines (when rendered in their ASCII encoding)
//
// File Batches can only be merged if they are unique and routed to and from the same ABA routing numbers.
func MergeFiles(files []*File) ([]*File, error) {
	fs := &mergableFiles{infiles: files}
	for i := range fs.infiles {
		outf := fs.lookupByHeader(fs.infiles[i])
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

	// TODO(adam): We should also look at consolidating EntryDetail records inside Batches

	return append(fs.locMaxed, fs.outfiles...), nil // return LOC-maxed files and merged files
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

// lookupByHeader optionally returns a File from fs.files if the FileHeaders match.
// This is done because we append batches into files to minimize the count of output files.
//
// lookupByHeader will return the existing file (stored in outfiles) if no matching file exists.
func (fs *mergableFiles) lookupByHeader(f *File) *File {
	for i := range fs.outfiles {
		if fs.outfiles[i].Header.ImmediateDestination == f.Header.ImmediateDestination &&
			fs.outfiles[i].Header.ImmediateOrigin == f.Header.ImmediateOrigin {
			// found a matching file, so return it
			return fs.outfiles[i]
		}
	}
	fs.outfiles = append(fs.outfiles, f)
	return f
}

func lineCount(f *File) (int, error) {
	if len(f.Batches) < 100 {
		// Ignore Files with low batch counts by returning a valid count.
		// Calling Writer.Write() is costly and so we're going to ignore it in easy cases.
		return 1, nil
	}

	var buf bytes.Buffer
	if err := NewWriter(&buf).Write(f); err != nil {
		return 0, err
	}
	lines := 0
	s := bufio.NewScanner(&buf)
	for s.Scan() {
		if v := s.Text(); v != "" {
			lines++
		}
	}
	return lines, nil
}
