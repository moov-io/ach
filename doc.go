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

// Package ach reads and writes Automated Clearing House (ACH) files. ACH is the
// primary method of electronic money movement through the United States.
//
// https://en.wikipedia.org/wiki/Automated_Clearing_House
//
// https://moov-io.github.io/ach/
//
// Read an ACH File
//
//	fd, err := os.Open("name-of-your-ach-file.ach")
//	if err != nil {
//	    log.Fatalf("problem opening file: %v", err)
//	}
//	file, err := ach.NewReader(fd).Read()
//	if err != nil {
//	    log.Fatalf("problem parsing ACH file: %v", err)
//	}
//	if err := file.Validate(); err != nil {
//	    log.Fatalf("ACH file isn't valid: %v", err)
//	}
package ach
