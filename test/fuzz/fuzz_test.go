package main

import (
	"bytes"
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
		if len(contents) > 1<<20 {
			t.Skip()
		}

		// Read the sample file with validation skipped so more inputs parse.
		r := ach.NewReader(strings.NewReader(contents))
		r.SetValidation(&ach.ValidateOpts{
			SkipAll: true,
		})
		file, err := r.Read()
		if err != nil {
			return
		}

		// Write the file
		var buf bytes.Buffer
		if err := ach.NewWriter(&buf).Write(&file); err != nil {
			// Write errors are fine; panics are not.
			_ = ach.NewWriter(io.Discard).Write(&file)
		}

		// Re-read what we wrote (round trip)
		if buf.Len() > 0 {
			r2 := ach.NewReader(bytes.NewReader(buf.Bytes()))
			r2.SetValidation(&ach.ValidateOpts{SkipAll: true})
			_, _ = r2.Read()
		}

		// Validate without SkipAll — must not panic
		file.SetValidation(&ach.ValidateOpts{SkipAll: false})
		_ = file.Validate()
	})
}

func FuzzReaderWriterJSON(f *testing.F) {
	populateCorpus(f, false)

	f.Fuzz(func(t *testing.T, contents string) {
		if len(contents) > 1<<20 {
			t.Skip()
		}

		file, err := ach.FileFromJSONWith([]byte(contents), &ach.ValidateOpts{
			SkipAll: true,
		})
		if err != nil || file == nil {
			return
		}

		_, _ = file.MarshalJSON()

		file.SetValidation(&ach.ValidateOpts{SkipAll: false})
		_ = file.Validate()

		_ = ach.NewWriter(io.Discard).Write(file)
	})
}

func populateCorpus(f *testing.F, achFiles bool) {
	f.Helper()

	f.Add("")
	f.Add("{}")

	err := filepath.Walk(filepath.Join("..", ".."), func(path string, info fs.FileInfo, err error) error {
		if err != nil || info == nil || info.IsDir() {
			return nil
		}
		// Skip generated fuzz cache and vendor-like trees
		if strings.Contains(path, string(filepath.Separator)+"testdata"+string(filepath.Separator)+"fuzz") {
			return nil
		}
		if strings.Contains(path, ".git") {
			return nil
		}

		ext := filepath.Ext(strings.ToLower(path))
		if (ext == ".ach" && achFiles) || (ext == ".json" && !achFiles) {
			bs, err := os.ReadFile(path)
			if err != nil {
				return nil
			}
			if len(bs) > 256*1024 {
				return nil
			}
			f.Add(string(bs))
		}
		return nil
	})
	if err != nil {
		f.Fatal(err)
	}
}
