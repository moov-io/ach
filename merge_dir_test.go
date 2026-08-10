// Licensed to The Moov Authors under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. The Moov Authors licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package ach

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMergeDir__DefaultFileAcceptor(t *testing.T) {
	output := DefaultFileAcceptor("")
	require.Equal(t, AcceptFile, output)

	output = DefaultFileAcceptor("foo.ach")
	require.Equal(t, AcceptFile, output)

	output = DefaultFileAcceptor("foo.txt")
	require.Equal(t, AcceptFile, output)

	output = DefaultFileAcceptor("foo.json")
	require.Equal(t, AcceptAsJSON, output)

	output = DefaultFileAcceptor("foo.mp3")
	require.Equal(t, SkipFile, output)
}

func TestMergeDir(t *testing.T) {
	dir := t.TempDir()

	src, err := os.Open(filepath.Join("test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)
	t.Cleanup(func() { src.Close() })

	dst, err := os.Create(filepath.Join(dir, "input.ach"))
	require.NoError(t, err)

	_, err = io.Copy(dst, src)
	require.NoError(t, err)
	require.NoError(t, dst.Close())

	var conditions Conditions
	merged, err := MergeDir(dir, conditions, nil)
	require.NoError(t, err)
	require.Len(t, merged, 1)
}

func TestMergeDir_WithFS(t *testing.T) {
	dir := t.TempDir()
	sub := filepath.Join("a", "b", "c")
	require.NoError(t, os.MkdirAll(filepath.Join(dir, sub), 0777))

	src, err := os.Open(filepath.Join("test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)
	t.Cleanup(func() { src.Close() })

	dst, err := os.Create(filepath.Join(dir, sub, "input.ach"))
	require.NoError(t, err)

	_, err = io.Copy(dst, src)
	require.NoError(t, err)
	require.NoError(t, dst.Close())

	var conditions Conditions

	// partial dir + sub
	merged, err := MergeDir(sub, conditions, &MergeDirOptions{
		FS: os.DirFS(dir),
	})
	require.NoError(t, err)
	require.Len(t, merged, 1)

	// full dir
	merged, err = MergeDir(".", conditions, &MergeDirOptions{
		FS: os.DirFS(filepath.Join(dir, sub)),
	})
	require.NoError(t, err)
	require.Len(t, merged, 1)
}

func TestMergeDir_Nested(t *testing.T) {
	dir := t.TempDir()

	sub := filepath.Join("aaaa")
	require.NoError(t, os.MkdirAll(filepath.Join(dir, sub), 0777))
	sub = filepath.Join("inner")
	require.NoError(t, os.MkdirAll(filepath.Join(dir, sub), 0777))

	src, err := os.Open(filepath.Join("test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)
	t.Cleanup(func() { src.Close() })

	dst, err := os.Create(filepath.Join(dir, sub, "input.ach"))
	require.NoError(t, err)

	_, err = io.Copy(dst, src)
	require.NoError(t, err)
	require.NoError(t, dst.Close())

	var conditions Conditions

	// nothing in the top level
	merged, err := MergeDir(dir, conditions, nil)
	require.NoError(t, err)
	require.Empty(t, merged)

	// found files in sub directories
	merged, err = MergeDir(dir, conditions, &MergeDirOptions{
		SubDirectories: true,
	})
	require.NoError(t, err)
	require.Len(t, merged, 1)
}

func TestMergeDirHelpers(t *testing.T) {
	dir := os.DirFS(filepath.Join("test", "testdata"))

	t.Run("readFile", func(t *testing.T) {
		t.Run("ach", func(t *testing.T) {
			file, err := readFile(dir, "web-debit.ach", AcceptFile, nil)
			require.NoError(t, err)
			require.NotNil(t, file)
		})

		t.Run("json", func(t *testing.T) {
			file, err := readFile(dir, "ppd-valid.json", AcceptAsJSON, nil)
			require.NoError(t, err)
			require.NotNil(t, file)
		})
	})

	t.Run("readValidateOptsFromFile", func(t *testing.T) {
		opts := &MergeDirOptions{
			FS: dir,

			ValidateOptsExtension: ".json",
		}
		output := readValidateOptsFromFile("web-debit.ach", opts)
		require.NotNil(t, output)
		require.True(t, output.RequireABAOrigin)
		require.True(t, output.AllowMissingFileHeader)
		require.True(t, output.UnequalServiceClassCode)
	})
}

func TestMergeDir_DeadlockPrevention(t *testing.T) {
	dir := t.TempDir()

	// Copy a valid file
	src1, err := os.Open(filepath.Join("test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)
	defer src1.Close()

	dst1, err := os.Create(filepath.Join(dir, "valid.ach"))
	require.NoError(t, err)
	_, err = io.Copy(dst1, src1)
	require.NoError(t, err)
	require.NoError(t, dst1.Close())

	// Copy an ADV file which will cause sorted.add to fail
	src2, err := os.Open(filepath.Join("test", "testdata", "adv.ach"))
	require.NoError(t, err)
	defer src2.Close()

	dst2, err := os.Create(filepath.Join(dir, "adv.ach"))
	require.NoError(t, err)
	_, err = io.Copy(dst2, src2)
	require.NoError(t, err)
	require.NoError(t, dst2.Close())

	// Copy another valid file to ensure concurrency
	src3, err := os.Open(filepath.Join("test", "testdata", "web-debit.ach"))
	require.NoError(t, err)
	defer src3.Close()

	dst3, err := os.Create(filepath.Join(dir, "valid2.ach"))
	require.NoError(t, err)
	_, err = io.Copy(dst3, src3)
	require.NoError(t, err)
	require.NoError(t, dst3.Close())

	// MergeDir should fail due to ADV file, but not deadlock
	_, err = MergeDir(dir, Conditions{}, nil)
	require.Error(t, err)
	require.Contains(t, err.Error(), "merging ADV batches is not supported")
}

func TestMergeDir_HeaderFromFirstFileNoRace(t *testing.T) {
	// Stress concurrent parse + merge; header must come from a real file.
	dir := t.TempDir()
	src, err := os.Open(filepath.Join("test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)
	defer src.Close()

	for i := 0; i < 50; i++ {
		dst, err := os.Create(filepath.Join(dir, fmt.Sprintf("f-%d.ach", i)))
		require.NoError(t, err)
		_, err = src.Seek(0, 0)
		require.NoError(t, err)
		_, err = io.Copy(dst, src)
		require.NoError(t, err)
		require.NoError(t, dst.Close())
	}

	merged, err := MergeDir(dir, Conditions{}, &MergeDirOptions{ParseWorkers: 8})
	require.NoError(t, err)
	require.NotEmpty(t, merged)
	require.NotEmpty(t, merged[0].Header.ImmediateOrigin)
	require.NotEmpty(t, merged[0].Header.ImmediateDestination)
}

func TestDefaultParseWorkers(t *testing.T) {
	n := defaultParseWorkers()
	require.GreaterOrEqual(t, n, 8)
	require.LessOrEqual(t, n, 50)

	// Force the lower clamp (GOMAXPROCS*4 < 8)
	prev := runtime.GOMAXPROCS(1)
	t.Cleanup(func() { runtime.GOMAXPROCS(prev) })
	require.Equal(t, 8, defaultParseWorkers())

	// Force the upper clamp (GOMAXPROCS*4 > 50)
	runtime.GOMAXPROCS(20)
	require.Equal(t, 50, defaultParseWorkers())
}

func TestMergeDir_ParseError(t *testing.T) {
	dir := t.TempDir()

	// Valid file so parsers have work after the bad one
	src, err := os.Open(filepath.Join("test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)
	defer src.Close()
	for i := 0; i < 5; i++ {
		dst, err := os.Create(filepath.Join(dir, fmt.Sprintf("valid-%d.ach", i)))
		require.NoError(t, err)
		_, err = src.Seek(0, 0)
		require.NoError(t, err)
		_, err = io.Copy(dst, src)
		require.NoError(t, err)
		require.NoError(t, dst.Close())
	}

	// Corrupt Nacha content
	require.NoError(t, os.WriteFile(filepath.Join(dir, "zzz-bad.ach"), []byte("not-a-valid-ach-file\n"), 0644))

	_, err = MergeDir(dir, Conditions{}, &MergeDirOptions{ParseWorkers: 2})
	require.Error(t, err)
	require.Contains(t, err.Error(), "merging")
}

func TestMergeDir_SkipAndJSON(t *testing.T) {
	dir := t.TempDir()

	// Copy a JSON ACH file
	src, err := os.Open(filepath.Join("test", "testdata", "ppd-valid.json"))
	require.NoError(t, err)
	defer src.Close()
	dst, err := os.Create(filepath.Join(dir, "file.json"))
	require.NoError(t, err)
	_, err = io.Copy(dst, src)
	require.NoError(t, err)
	require.NoError(t, dst.Close())

	// Noise files that must be skipped
	require.NoError(t, os.WriteFile(filepath.Join(dir, "notes.md"), []byte("# hi"), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "image.png"), []byte{0x89, 0x50}, 0644))

	merged, err := MergeDir(dir, Conditions{}, nil)
	require.NoError(t, err)
	require.Len(t, merged, 1)
}

func TestMergeDir_ValidateOptsExtension(t *testing.T) {
	dir := t.TempDir()

	src, err := os.Open(filepath.Join("test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)
	defer src.Close()
	dst, err := os.Create(filepath.Join(dir, "input.ach"))
	require.NoError(t, err)
	_, err = io.Copy(dst, src)
	require.NoError(t, err)
	require.NoError(t, dst.Close())

	// Sidecar ValidateOpts next to the ACH file (no FS — exercises os.Open path)
	optsJSON := `{"requireABAOrigin":true,"customReturnCodes":true}`
	require.NoError(t, os.WriteFile(filepath.Join(dir, "input.validate"), []byte(optsJSON), 0644))

	// Missing sidecar for a second file should just yield nil opts
	dst2, err := os.Create(filepath.Join(dir, "other.ach"))
	require.NoError(t, err)
	_, err = src.Seek(0, 0)
	require.NoError(t, err)
	_, err = io.Copy(dst2, src)
	require.NoError(t, err)
	require.NoError(t, dst2.Close())

	merged, err := MergeDir(dir, Conditions{}, &MergeDirOptions{
		ValidateOptsExtension: ".validate",
	})
	require.NoError(t, err)
	require.NotEmpty(t, merged)
}

func TestMergeDir_FSSubError(t *testing.T) {
	dir := t.TempDir()
	// fs.Sub rejects paths that escape the FS root
	_, err := MergeDir("..", Conditions{}, &MergeDirOptions{
		FS: os.DirFS(dir),
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "fs.Sub")
}

func TestMergeDir_CustomAcceptFile(t *testing.T) {
	dir := t.TempDir()
	src, err := os.Open(filepath.Join("test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)
	defer src.Close()
	dst, err := os.Create(filepath.Join(dir, "keep.ach"))
	require.NoError(t, err)
	_, err = io.Copy(dst, src)
	require.NoError(t, err)
	require.NoError(t, dst.Close())
	require.NoError(t, os.WriteFile(filepath.Join(dir, "drop.ach"), []byte("x"), 0644))

	merged, err := MergeDir(dir, Conditions{}, &MergeDirOptions{
		AcceptFile: func(path string) FileAcceptance {
			if strings.Contains(path, "drop") {
				return SkipFile
			}
			return AcceptFile
		},
	})
	require.NoError(t, err)
	require.Len(t, merged, 1)
}

func TestMergeDir_SubDirectoriesDoesNotEnqueueDirs(t *testing.T) {
	// Nested ACH files are merged; directory paths themselves must not be treated
	// as inputs (walkDir continues past dirs after recursing).
	dir := t.TempDir()
	sub := filepath.Join(dir, "nested")
	require.NoError(t, os.MkdirAll(sub, 0755))
	// Empty nested dir should be harmless
	require.NoError(t, os.MkdirAll(filepath.Join(dir, "empty-nested"), 0755))

	src, err := os.Open(filepath.Join("test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)
	defer src.Close()
	dst, err := os.Create(filepath.Join(sub, "input.ach"))
	require.NoError(t, err)
	_, err = io.Copy(dst, src)
	require.NoError(t, err)
	require.NoError(t, dst.Close())

	merged, err := MergeDir(dir, Conditions{}, &MergeDirOptions{SubDirectories: true})
	require.NoError(t, err)
	require.Len(t, merged, 1)
}

func TestWalkDir_Errors(t *testing.T) {
	t.Run("os.ReadDir missing", func(t *testing.T) {
		ch := make(chan string, 4)
		err := walkDir(context.Background(), nil, filepath.Join(t.TempDir(), "nope"), &MergeDirOptions{}, ch)
		require.Error(t, err)
		require.Contains(t, err.Error(), "os.readdir")
	})

	t.Run("fs.ReadDir error", func(t *testing.T) {
		ch := make(chan string, 4)
		err := walkDir(context.Background(), errReadDirFS{}, ".", &MergeDirOptions{}, ch)
		require.Error(t, err)
		require.Contains(t, err.Error(), "fs.readdir")
	})

	t.Run("nested walk error", func(t *testing.T) {
		// Parent lists a subdirectory; nested ReadDir fails.
		fsys := &nestedErrFS{failOn: "child"}
		ch := make(chan string, 4)
		err := walkDir(context.Background(), fsys, ".", &MergeDirOptions{SubDirectories: true}, ch)
		require.Error(t, err)
	})

	t.Run("cancelled context", func(t *testing.T) {
		dir := t.TempDir()
		require.NoError(t, os.WriteFile(filepath.Join(dir, "a.ach"), []byte("x"), 0644))
		require.NoError(t, os.WriteFile(filepath.Join(dir, "b.ach"), []byte("y"), 0644))

		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		// Unbuffered so send blocks and selects ctx.Done
		ch := make(chan string)
		done := make(chan error, 1)
		go func() {
			done <- walkDir(ctx, nil, dir, &MergeDirOptions{}, ch)
		}()
		select {
		case err := <-done:
			require.NoError(t, err)
		case <-time.After(2 * time.Second):
			t.Fatal("walkDir did not return after cancel")
		}
	})
}

type errReadDirFS struct{}

func (errReadDirFS) Open(name string) (fs.File, error) {
	return nil, errors.New("open not supported")
}
func (errReadDirFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return nil, errors.New("readdir fail")
}

// nestedErrFS lists one child dir at "."; ReadDir("child") fails.
type nestedErrFS struct{ failOn string }

func (n *nestedErrFS) Open(name string) (fs.File, error) {
	return nil, errors.New("open not supported")
}
func (n *nestedErrFS) ReadDir(name string) ([]fs.DirEntry, error) {
	name = filepath.Clean(name)
	if name == "." || name == "" {
		return []fs.DirEntry{fakeDirEntry{name: "child", dir: true}}, nil
	}
	if name == n.failOn || name == filepath.Join(".", n.failOn) {
		return nil, errors.New("nested readdir fail")
	}
	return nil, fs.ErrNotExist
}

type fakeDirEntry struct {
	name string
	dir  bool
}

func (f fakeDirEntry) Name() string               { return f.name }
func (f fakeDirEntry) IsDir() bool                { return f.dir }
func (f fakeDirEntry) Type() fs.FileMode          { return 0 }
func (f fakeDirEntry) Info() (fs.FileInfo, error) { return nil, nil }

func TestReadFileForMerging_Helpers(t *testing.T) {
	t.Run("empty path skipped", func(t *testing.T) {
		ctx := context.Background()
		paths := make(chan string, 2)
		files := make(chan *File, 2)
		paths <- ""
		close(paths)
		require.NoError(t, readFileForMerging(ctx, paths, files, &MergeDirOptions{
			AcceptFile: DefaultFileAcceptor,
		}))
		close(files)
		require.Nil(t, <-files)
	})

	t.Run("cancelled drains paths", func(t *testing.T) {
		// Unbuffered paths with no sender means only ctx.Done is selectable.
		// Close paths shortly after so the drain loop can finish.
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		paths := make(chan string)
		files := make(chan *File, 1)
		go func() {
			time.Sleep(20 * time.Millisecond)
			close(paths)
		}()
		require.NoError(t, readFileForMerging(ctx, paths, files, &MergeDirOptions{}))
	})

	t.Run("skip file", func(t *testing.T) {
		ctx := context.Background()
		paths := make(chan string, 1)
		files := make(chan *File, 1)
		paths <- "noise.bin"
		close(paths)
		require.NoError(t, readFileForMerging(ctx, paths, files, &MergeDirOptions{
			AcceptFile: func(string) FileAcceptance { return SkipFile },
		}))
		close(files)
		require.Nil(t, <-files)
	})

	t.Run("directory path yields nil file", func(t *testing.T) {
		// walkDir no longer enqueues dirs, but a directory path must still be
		// skipped cleanly if it appears on the channel.
		ctx := context.Background()
		paths := make(chan string, 1)
		files := make(chan *File, 1)
		paths <- t.TempDir()
		close(paths)
		require.NoError(t, readFileForMerging(ctx, paths, files, &MergeDirOptions{
			AcceptFile: func(string) FileAcceptance { return AcceptFile },
		}))
		close(files)
		require.Nil(t, <-files)
	})

	t.Run("read error", func(t *testing.T) {
		ctx := context.Background()
		paths := make(chan string, 1)
		files := make(chan *File, 1)
		paths <- filepath.Join(t.TempDir(), "missing.ach")
		close(paths)
		err := readFileForMerging(ctx, paths, files, &MergeDirOptions{
			AcceptFile: func(string) FileAcceptance { return AcceptFile },
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "reading")
	})

	t.Run("cancel while sending file", func(t *testing.T) {
		// Fill a one-slot files channel so send blocks, then cancel.
		dir := t.TempDir()
		src, err := os.Open(filepath.Join("test", "testdata", "ppd-debit.ach"))
		require.NoError(t, err)
		defer src.Close()
		dst, err := os.Create(filepath.Join(dir, "a.ach"))
		require.NoError(t, err)
		_, err = io.Copy(dst, src)
		require.NoError(t, err)
		require.NoError(t, dst.Close())

		ctx, cancel := context.WithCancel(context.Background())
		paths := make(chan string, 2)
		// Unbuffered files channel — first send blocks until we cancel
		files := make(chan *File)
		paths <- filepath.Join(dir, "a.ach")
		paths <- filepath.Join(dir, "a.ach") // second path to drain after cancel
		close(paths)

		done := make(chan error, 1)
		go func() {
			done <- readFileForMerging(ctx, paths, files, &MergeDirOptions{
				AcceptFile: DefaultFileAcceptor,
			})
		}()

		// Give the reader time to parse and block on send
		time.Sleep(50 * time.Millisecond)
		cancel()

		select {
		case err := <-done:
			require.NoError(t, err)
		case <-time.After(3 * time.Second):
			t.Fatal("readFileForMerging did not return after cancel")
		}
	})
}

func TestReadFile_ErrorPaths(t *testing.T) {
	t.Run("skip", func(t *testing.T) {
		f, err := readFile(nil, "x", SkipFile, nil)
		require.NoError(t, err)
		require.Nil(t, f)
	})

	t.Run("open failure", func(t *testing.T) {
		f, err := readFile(nil, filepath.Join(t.TempDir(), "nope.ach"), AcceptFile, nil)
		require.Error(t, err)
		require.Nil(t, f)
		require.Contains(t, err.Error(), "opening")
	})

	t.Run("directory returns nil", func(t *testing.T) {
		dir := t.TempDir()
		f, err := readFile(nil, dir, AcceptFile, nil)
		require.NoError(t, err)
		require.Nil(t, f)
	})

	t.Run("invalid nacha", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "bad.ach")
		require.NoError(t, os.WriteFile(path, []byte("garbage"), 0644))
		f, err := readFile(nil, path, AcceptFile, nil)
		require.Error(t, err)
		require.Nil(t, f)
		require.Contains(t, err.Error(), "nacha")
	})

	t.Run("invalid json", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "bad.json")
		require.NoError(t, os.WriteFile(path, []byte("{not json"), 0644))
		f, err := readFile(nil, path, AcceptAsJSON, nil)
		require.Error(t, err)
		require.Nil(t, f)
	})

	t.Run("unknown acceptance", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "x.ach")
		require.NoError(t, os.WriteFile(path, []byte("x"), 0644))
		f, err := readFile(nil, path, FileAcceptance("nope"), nil)
		require.Error(t, err)
		require.Nil(t, f)
		require.Contains(t, err.Error(), "unknown")
	})

	t.Run("json read error via FS", func(t *testing.T) {
		f, err := readFile(errReadFS{err: errors.New("read boom")}, "x.json", AcceptAsJSON, nil)
		require.Error(t, err)
		require.Nil(t, f)
	})

	t.Run("open via FS failure", func(t *testing.T) {
		f, err := readFile(openErrFS{}, "x.ach", AcceptFile, nil)
		require.Error(t, err)
		require.Nil(t, f)
		require.Contains(t, err.Error(), "opening")
	})

	t.Run("stat failure", func(t *testing.T) {
		f, err := readFile(statErrFS{}, "x.ach", AcceptFile, nil)
		require.Error(t, err)
		require.Nil(t, f)
		require.Contains(t, err.Error(), "stat")
	})
}

// errReadFS opens files whose Read always fails.
type errReadFS struct{ err error }

func (e errReadFS) Open(name string) (fs.File, error) {
	return &errFile{err: e.err}, nil
}

// openErrFS always fails Open.
type openErrFS struct{}

func (openErrFS) Open(name string) (fs.File, error) {
	return nil, errors.New("no open")
}

type errFile struct{ err error }

func (e *errFile) Stat() (fs.FileInfo, error) {
	return fakeFileInfo{name: "x", size: 10}, nil
}
func (e *errFile) Read(p []byte) (int, error) { return 0, e.err }
func (e *errFile) Close() error               { return nil }

type fakeFileInfo struct {
	name string
	size int64
	dir  bool
}

func (f fakeFileInfo) Name() string       { return f.name }
func (f fakeFileInfo) Size() int64        { return f.size }
func (f fakeFileInfo) Mode() fs.FileMode  { return 0644 }
func (f fakeFileInfo) ModTime() time.Time { return time.Time{} }
func (f fakeFileInfo) IsDir() bool        { return f.dir }
func (f fakeFileInfo) Sys() any           { return nil }

// statErrFS returns a file whose Stat fails.
type statErrFS struct{}

func (statErrFS) Open(name string) (fs.File, error) { return &statErrFile{}, nil }

type statErrFile struct{}

func (statErrFile) Stat() (fs.FileInfo, error) { return nil, errors.New("stat boom") }
func (statErrFile) Read([]byte) (int, error)   { return 0, io.EOF }
func (statErrFile) Close() error               { return nil }

func TestReadValidateOptsFromFile_OSOpen(t *testing.T) {
	dir := t.TempDir()
	achPath := filepath.Join(dir, "f.ach")
	require.NoError(t, os.WriteFile(achPath, []byte("x"), 0644))

	// Missing sidecar → nil
	opts := readValidateOptsFromFile(achPath, &MergeDirOptions{
		ValidateOptsExtension: ".json",
	})
	require.Nil(t, opts)

	// Present sidecar via os.Open (FS == nil)
	require.NoError(t, os.WriteFile(filepath.Join(dir, "f.json"), []byte(`{"skipAll":true}`), 0644))
	opts = readValidateOptsFromFile(achPath, &MergeDirOptions{
		ValidateOptsExtension: ".json",
	})
	require.NotNil(t, opts)
	require.True(t, opts.SkipAll)

	// Empty extension → nil
	require.Nil(t, readValidateOptsFromFile(achPath, &MergeDirOptions{}))
}

func TestMergeDir_ParseWorkersOverride(t *testing.T) {
	dir := t.TempDir()
	src, err := os.Open(filepath.Join("test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)
	defer src.Close()
	dst, err := os.Create(filepath.Join(dir, "a.ach"))
	require.NoError(t, err)
	_, err = io.Copy(dst, src)
	require.NoError(t, err)
	require.NoError(t, dst.Close())

	merged, err := MergeDir(dir, Conditions{}, &MergeDirOptions{ParseWorkers: 1})
	require.NoError(t, err)
	require.Len(t, merged, 1)
}

func TestMergeDir_NoHangOnScatteredErrors(t *testing.T) {
	// Many valid files mixed with corrupt and ADV files: MergeDir must return an
	// error promptly (no deadlock) regardless of which worker hits the bad file.
	dir := t.TempDir()

	ppd, err := os.ReadFile(filepath.Join("test", "testdata", "ppd-debit.ach"))
	require.NoError(t, err)
	adv, err := os.ReadFile(filepath.Join("test", "testdata", "adv.ach"))
	require.NoError(t, err)

	for i := 0; i < 30; i++ {
		require.NoError(t, os.WriteFile(filepath.Join(dir, fmt.Sprintf("ok-%02d.ach", i)), ppd, 0644))
	}
	require.NoError(t, os.WriteFile(filepath.Join(dir, "bad-mid.ach"), []byte("%%%not-ach%%%"), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(dir, "adv-mid.ach"), adv, 0644))
	for i := 30; i < 40; i++ {
		require.NoError(t, os.WriteFile(filepath.Join(dir, fmt.Sprintf("ok-%02d.ach", i)), ppd, 0644))
	}

	done := make(chan error, 1)
	go func() {
		_, err := MergeDir(dir, Conditions{}, &MergeDirOptions{ParseWorkers: 8})
		done <- err
	}()

	select {
	case err := <-done:
		require.Error(t, err, "expected merge to fail on corrupt or ADV input")
	case <-time.After(10 * time.Second):
		t.Fatal("MergeDir hung when encountering errors among many files")
	}
}
