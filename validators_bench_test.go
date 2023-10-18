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
	"math/rand"
	"testing"
)

func BenchmarkCalculateCheckDigit(b *testing.B) {
	var routingNumbers = []string{
		"053902197", "067014822", "273976369", "026009593", "071101307", "322271627",
		"121000358", "028000121", "054001725", "011401533", "322271724", "121042882",
		"021000021", "063100277", "071000013", "314074269", "063107513", "111000025",
		"031201360", "031176110", "991221943", "994489520", "724524972", "726286823",
	}

	aba := func(b *testing.B) string {
		idx := rand.Int() % len(routingNumbers) //nolint:gosec
		out := routingNumbers[idx]
		return out
	}

	b.Run("valid", func(b *testing.B) {
		num := aba(b)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			got := CalculateCheckDigit(num)

			last := int(num[8] - '0')
			if last != got {
				b.Errorf("ABA=%v got=%v expected=%v", num, got, last)
			}
		}
	})
}
