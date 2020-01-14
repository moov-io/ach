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
