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

	"github.com/stretchr/testify/require"
)

// mockBatchDNEHeader creates a DNE batch header
func mockBatchDNEHeader() *BatchHeader {
	bh := NewBatchHeader()
	bh.ServiceClassCode = CreditsOnly
	bh.CompanyName = "Name on Account"
	bh.CompanyIdentification = "231380104"
	bh.StandardEntryClassCode = DNE
	bh.CompanyEntryDescription = "Death"
	bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1).Format("060102") // YYMMDD
	bh.ODFIIdentification = "23138010"
	bh.OriginatorStatusCode = 2
	return bh
}

// mockDNEEntryDetail creates a DNE entry detail
func mockDNEEntryDetail() *EntryDetail {
	entry := NewEntryDetail()
	entry.TransactionCode = CheckingPrenoteCredit
	entry.SetRDFI("031300012")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 0
	entry.SetOriginalTraceNumber("031300010000001")
	entry.SetReceivingCompany("Best. #1")
	entry.SetTraceNumber("23138010", 1)
	entry.AddendaRecordIndicator = 1

	addenda := NewAddenda05()
	addenda.PaymentRelatedInformation = `    DATE OF DEATH*010218*CUSTOMERSSN*#########*AMOUNT*$$$$.cc\`
	entry.AddAddenda05(addenda)

	return entry
}

// mockBatchDNE creates a DNE batch
func mockBatchDNE(t testing.TB) *BatchDNE {
	t.Helper()
	batch := NewBatchDNE(mockBatchDNEHeader())
	batch.AddEntry(mockDNEEntryDetail())
	if err := batch.Create(); err != nil {
		t.Fatalf("Unexpected error building batch: %s\n", err)
	}
	return batch
}

// testBatchDNEHeader creates a DNE batch header
func testBatchDNEHeader(t testing.TB) {
	batch, _ := NewBatch(mockBatchDNEHeader())
	_, ok := batch.(*BatchDNE)
	if !ok {
		t.Error("Expecting BatchDNE")
	}
}

// TestBatchDNEHeader tests creating a DNE batch header
func TestBatchDNEHeader(t *testing.T) {
	testBatchDNEHeader(t)
}

// BenchmarkBatchDNEHeader benchmark creating a DNE batch header
func BenchmarkBatchDNEHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchDNEHeader(b)
	}
}

// testBatchDNEAddendumCount batch control DNE can only have one addendum per entry detail
func testBatchDNEAddendumCount(t testing.TB) {
	mockBatch := mockBatchDNE(t)
	// Adding a second addenda to the mock entry
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	err := mockBatch.Validate()
	if !base.Match(err, NewErrBatchCalculatedControlEquality(3, 2)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchDNEAddendumCount tests batch control DNE can only have one addendum per entry detail
func TestBatchDNEAddendumCount(t *testing.T) {
	testBatchDNEAddendumCount(t)
}

// BenchmarkBatchDNEAddendumCount benchmarks batch control DNE can only have one addendum per entry detail
func BenchmarkBatchDNEAddendumCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchDNEAddendumCount(b)
	}
}

// TestBatchDNEAddendum98 validates Addenda05 returns an error
func TestBatchDNEAddendum98(t *testing.T) {
	mockBatch := NewBatchDNE(mockBatchDNEHeader())
	mockBatch.AddEntry(mockDNEEntryDetail())
	err := mockBatch.Create()
	// TODO: are we expecting there to be an error here?
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// testBatchDNEReceivingCompanyName validates Receiving company / Individual name is a mandatory field
func testBatchDNEReceivingCompanyName(t testing.TB) {
	mockBatch := mockBatchDNE(t)
	// modify the Individual name / receiving company to nothing
	mockBatch.GetEntries()[0].SetReceivingCompany("")
	err := mockBatch.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchDNEReceivingCompanyName tests validating receiving company / Individual name is a mandatory field
func TestBatchDNEReceivingCompanyName(t *testing.T) {
	testBatchDNEReceivingCompanyName(t)
}

// BenchmarkBatchDNEReceivingCompanyName benchmarks validating receiving company / Individual name is a mandatory field
func BenchmarkBatchDNEReceivingCompanyName(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchDNEReceivingCompanyName(b)
	}
}

// testBatchDNEAddendaTypeCode validates addenda type code is 05
func testBatchDNEAddendaTypeCode(t testing.TB) {
	mockBatch := mockBatchDNE(t)
	mockBatch.GetEntries()[0].Addenda05[0].TypeCode = "05"
	err := mockBatch.Validate()
	// no error expected
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchDNEAddendaTypeCode tests validating addenda type code is 05
func TestBatchDNEAddendaTypeCode(t *testing.T) {
	testBatchDNEAddendaTypeCode(t)
}

// BenchmarkBatchDNEAddendaTypeCod benchmarks validating addenda type code is 05
func BenchmarkBatchDNEAddendaTypeCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchDNEAddendaTypeCode(b)
	}
}

// testBatchDNESEC validates that the standard entry class code is DNE for batchDNE
func testBatchDNESEC(t testing.TB) {
	mockBatch := mockBatchDNE(t)
	mockBatch.Header.StandardEntryClassCode = ACK
	err := mockBatch.Validate()
	if !base.Match(err, ErrBatchSECType) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchDNESEC tests validating that the standard entry class code is DNE for batchDNE
func TestBatchDNESEC(t *testing.T) {
	testBatchDNESEC(t)
}

// BenchmarkBatchDNESEC benchmarks validating that the standard entry class code is DNE for batch DNE
func BenchmarkBatchDNESEC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchDNESEC(b)
	}
}

// testBatchDNEAddendaCount validates batch DNE addenda count
func testBatchDNEAddendaCount(t testing.TB) {
	mockBatch := mockBatchDNE(t)
	mockBatch.GetEntries()[0].AddAddenda05(mockAddenda05())
	err := mockBatch.Create()
	if !base.Match(err, NewErrBatchAddendaCount(0, 1)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchDNEAddendaCount tests validating batch DNE addenda count
func TestBatchDNEAddendaCount(t *testing.T) {
	testBatchDNEAddendaCount(t)
}

// BenchmarkBatchDNEAddendaCount benchmarks validating batch DNE addenda count
func BenchmarkBatchDNEAddendaCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchDNEAddendaCount(b)
	}
}

// testBatchDNEServiceClassCode validates ServiceClassCode
func testBatchDNEServiceClassCode(t testing.TB) {
	mockBatch := mockBatchDNE(t)
	// Batch Header information is required to Create a batch.
	mockBatch.GetHeader().ServiceClassCode = 0
	err := mockBatch.Create()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchDNEServiceClassCode tests validating ServiceClassCode
func TestBatchDNEServiceClassCode(t *testing.T) {
	testBatchDNEServiceClassCode(t)
}

// BenchmarkBatchDNEServiceClassCode benchmarks validating ServiceClassCode
func BenchmarkBatchDNEServiceClassCode(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testBatchDNEServiceClassCode(b)
	}
}

// TestBatchDNEAmount validates Amount
func TestBatchDNEAmount(t *testing.T) {
	mockBatch := mockBatchDNE(t)
	// Batch Header information is required to Create a batch.
	mockBatch.GetEntries()[0].Amount = 25000
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchAmountNonZero) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchDNETransactionCode validates TransactionCode
func TestBatchDNETransactionCode(t *testing.T) {
	mockBatch := mockBatchDNE(t)
	mockBatch.GetEntries()[0].TransactionCode = CheckingCredit
	err := mockBatch.Create()
	if !base.Match(err, ErrBatchTransactionCode) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestBatchDNE__ParseDNEPaymentInformation(t *testing.T) {
	mockBatch := mockBatchDNE(t)

	details, err := ParseDNEPaymentInformation(mockBatch.Entries[0].Addenda05[0])
	require.NoError(t, err)

	require.Equal(t, "010218", details.DateOfDeath.Format("010206"))
	require.Equal(t, "#########", details.CustomerSSN)
	require.Equal(t, "$$$$.cc", details.Amount)
}

func TestBatchDNE__nil(t *testing.T) {
	details, err := ParseDNEPaymentInformation(nil)
	require.Nil(t, details)
	require.NoError(t, err)

	t.Run("panic", func(t *testing.T) {
		batch := mockBatchDNE(t)
		batch.Entries[0].Addenda05[0].PaymentRelatedInformation = ""

		details, err := ParseDNEPaymentInformation(batch.Entries[0].Addenda05[0])
		require.ErrorContains(t, err, "unexpected 1 fields")
		require.Nil(t, details)
	})
}

func TestDNRPaymentInformation(t *testing.T) {
	info := DNEPaymentInformation{
		DateOfDeath: time.Date(2024, time.August, 28, 10, 30, 0, 0, time.UTC),
		CustomerSSN: "333224444",
		Amount:      "123.45",
	}
	expected := `DATE OF DEATH*082824*CUSTOMER SSN*333224444*AMOUNT*123.45\`
	require.Equal(t, expected, info.String())
}
