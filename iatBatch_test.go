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
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/base"
	"github.com/stretchr/testify/require"
)

// mockIATBatch
func mockIATBatch(t testing.TB) IATBatch {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATBatchHeaderFF())
	mockBatch.AddEntry(mockIATEntryDetailWithAddendas())
	if err := mockBatch.build(); err != nil {
		t.Fatal(err)
	}
	return mockBatch
}

func mockIATEntryDetailWithAddendas() *IATEntryDetail {
	ed := mockIATEntryDetail()
	ed.Addenda10 = mockAddenda10()
	ed.Addenda11 = mockAddenda11()
	ed.Addenda12 = mockAddenda12()
	ed.Addenda13 = mockAddenda13()
	ed.Addenda14 = mockAddenda14()
	ed.Addenda15 = mockAddenda15()
	ed.Addenda16 = mockAddenda16()
	return ed
}

// mockIATBatchManyEntries
func mockIATBatchManyEntries(t testing.TB) IATBatch {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATBatchHeaderFF())

	mockBatch.AddEntry(mockIATEntryDetail())

	mockBatch.Entries[0].Addenda10 = mockAddenda10()
	mockBatch.Entries[0].Addenda11 = mockAddenda11()
	mockBatch.Entries[0].Addenda12 = mockAddenda12()
	mockBatch.Entries[0].Addenda13 = mockAddenda13()
	mockBatch.Entries[0].Addenda14 = mockAddenda14()
	mockBatch.Entries[0].Addenda15 = mockAddenda15()
	mockBatch.Entries[0].Addenda16 = mockAddenda16()
	mockBatch.Entries[0].AddAddenda17(mockAddenda17())
	mockBatch.Entries[0].AddAddenda17(mockAddenda17B())
	mockBatch.Entries[0].AddAddenda18(mockAddenda18())
	mockBatch.Entries[0].AddAddenda18(mockAddenda18B())
	mockBatch.Entries[0].AddAddenda18(mockAddenda18C())
	mockBatch.Entries[0].AddAddenda18(mockAddenda18D())
	mockBatch.Entries[0].AddAddenda18(mockAddenda18E())

	mockBatch.AddEntry(mockIATEntryDetail2())

	mockBatch.Entries[1].Addenda10 = mockAddenda10()
	mockBatch.Entries[1].Addenda11 = mockAddenda11()
	mockBatch.Entries[1].Addenda12 = mockAddenda12()
	mockBatch.Entries[1].Addenda13 = mockAddenda13()
	mockBatch.Entries[1].Addenda14 = mockAddenda14()
	mockBatch.Entries[1].Addenda15 = mockAddenda15()
	mockBatch.Entries[1].Addenda16 = mockAddenda16()
	mockBatch.Entries[1].AddAddenda17(mockAddenda17())
	mockBatch.Entries[1].AddAddenda17(mockAddenda17B())
	mockBatch.Entries[1].AddAddenda18(mockAddenda18())
	mockBatch.Entries[1].AddAddenda18(mockAddenda18B())
	mockBatch.Entries[1].AddAddenda18(mockAddenda18C())
	mockBatch.Entries[1].AddAddenda18(mockAddenda18D())
	mockBatch.Entries[1].AddAddenda18(mockAddenda18E())

	if err := mockBatch.build(); err != nil {
		t.Fatal(err)
	}
	return mockBatch
}

// mockIATBatch
func mockInvalidIATBatch(t testing.TB) IATBatch {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATBatchHeaderFF())
	mockBatch.AddEntry(mockIATEntryDetail())
	mockBatch.Entries[0].Addenda10 = mockAddenda10()
	mockBatch.Entries[0].Addenda11 = mockAddenda11()
	mockBatch.Entries[0].Addenda12 = mockAddenda12()
	mockBatch.Entries[0].Addenda13 = mockAddenda13()
	mockBatch.Entries[0].Addenda14 = mockAddenda14()
	mockBatch.Entries[0].Addenda15 = mockAddenda15()
	mockBatch.Entries[0].Addenda16 = mockAddenda16()
	mockBatch.Entries[0].AddAddenda17(mockInvalidAddenda17())
	if err := mockBatch.build(); err != nil {
		t.Fatal(err)
	}
	return mockBatch
}

func mockInvalidAddenda17() *Addenda17 {
	addenda17 := NewAddenda17()
	addenda17.PaymentRelatedInformation = "Transfer of money from one country to another"
	addenda17.TypeCode = "02"
	addenda17.SequenceNumber = 2
	addenda17.EntryDetailSequenceNumber = 0000002
	return addenda17
}

func mockIATAddenda99() *Addenda99 {
	addenda99 := NewAddenda99()
	addenda99.ReturnCode = "R07"
	addenda99.OriginalTrace = "231380100000001"
	addenda99.OriginalDFI = "12104288"
	addenda99.IATPaymentAmount("0000100000")
	addenda99.IATAddendaInformation("Authorization Revoked")
	return addenda99
}

func mockIATAddenda98() *Addenda98 {
	addenda98 := NewAddenda98()
	addenda98.ChangeCode = "C01"
	addenda98.OriginalTrace = "231380100000001"
	addenda98.OriginalDFI = "12104288"
	addenda98.CorrectedData = "89722-C3"
	addenda98.TraceNumber = "121042880000001"
	return addenda98
}

// TestMockIATBatch validates mockIATBatch
func TestMockIATBatch(t *testing.T) {
	iatBatch := mockIATBatch(t)
	if err := iatBatch.verify(); err != nil {
		t.Error("mockIATBatch does not validate and will break other tests")
	}
}

// TestIATBatch__UnmarshalJSON reads an example File (with IAT Batches) and attempts to unmarshal it as JSON
func TestIATBatch__UnmarshalJSON(t *testing.T) {
	// Make sure we don't panic with nil in the mix
	var batch *IATBatch
	if err := batch.UnmarshalJSON(nil); err != nil && !strings.Contains(err.Error(), "unexpected end of JSON input") {
		t.Fatal(err)
	}

	// Read file, convert to JSON
	fd, err := os.Open(filepath.Join("test", "ach-iat-read", "iat-credit.ach"))
	if err != nil {
		t.Fatal(err)
	}
	f, err := NewReader(fd).Read()
	if err != nil {
		t.Fatal(err)
	}
	bs, err := json.Marshal(f)
	if err != nil {
		t.Fatal(err)
	}

	// Read as JSON
	file, err := FileFromJSON(bs)
	if err != nil {
		t.Fatal(err)
	}
	if file == nil {
		t.Error("file == nil")
	}
}

// testIATBatchAddenda10Error validates IATBatch returns an error if Addenda10 is not included
func testIATBatchAddenda10Error(t testing.TB) {
	iatBatch := mockIATBatch(t)
	iatBatch.GetEntries()[0].Addenda10 = nil
	err := iatBatch.verify()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchAddenda10Error tests validating IATBatch returns an error
// if Addenda10 is not included
func TestIATBatchAddenda10Error(t *testing.T) {
	testIATBatchAddenda10Error(t)
}

// BenchmarkIATBatchAddenda10Error benchmarks validating IATBatch returns an error
// if Addenda10 is not included
func BenchmarkIATBatchAddenda10Error(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda10Error(b)
	}
}

// testIATBatchAddenda11Error validates IATBatch returns an error if Addenda11 is not included
func testIATBatchAddenda11Error(t testing.TB) {
	iatBatch := mockIATBatch(t)
	iatBatch.GetEntries()[0].Addenda11 = nil
	err := iatBatch.verify()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchAddenda11Error tests validating IATBatch returns an error
// if Addenda11 is not included
func TestIATBatchAddenda11Error(t *testing.T) {
	testIATBatchAddenda11Error(t)
}

// BenchmarkIATBatchAddenda11Error benchmarks validating IATBatch returns an error
// if Addenda11 is not included
func BenchmarkIATBatchAddenda11Error(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda11Error(b)
	}
}

