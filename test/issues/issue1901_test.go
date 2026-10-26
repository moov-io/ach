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
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/moov-io/ach"

	"github.com/stretchr/testify/require"
)

func TestIssue1901(t *testing.T) {
	file, err := ach.ReadJSONFile(filepath.Join("testdata", "issue1901.json"))
	require.NoError(t, err)

	require.Len(t, file.Batches, 1)
	entries := file.Batches[0].GetEntries()
	require.Len(t, entries, 1)

	e := entries[0]
	require.Equal(t, "0000", e.CATXAddendaRecordsField())
	require.Equal(t, "GLOBAL LOGISTICS", e.CATXReceivingCompanyField())
	require.Equal(t, "0000GLOBAL LOGISTICS", e.IndividualName)

	bs, err := json.Marshal(file)
	require.NoError(t, err)
	file, err = ach.FileFromJSON(bs)
	require.NoError(t, err)
	e = file.Batches[0].GetEntries()[0]
	require.Equal(t, "0000", e.CATXAddendaRecordsField())
	require.Equal(t, "GLOBAL LOGISTICS", e.CATXReceivingCompanyField())

	var buf bytes.Buffer
	w := ach.NewWriter(&buf)
	require.NoError(t, w.Write(file))
	w.Flush()

	got, err := ach.NewReader(&buf).Read()
	require.NoError(t, err)
	e = got.Batches[0].GetEntries()[0]
	require.Equal(t, "0000", e.CATXAddendaRecordsField())
	require.Equal(t, "GLOBAL LOGISTICS  ", e.CATXReceivingCompanyField())
}
