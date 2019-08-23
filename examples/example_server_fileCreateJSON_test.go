// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package examples

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/moov-io/ach/server"
	lg "log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func Example_serverFileCreateJSON() {
	repo := server.NewRepositoryInMemory(24*time.Hour, nil)
	service := server.NewService(repo)
	logger := log.NewLogfmtLogger(os.Stderr)
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
