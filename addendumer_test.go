// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestAddendumer02__json(t *testing.T) {
	bs, err := json.Marshal(mockAddenda02())
	if err != nil {
		t.Fatal(err)
	}

	var add addenda
	if err := json.NewDecoder(bytes.NewReader(bs)).Decode(&add); err != nil {
		t.Fatal(err)
	}
	if err := add.Addenda.Validate(); err != nil {
		t.Error(err)
	}
}

func TestAddendumer05__json(t *testing.T) {
	bs, err := json.Marshal(mockAddenda05())
	if err != nil {
		t.Fatal(err)
	}

	var add addenda
	if err := json.NewDecoder(bytes.NewReader(bs)).Decode(&add); err != nil {
		t.Fatal(err)
	}
	if err := add.Addenda.Validate(); err != nil {
		t.Error(err)
	}
}

func TestAddendumer17__json(t *testing.T) {
	bs, err := json.Marshal(mockAddenda17())
	if err != nil {
		t.Fatal(err)
	}

	var add addenda
	if err := json.NewDecoder(bytes.NewReader(bs)).Decode(&add); err != nil {
		t.Fatal(err)
	}
	if err := add.Addenda.Validate(); err != nil {
		t.Error(err)
	}
}

func TestAddendumer18__json(t *testing.T) {
	bs, err := json.Marshal(mockAddenda18())
	if err != nil {
		t.Fatal(err)
	}

	var add addenda
	if err := json.NewDecoder(bytes.NewReader(bs)).Decode(&add); err != nil {
		t.Fatal(err)
	}
	if err := add.Addenda.Validate(); err != nil {
		t.Error(err)
	}
}

func TestAddendumer98__json(t *testing.T) {
	bs, err := json.Marshal(mockAddenda98())
	if err != nil {
		t.Fatal(err)
	}

	var add addenda
	if err := json.NewDecoder(bytes.NewReader(bs)).Decode(&add); err != nil {
		t.Fatal(err)
	}
	if err := add.Addenda.Validate(); err != nil {
		t.Error(err)
	}
}

func TestAddendumer99__json(t *testing.T) {
	bs, err := json.Marshal(mockAddenda99())
	if err != nil {
		t.Fatal(err)
	}

	var add addenda
	if err := json.NewDecoder(bytes.NewReader(bs)).Decode(&add); err != nil {
		t.Fatal(err)
	}
	if err := add.Addenda.Validate(); err != nil {
		t.Error(err)
	}
}
