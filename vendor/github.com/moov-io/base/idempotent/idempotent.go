// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package idempotent

import (
	"errors"
	"net/http"
	"unicode/utf8"
)

const (
	// maxIdempotencyKeyLength is the longest X-Idempotency-Key string legnth allowed.
	maxIdempotencyKeyLength = 50

	// HeaderKey is the HTTP header key used to stored idempotency keys
	HeaderKey = "X-Idempotency-Key"
)

var (
	// ErrSeenBefore is returned when our http ResponseWriter has seen the idempotency key before
	ErrSeenBefore = errors.New("X-Idempotency-Key seen before")
)

// Recorder offers a method to determine if a given key has been
// seen before or not. Each invocation of SeenBefore needs to
// record each key found, but there's no minimum duration required.
type Recorder interface {
	SeenBefore(key string) bool
}

// Header returns the idempotency key from a http.Request headers
func Header(r *http.Request) string {
	return truncate(r.Header.Get(HeaderKey))
}

// FromRequest extracts the idempotency key from HTTP headers and records its presence in
// the provided Recorder.
//
// A nil Recorder will always return idempotency keys as unseen.
func FromRequest(req *http.Request, rec Recorder) (key string, seen bool) {
	key = Header(req)
	if rec == nil {
		return key, false
	}
	if key == "" {
		return "", false
	}
	return key, rec.SeenBefore(key)
}

// SeenBefore sets a HTTP response code as an error for previously seen idempotency keys.
func SeenBefore(w http.ResponseWriter) {
	w.WriteHeader(http.StatusPreconditionFailed)
}

func truncate(s string) string {
	if utf8.RuneCountInString(s) > maxIdempotencyKeyLength {
		return s[:maxIdempotencyKeyLength]
	}
	return s
}
