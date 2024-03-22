// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/moov-io/ach"
)

func reformat(as string, filepath string, validateOpts *ach.ValidateOpts) error {
	if _, err := os.Stat(filepath); err != nil {
		return err
	}

	file, err := readIncomingFile(filepath, validateOpts)
	if err != nil {
		return err
	}

	switch as {
	case "ach":
		w := ach.NewWriter(os.Stdout)
		if err := w.Write(file); err != nil {
			return err
		}

	case "json":
		if err := json.NewEncoder(os.Stdout).Encode(file); err != nil {
			return err
		}

	default:
		return fmt.Errorf("unknown format %s", as)
	}
	return nil
}

func readIncomingFile(path string, validateOpts *ach.ValidateOpts) (*ach.File, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if json.Valid(bs) {
		return readJsonFile(bs, validateOpts)
	}
	return readACHFile(bs, validateOpts)
}
