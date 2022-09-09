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

package issues

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/ach"
)

func TestPull713(t *testing.T) {
	// A vendor has these health files which contain zero batches and are used
	// as a sort of daily heartbeat for software.
	file, err := ach.ReadFile(filepath.Join("..", "testdata", "FISERV-ZEROFILE-PIMRET825324_032720_110221.ach"))
	if err == nil {
		t.Error("expected error")
	}
	if !strings.Contains(err.Error(), `ImmediateDestination 100067554 routing number checksum mismatch`) {
		t.Error(err)
	}

	fh := file.Header
	if fh.ImmediateDestination != "100067554" {
		t.Errorf("ImmediateDestination=%s", fh.ImmediateDestination)
	}
	if fh.ImmediateOrigin != "182327390" {
		t.Errorf("ImmediateOrigin=%s", fh.ImmediateOrigin)
	}
	if fh.ImmediateDestinationName != "PIMRET825324" {
		t.Errorf("ImmediateDestinationName=%s", fh.ImmediateDestinationName)
	}
	if fh.ImmediateOriginName != "FISERV" {
		t.Errorf("ImmediateOriginName=%s", fh.ImmediateOriginName)
	}

	if file.Control.BatchCount != 0 {
		t.Errorf("control batch count: %d", file.Control.BatchCount)
	}
}
