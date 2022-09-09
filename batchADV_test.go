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

package ach

import (
	"testing"
	"time"

	"github.com/moov-io/base"
)

// mockBatchADVHeader creates a ADV batch header
func mockBatchADVHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = AutomatedAccountingAdvices
	bh.StandardEntryClassCode = ADV
	bh.CompanyName = "Your Company, inc"
	bh.CompanyIdentification = "121042882"
	bh.CompanyEntryDescription = "Accounting"
	bh.ODFIIdentification = "12104288"
	bh.OriginatorStatusCode = 0
	bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1).Format("060102") // YYMMDD
	return bh
}

// mockBatchADV creates a ADV batch
func mockBatchADV() *BatchADV {
	mockBatch := NewBatchADV(mockBatchADVHeader())
	mockBatch.AddADVEntry(mockADVEntryDetail())
	if err := mockBatch.Create(); err != nil {
		panic(err)
	}
	return mockBatch
}

// testBatchADVHeader creates a ADV batch header
func testBatchADVHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchADVHeader())
	_, ok := batch.(*BatchADV)
	if !ok {
		t.Error("Expecting BatchADV")
	}
}

// TestBatchADVHeader tests creating a ADV batch header
func TestBatchADVHeader(t *testing.T) {
	testBatchADVHeader(t)
}

// BenchmarkBatchADVHeader benchmark creating a ADV batch header
func BenchmarkBatchADVHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchADVHeader(b)
	}
}

// TestBatchADVAddendum99 validates Addenda99 returns an error
func TestBatchADVAddendum99(t *testing.T) {
	mockBatch := NewBatchADV(mockBatchADVHeader())
	mockBatch.AddADVEntry(mockADVEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockAddenda99.TypeCode = "05"
	mockBatch.GetADVEntries()[0].Category = CategoryReturn
	mockBatch.GetADVEntries()[0].Addenda99 = mockAddenda99
	err := mockBatch.Create()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// testBatchADVSEC validates that the standard entry class code is ADV for batchADV
func testBatchADVSEC(t testing.TB) {
	mockBatch := mockBatchADV()
	mockBatch.Header.StandardEntryClassCode = RCK
	err := mockBatch.Validate()
	if !base.Match(err, ErrBatchSECType) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchADVSEC tests validating that the standard entry class code is ADV for batchADV
func TestBatchADVSEC(t *testing.T) {
	testBatchADVSEC(t)
}

// BenchmarkBatchADVSEC benchmarks validating that the standard entry class code is ADV for batch ADV
func BenchmarkBatchADVSEC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchADVSEC(b)
	}
}

// testBatchADVServiceClassCode validates ServiceClassCode
func testBatchADVServiceClassCode(t testing.TB) {
	mockBatch := mockBatchADV()
	// Batch Header information is required to Create a batch.
	mockBatch.GetHeader().ServiceClassCode = 220
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchServiceClassCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchADVServiceClassCode tests validating ServiceClassCode
func TestBatchADVServiceClassCode(t *testing.T) {
	testBatchADVServiceClassCode(t)
}

// BenchmarkBatchADVServiceClassCode benchmarks validating ServiceClassCode
func BenchmarkBatchADVServiceClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchADVServiceClassCode(b)
	}
}

// TestBatchADVAddendum99Category validates Addenda99 returns an error
func TestBatchADVAddendum99Category(t *testing.T) {
	mockBatch := NewBatchADV(mockBatchADVHeader())
	mockBatch.AddADVEntry(mockADVEntryDetail())
	mockAddenda99 := mockAddenda99()
	mockBatch.GetADVEntries()[0].Category = CategoryForward
	mockBatch.GetADVEntries()[0].Addenda99 = mockAddenda99
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAddendaCategory) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchADVInvalidTransactionCode validates TransactionCode
func TestBatchADVInvalidTransactionCode(t *testing.T) {
	mockBatch := mockBatchADV()
	// Batch Header information is required to Create a batch.
	mockBatch.GetADVEntries()[0].TransactionCode = CheckingCredit
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchTransactionCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestADVMaximumEntries validates maximum entries for an ADV ACH file
func TestADVMaximumEntries(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.AddendaRecordIndicator = 0
	batch := NewBatchADV(mockBatchADVHeader())
	batch.SetHeader(mockBatchADVHeader())

	for i := 0; i < 10000; i++ {
		batch.AddADVEntry(entry)
	}
	err := batch.Create()
	if !base.Match(err, ErrBatchADVCount) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchADVOriginatorStatusCode validates the originator status code
func TestBatchADVOriginatorStatusCode(t *testing.T) {
	mockBatch := mockBatchADV()
	mockBatch.Header.OriginatorStatusCode = 1
	err := mockBatch.Create()
	if !base.Match(err, ErrOrigStatusCode) {
		t.Errorf("%T: %s", err, err)
	}
}
