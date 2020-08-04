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
	"testing"

	"github.com/moov-io/ach"
)

func TestIssue751(t *testing.T) {
	fd, err := os.Open(filepath.Join("testdata", "issue751.ach"))
	if err != nil {
		t.Fatal(err)
	}
	file, err := ach.NewReader(fd).Read()
	if err == nil {
		t.Error("expected error")
	}
	if len(file.Batches) != 1 || len(file.IATBatches) != 0 {
		t.Errorf("got %d Batches and %d IAT Batches", len(file.Batches), len(file.IATBatches))
	}

	entries := file.Batches[0].GetEntries()
	if len(entries) != 2 {
		t.Fatalf("got %d Entries", len(entries))
	}

	if acct := entries[0].DFIAccountNumber; acct != "82111184         " {
		t.Errorf("got %q", acct)
	}
	if acct := entries[1].DFIAccountNumber; acct != "0110             " {
		t.Errorf("got %q", acct)
	}
}
