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

// TestIssue1647_AddendaRecordIndicatorValidation tests that batch validation
// catches when AddendaRecordIndicator doesn't match the actual number of addendas
func TestIssue1647_AddendaRecordIndicatorValidation(t *testing.T) {
	// Test case 1: AddendaRecordIndicator=1 but no addenda present
	t.Run("AddendaRecordIndicator1WithNoAddenda", func(t *testing.T) {
		bh := ach.NewBatchHeader()
		bh.StandardEntryClassCode = ach.PPD
		bh.CompanyName = "Test Company"
		bh.CompanyIdentification = "123456789"
		bh.ODFIIdentification = "12104288"
		bh.BatchNumber = 1
		bh.ServiceClassCode = ach.MixedDebitsAndCredits
		bh.CompanyEntryDescription = "PAYMENT"

		entry := ach.NewEntryDetail()
		entry.TransactionCode = ach.CheckingCredit
		entry.SetRDFI("121042882")
		entry.DFIAccountNumber = "123456789"
		entry.Amount = 10000
		entry.IndividualName = "John Doe"
		entry.SetTraceNumber(bh.ODFIIdentification, 1)
		entry.Category = ach.CategoryForward
		entry.AddendaRecordIndicator = 1 // Set to 1 but no addenda

		batch := ach.NewBatchPPD(bh)
		batch.AddEntry(entry)

		err := batch.Create()
		require.ErrorContains(t, err, ach.ErrBatchAddendaRequired.Error())
	})

	// Test case 2: AddendaRecordIndicator=0 but addenda present
	t.Run("AddendaRecordIndicator0WithAddenda", func(t *testing.T) {
		bh := ach.NewBatchHeader()
		bh.StandardEntryClassCode = ach.PPD
		bh.CompanyName = "Test Company"
		bh.CompanyIdentification = "123456789"
		bh.ODFIIdentification = "12104288"
		bh.BatchNumber = 1
		bh.ServiceClassCode = ach.MixedDebitsAndCredits
		bh.CompanyEntryDescription = "PAYMENT"

		entry := ach.NewEntryDetail()
		entry.TransactionCode = ach.CheckingCredit
		entry.SetRDFI("121042882")
		entry.DFIAccountNumber = "123456789"
		entry.Amount = 10000
		entry.IndividualName = "John Doe"
		entry.SetTraceNumber(bh.ODFIIdentification, 1)
		entry.Category = ach.CategoryForward
		entry.AddendaRecordIndicator = 0 // Set to 0 but addenda present

		addenda := ach.NewAddenda05()
		addenda.PaymentRelatedInformation = "Test payment info"
		entry.AddAddenda05(addenda)

		batch := ach.NewBatchPPD(bh)
		batch.AddEntry(entry)

		err := batch.Create()
		require.ErrorContains(t, err, ach.ErrBatchAddendaIndicator.Error())
	})

	// Test case 3: Valid case - AddendaRecordIndicator=1 with addenda present
	t.Run("ValidAddendaRecordIndicator1WithAddenda", func(t *testing.T) {
		bh := ach.NewBatchHeader()
		bh.StandardEntryClassCode = ach.PPD
		bh.CompanyName = "Test Company"
		bh.CompanyIdentification = "123456789"
		bh.ODFIIdentification = "12104288"
		bh.BatchNumber = 1
		bh.ServiceClassCode = ach.MixedDebitsAndCredits
		bh.CompanyEntryDescription = "PAYMENT"

		entry := ach.NewEntryDetail()
		entry.TransactionCode = ach.CheckingCredit
		entry.SetRDFI("121042882")
		entry.DFIAccountNumber = "123456789"
		entry.Amount = 10000
		entry.IndividualName = "John Doe"
		entry.SetTraceNumber(bh.ODFIIdentification, 1)
		entry.Category = ach.CategoryForward
		entry.AddendaRecordIndicator = 1 // Set to 1 with addenda present

		addenda := ach.NewAddenda05()
		addenda.PaymentRelatedInformation = "Test payment info"
		entry.AddAddenda05(addenda)

		batch := ach.NewBatchPPD(bh)
		batch.AddEntry(entry)

		err := batch.Create()
		require.NoError(t, err)
	})

	// Test case 4: Valid case - AddendaRecordIndicator=0 with no addenda
	t.Run("ValidAddendaRecordIndicator0WithNoAddenda", func(t *testing.T) {
		bh := ach.NewBatchHeader()
		bh.StandardEntryClassCode = ach.PPD
		bh.CompanyName = "Test Company"
		bh.CompanyIdentification = "123456789"
		bh.ODFIIdentification = "12104288"
		bh.BatchNumber = 1
		bh.ServiceClassCode = ach.MixedDebitsAndCredits
		bh.CompanyEntryDescription = "PAYMENT"

		entry := ach.NewEntryDetail()
		entry.TransactionCode = ach.CheckingCredit
		entry.SetRDFI("121042882")
		entry.DFIAccountNumber = "123456789"
		entry.Amount = 10000
		entry.IndividualName = "John Doe"
		entry.SetTraceNumber(bh.ODFIIdentification, 1)
		entry.Category = ach.CategoryForward
		entry.AddendaRecordIndicator = 0 // Set to 0 with no addenda

		batch := ach.NewBatchPPD(bh)
		batch.AddEntry(entry)

		err := batch.Create()
		require.NoError(t, err)
	})
}
