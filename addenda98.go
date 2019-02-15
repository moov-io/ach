// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"strings"
)

// Addenda98 is a Addendumer addenda record format for Notification OF Change(98)
// The field contents for Notification of Change Entries must match the field contents of the original Entries
type Addenda98 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record in the block. entryAddendaPos 7
	recordType string
	// TypeCode Addenda types code '98'
	TypeCode string `json:"typeCode"`
	// ChangeCode field contains a standard code used by an ACH Operator or RDFI to describe the reason for a change Entry.
	// Must exist in changeCodeDict
	ChangeCode string `json:"changeCode"`
	// OriginalTrace This field contains the Trace Number as originally included on the forward Entry or Prenotification.
	// The RDFI must include the Original Entry Trace Number in the Addenda Record of an Entry being returned to an ODFI,
	// in the Addenda Record of an 98, within an Acknowledgment Entry, or with an RDFI request for a copy of an authorization.
	OriginalTrace string `json:"originalTrace"`
	// OriginalDFI field contains the Receiving DFI Identification (addenda.RDFIIdentification) as originally included on the forward Entry or Prenotification that the RDFI is returning or correcting.
	OriginalDFI string `json:"originalDFI"`
	// CorrectedData
	CorrectedData string `json:"correctedData"`
	// TraceNumber matches the Entry Detail Trace Number of the entry being returned.
	//
	// Use TraceNumberField() for a properly formatted string representation.
	TraceNumber string `json:"traceNumber,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters
}

var (
	changeCodeDict = map[string]*changeCode{}
)

func init() {
	// populate the changeCode map with lookup values
	changeCodeDict = makeChangeCodeDict()
}

// changeCode holds a change Code, Reason/Title, and Description
// table of return codes exists in Part 4.2 of the NACHA corporate rules and guidelines
type changeCode struct {
	Code, Reason, Description string
}

// NewAddenda98 returns an reference to an instantiated Addenda98 with default values
func NewAddenda98() *Addenda98 {
	addenda98 := &Addenda98{
		recordType: "7",
		TypeCode:   "98",
	}
	return addenda98
}

// Parse takes the input record string and parses the Addenda98 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm successful parsing and data validity.
func (addenda98 *Addenda98) Parse(record string) {
	// 1-1 Always "7"
	addenda98.recordType = "7"
	// 2-3 Always "98"
	addenda98.TypeCode = record[1:3]
	// 4-6
	addenda98.ChangeCode = record[3:6]
	// 7-21
	addenda98.OriginalTrace = strings.TrimSpace(record[6:21])
	// 28-35
	addenda98.OriginalDFI = addenda98.parseStringField(record[27:35])
	// 36-64
	addenda98.CorrectedData = strings.TrimSpace(record[35:64])
	// 80-94
	addenda98.TraceNumber = strings.TrimSpace(record[79:94])
}

// String writes the Addenda98 struct to a 94 character string
func (addenda98 *Addenda98) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(addenda98.recordType)
	buf.WriteString(addenda98.TypeCode)
	buf.WriteString(addenda98.ChangeCode)
	buf.WriteString(addenda98.OriginalTraceField())
	buf.WriteString("      ") // 6 char reserved field
	buf.WriteString(addenda98.OriginalDFIField())
	buf.WriteString(addenda98.CorrectedDataField())
	buf.WriteString("               ") // 15 char reserved field
	buf.WriteString(addenda98.TraceNumberField())
	return buf.String()
}

// Validate verifies NACHA rules for Addenda98
func (addenda98 *Addenda98) Validate() error {
	if addenda98.recordType != "7" {
		return fieldError("recordType", NewErrRecordType(7), addenda98.recordType)
	}
	if addenda98.TypeCode == "" {
		return fieldError("TypeCode", ErrConstructor, addenda98.TypeCode)
	}
	// Type Code must be 98
	if addenda98.TypeCode != "98" {
		return fieldError("TypeCode", ErrAddendaTypeCode, addenda98.TypeCode)
	}

	// Addenda98 requires a valid ChangeCode
	_, ok := changeCodeDict[addenda98.ChangeCode]
	if !ok {
		return fieldError("ChangeCode", ErrAddenda98ChangeCode, addenda98.ChangeCode)
	}

	// Addenda98 Record must contain the corrected information corresponding to the Change Code used
	if addenda98.CorrectedData == "" {
		return fieldError("CorrectedData", ErrAddenda98CorrectedData, addenda98.CorrectedData)
	}

	return nil
}

// OriginalTraceField returns a zero padded OriginalTrace string
func (addenda98 *Addenda98) OriginalTraceField() string {
	return addenda98.stringField(addenda98.OriginalTrace, 15)
}

// OriginalDFIField returns a zero padded OriginalDFI string
func (addenda98 *Addenda98) OriginalDFIField() string {
	return addenda98.stringField(addenda98.OriginalDFI, 8)
}

//CorrectedDataField returns a space padded CorrectedData string
func (addenda98 *Addenda98) CorrectedDataField() string {
	return addenda98.alphaField(addenda98.CorrectedData, 29)
}

// TraceNumberField returns a zero padded traceNumber string
func (addenda98 *Addenda98) TraceNumberField() string {
	return addenda98.stringField(addenda98.TraceNumber, 15)
}

func makeChangeCodeDict() map[string]*changeCode {
	dict := make(map[string]*changeCode)

	codes := []changeCode{
		{"C01", "Incorrect bank account number", "Bank account number incorrect or formatted incorrectly"},
		{"C02", "Incorrect transit/routing number", "Once valid transit/routing number must be changed"},
		{"C03", "Incorrect transit/routing number and bank account number", "Once valid transit/routing number must be changed and causes a change to bank account number structure"},
		{"C04", "Bank account name change", "Customer has changed name or ODFI submitted name incorrectly"},
		{"C05", "Incorrect payment code", "Entry posted to demand account should contain savings payment codes or vice versa"},
		{"C06", "Incorrect bank account number and transit code", "Bank account number must be changed and payment code should indicate posting to another account type (demand/savings)"},
		{"C07", "Incorrect transit/routing number, bank account number and payment code", "Changes required in three fields indicated"},
		{"C09", "Incorrect individual ID number", "Individual's ID number is incorrect"},
		{"C10", "Incorrect company name", "Company name is no longer valid and should be changed."},
		{"C11", "Incorrect company identification", "Company ID is no longer valid and should be changed"},
		{"C12", "Incorrect company name and company ID", "Both the company name and company id are no longer valid and must be changed"},
	}
	// populate the map
	for _, code := range codes {
		dict[code.Code] = &code
	}
	return dict
}
