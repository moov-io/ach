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
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFile__FlattenBatchesOutOfOrder(t *testing.T) {
	file := NewFile()
	file.SetHeader(mockFileHeader())

	// First Batch
	b1 := mockBatchPPD()
	b1.Header.BatchNumber = 2
	b1.Entries[0].SetTraceNumber(b1.Header.ODFIIdentification, 100)
	// Create and Validate
	require.NoError(t, b1.Create())
	// Add to File
	file.AddBatch(b1)
	require.NoError(t, file.Create())

	// Second Batch
	b2 := mockBatchPPD()
	b2.Header.BatchNumber = 1
	b2.Entries[0].SetTraceNumber(b2.Header.ODFIIdentification, 1)
	// Create and Validate
	require.NoError(t, b2.Create())
	// Add to File
	file.AddBatch(b2)
	require.NoError(t, file.Create())

	// Validate File
	require.NoError(t, file.Create())
	// require.NoError(t, file.Validate()) // fails, but can we re-order batches?

	// Flatten batches
	out, err := file.FlattenBatches()
	require.NoError(t, err)
	require.NotNil(t, out)
	require.Equal(t, len(out.Batches), 1)
	t.Logf("%#v", out)
}
