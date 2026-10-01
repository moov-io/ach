// Copyright 2026 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/moov-io/ach"
)

func validateTotalsFlags() error {
	if *flagSkipValidation {
		return errors.New("-validate-totals cannot be combined with -skip-validation")
	}
	if *flagVersion || *flagDiff || *flagFix || *flagMerge || *flagFlatten || *flagReformat != "" || *flagUpdateEED != "" ||
		*flagMask || *flagMaskAccounts || *flagMaskCorrectedData || *flagMaskNames || *flagPretty || *flagPrettyAmounts {
		return errors.New("-validate-totals cannot be combined with version, transformation or presentation flags")
	}
	return nil
}

func validateTotalsFiles(paths []string, opts *ach.ValidateOpts) error {
	if len(paths) == 0 {
		return errors.New("-validate-totals requires at least one ACH file")
	}
	if opts != nil {
		if opts.SkipAll {
			return errors.New("-validate-totals cannot be used with SkipAll")
		}
		if opts.UnequalAddendaCounts {
			return errors.New("-validate-totals cannot be used with UnequalAddendaCounts")
		}
	}

	var failures []error
	for _, path := range paths {
		if err := validateTotalsFile(path, opts); err != nil {
			failures = append(failures, err)
		}
	}
	return errors.Join(failures...)
}

func validateTotalsFile(path string, opts *ach.ValidateOpts) error {
	input, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("%s: reading input: %w", path, err)
	}
	if json.Valid(input) {
		return fmt.Errorf("%s: -validate-totals accepts raw ACH only; JSON parsing recalculates totals", path)
	}
	file, err := readACHFile(input, opts)
	if err != nil {
		return fmt.Errorf("%s: reading ACH: %w", path, err)
	}
	// Missing-record options must not waive the presence of original control totals.
	if file.Control == (ach.FileControl{}) && file.ADVControl == (ach.ADVFileControl{}) {
		return fmt.Errorf("%s: reading ACH: FileControl: %w", path, ach.ErrFileControl)
	}
	if err := file.ValidateTotals(); err != nil {
		return fmt.Errorf("%s: validating totals: %w", path, err)
	}
	return nil
}
