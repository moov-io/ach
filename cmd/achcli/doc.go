// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/moov-io/ach"
)

func help() {
	fmt.Printf(strings.TrimSpace(`
achcli is a tool for displaying Nacha formatted ACH files in a human readable format.

USAGE
   achcli [-mask] [-pretty] [-validate opts.json] path/to/file.ach

EXAMPLES
  achcli -diff first.ach second.ach    Show the difference between two ACH files
  achcli -mask file.ach                Print file details with personally identifiable information partially removed
  achcli -reformat=json first.ach      Convert an incoming ACH file into another format (options: ach, json)
  achcli -validate opts.json file.ach  Read an ACH File with the provided ValidateOpts
  achcli -version                      Print the version of achcli (Example: %s)
  achcli 20060102.ach                  Summarize an ACH file for human readability

FLAGS
`), ach.Version)
	fmt.Println("")
	flag.PrintDefaults()
}
