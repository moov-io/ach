package examples

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/moov-io/ach"
)

func Example_ppdReadCredit() {
	f, err := os.Open(filepath.Join("examplesData", "ppd-credit.ach"))
	if err != nil {
		log.Fatal(err)
	}
	r := ach.NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		fmt.Printf("Issue reading file: %+v \n", err)
	}

	if achFile.Validate(); err != nil {
		fmt.Printf("Could not validate entire read file: %v", err)
	}

	if err := achFile.Create(); err != nil {
		fmt.Printf("Could not create file with read properties: %v", err)
	}

	fmt.Printf("SEC Code: %s", achFile.Batches[0].GetHeader().StandardEntryClassCode+"\n")
	fmt.Printf("Total File Debit Amount: %s", strconv.Itoa(achFile.Control.TotalDebitEntryDollarAmountInFile)+"\n")
	fmt.Printf("Total File Credit Amount: %s", strconv.Itoa(achFile.Control.TotalCreditEntryDollarAmountInFile)+"\n")

	// Output:
	// SEC Code: PPD
	// Total File Debit Amount: 0
	// Total File Credit Amount: 100000000
}
