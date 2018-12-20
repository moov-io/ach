// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package test

import (
	"encoding/json"
	"fmt"
	"github.com/moov-io/ach"
	"github.com/moov-io/ach/server"
	"github.com/moov-io/base"
	"log"
	"os"
	"testing"
	"time"
)

// TestCreateJSONFromACHFiles creates JSON from existing ACH Files
func TestCreateJSONFromACHFiles(t *testing.T) {
	files := getFileNames()

	for _, file := range files {
		f, err := os.Open(file.AchFileName)
		if err != nil {
			log.Fatal(err)
		}
		r := ach.NewReader(f)
		achFile, err := r.Read()
		if err != nil {
			fmt.Printf("Issue reading file: %+v \n", err)
		}
		// ensure we have a validated file structure
		if achFile.Validate(); err != nil {
			fmt.Printf("Could not validate entire read file: %v", err)
		}
		// If you trust the file but it's formatting is off building will probably resolve the malformed file.
		if achFile.Create(); err != nil {
			fmt.Printf("Could not build file with read properties: %v", err)
		}

		f.Close()

		fmt.Printf("ACH File: %v \r\n", file.AchFileName)

		// ENR ACH Files does not have BatchHeader.EffectiveEntryDate, so setting this to Today +1 to be included
		// in the JSON File.  For this test after the ACH file is converted to JSON, the test validates the JSON by
		// calling ach.FileFromJSON(bs) and it fails with an empty date time.
		if file.SECCode == "ENR" {
			for _, batch := range achFile.Batches {
				batch.GetHeader().EffectiveEntryDate = base.NewTime(time.Now().AddDate(0, 0, 1))
			}

		}

		bs, err := json.MarshalIndent(achFile, "", " ")
		if err != nil {
			fmt.Println("error:", err)
		}

		fmt.Printf("JSON Output: %v \r\n", string(bs))

		fmt.Printf("Validating JSON byte stream %v created \r\n", file.jsonName)

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
}

type fileNames struct {
	SECCode     string
	AchFileName string
	jsonName    string
}

func getFileNames() []fileNames {
	fn := []fileNames{
		{"ACK", "../../test/ach-ack-read/ack-read.ach", "ack-credit"},
		{"ADV", "../../test/ach-adv-read/adv-read.ach", "adv-debitForCreditsOriginated"},
		{"ARC", "../../test/ach-arc-read/arc-debit.ach", "arc-debit"},
		{"ATX", "../../test/ach-atx-read/atx-read.ach", "atx-credit"},
		{"BOC", "../../test/ach-boc-read/boc-debit.ach", "boc-debit"},
		{"CCD", "../../test/ach-ccd-read/ccd-debit.ach", "ccd-debit"},
		{"CIE", "../../test/ach-cie-read/cie-credit.ach", "cie-credit"},
		{"CTX", "../../test/ach-ctx-read/ctx-debit.ach", "ctx-debit"},
		{"DNE", "../../test/ach-dne-read/dne-read.ach", "dne-debit"},
		{"ENR", "../../test/ach-enr-read/enr-read.ach", "enr-debit"},
		{"IAT", "../../test/ach-iat-read/iat-credit.ach", "iat-credit"},
		{"MTE", "../../test/ach-mte-read/mte-read.ach", "mte-debit"},
		{"POP", "../../test/ach-pop-read/pop-debit.ach", "pop-debit"},
		{"POS", "../../test/ach-pos-read/pos-debit.ach", "pos-debit"},
		{"PPD", "../../test/ach-ppd-read/ppd-credit.ach", "ppd-credit"},
		{"RCK", "../../test/ach-rck-read/rck-debit.ach", "rck-debit"},
		{"SHR", "../../test/ach-shr-read/shr-debit.ach", "shr-debit"},
		{"TEL", "../../test/ach-tel-read/tel-debit.ach", "tel-debit"},
		{"TRC", "../../test/ach-trc-read/trc-debit.ach", "trc-debit"},
		{"TRX", "../../test/ach-trx-read/trx-debit.ach", "trx-debit"},
		{"WEB", "../../test/ach-web-read/web-credit.ach", "web-credit"},
		{"XCK", "../../test/ach-xck-read/xck-debit.ach", "xck-debit"},
	}
	return fn
}
