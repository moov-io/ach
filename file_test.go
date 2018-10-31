// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

// mockFilePPD creates an ACH file with PPD batch and entry
func mockFilePPD() *File {
	mockFile := NewFile()
	mockFile.SetHeader(mockFileHeader())
	mockBatch := mockBatchPPD()
	mockFile.AddBatch(mockBatch)
	if err := mockFile.Create(); err != nil {
		panic(err)
	}
	return mockFile
}

// testFileError validates a file error
func testFileError(t testing.TB) {
	err := &FileError{FieldName: "mock", Msg: "test message"}
	if err.Error() != "mock test message" {
		t.Error("FileError Error has changed formatting")
	}
}

// TestFileError tests validating a file error
func TestFileError(t *testing.T) {
	testFileError(t)
}

// BenchmarkFileError benchmarks validating a file error
func BenchmarkFileError(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileError(b)
	}
}

// TestFileEmptyError tests an empty file error
func TestFileEmptyError(t *testing.T) {
	file := &File{}
	if err := file.Create(); err == nil {
		t.Error("expected error")
	}
	err := file.Validate()
	msg := err.Error()
	if !strings.HasPrefix(msg, "recordType") || !strings.Contains(msg, "is a mandatory field") {
		t.Errorf("got %q", err)
	}
}

