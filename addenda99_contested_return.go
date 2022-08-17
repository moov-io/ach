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

type Addenda99Contested struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`

	// TypeCode Addenda types code '99'
	TypeCode string `json:"typeCode"`

	// ContestedReturnCode is the return code explaining the contested dishonorment
	ContestedReturnCode string `json:"contestedReturnCode"`

	// OriginalEntryTraceNumber is the trace number specifieid in the initial entry
	OriginalEntryTraceNumber string `json:"originalEntryTraceNumber"`

	// DateOriginalEntryReturned is the original entry's date
	DateOriginalEntryReturned string `json:"dateOriginalEntryReturned"`

	// OriginalReceivingDFIIdentification is the DFI Identification specifieid in the initial entry
	OriginalReceivingDFIIdentification string `json:"originalReceivingDFIIdentification"`

	// OriginalSettlementDate is the initial date of settlement
	OriginalSettlementDate string `json:"originalSettlementDate"`

	// ReturnTraceNumber is the original returns trace number
	ReturnTraceNumber string `json:"returnTraceNumber"`

	// ReturnSettlementDate is the original return's settlement date
	ReturnSettlementDate string `json:"returnSettlementDate"`

	// ReturnReasonCode is the original return's code
	ReturnReasonCode string `json:"returnReasonCode"`

	// DishonoredReturnTraceNumber is the dishonorment's trace number
	DishonoredReturnTraceNumber string `json:"dishonoredReturnTraceNumber"`

	// DishonoredReturnSettlementDate is the dishonorment's settlement date
	DishonoredReturnSettlementDate string `json:"dishonoredReturnSettlementDate"`

	// DishonoredReturnReasonCode is the dishonorment's return code
	DishonoredReturnReasonCode string `json:"dishonoredReturnReasonCode"`

	// TraceNumber is the trace number for contesting
	TraceNumber string `json:"traceNumber"`

	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters

	validateOpts *ValidateOpts
}

// NewAddenda99Contested returns a new Addenda99Contested with default values for none exported fields
func NewAddenda99Contested() *Addenda99Contested {
	Addenda99Contested := &Addenda99Contested{
		TypeCode: "99",
	}
	return Addenda99Contested
}

func (Addenda99Contested *Addenda99Contested) Parse(record string) {
	if utf8.RuneCountInString(record) != 94 {
		return
	}

	Addenda99Contested.TypeCode = record[1:3]
	Addenda99Contested.ContestedReturnCode = record[3:6]
	Addenda99Contested.OriginalEntryTraceNumber = record[6:21]
	Addenda99Contested.DateOriginalEntryReturned = record[21:27]
	Addenda99Contested.OriginalReceivingDFIIdentification = record[27:35]
	Addenda99Contested.OriginalSettlementDate = record[35:38]
	Addenda99Contested.ReturnTraceNumber = record[38:53]
	Addenda99Contested.ReturnSettlementDate = record[53:56]
	Addenda99Contested.ReturnReasonCode = record[56:58]
	Addenda99Contested.DishonoredReturnTraceNumber = record[58:73]
	Addenda99Contested.DishonoredReturnSettlementDate = record[73:76]
	Addenda99Contested.DishonoredReturnReasonCode = record[76:78]
	Addenda99Contested.TraceNumber = record[79:94]
}

func (Addenda99Contested *Addenda99Contested) String() string {
	var buf strings.Builder
	buf.Grow(94)

	buf.WriteString(entryAddendaPos)
	buf.WriteString(Addenda99Contested.TypeCode)
	buf.WriteString(Addenda99Contested.ContestedReturnCodeField())
	buf.WriteString(Addenda99Contested.OriginalEntryTraceNumberField())
	buf.WriteString(Addenda99Contested.DateOriginalEntryReturnedField())
	buf.WriteString(Addenda99Contested.OriginalReceivingDFIIdentificationField())
	buf.WriteString(Addenda99Contested.OriginalSettlementDateField())
	buf.WriteString(Addenda99Contested.ReturnTraceNumberField())
	buf.WriteString(Addenda99Contested.ReturnSettlementDateField())
	buf.WriteString(Addenda99Contested.ReturnReasonCodeField())
	buf.WriteString(Addenda99Contested.DishonoredReturnTraceNumberField())
	buf.WriteString(Addenda99Contested.DishonoredReturnSettlementDateField())
	buf.WriteString(Addenda99Contested.DishonoredReturnReasonCodeField())
	buf.WriteString(" ")
	buf.WriteString(Addenda99Contested.TraceNumberField())

	return buf.String()
}

// SetValidation stores ValidateOpts on the Batch which are to be used to override
// the default NACHA validation rules.
func (Addenda99Contested *Addenda99Contested) SetValidation(opts *ValidateOpts) {
	if Addenda99Contested == nil {
		return
	}
	Addenda99Contested.validateOpts = opts
}

// Validate verifies NACHA rules for Addenda99Contested
func (Addenda99Contested *Addenda99Contested) Validate() error {
	if Addenda99Contested.TypeCode == "" {
		return fieldError("TypeCode", ErrConstructor, Addenda99Contested.TypeCode)
	}
	if Addenda99Contested.TypeCode != "99" {
		return fieldError("TypeCode", ErrAddendaTypeCode, Addenda99Contested.TypeCode)
	}

	// Verify the ContestedReturnReasonCode matches expected values
	if Addenda99Contested.validateOpts == nil || !Addenda99Contested.validateOpts.CustomReturnCodes {
		// We can validate the Contested ReturnCode
		if !IsContestedReturnCode(Addenda99Contested.ContestedReturnCode) {
			return fieldError("ContestedReturnCode", ErrAddenda99ContestedReturnCode, Addenda99Contested.ContestedReturnCode)
		}
	}

	return nil
}

func IsContestedReturnCode(code string) bool {
	switch code {
	case "R71", "R72", "R73", "R74", "R75", "R76":
		return true
	}
	return false
}

func (Addenda99Contested *Addenda99Contested) ContestedReturnCodeField() string {
	return Addenda99Contested.stringField(Addenda99Contested.ContestedReturnCode, 3)
}

func (Addenda99Contested *Addenda99Contested) OriginalEntryTraceNumberField() string {
	return Addenda99Contested.stringField(Addenda99Contested.OriginalEntryTraceNumber, 15)
}

func (Addenda99Contested *Addenda99Contested) DateOriginalEntryReturnedField() string {
	return Addenda99Contested.stringField(Addenda99Contested.DateOriginalEntryReturned, 6)
}

func (Addenda99Contested *Addenda99Contested) OriginalReceivingDFIIdentificationField() string {
	return Addenda99Contested.stringField(Addenda99Contested.OriginalReceivingDFIIdentification, 8)
}

func (Addenda99Contested *Addenda99Contested) OriginalSettlementDateField() string {
	return Addenda99Contested.stringField(Addenda99Contested.OriginalSettlementDate, 3)
}

func (Addenda99Contested *Addenda99Contested) ReturnTraceNumberField() string {
	return Addenda99Contested.stringField(Addenda99Contested.ReturnTraceNumber, 15)
}

func (Addenda99Contested *Addenda99Contested) ReturnSettlementDateField() string {
	return Addenda99Contested.stringField(Addenda99Contested.ReturnSettlementDate, 3)
}

func (Addenda99Contested *Addenda99Contested) ReturnReasonCodeField() string {
	return Addenda99Contested.stringField(Addenda99Contested.ReturnReasonCode, 2)
}

func (Addenda99Contested *Addenda99Contested) DishonoredReturnTraceNumberField() string {
	return Addenda99Contested.stringField(Addenda99Contested.DishonoredReturnTraceNumber, 15)
}

func (Addenda99Contested *Addenda99Contested) DishonoredReturnSettlementDateField() string {
	return Addenda99Contested.stringField(Addenda99Contested.DishonoredReturnSettlementDate, 3)
}

func (Addenda99Contested *Addenda99Contested) DishonoredReturnReasonCodeField() string {
	return Addenda99Contested.stringField(Addenda99Contested.DishonoredReturnReasonCode, 2)
}

func (Addenda99Contested *Addenda99Contested) TraceNumberField() string {
	return Addenda99Contested.stringField(Addenda99Contested.TraceNumber, 15)
}
