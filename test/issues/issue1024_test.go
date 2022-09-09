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

package issues

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/moov-io/ach"
	"github.com/moov-io/ach/server"
	"github.com/moov-io/base/log"

	kitlog "github.com/go-kit/log"
	"github.com/stretchr/testify/require"
)

func TestIssue1024__Read(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("testdata", "issue1024.json"))
	require.NoError(t, err)

	file, err := ach.FileFromJSON(bs)
	require.NoError(t, err)
	require.Len(t, file.Batches, 1)

	entries := file.Batches[0].GetEntries()
	require.Len(t, entries, 1)

	// I expected the traceNumber field in addenda99 to be automatically
	// populated equal to the traceNumber of the entry
	ed := entries[0]
	require.Equal(t, "084106760000001", ed.TraceNumber)
	require.NotNil(t, ed.Addenda99)
	require.Equal(t, "084106760000001", ed.Addenda99.TraceNumber)
}

func TestIssue1024__Server(t *testing.T) {
	repo := server.NewRepositoryInMemory(0*time.Second, log.NewNopLogger())
	svc := server.NewService(repo)
	handler := server.MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	fd, err := os.Open(filepath.Join("testdata", "issue1024.json"))
	require.NoError(t, err)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/files/create", fd)
	req.Header.Set("Content-Type", "application/json")
	handler.ServeHTTP(w, req)
	w.Flush()
	require.Equal(t, http.StatusOK, w.Code)

	var createResponse struct {
		ID string `json:"id"`
	}
	err = json.NewDecoder(w.Body).Decode(&createResponse)
	require.NoError(t, err)

	// Read the full file in JSON format
	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", fmt.Sprintf("/files/%s/build", createResponse.ID), nil)
	req.Header.Set("Content-Type", "application/json")
	handler.ServeHTTP(w, req)
	w.Flush()
	require.Equal(t, http.StatusOK, w.Code)

	var wrapper struct {
		File ach.File `json:"file"`
		Err  error    `json:"error"`
	}
	err = json.NewDecoder(w.Body).Decode(&wrapper)
	require.NoError(t, err)

	require.Len(t, wrapper.File.Batches, 1)

	entries := wrapper.File.Batches[0].GetEntries()
	require.Len(t, entries, 1)

	// I expected the traceNumber field in addenda99 to be automatically
	// populated equal to the traceNumber of the entry
	ed := entries[0]
	require.Equal(t, "084106760000001", ed.TraceNumber)
	require.NotNil(t, ed.Addenda99)
	require.Equal(t, "084106760000001", ed.Addenda99.TraceNumber)
}