// testFileBatchCount validates if calculated count is different from control
func testFileBatchCount(t testing.TB) {
	file := mockFilePPD()

	// More batches than the file control count.
	file.AddBatch(mockBatchPPD())
	if err := file.Validate(); err != nil {
		if e, ok := err.(*FileError); ok {
			if e.FieldName != "BatchCount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFileBatchCount tests validating if calculated count is different from control
func TestFileBatchCount(t *testing.T) {
	testFileBatchCount(t)
}

// BenchmarkFileBatchCount benchmarks validating if calculated count is different from control
func BenchmarkFileBatchCount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileBatchCount(b)
	}
}

// testFileEntryAddenda validates an addenda entry
func testFileEntryAddenda(t testing.TB) {
	file := mockFilePPD()

	// more entries than the file control
	file.Control.EntryAddendaCount = 5
	if err := file.Validate(); err != nil {
		if e, ok := err.(*FileError); ok {
			if e.FieldName != "EntryAddendaCount" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFileEntryAddenda tests validating an addenda entry
func TestFileEntryAddenda(t *testing.T) {
	testFileEntryAddenda(t)
}

// BenchmarkFileEntryAddenda benchmarks validating an addenda entry
func BenchmarkFileEntryAddenda(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileEntryAddenda(b)
	}
}

// testFileDebitAmount validates file total debit amount
func testFileDebitAmount(t testing.TB) {
	file := mockFilePPD()

	// inequality in total debit amount
	file.Control.TotalDebitEntryDollarAmountInFile = 63
	if err := file.Validate(); err != nil {
		if e, ok := err.(*FileError); ok {
			if e.FieldName != "TotalDebitEntryDollarAmountInFile" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFileDebitAmount tests validating file total debit amount
func TestFileDebitAmount(t *testing.T) {
	testFileDebitAmount(t)
}

// BenchmarkFileDebitAmount benchmarks validating file total debit amount
func BenchmarkFileDebitAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileDebitAmount(b)
	}
}

// testFileCreditAmount validates file total credit amount
func testFileCreditAmount(t testing.TB) {
	file := mockFilePPD()

	// inequality in total credit amount
	file.Control.TotalCreditEntryDollarAmountInFile = 63
	if err := file.Validate(); err != nil {
		if e, ok := err.(*FileError); ok {
			if e.FieldName != "TotalCreditEntryDollarAmountInFile" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFileCreditAmount tests validating file total credit amount
func TestFileCreditAmount(t *testing.T) {
	testFileCreditAmount(t)
}

// BenchmarkFileCreditAmount benchmarks validating file total credit amount
func BenchmarkFileCreditAmount(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileCreditAmount(b)
	}
}

// testFileEntryHash validates entry hash
func testFileEntryHash(t testing.TB) {
	file := mockFilePPD()
	file.AddBatch(mockBatchPPD())
	file.Create()
	file.Control.EntryHash = 63
	if err := file.Validate(); err != nil {
		if e, ok := err.(*FileError); ok {
			if e.FieldName != "EntryHash" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFileEntryHash tests validating entry hash
func TestFileEntryHash(t *testing.T) {
	testFileEntryHash(t)
}

// BenchmarkFileEntryHash benchmarks validating entry hash
func BenchmarkFileEntryHash(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileEntryHash(b)
	}
}

// testFileBlockCount10 validates file block count
func testFileBlockCount10(t testing.TB) {
	file := NewFile().SetHeader(mockFileHeader())
	batch := NewBatchPPD(mockBatchPPDHeader())
	batch.AddEntry(mockEntryDetail())
	batch.AddEntry(mockEntryDetail())
	batch.AddEntry(mockEntryDetail())
	batch.AddEntry(mockEntryDetail())
	batch.AddEntry(mockEntryDetail())
	batch.AddEntry(mockEntryDetail())
	batch.Create()
	file.AddBatch(batch)
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	// ensure with 10 records in file we don't get 2 for a block count
	if file.Control.BlockCount != 1 {
		t.Error("BlockCount on 10 records is not equal to 1")
	}
	// make 11th record which should produce BlockCount of 2
	file.Batches[0].AddEntry(mockEntryDetail())
	file.Batches[0].Create() // File.Build does not re-build Batches
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if file.Control.BlockCount != 2 {
		t.Error("BlockCount on 11 records is not equal to 2")
	}
}

// TestFileBlockCount10 tests validating file block count
func TestFileBlockCount10(t *testing.T) {
	testFileBlockCount10(t)
}

// BenchmarkFileBlockCount10 benchmarks validating file block count
func BenchmarkFileBlockCount10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileBlockCount10(b)
	}
}

// testFileBuildBadFileHeader validates a bad file header
func testFileBuildBadFileHeader(t testing.TB) {
	file := NewFile().SetHeader(FileHeader{})
	if err := file.Create(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if !strings.Contains(e.Msg, msgFieldInclusion) {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFileBuildBadFileHeader tests validating a bad file header
func TestFileBuildBadFileHeader(t *testing.T) {
	testFileBuildBadFileHeader(t)
}

// BenchmarkFileBuildBadFileHeader benchmarks validating a bad file header
func BenchmarkFileBuildBadFileHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileBuildBadFileHeader(b)
	}
}

// testFileBuildNoBatch validates a file with no batches
func testFileBuildNoBatch(t testing.TB) {
	file := NewFile().SetHeader(mockFileHeader())
	if err := file.Create(); err != nil {
		if e, ok := err.(*FileError); ok {
			if e.FieldName != "Batches" {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFileBuildNoBatch tests validating a file with no batches
func TestFileBuildNoBatch(t *testing.T) {
	testFileBuildNoBatch(t)
}

// BenchmarkFileBuildNoBatch benchmarks validating a file with no batches
func BenchmarkFileBuildNoBatch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileBuildNoBatch(b)
	}
}

// testFileNotificationOfChange validates if a file contains
// BatchNOC notification of change
func testFileNotificationOfChange(t testing.TB) {
	file := NewFile().SetHeader(mockFileHeader())
	file.AddBatch(mockBatchPPD())
	bCOR := mockBatchCOR()
	file.AddBatch(bCOR)
	file.Create()

	if file.NotificationOfChange[0] != bCOR {
		t.Error("BatchCOR added to File.AddBatch should exist in NotificationOfChange")
	}
}

// TestFileNotificationOfChange tests validating if a file contains
// BatchNOC notification of change
func TestFileNotificationOfChange(t *testing.T) {
	testFileNotificationOfChange(t)
}

// BenchmarkFileNotificationOfChange benchmarks validating if a file contains
// BatchNOC notification of change
func BenchmarkFileNotificationOfChange(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileNotificationOfChange(b)
	}
}

// testFileReturnEntries validates file return entries
func testFileReturnEntries(t testing.TB) {
	// create or copy the entry to be returned record
	entry := mockEntryDetail()
	// Add the addenda return with appropriate ReturnCode and addenda information
	entry.AddAddenda(mockAddenda99())
	// create or copy the previous batch header of the item being returned
	batchHeader := mockBatchHeader()
	// create or copy the batch to be returned

	//batch, err := NewBatch(BatchParam{StandardEntryClass: batchHeader.StandardEntryClassCode})
	batch, err := NewBatch(batchHeader)

	if err != nil {
		t.Error(err.Error())
	}
	// Add the entry to be returned to the batch
	batch.AddEntry(entry)
	// Create the batch
	batch.Create()
	// Add the batch to your file.
	file := NewFile().SetHeader(mockFileHeader())
	file.AddBatch(batch)
	// Create the return file
	if err := file.Create(); err != nil {
		t.Error(err.Error())
	}

	if len(file.ReturnEntries) != 1 {
		t.Errorf("1 file.ReturnEntries added and %v exist", len(file.ReturnEntries))
	}
}

// TestFileReturnEntries tests validating file return entries
func TestFileReturnEntries(t *testing.T) {
	testFileReturnEntries(t)
}

// BenchmarkFileReturnEntries benchmarks validating file return entries
func BenchmarkFileReturnEntries(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileReturnEntries(b)
	}
}

func TestFile__readFromJson(t *testing.T) {
	path := filepath.Join("test", "testdata", "ppd-valid.json")
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	file, err := FileFromJson(bs)
	if err != nil {
		t.Fatal(err)
	}

	if file.ID != "adam-01" {
		t.Errorf("file.ID: %s", file.ID)
	}

	// Header
	if file.Header.ImmediateOrigin != "121042882" || file.Header.ImmediateOriginName != "Wells Fargo" {
		t.Errorf("origin=%s name=%s", file.Header.ImmediateOrigin, file.Header.ImmediateOriginName)
	}
	if file.Header.ImmediateDestination != "231380104" || file.Header.ImmediateDestinationName != "Citadel" {
		t.Errorf("destination=%s name=%s", file.Header.ImmediateDestination, file.Header.ImmediateDestinationName)
	}
	if file.Header.FileCreationTime.IsZero() || file.Header.FileCreationDate.IsZero() {
		t.Errorf("time=%v date=%v", file.Header.FileCreationTime, file.Header.FileCreationDate)
	}

	// Batches
	if len(file.Batches) != 1 {
		t.Errorf("got %d batches: %v", len(file.Batches), file.Batches)
	}

	// Control
	if file.Control.BatchCount != 1 {
		t.Errorf("BatchCount: %d", file.Control.BatchCount)
	}
	if file.Control.TotalDebitEntryDollarAmountInFile != 0 || file.Control.TotalCreditEntryDollarAmountInFile != 100000 {
		t.Errorf("debit=%d credit=%d", file.Control.TotalDebitEntryDollarAmountInFile, file.Control.TotalCreditEntryDollarAmountInFile)
	}

	// ensure we error on struct tag unmarshal
	var f File
	err = json.Unmarshal(bs, &f)
	if !strings.Contains(err.Error(), "use ach.FileFromJSON instead") {
		t.Error("expected error, see FileFromJson definition")
	}

	if err := file.Validate(); err != nil {
		t.Error(err)
	}
}
