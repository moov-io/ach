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

	// BatchHeader identifies the originating entity and the type of transactions contained in the batch
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

	// Identifies the receiver's account information
	// can be multiple entries per batch
	entry := ach.NewIATEntryDetail()
	entry.TransactionCode = ach.CheckingDebit
	entry.SetRDFI("121042882")
	entry.AddendaRecords = 007
	entry.DFIAccountNumber = "123456789"
	entry.Amount = 100000 // 1000.00
	entry.SetTraceNumber("23138010", 1)
	entry.Category = ach.CategoryForward

	//addenda

	addenda10 := ach.NewAddenda10()
	addenda10.TransactionTypeCode = "ANN"
	addenda10.ForeignPaymentAmount = 100000
	addenda10.ForeignTraceNumber = "928383-23938"
	addenda10.Name = "BEK Enterprises"
	addenda10.EntryDetailSequenceNumber = 00000001
	entry.Addenda10 = addenda10

	addenda11 := ach.NewAddenda11()
	addenda11.OriginatorName = "BEK Solutions"
	addenda11.OriginatorStreetAddress = "15 West Place Street"
	addenda11.EntryDetailSequenceNumber = 00000001
	entry.Addenda11 = addenda11

	addenda12 := ach.NewAddenda12()
	addenda12.OriginatorCityStateProvince = "JacobsTown*PA\\"
	addenda12.OriginatorCountryPostalCode = "US*19305\\"
	addenda12.EntryDetailSequenceNumber = 00000001
	entry.Addenda12 = addenda12

	addenda13 := ach.NewAddenda13()
	addenda13.ODFIName = "Wells Fargo"
	addenda13.ODFIIDNumberQualifier = "01"
	addenda13.ODFIIdentification = "231380104"
	addenda13.ODFIBranchCountryCode = "US"
	addenda13.EntryDetailSequenceNumber = 00000001
	entry.Addenda13 = addenda13

	addenda14 := ach.NewAddenda14()
	addenda14.RDFIName = "Citadel Bank"
	addenda14.RDFIIDNumberQualifier = "01"
	addenda14.RDFIIdentification = "121042882"
	addenda14.RDFIBranchCountryCode = "CA"
	addenda14.EntryDetailSequenceNumber = 00000001
	entry.Addenda14 = addenda14

	addenda15 := ach.NewAddenda15()
	addenda15.ReceiverIDNumber = "987465493213987"
	addenda15.ReceiverStreetAddress = "2121 Front Street"
	addenda15.EntryDetailSequenceNumber = 00000001
	entry.Addenda15 = addenda15

	addenda16 := ach.NewAddenda16()
	addenda16.ReceiverCityStateProvince = "LetterTown*AB\\"
	addenda16.ReceiverCountryPostalCode = "CA*80014\\"
	addenda16.EntryDetailSequenceNumber = 00000001
	entry.Addenda16 = addenda16

	addenda17 := ach.NewAddenda17()
	addenda17.PaymentRelatedInformation = "This is an international payment"
	addenda17.SequenceNumber = 1
	addenda17.EntryDetailSequenceNumber = 0000001
	entry.AddAddenda17(addenda17)

	addenda18 := ach.NewAddenda18()
	addenda18.ForeignCorrespondentBankName = "Bank of France"
	addenda18.ForeignCorrespondentBankIDNumberQualifier = "01"
	addenda18.ForeignCorrespondentBankIDNumber = "456456456987987"
	addenda18.ForeignCorrespondentBankBranchCountryCode = "FR"
	addenda18.SequenceNumber = 3
	addenda18.EntryDetailSequenceNumber = 0000001
	entry.AddAddenda18(addenda18)

	// Identifies the receiver's account information
	// can be multiple entries per batch
	entryTwo := ach.NewIATEntryDetail()
	entryTwo.TransactionCode = ach.CheckingCredit
	entryTwo.SetRDFI("121042882")
	entryTwo.AddendaRecords = 007
	entryTwo.DFIAccountNumber = "123456789"
	entryTwo.Amount = 100000 // 1000.00
	entryTwo.SetTraceNumber("23138010", 2)
	entryTwo.Category = ach.CategoryForward

	//addenda

	addenda10Two := ach.NewAddenda10()
	addenda10Two.TransactionTypeCode = "ANN"
	addenda10Two.ForeignPaymentAmount = 100000
	addenda10Two.ForeignTraceNumber = "928383-23938"
	addenda10Two.Name = "ADCAF Enterprises"
	addenda10Two.EntryDetailSequenceNumber = 00000002
	entryTwo.Addenda10 = addenda10Two

	addenda11Two := ach.NewAddenda11()
	addenda11Two.OriginatorName = "ADCAF Solutions"
	addenda11Two.OriginatorStreetAddress = "15 West Place Street"
	addenda11Two.EntryDetailSequenceNumber = 00000002
	entryTwo.Addenda11 = addenda11Two

	addenda12Two := ach.NewAddenda12()
	addenda12Two.OriginatorCityStateProvince = "JacobsTown*PA\\"
	addenda12Two.OriginatorCountryPostalCode = "US*19305\\"
	addenda12Two.EntryDetailSequenceNumber = 00000002
	entryTwo.Addenda12 = addenda12Two

	addenda13Two := ach.NewAddenda13()
	addenda13Two.ODFIName = "Wells Fargo"
	addenda13Two.ODFIIDNumberQualifier = "01"
	addenda13Two.ODFIIdentification = "231380104"
	addenda13Two.ODFIBranchCountryCode = "US"
	addenda13Two.EntryDetailSequenceNumber = 00000002
	entryTwo.Addenda13 = addenda13Two

	addenda14Two := ach.NewAddenda14()
	addenda14Two.RDFIName = "Citadel Bank"
	addenda14Two.RDFIIDNumberQualifier = "01"
	addenda14Two.RDFIIdentification = "121042882"
	addenda14Two.RDFIBranchCountryCode = "CA"
	addenda14Two.EntryDetailSequenceNumber = 00000002
	entryTwo.Addenda14 = addenda14Two

	addenda15Two := ach.NewAddenda15()
	addenda15Two.ReceiverIDNumber = "987465493213987"
	addenda15Two.ReceiverStreetAddress = "18 Fifth Street"
	addenda15Two.EntryDetailSequenceNumber = 00000002
	entryTwo.Addenda15 = addenda15Two

	addenda16Two := ach.NewAddenda16()
	addenda16Two.ReceiverCityStateProvince = "LetterTown*AB\\"
	addenda16Two.ReceiverCountryPostalCode = "CA*80014\\"
	addenda16Two.EntryDetailSequenceNumber = 00000002
	entryTwo.Addenda16 = addenda16Two

	addenda17Two := ach.NewAddenda17()
	addenda17Two.PaymentRelatedInformation = "This is an international payment"
	addenda17Two.SequenceNumber = 1
	addenda17Two.EntryDetailSequenceNumber = 0000002
	entry.AddAddenda17(addenda17)

	addenda18Two := ach.NewAddenda18()
	addenda18Two.ForeignCorrespondentBankName = "Bank of France"
	addenda18Two.ForeignCorrespondentBankIDNumberQualifier = "01"
	addenda18Two.ForeignCorrespondentBankIDNumber = "456456456987987"
	addenda18Two.ForeignCorrespondentBankBranchCountryCode = "FR"
	addenda18Two.SequenceNumber = 3
	addenda18Two.EntryDetailSequenceNumber = 0000002
	entryTwo.AddAddenda18(addenda18Two)

	// build the batch
	batch := ach.NewIATBatch(bh)
	batch.AddEntry(entry)
	batch.AddEntry(entryTwo)
	if err := batch.Create(); err != nil {
		log.Fatalf("Unexpected error building batch: %s\n", err)
	}

	// build the file
	file := ach.NewFile()
	file.SetHeader(fh)
	file.AddIATBatch(batch)
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
