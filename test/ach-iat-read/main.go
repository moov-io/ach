package main

import (
	"fmt"
	"github.com/moov-io/ach"
	"log"
	"os"
)

func main() {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open("iat-ach.ach")
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
	if achFile.Create(); err != nil {
		fmt.Printf("Could not build file with read properties: %v", err)
	}

	fmt.Printf("Total Amount Debit: %v \n", achFile.Control.TotalDebitEntryDollarAmountInFile)
	fmt.Printf("SEC Code: %v \n", achFile.IATBatches[0].GetHeader().StandardEntryClassCode)
	fmt.Printf("Entry: %v \n", achFile.IATBatches[0].GetEntries()[0])
	fmt.Printf("Addenda Record Indicator: %v \n", achFile.IATBatches[0].GetEntries()[0].AddendaRecordIndicator)
	fmt.Printf("Addenda10: %v \n", achFile.IATBatches[0].GetEntries()[0].Addenda10)
	fmt.Printf("Addenda11: %v \n", achFile.IATBatches[0].GetEntries()[0].Addenda11)
	fmt.Printf("Addenda12: %v \n", achFile.IATBatches[0].GetEntries()[0].Addenda12)
	fmt.Printf("Addenda13: %v \n", achFile.IATBatches[0].GetEntries()[0].Addenda13)
	fmt.Printf("Addenda14: %v \n", achFile.IATBatches[0].GetEntries()[0].Addenda14)
	fmt.Printf("Addenda15: %v \n", achFile.IATBatches[0].GetEntries()[0].Addenda15)
	fmt.Printf("Addenda16: %v \n", achFile.IATBatches[0].GetEntries()[0].Addenda16)
	fmt.Printf("Addenda17: %v \n", achFile.IATBatches[0].GetEntries()[0].Addendum[0].String())
	fmt.Printf("Addenda18: %v \n", achFile.IATBatches[0].GetEntries()[0].Addendum[1].String())
}
