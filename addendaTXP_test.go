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

	"github.com/moov-io/base"
)

// mockAddendaTXP creates a mock AddendaTXP record
func mockAddendaTXP() *AddendaTXP {
	addendaTXP := NewAddendaTXP()
	addendaTXP.PaymentRelatedInformation = "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\"
	addendaTXP.SequenceNumber = 1
	addendaTXP.EntryDetailSequenceNumber = 0000001
	return addendaTXP
}

// TestMockAddendaTXP validates mockAddendaTXP
func TestMockAddendaTXP(t *testing.T) {
	addendaTXP := mockAddendaTXP()
	if err := addendaTXP.Validate(); err != nil {
		t.Error("mockAddendaTXP does not validate and will break other tests")
	}
}

// testAddendaTXPParse parses AddendaTXP record
func testAddendaTXPParse(t testing.TB) {
	AddendaTXP := NewAddendaTXP()
	line := "705TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\                   00010000001"
	AddendaTXP.Parse(line)
	// walk the AddendaTXP struct
	if AddendaTXP.TypeCode != "05" {
		t.Errorf("expected %v got %v", "05", AddendaTXP.TypeCode)
	}
	if AddendaTXP.PaymentRelatedInformation != "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\" {
		t.Errorf("expected %v got %v", "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\", AddendaTXP.PaymentRelatedInformation)
	}
	if AddendaTXP.SequenceNumber != 1 {
		t.Errorf("expected %v got %v", 1, AddendaTXP.SequenceNumber)
	}
	if AddendaTXP.EntryDetailSequenceNumber != 1 {
		t.Errorf("expected %v got %v", 1, AddendaTXP.EntryDetailSequenceNumber)
	}
}

// TestAddendaTXPParse tests parsing AddendaTXP record
func TestAddendaTXPParse(t *testing.T) {
	testAddendaTXPParse(t)
}

// BenchmarkAddendaTXPParse benchmarks AddendaTXP parsing
func BenchmarkAddendaTXPParse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddendaTXPParse(b)
	}
}

// testAddendaTXPString validates that a known parsed AddendaTXP record can be written to a string
func testAddendaTXPString(t testing.TB) {
	AddendaTXP := NewAddendaTXP()
	line := "705TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\                   00010000001"
	AddendaTXP.Parse(line)

	if AddendaTXP.String() != line {
		t.Errorf("expected %v got %v", line, AddendaTXP.String())
	}
}

// TestAddendaTXPString tests AddendaTXP string
func TestAddendaTXPString(t *testing.T) {
	testAddendaTXPString(t)
}

// BenchmarkAddendaTXPString benchmarks AddendaTXP string
func BenchmarkAddendaTXPString(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddendaTXPString(b)
	}
}

// testAddendaTXPValidate validates AddendaTXP record
func testAddendaTXPValidate(t testing.TB) {
	AddendaTXP := NewAddendaTXP()
	AddendaTXP.PaymentRelatedInformation = "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\"
	AddendaTXP.SequenceNumber = 1
	AddendaTXP.EntryDetailSequenceNumber = 1

	if err := AddendaTXP.Validate(); err != nil {
		t.Errorf("AddendaTXP Validate() Error: %v", err)
	}
}

// TestAddendaTXPValidate tests AddendaTXP validation
func TestAddendaTXPValidate(t *testing.T) {
	testAddendaTXPValidate(t)
}

// BenchmarkAddendaTXPValidate benchmarks AddendaTXP validation
func BenchmarkAddendaTXPValidate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddendaTXPValidate(b)
	}
}

// testAddendaTXPFieldInclusion validates AddendaTXP field inclusion
func testAddendaTXPFieldInclusion(t testing.TB) {
	AddendaTXP := NewAddendaTXP()
	AddendaTXP.TypeCode = ""
	AddendaTXP.PaymentRelatedInformation = "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\"
	AddendaTXP.SequenceNumber = 1
	AddendaTXP.EntryDetailSequenceNumber = 1

	if err := AddendaTXP.Validate(); err != nil {
		if !base.Match(err, ErrConstructor) {
			t.Errorf("AddendaTXP Validate() Error: %v", err)
		}
	}
}

