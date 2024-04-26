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
	"net/http"
	"strings"

	"github.com/moov-io/ach"
	"github.com/moov-io/base"
	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/base/log"

	"github.com/go-kit/kit/endpoint"
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
	File         *ach.File
	parseError   error
	requestID    string
	validateOpts *ach.ValidateOpts
}

type createFileResponse struct {
	ID   string    `json:"id"`
	File *ach.File `json:"file"`

	Err error `json:"error"`
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

		if req.validateOpts != nil {
			req.File.SetValidation(req.validateOpts)
		}

		err := r.StoreFile(req.File)
		if logger != nil {
			logger := logger.With(log.Fields{
				"files":     log.String("createFile"),
				"requestID": log.String(req.requestID),
			})
			if err != nil {
				logger.Error().LogError(err)
			} else {
				logger.Info().Log("create file")
			}
		}

		resp := createFileResponse{
			ID:   req.File.ID,
			File: req.File,
			Err:  err,
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

	body, validateOpts, err := readValidateOpts(request)
	if err != nil {
		return nil, err
	}
	req.validateOpts = validateOpts

	bs, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	h := strings.ToLower(request.Header.Get("Content-Type"))
	if strings.Contains(h, "application/json") {
		// Read body as ACH file in JSON
		f, err := ach.FileFromJSONWith(bs, req.validateOpts)
		if f != nil {
			req.File = f
		}
		req.parseError = err
	} else {
		// Attempt parsing body as an ACH File
		r = bytes.NewReader(bs)
		achReader := ach.NewReader(r)
		achReader.SetValidation(req.validateOpts)

		f, err := achReader.Read()
		req.File = &f
		req.parseError = err
	}

	// Set the fileID from the request
	fileID, ok := mux.Vars(request)["fileID"]
	if ok && fileID != "" && fileID != "create" {
		if req.File != nil {
			req.File.ID = fileID
		}
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
			logger := logger.With(log.Fields{
				"files":     log.String("getFile"),
				"requestID": log.String(req.requestID),
			})
			if err != nil {
				logger.Error().LogError(err)
			} else {
				logger.Info().Log("get file")
			}
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
			logger := logger.With(log.Fields{
				"files":     log.String("deleteFile"),
				"requestID": log.String(req.requestID),
			})
			if err != nil {
				logger.Error().LogError(err)
			} else {
				logger.Info().Log("delete file")
			}
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

type buildFileRequest struct {
	ID string

	requestID string
}

type buildFileResponse struct {
	File *ach.File `json:"file"`
	Err  error     `json:"error"`
}

func (v buildFileResponse) error() error { return v.Err }

func buildFileEndpoint(s Service, r Repository, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(buildFileRequest)
		if !ok {
			return buildFileResponse{Err: ErrFoundABug}, ErrFoundABug
		}

		file, err := s.BuildFile(req.ID)

		logger := logger.With(log.Fields{
			"files":     log.String("buildFile"),
			"fileID":    log.String(req.ID),
			"requestID": log.String(req.requestID),
		})
		if err != nil {
			logger.Error().LogError(err)
		} else {
			logger.Info().Log("build file")
		}

		return buildFileResponse{
			File: file,
			Err:  err,
		}, nil
	}
}

func decodeBuildFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return buildFileRequest{
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
			logger := logger.With(log.Fields{
				"files":     log.String("getFileContents"),
				"requestID": log.String(req.requestID),
			})
			if err != nil {
				logger.Error().LogError(err)
			} else {
				logger.Info().Log("get file contents")
			}
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
			logger := logger.With(log.Fields{
				"files":     log.String("validateFile"),
				"requestID": log.String(req.requestID),
			})
			if err != nil {
				logger.Error().LogError(err)
			} else {
				logger.Info().Log("validate file")
			}
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

	_, validateOpts, err := readValidateOpts(r)
	if err != nil {
		return nil, err
	}
	req.opts = validateOpts

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
			logger := logger.With(log.Fields{
				"files":     log.String(fmt.Sprintf("balance file created %s", balancedFile.ID)),
				"requestID": log.String(req.requestID),
			})
			if err != nil {
				logger.Error().LogError(err)
			} else {
				logger.Info().Log("balance file")
			}
		}
		if err != nil {
			if logger != nil {
				logger := logger.With(log.Fields{
					"files":     log.String(fmt.Sprintf("problem balancing %s: %v", req.fileID, err)),
					"requestID": log.String(req.requestID),
				})
				if err != nil {
					logger.Error().LogError(err)
				} else {
					logger.Info().Log("balance file")
				}

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

type segmentFileIDRequest struct {
	fileID    string
	requestID string

	opts *ach.SegmentFileConfiguration
}

type segmentedFilesResponse struct {
	CreditFileID string    `json:"creditFileID"`
	CreditFile   *ach.File `json:"creditFile"`

	DebitFileID string    `json:"debitFileID"`
	DebitFile   *ach.File `json:"debitFile"`

	Err error `json:"error"`
}

func segmentFileIDEndpoint(s Service, r Repository, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(segmentFileIDRequest)
		if !ok {
			return segmentedFilesResponse{Err: ErrFoundABug}, ErrFoundABug
		}

		creditFile, debitFile, err := s.SegmentFileID(req.fileID, req.opts)

		if logger != nil {
			logger.With(log.Fields{
				"files":     log.String("segmentFileID"),
				"requestID": log.String(req.requestID),
			})
			if err != nil {
				logger.Error().LogError(err)
			} else {
				logger.Info().Log("segment fileID")
			}
		}
		if err != nil {
			return segmentedFilesResponse{Err: err}, err
		}

		var resp segmentedFilesResponse

		if creditFile.ID != "" {
			err = r.StoreFile(creditFile)
			if logger != nil && err != nil {
				logger.With(log.Fields{
					"files":     log.String("storeCreditFile"),
					"requestID": log.String(req.requestID),
				}).LogError(err)
			}
			resp.CreditFile = creditFile
			resp.CreditFileID = creditFile.ID
		}

		if debitFile.ID != "" {
			err = r.StoreFile(debitFile)
			if logger != nil && err != nil {
				logger.With(log.Fields{
					"files":     log.String("storeDebitFile"),
					"requestID": log.String(req.requestID),
				}).LogError(err)
			}
			resp.DebitFile = debitFile
			resp.DebitFileID = debitFile.ID
		}

		resp.Err = err

		return resp, nil
	}
}

func decodeSegmentFileIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	fileID, ok := vars["fileID"]
	if !ok {
		return nil, ErrBadRouting
	}

	req := segmentFileIDRequest{
		fileID:    fileID,
		requestID: moovhttp.GetRequestID(r),
	}

	var opts ach.SegmentFileConfiguration
	if err := json.NewDecoder(r.Body).Decode(&opts); err == nil {
		req.opts = &opts
	}

	return req, nil
}

type segmentFileRequest struct {
	File      *ach.File
	requestID string

	opts         *ach.SegmentFileConfiguration
	validateOpts *ach.ValidateOpts
}

func segmentFileEndpoint(s Service, r Repository, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(segmentFileRequest)
		if !ok {
			return segmentedFilesResponse{Err: ErrFoundABug}, ErrFoundABug
		}

		if req.File != nil && req.validateOpts != nil {
			req.File.SetValidation(req.validateOpts)
		}

		creditFile, debitFile, err := s.SegmentFile(req.File, req.opts)
		if logger != nil {
			logger.With(log.Fields{
				"files":     log.String("segmentFile"),
				"requestID": log.String(req.requestID),
			})
			if err != nil {
				logger.Error().LogError(err)
			} else {
				logger.Info().Log("segment file")
			}
		}
		if err != nil {
			return segmentedFilesResponse{Err: err}, err
		}

		var resp segmentedFilesResponse

		if creditFile.ID != "" {
			err = r.StoreFile(creditFile)
			if logger != nil && err != nil {
				logger.With(log.Fields{
					"files":     log.String("storeCreditFile"),
					"requestID": log.String(req.requestID),
				}).LogError(err)
			}
			resp.CreditFile = creditFile
			resp.CreditFileID = creditFile.ID
		}

		if debitFile.ID != "" {
			err = r.StoreFile(debitFile)
			if logger != nil && err != nil {
				logger.With(log.Fields{
					"files":     log.String("storeDebitFile"),
					"requestID": log.String(req.requestID),
				}).LogError(err)
			}
			resp.DebitFile = debitFile
			resp.DebitFileID = debitFile.ID
		}

		resp.Err = err

		return resp, nil
	}
}

func decodeSegmentFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var file *ach.File

	var wrapper struct {
		Opts         *ach.SegmentFileConfiguration `json:"opts"`
		ValidateOpts *ach.ValidateOpts             `json:"validateOpts"`
	}

	header := strings.ToLower(r.Header.Get("content-type"))
	if strings.Contains(header, "application/json") {
		kv := make(map[string]json.RawMessage)

		err := json.NewDecoder(r.Body).Decode(&kv)
		if err != nil {
			return segmentedFilesResponse{Err: err}, fmt.Errorf("A : %v", err)
		}

		if vv, exists := kv["opts"]; vv != nil && exists {
			err = json.Unmarshal(vv, &wrapper.Opts)
			if err != nil {
				return segmentedFilesResponse{Err: err}, fmt.Errorf("B1 : %v", err)
			}
		}
		if vv, exists := kv["validateOpts"]; vv != nil && exists {
			err = json.Unmarshal(vv, &wrapper.ValidateOpts)
			if err != nil {
				return segmentedFilesResponse{Err: err}, fmt.Errorf("B2 : %v", err)
			}
		}
		if vv, exists := kv["file"]; exists {
			file, err = ach.FileFromJSONWith(vv, wrapper.ValidateOpts)
			if err != nil {
				return segmentedFilesResponse{Err: err}, fmt.Errorf("C : %v", err)
			}
		}
	} else {
		ff, err := ach.NewReader(r.Body).Read()
		if err != nil {
			return segmentedFilesResponse{Err: err}, fmt.Errorf("D : %v", err)
		}
		file = &ff
	}

	return segmentFileRequest{
		File:         file,
		requestID:    moovhttp.GetRequestID(r),
		opts:         wrapper.Opts,
		validateOpts: wrapper.ValidateOpts,
	}, nil
}

