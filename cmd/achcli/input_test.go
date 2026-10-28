// Copyright 2026 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func testdataPath(name string) string {
	return filepath.Join("..", "..", "test", "testdata", name)
}

func testdataBytes(t *testing.T, name string) []byte {
	t.Helper()
	data, err := os.ReadFile(testdataPath(name))
	require.NoError(t, err)
	return data
}

func setStdinReader(t *testing.T, r io.Reader) {
	t.Helper()
	prev := stdin
	stdin = &stdinCache{r: r}
	t.Cleanup(func() {
		stdin = prev
	})
}

func setStdinFile(t *testing.T, name string) {
	t.Helper()
	f, err := os.Open(testdataPath(name))
	require.NoError(t, err)
	prev := stdin
	stdin = &stdinCache{r: f}
	t.Cleanup(func() {
		stdin = prev
		f.Close()
	})
}

func TestCanReadStdin(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		require.False(t, canReadStdin(nil))
	})
	t.Run("injected reader", func(t *testing.T) {
		require.True(t, canReadStdin(bytes.NewReader([]byte("x"))))
	})
	t.Run("redirected file", func(t *testing.T) {
		f, err := os.Open(testdataPath("ppd-debit.ach"))
		require.NoError(t, err)
		t.Cleanup(func() { f.Close() })
		require.True(t, canReadStdin(f))
	})
	t.Run("empty file", func(t *testing.T) {
		f, err := os.CreateTemp(t.TempDir(), "empty")
		require.NoError(t, err)
		t.Cleanup(func() { f.Close() })
		require.False(t, canReadStdin(f))
	})
	t.Run("named pipe", func(t *testing.T) {
		r, w, err := os.Pipe()
		require.NoError(t, err)
		t.Cleanup(func() {
			r.Close()
			w.Close()
		})
		require.True(t, canReadStdin(r))
	})
}

func TestReadInput(t *testing.T) {
	ppd := testdataBytes(t, "ppd-debit.ach")

	t.Run("file path", func(t *testing.T) {
		got, err := readInput(testdataPath("ppd-debit.ach"))
		require.NoError(t, err)
		require.Equal(t, ppd, got)
	})
	t.Run("missing file", func(t *testing.T) {
		_, err := readInput(filepath.Join(t.TempDir(), "missing.ach"))
		require.Error(t, err)
	})
	t.Run("injected reader", func(t *testing.T) {
		r := bytes.NewReader(ppd)
		setStdinReader(t, r)
		got, err := readInput("-")
		require.NoError(t, err)
		require.Equal(t, ppd, got)
	})
	t.Run("caches stdin", func(t *testing.T) {
		r := bytes.NewReader(ppd)
		setStdinReader(t, r)
		first, err := readInput("-")
		require.NoError(t, err)
		require.Equal(t, 0, r.Len())
		second, err := readInput("-")
		require.NoError(t, err)
		require.Equal(t, first, second)
		require.Equal(t, ppd, second)
	})
	t.Run("redirected file", func(t *testing.T) {
		setStdinFile(t, "ppd-debit.ach")
		got, err := readInput("-")
		require.NoError(t, err)
		require.Equal(t, ppd, got)
	})
	t.Run("named pipe", func(t *testing.T) {
		pr, pw, err := os.Pipe()
		require.NoError(t, err)
		t.Cleanup(func() {
			pr.Close()
			pw.Close()
		})
		go func() {
			_, _ = pw.Write(ppd)
			_ = pw.Close()
		}()
		setStdinReader(t, pr)
		got, err := readInput("-")
		require.NoError(t, err)
		require.Equal(t, ppd, got)
	})
	t.Run("empty injected reader", func(t *testing.T) {
		setStdinReader(t, bytes.NewReader(nil))
		_, err := readInput("-")
		require.ErrorIs(t, err, errStdinEmpty)
		_, err = readInput("-")
		require.ErrorIs(t, err, errStdinEmpty)
	})
	t.Run("empty file does not block", func(t *testing.T) {
		f, err := os.CreateTemp(t.TempDir(), "empty")
		require.NoError(t, err)
		t.Cleanup(func() { f.Close() })
		setStdinReader(t, f)
		_, err = readInput("-")
		require.ErrorIs(t, err, errStdinEmpty)
	})
}

