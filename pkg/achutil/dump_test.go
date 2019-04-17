// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package achutil

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/ach"
)

func TestDumpFile(t *testing.T) {
	checkPlain := func(t *testing.T, path string) {
		t.Helper()
		fd, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}
		file, err := ach.NewReader(fd).Read()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Print(DumpFile(&file).String())
	}
	checkPlain(t, filepath.Join("..", "..", "test", "testdata", "20180716-IAT-A17-A18.ach"))
	checkPlain(t, filepath.Join("..", "..", "test", "testdata", "ppd-debit.ach"))
	checkPlain(t, filepath.Join("..", "..", "test", "testdata", "return-WEB.ach"))     // Addenda99
	checkPlain(t, filepath.Join("..", "..", "test", "ach-mte-read", "mte-read.ach"))   // Addenda02
	checkPlain(t, filepath.Join("..", "..", "test", "ach-web-read", "web-credit.ach")) // Addenda05

	// Some test files are in JSON, so let's check those too
	checkJSON := func(t *testing.T, path string) {
		t.Helper()
		bs, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		file, err := ach.FileFromJSON(bs)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Print(DumpFile(file).String())
	}
	checkJSON(t, filepath.Join("..", "..", "test", "testdata", "adv-valid.json")) // ADV
}
