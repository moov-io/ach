// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

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

	sfc := ach.NewSegmentFileConfiguration()
	creditFile, debitFile, err := achFile.SegmentFile(sfc)

	if err != nil {
		fmt.Printf("Could not segment the file: %v", err)
	}

	// write the file to std out. Anything io.Writer
	w := ach.NewWriter(os.Stdout)
	if err := w.Write(creditFile); err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}
	if err := w.Write(debitFile); err != nil {
		log.Fatalf("Unexpected error: %s\n", err)
	}

	w.Flush()

}
