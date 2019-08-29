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
