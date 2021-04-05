// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/moov-io/ach"
)

var (
	flagVerbose = flag.Bool("v", false, "Print verbose details about each ACH file")
	flagVersion = flag.Bool("version", false, "Print moov-io/ach cli version")

	flagDiff     = flag.Bool("diff", false, "Compare two files against each other")
	flagMerge    = flag.Bool("merge", false, "Merge files before describing")
	flagMask     = flag.Bool("mask", false, "Mask/redact full account numbers")
	flagReformat = flag.String("reformat", "", "Reformat an incoming ACH file to another format")

	// todo: add dry-run option to print to stdout
	// todo: add file output option; flagOutput = flag.String("output", "", "output file")
	flagRedact = flag.Bool("redact", false, "redact stuff")
	flagDryRun = flag.Bool("dry-run", false, "dry run to print to stdout")

	programName = filepath.Base(os.Args[0])
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage of ach (%s):\n", ach.Version)
		fmt.Printf("   usage: %s [<flags>] <files>", programName)
		fmt.Println("")
		fmt.Println("Commands: ")
		fmt.Printf("  %s -diff first.ach second.ach", programName)
		fmt.Println("    Show the difference between two ACH files")
		fmt.Printf("  %s -mask file.ach", programName)
		fmt.Println("                Print file details with personally identifiable information partially removed")
		fmt.Printf("  %s -reformat=json first.ach", programName)
		fmt.Println("      Convert an incoming ACH file into another format (options: ach, json)")
		fmt.Printf("  %s 20060102.ach", programName)
		fmt.Println("                  Summarize an ACH file for human readability")
		fmt.Println("")
		fmt.Println("Flags: ")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	switch {
	case *flagVersion:
		fmt.Printf("moov-io/ach:%s cli tool\n", ach.Version)
		return
	case *flagVerbose:
		fmt.Printf("moov-io/ach:%s cli tool\n", ach.Version)
	}

	args := flag.Args()

	// error conditions, verify we're okay for whatever the task at hand is
	switch {
	case *flagDiff && len(args) != 2:
		fmt.Printf("with -diff exactly two files are expected, found %d files\n", len(args))
		os.Exit(1)
		//todo: add redact here
	}

	// minor debugging
	if *flagVerbose {
		fmt.Printf("found %d ACH files to describe: %s\n", len(args), strings.Join(args, ", "))
	}

	// pick our command to do
	switch {
	case *flagDiff:
		if err := diffFiles(args); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}

	case *flagReformat != "" && len(args) == 1:
		if err := reformat(*flagReformat, args[0]); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}

	case *flagRedact:
		if err := redact(args, *flagDryRun); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}

	default:
		if err := dumpFiles(args); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}
	}
}
