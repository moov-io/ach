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
	"fmt"
	"io"
	"net/http"

	"github.com/moov-io/ach"
	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/base/log"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

type createBatchRequest struct {
	FileID string
	Batch  ach.Batcher

	requestID string
}

type createBatchResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error"`
}

func (r createBatchResponse) error() error { return r.Err }

func createBatchEndpoint(s Service, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(createBatchRequest)
		if !ok {
			err := errors.New("invalid request")
			return createBatchResponse{
				Err: err,
			}, err
		}

		id, err := s.CreateBatch(req.FileID, req.Batch)

		if logger != nil {
			logger := logger.With(log.Fields{
				"batches":   log.String("createBatch"),
				"file":      log.String(req.FileID),
				"requestID": log.String(req.requestID),
			})
			if err != nil {
				logger.Error().LogError(err)
			} else {
				logger.Info().Log("creating batch")
			}
		}

		return createBatchResponse{
			ID:  id,
			Err: err,
		}, nil
	}
}

func decodeCreateBatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createBatchRequest
	req.requestID = moovhttp.GetRequestID(r)

	vars := mux.Vars(r)
	id, ok := vars["fileID"]
	if !ok {
		return nil, ErrBadRouting
	}
	req.FileID = id

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	// In order to use FileFromJSON we need a populated JSON structure that can be parsed.
	// We're going to copy the body into this shim to parse the Batch, otherwise we'd have
	// to copy/export the logic of reading batches from their JSON representation.
	fileContentsShim := `{"fileHeader": {
  "immediateOriginName": "Test Sender",
  "immediateDestinationName": "Test Dest",
  "fileIDModifier": "1",
  "fileCreationTime": "0437",
  "fileCreationDate": "200217",
  "immediateOrigin": "123456780",
  "immediateDestination": "987654320",
  "id": ""
}, "batches":[%v] }`
	file, err := ach.FileFromJSON([]byte(fmt.Sprintf(fileContentsShim, string(body))))
	if err != nil {
		return nil, err
	}
	if len(file.Batches) == 1 {
		req.Batch = file.Batches[0]
	}
	if req.Batch == nil {
		return nil, errors.New("no Batch provided")
	}
	if err := req.Batch.Validate(); err != nil {
		return nil, err
	}
	return req, nil
}

type getBatchesRequest struct {
	fileID string

	requestID string
}

type getBatchesResponse struct {
	// TODO(adam): change this to JSON encode without wrapper {"batches": [..]}
	// We don't wrap json objects in other responses, so why here?
	Batches []ach.Batcher `json:"batches"`
	Err     error         `json:"error"`
}

func (r getBatchesResponse) count() int { return len(r.Batches) }

func (r getBatchesResponse) error() error { return r.Err }

func decodeGetBatchesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req getBatchesRequest
	req.requestID = moovhttp.GetRequestID(r)

	vars := mux.Vars(r)
	id, ok := vars["fileID"]
	if !ok {
		return nil, ErrBadRouting
	}
	req.fileID = id
	return req, nil
}

func getBatchesEndpoint(s Service, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(getBatchesRequest)
		if !ok {
			err := errors.New("invalid request")
			return getBatchesResponse{
				Err: err,
			}, err
		}

		if logger != nil {
			logger.With(log.Fields{
				"batches":   log.String("getBatches"),
				"file":      log.String(req.fileID),
				"requestID": log.String(req.requestID),
			}).Log("get batches")
		}

		return getBatchesResponse{
			Batches: s.GetBatches(req.fileID),
			Err:     nil,
		}, nil
	}
}

type getBatchRequest struct {
	fileID  string
	batchID string

	requestID string
}

type getBatchResponse struct {
	Batch ach.Batcher `json:"batch"`
	Err   error       `json:"error"`
}

func (r getBatchResponse) error() error { return r.Err }

func decodeGetBatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req getBatchRequest
	req.requestID = moovhttp.GetRequestID(r)

	vars := mux.Vars(r)
	fileID, ok := vars["fileID"]
	if !ok {
		return nil, ErrBadRouting
	}
	batchID, ok := vars["batchID"]
	if !ok {
		return nil, ErrBadRouting
	}

	req.fileID = fileID
	req.batchID = batchID
	return req, nil
}

func getBatchEndpoint(s Service, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(getBatchRequest)
		if !ok {
			err := errors.New("invalid request")
			return getBatchResponse{
				Err: err,
			}, err
		}

		batch, err := s.GetBatch(req.fileID, req.batchID)

		if logger != nil {
			logger := logger.With(log.Fields{
				"batches":   log.String("getBatch"),
				"file":      log.String(req.fileID),
				"requestID": log.String(req.requestID),
			})
			if err != nil {
				logger.Error().LogError(err)
			} else {
				logger.Info().Log("get batch")
			}
		}

		return getBatchResponse{
			Batch: batch,
			Err:   err,
		}, nil
	}
}

type deleteBatchRequest struct {
	fileID  string
	batchID string

	requestID string
}

type deleteBatchResponse struct {
	Err error `json:"error"`
}

func (r deleteBatchResponse) error() error { return r.Err }

func decodeDeleteBatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req deleteBatchRequest
	req.requestID = moovhttp.GetRequestID(r)

	vars := mux.Vars(r)
	fileID, ok := vars["fileID"]
	if !ok {
		return nil, ErrBadRouting
	}
	batchID, ok := vars["batchID"]
	if !ok {
		return nil, ErrBadRouting
	}

	req.fileID = fileID
	req.batchID = batchID
	return req, nil
}

func deleteBatchEndpoint(s Service, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(deleteBatchRequest)
		if !ok {
			err := errors.New("invalid request")
			return deleteBatchResponse{
				Err: err,
			}, err
		}

		err := s.DeleteBatch(req.fileID, req.batchID)

		if logger != nil {
			logger := logger.With(log.Fields{
				"batches":   log.String("deleteBatch"),
				"file":      log.String(req.fileID),
				"requestID": log.String(req.requestID),
			})
			if err != nil {
				logger.Error().LogError(err)
			} else {
				logger.Info().Log("delete batch")
			}
		}

		return deleteBatchResponse{
			Err: err,
		}, nil
	}
}
