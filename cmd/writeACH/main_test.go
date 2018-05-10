package main

import (
	"testing"
)

// TestFileCreate tests creating an ACH File
func TestFileWrite(t *testing.T) {
	FileWrite(t)
}

//BenchmarkTestFileCreate benchmarks creating an ACH File
func BenchmarkTestFileWrite(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FileWrite(b)
	}
}

// FileCreate creates an ACH File
func FileWrite(t testing.TB) {
	main()
}
