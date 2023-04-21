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
	// ErrBatchNoEntries is the error given when a batch doesn't have any entries
	ErrBatchNoEntries = errors.New("must have Entry Record(s) to be built")
	// ErrBatchADVCount is the error given when an ADV batch has too many entries
	ErrBatchADVCount = errors.New("there can be a maximum of 9999 ADV Sequence Numbers (ADV Entry Detail Records)")
	// ErrBatchAddendaIndicator is the error given when the addenda indicator is incorrectly set
	ErrBatchAddendaIndicator = errors.New("is 0 but found addenda record(s)")
	// ErrBatchOriginatorDNE is the error given when a non-government agency tries to originate a DNE
	ErrBatchOriginatorDNE = errors.New("only government agencies (originator status code 2) can originate a DNE")
	// ErrBatchInvalidCardTransactionType is the error given when a card transaction type is invalid
	ErrBatchInvalidCardTransactionType = errors.New("invalid card transaction type")
	// ErrBatchDebitOnly is the error given when a batch which can only have debits has a credit
	ErrBatchDebitOnly = errors.New("this batch type does not allow credit transaction codes")
	// ErrBatchCheckSerialNumber is the error given when a batch requires check serial numbers, but it is missing
	ErrBatchCheckSerialNumber = errors.New("this batch type requires entries to have Check Serial Numbers")
	// ErrBatchSECType is the error given when the batch's header has the wrong SEC for its type
	ErrBatchSECType = errors.New("header SEC does not match this batch's type")
	// ErrBatchServiceClassCode is the error given when the batch's header has the wrong SCC for its type
	ErrBatchServiceClassCode = errors.New("header SCC is not valid for this batch's type")
	// ErrBatchTransactionCode is the error given when a batch has an invalid transaction code
	ErrBatchTransactionCode = errors.New("transaction code is not valid for this batch's type")
	// ErrBatchTransactionCodeAddenda is the error given when a batch has an addenda on a transaction code which doesn't allow it
	ErrBatchTransactionCodeAddenda = errors.New("this batch type does not allow an addenda for this transaction code")
	// ErrBatchAmountNonZero is the error given when an entry for a non-zero amount is in a batch that requires zero amount entries
	ErrBatchAmountNonZero = errors.New("this batch type requires that the amount is zero")
	// ErrBatchAmountZero is the error given when an entry for zero amount is in a batch that requires non-zero amount entries
	ErrBatchAmountZero = errors.New("this batch type requires that the amount is non-zero")
	// ErrBatchCompanyEntryDescriptionAutoenroll is the error given when the Company Entry Description is invalid (needs to be 'Autoenroll')
	ErrBatchCompanyEntryDescriptionAutoenroll = errors.New("this batch type requires that the Company Entry Description is AUTOENROLL")
	// ErrBatchCompanyEntryDescriptionREDEPCHECK is the error given when the Company Entry Description is invalid (needs to be 'REDEPCHECK')
	ErrBatchCompanyEntryDescriptionREDEPCHECK = errors.New("this batch type requires that the Company Entry Description is REDEPCHECK")
	// ErrBatchAddendaCategory is the error given when the addenda isn't allowed for the batch's type and category
	ErrBatchAddendaCategory = errors.New("this batch type does not allow this addenda for category")
)

// BatchError is an Error that describes batch validation issues
type BatchError struct {
	BatchNumber int
	BatchType   string
	FieldName   string
	FieldValue  interface{}
	Err         error
}

func (e *BatchError) Error() string {
	if e.FieldValue == nil {
		return fmt.Sprintf("batch #%d (%v) %s %v", e.BatchNumber, e.BatchType, e.FieldName, e.Err)
	}
	return fmt.Sprintf("batch #%d (%v) %s %v: %v", e.BatchNumber, e.BatchType, e.FieldName, e.Err, e.FieldValue)
}

// Unwrap implements the base.UnwrappableError interface for BatchError
func (e *BatchError) Unwrap() error {
	return e.Err
}

// error returns a new BatchError based on err
func (b *Batch) Error(field string, err error, values ...interface{}) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(*BatchError); ok {
		return err
	}
	be := BatchError{
		BatchNumber: b.Header.BatchNumber,
		BatchType:   b.Header.StandardEntryClassCode,
		FieldName:   field,
		Err:         err,
	}
	// only the first value counts
	if len(values) > 0 {
		be.FieldValue = values[0]
	}
	return &be
}

