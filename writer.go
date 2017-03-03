// Copyright 2016 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"bufio"
	"io"
)

// A Writer writes an ach.file to a NACHA encoded file.
//
// As returned by NewWriter, a Writer writes ach.file structs into
// NACHA formted files.
//
type Writer struct {
	w *bufio.Writer
}

// NewWriter returns a new Writer that writes to w.
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		w: bufio.NewWriter(w),
	}
}

// Writer writes a single ach.file record to w
func (w *Writer) Write(file File) error {
	var err error
	// TODO: add ValidateAll to ach.file to recursively ensure we have a valid records
	if err = file.Validate(); err != nil {
		return err
	}

	// Iterate over all records in the file
	_, err = w.w.WriteString(file.Header.String())
	for _, batch := range file.Batches {
		_, err = w.w.WriteString(batch.Header.String())
		for _, entry := range batch.Entries {
			_, err = w.w.WriteString(entry.String())
			for _, addenda := range entry.Addendums {
				_, err = w.w.WriteString(addenda.String())
			}
		}
		_, err = w.w.WriteString(batch.Control.String())
	}
	_, err = w.w.WriteString(file.Control.String())

	if err != nil {
		return err
	}
	return nil
}

// Flush writes any buffered data to the underlying io.Writer.
// To check if an error occurred during the Flush, call Error.
func (w *Writer) Flush() {
	w.w.Flush()
}

// Error reports any error that has occurred during a previous Write or Flush.
func (w *Writer) Error() error {
	_, err := w.w.Write(nil)
	return err
}

// WriteAll writes multiple ach.fieles to w using Write and then calls Flush.
func (w *Writer) WriteAll(files []File) error {
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
