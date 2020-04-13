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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/moov-io/ach"
	"github.com/moov-io/base"
	moovhttp "github.com/moov-io/base/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/gorilla/mux"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

var (
	filesCreated = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Name: "ach_files_created",
		Help: "The number of ACH files created",
	}, []string{"destination", "origin"})

	filesDeleted = prometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Name: "ach_files_deleted",
		Help: "The number of ACH files deleted",
	}, nil)
)

type createFileRequest struct {
	File       *ach.File
	parseError error
	requestID  string
}

type createFileResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error"`
}

func (r createFileResponse) error() error { return r.Err }

func createFileEndpoint(s Service, r Repository, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(createFileRequest)
		if !ok {
			return createFileResponse{Err: ErrFoundABug}, ErrFoundABug
		}

		// record a metric for files created
		if req.File != nil && req.File.Header.ImmediateDestination != "" && req.File.Header.ImmediateOrigin != "" {
			filesCreated.With("destination", req.File.Header.ImmediateDestination, "origin", req.File.Header.ImmediateOrigin).Add(1)
		}

		// Create a random file ID if none was provided
		if req.File.ID == "" {
			req.File.ID = base.ID()
		}

		err := r.StoreFile(req.File)
		if logger != nil {
			logger.Log("files", "createFile", "requestID", req.requestID, "error", err)
		}

		resp := createFileResponse{
			ID:  req.File.ID,
			Err: err,
		}
		if req.parseError != nil {
			resp.Err = req.parseError
		}

		return resp, nil
	}
}

func decodeCreateFileRequest(_ context.Context, request *http.Request) (interface{}, error) {
	var r io.Reader
	req := createFileRequest{
		File:      ach.NewFile(),
		requestID: moovhttp.GetRequestID(request),
	}

	bs, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	h := strings.ToLower(request.Header.Get("Content-Type"))
	if strings.Contains(h, "application/json") {
		// Read body as ACH file in JSON
		f, err := ach.FileFromJSON(bs)
		if f != nil {
			req.File = f
		}
		req.parseError = err
	} else {
		// Attempt parsing body as an ACH File
		r = bytes.NewReader(bs)
		f, err := ach.NewReader(r).Read()
		req.File = &f
		req.parseError = err
	}
	return req, nil
}

type getFilesRequest struct {
	requestID string
}

type getFilesResponse struct {
	Files []*ach.File `json:"files"`
	Err   error       `json:"error"`
}

func (r getFilesResponse) count() int { return len(r.Files) }

func (r getFilesResponse) error() error { return r.Err }

func getFilesEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (interface{}, error) {
		return getFilesResponse{
			Files: s.GetFiles(),
			Err:   nil,
		}, nil
	}
}

func decodeGetFilesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return getFilesRequest{
		requestID: moovhttp.GetRequestID(r),
	}, nil
}

type getFileRequest struct {
	ID string

	requestID string
}

type getFileResponse struct {
	File *ach.File `json:"file"`
	Err  error     `json:"error"`
}

func (r getFileResponse) error() error { return r.Err }

func getFileEndpoint(s Service, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(getFileRequest)
		if !ok {
			return getFileResponse{Err: ErrFoundABug}, ErrFoundABug
		}

		f, err := s.GetFile(req.ID)

		if logger != nil {
			logger.Log("files", "getFile", "requestID", req.requestID, "error", err)
		}

		return getFileResponse{
			File: f,
			Err:  err,
		}, nil
	}
}

func decodeGetFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return getFileRequest{
		ID:        id,
		requestID: moovhttp.GetRequestID(r),
	}, nil
}

type deleteFileRequest struct {
	ID string

	requestID string
}

type deleteFileResponse struct {
	Err error `json:"err"`
}

func (r deleteFileResponse) error() error { return r.Err }

func deleteFileEndpoint(s Service, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(deleteFileRequest)
		if !ok {
			return deleteFileResponse{Err: ErrFoundABug}, ErrFoundABug
		}

		filesDeleted.Add(1)

		err := s.DeleteFile(req.ID)

		if logger != nil {
			logger.Log("files", "deleteFile", "requestID", req.requestID, "error", err)
		}

		return deleteFileResponse{
			Err: err,
		}, nil
	}
}

func decodeDeleteFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return deleteFileRequest{
		ID:        id,
		requestID: moovhttp.GetRequestID(r),
	}, nil
}

type getFileContentsRequest struct {
	ID string

	requestID string
}

type getFileContentsResponse struct {
	Err error `json:"error"`
}

func (v getFileContentsResponse) error() error { return v.Err }

func getFileContentsEndpoint(s Service, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(getFileContentsRequest)
		if !ok {
			return getFileContentsResponse{Err: ErrFoundABug}, ErrFoundABug
		}

		r, err := s.GetFileContents(req.ID)

		if logger != nil {
			logger.Log("files", "getFileContents", "requestID", req.requestID, "error", err)
		}
		if err != nil {
			return getFileContentsResponse{Err: err}, nil
		}

		return r, nil
	}
}

func decodeGetFileContentsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return getFileContentsRequest{
		ID:        id,
		requestID: moovhttp.GetRequestID(r),
	}, nil
}

type validateFileRequest struct {
	ID        string
	requestID string

	opts *ach.ValidateOpts
}

type validateFileResponse struct {
	Err error `json:"error"`
}

