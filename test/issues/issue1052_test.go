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
// KIND, either express or implied. See the License for the
// specific language governing permissions and limitations
// under the License.

package issues

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/ach"
	"github.com/stretchr/testify/require"
)

// Prenote JSON omits advBatchControl. Decoding used to install one with service class 280.
const prenoteFileJSON = `{
  "fileHeader": {
    "immediateDestination": "231380104",
    "immediateOrigin": "121042882",
    "fileCreationDate": "190816",
    "fileIDModifier": "A",
    "immediateDestinationName": "Federal Reserve Bank",
    "immediateOriginName": "My Bank Name"
  },
  "batches": [
    {
      "batchHeader": {
        "serviceClassCode": 220,
        "companyName": "Company Name, Inc",
        "companyIdentification": "121042882",
        "standardEntryClassCode": "PPD",
        "companyEntryDescription": "PRENOTE",
        "effectiveEntryDate": "190817",
        "ODFIIdentification": "12104288",
        "originatorStatusCode": 1,
        "batchNumber": 1
      },
      "entryDetails": [
        {
          "transactionCode": 23,
          "RDFIIdentification": "23138010",
          "checkDigit": "4",
          "DFIAccountNumber": "744-5678-99",
          "amount": 0,
          "identificationNumber": "45689033",
          "individualName": "John Doe",
          "addendaRecordIndicator": 0,
          "traceNumber": "121042880000001",
          "category": "Forward"
        }
      ]
    }
  ]
}`

func TestIssue1052(t *testing.T) {
	t.Run("json prenote omits generated ADV control", func(t *testing.T) {
		file, err := ach.FileFromJSON([]byte(prenoteFileJSON))
		require.NoError(t, err)
		require.Len(t, file.Batches, 1)

		batch := file.Batches[0]
		require.Equal(t, ach.PPD, batch.GetHeader().StandardEntryClassCode)
		require.Equal(t, ach.CreditsOnly, batch.GetHeader().ServiceClassCode)
		require.Equal(t, ach.CreditsOnly, batch.GetControl().ServiceClassCode)
		require.Nil(t, batch.GetADVControl())

		bs, err := json.Marshal(file)
		require.NoError(t, err)
		require.NotContains(t, string(bs), `"advBatchControl"`)
	})

	t.Run("json prenote drops a submitted ADV control", func(t *testing.T) {
		// A caller echoing an older response may send the generated control back.
		raw := strings.Replace(prenoteFileJSON, `"entryDetails"`, `"advBatchControl": {"serviceClassCode": 280}, "entryDetails"`, 1)
		require.Contains(t, raw, `"advBatchControl"`)

		file, err := ach.FileFromJSON([]byte(raw))
		require.NoError(t, err)
		require.Len(t, file.Batches, 1)
		require.Nil(t, file.Batches[0].GetADVControl())
		require.Equal(t, ach.CreditsOnly, file.Batches[0].GetControl().ServiceClassCode)

		bs, err := json.Marshal(file)
		require.NoError(t, err)
		require.NotContains(t, string(bs), `"advBatchControl"`)
	})

	t.Run("decoded prenote batch does not invent an ADV control", func(t *testing.T) {
		var raw struct {
			Batches []json.RawMessage `json:"batches"`
		}
		require.NoError(t, json.Unmarshal([]byte(prenoteFileJSON), &raw))
		require.Len(t, raw.Batches, 1)

		batch := &ach.Batch{}
		require.NoError(t, batch.UnmarshalJSON(raw.Batches[0]))
		require.Nil(t, batch.GetADVControl())
		require.Equal(t, ach.PPD, batch.GetHeader().StandardEntryClassCode)
		require.Equal(t, ach.CreditsOnly, batch.GetHeader().ServiceClassCode)
	})

	t.Run("ADV batch keeps service class 280", func(t *testing.T) {
		bs, err := os.ReadFile(filepath.Join("..", "testdata", "adv-valid.json"))
		require.NoError(t, err)

		file, err := ach.FileFromJSON(bs)
		require.NoError(t, err)
		require.True(t, file.IsADV())
		require.Len(t, file.Batches, 1)

		adv := file.Batches[0].GetADVControl()
		require.NotNil(t, adv)
		require.Equal(t, ach.AutomatedAccountingAdvices, adv.ServiceClassCode)
		require.Equal(t, ach.AutomatedAccountingAdvices, file.Batches[0].GetHeader().ServiceClassCode)
	})
}
