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

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

type createBatchRequest struct {
	FileID string
	Batch  *ach.Batch
}

type createBatchResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error"`
}

func (r createBatchResponse) error() error { return r.Err }

func createBatchEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(createBatchRequest)
		id, e := s.CreateBatch(req.FileID, req.Batch)
		return createBatchResponse{
			ID:  id,
			Err: e,
		}, nil
	}
}

func decodeCreateBatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createBatchRequest
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
	vars := mux.Vars(r)
	id, ok := vars["fileID"]
	if !ok {
		return nil, ErrBadRouting
	}
	req.fileID = id
	return req, nil
}

func getBatchesEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getBatchesRequest)
		return getBatchesResponse{
			Batches: s.GetBatches(req.fileID),
			Err:     nil,
		}, nil
	}
}

type getBatchRequest struct {
	fileID  string
	batchID string
}

type getBatchResponse struct {
	Batch ach.Batcher `json:"batch"`
	Err   error       `json:"error"`
}

func (r getBatchResponse) error() error { return r.Err }

func decodeGetBatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req getBatchRequest
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

func getBatchEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getBatchRequest)
		batch, e := s.GetBatch(req.fileID, req.batchID)
		return getBatchResponse{
			Batch: batch,
			Err:   e,
		}, nil
	}
}

type deleteBatchRequest struct {
	fileID  string
	batchID string
}

type deleteBatchResponse struct {
	Err error `json:"error"`
}

func (r deleteBatchResponse) error() error { return r.Err }

func decodeDeleteBatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req deleteBatchRequest
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

func deleteBatchEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteBatchRequest)
		return deleteBatchResponse{
			Err: s.DeleteBatch(req.fileID, req.batchID),
		}, nil
	}
}