func (v validateFileResponse) error() error { return v.Err }

func validateFileEndpoint(s Service, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(validateFileRequest)
		if !ok {
			return validateFileResponse{Err: ErrFoundABug}, ErrFoundABug
		}

		err := s.ValidateFile(req.ID, req.opts)
		if logger != nil {
			logger.Log("files", "validateFile", "requestID", req.requestID, "error", err)
		}
		if err != nil { // wrap err with context
			err = fmt.Errorf("%v: %v", errInvalidFile, err)
		}
		return validateFileResponse{err}, nil
	}
}

func decodeValidateFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}

	req := validateFileRequest{
		ID:        id,
		requestID: moovhttp.GetRequestID(r),
	}

	var opts ach.ValidateOpts
	if err := json.NewDecoder(r.Body).Decode(&opts); err == nil {
		req.opts = &opts
	}

	return req, nil
}

type balanceFileRequest struct {
	fileID    string
	offset    *ach.Offset
	requestID string
}

type balanceFileResponse struct {
	FileID string `json:"id"`
	Err    error  `json:"error"`
}

func balanceFileEndpoint(s Service, r Repository, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(balanceFileRequest)
		if !ok {
			return balanceFileResponse{Err: ErrFoundABug}, ErrFoundABug
		}

		balancedFile, err := s.BalanceFile(req.fileID, req.offset)
		if balancedFile != nil && logger != nil {
			logger.Log("files", fmt.Sprintf("balance file created %s", balancedFile.ID), "requestID", req.requestID, "error", err)
		}
		if err != nil {
			if logger != nil {
				logger.Log("files", fmt.Sprintf("problem balancing %s: %v", req.fileID, err), "requestID", req.requestID, "error", err)
			}
			return balanceFileResponse{Err: err}, err
		}
		return balanceFileResponse{
			FileID: balancedFile.ID,
		}, nil
	}
}

func decodeBalanceFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	fileID, ok := vars["fileID"]
	if !ok {
		return nil, ErrBadRouting
	}

	var off ach.Offset
	if err := json.NewDecoder(r.Body).Decode(&off); err != nil {
		return nil, err
	}
	if off.RoutingNumber == "" || off.AccountNumber == "" || string(off.AccountType) == "" {
		return nil, errors.New("missing some offset json fields")
	}
	return balanceFileRequest{
		fileID:    fileID,
		offset:    &off,
		requestID: moovhttp.GetRequestID(r),
	}, nil
}

type segmentFileRequest struct {
	fileID    string
	requestID string

	opts *ach.SegmentFileConfiguration
}

type segmentFileResponse struct {
	CreditFileID string `json:"creditFileID"`
	DebitFileID  string `json:"debitFileID"`
	Err          error  `json:"error"`
}

func segmentFileEndpoint(s Service, r Repository, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(segmentFileRequest)
		if !ok {
			return segmentFileResponse{Err: ErrFoundABug}, ErrFoundABug
		}

		creditFile, debitFile, err := s.SegmentFile(req.fileID, req.opts)

		if logger != nil {
			logger.Log("files", "segmentFile", "requestID", req.requestID, "error", err)
		}
		if err != nil {
			return segmentFileResponse{Err: err}, err
		}

		if creditFile.ID != "" {
			err = r.StoreFile(creditFile)
			if logger != nil {
				logger.Log("files", "storeCreditFile", "requestID", req.requestID, "error", err)
			}
		}

		if debitFile.ID != "" {
			err = r.StoreFile(debitFile)
			if logger != nil {
				logger.Log("files", "storeDebitFile", "requestID", req.requestID, "error", err)
			}
		}
		return segmentFileResponse{
			CreditFileID: creditFile.ID,
			DebitFileID:  debitFile.ID,
			Err:          err,
		}, nil
	}
}

func decodeSegmentFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	fileID, ok := vars["fileID"]
	if !ok {
		return nil, ErrBadRouting
	}

	req := segmentFileRequest{
		fileID:    fileID,
		requestID: moovhttp.GetRequestID(r),
	}

	var opts ach.SegmentFileConfiguration
	if err := json.NewDecoder(r.Body).Decode(&opts); err == nil {
		req.opts = &opts
	}

	return req, nil
}

type flattenBatchesRequest struct {
	fileID    string
	requestID string
}

type flattenBatchesResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error"`
}

func flattenBatchesEndpoint(s Service, r Repository, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(flattenBatchesRequest)
		if !ok {
			return flattenBatchesResponse{Err: ErrFoundABug}, ErrFoundABug
		}
		flattenFile, err := s.FlattenBatches(req.fileID)
		if logger != nil {
			logger.Log("files", "FlattenBatches", "requestID", req.requestID, "error", err)
		}
		if err != nil {
			return flattenBatchesResponse{Err: err}, err
		}
		if flattenFile.ID != "" {
			err = r.StoreFile(flattenFile)
			if logger != nil {
				logger.Log("files", "storeFlattenFile", "requestID", req.requestID, "error", err)
			}
		}
		return flattenBatchesResponse{
			ID:  flattenFile.ID,
			Err: err,
		}, nil
	}
}

func decodeFlattenBatchesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	fileID, ok := vars["fileID"]
	if !ok {
		return nil, ErrBadRouting
	}
	return flattenBatchesRequest{
		fileID:    fileID,
		requestID: moovhttp.GetRequestID(r),
	}, nil
}