// error returns a new BatchError based on err
func (iatBatch *IATBatch) Error(field string, err error, values ...interface{}) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(*BatchError); ok {
		return err
	}
	be := BatchError{
		BatchNumber: iatBatch.Header.BatchNumber,
		BatchType:   iatBatch.Header.StandardEntryClassCode,
		FieldName:   field,
		Err:         err,
	}
	// only the first value counts
	if len(values) > 0 {
		be.FieldValue = values[0]
	}
	return &be
}

// ErrBatchHeaderControlEquality is the error given when the control record does not match the calculated value
type ErrBatchHeaderControlEquality struct {
	Message      string
	HeaderValue  interface{}
	ControlValue interface{}
}

// NewErrBatchHeaderControlEquality creates a new error of the ErrBatchHeaderControlEquality type
func NewErrBatchHeaderControlEquality(header, control interface{}) ErrBatchHeaderControlEquality {
	return ErrBatchHeaderControlEquality{
		Message:      fmt.Sprintf("header %v is not equal to control %v", header, control),
		HeaderValue:  header,
		ControlValue: control,
	}
}

func (e ErrBatchHeaderControlEquality) Error() string {
	return e.Message
}

// ErrBatchCalculatedControlEquality is the error given when the control record does not match the calculated value
type ErrBatchCalculatedControlEquality struct {
	Message         string
	CalculatedValue interface{}
	ControlValue    interface{}
}

// NewErrBatchCalculatedControlEquality creates a new error of the ErrBatchCalculatedControlEquality type
func NewErrBatchCalculatedControlEquality(calculated, control interface{}) ErrBatchCalculatedControlEquality {
	return ErrBatchCalculatedControlEquality{
		Message:         fmt.Sprintf("calculated %v is out-of-balance with batch control %v", calculated, control),
		CalculatedValue: calculated,
		ControlValue:    control,
	}
}

func (e ErrBatchCalculatedControlEquality) Error() string {
	return e.Message
}

// ErrBatchAscending is the error given when the trace numbers in a batch are not in ascending order
type ErrBatchAscending struct {
	Message       string
	PreviousTrace interface{}
	CurrentTrace  interface{}
}

// NewErrBatchAscending creates a new error of the ErrBatchAscending type
func NewErrBatchAscending(previous, current interface{}) ErrBatchAscending {
	return ErrBatchAscending{
		Message:       fmt.Sprintf("must be in ascending order, %v is less than or equal to last number %v", current, previous),
		PreviousTrace: previous,
		CurrentTrace:  current,
	}
}

func (e ErrBatchAscending) Error() string {
	return e.Message
}

// ErrBatchCategory is the error given when a batch has entires with two different categories
type ErrBatchCategory struct {
	Message   string
	CategoryA string
	CategoryB string
}

// NewErrBatchCategory creates a new error of the ErrBatchCategory type
func NewErrBatchCategory(categoryA, categoryB string) ErrBatchCategory {
	return ErrBatchCategory{
		Message:   fmt.Sprintf("%v category found in batch with category %v", categoryA, categoryB),
		CategoryA: categoryA,
		CategoryB: categoryB,
	}
}

func (e ErrBatchCategory) Error() string {
	return e.Message
}

// ErrBatchTraceNumberNotODFI is the error given when a batch's ODFI does not match an entry's trace number
type ErrBatchTraceNumberNotODFI struct {
	Message     string
	ODFI        string
	TraceNumber string
}

// NewErrBatchTraceNumberNotODFI creates a new error of the ErrBatchTraceNumberNotODFI type
func NewErrBatchTraceNumberNotODFI(odfi, trace string) ErrBatchTraceNumberNotODFI {
	return ErrBatchTraceNumberNotODFI{
		Message:     fmt.Sprintf("%v in header does not match entry trace number %v", odfi, trace),
		ODFI:        odfi,
		TraceNumber: trace,
	}
}

func (e ErrBatchTraceNumberNotODFI) Error() string {
	return e.Message
}

// ErrBatchAddendaTraceNumber is the error given when the entry detail sequence number doesn't match the trace number
type ErrBatchAddendaTraceNumber struct {
	Message           string
	EntryDetailNumber string
	TraceNumber       string
}

// NewErrBatchAddendaTraceNumber creates a new error of the ErrBatchAddendaTraceNumber type
func NewErrBatchAddendaTraceNumber(entryDetail, trace string) ErrBatchAddendaTraceNumber {
	return ErrBatchAddendaTraceNumber{
		Message:           fmt.Sprintf("%v does not match proceeding entry detail trace number %v", entryDetail, trace),
		EntryDetailNumber: entryDetail,
		TraceNumber:       trace,
	}
}

