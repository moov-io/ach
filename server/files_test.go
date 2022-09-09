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
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/moov-io/ach"
	"github.com/moov-io/base"
	"github.com/moov-io/base/log"

	httptransport "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestFiles__decodeCreateFileRequest(t *testing.T) {
	f := ach.NewFile()
	f.ID = "foo"
	f.Header = *mockFileHeader()
	batch := mockBatchWEB()
	f.AddBatch(batch)

	// Setup our persistence
	repo := NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	svc := NewService(repo)

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(f)
	require.NoError(t, err)

	req := httptest.NewRequest("POST", "/files/create", &body)
	req.Header.Set("x-request-id", "test")
	req.Header.Set("content-type", "application/json")

	// setup our HTTP handler
	handler := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

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
	require.Equal(t, "foo", f.ID)
	require.NotNil(t, resp.File)
	require.Equal(t, "foo", resp.File.ID)

	// Check stored file state
	got, _ := svc.GetFile(f.ID)
	if got.Batches[0].ID() != batch.ID() {
		t.Fatalf("batch ID: got %v, want %v", got.Batches[0].ID(), batch.ID())
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

func TestFiles__CreateFileNacha(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	svc := NewService(repo)

	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)
	defer fd.Close()

	fileID := base.ID()
	req := httptest.NewRequest("POST", fmt.Sprintf("/files/%s", fileID), fd)
	req.Header.Set("content-type", "text/html")

	handler := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp createFileResponse
	err = json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	require.Equal(t, fileID, resp.ID)
	require.NotNil(t, resp.File)
	require.Equal(t, fileID, resp.File.ID)

	got, _ := svc.GetFile(fileID)
	require.Len(t, got.Batches, 1)

	entries := got.Batches[0].GetEntries()
	require.Len(t, entries, 1)
	require.Equal(t, "121042880000001", entries[0].TraceNumber)
}

func TestFiles__CustomJsonValidation(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, nil)
	svc := NewService(repo)

	body, err := os.Open(filepath.Join("..", "test", "testdata", "json-bypass-origin-and-destination.json"))
	require.NoError(t, err)

	req := httptest.NewRequest("POST", "/files/create?bypassDestination=true&bypassOrigin=true", body)
	req.Header.Set("content-type", "application/json")

	// setup our HTTP handler
	handler := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// execute our HTTP request
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusOK, w.Code)

	var resp createFileResponse
	err = json.NewDecoder(w.Body).Decode(&resp)
	require.ErrorContains(t, err, "ImmediateDestination 000000000 is a mandatory field")
	require.Equal(t, "adam-01", resp.ID)
	require.NotNil(t, resp.File)
	require.Equal(t, nil, resp.Err)
}

