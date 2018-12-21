package main

import (
	"fmt"
	"github.com/moov-io/ach"
	"log"
	"os"
)

func main() {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open("shr-debit.ach")
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
	fmt.Printf("SEC Code: %v \n", achFile.Batches[0].GetHeader().StandardEntryClassCode)
	fmt.Printf("SHR Card Expiration Date: %v \n", achFile.Batches[0].GetEntries()[0].SHRCardExpirationDateField())
	fmt.Printf("SHR Document Reference Number: %v \n", achFile.Batches[0].GetEntries()[0].SHRDocumentReferenceNumberField())
	fmt.Printf("SHR Individual Card Account Number: %v \n", achFile.Batches[0].GetEntries()[0].SHRIndividualCardAccountNumberField())
}
