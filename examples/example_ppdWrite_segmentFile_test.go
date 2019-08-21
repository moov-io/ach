// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package examples

import (
	"fmt"
	"github.com/moov-io/ach"
	"log"
	"os"
	"path/filepath"
)

func Example_ppdWriteSegmentFile() {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("testdata", "ppd-mixedDebitCredit.ach"))
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

	sfc := ach.NewSegmentFileConfiguration()
	creditFile, debitFile, err := achFile.SegmentFile(sfc)

	if err != nil {
		fmt.Printf("Could not segment the file: %v", err)
	}

	fmt.Printf("%s", creditFile.Batches[0].GetEntries()[0].String()+"\n")
	fmt.Printf("%s", debitFile.Batches[0].GetEntries()[0].String()+"\n")

	// Output:
	// 622231380104987654321        0100000000               Credit Account 1        0121042880000001
	// 627231380104123456789        0200000000               Debit Account           0121042880000001

}
