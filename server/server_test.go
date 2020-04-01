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
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/moov-io/ach"
)

// TestServer__CreateFileEndpoint creates JSON from existing ACH Files and submits them to our
// HTTP API. We do this to ensure all SEC codes can be submitted and created via the HTTP API.
func TestServer__CreateFileEndpoint(t *testing.T) {
	files := getTestFiles()
	if len(files) == 0 {
		t.Fatal("got no test ACH files to process")
	}
	t.Logf("read %d test files", len(files))

	for _, file := range files {
		f, err := os.Open(file.ACHFilepath)
		if err != nil {
			log.Fatal(err)
		}

		achFile, err := ach.NewReader(f).Read()
		if err != nil {
			fmt.Printf("Issue reading file: %+v \n", err)
		}

		// ensure we have a validated file structure
		if err := achFile.Validate(); err != nil {
			t.Errorf("Could not validate entire read file: %v", err)
		}

		// If you trust the file but it's formatting is off building will probably resolve the malformed file.
		if err := achFile.Create(); err != nil {
			t.Errorf("Could not build file with read properties: %v", err)
		}

		if err := f.Close(); err != nil {
			t.Errorf("Problem closing %s: %v", file.ACHFilepath, err)
		}

		// ENR ACH Files does not have BatchHeader.EffectiveEntryDate, so setting this to Today +1 to be included
		// in the JSON File.  For this test after the ACH file is converted to JSON, the test validates the JSON by
		// calling ach.FileFromJSON(bs) and it fails with an empty date time.
		if file.SECCode == "ENR" {
			for _, batch := range achFile.Batches {
				batch.GetHeader().EffectiveEntryDate = time.Now().AddDate(0, 0, 1).Format("060102")
			}

		}

		// Marshal the ach.File into JSON for HTTP API submission
		bs, err := json.Marshal(achFile)
		if err != nil {
			t.Fatalf("Problem converting %s to JSON: %v", file.ACHFilepath, err)
		}

		httpReq, err := http.NewRequest("POST", "/files/create", bytes.NewReader(bs))
		if err != nil {
			t.Fatal(err)
		}
		httpReq.Header.Set("Content-Type", "application/json; charset=utf-8")

		createFileReq, err := decodeCreateFileRequest(context.TODO(), httpReq)
		if err != nil {
			t.Error(string(bs))
			t.Fatalf("file %s had error against HTTP decode: %v", file.ACHFilepath, err)
		}

		repo := NewRepositoryInMemory(testTTLDuration, nil)
		s := NewService(repo)

		endpoint := createFileEndpoint(s, repo, nil) // nil logger

		resp, err := endpoint(context.TODO(), createFileReq)
		if err != nil {
			t.Fatalf("%s couldn't be created against our HTTP API: %v", file.ACHFilepath, err)
		}
		if resp == nil {
			t.Fatalf("resp == nil")
		}
		createFileResponse, ok := resp.(createFileResponse)
		if !ok {
			t.Fatalf("couldn't convert %#v to createFileResponse", resp)
		}
		if createFileResponse.ID == "" || createFileResponse.Err != nil {
			t.Fatalf("%s failed HTTP API creation: %v", file.ACHFilepath, createFileResponse.Err)
		}
	}
}

type testFile struct {
	SECCode     string
	ACHFilepath string
	Filename    string
}

func getTestFiles() []testFile {
	matches, err := filepath.Glob(filepath.Join("..", "test", "ach-*-read", "*.ach"))
	if err != nil {
		return nil
	}

	var testFiles []testFile
	for i := range matches {
		filename := filepath.Base(matches[i])

		testFiles = append(testFiles, testFile{
			SECCode:     strings.ToUpper(filename[:3]),
			ACHFilepath: matches[i],
			Filename:    strings.TrimSuffix(filename, ".ach"),
		})
	}

	return testFiles
}
