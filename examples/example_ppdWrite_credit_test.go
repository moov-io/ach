// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package examples

import (
	"fmt"
	"github.com/moov-io/ach"
	"log"
)

func Example_ppdWriteCredit() {
	fh := mockFileHeader()

	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.CreditsOnly
	bh.CompanyName = "Name on Account"
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.PPD
	bh.CompanyEntryDescription = "REG.SALARY"
	// need EffectiveEntryDate to be fixed so it can match output
	bh.EffectiveEntryDate = "190816"
	bh.ODFIIdentification = "121042882"

	entry := ach.NewEntryDetail()
	entry.TransactionCode = ach.CheckingCredit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "987654321"
	entry.Amount = 100000000
	entry.SetTraceNumber(bh.ODFIIdentification, 2)
	entry.IndividualName = "Credit Account 1"

	// build the batch
	batch := ach.NewBatchPPD(bh)
	batch.AddEntry(entry)

	if err := batch.Create(); err != nil {
		log.Fatalf("Unexpected error building batch: %s\n", err)
	}

	// build the file
	file := ach.NewFile()
	file.SetHeader(fh)
	file.AddBatch(batch)
	if err := file.Create(); err != nil {
		log.Fatalf("Unexpected error building file: %s\n", err)
	}

	fmt.Printf("%s", file.Header.String()+"\n")
	fmt.Printf("%s", file.Batches[0].GetHeader().String()+"\n")
	fmt.Printf("%s", file.Batches[0].GetEntries()[0].String()+"\n")
	fmt.Printf("%s", file.Batches[0].GetControl().String()+"\n")
	fmt.Printf("%s", file.Control.String()+"\n")

	// Output:
	// 101 03130001202313801041908161055A094101Federal Reserve Bank   My Bank Name           12345678
	// 5220Name on Account                     231380104 PPDREG.SALARY      190816   1121042880000001
	// 622231380104987654321        0100000000               Credit Account 1        0121042880000002
	// 82200000010023138010000000000000000100000000231380104                          121042880000001
	// 9000001000001000000010023138010000000000000000100000000
}
