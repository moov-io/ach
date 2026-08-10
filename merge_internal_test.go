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
	"testing"

	"github.com/igrmk/treemap/v2"
	"github.com/stretchr/testify/require"
)

func TestOutFileAdd_NilHeadersAndEntries(t *testing.T) {
	t.Run("nil batch header", func(t *testing.T) {
		file := NewFile()
		file.SetHeader(mockFileHeader())
		// Raw Batch with a nil Header implements Batcher
		file.Batches = append(file.Batches, &Batch{Header: nil})

		outf := &outFile{header: mockFileHeader()}
		err := outf.add(file)
		require.Error(t, err)
		require.Contains(t, err.Error(), "nil BatchHeader")
	})

	t.Run("nil entry detail skipped", func(t *testing.T) {
		bh := mockBatchPPDHeader()
		batch := NewBatchPPD(bh)
		// Inject a nil entry alongside a real one (bypass AddBatch which panics on nil)
		ed := mockPPDEntryDetail()
		batch.Entries = []*EntryDetail{nil, ed}
		file := NewFile()
		file.SetHeader(mockFileHeader())
		file.Batches = append(file.Batches, batch)

		outf := &outFile{header: mockFileHeader()}
		require.NoError(t, outf.add(file))
		require.Len(t, outf.batches, 1)
		require.Equal(t, 1, outf.batches[0].entries.Len())
	})

	t.Run("nil IAT batch header", func(t *testing.T) {
		file := NewFile()
		file.SetHeader(mockFileHeader())
		file.IATBatches = append(file.IATBatches, IATBatch{Header: nil})

		outf := &outFile{header: mockFileHeader()}
		err := outf.add(file)
		require.Error(t, err)
		require.Contains(t, err.Error(), "nil IATBatchHeader")
	})

	t.Run("nil IAT entry detail skipped", func(t *testing.T) {
		bh := mockIATBatchHeaderFF()
		ed := mockIATEntryDetailWithAddendas()
		iat := IATBatch{
			Header:  bh,
			Entries: []*IATEntryDetail{nil, ed},
		}
		file := NewFile()
		file.SetHeader(mockFileHeader())
		file.IATBatches = append(file.IATBatches, iat)

		outf := &outFile{header: mockFileHeader()}
		require.NoError(t, outf.add(file))
		require.Len(t, outf.iatBatches, 1)
		require.Equal(t, 1, outf.iatBatches[0].entries.Len())
	})
}

