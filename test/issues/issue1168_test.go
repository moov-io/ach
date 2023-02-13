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
	"time"

	"github.com/moov-io/ach"

	"github.com/stretchr/testify/require"
)

func TestIssue1168(t *testing.T) {
	// Batch Header
	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = 200
	bh.CompanyName = "COMPANY"
	bh.CompanyIdentification = "IDENTIFICATION"
	bh.StandardEntryClassCode = ach.WEB
	bh.CompanyEntryDescription = "DESCRIPTION"
	bh.CompanyDescriptiveDate = time.Now().Format("060102")
	bh.EffectiveEntryDate = time.Now().Add(time.Hour * 24).Format("060102")
	bh.ODFIIdentification = "1"

	ed := &ach.EntryDetail{
		TransactionCode:    ach.SavingsCredit,
		DFIAccountNumber:   "54321",
		RDFIIdentification: "12345678",
		CheckDigit:         "0",
		Amount:             0,
		IndividualName:     "Jane Doe",
		TraceNumber:        "445566778899",
	}

	b1 := ach.NewBatchWEB(bh)
	b1.AddEntry(ed)

	// Require that Batches cannot have zero amounts
	require.ErrorContains(t, b1.Create(), ach.ErrBatchAmountZero.Error())

	// Use a pre-note transaction code and verify success
	ed.TransactionCode = ach.SavingsPrenoteCredit
	require.NoError(t, b1.Create())
	require.NoError(t, b1.Validate())

	// Set a non-zero amount and catch error
	ed.Amount = 1223
	require.ErrorContains(t, b1.Create(), ach.ErrBatchAmountNonZero.Error())
}
