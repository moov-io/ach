// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package http

import (
	"errors"
	"net/http"
	"time"

	"github.com/moov-io/base/idempotent"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
)

var (
	ErrNoUserId = errors.New("no X-User-Id header provided")
)

type ResponseWriter struct {
	http.ResponseWriter

	start   time.Time
	request *http.Request
	metric  metrics.Histogram

	headersWritten bool // set on WriteHeader

	log log.Logger
}

func (w *ResponseWriter) WriteHeader(code int) {
	if w.headersWritten {
		return
	}
	w.headersWritten = true

	// Headers
	SetAccessControlAllowHeaders(w, w.request.Header.Get("Origin"))
	defer w.ResponseWriter.WriteHeader(code)

	// Record route timing
	diff := time.Since(w.start)
	if w != nil && w.metric != nil {
		w.metric.Observe(diff.Seconds())
	}

	// Skip Go's content sniff here to speed up response timing for client
	if w.ResponseWriter.Header().Get("Content-Type") == "" {
		w.ResponseWriter.Header().Set("Content-Type", "text/plain")
		w.ResponseWriter.Header().Set("X-Content-Type-Options", "nosniff")
	}

	if requestId := GetRequestId(w.request); requestId != "" && w.log != nil {
		w.log.Log("method", w.request.Method, "path", w.request.URL.Path, "status", code, "duration", diff, "requestId", requestId)
	}
}

// Wrap returns a ResponseWriter usable by applications. No parts of the Request are inspected or ResponseWriter modified.
func Wrap(logger log.Logger, m metrics.Histogram, w http.ResponseWriter, r *http.Request) *ResponseWriter {
	now := time.Now()
	return &ResponseWriter{
		ResponseWriter: w,
		start:          now,
		request:        r,
		metric:         m,
		log:            logger,
	}
}

// EnsureHeaders wraps the http.ResponseWriter but also checks Moov specific headers.
//
// X-User-Id is required, and requests without one will be completed with a 403 forbidden.
// No lookup is done to ensure the value exists and is valid for a Moov user.
//
// X-Request-Id is optional, but if used we will emit a log line with that request fulfillment timing
// and the status code.
//
// X-Idempotency-Key is optional, but recommended to ensure requests only execute once. Clients are
// assumed to resend requests many times with the same key. We just need to reply back "already done".
func EnsureHeaders(logger log.Logger, m metrics.Histogram, rec idempotent.Recorder, w http.ResponseWriter, r *http.Request) (*ResponseWriter, error) {
	writer := Wrap(logger, m, w, r)
	return writer, writer.ensureHeaders(rec)
}

// ensureHeaders verifies the headers which Moov apps all cares about.
func (w *ResponseWriter) ensureHeaders(rec idempotent.Recorder) error {
	if userId := GetUserId(w.request); userId == "" {
		if !w.headersWritten {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusForbidden)
		}
		return ErrNoUserId
	}
	if rec != nil {
		if _, seen := idempotent.FromRequest(w.request, rec); seen {
			idempotent.SeenBefore(w)
			return idempotent.ErrSeenBefore
		}
	}
	return nil
}
