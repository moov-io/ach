// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
	"unicode/utf8"

	"github.com/gorilla/mux"
)

const (
	maxHeaderLength = 36
)

// Problem writes err to w while also setting the HTTP status code, content-type and marshaling
// err as the response body.
func Problem(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

// Problem writes err to w while also setting the HTTP status code, content-type and marshaling
// err as the response body.
//
// Returned is the calling file and line number: server.go:33
func InternalError(w http.ResponseWriter, err error) string {
	w.WriteHeader(http.StatusInternalServerError)

	pcs := make([]uintptr, 5) // some limit
	_ = runtime.Callers(1, pcs)

	file, line := "", 0

	// Sometimes InternalError will be wrapped by helper methods inside an application.
	// We should linear search our callers until we find one outside github.com/moov-io
	// because that likely represents the stdlib.
	//
	// Note: This might not work for code already outside github.com/moov-io, please report
	// feedback if this works or not.
	i, frames := 0, runtime.CallersFrames(pcs)
	for {
		f, more := frames.Next()
		if !more {
			break
		}

		// f.Function can either be an absolute path (/Users/...) or a package
		// (i.e. github.com/moov-io/...) so check for either.
		if strings.Contains(f.Function, "github.com/moov-io") || strings.HasPrefix(f.Function, "main.") {
			_, file, line, _ = runtime.Caller(i) // next caller
		}
		i++
	}

	// Get the filename, file was a full path
	_, file = filepath.Split(file)
	return fmt.Sprintf("%s:%d", file, line)
}

// AddCORSHandler captures Corss Origin Resource Sharing (CORS) requests
// by looking at all OPTIONS requests for the Origin header, parsing that
// and responding back with the other Access-Control-Allow-* headers.
//
// Docs: https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
func AddCORSHandler(r *mux.Router) {
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		SetAccessControlAllowHeaders(w, r.Header.Get("Origin"))
		w.WriteHeader(http.StatusOK)
	})
}

// SetAccessControlAllowHeaders writes Access-Control-Allow-* headers to a response to allow
// for further CORS-allowed requests.
func SetAccessControlAllowHeaders(w http.ResponseWriter, origin string) {
	// Access-Control-Allow-Origin can't be '*' with requests that send credentials.
	// Instead, we need to explicitly set the domain (from request's Origin header)
	//
	// Allow requests from anyone's localhost and only from secure pages.
	if strings.HasPrefix(origin, "http://localhost:") || strings.HasPrefix(origin, "https://") {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PATCH,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Cookie,X-User-Id,X-Request-Id,Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	}
}

// GetRequestId returns the Moov header value for request IDs
func GetRequestId(r *http.Request) string {
	return r.Header.Get("X-Request-Id")
}

// GetUserId returns the Moov userId from HTTP headers
func GetUserId(r *http.Request) string {
	return r.Header.Get("X-User-Id")
}

func truncate(s string) string {
	if utf8.RuneCountInString(s) > maxHeaderLength {
		return s[:maxHeaderLength]
	}
	return s
}
