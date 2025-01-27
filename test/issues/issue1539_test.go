package issues

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/ach"
	"github.com/moov-io/ach/cmd/achcli/describe"

	"github.com/stretchr/testify/require"
)

func TestIssue1539(t *testing.T) {
	checkFile := func(t *testing.T, file *ach.File) {
		t.Helper()

		for _, b := range file.Batches {
			for _, e := range b.GetEntries() {
				// Check common fields
				require.Contains(t, e.TraceNumber, "29147102")

				switch e.TransactionCode {
				case ach.CheckingCredit:
					require.Equal(t, "27397636", e.RDFIIdentification)
					require.Equal(t, "0123456789", e.DFIAccountNumber)

				case ach.CheckingDebit:
					require.Equal(t, "43219876", e.RDFIIdentification)
					require.Equal(t, "3323523523", e.DFIAccountNumber)
				}
			}
		}
	}

	t.Run("json.Unmarshal", func(t *testing.T) {
		fd, err := os.Open(filepath.Join("testdata", "issue1539.json"))
		require.NoError(t, err)
		defer fd.Close()

		var file ach.File
		err = json.NewDecoder(fd).Decode(&file)
		require.NoError(t, err)

		// TraceNumbers should persist through Create()
		require.NoError(t, file.Create())
		for idx := range file.Batches {
			require.NoError(t, file.Batches[idx].Create())
		}

		if testing.Verbose() {
			describe.File(os.Stdout, &file, nil)
		}

		checkFile(t, &file)
	})

	t.Run("FileFromJSON", func(t *testing.T) {
		file, err := ach.ReadJSONFile(filepath.Join("testdata", "issue1539.json"))
		require.NoError(t, err)

		// TraceNumbers should persist through Create()
		require.NoError(t, file.Create())
		for idx := range file.Batches {
			require.NoError(t, file.Batches[idx].Create())
		}

		if testing.Verbose() {
			describe.File(os.Stdout, file, nil)
		}

		checkFile(t, file)
	})
}
