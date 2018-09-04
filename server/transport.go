package server

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"

	"net/http"
	"net/url"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/moov-io/ach"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")

	ErrFoundABug = errors.New("Snuck into encodeError with err == nil, please report this as a bug -- https://github.com/moov-io/ach/issues/new")
)

func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	// TODO(Adam): endpoints still needed

	// GET    /files/:id/validate	validates the supplied file id for nacha compliance

	// MUTATES FILE
	// PATCH  /files/:id/build	build batch and file controls in ach file with supplied values

	// POST	  /files/upload	        Parse body as an ACH File and add to repository, replaces any existing file by ID

	// HTTP Methods
	r.Methods("POST").Path("/files/").Handler(httptransport.NewServer(
		e.CreateFileEndpoint,
		decodeCreateFileRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/files/").Handler(httptransport.NewServer(
		e.GetFilesEndpoint,
		decodeGetFilesRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/files/{id}").Handler(httptransport.NewServer(
		e.GetFileEndpoint,
		decodeGetFileRequest,
		encodeResponse,
		options...,
	))
	r.Methods("DELETE").Path("/files/{id}").Handler(httptransport.NewServer(
		e.DeleteFileEndpoint,
		decodeDeleteFileRequest,
		encodeResponse,
		options...,
	))
	r.Methods("POST").Path("/files/{fileID}/batches/").Handler(httptransport.NewServer(
		e.CreateBatchEndpoint,
		decodeCreateBatchRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/files/{fileID}/batches/").Handler(httptransport.NewServer(
		e.GetBatchesEndpoint,
		decodeGetBatchesRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/files/{fileID}/batches/{batchID}").Handler(httptransport.NewServer(
		e.GetBatchEndpoint,
		decodeGetBatchRequest,
		encodeResponse,
		options...,
	))
	r.Methods("DELETE").Path("/files/{fileID}/batches/{batchID}").Handler(httptransport.NewServer(
		e.DeleteBatchEndpoint,
		decodeDeleteBatchRequest,
		encodeResponse,
		options...,
	))
	return r
}

//** FILES ** //
func decodeCreateFileRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req createFileRequest
	// Sets default values
	req.FileHeader = ach.NewFileHeader()
	if e := json.NewDecoder(r.Body).Decode(&req.FileHeader); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetFileRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return getFileRequest{ID: id}, nil
}

func decodeDeleteFileRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return deleteFileRequest{ID: id}, nil
}

func encodeCreateFileRequest(ctx context.Context, req *http.Request, request interface{}) error {
	req.Method, req.URL.Path = "POST", "/files/"
	return encodeRequest(ctx, req, request)
}
func decodeGetFilesRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return getFilesRequest{}, nil
}

func encodeGetFileRequest(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(getFileRequest)
	fileID := url.QueryEscape(r.ID)
	req.Method, req.URL.Path = "GET", "/files/"+fileID
	return encodeRequest(ctx, req, request)
}

//** BATCHES **//

func decodeCreateBatchRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req createBatchRequest
	vars := mux.Vars(r)
	id, ok := vars["fileID"]
	if !ok {
		return nil, ErrBadRouting
	}
	req.FileID = id
	req.BatchHeader = *ach.NewBatchHeader()
	if e := json.NewDecoder(r.Body).Decode(&req.BatchHeader); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetBatchesRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req getBatchesRequest
	vars := mux.Vars(r)
	id, ok := vars["fileID"]
	if !ok {
		return nil, ErrBadRouting
	}
	req.fileID = id
	return req, nil
}

func decodeGetBatchRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
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

func decodeDeleteBatchRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
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

// errorer is implemented by all concrete response types that may contain
// errors. There are a few well-known values which are used to change the
// HTTP response code without needing to trigger an endpoint (transport-level)
// error.
type errorer interface {
	error() error
}

// encodeResponse is the common method to encode all response types to the
// client. I chose to do it this way because, since we're using JSON, there's no
// reason to provide anything more specific. It's certainly possible to
// specialize on a per-response (per-method) basis.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// encodeRequest likewise JSON-encodes the request to the HTTP request body.
// Don't use it directly as a transport/http.Client EncodeRequestFunc:
// Service endpoints require mutating the HTTP method and request path.
func encodeRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}

// encodeError JSON encodes the supplied error
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		err = ErrFoundABug
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrAlreadyExists:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
