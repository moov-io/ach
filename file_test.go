// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/base"
)

// mockFilePPD creates an ACH file with PPD batch and entry
func mockFilePPD() *File {
	mockFile := NewFile()
	mockFile.ID = "fileId"
	mockFile.SetHeader(mockFileHeader())
	mockFile.Header.ID = mockFile.ID
	mockFile.Control = mockFileControl()
	mockFile.Control.ID = mockFile.ID
	mockBatch := mockBatchPPD()
	mockFile.AddBatch(mockBatch)
	if err := mockFile.Create(); err != nil {
		panic(err)
	}
	if mockFile.ID != mockFile.Header.ID {
		panic(fmt.Sprintf("mockFile.ID=%s mockFile.Header.ID=%s", mockFile.ID, mockFile.Header.ID))
	}
	if mockFile.ID != mockFile.Control.ID {
		panic(fmt.Sprintf("mockFile.ID=%s mockFile.Control.ID=%s", mockFile.ID, mockFile.Control.ID))
	}
	return mockFile
}

func mockFileADV() *File {
	mockFile := NewFile()
	mockFile.ID = "fileId"
	mockFile.SetHeader(mockFileHeader())
	mockFile.Header.ID = mockFile.ID
	mockFile.ADVControl = mockADVFileControl()
	mockFile.Control.ID = mockFile.ID
	mockBatchADV := mockBatchADV()
	mockFile.AddBatch(mockBatchADV)
	if err := mockFile.Create(); err != nil {
		panic(err)
	}
	if mockFile.ID != mockFile.Header.ID {
		panic(fmt.Sprintf("mockFile.ID=%s mockFile.Header.ID=%s", mockFile.ID, mockFile.Header.ID))
	}
	if mockFile.ID != mockFile.Control.ID {
		panic(fmt.Sprintf("mockFile.ID=%s mockFile.Control.ID=%s", mockFile.ID, mockFile.Control.ID))
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
	err := file.Validate()
	if err != NewErrFileCalculatedControlEquality("BatchCount", 2, 1) {
		t.Errorf("%T: %s", err, err)
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
	err := file.Validate()
	if err != NewErrFileCalculatedControlEquality("EntryAddendaCount", 1, 5) {
		t.Errorf("%T: %s", err, err)
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
	err := file.Validate()
	if err != NewErrFileCalculatedControlEquality("TotalDebitEntryDollarAmountInFile", 0, 63) {
		t.Errorf("%T: %s", err, err)
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
	err := file.Validate()
	if err != NewErrFileCalculatedControlEquality("TotalCreditEntryDollarAmountInFile", 100000000, 63) {
		t.Errorf("%T: %s", err, err)
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
	if err := file.Create(); err != nil {
		t.Fatal(err)
	}
	file.Control.EntryHash = 63
	err := file.Validate()
	if err != NewErrFileCalculatedControlEquality("EntryHash", 46276020, 63) {
		t.Errorf("%T: %s", err, err)
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

	ed1 := mockEntryDetail()
	ed1.SetTraceNumber(mockBatchHeader().ODFIIdentification, 1)
	batch.AddEntry(ed1)
	ed2 := mockEntryDetail()
	ed2.SetTraceNumber(mockBatchHeader().ODFIIdentification, 2)
	batch.AddEntry(ed2)
	ed3 := mockEntryDetail()
	ed3.SetTraceNumber(mockBatchHeader().ODFIIdentification, 3)
	batch.AddEntry(ed3)
	ed4 := mockEntryDetail()
	ed4.SetTraceNumber(mockBatchHeader().ODFIIdentification, 4)
	batch.AddEntry(ed4)
	ed5 := mockEntryDetail()
	ed5.SetTraceNumber(mockBatchHeader().ODFIIdentification, 5)
	batch.AddEntry(ed5)
	ed6 := mockEntryDetail()
	ed6.SetTraceNumber(mockBatchHeader().ODFIIdentification, 6)
	batch.AddEntry(ed6)

	if err := batch.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	file.AddBatch(batch)
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	// ensure with 10 records in file we don't get 2 for a block count
	if file.Control.BlockCount != 1 {
		t.Error("BlockCount on 10 records is not equal to 1")
	}
	// make 11th record which should produce BlockCount of 2
	ed7 := mockEntryDetail()
	ed7.SetTraceNumber(mockBatchHeader().ODFIIdentification, 7)
	file.Batches[0].AddEntry(ed7)

	if err := file.Batches[0].Create(); err != nil {
		t.Fatal(err)
	} // File.Build does not re-build Batches
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
	err := file.Create()
	if !base.Match(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
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
	err := file.Create()
	if !base.Match(err, ErrFileNoBatches) {
		t.Errorf("%T: %s", err, err)
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
	mockBatch := NewBatchPPD(mockBatchPPDHeader())
	mockBatch.AddEntry(mockPPDEntryDetailNOC())
	file.AddBatch(mockBatch)
	if err := file.Create(); err != nil {
		t.Fatal(err)
	}
	if file.NotificationOfChange[0] != mockBatch {
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
	entry.AddendaRecordIndicator = 1
	// Add the addenda return with appropriate ReturnCode and addenda information
	entry.Addenda99 = mockAddenda99()
	entry.Category = CategoryReturn
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
	if err := batch.Create(); err != nil {
		t.Fatal(err)
	}
	// Add the batch to your file.
	file := NewFile().SetHeader(mockFileHeader())
	file.AddBatch(batch)
	// Create the return file
	if err := file.Create(); err != nil {
		t.Error(err.Error())
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

	file, err := FileFromJSON(bs)
	if err != nil {
		t.Fatal(err)
	}

	// Ensure the file is valid
	if err := file.Create(); err != nil {
		t.Error(err)
	}
	if err := file.Validate(); err != nil {
		t.Error(err)
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
	if file.Header.FileCreationTime == "" || file.Header.FileCreationDate == "" {
		t.Errorf("time=%v date=%v", file.Header.FileCreationTime, file.Header.FileCreationDate)
	}

	// Batches
	if len(file.Batches) != 1 {
		t.Errorf("got %d batches: %v", len(file.Batches), file.Batches)
	}
	batch := file.Batches[0]
	batchControl := batch.GetControl()
	if batchControl.EntryAddendaCount != 1 {
		t.Errorf("EntryAddendaCount: %d", batchControl.EntryAddendaCount)
	}

	// Control
	if file.Control.BatchCount != 1 {
		t.Errorf("BatchCount: %d", file.Control.BatchCount)
	}
	if file.Control.EntryAddendaCount != 1 {
		t.Errorf("File Control EntryAddendaCount: %d", file.Control.EntryAddendaCount)
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

// TestFile__jsonFileNoControlBlobs will read an ach.File from its JSON form, but the JSON has no
// batchControl or fileControl sub-objects.
func TestFile__jsonFileNoControlBlobs(t *testing.T) {
	path := filepath.Join("test", "testdata", "ppd-no-control-blobs-valid.json")
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	file, err := FileFromJSON(bs)
	if err != nil {
		t.Fatal(err)
	}

	if err := file.Create(); err != nil {
		t.Fatal(err)
	}
	if err := file.Validate(); err != nil {
		t.Fatal(err)
	}

	if file.ID != "adam-01" {
		t.Errorf("file.ID: %s", file.ID)
	}
}

func TestFileADV__Success(t *testing.T) {
	fh := mockFileHeader()
	bh := mockBatchADVHeader()

	entryOne := mockADVEntryDetail()
	entryTwo := mockADVEntryDetail()
	entryTwo.SequenceNumber = 2

	// build the batch
	batch := NewBatchADV(bh)
	batch.AddADVEntry(entryOne)
	batch.AddADVEntry(entryTwo)
	if err := batch.Create(); err != nil {
		t.Fatalf("Unexpected error building batch: %s\n", err)
	}

	// build the file
	file := NewFile()
	file.SetHeader(fh)
	file.AddBatch(batch)
	if err := file.Create(); err != nil {
		t.Fatalf("Unexpected error building file: %s\n", err)
	}
}

func TestFileADVInvalid__StandardEntryClassCode(t *testing.T) {
	fh := mockFileHeader()

	// ADV
	bhADV := mockBatchADVHeader()
	entryADV := mockADVEntryDetail()

	// build the ADV batch
	batchADV := NewBatchADV(bhADV)
	batchADV.AddADVEntry(entryADV)

	if err := batchADV.Create(); err != nil {
		t.Fatalf("Unexpected error building batch: %s\n", err)
	}

	// PPD
	bhPPD := mockBatchPPDHeader()
	entryPPD := mockPPDEntryDetail()

	// build the PPD batch
	batchPPD, err := NewBatch(bhPPD)
	if err != nil {
		t.Fatalf("Unexpected error with NewBatch: %s\n", err)
	}
	batchPPD.AddEntry(entryPPD)

	if err := batchPPD.Create(); err != nil {
		t.Fatalf("Unexpected error building batch: %s\n", err)
	}

	// build the file
	file := NewFile()
	file.SetHeader(fh)
	file.AddBatch(batchADV)
	file.AddBatch(batchPPD)
	err = file.Create()
	if err != ErrFileADVOnly {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileADVEntryHash validates entry hash
func TestFileADVEntryHash(t *testing.T) {
	file := mockFileADV()
	file.AddBatch(mockBatchADV())
	if err := file.Create(); err != nil {
		t.Fatal(err)
	}
	file.ADVControl.EntryHash = 63
	err := file.Validate()
	if err != NewErrFileCalculatedControlEquality("EntryHash", 46276020, 63) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileADVDebitAmount validates file total debit amount
func TestFileADVDebitAmount(t *testing.T) {
	file := mockFileADV()

	// inequality in total debit amount
	file.ADVControl.TotalDebitEntryDollarAmountInFile = 06
	err := file.Validate()
	if err != NewErrFileCalculatedControlEquality("TotalDebitEntryDollarAmountInFile", 0, 6) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileADVCreditAmount validates file total credit amount
func TestFileADVCreditAmount(t *testing.T) {
	file := mockFileADV()

	// inequality in total credit amount
	file.ADVControl.TotalCreditEntryDollarAmountInFile = 07
	err := file.Validate()
	if err != NewErrFileCalculatedControlEquality("TotalCreditEntryDollarAmountInFile", 50000, 7) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileADVEntryAddenda validates an addenda entry
func TestFileADVEntryAddenda(t *testing.T) {
	file := mockFileADV()

	// more entries than the file control
	file.ADVControl.EntryAddendaCount = 5
	err := file.Validate()
	if err != NewErrFileCalculatedControlEquality("EntryAddendaCount", 1, 5) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileADVBatchCount validates if calculated count is different from control
func TestFileADVBatchCount(t *testing.T) {
	file := mockFileADV()

	// More batches than the file control count.
	file.AddBatch(mockBatchADV())
	err := file.Validate()
	if err != NewErrFileCalculatedControlEquality("BatchCount", 2, 1) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileADVBlockCount10 validates file block count
func TestFileADVBlockCount10(t *testing.T) {
	file := NewFile().SetHeader(mockFileHeader())
	batchADV := NewBatchADV(mockBatchADVHeader())

	ed1 := mockADVEntryDetail()
	batchADV.AddADVEntry(ed1)

	ed2 := mockADVEntryDetail()
	ed2.SequenceNumber = 2
	batchADV.AddADVEntry(ed2)

	ed3 := mockADVEntryDetail()
	ed3.SequenceNumber = 3
	batchADV.AddADVEntry(ed3)

	ed4 := mockADVEntryDetail()
	ed4.SequenceNumber = 4
	batchADV.AddADVEntry(ed4)

	ed5 := mockADVEntryDetail()
	ed5.SequenceNumber = 5
	batchADV.AddADVEntry(ed5)

	ed6 := mockADVEntryDetail()
	ed6.SequenceNumber = 6
	batchADV.AddADVEntry(ed6)

	if err := batchADV.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	file.AddBatch(batchADV)
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	// ensure with 10 records in file we don't get 2 for a block count
	if file.ADVControl.BlockCount != 1 {
		t.Error("BlockCount on 10 records is not equal to 1")
	}
	// make 11th record which should produce BlockCount of 2
	ed7 := mockADVEntryDetail()
	ed7.SequenceNumber = 7
	file.Batches[0].AddADVEntry(ed7)

	if err := file.Batches[0].Create(); err != nil {
		t.Fatal(err)
	} // File.Build does not re-build Batches
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if file.ADVControl.BlockCount != 2 {
		t.Error("BlockCount on 11 records is not equal to 2")
	}
}

// TestFileADVControlValidate validates ADV File Control
func TestFileADVControlValidate(t *testing.T) {
	file := mockFileADV()

	file.ADVControl.recordType = "22"
	err := file.Validate()
	// TODO: are we expecting there to be an error here?
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileControlValidate validates PPD File Control
func TestFileControlValidate(t *testing.T) {
	file := mockFilePPD()

	file.Control.recordType = "22"
	err := file.Validate()
	// TODO: are we expecting there to be an error here?
	if !base.Match(err, nil) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestBatchHeaderNil Batch Header Nil
func TestBatchHeaderNil(t *testing.T) {
	fh := mockFileHeader()
	bh := mockBatchPPDHeader()

	entryOne := mockPPDEntryDetail()

	// build the batch
	batch := NewBatchPPD(bh)
	batch.AddEntry(entryOne)

	if err := batch.Create(); err != nil {
		t.Fatalf("Unexpected error building batch: %s\n", err)
	}

	batch.Header = nil

	// build the file
	file := NewFile()
	file.SetHeader(fh)
	file.AddBatch(batch)
	if err := file.Create(); err != nil {
		t.Fatalf("Unexpected error building file: %s\n", err)
	}

}

// TestBatchControlrNil Batch Control Nil
func TestBatchControlNil(t *testing.T) {
	fh := mockFileHeader()
	bh := mockBatchPPDHeader()

	entryOne := mockPPDEntryDetail()

	// build the batch
	batch := NewBatchPPD(bh)
	batch.AddEntry(entryOne)

	if err := batch.Create(); err != nil {
		t.Fatalf("Unexpected error building batch: %s\n", err)
	}

	batch.Control = nil

	// build the file
	file := NewFile()
	file.SetHeader(fh)
	file.AddBatch(batch)
	if err := file.Create(); err != nil {
		t.Fatalf("Unexpected error building file: %s\n", err)
	}

}

func TestFileADV__readFromJson(t *testing.T) {
	path := filepath.Join("test", "testdata", "adv-valid.json")
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	file, err := FileFromJSON(bs)
	if err != nil {
		t.Fatal(err)
	}

	// Ensure the file is valid
	if err := file.Create(); err != nil {
		t.Error(err)
	}
	if err := file.Validate(); err != nil {
		t.Error(err)
	}

	if file.ID != "adv-01" {
		t.Errorf("file.ID: %s", file.ID)
	}

	// Header
	if file.Header.ImmediateOrigin != "121042882" || file.Header.ImmediateOriginName != "Wells Fargo" {
		t.Errorf("origin=%s name=%s", file.Header.ImmediateOrigin, file.Header.ImmediateOriginName)
	}
	if file.Header.ImmediateDestination != "231380104" || file.Header.ImmediateDestinationName != "Citadel" {
		t.Errorf("destination=%s name=%s", file.Header.ImmediateDestination, file.Header.ImmediateDestinationName)
	}
	if file.Header.FileCreationTime == "" || file.Header.FileCreationDate == "" {
		t.Errorf("time=%v date=%v", file.Header.FileCreationTime, file.Header.FileCreationDate)
	}

	// Batches
	if len(file.Batches) != 1 {
		t.Errorf("got %d batches: %v", len(file.Batches), file.Batches)
	}
	batch := file.Batches[0]
	batchADVControl := batch.GetADVControl()
	if batchADVControl.EntryAddendaCount != 1 {
		t.Errorf("EntryAddendaCount: %d", batchADVControl.EntryAddendaCount)
	}

	// Control
	if file.ADVControl.BatchCount != 1 {
		t.Errorf("BatchCount: %d", file.ADVControl.BatchCount)
	}
	if file.ADVControl.EntryAddendaCount != 1 {
		t.Errorf("File Control EntryAddendaCount: %d", file.ADVControl.EntryAddendaCount)
	}
	if file.ADVControl.TotalDebitEntryDollarAmountInFile != 0 || file.ADVControl.TotalCreditEntryDollarAmountInFile != 100000 {
		t.Errorf("debit=%d credit=%d", file.ADVControl.TotalDebitEntryDollarAmountInFile, file.ADVControl.TotalCreditEntryDollarAmountInFile)
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

func TestFile__readInvalidJson(t *testing.T) {
	path := filepath.Join("test", "testdata", "ppd-invalid.json")
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	_, err = FileFromJSON(bs)

	if err != nil {
		if strings.Contains(err.Error(), "problem reading File") {
		} else {
			t.Fatal(err)
		}
	}
}

func TestFile__readEmptyJson(t *testing.T) {
	bs := make([]byte, 0)
	_, err := FileFromJSON(bs)

	if err != nil {
		if strings.Contains(err.Error(), "no JSON data provided") {
		} else {
			t.Fatal(err)
		}
	}
}

func TestFile__readNoBatchesJson(t *testing.T) {
	path := filepath.Join("test", "testdata", "ppd-noBatches.json")
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	_, err = FileFromJSON(bs)

	if !base.Match(err, ErrConstructor) {
		t.Fatal(err)
	}
}

func TestFile__readInvalidFilesJson(t *testing.T) {
	path := filepath.Join("test", "testdata", "ppd-invalidFile.json")
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	_, err = FileFromJSON(bs)

	if !base.Match(err, ErrUpperAlpha) {
		t.Errorf("%T: %s", err, err)
	}
}
