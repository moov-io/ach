package main

import (
	"log"
	"os"
	"time"
	
	"github.com/moov-io/ach"
)

func main() {
	// Example transfer to write an ACH PPD file to send/credit a external institutions account
	// Important: All financial institutions are different and will require registration and exact field values.

	// Set originator bank ODFI and destination Operator for the financial institution
	// this is the funding/receiving source of the transfer
	fh := ach.NewFileHeader()
	fh.ImmediateDestination = "231380104" // Routing Number of the ACH Operator or receiving point to which the file is being sent
	fh.ImmediateOrigin = "121042882"      // Routing Number of the ACH Operator or sending point that is sending the file
	fh.FileCreationDate = time.Now()      // Today's Date
	fh.ImmediateDestinationName = "Bank"
	fh.ImmediateOriginName = "My Bank Name"

	// BatchHeader identifies the originating entity and the type of transactions contained in the batch
	bh := ach.NewIATBatchHeader()
	bh.ServiceClassCode = 220
	bh.ForeignExchangeIndicator = "FF"
	bh.ForeignExchangeReferenceIndicator = 3
	bh.ISODestinationCountryCode = "US"
	bh.OriginatorIdentification = "123456789"
	bh.StandardEntryClassCode = "IAT"
	bh.CompanyEntryDescription = "TRADEPAYMT"
	bh.ISOOriginatingCurrencyCode = "CAD"
	bh.ISODestinationCurrencyCode = "USD"
	bh.ODFIIdentification = "23138010"

	// Identifies the receivers account information
	// can be multiple entry's per batch
	entry := ach.NewIATEntryDetail()
	entry.TransactionCode = 27
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
	addenda13.ODFIIdentification = "121042882"
	addenda13.ODFIBranchCountryCode = "US"
	addenda13.EntryDetailSequenceNumber = 00000001
	entry.Addenda13 = addenda13

	addenda14 := ach.NewAddenda14()
	addenda14.RDFIName = "Citadel Bank"
	addenda14.RDFIIDNumberQualifier = "01"
	addenda14.RDFIIdentification = "231380104"
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
	entry.AddIATAddenda(addenda17)

	addenda18 := ach.NewAddenda18()
	addenda18.ForeignCorrespondentBankName = "Bank of France"
	addenda18.ForeignCorrespondentBankIDNumberQualifier = "01"
	addenda18.ForeignCorrespondentBankIDNumber = "456456456987987"
	addenda18.ForeignCorrespondentBankBranchCountryCode = "FR"
	addenda18.SequenceNumber = 3
	addenda18.EntryDetailSequenceNumber = 0000001
	entry.AddIATAddenda(addenda18)

	// build the batch
	batch := ach.NewIATBatch(bh)
	batch.AddEntry(entry)
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

	// write the file to std out. Anything io.Writer
	w := ach.NewWriter(os.Stdout)
	if err := w.Write(file); err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}
	w.Flush()
}
