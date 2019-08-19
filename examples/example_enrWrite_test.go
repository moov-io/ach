package examples

import (
	"fmt"
	"github.com/moov-io/ach"
	"log"
)

// Example_enrWrite writes and ENR file
func Example_enrWrite() {
	fh := mockFileHeader()

	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.DebitsOnly
	bh.CompanyName = "Name on Account"
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.ENR
	bh.CompanyEntryDescription = "AUTOENROLL"
	bh.ODFIIdentification = "23138010"

	entry := ach.NewEntryDetail()

	entry.TransactionCode = ach.CheckingDebit
	entry.SetRDFI("031300012")
	entry.DFIAccountNumber = "744-5678-99"
	entry.Amount = 0
	entry.SetOriginalTraceNumber("031300010000001")
	entry.SetReceivingCompany("Best. #1")
	entry.SetTraceNumber(bh.ODFIIdentification, 1)

	addenda05 := ach.NewAddenda05()
	addenda05.PaymentRelatedInformation = `22*12200004*3*123987654321*777777777*DOE*JOHN*1\` // From NACHA 2013 Official Rules
	entry.AddAddenda05(addenda05)
	entry.AddendaRecordIndicator = 1

	// build the batch
	batch := ach.NewBatchENR(bh)
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
	fmt.Printf("%s", file.Batches[0].GetEntries()[0].Addenda05[0].String()+"\n")
	fmt.Printf("%s", file.Batches[0].GetControl().String()+"\n")
	fmt.Printf("%s", file.Control.String()+"\n")

	// Output: 101 03130001202313801041908161055A094101Federal Reserve Bank   My Bank Name           12345678
	// 5225Name on Account                     231380104 ENRAUTOENROLL               1231380100000001
	// 627031300012744-5678-99      0000000000031300010000001Best. #1                1231380100000001
	// 70522*12200004*3*123987654321*777777777*DOE*JOHN*1\                                00010000001
	// 82250000020003130001000000000000000000000000231380104                          231380100000001
	// 9000001000001000000020003130001000000000000000000000000
}
