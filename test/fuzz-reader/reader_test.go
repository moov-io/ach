// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package fuzzreader

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestCorpusSymlinks(t *testing.T) {
	// avoid symbolic link error on windows
	if runtime.GOOS == "windows" {
		t.Skip()
	}
	fds, err := ioutil.ReadDir("corpus")
	if err != nil {
		t.Fatal(err)
	}
	if len(fds) == 0 {
		t.Fatal("no file descriptors found in corpus/")
	}

	for i := range fds {
		if fds[i].Mode()&os.ModeSymlink != 0 {
			if path, err := os.Readlink(filepath.Join("corpus", fds[i].Name())); err != nil {
				t.Errorf("broken symlink: %v", err)
			} else {
				if _, err := os.Stat(filepath.Join("corpus", path)); err != nil {
					t.Errorf("broken symlink: %v", err)
				}
			}
		} else {
			t.Errorf("%s isn't a symlink, move outside corpus/ and symlink into directory", fds[i].Name())
		}
	}
}