// TestAddendaTXPFieldInclusion tests AddendaTXP field inclusion
func TestAddendaTXPFieldInclusion(t *testing.T) {
	testAddendaTXPFieldInclusion(t)
}

// testAddendaTXPTypeCode validates AddendaTXP type code
func testAddendaTXPTypeCode(t testing.TB) {
	AddendaTXP := NewAddendaTXP()
	AddendaTXP.TypeCode = "99"
	AddendaTXP.PaymentRelatedInformation = "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\"
	AddendaTXP.SequenceNumber = 1
	AddendaTXP.EntryDetailSequenceNumber = 1

	if err := AddendaTXP.Validate(); err != nil {
		if !base.Match(err, ErrAddendaTypeCode) {
			t.Errorf("AddendaTXP Validate() Error: %v", err)
		}
	}
}

// TestAddendaTXPTypeCode tests AddendaTXP type code
func TestAddendaTXPTypeCode(t *testing.T) {
	testAddendaTXPTypeCode(t)
}

// testAddendaTXPPaymentRelatedInformation validates AddendaTXP payment related information
func testAddendaTXPPaymentRelatedInformation(t testing.TB) {
	AddendaTXP := NewAddendaTXP()
	AddendaTXP.TypeCode = "05"
	AddendaTXP.PaymentRelatedInformation = "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\"
	AddendaTXP.SequenceNumber = 1
	AddendaTXP.EntryDetailSequenceNumber = 1

	if err := AddendaTXP.Validate(); err != nil {
		t.Errorf("AddendaTXP Validate() Error: %v", err)
	}
}

// TestAddendaTXPPaymentRelatedInformation tests AddendaTXP payment related information
func TestAddendaTXPPaymentRelatedInformation(t *testing.T) {
	testAddendaTXPPaymentRelatedInformation(t)
}

// TestAddendaTXPPaymentRelatedInformationInvalidChars tests TXP allowed characters validation
func TestAddendaTXPPaymentRelatedInformationInvalidChars(t *testing.T) {
	addendaTXP := NewAddendaTXP()
	addendaTXP.TypeCode = "05"
	addendaTXP.SequenceNumber = 1
	addendaTXP.EntryDetailSequenceNumber = 1

	// Test invalid characters that should trigger ErrVariableFields (format check) or ErrInvalidProperty (character check)
	invalidChars := []string{
		"TXP*123@456*FEDERAL",  // @ symbol not allowed
		"TXP*123#456*FEDERAL",  // # symbol not allowed
		"TXP*123$456*FEDERAL",  // $ symbol not allowed
		"TXP*123%456*FEDERAL",  // % symbol not allowed
		"TXP*123^456*FEDERAL",  // ^ symbol not allowed
		"TXP*123&456*FEDERAL",  // & symbol not allowed
		"TXP*123(456*FEDERAL",  // ( symbol not allowed
		"TXP*123)456*FEDERAL",  // ) symbol not allowed
		"TXP*123[456*FEDERAL",  // [ symbol not allowed
		"TXP*123]456*FEDERAL",  // ] symbol not allowed
		"TXP*123{456*FEDERAL",  // { symbol not allowed
		"TXP*123}456*FEDERAL",  // } symbol not allowed
		"TXP*123|456*FEDERAL",  // | symbol not allowed
		"TXP*123~456*FEDERAL",  // ~ symbol not allowed
		"TXP*123`456*FEDERAL",  // ` symbol not allowed
		"TXP*123!456*FEDERAL",  // ! symbol not allowed
		"TXP*123?456*FEDERAL",  // ? symbol not allowed
		"TXP*123<456*FEDERAL",  // < symbol not allowed
		"TXP*123=456*FEDERAL",  // = symbol not allowed
		"TXP*123+456*FEDERAL",  // + symbol not allowed
		"TXP*123,456*FEDERAL",  // , symbol not allowed
		"TXP*123;456*FEDERAL",  // ; symbol not allowed
		"TXP*123\"456*FEDERAL", // " symbol not allowed
		"TXP*123'456*FEDERAL",  // ' symbol not allowed
		"TXP*123\n456*FEDERAL", // newline not allowed
		"TXP*123\r456*FEDERAL", // carriage return not allowed
		"TXP*123\t456*FEDERAL", // tab not allowed
	}

	for _, invalidChar := range invalidChars {
		addendaTXP.PaymentRelatedInformation = invalidChar
		if err := addendaTXP.Validate(); err != nil {
			// The validation may return either ErrVariableFields (format check) or ErrInvalidProperty (character check)
			if !base.Match(err, ErrVariableFields) && !base.Match(err, ErrInvalidProperty) {
				t.Errorf("expected ErrVariableFields or ErrInvalidProperty for '%s', got %T: %s", invalidChar, err, err)
			}
		} else {
			t.Errorf("expected validation error for invalid character '%s', but validation passed", invalidChar)
		}
	}
}

