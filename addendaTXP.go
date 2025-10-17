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
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// AddendaTXP is an addenda which provides tax payment information for Addenda Type
// Code 05 in a machine readable format. It is used for TXP (Tax Payment) transactions
// with StandardEntryClassCode: CCD+ for domestic tax payments.
type AddendaTXP struct {
	// ID is an identifier only used by the moov-io/ach HTTP server as a way to identify a batch.
	ID string `json:"id"`
	// TypeCode AddendaTXP type code '05'
	TypeCode string `json:"typeCode"`
	// PaymentRelatedInformation contains tax payment information
	// This field contains structured tax payment data such as:
	// - Taxpayer identification
	// - Tax period information
	// - Tax authority details
	// - Payment type codes
	PaymentRelatedInformation string `json:"paymentRelatedInformation"`
	// SequenceNumber is consecutively assigned to each AddendaTXP Record following
	// an Entry Detail Record. The first addendaTXP sequence number must always
	// be a "1".
	SequenceNumber int `json:"sequenceNumber"`
	// EntryDetailSequenceNumber contains the ascending sequence number section of the Entry
	// Detail or Corporate Entry Detail Record's trace number This number is
	// the same as the last seven digits of the trace number of the related
	// Entry Detail Record or Corporate Entry Detail Record.
	EntryDetailSequenceNumber int `json:"entryDetailSequenceNumber"`
	// Line number at which the record appears in the file
	LineNumber int `json:"lineNumber,omitempty"`
	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters
	// validateOpts defines optional overrides for record validation
	validateOpts *ValidateOpts
}

// NewAddendaTXP returns a new AddendaTXP with default values for none exported fields
func NewAddendaTXP() *AddendaTXP {
	addendaTXP := new(AddendaTXP)
	addendaTXP.TypeCode = "05"
	return addendaTXP
}

// Parse takes the input record string and parses the AddendaTXP values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addendaTXP *AddendaTXP) Parse(record string) {
	runeCount := utf8.RuneCountInString(record)
	if runeCount != 94 {
		return
	}

	buf := getBuffer()
	defer saveBuffer(buf)

	reset := func() string {
		out := buf.String()
		buf.Reset()
		return out
	}

	// We're going to process the record rune-by-rune and at each field cutoff save the value.
	var idx int
	for _, r := range record {
		idx++

		// Append rune to buffer
		buf.WriteRune(r)

		// At each cutoff save the buffer and reset
		switch idx {
		case 0, 1:
			// 1-1 Always 7
			reset()
		case 3:
			// 2-3 Always 05
			addendaTXP.TypeCode = reset()
		case 83:
			// 4-83 Based on the information entered (04-83) 80 alphanumeric
			addendaTXP.PaymentRelatedInformation = strings.TrimRight(reset(), " ")
		case 87:
			// 84-87 SequenceNumber is consecutively assigned to each AddendaTXP Record following
			// an Entry Detail Record
			addendaTXP.SequenceNumber = addendaTXP.parseNumField(reset())
		case 94:
			// 88-94 Contains the last seven digits of the number entered in the Trace Number field in the corresponding Entry Detail Record
			addendaTXP.EntryDetailSequenceNumber = addendaTXP.parseNumField(reset())
		}
	}
}

func (a *AddendaTXP) SetValidation(opts *ValidateOpts) {
	if a != nil {
		a.validateOpts = opts
	}
}

// String writes the AddendaTXP struct to a 80 character string.
func (addendaTXP *AddendaTXP) String() string {
	if addendaTXP == nil {
		return ""
	}

	buf := getBuffer()
	defer saveBuffer(buf)

	buf.WriteString(entryAddendaPos)
	buf.WriteString(addendaTXP.TypeCode)
	buf.WriteString(addendaTXP.PaymentRelatedInformationField())
	buf.WriteString(addendaTXP.SequenceNumberField())
	buf.WriteString(addendaTXP.EntryDetailSequenceNumberField())

	return buf.String()
}

