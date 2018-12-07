// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import "testing"

func TestFileRead(t *testing.T) {
	testFileRead(t)
}

/*//BenchmarkTestFileCreate benchmarks creating an ACH File
func BenchmarkTestFileRead(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileRead(b)
	}
}*/

// FileCreate creates an ACH File
func testFileRead(t testing.TB) {
	main()
}