// TestAddendaTXPPaymentRelatedInformationValidChars tests TXP allowed characters validation
func TestAddendaTXPPaymentRelatedInformationValidChars(t *testing.T) {
	addendaTXP := NewAddendaTXP()
	addendaTXP.TypeCode = "05"
	addendaTXP.SequenceNumber = 1
	addendaTXP.EntryDetailSequenceNumber = 1

	// Test valid characters that should pass validation
	validChars := []string{
		"TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\", // Standard TXP format
		"TXP*ABC123*STATE*20231231*CA2023Q4*250000*****VERIFIED\\",       // Mixed alphanumeric
		"TXP*123-456*LOCAL*20231231*NY2023Q4*100000*****VERIFIED\\",      // Hyphen allowed
		"TXP*123.456*COUNTY*20231231*TX2023Q4*75000*****VERIFIED\\",      // Period allowed
		"TXP*123/456*CITY*20231231*FL2023Q4*50000*****VERIFIED\\",        // Forward slash allowed
		"TXP*123:456*SPECIAL*20231231*WA2023Q4*30000*****VERIFIED\\",     // Colon allowed
		"TXP*123>456*PRIORITY*20231231*OR2023Q4*20000*****VERIFIED\\",    // Greater than allowed
		"TXP*123 456*SPACE*20231231*NV2023Q4*15000*****VERIFIED\\",       // Space allowed
		"TXP*123*456*ASTERISK*20231231*UT2023Q4*10000*****VERIFIED\\",    // Asterisk allowed
		"TXP*123\\456*BACKSLASH*20231231*CO2023Q4*5000*****VERIFIED\\",   // Backslash allowed
	}

	for _, validChar := range validChars {
		addendaTXP.PaymentRelatedInformation = validChar
		if err := addendaTXP.Validate(); err != nil {
			t.Errorf("unexpected validation error for valid character string '%s': %v", validChar, err)
		}
	}
}

// testAddendaTXPSequenceNumber validates AddendaTXP sequence number
func testAddendaTXPSequenceNumber(t testing.TB) {
	AddendaTXP := NewAddendaTXP()
	AddendaTXP.TypeCode = "05"
	AddendaTXP.PaymentRelatedInformation = "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\"
	AddendaTXP.SequenceNumber = 0
	AddendaTXP.EntryDetailSequenceNumber = 1

	if err := AddendaTXP.Validate(); err != nil {
		if !base.Match(err, ErrConstructor) {
			t.Errorf("AddendaTXP Validate() Error: %v", err)
		}
	}
}

