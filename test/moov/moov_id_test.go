package moov_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/ach"
	"github.com/moov-io/ach/cmd/achcli/describe"

	"github.com/stretchr/testify/require"
)

func TestReadWithMoovIDs(t *testing.T) {
	fd, err := os.Open(filepath.Join("..", "testdata", "moov-ids.ach"))
	require.NoError(t, err)

	t.Cleanup(func() { fd.Close() })

	r := ach.NewReader(fd)
	r.SetValidation(&ach.ValidateOpts{
		CustomTraceNumbers: true,
	})

	file, err := r.Read()
	require.NoError(t, err)

	if testing.Verbose() {
		describe.File(os.Stdout, &file, nil)
	}

	// File Header changes
	fh := file.Header
	require.Equal(t, "123456789", fh.ImmediateOrigin)
	require.Equal(t, "228591016", fh.ImmediateDestination)
	require.Equal(t, "48104ccf-d5c9-49cc-a273-aaac7d14cd68", fh.ImmediateDestinationName+fh.ImmediateOriginName)

	// Batch Header changes
	require.Len(t, file.Batches, 2)

	bh0 := file.Batches[0].GetHeader()
	require.Equal(t, ach.CreditsOnly, bh0.ServiceClassCode)
	require.Equal(t, ach.PPD, bh0.StandardEntryClassCode)
	require.Equal(t, "2ada48fb-92bc-4882-890e-e39bc1bebc4f", bh0.CompanyName+bh0.CompanyDiscretionaryData)
	require.Equal(t, "COMPANY", bh0.CompanyIdentification)

	bh1 := file.Batches[1].GetHeader()
	require.Equal(t, ach.CreditsOnly, bh1.ServiceClassCode)
	require.Equal(t, ach.PPD, bh1.StandardEntryClassCode)
	require.Equal(t, "f36eb0f2-65f6-4481-bce3-4c1c6e6a586c", bh1.CompanyName+bh1.CompanyDiscretionaryData)
	require.Equal(t, "COMPANY2", bh1.CompanyIdentification)

	// Entry changes
	b0entries := file.Batches[0].GetEntries()
	require.Len(t, b0entries, 2)
	b1entries := file.Batches[1].GetEntries()
	require.Len(t, b0entries, 2)

	// Batch 0, Entry 0
	require.Equal(t, "11122233", b0entries[0].RDFIIdentification)
	require.Equal(t, "12345678901", b0entries[0].DFIAccountNumber)

	// Batch 0, Entry 1
	require.Equal(t, "11122233", b0entries[1].RDFIIdentification)
	require.Equal(t, "12345678902", b0entries[1].DFIAccountNumber)

	// Batch 1, Entry 0
	require.Equal(t, "44455566", b1entries[0].RDFIIdentification)
	require.Equal(t, "12345678903", b1entries[0].DFIAccountNumber)

	// Batch 1, Entry 1
	require.Equal(t, "77788899", b1entries[1].RDFIIdentification)
	require.Equal(t, "12345678904", b1entries[1].DFIAccountNumber)

	// File Control
	require.Equal(t, "8c365fa8-460f-4665-99d0-735ec25e346b   ", file.Control.Reserved)
}
