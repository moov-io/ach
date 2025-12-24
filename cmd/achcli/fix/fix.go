package fix

import (
	"bytes"
	"fmt"
	"os"

	"github.com/moov-io/ach"
	"github.com/moov-io/ach/cmd/achcli/internal/read"
	"github.com/moov-io/ach/cmd/achcli/internal/write"
)

func Perform(path string, validateOptsPath *string, skipAll *bool, conf Config) (string, error) {
	file, format, err := read.Filepath(path, validateOptsPath, skipAll)
	if err != nil {
		return "", fmt.Errorf("reading %s failed: %w", path, err)
	}

	// Build up our fixers
	var batchHeaderFixers []batchHeaderFixer
	if conf.UpdateEED != "" {
		batchHeaderFixers = append(batchHeaderFixers, updateEED(conf))
	}

	// Fix the file
	for idx := range file.Batches {
		// Batch headers
		bh := file.Batches[idx].GetHeader()
		for _, fn := range batchHeaderFixers {
			if err := fn(bh); err != nil {
				return "", fmt.Errorf("applying %T to batch header: %w", fn, err)
			}
		}
		file.Batches[idx].SetHeader(bh)
	}

	// Write file
	newpath := path + ".fix"

	var buf bytes.Buffer
	err = write.File(&buf, file, format)
	if err != nil {
		return "", fmt.Errorf("encoding fixed file as %s: %w", format, err)
	}

	err = os.WriteFile(newpath, buf.Bytes(), 0600)
	if err != nil {
		return "", fmt.Errorf("writing %s failed: %w", newpath, err)
	}

	return newpath, nil
}

type batchHeaderFixer func(bh *ach.BatchHeader) error
