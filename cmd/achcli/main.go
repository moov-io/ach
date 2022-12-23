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
	flagVerbose = flag.Bool("v", false, "Print verbose details about each ACH file\n")
	flagVersion = flag.Bool("version", false, "Print moov-io/ach cli version\n")

	flagDiff     = flag.Bool("diff", false, "Compare two files against each other\n")
	flagFlatten  = flag.Bool("flatten", false, "Flatten batches in each file\n")
	flagMerge    = flag.Bool("merge", false, "Merge files before describing\n")
	flagReformat = flag.String("reformat", "", "Reformat an incoming ACH file to another format\n")

	flagMask              = flag.Bool("mask", false, "Mask/hide full account numbers and individual names")
	flagMaskAccounts      = flag.Bool("mask.accounts", false, "Mask/hide full account numbers")
	flagMaskCorrectedData = flag.Bool("mask.corrections", false, "Mask/Hide Corrected Data in Addenda98 records")
	flagMaskNames         = flag.Bool("mask.names", false, "Mask/hide full individual names\n")

	flagPretty        = flag.Bool("pretty", false, "Display all values in their human readable format")
	flagPrettyAmounts = flag.Bool("pretty.amounts", false, "Display human readable amounts instead of exact values\n")

	flagValidate                                 = flag.String("validate", "", "Path to config file in json format to enable validation opts")
	flagValidateRequireABAOrigin                 = flag.Bool("validate.requireABAOrigin", false, "Enable routing number validation over the ImmediateOrigin file header field")
	flagValidateBypassOriginValidation           = flag.Bool("validate.bypassOriginValidation", false, "Skip validation for the ImmediateOrigin file header field")
	flagValidateBypassDestinationValidation      = flag.Bool("validate.bypassDestinationValidation", false, "Skip validation for the ImmediateDestination file header field")
	flagValidateCustomTraceNumbers               = flag.Bool("validate.customTraceNumbers", false, "Disable Nacha specified checks of TraceNumbers")
	flagValidateAllowZeroBatches                 = flag.Bool("validate.allowZeroBatches", false, "Allow the file to have zero batches")
	flagValidateAllowMissingFileHeader           = flag.Bool("validate.allowMissingFileHeader", false, "Allow a file to be read without a FileHeader record")
	flagValidateAllowMissingFileControl          = flag.Bool("validate.allowMissingFileControl", false, "Allow a file to be read without a FileControl record")
	flagValidateBypassCompanyIdentificationMatch = flag.Bool("validate.bypassCompanyIdentificationMatch", false, "Allow batches in which the Company Identification field in the batch header and control do not match")
	flagValidateCustomReturnCodes                = flag.Bool("validate.customReturnCodes", false, "Skip validation for the Return Code field in an Addenda99")
	flagValidateAllowUnorderedBatchNumbers       = flag.Bool("validate.allowUnorderedBatchNumbers", false, "Allow a file to be read with unordered batch numbers")
	flagValidateAllowInvalidCheckDigit           = flag.Bool("validate.allowInvalidCheckDigit", false, "Allow the CheckDigit field in EntryDetail to differ from the expected calculation")
	flagValidateUnequalAddendaCounts             = flag.Bool("validate.unequalAddendaCounts", false, "Skip checking that Addenda Count fields match their expected and computed values")
	flagValidateUnequalServiceClassCode          = flag.Bool("validate.unequalServiceClassCode", false, "Skips equality checks for the ServiceClassCode in each pair of BatchHeader\n")
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

	// getting validation opts
	validateOpts := getValidationOpts()

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

func getValidationOpts() (opts *ach.ValidateOpts) {

	var newOpts ach.ValidateOpts
	var existFlg bool

	if *flagValidate != "" {
		// read config file
		buf, readErr := os.ReadFile(*flagValidate)
		if readErr == nil {
			if readErr = json.Unmarshal(buf, &newOpts); readErr == nil {
				opts = &newOpts
				return
			}
		}

		if readErr != nil {
			validationHelp()
			os.Exit(1)
		}
	}

	if *flagValidateRequireABAOrigin {
		existFlg = true
		newOpts.RequireABAOrigin = true
	}
	if *flagValidateBypassOriginValidation {
		existFlg = true
		newOpts.BypassOriginValidation = true
	}
	if *flagValidateBypassDestinationValidation {
		existFlg = true
		newOpts.BypassDestinationValidation = true
	}
	if *flagValidateCustomTraceNumbers {
		existFlg = true
		newOpts.CustomTraceNumbers = true
	}

	if *flagValidateAllowZeroBatches {
		existFlg = true
		newOpts.AllowZeroBatches = true
	}
	if *flagValidateAllowMissingFileHeader {
		existFlg = true
		newOpts.AllowMissingFileHeader = true
	}
	if *flagValidateAllowMissingFileControl {
		existFlg = true
		newOpts.AllowMissingFileControl = true
	}
	if *flagValidateBypassCompanyIdentificationMatch {
		existFlg = true
		newOpts.BypassCompanyIdentificationMatch = true
	}
	if *flagValidateCustomReturnCodes {
		existFlg = true
		newOpts.CustomReturnCodes = true
	}
	if *flagValidateUnequalServiceClassCode {
		existFlg = true
		newOpts.UnequalServiceClassCode = true
	}
	if *flagValidateAllowUnorderedBatchNumbers {
		existFlg = true
		newOpts.AllowUnorderedBatchNumbers = true
	}
	if *flagValidateAllowInvalidCheckDigit {
		existFlg = true
		newOpts.AllowInvalidCheckDigit = true
	}
	if *flagValidateUnequalAddendaCounts {
		existFlg = true
		newOpts.UnequalAddendaCounts = true
	}

	if !existFlg {
		return
	}

	opts = &newOpts
	return
}
