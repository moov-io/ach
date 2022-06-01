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
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/moov-io/ach"

	"github.com/stretchr/testify/require"
)

func TestIssue1024__Read(t *testing.T) {
	bs, err := ioutil.ReadFile(filepath.Join("testdata", "issue1024.json"))
	require.NoError(t, err)

	file, err := ach.FileFromJSON(bs)
	require.NoError(t, err)
	require.Len(t, file.Batches, 1)

	entries := file.Batches[0].GetEntries()
	require.Len(t, entries, 1)

	// I expected the traceNumber field in addenda99 to be automatically
	// populated equal to the traceNumber of the entry
	ed := entries[0]
	require.Equal(t, "084106760000001", ed.TraceNumber)
	require.NotNil(t, ed.Addenda99)
	require.Equal(t, "084106760000001", ed.Addenda99.TraceNumber)
}
