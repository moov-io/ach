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
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/ach"
	"github.com/stretchr/testify/require"
)

func TestIssue1275(t *testing.T) {
	file, err := ach.ReadFile(filepath.Join("..", "..", "test", "testdata", "iat-mixedDebitCredit.ach"))
	require.NoError(t, err)

	b1 := file.IATBatches[0]
	entries := b1.GetEntries()

	// Remove some batches to force new lines
	entries[0].Addenda11 = nil
	entries[0].Addenda12 = nil
	entries[0].Addenda14 = nil

	expectedLines := []string{
		`101 121042882 2313801041908071540A094101Bank                   My Bank Name                   `,
		`5200                FF3               US123456789 IATTRADEPAYMTCADUSD190808   0231380100000001`,
		`6271210428820007             0000100000123456789                              1231380100000001`,
		`710ANN000000000000100000928383-23938          BEK Enterprises                          0000001`,
		`713Wells Fargo                        01231380104                         US           0000001`,
		`7159874654932139872121 Front Street                                                    0000001`,
		`716LetterTown*AB\                     CA*80014\                                        0000001`, // [6]
		`717This is an international payment                                                00020000001`,
		`717This is an international payment                                                00020000001`,
		`718Bank of Fr` + "`" + `nce                     01456456456987987                   FR       00010000001`,
		`6221210428820007             0000100000123456789                              1231380100000002`,
		`710ANN000000000000100000928383-23938          ADCAF Enterprises                        0000002`,
		`711ADCAF Solutions                    15 West Place Street                             0000002`,
		`712JacobsTown*PA\                     US*19305\                                        0000002`,
		`713Wells Fargo                        01231380104                         US           0000002`,
		`714Citadel Bank                       01121042882                         CA           0000002`,
		`71598746549321398718 Fifth Street                                                      0000002`,
		`716LetterTown*AB\                     CA*80014\                                        0000002`,
		`718Bank of Fr` + "`" + `nce                     01456456456987987                   FR       00010000002`,
		`82000000200024208576000000100000000000100000                                   231380100000001`,
		`9000001000003000000200024208576000000100000000000100000                                       `,
	}

	t.Run("with padding", func(t *testing.T) {
		var buf bytes.Buffer
		err = ach.NewWriter(&buf).Write(file)
		require.NoError(t, err)

		lines := expectedLines
		for i := 0; i < 9; i++ {
			lines = append(lines, strings.Repeat("9", 94))
		}
		expected := strings.Join(append(lines, ""), "\n")
		require.Equal(t, expected, buf.String())
	})

	t.Run("no padding", func(t *testing.T) {
		entries[0].Addenda16 = nil
		lines := append(expectedLines[:6], expectedLines[7:]...)

		var buf bytes.Buffer
		err = ach.NewWriter(&buf).Write(file)
		require.NoError(t, err)

		expected := strings.Join(append(lines, ""), "\n")
		require.Equal(t, expected, buf.String())
	})
}