func TestFiles__decodeCreateFileRequest__validateOpts(t *testing.T) {
	f := ach.NewFile()
	f.ID = "foo"
	f.Header = *mockFileHeader()
	f.AddBatch(mockBatchWEB())

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(f); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		query  string
		expect ach.ValidateOpts
	}{
		{
			query: "?bypassCompanyIdentificationMatch=true",
			expect: ach.ValidateOpts{
				BypassCompanyIdentificationMatch: true,
			},
		},
		{
			query: "?bypassOrigin=true&requireABAOrigin=true",
			expect: ach.ValidateOpts{
				RequireABAOrigin:       true,
				BypassOriginValidation: true,
			},
		},
		{
			query: "?bypassDestination=1&requireABAOrigin=0",
			expect: ach.ValidateOpts{
				BypassDestinationValidation: true,
			},
		},
		{
			query: "?requireABAOrigin=TRUE&bypassOrigin=true&bypassDestination=1",
			expect: ach.ValidateOpts{
				RequireABAOrigin:            true,
				BypassOriginValidation:      true,
				BypassDestinationValidation: true,
			},
		},
		{
			query: "?requireABAOrigin=false&bypassOrigin=true&bypassDestination=true",
			expect: ach.ValidateOpts{
				RequireABAOrigin:            false,
				BypassOriginValidation:      true,
				BypassDestinationValidation: true,
			},
		},
		{
			query: "?customTraceNumbers=true",
			expect: ach.ValidateOpts{
				CustomTraceNumbers: true,
			},
		},
		{
			query: "?allowMissingFileHeader=true&allowMissingFileControl=true",
			expect: ach.ValidateOpts{
				AllowMissingFileHeader:  true,
				AllowMissingFileControl: true,
			},
		},
		{
			query: "?unorderedBatchNumbers=true",
			expect: ach.ValidateOpts{
				AllowUnorderedBatchNumbers: true,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.query, func(t *testing.T) {
			httpRequest := httptest.NewRequest("POST", "/files/create"+tc.query, &body)
			httpRequest.Header.Set("x-request-id", "test")

			decodedReq, err := decodeCreateFileRequest(context.TODO(), httpRequest)
			if err != nil {
				t.Fatal(err)
			}
			req := decodedReq.(createFileRequest)
			if !reflect.DeepEqual(tc.expect, *req.validateOpts) {
				t.Fatalf("validateOpts: want %v, got %v", tc.expect, *req.validateOpts)
			}
		})
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
	handler := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

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
	handler := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

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

func TestFiles__buildFileEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, nil)
	svc := NewService(repo)

	f := ach.NewFile()
	f.ID = "foo"
	f.Header = *mockFileHeader()
	f.AddBatch(mockBatchWEB())
	if err := repo.StoreFile(f); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("GET", "/files/foo/build", nil)
	req.Header.Set("x-request-id", "test")
	req.Header.Set("content-type", "application/json")

	// setup our HTTP handler
	handler := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// execute our HTTP request
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusOK, w.Code)

	var response struct {
		File  ach.File `json:"file"`
		Error error    `json:"error"`
	}
	err := json.NewDecoder(w.Body).Decode(&response)
	require.NoError(t, err)

	require.Len(t, response.File.Batches, 1)

	entries := response.File.Batches[0].GetEntries()
	require.Len(t, entries, 1)
	require.Equal(t, "121042880000001", entries[0].TraceNumber)
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
	handler := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

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
	bs, _ := io.ReadAll(fd)
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
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)

	// Write file into storage
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-valid.json"))
	if fd == nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()

	bs, _ := io.ReadAll(fd)
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
	router := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// write an ACH file into the repository
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-mixedDebitCredit-valid.json"))
	if err != nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()

	bs, _ := io.ReadAll(fd)
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
	router := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

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
	router := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// write an ACH file into the repository
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-mixedDebitCredit-valid.json"))
	if err != nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()

	bs, _ := io.ReadAll(fd)
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

// TestFilesError__segmentFileIDEndpoint test an error returned from segmentFileIDEndpoint
func TestFilesError__segmentFileIDEndpoint(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, nil)
	svc := NewService(repo)

	resp, err := segmentFileIDEndpoint(svc, repo, nil)(context.TODO(), nil)
	r, ok := resp.(segmentedFilesResponse)
	if !ok {
		t.Errorf("got %#v", resp)
	}
	if err == nil || r.Err == nil {
		t.Errorf("expected error: err=%v resp.Err=%v", err, r.Err)
	}
}

// TestFiles__segmentFileIDEndpoint tests segmentFileIDEndpoints
func TestFiles__segmentFileIDEndpoint(t *testing.T) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)
	router := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// write an ACH file into repository
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-mixedDebitCredit-valid.json"))
	if fd == nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()
	bs, _ := io.ReadAll(fd)
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

	var resp segmentedFilesResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	require.NotEmpty(t, resp.CreditFileID)
	require.NotNil(t, resp.CreditFile)
	require.NotEmpty(t, resp.DebitFileID)
	require.NotNil(t, resp.DebitFile)
}

func TestFiles__segmentFileEndpoint(t *testing.T) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)
	router := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-mixedDebitCredit.ach"))
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/segment", fd)
	req.Header.Set("Origin", "https://moov.io")
	req.Header.Set("X-Request-Id", "222222")

	router.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusOK, w.Code)

	var resp segmentedFilesResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	require.NotEmpty(t, resp.CreditFileID)
	require.NotNil(t, resp.CreditFile)
	require.NotEmpty(t, resp.DebitFileID)
	require.NotNil(t, resp.DebitFile)
}

func TestFiles__segmentFileEndpointJSON(t *testing.T) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)
	router := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	file, err := ach.ReadFile(filepath.Join("..", "test", "testdata", "ppd-mixedDebitCredit.ach"))
	require.NoError(t, err)

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(struct {
		File *ach.File `json:"file"`
	}{
		File: file,
	})
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/segment", &buf)
	req.Header.Set("Origin", "https://moov.io")
	req.Header.Set("X-Request-Id", "222222")
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusOK, w.Code)

	var resp segmentedFilesResponse
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatal(err)
	}
	require.NotEmpty(t, resp.CreditFileID)
	require.NotNil(t, resp.CreditFile)
	require.NotEmpty(t, resp.DebitFileID)
	require.NotNil(t, resp.DebitFile)
}

