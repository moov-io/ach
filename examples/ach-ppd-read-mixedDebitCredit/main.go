package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/moov-io/ach"
)

func main() {

	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("examples", "ach-ppd-read-mixedDebitCredit", "ppd-mixedDebitCredit.ach"))
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

	fmt.Printf("File Name: %s \n\n", f.Name())
	fmt.Printf("Total Debit Amount: %v \n", achFile.Control.TotalDebitEntryDollarAmountInFile)
	fmt.Printf("Total Credit Amount: %v \n", achFile.Control.TotalCreditEntryDollarAmountInFile)
	fmt.Printf("File Header: %v \n", achFile.Header.String())
	fmt.Printf("Batch Header: %v \n", achFile.Batches[0].GetHeader().String())
	fmt.Printf("Entry Detail 1: %v \n", achFile.Batches[0].GetEntries()[0].String())
	fmt.Printf("Entry Detail 2: %v \n", achFile.Batches[0].GetEntries()[1].String())
	fmt.Printf("Entry Detail 3: %v \n", achFile.Batches[0].GetEntries()[2].String())
	fmt.Printf("Batch Control: %v \n", achFile.Batches[0].GetControl().String())
	fmt.Printf("File Header: %v \n", achFile.Control.String())

	// ToDo:  This may be temporary: Create a specific example for creating the Segmented File

	creditFile, debitFile, err := achFile.SegmentFile()

	if err != nil {
		fmt.Printf("Could not segment the file: %v", err)
	}

	fmt.Printf("Total Debit Amount: %v \n", creditFile.Control.TotalDebitEntryDollarAmountInFile)
	fmt.Printf("Total Credit Amount: %v \n", creditFile.Control.TotalCreditEntryDollarAmountInFile)
	fmt.Printf("File Header: %v \n", creditFile.Header.String())
	fmt.Printf("Batch Header: %v \n", creditFile.Batches[0].GetHeader().String())
	fmt.Printf("Entry Detail 1: %v \n", creditFile.Batches[0].GetEntries()[0].String())
	fmt.Printf("Entry Detail 2: %v \n", creditFile.Batches[0].GetEntries()[1].String())
	fmt.Printf("Batch Control: %v \n", creditFile.Batches[0].GetControl().String())
	fmt.Printf("File Header: %v \n", creditFile.Control.String())

	fmt.Printf("Total Debit Amount: %v \n", debitFile.Control.TotalDebitEntryDollarAmountInFile)
	fmt.Printf("Total Credit Amount: %v \n", debitFile.Control.TotalCreditEntryDollarAmountInFile)
	fmt.Printf("File Header: %v \n", debitFile.Header.String())
	fmt.Printf("Batch Header: %v \n", debitFile.Batches[0].GetHeader().String())
	fmt.Printf("Entry Detail 1: %v \n", debitFile.Batches[0].GetEntries()[0].String())
	fmt.Printf("Batch Control: %v \n", debitFile.Batches[0].GetControl().String())
	fmt.Printf("File Header: %v \n", debitFile.Control.String())
}
