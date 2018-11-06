// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package serverexample

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/moov-io/ach/server"

	"github.com/go-kit/kit/log"
)

func TestServer__CreateFile(t *testing.T) {
	// Local server setup - usually ach would be running on another machine.
	repo := server.NewRepositoryInMemory()
	service := server.NewService(repo)
	logger := log.NewLogfmtLogger(os.Stderr)
	handler := server.MakeHTTPHandler(service, repo, logger)

	// Spin up a local HTTP server
	server := httptest.NewServer(handler)
	defer server.Close()

	// Read an Example ach.File in JSON format
	file, err := os.Open("../testdata/ppd-valid.json")
	if err != nil {
		t.Fatal(err)
	}

	// Make our request
	req, err := http.NewRequest("POST", server.URL+"/files/create", file)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := server.Client().Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("got %d HTTP status code", resp.StatusCode)
	}
}
