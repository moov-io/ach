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
	"path/filepath"
	"testing"

	"github.com/moov-io/ach"

	"github.com/stretchr/testify/require"
)

func TestIssue927(t *testing.T) {
	before, err := ach.ReadDir(filepath.Join("testdata", "issue927"))
	require.NoError(t, err)

	after, err := ach.MergeFiles(before)
	require.NoError(t, err)

	if len(after) != 1 {
		t.Fatalf("merged %d files into %d files", len(before), len(after))
	}

	// batches are flattened
	require.Len(t, after[0].Batches, 2)

	// Verify all the entries are present
	var entryCount int
	for i := range after {
		for j := range after[i].Batches {
			entryCount += len(after[i].Batches[j].GetEntries())
		}
	}
	require.Equal(t, 18, entryCount)
}
