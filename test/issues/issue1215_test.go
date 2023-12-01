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

func TestIssue1215(t *testing.T) {
	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.MixedDebitsAndCredits
	bh.CompanyName = "WhlBodyFit"
	bh.CompanyIdentification = "MOOVTESTER"
	bh.StandardEntryClassCode = ach.CCD
	bh.CompanyEntryDescription = "PRENOTE"
	bh.EffectiveEntryDate = "230411"
	bh.SettlementDate = "102"
	bh.ODFIIdentification = "05310055"

	entry := ach.NewEntryDetail()
	entry.TransactionCode = ach.CheckingReturnNOCDebit
	entry.RDFIIdentification = "12345678"
	entry.CheckDigit = "0"
	entry.DFIAccountNumber = "2345532322345"
	entry.Amount = 0
	entry.IdentificationNumber = "               "
	entry.IndividualName = "Whole Body Fitness    "
	entry.AddendaRecordIndicator = 1
	entry.TraceNumber = "053100552202922"
	entry.Category = ach.CategoryReturn

	addenda99 := ach.NewAddenda99()
	addenda99.ReturnCode = "R03"
	addenda99.OriginalTrace = "273976363256887"
	addenda99.DateOfDeath = ""
	addenda99.OriginalDFI = "12100035"
	addenda99.AddendaInformation = ""
	addenda99.TraceNumber = "053100552202922"
	entry.Addenda99 = addenda99

	batch := ach.NewBatchCCD(bh)
	batch.AddEntry(entry)

	require.NoError(t, batch.Create())
	require.NoError(t, batch.Validate())
}

func TestIssue1332(t *testing.T) {
	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.MixedDebitsAndCredits
	bh.CompanyName = "WhlBodyFit"
	bh.CompanyIdentification = "MOOVTESTER"
	bh.StandardEntryClassCode = ach.CCD
	bh.CompanyEntryDescription = "PMTCHECK"
	bh.EffectiveEntryDate = "230411"
	bh.SettlementDate = "102"
	bh.ODFIIdentification = "05310055"

	entry := ach.NewEntryDetail()
	entry.TransactionCode = ach.CheckingReturnNOCDebit
	entry.RDFIIdentification = "12345678"
	entry.CheckDigit = "0"
	entry.DFIAccountNumber = "2345532322345"
	entry.Amount = 0
	entry.IdentificationNumber = "               "
	entry.IndividualName = "Whole Body Fitness    "
	entry.AddendaRecordIndicator = 1
	entry.TraceNumber = "053100552202922"
	entry.Category = ach.CategoryReturn

	addenda99 := ach.NewAddenda99()
	addenda99.ReturnCode = "R03"
	addenda99.OriginalTrace = "273976363256887"
	addenda99.DateOfDeath = ""
	addenda99.OriginalDFI = "12100035"
	addenda99.AddendaInformation = ""
	addenda99.TraceNumber = "053100552202922"
	entry.Addenda99 = addenda99

	batch := ach.NewBatchCCD(bh)
	batch.AddEntry(entry)
	batch.SetValidation(&ach.ValidateOpts{
		// Needs to be set because we lose all information the original entry was a prenote.
		// Typically the CompanyEntryDescription is set to PRENOTE, which is carried through
		// to the return.
		AllowInvalidAmounts: true,
	})

	require.NoError(t, batch.Create())
	require.NoError(t, batch.Validate())
}
