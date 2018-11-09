// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "testing"

// mockEntryDetailADV creates a ADV entry detail
func mockEntryDetailADV() *EntryDetailADV {
	entry := NewEntryDetailADV()
	entry.TransactionCode = 81
	entry.SetRDFI("231380104")
	entry.AdviceRoutingNumber = "121042882"
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 50000
	entry.IndividualName = "Name"
	entry.ACHOperatorRoutingNumber = "01100001"
	entry.DiscretionaryData = ""
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
	if entry.RDFIIdentification != "231380104" {
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
