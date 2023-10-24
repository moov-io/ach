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

type Addenda99Dishonored struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`

	// TypeCode Addenda types code '99'
	TypeCode string `json:"typeCode"`

	// DishonoredReturnReasonCode is the return code explaining the dishonorment
	DishonoredReturnReasonCode string `json:"dishonoredReturnReasonCode"`

	// OriginalEntryTraceNumber is the trace number specifieid in the initial entry
	OriginalEntryTraceNumber string `json:"originalEntryTraceNumber"`

	// OriginalReceivingDFIIdentification is the DFI Identification specifieid in the initial entry
	OriginalReceivingDFIIdentification string `json:"originalReceivingDFIIdentification"`

	// ReturnTraceNumber is the TraceNumber used when issuing the return
	ReturnTraceNumber string `json:"returnTraceNumber"`

	// ReturnSettlementDate is the date of return issuing
	ReturnSettlementDate string `json:"returnSettlementDate"`

	// ReturnReasonCode is the initial return code
	ReturnReasonCode string `json:"returnReasonCode"`

	// AddendaInformation is additional data
	AddendaInformation string `json:"addendaInformation"`

	// TraceNumber is the trace number for dishonorment
	TraceNumber string `json:"traceNumber"`

	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters

	validateOpts *ValidateOpts
}

// NewAddenda99Dishonored returns a new Addenda99Dishonored with default values for none exported fields
func NewAddenda99Dishonored() *Addenda99Dishonored {
	Addenda99Dishonored := &Addenda99Dishonored{
		TypeCode: "99",
	}
	return Addenda99Dishonored
}

func (Addenda99Dishonored *Addenda99Dishonored) Parse(record string) {
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
			Addenda99Dishonored.TypeCode = reset()
		case 6:
			Addenda99Dishonored.DishonoredReturnReasonCode = reset()
		case 21:
			Addenda99Dishonored.OriginalEntryTraceNumber = reset()
		case 35:
			Addenda99Dishonored.OriginalReceivingDFIIdentification = reset()
		case 53:
			Addenda99Dishonored.ReturnTraceNumber = reset()
		case 56:
			Addenda99Dishonored.ReturnSettlementDate = reset()
		case 58:
			Addenda99Dishonored.ReturnReasonCode = reset()
		case 79:
			Addenda99Dishonored.AddendaInformation = reset()
		case 94:
			Addenda99Dishonored.TraceNumber = reset()
		}
	}
}

func (Addenda99Dishonored *Addenda99Dishonored) String() string {
	if Addenda99Dishonored == nil {
		return ""
	}

	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(entryAddendaPos)
	buf.WriteString(Addenda99Dishonored.TypeCode)
	buf.WriteString(Addenda99Dishonored.DishonoredReturnReasonCodeField())
	buf.WriteString(Addenda99Dishonored.OriginalEntryTraceNumberField())
	buf.WriteString("      ")
	buf.WriteString(Addenda99Dishonored.OriginalReceivingDFIIdentificationField())
	buf.WriteString("   ")
	buf.WriteString(Addenda99Dishonored.ReturnTraceNumberField())
	buf.WriteString(Addenda99Dishonored.ReturnSettlementDateField())
	buf.WriteString(Addenda99Dishonored.ReturnReasonCodeField())
	buf.WriteString(Addenda99Dishonored.AddendaInformationField())
	buf.WriteString(Addenda99Dishonored.TraceNumberField())
	return buf.String()
}

// SetValidation stores ValidateOpts on the Batch which are to be used to override
// the default NACHA validation rules.
func (Addenda99Dishonored *Addenda99Dishonored) SetValidation(opts *ValidateOpts) {
	if Addenda99Dishonored == nil {
		return
	}
	Addenda99Dishonored.validateOpts = opts
}

func IsDishonoredReturnCode(code string) bool {
	switch code {
	case "R61", "R67", "R68", "R69", "R70":
		return true
	}
	return false
}

// Validate verifies NACHA rules for Addenda99Dishonored
func (Addenda99Dishonored *Addenda99Dishonored) Validate() error {
	if Addenda99Dishonored.TypeCode == "" {
		return fieldError("TypeCode", ErrConstructor, Addenda99Dishonored.TypeCode)
	}
	if Addenda99Dishonored.TypeCode != "99" {
		return fieldError("TypeCode", ErrAddendaTypeCode, Addenda99Dishonored.TypeCode)
	}

	// Verify the DishonoredReturnReasonCode matches expected values
	if Addenda99Dishonored.validateOpts == nil || !Addenda99Dishonored.validateOpts.CustomReturnCodes {
		// We can validate the Dishonored ReturnCode
		if !IsDishonoredReturnCode(Addenda99Dishonored.DishonoredReturnReasonCode) {
			return fieldError("DishonoredReturnReasonCode", ErrAddenda99DishonoredReturnCode, Addenda99Dishonored.DishonoredReturnReasonCode)
		}
	}

	return nil
}

func (Addenda99Dishonored *Addenda99Dishonored) DishonoredReturnReasonCodeField() string {
	return Addenda99Dishonored.stringField(Addenda99Dishonored.DishonoredReturnReasonCode, 3)
}

// OriginalEntryTraceNumberField returns a zero padded TraceNumber string
func (Addenda99Dishonored *Addenda99Dishonored) OriginalEntryTraceNumberField() string {
	return Addenda99Dishonored.stringField(Addenda99Dishonored.OriginalEntryTraceNumber, 15)
}

func (Addenda99Dishonored *Addenda99Dishonored) OriginalReceivingDFIIdentificationField() string {
	return Addenda99Dishonored.stringField(Addenda99Dishonored.OriginalReceivingDFIIdentification, 8)
}

func (Addenda99Dishonored *Addenda99Dishonored) ReturnTraceNumberField() string {
	return Addenda99Dishonored.stringField(Addenda99Dishonored.ReturnTraceNumber, 15)
}

func (Addenda99Dishonored *Addenda99Dishonored) ReturnSettlementDateField() string {
	return Addenda99Dishonored.stringField(Addenda99Dishonored.ReturnSettlementDate, 3)
}

func (Addenda99Dishonored *Addenda99Dishonored) ReturnReasonCodeField() string {
	return Addenda99Dishonored.stringField(Addenda99Dishonored.ReturnReasonCode, 2)
}

func (Addenda99Dishonored *Addenda99Dishonored) AddendaInformationField() string {
	return Addenda99Dishonored.alphaField(Addenda99Dishonored.AddendaInformation, 21)
}

// TraceNumberField returns a zero padded TraceNumber string
func (Addenda99Dishonored *Addenda99Dishonored) TraceNumberField() string {
	return Addenda99Dishonored.stringField(Addenda99Dishonored.TraceNumber, 15)
}
