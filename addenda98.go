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

// Addenda98 is a Addendumer addenda record format for Notification OF Change(98)
// The field contents for Notification of Change Entries must match the field contents of the original Entries
type Addenda98 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
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
	// CorrectedData is the corrected data
	CorrectedData string `json:"correctedData"`
	// TraceNumber matches the Entry Detail Trace Number of the entry being returned.
	//
	// Use TraceNumberField for a properly formatted string representation.
	TraceNumber string `json:"traceNumber,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters
}

var (
	changeCodeDict = map[string]*ChangeCode{}
)

func init() {
	// populate the changeCode map with lookup values
	changeCodeDict = makeChangeCodeDict()
}

// ChangeCode holds a change Code, Reason/Title, and Description
// table of return codes exists in Part 4.2 of the NACHA corporate rules and guidelines
type ChangeCode struct {
	Code        string `json:"code"`
	Reason      string `json:"reason"`
	Description string `json:"description"`
}

// NewAddenda98 returns an reference to an instantiated Addenda98 with default values
func NewAddenda98() *Addenda98 {
	addenda98 := &Addenda98{
		TypeCode: "98",
	}
	return addenda98
}

// Parse takes the input record string and parses the Addenda98 values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda98 *Addenda98) Parse(record string) {
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
			// 2-3 Always "98"
			addenda98.TypeCode = reset()
		case 6:
			// 4-6
			addenda98.ChangeCode = reset()
		case 21:
			// 7-21
			addenda98.OriginalTrace = strings.TrimSpace(reset())
		case 35:
			// 28-35
			addenda98.OriginalDFI = addenda98.parseStringField(reset())
		case 64:
			// 36-64
			addenda98.CorrectedData = strings.TrimSpace(reset())
		case 94:
			// 80-94
			addenda98.TraceNumber = strings.TrimSpace(reset())
		}
	}
}

// String writes the Addenda98 struct to a 94 character string
func (addenda98 *Addenda98) String() string {
	if addenda98 == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(entryAddendaPos)
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

// CorrectedDataField returns a space padded CorrectedData string
func (addenda98 *Addenda98) CorrectedDataField() string {
	return addenda98.alphaField(addenda98.CorrectedData, 29)
}

// TraceNumberField returns a zero padded traceNumber string
func (addenda98 *Addenda98) TraceNumberField() string {
	return addenda98.stringField(addenda98.TraceNumber, 15)
}

func (addenda98 *Addenda98) ChangeCodeField() *ChangeCode {
	code, ok := changeCodeDict[addenda98.ChangeCode]
	if ok {
		return code
	}
	return nil
}

// LookupChangeCode will return a struct representing the reason and description for
// the provided NACHA change code.
func LookupChangeCode(code string) *ChangeCode {
	if code, exists := changeCodeDict[strings.ToUpper(code)]; exists {
		return code
	}
	return nil
}

func makeChangeCodeDict() map[string]*ChangeCode {
	dict := make(map[string]*ChangeCode)

	codes := []ChangeCode{
		{"C01", "Incorrect bank account number", "Bank account number incorrect or formatted incorrectly"},
		{"C02", "Incorrect transit/routing number", "Once valid transit/routing number must be changed"},
		{"C03", "Incorrect transit/routing number and bank account number", "Once valid transit/routing number must be changed and causes a change to bank account number structure"},
		{"C04", "Bank account name change", "Customer has changed name or ODFI submitted name incorrectly"},
		{"C05", "Incorrect transaction code", "Entry posted to demand account should contain savings transaction codes or vice versa"},
		{"C06", "Incorrect bank account number and transit code", "Bank account number must be changed and transaction code should indicate posting to another account type (demand/savings)"},
		{"C07", "Incorrect transit/routing number, bank account number and transaction code", "Changes required in three fields indicated"},
		{"C09", "Incorrect individual ID number", "Individual's ID number is incorrect"},
		{"C10", "Incorrect company name", "Company name is no longer valid and should be changed."},
		{"C11", "Incorrect company identification", "Company ID is no longer valid and should be changed"},
		{"C12", "Incorrect company name and company ID", "Both the company name and company id are no longer valid and must be changed"},
		// Change codes used when refusing a Notification of Change
		{"C61", "Misrouted Notification of Change", ""},
		{"C62", "Incorrect Trace Number", ""},
		{"C63", "Incorrect Company Identification Number", ""},
		{"C64", "Incorrect Individual Identification Number or Identification Number", ""},
		{"C65", "Incorrectly Formatted Corrected Data", ""},
		{"C66", "Incorrect Discretionary Data", ""},
		{"C67", "Routing Number not from Original Entry Detail Record", ""},
		{"C68", "DFI Account Number not from Original Entry Detail Record", ""},
		{"C69", "Incorrect Transaction Code", ""},
	}
	// populate the map
	for i := range codes {
		dict[codes[i].Code] = &codes[i]
	}
	return dict
}

func IsRefusedChangeCode(code string) bool {
	switch strings.ToUpper(code) {
	case "C61", "C62", "C63", "C64", "C65", "C66", "C67", "C68", "C69":
		return true
	}
	return false
}

// CorrectedData is a struct returned from our helper method for parsing the NOC/COR
// corrected data from Addenda98 records.
//
// All fields are optional and a valid code may not have populated data in this struct.
type CorrectedData struct {
	AccountNumber   string
	RoutingNumber   string
	Name            string
	TransactionCode int
	Identification  string
}

// ParseCorrectedData returns a struct with some fields filled in depending on the Addenda98's
// Code and CorrectedData. Fields are trimmed when populated in this struct.
func (addenda98 *Addenda98) ParseCorrectedData() *CorrectedData {
	if addenda98 == nil {
		return nil
	}
	cc := addenda98.ChangeCodeField()
	if cc == nil {
		return nil
	}
	switch cc.Code {
	case "C01": // Incorrect DFI Account Number
		if v := first(17, addenda98.CorrectedData); v != "" {
			return &CorrectedData{AccountNumber: v}
		}
	case "C02": // Incorrect Routing Number
		if v := first(9, addenda98.CorrectedData); v != "" {
			return &CorrectedData{RoutingNumber: v}
		}
	case "C03": // Incorrect Routing Number and Incorrect DFI Account Number
		parts := strings.Fields(addenda98.CorrectedData)
		if len(parts) == 2 {
			return &CorrectedData{
				RoutingNumber: parts[0],
				AccountNumber: parts[1],
			}
		}
	case "C04": // Incorrect Individual Name
		if v := first(22, addenda98.CorrectedData); v != "" {
			return &CorrectedData{Name: v}
		}
	case "C05": // Incorrect Transaction Code
		if n, err := strconv.Atoi(first(2, addenda98.CorrectedData)); err == nil {
			return &CorrectedData{TransactionCode: n}
		}
	case "C06": // Incorrect DFI Account Number and Incorrect Transaction Code
		parts := strings.Fields(addenda98.CorrectedData)
		if len(parts) == 2 {
			if n, err := strconv.Atoi(parts[1]); err == nil {
				return &CorrectedData{
					AccountNumber:   parts[0],
					TransactionCode: n,
				}
			}
		}
	case "C07": // Incorrect Routing Number, Incorrect DFI Account Number, and Incorrect Tranaction Code
		var cd CorrectedData
		if n := len(addenda98.CorrectedData); n > 9 {
			cd.RoutingNumber = addenda98.CorrectedData[:9]
		} else {
			return nil
		}
		parts := strings.Fields(addenda98.CorrectedData[9:])
		if len(parts) == 2 {
			if n, err := strconv.Atoi(parts[1]); err == nil {
				cd.AccountNumber = parts[0]
				cd.TransactionCode = n
				return &cd
			}
		} else {
			return nil
		}
	case "C09": // Incorrect Individual Identification Number
		if v := first(22, addenda98.CorrectedData); v != "" {
			return &CorrectedData{Identification: v}
		}
	}
	// The Code/Correction is either unsupported or wasn't parsed correctly
	return nil
}

func first(size int, data string) string {
	if utf8.RuneCountInString(data) < size {
		if data != "" {
			return strings.TrimSpace(data)
		} else {
			return ""
		}
	}
	return strings.TrimSpace(data[:size])
}

const correctedDataCharLength = 29

// ParseCorrectedData returns the string properlty formatted and justified for an
// Addenda98.CorrectedData field. The code must be an official NACHA change code.
func WriteCorrectionData(code string, data *CorrectedData) string {
	pad := &converters{}
	switch strings.ToUpper(code) {
	case "C01":
		return pad.alphaField(data.AccountNumber, correctedDataCharLength)
	case "C02":
		return pad.alphaField(data.RoutingNumber, correctedDataCharLength)
	case "C03":
		spaces := strings.Repeat(" ", correctedDataCharLength-len(data.RoutingNumber)-len(data.AccountNumber))
		return fmt.Sprintf("%s%s%s", data.RoutingNumber, spaces, data.AccountNumber)
	case "C04":
		return pad.alphaField(data.Name, correctedDataCharLength)
	case "C05":
		return pad.alphaField(strconv.Itoa(data.TransactionCode), correctedDataCharLength)
	case "C06":
		txcode := strconv.Itoa(data.TransactionCode)
		spaces := strings.Repeat(" ", correctedDataCharLength-len(data.AccountNumber)-len(txcode))
		return fmt.Sprintf("%s%s%s", data.AccountNumber, spaces, txcode)
	case "C07":
		txcode := strconv.Itoa(data.TransactionCode)
		spaces := strings.Repeat(" ", correctedDataCharLength-9-len(data.AccountNumber)-len(txcode))
		return fmt.Sprintf("%s%s%s%s", data.RoutingNumber, data.AccountNumber, spaces, txcode)
	case "C09":
		return pad.alphaField(data.Identification, correctedDataCharLength)
	}
	return pad.alphaField("", correctedDataCharLength)
}
