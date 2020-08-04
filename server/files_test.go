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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/ach"
	"github.com/moov-io/base"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func TestFiles__decodeCreateFileRequest(t *testing.T) {
	f := ach.NewFile()
	f.ID = "foo"
	f.Header = *mockFileHeader()
	f.AddBatch(mockBatchWEB())

	// Setup our persistence
	repo := NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	svc := NewService(repo)

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(f); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("POST", "/files/create", &body)
	req.Header.Set("x-request-id", "test")
	req.Header.Set("content-type", "application/json")

	// setup our HTTP handler
	handler := MakeHTTPHandler(svc, repo, log.NewNopLogger())

	// execute our HTTP request
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status code: %d: %s", w.Code, w.Body.String())
	}

	var resp createFileResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	if resp.ID == "" || resp.Err != nil {
		t.Errorf("id=%q error=%v", resp.ID, resp.Err)
	}
}

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

	f := ach.NewFile()
	f.ID = "foo"
	f.Header = *mockFileHeader()
	f.AddBatch(mockBatchWEB())
	if err := repo.StoreFile(f); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("GET", "/files", nil)
	req.Header.Set("x-request-id", "test")
	req.Header.Set("content-type", "application/json")

	// setup our HTTP handler
	handler := MakeHTTPHandler(svc, repo, log.NewNopLogger())

	// execute our HTTP request
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status code: %d: %s", w.Code, w.Body.String())
	}

	// sad path
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

	f := ach.NewFile()
	f.ID = "foo"
	f.Header = *mockFileHeader()
	f.AddBatch(mockBatchWEB())
	if err := repo.StoreFile(f); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("GET", "/files/foo", nil)
	req.Header.Set("x-request-id", "test")
	req.Header.Set("content-type", "application/json")

	// setup our HTTP handler
	handler := MakeHTTPHandler(svc, repo, log.NewNopLogger())

	// execute our HTTP request
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status code: %d: %s", w.Code, w.Body.String())
	}

	// sad path
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

	f := ach.NewFile()
	f.ID = "foo"
	f.Header = *mockFileHeader()
	f.AddBatch(mockBatchWEB())
	if err := repo.StoreFile(f); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("GET", "/files/foo/contents", nil)
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
	if v := w.Header().Get("content-type"); v != "text/plain" {
		t.Errorf("content-type: %s", v)
	}

	// sad path
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
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)

	rawBody := `{"random":"json"}`

	resp, err := validateFileEndpoint(svc, logger)(context.TODO(), strings.NewReader(rawBody))
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
		httptransport.NewServer(validateFileEndpoint(svc, logger), decodeValidateFileRequest, encodeResponse),
	)

	req.Header.Set("Origin", "https://moov.io")
	req.Header.Set("X-Request-Id", "55555")
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}
	if !strings.HasPrefix(w.Body.String(), `{"error":"invalid ACH file: ImmediateDestination`) {
		t.Errorf("unknown error: %v\n%v", err, w.Body.String())
	}
}

func TestFiles__ValidateOpts(t *testing.T) {
	logger := log.NewLogfmtLogger(ioutil.Discard)
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)

	// Write file into storage
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-valid.json"))
	if fd == nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()

	bs, _ := ioutil.ReadAll(fd)
	file, _ := ach.FileFromJSON(bs)
	file.Header.ImmediateOrigin = "123456789" // invalid routing number
	repo.StoreFile(file)

	// validate, expect failure
	w := httptest.NewRecorder()
	body := strings.NewReader(`{"requireABAOrigin": true}`)
	req := httptest.NewRequest("POST", fmt.Sprintf("/files/%s/validate", file.ID), body)

	router := mux.NewRouter()
	router.Methods("POST").Path("/files/{id}/validate").Handler(
		httptransport.NewServer(validateFileEndpoint(svc, logger), decodeValidateFileRequest, encodeResponse),
	)

	req.Header.Set("X-Request-Id", "55555")
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}

	// correct file
	file.Header.ImmediateOrigin = "987654320" // routing number
	repo.StoreFile(file)

	// retry, but with different ValidateOpts
	w = httptest.NewRecorder()
	body = strings.NewReader(`{"requireABAOrigin": true}`)
	req = httptest.NewRequest("POST", fmt.Sprintf("/files/%s/validate", file.ID), body)
	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d: %s", w.Code, w.Body.String())
	}
}

func TestFilesErr__balanceFileEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, nil)
	svc := NewService(repo)

	resp, err := balanceFileEndpoint(svc, repo, log.NewNopLogger())(context.TODO(), nil)
	r, ok := resp.(balanceFileResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil || r.Err == nil {
		t.Errorf("expected error: err=%v resp.Err=%v", err, r.Err)
	}
}

