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
	"testing"

	"github.com/stretchr/testify/require"
)

func mockAddenda98Refused() *Addenda98Refused {
	add := NewAddenda98Refused()
	add.RefusedChangeCode = "C62"
	add.OriginalTrace = "059999990000003"
	add.OriginalDFI = "05999999"
	add.CorrectedData = "68-6547"
	add.ChangeCode = "C01"
	add.TraceSequenceNumber = "0000002"
	add.TraceNumber = "059999990000001"
	return add
}

func TestAddenda98Refused_Fields(t *testing.T) {
	add := mockAddenda98Refused()

	// shorten some fields
	add.OriginalTrace = "059993"
	add.TraceNumber = "000123"

	require.Equal(t, "C62", add.RefusedChangeCodeField().Code)
	require.Equal(t, "000000000059993", add.OriginalTraceField())
	require.Equal(t, "05999999", add.OriginalDFIField())
	require.Equal(t, "68-6547                      ", add.CorrectedDataField())
	require.Equal(t, "C01", add.ChangeCodeField().Code)
	require.Equal(t, "0000002", add.TraceSequenceNumberField())
	require.Equal(t, "000000000000123", add.TraceNumberField())
}

func TestAddenda98Refused_Read(t *testing.T) {
	original := mockAddenda98Refused()

	read := &Addenda98Refused{}
	read.Parse(original.String())

	require.Equal(t, "C62", read.RefusedChangeCodeField().Code)
	require.Equal(t, "059999990000003", read.OriginalTraceField())
	require.Equal(t, "05999999", read.OriginalDFIField())
	require.Equal(t, "68-6547                      ", read.CorrectedDataField())
	require.Equal(t, "C01", read.ChangeCodeField().Code)
	require.Equal(t, "0000002", read.TraceSequenceNumberField())
	require.Equal(t, "059999990000001", read.TraceNumberField())
}

func TestAddenda98Refused_File(t *testing.T) {
	file := NewFile()
	file.Header = mockFileHeader()

	b := mockBatchCOR(t)
	b.Entries[0].AddendaRecordIndicator = 1
	b.Entries[0].Addenda98Refused = mockAddenda98Refused()
	b.Entries[0].TraceNumber = "121042880000002"
	require.NoError(t, b.Create())
	file.AddBatch(b)
	require.NoError(t, file.Create())

	var buf bytes.Buffer
	err := NewWriter(&buf).Write(file)
	require.NoError(t, err)
	require.Contains(t, buf.String(), "C62")

	read, err := NewReader(&buf).Read()
	require.NoError(t, err)

	require.Len(t, read.Batches, 1)
	entries := read.Batches[0].GetEntries()
	require.Len(t, entries, 1)

	ed := entries[0]
	require.NotNil(t, ed.Addenda98)
	require.Equal(t, "C01", ed.Addenda98.ChangeCode)
	require.NotNil(t, ed.Addenda98Refused)
	require.Equal(t, "C62", ed.Addenda98Refused.RefusedChangeCode)
	require.Equal(t, "C01", ed.Addenda98Refused.ChangeCode)
}
