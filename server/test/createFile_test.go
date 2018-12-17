// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package test

import (
	"fmt"
	"github.com/moov-io/ach/server"
	"io/ioutil"
	"os"
	"testing"

	"github.com/moov-io/ach"
)

//
// TestCreatFile Tests Creating an ACH File From Json
func TestCreatFile(t *testing.T) {
	createFileError := func(path string, msg string) {
		_, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}

		bs, err := ioutil.ReadFile("ppd-debit.json")
		if err != nil {
			t.Fatal(err)
		}

		// Create ACH File from JSON
		file, err := ach.FileFromJSON(bs)
		if err != nil {
			t.Fatal(err.Error())
		}

		// Validate ACH File
		if file.Validate(); err != nil {
			fmt.Printf("Could not validate file: %v", err)
		}

		repository := server.NewRepositoryInMemory()
		s := server.NewService(repository)

		// Create and store ACH File in repository
		fileID, err := s.CreateFile(&file.Header)
		if err != nil {
			t.Fatal(err.Error())
		}

		// Get the stored ACH File from repository
		getFile, err := s.GetFile(fileID)
		if err == server.ErrNotFound {
			t.Errorf("expected %s received %s w/ error %s", "ErrNotFound", getFile.ID, err)
		}

	}

	// test cases
	createFileError("ppd-debit.json", "PPD Debit failed to create")

}
