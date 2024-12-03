// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"os"

	// "text/tabwriter"

	"github.com/moov-io/ach"

	"github.com/juju/ansiterm"
)

func diffFiles(paths []string, validateOpts *ach.ValidateOpts) error {
	if len(paths) != 2 {
		return fmt.Errorf("expected 2 files, but got %d", len(paths))
	}
	f1, f2, err := readTwoFiles(paths, validateOpts)
	if err != nil {
		return err
	}

	// w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	// defer w.Flush()

	w := ansiterm.NewTabWriter(os.Stdout, 0, 0, 1, ' ', 0)
	w = w.Init(os.Stdout, 0, 0, 1, ' ', 0)
	defer w.Flush()

	switch {
	case f1.Header.ImmediateOrigin != f2.Header.ImmediateOrigin:
		printDiffedFileHeader(w, f1, f2)
	case f1.Header.ImmediateOriginName != f2.Header.ImmediateOriginName:
		printDiffedFileHeader(w, f1, f2)
	case f1.Header.ImmediateDestination != f2.Header.ImmediateDestination:
		printDiffedFileHeader(w, f1, f2)
	case f1.Header.ImmediateDestinationName != f2.Header.ImmediateDestinationName:
		printDiffedFileHeader(w, f1, f2)
	}

	return nil
}

func readTwoFiles(paths []string, validateOpts *ach.ValidateOpts) (*ach.File, *ach.File, error) {
	f1, err := readIncomingFile(paths[0], validateOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("problem reading %s: %v", paths[0], err)
	}
	f2, err := readIncomingFile(paths[1], validateOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("problem reading %s: %v", paths[1], err)
	}
	return f1, f2, nil
}

// TODO(adam): we should lookup batches which are in f1 against those in f2 and if a similar
// batch is found (TODO: what, if any fields can we do exact matches on?) print that in the
// order of f1. Otherwise show missing batches (from f1's view) and additional batches (from
// f2's view) at the end.

func printDiffedFileHeader(w *ansiterm.TabWriter, f1, f2 *ach.File) {
	fmt.Fprintln(w, "  Origin\tOriginName\tDestination\tDestinationName")

	var minusBuf bytes.Buffer
	var plusBuf bytes.Buffer

	minusBuf.WriteString("- ")
	plusBuf.WriteString("+ ")

	printColumn(&minusBuf, &plusBuf, f1.Header.ImmediateOrigin, f2.Header.ImmediateOrigin)
	printColumn(&minusBuf, &plusBuf, f1.Header.ImmediateOriginName, f2.Header.ImmediateOriginName)
	printColumn(&minusBuf, &plusBuf, f1.Header.ImmediateDestination, f2.Header.ImmediateDestination)
	printColumn(&minusBuf, &plusBuf, f1.Header.ImmediateDestinationName, f2.Header.ImmediateDestinationName)

	fmt.Fprintln(w, minusBuf.String())
	fmt.Fprintln(w, plusBuf.String())

	// fmt.Fprintln(w, fmt.Sprintf("- %s\t%s\t%s\t%s\t", f2.Header.ImmediateOrigin, f2.Header.ImmediateOriginName, f2.Header.ImmediateDestination, f2.Header.ImmediateDestinationName))
	// fmt.Fprintln(w, fmt.Sprintf("+ %s\t%s\t%s\t%s\t", f1.Header.ImmediateOrigin, f1.Header.ImmediateOriginName, f1.Header.ImmediateDestination, f1.Header.ImmediateDestinationName))
}

func printColumn(minusBuf, plusBuf *bytes.Buffer, v1, v2 string) {
	if v1 != v2 {
		w := ansiterm.NewWriter(minusBuf)
		ctx := ansiterm.Foreground(ansiterm.Green)
		ctx.Fprintf(w, v2)
		ctx.SetForeground(ansiterm.Default)
		ctx.Fprintf(w, "\t")

		w = ansiterm.NewWriter(plusBuf)
		ctx = ansiterm.Foreground(ansiterm.Red)
		ctx.Fprintf(w, v1)
		ctx.SetForeground(ansiterm.Default)
		ctx.Fprintf(w, "\t")

		// minusBuf.WriteString(

		// ansiterm.Fprint(minusBuf, fmt.Sprintf("[red]%s[reset]\t", v2))
		// ansiterm.Fprint(plusBuf, fmt.Sprintf("[green]%s[reset]\t", v1))

		// ansiterm.Write([]byte(fmt.Sprintf("[red]%s[reset]\t", v2)))
		// ansiterm.Write([]byte(fmt.Sprintf("[green]%s[reset]\t", v1)))
	} else {
		minusBuf.WriteString(v2 + "\t")
		plusBuf.WriteString(v1 + "\t")
	}
}
