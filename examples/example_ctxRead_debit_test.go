package examples

import (
	"fmt"
	"github.com/moov-io/ach"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func Example_ctxReadDebit() {
	f, err := os.Open(filepath.Join("testdata", "ctx-debit.ach"))
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

	fmt.Printf("Total Amount Debit: %s", strconv.Itoa(achFile.Control.TotalDebitEntryDollarAmountInFile)+"\n")
	fmt.Printf("Total Amount Credit: %s", strconv.Itoa(achFile.Control.TotalCreditEntryDollarAmountInFile)+"\n")
	fmt.Printf("SEC Code: %s", achFile.Batches[0].GetHeader().StandardEntryClassCode+"\n")
	fmt.Printf("Addenda1: %s", achFile.Batches[0].GetEntries()[0].Addenda05[0].String()+"\n")
	fmt.Printf("Addenda2: %s", achFile.Batches[0].GetEntries()[0].Addenda05[1].String()+"\n")

	// Output:
	// Total Amount Debit: 100000000
	// Total Amount Credit: 0
	// SEC Code: CTX
	// Addenda1: 705Debit First Account                                                             00010000001
	// Addenda2: 705Debit Second Account                                                            00020000001

}
