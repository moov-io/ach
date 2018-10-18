// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package iso3166

import (
	"strings"
)

// Valid returns successful if code is a valid ISO 3166-1-alpha-2
// country code. Example: US
func Valid(code string) bool {
	_, ok := countryCodes[strings.ToUpper(code)]
	return ok
}
