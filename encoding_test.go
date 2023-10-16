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

	"golang.org/x/net/html/charset"

	"github.com/stretchr/testify/require"
)

func TestExtendedCharacters(t *testing.T) {
	t.Run("IAT File", func(t *testing.T) {
		b := mockIATBatch(t)
		b.Entries[0].Addenda10.Name = "John¦Smith"
		b.Entries[0].Addenda11.OriginatorName = "My¦Bank"
		b.Entries[0].Addenda12.OriginatorCountryPostalCode = `US*10036\`
		b.Entries[0].Addenda13.ODFIName = "My¦Bank2"
		b.Entries[0].Addenda14.RDFIName = "Other¦Bank"
		b.Entries[0].Addenda15.ReceiverIDNumber = "123¦456"
		b.Entries[0].Addenda16.ReceiverCountryPostalCode = `US*10036\`

		addenda17 := NewAddenda17()
		addenda17.PaymentRelatedInformation = "Thing1¦Thing2"
		addenda17.SequenceNumber = 1
		b.Entries[0].Addenda17 = append(b.Entries[0].Addenda17, addenda17)

		addenda18 := mockAddenda18()
		addenda18.ForeignCorrespondentBankIDNumber = "456¦123"
		b.Entries[0].Addenda18 = append(b.Entries[0].Addenda18, addenda18)

		require.NoError(t, b.Create())

		file := NewFile()
		file.SetHeader(mockFileHeader())
		file.AddIATBatch(b)
		require.NoError(t, file.Create())

		// Cycle file through write/read
		var buf bytes.Buffer
		err := NewWriter(&buf).Write(file)
		require.NoError(t, err)

		parsed, err := NewReader(&buf).Read()
		require.NoError(t, err)

		// Verify field details
		entries := parsed.IATBatches[0].GetEntries()
		require.Equal(t, "John¦Smith", entries[0].Addenda10.Name)
		require.Equal(t, "My¦Bank", entries[0].Addenda11.OriginatorName)
		require.Equal(t, `US*10036\`, entries[0].Addenda12.OriginatorCountryPostalCode)
		require.Equal(t, "My¦Bank2", entries[0].Addenda13.ODFIName)
		require.Equal(t, "Other¦Bank", entries[0].Addenda14.RDFIName)
		require.Equal(t, "123¦456", entries[0].Addenda15.ReceiverIDNumber)
		require.Equal(t, `US*10036\`, entries[0].Addenda16.ReceiverCountryPostalCode)
		require.Equal(t, "Thing1¦Thing2", entries[0].Addenda17[0].PaymentRelatedInformation)
		require.Equal(t, "456¦123", entries[0].Addenda18[0].ForeignCorrespondentBankIDNumber)
	})

	t.Run("MTE File", func(t *testing.T) {
		bh := mockBatchMTEHeader()
		bh.CompanyName = "Merchant | ATM"

		b := NewBatchMTE(bh)
		ed := mockMTEEntryDetail()
		ed.IndividualName = "My {Store}"
		ed.Addenda02.ReferenceInformationOne = `RF1¦RF2`
		b.AddEntry(ed)
		require.NoError(t, b.Create())

		file := NewFile()
		file.SetHeader(mockFileHeader())
		file.AddBatch(b)
		require.NoError(t, file.Create())

		// Cycle file through write/read
		var buf bytes.Buffer
		err := NewWriter(&buf).Write(file)
		require.NoError(t, err)

		parsed, err := NewReader(&buf).Read()
		require.NoError(t, err)

		require.Equal(t, "Federal Reserve Bank", parsed.Header.ImmediateDestinationName)
		require.Equal(t, "My Bank Name", parsed.Header.ImmediateOriginName)

		// Verify file details
		b1 := parsed.Batches[0]
		require.Equal(t, `Merchant | ATM`, b1.GetHeader().CompanyName)

		entries := b1.GetEntries()
		require.Equal(t, `My {Store}            `, entries[0].IndividualName)
		require.Equal(t, `RF1¦RF2`, entries[0].Addenda02.ReferenceInformationOne)
	})

	t.Run("detect", func(t *testing.T) {
		bs, err := os.ReadFile(filepath.Join("test", "testdata", "nonascii-utf8.ach"))
		require.NoError(t, err)
		_, name, _ := charset.DetermineEncoding(bs, "plain/text")
		require.Equal(t, "utf-8", name)

		bs, err = os.ReadFile(filepath.Join("test", "testdata", "nonascii.ach"))
		require.NoError(t, err)
		_, name, _ = charset.DetermineEncoding(bs, "plain/text")
		require.Equal(t, "windows-1252", name)
	})

	t.Run("parse nonaschii-utf8", func(t *testing.T) {
		file, err := ReadFile(filepath.Join("test", "testdata", "nonascii-utf8.ach"))
		// require.NoError(t, err)
		if err != nil {
			t.Log(err)
		}

		require.Len(t, file.Batches, 1)
		bh := file.Batches[0].GetHeader()
		require.Equal(t, "REG.SALARY", bh.CompanyEntryDescription)

		entries := file.Batches[0].GetEntries()
		require.Len(t, entries, 1)
		require.Equal(t, "0012Receiver¦Acc Name ", entries[0].IndividualName)

		require.Len(t, entries[0].Addenda05, 12)
		require.Contains(t, entries[0].Addenda05[0].PaymentRelatedInformation, "¦ZZ¦PAYEXPENSEPAY")
	})

	t.Run("parse windows-1252", func(t *testing.T) {
		file, err := ReadFile(filepath.Join("test", "testdata", "nonascii.ach"))
		// require.NoError(t, err)
		if err != nil {
			t.Log(err)
		}

		require.Len(t, file.Batches, 1)
		bh := file.Batches[0].GetHeader()
		require.Equal(t, "REG.SALARY", bh.CompanyEntryDescription)

		entries := file.Batches[0].GetEntries()
		require.Len(t, entries, 1)
		require.Equal(t, "0012Receiver¦Acc Name ", entries[0].IndividualName)

		require.Len(t, entries[0].Addenda05, 12)
		require.Contains(t, entries[0].Addenda05[0].PaymentRelatedInformation, "¦ZZ¦PAYEXPENSEPAY")
	})
}
