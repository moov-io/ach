// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/moov-io/ach"
	moovhttp "github.com/moov-io/base/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

type createBatchRequest struct {
	FileID string
	Batch  *ach.Batch

	requestId string
}

type createBatchResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error"`
}

func (r createBatchResponse) error() error { return r.Err }

func createBatchEndpoint(s Service, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(createBatchRequest)
		id, err := s.CreateBatch(req.FileID, req.Batch)

		if req.requestId != "" && logger != nil {
			logger.Log("batches", "createBatch", "file", req.FileID, "requestId", req.requestId, "error", err)
		}

		return createBatchResponse{
			ID:  id,
			Err: err,
		}, nil
	}
}

func decodeCreateBatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createBatchRequest
	req.requestId = moovhttp.GetRequestId(r)

	vars := mux.Vars(r)
	id, ok := vars["fileID"]
	if !ok {
		return nil, ErrBadRouting
	}
	req.FileID = id
	if err := json.NewDecoder(r.Body).Decode(req.Batch); err != nil {
		return nil, err
	}
	if req.Batch == nil {
		return nil, errors.New("no Batch provided")
	}
	return req, nil
}

type getBatchesRequest struct {
	fileID string

	requestId string
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
	req.requestId = moovhttp.GetRequestId(r)

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
		req := request.(getBatchesRequest)
		if req.requestId != "" && logger != nil {
			logger.Log("batches", "getBatches", "file", req.fileID, "requestId", req.requestId)
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

	requestId string
}

type getBatchResponse struct {
	Batch ach.Batcher `json:"batch"`
	Err   error       `json:"error"`
}

func (r getBatchResponse) error() error { return r.Err }

func decodeGetBatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req getBatchRequest
	req.requestId = moovhttp.GetRequestId(r)

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
		req := request.(getBatchRequest)
		batch, err := s.GetBatch(req.fileID, req.batchID)

		if req.requestId != "" && logger != nil {
			logger.Log("batches", "getBatche", "file", req.fileID, "requestId", req.requestId, "error", err)
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

	requestId string
}

type deleteBatchResponse struct {
	Err error `json:"error"`
}

func (r deleteBatchResponse) error() error { return r.Err }

func decodeDeleteBatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req deleteBatchRequest
	req.requestId = moovhttp.GetRequestId(r)

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
		req := request.(deleteBatchRequest)
		err := s.DeleteBatch(req.fileID, req.batchID)

		if req.requestId != "" && logger != nil {
			logger.Log("batches", "deleteBatch", "file", req.fileID, "requestId", req.requestId, "error", err)
		}

		return deleteBatchResponse{
			Err: err,
		}, nil
	}
}
