package fix_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/ach"
	"github.com/moov-io/ach/cmd/achcli/fix"

	"github.com/stretchr/testify/require"
)

func TestPerform(t *testing.T) {
	cases := []struct {
		inputFilepath    string
		validateOpts     *ach.ValidateOpts
		config           fix.Config
		expectedFilepath string
	}{
		{
			inputFilepath: filepath.Join("..", "..", "..", "test", "testdata", "ppd-debit.ach"),
			config: fix.Config{
				UpdateEED: "20251224",
			},
			expectedFilepath: filepath.Join("..", "..", "..", "test", "testdata", "ppd-debit.ach.fix.expected"),
		},
	}
	for _, tc := range cases {
		_, filename := filepath.Split(tc.inputFilepath)

		t.Run(filename, func(t *testing.T) {
			newpath, err := fix.Perform(tc.inputFilepath, tc.validateOpts, tc.config)
			require.NoError(t, err)

			got, err := os.ReadFile(newpath)
			require.NoError(t, err)

			expected, err := os.ReadFile(tc.expectedFilepath)
			require.NoError(t, err)

			got = bytes.TrimSpace(got)
			expected = bytes.TrimSpace(expected)

			require.Equal(t, string(expected), string(got))
		})
	}
}
