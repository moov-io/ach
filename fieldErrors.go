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
	"errors"
	"fmt"
)

var (
	// Errors specific to validation

	//ErrNonAlphanumeric is given when a field has non-alphanumeric characters
	ErrNonAlphanumeric = errors.New("has non alphanumeric characters")
	//ErrUpperAlpha is given when a field is not in uppercase
	ErrUpperAlpha = errors.New("is not uppercase A-Z or 0-9")
	//ErrFieldInclusion is given when a field is mandatory and has a default value
	ErrFieldInclusion = errors.New("is a mandatory field and has a default value")
	//ErrConstructor is given when there's a mandatory field is not initialized correctly, and prompts to use the constructor
	ErrConstructor = errors.New("is a mandatory field and has a default value, did you use the constructor?")
	//ErrFieldRequired is given when a field is required
	ErrFieldRequired = errors.New("is a required field")
	//ErrServiceClass is given when there's an invalid service class code
	ErrServiceClass = errors.New("is an invalid Service Class Code")
	//ErrSECCode is given when there's an invalid standard entry class code
	ErrSECCode = errors.New("is an invalid Standard Entry Class Code")
	//ErrOrigStatusCode is given when there's an invalid originator status code
	ErrOrigStatusCode = errors.New("is an invalid Originator Status Code")
	//ErrAddendaTypeCode is given when there's an invalid addenda type code
	ErrAddendaTypeCode = errors.New("is an invalid Addenda Type Code")
	//ErrTransactionCode is given when there's an invalid transaction code
	ErrTransactionCode = errors.New("is an invalid Transaction Code")
	//ErrIdentificationNumber is given when there's an invalid identification number
	ErrIdentificationNumber = errors.New("is an invalid identification number")
	//ErrCardTransactionType  is given when there's an invalid card transaction type
	ErrCardTransactionType = errors.New("is an invalid Card Transaction Type")
	//ErrValidMonth is given when there's an invalid month
	ErrValidMonth = errors.New("is an invalid month")
	//ErrValidDay is given when there's an invalid day
	ErrValidDay = errors.New("is an invalid day")
	//ErrValidYear is given when there's an invalid year
	ErrValidYear = errors.New("is an invalid year")
	// ErrValidState is the error given when a field has an invalid US state or territory
	ErrValidState = errors.New("is an invalid US state or territory")
	// ErrValidISO3166 is the error given when a field has an invalid ISO 3166-1-alpha-2 code
	ErrValidISO3166 = errors.New("is an invalid ISO 3166-1-alpha-2 code")
	// ErrValidISO4217 is the error given when a field has an invalid ISO 4217 code
	ErrValidISO4217 = errors.New("is an invalid ISO 4217 code")

	// EntryDetail errors

	// ErrNegativeAmount is the error given when an Amount value is negaitve, which is
	// against NACHA rules and guidelines.
	ErrNegativeAmount = errors.New("amounts cannot be negative")

	// Addenda errors

	// ErrAddenda98ChangeCode is given when there's an invalid addenda change code
	ErrAddenda98ChangeCode = errors.New("found is not a valid addenda Change Code")
	// ErrAddenda98RefusedChangeCode is given when there's an invalid addenda refused change code
	ErrAddenda98RefusedChangeCode = errors.New("found is not a valid addenda Refused Change Code")
	// ErrAddenda98RefusedTraceSequenceNumber is given when there's an invalid addenda trace sequence number
	ErrAddenda98RefusedTraceSequenceNumber = errors.New("found is not a valid addenda trace sequence number")
	// ErrAddenda98CorrectedData is given when the corrected data does not corespond to the change code
	ErrAddenda98CorrectedData = errors.New("must contain the corrected information corresponding to the Change Code")
	// ErrAddenda99ReturnCode is given when there's an invalid return code
	ErrAddenda99ReturnCode = errors.New("found is not a valid return code")
	// ErrAddenda99DishonoredReturnCode is given when there's an invalid dishonored return code
	ErrAddenda99DishonoredReturnCode = errors.New("found is not a valid dishonored return code")
	// ErrAddenda99ContestedReturnCode is given when there's an invalid dishonored return code
	ErrAddenda99ContestedReturnCode = errors.New("found is not a valid contested dishonored return code")
	// ErrBatchCORAddenda is given when an entry in a COR batch does not have an addenda98
	ErrBatchCORAddenda = errors.New("one Addenda98 or Addenda98Refused record is required for each entry in SEC Type COR")

	// FileHeader errors

	// ErrRecordSize is given when there's an invalid record size
	ErrRecordSize = errors.New("is not 094")
	// ErrBlockingFactor is given when there's an invalid blocking factor
	ErrBlockingFactor = errors.New("is not 10")
	// ErrFormatCode is given when there's an invalid format code
	ErrFormatCode = errors.New("is not 1")

	// IAT

	// ErrForeignExchangeIndicator is given when there's an invalid foreign exchange indicator
	ErrForeignExchangeIndicator = errors.New("is an invalid Foreign Exchange Indicator")
	// ErrForeignExchangeReferenceIndicator is given when there's an invalid foreign exchange reference indicator
	ErrForeignExchangeReferenceIndicator = errors.New("is an invalid Foreign Exchange Reference Indicator")
	// ErrTransactionTypeCode is given when there's an invalid transaction type code
	ErrTransactionTypeCode = errors.New("is an invalid Addenda10 Transaction Type Code")
	// ErrIDNumberQualifier is given when there's an invalid identification number qualifier
	ErrIDNumberQualifier = errors.New("is an invalid Identification Number Qualifier")
	// ErrIATBatchAddendaIndicator is given when there's an invalid addenda record for an IAT batch
	ErrIATBatchAddendaIndicator = errors.New("is invalid for addenda record(s) found")
)

