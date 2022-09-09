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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/moov-io/base"
	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/base/log"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	gokitlog "github.com/go-kit/log"
	"github.com/gorilla/mux"
)

var (
	bugReportHelp = "please report this as a bug -- https://github.com/moov-io/ach/issues/new"

	// ErrBadRouting is returned when an expected path variable is missing, which is always programmer error.
	ErrBadRouting = fmt.Errorf("inconsistent mapping between route and handler, %s", bugReportHelp)
	ErrFoundABug  = fmt.Errorf("snuck into encodeError with err == nil, %s", bugReportHelp)

	errInvalidFile = errors.New("invalid ACH file")
)

// contextKey is a unique (and compariable) type we use
// to store and retrieve additional information in the
// go-kit context.
var contextKey struct{}

// saveCORSHeadersIntoContext saves CORS headers into the go-kit context.
//
// This is designed to be added as a ServerOption in our main http handler.
func saveCORSHeadersIntoContext() httptransport.RequestFunc {
	return func(ctx context.Context, r *http.Request) context.Context {
		origin := r.Header.Get("Origin")
		return context.WithValue(ctx, contextKey, origin)
	}
}

// respondWithSavedCORSHeaders looks in the go-kit request context
// for our own CORS headers. (Stored with our context key in
// saveCORSHeadersIntoContext.)
//
// This is designed to be added as a ServerOption in our main http handler.
func respondWithSavedCORSHeaders() httptransport.ServerResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter) context.Context {
		v, ok := ctx.Value(contextKey).(string)
		if ok && v != "" {
			moovhttp.SetAccessControlAllowHeaders(w, v) // set CORS headers
		}
		return ctx
	}
}

// preflightHandler captures Corss Origin Resource Sharing (CORS) requests
// by looking at all OPTIONS requests for the Origin header, parsing that
// and responding back with the other Access-Control-Allow-* headers.
//
// Docs: https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
func preflightHandler(options []httptransport.ServerOption) http.Handler {
	return httptransport.NewServer(
		endpoint.Nop,
		httptransport.NopRequestDecoder,
		func(_ context.Context, w http.ResponseWriter, _ interface{}) error {
			if v := w.Header().Get("Content-Type"); v == "" {
				w.Header().Set("Content-Type", "text/plain")
			}
			return nil
		},
		options...,
	)
}

func MakeHTTPHandler(s Service, repo Repository, kitlog gokitlog.Logger) http.Handler {
	logger := log.NewLogger(kitlog)

	r := mux.NewRouter()
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(kitlog),
		httptransport.ServerErrorEncoder(encodeError),
		httptransport.ServerBefore(saveCORSHeadersIntoContext()),
		httptransport.ServerAfter(respondWithSavedCORSHeaders()),
	}

	// HTTP Methods
	r.Methods("OPTIONS").Handler(preflightHandler(options)) // CORS pre-flight handler
	r.Methods("GET").Path("/ping").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		moovhttp.SetAccessControlAllowHeaders(w, r.Header.Get("Origin"))
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("PONG"))
	})
	r.Methods("GET").Path("/files").Handler(httptransport.NewServer(
		getFilesEndpoint(s),
		decodeGetFilesRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/files/{id}/build").Handler(httptransport.NewServer(
		buildFileEndpoint(s, repo, logger),
		decodeBuildFileRequest,
		encodeResponse,
		options...,
	))
	r.Methods("POST").Path("/files/{fileID}").Handler(httptransport.NewServer(
		createFileEndpoint(s, repo, logger),
		decodeCreateFileRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/files/{id}").Handler(httptransport.NewServer(
		getFileEndpoint(s, logger),
		decodeGetFileRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/files/{id}/contents").Handler(httptransport.NewServer(
		getFileContentsEndpoint(s, logger),
		decodeGetFileContentsRequest,
		encodeTextResponse,
		options...,
	))
	r.Methods("GET").Path("/files/{id}/validate").Handler(httptransport.NewServer(
		validateFileEndpoint(s, logger),
		decodeValidateFileRequest,
		encodeResponse,
		options...,
	))
	r.Methods("POST").Path("/files/{id}/validate").Handler(httptransport.NewServer(
		validateFileEndpoint(s, logger),
		decodeValidateFileRequest,
		encodeResponse,
		options...,
	))
	r.Methods("DELETE").Path("/files/{id}").Handler(httptransport.NewServer(
		deleteFileEndpoint(s, logger),
		decodeDeleteFileRequest,
		encodeResponse,
		options...,
	))
	r.Methods("POST").Path("/files/{fileID}/batches").Handler(httptransport.NewServer(
		createBatchEndpoint(s, logger),
		decodeCreateBatchRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/files/{fileID}/batches").Handler(httptransport.NewServer(
		getBatchesEndpoint(s, logger),
		decodeGetBatchesRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/files/{fileID}/batches/{batchID}").Handler(httptransport.NewServer(
		getBatchEndpoint(s, logger),
		decodeGetBatchRequest,
		encodeResponse,
		options...,
	))
	r.Methods("DELETE").Path("/files/{fileID}/batches/{batchID}").Handler(httptransport.NewServer(
		deleteBatchEndpoint(s, logger),
		decodeDeleteBatchRequest,
		encodeResponse,
		options...,
	))
	r.Methods("POST").Path("/files/{fileID}/balance").Handler(httptransport.NewServer(
		balanceFileEndpoint(s, repo, logger),
		decodeBalanceFileRequest,
		encodeResponse,
		options...,
	))
	r.Methods("POST").Path("/files/{fileID}/segment").Handler(httptransport.NewServer(
		segmentFileIDEndpoint(s, repo, logger),
		decodeSegmentFileIDRequest,
		encodeResponse,
		options...,
	))
	r.Methods("POST").Path("/segment").Handler(httptransport.NewServer(
		segmentFileEndpoint(s, repo, logger),
		decodeSegmentFileRequest,
		encodeResponse,
		options...,
	))
	r.Methods("POST").Path("/files/{fileID}/flatten").Handler(httptransport.NewServer(
		flattenBatchesEndpoint(s, repo, logger),
		decodeFlattenBatchesRequest,
		encodeResponse,
		options...,
	))
	return r
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

