package examples

import (
	"fmt"
	"github.com/moov-io/ach"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func Example_popReadDebit() {
	f, err := os.Open(filepath.Join("testdata", "pop-debit.ach"))
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
	fmt.Printf("SEC Code: %s", achFile.Batches[0].GetHeader().StandardEntryClassCode+"\n")
	fmt.Printf("POP Check Serial Number: %s", achFile.Batches[0].GetEntries()[0].IdentificationNumber[0:9]+"\n")
	fmt.Printf("POP Terminal City: %s", achFile.Batches[0].GetEntries()[0].IdentificationNumber[9:13]+"\n")
	fmt.Printf("POP Terminal State: %s", achFile.Batches[0].GetEntries()[0].IdentificationNumber[13:15]+"\n")

	// Output: Total Amount Debit: 250500
	// SEC Code: POP
	// POP Check Serial Number: 123456789
	// POP Terminal City: PHIL
	// POP Terminal State: PA
}
