// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package server

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
)

// NextID generates a new resource ID.
// Do not assume anything about the data structure.
//
// Multiple calls to NextID() have no concern about producing
// lexicographically ordered output.
func NextID() string {
	bs := make([]byte, 20)
	rand.Reader.Read(bs)

	h := sha1.New()
	h.Write(bs)
	return hex.EncodeToString(h.Sum(nil))[:16]
}
