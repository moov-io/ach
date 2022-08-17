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
	"strings"
	"testing"

	"github.com/moov-io/base"
)

// mockADVEntryDetail creates a ADV entry detail
func mockADVEntryDetail() *ADVEntryDetail {
	entry := NewADVEntryDetail()
	entry.TransactionCode = CreditForDebitsOriginated
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 50000
	entry.AdviceRoutingNumber = "121042882"
	entry.FileIdentification = "11131"
	entry.ACHOperatorData = ""
	entry.IndividualName = "Name"
	entry.DiscretionaryData = ""
	entry.AddendaRecordIndicator = 0
	entry.ACHOperatorRoutingNumber = "01100001"
	entry.JulianDay = 50
	entry.SequenceNumber = 1
	return entry
}

// testMockADVEntryDetail validates an ADV entry detail record
func testMockADVEntryDetail(t testing.TB) {
	entry := mockADVEntryDetail()
	if err := entry.Validate(); err != nil {
		t.Error("mockADVEntryDetail does not validate and will break other tests")
	}
	if entry.TransactionCode != CreditForDebitsOriginated {
		t.Error("TransactionCode dependent default value has changed")
	}
	if entry.RDFIIdentification != "23138010" {
		t.Error("RDFIIdentification dependent default value has changed")
	}
	if entry.AdviceRoutingNumber != "121042882" {
		t.Error("AdviceRoutingNumber dependent default value has changed")
	}
	if entry.DFIAccountNumber != "744-5678-99" {
		t.Error("DFIAccountNumber dependent default value has changed")
	}
	if entry.FileIdentification != "11131" {
		t.Error("FileIdentification dependent default value has changed")
	}
	if entry.Amount != 50000 {
		t.Error("Amount dependent default value has changed")
	}
	if entry.IndividualName != "Name" {
		t.Error("IndividualName dependent default value has changed")
	}
	if entry.ACHOperatorRoutingNumber != "01100001" {
		t.Error("ACHOperatorRoutingNumber dependent default value has changed")
	}
	if entry.DiscretionaryData != "" {
		t.Error("DiscretionaryData dependent default value has changed")
	}
	if entry.JulianDay != 50 {
		t.Error("JulianDay dependent default value has changed")
	}
	if entry.SequenceNumber != 1 {
		t.Error("SequenceNumber dependent default value has changed")
	}

}

// TestMockADVEntryDetail tests validating an entry detail record
func TestMockADVEntryDetail(t *testing.T) {
	testMockADVEntryDetail(t)
}

// BenchmarkMockEntryDetail benchmarks validating an entry detail record
func BenchmarkMockADVEntryDetail(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockADVEntryDetail(b)
	}
}

// testADVEDString validates that a known parsed ADV entry
// detail can be returned to a string of the same value
func testADVEDString(t testing.TB) {
	var line = "681231380104744-5678-99    00000005000012104288211131 Name                    0011000010500001"
	r := NewReader(strings.NewReader(line))
	r.line = line
	bh := BatchHeader{BatchNumber: 1,
		StandardEntryClassCode: ADV,
		ServiceClassCode:       AutomatedAccountingAdvices,
		CompanyIdentification:  "origid",
		ODFIIdentification:     "121042882"}
	r.addCurrentBatch(NewBatchADV(&bh))

	r.currentBatch.AddADVEntry(mockADVEntryDetail())
	if err := r.parseEntryDetail(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentBatch.GetADVEntries()[0]

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestADVEDString tests validating that a known parsed entry
// detail can be returned to a string of the same value
func TestADVEDString(t *testing.T) {
	testADVEDString(t)
}

// BenchmarkADVEDString benchmarks validating that a known parsed entry
// detail can be returned to a string of the same value
func BenchmarkADVEDString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testADVEDString(b)
	}
}