// testIATBatchAddenda12Error validates IATBatch returns an error if Addenda12 is not included
func testIATBatchAddenda12Error(t testing.TB) {
	iatBatch := mockIATBatch(t)
	iatBatch.GetEntries()[0].Addenda12 = nil
	err := iatBatch.verify()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchAddenda12Error tests validating IATBatch returns an error
// if Addenda12 is not included
func TestIATBatchAddenda12Error(t *testing.T) {
	testIATBatchAddenda12Error(t)
}

// BenchmarkIATBatchAddenda12Error benchmarks validating IATBatch returns an error
// if Addenda12 is not included
func BenchmarkIATBatchAddenda12Error(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda12Error(b)
	}
}

// testIATBatchAddenda13Error validates IATBatch returns an error if Addenda13 is not included
func testIATBatchAddenda13Error(t testing.TB) {
	iatBatch := mockIATBatch(t)
	iatBatch.GetEntries()[0].Addenda13 = nil
	err := iatBatch.verify()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchAddenda13Error tests validating IATBatch returns an error
// if Addenda13 is not included
func TestIATBatchAddenda13Error(t *testing.T) {
	testIATBatchAddenda13Error(t)
}

// BenchmarkIATBatchAddenda13Error benchmarks validating IATBatch returns an error
// if Addenda13 is not included
func BenchmarkIATBatchAddenda13Error(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda13Error(b)
	}
}

// testIATBatchAddenda14Error validates IATBatch returns an error if Addenda14 is not included
func testIATBatchAddenda14Error(t testing.TB) {
	iatBatch := mockIATBatch(t)
	iatBatch.GetEntries()[0].Addenda14 = nil
	err := iatBatch.verify()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchAddenda14Error tests validating IATBatch returns an error
// if Addenda14 is not included
func TestIATBatchAddenda14Error(t *testing.T) {
	testIATBatchAddenda14Error(t)
}

// BenchmarkIATBatchAddenda14Error benchmarks validating IATBatch returns an error
// if Addenda14 is not included
func BenchmarkIATBatchAddenda14Error(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda14Error(b)
	}
}

// testIATBatchAddenda15Error validates IATBatch returns an error if Addenda15 is not included
func testIATBatchAddenda15Error(t testing.TB) {
	iatBatch := mockIATBatch(t)
	iatBatch.GetEntries()[0].Addenda15 = nil
	err := iatBatch.verify()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchAddenda15Error tests validating IATBatch returns an error
// if Addenda15 is not included
func TestIATBatchAddenda15Error(t *testing.T) {
	testIATBatchAddenda15Error(t)
}

// BenchmarkIATBatchAddenda15Error benchmarks validating IATBatch returns an error
// if Addenda15 is not included
func BenchmarkIATBatchAddenda15Error(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda15Error(b)
	}
}

// testIATBatchAddenda16Error validates IATBatch returns an error if Addenda16 is not included
func testIATBatchAddenda16Error(t testing.TB) {
	iatBatch := mockIATBatch(t)
	iatBatch.GetEntries()[0].Addenda16 = nil
	err := iatBatch.verify()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchAddenda16Error tests validating IATBatch returns an error
// if Addenda16 is not included
func TestIATBatchAddenda16Error(t *testing.T) {
	testIATBatchAddenda16Error(t)
}

// BenchmarkIATBatchAddenda16Error benchmarks validating IATBatch returns an error
// if Addenda16 is not included
func BenchmarkIATBatchAddenda16Error(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda16Error(b)
	}
}

// testAddenda10EntryDetailSequenceNumber validates IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func testAddenda10EntryDetailSequenceNumber(t testing.TB) {
	iatBatch := mockIATBatch(t)
	iatBatch.GetEntries()[0].Addenda10.EntryDetailSequenceNumber = 00000005
	err := iatBatch.verify()
	if !base.Match(err, NewErrBatchAddendaTraceNumber("0000005", "0000001")) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda10EntryDetailSequenceNumber tests validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func TestAddenda10EntryDetailSequenceNumber(t *testing.T) {
	testAddenda10EntryDetailSequenceNumber(t)
}

// BenchmarkAddenda10EntryDetailSequenceNumber benchmarks validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func BenchmarkAddenda10EntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda10EntryDetailSequenceNumber(b)
	}
}

// testAddenda11EntryDetailSequenceNumber validates IATBatch returns an error if EntryDetailSequenceNumber
// is not valid
func testAddenda11EntryDetailSequenceNumber(t testing.TB) {
	iatBatch := mockIATBatch(t)
	iatBatch.GetEntries()[0].Addenda11.EntryDetailSequenceNumber = 00000005
	err := iatBatch.verify()
	if !base.Match(err, NewErrBatchAddendaTraceNumber("0000005", "0000001")) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda11EntryDetailSequenceNumber tests validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func TestAddenda11EntryDetailSequenceNumber(t *testing.T) {
	testAddenda11EntryDetailSequenceNumber(t)
}

// BenchmarkAddenda11EntryDetailSequenceNumber benchmarks validating IATBatch returns an error
// if EntryDetailSequenceNumber is not valid
func BenchmarkAddenda11EntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda11EntryDetailSequenceNumber(b)
	}
}

// testAddenda12EntryDetailSequenceNumber validates IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func testAddenda12EntryDetailSequenceNumber(t testing.TB) {
	iatBatch := mockIATBatch(t)
	iatBatch.GetEntries()[0].Addenda12.EntryDetailSequenceNumber = 00000005
	err := iatBatch.verify()
	if !base.Match(err, NewErrBatchAddendaTraceNumber("0000005", "0000001")) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda12EntryDetailSequenceNumber tests validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func TestAddenda12EntryDetailSequenceNumber(t *testing.T) {
	testAddenda12EntryDetailSequenceNumber(t)
}

// BenchmarkAddenda12EntryDetailSequenceNumber benchmarks validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func BenchmarkAddenda12EntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda12EntryDetailSequenceNumber(b)
	}
}

// testAddenda13EntryDetailSequenceNumber validates IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func testAddenda13EntryDetailSequenceNumber(t testing.TB) {
	iatBatch := mockIATBatch(t)
	iatBatch.GetEntries()[0].Addenda13.EntryDetailSequenceNumber = 00000005
	err := iatBatch.verify()
	if !base.Match(err, NewErrBatchAddendaTraceNumber("0000005", "0000001")) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda13EntryDetailSequenceNumber tests validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func TestAddenda13EntryDetailSequenceNumber(t *testing.T) {
	testAddenda13EntryDetailSequenceNumber(t)
}

// BenchmarkAddenda13EntryDetailSequenceNumber benchmarks validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func BenchmarkAddenda13EntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda13EntryDetailSequenceNumber(b)
	}
}

// testAddenda14EntryDetailSequenceNumber validates IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func testAddenda14EntryDetailSequenceNumber(t testing.TB) {
	iatBatch := mockIATBatch(t)
	iatBatch.GetEntries()[0].Addenda14.EntryDetailSequenceNumber = 00000005
	err := iatBatch.verify()
	if !base.Match(err, NewErrBatchAddendaTraceNumber("0000005", "0000001")) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda14EntryDetailSequenceNumber tests validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func TestAddenda14EntryDetailSequenceNumber(t *testing.T) {
	testAddenda14EntryDetailSequenceNumber(t)
}

// BenchmarkAddenda14EntryDetailSequenceNumber benchmarks validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func BenchmarkAddenda14EntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda14EntryDetailSequenceNumber(b)
	}
}

// testAddenda15EntryDetailSequenceNumber validates IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func testAddenda15EntryDetailSequenceNumber(t testing.TB) {
	iatBatch := mockIATBatch(t)
	iatBatch.GetEntries()[0].Addenda15.EntryDetailSequenceNumber = 00000005
	err := iatBatch.verify()
	if !base.Match(err, NewErrBatchAddendaTraceNumber("0000005", "0000001")) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda15EntryDetailSequenceNumber tests validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func TestAddenda15EntryDetailSequenceNumber(t *testing.T) {
	testAddenda15EntryDetailSequenceNumber(t)
}

// BenchmarkAddenda15EntryDetailSequenceNumber benchmarks validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func BenchmarkAddenda15EntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda15EntryDetailSequenceNumber(b)
	}
}

