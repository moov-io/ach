// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

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
	expectError("issue403-addenda02.json", "EntryDetail.addenda02")
	expectError("issue403-amount.json", "EntryDetail.amount")
}
