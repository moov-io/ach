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
	"testing"

	"github.com/moov-io/ach"
	"github.com/stretchr/testify/require"
)

// TestIssue1052 tests that prenote files don't get ADV batch control with ServiceClassCode 280
// 
// Issue #1052: Users were incorrectly creating ADV batches when they intended to create
// prenote files. This was due to misleading documentation that labeled ADV as "Prenote Debit".
// 
// The fix corrects the documentation to clarify that:
// - ADV (Automated Accounting Advice) is for accounting advice entries from ACH operators
// - Prenotes are regular entries with transaction codes like 23, 28, 33, 38 used with SEC codes like PPD, CCD, CTX
// - Prenote batches should NOT use ServiceClassCode 280 (AutomatedAccountingAdvices)
func TestIssue1052(t *testing.T) {
	// Create a file with prenote entries
	fh := ach.NewFileHeader()
	fh.ImmediateDestination = "121042882"
	fh.ImmediateOrigin = "231380104"
	fh.FileCreationDate = "190816"
	fh.ImmediateDestinationName = "Federal Reserve Bank"
	fh.ImmediateOriginName = "My Bank Name"

	file := ach.NewFile()
	file.SetHeader(fh)

	// Create a batch header for PPD (not ADV)
	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.CreditsOnly
	bh.CompanyName = "Company Name, Inc"
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.PPD // This should be PPD, not ADV
	bh.CompanyEntryDescription = "PRENOTE"
	bh.ODFIIdentification = "121042882"
	bh.OriginatorStatusCode = 1

	// Create a batch
	batch := ach.NewBatchPPD(bh)

	// Add a prenote entry
	entry := ach.NewEntryDetail()
	entry.TransactionCode = ach.CheckingPrenoteCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 0 // Prenotes must have zero amount
	entry.IdentificationNumber = "45689033"
	entry.IndividualName = "John Doe"
	entry.SetTraceNumber(bh.ODFIIdentification, 1)
	entry.Category = ach.CategoryForward

	batch.AddEntry(entry)

	// Add batch to file
	file.AddBatch(batch)

	// Create the file
	err := file.Create()
	require.NoError(t, err)

	// Verify that the batch control does NOT have ServiceClassCode 280
	control := batch.GetControl()
	require.NotNil(t, control)

	// The batch control should have the same ServiceClassCode as the header
	// It should NOT be 280 (AutomatedAccountingAdvices)
	require.NotEqual(t, ach.AutomatedAccountingAdvices, control.ServiceClassCode,
		"Batch control ServiceClassCode should not be %d (AutomatedAccountingAdvices)",
		ach.AutomatedAccountingAdvices)

	// Verify it's not an ADV batch
	require.False(t, batch.IsADV(), "Batch should not be ADV type")

	// Verify there's no ADV control
	require.Nil(t, batch.GetADVControl(), "Batch should not have ADV control")
}
