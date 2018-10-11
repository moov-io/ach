// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package iso4217

import (
	"testing"
)

func TestValidate(t *testing.T) {
	if !Valid("USD") {
		t.Error("expected valid")
	}

	if !Valid("eur") {
		t.Error("expected valid")
	}

	if Valid("") {
		t.Errorf("invalid")
	}

	if Valid("QZA") {
		t.Errorf("invalid")
	}
}
