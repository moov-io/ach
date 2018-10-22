// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"
)

func TestValidators__checkDigit(t *testing.T) {
	cases := map[string]int{
		// invalid
		"":       -1,
		"123456": -1,
		// valid
		"07300022": 8, // Wells Fargo - Iowa
		"10200007": 6, // Wells Fargo - Colorado
	}

	v := validator{}
	for rtn, check := range cases {
		answer := v.CalculateCheckDigit(rtn)
		if check != answer {
			t.Errorf("input=%s answer=%d expected=%d", rtn, answer, check)
		}
	}
}
