// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

// Package ach reads and writes Automated Clearing House (ACH) files. ACH is the
// primary method of electronic money movement through the United States.
//
// https://en.wikipedia.org/wiki/Automated_Clearing_House
//
// https://github.com/moov-io/ach/tree/master/documentation
//
// Read an ACH File
//     fd, err := os.Open("name-of-your-ach-file.ach")
//     if err != nil {
//         log.Fatalf("problem opening file: %v", err)
//     }
//     file, err := ach.NewReader(fd).Read()
//     if err != nil {
//         log.Fatalf("problem parsing ACH file: %v", err)
//     }
//     if err := file.Validate(); err != nil {
//         log.Fatalf("ACH file isn't valid: %v", err)
//     }
package ach
