// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/moov-io/ach"
)

func dumpFiles(paths []string) error {
	for i := range paths {
		if i > 0 {
			fmt.Println("") // extra newline between multiple ACH files
		}
		fmt.Printf("Describing ACH file '%s'\n", paths[i])
		file, err := readACHFile(paths[i])
		if err != nil {
			return fmt.Errorf("problem reading %s: %v", paths[i], err)
		}
		dumpFile(file)
	}
	return nil
}

func dumpFile(file *ach.File) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	defer w.Flush()

	fh := file.Header

	// FileHeader
	fmt.Fprintln(w, "  Origin\tOriginName\tDestination\tDestinationName")
	fmt.Fprintln(w, fmt.Sprintf("  %s\t%s\t%s\t%s\t", fh.ImmediateOrigin, fh.ImmediateOriginName, fh.ImmediateDestination, fh.ImmediateDestinationName))

	// Batches
	for i := range file.Batches {
		fmt.Fprintln(w, "\n  BatchNumber\tType\tServiceClass\tAddendas\tTotalDebits\tTotalCredits")

		bh := file.Batches[i].GetHeader()
		bc := file.Batches[i].GetControl()
		fmt.Fprintln(w, fmt.Sprintf("  %d\t%s\t%d\t%d\t%d\t%d", bh.BatchNumber, bh.StandardEntryClassCode, bh.ServiceClassCode, bc.EntryAddendaCount, bc.TotalDebitEntryDollarAmount, bc.TotalCreditEntryDollarAmount))

		entries := file.Batches[i].GetEntries()
		if len(entries) > 0 {
			fmt.Fprintln(w, "\n    TxCode\tAccountNumber\tAmount\tName\tTraceNumber")
		}
		for j := range entries {
			e := entries[j]
			fmt.Fprintln(w, fmt.Sprintf("    %d\t%s\t%d\t%s\t%s", e.TransactionCode, e.DFIAccountNumber, e.Amount, e.IndividualName, e.TraceNumber))
		}
	}
	// TODO(adam): Do different stuff with -v enabeld
}

func readACHFile(path string) (*ach.File, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("problem opening %s: %v", path, err)
	}
	defer fd.Close()

	f, err := ach.NewReader(fd).Read()
	return &f, err
}
