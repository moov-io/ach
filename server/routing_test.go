// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"encoding/json"
	"github.com/moov-io/ach"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	httptransport "github.com/go-kit/kit/transport/http"
)

func TestEncodeResponse(t *testing.T) {
	ctx := context.TODO()
	w := httptest.NewRecorder()
	if err := encodeResponse(ctx, w, "hi mom"); err != nil {
		t.Fatal(err)
	}
	w.Flush()

	var resp string
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Error(err)
	}
	if resp != "hi mom" {
		t.Errorf("got %q", resp)
	}

	v := w.Header().Get("content-type")
	if v != "application/json; charset=utf-8" {
		t.Errorf("got %q", v)
	}
}

func TestEncodeTextResponse(t *testing.T) {
	ctx := context.TODO()
	w := httptest.NewRecorder()
	if err := encodeTextResponse(ctx, w, strings.NewReader("hi mom")); err != nil {
		t.Fatal(err)
	}
	if v := w.Body.String(); v != "hi mom" {
		t.Errorf("got %q", v)
	}

	if v := w.Header().Get("content-type"); v != "text/plain" {
		t.Errorf("got %q", v)
	}
}

func TestFilesXTotalCountHeader(t *testing.T) {
	counter := getFilesResponse{
		Files: []*ach.File{ach.NewFile()},
		Err:   nil,
	}

	w := httptest.NewRecorder()
	encodeResponse(context.Background(), w, counter)

	actual, ok := w.Result().Header["X-Total-Count"]
	if !ok {
		t.Fatal("should have count")
	}
	if actual[0] != "1" {
		t.Errorf("should be 1, got %v", actual[0])
	}
}

func TestBatchesXTotalCountHeader(t *testing.T) {
	bh := mockBatchHeaderWeb()
	entry := mockWEBEntryDetail()
	// build the batch
	batch := ach.NewBatchWEB(bh)
	batch.SetID(batch.Header.ID)
	batch.AddEntry(entry)

	counter := getBatchesResponse{
		Batches: []ach.Batcher{batch},
		Err:     nil,
	}

	w := httptest.NewRecorder()
	encodeResponse(context.Background(), w, counter)

	actual, ok := w.Result().Header["X-Total-Count"]
	if !ok {
		t.Fatal("should have count")
	}
	if actual[0] != "1" {
		t.Errorf("should be 1, got %v", actual[0])
	}
}

func TestRouting__CORSHeaders(t *testing.T) {
	ctx := context.TODO()
	req := httptest.NewRequest("GET", "/files/create", nil)
	req.Header.Set("Origin", "https://api.moov.io")

	ctx = saveCORSHeadersIntoContext()(ctx, req)

	w := httptest.NewRecorder()
	respondWithSavedCORSHeaders()(ctx, w)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("expected no status code, but got %d", w.Code)
	}
	if v := w.Header().Get("Content-Type"); v != "" {
		t.Errorf("expected no Content-Type, but got %q", v)
	}

	// check CORS headers
	if v := w.Header().Get("Access-Control-Allow-Origin"); v != "https://api.moov.io" {
		t.Errorf("got %q", v)
	}
	if v := w.Header().Get("Access-Control-Allow-Methods"); v == "" {
		t.Error("missing Access-Control-Allow-Methods")
	}
	if v := w.Header().Get("Access-Control-Allow-Headers"); v == "" {
		t.Error("missing Access-Control-Allow-Headers")
	}
	if v := w.Header().Get("Access-Control-Allow-Credentials"); v == "" {
		t.Error("missing Access-Control-Allow-Credentials")
	}
}

func TestPreflightHandler(t *testing.T) {
	options := []httptransport.ServerOption{
		httptransport.ServerBefore(saveCORSHeadersIntoContext()),
		httptransport.ServerAfter(respondWithSavedCORSHeaders()),
	}

	handler := preflightHandler(options)

	// Make our pre-flight request
	w := httptest.NewRecorder()
	r := httptest.NewRequest("OPTIONS", "/files/create", nil)
	r.Header.Set("Origin", "https://moov.io")

	// Make the request
	handler.ServeHTTP(w, r)
	w.Flush()

	// Check response
	if v := w.Header().Get("Access-Control-Allow-Origin"); v != "https://moov.io" {
		t.Errorf("got %s", v)
	}
	if v := w.Header().Get("Access-Control-Allow-Methods"); v == "" {
		t.Error("missing Access-Control-Allow-Methods")
	}
	if v := w.Header().Get("Access-Control-Allow-Headers"); v == "" {
		t.Error("missing Access-Control-Allow-Headers")
	}
	if v := w.Header().Get("Access-Control-Allow-Credentials"); v == "" {
		t.Error("missing Access-Control-Allow-Credentials")
	}
}
