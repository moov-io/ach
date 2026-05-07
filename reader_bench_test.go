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
	"compress/gzip"
	"fmt"
	"io"
	"math/rand/v2"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// Benchmark_ReadLargeFile generates/uses large files to measure read performance and memory usage.
//
// goos: darwin
// goarch: arm64
// pkg: github.com/moov-io/ach
// cpu: Apple M4 Max
// Benchmark_ReadLargeFile/200_batches_100000_entries/reader-16  	       8	 129170396 ns/op	42461288 B/op	  203661 allocs/op
// Benchmark_ReadLargeFile/200_batches_100000_entries/iterator-16  	       30	  38794846 ns/op	42431329 B/op	  203605 allocs/op
// Benchmark_ReadLargeFile/2500_batches_500000_entries/reader-16  	       2	 680684312 ns/op	215550788 B/op	 1042864 allocs/op
// Benchmark_ReadLargeFile/2500_batches_500000_entries/iterator-16  	       6	 199791917 ns/op	215435520 B/op	 1042806 allocs/op
func Benchmark_ReadLargeFile(b *testing.B) {
	cases := []struct {
		batches int
		entries int
	}{
		{batches: 200, entries: 100_000},
		{batches: 2500, entries: 500_000},
	}
	for _, tc := range cases {
		name := fmt.Sprintf("%d batches %d entries", tc.batches, tc.entries)

		b.Run(name, func(b *testing.B) {
			path := writeLargeACHFile(b, tc.batches, tc.entries)

			b.Run("reader", func(b *testing.B) {
				for b.Loop() {
					file, err := NewReader(largeFileReader(b, path)).Read()
					require.NoError(b, err)

					err = file.Validate()
					require.NoError(b, err)
				}
			})

			b.Run("iterator", func(b *testing.B) {
				for b.Loop() {
					iter := NewIterator(largeFileReader(b, path))

					var bh *BatchHeader
					var entry *EntryDetail
					var err error

					for {
						bh, entry, err = iter.NextEntry()
						require.NoError(b, err)

						if bh == nil && entry == nil {
							break
						}
					}
				}
			})
		})
	}
}

func largeFileReader(b *testing.B, path string) io.Reader {
	b.Helper()

	fd, err := os.Open(path)
	require.NoError(b, err)

	gr, err := gzip.NewReader(fd)
	require.NoError(b, err)

	return gr
}

func writeLargeACHFile(tb testing.TB, batchCount, entryCount int) string {
	tb.Helper()

	filename := fmt.Sprintf("PPD-%d-batches-%d-entries.ach.gz", batchCount, entryCount)
	path := filepath.Join("test", "testdata", "large-files", filename)

	// If the file exists already skip creating it
	_, err := os.Stat(path)
	if err == nil {
		return path
	}

	// Create the file
	file := NewFile()
	file.SetHeader(mockFileHeader())

	var batchesCreated int
	var entriesCreated int

	for {
		if batchesCreated >= batchCount || entriesCreated >= entryCount {
			break
		}

		// Create a batch with entries
		batch := mockBatchPPD(tb)

		entriesToCreate := rand.IntN(entryCount - entriesCreated)
		for e := 0; e < entriesToCreate; e++ {
			entry := mockPPDEntryDetail()
			entry.Amount = 1234
			entry.SetTraceNumber(batch.Header.ODFIIdentification, e+10)

			batch.AddEntry(entry)
		}

		err := batch.Create()
		require.NoError(tb, err)

		file.AddBatch(batch)

		batchesCreated++
		entriesCreated += entriesToCreate
	}

	err = file.Create()
	require.NoError(tb, err)

	// write the file as gzipped
	fd, err := os.Create(path)
	require.NoError(tb, err)

	gz, err := gzip.NewWriterLevel(fd, gzip.BestCompression)
	require.NoError(tb, err)

	err = NewWriter(gz).Write(file)
	require.NoError(tb, err)

	err = gz.Close()
	require.NoError(tb, err)

	err = fd.Close()
	require.NoError(tb, err)

	return path
}
