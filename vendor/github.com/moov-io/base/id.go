// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package base

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

// ID creates a new random string for Moov systems.
// Do not assume anything about these ID's other than they are non-empty strings.
func ID() string {
	// NOTE(adam): Moov's apps depend on the length and hex encoding of these ID's to cleanup HTTP Prometheus metrics.
	bs := make([]byte, 20)
	n, err := rand.Read(bs)
	if err != nil || n == 0 {
		return ""
	}
	return strings.ToLower(hex.EncodeToString(bs))
}
