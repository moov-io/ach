package main

import (
	"fmt"
	"github.com/moov-io/ach"
	"log"
	"os"
	"path/filepath"
)

func main() {
	f, err := os.Open(filepath.Join("test", "ach-cie-read", "cie-credit.ach"))
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
	fmt.Printf("Addenda05: %v \n", achFile.Batches[0].GetEntries()[0].Addenda05[0].String())
}
