// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"errors"
	"fmt"
)

var (
	// ErrBatchNoEntries is the error given when a batch doesn't have any entries
	ErrBatchNoEntries = errors.New("must have Entry Record(s) to be built")
	// ErrBatchADVCount is the error given when an ADV batch has too many entries
	ErrBatchADVCount = errors.New("There can be a maximum of 9999 ADV Sequence Numbers (ADV Entry Detail Records)")
	// ErrBatchAddendaIndicator is the error given when the addenda indicator is incorrectly set
	ErrBatchAddendaIndicator = errors.New("is 0 but found addenda record(s)")
)

// BatchError is an Error that describes batch validation issues
type BatchError struct {
	BatchNumber int
	BatchType   string
	FieldName   string
	FieldValue  interface{}
	Msg         string // deprecated
	Err         error
}

func (e *BatchError) Error() string {
	if e.FieldValue == nil {
		return fmt.Sprintf("BatchNumber %d (%v) %s %v", e.BatchNumber, e.BatchType, e.FieldName, e.Err)
	}
	return fmt.Sprintf("BatchNumber %d (%v) %s %v: %v", e.BatchNumber, e.BatchType, e.FieldName, e.Err, e.FieldValue)
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
func (b *IATBatch) Error(field string, err error, values ...interface{}) error {
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
		Message:       fmt.Sprintf("%v is less than last %v. Must be in ascending order", current, previous),
		PreviousTrace: previous,
		CurrentTrace:  current,
	}
}

func (e ErrBatchAscending) Error() string {
	return e.Message
}
