// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package examples

import (
	"fmt"
	"log"

	"github.com/moov-io/ach"
)

func Example_mteWriteDebit() {
	fh := mockFileHeader()

	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.DebitsOnly
	bh.CompanyName = "Merchant with ATM" // Merchant with the ATM
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.MTE
	bh.CompanyEntryDescription = "CASH WITHDRAW"
	bh.EffectiveEntryDate = "190816" // need EffectiveEntryDate to be fixed so it can match output
	bh.ODFIIdentification = "23138010"

	entry := ach.NewEntryDetail()
	entry.TransactionCode = ach.CheckingDebit
	entry.SetRDFI("031300012")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 10000
	entry.SetOriginalTraceNumber("031300010000001")
	entry.SetReceivingCompany("JANE DOE")
	entry.SetTraceNumber(bh.ODFIIdentification, 1)

	addenda02 := ach.NewAddenda02()
	addenda02.TerminalIdentificationCode = "200509"
	addenda02.TerminalLocation = "321 East Market Street"
	addenda02.TerminalCity = "ANYTOWN"
	addenda02.TerminalState = "VA"

	addenda02.TransactionSerialNumber = "123456"
	addenda02.TransactionDate = "1224"
	addenda02.TraceNumber = entry.TraceNumber
	entry.Addenda02 = addenda02
	entry.AddendaRecordIndicator = 1

	// build the batch
	batch := ach.NewBatchMTE(bh)
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

	fmt.Println(file.Header.String())
	fmt.Println(file.Batches[0].GetHeader().String())
	fmt.Println(file.Batches[0].GetEntries()[0].String())
	fmt.Println(file.Batches[0].GetControl().String())
	fmt.Println(file.Control.String())

	// Output:
	// 101 031300012 2313801041908161055A094101Federal Reserve Bank   My Bank Name           12345678
	// 5225Merchant with AT                    231380104 MTECASH WITHD      190816   1231380100000001
	// 627031300012744-5678-99      0000010000031300010000001JANE DOE                1231380100000001
	// 82250000020003130001000000010000000000000000231380104                          231380100000001
	// 9000001000001000000020003130001000000010000000000000000
}
