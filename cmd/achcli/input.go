// Copyright 2026 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"io"
	"os"
	"sync"
)

var (
	errStdinEmpty = errors.New("stdin is empty")
	errStdinFix   = errors.New("stdin (-) is not supported with -fix")
)

// stdinCache reads stdin once and reuses the bytes, matching csvq's STDIN view.
type stdinCache struct {
	r    io.Reader
	once sync.Once
	data []byte
	err  error
}

var stdin = &stdinCache{r: os.Stdin}

func (s *stdinCache) read() ([]byte, error) {
	s.once.Do(func() {
		if !canReadStdin(s.r) {
			s.err = errStdinEmpty
			return
		}
		s.data, s.err = io.ReadAll(s.r)
		if s.err != nil {
			return
		}
		if len(s.data) == 0 {
			s.data = nil
			s.err = errStdinEmpty
		}
	})
	return s.data, s.err
}

// canReadStdin reports whether r has data to read without blocking on a TTY.
// Non-file readers (tests) are readable. Files are readable when they are a
// named pipe or a redirected file with size > 0, the same check csvq uses.
func canReadStdin(r io.Reader) bool {
	if r == nil {
		return false
	}
	f, ok := r.(*os.File)
	if !ok {
		return true
	}
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeNamedPipe != 0 || fi.Size() > 0
}

// readInput reads a file, or stdin when path is "-". Stdin is cached so a
// second "-" reuses the same bytes.
func readInput(path string) ([]byte, error) {
	if path != "-" {
		return os.ReadFile(path)
	}
	return stdin.read()
}

func rejectStdinFix(path string) error {
	if path == "-" {
		return errStdinFix
	}
	return nil
}
