// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/moov-io/ach"
	"github.com/moov-io/ach/cmd/achcli/describe"
)

func dumpFiles(paths []string) error {
	var files []*ach.File
	for i := range paths {
		f, err := readACHFile(paths[i])
		if err != nil {
			fmt.Printf("WARN: problem reading %s:\n %v\n\n", paths[i], err)
		}
		if *flagMerge {
			files = append(files, f)
		} else {
			if i > 0 {
				fmt.Println("") // extra newline between multiple ACH files
			}
			fmt.Printf("Describing ACH file '%s'\n\n", paths[i])
			describe.File(os.Stdout, f, &describe.Opts{
				MaskAccountNumbers: *flagMask,
			})
		}
	}

	if *flagMerge {
		merged, err := ach.MergeFiles(files)
		if err != nil {
			fmt.Printf("ERROR: merging files: %v\n", err)
		}
		fmt.Printf("Describing %d file(s) merged into %d file(s)\n", len(paths), len(files))
		files = merged
	}

	if *flagFlatten {
		for i := range files {
			file, err := files[i].FlattenBatches()
			if err != nil {
				fmt.Printf("ERROR: problem flattening file: %v\n", err)
			}
			files[i] = file
		}
	}

	for i := range files {
		if i > 0 && len(files) > 1 {
			fmt.Println("") // extra newline between multiple ACH files
		}
		if files[i] != nil {
			describe.File(os.Stdout, files[i], &describe.Opts{
				MaskAccountNumbers: *flagMask,
			})
		} else {
			fmt.Printf("nil ACH file in position %d\n", i)
		}
	}

	return nil
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
