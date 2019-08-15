// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/moov-io/ach"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// open a file for reading. Any io.Reader Can be used
	fCredit, err := os.Open(filepath.Join("examples", "ach-iat-segmentFile-read", "segmentFile-iat-credit.ach"))
	if err != nil {
		log.Fatal(err)
	}
	r := ach.NewReader(fCredit)
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

	fmt.Printf("File Name: %s \n\n", fCredit.Name())
	fmt.Printf("Total Credit Amount: %v \n", achFile.Control.TotalCreditEntryDollarAmountInFile)
	fmt.Printf("SEC Code: %v \n\n", achFile.IATBatches[0].GetHeader().StandardEntryClassCode)
	fmt.Printf("Batch Total CRedit Amount: %v \n", achFile.IATBatches[0].GetControl().TotalCreditEntryDollarAmount)

	// open a file for reading. Any io.Reader Can be used
	fDebit, err := os.Open(filepath.Join("examples", "ach-iat-segmentFile-read", "segmentFile-iat-debit.ach"))
	if err != nil {
		log.Fatal(err)
	}
	rDebit := ach.NewReader(fDebit)
	achFileDebit, err := rDebit.Read()
	if err != nil {
		fmt.Printf("Issue reading file: %+v \n", err)
	}
	// ensure we have a validated file structure
	if achFileDebit.Validate(); err != nil {
		fmt.Printf("Could not validate entire read file: %v", err)
	}
	// If you trust the file but it's formatting is off building will probably resolve the malformed file.
	if achFileDebit.Create(); err != nil {
		fmt.Printf("Could not create file with read properties: %v", err)
	}

	fmt.Printf("File Name: %s \n\n", fDebit.Name())
	fmt.Printf("Total Debit Amount: %v \n", achFileDebit.Control.TotalDebitEntryDollarAmountInFile)
	fmt.Printf("SEC Code: %v \n", achFileDebit.IATBatches[0].GetHeader().StandardEntryClassCode)
	fmt.Printf("Batch Total Debit Amount: %v \n", achFileDebit.IATBatches[0].GetControl().TotalDebitEntryDollarAmount)
}
