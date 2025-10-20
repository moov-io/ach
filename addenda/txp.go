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
	// TaxAmount is the tax amount in cents
	TaxAmount string
	// TaxpayerVerification is the verification information
	TaxpayerVerification string
}

// ErrInvalidTXPFormat is returned when the TXP format is invalid
var ErrInvalidTXPFormat = errors.New("invalid TXP format")

// ErrInvalidTXPCharacter is returned when TXP contains invalid characters
var ErrInvalidTXPCharacter = errors.New("invalid TXP character")

// ParseTXP parses a TXP-formatted PaymentRelatedInformation string
// Expected format: TXP*tax identification number*tax payment type code*date*tax information ID number*tax amount*****taxpayer verification\
func ParseTXP(paymentInfo string) (*TXP, error) {
	if paymentInfo == "" {
		return nil, ErrInvalidTXPFormat
	}

	// Must start with "TXP*"
	if !strings.HasPrefix(paymentInfo, "TXP*") {
		return nil, ErrInvalidTXPFormat
	}

	// Must end with backslash terminator
	if !strings.HasSuffix(strings.TrimSpace(paymentInfo), "\\") {
		return nil, ErrInvalidTXPFormat
	}

	// Remove the "TXP*" prefix and "\" suffix for parsing
	content := strings.TrimSpace(paymentInfo[4:]) // Remove "TXP*"
	content = strings.TrimSuffix(content, "\\")   // Remove trailing "\"

	// Split by asterisk delimiter
	parts := strings.Split(content, "*")
	if len(parts) < 10 {
		return nil, ErrInvalidTXPFormat
	}

	// Validate allowed characters
	if err := validateTXPCharacters(paymentInfo); err != nil {
		return nil, err
	}

	// The format is: tax_id*tax_type*date*tax_info_id*tax_amount*****verification
	// When split by "*", "*****" creates empty parts at indices 5,6,7,8
	// So verification is at index 9
	txp := &TXP{
		TaxIdentificationNumber: parts[0],
		TaxPaymentTypeCode:      parts[1],
		Date:                    parts[2],
		TaxInformationIDNumber:  parts[3],
		TaxAmount:               parts[4],
		TaxpayerVerification:    parts[9], // Skip empty parts 5,6,7,8 from "*****"
	}

	return txp, nil
}

// FormatTXP creates a TXP-formatted PaymentRelatedInformation string
func FormatTXP(txp *TXP) string {
	if txp == nil {
		return ""
	}

	// Format: TXP*tax identification number*tax payment type code*date*tax information ID number*tax amount*****taxpayer verification\
	return strings.Join([]string{
		"TXP",
		txp.TaxIdentificationNumber,
		txp.TaxPaymentTypeCode,
		txp.Date,
		txp.TaxInformationIDNumber,
		txp.TaxAmount,
		"",
		"",
		"",
		"",
		txp.TaxpayerVerification + "\\",
	}, "*")
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
