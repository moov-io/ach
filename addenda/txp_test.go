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

package addenda

import (
	"testing"
)

// TestParseTXP tests parsing valid TXP format strings
func TestParseTXP(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    *TXP
		expectError bool
	}{
		{
			name:  "Valid TXP format",
			input: "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\",
			expected: &TXP{
				TaxIdentificationNumber: "123456789",
				TaxPaymentTypeCode:      "FEDERAL",
				Date:                    "20231231",
				TaxInformationIDNumber:  "TAX2023Q4",
				TaxAmount:               "500000",
				TaxpayerVerification:    "VERIFIED",
			},
			expectError: false,
		},
		{
			name:  "Valid TXP with different values",
			input: "TXP*987654321*STATE*20231231*CA2023Q4*250000*****CONFIRMED\\",
			expected: &TXP{
				TaxIdentificationNumber: "987654321",
				TaxPaymentTypeCode:      "STATE",
				Date:                    "20231231",
				TaxInformationIDNumber:  "CA2023Q4",
				TaxAmount:               "250000",
				TaxpayerVerification:    "CONFIRMED",
			},
			expectError: false,
		},
		{
			name:        "Empty string",
			input:       "",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Missing TXP prefix",
			input:       "123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Missing backslash terminator",
			input:       "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Too few parts",
			input:       "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000\\",
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseTXP(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if result == nil {
				t.Errorf("expected result but got nil")
				return
			}

			if result.TaxIdentificationNumber != tt.expected.TaxIdentificationNumber {
				t.Errorf("TaxIdentificationNumber: expected %s, got %s", tt.expected.TaxIdentificationNumber, result.TaxIdentificationNumber)
			}
			if result.TaxPaymentTypeCode != tt.expected.TaxPaymentTypeCode {
				t.Errorf("TaxPaymentTypeCode: expected %s, got %s", tt.expected.TaxPaymentTypeCode, result.TaxPaymentTypeCode)
			}
			if result.Date != tt.expected.Date {
				t.Errorf("Date: expected %s, got %s", tt.expected.Date, result.Date)
			}
			if result.TaxInformationIDNumber != tt.expected.TaxInformationIDNumber {
				t.Errorf("TaxInformationIDNumber: expected %s, got %s", tt.expected.TaxInformationIDNumber, result.TaxInformationIDNumber)
			}
			if result.TaxAmount != tt.expected.TaxAmount {
				t.Errorf("TaxAmount: expected %s, got %s", tt.expected.TaxAmount, result.TaxAmount)
			}
			if result.TaxpayerVerification != tt.expected.TaxpayerVerification {
				t.Errorf("TaxpayerVerification: expected %s, got %s", tt.expected.TaxpayerVerification, result.TaxpayerVerification)
			}
		})
	}
}

// TestFormatTXP tests formatting TXP structs to strings
func TestFormatTXP(t *testing.T) {
	tests := []struct {
		name     string
		input    *TXP
		expected string
	}{
		{
			name: "Valid TXP struct",
			input: &TXP{
				TaxIdentificationNumber: "123456789",
				TaxPaymentTypeCode:      "FEDERAL",
				Date:                    "20231231",
				TaxInformationIDNumber:  "TAX2023Q4",
				TaxAmount:               "500000",
				TaxpayerVerification:    "VERIFIED",
			},
			expected: "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\",
		},
		{
			name:     "Nil TXP",
			input:    nil,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatTXP(tt.input)
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

// TestValidateTXPCharacters tests character validation
func TestValidateTXPCharacters(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
	}{
		{
			name:        "Valid characters",
			input:       "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\",
			expectError: false,
		},
		{
			name:        "Invalid @ symbol",
			input:       "TXP*123@456*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\",
			expectError: true,
		},
		{
			name:        "Invalid # symbol",
			input:       "TXP*123#456*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\",
			expectError: true,
		},
		{
			name:        "Invalid newline",
			input:       "TXP*123\n456*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\",
			expectError: true,
		},
		{
			name:        "Invalid tab",
			input:       "TXP*123\t456*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\",
			expectError: true,
		},
		{
			name:        "Valid hyphen",
			input:       "TXP*123-456*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\",
			expectError: false,
		},
		{
			name:        "Valid period",
			input:       "TXP*123.456*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\",
			expectError: false,
		},
		{
			name:        "Valid colon",
			input:       "TXP*123:456*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTXPCharacters(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			}
		})
	}
}

// TestIsTXPFormat tests format detection
func TestIsTXPFormat(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid TXP format",
			input:    "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\",
			expected: true,
		},
		{
			name:     "Invalid format - missing TXP prefix",
			input:    "123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\",
			expected: false,
		},
		{
			name:     "Invalid format - missing backslash",
			input:    "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED",
			expected: false,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsTXPFormat(tt.input)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// TestTXPParseFormatRoundTrip tests that parsing and formatting are consistent
func TestTXPParseFormatRoundTrip(t *testing.T) {
	original := "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*****VERIFIED\\"

	// Parse the original string
	txp, err := ParseTXP(original)
	if err != nil {
		t.Fatalf("unexpected error parsing: %v", err)
	}

	// Format it back to string
	formatted := FormatTXP(txp)
	if formatted != original {
		t.Errorf("round trip failed: expected %s, got %s", original, formatted)
	}
}
