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

package issues

import (
	"bytes"
	"testing"
	"time"

	"github.com/moov-io/ach"

	"github.com/stretchr/testify/require"
)

func TestIssue1385(t *testing.T) {
	bh := ach.NewIATBatchHeader()
	bh.ServiceClassCode = ach.CreditsOnly
	bh.IATIndicator = ach.IATCOR
	bh.ForeignExchangeIndicator = "FF"
	bh.ForeignExchangeReferenceIndicator = 3
	bh.ISODestinationCountryCode = "US"
	bh.OriginatorIdentification = "123456789"
	bh.StandardEntryClassCode = ach.COR
	bh.CompanyEntryDescription = "TRADEPAYMT"
	bh.ISOOriginatingCurrencyCode = "CAD"
	bh.ISODestinationCurrencyCode = "USD"
	bh.ODFIIdentification = "91012980"

	batch := ach.NewIATBatch(bh)

	entry := ach.NewIATEntryDetail()
	entry.TransactionCode = ach.CheckingReturnNOCCredit
	entry.SetRDFI("121042882")
	entry.AddendaRecords = 1
	entry.DFIAccountNumber = "123456789"
	entry.Amount = 0
	entry.SetTraceNumber(bh.ODFIIdentification, 1)
	// entry.Category = ach.CategoryNOC // not set, but an IATBatch can detect COR/NOC
	require.Nil(t, entry.Addenda10)

	entry.Addenda98 = ach.NewAddenda98()
	entry.Addenda98.ChangeCode = "C01"
	entry.Addenda98.OriginalTrace = "12345"
	entry.Addenda98.OriginalDFI = "9101298"
	entry.Addenda98.CorrectedData = "1918171614"
	entry.Addenda98.TraceNumber = "91012980000088"

	batch.AddEntry(entry)

	// Before the fix this returns
	//  "Addenda10 <nil> is a mandatory field and has a default value"
	require.NoError(t, batch.Create())

	file := ach.NewFile()
	file.Header = ach.NewFileHeader()
	file.Header.ImmediateDestination = "121042882"
	file.Header.ImmediateOrigin = "91012980"
	file.Header.FileCreationDate = time.Now().AddDate(0, 0, 1).Format("060102") // YYMMDD
	file.Header.ImmediateDestinationName = "Federal Reserve Bank"
	file.Header.ImmediateOriginName = "My Bank Name"

	file.AddIATBatch(batch)
	require.NoError(t, file.Create())
	require.NoError(t, file.Validate())

	// Write so we can re-read
	var buf bytes.Buffer
	err := ach.NewWriter(&buf).Write(file)
	require.NoError(t, err)

	_, err = ach.NewReader(&buf).Read()
	require.NoError(t, err)
}
