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

// TestBatchCATX_RepliesKeepOriginalAddendaCount validates replies to CTX, ATX, and TRX entries.
// A reply copies Number of Addenda Records from the entry it answers and carries one reply addenda,
// so the field is only checked for being numeric.
//
// See https://github.com/moov-io/ach/issues/1875
func TestBatchCATX_RepliesKeepOriginalAddendaCount(t *testing.T) {
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
				newBatch := func(addendaRecords string) (Batcher, error) {
					entry := sec.entry()
					entry.Addenda05 = nil
					entry.TransactionCode = sec.transactionCode
					entry.Category = reply.category
					reply.attach(entry)
					entry.IndividualName = addendaRecords + entry.CATXReceivingCompanyField() + "  "
					entry.AddendaRecordIndicator = 1

					batch, err := NewBatch(sec.header())
					require.NoError(t, err)
					batch.AddEntry(entry)
					return batch, batch.Create()
				}

				for _, count := range []string{"0000", "0001", "0002", "0012"} {
					batch, err := newBatch(count)
					require.NoError(t, err, count)
					require.NoError(t, batch.Validate(), count)
					require.Equal(t, count, batch.GetEntries()[0].CATXAddendaRecordsField())
				}

				_, err := newBatch("00A1")
				require.ErrorIs(t, err, ErrNonNumeric)
			})
		}
	}
}

// TestBatchCATX_ForwardAddendaCount verifies forward CTX, ATX, and TRX entries still compare
// Number of Addenda Records to their Addenda05 records.
func TestBatchCATX_ForwardAddendaCount(t *testing.T) {
	secs := []struct {
		sec    string
		header func() *BatchHeader
		entry  func() *EntryDetail
	}{
		{CTX, mockBatchCTXHeader, mockCTXEntryDetail},
		{ATX, mockBatchATXHeader, mockATXEntryDetail},
		{TRX, mockBatchTRXHeader, mockTRXEntryDetail},
	}
	for _, sec := range secs {
		t.Run(sec.sec, func(t *testing.T) {
			entry := sec.entry()
			entry.Addenda05 = nil
			entry.AddAddenda05(mockAddenda05())
			entry.SetCATXAddendaRecords(2)
			entry.AddendaRecordIndicator = 1

			batch, err := NewBatch(sec.header())
			require.NoError(t, err)
			batch.AddEntry(entry)
			require.ErrorContains(t, batch.Create(), "AddendaCount 1 addendum found where 2 are expected")
		})
	}
}

// TestBatchCTX_ReadReturnKeepsOriginalAddendaCount reads a file with a CTX return that keeps
// the original entry's Number of Addenda Records.
func TestBatchCTX_ReadReturnKeepsOriginalAddendaCount(t *testing.T) {
	entry := mockCTXEntryDetail()
	entry.Addenda05 = nil
	entry.TransactionCode = CheckingReturnNOCCredit
	entry.Category = CategoryReturn
	entry.Addenda99 = mockAddenda99()
	entry.SetCATXAddendaRecords(12)
	entry.AddendaRecordIndicator = 1

	batch, err := NewBatch(mockBatchCTXHeader())
	require.NoError(t, err)
	batch.AddEntry(entry)
	require.NoError(t, batch.Create())

	file := NewFile()
	file.SetHeader(mockFileHeader())
	file.AddBatch(batch)
	require.NoError(t, file.Create())

	var buf bytes.Buffer
	require.NoError(t, NewWriter(&buf).Write(file))

	read, err := NewReader(&buf).Read()
	require.NoError(t, err)
	require.Len(t, read.Batches, 1)

	got := read.Batches[0].GetEntries()[0]
	require.Equal(t, CategoryReturn, got.Category)
	require.Equal(t, "0012", got.CATXAddendaRecordsField())
	require.NotNil(t, got.Addenda99)
}

// TestBatchCTX_ReturnKeepsAddenda05AndCopiedCount validates a CTX return that still carries
// Addenda05 records and copies Number of Addenda Records from the original entry.
// addendaFieldInclusionReturn allows Addenda05 on CTX returns; the count is only checked
// for being numeric.
func TestBatchCTX_ReturnKeepsAddenda05AndCopiedCount(t *testing.T) {
	entry := mockCTXEntryDetail()
	entry.Addenda05 = nil
	entry.AddAddenda05(mockAddenda05())
	entry.AddAddenda05(mockAddenda05())
	entry.TransactionCode = CheckingReturnNOCCredit
	entry.Category = CategoryReturn
	entry.Addenda99 = mockAddenda99()
	entry.SetCATXAddendaRecords(12)
	entry.AddendaRecordIndicator = 1

	batch, err := NewBatch(mockBatchCTXHeader())
	require.NoError(t, err)
	batch.AddEntry(entry)
	require.NoError(t, batch.Create())
	require.NoError(t, batch.Validate())

	got := batch.GetEntries()[0]
	require.Equal(t, "0012", got.CATXAddendaRecordsField())
	require.Len(t, got.Addenda05, 2)
	require.NotNil(t, got.Addenda99)
}

