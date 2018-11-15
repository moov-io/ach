// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package server

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/moov-io/ach"

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
)

type createFileRequest struct {
	File *ach.File
}

type createFileResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error"`
}

func (r createFileResponse) error() error { return r.Err }

func createFileEndpoint(s Service, r Repository) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(createFileRequest)

		// record a metric for files created
		if req.File != nil && req.File.Header.ImmediateDestination != "" && req.File.Header.ImmediateOrigin != "" {
			filesCreated.With("destination", req.File.Header.ImmediateDestination, "origin", req.File.Header.ImmediateOrigin).Add(1)
		}

		if req.File.ID == "" {
			// No File ID, so create the file
			id, e := s.CreateFile(&req.File.Header)
			return createFileResponse{
				ID:  id,
				Err: e,
			}, nil
		} else {
			return createFileResponse{
				ID:  req.File.ID,
				Err: r.StoreFile(req.File),
			}, nil
		}
	}
}

func decodeCreateFileRequest(_ context.Context, request *http.Request) (interface{}, error) {
	var r io.Reader
	var req createFileRequest

	// Sets default values
	req.File = &ach.File{
		Header: ach.NewFileHeader(),
	}

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

type getFilesRequest struct{}

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
	return getFilesRequest{}, nil
}

type getFileRequest struct {
	ID string
}

type getFileResponse struct {
	File *ach.File `json:"file"`
	Err  error     `json:"error"`
}

func (r getFileResponse) error() error { return r.Err }

func getFileEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getFileRequest)
		f, e := s.GetFile(req.ID)
		return getFileResponse{
			File: f,
			Err:  e,
		}, nil
	}
}

func decodeGetFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return getFileRequest{ID: id}, nil
}

type deleteFileRequest struct {
	ID string
}

type deleteFileResponse struct {
	Err error `json:"err"`
}

func (r deleteFileResponse) error() error { return r.Err }

func deleteFileEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteFileRequest)
		return deleteFileResponse{
			Err: s.DeleteFile(req.ID),
		}, nil
	}
}

func decodeDeleteFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return deleteFileRequest{ID: id}, nil
}

type getFileContentsRequest struct {
	ID string
}

type getFileContentsResponse struct {
	Err error `json:"error"`
}

func (v getFileContentsResponse) error() error { return v.Err }

func getFileContentsEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getFileContentsRequest)
		r, err := s.GetFileContents(req.ID)
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
	return getFileContentsRequest{ID: id}, nil
}

type validateFileRequest struct {
	ID string
}

type validateFileResponse struct {
	Err error `json:"error"`
}

func (v validateFileResponse) error() error { return v.Err }

func validateFileEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(validateFileRequest)
		return validateFileResponse{
			Err: s.ValidateFile(req.ID),
		}, nil
	}
}

func decodeValidateFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return validateFileRequest{ID: id}, nil
}
