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
			name:  "Valid TXP with 3 amounts and verification (NACHA Example 1)",
			input: "TXP*123456789*606*960331*T*100000*P*12000*I*4567*SML2A\\",
			expected: &TXP{
				TaxIdentificationNumber: "123456789",
				TaxPaymentTypeCode:      "606",
				Date:                    "960331",
				TaxAmounts: []TaxAmount{
					{AmountType: "T", AmountCents: "100000"},
					{AmountType: "P", AmountCents: "12000"},
					{AmountType: "I", AmountCents: "4567"},
				},
				TaxpayerVerification: "SML2A",
			},
			expectError: false,
		},
		{
			name:  "Valid TXP with 1 amount and verification (NACHA Example 2)",
			input: "TXP*12345678934*526*960930*T*100000***SML2A\\",
			expected: &TXP{
				TaxIdentificationNumber: "12345678934",
				TaxPaymentTypeCode:      "526",
				Date:                    "960930",
				TaxAmounts: []TaxAmount{
					{AmountType: "T", AmountCents: "100000"},
				},
				TaxpayerVerification: "SML2A",
			},
			expectError: false,
		},
		{
			name:  "Valid TXP with 3 amounts and no verification (NACHA Example 3)",
			input: "TXP*123456789*94105*960301*1*10000*2*5000*3*15000\\",
			expected: &TXP{
				TaxIdentificationNumber: "123456789",
				TaxPaymentTypeCode:      "94105",
				Date:                    "960301",
				TaxAmounts: []TaxAmount{
					{AmountType: "1", AmountCents: "10000"},
					{AmountType: "2", AmountCents: "5000"},
					{AmountType: "3", AmountCents: "15000"},
				},
				TaxpayerVerification: "",
			},
			expectError: false,
		},
		{
			name:  "Valid TXP with 2 amounts and *** delimiter",
			input: "TXP*123456789*941*250901*1*1000*2*500***VERIFIED\\",
			expected: &TXP{
				TaxIdentificationNumber: "123456789",
				TaxPaymentTypeCode:      "941",
				Date:                    "250901",
				TaxAmounts: []TaxAmount{
					{AmountType: "1", AmountCents: "1000"},
					{AmountType: "2", AmountCents: "500"},
				},
				TaxpayerVerification: "VERIFIED",
			},
			expectError: false,
		},
		{
			name:  "Valid TXP with 1 amount and no verification",
			input: "TXP*123456789*941*250901*T*100000\\",
			expected: &TXP{
				TaxIdentificationNumber: "123456789",
				TaxPaymentTypeCode:      "941",
				Date:                    "250901",
				TaxAmounts: []TaxAmount{
					{AmountType: "T", AmountCents: "100000"},
				},
				TaxpayerVerification: "",
			},
			expectError: false,
		},
		{
			name:  "Valid TXP with 2 amounts and no verification",
			input: "TXP*123456789*941*250901*1*1000*2*500\\",
			expected: &TXP{
				TaxIdentificationNumber: "123456789",
				TaxPaymentTypeCode:      "941",
				Date:                    "250901",
				TaxAmounts: []TaxAmount{
					{AmountType: "1", AmountCents: "1000"},
					{AmountType: "2", AmountCents: "500"},
				},
				TaxpayerVerification: "",
			},
			expectError: false,
		},
		{
			name:  "Valid TXP with mixed amount types",
			input: "TXP*888999000*STATE*20231231*S*200000*2*10000**CONFIRMED\\",
			expected: &TXP{
				TaxIdentificationNumber: "888999000",
				TaxPaymentTypeCode:      "STATE",
				Date:                    "20231231",
				TaxAmounts: []TaxAmount{
					{AmountType: "S", AmountCents: "200000"},
					{AmountType: "2", AmountCents: "10000"},
				},
				TaxpayerVerification: "CONFIRMED",
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
			input:       "123456789*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED\\",
			expected:    nil,
			expectError: true,
		},
		{
			name:  "Valid TXP without backslash terminator",
			input: "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED",
			expected: &TXP{
				TaxIdentificationNumber: "123456789",
				TaxPaymentTypeCode:      "FEDERAL",
				Date:                    "20231231",
				TaxAmounts: []TaxAmount{
					{AmountType: "TAX2023Q4", AmountCents: "500000"},
				},
				TaxpayerVerification: "VERIFIED",
			},
			expectError: false,
		},
		{
			name:        "Too few parts",
			input:       "TXP*123456789*FEDERAL*20231231*TAX2023Q4\\",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Odd number of amount parts",
			input:       "TXP*123456789*FEDERAL*20231231*500000*1*25000*VERIFIED\\",
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
			if len(result.TaxAmounts) != len(tt.expected.TaxAmounts) {
				t.Errorf("TaxAmounts length: expected %d, got %d", len(tt.expected.TaxAmounts), len(result.TaxAmounts))
			} else {
				for i, expectedAmount := range tt.expected.TaxAmounts {
					if result.TaxAmounts[i].AmountCents != expectedAmount.AmountCents {
						t.Errorf("TaxAmounts[%d].AmountCents: expected %s, got %s", i, expectedAmount.AmountCents, result.TaxAmounts[i].AmountCents)
					}
					if result.TaxAmounts[i].AmountType != expectedAmount.AmountType {
						t.Errorf("TaxAmounts[%d].AmountType: expected %s, got %s", i, expectedAmount.AmountType, result.TaxAmounts[i].AmountType)
					}
				}
			}
			if result.TaxpayerVerification != tt.expected.TaxpayerVerification {
				t.Errorf("TaxpayerVerification: expected %s, got %s", tt.expected.TaxpayerVerification, result.TaxpayerVerification)
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
			input:       "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*1*25000*2**VERIFIED\\",
			expectError: false,
		},
		{
			name:        "Invalid @ symbol",
			input:       "TXP*123@456*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED\\",
			expectError: true,
		},
		{
			name:        "Invalid # symbol",
			input:       "TXP*123#456*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED\\",
			expectError: true,
		},
		{
			name:        "Invalid newline",
			input:       "TXP*123\n456*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED\\",
			expectError: true,
		},
		{
			name:        "Invalid tab",
			input:       "TXP*123\t456*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED\\",
			expectError: true,
		},
		{
			name:        "Valid hyphen",
			input:       "TXP*123-456*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED\\",
			expectError: false,
		},
		{
			name:        "Valid period",
			input:       "TXP*123.456*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED\\",
			expectError: false,
		},
		{
			name:        "Valid colon",
			input:       "TXP*123:456*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED\\",
			expectError: false,
		},
		{
			name:        "Valid integer amount type",
			input:       "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*1*25000*2**VERIFIED\\",
			expectError: false,
		},
		{
			name:        "Valid single letter amount type",
			input:       "TXP*123456789*STATE*20231231*TAX2023Q4*500000*S**VERIFIED\\",
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
			input:    "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*1*25000*2**VERIFIED\\",
			expected: true,
		},
		{
			name:     "Invalid format - missing TXP prefix",
			input:    "123456789*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED\\",
			expected: false,
		},
		{
			name:     "Valid format - without backslash",
			input:    "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED",
			expected: true,
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
