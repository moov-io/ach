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

package issue1082

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/moov-io/ach/server"
	"github.com/moov-io/base"
	"github.com/moov-io/base/log"

	kitlog "github.com/go-kit/log"
	"github.com/stretchr/testify/require"
)

func TestIssue1082(t *testing.T) {
	logger := log.NewTestLogger()
	r := server.NewRepositoryInMemory(1*time.Minute, logger)
	svc := server.NewService(r)
	router := server.MakeHTTPHandler(svc, r, kitlog.NewNopLogger())

	// Create file
	fileID := base.ID()
	createFile, err := os.Open("1-create.json")
	require.NoError(t, err)
	defer createFile.Close()

	req := httptest.NewRequest("POST", fmt.Sprintf("/files/%s", fileID), createFile)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Body.String(), fileID)

	// Add second batch
	addBatch, err := os.Open("2-add.json")
	require.NoError(t, err)
	defer createFile.Close()

	req = httptest.NewRequest("POST", fmt.Sprintf("/files/%s/batches", fileID), addBatch)
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	require.Equal(t, http.StatusOK, w.Code)

	// Flatten batches together
	req = httptest.NewRequest("POST", fmt.Sprintf("/files/%s/flatten", fileID), nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	w.Flush()

	// require.Equal(t, http.StatusOK, w.Code)

	fmt.Printf("\n\n%s\n\n", w.Body.String())
}