// testAddenda16EntryDetailSequenceNumber validates IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func testAddenda16EntryDetailSequenceNumber(t testing.TB) {
	iatBatch := mockIATBatch(t)
	iatBatch.GetEntries()[0].Addenda16.EntryDetailSequenceNumber = 00000005
	err := iatBatch.verify()
	if !base.Match(err, NewErrBatchAddendaTraceNumber("0000005", "0000001")) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestAddenda16EntryDetailSequenceNumber tests validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func TestAddenda16EntryDetailSequenceNumber(t *testing.T) {
	testAddenda16EntryDetailSequenceNumber(t)
}

// BenchmarkAddenda16EntryDetailSequenceNumber benchmarks validating IATBatch returns an error if
// EntryDetailSequenceNumber is not valid
func BenchmarkAddenda16EntryDetailSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddenda16EntryDetailSequenceNumber(b)
	}
}

// testIATBatchNumberMismatch validates BatchNumber mismatch
func testIATBatchNumberMismatch(t testing.TB) {
	mockBatch := mockIATBatch(t)
	mockBatch.GetControl().BatchNumber = 2
	err := mockBatch.verify()
	if !base.Match(err, NewErrBatchHeaderControlEquality(1, 2)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchNumberMismatch tests validating BatchNumber mismatch
func TestIATBatchNumberMismatch(t *testing.T) {
	testIATBatchNumberMismatch(t)
}

// BenchmarkIATBatchNumberMismatch benchmarks validating BatchNumber mismatch
func BenchmarkIATBatchNumberMismatch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchNumberMismatch(b)
	}
}

// testIATServiceClassCodeMismatch validates ServiceClassCode mismatch
func testIATServiceClassCodeMismatch(t testing.TB) {
	mockBatch := mockIATBatch(t)
	mockBatch.GetControl().ServiceClassCode = DebitsOnly
	err := mockBatch.verify()
	if !base.Match(err, NewErrBatchHeaderControlEquality(CreditsOnly, DebitsOnly)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATServiceClassCodeMismatch tests validating ServiceClassCode mismatch
func TestServiceClassCodeMismatch(t *testing.T) {
	testIATServiceClassCodeMismatch(t)
}

// BenchmarkIATServiceClassCoderMismatch benchmarks validating ServiceClassCode mismatch
func BenchmarkIATServiceClassCodeMismatch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATServiceClassCodeMismatch(b)
	}
}

// testIATBatchCreditIsBatchAmount validates credit isBatchAmount
func testIATBatchCreditIsBatchAmount(t testing.TB) {
	mockBatch := mockIATBatch(t)
	e1 := mockBatch.GetEntries()[0]
	e2 := mockIATEntryDetail()
	e2.TransactionCode = CheckingCredit
	e2.Amount = 5000
	// replace last 2 of TraceNumber
	e2.TraceNumber = e1.TraceNumber[:13] + "10"
	mockBatch.AddEntry(e2)
	mockBatch.Entries[1].Addenda10 = mockAddenda10()
	mockBatch.Entries[1].Addenda11 = mockAddenda11()
	mockBatch.Entries[1].Addenda12 = mockAddenda12()
	mockBatch.Entries[1].Addenda13 = mockAddenda13()
	mockBatch.Entries[1].Addenda14 = mockAddenda14()
	mockBatch.Entries[1].Addenda15 = mockAddenda15()
	mockBatch.Entries[1].Addenda16 = mockAddenda16()
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	mockBatch.GetControl().TotalCreditEntryDollarAmount = 1000
	err := mockBatch.verify()
	if !base.Match(err, NewErrBatchCalculatedControlEquality(105000, 1000)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchCreditIsBatchAmount tests validating credit isBatchAmount
func TestIATBatchCreditIsBatchAmount(t *testing.T) {
	testIATBatchCreditIsBatchAmount(t)
}

// BenchmarkIATBatchCreditIsBatchAmount benchmarks validating credit isBatchAmount
func BenchmarkIATBatchCreditIsBatchAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchCreditIsBatchAmount(b)
	}

}

// testIATBatchDebitIsBatchAmount validates debit isBatchAmount
func testIATBatchDebitIsBatchAmount(t testing.TB) {
	mockBatch := mockIATBatch(t)
	e1 := mockBatch.GetEntries()[0]
	e1.TransactionCode = CheckingDebit
	e2 := mockIATEntryDetail()
	e2.TransactionCode = CheckingDebit
	e2.Amount = 5000
	// replace last 2 of TraceNumber
	e2.TraceNumber = e1.TraceNumber[:13] + "10"
	mockBatch.AddEntry(e2)
	mockBatch.Entries[1].Addenda10 = mockAddenda10()
	mockBatch.Entries[1].Addenda11 = mockAddenda11()
	mockBatch.Entries[1].Addenda12 = mockAddenda12()
	mockBatch.Entries[1].Addenda13 = mockAddenda13()
	mockBatch.Entries[1].Addenda14 = mockAddenda14()
	mockBatch.Entries[1].Addenda15 = mockAddenda15()
	mockBatch.Entries[1].Addenda16 = mockAddenda16()
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	mockBatch.GetControl().TotalDebitEntryDollarAmount = 1000
	err := mockBatch.verify()
	if !base.Match(err, NewErrBatchCalculatedControlEquality(105000, 1000)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchDebitIsBatchAmount tests validating debit isBatchAmount
func TestIATBatchDebitIsBatchAmount(t *testing.T) {
	testIATBatchDebitIsBatchAmount(t)
}

// BenchmarkIATBatchDebitIsBatchAmount benchmarks validating debit isBatchAmount
func BenchmarkIATBatchDebitIsBatchAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchDebitIsBatchAmount(b)
	}

}

// testIATBatchFieldInclusion validates IATBatch FieldInclusion
func testIATBatchFieldInclusion(t testing.TB) {
	mockBatch := mockIATBatch(t)
	mockBatch2 := mockIATBatch(t)
	mockBatch2.Header.ServiceClassCode = 4

	err := mockBatch.verify()
	// no errors expected
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
	err = mockBatch2.verify()
	if !base.Match(err, ErrServiceClass) {
		t.Errorf("%T: %s", err, err)
	}
	err = mockBatch2.build()
	if !base.Match(err, ErrServiceClass) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchFieldInclusion tests validating IATBatch FieldInclusion
func TestIATBatchFieldInclusion(t *testing.T) {
	testIATBatchFieldInclusion(t)
}

// BenchmarkIATBatchFieldInclusion benchmarks validating IATBatch FieldInclusion
func BenchmarkIATBatchFieldInclusion(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchFieldInclusion(b)
	}

}

// testIATBatchBuild validates IATBatch build error
func testIATBatchBuild(t testing.TB) {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATBatchHeaderFF())

	err := mockBatch.build()
	if !base.Match(err, ErrBatchNoEntries) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchBuild tests validating IATBatch build error
func TestIATBatchBuild(t *testing.T) {
	testIATBatchBuild(t)
}

// BenchmarkIATBatchBuild benchmarks validating IATBatch build error
func BenchmarkIATBatchBuild(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchBuild(b)
	}

}

// testIATODFIIdentificationMismatch validates ODFIIdentification mismatch
func testIATODFIIdentificationMismatch(t testing.TB) {
	mockBatch := mockIATBatch(t)
	mockBatch.GetControl().ODFIIdentification = "53158020"
	err := mockBatch.verify()
	if !base.Match(err, NewErrBatchHeaderControlEquality(23138010, 53158020)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATODFIIdentificationMismatch tests validating ODFIIdentification mismatch
func TestODFIIdentificationMismatch(t *testing.T) {
	testIATODFIIdentificationMismatch(t)
}

// BenchmarkIATODFIIdentificationMismatch benchmarks validating ODFIIdentification mismatch
func BenchmarkIATODFIIdentificationMismatch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATODFIIdentificationMismatch(b)
	}
}

