package bench

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/ach"
	"github.com/moov-io/base"
)

func BenchmarkParsing(b *testing.B) {
	b.ReportAllocs()

	files := []string{
		filepath.Join("..", "testdata", "20110805A.ach"),
	}
	opts := &ach.ValidateOpts{
		AllowMissingFileControl:    true,
		AllowMissingFileHeader:     true,
		AllowUnorderedBatchNumbers: true,
	}

	for i := range files {
		path := files[i]
		b.Run("Read_"+path, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fd, err := os.Open(path)
				if err != nil {
					b.Fatal(err)
				}

				r := ach.NewReader(fd)
				r.SetValidation(opts)
				file, err := r.Read()
				fd.Close()
				if err != nil {
					if !base.Has(err, ach.ErrFileHeader) {
						b.Fatal(err)
					}
				}
				if len(file.Batches) == 0 {
					b.Error("no batches read")
				}
			}
		})
		b.Run("Iterator_"+path, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fd, err := os.Open(path)
				if err != nil {
					b.Fatal(err)
				}

				iter := ach.NewIterator(fd)
				iter.SetValidation(opts)

				entries := 0
				for {
					bh, entry, err := iter.NextEntry()
					if err != nil {
						if !base.Has(err, ach.ErrFileHeader) {
							b.Fatal(err)
						}
					}
					if bh != nil && entry != nil {
						entries += 1
					}
					if bh == nil && entry == nil {
						break
					}
				}

				fd.Close()
			}
		})
	}
}
