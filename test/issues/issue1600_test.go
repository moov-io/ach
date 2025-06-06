package issues

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/ach"

	"github.com/stretchr/testify/require"
)

func TestIssue1600(t *testing.T) {
	t.Run("zero fields", func(t *testing.T) {
		t.Run("2lCompanyNameAllZeros.txt", func(t *testing.T) {
			file, err := ach.ReadFile(filepath.Join("testdata", "issue1600", "2lCompanyNameAllZeros.txt"))
			require.ErrorAs(t, err, &ach.ErrOnlyZeros)
			require.ErrorAs(t, err, &ach.ErrFileEntryOutsideBatch)

			require.NoError(t, file.Create())
			require.NoError(t, file.Validate())

			var buf bytes.Buffer
			err = ach.NewWriter(&buf).Write(file)
			require.NoError(t, err)
			require.Equal(t, 950, buf.Len())
		})

		t.Run("2mCompanyEntryDescAllZeros.txt", func(t *testing.T) {
			file, err := ach.ReadFile(filepath.Join("testdata", "issue1600", "2mCompanyEntryDescAllZeros.txt"))
			require.ErrorAs(t, err, &ach.ErrOnlyZeros)
			require.ErrorAs(t, err, &ach.ErrFileEntryOutsideBatch)

			err = file.Create()
			require.ErrorContains(t, err, "must have []*Batches or []*IATBatches to be built")

			err = file.Validate()
			require.ErrorContains(t, err, "BatchCount calculated 0 is out-of-balance with file control 2")

			var buf bytes.Buffer
			err = ach.NewWriter(&buf).Write(file)
			require.ErrorContains(t, err, "BatchCount calculated 0 is out-of-balance with file control 2")
			require.Equal(t, 0, buf.Len())
		})

	})

	t.Run("out of order", func(t *testing.T) {
		t.Run("1e5RecordBefore1or8.txt", func(t *testing.T) {
			file, err := ach.ReadFile(filepath.Join("testdata", "issue1600", "1e5RecordBefore1or8.txt"))
			require.ErrorAs(t, err, &ach.ErrMisplacedFileHeader)

			err = file.Create()
			require.ErrorContains(t, err, "ImmediateDestination            is a mandatory field")
		})

		t.Run("1k8RecordFollowing9Record.txt", func(t *testing.T) {
			_, err := ach.ReadFile(filepath.Join("testdata", "issue1600", "1k8RecordFollowing9Record.txt"))
			require.ErrorAs(t, err, &ach.ErrExtraRecordsAfterFileControl)
		})
	})
}

func TestIssue1600_TEL_Validate(t *testing.T) {
	t.Run("3fTELIndvNameAllBlank.txt", func(t *testing.T) {
		file, err := ach.ReadFile(filepath.Join("testdata", "issue1600", "3fTELIndvNameAllBlank.txt"))
		require.ErrorAs(t, err, &ach.ErrOnlyZeros)

		for idx := range file.Batches {
			bh := file.Batches[idx].GetHeader()

			if bh.StandardEntryClassCode == ach.TEL {
				err := file.Batches[idx].Validate()
				require.ErrorAs(t, err, &ach.ErrOnlyZeros)
			}
		}
	})

	t.Run("3fTELIndvNameAllZeros.txt", func(t *testing.T) {
		file, err := ach.ReadFile(filepath.Join("testdata", "issue1600", "3fTELIndvNameAllZeros.txt"))
		require.ErrorAs(t, err, &ach.ErrOnlyZeros)

		for idx := range file.Batches {
			bh := file.Batches[idx].GetHeader()

			if bh.StandardEntryClassCode == ach.TEL {
				err := file.Batches[idx].Validate()
				require.ErrorAs(t, err, &ach.ErrOnlyZeros)
			}
		}
	})

	t.Run("AllowEmptyIndividualName", func(t *testing.T) {
		bs, err := os.ReadFile(filepath.Join("testdata", "issue1600", "3fTELIndvNameAllBlank.txt"))
		require.NoError(t, err)

		rdr := ach.NewReader(bytes.NewReader(bs))
		rdr.SetValidation(&ach.ValidateOpts{
			AllowEmptyIndividualName: true,
		})

		file, err := rdr.Read()
		require.NoError(t, err)
		require.NotNil(t, file)
	})
}

func TestIssue1600_FormatCode(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("testdata", "issue1600", "1qFormatCodeIncorrect.txt"))
	require.NoError(t, err)

	rdr := ach.NewReader(bytes.NewReader(bs))
	rdr.SetValidation(&ach.ValidateOpts{
		SkipAll: false,
	})

	file, err := rdr.Read()
	require.ErrorAs(t, err, &ach.ErrFormatCode)
	require.Equal(t, "011000015", file.Header.ImmediateOrigin)

	err = file.Header.Validate()
	require.ErrorAs(t, err, &ach.ErrFormatCode)
}
