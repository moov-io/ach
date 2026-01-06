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

package addenda

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func FuzzParseTXP(f *testing.F) {
	examples := []string{
		"123456789*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED\\",
		"TXP**941*250901*T*1000\\",
		"TXP*123#456*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED\\",
		"TXP*123-456*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED\\",
		"TXP*123.456*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED\\",
		"TXP*123456789**250901*T*1000\\",
		"TXP*123456789*606*960331*T*100000*P*12000*I*4567*SML2A\\",
		"TXP*123456789*941**T*1000\\",
		"TXP*123456789*941*2509*T*1000\\",
		"TXP*123456789*941*250901**1000\\",
		"TXP*123456789*941*250901*1*1000*2*500***VERIFIED\\",
		"TXP*123456789*941*250901*1*1000*2*500\\",
		"TXP*123456789*941*250901*T* 1000 \\",
		"TXP*123456789*941*250901*T*1000****VERIFIED\\",
		"TXP*123456789*941*250901*T*1000*VERIFIED\\",
		"TXP*123456789*941*250901*T*1000*\\",
		"TXP*123456789*941*250901*T*100000\\",
		"TXP*123456789*941*250901*T*\\",
		"TXP*123456789*941*2509011*T*1000\\",
		"TXP*123456789*941*ABCDEF*T*1000\\",
		"TXP*123456789*94105*960301*1*10000*2*5000*3*15000\\",
		"TXP*123456789*FEDERAL*20231231*500000*1*25000*VERIFIED\\",
		"TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED",
		"TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*1*25000*2**VERIFIED\\",
		"TXP*123456789*FEDERAL*20231231*TAX2023Q4*500000*1*25000*2*10000*VERIFIED_WITH_VERY_LONG_VERIFICATION_STRING\\",
		"TXP*123456789*FEDERAL*20231231*TAX2023Q4\\",
		"TXP*123456789*STATE*20231231*TAX2023Q4*500000*S**VERIFIED\\",
		"TXP*12345678934*526*960930*T*100000***SML2A\\",
		"TXP*123:456*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED\\",
		"TXP*123@456*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED\\",
		"TXP*123\n456*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED\\",
		"TXP*123\t456*FEDERAL*20231231*TAX2023Q4*500000*1**VERIFIED\\",
		"TXP*888999000*STATE*20231231*S*200000*2*10000**CONFIRMED\\",
		"txp*123456789*941*250901*T*1000\\",
	}

	for i := range examples {
		f.Add(examples[i])
	}

	f.Fuzz(func(t *testing.T, input string) {
		require.NotPanics(t, func() {
			rec, _ := ParseTXP(input)
			if rec != nil {
				rec.String()
			}
		})
	})
}
