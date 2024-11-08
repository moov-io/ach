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
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/ach"
	"github.com/moov-io/ach/cmd/achcli/describe"
	"github.com/stretchr/testify/require"
)

func TestIssue1501(t *testing.T) {
	file, err := ach.ReadJSONFile(filepath.Join("testdata", "issue1501.json"))
	require.NoError(t, err)

	if testing.Verbose() {
		describe.File(os.Stdout, file, nil)
	}

	require.Len(t, file.Batches, 1)

	entries := file.Batches[0].GetEntries()
	require.Len(t, entries, 2)

	e1 := entries[0]
	require.Equal(t, "0001", e1.CATXAddendaRecordsField())
	require.Equal(t, "Test              ", e1.CATXReceivingCompanyField())

	e2 := entries[1]
	require.Equal(t, "0001", e2.CATXAddendaRecordsField())
	require.Equal(t, "Other", e2.CATXReceivingCompanyField())
}
