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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/moov-io/ach"
	"github.com/moov-io/base"
)

type StructuredErrorResponse struct {
	Errors []StructuredError `json:"errors"`
}

type StructuredError struct {
	LineNumber int    `json:"lineNumber"`
	RecordType string `json:"recordType"`
	ErrorType  string `json:"errorType"`
	FieldName  string `json:"fieldName"`
	Message    string `json:"message"`
}

func buildStructuredErrors(err error) *StructuredErrorResponse {
	if err == nil {
		return nil
	}

	var el base.ErrorList
	if errors.As(err, &el) {
		var errs []StructuredError
		for i := range el {
			if list, ok := el[i].(base.ErrorList); ok {
				resp := buildStructuredErrors(list)
				if resp != nil {
					errs = append(errs, resp.Errors...)
				}
			}

			var fe *ach.FieldError
			if errors.As(el[i], &fe) {
				errs = append(errs, StructuredError{
					ErrorType: "FieldError",
					FieldName: fe.FieldName,
					Message:   fe.Err.Error(),
				})
			}
		}
		return &StructuredErrorResponse{Errors: errs}
	}

	var fe *ach.FieldError
	if errors.As(err, &fe) {
		return &StructuredErrorResponse{
			Errors: []StructuredError{
				{
					ErrorType: "FieldError",
					FieldName: fe.FieldName,
					Message:   fe.Err.Error(),
				},
			},
		}
	}

	// fallback
	return &StructuredErrorResponse{
		Errors: []StructuredError{
			{
				Message: err.Error(),
			},
		},
	}
}

func encodeStructuredError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		err = ErrFoundABug
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))

	resp := buildStructuredErrors(err)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.Write([]byte(fmt.Sprintf("problem rendering structured json: %v", err)))
	}
}
