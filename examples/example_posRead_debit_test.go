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
	"strconv"
)

func Example_posReadDebit() {
	f, err := os.Open(filepath.Join("testdata", "pos-debit.ach"))
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
	if err := achFile.Create(); err != nil {
		fmt.Printf("Could not create file with read properties: %v", err)
	}

	fmt.Printf("Total Amount Debit: %s", strconv.Itoa(achFile.Control.TotalDebitEntryDollarAmountInFile)+"\n")
	fmt.Printf("SEC Code: %s", achFile.Batches[0].GetHeader().StandardEntryClassCode+"\n")
	fmt.Printf("POS Card Transaction Type: %s", achFile.Batches[0].GetEntries()[0].DiscretionaryData+"\n")
	fmt.Printf("POS Trace Number: %s", achFile.Batches[0].GetEntries()[0].TraceNumber+"\n")

	// Output:
	// Total Amount Debit: 100000000
	// SEC Code: POS
	// POS Card Transaction Type: 01
	// POS Trace Number: 121042880000001
}
