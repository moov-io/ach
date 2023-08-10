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
	"strconv"
	"testing"
	"time"
)

func BenchmarkMergeFiles(b *testing.B) {
	randomFile := func(B *testing.B) *File {
		B.Helper()

		file := NewFile()
		file.SetHeader(mockFileHeader())
		file.Control = mockFileControl()

		base := fmt.Sprintf("%d", time.Now().UnixMilli())[:8]

		for b := 0; b < 10; b++ {
			batch := mockBatchPPD()
			for e := 1; e < 20; e++ {
				entry := mockPPDEntryDetail()

				traceNumber, err := strconv.ParseInt(fmt.Sprintf("%s%d%d", base, b, e), 10, 64)
				if err != nil {
					B.Error(err)
				}
				entry.SetTraceNumber(batch.GetHeader().ODFIIdentification, int(traceNumber))

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

	b.Run("MergeFiles", func(b *testing.B) {
		files := randomFiles(b)
		b.ReportAllocs()
		b.ResetTimer()

		out, err := MergeFiles(files)
		b.StopTimer()

		if err != nil {
			b.Error(err)
		}
		if n := len(out); n == 0 {
			b.Error("no files merged")
		}
	})

	b.Run("Merger.MergeFile", func(b *testing.B) {
		files := randomFiles(b)
		m := NewMerger(nil)
		var cond Conditions
		b.ReportAllocs()
		b.ResetTimer()

		for i := range files {
			err := m.MergeFile(files[i], cond)
			if err != nil {
				b.Error(err)
			}
		}
		b.StopTimer()

		if n := len(m.Files()); n == 0 {
			b.Error("no files merged")
		}
	})
}