// testIATBatchAddendaRecordIndicator validates IATEntryDetail AddendaRecordIndicator
func testIATBatchAddendaRecordIndicator(t testing.TB) {
	mockBatch := mockIATBatch(t)
	mockBatch.GetEntries()[0].AddendaRecordIndicator = 2
	err := mockBatch.verify()
	if !base.Match(err, ErrIATBatchAddendaIndicator) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchAddendaRecordIndicator tests validating IATEntryDetail AddendaRecordIndicator
func TestIATBatchAddendaRecordIndicator(t *testing.T) {
	testIATBatchAddendaRecordIndicator(t)
}

// BenchmarkIATBatchAddendaRecordIndicator benchmarks IATEntryDetail AddendaRecordIndicator
func BenchmarkIATBatchAddendaRecordIndicator(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddendaRecordIndicator(b)
	}
}

// testIATBatchInvalidTraceNumberODFI validates TraceNumberODFI
func testIATBatchInvalidTraceNumberODFI(t testing.TB) {
	mockBatch := mockIATBatch(t)
	mockBatch.GetEntries()[0].SetTraceNumber("9928272", 1)
	err := mockBatch.verify()
	if !base.Match(err, NewErrBatchTraceNumberNotODFI("23138010", "09928272")) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchInvalidTraceNumberODFI tests validating TraceNumberODFI
func TestIATBatchInvalidTraceNumberODFI(t *testing.T) {
	testIATBatchInvalidTraceNumberODFI(t)
}

// BenchmarkIATBatchInvalidTraceNumberODFI benchmarks validating TraceNumberODFI
func BenchmarkIATBatchInvalidTraceNumberODFI(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchInvalidTraceNumberODFI(b)
	}
}

// testIATBatchControl validates BatchControl ODFIIdentification
func testIATBatchControl(t testing.TB) {
	mockBatch := mockIATBatch(t)
	mockBatch.Control.ODFIIdentification = ""
	err := mockBatch.verify()
	if !base.Match(err, NewErrBatchHeaderControlEquality("23138010", "")) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchControl tests validating BatchControl ODFIIdentification
func TestIATBatchControl(t *testing.T) {
	testIATBatchControl(t)
}

// BenchmarkIATBatchControl benchmarks validating BatchControl ODFIIdentification
func BenchmarkIATBatchControl(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchControl(b)
	}
}

// testIATBatchEntryCountEquality validates IATBatch EntryAddendaCount
func testIATBatchEntryCountEquality(t testing.TB) {
	mockBatch := mockIATBatch(t)
	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	mockBatch.GetControl().EntryAddendaCount = 1
	err := mockBatch.verify()
	if !base.Match(err, NewErrBatchCalculatedControlEquality(8, 1)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchEntryCountEquality tests validating IATBatch EntryAddendaCount
func TestIATBatchEntryCountEquality(t *testing.T) {
	testIATBatchEntryCountEquality(t)
}

// BenchmarkIATBatchEntryCountEquality benchmarks validating IATBatch EntryAddendaCount
func BenchmarkIATBatchEntryCountEquality(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchEntryCountEquality(b)
	}
}

// testIATBatchisEntryHash validates IATBatch EntryHash
func testIATBatchisEntryHash(t testing.TB) {
	mockBatch := mockIATBatch(t)
	mockBatch.GetControl().EntryHash = 1
	err := mockBatch.verify()
	if !base.Match(err, NewErrBatchCalculatedControlEquality("0012104288", "1")) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchisEntryHash tests validating IATBatch EntryHash
func TestIATBatchisEntryHash(t *testing.T) {
	testIATBatchisEntryHash(t)
}

// BenchmarkIATBatchisEntryHash benchmarks validating IATBatch EntryHash
func BenchmarkIATBatchisEntryHash(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchisEntryHash(b)
	}
}

// testIATBatchIsSequenceAscending validates sequence ascending
func testIATBatchIsSequenceAscending(t testing.TB) {
	mockBatch := mockIATBatch(t)
	e2 := mockIATEntryDetail()
	e2.TraceNumber = "1"
	mockBatch.AddEntry(e2)
	mockBatch.Entries[1].Addenda10 = mockAddenda10()
	mockBatch.Entries[1].Addenda11 = mockAddenda11()
	mockBatch.Entries[1].Addenda12 = mockAddenda12()
	mockBatch.Entries[1].Addenda13 = mockAddenda13()
	mockBatch.Entries[1].Addenda14 = mockAddenda14()
	mockBatch.Entries[1].Addenda15 = mockAddenda15()
	mockBatch.Entries[1].Addenda16 = mockAddenda16()
	mockBatch.GetControl().EntryAddendaCount = 16
	err := mockBatch.verify()
	if !base.Match(err, NewErrBatchAscending("231380100000001", "1")) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchIsSequenceAscending tests validating sequence ascending
func TestIATBatchIsSequenceAscending(t *testing.T) {
	testIATBatchIsSequenceAscending(t)
}

// BenchmarkIATBatchIsSequenceAscending tests validating sequence ascending
func BenchmarkIATBatchIsSequenceAscending(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchIsSequenceAscending(b)
	}
}

// testIATBatchIsCategory validates category
func testIATBatchIsCategory(t testing.TB) {
	mockBatch := mockIATBatchManyEntries(t)
	mockBatch.GetEntries()[1].Category = CategoryReturn

	err := mockBatch.verify()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchIsCategory tests validating category
func TestIATBatchIsCategory(t *testing.T) {
	testIATBatchIsCategory(t)
}

// BenchmarkIATBatchIsCategory tests validating category
func BenchmarkIATBatchIsCategory(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchIsCategory(b)
	}
}

// testIATBatchCategory tests IATBatch Category
func testIATBatchCategory(t testing.TB) {
	mockBatch := mockIATBatch(t)

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	if mockBatch.Category() != CategoryForward {
		t.Errorf("No returns and Category is %s", mockBatch.Category())
	}
}

// TestIATBatchCategory tests IATBatch Category
func TestIATBatchCategory(t *testing.T) {
	testIATBatchCategory(t)
}

// BenchmarkIATBatchCategory benchmarks IATBatch Category
func BenchmarkIATBatchCategory(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchCategory(b)
	}
}

// testIATBatchValidateEntry validates EntryDetail
func testIATBatchValidateEntry(t testing.TB) {
	mockBatch := mockIATBatch(t)
	mockBatch.GetEntries()[0].TransactionCode = 5

	err := mockBatch.verify()
	if !base.Match(err, ErrTransactionCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchValidateEntry tests validating Entry
func TestIATBatchValidateEntry(t *testing.T) {
	testIATBatchValidateEntry(t)
}

// BenchmarkIATBatchValidateEntry tests validating Entry
func BenchmarkIATBatchValidateEntry(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidateEntry(b)
	}
}

// testIATBatchValidateAddenda10 validates Addenda10
func testIATBatchValidateAddenda10(t testing.TB) {
	mockBatch := mockIATBatchManyEntries(t)
	mockBatch.GetEntries()[1].Addenda10.TypeCode = "02"

	err := mockBatch.verify()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchValidateAddenda10 tests validating Addenda10
func TestIATBatchValidateAddenda10(t *testing.T) {
	testIATBatchValidateAddenda10(t)
}

// BenchmarkIATBatchValidateAddenda10 tests validating Addenda10
func BenchmarkIATBatchValidateAddenda10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidateAddenda10(b)
	}
}

// testIATBatchValidateAddenda11 validates Addenda11
func testIATBatchValidateAddenda11(t testing.TB) {
	mockBatch := mockIATBatchManyEntries(t)
	mockBatch.GetEntries()[1].Addenda11.TypeCode = "02"

	err := mockBatch.verify()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchValidateAddenda11 tests validating Addenda11
func TestIATBatchValidateAddenda11(t *testing.T) {
	testIATBatchValidateAddenda11(t)
}

// BenchmarkIATBatchValidateAddenda11 tests validating Addenda11
func BenchmarkIATBatchValidateAddenda11(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidateAddenda11(b)
	}
}

// testIATBatchValidateAddenda12 validates Addenda12
func testIATBatchValidateAddenda12(t testing.TB) {
	mockBatch := mockIATBatchManyEntries(t)
	mockBatch.GetEntries()[1].Addenda12.TypeCode = "02"

	err := mockBatch.verify()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchValidateAddenda12 tests validating Addenda12
func TestIATBatchValidateAddenda12(t *testing.T) {
	testIATBatchValidateAddenda12(t)
}

// BenchmarkIATBatchValidateAddenda12 tests validating Addenda12
func BenchmarkIATBatchValidateAddenda12(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidateAddenda12(b)
	}
}