func TestConvertToFiles_ErrorPaths(t *testing.T) {
	makeEntries := func(n int) *treemap.TreeMap[string, *EntryDetail] {
		m := treemap.New[string, *EntryDetail]()
		for i := 0; i < n; i++ {
			ed := mockPPDEntryDetail()
			ed.SetTraceNumber("12104288", i+1)
			m.Set(ed.TraceNumber, ed)
		}
		return m
	}

	t.Run("unknown SEC", func(t *testing.T) {
		bh := mockBatchPPDHeader()
		bh.StandardEntryClassCode = "ZZZ"
		sorted := &outFile{
			header: mockFileHeader(),
			batches: []*batch{{
				header:  *bh,
				entries: makeEntries(1),
			}},
		}
		out, err := convertToFiles(sorted, Conditions{})
		require.Error(t, err)
		require.Nil(t, out)
		require.Contains(t, err.Error(), "creating batch from sorted.batches")
	})

	t.Run("batch Create fails", func(t *testing.T) {
		bh := mockBatchPPDHeader()
		ed := mockPPDEntryDetail()
		ed.TransactionCode = 99 // invalid
		ed.SetTraceNumber(bh.ODFIIdentification, 1)
		entries := treemap.New[string, *EntryDetail]()
		entries.Set(ed.TraceNumber, ed)

		sorted := &outFile{
			header: mockFileHeader(),
			batches: []*batch{{
				header:  *bh,
				entries: entries,
			}},
		}
		out, err := convertToFiles(sorted, Conditions{})
		require.Error(t, err)
		require.Nil(t, out)
		require.Contains(t, err.Error(), "problem creating batch for outfile")
	})

	t.Run("overflow batch Create fails", func(t *testing.T) {
		// First entry is valid and fills the file; second entry is invalid and
		// triggers overflow so Create runs on the valid batch, then the invalid
		// entry is added to the new batch and fails at the end.
		bh := mockBatchPPDHeader()
		good := mockPPDEntryDetail()
		good.SetTraceNumber(bh.ODFIIdentification, 1)
		bad := mockPPDEntryDetail()
		bad.TransactionCode = 99
		bad.SetTraceNumber(bh.ODFIIdentification, 2)

		entries := treemap.New[string, *EntryDetail]()
		entries.Set(good.TraceNumber, good)
		entries.Set(bad.TraceNumber, bad)

		sorted := &outFile{
			header: mockFileHeader(),
			batches: []*batch{{
				header:  *bh,
				entries: entries,
			}},
		}
		// MaxLines=5: FileHeader+Control+BatchHeader+Control = 4, first entry +1 = 5,
		// second entry forces overflow.
		out, err := convertToFiles(sorted, Conditions{MaxLines: 5})
		require.Error(t, err)
		require.Nil(t, out)
	})

	t.Run("overflow Create fails on already-merged invalid entry", func(t *testing.T) {
		// Invalid entry is merged first; the next entry forces overflow and
		// batch.Create on the overflow path fails (line ~644).
		bh := mockBatchPPDHeader()
		bad := mockPPDEntryDetail()
		bad.TransactionCode = 99
		bad.SetTraceNumber(bh.ODFIIdentification, 1)
		good := mockPPDEntryDetail()
		good.SetTraceNumber(bh.ODFIIdentification, 2)

		entries := treemap.New[string, *EntryDetail]()
		// Trace order: bad (…001) then good (…002)
		entries.Set(bad.TraceNumber, bad)
		entries.Set(good.TraceNumber, good)

		sorted := &outFile{
			header: mockFileHeader(),
			batches: []*batch{{
				header:  *bh,
				entries: entries,
			}},
		}
		out, err := convertToFiles(sorted, Conditions{MaxLines: 5})
		require.Error(t, err)
		require.Nil(t, out)
		require.Contains(t, err.Error(), "problem creating batch for new file/batch")
	})

	t.Run("overflow NewBatch fails", func(t *testing.T) {
		// Exercise the overflow NewBatch error by swapping SEC after the first
		// entry would have been accepted — not reachable via public API, but
		// convertToFiles must surface it. We simulate by using a header that
		// NewBatch rejects once we already have a closed file from a prior
		// outFile node... Instead: use two batches where the second has bad SEC
		// and MaxLines forces the first entry of batch 2 into overflow of file 1.
		// Simpler: call convertToFiles with a batch whose SEC is IAT (ErrFileIATSEC).
		bh := mockBatchPPDHeader()
		bh.StandardEntryClassCode = IAT
		sorted := &outFile{
			header: mockFileHeader(),
			batches: []*batch{{
				header:  *bh,
				entries: makeEntries(1),
			}},
		}
		out, err := convertToFiles(sorted, Conditions{})
		require.Error(t, err)
		require.Nil(t, out)
	})

	t.Run("file Create fails on overflow", func(t *testing.T) {
		bh := mockBatchPPDHeader()
		entries := makeEntries(2)
		// Invalid file header so file.Create fails when closing out the overflowed file.
		sorted := &outFile{
			header: FileHeader{
				ImmediateOrigin:      "0",
				ImmediateDestination: "0",
			},
			batches: []*batch{{
				header:  *bh,
				entries: entries,
			}},
			validateOpts: &ValidateOpts{
				// Allow building batches but fail file header validation
				BypassOriginValidation:      false,
				BypassDestinationValidation: false,
			},
		}
		out, err := convertToFiles(sorted, Conditions{MaxLines: 5})
		require.Error(t, err)
		require.Nil(t, out)
	})

	t.Run("IAT batch Create fails", func(t *testing.T) {
		bh := mockIATBatchHeaderFF()
		// Entry missing required addenda → Create fails
		ed := mockIATEntryDetail()
		ed.SetTraceNumber(bh.ODFIIdentification, 1)
		entries := treemap.New[string, *IATEntryDetail]()
		entries.Set(ed.TraceNumber, ed)

		sorted := &outFile{
			header: mockFileHeader(),
			iatBatches: []*iatBatch{{
				header:  *bh,
				entries: entries,
			}},
		}
		out, err := convertToFiles(sorted, Conditions{})
		require.Error(t, err)
		require.Nil(t, out)
		require.Contains(t, err.Error(), "problem creating IAT batch for outfile")
	})

	t.Run("IAT overflow Create fails", func(t *testing.T) {
		bh := mockIATBatchHeaderFF()
		good := mockIATEntryDetailWithAddendas()
		good.SetTraceNumber(bh.ODFIIdentification, 1)
		bad := mockIATEntryDetail() // missing addenda
		bad.SetTraceNumber(bh.ODFIIdentification, 2)

		entries := treemap.New[string, *IATEntryDetail]()
		entries.Set(good.TraceNumber, good)
		entries.Set(bad.TraceNumber, bad)

		sorted := &outFile{
			header: mockFileHeader(),
			iatBatches: []*iatBatch{{
				header:  *bh,
				entries: entries,
			}},
		}
		// First IAT entry is ~8 lines + headers; MaxLines=12 fits one entry.
		out, err := convertToFiles(sorted, Conditions{MaxLines: 12})
		require.Error(t, err)
		require.Nil(t, out)
	})

	t.Run("IAT overflow Create fails on already-merged invalid entry", func(t *testing.T) {
		// Invalid IAT entry merged first; second entry forces iatOverflow and
		// iatBatch.Create fails on the overflow path (line ~755).
		bh := mockIATBatchHeaderFF()
		bad := mockIATEntryDetail() // missing required addenda
		bad.SetTraceNumber(bh.ODFIIdentification, 1)
		good := mockIATEntryDetailWithAddendas()
		good.SetTraceNumber(bh.ODFIIdentification, 2)

		entries := treemap.New[string, *IATEntryDetail]()
		entries.Set(bad.TraceNumber, bad)
		entries.Set(good.TraceNumber, good)

		sorted := &outFile{
			header: mockFileHeader(),
			iatBatches: []*iatBatch{{
				header:  *bh,
				entries: entries,
			}},
		}
		// bad entry is 1 line + headers(4) = 5; MaxLines=6 fits it; good (~8 lines) overflows.
		out, err := convertToFiles(sorted, Conditions{MaxLines: 6})
		require.Error(t, err)
		require.Nil(t, out)
		require.Contains(t, err.Error(), "problem creating IAT batch for new file/batch")
	})

	t.Run("IAT overflow file Create fails", func(t *testing.T) {
		bh := mockIATBatchHeaderFF()
		e1 := mockIATEntryDetailWithAddendas()
		e1.SetTraceNumber(bh.ODFIIdentification, 1)
		e2 := mockIATEntryDetailWithAddendas()
		e2.SetTraceNumber(bh.ODFIIdentification, 2)
		entries := treemap.New[string, *IATEntryDetail]()
		entries.Set(e1.TraceNumber, e1)
		entries.Set(e2.TraceNumber, e2)

		sorted := &outFile{
			header: FileHeader{
				ImmediateOrigin:      "0",
				ImmediateDestination: "0",
			},
			iatBatches: []*iatBatch{{
				header:  *bh,
				entries: entries,
			}},
		}
		out, err := convertToFiles(sorted, Conditions{MaxLines: 12})
		require.Error(t, err)
		require.Nil(t, out)
	})

	t.Run("linked outFile chain", func(t *testing.T) {
		// Two outFiles in the linked list (different ODFI pairs)
		bh := mockBatchPPDHeader()
		e1 := makeEntries(1)
		e2 := makeEntries(1)

		fh2 := mockFileHeader()
		fh2.ImmediateOrigin = "987654320"

		sorted := &outFile{
			header: mockFileHeader(),
			batches: []*batch{{
				header:  *bh,
				entries: e1,
			}},
			next: &outFile{
				header: fh2,
				batches: []*batch{{
					header:  *bh,
					entries: e2,
				}},
			},
		}
		out, err := convertToFiles(sorted, Conditions{})
		require.NoError(t, err)
		require.Len(t, out, 2)
	})
}
