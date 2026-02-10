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
	"errors"

	"github.com/moov-io/ach"
	"github.com/moov-io/base"
	"github.com/moov-io/base/log"

	"github.com/go-kit/kit/endpoint"
)

// v2 API response types with structured errors

// createFileResponseV2 is the v2 response for file creation with structured errors
type createFileResponseV2 struct {
	ID     string            `json:"id,omitempty"`
	File   *ach.File         `json:"file,omitempty"`
	Errors []ValidationError `json:"errors,omitempty"`
}

func (r createFileResponseV2) error() error {
	if len(r.Errors) > 0 {
		return errors.New("validation errors found")
	}
	return nil
}

// validateFileResponseV2 is the v2 response for file validation with structured errors
type validateFileResponseV2 struct {
	Valid  bool              `json:"valid"`
	Errors []ValidationError `json:"errors,omitempty"`
}

func (r validateFileResponseV2) error() error {
	if len(r.Errors) > 0 {
		return errors.New("validation errors found")
	}
	return nil
}

// createFileEndpointV2 creates a file and returns ALL validation errors in structured format
func createFileEndpointV2(s Service, r Repository, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(createFileRequest)
		if !ok {
			return createFileResponseV2{
				Errors: []ValidationError{{
					ErrorType: "RequestError",
					Message:   "invalid request type",
				}},
			}, nil
		}

		var validationErrors []ValidationError

		// Handle parse errors first
		if req.parseError != nil {
			if el, ok := req.parseError.(base.ErrorList); ok {
				validationErrors = ConvertErrorList(el)
			} else {
				validationErrors = append(validationErrors, ConvertError(req.parseError))
			}
		}

		// Run full validation if file was parsed
		if req.File != nil {
			// Record metrics
			if req.File.Header.ImmediateDestination != "" && req.File.Header.ImmediateOrigin != "" {
				filesCreated.With("destination", req.File.Header.ImmediateDestination, "origin", req.File.Header.ImmediateOrigin).Add(1)
			}

			// Create a random file ID if none was provided
			if req.File.ID == "" {
				req.File.ID = base.ID()
			}

			if req.validateOpts != nil {
				req.File.SetValidation(req.validateOpts)
			}

			// Run ValidateAll to collect all errors
			if errs := req.File.ValidateAll(); errs != nil {
				for _, err := range errs {
					validationErrors = append(validationErrors, ConvertError(err))
				}
			}

			// Store the file even if there are validation errors
			err := r.StoreFile(req.File)
			if err != nil {
				validationErrors = append(validationErrors, ValidationError{
					ErrorType: "StorageError",
					Message:   err.Error(),
				})
			}

			if logger != nil {
				logger := logger.With(log.Fields{
					"files":     log.String("createFileV2"),
					"requestID": log.String(req.requestID),
				})
				if len(validationErrors) > 0 {
					logger.Info().Log("create file with validation errors")
				} else {
					logger.Info().Log("create file")
				}
			}
		}

		return createFileResponseV2{
			ID:     req.File.ID,
			File:   req.File,
			Errors: validationErrors,
		}, nil
	}
}

// validateFileEndpointV2 validates a file and returns ALL errors in structured format
func validateFileEndpointV2(s Service, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(validateFileRequest)
		if !ok {
			return validateFileResponseV2{
				Valid: false,
				Errors: []ValidationError{{
					ErrorType: "RequestError",
					Message:   "invalid request type",
				}},
			}, nil
		}

		file, err := s.GetFile(req.ID)
		if err != nil {
			return validateFileResponseV2{
				Valid: false,
				Errors: []ValidationError{{
					ErrorType: "FileError",
					Message:   err.Error(),
				}},
			}, nil
		}

		// Use ValidateAllWith to collect all errors
		errs := file.ValidateAllWith(req.opts)
		if errs == nil || errs.Empty() {
			if logger != nil {
				logger.With(log.Fields{
					"files":     log.String("validateFileV2"),
					"requestID": log.String(req.requestID),
				}).Info().Log("file is valid")
			}
			return validateFileResponseV2{Valid: true}, nil
		}

		var validationErrors []ValidationError
		for _, e := range errs {
			validationErrors = append(validationErrors, ConvertError(e))
		}

		if logger != nil {
			logger.With(log.Fields{
				"files":     log.String("validateFileV2"),
				"requestID": log.String(req.requestID),
			}).Info().Logf("file has %d validation errors", len(validationErrors))
		}

		return validateFileResponseV2{
			Valid:  false,
			Errors: validationErrors,
		}, nil
	}
}