// testIATBatchValidateAddenda13 validates Addenda13
func testIATBatchValidateAddenda13(t testing.TB) {
	mockBatch := mockIATBatchManyEntries(t)
	mockBatch.GetEntries()[1].Addenda13.TypeCode = "02"

	err := mockBatch.verify()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchValidateAddenda13 tests validating Addenda13
func TestIATBatchValidateAddenda13(t *testing.T) {
	testIATBatchValidateAddenda13(t)
}

// BenchmarkIATBatchValidateAddenda13 tests validating Addenda13
func BenchmarkIATBatchValidateAddenda13(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidateAddenda13(b)
	}
}

// testIATBatchValidateAddenda14 validates Addenda14
func testIATBatchValidateAddenda14(t testing.TB) {
	mockBatch := mockIATBatchManyEntries(t)
	mockBatch.GetEntries()[1].Addenda14.TypeCode = "02"

	err := mockBatch.verify()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchValidateAddenda14 tests validating Addenda14
func TestIATBatchValidateAddenda14(t *testing.T) {
	testIATBatchValidateAddenda14(t)
}

// BenchmarkIATBatchValidateAddenda14 tests validating Addenda14
func BenchmarkIATBatchValidateAddenda14(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidateAddenda14(b)
	}
}

// testIATBatchValidateAddenda15 validates Addenda15
func testIATBatchValidateAddenda15(t testing.TB) {
	mockBatch := mockIATBatchManyEntries(t)
	mockBatch.GetEntries()[1].Addenda15.TypeCode = "02"

	err := mockBatch.verify()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchValidateAddenda15 tests validating Addenda15
func TestIATBatchValidateAddenda15(t *testing.T) {
	testIATBatchValidateAddenda15(t)
}

// BenchmarkIATBatchValidateAddenda15 tests validating Addenda15
func BenchmarkIATBatchValidateAddenda15(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidateAddenda15(b)
	}
}

// testIATBatchValidateAddenda16 validates Addenda16
func testIATBatchValidateAddenda16(t testing.TB) {
	mockBatch := mockIATBatchManyEntries(t)
	mockBatch.GetEntries()[1].Addenda16.TypeCode = "02"

	err := mockBatch.verify()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchValidateAddenda16 tests validating Addenda16
func TestIATBatchValidateAddenda16(t *testing.T) {
	testIATBatchValidateAddenda16(t)
}

// BenchmarkIATBatchValidateAddenda16 tests validating Addenda16
func BenchmarkIATBatchValidateAddenda16(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidateAddenda16(b)
	}
}

// testIATBatchValidateAddenda17 validates Addenda17
func testIATBatchValidateAddenda17(t testing.TB) {
	mockBatch := mockInvalidIATBatch(t)

	err := mockBatch.verify()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchValidateAddenda17 tests validating Addenda17
func TestIATBatchValidateAddenda17(t *testing.T) {
	testIATBatchValidateAddenda17(t)
}

// BenchmarkIATBatchValidateAddenda17 tests validating Addenda17
func BenchmarkIATBatchValidateAddenda17(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidateAddenda17(b)
	}
}

// testIATBatchCreateError validates IATBatch create error
func testIATBatchCreate(t testing.TB) {
	file := NewFile().SetHeader(mockFileHeader())
	mockBatch := mockIATBatch(t)
	mockBatch.GetHeader().ServiceClassCode = 7

	err := mockBatch.Create()
	if !base.Match(err, ErrServiceClass) {
		t.Errorf("%T: %s", err, err)
	}

	file.AddIATBatch(mockBatch)
}

// TestIATBatchCreate tests validating IATBatch create error
func TestIATBatchCreate(t *testing.T) {
	testIATBatchCreate(t)
}

// BenchmarkIATBatchCreate benchmarks validating IATBatch create error
func BenchmarkIATBatchCreate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchCreate(b)
	}

}

// testIATBatchValidate validates IATBatch validate error
func testIATBatchValidate(t testing.TB) {
	file := NewFile().SetHeader(mockFileHeader())
	mockBatch := mockIATBatch(t)
	mockBatch.GetHeader().ServiceClassCode = DebitsOnly

	err := mockBatch.verify()
	if !base.Match(err, NewErrBatchHeaderControlEquality(DebitsOnly, CreditsOnly)) {
		t.Errorf("%T: %s", err, err)
	}

	file.AddIATBatch(mockBatch)
}

// TestIATBatchValidate tests validating IATBatch validate error
func TestIATBatchValidate(t *testing.T) {
	testIATBatchValidate(t)
}

// BenchmarkIATBatchValidate benchmarks validating IATBatch validate error
func BenchmarkIATBatchValidate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchValidate(b)
	}

}

// testIATBatchEntryAddendum validates IATBatch EntryAddendum error
func testIATBatchEntryAddendum(t testing.TB) {
	file := NewFile().SetHeader(mockFileHeader())
	mockBatch := mockIATBatch(t)
	mockBatch.Entries[0].AddAddenda17(mockAddenda17())
	mockBatch.Entries[0].AddAddenda17(mockAddenda17B())
	mockBatch.Entries[0].AddAddenda18(mockAddenda18())
	mockBatch.Entries[0].AddAddenda18(mockAddenda18B())
	mockBatch.Entries[0].AddAddenda18(mockAddenda18C())
	mockBatch.Entries[0].AddAddenda18(mockAddenda18D())
	mockBatch.Entries[0].AddAddenda18(mockAddenda18E())
	mockBatch.Entries[0].AddAddenda18(mockAddenda18F())

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchAddendaCount(6, 5)) {
		t.Errorf("%T: %s", err, err)
	}

	file.AddIATBatch(mockBatch)
}

// TestIATBatchEntryAddendum tests validating IATBatch EntryAddendum error
func TestIATBatchEntryAddendum(t *testing.T) {
	testIATBatchEntryAddendum(t)
}

// BenchmarkIATBatchEntryAddendum benchmarks validating IATBatch EntryAddendum error
func BenchmarkIATBatchEntryAddendum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchEntryAddendum(b)
	}
}

// testIATBatchAddenda17EDSequenceNumber validates IATBatch Addenda17 Entry Detail Sequence Number error
func testIATBatchAddenda17EDSequenceNumber(t testing.TB) {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATBatchHeaderFF())
	mockBatch.AddEntry(mockIATEntryDetail())
	addenda17 := NewAddenda17()
	addenda17.PaymentRelatedInformation = "This is an international payment"
	addenda17.SequenceNumber = 1
	addenda17.EntryDetailSequenceNumber = 0000001
	addenda17B := NewAddenda17()
	addenda17B.PaymentRelatedInformation = "Transfer of money from one country to another"
	addenda17B.SequenceNumber = 2
	addenda17B.EntryDetailSequenceNumber = 0000001
	mockBatch.Entries[0].Addenda10 = mockAddenda10()
	mockBatch.Entries[0].Addenda11 = mockAddenda11()
	mockBatch.Entries[0].Addenda12 = mockAddenda12()
	mockBatch.Entries[0].Addenda13 = mockAddenda13()
	mockBatch.Entries[0].Addenda14 = mockAddenda14()
	mockBatch.Entries[0].Addenda15 = mockAddenda15()
	mockBatch.Entries[0].Addenda16 = mockAddenda16()
	mockBatch.Entries[0].AddAddenda17(addenda17)
	mockBatch.Entries[0].AddAddenda17(addenda17B)

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	addenda17B.SequenceNumber = 1
	addenda17B.EntryDetailSequenceNumber = 0000002
	err := mockBatch.verify()
	if !base.Match(err, NewErrBatchAddendaTraceNumber("0000002", "0000001")) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchAddenda17EDSequenceNumber tests validating IATBatch Addenda17 Entry Detail Sequence Number error
func TestIATBatchAddenda17EDSequenceNumber(t *testing.T) {
	testIATBatchAddenda17EDSequenceNumber(t)
}

