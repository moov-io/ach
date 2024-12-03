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
	"archive/tar"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/ach"
	"github.com/stretchr/testify/require"
)

func TestIssue863(t *testing.T) {
	fd, err := os.Open(filepath.Join("testdata", "storage.tar"))
	require.NoError(t, err)

	files, err := readFiles(t, fd)
	require.NoError(t, err)
	require.Len(t, files, 11)

	merged, err := ach.MergeFiles(files)
	require.NoError(t, err)
	require.Len(t, merged, 1)
	require.NoError(t, merged[0].Validate())

	final, err := merged[0].FlattenBatches()
	if err != nil {
		t.Fatal(err)
	}
	if err := final.Validate(); err != nil {
		t.Fatal(err)
	}
}

func readFiles(t *testing.T, r io.Reader) ([]*ach.File, error) {
	t.Helper()

	var out []*ach.File
	rdr := tar.NewReader(r)
	for {
		header, err := rdr.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatal(err)
		}

		if fd := header.FileInfo(); fd.IsDir() || strings.Contains(fd.Name(), ".json") {
			continue
		}

		f, err := ach.NewReader(rdr).Read()
		if err != nil {
			t.Fatal(err)
		}
		out = append(out, &f)
	}
	return out, nil
}
