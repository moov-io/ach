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

package ach

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFlattenFile(t *testing.T) {
	// Represent a file as groupings of trace numbers (or, hackily, amounts for ADV entries)
	getFileStructure := func(file *File) [][]string {
		var fileStructure [][]string
		for _, batcher := range file.Batches {
			var batchStructure []string
			for _, entry := range batcher.GetEntries() {
				batchStructure = append(batchStructure, entry.TraceNumber)
			}
			for _, entry := range batcher.GetADVEntries() {
				batchStructure = append(batchStructure, fmt.Sprintf("%d", entry.Amount))
			}
			fileStructure = append(fileStructure, batchStructure)
		}
		for _, batch := range file.IATBatches {
			var batchStructure []string
			for _, entry := range batch.Entries {
				batchStructure = append(batchStructure, entry.TraceNumber)
			}
			fileStructure = append(fileStructure, batchStructure)
		}

		return fileStructure
	}

	testCases := []struct {
		label string

		inputFixtureName string

		inputStructure     [][]string
		flattenedStructure [][]string
	}{
		{
			label:            "Single batch",
			inputFixtureName: "flattenBatchesOneBatchHeader.ach",
			inputStructure: [][]string{
				{"121042880000001", "121042880000002", "121042880000003"},
			},
			flattenedStructure: [][]string{
				{"121042880000001", "121042880000002", "121042880000003"},
			},
		},
		{
			label:            "Multiple batches (no trace collision)",
			inputFixtureName: "flattenBatchesMultipleBatchHeaders.ach",
			inputStructure: [][]string{
				{"121042880000001", "121042880000002", "121042880000003"},
				{"121042880000004", "121042880000005", "121042880000006"},
				{"121042880000007", "121042880000008", "121042880000009"},
				{"121042880000010", "121042880000011", "121042880000012"},
			},
			flattenedStructure: [][]string{
				{"121042880000001", "121042880000002", "121042880000003", "121042880000004", "121042880000005", "121042880000006"},
				{"121042880000007", "121042880000008", "121042880000009"},
				{"121042880000010", "121042880000011", "121042880000012"},
			},
		},
		{
			label:            "Trace number collision",
			inputFixtureName: "flattenBatchesTraceNumberCollision.ach",
			inputStructure: [][]string{
				{"121042880000001", "121042880000002", "121042880000022"},
				{"121042880000001", "121042880000002", "121042880000003"},
				{"121042880000004", "121042880000005", "121042880000006"},
				{"121042880000007", "121042880000008", "121042880000009"},
				{"121042880000010", "121042880000011", "121042880000012"},
			},
			flattenedStructure: [][]string{
				{"121042880000001", "121042880000002", "121042880000004", "121042880000005", "121042880000006", "121042880000022"},
				{"121042880000001", "121042880000002", "121042880000003"},
				{"121042880000007", "121042880000008", "121042880000009"},
				{"121042880000010", "121042880000011", "121042880000012"},
			},
		},
		{
			label:            "IAT batches",
			inputFixtureName: "flattenIATBatchesMultipleBatchHeaders.ach",
			inputStructure: [][]string{
				{"231380100000001", "231380100000002", "231380100000003"},
				{"231380100000004", "231380100000005", "231380100000006"},
				{"231380100000007", "231380100000008", "231380100000009"},
				{"231380100000010", "231380100000011", "231380100000012"},
			},
			flattenedStructure: [][]string{
				{"231380100000001", "231380100000002", "231380100000003", "231380100000004", "231380100000005", "231380100000006"},
				{"231380100000007", "231380100000008", "231380100000009", "231380100000010", "231380100000011", "231380100000012"},
			},
		},
		{
			label:            "ADV batches",
			inputFixtureName: "flattenADVBatchesMultipleBatchHeaders.ach",
			inputStructure: [][]string{
				{"50001", "50002", "50003"},
				{"50004", "50005", "50006"},
				{"50007", "50008", "50009"},
				{"50010", "50011", "50012"},
			},
			flattenedStructure: [][]string{
				{"50001", "50002", "50003", "50004", "50005", "50006"},
				{"50007", "50008", "50009", "50010", "50011", "50012"},
			},
		},
		{
			label:            "Micro deposits",
			inputFixtureName: "two-micro-deposits.ach",
			inputStructure: [][]string{
				{"121042886829038", "121042886829039", "121042886829040"},
				{"121042889211556", "121042889211557", "121042889211558"},
			},
			flattenedStructure: [][]string{
				{"121042886829038", "121042886829039", "121042886829040", "121042889211556", "121042889211557", "121042889211558"},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.label, func(t *testing.T) {
			inputRaw, _ := os.Open(filepath.Join("test", "testdata", testCase.inputFixtureName))
			inputFile, _ := NewReader(inputRaw).Read()

			flattenedFile, err := Flatten(&inputFile)
			if err != nil {
				require.NoError(t, err)
			}

			inputFileStructureJson, _ := json.MarshalIndent(getFileStructure(&inputFile), "", "  ")
			flattenedFileStructureJson, _ := json.MarshalIndent(getFileStructure(flattenedFile), "", "  ")
			expectedInputFileStructureJson, _ := json.MarshalIndent(testCase.inputStructure, "", "  ")
			expectedFlattenedFileStructureJson, _ := json.MarshalIndent(testCase.flattenedStructure, "", "  ")

			if string(inputFileStructureJson) != string(expectedInputFileStructureJson) {
				t.Fatalf(
					"Input file structure doesn't match expectation. \n\nActual:\n%s\n\nExpected:\n%s",
					string(inputFileStructureJson),
					string(expectedInputFileStructureJson),
				)
			}

			if string(flattenedFileStructureJson) != string(expectedFlattenedFileStructureJson) {
				t.Fatalf(
					"Flattened file structure doesn't match expectation. \n\nActual:\n%s\n\nExpected:\n%s",
					string(flattenedFileStructureJson),
					string(expectedFlattenedFileStructureJson),
				)
			}

			// Ensure that every entry from the input file is present in the flattened file, and is associated with the same batch header
			getEntrySignatures := func(file File) []string {
				var entrySignatureSet []string
				for _, batcher := range file.Batches {
					for _, entry := range batcher.GetEntries() {
						entrySignatureSet = append(entrySignatureSet, fmt.Sprintf("%s|%s", batcher.GetHeader().String()[4:87], entry.String()))
					}
					for _, entry := range batcher.GetADVEntries() {
						entrySignatureSet = append(entrySignatureSet, fmt.Sprintf("%s|%s", batcher.GetHeader().String()[4:87], entry.String()[:90]))
					}
				}
				for _, batch := range file.IATBatches {
					for _, entry := range batch.Entries {
						entrySignatureSet = append(entrySignatureSet, fmt.Sprintf("%s|%s", batch.Header.String()[4:87], entry.String()))
					}
				}

				sort.Strings(entrySignatureSet)
				return entrySignatureSet
			}

			if strings.Join(getEntrySignatures(inputFile), ", ") != strings.Join(getEntrySignatures(*flattenedFile), ", ") {
				t.Fatal("Discrepancy between input and flattened file entry signatures")
			}
		})
	}
}
