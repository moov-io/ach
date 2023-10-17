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
	// ErrFileTooLong is the error given when a file exceeds the maximum possible length
	ErrFileTooLong = errors.New("file exceeds maximum possible number of lines")
	// ErrFileHeader is the error given if there is the wrong number of file headers
	ErrFileHeader = errors.New("none or more than one file headers exists")
	// ErrFileControl is the error given if there is the wrong number of file control records
	ErrFileControl = errors.New("none or more than one file control exists")
	// ErrFileEntryOutsideBatch is the error given if an entry is outside of a batch
	ErrFileEntryOutsideBatch = errors.New("entry outside of batch")
	// ErrFileAddendaOutsideBatch is the error given if an addenda is outside of a batch
	ErrFileAddendaOutsideBatch = errors.New("addenda outside of batch")
	// ErrFileAddendaOutsideEntry is the error given if an addenda is outside of an entry
	ErrFileAddendaOutsideEntry = errors.New("addenda outside of entry")
	// ErrFileBatchControlOutsideBatch is the error given if a batch control record is outside of a batch
	ErrFileBatchControlOutsideBatch = errors.New("batch control outside of batch")
	// ErrFileConsecutiveBatchHeaders is the error given when multiple batch header records occur in sequence
	ErrFileConsecutiveBatchHeaders = errors.New("consecutive Batch Headers in file")
	// ErrFileADVOnly is the error given if an ADV only file has a non-ADV batch
	ErrFileADVOnly = errors.New("file can only have ADV Batches")
	// ErrFileIATSEC is the error given if an IAT batch uses the normal NewBatch
	ErrFileIATSEC = errors.New("IAT Standard Entry Class Code should use iatBatch")
	// ErrFileNoBatches is the error given if a file has no batches
	ErrFileNoBatches = errors.New("must have []*Batches or []*IATBatches to be built")

	ErrInvalidJSON = errors.New("invalid JSON")
)

// RecordWrongLengthErr is the error given when a record is the wrong length
type RecordWrongLengthErr struct {
	Message string
	Length  int
}

// NewRecordWrongLengthErr creates a new error of the RecordWrongLengthErr type
func NewRecordWrongLengthErr(length int) RecordWrongLengthErr {
	return RecordWrongLengthErr{
		Message: fmt.Sprintf("must be 94 characters and found %d", length),
		Length:  length,
	}
}

func (e RecordWrongLengthErr) Error() string {
	return e.Message
}

// ErrUnknownRecordType is the error given when a record does not have a known type
type ErrUnknownRecordType struct {
	Message string
	Type    string
}

// NewErrUnknownRecordType creates a new error of the ErrUnknownRecordType type
func NewErrUnknownRecordType(recordType string) ErrUnknownRecordType {
	return ErrUnknownRecordType{
		Message: fmt.Sprintf("%s is an unknown record type", recordType),
		Type:    recordType,
	}
}

func (e ErrUnknownRecordType) Error() string {
	return e.Message
}

// ErrFileUnknownSEC is the error given when a record does not have a known type
type ErrFileUnknownSEC struct {
	Message string
	SEC     string
}

// NewErrFileUnknownSEC creates a new error of the ErrFileUnknownSEC type
func NewErrFileUnknownSEC(secType string) ErrFileUnknownSEC {
	return ErrFileUnknownSEC{
		Message: fmt.Sprintf("%s Standard Entry Class Code is not implemented", secType),
		SEC:     secType,
	}
}

func (e ErrFileUnknownSEC) Error() string {
	return e.Message
}

// ErrFileCalculatedControlEquality is the error given when the control record does not match the calculated value
type ErrFileCalculatedControlEquality struct {
	Message         string
	Field           string
	CalculatedValue int
	ControlValue    int
}

// NewErrFileCalculatedControlEquality creates a new error of the ErrFileCalculatedControlEquality type
func NewErrFileCalculatedControlEquality(field string, calculated, control int) ErrFileCalculatedControlEquality {
	return ErrFileCalculatedControlEquality{
		Message:         fmt.Sprintf("%v calculated %v is out-of-balance with file control %v", field, calculated, control),
		Field:           field,
		CalculatedValue: calculated,
		ControlValue:    control,
	}
}

func (e ErrFileCalculatedControlEquality) Error() string {
	return e.Message
}

// ErrFileBatchNumberAscending is the error given when the batch numbers in a file are not in ascending order
type ErrFileBatchNumberAscending struct {
	Message       string
	PreviousBatch int
	CurrentBatch  int
}

// NewErrFileBatchNumberAscending creates a new error of the ErrFileBatchNumberAscending type
func NewErrFileBatchNumberAscending(previous, current int) ErrFileBatchNumberAscending {
	return ErrFileBatchNumberAscending{
		Message:       fmt.Sprintf("Batch numbers must be in ascending order, batch %v is less than or equal to the previous batch: %v", current, previous),
		PreviousBatch: previous,
		CurrentBatch:  current,
	}
}

func (e ErrFileBatchNumberAscending) Error() string {
	return e.Message
}