// TestAddendaTXPSequenceNumber tests AddendaTXP sequence number
func TestAddendaTXPSequenceNumber(t *testing.T) {
	testAddendaTXPSequenceNumber(t)
}

// testAddendaTXPEntryDetailSequenceNumber validates AddendaTXP entry detail sequence number
func testAddendaTXPEntryDetailSequenceNumber(t testing.TB) {
	AddendaTXP := NewAddendaTXP()
	AddendaTXP.TypeCode = "05"
	AddendaTXP.PaymentRelatedInformation = "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\"
	AddendaTXP.SequenceNumber = 1
	AddendaTXP.EntryDetailSequenceNumber = -1

	if err := AddendaTXP.Validate(); err != nil {
		if !base.Match(err, ErrConstructor) {
			t.Errorf("AddendaTXP Validate() Error: %v", err)
		}
	}
}

// TestAddendaTXPEntryDetailSequenceNumber tests AddendaTXP entry detail sequence number
func TestAddendaTXPEntryDetailSequenceNumber(t *testing.T) {
	testAddendaTXPEntryDetailSequenceNumber(t)
}

// testAddendaTXPPaymentRelatedInformationField validates AddendaTXP payment related information field
func testAddendaTXPPaymentRelatedInformationField(t testing.TB) {
	AddendaTXP := NewAddendaTXP()
	AddendaTXP.PaymentRelatedInformation = "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\"
	expected := "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\                   "
	actual := AddendaTXP.PaymentRelatedInformationField()

	if actual != expected {
		t.Errorf("expected %v got %v", expected, actual)
	}
}

// TestAddendaTXPPaymentRelatedInformationField tests AddendaTXP payment related information field
func TestAddendaTXPPaymentRelatedInformationField(t *testing.T) {
	testAddendaTXPPaymentRelatedInformationField(t)
}

// testAddendaTXPSequenceNumberField validates AddendaTXP sequence number field
func testAddendaTXPSequenceNumberField(t testing.TB) {
	AddendaTXP := NewAddendaTXP()
	AddendaTXP.SequenceNumber = 1
	expected := "0001"
	actual := AddendaTXP.SequenceNumberField()

	if actual != expected {
		t.Errorf("expected %v got %v", expected, actual)
	}
}

// TestAddendaTXPSequenceNumberField tests AddendaTXP sequence number field
func TestAddendaTXPSequenceNumberField(t *testing.T) {
	testAddendaTXPSequenceNumberField(t)
}

// testAddendaTXPEntryDetailSequenceNumberField validates AddendaTXP entry detail sequence number field
func testAddendaTXPEntryDetailSequenceNumberField(t testing.TB) {
	AddendaTXP := NewAddendaTXP()
	AddendaTXP.EntryDetailSequenceNumber = 1
	expected := "0000001"
	actual := AddendaTXP.EntryDetailSequenceNumberField()

	if actual != expected {
		t.Errorf("expected %v got %v", expected, actual)
	}
}

// TestAddendaTXPEntryDetailSequenceNumberField tests AddendaTXP entry detail sequence number field
func TestAddendaTXPEntryDetailSequenceNumberField(t *testing.T) {
	testAddendaTXPEntryDetailSequenceNumberField(t)
}

// testAddendaTXPRecordLength validates AddendaTXP record length
func testAddendaTXPRecordLength(t testing.TB) {
	AddendaTXP := NewAddendaTXP()
	line := "705TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\                   00010000001"
	AddendaTXP.Parse(line)

	if len(AddendaTXP.String()) != 94 {
		t.Errorf("expected %v got %v", 94, len(AddendaTXP.String()))
	}
}

// TestAddendaTXPRecordLength tests AddendaTXP record length
func TestAddendaTXPRecordLength(t *testing.T) {
	testAddendaTXPRecordLength(t)
}

// BenchmarkAddendaTXPRecordLength benchmarks AddendaTXP record length
func BenchmarkAddendaTXPRecordLength(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testAddendaTXPRecordLength(b)
	}
}
