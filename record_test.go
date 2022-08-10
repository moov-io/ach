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
	"bytes"
	"testing"
)

// testFileRecord validates a file record
func testFileRecord(t testing.TB) {
	f := NewFile()
	f.SetHeader(mockFileHeader())
	if err := f.Header.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	if f.Header.ImmediateOriginName != "My Bank Name" {
		t.Errorf("FileParam value was not copied to file.Header")
	}
}

// TestFileRecord tests validating a file record
func TestFileRecord(t *testing.T) {
	testFileRecord(t)
}

// BenchmarkFileRecord benchmarks validating a file record
func BenchmarkFileRecord(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileRecord(b)
	}
}

// testBatchRecord validates a batch record
func testBatchRecord(t testing.TB) {
	companyName := "ACME Corporation"
	batch, _ := NewBatch(mockBatchPPDHeader())

	bh := batch.GetHeader()
	if err := bh.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if bh.CompanyName != companyName {
		t.Errorf("BatchParam value was not copied to batch.Header.CompanyName")
	}
}

// TestBatchRecord tests validating a batch record
func TestBatchRecord(t *testing.T) {
	testBatchRecord(t)
}

// BenchmarkBatchRecord benchmarks validating a batch record
func BenchmarkBatchRecord(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchRecord(b)
	}
}

// testEntryDetail validates an entry detail record
func testEntryDetail(t testing.TB) {
	entry := mockEntryDetail()
	//override mockEntryDetail
	entry.TransactionCode = CheckingDebit

	if err := entry.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestEntryDetail tests validating an entry detail record
func TestEntryDetail(t *testing.T) {
	testEntryDetail(t)
}

// BenchmarkEntryDetail benchmarks validating an entry detail record
func BenchmarkEntryDetail(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEntryDetail(b)
	}
}

// testEntryDetailPaymentType validates an entry detail record payment type
func testEntryDetailPaymentType(t testing.TB) {
	entry := mockEntryDetail()
	//override mockEntryDetail
	entry.TransactionCode = CheckingDebit
	entry.DiscretionaryData = "R"
	if err := entry.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestEntryDetailPaymentType tests validating an entry detail record payment type
func TestEntryDetailPaymentType(t *testing.T) {
	testEntryDetailPaymentType(t)
}

// BenchmarkEntryDetailPaymentType benchmarks validating an entry detail record payment type
func BenchmarkEntryDetailPaymentType(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEntryDetailPaymentType(b)
	}
}

// testEntryDetailReceivingCompany validates an entry detail record receiving company
func testEntryDetailReceivingCompany(t testing.TB) {
	entry := mockEntryDetail()
	//override mockEntryDetail
	entry.TransactionCode = CheckingDebit
	entry.IdentificationNumber = "location #23"
	entry.IndividualName = "Best Co. #23"

	if err := entry.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestEntryDetailReceivingCompany tests validating an entry detail record receiving company
func TestEntryDetailReceivingCompany(t *testing.T) {
	testEntryDetailReceivingCompany(t)
}

// BenchmarkEntryDetailReceivingCompany benchmarks validating an entry detail record receiving company
func BenchmarkEntryDetailReceivingCompany(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEntryDetailReceivingCompany(b)
	}
}

// testAddendaRecord validates an addenda record
func testAddendaRecord(t testing.TB) {
	addenda05 := NewAddenda05()
	addenda05.PaymentRelatedInformation = "Currently string needs ASC X12 Interchange Control Structures"
	addenda05.SequenceNumber = 1
	addenda05.EntryDetailSequenceNumber = 1234567

	if err := addenda05.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddendaRecord tests validating an addenda record
func TestAddendaRecord(t *testing.T) {
	testAddendaRecord(t)
}

// BenchmarkAddendaRecord benchmarks validating an addenda record
func BenchmarkAddendaRecord(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddendaRecord(b)
	}
}

// testBuildFile validates building a file
func testBuildFile(t testing.TB) {
	// To create a file
	file := NewFile()
	file.SetHeader(mockFileHeader())

	// To create a batch. Errors only if payment type is not supported.
	batch, _ := NewBatch(mockBatchHeader())

	// To create an entry
	entry := mockPPDEntryDetail()
	entry.AddendaRecordIndicator = 1

	// To add one or more optional addenda records for an entry
	addendaPPD := NewAddenda05()
	addendaPPD.PaymentRelatedInformation = "Currently string needs ASC X12 Interchange Control Structures"

	// Add the addenda record to the detail entry
	entry.AddAddenda05(addendaPPD)

	// Entries are added to batches like so:
	batch.AddEntry(entry)

	// When all of the Entries are added to the batch we must build it.

	if err := batch.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	// And batches are added to files much the same way:

	file.AddBatch(batch)

	// Now add a new batch for accepting payments on the web

	batch, _ = NewBatch(mockBatchWEBHeader())

	// Add an entry and define if it is a single or recurring payment
	// The following is a reoccuring payment for $7.99

	entry = mockWEBEntryDetail()
	entry.AddendaRecordIndicator = 1

	addendaWEB := NewAddenda05()
	addendaWEB.PaymentRelatedInformation = "Monthly Membership Subscription"

	// Add the addenda record to the detail entry
	entry.AddAddenda05(addendaWEB)

	// Add the entry to the batch
	batch.AddEntry(entry)

	// Now we must build this batch
	if err := batch.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	// And add the second batch to the file
	file.AddBatch(batch)

	// Once we added all our batches we must build the file
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	// Finally we write the file to an io.Writer
	var b bytes.Buffer
	w := NewWriter(&b)
	if err := w.Write(file); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	w.Flush()
}

// TestBuildFile tests validating building a file
func TestBuildFile(t *testing.T) {
	testBuildFile(t)
}

// BenchmarkBuildFile benchmarks validating building a file
func BenchmarkBuildFile(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBuildFile(b)
	}
}
