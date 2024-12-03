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
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/base"
	"github.com/stretchr/testify/require"
)

func TestIterator(t *testing.T) {
	t.Run("valid PPD file", func(t *testing.T) {
		where := filepath.Join("test", "testdata", "ppd-mixedDebitCredit.ach")

		file := openFile(t, where, nil)
		iter := iteratorFromFile(t, where, nil)

		ensureFileEqualsIterator(t, file, iter)
	})

	t.Run("more valid files", func(t *testing.T) {
		paths := []string{
			filepath.Join("test", "testdata", "two-micro-deposits.ach"),
			filepath.Join("test", "testdata", "web-debit.ach"),
			filepath.Join("test", "testdata", "20110805A.ach"),
		}
		for i := range paths {
			t.Logf("checking %s", paths[i])

			file := openFile(t, paths[i], nil)
			iter := iteratorFromFile(t, paths[i], nil)
			ensureFileEqualsIterator(t, file, iter)
		}
	})

	t.Run("bh-ed-ad-bh-ed-ad-ed-ad", func(t *testing.T) {
		paths := []string{
			filepath.Join("test", "testdata", "bh-ed-ad-bh-ed-ad-ed-ad.ach"),
		}
		for i := range paths {
			t.Logf("checking %s", paths[i])

			file := openFile(t, paths[i], nil)
			iter := iteratorFromFile(t, paths[i], nil)
			ensureFileEqualsIterator(t, file, iter)
		}
	})

	t.Run("skip IAT batches/entries for now", func(t *testing.T) {
		path := filepath.Join("test", "testdata", "iat-debit.ach")

		iter := iteratorFromFile(t, path, nil)

		entries := collectEntries(t, iter)
		require.Empty(t, entries)
	})

	t.Run("return examples", func(t *testing.T) {
		paths := []string{
			filepath.Join("test", "testdata", "return-WEB.ach"),
			filepath.Join("test", "testdata", "return-no-batch-controls.ach"),
			filepath.Join("test", "testdata", "return-no-file-header-control.ach"),
		}

		for i := range paths {
			t.Logf("checking %s", paths[i])

			file := openFile(t, paths[i], nil)
			iter := iteratorFromFile(t, paths[i], nil)
			ensureFileEqualsIterator(t, file, iter)
		}
	})

	t.Run("return without batch header or control", func(t *testing.T) {
		path := filepath.Join("test", "testdata", "return-no-batch-header.ach")

		iter := iteratorFromFile(t, path, nil)

		entries := collectEntries(t, iter)
		require.Len(t, entries, 2)

		// Check first EntryDetail
		ed := entries[0]
		require.Equal(t, "091400606", ed.RDFIIdentification+ed.CheckDigit)
		require.Equal(t, "Paul Jones            ", ed.IndividualName)
		require.Equal(t, "091000017611242", ed.TraceNumber)

		require.Nil(t, ed.Addenda98)
		require.Equal(t, "R01", ed.Addenda99.ReturnCode)
		require.Equal(t, "091400600000001", ed.Addenda99.OriginalTrace)
		require.Equal(t, "091000017611242", ed.Addenda99.TraceNumber)

		// Check second EntryDetail
		ed = entries[1]
		require.Equal(t, "231380104", ed.RDFIIdentification+ed.CheckDigit)
		require.Equal(t, "Best Co. #23          ", ed.IndividualName)
		require.Equal(t, "121042880000001", ed.TraceNumber)

		require.Equal(t, "C01", ed.Addenda98.ChangeCode)
		require.Equal(t, "1918171614", ed.Addenda98.CorrectedData)
		require.Equal(t, "121042880000001", ed.Addenda98.OriginalTrace)
		require.Equal(t, "091012980000088", ed.Addenda98.TraceNumber)
		require.Nil(t, ed.Addenda99)
	})

	t.Run("custom return codes", func(t *testing.T) {
		where := filepath.Join("test", "testdata", "return-PPD-custom-reason-code.ach")

		opts := &ValidateOpts{
			CustomReturnCodes: true,
		}
		file := openFile(t, where, opts)
		iter := iteratorFromFile(t, where, opts)

		ensureFileEqualsIterator(t, file, iter)
	})

	t.Run("blank file", func(t *testing.T) {
		iter := NewIterator(strings.NewReader(""))

		entries := collectEntries(t, iter)
		require.Empty(t, entries)
	})

	t.Run("short lines and has padding", func(t *testing.T) {
		fd, err := os.Open(filepath.Join("test", "testdata", "short-line.ach"))
		require.NoError(t, err)
		t.Cleanup(func() { fd.Close() })

		iter := NewIterator(fd)

		entries := collectEntries(t, iter)
		require.Len(t, entries, 1)
	})

	t.Run("whitespace in file", func(t *testing.T) {
		firstBytes, err := os.ReadFile(filepath.Join(filepath.Join("test", "testdata", "bh-ed-ad-bh-ed-ad-ed-ad.ach")))
		require.NoError(t, err)

		secondBytes, err := os.ReadFile(filepath.Join(filepath.Join("test", "testdata", "return-no-file-header-control.ach")))
		require.NoError(t, err)

		checkBytes := func(t *testing.T, data []byte) {
			t.Helper()

			iter := NewIterator(bytes.NewReader(data))

			entries := collectEntries(t, iter)
			require.Len(t, entries, 4)
		}

		t.Run("first then second", func(t *testing.T) {
			data := append(firstBytes, append([]byte("\n  \n"), secondBytes...)...)
			checkBytes(t, data)
		})

		t.Run("second then first", func(t *testing.T) {
			data := append(secondBytes, append([]byte("\n  \n"), firstBytes...)...)
			checkBytes(t, data)
		})
	})

	t.Run("allSpaces", func(t *testing.T) {
		require.True(t, allSpaces("\n"))
		require.True(t, allSpaces("\n \r\n"))

		require.False(t, allSpaces(""))
		require.False(t, allSpaces("abc"))
	})
}

