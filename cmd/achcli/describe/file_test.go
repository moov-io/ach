// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package describe

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/moov-io/ach"
	"github.com/stretchr/testify/require"
)

func TestDescribeFile(t *testing.T) {
	file, err := ach.ReadFile(filepath.Join("..", "..", "..", "test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)

	var buf bytes.Buffer
	File(&buf, file, nil) // No Options
	if testing.Verbose() {
		os.Stdout.Write(buf.Bytes())
	}
	require.Equal(t, 1218, buf.Len())
}

func TestDescribeIAT(t *testing.T) {
	file, err := ach.ReadFile(filepath.Join("..", "..", "..", "test", "testdata", "iat-debit.ach"))
	require.NoError(t, err)

	var buf bytes.Buffer
	File(&buf, file, nil) // No Options
	if testing.Verbose() {
		os.Stdout.Write(buf.Bytes())
	}
	require.Equal(t, 5262, buf.Len())
}

func TestDescribeReturn(t *testing.T) {
	file, err := ach.ReadFile(filepath.Join("..", "..", "..", "test", "testdata", "return-WEB.ach"))
	require.NoError(t, err)

	var buf bytes.Buffer
	File(&buf, file, nil) // No Options
	if testing.Verbose() {
		os.Stdout.Write(buf.Bytes())
	}
	require.Equal(t, 2692, buf.Len())
}

func TestDescribeCorrection(t *testing.T) {
	file, err := ach.ReadFile(filepath.Join("..", "..", "..", "test", "testdata", "cor-example.ach"))
	require.NoError(t, err)

	var buf bytes.Buffer
	File(&buf, file, nil) // No Options
	if testing.Verbose() {
		os.Stdout.Write(buf.Bytes())
	}
	require.Equal(t, 1401, buf.Len())
}

func TestFormatAmount(t *testing.T) {
	require.Equal(t, "12345", formatAmount(false, 12345))
	require.Equal(t, "123.45", formatAmount(true, 12345))
	require.Equal(t, "1,234,567.89", formatAmount(true, 123456789))
}

func TestMaskNumber(t *testing.T) {
	require.Equal(t, "*****", maskNumber(""))
	require.Equal(t, "*****", maskNumber("1"))
	require.Equal(t, "*****", maskNumber("12"))
	require.Equal(t, "*****", maskNumber("123"))
	require.Equal(t, "*****", maskNumber("1234"))
	require.Equal(t, "**34 ", maskNumber("1234 "))
	require.Equal(t, "**345", maskNumber("12345"))
	require.Equal(t, "***2345", maskNumber("  12345"))
	require.Equal(t, "**2345 ", maskNumber(" 12345 "))
	require.Equal(t, "**3456", maskNumber("123456"))
	require.Equal(t, "***4567", maskNumber("1234567"))
	require.Equal(t, "****5678", maskNumber("12345678"))
	require.Equal(t, "*******5678", maskNumber("   12345678"))
	require.Equal(t, "*****6789", maskNumber("123456789"))
	require.Equal(t, "******7890", maskNumber("1234567890"))
	require.Equal(t, "********7890", maskNumber("  1234567890"))
	require.Equal(t, "******7890  ", maskNumber("1234567890  "))
	require.Equal(t, "*******7890 ", maskNumber(" 1234567890 "))

	// Verify we mask .DFIAccountNumberField() as expected
	ed := &ach.EntryDetail{
		DFIAccountNumber: "12345678",
	}
	require.Equal(t, "****5678         ", maskNumber(ed.DFIAccountNumberField()))
}

func TestMaskName(t *testing.T) {
	require.Equal(t, "", maskName(""))
	require.Equal(t, "* * *", maskName(`a b c`))
	require.Equal(t, "* * *", maskName(` a  b  c `))
	require.Equal(t, "Jo** ***", maskName("John Doe"))
	require.Equal(t, "Jo** Sm*** **", maskName("John Smith Jr"))
	require.Equal(t, "Al******* Lo********** ** ***", maskName("Alexander Longnameiton Jr III"))

	// Verify we mask .IndividualNameField() as expected
	ed := &ach.EntryDetail{
		IndividualName: "Jane Smith Jr",
	}
	require.Equal(t, "Ja** Sm*** **", maskName(ed.IndividualNameField()))
}

func TestMaskIdentification(t *testing.T) {
	ed := &ach.EntryDetail{
		IdentificationNumber: "abc123",
	}
	require.Equal(t, "**c123", maskNumber(ed.IdentificationNumber))
	require.Equal(t, "**c123         ", maskNumber(ed.IdentificationNumberField()))
}
