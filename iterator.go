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
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"

	"github.com/moov-io/base"
)

// Iterator provides a way to read an ACH file one entry at a time without loading the entire file into memory.
// It is useful for processing large ACH files efficiently.
// The iterator maintains internal state to track the current position in the file.
type Iterator struct {
	reader     *Reader
	scanner    *bufio.Scanner
	cachedLine string
}

// NewIterator creates a new Iterator for reading ACH files from the provided io.Reader.
// The iterator processes the file incrementally, returning one EntryDetail at a time.
func NewIterator(r io.Reader) *Iterator {
	reader := NewReader(strings.NewReader("")) // the input is not used, we rely on .readLine()
	reader.skipBatchAccumulation = true        // don't call .AddBatch(..)

	out := &Iterator{
		reader:  reader,
		scanner: bufio.NewScanner(r),
	}
	return out
}

// SetValidation configures validation options for the iterator's internal reader.
// This affects how strictly the ACH file format is enforced during parsing.
func (i *Iterator) SetValidation(opts *ValidateOpts) {
	if i.reader != nil {
		i.reader.SetValidation(opts)
	}
}

// SetMaxLines limits the number of lines the iterator will process to prevent excessive memory usage or processing time.
// If the limit is exceeded, NextEntry returns an error.
// Set to 0 for no limit (default).
func (i *Iterator) SetMaxLines(max int) {
	if i.reader != nil {
		i.reader.SetMaxLines(max)
	}
}

// GetHeader returns the FileHeader record from the ACH file.
// Returns nil if NextEntry has not been called yet or if the file has no header.
func (i *Iterator) GetHeader() *FileHeader {
	if i.reader != nil {
		return &i.reader.File.Header
	}
	return nil
}

// GetControl returns the FileControl record from the ACH file.
// Returns nil if the end of the file has not been reached yet.
func (i *Iterator) GetControl() *FileControl {
	if i.reader != nil {
		return &i.reader.File.Control
	}
	return nil
}

// NextEntry advances the iterator and returns the next EntryDetail record along with its associated BatchHeader.
// Returns (nil, nil, nil) when there are no more entries.
// Returns an error if the file is malformed or if the max lines limit is exceeded.
// IAT entries are not currently supported.
func (i *Iterator) NextEntry() (*BatchHeader, *EntryDetail, error) {
start:
	// Read the file one line at a time
	line := i.cachedLine
	if line != "" {
		i.cachedLine = "" // clear cache
	} else {
		// Consume lines until we reach a non-empty line
		for i.scanner.Scan() {
			line = i.scanner.Text()
			i.reader.lineNum++
			if i.reader.maxLines > 0 && i.reader.lineNum > i.reader.maxLines {
				return nil, nil, fmt.Errorf("line %d: %w", i.reader.lineNum, ErrFileTooLong)
			}
			if allSpaces(line) {
				continue
			}
			if line != "" {
				break
			}
		}
		// If we've exhausted all lines in the reader then quit
		if line == "" || allSpaces(line) {
			return nil, nil, nil
		}
	}

	if err := i.reader.readLine(line); err != nil {
		if base.Match(err, ErrFileEntryOutsideBatch) {
			// Fake a Batch so we can parse entries
			bh := NewBatchHeader()
			bh.StandardEntryClassCode = PPD
			i.reader.currentBatch, err = NewBatch(bh)
			if err != nil {
				return nil, nil, fmt.Errorf("faking batch for line %d failed: %w", i.reader.lineNum, err)
			}
			if i.reader.currentBatch == nil {
				return nil, nil, fmt.Errorf("failed to create %s batch: %v", bh.StandardEntryClassCode, err)
			}
			if err := i.reader.readLine(line); err != nil {
				return nil, nil, fmt.Errorf("reading line %d with fake BatchHeader failed: %w", i.reader.lineNum, err)
			}
		} else {
			return nil, nil, fmt.Errorf("reading line %d failed: %w", i.reader.lineNum, err)
		}
	}

	if i.reader.currentBatch != nil {
		bh := i.reader.currentBatch.GetHeader()
		entries := i.reader.currentBatch.GetEntries()
		if len(entries) > 0 {
			// Find the next entry to return and consume the file until we run out of
			// addenda records or encounter a batch control/header record.
			returnableEntry := entries[len(entries)-1]

			// Read lines so long as we encounter an addenda or batch control record
			for {
				if i.scanner.Scan() {
					foundLine := i.scanner.Text()
					i.reader.lineNum++
					if i.reader.maxLines > 0 && i.reader.lineNum > i.reader.maxLines {
						return nil, nil, fmt.Errorf("line %d: %w", i.reader.lineNum, ErrFileTooLong)
					}
					if foundLine == "" {
						break
					}
					switch {
					case strings.HasPrefix(foundLine, entryDetailPos):
						i.cachedLine = foundLine
						return bh, returnableEntry, nil

					case strings.HasPrefix(foundLine, entryAddendaPos):
						i.reader.line = foundLine
						if err := i.reader.parseEDAddenda(); err != nil {
							return nil, nil, fmt.Errorf("reading addenda on line %d failed: %w", i.reader.lineNum, err)
						}

						entries := i.reader.currentBatch.GetEntries()
						ed := entries[len(entries)-1]

						return bh, ed, nil

					case strings.HasPrefix(foundLine, batchControlPos):
						// Do nothing with the Batch Control record
						i.reader.currentBatch = nil
						return bh, returnableEntry, nil

					default:
						i.cachedLine = foundLine
						i.reader.currentBatch = nil
						return bh, returnableEntry, nil
					}
				} else {
					break // quit processing if we can't read another line
				}
			}

			return bh, returnableEntry, nil
		} else {
			// We processed the BatchHeader, but need to find an Entry Detail record
			goto start
		}
	}

	goto start
}

func allSpaces(input string) bool {
	for _, r := range input {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return len(input) > 0
}
