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
// TestCreateFile Tests Creating an ACH File From Json
func TestCreateFile(t *testing.T) {
	createFileError := func(path string, msg string) {
		_, err := os.Open(path)
		if err != nil {
			t.Fatal(err)
		}

		bs, err := ioutil.ReadFile(path)

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
	createFileError("ack-credit.json", "ACK credit zero dollar remittance failed to create")
	createFileError("adv-debitForCreditsOriginated.json", "ADV debit for credits originated failed to create")
	createFileError("arc-debit.json", "ARC debit failed to create")
	createFileError("atx-credit.json", "ATX credit failed to create")
	createFileError("boc-debit.json", "BOC debit failed to create")
	createFileError("ccd-debit.json", "CCD debit failed to create")
	createFileError("cie-credit.json", "CIE credit failed to create")
	createFileError("ctx-debit.json", "CTX debit failed to create")
	createFileError("dne-debit.json", "DNE debit failed to create")
	//createFileError("enr-debit.json", "ENR debit failed to create")
	//createFileError("iat-credit.json", "IAT credit failed to create")
	createFileError("mte-debit.json", "MTE debit failed to create")
	createFileError("ppd-debit.json", "PPD debit failed to create")
}
