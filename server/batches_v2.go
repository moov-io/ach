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

	"github.com/moov-io/base/log"

	"github.com/go-kit/kit/endpoint"
)

// v2 API batch response types with structured errors

// createBatchResponseV2 is the v2 response for batch creation with structured errors
type createBatchResponseV2 struct {
	ID     string            `json:"id,omitempty"`
	Errors []ValidationError `json:"errors,omitempty"`
}

func (r createBatchResponseV2) error() error {
	if len(r.Errors) > 0 {
		return errors.New("validation errors found")
	}
	return nil
}

// createBatchEndpointV2 creates a batch and returns ALL validation errors in structured format
func createBatchEndpointV2(s Service, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(createBatchRequest)
		if !ok {
			return createBatchResponseV2{
				Errors: []ValidationError{{
					ErrorType: "RequestError",
					Message:   "invalid request type",
				}},
			}, nil
		}

		var validationErrors []ValidationError

		// Run ValidateAll on the batch to collect all errors
		if req.Batch != nil {
			if errs := req.Batch.ValidateAll(); errs != nil {
				for _, err := range errs {
					validationErrors = append(validationErrors, ConvertError(err))
				}
			}
		}

		// If there are validation errors, return them without creating the batch
		if len(validationErrors) > 0 {
			if logger != nil {
				logger.With(log.Fields{
					"batches":   log.String("createBatchV2"),
					"file":      log.String(req.FileID),
					"requestID": log.String(req.requestID),
				}).Info().Logf("batch has %d validation errors", len(validationErrors))
			}
			return createBatchResponseV2{
				Errors: validationErrors,
			}, nil
		}

		// Create the batch
		id, err := s.CreateBatch(req.FileID, req.Batch)
		if err != nil {
			validationErrors = append(validationErrors, ValidationError{
				ErrorType: "ServiceError",
				Message:   err.Error(),
			})
			if logger != nil {
				logger.With(log.Fields{
					"batches":   log.String("createBatchV2"),
					"file":      log.String(req.FileID),
					"requestID": log.String(req.requestID),
				}).Error().LogError(err)
			}
			return createBatchResponseV2{
				Errors: validationErrors,
			}, nil
		}

		if logger != nil {
			logger.With(log.Fields{
				"batches":   log.String("createBatchV2"),
				"file":      log.String(req.FileID),
				"requestID": log.String(req.requestID),
			}).Info().Log("batch created")
		}

		return createBatchResponseV2{
			ID: id,
		}, nil
	}
}
