package issues

import (
	"path/filepath"
	"testing"

	"github.com/moov-io/ach"

	"github.com/stretchr/testify/require"
)

func TestIssue1595(t *testing.T) {
	file, err := ach.ReadFile(filepath.Join("testdata", "issue1595.ach"))
	require.NoError(t, err)
	require.NoError(t, file.Validate())
}
