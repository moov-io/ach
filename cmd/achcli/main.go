// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/moov-io/ach"
)

var (
	flagVerbose = flag.Bool("v", false, "Print verbose details about each ACH file")
	flagVersion = flag.Bool("version", false, "Print moov-io/ach cli version")

	flagDiff     = flag.Bool("diff", false, "Compare two files against each other")
	flagFlatten  = flag.Bool("flatten", false, "Flatten batches in each file")
	flagMerge    = flag.Bool("merge", false, "Merge files before describing")
	flagReformat = flag.String("reformat", "", "Reformat an incoming ACH file to another format")

	flagMask              = flag.Bool("mask", false, "Mask/hide full account numbers and individual names")
	flagMaskAccounts      = flag.Bool("mask.accounts", false, "Mask/hide full account numbers")
	flagMaskCorrectedData = flag.Bool("mask.corrections", false, "Mask/Hide Corrected Data in Addenda98 records")
	flagMaskNames         = flag.Bool("mask.names", false, "Mask/hide full individual names")

	flagPretty        = flag.Bool("pretty", false, "Display all values in their human readable format")
	flagPrettyAmounts = flag.Bool("pretty.amounts", false, "Display human readable amounts instead of exact values")

	flagSkipValidation = flag.Bool("skip-validation", false, "Skip all validation checks")
	flagValidateOpts   = flag.String("validate", "", "Path to config file in json format to enable validation opts")
)

func main() {
	flag.Usage = help
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
	}

	// minor debugging
	if *flagVerbose {
		fmt.Printf("found %d ACH files to describe: %s\n", len(args), strings.Join(args, ", "))
	}

	// checking input files
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Read validation options from the command
	validateOpts := readValidationOpts(*flagValidateOpts)

	// pick our command to do
	switch {
	case *flagDiff:
		if err := diffFiles(args, validateOpts); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}

	case *flagReformat != "" && len(args) == 1:
		if err := reformat(*flagReformat, args[0], validateOpts); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}

	default:
		if err := dumpFiles(args, validateOpts); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}
	}
}

func readValidationOpts(path string) *ach.ValidateOpts {
	var opts ach.ValidateOpts

	if *flagSkipValidation {
		opts.SkipAll = true
		return &opts
	}

	if path != "" {
		// read config file
		bs, readErr := os.ReadFile(path)
		if readErr != nil {
			fmt.Printf("ERROR: reading validate opts failed: %v\n", readErr)
			os.Exit(1)
		}

		if err := json.Unmarshal(bs, &opts); err != nil {
			fmt.Printf("ERROR: unmarshal of validate opts failed: %v\n", err)
			os.Exit(1)
		}
		return &opts
	}
	return nil
}