// BenchmarkIATBatchAddenda17EDSequenceNumber benchmarks validating IATBatch Addenda17 Entry Detail Sequence Number error
func BenchmarkIATBatchAddenda17EDSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda17EDSequenceNumber(b)
	}
}

// testIATBatchAddenda17Sequence validates IATBatch Addenda17 Sequence Number error
func testIATBatchAddenda17Sequence(t testing.TB) {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATBatchHeaderFF())
	mockBatch.AddEntry(mockIATEntryDetail())
	addenda17 := NewAddenda17()
	addenda17.PaymentRelatedInformation = "This is an international payment"
	addenda17.SequenceNumber = 2
	addenda17.EntryDetailSequenceNumber = 0000001
	addenda17B := NewAddenda17()
	addenda17B.PaymentRelatedInformation = "Transfer of money from one country to another"
	addenda17B.SequenceNumber = 1
	addenda17B.EntryDetailSequenceNumber = 0000001
	mockBatch.Entries[0].Addenda10 = mockAddenda10()
	mockBatch.Entries[0].Addenda11 = mockAddenda11()
	mockBatch.Entries[0].Addenda12 = mockAddenda12()
	mockBatch.Entries[0].Addenda13 = mockAddenda13()
	mockBatch.Entries[0].Addenda14 = mockAddenda14()
	mockBatch.Entries[0].Addenda15 = mockAddenda15()
	mockBatch.Entries[0].Addenda16 = mockAddenda16()
	mockBatch.Entries[0].AddAddenda17(addenda17)
	mockBatch.Entries[0].AddAddenda17(addenda17B)

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	addenda17B.SequenceNumber = -1
	err := mockBatch.verify()
	if !base.Match(err, NewErrBatchAscending("-1", "1")) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchAddenda17Sequence tests validating IATBatch Addenda17 Sequence Number error
func TestIATBatchAddenda17Sequence(t *testing.T) {
	testIATBatchAddenda17Sequence(t)
}

// BenchmarkIATBatchAddenda17Sequence benchmarks validating IATBatch Addenda17 Sequence Number error
func BenchmarkIATBatchAddenda17Sequence(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda17Sequence(b)
	}
}

// testIATBatchAddenda18EDSequenceNumber validates IATBatch Addenda18 Entry Detail Sequence Number error
func testIATBatchAddenda18EDSequenceNumber(t testing.TB) {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATBatchHeaderFF())
	mockBatch.AddEntry(mockIATEntryDetail())
	addenda17 := NewAddenda17()
	addenda17.PaymentRelatedInformation = "This is an international payment"
	addenda17.SequenceNumber = 1
	addenda17.EntryDetailSequenceNumber = 0000001
	addenda17B := NewAddenda17()
	addenda17B.PaymentRelatedInformation = "Transfer of money from one country to another"
	addenda17B.SequenceNumber = 2
	addenda17B.EntryDetailSequenceNumber = 0000001
	addenda18 := NewAddenda18()
	addenda18.ForeignCorrespondentBankName = "Bank of Turkey"
	addenda18.ForeignCorrespondentBankIDNumberQualifier = "01"
	addenda18.ForeignCorrespondentBankIDNumber = "12312345678910"
	addenda18.ForeignCorrespondentBankBranchCountryCode = "TR"
	addenda18.SequenceNumber = 1
	addenda18.EntryDetailSequenceNumber = 0000001
	addenda18B := NewAddenda18()
	addenda18B.ForeignCorrespondentBankName = "Bank of United Kingdom"
	addenda18B.ForeignCorrespondentBankIDNumberQualifier = "01"
	addenda18B.ForeignCorrespondentBankIDNumber = "1234567890123456789012345678901234"
	addenda18B.ForeignCorrespondentBankBranchCountryCode = "GB"
	addenda18B.SequenceNumber = 2
	addenda18B.EntryDetailSequenceNumber = 0000001
	mockBatch.Entries[0].Addenda10 = mockAddenda10()
	mockBatch.Entries[0].Addenda11 = mockAddenda11()
	mockBatch.Entries[0].Addenda12 = mockAddenda12()
	mockBatch.Entries[0].Addenda13 = mockAddenda13()
	mockBatch.Entries[0].Addenda14 = mockAddenda14()
	mockBatch.Entries[0].Addenda15 = mockAddenda15()
	mockBatch.Entries[0].Addenda16 = mockAddenda16()
	mockBatch.Entries[0].AddAddenda17(addenda17)
	mockBatch.Entries[0].AddAddenda17(addenda17B)
	mockBatch.Entries[0].AddAddenda18(addenda18)
	mockBatch.Entries[0].AddAddenda18(addenda18B)

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	addenda18B.SequenceNumber = 1
	addenda18B.EntryDetailSequenceNumber = 0000002
	err := mockBatch.verify()
	if !base.Match(err, NewErrBatchAddendaTraceNumber("0000002", "0000001")) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchAddenda18EDSequenceNumber tests validating IATBatch Addenda18 Entry Detail Sequence Number error
func TestIATBatchAddenda18EDSequenceNumber(t *testing.T) {
	testIATBatchAddenda18EDSequenceNumber(t)
}

// BenchmarkIATBatchAddenda18EDSequenceNumber benchmarks validating IATBatch Addenda18 Entry Detail Sequence Number error
func BenchmarkIATBatchAddenda18EDSequenceNumber(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda18EDSequenceNumber(b)
	}
}

// testIATBatchAddenda18Sequence validates IATBatch Addenda18 Sequence Number error
func testIATBatchAddenda18Sequence(t testing.TB) {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATBatchHeaderFF())
	mockBatch.AddEntry(mockIATEntryDetail())
	addenda17 := NewAddenda17()
	addenda17.PaymentRelatedInformation = "This is an international payment"
	addenda17.SequenceNumber = 1
	addenda17.EntryDetailSequenceNumber = 0000001
	addenda17B := NewAddenda17()
	addenda17B.PaymentRelatedInformation = "Transfer of money from one country to another"
	addenda17B.SequenceNumber = 2
	addenda17B.EntryDetailSequenceNumber = 0000001
	addenda18 := NewAddenda18()
	addenda18.ForeignCorrespondentBankName = "Bank of Turkey"
	addenda18.ForeignCorrespondentBankIDNumberQualifier = "01"
	addenda18.ForeignCorrespondentBankIDNumber = "12312345678910"
	addenda18.ForeignCorrespondentBankBranchCountryCode = "TR"
	addenda18.SequenceNumber = 1
	addenda18.EntryDetailSequenceNumber = 0000001
	addenda18B := NewAddenda18()
	addenda18B.ForeignCorrespondentBankName = "Bank of United Kingdom"
	addenda18B.ForeignCorrespondentBankIDNumberQualifier = "01"
	addenda18B.ForeignCorrespondentBankIDNumber = "1234567890123456789012345678901234"
	addenda18B.ForeignCorrespondentBankBranchCountryCode = "GB"
	addenda18B.SequenceNumber = 2
	addenda18B.EntryDetailSequenceNumber = 0000001
	mockBatch.Entries[0].Addenda10 = mockAddenda10()
	mockBatch.Entries[0].Addenda11 = mockAddenda11()
	mockBatch.Entries[0].Addenda12 = mockAddenda12()
	mockBatch.Entries[0].Addenda13 = mockAddenda13()
	mockBatch.Entries[0].Addenda14 = mockAddenda14()
	mockBatch.Entries[0].Addenda15 = mockAddenda15()
	mockBatch.Entries[0].Addenda16 = mockAddenda16()
	mockBatch.Entries[0].AddAddenda17(addenda17)
	mockBatch.Entries[0].AddAddenda17(addenda17B)
	mockBatch.Entries[0].AddAddenda18(addenda18)
	mockBatch.Entries[0].AddAddenda18(addenda18B)

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	addenda18B.SequenceNumber = -1
	err := mockBatch.verify()
	if !base.Match(err, NewErrBatchAscending("-1", "1")) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchAddenda18Sequence tests validating IATBatch Addenda18 Sequence Number error
func TestIATBatchAddenda18Sequence(t *testing.T) {
	testIATBatchAddenda18Sequence(t)
}

