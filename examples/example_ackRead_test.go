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
	"github.com/moov-io/ach"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// Example_ackRead reads a ACK file
func Example_ackRead() {
	f, err := os.Open(filepath.Join("testdata", "ack-read.ach"))
	if err != nil {
		log.Fatal(err)
	}
	r := ach.NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		fmt.Printf("Issue reading file: %+v \n", err)
	}

	if err := achFile.Validate(); err != nil {
		fmt.Printf("Could not validate entire read file: %v", err)
	}

	if err := achFile.Create(); err != nil {
		fmt.Printf("Could not create file with read properties: %v", err)
	}

	fmt.Printf("Credit Total Amount: %v", strconv.Itoa(achFile.Control.TotalCreditEntryDollarAmountInFile)+"\n")
	fmt.Printf("Batch Credit Total Amount: %v", strconv.Itoa(achFile.Batches[0].GetEntries()[0].Amount)+"\n")
	fmt.Printf("SEC Code: %v", achFile.Batches[0].GetHeader().StandardEntryClassCode+"\n")
	fmt.Printf("Original Trace Number: %v", achFile.Batches[0].GetEntries()[0].OriginalTraceNumberField())

	// Output:
	// Credit Total Amount: 0
	// Batch Credit Total Amount: 0
	// SEC Code: ACK
	// Original Trace Number: 031300010000001
}
