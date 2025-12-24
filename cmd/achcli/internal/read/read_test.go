package read_test

import (
	"path/filepath"
	"testing"

	"github.com/moov-io/ach/cmd/achcli/internal/read"

	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	cases := []struct {
		inputFilepath    string
		validateOptsPath *string
		skipAll          *bool

		expectedFormat read.Format
	}{
		{
			inputFilepath:    filepath.Join("..", "..", "..", "..", "test", "testdata", "ppd-debit-invalid-entryDetail-checkDigit.ach"),
			validateOptsPath: ptr(filepath.Join("..", "..", "..", "..", "test", "testdata", "ppd-debit-invalid-entryDetail-checkDigit.json")),
			expectedFormat:   read.FormatNacha,
		},
		{
			inputFilepath:  filepath.Join("..", "..", "..", "..", "test", "testdata", "ppd-valid-preserve-spaces.json"),
			expectedFormat: read.FormatJSON,
		},
	}
	for _, tc := range cases {
		_, filename := filepath.Split(tc.inputFilepath)

		t.Run(filename, func(t *testing.T) {
			file, format, err := read.Filepath(tc.inputFilepath, tc.validateOptsPath, tc.skipAll)
			require.NoError(t, err)
			require.Equal(t, tc.expectedFormat, format)
			require.NotNil(t, file)
		})
	}
}

func ptr[T any](in T) *T { return &in }
