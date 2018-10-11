// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package iso4217

import (
	"strings"
)

// Valid returns successful if code is a valid ISO 4217
// currency code. Example: USD
func Valid(code string) bool {
	_, ok := currencyCodes[strings.ToUpper(code)]
	return ok
}
