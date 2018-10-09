// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

// Addendumer abstracts the different ACH addendum types that can be added to an EntryDetail record
type Addendumer interface {
	Parse(string)
	//TypeCode Defines the specific explanation and format for the addenda information
	typeCode() string
	String() string
	Validate() error
}
