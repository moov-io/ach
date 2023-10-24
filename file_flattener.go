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
	"errors"
	"fmt"
	"sort"
)

var (
	ErrFlattenChangedEntryCount   = errors.New("Flatten operation changed entry and addenda count")
	ErrFlattenChangedDebitAmount  = errors.New("Flatten operation changed total debit entry amount")
	ErrFlattenChangedCreditAmount = errors.New("Flatten operation changed total credit entry amount")
)

// Flatten returns a flattened version of a File, where batches with similar batch headers are consolidated.
//
// Two batches are eligible to be combined if:
//   - their headers match, excluding the batch number (which isn't used in return matching and reflects
//     the final composition of the file.)
//   - they don't contain any entries with common trace numbers, since trace numbers must be unique
//     within a batch.
func Flatten(originalFile *File) (*File, error) {
	originalBatches := make([]mergeable, 0, len(originalFile.Batches)+len(originalFile.IATBatches))

	// Convert batches and IAT batches to "mergeables" for consistent flattening logic
	for i := range originalFile.Batches {
		originalBatches = append(originalBatches, mergeableBatcher{originalFile.Batches[i], nil})
	}
	for i := range originalFile.IATBatches {
		originalBatches = append(originalBatches, mergeableIATBatch{&originalFile.IATBatches[i], nil})
	}

	// Considering bigger batches first allows for the least number of flattened batches
	sort.Slice(originalBatches, func(i, j int) bool {
		return originalBatches[i].GetEntryCount() < originalBatches[j].GetEntryCount()
	})

	// Merge each original batch into a new batch
	newBatchesByHeader := map[string][]mergeable{}
	for i := range originalBatches {
		batch := originalBatches[i]

		var batchToMergeWith mergeable

		batchesWithMatchingHeader, found := newBatchesByHeader[batch.GetHeaderSignature()]
		if found {
			for _, batchWithMatchingHeader := range batchesWithMatchingHeader {
				if canMerge(batch, batchWithMatchingHeader) {
					batchToMergeWith = batchWithMatchingHeader
					break
				}
			}
		}

		if batchToMergeWith == nil {
			newBatchesByHeader[batch.GetHeaderSignature()] = append(newBatchesByHeader[batch.GetHeaderSignature()], batch.Copy())
		} else {
			batchToMergeWith.Consume(batch)
		}
	}

	// Create a new file containing each of our new batches
	newFile := originalFile.addFileHeaderData(NewFile())
	var allBatches []mergeable
	for i := range newBatchesByHeader {
		allBatches = append(allBatches, newBatchesByHeader[i]...)
	}

	// Sort batches by original batch number to roughly maintain batch order in the flattened file
	sort.Slice(allBatches, func(i int, j int) bool { return allBatches[i].GetBatchNumber() < allBatches[j].GetBatchNumber() })

	for i := range allBatches {
		allBatches[i].AddToFile(newFile)
	}

	if err := newFile.Create(); err != nil {
		return nil, err
	}
	if err := newFile.Validate(); err != nil {
		return nil, err
	}

	// Sanity checks; this is kind of a scary operation!
	if originalFile.Control.EntryAddendaCount != newFile.Control.EntryAddendaCount {
		return nil, askForBugReports(ErrFlattenChangedEntryCount)
	}
	if originalFile.Control.TotalDebitEntryDollarAmountInFile != newFile.Control.TotalDebitEntryDollarAmountInFile {
		return nil, askForBugReports(ErrFlattenChangedDebitAmount)
	}
	if originalFile.Control.TotalCreditEntryDollarAmountInFile != newFile.Control.TotalCreditEntryDollarAmountInFile {
		return nil, askForBugReports(ErrFlattenChangedCreditAmount)
	}

	return newFile, nil
}

// FlattenBatches flattens the file's batches by consolidating batches with the same BatchHeader data into one Batch.
// Entries within each flattened batch will be sorted by their TraceNumber field.
func (f *File) FlattenBatches() (*File, error) {
	return Flatten(f)
}

// Determine if two batches can be combined (ie, have the same header and no common trace numbers)
func canMerge(a mergeable, b mergeable) bool {
	traceNumbers := b.GetTraceNumbers()
	for traceNumber := range a.GetTraceNumbers() {
		_, found := traceNumbers[traceNumber]
		if found {
			return false
		}
	}

	return a.GetHeaderSignature() == b.GetHeaderSignature()
}

// Represents either a "normal" batch or an IAT batch
type mergeable interface {
	GetHeaderSignature() string
	GetTraceNumbers() map[string]bool
	Consume(mergeable) error
	GetBatch() interface{}
	GetBatchNumber() int
	Copy() mergeable
	GetEntryCount() int
	AddToFile(*File) error
}

type mergeableBatcher struct {
	batcher      Batcher
	traceNumbers map[string]bool
}

