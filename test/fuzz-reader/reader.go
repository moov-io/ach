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

package fuzzreader

import (
	"bytes"

	"github.com/moov-io/ach"
)

// Return codes (from go-fuzz docs)
//
// The function must return 1 if the fuzzer should increase priority
// of the given input during subsequent fuzzing (for example, the input is
// lexically correct and was parsed successfully); -1 if the input must not be
// added to corpus even if gives new coverage; and 0 otherwise; other values are
// reserved for future use.
func Fuzz(data []byte) int {
	r := ach.NewReader(bytes.NewReader(data))
	f, err := r.Read()
	if err != nil {
		// if f != nil {
		// 	panic(fmt.Sprintf("f != nil on err != nil: %v", f))
		// }
		return 0
	}

	// Check f (as ach.File)
	if f.ID != "" {
		return 1
	}

	// FileHeader
	if n := checkFileHeader(&f); n > 0 {
		return n
	}

	// Batches
	if n := checkFileBatches(&f); n > 0 {
		return n
	}

	// FileControl
	if n := checkFileControl(&f); n > 0 {
		return n
	}

	// Changes / Returns
	if len(f.NotificationOfChange) > 0 {
		return 1
	}
	if len(f.ReturnEntries) > 0 {
		return 1
	}

	return 0 // increase priority of input
}

func checkFileHeader(f *ach.File) int {
	if f.Header.ID != "" {
		return 1
	}
	if f.Header.ImmediateDestination != "" || f.Header.ImmediateOrigin != "" {
		return 1
	}
	if f.Header.FileCreationDate != "" || f.Header.FileCreationTime != "" {
		return 1
	}
	if f.Header.FileIDModifier != "" {
		return 1
	}
	if f.Header.ImmediateDestinationName != "" || f.Header.ImmediateOriginName != "" {
		return 1
	}
	if f.Header.ReferenceCode != "" {
		return 1
	}
	return -2
}

func checkFileBatches(f *ach.File) int {
	if len(f.Batches) > 0 || len(f.IATBatches) > 0 {
		return 1
	}
	return -2
}

func checkFileControl(f *ach.File) int {
	if f.Control.ID != "" {
		return 1
	}
	if f.Control.BatchCount > 0 || f.Control.BlockCount > 0 {
		return 1
	}
	if f.Control.EntryAddendaCount > 0 || f.Control.EntryHash > 0 {
		return 1
	}
	if f.Control.TotalDebitEntryDollarAmountInFile > 0 {
		return 1
	}
	if f.Control.TotalCreditEntryDollarAmountInFile > 0 {
		return 1
	}
	return -2
}
