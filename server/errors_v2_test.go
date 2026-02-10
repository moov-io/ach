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
	"testing"

	"github.com/moov-io/ach"
	"github.com/moov-io/base"

	"github.com/stretchr/testify/require"
)

func TestConvertError__FieldError(t *testing.T) {
	fieldErr := &ach.FieldError{
		FieldName: "Amount",
		Value:     "abc",
		Msg:       "invalid amount",
	}

	result := ConvertError(fieldErr)

	require.Equal(t, "FieldError", result.ErrorType)
	require.Equal(t, "Amount", result.FieldName)
	require.Equal(t, "abc", result.FieldValue)
	require.NotEmpty(t, result.Message)
}

func TestConvertError__BatchError(t *testing.T) {
	batchErr := &ach.BatchError{
		BatchNumber: 1,
		FieldName:   "EntryCount",
		Err:         errors.New("count mismatch"),
	}

	result := ConvertError(batchErr)

	require.Equal(t, "BatchError", result.ErrorType)
	require.Equal(t, 1, result.BatchNumber)
	require.Equal(t, "EntryCount", result.FieldName)
	require.Contains(t, result.Message, "count mismatch")
}

func TestConvertError__FileError(t *testing.T) {
	fileErr := &ach.FileError{
		FieldName: "ImmediateDestination",
		Msg:       "invalid routing number",
	}

	result := ConvertError(fileErr)

	require.Equal(t, "FileError", result.ErrorType)
	require.Equal(t, "ImmediateDestination", result.FieldName)
	require.Contains(t, result.Message, "invalid routing number")
}

func TestConvertError__ParseError(t *testing.T) {
	parseErr := &base.ParseError{
		Line:   5,
		Record: "EntryDetail",
		Err:    errors.New("parse failed"),
	}

	result := ConvertError(parseErr)

	require.Equal(t, "ParseError", result.ErrorType)
	require.Equal(t, 5, result.LineNumber)
	require.Equal(t, "EntryDetail", result.RecordType)
	require.Contains(t, result.Message, "parse failed")
}

func TestConvertError__GenericError(t *testing.T) {
	err := errors.New("generic error")

	result := ConvertError(err)

	require.Equal(t, "Error", result.ErrorType)
	require.Equal(t, "generic error", result.Message)
}

func TestConvertErrors(t *testing.T) {
	errs := []error{
		&ach.FieldError{FieldName: "Amount", Msg: "error 1"},
		&ach.BatchError{BatchNumber: 2, Err: errors.New("error 2")},
		errors.New("error 3"),
	}

	result := ConvertErrors(errs)

	require.Len(t, result, 3)
	require.Equal(t, "FieldError", result[0].ErrorType)
	require.Equal(t, "BatchError", result[1].ErrorType)
	require.Equal(t, "Error", result[2].ErrorType)
}

func TestConvertErrorList(t *testing.T) {
	var el base.ErrorList
	el.Add(&ach.FieldError{FieldName: "Field1", Msg: "msg1"})
	el.Add(&ach.FieldError{FieldName: "Field2", Msg: "msg2"})

	result := ConvertErrorList(el)

	require.Len(t, result, 2)
	require.Equal(t, "Field1", result[0].FieldName)
	require.Equal(t, "Field2", result[1].FieldName)
}

func TestConvertError__NilError(t *testing.T) {
	result := ConvertError(nil)

	require.Empty(t, result.ErrorType)
	require.Empty(t, result.Message)
}
