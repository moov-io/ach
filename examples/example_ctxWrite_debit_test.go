package examples

import (
	"fmt"
	"github.com/moov-io/ach"
	"log"
)

// Example_ctxWriteDebit writes a CTX debit file
func Example_ctxWriteDebit() {

	fh := mockFileHeader()

	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.DebitsOnly
	bh.CompanyName = "Name on Account"
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.CTX
	bh.CompanyEntryDescription = "ACH CTX"
	bh.EffectiveEntryDate = "190816"
	bh.ODFIIdentification = "121042882"

	entry := ach.NewEntryDetail()

	entry.TransactionCode = ach.CheckingDebit
	entry.SetRDFI("231380104")
	entry.DFIAccountNumber = "12345678"
	entry.Amount = 100000000
	entry.IdentificationNumber = "45689033"
	entry.SetCATXAddendaRecords(2)
	entry.SetCATXReceivingCompany("Receiver Company")
	entry.SetTraceNumber(bh.ODFIIdentification, 1)
	entry.DiscretionaryData = "01"

	addenda1 := ach.NewAddenda05()
	addenda1.PaymentRelatedInformation = "Debit First Account"
	addenda1.SequenceNumber = 1
	addenda1.EntryDetailSequenceNumber = 0000001

	addenda2 := ach.NewAddenda05()
	addenda2.PaymentRelatedInformation = "Debit Second Account"
	addenda2.SequenceNumber = 2
	addenda2.EntryDetailSequenceNumber = 0000001

	// build the batch
	batch := ach.NewBatchCTX(bh)
	batch.AddEntry(entry)
	batch.Entries[0].AddendaRecordIndicator = 1
	batch.GetEntries()[0].AddAddenda05(addenda1)
	batch.GetEntries()[0].AddAddenda05(addenda2)
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
	fmt.Printf("%s", file.Batches[0].GetEntries()[0].Addenda05[0].String()+"\n")
	fmt.Printf("%s", file.Batches[0].GetEntries()[0].Addenda05[1].String()+"\n")
	fmt.Printf("%s", file.Batches[0].GetControl().String()+"\n")
	fmt.Printf("%s", file.Control.String()+"\n")

	// Output:
	// 101 03130001202313801041908161055A094101Federal Reserve Bank   My Bank Name           12345678
	// 5225Name on Account                     231380104 CTXACH CTX         190816   1121042880000001
	// 62723138010412345678         010000000045689033       0002Receiver Company  011121042880000001
	// 705Debit First Account                                                             00010000001
	// 705Debit Second Account                                                            00020000001
	// 82250000030023138010000100000000000000000000231380104                          121042880000001
	// 9000001000001000000030023138010000100000000000000000000

}
