package main

import (
	"fmt"
	"log"
	"os"

	"github.com/moov-io/ach"
)

func main() {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open("ppd-debit.ach")
	if err != nil {
		log.Panicf("Can not open file: %s: \n", err)
	}
	r := ach.NewReader(f)
	_, err = r.Read()
	if err != nil {
		fmt.Printf("Issue reading file: %+v \n", err)
	}
	// ensure we have a validated file structure
	if r.File.ValidateAll(); err != nil {
		fmt.Printf("Could not validate entire read file: %v", err)
	}
	// If you trust the file but it's formating is off building will probably resolve the malformed file.
	if r.File.Build(); err != nil {
		fmt.Printf("Could not build file with read properties: %v", err)
	}

	fmt.Printf("total amount debit: %v \n", r.File.Control.TotalDebitEntryDollarAmountInFile)
}
