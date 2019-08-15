// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package server

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/moov-io/ach"

	"github.com/go-kit/kit/log"
)

func TestFiles__decodeCreateBatchRequest(t *testing.T) {
	f := ach.NewFile()
	f.ID = "foo"

	// Setup our persistence
	repo := NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	svc := NewService(repo)
	if err := repo.StoreFile(f); err != nil {
		t.Fatal(err)
	}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(mockBatchWEB()); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("POST", "/files/foo/batches", &body)
	req.Header.Set("x-request-id", "test")

	// setup our HTTP handler
	handler := MakeHTTPHandler(svc, repo, log.NewNopLogger())

	// execute our HTTP request
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status code: %d: %s", w.Code, w.Body.String())
	}

	// bad JSON body
	body.Reset()
	if _, err := body.WriteString(`{"batchHeader": "expected-an-object"}`); err != nil {
		t.Fatal(err)
	}
	req = httptest.NewRequest("POST", "/files/foo/batches", &body)
	req.Header.Set("x-request-id", "test")

	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusInternalServerError {
		t.Errorf("bogus HTTP status code: %d: %s", w.Code, w.Body.String())
	}
}

func TestFiles__createBatchEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	svc := NewService(repo)

	body := strings.NewReader(`{"random":"json"}`)

	resp, err := createBatchEndpoint(svc, log.NewNopLogger())(context.TODO(), body)
	r, ok := resp.(createBatchResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil || r.Err == nil {
		t.Errorf("expected error: err=%v resp.Err=%v", err, r.Err)
	}

	f := ach.NewFile()
	f.ID = "create-batch"
	if err := repo.StoreFile(f); err != nil {
		t.Fatal(err)
	}

	// successful batch
	resp, err = createBatchEndpoint(svc, log.NewNopLogger())(context.TODO(), createBatchRequest{
		FileID: f.ID,
		Batch:  &mockBatchWEB().Batch,
	})
	if r, ok := resp.(createBatchResponse); ok {
		if r.ID != "54321" || err != nil {
			t.Errorf("id=%s error=%v", r.ID, r.Err)
		}
	} else {
		t.Errorf("%T %#v", resp, resp)
	}
}

func TestFiles__decodeGetBatchesRequest(t *testing.T) {
	f := ach.NewFile()
	f.ID = "foo"

	// Setup our persistence
	repo := NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	svc := NewService(repo)
	if err := repo.StoreFile(f); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("GET", "/files/foo/batches", nil)
	req.Header.Set("x-request-id", "test")

	// setup our HTTP handler
	handler := MakeHTTPHandler(svc, repo, log.NewNopLogger())

	// execute our HTTP request
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status code: %d: %s", w.Code, w.Body.String())
	}
}

func TestFiles__getBatchesEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	svc := NewService(repo)

	body := strings.NewReader(`{"random":"json"}`)

	resp, err := getBatchesEndpoint(svc, log.NewNopLogger())(context.TODO(), body)
	r, ok := resp.(getBatchesResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil || r.Err == nil {
		t.Errorf("expected error: err=%v resp.Err=%v", err, r.Err)
	}

	// successful batch
	f := ach.NewFile()
	f.ID = "get-batches"
	f.AddBatch(mockBatchWEB())
	if err := repo.StoreFile(f); err != nil {
		t.Fatal(err)
	}
	resp, err = getBatchesEndpoint(svc, log.NewNopLogger())(context.TODO(), getBatchesRequest{
		fileID: f.ID,
	})
	if r, ok := resp.(getBatchesResponse); ok {
		if len(r.Batches) != 1 {
			t.Errorf("got %d Batches=%#v", len(r.Batches), r.Batches)
		}
		if err != nil {
			t.Error(r.Err)
		}
	} else {
		t.Errorf("%T %#v", resp, resp)
	}
}

func TestFiles__decodeGetBatchRequest(t *testing.T) {
	f := ach.NewFile()
	f.ID = "foo"
	b := mockBatchWEB()
	b.SetID("foo2")
	f.AddBatch(b)

	// Setup our persistence
	repo := NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	svc := NewService(repo)
	if err := repo.StoreFile(f); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("GET", "/files/foo/batches/foo2", nil)
	req.Header.Set("x-request-id", "test")

	// setup our HTTP handler
	handler := MakeHTTPHandler(svc, repo, log.NewNopLogger())

	// execute our HTTP request
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status code: %d: %s", w.Code, w.Body.String())
	}
}

func TestFiles__getBatchEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	svc := NewService(repo)

	body := strings.NewReader(`{"random":"json"}`)

	resp, err := getBatchEndpoint(svc, log.NewNopLogger())(context.TODO(), body)
	r, ok := resp.(getBatchResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil || r.Err == nil {
		t.Errorf("expected error: err=%v resp.Err=%v", err, r.Err)
	}

	// successful batch
	f := ach.NewFile()
	f.ID = "get-batch"
	b := mockBatchWEB()
	f.AddBatch(b)
	if err := repo.StoreFile(f); err != nil {
		t.Fatal(err)
	}
	resp, err = getBatchEndpoint(svc, log.NewNopLogger())(context.TODO(), getBatchRequest{
		fileID:  f.ID,
		batchID: b.ID(),
	})
	if r, ok := resp.(getBatchResponse); ok {
		if r.Batch == nil {
			t.Error("nil ach.Batcher")
		}
		if err != nil {
			t.Error(r.Err)
		}
	} else {
		t.Errorf("%T %#v", resp, resp)
	}
}

func TestFiles__decodeDeleteBatchRequest(t *testing.T) {
	f := ach.NewFile()
	f.ID = "foo"
	b := mockBatchWEB()
	b.SetID("foo2")
	f.AddBatch(b)

	// Setup our persistence
	repo := NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	svc := NewService(repo)
	if err := repo.StoreFile(f); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("DELETE", "/files/foo/batches/foo2", nil)
	req.Header.Set("x-request-id", "test")

	// setup our HTTP handler
	handler := MakeHTTPHandler(svc, repo, log.NewNopLogger())

	// execute our HTTP request
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status code: %d: %s", w.Code, w.Body.String())
	}
}

func TestFiles__deleteBatchEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	svc := NewService(repo)

	body := strings.NewReader(`{"random":"json"}`)

	resp, err := deleteBatchEndpoint(svc, log.NewNopLogger())(context.TODO(), body)
	r, ok := resp.(deleteBatchResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil || r.Err == nil {
		t.Errorf("expected error: err=%v resp.Err=%v", err, r.Err)
	}

	// successful batch
	f := ach.NewFile()
	f.ID = "delete-batch"
	b := mockBatchWEB()
	f.AddBatch(b)
	if err := repo.StoreFile(f); err != nil {
		t.Fatal(err)
	}
	resp, err = deleteBatchEndpoint(svc, log.NewNopLogger())(context.TODO(), deleteBatchRequest{
		fileID:  f.ID,
		batchID: b.ID(),
	})
	if r, ok := resp.(deleteBatchResponse); ok {
		if err != nil {
			t.Error(r.Err)
		}
	} else {
		t.Errorf("%T %#v", resp, resp)
	}
}