// BenchmarkIATBatchAddenda18Sequence benchmarks validating IATBatch Addenda18 Sequence Number error
func BenchmarkIATBatchAddenda18Sequence(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda18Sequence(b)
	}
}

// testIATNoEntry validates error for no entries
func testIATNoEntry(t testing.TB) {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATBatchHeaderFF())
	err := mockBatch.verify()
	if !base.Match(err, ErrBatchNoEntries) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATNoEntry tests validating error for no entries
func TestIATNoEntry(t *testing.T) {
	testIATNoEntry(t)
}

// BenchmarkIATNoEntry benchmarks validating error for no entries
func BenchmarkIATNoEntry(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATNoEntry(b)
	}
}

// testIATBatchAddendumTypeCode validates IATBatch Addendum TypeCode
func testIATBatchAddendumTypeCode(t testing.TB) {
	mockBatch := mockIATBatch(t)
	mockBatch.GetEntries()[0].AddAddenda17(mockAddenda17())

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	err := mockBatch.Validate()
	// no error expected
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchAddendumTypeCode tests validating IATBatch Addendum TypeCode
func TestIATBatchAddendumTypeCode(t *testing.T) {
	testIATBatchAddendumTypeCode(t)
}

// BenchmarkIATBatchAddendumTypeCode benchmarks validating IATBatch Addendum TypeCode
func BenchmarkIATBatchAddendumTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddendumTypeCode(b)
	}

}

// testIATBatchAddenda17Count validates IATBatch Addenda17 Count
func testIATBatchAddenda17Count(t testing.TB) {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATBatchHeaderFF())
	mockBatch.AddEntry(mockIATEntryDetail())
	addenda17 := NewAddenda17()
	addenda17.PaymentRelatedInformation = "This is an international payment"
	addenda17.SequenceNumber = 1
	addenda17.EntryDetailSequenceNumber = 0000001
	addenda17B := NewAddenda17()
	addenda17B.PaymentRelatedInformation = "Transfer of money from one country to another"
	addenda17B.SequenceNumber = 2
	addenda17B.EntryDetailSequenceNumber = 0000001
	addenda17C := NewAddenda17()
	addenda17C.PaymentRelatedInformation = "Send money Internationally"
	addenda17C.SequenceNumber = 3
	addenda17C.EntryDetailSequenceNumber = 0000001
	mockBatch.Entries[0].Addenda10 = mockAddenda10()
	mockBatch.Entries[0].Addenda11 = mockAddenda11()
	mockBatch.Entries[0].Addenda12 = mockAddenda12()
	mockBatch.Entries[0].Addenda13 = mockAddenda13()
	mockBatch.Entries[0].Addenda14 = mockAddenda14()
	mockBatch.Entries[0].Addenda15 = mockAddenda15()
	mockBatch.Entries[0].Addenda16 = mockAddenda16()
	mockBatch.Entries[0].AddAddenda17(addenda17)
	mockBatch.Entries[0].AddAddenda17(addenda17B)
	mockBatch.Entries[0].AddAddenda17(addenda17C)

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchAddendaCount(3, 2)) {
		t.Errorf("%T: %s", err, err)
	}

}

// TestIATBatchAddenda17Count tests validating IATBatch Addenda17 Count
func TestIATBatchAddenda17Count(t *testing.T) {
	testIATBatchAddenda17Count(t)
}

// BenchmarkIATBatchAddenda17Count benchmarks validating IATBatch Addenda17 Count
func BenchmarkIATBatchAddenda17Count(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda17Count(b)
	}
}

// testIATBatchAddenda18Count validates IATBatch Addenda18 Count
func testIATBatchAddenda18Count(t testing.TB) {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATBatchHeaderFF())
	mockBatch.AddEntry(mockIATEntryDetail())
	mockBatch.Entries[0].Addenda10 = mockAddenda10()
	mockBatch.Entries[0].Addenda11 = mockAddenda11()
	mockBatch.Entries[0].Addenda12 = mockAddenda12()
	mockBatch.Entries[0].Addenda13 = mockAddenda13()
	mockBatch.Entries[0].Addenda14 = mockAddenda14()
	mockBatch.Entries[0].Addenda15 = mockAddenda15()
	mockBatch.Entries[0].Addenda16 = mockAddenda16()
	mockBatch.Entries[0].AddAddenda17(mockAddenda17())
	mockBatch.Entries[0].AddAddenda18(mockAddenda18())
	mockBatch.Entries[0].AddAddenda18(mockAddenda18B())
	mockBatch.Entries[0].AddAddenda18(mockAddenda18C())
	mockBatch.Entries[0].AddAddenda18(mockAddenda18D())
	mockBatch.Entries[0].AddAddenda18(mockAddenda18E())

	addenda18F := NewAddenda18()
	addenda18F.ForeignCorrespondentBankName = "Russian Federation Bank"
	addenda18F.ForeignCorrespondentBankIDNumberQualifier = "01"
	addenda18F.ForeignCorrespondentBankIDNumber = "123123456789943874"
	addenda18F.ForeignCorrespondentBankBranchCountryCode = "RU"
	addenda18F.SequenceNumber = 6
	addenda18F.EntryDetailSequenceNumber = 0000001

	mockBatch.Entries[0].AddAddenda18(mockAddenda18F())

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchAddendaCount(6, 5)) {
		t.Errorf("%T: %s", err, err)
	}

}

// TestIATBatchAddenda18Count tests validating IATBatch Addenda18 Count
func TestIATBatchAddenda18Count(t *testing.T) {
	testIATBatchAddenda18Count(t)
}

// BenchmarkIATBatchAddenda18Count benchmarks validating IATBatch Addenda18 Count
func BenchmarkIATBatchAddenda18Count(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda18Count(b)
	}
}

// testIATBatchBuildAddendaError validates IATBatch build Addenda error
func testIATBatchBuildAddendaError(t testing.TB) {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATBatchHeaderFF())
	mockBatch.AddEntry(mockIATEntryDetail())

	err := mockBatch.build()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchBuildAddendaError tests validating IATBatch build Addenda error
func TestIATBatchBuildAddendaError(t *testing.T) {
	testIATBatchBuildAddendaError(t)
}

// BenchmarkIATBatchBuildAddendaError benchmarks validating IATBatch build Addenda error
func BenchmarkIATBatchBuildAddendaError(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchBuildAddendaError(b)
	}

}

// testIATBatchBHODFI validates IATBatchHeader ODFI error
func testIATBatchBHODFI(t testing.TB) {
	mockBatch := mockIATBatch(t)
	mockBatch.GetEntries()[0].SetTraceNumber("39387337", 1)

	err := mockBatch.verify()
	if !base.Match(err, NewErrBatchTraceNumberNotODFI("23138010", "39387337")) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchBHODFI tests validating IATBatchHeader ODFI error
func TestIATBatchBHODFI(t *testing.T) {
	testIATBatchBHODFI(t)
}

// BenchmarkIATBatchBHODFI benchmarks validating IATBatchHeader ODFI error
func BenchmarkIATBatchBHODFI(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchBHODFI(b)
	}

}

// testIATBatchAddenda99Count validates IATBatch Addenda99 Count
func testIATBatchAddenda99Count(t testing.TB) {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATReturnBatchHeaderFF())
	mockBatch.AddEntry(mockIATEntryDetail())
	mockBatch.GetEntries()[0].Category = CategoryReturn
	mockBatch.Entries[0].Addenda10 = mockAddenda10()
	mockBatch.Entries[0].Addenda11 = mockAddenda11()
	mockBatch.Entries[0].Addenda12 = mockAddenda12()
	mockBatch.Entries[0].Addenda13 = mockAddenda13()
	mockBatch.Entries[0].Addenda14 = mockAddenda14()
	mockBatch.Entries[0].Addenda15 = mockAddenda15()
	mockBatch.Entries[0].Addenda16 = mockAddenda16()
	mockBatch.Entries[0].AddAddenda17(mockAddenda17())
	mockBatch.Entries[0].Addenda99 = mockIATAddenda99()
	mockBatch.category = CategoryReturn

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	err := mockBatch.Validate()
	// TODO: are we expecting there to be no errors here?
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}

}