// Validate performs NACHA format rule checks on the record and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (addendaTXP *AddendaTXP) Validate() error {
	if err := addendaTXP.fieldInclusion(); err != nil {
		return err
	}

	if err := addendaTXP.isTypeCode(addendaTXP.TypeCode); err != nil {
		return fieldError("TypeCode", err, addendaTXP.TypeCode)
	}
	// Type Code must be 05
	if addendaTXP.TypeCode != "05" {
		return fieldError("TypeCode", ErrAddendaTypeCode, addendaTXP.TypeCode)
	}

	// Enforce TXP segment framing and allowed characters per TXP rules
	pri := addendaTXP.PaymentRelatedInformation
	if len(pri) == 0 {
		return fieldError("PaymentRelatedInformation", ErrConstructor, pri)
	}
	// Must start with "TXP*"
	if !strings.HasPrefix(pri, "TXP*") {
		return fieldError("PaymentRelatedInformation", ErrVariableFields, pri)
	}
	// Must end with backslash terminator ("\\") (allow trailing spaces which are outside the 80-char PRI field)
	if !strings.HasSuffix(strings.TrimRight(pri, " "), "\\") {
		return fieldError("PaymentRelatedInformation", ErrVariableFields, pri)
	}
	if err := addendaTXP.isTXPAllowedChars(pri); err != nil {
		return fieldError("PaymentRelatedInformation", err, pri)
	}

	infoLength := utf8.RuneCountInString(addendaTXP.PaymentRelatedInformation)
	if infoLength > 80 {
		return fieldError("PaymentRelatedInformation", ErrExceedsFieldLength, addendaTXP.PaymentRelatedInformation)
	}
	// Sequence numbers must be positive (non-zero)
	if addendaTXP.SequenceNumber <= 0 {
		return fieldError("SequenceNumber", ErrInvalidProperty, addendaTXP.SequenceNumberField())
	}
	if addendaTXP.EntryDetailSequenceNumber <= 0 {
		return fieldError("EntryDetailSequenceNumber", ErrInvalidProperty, addendaTXP.EntryDetailSequenceNumberField())
	}

	return nil
}

// fieldInclusion validate mandatory fields are not default values. If fields are
// invalid the ACH transfer will be returned.
func (addendaTXP *AddendaTXP) fieldInclusion() error {
	if addendaTXP.TypeCode == "" {
		return fieldError("TypeCode", ErrConstructor, addendaTXP.TypeCode)
	}
	if addendaTXP.SequenceNumber == 0 {
		return fieldError("SequenceNumber", ErrConstructor, addendaTXP.SequenceNumberField())
	}
	if addendaTXP.EntryDetailSequenceNumber <= 0 {
		return fieldError("EntryDetailSequenceNumber", ErrConstructor, addendaTXP.EntryDetailSequenceNumberField())
	}
	return nil
}

// PaymentRelatedInformationField returns a left-justified, space-padded 80-character PaymentRelatedInformation field
func (addendaTXP *AddendaTXP) PaymentRelatedInformationField() string {
	return addendaTXP.alphaField(addendaTXP.PaymentRelatedInformation, 80)
}

// SequenceNumberField returns a zero padded SequenceNumber string
func (addendaTXP *AddendaTXP) SequenceNumberField() string {
	return addendaTXP.numericField(addendaTXP.SequenceNumber, 4)
}

// EntryDetailSequenceNumberField returns a zero padded EntryDetailSequenceNumber string
func (addendaTXP *AddendaTXP) EntryDetailSequenceNumberField() string {
	return addendaTXP.numericField(addendaTXP.EntryDetailSequenceNumber, 7)
}

// ToAddenda05 converts AddendaTXP to Addenda05 for compatibility
func (addendaTXP *AddendaTXP) ToAddenda05() *Addenda05 {
	addenda05 := NewAddenda05()
	addenda05.PaymentRelatedInformation = addendaTXP.PaymentRelatedInformation
	addenda05.SequenceNumber = addendaTXP.SequenceNumber
	addenda05.EntryDetailSequenceNumber = addendaTXP.EntryDetailSequenceNumber
	return addenda05
}

// isTXPAllowedChars ensures PaymentRelatedInformation only contains characters
// permitted by TXP addenda conventions (printable set and delimiters).
func (addendaTXP *AddendaTXP) isTXPAllowedChars(s string) error {
	for _, r := range s {
		if r == '\n' || r == '\r' || r == '\t' {
			return ErrInvalidProperty
		}
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			continue
		}
		switch r {
		case ' ', '*', '\\', '>', '-', '.', '/', ':':
			continue
		default:
			return ErrInvalidProperty
		}
	}
	return nil
}

// TXPAmount represents a (code, amount) pair, where Code may be T/P/I, a numeric subcategory (e.g., 1,14,16,20),
// or the repeated tax code (e.g., 09405) used in the no-subcategory pattern.
type TXPAmount struct {
	Code        string
	AmountCents int64
}

