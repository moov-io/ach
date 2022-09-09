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
	// Example transfer to write an ACH IAT file to debit a external institution's account
	// Important: All financial institutions are different and will require registration and exact field values.

	// Set originator bank ODFI and destination Operator for the financial institution
	// this is the funding/receiving source of the transfer
	fh := ach.NewFileHeader()
	fh.ImmediateDestination = "121042882"             // Routing Number of the ACH Operator or receiving point to which the file is being sent
	fh.ImmediateOrigin = "231380104"                  // Routing Number of the ACH Operator or sending point that is sending the file
	fh.FileCreationDate = time.Now().Format("060102") // Today's Date
	fh.ImmediateDestinationName = "Bank"
	fh.ImmediateOriginName = "My Bank Name"

	file := ach.NewFile()
	file.SetHeader(fh)

	for i := 0; i < 4; i++ {
		bh := ach.NewIATBatchHeader()
		bh.ServiceClassCode = ach.MixedDebitsAndCredits
		bh.ForeignExchangeIndicator = "FF"
		bh.ForeignExchangeReferenceIndicator = 3
		bh.ISODestinationCountryCode = "US"
		bh.OriginatorIdentification = "123456789"
		bh.StandardEntryClassCode = ach.IAT
		bh.CompanyEntryDescription = "TRADEPAYMT"
		bh.ISOOriginatingCurrencyCode = "CAD"
		bh.ISODestinationCurrencyCode = "USD"
		bh.ODFIIdentification = "23138010"
		bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1).Format("060102") // YYMMDD

		batch := ach.NewIATBatch(bh)

		// Create Entry
		entrySeq := 0
		for i := 0; i < 3; i++ {
			entrySeq = entrySeq + 1
			// Identifies the receiver's account information
			// can be multiple entries per batch
			entry := ach.NewIATEntryDetail()
			entry.TransactionCode = ach.CheckingDebit
			entry.SetRDFI("121042882")
			entry.AddendaRecords = 007
			entry.DFIAccountNumber = "123456789"
			entry.Amount = 100000 // 1000.00
			entry.SetTraceNumber("23138010", entrySeq)
			entry.Category = ach.CategoryForward

			addenda10 := ach.NewAddenda10()
			addenda10.TransactionTypeCode = "ANN"
			addenda10.ForeignPaymentAmount = 100000
			addenda10.ForeignTraceNumber = "928383-23938"
			addenda10.Name = "BEK Enterprises"
			entry.Addenda10 = addenda10

			addenda11 := ach.NewAddenda11()
			addenda11.OriginatorName = "BEK Solutions"
			addenda11.OriginatorStreetAddress = "15 West Place Street"
			entry.Addenda11 = addenda11

			addenda12 := ach.NewAddenda12()
			addenda12.OriginatorCityStateProvince = "JacobsTown*PA\\"
			addenda12.OriginatorCountryPostalCode = "US*19305\\"
			entry.Addenda12 = addenda12

			addenda13 := ach.NewAddenda13()
			addenda13.ODFIName = "Wells Fargo"
			addenda13.ODFIIDNumberQualifier = "01"
			addenda13.ODFIIdentification = "231380104"
			addenda13.ODFIBranchCountryCode = "US"
			entry.Addenda13 = addenda13

			addenda14 := ach.NewAddenda14()
			addenda14.RDFIName = "Citadel Bank"
			addenda14.RDFIIDNumberQualifier = "01"
			addenda14.RDFIIdentification = "121042882"
			addenda14.RDFIBranchCountryCode = "CA"
			entry.Addenda14 = addenda14

			addenda15 := ach.NewAddenda15()
			addenda15.ReceiverIDNumber = "987465493213987"
			addenda15.ReceiverStreetAddress = "2121 Front Street"
			entry.Addenda15 = addenda15

			addenda16 := ach.NewAddenda16()
			addenda16.ReceiverCityStateProvince = "LetterTown*AB\\"
			addenda16.ReceiverCountryPostalCode = "CA*80014\\"
			entry.Addenda16 = addenda16

			addenda17 := ach.NewAddenda17()
			addenda17.PaymentRelatedInformation = "This is an international payment"
			addenda17.SequenceNumber = 1
			entry.AddAddenda17(addenda17)

			addenda18 := ach.NewAddenda18()
			addenda18.ForeignCorrespondentBankName = "Bank of France"
			addenda18.ForeignCorrespondentBankIDNumberQualifier = "01"
			addenda18.ForeignCorrespondentBankIDNumber = "456456456987987"
			addenda18.ForeignCorrespondentBankBranchCountryCode = "FR"
			addenda18.SequenceNumber = 3
			entry.AddAddenda18(addenda18)

			batch.AddEntry(entry)
		}
		if err := batch.Create(); err != nil {
			log.Fatalf("Unexpected error building batch: %s\n", err)
		}

		file.AddIATBatch(batch)
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
