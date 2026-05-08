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
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIterator_MalformedFiles(t *testing.T) {
	type batchExpectation struct {
		entries int
	}

	cases := []struct {
		path            string
		expectedBatches []batchExpectation
	}{
		{
			path: filepath.Join("test", "testdata", "20110729A-invalid.ach"),
			expectedBatches: []batchExpectation{
				{entries: 223},
			},
		},
		{
			path: filepath.Join("test", "testdata", "invalid-two-micro-deposits.ach"),
			expectedBatches: []batchExpectation{
				{entries: 3},
				{entries: 3},
			},
		},
		{
			path: filepath.Join("test", "testdata", "ppd-debit-invalid-entryDetail-checkDigit.ach"),
			expectedBatches: []batchExpectation{
				{entries: 1},
			},
		},
		{
			path: filepath.Join("test", "testdata", "ppd-mixedDebitCredit-invalid.ach"),
			expectedBatches: []batchExpectation{
				{entries: 2},
			},
		},
	}
	for _, tc := range cases {
		t.Run(filepath.Base(tc.path), func(t *testing.T) {
			iter := iteratorFromFile(t, tc.path, &ValidateOpts{
				SkipAll: true,
			})

			batchesFound := make([]string, 0)
			batchStats := make(map[string]int, 0)

			for {
				bh, entry, err := iter.NextEntry()
				if err != nil {
					t.Log(err)
				} else {
					if bh == nil && entry == nil {
						break // iterator finished
					}
				}

				if bh != nil && entry != nil {
					batchKey := bh.String()
					_, exists := batchStats[batchKey]
					if !exists {
						batchesFound = append(batchesFound, batchKey)
					}
					batchStats[batchKey]++
				}
			}

			// compare
			var expected []batchExpectation
			for _, key := range batchesFound {
				expected = append(expected, batchExpectation{
					entries: batchStats[key],
				})
			}
			// exact order and count
			require.Equal(t, expected, tc.expectedBatches)
		})
	}
}
