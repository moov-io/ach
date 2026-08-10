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
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/igrmk/treemap/v2"
	"github.com/stretchr/testify/require"
)

func TestMergeFiles__IAT(t *testing.T) {
	iatFile, err := readACHFilepath(filepath.Join("test", "testdata", "iat-debit.ach"))
	require.NoError(t, err)
	require.NotEmpty(t, iatFile.IATBatches)

	out, err := MergeFiles([]*File{iatFile})
	require.NoError(t, err)
	require.Len(t, out, 1)
	require.Len(t, out[0].IATBatches, 1)
	require.NoError(t, out[0].Validate())
}

func TestMergeFiles__IATMultiple(t *testing.T) {
	iatFile, err := readACHFilepath(filepath.Join("test", "testdata", "iat-debit.ach"))
	require.NoError(t, err)

	out, err := MergeFiles([]*File{iatFile, iatFile})
	require.NoError(t, err)
	require.Len(t, out, 1)
	require.Len(t, out[0].IATBatches, 2)
	require.NoError(t, out[0].Validate())
}

func TestMergeFiles__IATAndNonIAT(t *testing.T) {
	ppdFile, err := readACHFilepath(filepath.Join("test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)

	iatFile, err := readACHFilepath(filepath.Join("test", "testdata", "iat-debit.ach"))
	require.NoError(t, err)

	iatFile.Header = ppdFile.Header

	out, err := MergeFiles([]*File{ppdFile, iatFile})
	require.NoError(t, err)
	require.Len(t, out, 1)
	require.Len(t, out[0].Batches, 1)
	require.Len(t, out[0].IATBatches, 1)
	require.NoError(t, out[0].Validate())
}

func TestMergeFiles__IATSameBatchAndTrace(t *testing.T) {
	iatFile1, err := readACHFilepath(filepath.Join("test", "testdata", "iat-debit.ach"))
	require.NoError(t, err)

	iatFile2, err := readACHFilepath(filepath.Join("test", "testdata", "iat-debit.ach"))
	require.NoError(t, err)

	out, err := MergeFiles([]*File{iatFile1, iatFile2})
	require.NoError(t, err)
	require.Len(t, out, 1)
	require.Len(t, out[0].IATBatches, 2)
}

func TestMergeFiles__IATAndNonIATApart(t *testing.T) {
	ppdFile, err := readACHFilepath(filepath.Join("test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)

	iatFile, err := readACHFilepath(filepath.Join("test", "testdata", "iat-debit.ach"))
	require.NoError(t, err)

	out, err := MergeFiles([]*File{ppdFile, iatFile})
	require.NoError(t, err)
	require.Len(t, out, 2)
}

func TestMergeDir__IAT(t *testing.T) {
	dir := t.TempDir()

	src, err := os.Open(filepath.Join("test", "testdata", "iat-debit.ach"))
	require.NoError(t, err)
	t.Cleanup(func() { src.Close() })

	dst, err := os.Create(filepath.Join(dir, "input.ach"))
	require.NoError(t, err)

	_, err = io.Copy(dst, src)
	require.NoError(t, err)
	require.NoError(t, dst.Close())

	var conditions Conditions
	merged, err := MergeDir(dir, conditions, nil)
	require.NoError(t, err)
	require.Len(t, merged, 1)
	require.Len(t, merged[0].IATBatches, 1)
	require.NoError(t, merged[0].Validate())
}

func TestMergeDir__MixedIATAndNonIAT(t *testing.T) {
	dir := t.TempDir()

	ppdSrc, err := os.Open(filepath.Join("test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)
	t.Cleanup(func() { ppdSrc.Close() })

	ppdDst, err := os.Create(filepath.Join(dir, "ppd.ach"))
	require.NoError(t, err)
	_, err = io.Copy(ppdDst, ppdSrc)
	require.NoError(t, err)
	require.NoError(t, ppdDst.Close())

	iatSrc, err := os.Open(filepath.Join("test", "testdata", "iat-debit.ach"))
	require.NoError(t, err)
	t.Cleanup(func() { iatSrc.Close() })

	iatDst, err := os.Create(filepath.Join(dir, "iat.ach"))
	require.NoError(t, err)
	_, err = io.Copy(iatDst, iatSrc)
	require.NoError(t, err)
	require.NoError(t, iatDst.Close())

	var conditions Conditions
	merged, err := MergeDir(dir, conditions, nil)
	require.NoError(t, err)
	require.Len(t, merged, 2)

	var hasIAT, hasNonIAT bool
	for _, f := range merged {
		if len(f.IATBatches) > 0 {
			hasIAT = true
		}
		if len(f.Batches) > 0 {
			hasNonIAT = true
		}
		require.NoError(t, f.Validate())
	}
	require.True(t, hasIAT, "expected at least one file with IAT batches")
	require.True(t, hasNonIAT, "expected at least one file with non-IAT batches")
}

func TestMergeDir__MixedIATSameHeader(t *testing.T) {
	ppdFile, err := readACHFilepath(filepath.Join("test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)

	iatFile, err := readACHFilepath(filepath.Join("test", "testdata", "iat-debit.ach"))
	require.NoError(t, err)

	iatFile.Header = ppdFile.Header

	out, err := MergeFiles([]*File{ppdFile, iatFile})
	require.NoError(t, err)
	require.Len(t, out, 1)

	dir := t.TempDir()
	for i, f := range out {
		var buf bytes.Buffer
		require.NoError(t, NewWriter(&buf).Write(f))

		dst, err := os.Create(filepath.Join(dir, fmt.Sprintf("file-%d.ach", i)))
		require.NoError(t, err)
		_, err = dst.Write(buf.Bytes())
		require.NoError(t, err)
		require.NoError(t, dst.Close())
	}

	var conditions Conditions
	merged, err := MergeDir(dir, conditions, nil)
	require.NoError(t, err)
	require.Len(t, merged, 1)
	require.Len(t, merged[0].Batches, 1)
	require.Len(t, merged[0].IATBatches, 1)
	require.NoError(t, merged[0].Validate())
}

func TestIATBatchHeader__Equal(t *testing.T) {
	bh1 := mockIATBatchHeaderFF()
	bh2 := mockIATBatchHeaderFF()
	require.True(t, bh1.Equal(bh2))
	require.True(t, bh1.Equal(bh1))

	bh2.ServiceClassCode = DebitsOnly
	require.False(t, bh1.Equal(bh2))

	require.False(t, bh1.Equal(nil))
	var nilBh *IATBatchHeader
	require.False(t, nilBh.Equal(bh1))
}

func TestIATEntryDetail__addendaCount(t *testing.T) {
	ed := mockIATEntryDetailWithAddendas()
	require.Equal(t, 7, ed.addendaCount())

	ed.Addenda17 = append(ed.Addenda17, mockAddenda17())
	require.Equal(t, 8, ed.addendaCount())

	ed.Addenda18 = append(ed.Addenda18, mockAddenda18())
	require.Equal(t, 9, ed.addendaCount())

	ed.Addenda99 = mockAddenda99()
	require.Equal(t, 10, ed.addendaCount())
}

func TestFindOutIATBatch(t *testing.T) {
	bh := mockIATBatchHeaderFF()
	var batches []*iatBatch
	var entry *IATEntryDetail

	output := findOutIATBatch(bh, batches, entry)
	require.Nil(t, output)

	batches = append(batches, &iatBatch{
		header:  *bh,
		entries: treemap.New[string, *IATEntryDetail](),
	})
	output = findOutIATBatch(bh, batches, entry)
	require.Equal(t, batches[0], output)

	traceNumber := "231380100000001"
	batches[0].entries.Set(traceNumber, &IATEntryDetail{
		ID:          "1",
		TraceNumber: traceNumber,
	})
	require.True(t, batches[0].entries.Contains(traceNumber))

	output = findOutIATBatch(bh, batches, &IATEntryDetail{
		TraceNumber: traceNumber,
	})
	require.Nil(t, output)
}

func makeIATMergeFile(t *testing.T, entries int, amount int) *File {
	t.Helper()
	file := NewFile()
	file.SetHeader(mockFileHeader())

	bh := mockIATBatchHeaderFF()
	// Align ODFI with file header used above
	bh.ODFIIdentification = "12104288"
	batch := NewIATBatch(bh)

	for i := 0; i < entries; i++ {
		ed := mockIATEntryDetailWithAddendas()
		ed.Amount = amount
		ed.SetTraceNumber(bh.ODFIIdentification, i+1)
		// Keep required addenda sequence numbers consistent with the entry
		ed.Addenda10.EntryDetailSequenceNumber = i + 1
		ed.Addenda11.EntryDetailSequenceNumber = i + 1
		ed.Addenda12.EntryDetailSequenceNumber = i + 1
		ed.Addenda13.EntryDetailSequenceNumber = i + 1
		ed.Addenda14.EntryDetailSequenceNumber = i + 1
		ed.Addenda15.EntryDetailSequenceNumber = i + 1
		ed.Addenda16.EntryDetailSequenceNumber = i + 1
		batch.AddEntry(ed)
	}
	require.NoError(t, batch.Create())
	file.AddIATBatch(batch)
	require.NoError(t, file.Create())
	return file
}

func TestMergeFiles__IATLineLimit(t *testing.T) {
	// Each IAT entry with addenda10-16 is 1+7=8 lines. FileHeader+Control=2,
	// IATBatchHeader+Control=2 → first entry brings us to 12 lines; a second
	// entry at MaxLines=15 must overflow into a new file.
	file := makeIATMergeFile(t, 2, 1000)

	out, err := MergeFilesWith([]*File{file}, Conditions{MaxLines: 15})
	require.NoError(t, err)
	require.Len(t, out, 2)
	for _, f := range out {
		require.Len(t, f.IATBatches, 1)
		require.Len(t, f.IATBatches[0].Entries, 1)
		require.NoError(t, f.Validate())
	}
}

func TestMergeFiles__IATDollarLimit(t *testing.T) {
	// Two IAT credits of $10.00 under a $15.00 max must split into two files.
	file := makeIATMergeFile(t, 2, 1000)

	out, err := MergeFilesWith([]*File{file}, Conditions{MaxDollarAmount: 1500})
	require.NoError(t, err)
	require.Len(t, out, 2)
	require.Equal(t, 1000, out[0].Control.TotalCreditEntryDollarAmountInFile)
	require.Equal(t, 1000, out[1].Control.TotalCreditEntryDollarAmountInFile)
}

func TestMergeFiles__IATFirstEntryOverflow(t *testing.T) {
	// MaxLines smaller than a single IAT entry still produces a file (overflow
	// on the first entry with an empty current batch).
	file := makeIATMergeFile(t, 1, 1000)

	out, err := MergeFilesWith([]*File{file}, Conditions{MaxLines: 5})
	require.NoError(t, err)
	require.Len(t, out, 1)
	require.Len(t, out[0].IATBatches, 1)
}

func TestCountTraceNumbers_IAT(t *testing.T) {
	// Ensure IAT entries are not lost across a merge that splits on conditions.
	file := makeIATMergeFile(t, 3, 500)
	before := 0
	for i := range file.IATBatches {
		before += len(file.IATBatches[i].GetEntries())
	}
	require.Equal(t, 3, before)

	out, err := MergeFilesWith([]*File{file}, Conditions{MaxDollarAmount: 600})
	require.NoError(t, err)
	require.Len(t, out, 3)

	after := 0
	for _, f := range out {
		for i := range f.IATBatches {
			after += len(f.IATBatches[i].GetEntries())
		}
		require.NoError(t, f.Validate())
	}
	require.Equal(t, before, after, "IAT entries must not be dropped when splitting files")
}
