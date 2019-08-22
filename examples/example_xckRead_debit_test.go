package examples

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/moov-io/ach"
)

func Example_xckReadDebit() {
	f, err := os.Open(filepath.Join("testdata", "xck-debit.ach"))
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
	fmt.Printf("Check Serial Number: %s", achFile.Batches[0].GetEntries()[0].IdentificationNumber+"\n")
	fmt.Printf("Process Control Field: %s", achFile.Batches[0].GetEntries()[0].IndividualName[0:6]+"\n")
	fmt.Printf("Item Research Number: %s", achFile.Batches[0].GetEntries()[0].IndividualName[6:22]+"\n")

	// Output:
	// SEC Code: XCK
	// Check Serial Number: 123456789012345
	// Process Control Field: CHECK1
	// Item Research Number: 182726

}
