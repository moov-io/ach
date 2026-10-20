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

package issues

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/ach"

	"github.com/stretchr/testify/require"
)

func TestIssue1891(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("..", "..", "test", "testdata", "iat-debit.ach"))
	require.NoError(t, err)

	bs = bytes.Replace(bs, []byte("710ANN"), []byte("710PAY"), 1)

	file, err := ach.NewReader(bytes.NewReader(bs)).Read()
	require.NoError(t, err)
	require.NoError(t, file.Validate())

	require.Len(t, file.IATBatches, 1)
	entries := file.IATBatches[0].GetEntries()
	require.NotEmpty(t, entries)
	require.Equal(t, "PAY", entries[0].Addenda10.TransactionTypeCode)
}
