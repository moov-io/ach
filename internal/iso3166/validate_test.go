// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package iso3166

import (
	"testing"
)

func TestValidate(t *testing.T) {
	if !Valid("US") {
		t.Error("expected valid")
	}

	if !Valid("SS") {
		t.Error("expected valid")
	}

	if Valid("") {
		t.Errorf("invalid")
	}

	if Valid("QZ") {
		t.Errorf("invalid")
	}
}
