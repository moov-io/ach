package issues

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/ach"
	"github.com/moov-io/ach/cmd/achcli/describe"
	"github.com/stretchr/testify/require"
)

func TestIssue1202(t *testing.T) {
	file, err := ach.ReadFile(filepath.Join("..", "..", "test", "testdata", "return-no-batch-controls.ach"))
	require.ErrorContains(t, err, ach.ErrFileHeader.Error())

	if testing.Verbose() {
		describe.File(os.Stdout, file, nil)
	}

	require.Len(t, file.Batches, 2)
	require.Len(t, file.NotificationOfChange, 1)
	require.Len(t, file.ReturnEntries, 1)

	b1Entries := file.Batches[0].GetEntries()
	require.Len(t, b1Entries, 1)
	require.Nil(t, b1Entries[0].Addenda98)
	require.NotNil(t, b1Entries[0].Addenda99)

	b2Entries := file.Batches[1].GetEntries()
	require.Len(t, b2Entries, 1)
	require.NotNil(t, b2Entries[0].Addenda98)
	require.Nil(t, b2Entries[0].Addenda99)
}
