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

func TestIssue1861(t *testing.T) {
	t.Run("zero dollar CCD remittance with addenda parses", func(t *testing.T) {
		bh := ach.NewBatchHeader()
		bh.ServiceClassCode = ach.CreditsOnly
		bh.StandardEntryClassCode = ach.CCD
		bh.CompanyName = "Your Company, inc"
		bh.CompanyIdentification = "121042882"
		bh.CompanyEntryDescription = "REMIT"
		bh.ODFIIdentification = "121042882"

		entry := ach.NewEntryDetail()
		entry.TransactionCode = ach.CheckingZeroDollarRemittanceCredit
		entry.SetRDFI("231380104")
		entry.DFIAccountNumber = "744-5678-99"
		entry.Amount = 0
		entry.IdentificationNumber = "location #23"
		entry.SetReceivingCompany("Best Co. #23")
		entry.SetTraceNumber(bh.ODFIIdentification, 1)
		entry.AddendaRecordIndicator = 1

		addenda := ach.NewAddenda05()
		addenda.PaymentRelatedInformation = `NTE*REMITTANCE INFORMATION\`
		entry.AddAddenda05(addenda)

		batch := ach.NewBatchCCD(bh)
		batch.AddEntry(entry)
		require.NoError(t, batch.Create())
		require.NoError(t, batch.Validate())
	})

	t.Run("non-zero amount on remittance code is rejected", func(t *testing.T) {
		bh := ach.NewBatchHeader()
		bh.ServiceClassCode = ach.CreditsOnly
		bh.StandardEntryClassCode = ach.CCD
		bh.CompanyName = "Your Company, inc"
		bh.CompanyIdentification = "121042882"
		bh.CompanyEntryDescription = "REMIT"
		bh.ODFIIdentification = "121042882"

		entry := ach.NewEntryDetail()
		entry.TransactionCode = ach.CheckingZeroDollarRemittanceCredit
		entry.SetRDFI("231380104")
		entry.DFIAccountNumber = "744-5678-99"
		entry.Amount = 100
		entry.SetReceivingCompany("Best Co. #23")
		entry.SetTraceNumber(bh.ODFIIdentification, 1)
		entry.AddendaRecordIndicator = 1
		entry.AddAddenda05(ach.NewAddenda05())

		batch := ach.NewBatchCCD(bh)
		batch.AddEntry(entry)
		require.ErrorContains(t, batch.Create(), ach.ErrBatchAmountNonZero.Error())
	})
}