func TestFiles__balanceFileEndpoint(t *testing.T) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)
	router := MakeHTTPHandler(svc, repo, logger)

	// write an ACH file into the repository
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-mixedDebitCredit-valid.json"))
	if err != nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()

	bs, _ := ioutil.ReadAll(fd)
	file, _ := ach.FileFromJSON(bs)
	repo.StoreFile(file)

	w := httptest.NewRecorder()
	body := strings.NewReader(`{"routingNumber": "987654320", "accountNumber": "216112", "accountType": "checking", "description": "OFFSET"}`)
	req := httptest.NewRequest("POST", fmt.Sprintf("/files/%s/balance", file.ID), body)
	req.Header.Set("X-Request-ID", base.ID())

	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}

	var resp balanceFileResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	if resp.FileID == "" {
		t.Error("empty FileID")
	}

	// check for ErrBadRouting
	if _, err := decodeBalanceFileRequest(context.TODO(), &http.Request{}); err != ErrBadRouting {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestFilesErr__balanceInvalidFile(t *testing.T) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)
	router := MakeHTTPHandler(svc, repo, logger)

	// write an invalid (partial) file
	fh := ach.NewFileHeader()
	fileID, err := svc.CreateFile(&fh)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	body := strings.NewReader(`{"routingNumber": "987654320", "accountNumber": "216112", "accountType": "checking", "description": "OFFSET"}`)
	req := httptest.NewRequest("POST", fmt.Sprintf("/files/%s/balance", fileID), body)
	req.Header.Set("X-Request-ID", base.ID())

	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}

	var resp struct {
		Err string `json:"error"`
	}
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	if resp.Err == "" {
		t.Errorf("resp.Err=%q", resp.Err)
	}
}

func TestFilesErr__balanceFileEndpointJSON(t *testing.T) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)
	router := MakeHTTPHandler(svc, repo, logger)

	// write an ACH file into the repository
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-mixedDebitCredit-valid.json"))
	if err != nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()

	bs, _ := ioutil.ReadAll(fd)
	file, _ := ach.FileFromJSON(bs)
	repo.StoreFile(file)

	w := httptest.NewRecorder()

	body := strings.NewReader(`{"routingNumber": "987654320"}`) // partial JSON, but we left off fields
	req := httptest.NewRequest("POST", fmt.Sprintf("/files/%s/balance", file.ID), body)
	req.Header.Set("X-Request-ID", base.ID())

	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusInternalServerError {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}

	var resp struct {
		Err string `json:"error"`
	}
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	if resp.Err == "" {
		t.Errorf("resp.Err=%q", resp.Err)
	}

	// send totally invalid JSON
	w = httptest.NewRecorder()
	body = strings.NewReader(`invalid-json asdlsk`)
	req = httptest.NewRequest("POST", fmt.Sprintf("/files/%s/balance", file.ID), body)
	req.Header.Set("X-Request-ID", base.ID())

	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusInternalServerError {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}
}

// TestFilesError__segmentFileEndpoint test an error returned from segmentFileEndpoint
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

// TestFiles__segmentFileEndpoint tests segmentFileEndpoints
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

	body := strings.NewReader(`{}`)

	// test status code
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", fmt.Sprintf("/files/%s/segment", file.ID), body)
	req.Header.Set("Origin", "https://moov.io")
	req.Header.Set("X-Request-Id", "11111")

	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}

	var resp segmentFileResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	if resp.CreditFileID == "" {
		t.Errorf("empty CreditFileID")
	}
	if resp.DebitFileID == "" {
		t.Errorf("empty DebitFileID")
	}
}

// TestFiles__decodeSegmentFileRequest tests segmentFileEndpoints
func TestFiles__decodeSegmentFileRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "/files/segment", nil)
	req.Header.Set("Origin", "https://moov.io")
	req.Header.Set("X-Request-Id", "11111")

	_, err := decodeSegmentFileRequest(context.TODO(), req)

	if !base.Match(err, ErrBadRouting) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFiles__flattenFileEndpoint tests flattenFileEndpoints
func TestFiles__flattenFileEndpoint(t *testing.T) {
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
	req := httptest.NewRequest("POST", fmt.Sprintf("/files/%s/flatten", file.ID), nil)
	req.Header.Set("Origin", "https://moov.io")
	req.Header.Set("X-Request-Id", "11111")

	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}
}

