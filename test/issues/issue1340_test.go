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

	"github.com/moov-io/ach/server"
	"github.com/moov-io/base/log"

	kitlog "github.com/go-kit/log"
	"github.com/stretchr/testify/require"
)

// TestIssue1340 reproduces an issue with setting the offset on a batch with the HTTP server.
//
// Here is what I did:
//   - I created an ACH JSON structure as expected by the ACH API
//   - I added the offset instructions
//   - I called the "create" API endpoint
//   - I called the "contents" API endpoint which gives me the raw ACH file
func TestIssue1340(t *testing.T) {
	repo := server.NewRepositoryInMemory(0*time.Second, log.NewNopLogger())
	svc := server.NewService(repo)
	handler := server.MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	fd, err := os.Open(filepath.Join("testdata", "issue1340.json"))
	require.NoError(t, err)

	// "I created an ACH JSON structure as expected by the ACH API"
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
	require.NotEmpty(t, createResponse.ID)

	t.Logf("file ID: %v", createResponse.ID)

	// Verify the file's contents contain the Offset
	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", fmt.Sprintf("/files/%s/contents", createResponse.ID), nil)
	handler.ServeHTTP(w, req)
	w.Flush()
	require.Equal(t, http.StatusOK, w.Code)

	require.Contains(t, w.Body.String(), "Jane Smith")
	require.Contains(t, w.Body.String(), "OFFSET")
}
