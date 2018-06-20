// Copyright 2018 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"bufio"
	"io"
	"strings"
)

// A Writer writes an ach.file to a NACHA encoded file.
//
// As returned by NewWriter, a Writer writes ach.file structs into
// NACHA formatted files.
//
type Writer struct {
	w       *bufio.Writer
	lineNum int //current line being written
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w: bufio.NewWriter(w),
	}
}

// Writer writes a single ach.file record to w
func (w *Writer) Write(file *File) error {
	if err := file.Validate(); err != nil {
		return err
	}

	w.lineNum = 0
	// Iterate over all records in the file
	w.w.WriteString(file.Header.String() + "\n")
	w.lineNum++

	for _, batch := range file.Batches {
		w.w.WriteString(batch.GetHeader().String() + "\n")
		w.lineNum++
		for _, entry := range batch.GetEntries() {
			w.w.WriteString(entry.String() + "\n")
			w.lineNum++
			for _, addenda := range entry.Addendum {
				w.w.WriteString(addenda.String() + "\n")
				w.lineNum++
			}
		}
		w.w.WriteString(batch.GetControl().String() + "\n")
		w.lineNum++
	}
	w.w.WriteString(file.Control.String() + "\n")
	w.lineNum++

	// pad the final block
	for i := 0; i < (10-(w.lineNum%10)) && w.lineNum%10 != 0; i++ {
		w.w.WriteString(strings.Repeat("9", 94) + "\n")
	}

	return nil
}

// Flush writes any buffered data to the underlying io.Writer.
// To check if an error occurred during the Flush, call Error.
func (w *Writer) Flush() {
	w.w.Flush()
}

// WriteAll writes multiple ach.files to w using Write and then calls Flush.
func (w *Writer) WriteAll(files []*File) error {
	for _, file := range files {
		err := w.Write(file)
		// TODO if one of the files errors at a Writer struct flag to decide if
		// the other files should still be written
		if err != nil {
			return err
		}
	}
	return w.w.Flush()
}
