// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
	"testing"
)

// mockADVEntryDetail creates a ADV entry detail
func mockADVEntryDetail() *ADVEntryDetail {
	entry := NewADVEntryDetail()
	entry.TransactionCode = 81
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
	entry.JulianDateDay = 50
	entry.SequenceNumber = 1
	return entry
}

// testMockADVEntryDetail validates an ADV entry detail record
func testMockADVEntryDetail(t testing.TB) {
	entry := mockADVEntryDetail()
	if err := entry.Validate(); err != nil {
		t.Error("mockADVEntryDetail does not validate and will break other tests")
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
	if entry.JulianDateDay != 50 {
		t.Error("JulianDateDay dependent default value has changed")
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
		StandardEntryClassCode: "ADV",
		ServiceClassCode:       280,
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

// TestValidateADVEDRecordType validates error if recordType is not 6
func TestValidateADVEDRecordType(t *testing.T) {
	ed := mockADVEntryDetail()
	ed.recordType = "2"
	if err := ed.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "recordType" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestValidateADVEDTransactionCode validates error if transaction code is not valid
func TestValidateADVEDTransactionCode(t *testing.T) {
	ed := mockADVEntryDetail()
	ed.TransactionCode = 63
	if err := ed.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TransactionCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVEDDFIAccountNumberAlphaNumeric validates DFI account number is alpha numeric
func TestADVEDDFIAccountNumberAlphaNumeric(t *testing.T) {
	ed := mockADVEntryDetail()
	ed.DFIAccountNumber = "®"
	if err := ed.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "DFIAccountNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVEDAdviceRoutingNumberAlphaNumeric validates Advice Routing Number is alpha numeric
func TestADVEDAdviceRoutingNumberAlphaNumeric(t *testing.T) {
	ed := mockADVEntryDetail()
	ed.AdviceRoutingNumber = "®"
	if err := ed.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "AdviceRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVEDIndividualNameAlphaNumeric validates IndividualName is alpha numeric
func TestADVEDIndividualNameAlphaNumeric(t *testing.T) {
	ed := mockADVEntryDetail()
	ed.IndividualName = "®"
	if err := ed.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "IndividualName" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVEDDiscretionaryDataAlphaNumeric validates DiscretionaryData is alpha numeric
func TestADVEDDiscretionaryDataAlphaNumeric(t *testing.T) {
	ed := mockADVEntryDetail()
	ed.DiscretionaryData = "®"
	if err := ed.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "DiscretionaryData" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVEDACHOperatorRoutingNumberAlphaNumeric validates ACHOperatorRoutingNumber is alpha numeric
func TestADVEDACHOperatorRoutingNumberAlphaNumeric(t *testing.T) {
	ed := mockADVEntryDetail()
	ed.ACHOperatorRoutingNumber = "®"
	if err := ed.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "ACHOperatorRoutingNumber" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVEDisCheckDigit validates check digit
func TestADVEDisCheckDigit(t *testing.T) {
	ed := mockADVEntryDetail()
	ed.CheckDigit = "1"
	if err := ed.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "RDFIIdentification" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVEDFieldInclusionRecordType validates recordType field inclusion
func TestADVEDFieldInclusionRecordType(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.recordType = ""
	if err := entry.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVEDFieldInclusionTransactionCode validates TransactionCode field inclusion
func TestADVEDFieldInclusionTransactionCode(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.TransactionCode = 0
	if err := entry.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVEDFieldInclusionRDFIIdentification validates RDFIIdentification field inclusion
func TestADVEDFieldInclusionRDFIIdentification(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.RDFIIdentification = ""
	if err := entry.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVEDFieldInclusionDFIAccountNumber validates DFIAccountNumber field inclusion
func TestADVEDFieldInclusionDFIAccountNumber(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.DFIAccountNumber = ""
	if err := entry.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVEDFieldInclusionAdviceRoutingNumber validates AdviceRoutingNumber field inclusion
func TestADVEDFieldInclusionAdviceRoutingNumber(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.AdviceRoutingNumber = ""
	if err := entry.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVEDFieldInclusionIndividualName validates IndividualName field inclusion
func TestADVEDFieldInclusionIndividualName(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.IndividualName = ""
	if err := entry.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldRequired) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVEDFieldInclusionACHOperatorRoutingNumber validates ACHOperatorRoutingNumber field inclusion
func TestADVEDFieldInclusionACHOperatorRoutingNumber(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.ACHOperatorRoutingNumber = ""
	if err := entry.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVEDFieldInclusionJulianDateDay validates JulianDateDay field inclusion
func TestADVEDFieldInclusionJulianDateDay(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.JulianDateDay = 0
	if err := entry.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVEDFieldInclusionSequenceNumber validates SequenceNumber field inclusion
func TestADVEDFieldInclusionSequenceNumber(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.SequenceNumber = 0
	if err := entry.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestADVEDBadTransactionCode validates TransactionCode field inclusion
func TestBadTransactionCode(t *testing.T) {
	entry := mockADVEntryDetail()
	entry.TransactionCode = 1
	if err := entry.Validate(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.FieldName != "TransactionCode" {
				t.Errorf("%T: %s", err, err)
			}
		}
	}
}

// TestInvalidADVEDParse returns an error when parsing an ADV Entry Detail
func TestInvalidADVEDParse(t *testing.T) {
	var line = "681231380104744-5678-99    000000050000121042882FILE1 Name"
	r := NewReader(strings.NewReader(line))
	r.line = line
	bh := BatchHeader{BatchNumber: 1,
		StandardEntryClassCode: "ADV",
		ServiceClassCode:       280,
		CompanyIdentification:  "origid",
		ODFIIdentification:     "121042882"}
	r.addCurrentBatch(NewBatchADV(&bh))

	if err := r.parseEntryDetail(); err != nil {
		if p, ok := err.(*ParseError); ok {
			if p.Record != "EntryDetail" {
				t.Errorf("%T: %s", p, p)
			}
		} else {
			t.Errorf("%T: %s", p.Err, p.Err)
		}
	}
}
