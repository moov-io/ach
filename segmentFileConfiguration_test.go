// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import "testing"

// mockSegmentFileConfiguration creates a Segment File Configuration
func mockSegmentFileConfiguration() *SegmentFileConfiguration {
	sfc := NewSegmentFileConfiguration()
	return sfc
}

// TestSegmentFileConfiguration validates Segment File Configuration
func TestSegmentFileConfiguration(t *testing.T) {
	sfc := mockSegmentFileConfiguration()
	if sfc == nil {
		t.Error("mockSegmentFileConfiguration does not validate and will break other tests")
	}
}
