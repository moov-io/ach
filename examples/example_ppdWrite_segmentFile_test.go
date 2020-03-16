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

package examples

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/moov-io/ach"
)

func Example_ppdWriteSegmentFile() {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("testdata", "ppd-mixedDebitCredit.ach"))
	if err != nil {
		log.Fatal(err)
	}
	r := ach.NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		fmt.Printf("Issue reading file: %+v \n", err)
	}

	// verify a TraceNumber is set
	for i := range achFile.Batches {
		ed := achFile.Batches[i].GetEntries()
		for j := range ed {
			if ed[j].TraceNumber == "" {
				log.Fatalf("Batch[%d].Entries[%d] is missing a TraceNumber", i, j)
			}
		}
	}

	// ensure we have a validated file structure
	if achFile.Validate(); err != nil {
		fmt.Printf("Could not validate entire read file: %v", err)
	}

	sfc := ach.NewSegmentFileConfiguration()
	creditFile, debitFile, err := achFile.SegmentFile(sfc)

	if err != nil {
		fmt.Printf("Could not segment the file: %v", err)
	}

	if err := creditFile.Validate(); err != nil {
		log.Fatal(err)
	}
	if err := debitFile.Validate(); err != nil {
		log.Fatal(err)
	}

	ed := creditFile.Batches[0].GetEntries()
	for i := range ed {
		fmt.Printf("credit %d: %s\n", i, ed[i].String())
	}

	ed = debitFile.Batches[0].GetEntries()
	for i := range ed {
		fmt.Printf("debit %d: %s\n", i, ed[i].String())
	}

	// Output:
	// credit 0: 622231380104987654321        0100000000               Credit Account 1        0121042880000002
	// credit 1: 622231380104837098765        0100000000               Credit Account 2        0121042880000003
	// debit 0: 627231380104123456789        0200000000               Debit Account           0121042880000001

}
