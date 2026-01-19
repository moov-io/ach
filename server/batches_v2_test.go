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
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/moov-io/ach"
	"github.com/moov-io/base/log"

	kitlog "github.com/go-kit/log"
	"github.com/stretchr/testify/require"
)

func TestBatchesV2__CreateBatchWithErrors(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	svc := NewService(repo)
	handler := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// Create a file first
	f := ach.NewFile()
	f.ID = "test-file-for-batch-v2"
	f.Header = *mockFileHeader()
	err := repo.StoreFile(f)
	require.NoError(t, err)

	// Create an invalid batch
	batch := ach.NewBatchWEB(mockBatchHeaderWeb())
	entry := mockWEBEntryDetail()
	entry.TransactionCode = 0 // invalid
	batch.AddEntry(entry)

	var body bytes.Buffer
	err = json.NewEncoder(&body).Encode(batch)
	require.NoError(t, err)

	req := httptest.NewRequest("POST", "/v2/files/test-file-for-batch-v2/batches", &body)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-request-id", "test-v2-batch-create")

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusBadRequest, w.Code)

	var resp createBatchResponseV2
	err = json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.NotEmpty(t, resp.Errors)
	require.Empty(t, resp.ID)

	// Check structured error format
	for _, verr := range resp.Errors {
		require.NotEmpty(t, verr.ErrorType)
		require.NotEmpty(t, verr.Message)
	}
}

func TestBatchesV2__CreateBatchValid(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	svc := NewService(repo)
	handler := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// Create a file first
	f := ach.NewFile()
	f.ID = "test-file-for-batch-v2-valid"
	f.Header = *mockFileHeader()
	err := repo.StoreFile(f)
	require.NoError(t, err)

	// Create a valid batch
	batch := mockBatchWEB(t)

	var body bytes.Buffer
	err = json.NewEncoder(&body).Encode(batch)
	require.NoError(t, err)

	req := httptest.NewRequest("POST", "/v2/files/test-file-for-batch-v2-valid/batches", &body)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-request-id", "test-v2-batch-create-valid")

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusOK, w.Code)

	var resp createBatchResponseV2
	err = json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.NotEmpty(t, resp.ID)
	require.Empty(t, resp.Errors)
}

func TestBatchesV2__CreateBatchFileNotFound(t *testing.T) {
	repo := NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	svc := NewService(repo)
	handler := MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	batch := mockBatchWEB(t)

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(batch)
	require.NoError(t, err)

	req := httptest.NewRequest("POST", "/v2/files/non-existent-file/batches", &body)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-request-id", "test-v2-batch-not-found")

	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusBadRequest, w.Code)

	var resp createBatchResponseV2
	err = json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.NotEmpty(t, resp.Errors)
	require.Equal(t, "ServiceError", resp.Errors[0].ErrorType)
}
