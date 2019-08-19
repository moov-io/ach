// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package examples

import (
	"fmt"
	"github.com/moov-io/ach"
	"log"
)

// Example_ackWrite writes an ACK File
func Example_ackWrite() {
	fh := mockFileHeader()
	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.CreditsOnly
	bh.CompanyName = "Name on Account"
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.ACK
	bh.CompanyEntryDescription = "Vndr Pay"
	// need EffectiveEntryDate to be fixed so it can match output
	bh.EffectiveEntryDate = "190816"
	bh.ODFIIdentification = "23138010"

	entry := ach.NewEntryDetail()
	entry.TransactionCode = ach.CheckingZeroDollarRemittanceCredit
	entry.SetRDFI("031300012")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 0
	entry.SetOriginalTraceNumber("031300010000001")
	entry.SetReceivingCompany("Best. #1")
	entry.SetTraceNumber(bh.ODFIIdentification, 1)

	entryOne := ach.NewEntryDetail() // Fee Entry
	entryOne.TransactionCode = ach.CheckingZeroDollarRemittanceCredit
	entryOne.SetRDFI("031300012")
	entryOne.DFIAccountNumber = "744-5678-99"
	entryOne.Amount = 0
	entryOne.SetOriginalTraceNumber("031300010000002")
	entryOne.SetReceivingCompany("Best. #1")
	entryOne.SetTraceNumber(bh.ODFIIdentification, 2)

	// build the batch
	batch := ach.NewBatchACK(bh)
	batch.AddEntry(entry)
	batch.AddEntry(entryOne)
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
	fmt.Printf("%s", file.Batches[0].GetEntries()[1].String()+"\n")
	fmt.Printf("%s", file.Batches[0].GetControl().String()+"\n")
	fmt.Printf("%s", file.Control.String()+"\n")

	// Output:
	// 101 03130001202313801041908161055A094101Federal Reserve Bank   My Bank Name           12345678
	// 5220Name on Account                     231380104 ACKVndr Pay        190816   1231380100000001
	// 624031300012744-5678-99      0000000000031300010000001Best. #1                0231380100000001
	// 624031300012744-5678-99      0000000000031300010000002Best. #1                0231380100000002
	// 82200000020006260002000000000000000000000000231380104                          231380100000001
	// 9000001000001000000020006260002000000000000000000000000
}
