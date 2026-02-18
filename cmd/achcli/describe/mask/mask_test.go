// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package mask

import (
	"testing"

	"github.com/moov-io/ach"

	"github.com/stretchr/testify/require"
)

func TestMaskNumber(t *testing.T) {
	require.Equal(t, "*****", Number(""))
	require.Equal(t, "*****", Number("1"))
	require.Equal(t, "*****", Number("12"))
	require.Equal(t, "*****", Number("123"))
	require.Equal(t, "*****", Number("1234"))
	require.Equal(t, "**34 ", Number("1234 "))
	require.Equal(t, "**345", Number("12345"))
	require.Equal(t, "***2345", Number("  12345"))
	require.Equal(t, "**2345 ", Number(" 12345 "))
	require.Equal(t, "**3456", Number("123456"))
	require.Equal(t, "***4567", Number("1234567"))
	require.Equal(t, "****5678", Number("12345678"))
	require.Equal(t, "*******5678", Number("   12345678"))
	require.Equal(t, "*****6789", Number("123456789"))
	require.Equal(t, "******7890", Number("1234567890"))
	require.Equal(t, "********7890", Number("  1234567890"))
	require.Equal(t, "******7890  ", Number("1234567890  "))
	require.Equal(t, "*******7890 ", Number(" 1234567890 "))

	// Verify we mask .DFIAccountNumberField() as expected
	ed := &ach.EntryDetail{
		DFIAccountNumber: "12345678",
	}
	require.Equal(t, "****5678         ", Number(ed.DFIAccountNumberField()))

	// Test with multi-byte UTF-8 characters
	require.Equal(t, "**3é4", Number("123é4"))
	require.Equal(t, "*****5678", Number("12😊345678"))
}

func TestMaskName(t *testing.T) {
	require.Equal(t, "", Name(""))
	require.Equal(t, "* * *", Name(`a b c`))
	require.Equal(t, "* * *", Name(` a  b  c `))
	require.Equal(t, "Jo** ***", Name("John Doe"))
	require.Equal(t, "Jo** Sm*** **", Name("John Smith Jr"))
	require.Equal(t, "Al******* Lo********** ** ***", Name("Alexander Longnameiton Jr III"))

	// Verify we mask .IndividualNameField() as expected
	ed := &ach.EntryDetail{
		IndividualName: "Jane Smith Jr",
	}
	require.Equal(t, "Ja** Sm*** **", Name(ed.IndividualNameField()))
}

func TestMaskIdentification(t *testing.T) {
	ed := &ach.EntryDetail{
		IdentificationNumber: "abc123",
	}
	require.Equal(t, "**c123", Number(ed.IdentificationNumber))
	require.Equal(t, "**c123         ", Number(ed.IdentificationNumberField()))
}
