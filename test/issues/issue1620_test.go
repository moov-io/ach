package issues

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/ach"

	"github.com/stretchr/testify/require"
)

func TestIssue1620(t *testing.T) {
	t.Run("Forward - no error", func(t *testing.T) {
		file, err := ach.ReadFile(filepath.Join("testdata", "issue1620.ach"))
		require.NoError(t, err)
		require.NotNil(t, file)

		require.NoError(t, file.Create())
		require.NoError(t, file.Validate())
	})

	t.Run("Return - no error", func(t *testing.T) {
		file, err := ach.ReadFile(filepath.Join("testdata", "issue1620_return.ach"))
		require.NoError(t, err)
		require.NotNil(t, file)

		require.NoError(t, file.Create())
		require.NoError(t, file.Validate())
	})

	t.Run("Return - AllowInvalidAmounts", func(t *testing.T) {
		fd, err := os.Open(filepath.Join("testdata", "issue1620.ach"))
		require.NoError(t, err)
		t.Cleanup(func() { fd.Close() })

		r := ach.NewReader(fd)
		r.SetValidation(&ach.ValidateOpts{
			AllowInvalidAmounts: true,
		})

		file, err := r.Read()
		require.NoError(t, err)
		require.NotNil(t, file)

		require.NoError(t, file.Create())
		require.NoError(t, file.Validate())
	})

}