func (e ErrBatchAddendaTraceNumber) Error() string {
	return e.Message
}

// ErrBatchAddendaCount is the error given when there are too many addenda than allowed for the batch type
type ErrBatchAddendaCount struct {
	Message      string
	FoundCount   int
	AllowedCount int
}

// NewErrBatchAddendaCount creates a new error of the ErrBatchAddendaCount type
func NewErrBatchAddendaCount(found, allowed int) ErrBatchAddendaCount {
	return ErrBatchAddendaCount{
		Message:      fmt.Sprintf("%v addendum found where %v is allowed for this batch type", found, allowed),
		FoundCount:   found,
		AllowedCount: allowed,
	}
}

func (e ErrBatchAddendaCount) Error() string {
	return e.Message
}

// ErrBatchRequiredAddendaCount is the error given when the batch type requires a certain number of addenda, which is not met
type ErrBatchRequiredAddendaCount struct {
	Message       string
	FoundCount    int
	RequiredCount int
}

// NewErrBatchRequiredAddendaCount creates a new error of the ErrBatchRequiredAddendaCount type
func NewErrBatchRequiredAddendaCount(found, required int) ErrBatchRequiredAddendaCount {
	return ErrBatchRequiredAddendaCount{
		Message:       fmt.Sprintf("%v addendum found where %v are required for this batch type", found, required),
		FoundCount:    found,
		RequiredCount: required,
	}
}

func (e ErrBatchRequiredAddendaCount) Error() string {
	return e.Message
}

// ErrBatchExpectedAddendaCount is the error given when the batch type has entries with a field
// for the number of addenda, and a different number of addenda are foound
type ErrBatchExpectedAddendaCount struct {
	Message       string
	FoundCount    int
	ExpectedCount int
}

// NewErrBatchExpectedAddendaCount creates a new error of the ErrBatchExpectedAddendaCount type
func NewErrBatchExpectedAddendaCount(found, expected int) ErrBatchExpectedAddendaCount {
	return ErrBatchExpectedAddendaCount{
		Message:       fmt.Sprintf("%v addendum found where %v are expected for this batch type", found, expected),
		FoundCount:    found,
		ExpectedCount: expected,
	}
}

func (e ErrBatchExpectedAddendaCount) Error() string {
	return e.Message
}

// ErrBatchServiceClassTranCode is the error given when the transaction code is not valid for the batch's service class
type ErrBatchServiceClassTranCode struct {
	Message          string
	ServiceClassCode int
	TransactionCode  int
}

// NewErrBatchServiceClassTranCode creates a new error of the ErrBatchServiceClassTranCode type
func NewErrBatchServiceClassTranCode(serviceClassCode, transactionCode int) ErrBatchServiceClassTranCode {
	return ErrBatchServiceClassTranCode{
		Message:          fmt.Sprintf("service class code %v does not support transaction code %v", serviceClassCode, transactionCode),
		ServiceClassCode: serviceClassCode,
		TransactionCode:  transactionCode,
	}
}

func (e ErrBatchServiceClassTranCode) Error() string {
	return e.Message
}

// ErrBatchAmount is the error given when the amount exceeds the batch type's limit
type ErrBatchAmount struct {
	Message string
	Amount  int
	Limit   int
}

// NewErrBatchAmount creates a new error of the ErrBatchAmount type
func NewErrBatchAmount(amount, limit int) ErrBatchAmount {
	// TODO: pretty format the amounts to make it more readable
	return ErrBatchAmount{
		Message: fmt.Sprintf("amounts in this batch type are limited to %v, found amount of %v", limit, amount),
		Amount:  amount,
		Limit:   limit,
	}
}

func (e ErrBatchAmount) Error() string {
	return e.Message
}

// ErrBatchIATNOC is the error given when an IAT batch has an NOC, and there are invalid values
type ErrBatchIATNOC struct {
	Message  string
	Found    interface{}
	Expected interface{}
}

// NewErrBatchIATNOC creates a new error of the ErrBatchIATNOC type
func NewErrBatchIATNOC(found, expected interface{}) ErrBatchIATNOC {
	// TODO: pretty format the amounts to make it more readable
	return ErrBatchIATNOC{
		Message:  fmt.Sprintf("%v invalid for IAT NOC, should be %v", found, expected),
		Found:    found,
		Expected: expected,
	}
}

func (e ErrBatchIATNOC) Error() string {
	return e.Message
}