// TestFiles__decodeSegmentFileIDRequest tests segmentFileIDEndpoints
func TestFiles__decodeSegmentFileIDRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "/files/segment", nil)
	req.Header.Set("Origin", "https://moov.io")
	req.Header.Set("X-Request-Id", "11111")

	_, err := decodeSegmentFileIDRequest(context.TODO(), req)

	if !base.Match(err, ErrBadRouting) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFiles__flattenFileEndpoint tests flattenFileEndpoints
func TestFiles__flattenFileEndpoint(t *testing.T) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)
	router := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// write an ACH file into repository
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-mixedDebitCredit-valid.json"))
	if fd == nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()
	bs, _ := io.ReadAll(fd)
	file, _ := ach.FileFromJSON(bs)
	repo.StoreFile(file)

	// test status code
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", fmt.Sprintf("/files/%s/flatten", file.ID), nil)
	req.Header.Set("Origin", "https://moov.io")
	req.Header.Set("X-Request-Id", "11111")

	router.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusOK, w.Code)

	var resp flattenBatchesResponse
	err = json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	require.NotEqual(t, file.ID, resp.ID)
	require.NotNil(t, resp.File)
	require.NotEqual(t, file.ID, resp.File.ID)
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

// TestFiles__decodeFlattenFileRequest tests segmentFileIDEndpoints
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
	router := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// write an ACH file into repository
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-mixedDebitCredit-valid.json"))
	if fd == nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()
	bs, _ := io.ReadAll(fd)
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
	router := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// write an ACH file into repository
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-mixedDebitCredit-valid.json"))
	if fd == nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()
	bs, _ := io.ReadAll(fd)
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
	router := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// write an ACH file into repository
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-mixedDebitCredit-valid.json"))
	if fd == nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()
	bs, _ := io.ReadAll(fd)
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
	router := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	testCases := []struct {
		filename           string
		queryParams        string
		expectedStatusCode int
	}{
		{
			filename:           "ppd-debit.ach",
			queryParams:        "",
			expectedStatusCode: http.StatusOK,
		},
		{
			filename:           "ppd-debit-customTraceNumber.ach",
			queryParams:        "?bypassOrigin=true",
			expectedStatusCode: http.StatusOK,
		},
		{
			filename:           "ppd-debit-customTraceNumber.ach",
			queryParams:        "?bypassOrigin=false",
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		// write an ACH file into repository
		fd, err := os.Open(filepath.Join("..", "test", "testdata", tc.filename))
		if fd == nil {
			t.Fatalf("empty ACH file: %v", err)
		}
		defer fd.Close()

		// test status code
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", fmt.Sprintf("/files/create%s", tc.queryParams), fd)
		req.Header.Set("Origin", "https://moov.io")
		req.Header.Set("X-Request-Id", "11114")

		router.ServeHTTP(w, req)
		w.Flush()

		require.Equal(t, tc.expectedStatusCode, w.Code)

		var resp struct {
			ID string `json:"id"`
		}
		err = json.NewDecoder(w.Body).Decode(&resp)
		require.NoError(t, err)

		require.NotEmpty(t, resp.ID)
		require.NotEqual(t, "create", resp.ID)
	}
}

func TestFiles__CreateFileWithZeroBatches(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	svc := NewService(repo)
	handler := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	f := ach.NewFile()
	f.ID = "foo"
	f.Header = *mockFileHeader()

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(f); err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest("POST", "/files/create?allowZeroBatches=true", &body)
	req.Header.Set("content-type", "application/json")

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	// Get contents back
	req = httptest.NewRequest("GET", fmt.Sprintf("/files/%s/contents", f.ID), nil)
	w = httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	got, err := ach.NewReader(w.Body).Read()
	if err != nil {
		t.Fatalf("reading file: %v", err)
	}
	if err := got.Validate(); err != nil {
		t.Fatalf("validating file: %v", err)
	}
}

func TestFiles__CreateFileEndpointErr(t *testing.T) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)
	router := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

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
	router := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

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

// TestFiles_segmentFileIDEndpointError tests segmentFileIDEndpoints
func TestFiles__segmentFileIDEndpointError(t *testing.T) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)
	router := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// write an ACH file into repository
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-mixedDebitCredit-valid.json"))
	if fd == nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()
	bs, _ := io.ReadAll(fd)
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
	router := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// write an ACH file into repository
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-mixedDebitCredit-valid.json"))
	if fd == nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()
	bs, _ := io.ReadAll(fd)
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
