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
	"os"
	"strings"
	"testing"

	"github.com/moov-io/ach"
)

// Testissue403 attempts to parse a few JSON files as ach.Batch objects, but
// each JSON object is malformed and we're matching on the error returned.
//
// See: https://github.com/moov-io/ach/issues/403
func TestIssue403(t *testing.T) {
	expectError := func(path string, msg string) {
		t.Helper()

		fd, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}
		var batch ach.Batch
		if err := json.NewDecoder(fd).Decode(&batch); err != nil {
			if !strings.Contains(err.Error(), msg) {
				t.Errorf("(file: %s) %q doesn't contain expected %q", path, err.Error(), msg)
			}
		} else {
			t.Error("expected error, but got none")
		}
	}

	// test cases
	expectError("issue403-addenda02.json", "addenda02")
	expectError("issue403-amount.json", "amount")
}
