// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strings"
	"testing"
)

// mockEntryDetailADV creates a ADV entry detail
func mockEntryDetailADV() *EntryDetailADV {
	entry := NewEntryDetailADV()
	entry.TransactionCode = 81
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 50000
	entry.AdviceRoutingNumber = "121042882"
	entry.FileIdentification = "FILE1"
	entry.ACHOperatorData = ""
	entry.IndividualName = "Name"
	entry.DiscretionaryData = ""
	entry.AddendaRecordIndicator = 0
	entry.ACHOperatorRoutingNumber = "01100001"
	entry.JulianDateDay = 50
	entry.SequenceNumber = 1
	return entry
}

// testMockEntryDetailADV validates an ADV entry detail record
func testMockEntryDetailADV(t testing.TB) {
	entry := mockEntryDetailADV()
	if err := entry.Validate(); err != nil {
		t.Error("mockEntryDetailADV does not validate and will break other tests")
	}
	if entry.TransactionCode != 81 {
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
}

// TestMockEntryDetailADV tests validating an entry detail record
func TestMockEntryDetailADV(t *testing.T) {
	testMockEntryDetailADV(t)
}

// BenchmarkMockEntryDetail benchmarks validating an entry detail record
func BenchmarkMockEntryDetailADV(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testMockEntryDetailADV(b)
	}
}

// testEDAVString validates that a known parsed entry
// detail can be returned to a string of the same value
func testEDADVString(t testing.TB) {
	var line = "681231380104744-5678-99    000000050000121042882FILE1 Name                    0011000010500001"
	r := NewReader(strings.NewReader(line))
	r.line = line
	bh := BatchHeader{BatchNumber: 1,
		StandardEntryClassCode: "ADV",
		ServiceClassCode:       280,
		CompanyIdentification:  "origid",
		ODFIIdentification:     "121042882"}
	r.addCurrentBatch(NewBatchADV(&bh))

	r.currentBatch.AddADVEntry(mockEntryDetailADV())
	if err := r.parseEntryDetail(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	record := r.currentBatch.GetADVEntries()[0]

	fmt.Printf("line: %v \n", line)
	fmt.Printf("stri: %v \n", record.String())

	if record.String() != line {
		t.Errorf("Strings do not match")
	}
}

// TestEDADVString tests validating that a known parsed entry
// detail can be returned to a string of the same value
func TestEDADVString(t *testing.T) {
	testEDADVString(t)
}

// BenchmarkEDADVString benchmarks validating that a known parsed entry
// detail can be returned to a string of the same value
func BenchmarkEDADVString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testEDADVString(b)
	}
}
