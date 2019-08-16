package examples

import (
	"fmt"
	"github.com/moov-io/ach"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func Example_webReadCredit() {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("examplesData", "web-credit.ach"))
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

	fmt.Printf("SEC Code: %s", achFile.Batches[0].GetHeader().StandardEntryClassCode+"\n")
	fmt.Printf("Total File Debit Amount: %s", strconv.Itoa(achFile.Control.TotalDebitEntryDollarAmountInFile)+"\n")
	fmt.Printf("Total File Credit Amount: %s", strconv.Itoa(achFile.Control.TotalCreditEntryDollarAmountInFile)+"\n")

	// Output:
	// SEC Code: WEB
	// Total File Debit Amount: 0
	// Total File Credit Amount: 10000

}
