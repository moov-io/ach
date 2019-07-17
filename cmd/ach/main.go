// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/moov-io/ach"
)

var (
	flagVerbose = flag.Bool("v", false, "Print verbose details about each ACH file")
	flagVersion = flag.Bool("version", false, "Print moov-io/ach cli version")
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage of ach (%s):\n", ach.Version)
		fmt.Println("   usage: ach [<flags>] <files>")
		fmt.Println("")
		fmt.Println("Examples: ")
		fmt.Println("  ach -v 20060102.ach")
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
	if len(args) == 0 {
		fmt.Println("No command or ACH files provided, see -help")
		os.Exit(1)
	}
	if err := describeACHFiles(args); err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}
