// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package server

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/ach"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func TestFiles__createFileEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, nil)
	svc := NewService(repo)

	body := strings.NewReader(`{"random":"json"}`)

	resp, err := createFileEndpoint(svc, repo, nil)(context.TODO(), body)
	r, ok := resp.(createFileResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil || r.Err == nil {
		t.Errorf("expected error: err=%v resp.Err=%v", err, r.Err)
	}

}

func TestFiles__getFilesEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, nil)
	svc := NewService(repo)

	body := strings.NewReader(`{"random":"json"}`)

	resp, err := getFilesEndpoint(svc)(context.TODO(), body)
	_, ok := resp.(getFilesResponse)
	if !ok || err != nil {
		t.Errorf("got %#v : err=%v", resp, err)
	}
}

func TestFiles__getFileEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, nil)
	svc := NewService(repo)

	body := strings.NewReader(`{"random":"json"}`)

	resp, err := getFileEndpoint(svc, nil)(context.TODO(), body)
	r, ok := resp.(getFileResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil || r.Err == nil {
		t.Errorf("expected error: err=%v resp.Err=%v", err, r.Err)
	}

}

func TestFiles__getFileContentsEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, nil)
	svc := NewService(repo)

	body := strings.NewReader(`{"random":"json"}`)

	resp, err := getFileContentsEndpoint(svc, nil)(context.TODO(), body)
	_, ok := resp.(getFileContentsResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil {
		t.Errorf("expected error: err=%v", err)
	}

}

func TestFiles__validateFileEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, nil)
	svc := NewService(repo)

	rawBody := `{"random":"json"}`

	resp, err := validateFileEndpoint(svc, nil)(context.TODO(), strings.NewReader(rawBody))
	r, ok := resp.(validateFileResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil || r.Err == nil {
		t.Errorf("expected error: err=%v resp.Err=%v", err, r.Err)
	}

	// write an ACH file into repository
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-valid.json"))
	if fd == nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()
	bs, _ := ioutil.ReadAll(fd)
	file, _ := ach.FileFromJSON(bs)
	file.Header.ImmediateDestination = "" // invalid routing number
	repo.StoreFile(file)

	// test status code
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", fmt.Sprintf("/files/%s/validate", file.ID), strings.NewReader(rawBody))

	router := mux.NewRouter()
	router.Methods("GET").Path("/files/{id}/validate").Handler(
		httptransport.NewServer(validateFileEndpoint(svc, nil), decodeValidateFileRequest, encodeResponse),
	)

	req.Header.Set("Origin", "https://moov.io")
	req.Header.Set("X-Request-Id", "55555")
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}
	if !strings.HasPrefix(w.Body.String(), `{"error":"invalid ACH file: ImmediateDestination`) {
		t.Errorf("unknown error: %v", err)
	}
}

func TestFilesError__segmentFileEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, nil)
	svc := NewService(repo)

	resp, err := segmentFileEndpoint(svc, repo, nil)(context.TODO(), nil)
	r, ok := resp.(segmentFileResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil || r.Err == nil {
		t.Errorf("expected error: err=%v resp.Err=%v", err, r.Err)
	}

}

func TestFiles__segmentFileEndpoint(t *testing.T) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)
	router := MakeHTTPHandler(svc, repo, logger)

	// write an ACH file into repository
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-mixedDebitCredit-valid.json"))
	if fd == nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()
	bs, _ := ioutil.ReadAll(fd)
	file, _ := ach.FileFromJSON(bs)
	repo.StoreFile(file)

	// test status code
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", fmt.Sprintf("/files/%s/segment", file.ID), nil)
	req.Header.Set("Origin", "https://moov.io")
	req.Header.Set("X-Request-Id", "11111")

	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}
}

func TestFilesByID__getFileEndpoint(t *testing.T) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)
	router := MakeHTTPHandler(svc, repo, logger)

	// write an ACH file into repository
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-mixedDebitCredit-valid.json"))
	if fd == nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()
	bs, _ := ioutil.ReadAll(fd)
	file, _ := ach.FileFromJSON(bs)
	repo.StoreFile(file)

	// test status code
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", fmt.Sprintf("/files/%s", file.ID), nil)
	req.Header.Set("Origin", "https://moov.io")
	req.Header.Set("X-Request-Id", "11112")

	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}

}

func TestFilesByID__deleteFileEndpoint(t *testing.T) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)
	router := MakeHTTPHandler(svc, repo, logger)

	// write an ACH file into repository
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-mixedDebitCredit-valid.json"))
	if fd == nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()
	bs, _ := ioutil.ReadAll(fd)
	file, _ := ach.FileFromJSON(bs)
	repo.StoreFile(file)

	// test status code
	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", fmt.Sprintf("/files/%s", file.ID), nil)
	req.Header.Set("Origin", "https://moov.io")
	req.Header.Set("X-Request-Id", "11113")

	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}

}

func TestFiles__deleteFileEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, nil)
	svc := NewService(repo)

	body := strings.NewReader(`{"random":"json"}`)

	resp, err := deleteFileEndpoint(svc, nil)(context.TODO(), body)
	r, ok := resp.(deleteFileResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil || r.Err == nil {
		t.Errorf("expected error: err=%v resp.Err=%v", err, r.Err)
	}

}
