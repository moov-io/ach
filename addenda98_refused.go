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
	"strings"
	"unicode/utf8"
)

type Addenda98Refused struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`

	// TypeCode Addenda types code '98'
	TypeCode string `json:"typeCode"`

	// RefusedChangeCode is the code specifying why the Notification of Change is being refused.
	RefusedChangeCode string `json:"refusedChangeCode"`

	// OriginalTrace This field contains the Trace Number as originally included on the forward Entry or Prenotification.
	// The RDFI must include the Original Entry Trace Number in the Addenda Record of an Entry being returned to an ODFI,
	// in the Addenda Record of an 98, within an Acknowledgment Entry, or with an RDFI request for a copy of an authorization.
	OriginalTrace string `json:"originalTrace"`

	// OriginalDFI field contains the Receiving DFI Identification (addenda.RDFIIdentification) as originally included on the
	// forward Entry or Prenotification that the RDFI is returning or correcting.
	OriginalDFI string `json:"originalDFI"`

	// CorrectedData is the corrected data
	CorrectedData string `json:"correctedData"`

	// ChangeCode field contains a standard code used by an ACH Operator or RDFI to describe the reason for a change Entry.
	ChangeCode string `json:"changeCode"`

	// TraceSequenceNumber is the last seven digits of the TraceNumber in the original Notification of Change
	TraceSequenceNumber string `json:"traceSequenceNumber"`

	// TraceNumber matches the Entry Detail Trace Number of the entry being returned.
	//
	// Use TraceNumberField for a properly formatted string representation.
	TraceNumber string `json:"traceNumber"`

	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters
}

// NewAddenda98Refused returns an reference to an instantiated Addenda98Refused with default values
func NewAddenda98Refused() *Addenda98Refused {
	addenda98Refused := &Addenda98Refused{
		TypeCode: "98",
	}
	return addenda98Refused
}

// Parse takes the input record string and parses the Addenda98Refused values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate call to confirm successful parsing and data validity.
func (addenda98Refused *Addenda98Refused) Parse(record string) {
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
			addenda98Refused.TypeCode = strings.TrimSpace(reset())
		case 6:
			addenda98Refused.RefusedChangeCode = strings.TrimSpace(reset())
		case 21:
			addenda98Refused.OriginalTrace = strings.TrimSpace(reset())
		case 27:
			// Positions 22-27 are Reserved
			reset()
		case 35:
			addenda98Refused.OriginalDFI = addenda98Refused.parseStringField(reset())
		case 64:
			addenda98Refused.CorrectedData = strings.TrimSpace(reset())
		case 67:
			addenda98Refused.ChangeCode = strings.TrimSpace(reset())
		case 74:
			addenda98Refused.TraceSequenceNumber = strings.TrimSpace(reset())
		case 79:
			// Positions 75-79 are Reserved
			reset()
		case 94:
			addenda98Refused.TraceNumber = strings.TrimSpace(reset())
		}
	}
}

// String writes the Addenda98 struct to a 94 character string
func (addenda98Refused *Addenda98Refused) String() string {
	if addenda98Refused == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(entryAddendaPos)
	buf.WriteString(addenda98Refused.TypeCode)
	buf.WriteString(addenda98Refused.RefusedChangeCode)
	buf.WriteString(addenda98Refused.OriginalTraceField())
	buf.WriteString(strings.Repeat(" ", 6))
	buf.WriteString(addenda98Refused.OriginalDFIField())
	buf.WriteString(addenda98Refused.CorrectedDataField())
	buf.WriteString(addenda98Refused.ChangeCode)
	buf.WriteString(addenda98Refused.TraceSequenceNumberField())
	buf.WriteString(strings.Repeat(" ", 5))
	buf.WriteString(addenda98Refused.TraceNumberField())
	return buf.String()
}

// Validate verifies NACHA rules for Addenda98
func (addenda98Refused *Addenda98Refused) Validate() error {
	if addenda98Refused.TypeCode == "" {
		return fieldError("TypeCode", ErrConstructor, addenda98Refused.TypeCode)
	}
	// Type Code must be 98
	if addenda98Refused.TypeCode != "98" {
		return fieldError("TypeCode", ErrAddendaTypeCode, addenda98Refused.TypeCode)
	}

	// RefusedChangeCode must be valid
	_, ok := changeCodeDict[addenda98Refused.RefusedChangeCode]
	if !ok {
		return fieldError("RefusedChangeCode", ErrAddenda98RefusedChangeCode, addenda98Refused.RefusedChangeCode)
	}

	// Addenda98 Record must contain the corrected information corresponding to the Change Code used
	if addenda98Refused.CorrectedData == "" {
		return fieldError("CorrectedData", ErrAddenda98CorrectedData, addenda98Refused.CorrectedData)
	}

	// ChangeCode must be valid
	_, ok = changeCodeDict[addenda98Refused.ChangeCode]
	if !ok {
		return fieldError("ChangeCode", ErrAddenda98ChangeCode, addenda98Refused.ChangeCode)
	}

	// TraceSequenceNumber must be valid
	if addenda98Refused.TraceSequenceNumber == "" {
		return fieldError("TraceSequenceNumber", ErrAddenda98RefusedTraceSequenceNumber, addenda98Refused.TraceSequenceNumber)
	}

	return nil
}

func (addenda98Refused *Addenda98Refused) RefusedChangeCodeField() *ChangeCode {
	code, ok := changeCodeDict[addenda98Refused.RefusedChangeCode]
	if ok {
		return code
	}
	return nil
}

// OriginalTraceField returns a zero padded OriginalTrace string
func (addenda98Refused *Addenda98Refused) OriginalTraceField() string {
	return addenda98Refused.stringField(addenda98Refused.OriginalTrace, 15)
}

// OriginalDFIField returns a zero padded OriginalDFI string
func (addenda98Refused *Addenda98Refused) OriginalDFIField() string {
	return addenda98Refused.stringField(addenda98Refused.OriginalDFI, 8)
}

// CorrectedDataField returns a space padded CorrectedData string
func (addenda98Refused *Addenda98Refused) CorrectedDataField() string {
	return addenda98Refused.alphaField(addenda98Refused.CorrectedData, 29)
}

func (addenda98Refused *Addenda98Refused) ChangeCodeField() *ChangeCode {
	code, ok := changeCodeDict[addenda98Refused.ChangeCode]
	if ok {
		return code
	}
	return nil
}

func (addenda98Refused *Addenda98Refused) TraceSequenceNumberField() string {
	return addenda98Refused.stringField(addenda98Refused.TraceSequenceNumber, 7)
}

// TraceNumberField returns a zero padded traceNumber string
func (addenda98Refused *Addenda98Refused) TraceNumberField() string {
	return addenda98Refused.stringField(addenda98Refused.TraceNumber, 15)
}