func TestReadIncomingFileStdin(t *testing.T) {
	for _, name := range []string{"ppd-debit.ach", "ppd-valid.json"} {
		t.Run(name, func(t *testing.T) {
			setStdinReader(t, bytes.NewReader(testdataBytes(t, name)))
			file, err := readIncomingFile("-", nil)
			require.NoError(t, err)
			require.Len(t, file.Batches, 1)
		})
	}
}

func TestValidateTotalsFileStdin(t *testing.T) {
	t.Run("ach", func(t *testing.T) {
		setStdinReader(t, bytes.NewReader(testdataBytes(t, "ppd-debit.ach")))
		_, err := validateTotalsFile(totalsValidation, "-", nil)
		require.NoError(t, err)
	})
	t.Run("json", func(t *testing.T) {
		setStdinReader(t, bytes.NewReader(testdataBytes(t, "ppd-valid.json")))
		_, err := validateTotalsFile(totalsValidation, "-", nil)
		require.Error(t, err)
		require.ErrorContains(t, err, "raw ACH only")
	})
	t.Run("whole file", func(t *testing.T) {
		setStdinReader(t, bytes.NewReader(testdataBytes(t, "ppd-debit.ach")))
		_, err := validateTotalsFile(fileValidation, "-", nil)
		require.NoError(t, err)
	})
	t.Run("empty", func(t *testing.T) {
		setStdinReader(t, bytes.NewReader(nil))
		_, err := validateTotalsFile(totalsValidation, "-", nil)
		require.ErrorIs(t, err, errStdinEmpty)
	})
}

func TestRejectStdinFix(t *testing.T) {
	require.NoError(t, rejectStdinFix("file.ach"))
	require.ErrorIs(t, rejectStdinFix("-"), errStdinFix)
}

func TestStdinCLI(t *testing.T) {
	if runtime.GOOS == "js" || runtime.GOOS == "wasip1" {
		t.Skip("executable CLI tests require native subprocesses")
	}
	bin := filepath.Join(t.TempDir(), "achcli")
	if runtime.GOOS == "windows" {
		bin += ".exe"
	}
	build := exec.CommandContext(t.Context(), "go", "build", "-o", bin, ".")
	buildOutput, err := build.CombinedOutput()
	require.NoError(t, err, "%s", buildOutput)

	ppd := testdataBytes(t, "ppd-debit.ach")
	run := func(t *testing.T, stdin []byte, args ...string) (string, int) {
		t.Helper()
		cmd := exec.CommandContext(t.Context(), bin, args...)
		if stdin != nil {
			cmd.Stdin = bytes.NewReader(stdin)
		} else {
			cmd.Stdin = bytes.NewReader(nil)
		}
		output, err := cmd.CombinedOutput()
		if err == nil {
			return string(output), 0
		}
		var exitErr *exec.ExitError
		require.True(t, errors.As(err, &exitErr), "expected an exited CLI, got %v: %s", err, output)
		return string(output), exitErr.ExitCode()
	}

	t.Run("reformat json", func(t *testing.T) {
		out, code := run(t, ppd, "-reformat=json", "-")
		require.Equal(t, 0, code, out)
		require.Contains(t, out, `"fileHeader"`)
	})
	t.Run("validate totals", func(t *testing.T) {
		out, code := run(t, ppd, "-validate-totals", "-")
		require.Equal(t, 0, code, out)
		require.Empty(t, out)
	})
	t.Run("describe", func(t *testing.T) {
		out, code := run(t, ppd, "-")
		require.Equal(t, 0, code, out)
		require.Contains(t, out, "Describing ACH file '-'")
	})
	t.Run("diff caches stdin", func(t *testing.T) {
		out, code := run(t, ppd, "-diff", "-", "-")
		require.Equal(t, 0, code, out)
	})
	t.Run("diff file against stdin", func(t *testing.T) {
		out, code := run(t, ppd, "-diff", testdataPath("ppd-debit.ach"), "-")
		require.Equal(t, 0, code, out)
	})
	t.Run("empty stdin", func(t *testing.T) {
		out, code := run(t, nil, "-reformat=json", "-")
		require.Equal(t, 1, code, out)
		require.Contains(t, out, "stdin is empty")
	})
	t.Run("fix rejects stdin", func(t *testing.T) {
		out, code := run(t, ppd, "-fix", "-")
		require.Equal(t, 1, code, out)
		require.Contains(t, out, "not supported with -fix")
	})
}
