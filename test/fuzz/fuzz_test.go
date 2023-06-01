package main

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/ach"
)

func FuzzReaderWriterACH(f *testing.F) {
	populateCorpus(f, true)

	f.Fuzz(func(t *testing.T, contents string) {
		// Read the sample file
		r := ach.NewReader(strings.NewReader(contents))
		r.SetValidation(&ach.ValidateOpts{
			SkipAll: true,
		})
		file, err := r.Read()
		if err != nil {
			t.Skip()
		}

		// Write the file
		ach.NewWriter(io.Discard).Write(&file)

		// Remove Validation override
		file.SetValidation(&ach.ValidateOpts{
			SkipAll: false,
		})
		file.Validate()
	})
}

func FuzzReaderWriterJSON(f *testing.F) {
	populateCorpus(f, false)

	f.Fuzz(func(t *testing.T, contents string) {
		file, err := ach.FileFromJSONWith([]byte(contents), &ach.ValidateOpts{
			SkipAll: true,
		})
		if err != nil {
			t.Skip()
		}

		// Write the file
		file.MarshalJSON()

		// Remove Validation override
		file.SetValidation(&ach.ValidateOpts{
			SkipAll: false,
		})
		file.Validate()
	})
}

func populateCorpus(f *testing.F, ach bool) {
	f.Helper()

	err := filepath.Walk(filepath.Join("..", ".."), func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(strings.ToLower(path))
		if (ext == ".ach" && ach) || (ext == ".json" && !ach) {
			bs, err := os.ReadFile(path)
			if err != nil {
				f.Fatal(err)
			}
			f.Add(string(bs))
		}
		return nil
	})
	if err != nil {
		f.Fatal(err)
	}
}
