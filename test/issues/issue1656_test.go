package issues

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/ach"
	"github.com/moov-io/ach/cmd/achcli/describe"

	"github.com/stretchr/testify/require"
)

func TestIssue1656(t *testing.T) {
	file, err := ach.ReadJSONFile(filepath.Join("testdata", "issue1656.json"))
	require.NoError(t, err)

	if testing.Verbose() {
		describe.File(os.Stdout, file, nil)
	}

	require.Len(t, file.Batches, 1)

	entries := file.Batches[0].GetEntries()
	require.Len(t, entries, 2)

	// First Entry
	require.Equal(t, "Jane Smith", entries[0].IndividualName)
	require.Len(t, entries[0].Addenda05, 1)
	require.Equal(t, "abc", entries[0].Addenda05[0].PaymentRelatedInformation)

	// Second Entry
	require.Equal(t, "OFFSET", entries[1].IndividualName)
	require.Len(t, entries[1].Addenda05, 1)
	require.Equal(t, "def", entries[1].Addenda05[0].PaymentRelatedInformation)
}
