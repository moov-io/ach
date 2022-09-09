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

package main

import (
	"log"
	"os"
	"time"

	"github.com/moov-io/ach"
)

func main() {
	fh := ach.NewFileHeader()
	fh.ImmediateDestination = "231380104"
	fh.ImmediateOrigin = "121042882"
	fh.FileCreationDate = time.Now().Format("060102")
	fh.ImmediateDestinationName = "Federal Reserve Bank"
	fh.ImmediateOriginName = "My Bank Name"

	file := ach.NewFile()
	file.SetHeader(fh)

	for i := 0; i < 4; i++ {
		bh := ach.NewBatchHeader()
		bh.ServiceClassCode = ach.AutomatedAccountingAdvices
		bh.CompanyName = "Company Name, Inc"
		bh.CompanyIdentification = fh.ImmediateOrigin
		bh.StandardEntryClassCode = ach.ADV
		bh.CompanyEntryDescription = "Accounting"
		bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1).Format("060102")
		bh.ODFIIdentification = "121042882"
		bh.OriginatorStatusCode = 0

		batch, _ := ach.NewBatch(bh)

		// Create Entry
		entrySeq := 0
		for i := 0; i < 3; i++ {
			entry := ach.NewADVEntryDetail()
			entry.TransactionCode = ach.CreditForDebitsOriginated
			entry.SetRDFI("231380104")
			entry.DFIAccountNumber = "744-5678-99"
			entry.Amount = 50000
			entry.AdviceRoutingNumber = "121042882"
			entry.FileIdentification = "11131"
			entry.ACHOperatorData = ""
			entry.IndividualName = "Name"
			entry.DiscretionaryData = ""
			entry.AddendaRecordIndicator = 0
			entry.ACHOperatorRoutingNumber = "01100001"
			entry.JulianDay = 50
			entry.SequenceNumber = entrySeq

			batch.AddADVEntry(entry)
		}
		if err := batch.Create(); err != nil {
			log.Fatalf("Unexpected error building batch: %s\n", err)
		}
		file.AddBatch(batch)
	}

	if err := file.Create(); err != nil {
		log.Fatalf("Unexpected error building file: %s\n", err)
	}

	// Write the file to stdout, any io.Writer can be used
	w := ach.NewWriter(os.Stdout)
	if err := w.Write(file); err != nil {
		log.Fatalf("Unexpected error writing file: %s\n", err)
	}
	w.Flush()
}
