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

package server

import (
	"errors"

	"github.com/moov-io/ach"
	"github.com/moov-io/base"
)

// ValidationError represents a single structured validation error for the v2 API.
// It provides detailed, machine-readable information about validation failures.
type ValidationError struct {
	// LineNumber is the line number where the error occurred (from parsing)
	LineNumber int `json:"lineNumber,omitempty"`

	// RecordType is the type of record being processed when the error occurred
	RecordType string `json:"recordType,omitempty"`

	// ErrorType categorizes the error (e.g., "FieldError", "BatchError", "FileError")
	ErrorType string `json:"errorType"`

	// BatchNumber identifies which batch in the file had the error
	BatchNumber int `json:"batchNumber,omitempty"`

	// BatchType is the SEC code of the batch (e.g., "PPD", "WEB", "CCD")
	BatchType string `json:"batchType,omitempty"`

	// FieldName is the name of the field that failed validation
	FieldName string `json:"fieldName,omitempty"`

	// FieldValue is the value that caused the validation error
	FieldValue interface{} `json:"fieldValue,omitempty"`

	// Message is a human-readable description of the error
	Message string `json:"message"`
}

// ConvertError transforms various ACH error types into a ValidationError.
// It handles FieldError, BatchError, FileError, ParseError, and generic errors.
func ConvertError(err error) ValidationError {
	if err == nil {
		return ValidationError{}
	}

	ve := ValidationError{
		Message:   err.Error(),
		ErrorType: "Error",
	}

	// Check for ParseError wrapper first and extract line/record info
	var parseErr *base.ParseError
	if errors.As(err, &parseErr) {
		ve.LineNumber = parseErr.Line
		ve.RecordType = parseErr.Record
		ve.ErrorType = "ParseError"
		// Unwrap and continue processing the inner error
		if parseErr.Err != nil {
			err = parseErr.Err
		}
	}

	// Check specific ACH error types
	var fieldErr *ach.FieldError
	if errors.As(err, &fieldErr) {
		ve.ErrorType = "FieldError"
		ve.FieldName = fieldErr.FieldName
		ve.FieldValue = fieldErr.Value
		ve.Message = fieldErr.Error()
		return ve
	}

	var batchErr *ach.BatchError
	if errors.As(err, &batchErr) {
		ve.ErrorType = "BatchError"
		ve.BatchNumber = batchErr.BatchNumber
		ve.BatchType = batchErr.BatchType
		ve.FieldName = batchErr.FieldName
		ve.FieldValue = batchErr.FieldValue
		ve.Message = batchErr.Error()
		return ve
	}

	var fileErr *ach.FileError
	if errors.As(err, &fileErr) {
		ve.ErrorType = "FileError"
		ve.FieldName = fileErr.FieldName
		ve.Message = fileErr.Error()
		return ve
	}

	// Check for specific error types that have additional structured info
	var checkDigitErr ach.ErrValidCheckDigit
	if errors.As(err, &checkDigitErr) {
		ve.ErrorType = "FieldError"
		ve.Message = checkDigitErr.Error()
		return ve
	}

	var fieldLengthErr ach.ErrValidFieldLength
	if errors.As(err, &fieldLengthErr) {
		ve.ErrorType = "FieldError"
		ve.Message = fieldLengthErr.Error()
		return ve
	}

	var recordTypeErr ach.ErrRecordType
	if errors.As(err, &recordTypeErr) {
		ve.ErrorType = "FieldError"
		ve.Message = recordTypeErr.Error()
		return ve
	}

	// Batch-specific structured errors
	var batchHeaderControlErr ach.ErrBatchHeaderControlEquality
	if errors.As(err, &batchHeaderControlErr) {
		ve.ErrorType = "BatchError"
		ve.Message = batchHeaderControlErr.Error()
		return ve
	}

	var batchCalcControlErr ach.ErrBatchCalculatedControlEquality
	if errors.As(err, &batchCalcControlErr) {
		ve.ErrorType = "BatchError"
		ve.Message = batchCalcControlErr.Error()
		return ve
	}

	var batchAscendingErr ach.ErrBatchAscending
	if errors.As(err, &batchAscendingErr) {
		ve.ErrorType = "BatchError"
		ve.Message = batchAscendingErr.Error()
		return ve
	}

	var batchCategoryErr ach.ErrBatchCategory
	if errors.As(err, &batchCategoryErr) {
		ve.ErrorType = "BatchError"
		ve.Message = batchCategoryErr.Error()
		return ve
	}

	var batchTraceErr ach.ErrBatchTraceNumberNotODFI
	if errors.As(err, &batchTraceErr) {
		ve.ErrorType = "BatchError"
		ve.Message = batchTraceErr.Error()
		return ve
	}

	var batchAddendaTraceErr ach.ErrBatchAddendaTraceNumber
	if errors.As(err, &batchAddendaTraceErr) {
		ve.ErrorType = "BatchError"
		ve.Message = batchAddendaTraceErr.Error()
		return ve
	}

	var batchAddendaCountErr ach.ErrBatchAddendaCount
	if errors.As(err, &batchAddendaCountErr) {
		ve.ErrorType = "BatchError"
		ve.Message = batchAddendaCountErr.Error()
		return ve
	}

	var batchRequiredAddendaErr ach.ErrBatchRequiredAddendaCount
	if errors.As(err, &batchRequiredAddendaErr) {
		ve.ErrorType = "BatchError"
		ve.Message = batchRequiredAddendaErr.Error()
		return ve
	}

	var batchExpectedAddendaErr ach.ErrBatchExpectedAddendaCount
	if errors.As(err, &batchExpectedAddendaErr) {
		ve.ErrorType = "BatchError"
		ve.Message = batchExpectedAddendaErr.Error()
		return ve
	}

	var batchServiceClassErr ach.ErrBatchServiceClassTranCode
	if errors.As(err, &batchServiceClassErr) {
		ve.ErrorType = "BatchError"
		ve.Message = batchServiceClassErr.Error()
		return ve
	}

	var batchAmountErr ach.ErrBatchAmount
	if errors.As(err, &batchAmountErr) {
		ve.ErrorType = "BatchError"
		ve.Message = batchAmountErr.Error()
		return ve
	}

	var batchIATNOCErr ach.ErrBatchIATNOC
	if errors.As(err, &batchIATNOCErr) {
		ve.ErrorType = "BatchError"
		ve.Message = batchIATNOCErr.Error()
		return ve
	}

	return ve
}

// ConvertErrors transforms a slice of errors into ValidationErrors.
func ConvertErrors(errs []error) []ValidationError {
	if len(errs) == 0 {
		return nil
	}

	result := make([]ValidationError, 0, len(errs))
	for _, err := range errs {
		if err != nil {
			result = append(result, ConvertError(err))
		}
	}
	return result
}

// ConvertErrorList transforms a base.ErrorList into ValidationErrors.
func ConvertErrorList(el base.ErrorList) []ValidationError {
	if el.Empty() {
		return nil
	}
	return ConvertErrors(el)
}
