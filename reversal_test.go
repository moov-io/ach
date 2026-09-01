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
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestFileCreditReversal(t *testing.T) {
	file, err := ReadJSONFile(filepath.Join("test", "testdata", "ppd-valid.json"))
	require.NoError(t, err)

	effectiveEntryDate := time.Now().In(time.UTC)
	err = file.Reversal(effectiveEntryDate)
	require.NoError(t, err)

	b1 := file.Batches[0]
	require.Equal(t, "REVERSAL", b1.GetHeader().CompanyEntryDescription)

	entries := b1.GetEntries()
	require.Len(t, entries, 1)
	require.Equal(t, CheckingDebit, entries[0].TransactionCode)
}

func TestFileDebitReversal(t *testing.T) {
	file, err := ReadFile(filepath.Join("test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)

	effectiveEntryDate := time.Now().In(time.UTC)
	err = file.Reversal(effectiveEntryDate)
	require.NoError(t, err)

	b1 := file.Batches[0]
	require.Equal(t, "REVERSAL", b1.GetHeader().CompanyEntryDescription)

	entries := b1.GetEntries()
	require.Len(t, entries, 1)
	require.Equal(t, CheckingCredit, entries[0].TransactionCode)
}

func TestReversal_GL(t *testing.T) {
	// Debit
	file, err := ReadFile(filepath.Join("test", "testdata", "gl-debit.ach"))
	require.NoError(t, err)

	b1 := file.Batches[0]
	entries := b1.GetEntries()
	require.Len(t, entries, 1)
	require.Equal(t, GLDebit, entries[0].TransactionCode)

	// Reverse
	effectiveEntryDate := time.Now().In(time.UTC)
	err = file.Reversal(effectiveEntryDate)
	require.NoError(t, err)

	b1 = file.Batches[0]
	require.Equal(t, "REVERSAL", b1.GetHeader().CompanyEntryDescription)

	entries = b1.GetEntries()
	require.Len(t, entries, 1)
	require.Equal(t, GLCredit, entries[0].TransactionCode)

	// Reverse the REVERSAL
	err = file.Reversal(effectiveEntryDate)
	require.NoError(t, err)

	b1 = file.Batches[0]
	require.Equal(t, "REVERSAL", b1.GetHeader().CompanyEntryDescription)

	entries = b1.GetEntries()
	require.Len(t, entries, 1)
	require.Equal(t, GLDebit, entries[0].TransactionCode)
}

func TestReversal_LoanCredit(t *testing.T) {
	// Credit
	file, err := ReadFile(filepath.Join("test", "testdata", "loan-credit.ach"))
	require.NoError(t, err)

	b1 := file.Batches[0]
	entries := b1.GetEntries()
	require.Len(t, entries, 1)
	require.Equal(t, LoanCredit, entries[0].TransactionCode)

	// Reverse
	effectiveEntryDate := time.Now().In(time.UTC)
	err = file.Reversal(effectiveEntryDate)
	require.NoError(t, err)

	b1 = file.Batches[0]
	require.Equal(t, "REVERSAL", b1.GetHeader().CompanyEntryDescription)

	entries = b1.GetEntries()
	require.Len(t, entries, 1)
	require.Equal(t, LoanDebit, entries[0].TransactionCode)

	// Reverse the REVERSAL
	err = file.Reversal(effectiveEntryDate)
	require.NoError(t, err)

	b1 = file.Batches[0]
	require.Equal(t, "REVERSAL", b1.GetHeader().CompanyEntryDescription)

	entries = b1.GetEntries()
	require.Len(t, entries, 1)
	require.Equal(t, LoanCredit, entries[0].TransactionCode)
}

func TestReversal_TEL(t *testing.T) {
	file := NewFile()
	file.Header = mockFileHeader()

	batch := mockBatchTEL(t)
	file.AddBatch(batch)

	err := file.Create()
	require.NoError(t, err)

	err = file.Reversal(time.Now())
	require.NoError(t, err)

	err = file.Create()
	require.NoError(t, err)

	err = file.Validate()
	require.NoError(t, err)
}

// The debit-only SEC types all reject credits, which is right for a forward
// batch and wrong for a reversal: File.Reversal turns the debits into credits
// and sets ServiceClassCode to CreditsOnly, so the batch the library just
// produced failed its own validation. TEL was the only type that knew.
//
// Reversal returns nil either way, so nothing told the caller. The file was
// simply invalid from then on.
func TestReversal_DebitOnlySECTypes(t *testing.T) {
	tests := []struct {
		name  string
		batch func(*testing.T) Batcher
	}{
		{"ARC", func(t *testing.T) Batcher { return mockBatchARC(t) }},
		{"BOC", func(t *testing.T) Batcher { return mockBatchBOC(t) }},
		{"POP", func(t *testing.T) Batcher { return mockBatchPOP(t) }},
		{"RCK", func(t *testing.T) Batcher { return mockBatchRCK(t) }},
		{"XCK", func(t *testing.T) Batcher { return mockBatchXCK(t) }},
		{"TRC", func(t *testing.T) Batcher { return mockBatchTRC(t) }},
		{"TRX", func(t *testing.T) Batcher { return mockBatchTRX(t) }},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			file := NewFile()
			file.Header = mockFileHeader()
			file.AddBatch(tc.batch(t))

			err := file.Create()
			require.NoError(t, err)
			require.NoError(t, file.Validate())

			err = file.Reversal(time.Now())
			require.NoError(t, err)

			err = file.Create()
			require.NoError(t, err)

			err = file.Validate()
			require.NoError(t, err)
		})
	}
}

func TestBatch_IsReversal(t *testing.T) {
	batch := mockBatchARC(t)

	require.False(t, batch.IsReversal())

	// Nacha spells it in capitals and left justified, so the comparison has to
	// survive both a different case and the padding a fixed-width file carries.
	for _, description := range []string{"REVERSAL", "reversal", "  REVERSAL  "} {
		batch.GetHeader().CompanyEntryDescription = description
		require.True(t, batch.IsReversal(), description)
	}

	batch.GetHeader().CompanyEntryDescription = "REVERSALS"
	require.False(t, batch.IsReversal())
}
