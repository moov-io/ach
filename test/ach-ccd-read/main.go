package main

import (
	"fmt"
	"log"
	"os"

	"github.com/moov-io/ach"
)

func main() {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open("ccd-debit.ach")
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

	fmt.Printf("Total Amount Debit: %v \n", achFile.Control.TotalDebitEntryDollarAmountInFile)
	fmt.Printf("Total Amount Credit: %v \n", achFile.Control.TotalCreditEntryDollarAmountInFile)
	fmt.Printf("SEC Code: %v \n", achFile.Batches[0].GetHeader().StandardEntryClassCode)
	fmt.Printf("CCD Entry Discretionary Data: %v \n", achFile.Batches[0].GetEntries()[0].DiscretionaryDataField())
	fmt.Printf("CCD Entry Identification Number: %v \n", achFile.Batches[0].GetEntries()[0].IdentificationNumberField())
	fmt.Printf("CCD Entry Receiving Company: %v \n", achFile.Batches[0].GetEntries()[0].ReceivingCompanyField())
	fmt.Printf("CCD Entry Trace Number: %v \n", achFile.Batches[0].GetEntries()[0].TraceNumberField())
	fmt.Printf("CCD Fee Discretionary Data: %v \n", achFile.Batches[0].GetEntries()[1].DiscretionaryDataField())
	fmt.Printf("CCD Fee Identification Number: %v \n", achFile.Batches[0].GetEntries()[1].IdentificationNumberField())
	fmt.Printf("CCD Fee Receiving Company: %v \n", achFile.Batches[0].GetEntries()[1].ReceivingCompanyField())
	fmt.Printf("CCD Fee Trace Number: %v \n", achFile.Batches[0].GetEntries()[1].TraceNumberField())

}
