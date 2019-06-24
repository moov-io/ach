// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"path/filepath"
	"testing"
)

func filesAreEqual(f1, f2 *File) error {
	// File Header
	if f1.Header.ImmediateOrigin != f2.Header.ImmediateOrigin {
		return fmt.Errorf("f1.Header.ImmediateOrigin=%s vs f2.Header.ImmediateOrigin=%s", f1.Header.ImmediateOrigin, f2.Header.ImmediateOrigin)
	}
	if f1.Header.ImmediateDestination != f2.Header.ImmediateDestination {
		return fmt.Errorf("f1.Header.ImmediateDestination=%s vs f2.Header.ImmediateDestination=%s", f1.Header.ImmediateDestination, f2.Header.ImmediateDestination)
	}

	// Batches
	if len(f1.Batches) != len(f2.Batches) {
		return fmt.Errorf("len(f1.Batches)=%d vs len(f2.Batches)=%d", len(f1.Batches), len(f2.Batches))
	}
	for i := range f1.Batches {
		for j := range f2.Batches {
			if f1.Batches[i].Equal(f2.Batches[j]) {
				goto next
			}
		}
		return fmt.Errorf("unable to find batch in f2: %v", f1.Batches[i])
	next:
		// check the next batch
	}

	// IATBatches
	if len(f1.IATBatches) != len(f2.IATBatches) {
		return fmt.Errorf("len(f1.IATBatches)=%d vs len(f2.IATBatches)=%d", len(f1.IATBatches), len(f2.IATBatches))
	}

	// File Control
	if f1.Control.EntryAddendaCount != f2.Control.EntryAddendaCount {
		return fmt.Errorf("f1.Control.EntryAddendaCount=%d vs f2.Control.EntryAddendaCount=%d", f1.Control.EntryAddendaCount, f2.Control.EntryAddendaCount)
	}
	if f1.Control.TotalDebitEntryDollarAmountInFile != f2.Control.TotalDebitEntryDollarAmountInFile {
		return fmt.Errorf("f1.Control.TotalDebitEntryDollarAmountInFile=%d vs f2.Control.TotalDebitEntryDollarAmountInFile=%d", f1.Control.TotalDebitEntryDollarAmountInFile, f2.Control.TotalDebitEntryDollarAmountInFile)
	}
	if f1.Control.TotalCreditEntryDollarAmountInFile != f2.Control.TotalCreditEntryDollarAmountInFile {
		return fmt.Errorf("f1.Control.TotalCreditEntryDollarAmountInFile=%d vs f2.Control.TotalCreditEntryDollarAmountInFile=%d", f1.Control.TotalCreditEntryDollarAmountInFile, f2.Control.TotalCreditEntryDollarAmountInFile)
	}

	return nil
}

func TestMergeFiles__filesAreEqual(t *testing.T) {
	file, err := readACHFilepath(filepath.Join("test", "testdata", "ppd-debit.ach"))
	if err != nil {
		t.Fatal(err)
	}

	// compare a file against itself
	if err := filesAreEqual(file, file); err != nil {
		t.Fatalf("same file: %v", err)
	}

	// break the equality
	f2 := *file
	f2.Header.ImmediateOrigin = "12"
	if err := filesAreEqual(file, &f2); err == nil {
		t.Fatal("expected error")
	}
}

func TestMergeFiles__identity(t *testing.T) {
	file, err := readACHFilepath(filepath.Join("test", "testdata", "ppd-debit.ach"))
	if err != nil {
		t.Fatal(err)
	}

	out, err := MergeFiles([]*File{file})
	if err != nil {
		t.Fatal(err)
	}
	if len(out) != 1 {
		t.Errorf("got %d merged ACH files", len(out))
	}
	if err := filesAreEqual(file, out[0]); err != nil {
		t.Errorf("unequal files:%v", err)
	}
}

func TestMergeFiles__together(t *testing.T) {
	f1, err := readACHFilepath(filepath.Join("test", "testdata", "ppd-debit.ach"))
	if err != nil {
		t.Fatal(err)
	}
	f2, err := readACHFilepath(filepath.Join("test", "testdata", "web-debit.ach"))
	if err != nil {
		t.Fatal(err)
	}
	f2.Header = f1.Header // replace Header so they're merged into one file

	if len(f1.Batches) != 1 || len(f2.Batches) != 3 {
		t.Errorf("did batch counts change? f1:%d f2:%d", len(f1.Batches), len(f2.Batches))
	}

	out, err := MergeFiles([]*File{f1, f2})
	if err != nil {
		t.Fatal(err)
	}
	if len(out) != 1 {
		t.Errorf("got %d merged ACH files", len(out))
	}
	if len(out[0].Batches) != 4 {
		t.Errorf("got %d batches", len(out[0].Batches))
	}
}

