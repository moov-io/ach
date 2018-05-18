package main

import (
	"fmt"
	"log"
	"os"

	"flag"
	"github.com/moov-io/ach"
	"runtime/pprof"
)

func main() {

	var fPath = flag.String("fPath", "201805101354.ach", "File Path")
	var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	path := *fPath

	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(path)

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
	// If you trust the file but it's formating is off building will probably resolve the malformed file.
	if achFile.Create(); err != nil {
		fmt.Printf("Could not build file with read properties: %v", err)
	}

	fmt.Printf("total amount debit: %v \n", achFile.Control.TotalDebitEntryDollarAmountInFile)
	fmt.Printf("total amount credit: %v \n", achFile.Control.TotalCreditEntryDollarAmountInFile)
}
