package bench

import (
	"io"
	"path/filepath"
	"testing"

	"github.com/moov-io/ach"
)

func BenchmarkFile(b *testing.B) {
	file, err := ach.ReadFile(filepath.Join("..", "testdata", "20110805A.ach"))
	if err != nil {
		b.Fatal(err)
	}
	err = file.Create()
	if err != nil {
		b.Fatal(err)
	}

	b.Run("String", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			err := ach.NewWriter(io.Discard).Write(file)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
