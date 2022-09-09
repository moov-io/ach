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

package examples

import (
	"fmt"
	lg "log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/moov-io/ach/server"

	"github.com/go-kit/log"
)

func Example_serverFileCreateJSON() {
	repo := server.NewRepositoryInMemory(24*time.Hour, nil)
	service := server.NewService(repo)
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	handler := server.MakeHTTPHandler(service, repo, logger)

	// Spin up a local HTTP server
	server := httptest.NewServer(handler)
	defer server.Close()

	// Read an Example ach.File in JSON format
	file, err := os.Open(filepath.Join("testdata", "ppd-valid.json"))
	if err != nil {
		lg.Fatal(err)
	}

	// Make our request
	req, err := http.NewRequest("POST", server.URL+"/files/create", file)
	if err != nil {
		lg.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := server.Client().Do(req)
	if err != nil {
		lg.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		lg.Fatalf("got %d HTTP status code", resp.StatusCode)
	}

	fmt.Printf("%s", strconv.Itoa(resp.StatusCode)+"\n")

	// Output: 200
}