// TestIATBatchAddenda99Count tests validating IATBatch Addenda99 Count
func TestIATBatchAddenda99Count(t *testing.T) {
	testIATBatchAddenda99Count(t)
}

// BenchmarkIATBatchAddenda99Count benchmarks validating IATBatch Addenda99 Count
func BenchmarkIATBatchAddenda99Count(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testIATBatchAddenda99Count(b)
	}
}

// TestIATBatchAddenda98TotalCount validates IATBatch Addenda98 TotalCount
func TestIATBatchAddenda98TotalCount(t *testing.T) {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATNOCBatchHeaderFF())
	mockBatch.AddEntry(mockIATEntryDetail())
	mockBatch.GetEntries()[0].TransactionCode = CheckingReturnNOCCredit
	mockBatch.GetEntries()[0].AddendaRecords = 2
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.Entries[0].Addenda98 = mockIATAddenda98()
	mockBatch.category = CategoryNOC

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	err := mockBatch.Validate()
	// TODO: are we expecting there to be no errors here?
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchAddenda98Nil validates IATBatch Addenda98 is not nil
func TestIATBatchAddenda98Nil(t *testing.T) {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATNOCBatchHeaderFF())
	mockBatch.AddEntry(mockIATEntryDetail())
	mockBatch.GetEntries()[0].TransactionCode = CheckingReturnNOCCredit
	mockBatch.GetEntries()[0].AddendaRecords = 2
	mockBatch.GetEntries()[0].Category = CategoryNOC

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	err := mockBatch.Validate()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchAddenda98RecordType validates IATBatch Addenda98 RecordType is valid
func TestIATBatchAddenda98RecordType(t *testing.T) {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATNOCBatchHeaderFF())
	mockBatch.AddEntry(mockIATEntryDetail())
	mockBatch.GetEntries()[0].TransactionCode = CheckingReturnNOCCredit
	mockBatch.GetEntries()[0].AddendaRecords = 2
	addenda98 := mockAddenda98()
	addenda98.TypeCode = "00"
	mockBatch.GetEntries()[0].Addenda98 = addenda98
	mockBatch.GetEntries()[0].Category = CategoryNOC

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	err := mockBatch.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchAddenda99RecordType validates IATBatch Addenda99 RecordType is valid
func TestIATBatchAddenda99RecordType(t *testing.T) {
	mockBatch := mockIATBatch(t)
	mockBatch.SetHeader(mockIATReturnBatchHeaderFF())
	mockBatch.GetEntries()[0].AddendaRecords = 1
	addenda99 := mockAddenda99()
	addenda99.TypeCode = "00"
	mockBatch.GetEntries()[0].Addenda99 = addenda99
	mockBatch.GetEntries()[0].Category = CategoryReturn

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	err := mockBatch.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchAddenda18RecordType validates IATBatch Addenda18 RecordType is valid
func TestIATBatchAddenda18RecordType(t *testing.T) {
	mockBatch := mockIATBatch(t)
	mockBatch.SetHeader(mockIATReturnBatchHeaderFF())
	addenda18 := mockAddenda18()
	addenda18.TypeCode = "00"
	mockBatch.GetEntries()[0].AddAddenda18(addenda18)

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	err := mockBatch.Validate()
	if !base.Match(err, ErrAddendaTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchAddenda98TransactionCode validates IATBatch Transaction Code Count
func TestIATBatchAddenda98TransactionCode(t *testing.T) {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATNOCBatchHeaderFF())
	mockBatch.AddEntry(mockIATEntryDetail())
	mockBatch.GetEntries()[0].TransactionCode = CheckingReturnNOCCredit
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.Entries[0].Addenda98 = mockIATAddenda98()
	mockBatch.category = CategoryNOC

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	err := mockBatch.Validate()
	// TODO: are we expecting there to be no errors here?
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestIATBatchAddenda98IATIndicator validates IATBatch Transaction Code Count
func TestIATBatchAddenda98IATIndicator(t *testing.T) {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATNOCBatchHeaderFF())
	mockBatch.GetHeader().IATIndicator = "B"
	mockBatch.AddEntry(mockIATEntryDetail())
	mockBatch.GetEntries()[0].TransactionCode = CheckingReturnNOCCredit
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.Entries[0].Addenda98 = mockIATAddenda98()
	mockBatch.category = CategoryNOC

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchIATNOC("B", "IATCOR")) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestIATBatchAddenda98SECCode(t *testing.T) {
	mockBatch := IATBatch{}
	mockBatch.SetHeader(mockIATNOCBatchHeaderFF())
	mockBatch.GetHeader().StandardEntryClassCode = IAT
	mockBatch.AddEntry(mockIATEntryDetail())
	mockBatch.GetEntries()[0].TransactionCode = CheckingReturnNOCCredit
	mockBatch.GetEntries()[0].Category = CategoryNOC
	mockBatch.Entries[0].Addenda98 = mockIATAddenda98()
	mockBatch.category = CategoryNOC

	if err := mockBatch.build(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchIATNOC("IAT", "COR")) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestMockISODestinationCountryCode validates mockIATBatch
func TestMockISODestinationCountryCode(t *testing.T) {
	iatBatch := mockIATBatch(t)
	iatBatch.Header.ISODestinationCountryCode = "®©"
	err := iatBatch.Validate()
	if !base.Match(err, ErrValidISO3166) {
		t.Errorf("%T: %s", err, err)
	}

}

// TestMockISOOriginatingCurrencyCode validates mockIATBatch
func TestMockISOOriginatingCurrencyCode(t *testing.T) {
	iatBatch := mockIATBatch(t)
	iatBatch.Header.ISOOriginatingCurrencyCode = "®©"
	err := iatBatch.Validate()
	if !base.Match(err, ErrValidISO4217) {
		t.Errorf("%T: %s", err, err)
	}

}

// TestMockISODestinationCurrencyCode validates mockIATBatch
func TestMockISODestinationCurrencyCode(t *testing.T) {
	iatBatch := mockIATBatch(t)
	iatBatch.Header.ISODestinationCurrencyCode = "®©"
	err := iatBatch.Validate()
	if !base.Match(err, ErrValidISO4217) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestParseRuneCountIATBatchHeader tests parsing an invalid RuneCount in IATBatchHeader
func TestParseRuneCountIATBatchHeader(t *testing.T) {
	line := "5220                FF3               US123456789 IATTRADEPAYMTCADUSD010101"
	r := NewReader(strings.NewReader(line))
	r.line = line
	err := r.parseIATBatchHeader()
	if !base.Match(err, ErrFieldInclusion) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestIATBatch_calculateEntryHash(t *testing.T) {
	b := &IATBatch{}
	b.SetHeader(mockIATBatchHeaderFF())

	ed1 := mockIATEntryDetailWithAddendas()
	ed1.RDFIIdentification = "05600507"
	b.AddEntry(ed1)

	ed2 := mockIATEntryDetailWithAddendas()
	ed2.RDFIIdentification = "05140225"
	b.AddEntry(ed2)

	ed3 := mockIATEntryDetailWithAddendas()
	ed3.RDFIIdentification = "11400065"
	b.AddEntry(ed3)

	require.NoError(t, b.build())

	hash := b.calculateEntryHash()
	require.Equal(t, 22140797, hash)
}

func TestDiscussion1077(t *testing.T) {
	b := IATBatch{}
	b.SetHeader(mockIATBatchHeaderFF())

	ed1 := mockIATEntryDetailWithAddendas()
	ed1.RDFIIdentification = "101201863"
	b.AddEntry(ed1)

	ed2 := mockIATEntryDetailWithAddendas()
	ed2.RDFIIdentification = "274973183"
	b.AddEntry(ed2)

	ed3 := mockIATEntryDetailWithAddendas()
	ed3.RDFIIdentification = "111900659"
	b.AddEntry(ed3)

	ed4 := mockIATEntryDetailWithAddendas()
	ed4.RDFIIdentification = "101918101"
	b.AddEntry(ed4)

	ed5 := mockIATEntryDetailWithAddendas()
	ed5.RDFIIdentification = "044000037"
	b.AddEntry(ed5)

	require.NoError(t, b.build())
	require.Equal(t, 63399382, b.calculateEntryHash())
}
