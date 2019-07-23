// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/moov-io/ach"
)

func main() {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("examples", "ach-ppd-read-mixedDebitCredit", "ppd-mixedDebitCredit.ach"))
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

	fmt.Printf("File Name: %s \n", f.Name())
	fmt.Printf("Total Debit Amount: %v \n", achFile.Control.TotalDebitEntryDollarAmountInFile)
	fmt.Printf("Total Credit Amount: %v \n", achFile.Control.TotalCreditEntryDollarAmountInFile)
	fmt.Printf("File Header: %v \n", achFile.Header.String())
	fmt.Printf("Batch Header: %v \n", achFile.Batches[0].GetHeader().String())
	fmt.Printf("Entry Detail 1: %v \n", achFile.Batches[0].GetEntries()[0].String())
	fmt.Printf("Entry Detail 2: %v \n", achFile.Batches[0].GetEntries()[1].String())
	fmt.Printf("Entry Detail 3: %v \n", achFile.Batches[0].GetEntries()[2].String())
	fmt.Printf("Batch Control: %v \n", achFile.Batches[0].GetControl().String())
	fmt.Printf("File Header: %v \n\n", achFile.Control.String())
}
