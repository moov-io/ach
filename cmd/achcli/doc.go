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

func help() {
	fmt.Printf(strings.TrimSpace(`
achcli is a tool for displaying Nacha formatted ACH files in a human readable format.

USAGE
   achcli [-mask] [-pretty] path/to/file.ach

EXAMPLES
  achcli -diff first.ach second.ach    Show the difference between two ACH files
  achcli -mask file.ach                Print file details with personally identifiable information partially removed
  achcli -reformat=json first.ach      Convert an incoming ACH file into another format (options: ach, json)
  achcli -version                      Print the version of achcli (Example: %s)
  achcli 20060102.ach                  Summarize an ACH file for human readability

FLAGS
`), ach.Version)
	fmt.Println("")
	flag.PrintDefaults()
}

func validationHelp() {
	fmt.Fprintf(os.Stdout, "\nSpecify validation config file in json foramt to enable valiation opts.\n\nEXAMPLE:\n")
	sampleJson, _ := json.MarshalIndent(ach.ValidateOpts{}, "", "  ")
	fmt.Fprintf(os.Stdout, "%s \n", string(sampleJson))
	fmt.Fprintf(os.Stdout, "\n")
}
