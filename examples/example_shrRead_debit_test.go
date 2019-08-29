package examples

import (
	"fmt"
	"github.com/moov-io/ach"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func Example_shrReadDebit() {
	f, err := os.Open(filepath.Join("testdata", "shr-debit.ach"))
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
	fmt.Printf("SHR Card Expiration Date: %s", achFile.Batches[0].GetEntries()[0].IdentificationNumber[0:4]+"\n")
	fmt.Printf("SHR Document Reference Number: %s", achFile.Batches[0].GetEntries()[0].IdentificationNumber[4:15]+"\n")
	fmt.Printf("SHR Individual Card Account Number: %s", achFile.Batches[0].GetEntries()[0].IndividualName)

	// Output:
	// Total Amount Debit: 100000000
	// SEC Code: SHR
	// SHR Card Expiration Date: 0722
	// SHR Document Reference Number: 12345678910
	// SHR Individual Card Account Number: 0001234567891123456789
}
