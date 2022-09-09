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

package test

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
)

func TestIssue786(t *testing.T) {
	repo := server.NewRepositoryInMemory(0*time.Second, log.NewNopLogger())
	svc := server.NewService(repo)
	handler := server.MakeHTTPHandler(svc, repo, kitlog.NewNopLogger())

	// create the file
	fd, err := os.Open(filepath.Join("testdata", "1-create.json"))
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/files/create", fd)
	req.Header.Set("Content-Type", "application/json")
	handler.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d", w.Code)
		t.Fatalf("body: %v", w.Body.String())
	}

	var file struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(w.Body).Decode(&file); err != nil || file.ID == "" {
		t.Fatal(err)
	}

	// file created, so add the batch now
	fd, err = os.Open(filepath.Join("testdata", "2-add-batch.json"))
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()

	w = httptest.NewRecorder()
	req = httptest.NewRequest("POST", fmt.Sprintf("/files/%s/batches", file.ID), fd)
	req.Header.Set("Content-Type", "application/json")
	handler.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d", w.Code)
		t.Fatalf("body: %v", w.Body.String())
	}

	// read the file and verify
	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", fmt.Sprintf("/files/%s/contents", file.ID), nil)
	handler.ServeHTTP(w, req)
	w.Flush()

	if w.Code != http.StatusOK {
		t.Errorf("bogus HTTP status: %d", w.Code)
	}

	created, err := ach.NewReader(w.Body).Read()
	if err != nil {
		t.Fatal(err)
	}
	if len(created.Batches) != 2 {
		t.Errorf("got %d batches", len(created.Batches))
	}
}
