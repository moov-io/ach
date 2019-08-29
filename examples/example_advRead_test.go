// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package examples

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/moov-io/ach"
)

func Example_advRead() {
	f, err := os.Open(filepath.Join("testdata", "adv-read.ach"))
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

	fmt.Printf("Credit Total Amount:%s", strconv.Itoa(achFile.ADVControl.TotalCreditEntryDollarAmountInFile)+"\n")
	fmt.Printf("Debit Total Amount:%s", strconv.Itoa(achFile.ADVControl.TotalDebitEntryDollarAmountInFile)+"\n")
	fmt.Printf("OriginatorStatusCode:%s", strconv.Itoa(achFile.Batches[0].GetHeader().OriginatorStatusCode)+"\n")
	fmt.Printf("Batch Credit Total Amount:%s", strconv.Itoa(achFile.Batches[0].GetADVControl().TotalCreditEntryDollarAmount)+"\n")
	fmt.Printf("Batch Debit Total Amount:%s", strconv.Itoa(achFile.Batches[0].GetADVControl().TotalDebitEntryDollarAmount)+"\n")
	fmt.Printf("SEC Code:%s", achFile.Batches[0].GetHeader().StandardEntryClassCode+"\n")
	fmt.Printf("Entry Amount:%s", strconv.Itoa(achFile.Batches[0].GetADVEntries()[0].Amount)+"\n")
	fmt.Printf("Sequence Number:%s", strconv.Itoa(achFile.Batches[0].GetADVEntries()[0].SequenceNumber)+"\n")
	fmt.Printf("EntryOne Amount:%s", strconv.Itoa(achFile.Batches[0].GetADVEntries()[1].Amount)+"\n")
	fmt.Printf("EntryOne Sequence Number:%s", strconv.Itoa(achFile.Batches[0].GetADVEntries()[1].SequenceNumber)+"\n")

	// Output:
	// Credit Total Amount:50000
	// Debit Total Amount:250000
	// OriginatorStatusCode:0
	// Batch Credit Total Amount:50000
	// Batch Debit Total Amount:250000
	// SEC Code:ADV
	// Entry Amount:50000
	// Sequence Number:1
	// EntryOne Amount:250000
	// EntryOne Sequence Number:2
}