type flattenBatchesRequest struct {
	fileID    string
	requestID string
}

type flattenBatchesResponse struct {
	ID   string    `json:"id"`
	File *ach.File `json:"file"`
	Err  error     `json:"error"`
}

func flattenBatchesEndpoint(s Service, r Repository, logger log.Logger) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(flattenBatchesRequest)
		if !ok {
			return flattenBatchesResponse{Err: ErrFoundABug}, ErrFoundABug
		}
		flattenFile, err := s.FlattenBatches(req.fileID)
		if logger != nil {
			logger := logger.With(log.Fields{
				"files":     log.String("FlattenBatches"),
				"requestID": log.String(req.requestID),
			})
			if err != nil {
				logger.Error().LogError(err)
			} else {
				logger.Info().Log("flatten batches")
			}
		}
		if err != nil {
			return flattenBatchesResponse{Err: err}, err
		}
		if flattenFile.ID != "" {
			err = r.StoreFile(flattenFile)
			if logger != nil {
				logger := logger.With(log.Fields{
					"files":     log.String("storeFlattenFile"),
					"requestID": log.String(req.requestID),
				})
				if err != nil {
					logger.Error().LogError(err)
				} else {
					logger.Info().Log("flatten batches")
				}
			}
		}
		return flattenBatchesResponse{
			ID:   flattenFile.ID,
			File: flattenFile,
			Err:  err,
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
