// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/moov-io/ach"
)

var (
	fPath      = flag.String("fPath", "201805101354.ach", "File Path")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

	flagJson = flag.Bool("json", false, "Output ACH File in JSON to stdout")
)

func main() {
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
	if err := achFile.Create(); err != nil {
		fmt.Printf("Could not create file with read properties: %v", err)
	}

	// Output file contents
	if *flagJson {
		if err := json.NewEncoder(os.Stdout).Encode(achFile); err != nil {
			fmt.Printf("ERROR: problem writing ACH File to stdout: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("total amount debit: %v \n", achFile.Control.TotalDebitEntryDollarAmountInFile)
		fmt.Printf("total amount credit: %v \n", achFile.Control.TotalCreditEntryDollarAmountInFile)
	}
}
