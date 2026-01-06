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

	"github.com/stretchr/testify/require"
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
		{
			name:        "TXP exceeds 80 byte limit",
			input:       "TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*1*25000*2*10000*VERIFIED_WITH_VERY_LONG_VERIFICATION_STRING\\",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Empty amount type in pair",
			input:       "TXP*123456789*941*250901**1000\\",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Empty amount cents in pair",
			input:       "TXP*123456789*941*250901*T*\\",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Empty TaxIdentificationNumber",
			input:       "TXP**941*250901*T*1000\\",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Empty TaxPaymentTypeCode",
			input:       "TXP*123456789**250901*T*1000\\",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Empty Date",
			input:       "TXP*123456789*941**T*1000\\",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Date too short",
			input:       "TXP*123456789*941*2509*T*1000\\",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Date too long",
			input:       "TXP*123456789*941*2509011*T*1000\\",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Date non-numeric",
			input:       "TXP*123456789*941*ABCDEF*T*1000\\",
			expected:    nil,
			expectError: true,
		},
		{
			name:  "Multiple consecutive empty delimiters",
			input: "TXP*123456789*941*250901*T*1000****VERIFIED\\",
			expected: &TXP{
				TaxIdentificationNumber: "123456789",
				TaxPaymentTypeCode:      "941",
				Date:                    "250901",
				TaxAmounts: []TaxAmount{
					{AmountType: "T", AmountCents: "1000"},
				},
				TaxpayerVerification: "VERIFIED",
			},
			expectError: false,
		},
		{
			name:  "Trailing asterisk without verification",
			input: "TXP*123456789*941*250901*T*1000*\\",
			expected: &TXP{
				TaxIdentificationNumber: "123456789",
				TaxPaymentTypeCode:      "941",
				Date:                    "250901",
				TaxAmounts: []TaxAmount{
					{AmountType: "T", AmountCents: "1000"},
				},
				TaxpayerVerification: "",
			},
			expectError: false,
		},
		{
			name:        "Non-numeric amount with spaces",
			input:       "TXP*123456789*941*250901*T* 1000 \\",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "Case-insensitive TXP prefix",
			input:       "txp*123456789*941*250901*T*1000\\",
			expected:    nil,
			expectError: true,
		},
		{
			name:  "Single asterisk after amounts",
			input: "TXP*123456789*941*250901*T*1000*VERIFIED\\",
			expected: &TXP{
				TaxIdentificationNumber: "123456789",
				TaxPaymentTypeCode:      "941",
				Date:                    "250901",
				TaxAmounts: []TaxAmount{
					{AmountType: "T", AmountCents: "1000"},
				},
				TaxpayerVerification: "VERIFIED",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseTXP(tt.input)

			if tt.expectError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, result)

			require.Equal(t, tt.expected.TaxIdentificationNumber, result.TaxIdentificationNumber)
			require.Equal(t, tt.expected.TaxPaymentTypeCode, result.TaxPaymentTypeCode)
			require.Equal(t, tt.expected.Date, result.Date)
			require.Len(t, result.TaxAmounts, len(tt.expected.TaxAmounts))
			for i, expectedAmount := range tt.expected.TaxAmounts {
				require.Equal(t, expectedAmount.AmountCents, result.TaxAmounts[i].AmountCents)
				require.Equal(t, expectedAmount.AmountType, result.TaxAmounts[i].AmountType)
			}
			require.Equal(t, tt.expected.TaxpayerVerification, result.TaxpayerVerification)
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
				require.Error(t, err)
			} else {
				require.NoError(t, err)
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
			require.Equal(t, tt.expected, result)
		})
	}
}

// TestTXPToString tests serialization of TXP objects
func TestTXPToString(t *testing.T) {
	tests := []struct {
		name     string
		txp      *TXP
		expected string
	}{
		{
			name: "TXP with 3 amounts and verification",
			txp: &TXP{
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
			expected: "TXP*123456789*606*960331*T*100000*P*12000*I*4567*SML2A\\",
		},
		{
			name: "TXP with 1 amount and verification",
			txp: &TXP{
				TaxIdentificationNumber: "12345678934",
				TaxPaymentTypeCode:      "526",
				Date:                    "960930",
				TaxAmounts: []TaxAmount{
					{AmountType: "T", AmountCents: "100000"},
				},
				TaxpayerVerification: "SML2A",
			},
			expected: "TXP*12345678934*526*960930*T*100000*SML2A\\",
		},
		{
			name: "TXP with 3 amounts and no verification",
			txp: &TXP{
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
			expected: "TXP*123456789*94105*960301*1*10000*2*5000*3*15000\\",
		},
		{
			name: "TXP with 2 amounts and no verification",
			txp: &TXP{
				TaxIdentificationNumber: "123456789",
				TaxPaymentTypeCode:      "941",
				Date:                    "250901",
				TaxAmounts: []TaxAmount{
					{AmountType: "1", AmountCents: "1000"},
					{AmountType: "2", AmountCents: "500"},
				},
				TaxpayerVerification: "",
			},
			expected: "TXP*123456789*941*250901*1*1000*2*500\\",
		},
		{
			name: "TXP with mixed amount types",
			txp: &TXP{
				TaxIdentificationNumber: "888999000",
				TaxPaymentTypeCode:      "STATE",
				Date:                    "20231231",
				TaxAmounts: []TaxAmount{
					{AmountType: "S", AmountCents: "200000"},
					{AmountType: "2", AmountCents: "10000"},
				},
				TaxpayerVerification: "CONFIRMED",
			},
			expected: "TXP*888999000*STATE*20231231*S*200000*2*10000*CONFIRMED\\",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.txp.String()
			require.Equal(t, tt.expected, result)
		})
	}
}

// TestTXPToStringRoundTrip tests that parsing and serializing produces consistent results
func TestTXPToStringRoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "TXP with 3 amounts and verification",
			input: "TXP*123456789*606*960331*T*100000*P*12000*I*4567*SML2A\\",
		},
		{
			name:  "TXP with 1 amount and verification",
			input: "TXP*12345678934*526*960930*T*100000***SML2A\\",
		},
		{
			name:  "TXP with 3 amounts and no verification",
			input: "TXP*123456789*94105*960301*1*10000*2*5000*3*15000\\",
		},
		{
			name:  "TXP with 2 amounts and no verification",
			input: "TXP*123456789*941*250901*1*1000*2*500\\",
		},
		{
			name:  "TXP with 1 amount and no verification",
			input: "TXP*123456789*941*250901*T*100000\\",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parse the input
			txp, err := ParseTXP(tt.input)
			require.NoError(t, err, "failed to parse TXP")

			// Serialize it back
			serialized := txp.String()

			// Parse the serialized string again
			reparsed, err := ParseTXP(serialized)
			require.NoError(t, err, "failed to reparse serialized TXP")

			// Verify the round-trip results match
			require.Equal(t, txp.TaxIdentificationNumber, reparsed.TaxIdentificationNumber)
			require.Equal(t, txp.TaxPaymentTypeCode, reparsed.TaxPaymentTypeCode)
			require.Equal(t, txp.Date, reparsed.Date)
			require.Len(t, reparsed.TaxAmounts, len(txp.TaxAmounts))
			for i, expectedAmount := range txp.TaxAmounts {
				require.Equal(t, expectedAmount.AmountCents, reparsed.TaxAmounts[i].AmountCents)
				require.Equal(t, expectedAmount.AmountType, reparsed.TaxAmounts[i].AmountType)
			}
			require.Equal(t, txp.TaxpayerVerification, reparsed.TaxpayerVerification)
		})
	}
}
