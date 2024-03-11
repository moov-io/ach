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
	"crypto/rand"
	"math/big"
	"reflect"
	"testing"
)

func BenchmarkMergeFiles(b *testing.B) {
	var (
		batchesPerFile  = 10
		entriesPerBatch = 2

		mergeConditions = Conditions{
			MaxLines:        10_000,
			MaxDollarAmount: 50_000_000_00,
		}
	)

	randomFile := func(B *testing.B) *File {
		B.Helper()

		file := NewFile()
		fh := staticFileHeader()
		file.SetHeader(fh)

		for b := 0; b < batchesPerFile; b++ {
			base, _ := rand.Int(rand.Reader, big.NewInt(1e7))
			traceNumber := int(base.Int64())

			batch := NewBatchPPD(mockBatchPPDHeader())
			for e := 0; e < entriesPerBatch; e++ {
				entry := mockPPDEntryDetail()

				traceNumber += 1
				entry.SetTraceNumber(batch.GetHeader().ODFIIdentification, traceNumber)

				batch.AddEntry(entry)
			}
			file.AddBatch(batch)
		}
		if err := file.Create(); err != nil {
			B.Error(err)
		}

		return file
	}

	randomFiles := func(b *testing.B) (out []*File) {
		b.Helper()
		for i := 0; i < b.N; i++ {
			out = append(out, randomFile(b))
		}
		return
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

	mergeInGroups := func(b *testing.B, groups int) []*File {
		b.Helper()

		files := randomFiles(b)
		indices := makeIndices(len(files), groups)

		b.ReportAllocs()
		b.ResetTimer()
		b.StopTimer()

		var out []*File
		var err error
		if len(indices) > 1 {
			var temp []*File
			for i := 0; i < len(indices)-1; i += 0 {
				b.StartTimer()
				fs, err := MergeFilesWith(files[indices[i]:indices[i+1]], mergeConditions)
				if err != nil {
					b.Fatal(err)
				}
				b.StopTimer()

				i += 1
				temp = append(temp, fs...)
			}
			b.StartTimer()
			out, err = MergeFilesWith(temp, mergeConditions)
		} else {
			b.StartTimer()
			out, err = MergeFilesWith(files, mergeConditions)
		}
		b.StopTimer()

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
		mergeInGroups(b, 1)
	})

	b.Run("MergeFiles_3Groups", func(b *testing.B) {
		mergeInGroups(b, 3)
	})
	b.Run("MergeFiles_5Groups", func(b *testing.B) {
		mergeInGroups(b, 5)
	})
	b.Run("MergeFiles_10Groups", func(b *testing.B) {
		mergeInGroups(b, 10)
	})
	b.Run("MergeFiles_100Groups", func(b *testing.B) {
		mergeInGroups(b, 100)
	})
}
