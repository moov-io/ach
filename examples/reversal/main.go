// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/moov-io/ach"
	"github.com/moov-io/ach/cmd/achcli/describe"
)

func main() {
	file, err := ach.ReadJSONFile(filepath.Join("test", "testdata", "ppd-valid.json"))
	if err != nil {
		log.Fatalf("ERROR: unable to read file: %v", err)
	}

	fmt.Println("Original:")
	describe.File(os.Stdout, file, nil)
	fmt.Printf("\n\n")

	effectiveEntryDate := time.Now().In(time.UTC)
	err = file.Reversal(effectiveEntryDate)
	if err != nil {
		log.Fatalf("ERROR: problem reversing file: %v", err)
	}

	fmt.Println("Reversal:")
	describe.File(os.Stdout, file, nil)
}
