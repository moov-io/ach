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
	"strconv"
	"strings"

	"github.com/moov-io/ach"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	bugReportHelp = "please report this as a bug -- https://github.com/moov-io/ach/issues/new"

	// ErrBadRouting is returned when an expected path variable is missing, which is always programmer error.
	ErrBadRouting = fmt.Errorf("inconsistent mapping between route and handler, %s", bugReportHelp)
	ErrFoundABug  = fmt.Errorf("Snuck into encodeError with err == nil, %s", bugReportHelp)

	MaxContentLength = 1 * 1024 * 1024 // bytes
)

func MakeHTTPHandler(s Service, repo Repository, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(s, repo)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	// HTTP Methods
	r.Methods("GET").Path("/ping").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("PONG"))
	})
	r.Methods("GET").Path("/files").Handler(httptransport.NewServer(
		e.GetFilesEndpoint,
		decodeGetFilesRequest,
		encodeResponse,
		options...,
	))
	r.Methods("POST").Path("/files/create").Handler(httptransport.NewServer(
		e.CreateFileEndpoint,
		decodeCreateFileRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/files/{id}").Handler(httptransport.NewServer(
		e.GetFileEndpoint,
		decodeGetFileRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/files/{id}/contents").Handler(httptransport.NewServer(
		e.GetFileContentsEndpoint,
		decodeGetFileContentsRequest,
		encodeTextResponse,
		options...,
	))
	r.Methods("GET").Path("/files/{id}/validate").Handler(httptransport.NewServer(
		e.ValidateFileEndpoint,
		decodeValidateFileRequest,
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
func decodeCreateFileRequest(_ context.Context, request *http.Request) (interface{}, error) {
	// Make sure content-length is small enough
	if !acceptableContentLength(request.Header) {
		return nil, errors.New("request body is too large")
	}

	var r io.Reader
	var req createFileRequest

	// Sets default values
	req.File = ach.File{
		Header: ach.NewFileHeader(),
	}

	bs, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	h := request.Header.Get("Content-Type")
	if strings.Contains(h, "application/json") {
		// Attempt to read file as json
		r = bytes.NewReader(bs)
		if err := json.NewDecoder(r).Decode(&req.File.Header); err == nil {
			return req, nil
		}
	} else {
		// Attempt parsing body as an ACH File
		r = bytes.NewReader(bs)
		f, err := ach.NewReader(r).Read()
		if err != nil {
			return nil, err
		}
		req.File = f
	}
	return req, nil
}

func decodeGetFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return getFileRequest{ID: id}, nil
}

func decodeDeleteFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return deleteFileRequest{ID: id}, nil
}

func decodeGetFilesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return getFilesRequest{}, nil
}

func decodeGetFileContentsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return getFileContentsRequest{ID: id}, nil
}

func decodeValidateFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return validateFileRequest{ID: id}, nil
}

//** BATCHES **//

func decodeCreateBatchRequest(_ context.Context, r *http.Request) (interface{}, error) {
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

// errorer is implemented by all concrete response types that may contain
// errors. There are a few well-known values which are used to change the
// HTTP response code without needing to trigger an endpoint (transport-level)
// error.
type errorer interface {
	error() error
}

// counter is implemented by any concrete response types that may contain
// some arbitrary count information.
type counter interface {
	count() int
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
	if e, ok := response.(counter); ok {
		w.Header().Set("X-Total-Count", strconv.Itoa(e.count()))
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeTextResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if r, ok := response.(io.Reader); ok {
		w.Header().Set("Content-Type", "text/plain")
		_, err := io.Copy(w, r)
		return err
	}
	return encodeResponse(ctx, w, response)
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
	// TODO(adam): this should really probably be a 4xx error
	// TODO(adam): on GET /files/:id/validate a "bad" file returns 500
	return http.StatusInternalServerError
}

func acceptableContentLength(headers http.Header) bool {
	h := headers.Get("Content-Length")
	if v, err := strconv.Atoi(h); err == nil {
		return v <= MaxContentLength
	}
	return false
}
