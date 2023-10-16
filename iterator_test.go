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
	"fmt"
	"os"
	"path/filepath"
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

	t.Run("custom return codes", func(t *testing.T) {
		where := filepath.Join("test", "testdata", "return-PPD-custom-reason-code.ach")

		opts := &ValidateOpts{
			CustomReturnCodes: true,
		}
		file := openFile(t, where, opts)
		iter := iteratorFromFile(t, where, opts)

		ensureFileEqualsIterator(t, file, iter)
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

			require.True(t, bh.Equal(ibh), fmt.Sprintf("batch[%d] headers", i))
			require.Equal(t, ed, ied, fmt.Sprintf("batch[%d] entry[%d] details", i, j))
		}
	}
}
