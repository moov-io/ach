package read

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/moov-io/ach"
)

type Format string

var (
	FormatUnknown Format = "unknown"
	FormatNacha   Format = "nacha"
	FormatJSON    Format = "json"
)

func Filepath(path string, validateOptsPath *string, skipAll *bool) (*ach.File, Format, error) {
	validateOpts := readValidationOpts(validateOptsPath, skipAll)

	return readFile(path, validateOpts)
}

func readValidationOpts(path *string, skipAll *bool) *ach.ValidateOpts {
	var opts ach.ValidateOpts

	if skipAll != nil && *skipAll {
		opts.SkipAll = true
		return &opts
	}

	if path != nil && *path != "" {
		// read config file
		bs, readErr := os.ReadFile(*path)
		if readErr != nil {
			fmt.Printf("ERROR: reading %s for validate opts failed: %v\n", *path, readErr)
			os.Exit(1)
		}

		if err := json.Unmarshal(bs, &opts); err != nil {
			fmt.Printf("ERROR: unmarshal of validate opts failed: %v\n", err)
			os.Exit(1)
		}
		return &opts
	}
	return nil
}

func readFile(path string, validateOpts *ach.ValidateOpts) (*ach.File, Format, error) {
	bs, err := os.ReadFile(path)
	if err != nil {
		return nil, FormatUnknown, err
	}
	if json.Valid(bs) {
		return readJsonFile(bs, validateOpts)
	}
	return readACHFile(bs, validateOpts)
}

func readACHFile(input []byte, validateOpts *ach.ValidateOpts) (*ach.File, Format, error) {
	r := ach.NewReader(bytes.NewReader(input))
	r.SetValidation(validateOpts)

	f, err := r.Read()

	return &f, FormatNacha, err
}

func readJsonFile(input []byte, validateOpts *ach.ValidateOpts) (*ach.File, Format, error) {
	file, err := ach.FileFromJSONWith(input, validateOpts)
	return file, FormatJSON, err
}
