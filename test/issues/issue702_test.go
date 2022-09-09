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
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/ach"
)

func TestIssue702(t *testing.T) {
	// A vendor gave issue702.ach to a customer of ours as a return but didn't properly
	// format the line lengths and included a non-routing number in the ImmediateDestination
	fd, err := os.Open(filepath.Join("testdata", "issue702.ach"))
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()

	r := ach.NewReader(fd)
	r.SetValidation(&ach.ValidateOpts{
		BypassDestinationValidation: true,
	})

	file, err := r.Read()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if err := file.Validate(); err != nil {
		t.Error(err)
	}

	if file.Header.ImmediateDestination != "YYYYYYYYY" {
		t.Errorf("file.Header.ImmediateDestination=%s", file.Header.ImmediateDestination)
	}
	if n := len(file.Batches); n != 3 {
		t.Errorf("got %d batches", n)
	}
	if n := len(file.ReturnEntries); n != 3 {
		t.Errorf("got %d NOC's", n)
	}
}

func TestIssue702_1(t *testing.T) {
	// This file was returned as a receipt from uploading an ACH file, but this file is
	// pretty useless as it contains zero EntryDetail's.
	fd, err := os.Open(filepath.Join("testdata", "issue702-1.ach"))
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()

	r := ach.NewReader(fd)
	r.SetValidation(&ach.ValidateOpts{
		BypassDestinationValidation: true,
	})

	file, err := r.Read()
	if err != nil {
		if !strings.Contains(err.Error(), "BatchCount 000000 is a mandatory field") {
			t.Error(err)
		}
	}
	if err := file.Validate(); err != nil {
		if !strings.Contains(err.Error(), "BatchCount 000000 is a mandatory field") {
			t.Error(err)
		}
	}

	if file.Header.ImmediateOrigin != "182327390" {
		t.Errorf("ImmediateOrigin=%s", file.Header.ImmediateOrigin)
	}
	if file.Header.ImmediateDestination != "10006XXXX" {
		t.Errorf("ImmediateDestination=%s", file.Header.ImmediateDestination)
	}
	if file.Header.ImmediateDestinationName != "PIMRET825324" {
		t.Errorf("ImmediateDestinationName=%s", file.Header.ImmediateDestinationName)
	}
}
