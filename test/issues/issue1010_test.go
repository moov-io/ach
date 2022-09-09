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
)

func TestIssue1010(t *testing.T) {
	// Batch Header
	batchHeader := ach.NewBatchHeader()
	batchHeader.ServiceClassCode = 200
	batchHeader.CompanyName = "COMPANY"
	batchHeader.CompanyIdentification = "IDENTIFICATION"
	batchHeader.StandardEntryClassCode = ach.WEB
	batchHeader.CompanyEntryDescription = "DESCRIPTION"
	batchHeader.CompanyDescriptiveDate = time.Now().Format("060102")
	batchHeader.EffectiveEntryDate = time.Now().Add(time.Hour * 24).Format("060102")
	batchHeader.ODFIIdentification = "1"

	// Addenda99
	addenda99 := ach.NewAddenda99()
	addenda99.ReturnCode = "R07"
	addenda99.OriginalTrace = "99912340000015"
	addenda99.AddendaInformation = "Authorization Revoked"
	addenda99.OriginalDFI = "9101298"

	// Entry
	entry := ach.NewEntryDetail()
	entry.TransactionCode = 22
	entry.RDFIIdentification = "1"
	entry.CheckDigit = "3"
	entry.DFIAccountNumber = "DUMMY"
	entry.Amount = 1000
	entry.IndividualName = "DOE, JOHN"
	entry.Category = ach.CategoryReturn
	entry.Addenda99 = addenda99
	entry.AddendaRecordIndicator = 1

	// Batches
	batch1 := ach.NewBatchWEB(batchHeader)
	batch1.AddEntry(entry)

	// Prints nil
	if err := batch1.Create(); err != nil {
		t.Fatal(err)
	}

	batch2 := ach.NewBatchWEB(batchHeader)
	batch2.AddEntry(entry)
	batch2.WithOffset(&ach.Offset{
		RoutingNumber: "100000007",
		AccountNumber: "DUMMY",
		AccountType:   "checking",
		Description:   "S",
	})

	// Prints error
	if err := batch2.Create(); err != nil {
		t.Fatal(err)
	}
}