// FieldError is returned for errors at a field level in a record
type FieldError struct {
	FieldName string      // field name where error happened
	Value     interface{} // value that cause error
	Err       error       // context of the error.
	Msg       string      // deprecated
}

// Error message is constructed
// FieldName Msg Value
// Example1: BatchCount $% has none alphanumeric characters
// Example2: BatchCount 5 is out-of-balance with file count 6
func (e *FieldError) Error() string {
	return fmt.Sprintf("%s %v %s", e.FieldName, e.Value, e.Err)
}

// Unwrap implements the base.UnwrappableError interface for FieldError
func (e *FieldError) Unwrap() error {
	return e.Err
}

func fieldError(field string, err error, values ...interface{}) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(*FieldError); ok {
		return err
	}
	fe := FieldError{
		FieldName: field,
		Err:       err,
	}
	// only the first value counts
	if len(values) > 0 {
		fe.Value = values[0]
	}
	return &fe
}

// ErrValidCheckDigit is the error given when the observed check digit does not match the calculated one
type ErrValidCheckDigit struct {
	Message              string
	CalculatedCheckDigit int
}

// NewErrValidCheckDigit creates a new error of the ErrValidCheckDigit type
func NewErrValidCheckDigit(digit int) ErrValidCheckDigit {
	return ErrValidCheckDigit{
		Message:              fmt.Sprintf("does not match calculated check digit %v", digit),
		CalculatedCheckDigit: digit,
	}
}

func (e ErrValidCheckDigit) Error() string {
	return e.Message
}

// ErrValidFieldLength is the error given when the field does not have the correct length
type ErrValidFieldLength struct {
	Message        string
	ExpectedLength int
}

// NewErrValidFieldLength creates a new error of the ErrValidFieldLength type
func NewErrValidFieldLength(expectedLength int) ErrValidFieldLength {
	return ErrValidFieldLength{
		Message:        fmt.Sprintf("is not length %v", expectedLength),
		ExpectedLength: expectedLength,
	}
}

func (e ErrValidFieldLength) Error() string {
	return e.Message
}

// ErrRecordType is the error given when the field does not have the right record type
type ErrRecordType struct {
	Message      string
	ExpectedType int
}

// NewErrRecordType creates a new error of the ErrRecordType type
func NewErrRecordType(expectedType int) ErrRecordType {
	return ErrRecordType{
		Message:      fmt.Sprintf("received expecting %v", expectedType),
		ExpectedType: expectedType,
	}
}

func (e ErrRecordType) Error() string {
	return e.Message
}
