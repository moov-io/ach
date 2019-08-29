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

package server

import (
	"time"

	"github.com/moov-io/ach"
)

func mockServiceInMemory() Service {
	repository := NewRepositoryInMemory(testTTLDuration, nil)
	repository.StoreFile(&ach.File{ID: "98765"})
	repository.StoreBatch("98765", mockBatchWEB())
	return NewService(repository)
}

func mockFileHeader() *ach.FileHeader {
	fh := ach.NewFileHeader()
	fh.ID = "12345"
	fh.ImmediateDestination = "231380104"
	fh.ImmediateOrigin = "121042882"
	fh.FileCreationDate = time.Now().Format("060102")
	fh.ImmediateDestinationName = "Federal Reserve Bank"
	fh.ImmediateOriginName = "My Bank Name"
	return &fh
}

func mockBatchHeaderWeb() *ach.BatchHeader {
	bh := ach.NewBatchHeader()
	bh.ID = "54321"
	bh.ServiceClassCode = ach.CreditsOnly
	bh.StandardEntryClassCode = ach.WEB
	bh.CompanyName = "Your Company, inc"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "Online Order"
	bh.ODFIIdentification = "12104288"
	return bh
}

// mockWEBEntryDetail creates a WEB entry detail
func mockWEBEntryDetail() *ach.EntryDetail {
	entry := ach.NewEntryDetail()
	entry.ID = "98765"
	entry.TransactionCode = ach.CheckingCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "123456789"
	entry.Amount = 100000000
	entry.IndividualName = "Wade Arnold"
	entry.SetTraceNumber(mockBatchHeaderWeb().ODFIIdentification, 1)
	entry.SetPaymentType("S")
	entry.AddAddenda05(mockAddenda05())
	return entry
}

// mockBatchWEB creates a WEB batch
func mockBatchWEB() *ach.BatchWEB {
	mockBatch := ach.NewBatchWEB(mockBatchHeaderWeb())
	mockBatch.SetID(mockBatch.Header.ID)
	mockBatch.AddEntry(mockWEBEntryDetail())
	mockBatch.Entries[0].AddendaRecordIndicator = 1
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

func mockAddenda05() *ach.Addenda05 {
	addenda05 := ach.NewAddenda05()
	addenda05.ID = "56789"
	addenda05.SequenceNumber = 1
	addenda05.EntryDetailSequenceNumber = 0000001
	return addenda05
}
