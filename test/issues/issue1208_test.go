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
	"testing"

	"github.com/moov-io/ach"

	"github.com/stretchr/testify/require"
)

func TestIssue1208(t *testing.T) {
	batchHeader := ach.NewBatchHeader()
	batchHeader.ServiceClassCode = 200
	batchHeader.CompanyDiscretionaryData = "1"
	batchHeader.CompanyName = "2"
	batchHeader.CompanyIdentification = "3"
	batchHeader.StandardEntryClassCode = "CTX"
	batchHeader.CompanyEntryDescription = "5"
	batchHeader.CompanyDescriptiveDate = "6"
	batchHeader.EffectiveEntryDate = "7"
	batchHeader.ODFIIdentification = "8"

	batch, err := ach.NewBatch(batchHeader)
	require.NoError(t, err)

	batch.SetValidation(&ach.ValidateOpts{
		CustomTraceNumbers: true,
	})

	entryDetail := ach.NewEntryDetail()

	// Verify some helpers for CTX
	require.Equal(t, "", entryDetail.CATXAddendaRecordsField())
	require.Equal(t, "", entryDetail.CATXReceivingCompanyField())
	require.Equal(t, "", entryDetail.IndividualName)

	// Set CTX examples from Github issue
	entryDetail.TransactionCode = 26
	entryDetail.RDFIIdentification = "87654321"
	entryDetail.CheckDigit = "2"
	entryDetail.DFIAccountNumber = "123456789"
	entryDetail.Amount = 123
	entryDetail.IdentificationNumber = "10"
	// CTX uses characters 1-4 of IndividualName to store AddendaRecords count
	entryDetail.SetCATXAddendaRecords(0)
	entryDetail.SetCATXReceivingCompany("Whatevs")
	entryDetail.DiscretionaryData = "DD"
	entryDetail.TraceNumber = "500"
	entryDetail.Category = ach.CategoryForward
	require.Equal(t, "0000", entryDetail.CATXAddendaRecordsField())
	require.Equal(t, "Whatevs           ", entryDetail.CATXReceivingCompanyField())
	require.Equal(t, "0000Whatevs           ", entryDetail.IndividualName)

	batch.AddEntry(entryDetail)

	require.NoError(t, batch.Create())
	require.NoError(t, batch.Validate())

	addenda99 := ach.NewAddenda99()
	addenda99.ReturnCode = "R05"
	addenda99.OriginalTrace = "500"
	addenda99.TraceNumber = "600"
	addenda99.OriginalDFI = "12345678"
	entryDetail.Addenda99 = addenda99
	entryDetail.SetCATXAddendaRecords(1)
	entryDetail.Category = ach.CategoryReturn
	require.Equal(t, "0001", entryDetail.CATXAddendaRecordsField())
	require.Equal(t, "Whatevs           ", entryDetail.CATXReceivingCompanyField())
	require.Equal(t, "0001Whatevs           ", entryDetail.IndividualName)

	require.NoError(t, batch.Create())
	require.NoError(t, batch.Validate())

	addenda05 := ach.NewAddenda05()
	addenda05.SequenceNumber = 1
	addenda05.PaymentRelatedInformation = "ASC X12 Data"
	addenda05.EntryDetailSequenceNumber = 0000001
	entryDetail.AddAddenda05(addenda05)
	entryDetail.SetCATXAddendaRecords(2)
	require.Equal(t, "0002", entryDetail.CATXAddendaRecordsField())
	require.Equal(t, "Whatevs           ", entryDetail.CATXReceivingCompanyField())
	require.Equal(t, "0002Whatevs           ", entryDetail.IndividualName)

	require.NoError(t, batch.Create())
	require.NoError(t, batch.Validate())
}
