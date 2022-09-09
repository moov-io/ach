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
	"testing"

	"github.com/stretchr/testify/require"
)

func mockAddenda99Contested() *Addenda99Contested {
	addenda99 := NewAddenda99Contested()
	addenda99.ContestedReturnCode = "R71"
	addenda99.OriginalEntryTraceNumber = "059999990000301"
	addenda99.DateOriginalEntryReturned = "167"
	addenda99.OriginalReceivingDFIIdentification = "12391871"
	addenda99.OriginalSettlementDate = "164"
	addenda99.ReturnTraceNumber = "779999990000301"
	addenda99.ReturnSettlementDate = "165"
	addenda99.ReturnReasonCode = "01"
	addenda99.DishonoredReturnTraceNumber = "889999990000301"
	addenda99.DishonoredReturnSettlementDate = "166"
	addenda99.DishonoredReturnReasonCode = "67"
	addenda99.TraceNumber = "123918710000001"
	return addenda99
}

func TestAddenda99Contested__Fields(t *testing.T) {
	addenda99 := mockAddenda99Contested()

	// shorten some fields
	addenda99.OriginalEntryTraceNumber = "0599999900301"
	addenda99.ReturnTraceNumber = "1239187101"
	addenda99.TraceNumber = "1239187100001"

	require.Equal(t, "R71", addenda99.ContestedReturnCodeField())
	require.Equal(t, "000599999900301", addenda99.OriginalEntryTraceNumberField())
	require.Equal(t, "000167", addenda99.DateOriginalEntryReturnedField())
	require.Equal(t, "12391871", addenda99.OriginalReceivingDFIIdentificationField())
	require.Equal(t, "164", addenda99.OriginalSettlementDateField())
	require.Equal(t, "000001239187101", addenda99.ReturnTraceNumberField())
	require.Equal(t, "165", addenda99.ReturnSettlementDateField())
	require.Equal(t, "01", addenda99.ReturnReasonCodeField())
	require.Equal(t, "889999990000301", addenda99.DishonoredReturnTraceNumberField())
	require.Equal(t, "166", addenda99.DishonoredReturnSettlementDateField())
	require.Equal(t, "67", addenda99.DishonoredReturnReasonCodeField())
	require.Equal(t, "001239187100001", addenda99.TraceNumberField())
}

func TestAddenda99Contested(t *testing.T) {
	file := NewFile()
	file.SetHeader(mockFileHeader())
	file.Control = mockFileControl()

	batch, err := NewBatch(mockBatchHeader())
	require.NoError(t, err)

	ed := mockEntryDetail()
	ed.AddendaRecordIndicator = 1
	ed.Category = CategoryDishonoredReturnContested

	addenda99 := mockAddenda99()
	ed.Addenda99 = addenda99
	ed.Addenda99Dishonored = mockAddenda99Dishonored()
	ed.Addenda99Contested = mockAddenda99Contested()
	require.Equal(t, 3, ed.addendaCount())

	batch.AddEntry(ed)
	require.NoError(t, batch.Create())
	require.Equal(t, 4, batch.GetControl().EntryAddendaCount)

	file.AddBatch(batch)
	require.NoError(t, file.Create())

	require.Equal(t, 1, len(file.Batches))
	require.Equal(t, 1, len(file.Batches[0].GetEntries()))
	require.Equal(t, 4, file.Batches[0].GetControl().EntryAddendaCount)
	require.Equal(t, 4, file.Control.EntryAddendaCount)

	require.NoError(t, file.Validate())

	var buf bytes.Buffer
	if err := NewWriter(&buf).Write(file); err != nil {
		t.Fatal(err)
	}

	path := filepath.Join("examples", "testdata", "contested-return.ach")
	err = os.WriteFile(path, buf.Bytes(), 0600)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddenda99Contested__Read(t *testing.T) {
	file, err := ReadFile(filepath.Join("test", "testdata", "contested_addenda.txt"))
	require.NoError(t, err)

	require.Len(t, file.Batches, 1)

	entries := file.Batches[0].GetEntries()
	require.Len(t, entries, 1)

	ed := entries[0]
	require.NotNil(t, ed.Addenda99Contested)

	addenda := ed.Addenda99Contested
	require.Equal(t, "R72", addenda.ContestedReturnCode)
	require.Equal(t, "123456780000069", addenda.OriginalEntryTraceNumber)
	require.Equal(t, "      ", addenda.DateOriginalEntryReturned)
	require.Equal(t, "75639218", addenda.OriginalReceivingDFIIdentification)
	require.Equal(t, "   ", addenda.OriginalSettlementDate)
	require.Equal(t, "756392180000001", addenda.ReturnTraceNumber)
	require.Equal(t, "067", addenda.ReturnSettlementDate)
	require.Equal(t, "01", addenda.ReturnReasonCode)
	require.Equal(t, "123456780000070", addenda.DishonoredReturnTraceNumber)
	require.Equal(t, "218", addenda.DishonoredReturnSettlementDate)
	require.Equal(t, "68", addenda.DishonoredReturnReasonCode)
	require.Equal(t, "364275034310088", addenda.TraceNumber)
}
