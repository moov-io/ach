package main

import (
	"testing"
)

// TestFileCreate tests creating an ACH File
func TestFileWrite(t *testing.T) {
	testFileWrite(t)
}

/*//BenchmarkTestFileCreate benchmarks creating an ACH File
func BenchmarkTestFileWrite(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileWrite(b)
	}
}*/

// FileCreate creates an ACH File
func testFileWrite(t testing.TB) {
	main()
}