// Batch header excluding the batch number, which isn't important to preserve
func (b mergeableBatcher) GetHeaderSignature() string { return b.batcher.GetHeader().String()[:87] }
func (b mergeableBatcher) GetBatch() interface{}      { return b.batcher }
func (b mergeableBatcher) GetEntryCount() int         { return len(b.batcher.GetEntries()) }
func (b mergeableBatcher) GetBatchNumber() int        { return b.batcher.GetHeader().BatchNumber }

func (b mergeableBatcher) GetTraceNumbers() map[string]bool {
	if b.traceNumbers != nil {
		return b.traceNumbers
	}

	b.traceNumbers = map[string]bool{}
	for _, entry := range b.batcher.GetEntries() {
		b.traceNumbers[entry.TraceNumber] = true
	}

	return b.traceNumbers
}

func (m mergeableBatcher) Consume(mergeableToConsume mergeable) error {
	batcherToConsume, ok := mergeableToConsume.GetBatch().(Batcher)
	if !ok {
		return fmt.Errorf("cannot consume %T - incompatible batch types", mergeableToConsume)
	}

	// Keep the lower of the two batch numbers, to roughly maintain batch order in the flattened file
	if batcherToConsume.GetHeader().BatchNumber < m.batcher.GetHeader().BatchNumber {
		m.batcher.GetHeader().BatchNumber = batcherToConsume.GetHeader().BatchNumber
	}

	entries := batcherToConsume.GetEntries()
	for i := range entries {
		m.batcher.AddEntry(entries[i])
	}
	advEntries := batcherToConsume.GetADVEntries()
	for i := range advEntries {
		m.batcher.AddADVEntry(advEntries[i])
	}

	return nil
}

func (m mergeableBatcher) Copy() mergeable {
	newBatcher, _ := NewBatch(m.batcher.GetHeader())
	newMergeable := mergeableBatcher{newBatcher, nil}
	newMergeable.Consume(m)

	return newMergeable
}

func (m mergeableBatcher) AddToFile(file *File) error {
	// Sort entries by trace number
	sort.Slice(m.batcher.GetEntries(), func(i, j int) bool {
		return m.batcher.GetEntries()[i].TraceNumber < m.batcher.GetEntries()[j].TraceNumber
	})

	err := m.batcher.Create()
	if err != nil {
		return askForBugReports(fmt.Errorf("mergeableBatcher - AddToFile: %v", err))
	}

	m.batcher.GetHeader().BatchNumber = 0

	file.AddBatch(m.batcher)

	return nil
}

type mergeableIATBatch struct {
	iatBatch     *IATBatch
	traceNumbers map[string]bool
}

// Batch header excluding the batch number, which isn't important to preserve
func (b mergeableIATBatch) GetHeaderSignature() string { return b.iatBatch.Header.String()[:87] }
func (b mergeableIATBatch) GetBatch() interface{}      { return *b.iatBatch }
func (b mergeableIATBatch) GetEntryCount() int         { return len(b.iatBatch.Entries) }
func (b mergeableIATBatch) GetBatchNumber() int        { return b.iatBatch.Header.BatchNumber }

func (b mergeableIATBatch) GetTraceNumbers() map[string]bool {
	if b.traceNumbers != nil {
		return b.traceNumbers
	}

	b.traceNumbers = map[string]bool{}
	for _, entry := range b.iatBatch.Entries {
		b.traceNumbers[entry.TraceNumber] = true
	}

	return b.traceNumbers
}

func (m mergeableIATBatch) Consume(mergeableToConsume mergeable) error {
	batchToConsume, ok := mergeableToConsume.GetBatch().(IATBatch)
	if !ok {
		return fmt.Errorf("IAT cannot consume %T - incompatible batch types", mergeableToConsume)
	}

	// Keep the lower of the two batch numbers, to roughly maintain batch order in the flattened file
	if batchToConsume.Header.BatchNumber < m.iatBatch.Header.BatchNumber {
		m.iatBatch.Header.BatchNumber = batchToConsume.Header.BatchNumber
	}

	for _, entry := range batchToConsume.Entries {
		m.iatBatch.AddEntry(entry)
	}

	return nil
}

func (m mergeableIATBatch) Copy() mergeable {
	newIATBatch := NewIATBatch(m.iatBatch.Header)
	newMergeable := mergeableIATBatch{&newIATBatch, nil}
	newMergeable.Consume(m)

	return newMergeable
}

func (m mergeableIATBatch) AddToFile(file *File) error {
	// Sort entries by trace number
	sort.Slice(m.iatBatch.Entries, func(i, j int) bool {
		return m.iatBatch.Entries[i].TraceNumber < m.iatBatch.Entries[j].TraceNumber
	})

	err := m.iatBatch.Create()
	if err != nil {
		return askForBugReports(fmt.Errorf("mergeableIATBatch - AddToFile: %v", err))
	}
	m.iatBatch.Header.BatchNumber = 0

	file.AddIATBatch(*m.iatBatch)

	return nil
}
