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

package ach

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	filenames := []string{
		"ppd-debit.ach",
		"ppd-valid-debit.json",
		"ppd-valid.json",
		"return-WEB.ach",
		"web-debit.ach",
	}

	dir := copyFilesToTempDir(t, filenames)
	defer os.RemoveAll(dir)

	files, err := ReadDir(dir)
	require.NoError(t, err)
	require.Len(t, files, 5)

	// Make a directory which is skipped
	err = os.Mkdir(filepath.Join(dir, "other"), 0777)
	require.NoError(t, err)

	files, err = ReadDir(dir)
	require.NoError(t, err)
	require.Len(t, files, 5)
}

func TestReadDirErr(t *testing.T) {
	filenames := []string{
		"ppd-debit.ach",
		"ppd-valid-debit.json",
	}

	dir := copyFilesToTempDir(t, filenames)
	defer os.RemoveAll(dir)

	// zzz- is a prefix as os.ReadDir seems to return file descriptors ordered alphabetically by filename
	err := os.WriteFile(filepath.Join(dir, "zzz-bad.ach"), []byte("bad data"), 0600)
	require.NoError(t, err)

	files, err := ReadDir(dir)
	require.Error(t, err)
	require.Len(t, files, 2)

	files, err = ReadDir("/not/exist/")
	if n := len(files); n != 0 || err == nil {
		t.Errorf("got %d files error=%v", n, err)
	}
}

func TestReadDirSymlinkErr(t *testing.T) {
	dir, err := os.MkdirTemp("", "readdir-symlink")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	// write an invalid symlink
	err = os.Symlink(filepath.Join("missing", "directory"), filepath.Join(dir, "foo.ach"))
	require.NoError(t, err)

	files, err := ReadDir(dir)
	require.Error(t, err)
	require.Len(t, files, 0)
}

func copyFilesToTempDir(t *testing.T, filenames []string) string {
	t.Helper()

	dir, err := os.MkdirTemp("", "ach-readdir")
	require.NoError(t, err)

	for i := range filenames {
		in, err := os.Open(filepath.Join("test", "testdata", filenames[i]))
		if err != nil {
			t.Fatalf("in: filename=%s error=%v", filenames[i], err)
		}
		out, err := os.Create(filepath.Join(dir, filenames[i]))
		if err != nil {
			t.Fatalf("out: filename=%s error=%v", filenames[i], err)
		}
		_, err = io.Copy(out, in)

		in.Close()
		out.Close()

		require.NoError(t, err)
	}

	return dir
}
