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
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func BenchmarkMergeFiles(b *testing.B) {
	var (
		batchesPerFile  = 3
		entriesPerBatch = 20

		mergeConditions = Conditions{
			MaxLines:        10_000,
			MaxDollarAmount: 50_000_000_00,
		}
	)

	type options struct {
		withValidateOpts bool
	}

	randomFile := func(B *testing.B, opts options) *File {
		B.Helper()

		file := NewFile()
		fh := staticFileHeader()
		file.SetHeader(fh)

		if opts.withValidateOpts {
			file.SetValidation(&ValidateOpts{
				PreserveSpaces: true,
			})
		}

		for b := 0; b < batchesPerFile; b++ {
			base, _ := rand.Int(rand.Reader, big.NewInt(1e7))
			traceNumber := int(base.Int64() - int64(entriesPerBatch+1))
			if traceNumber < 0 {
				traceNumber = 1
			}

			batch := NewBatchPPD(mockBatchPPDHeader())
			for e := 0; e < entriesPerBatch; e++ {
				entry := mockPPDEntryDetail()
				entry.Amount = (b + 1) * (e + 1) * 100

				traceNumber += 1
				entry.SetTraceNumber(batch.GetHeader().ODFIIdentification, traceNumber)

				batch.AddEntry(entry)
			}
			err := batch.Create()
			if err != nil {
				B.Fatal(err)
			}
			file.AddBatch(batch)
		}
		if err := file.Create(); err != nil {
			B.Error(err)
		}

		return file
	}

	randomFiles := func(b *testing.B, opts options) (out []*File) {
		b.Helper()
		for i := 0; i < b.N; i++ {
			out = append(out, randomFile(b, opts))
		}
		return
	}

	writeFiles := func(b *testing.B, dir string, files []*File) []string {
		b.Helper()

		out := make([]string, len(files))
		for i := range files {
			where := filepath.Join(dir, fmt.Sprintf("ACH-%d.txt", i))
			out[i] = where

			var buf bytes.Buffer
			err := NewWriter(&buf).Write(files[i])
			if err != nil {
				b.Fatal(err)
			}

			err = os.WriteFile(where, buf.Bytes(), 0600)
			if err != nil {
				b.Fatal(err)
			}
		}
		return out
	}

	makeIndices := func(total, groups int) []int {
		if groups == 1 || groups >= total {
			return []int{total}
		}
		xs := []int{0}
		i := 0
		for {
			if i > total {
				break
			}
			i += total / groups
			if i < total {
				xs = append(xs, i)
			}
		}
		return append(xs, total)
	}

	// Verify basics about makeIndices
	indices := makeIndices(122, 5)
	if len(indices) != 7 {
		b.Fatalf("unexpected number of indices: %#v", indices)
	}
	if !reflect.DeepEqual(indices, []int{0, 24, 48, 72, 96, 120, 122}) {
		b.Fatalf("unexpected indices: %#v", indices)
	}

	mergeInGroups := func(b *testing.B, groups int, opts options) []*File {
		b.Helper()

		// Write files to disk without capturing them in memory
		dir := b.TempDir()
		files := randomFiles(b, opts)
		paths := writeFiles(b, dir, files)

		b.ReportAllocs()
		b.ResetTimer()

		// We need to read files from disk, which needs to be accounted for in our cpu/memory
		files, err := ReadFiles(paths)
		if err != nil {
			b.Fatal(err)
		}
		indices := makeIndices(len(files), groups)

		b.ReportAllocs()

		var out []*File
		if len(indices) > 1 {
			var temp []*File
			for i := 0; i < len(indices)-1; i += 0 {
				fs, err := MergeFilesWith(files[indices[i]:indices[i+1]], mergeConditions)
				if err != nil {
					b.Fatal(err)
				}

				i += 1
				temp = append(temp, fs...)
			}
			out, err = MergeFilesWith(temp, mergeConditions)
		} else {
			out, err = MergeFilesWith(files, mergeConditions)
		}

		if err != nil {
			b.Error(err)
		}
		if n := len(out); n == 0 {
			b.Error("no files merged")
		} else {
			b.Logf("%d files merged into %d files", len(files), len(out))
		}

		return out
	}

	b.Run("MergeFiles", func(b *testing.B) {
		mergeInGroups(b, 1, options{})
	})

	b.Run("MergeFiles ValidateOpts", func(b *testing.B) {
		mergeInGroups(b, 1, options{
			withValidateOpts: true,
		})
	})

	b.Run("MergeDir", func(b *testing.B) {
		dir := b.TempDir()
		var opts options
		incoming := randomFiles(b, opts)
		writeFiles(b, dir, incoming)

		b.ReportAllocs()
		b.ResetTimer()

		merged, err := MergeDir(dir, mergeConditions)
		if err != nil {
			b.Fatal(err)
		}
		b.Logf("merged %d files into %d files", len(incoming), len(merged))
	})

	b.Run("MergeDir ValidateOpts", func(b *testing.B) {
		dir := b.TempDir()
		opts := options{
			withValidateOpts: true,
		}
		incoming := randomFiles(b, opts)
		writeFiles(b, dir, incoming)

		b.ReportAllocs()
		b.ResetTimer()

		merged, err := MergeDir(dir, mergeConditions)
		if err != nil {
			b.Fatal(err)
		}
		b.Logf("merged %d files into %d files", len(incoming), len(merged))
	})

	b.Run("MergeFiles 3Groups", func(b *testing.B) {
		mergeInGroups(b, 3, options{})
	})
	b.Run("MergeFiles 5Groups", func(b *testing.B) {
		mergeInGroups(b, 5, options{})
	})
	b.Run("MergeFiles 10Groups", func(b *testing.B) {
		mergeInGroups(b, 10, options{})
	})
	b.Run("MergeFiles 100Groups", func(b *testing.B) {
		mergeInGroups(b, 100, options{})
	})
}
