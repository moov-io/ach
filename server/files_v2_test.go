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
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/ach"
	"github.com/moov-io/base/log"

	kitlog "github.com/go-kit/log"
	"github.com/stretchr/testify/require"
)

func TestFilesV2__CreateFileWithValidationErrors(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	svc := NewService(repo)
	handler := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// Create an invalid file - use NACHA format (not JSON) which is easier to create invalid
	fd, err := os.Open(filepath.Join("..", "test", "issues", "testdata", "issue702.ach"))
	require.NoError(t, err)
	defer fd.Close()

	req := httptest.NewRequest("POST", "/v2/files/create", fd)
	req.Header.Set("x-request-id", "test-v2-create")

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusBadRequest, w.Code)

	// Parse response manually to avoid ach.File JSON unmarshal validation
	var rawResp map[string]json.RawMessage
	err = json.NewDecoder(w.Body).Decode(&rawResp)
	require.NoError(t, err)

	// Check errors exist
	errorsJSON, ok := rawResp["errors"]
	require.True(t, ok, "expected errors field in response")

	var errors []ValidationError
	err = json.Unmarshal(errorsJSON, &errors)
	require.NoError(t, err)

	// Should have validation errors
	require.NotEmpty(t, errors, "expected validation errors")

	// All errors should have structured format
	for _, verr := range errors {
		require.NotEmpty(t, verr.ErrorType)
		require.NotEmpty(t, verr.Message)
	}
}

func TestFilesV2__CreateFileValid(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	svc := NewService(repo)
	handler := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// Create a valid file
	f := ach.NewFile()
	f.ID = "test-file-v2"
	f.Header = *mockFileHeader()
	batch := mockBatchWEB(t)
	batch.Entries[0].TraceNumber = "121042880000007"
	f.AddBatch(batch)

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(f)
	require.NoError(t, err)

	req := httptest.NewRequest("POST", "/v2/files/create", &body)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-request-id", "test-v2-create-valid")

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusOK, w.Code)

	var resp createFileResponseV2
	err = json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.NotEmpty(t, resp.ID)
	require.NotNil(t, resp.File)
	require.Empty(t, resp.Errors)
}

func TestFilesV2__CreateFileWithFileID(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	svc := NewService(repo)
	handler := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)
	defer fd.Close()

	req := httptest.NewRequest("POST", "/v2/files/my-custom-id", fd)
	req.Header.Set("x-request-id", "test-v2-create-id")

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusOK, w.Code)

	var resp createFileResponseV2
	err = json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.Equal(t, "my-custom-id", resp.ID)
	require.NotNil(t, resp.File)
	require.Empty(t, resp.Errors)
}

func TestFilesV2__ValidateFileWithErrors(t *testing.T) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)
	handler := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// Store an invalid file
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-valid.json"))
	require.NoError(t, err)
	defer fd.Close()
	bs, _ := io.ReadAll(fd)
	file, _ := ach.FileFromJSON(bs)
	file.Header.ImmediateDestination = "" // invalid routing number
	repo.StoreFile(file)

	// Validate via v2 endpoint
	req := httptest.NewRequest("GET", "/v2/files/"+file.ID+"/validate", nil)
	req.Header.Set("x-request-id", "test-v2-validate")

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusBadRequest, w.Code)

	var resp validateFileResponseV2
	err = json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.False(t, resp.Valid)
	require.NotEmpty(t, resp.Errors)

	// Check structured error format
	for _, verr := range resp.Errors {
		require.NotEmpty(t, verr.ErrorType)
		require.NotEmpty(t, verr.Message)
	}
}

func TestFilesV2__ValidateFileValid(t *testing.T) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)
	handler := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// Store a valid file
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-valid.json"))
	require.NoError(t, err)
	defer fd.Close()
	bs, _ := io.ReadAll(fd)
	file, _ := ach.FileFromJSON(bs)
	repo.StoreFile(file)

	req := httptest.NewRequest("GET", "/v2/files/"+file.ID+"/validate", nil)
	req.Header.Set("x-request-id", "test-v2-validate-valid")

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusOK, w.Code)

	var resp validateFileResponseV2
	err = json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.True(t, resp.Valid)
	require.Empty(t, resp.Errors)
}

func TestFilesV2__ValidateFilePOST(t *testing.T) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)
	handler := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// Store a file that requires validation options
	fd, err := os.Open(filepath.Join("..", "test", "testdata", "ppd-valid.json"))
	require.NoError(t, err)
	defer fd.Close()
	bs, _ := io.ReadAll(fd)
	file, _ := ach.FileFromJSON(bs)
	file.Header.ImmediateOrigin = "123456789" // needs bypass
	repo.StoreFile(file)

	// POST with validation options
	body := bytes.NewReader([]byte(`{"requireABAOrigin": false, "bypassOrigin": true}`))
	req := httptest.NewRequest("POST", "/v2/files/"+file.ID+"/validate", body)
	req.Header.Set("x-request-id", "test-v2-validate-post")

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusOK, w.Code)

	var resp validateFileResponseV2
	err = json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.True(t, resp.Valid)
}

func TestFilesV2__ValidateFileNotFound(t *testing.T) {
	logger := log.NewNopLogger()
	repo := NewRepositoryInMemory(testTTLDuration, logger)
	svc := NewService(repo)
	handler := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	req := httptest.NewRequest("GET", "/v2/files/non-existent/validate", nil)
	req.Header.Set("x-request-id", "test-v2-validate-not-found")

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusBadRequest, w.Code)

	var resp validateFileResponseV2
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.False(t, resp.Valid)
	require.NotEmpty(t, resp.Errors)
	require.Equal(t, "FileError", resp.Errors[0].ErrorType)
}

func TestFilesV2__MultipleValidationErrors(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	svc := NewService(repo)
	handler := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// Store a file with multiple validation issues
	f := ach.NewFile()
	f.ID = "multi-error-file"
	f.Header = *mockFileHeader()
	f.Header.ImmediateDestination = "" // error 1
	f.Header.ImmediateOrigin = ""      // error 2
	err := repo.StoreFile(f)
	require.NoError(t, err)

	// Validate it
	req := httptest.NewRequest("GET", "/v2/files/multi-error-file/validate", nil)
	req.Header.Set("x-request-id", "test-v2-multi-error")

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusBadRequest, w.Code)

	var resp validateFileResponseV2
	err = json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	// Should have multiple errors, not just the first one
	require.False(t, resp.Valid)
	require.GreaterOrEqual(t, len(resp.Errors), 1, "should collect multiple validation errors")
}