// TestBatchCTX_ReturnWithOffset creates a CTX reply that keeps the original Number of Addenda
// Records and then balances the batch with WithOffset. The OFFSET entry copies the reply
// Category and uses IndividualName "OFFSET", so CATXAddendaRecordsField is "OFFS".
func TestBatchCTX_ReturnWithOffset(t *testing.T) {
	replies := []struct {
		name     string
		category string
		attach   func(*EntryDetail)
	}{
		{"return", CategoryReturn, func(e *EntryDetail) { e.Addenda99 = mockAddenda99() }},
		{"dishonored", CategoryDishonoredReturn, func(e *EntryDetail) { e.Addenda99Dishonored = mockAddenda99Dishonored() }},
		{"contested", CategoryDishonoredReturnContested, func(e *EntryDetail) { e.Addenda99Contested = mockAddenda99Contested() }},
	}
	for _, reply := range replies {
		t.Run(reply.name, func(t *testing.T) {
			entry := mockCTXEntryDetail()
			entry.Addenda05 = nil
			entry.TransactionCode = CheckingReturnNOCCredit
			entry.Category = reply.category
			reply.attach(entry)
			entry.SetCATXAddendaRecords(12)
			entry.AddendaRecordIndicator = 1

			batch, err := NewBatch(mockBatchCTXHeader())
			require.NoError(t, err)
			batch.AddEntry(entry)
			batch.WithOffset(&Offset{
				RoutingNumber: "121042882",
				AccountNumber: "123456789",
				AccountType:   OffsetChecking,
				Description:   "test offset",
			})
			require.NoError(t, batch.Create())
			require.NoError(t, batch.Validate())

			entries := batch.GetEntries()
			require.Len(t, entries, 2)
			require.Equal(t, "0012", entries[0].CATXAddendaRecordsField())
			require.Equal(t, offsetIndividualName, entries[1].IndividualName)
			require.Equal(t, reply.category, entries[1].Category)
			require.Equal(t, "OFFS", entries[1].CATXAddendaRecordsField())
		})
	}
}

// TestBatchCATX_OffsetEntryOnReply adds an OFFSET companion to CTX, ATX, and TRX returns
// that keep a copied Number of Addenda Records.
func TestBatchCATX_OffsetEntryOnReply(t *testing.T) {
	secs := []struct {
		sec             string
		header          func() *BatchHeader
		entry           func() *EntryDetail
		transactionCode int
		offset          func() *EntryDetail
	}{
		{
			sec:             CTX,
			header:          mockBatchCTXHeader,
			entry:           mockCTXEntryDetail,
			transactionCode: CheckingReturnNOCCredit,
			offset: func() *EntryDetail {
				ed := NewEntryDetail()
				ed.TransactionCode = CheckingDebit
				ed.SetRDFI("121042882")
				ed.DFIAccountNumber = "123456789"
				ed.Amount = 25000
				ed.IndividualName = offsetIndividualName
				ed.Category = CategoryReturn
				return ed
			},
		},
		{
			sec:             ATX,
			header:          mockBatchATXHeader,
			entry:           mockATXEntryDetail,
			transactionCode: CheckingReturnNOCCredit,
			offset: func() *EntryDetail {
				ed := NewEntryDetail()
				ed.TransactionCode = CheckingZeroDollarRemittanceCredit
				ed.SetRDFI("121042882")
				ed.DFIAccountNumber = "123456789"
				ed.Amount = 0
				ed.IndividualName = offsetIndividualName
				ed.Category = CategoryReturn
				ed.SetOriginalTraceNumber("121042880000002")
				return ed
			},
		},
		{
			sec:             TRX,
			header:          mockBatchTRXHeader,
			entry:           mockTRXEntryDetail,
			transactionCode: CheckingReturnNOCDebit,
			offset: func() *EntryDetail {
				ed := NewEntryDetail()
				ed.TransactionCode = CheckingDebit
				ed.SetRDFI("121042882")
				ed.DFIAccountNumber = "123456789"
				ed.Amount = 100
				ed.IndividualName = offsetIndividualName
				ed.Category = CategoryReturn
				return ed
			},
		},
	}
	for _, sec := range secs {
		t.Run(sec.sec, func(t *testing.T) {
			entry := sec.entry()
			entry.Addenda05 = nil
			entry.TransactionCode = sec.transactionCode
			entry.Category = CategoryReturn
			entry.Addenda99 = mockAddenda99()
			entry.SetCATXAddendaRecords(12)
			entry.AddendaRecordIndicator = 1

			header := sec.header()
			if sec.sec == CTX {
				header.ServiceClassCode = MixedDebitsAndCredits
			}
			batch, err := NewBatch(header)
			require.NoError(t, err)
			batch.AddEntry(entry)
			batch.AddEntry(sec.offset())
			require.NoError(t, batch.Create())
			require.NoError(t, batch.Validate())

			entries := batch.GetEntries()
			require.Len(t, entries, 2)
			require.Equal(t, "0012", entries[0].CATXAddendaRecordsField())
			require.Equal(t, offsetIndividualName, entries[1].IndividualName)
			require.Equal(t, "OFFS", entries[1].CATXAddendaRecordsField())
		})
	}
}
