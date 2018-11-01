// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"encoding/json"
)

// Addendumer abstracts the different ACH addendum types that can be added to an EntryDetail record
type Addendumer interface {
	Parse(string)
	//TypeCode Defines the specific explanation and format for the addenda information
	typeCode() string
	String() string
	Validate() error
}

type addenda struct {
	Addenda Addendumer
}

// UnmarshalJSON attempts to parse data as each AddendaXX record.
func (a *addenda) UnmarshalJSON(data []byte) error {
	addenda02 := NewAddenda02()
	if json.Unmarshal(data, &addenda02) == nil && addenda02.Validate() == nil {
		a.Addenda = addenda02
		return nil
	}

	addenda05 := NewAddenda05()
	if json.Unmarshal(data, &addenda05) == nil && addenda05.Validate() == nil {
		a.Addenda = addenda05
		return nil
	}

	addenda10 := NewAddenda10()
	if json.Unmarshal(data, &addenda10) == nil && addenda10.Validate() == nil {
		a.Addenda = addenda10
		return nil
	}

	addenda12 := NewAddenda12()
	if json.Unmarshal(data, &addenda12) == nil && addenda12.Validate() == nil {
		a.Addenda = addenda12
		return nil
	}

	addenda13 := NewAddenda13()
	if json.Unmarshal(data, &addenda13) == nil && addenda13.Validate() == nil {
		a.Addenda = addenda13
		return nil
	}

	addenda14 := NewAddenda14()
	if json.Unmarshal(data, &addenda14) == nil && addenda14.Validate() == nil {
		a.Addenda = addenda14
		return nil
	}

	addenda15 := NewAddenda15()
	if json.Unmarshal(data, &addenda15) == nil && addenda15.Validate() == nil {
		a.Addenda = addenda15
		return nil
	}

	addenda16 := NewAddenda16()
	if json.Unmarshal(data, &addenda16) == nil && addenda16.Validate() == nil {
		a.Addenda = addenda16
		return nil
	}

	addenda17 := NewAddenda17()
	if json.Unmarshal(data, &addenda17) == nil && addenda17.Validate() == nil {
		a.Addenda = addenda17
		return nil
	}

	addenda18 := NewAddenda18()
	if json.Unmarshal(data, &addenda18) == nil && addenda18.Validate() == nil {
		a.Addenda = addenda18
		return nil
	}

	addenda98 := NewAddenda98()
	if json.Unmarshal(data, &addenda98) == nil && addenda98.Validate() == nil {
		a.Addenda = addenda98
		return nil
	}

	addenda99 := NewAddenda99()
	if json.Unmarshal(data, &addenda99) == nil && addenda99.Validate() == nil {
		a.Addenda = addenda99
		return nil
	}

	return nil
}
