// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package examples

import (
	"github.com/moov-io/ach"
)

func mockFileHeader() ach.FileHeader {
	fh := ach.NewFileHeader()
	fh.ImmediateDestination = "031300012"
	fh.ImmediateOrigin = "231380104"
	// need FileCreationDate and FileCreationTime to be fixed so it can match output
	fh.FileCreationDate = "190816"
	fh.FileCreationTime = "1055"
	fh.ImmediateDestinationName = "Federal Reserve Bank"
	fh.ImmediateOriginName = "My Bank Name"
	fh.ReferenceCode = "12345678"
	return fh
}
