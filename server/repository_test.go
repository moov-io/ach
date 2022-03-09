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

package server

import (
	"testing"
	"time"

	"github.com/moov-io/ach"
	"github.com/moov-io/base"
	"github.com/moov-io/base/log"
)

var (
	testTTLDuration = 0 * time.Second // disable TTL expiry
)

func TestRepositoryFiles(t *testing.T) {
	r := NewRepositoryInMemory(testTTLDuration, nil)

	if v := len(r.FindAllFiles()); v != 0 {
		t.Errorf("unexpected length: %d", v)
	}

	header := mockFileHeader()
	f := &ach.File{
		ID:     base.ID(),
		Header: *header,
	}
	if err := r.StoreFile(f); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	found, err := r.FindFile(f.ID)
	if err != nil || found == nil {
		t.Errorf("found=%v, err=%v", found, err)
	}

	if v := len(r.FindAllFiles()); v != 1 {
		t.Errorf("unexpected length: %d", v)
	}

	if err := r.DeleteFile(f.ID); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRepositoryBatches(t *testing.T) {
	r := NewRepositoryInMemory(testTTLDuration, nil)

	// make sure our tests are setup
	if v := len(r.FindAllFiles()); v != 0 {
		t.Errorf("unexpected length: %d", v)
	}

	header := mockFileHeader()
	f := &ach.File{
		ID:     base.ID(),
		Header: *header,
	}
	if err := r.StoreFile(f); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// batch tests
	if v := len(r.FindAllBatches(f.ID)); v != 0 {
		t.Errorf("unexpected length: %d", v)
	}

	batch := mockBatchWEB()
	b, err := r.FindBatch(f.ID, batch.ID())
	if err == nil || b != nil {
		t.Errorf("b=%v, err=%v", b, err)
	}

	if err := r.StoreBatch(f.ID, batch); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if v := len(r.FindAllBatches(f.ID)); v != 1 {
		t.Errorf("unexpected length: %d", v)
	}

	if err := r.DeleteBatch(f.ID, batch.ID()); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if v := len(r.FindAllBatches(f.ID)); v != 0 {
		t.Errorf("unexpected length: %d", v)
	}
}

func TestRepository__cleanupOldFiles(t *testing.T) {
	r := NewRepositoryInMemory(testTTLDuration, nil)
	if repo, ok := r.(*repositoryInMemory); !ok {
		t.Fatalf("unexpected repository: %T %#v", r, r)
	} else {
		// write a file and later verify it's cleaned up
		file := ach.NewFile()
		file.Header.FileCreationDate = time.Now().Add(-1 * 24 * time.Hour).Format("060102") // YYMMDD of 24hrs ago
		repo.StoreFile(file)
		if n := len(repo.FindAllFiles()); n != 1 {
			t.Errorf("got %d ACH files", n)
		}
		repo.cleanupOldFiles() // make sure we don't panic
		if n := len(repo.FindAllFiles()); n != 0 {
			t.Errorf("got %d ACH files", n)
		}
	}

	// Create a repo with a logger
	r = NewRepositoryInMemory(testTTLDuration, log.NewNopLogger())
	if repo, ok := r.(*repositoryInMemory); !ok {
		t.Fatalf("unexpected repository: %T %#v", r, r)
	} else {
		repo.cleanupOldFiles() // make sure we don't panic
	}
}
