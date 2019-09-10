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
	File *ach.File

	requestID string
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
			err := errors.New("invalid request")
			return createFileResponse{
				Err: err,
			}, err
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

		return createFileResponse{
			ID:  req.File.ID,
			Err: err,
		}, nil
	}
}

func decodeCreateFileRequest(_ context.Context, request *http.Request) (interface{}, error) {
	var r io.Reader
	var req createFileRequest

	req.requestID = moovhttp.GetRequestID(request)

	// Sets default values
	req.File = ach.NewFile()
	bs, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	h := request.Header.Get("Content-Type")
	if strings.Contains(h, "application/json") {
		// Read body as ACH file in JSON
		f, err := ach.FileFromJSON(bs)
		if err != nil {
			return nil, err
		}
		req.File = f
	} else {
		// Attempt parsing body as an ACH File
		r = bytes.NewReader(bs)
		f, err := ach.NewReader(r).Read()
		if err != nil {
			return nil, err
		}
		req.File = &f
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
			err := errors.New("invalid request")
			return getFileResponse{
				Err: err,
			}, err
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
			err := errors.New("invalid request")
			return deleteFileResponse{
				Err: err,
			}, err
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
			err := errors.New("invalid request")
			return getFileContentsResponse{
				Err: err,
			}, err
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
	ID string

	requestID string
}

type validateFileResponse struct {
	Err error `json:"error"`
}

func (v validateFileResponse) error() error { return v.Err }

func validateFileEndpoint(s Service, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(validateFileRequest)
		if !ok {
			err := errors.New("invalid request")
			return validateFileResponse{
				Err: err,
			}, err
		}

		err := s.ValidateFile(req.ID)
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
	return validateFileRequest{
		ID:        id,
		requestID: moovhttp.GetRequestID(r),
	}, nil
}

type segmentFileRequest struct {
	fileID    string
	requestID string
}

type segmentFileResponse struct {
	creditFileID string `json:"creditFileID"`
	debitFileID  string `json:"debitFileID"`
	Err          error  `json:"error"`
}

func segmentFileEndpoint(s Service, r Repository, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(segmentFileRequest)
		if !ok {
			err := errors.New("invalid request")
			return segmentFileResponse{
				Err: err,
			}, err
		}

		creditFile, debitFile, err := s.SegmentFile(req.fileID)

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
			creditFileID: creditFile.ID,
			debitFileID:  debitFile.ID,
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
	return segmentFileRequest{
		fileID:    fileID,
		requestID: moovhttp.GetRequestID(r),
	}, nil
}

type flattenBatchesRequest struct {
	fileID    string
	requestID string
}

type flattenBatchesResponse struct {
	id  string `json:"id"`
	Err error  `json:"error"`
}

func flattenBatchesEndpoint(s Service, r Repository, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(flattenBatchesRequest)
		if !ok {
			err := errors.New("invalid request")
			return flattenBatchesResponse{
				Err: err,
			}, err
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
			id:  flattenFile.ID,
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