// TestValidateADVEDTransactionCode validates error if transaction code is not valid
func TestValidateADVEDTransactionCode(t *testing.T) {
	ed := mockADVEntryDetail()
	ed.TransactionCode = 63
	err := ed.Validate()
	if !base.Match(err, ErrTransactionCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestADVEDDFIAccountNumberAlphaNumeric validates DFI account number is alpha numeric
func TestADVEDDFIAccountNumberAlphaNumeric(t *testing.T) {
	ed := mockADVEntryDetail()
	ed.DFIAccountNumber = "®"
	err := ed.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestADVEDAdviceRoutingNumberAlphaNumeric validates Advice Routing Number is alpha numeric
func TestADVEDAdviceRoutingNumberAlphaNumeric(t *testing.T) {
	ed := mockADVEntryDetail()
	ed.AdviceRoutingNumber = "®"
	err := ed.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestADVEDIndividualNameAlphaNumeric validates IndividualName is alpha numeric
func TestADVEDIndividualNameAlphaNumeric(t *testing.T) {
	ed := mockADVEntryDetail()
	ed.IndividualName = "®"
	err := ed.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestADVEDDiscretionaryDataAlphaNumeric validates DiscretionaryData is alpha numeric
func TestADVEDDiscretionaryDataAlphaNumeric(t *testing.T) {
	ed := mockADVEntryDetail()
	ed.DiscretionaryData = "®"
	err := ed.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestADVEDACHOperatorRoutingNumberAlphaNumeric validates ACHOperatorRoutingNumber is alpha numeric
func TestADVEDACHOperatorRoutingNumberAlphaNumeric(t *testing.T) {
	ed := mockADVEntryDetail()
	ed.ACHOperatorRoutingNumber = "®"
	err := ed.Validate()
	if !base.Match(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestADVEDisCheckDigit validates check digit
func TestADVEDisCheckDigit(t *testing.T) {
	ed := mockADVEntryDetail()
	ed.CheckDigit = "1"
	err := ed.Validate()
	if !base.Match(err, NewErrValidCheckDigit(7)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestADVEDFieldInclusionTransactionCode validates TransactionCode field inclusion
func TestADVEDFieldInclusionTransactionCode(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.TransactionCode = 0
	err := entry.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestADVEDFieldInclusionRDFIIdentification validates RDFIIdentification field inclusion
func TestADVEDFieldInclusionRDFIIdentification(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.RDFIIdentification = ""
	err := entry.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestADVEDFieldInclusionDFIAccountNumber validates DFIAccountNumber field inclusion
func TestADVEDFieldInclusionDFIAccountNumber(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.DFIAccountNumber = ""
	err := entry.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestADVEDFieldInclusionAdviceRoutingNumber validates AdviceRoutingNumber field inclusion
func TestADVEDFieldInclusionAdviceRoutingNumber(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.AdviceRoutingNumber = ""
	err := entry.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestADVEDFieldInclusionIndividualName validates IndividualName field inclusion
func TestADVEDFieldInclusionIndividualName(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.IndividualName = ""
	err := entry.Validate()
	if !base.Match(err, ErrFieldRequired) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestADVEDFieldInclusionACHOperatorRoutingNumber validates ACHOperatorRoutingNumber field inclusion
func TestADVEDFieldInclusionACHOperatorRoutingNumber(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.ACHOperatorRoutingNumber = ""
	err := entry.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestADVEDFieldInclusionJulianDateDay validates JulianDay field inclusion
func TestADVEDFieldInclusionJulianDateDay(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.JulianDay = 0
	err := entry.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestADVEDFieldInclusionSequenceNumber validates SequenceNumber field inclusion
func TestADVEDFieldInclusionSequenceNumber(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.SequenceNumber = 0
	err := entry.Validate()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestADVEDBadTransactionCode validates TransactionCode field inclusion
func TestBadTransactionCode(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.TransactionCode = CheckingDebit
	err := entry.Validate()
	// TODO: are we expecting there to be no errors here?
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestInvalidADVEDParse returns an error when parsing an ADV Entry Detail
func TestInvalidADVEDParse(t *testing.T) {
	var line = "681231380104744-5678-99    000000050000121042882FILE1 Name"
	r := NewReader(strings.NewReader(line))
	r.line = line
	bh := BatchHeader{BatchNumber: 1,
		StandardEntryClassCode: ADV,
		ServiceClassCode:       AutomatedAccountingAdvices,
		CompanyIdentification:  "origid",
		ODFIIdentification:     "121042882"}
	r.addCurrentBatch(NewBatchADV(&bh))

	err := r.parseEntryDetail()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestADVUnmarshal(t *testing.T) {
	raw := []byte(`{
          "id": "adv-01",
          "transactionCode": 81,
          "RDFIIdentification": "23138010",
          "checkDigit": "4",
          "DFIAccountNumber": "81967038518      ",
          "amount": 100000,
          "adviceRoutingNumber": "121042882",
          "fileIdentification": "11131",
          "achOperatorData": "",
          "individualName": "Steven Tander         ",
          "discretionaryData": "  ",
          "addendaRecordIndicator": 0,
          "achOperatorRoutingNumber": "01100001",
          "julianDay": 2,
          "sequenceNumber": 1
        }`)

	var ed ADVEntryDetail
	if err := json.Unmarshal(raw, &ed); err != nil {
		t.Fatal(err)
	}

	if ed.JulianDay != 2 {
		t.Errorf("ed.JulianDay=%d", ed.JulianDay)
	}
}