// TXPParsed represents parsed fields from a TXP segment.
type TXPParsed struct {
	TIN       string
	TaxCode   string
	PeriodEnd string // YYMMDD as provided in the segment
	Amounts   []TXPAmount
	Verify    string // optional taxpayer verification token at end
}

// ParseTXP splits PaymentRelatedInformation into TXP components and supports the
// patterns shown in NACHA examples: state T/P/I, EFTPS numeric subcategories, and repeated-tax-code.
func (a *AddendaTXP) ParseTXP() (TXPParsed, error) {
	var out TXPParsed
	if a == nil {
		return out, fmt.Errorf("nil AddendaTXP")
	}
	pri := strings.TrimRight(a.PaymentRelatedInformation, " ")
	if !strings.HasPrefix(pri, "TXP*") || !strings.HasSuffix(pri, "\\") {
		return out, fieldError("PaymentRelatedInformation", ErrVariableFields, pri)
	}
	pri = strings.TrimSuffix(pri, "\\")
	parts := strings.Split(pri, "*")
	if len(parts) < 4 { // TXP, TIN, TaxCode, PeriodEnd
		return out, fieldError("PaymentRelatedInformation", ErrVariableFields, pri)
	}
	out.TIN = parts[1]
	out.TaxCode = parts[2]
	out.PeriodEnd = parts[3]

	// Walk remaining elements as Code/Amount pairs; skip empty elements; optional trailing Verify token.
	i := 4
	for i < len(parts) {
		// Skip blanks
		for i < len(parts) && parts[i] == "" {
			i++
		}
		if i >= len(parts) {
			break
		}
		// If this looks like a code and we have an amount following, parse a pair
		if isTXPAmountCode(parts[i]) && i+1 < len(parts) && parts[i+1] != "" && isAllDigits(parts[i+1]) {
			amt, err := parseCents(parts[i+1])
			if err != nil {
				return out, fieldError("PaymentRelatedInformation", ErrNonNumeric, parts[i+1])
			}
			out.Amounts = append(out.Amounts, TXPAmount{Code: parts[i], AmountCents: amt})
			i += 2
			continue
		}
		// Otherwise, treat the remainder as a verification token (common for state examples)
		out.Verify = parts[i]
		break
	}

	return out, nil
}

// BuildTXPSegment assembles a TXP segment string (without padding to 80 chars).
// Caller typically assigns it to PaymentRelatedInformation and relies on PaymentRelatedInformationField()
// to left-justify and space-pad to 80.
func BuildTXPSegment(p TXPParsed) (string, error) {
	if p.TIN == "" || p.TaxCode == "" || p.PeriodEnd == "" {
		return "", fmt.Errorf("missing required fields: TIN/TaxCode/PeriodEnd")
	}
	b := strings.Builder{}
	b.WriteString("TXP*")
	b.WriteString(p.TIN)
	b.WriteString("*")
	b.WriteString(p.TaxCode)
	b.WriteString("*")
	b.WriteString(p.PeriodEnd)
	for _, kv := range p.Amounts {
		if kv.Code == "" || kv.AmountCents < 0 {
			return "", fmt.Errorf("invalid amount pair")
		}
		b.WriteString("*")
		b.WriteString(kv.Code)
		b.WriteString("*")
		b.WriteString(strconv.FormatInt(kv.AmountCents, 10))
	}
	if p.Verify != "" {
		b.WriteString("*")
		b.WriteString(p.Verify)
	}
	b.WriteString("\\")
	seg := b.String()
	if utf8.RuneCountInString(seg) > 80 {
		return "", fmt.Errorf("TXP segment exceeds 80 characters")
	}
	return seg, nil
}

// NormalizePeriodYYMMToYYMM01 maps a 4-digit YYMM into a 6-digit YYMM01 string.
// If the input isn't 4 numeric characters, it is returned unchanged.
func NormalizePeriodYYMMToYYMM01(s string) string {
	if len(s) == 4 && isAllDigits(s) {
		return s + "01"
	}
	return s
}

func isTXPAmountCode(s string) bool {
	if s == "T" || s == "P" || s == "I" { // state T/P/I labels
		return true
	}
	// Numeric code (e.g., 1, 14, 20, 09405)
	if s != "" && isAllDigits(s) {
		return true
	}
	return false
}

func isAllDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return s != ""
}

func parseCents(s string) (int64, error) {
	if !isAllDigits(s) {
		return 0, ErrNonNumeric
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return v, nil
}
