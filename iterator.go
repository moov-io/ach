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

// Iterator is a data structure for processing an ACH file one entry at a time.
type Iterator struct {
	reader     *Reader
	scanner    *bufio.Scanner
	cachedLine string
}

// NewIterator returns an Iterator
func NewIterator(r io.Reader) *Iterator {
	reader := NewReader(strings.NewReader("")) // the input is not used, we rely on .readLine()
	reader.skipBatchAccumulation = true        // don't call .AddBatch(..)

	out := &Iterator{
		reader:  reader,
		scanner: bufio.NewScanner(r),
	}
	return out
}

func (i *Iterator) SetValidation(opts *ValidateOpts) {
	if i.reader != nil {
		i.reader.SetValidation(opts)
	}
}

// GetHeader will return the FileHeader once encountered by the iterator.
// Call NextEntry() at least once to populate the header.
func (i *Iterator) GetHeader() *FileHeader {
	if i.reader != nil {
		return &i.reader.File.Header
	}
	return nil
}

// GetControl will return the FileControl once encountered by the iterator.
// Call NextEntry() at least once to populate the control.
func (i *Iterator) GetControl() *FileControl {
	if i.reader != nil {
		return &i.reader.File.Control
	}
	return nil
}

// NextEntry will return the next available EntryDetail record and the BatchHeader the entry belongs to.
//
// IAT entries are not currently supported.
func (i *Iterator) NextEntry() (*BatchHeader, *EntryDetail, error) {
	// Clear the reader's File
	defer func() {
		i.reader.File = File{}
	}()

	// Read the file one line at a time
	line := i.cachedLine
	if line != "" {
		i.cachedLine = "" // clear cache
	} else {
		// Consume lines until we reach a non-empty line
		for i.scanner.Scan() {
			line = i.scanner.Text()
			i.reader.lineNum++
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
					if foundLine == "" {
						break
					}
					switch {
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
						return bh, returnableEntry, nil

					default:
						i.cachedLine = foundLine
						return bh, returnableEntry, nil
					}
				}
			}

			return bh, returnableEntry, nil
		} else {
			// We processed the BatchHeader, but need to find an Entry Detail record
			return i.NextEntry()
		}
	}

	return i.NextEntry()
}

func allSpaces(input string) bool {
	for _, r := range input {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return len(input) > 0
}
