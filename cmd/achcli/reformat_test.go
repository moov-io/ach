// Copyright 2026 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadInputStdin(t *testing.T) {
	setStdin := func(t *testing.T, name string) {
		t.Helper()
		f, err := os.Open(filepath.Join("..", "..", "test", "testdata", name))
		require.NoError(t, err)
		stdin := os.Stdin
		os.Stdin, stdinRead = f, false
		t.Cleanup(func() {
			os.Stdin, stdinRead = stdin, false
			f.Close()
		})
	}

	for _, name := range []string{"ppd-debit.ach", "ppd-valid.json"} {
		t.Run(name, func(t *testing.T) {
			setStdin(t, name)
			file, err := readIncomingFile("-", nil)
			require.NoError(t, err)
			require.Len(t, file.Batches, 1)
		})
	}
	t.Run("validate totals", func(t *testing.T) {
		setStdin(t, "ppd-debit.ach")
		require.NoError(t, validateTotalsFile("-", nil))
	})
	t.Run("only once", func(t *testing.T) {
		setStdin(t, "ppd-debit.ach")
		_, err := readInput("-")
		require.NoError(t, err)
		_, err = readInput("-")
		require.ErrorContains(t, err, "can only be read once")
	})
}
