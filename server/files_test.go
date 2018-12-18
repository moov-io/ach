// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package server

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/moov-io/ach"
	"github.com/moov-io/base"
)

func TestFiles__OldBatches(t *testing.T) {
	bs, err := ioutil.ReadFile(filepath.Join("..", "test", "testdata", "ppd-valid.json"))
	if err != nil {
		t.Fatal(err)
	}
	file, err := ach.FileFromJSON(bs)
	if err != nil {
		t.Fatal(err)
	}

	if err := file.Validate(); err != nil {
		t.Fatal(err)
	}

	// Set EffectiveEntryDate to some old date
	longAgo := base.Now().Add(-7 * 24 * time.Hour)
	file.Batches[0].GetHeader().EffectiveEntryDate = base.NewTime(longAgo)

	err = fileHasOldBatches(file)
	if err == nil {
		t.Error("expected error, but got nothing")
	}
	if err != nil && !strings.Contains(err.Error(), "EffectiveEntryDate before today") {
		t.Errorf("unexpected error: %v", err)
	}

	// Set the value for tomorrow and see it pass
	tomorrow := base.Now().Add(24 * time.Hour)
	file.Batches[0].GetHeader().EffectiveEntryDate = base.NewTime(tomorrow)
	if err := fileHasOldBatches(file); err != nil {
		t.Errorf("shouldn't have had error, but got %v", err)
	}
}
