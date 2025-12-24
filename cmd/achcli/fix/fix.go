package fix

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/moov-io/ach"
)

func Perform(path string, validateOpts *ach.ValidateOpts, conf Config) (string, error) {
	file, err := readFile(path, validateOpts)
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
	// TODO(adam): write back as json if we accepted json
	newpath := path + ".fix"

	var buf bytes.Buffer
	err = ach.NewWriter(&buf).Write(file)
	if err != nil {
		return "", fmt.Errorf("encoding fixed file: %w", err)
	}

	err = os.WriteFile(newpath, buf.Bytes(), 0644)
	if err != nil {
		return "", fmt.Errorf("writing %s failed: %w", newpath, err)
	}

	return newpath, nil
}

func readFile(path string, validateOpts *ach.ValidateOpts) (*ach.File, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if json.Valid(bs) {
		return readJsonFile(bs, validateOpts)
	}
	return readACHFile(bs, validateOpts)
}

func readACHFile(input []byte, validateOpts *ach.ValidateOpts) (*ach.File, error) {
	r := ach.NewReader(bytes.NewReader(input))
	r.SetValidation(validateOpts)
	f, err := r.Read()
	return &f, err
}

func readJsonFile(input []byte, validateOpts *ach.ValidateOpts) (*ach.File, error) {
	return ach.FileFromJSONWith(input, validateOpts)
}

type batchHeaderFixer func(bh *ach.BatchHeader) error

// TODO(adam): process file batch by batch
// TODO(adam): fix funcs take a file header, batch, entry, etc

// TODO(adam): updateEED()

// TODO(adam): write output to disk
