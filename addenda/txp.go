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
	"errors"
	"strings"
)

// TaxAmount represents a single tax amount with its type
type TaxAmount struct {
	// AmountCents is the amount in cents
	AmountCents string
	// AmountType is the tax information type ID for the tax amount
	// Can be an integer or letter (e.g., "1", "2", "3", "S", "P", "I", etc.)
	AmountType string
}

// TXP represents a Tax Payment addenda parsed from PaymentRelatedInformation
// of an Addenda05 record. TXP is not a NACHA standard addenda type, but rather
// a specific format for tax payment information within existing Addenda05 records.
type TXP struct {
	// TaxIdentificationNumber is the taxpayer's identification number
	TaxIdentificationNumber string
	// TaxPaymentTypeCode indicates the type of tax payment (e.g., FEDERAL, STATE, LOCAL)
	TaxPaymentTypeCode string
	// Date represents the tax period or payment date in YYYYMMDD format
	Date string
	// TaxInformationIDNumber is a unique identifier for the tax information
	TaxInformationIDNumber string
	// TaxAmounts is a slice of tax amounts with their types
	TaxAmounts []TaxAmount
	// TaxpayerVerification is the verification information
	TaxpayerVerification string
}

// isNumeric checks if a string contains only numeric characters
func isNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// ErrInvalidTXPCharacter is returned when TXP contains invalid characters
var ErrInvalidTXPCharacter = errors.New("invalid TXP character")

// ErrInvalidTXPFormat is returned when the TXP format is invalid
var ErrInvalidTXPFormat = errors.New("invalid TXP format")

// ParseTXP parses a TXP-formatted PaymentRelatedInformation string
// Expected format: TXP*tax identification number*tax payment type code*date*type1*amount1*type2*amount2*type3*amount3*taxpayer verification\
// Note: The total TXP addenda should be limited to 80 bytes
func ParseTXP(paymentInfo string) (*TXP, error) {
	if paymentInfo == "" {
		return nil, ErrInvalidTXPFormat
	}

	// Must start with "TXP*"
	if !strings.HasPrefix(paymentInfo, "TXP*") {
		return nil, ErrInvalidTXPFormat
	}

	// Remove the "TXP*" prefix for parsing
	content := strings.TrimSpace(paymentInfo[4:])

	// Remove backslash terminator if present
	content = strings.TrimSuffix(content, "\\")

	// Split by asterisk delimiter
	parts := strings.Split(content, "*")
	if len(parts) < 5 { // Minimum: tax_id*tax_type*date*amount_type*amount_cents
		return nil, ErrInvalidTXPFormat
	}

	// Validate allowed characters
	if err := validateTXPCharacters(paymentInfo); err != nil {
		return nil, err
	}

	txp := &TXP{}
	txp.TaxIdentificationNumber = parts[0]
	txp.TaxPaymentTypeCode = parts[1]
	txp.Date = parts[2]

	// Parse amount pairs sequentially
	i := 3
	for i < len(parts) {
		// Check if we have at least 2 more parts for an amount pair
		if i+1 >= len(parts) {
			break
		}

		// Check if the current part is empty (indicates delimiter)
		if parts[i] == "" {
			// We've hit a delimiter, look for verification after empty parts
			j := i
			for j < len(parts) && parts[j] == "" {
				j++
			}

			// If there's a non-empty part after the empty parts, it's verification
			if j < len(parts) {
				txp.TaxpayerVerification = parts[j]
			}
			break
		}

		// Check if the next part is empty (indicates we're at the end of amount pairs)
		if parts[i+1] == "" {
			// We have a type but no amount, this indicates we've hit a delimiter
			// Look for verification after the empty parts
			j := i + 1
			for j < len(parts) && parts[j] == "" {
				j++
			}

			// If there's a non-empty part after the empty parts, it's verification
			if j < len(parts) {
				txp.TaxpayerVerification = parts[j]
			}
			break
		}

		// We have a valid amount pair
		taxAmount := TaxAmount{
			AmountType:  parts[i],
			AmountCents: parts[i+1],
		}

		// Validate that amount cents is numeric
		if !isNumeric(taxAmount.AmountCents) {
			return nil, ErrInvalidTXPFormat
		}

		txp.TaxAmounts = append(txp.TaxAmounts, taxAmount)
		i += 2
	}

	// If we didn't find verification through delimiter detection,
	// check if there's a single remaining part after parsing amount pairs
	// This indicates the last part is verification (for cases like 3 amounts with verification)
	if txp.TaxpayerVerification == "" && i < len(parts) {
		// We have a single remaining part, so it's likely verification
		txp.TaxpayerVerification = parts[i]
	}

	// Validate that we have at least one amount
	if len(txp.TaxAmounts) == 0 {
		return nil, ErrInvalidTXPFormat
	}

	return txp, nil
}

// FormatTXP creates a TXP-formatted PaymentRelatedInformation string
func FormatTXP(txp *TXP) string {
	if txp == nil {
		return ""
	}

	// Start with basic fields
	parts := []string{
		"TXP",
		txp.TaxIdentificationNumber,
		txp.TaxPaymentTypeCode,
		txp.Date,
	}

	// Add tax amounts (pairs of type*amount)
	for _, taxAmount := range txp.TaxAmounts {
		parts = append(parts, taxAmount.AmountType, taxAmount.AmountCents)
	}

	// Join parts with asterisk delimiter
	result := strings.Join(parts, "*")

	// Add verification field if present, following NACHA pattern:
	// - 3 amount pairs: verification comes directly after (no delimiter)
	// - 2 amount pairs: verification comes after *** delimiter
	// - 1 amount pair: verification comes after *** delimiter
	if txp.TaxpayerVerification != "" {
		if len(txp.TaxAmounts) >= 3 {
			// Three or more amounts: no delimiter
			result += "*" + txp.TaxpayerVerification
		} else {
			// One or two amounts: use *** delimiter
			result += "***" + txp.TaxpayerVerification
		}
	}

	// Add backslash terminator
	result += "\\"

	return result
}

// validateTXPCharacters ensures PaymentRelatedInformation only contains characters
// permitted by TXP addenda conventions (printable set and delimiters).
func validateTXPCharacters(s string) error {
	for _, r := range s {
		if r == '\n' || r == '\r' || r == '\t' {
			return ErrInvalidTXPCharacter
		}
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			continue
		}
		switch r {
		case ' ', '*', '\\', '>', '-', '.', '/', ':':
			continue
		default:
			return ErrInvalidTXPCharacter
		}
	}
	return nil
}

// IsTXPFormat checks if a PaymentRelatedInformation string follows TXP format
func IsTXPFormat(paymentInfo string) bool {
	_, err := ParseTXP(paymentInfo)
	return err == nil
}
