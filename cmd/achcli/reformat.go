// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/moov-io/ach"
)

func reformat(as string, filepath string, validateOpts *ach.ValidateOpts) error {
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

var stdinRead bool

// readInput reads a file, or stdin when path is "-".
func readInput(path string) ([]byte, error) {
	if path != "-" {
		return os.ReadFile(path)
	}
	if stdinRead {
		return nil, errors.New("stdin (-) can only be read once")
	}
	stdinRead = true
	return io.ReadAll(os.Stdin)
}

func readIncomingFile(path string, validateOpts *ach.ValidateOpts) (*ach.File, error) {
	bs, err := readInput(path)
	if err != nil {
		return nil, err
	}
	if json.Valid(bs) {
		return readJsonFile(bs, validateOpts)
	}
	return readACHFile(bs, validateOpts)
}