// marshalStructWithError converts a struct into a JSON response with all fields of the struct
// with our expected error formats.
//
// There are a few reasons we need to do this.
//  1. base.ErrorList marshals to an object which breaks the string format our API declares
//     and isn't caught when we pass around interface{} values.
//  2. We want to return additional fields of structs (such as in createFileEndpoint)
func marshalStructWithError(in interface{}, w http.ResponseWriter) error {
	v := reflect.ValueOf(in)
	out := make(map[string]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		name := v.Type().Field(i).Name
		value := v.Field(i).Interface()

		if err, ok := value.(error); ok {
			out["error"] = err.Error()
		} else {
			out[name] = value
		}
	}

	return json.NewEncoder(w).Encode(out)
}

// encodeResponse is the common method to encode all response types to the
// client. I chose to do it this way because, since we're using JSON, there's no
// reason to provide anything more specific. It's certainly possible to
// specialize on a per-response (per-method) basis.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(codeFrom(e.error()))
		return marshalStructWithError(response, w)
	}

	// Used for pagination
	if e, ok := response.(counter); ok {
		w.Header().Set("X-Total-Count", strconv.Itoa(e.count()))
	}

	// Don't overwrite a header (i.e. called from encodeTextResponse)
	if v := w.Header().Get("Content-Type"); v == "" {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// Only write json body if we're setting response as json
		return json.NewEncoder(w).Encode(response)
	}
	return nil
}

// encodeTextResponse will marshal response into the HTTP Response
// This method is designed text/plain content-types and expects response
// to be an io.Reader.
func encodeTextResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if r, ok := response.(io.Reader); ok {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		_, err := io.Copy(w, r)
		return err
	}
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

	errString := fmt.Sprintf("%#v", err)
	if el, ok := err.(base.ErrorList); ok {
		errString = el.Error()
	}
	switch {
	case
		strings.Contains(errString, errInvalidFile.Error()), // This branch comes from validateFileEndpoint
		strings.Contains(errString, "*ach.FieldError"),
		strings.Contains(errString, "*ach.BatchError"),
		strings.Contains(errString, "*ach.ErrFile"),
		strings.Contains(errString, "ach.RecordWrongLengthErr"),
		strings.Contains(errString, "FieldName"): // FileFromJSON
		return http.StatusBadRequest
	}

	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrAlreadyExists:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
