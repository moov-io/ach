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
	"github.com/moov-io/ach/cmd/achcli/describe/mask"
	"github.com/stretchr/testify/require"
)

func TestDescribeFile(t *testing.T) {
	file, err := ach.ReadFile(filepath.Join("..", "..", "..", "test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)

	var buf bytes.Buffer
	File(&buf, file, &Opts{
		Options: mask.Options{
			MaskNames:          true,
			MaskAccountNumbers: true,
		},
	})
	if testing.Verbose() {
		os.Stdout.Write(buf.Bytes())
	}
	require.Equal(t, 1218, buf.Len())
	require.Contains(t, buf.String(), "Re****** Ac***** Na**")
}

func TestDescribeIAT(t *testing.T) {
	file, err := ach.ReadFile(filepath.Join("..", "..", "..", "test", "testdata", "iat-debit.ach"))
	require.NoError(t, err)

	var buf bytes.Buffer
	File(&buf, file, &Opts{
		Options: mask.Options{
			MaskAccountNumbers: true,
		},
	})
	if testing.Verbose() {
		os.Stdout.Write(buf.Bytes())
	}
	require.Equal(t, 5262, buf.Len())
	require.Contains(t, buf.String(), "*****6789")
}

func TestDescribeReturn(t *testing.T) {
	file, err := ach.ReadFile(filepath.Join("..", "..", "..", "test", "testdata", "return-WEB.ach"))
	require.NoError(t, err)

	var buf bytes.Buffer
	File(&buf, file, &Opts{
		Options: mask.Options{
			MaskIdentification: true,
		},
	})
	if testing.Verbose() {
		os.Stdout.Write(buf.Bytes())
	}
	require.Equal(t, 2692, buf.Len())
	require.Contains(t, buf.String(), "***********tMGN")
}

func TestDescribeCorrection(t *testing.T) {
	file, err := ach.ReadFile(filepath.Join("..", "..", "..", "test", "testdata", "cor-example.ach"))
	require.NoError(t, err)

	var buf bytes.Buffer
	File(&buf, file, &Opts{
		Options: mask.Options{
			MaskCorrectedData: true,
		},
	})
	if testing.Verbose() {
		os.Stdout.Write(buf.Bytes())
	}
	require.Equal(t, 1401, buf.Len())
	require.Contains(t, buf.String(), "******1614")
}

func TestFormatAmount(t *testing.T) {
	require.Equal(t, "12345", formatAmount(false, 12345))
	require.Equal(t, "123.45", formatAmount(true, 12345))
	require.Equal(t, "1,234,567.89", formatAmount(true, 123456789))
}
