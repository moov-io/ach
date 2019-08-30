package examples

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/moov-io/ach"
)

// Example_dneRead reads a DNR file
func Example_dneRead() {
	f, err := os.Open(filepath.Join("testdata", "dne-read.ach"))
	if err != nil {
		log.Fatal(err)
	}
	r := ach.NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		panic(fmt.Sprintf("Issue reading file: %+v \n", err))
	}
	// ensure we have a validated file structure
	if achFile.Validate(); err != nil {
		fmt.Printf("Could not validate entire read file: %v", err)
	}
	// If you trust the file but it's formatting is off building will probably resolve the malformed file.
	if err := achFile.Create(); err != nil {
		fmt.Printf("Could not create file with read properties: %v", err)
	}

	fmt.Printf("Total Amount: %s", strconv.Itoa(achFile.Batches[0].GetEntries()[0].Amount)+"\n")
	fmt.Printf("SEC Code: %s", achFile.Batches[0].GetHeader().StandardEntryClassCode+"\n")
	fmt.Printf("Payment Related Information: %s", achFile.Batches[0].GetEntries()[0].Addenda05[0].String()+"\n")

	// Output:
	// Total Amount: 0
	// SEC Code: DNE
	// Payment Related Information: 705DATE OF DEATH*010218*CUSTOMERSSN*#########*AMOUNT*$$$$.cc\                      00010000001

}
