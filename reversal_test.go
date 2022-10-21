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
