// Copyright 2017 The ACH Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"testing"
)

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

func TestFileError(t *testing.T) {
	err := &FileError{FieldName: "mock", Msg: "test message"}
	if err.Error() != "mock test message" {
		t.Error("FileError Error has changed formatting")
	}
}

// TestFileBatchCount if calculated count is different from control
func TestFileBatchCount(t *testing.T) {
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

func TestFileEntryAddenda(t *testing.T) {
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

func TestFileDebitAmount(t *testing.T) {
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

func TestFileCreditAmount(t *testing.T) {
	file := mockFilePPD()

	// inequality in total debit amount
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

func TestFileEntryHash(t *testing.T) {
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

func TestFileBlockCount10(t *testing.T) {
	file := NewFile().SetHeader(mockFileHeader())
	batch := NewBatchPPD()
	batch.SetHeader(mockBatchHeader())
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

func TestFileBuildBadFileHeader(t *testing.T) {
	file := NewFile().SetHeader(FileHeader{})
	if err := file.Create(); err != nil {
		if e, ok := err.(*FieldError); ok {
			if e.Msg != msgFieldInclusion {
				t.Errorf("%T: %s", err, err)
			}
		} else {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFileBuildNoBatch(t *testing.T) {
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

// Check if a file contains BatchNOC notification of change
func TestFileNotificationOfChange(t *testing.T) {
	file := NewFile().SetHeader(mockFileHeader())
	file.AddBatch(mockBatchPPD())
	bCOR := mockBatchCOR()
	file.AddBatch(bCOR)
	file.Create()

	if file.NotificationOfChange[0] != bCOR {
		t.Error("BatchCOR added to File.AddBatch should exist in NotificationOfChange")
	}
}

func TestFileReturnEntries(t *testing.T) {
	// create or copy the entry to be returned record
	entry := mockEntryDetail()
	// Add the addenda return with appropriate ReturnCode and addenda information
	entry.AddAddenda(mockAddendaReturn())
	// create or copy the previous batch header of the item being returned
	batchHeader := mockBatchHeader()
	// create or copy the batch to be returned
	batch, err := NewBatch(BatchParam{StandardEntryClass: batchHeader.StandardEntryClassCode})
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