func TestMergeFiles__apart(t *testing.T) {
	f1, err := readACHFilepath(filepath.Join("test", "testdata", "ppd-debit.ach"))
	if err != nil {
		t.Fatal(err)
	}
	f2, err := readACHFilepath(filepath.Join("test", "testdata", "web-debit.ach"))
	if err != nil {
		t.Fatal(err)
	}

	out, err := MergeFiles([]*File{f1, f2})
	if err != nil {
		t.Fatal(err)
	}
	if len(out) != 2 {
		t.Errorf("got %d merged ACH files", len(out))
	}
	if len(out[0].Batches) != 1 {
		t.Errorf("got %d batches", len(out[0].Batches))
	}
	if len(out[1].Batches) != 3 {
		t.Errorf("got %d batches", len(out[1].Batches))
	}
}

func TestMergeFiles__lineCount(t *testing.T) {
	file, err := readACHFilepath(filepath.Join("test", "testdata", "ppd-debit.ach"))
	if err != nil {
		t.Fatal(err)
	}
	if err := file.Create(); err != nil {
		t.Fatal(err)
	}

	if n, err := lineCount(file); n != 1 || err != nil {
		// We've optimized small file line counts to bypass writing out the file
		// into plain text as it's costly.
		t.Errorf("did we change optimizations? n=%d error=%v", n, err)
	}

	// Add 100 batches to file and get a real line count
	for i := 0; i < 100; i++ {
		file.AddBatch(file.Batches[0])
	}
	if err := file.Create(); err != nil {
		t.Fatal(err)
	}
	if n, err := lineCount(file); n != 310 || err != nil {
		t.Errorf("unexpected line count of %d: %v", n, err)
	}

	// make the file invalid and ensure we error
	file.Control.BatchCount = 0
	if n, err := lineCount(file); n != 0 || err == nil {
		t.Errorf("expected error n=%d error=%v", n, err)
	}
}

// TestMergeFiles__splitFiles generates a file over the 10k line limit and attempts to merge
// another file into it only to come away with two files after merging.
func TestMergeFiles__splitFiles(t *testing.T) {
	file, err := readACHFilepath(filepath.Join("test", "testdata", "ppd-debit.ach"))
	if err != nil {
		t.Fatal(err)
	}
	file.Control = NewFileControl()
	if err := file.Create(); err != nil {
		t.Fatal(err)
	}
	if len(file.Batches) != 1 {
		t.Fatalf("unexpected batch count of %d", len(file.Batches))
	}

	// Add a bunch of batches so it's over the line limit
	// somewhere between 3-4k Batches exceed the 10k line limit
	for i := 0; i < 4000; i++ {
		file.AddBatch(file.Batches[0])
	}
	if err := file.Create(); err != nil {
		t.Fatal(err)
	}

	// Read another file to merge
	f2, err := readACHFilepath(filepath.Join("test", "testdata", "web-debit.ach"))
	if err != nil {
		t.Fatal(err)
	}
	f2.Header = file.Header // replace Header so they're merged into one file
	if err := f2.Create(); err != nil {
		t.Fatal(err)
	}

	// read a third file
	f3, err := readACHFilepath(filepath.Join("test", "testdata", "20110805A.ach"))
	if err != nil {
		t.Fatal(err)
	}
	f3.Header = file.Header // replace Header so they're merged into one file
	if err := f3.Create(); err != nil {
		t.Fatal(err)
	}

	// Merge our big file into another file and verify we get two back
	// TODO(adam): We should probably recurse back on `file` to ensure we don't exceed the
	// 10k line limit. That shouldn't happen as MergeFiles processes one batch at a time, but
	// an incoming file might be invalid in that way.
	out, err := MergeFiles([]*File{file, f2, f3})
	if err != nil || len(out) != 2 {
		t.Fatalf("got %d files, error=%v", len(out), err)
	}
	if len(out[0].Batches) != 4001 || len(out[1].Batches) != 5 {
		// These batch counts will change when we recurse back through out[0]
		// so it doesn't exceed the 10k line limit.
		t.Errorf("out[0].Batches:%d out[1].Batches:%d", len(out[0].Batches), len(out[1].Batches))
	}
}

func TestMergeFiles__invalid(t *testing.T) {
	f1, err := readACHFilepath(filepath.Join("test", "testdata", "ppd-debit.ach"))
	if err != nil {
		t.Fatal(err)
	}
	f1.Header.ImmediateOrigin = "0000000000" // make file invalid

	f2, err := readACHFilepath(filepath.Join("test", "testdata", "web-debit.ach"))
	if err != nil {
		t.Fatal(err)
	}
	f2.Header = f1.Header

	out, err := MergeFiles([]*File{f1, f2})
	if len(out) != 0 || err == nil {
		t.Errorf("expected error: len(out)=%d error=%v", len(out), err)
	}
}
