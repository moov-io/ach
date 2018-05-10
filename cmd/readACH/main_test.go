package main

import "testing"

func TestFileRead(t *testing.T) {
	FileRead(t)
}

//BenchmarkTestFileCreate benchmarks creating an ACH File
func BenchmarkTestFileRead(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		FileRead(b)
	}
}

// FileCreate creates an ACH File
func FileRead(t testing.TB) {
	main()
}
