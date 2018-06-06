package main

import (
	"github.com/moov-io/ach"
	"fmt"
	"log"
	"os"


)

func main() {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open("pop-debit.ach")
	if err != nil {
		log.Panicf("Can not open file: %s: \n", err)
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

	fmt.Printf("total amount debit: %v \n", achFile.Control.TotalDebitEntryDollarAmountInFile)
	fmt.Printf("POP check serial number: %v \n", achFile.Batches[0].GetEntries()[0].POPCheckSerialNumberField())
	fmt.Printf("POP terminal city: %v \n", achFile.Batches[0].GetEntries()[0].POPTerminalCityField())
	fmt.Printf("POP terminal state: %v \n", achFile.Batches[0].GetEntries()[0].POPTerminalStateField())
}
