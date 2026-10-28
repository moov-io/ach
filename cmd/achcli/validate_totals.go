// Copyright 2026 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/moov-io/ach"
)

type validationMode struct {
	flag  string
	stage string
	check func(file *ach.File, opts *ach.ValidateOpts) error
}

var (
	totalsValidation = validationMode{
		flag:  "-validate-totals",
		stage: "validating totals",
		check: func(file *ach.File, _ *ach.ValidateOpts) error { return file.ValidateTotals() },
	}
	fileValidation = validationMode{
		flag:  "-validate-file",
		stage: "validating file",
		check: func(file *ach.File, opts *ach.ValidateOpts) error { return file.ValidateWith(opts) },
	}
)

func selectedValidation() (validationMode, bool) {
	switch {
	case *flagValidateFile:
		return fileValidation, true
	case *flagValidateTotals:
		return totalsValidation, true
	}
	return validationMode{}, false
}

func validateTotalsFlags(v validationMode) error {
	if *flagValidateTotals && *flagValidateFile {
		return errors.New("-validate-totals cannot be combined with -validate-file, which also checks totals")
	}
	if *flagSkipValidation {
		return fmt.Errorf("%s cannot be combined with -skip-validation", v.flag)
	}
	if *flagVersion || *flagDiff || *flagFix || *flagMerge || *flagFlatten || *flagUpdateEED != "" ||
		*flagMask || *flagMaskAccounts || *flagMaskCorrectedData || *flagMaskNames || *flagPretty || *flagPrettyAmounts {
		return fmt.Errorf("%s cannot be combined with version, transformation or presentation flags", v.flag)
	}
	return nil
}

func validateTotalsFiles(v validationMode, paths []string, opts *ach.ValidateOpts, reformatAs string) error {
	if len(paths) == 0 {
		return fmt.Errorf("%s requires at least one ACH file", v.flag)
	}
	if reformatAs != "" && len(paths) != 1 {
		return fmt.Errorf("%s with -reformat requires exactly one ACH file, found %d", v.flag, len(paths))
	}
	if opts != nil {
		if opts.SkipAll {
			return fmt.Errorf("%s cannot be used with SkipAll", v.flag)
		}
		if opts.UnequalAddendaCounts {
			return fmt.Errorf("%s cannot be used with UnequalAddendaCounts", v.flag)
		}
	}

	var failures []error
	for _, path := range paths {
		file, err := validateTotalsFile(v, path, opts)
		if err != nil {
			failures = append(failures, err)
		} else if reformatAs != "" {
			return writeFormat(reformatAs, file)
		}
	}
	return errors.Join(failures...)
}

func validateTotalsFile(v validationMode, path string, opts *ach.ValidateOpts) (*ach.File, error) {
	input, err := readInput(path)
	if err != nil {
		return nil, fmt.Errorf("%s: reading input: %w", path, err)
	}
	if json.Valid(input) {
		return nil, fmt.Errorf("%s: %s accepts raw ACH only; JSON parsing recalculates totals", path, v.flag)
	}
	file, err := readACHFile(input, opts)
	if err != nil {
		return nil, fmt.Errorf("%s: reading ACH: %w", path, err)
	}
	// Missing-record options must not waive the presence of original control totals.
	if file.Control == (ach.FileControl{}) && file.ADVControl == (ach.ADVFileControl{}) {
		return nil, fmt.Errorf("%s: reading ACH: FileControl: %w", path, ach.ErrFileControl)
	}
	if err := v.check(file, opts); err != nil {
		return nil, fmt.Errorf("%s: %s: %w", path, v.stage, err)
	}
	return file, nil
}