// TestFilesError__flattenFileEndpoint test an error returned from flattenFileEndpoint
func TestFilesError__flattenFileEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, nil)
	svc := NewService(repo)

	resp, err := flattenBatchesEndpoint(svc, repo, nil)(context.TODO(), nil)
	r, ok := resp.(flattenBatchesResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil || r.Err == nil {
		t.Errorf("expected error: err=%v resp.Err=%v", err, r.Err)
	}

}

// TestFiles__decodeFlattenFileRequest tests segmentFileEndpoints
func TestFiles__decodeFlattenFileRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "/files/flatten", nil)
	req.Header.Set("Origin", "https://moov.io")
	req.Header.Set("X-Request-Id", "11111")

	_, err := decodeFlattenBatchesRequest(context.TODO(), req)

	if !base.Match(err, ErrBadRouting) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFilesByID__getFileEndpoint tests getFileEndpoint by File ID
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

// TestFileContentsByID__getFileContentsEndpoint tests getFileContentsEndpoint by File ID
func TestFileContentsByID__getFileContentsEndpoint(t *testing.T) {
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
	req := httptest.NewRequest("GET", fmt.Sprintf("/files/%s/contents", file.ID), nil)
	req.Header.Set("Origin", "https://moov.io")
	req.Header.Set("X-Request-Id", "11112")

	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}

}

// TestFilesByID__deleteFileEndpoint tests by File ID
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

// TestFilesError__deleteFileEndpoint tests error returned for deleteFileEndpoint
func TestFilesError__deleteFileEndpoint(t *testing.T) {
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

// TestFiles__CreateFileEndpoint test CreateFileEndpoint
func TestFiles__CreateFileEndpoint(t *testing.T) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)
	router := MakeHTTPHandler(svc, repo, logger)

	// write an ACH file into repository
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-debit.ach"))
	if fd == nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()

	// test status code
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/files/create", fd)
	req.Header.Set("Origin", "https://moov.io")
	req.Header.Set("X-Request-Id", "11114")

	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}
}

func TestFiles__CreateFileEndpointErr(t *testing.T) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)
	router := MakeHTTPHandler(svc, repo, logger)

	// write an ACH file into repository
	fd, err := os.Open(filepath.Join("..", "test", "issues", "testdata", "issue702.ach"))
	if fd == nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()

	// test status code
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/files/create", fd)
	req.Header.Set("Origin", "https://moov.io")
	req.Header.Set("X-Request-Id", "11114")

	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d: %s", w.Code, w.Body.String())
	}

	var resp struct {
		ID  string `json:"id"`
		Err string `json:"error"`
	}
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	if resp.ID == "" || resp.Err == "" {
		t.Errorf("resp.ID=%q resp.Err=%q", resp.ID, resp.Err)
	}
}

func TestFiles__CreateFileEndpointJSONErr(t *testing.T) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)
	router := MakeHTTPHandler(svc, repo, logger)

	// write an ACH file into repository
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-noBatches.json"))
	if fd == nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()

	// test status code
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/files/create", fd)
	req.Header.Set("Origin", "https://moov.io")
	req.Header.Set("X-Request-Id", "11114")
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d: %s", w.Code, w.Body.String())
	}

	var resp struct {
		ID  string `json:"id"`
		Err string `json:"error"`
	}
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	if resp.ID == "" || resp.Err == "" {
		t.Errorf("resp.ID=%q resp.Err=%q", resp.ID, resp.Err)
	}
}

// TestFiles_segmentFileEndpointError tests segmentFileEndpoints
func TestFiles__segmentFileEndpointError(t *testing.T) {
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
	file.Header.ImmediateDestination = "" // invalid routing number
	repo.StoreFile(file)

	// test status code
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", fmt.Sprintf("/files/%s/segment", file.ID), nil)
	req.Header.Set("Origin", "https://moov.io")
	req.Header.Set("X-Request-Id", "11110")

	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}

	var resp struct {
		Err string `json:"error"`
	}
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	if resp.Err == "" {
		t.Errorf("resp.Err=%q", resp.Err)
	}
}

// TestFiles_flattenFileEndpointError tests flattenFileEndpoints
func TestFiles__flattenFileEndpointError(t *testing.T) {
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
	file.Header.ImmediateDestination = "" // invalid routing number
	repo.StoreFile(file)

	// test status code
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", fmt.Sprintf("/files/%s/flatten", file.ID), nil)
	req.Header.Set("Origin", "https://moov.io")
	req.Header.Set("X-Request-Id", "11110")

	router.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusBadRequest {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}

	var resp struct {
		Err string `json:"error"`
	}
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	if resp.Err == "" {
		t.Errorf("resp.Err=%q", resp.Err)
	}
}
