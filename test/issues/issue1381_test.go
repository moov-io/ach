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

package issues

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/moov-io/ach"

	"github.com/stretchr/testify/require"
)

func TestIssue1381(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "input.ach")

	entryCount := 1000
	makeLargeFileWithoutBreaks(t, filename, entryCount)

	input, err := os.Open(filename)
	require.NoError(t, err)
	t.Cleanup(func() { input.Close() })

	file, err := ach.NewReader(input).Read()
	require.NoError(t, err)
	require.NotNil(t, file)

	require.Len(t, file.Batches, 1)
	require.Len(t, file.Batches[0].GetEntries(), entryCount)
}

func makeLargeFileWithoutBreaks(t *testing.T, where string, entryCount int) {
	t.Helper()

	fd, err := os.Create(where)
	require.NoError(t, err)
	defer fd.Close()

	file := ach.NewFile()
	file.Header = ach.NewFileHeader()
	file.Header.ImmediateDestination = "121042882"
	file.Header.ImmediateOrigin = "91012980"
	file.Header.FileCreationDate = time.Now().AddDate(0, 0, 1).Format("060102") // YYMMDD
	file.Header.ImmediateDestinationName = "Federal Reserve Bank"
	file.Header.ImmediateOriginName = "My Bank Name"

	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.MixedDebitsAndCredits
	bh.CompanyName = "WhlBodyFit"
	bh.CompanyIdentification = "MOOVTESTER"
	bh.StandardEntryClassCode = ach.CCD
	bh.CompanyEntryDescription = "PMTCHECK"
	bh.EffectiveEntryDate = "230411"
	bh.SettlementDate = "102"
	bh.ODFIIdentification = "05310055"

	batch, err := ach.NewBatch(bh)
	require.NoError(t, err)

	for i := 0; i < entryCount; i++ {
		entry := ach.NewEntryDetail()
		entry.TransactionCode = ach.CheckingCredit
		entry.RDFIIdentification = "12345678"
		entry.CheckDigit = "0"
		entry.DFIAccountNumber = "2345532322345"
		entry.Amount = 10001
		entry.IndividualName = "Whole Body Fitness"
		entry.SetTraceNumber(bh.ODFIIdentification, i)
		entry.Category = ach.CategoryForward
		batch.AddEntry(entry)
	}
	require.NoError(t, batch.Create())

	file.AddBatch(batch)
	require.NoError(t, file.Create())

	// Now write it all to the
	fd.WriteString(file.Header.String())
	fd.WriteString(file.Batches[0].GetHeader().String())

	entries := file.Batches[0].GetEntries()
	for i := range entries {
		fd.WriteString(entries[i].String())
	}

	fd.WriteString(file.Batches[0].GetControl().String())
	fd.WriteString(file.Control.String())

	err = fd.Sync()
	require.NoError(t, err)
}
