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

	"github.com/stretchr/testify/require"
)

// TestBatchCATX_Replies validates replies to CTX, ATX, and TRX entries. These SEC codes
// count addenda records in the entry, so a reply must count its one Addenda99,
// Addenda99Dishonored, or Addenda99Contested record.
//
// See https://github.com/moov-io/ach/issues/1875
func TestBatchCATX_Replies(t *testing.T) {
	replies := []struct {
		name     string
		category string
		attach   func(*EntryDetail)
	}{
		{"return", CategoryReturn, func(e *EntryDetail) { e.Addenda99 = mockAddenda99() }},
		{"dishonored", CategoryDishonoredReturn, func(e *EntryDetail) { e.Addenda99Dishonored = mockAddenda99Dishonored() }},
		{"contested", CategoryDishonoredReturnContested, func(e *EntryDetail) { e.Addenda99Contested = mockAddenda99Contested() }},
	}
	secs := []struct {
		sec             string
		header          func() *BatchHeader
		entry           func() *EntryDetail
		transactionCode int
	}{
		{CTX, mockBatchCTXHeader, mockCTXEntryDetail, CheckingReturnNOCCredit},
		{ATX, mockBatchATXHeader, mockATXEntryDetail, CheckingReturnNOCCredit},
		{TRX, mockBatchTRXHeader, mockTRXEntryDetail, CheckingReturnNOCDebit},
	}

	for _, sec := range secs {
		for _, reply := range replies {
			t.Run(sec.sec+"/"+reply.name, func(t *testing.T) {
				newBatch := func(addendaRecords int, opts *ValidateOpts) (Batcher, error) {
					entry := sec.entry()
					entry.Addenda05 = nil
					entry.TransactionCode = sec.transactionCode
					entry.Category = reply.category
					reply.attach(entry)
					entry.SetCATXAddendaRecords(addendaRecords)
					entry.AddendaRecordIndicator = 1

					batch, err := NewBatch(sec.header())
					require.NoError(t, err)
					batch.SetValidation(opts)
					batch.AddEntry(entry)
					return batch, batch.Create()
				}

				batch, err := newBatch(1, nil)
				require.NoError(t, err)
				require.NoError(t, batch.Validate())
				require.Equal(t, "0001", batch.GetEntries()[0].CATXAddendaRecordsField())

				// The reply addenda record must be counted.
				_, err = newBatch(0, nil)
				require.ErrorContains(t, err, "AddendaCount 1 addendum found where 0 are expected")

				_, err = newBatch(0, &ValidateOpts{UnequalAddendaCounts: true})
				require.NoError(t, err)
			})
		}
	}
}