func openFile(t *testing.T, where string, opts *ValidateOpts) *File {
	t.Helper()

	fd, err := os.Open(where)
	require.NoError(t, err)

	t.Cleanup(func() { fd.Close() })

	r := NewReader(fd)
	r.SetValidation(opts)

	file, err := r.Read()
	if err != nil {
		if !base.Has(err, ErrFileHeader) {
			require.NoError(t, err)
			return nil
		}
	}
	return &file
}

func iteratorFromFile(t *testing.T, where string, opts *ValidateOpts) *Iterator {
	t.Helper()

	fd, err := os.Open(where)
	require.NoError(t, err)

	t.Cleanup(func() { fd.Close() })

	iter := NewIterator(fd)
	iter.SetValidation(opts)
	return iter
}

func ensureFileEqualsIterator(t *testing.T, file *File, iter *Iterator) {
	t.Helper()

	for i := range file.Batches {
		bh := file.Batches[i].GetHeader()
		entries := file.Batches[i].GetEntries()
		for j := range entries {
			ed := entries[j]

			ibh, ied, err := iter.NextEntry()
			require.NoError(t, err, "iterator batch[%d], entry[%d] error: %v", i, j, err)

			require.True(t, bh.Equal(ibh), "batch[%d] headers", i)
			require.Equal(t, ed, ied, "batch[%d] entry[%d] details", i, j)

			require.Equal(t, ed.Addenda02, ied.Addenda02)
			require.Equal(t, ed.Addenda05, ied.Addenda05)

			require.Equal(t, ed.Addenda98, ied.Addenda98)
			require.Equal(t, ed.Addenda98Refused, ied.Addenda98Refused)

			require.Equal(t, ed.Addenda99, ied.Addenda99)
			require.Equal(t, ed.Addenda99Contested, ied.Addenda99Contested)
			require.Equal(t, ed.Addenda99Dishonored, ied.Addenda99Dishonored)
		}
	}
}

func collectEntries(t *testing.T, iter *Iterator) []*EntryDetail {
	t.Helper()

	var entries []*EntryDetail
	for {
		bh, ed, err := iter.NextEntry()
		if err != nil {
			t.Fatal(err)
		}
		if bh == nil && ed == nil {
			break
		}
		if bh != nil && ed != nil {
			entries = append(entries, ed)
		}
	}
	return entries
}
