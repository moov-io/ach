// Copyright 2016 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// Package ach reads and writes (ACH) Automated Clearing House files. ACH is the
// primary method of electronic money movemenet through the United States.
//
// https://en.wikipedia.org/wiki/Automated_Clearing_House
package ach

// First position of all Record Types. These codes are uniquily assigned to
// the first byte of each row in a file.
const (
	headerPos       = "1"
	batchPos        = "5"
	entryDetailPos  = "6"
	entryAddendaPos = "7"
	batchControlPos = "8"
	fileControlPos  = "9"
)

// File contains the structures of a parsed ACH File.
type File struct {
	FileHeader
	BatchHeader  BatchHeaderRecord
	EntryDetail  EntryDetailRecord
	Addenda      AddendaRecord
	BatchControl BatchControlRecord
	FileControl  FileControlRecord
}
