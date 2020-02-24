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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

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
	if file.Header.FileCreationTime != "" || file.Header.FileCreationDate != "181008" {
		t.Errorf("time=%v date=%v", file.Header.FileCreationTime, file.Header.FileCreationDate)
	}
	if v := file.Header.FileCreationTimeField(); v == "" || len(v) != 4 {
		t.Errorf("time=%v", v)
	}
	if v := file.Header.FileCreationDateField(); v != "181008" {
		t.Errorf("date=%v", v)
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

func TestFile__rfc3339JSON(t *testing.T) {
	bs, err := ioutil.ReadFile(filepath.Join("test", "testdata", "rfc3339.json"))
	if err != nil {
		t.Fatal(err)
	}
	file, err := FileFromJSON(bs)
	if err != nil {
		t.Fatal(err)
	}

	if file.ID != "rfc3339" {
		t.Fatalf("file.ID=%s", file.ID)
	}

	if file.Header.FileCreationDate != "091110" {
		t.Errorf("file.Header.FileCreationDate=%s", file.Header.FileCreationDate)
	}
	if file.Header.FileCreationTime != "2300" {
		t.Errorf("file.Header.FileCreationTime=%s", file.Header.FileCreationTime)
	}

	if len(file.Batches) != 1 {
		t.Fatalf("got %d Batches", len(file.Batches))
	}

	header := file.Batches[0].GetHeader()
	if header.CompanyDescriptiveDate != "SD2300" {
		t.Errorf("header.CompanyDescriptiveDate=%s", header.CompanyDescriptiveDate)
	}
	if header.EffectiveEntryDate != "091110" {
		t.Errorf("header.EffectiveEntryDate=%s", header.EffectiveEntryDate)
	}
}

func TestFile__iso8601JSON(t *testing.T) {
	bs, err := ioutil.ReadFile(filepath.Join("test", "testdata", "iso8601.json"))
	if err != nil {
		t.Fatal(err)
	}
	file, err := FileFromJSON(bs)
	if err != nil {
		t.Fatal(err)
	}

	if file.ID != "iso8601" {
		t.Fatalf("file.ID=%s", file.ID)
	}

	if file.Header.FileCreationDate != "190920" {
		t.Errorf("file.Header.FileCreationDate=%s", file.Header.FileCreationDate)
	}
	if file.Header.FileCreationTime != "2114" {
		t.Errorf("file.Header.FileCreationTime=%s", file.Header.FileCreationTime)
	}

	if len(file.Batches) != 1 {
		t.Fatalf("got %d Batches", len(file.Batches))
	}

	header := file.Batches[0].GetHeader()
	if header.CompanyDescriptiveDate != "SD2114" {
		t.Errorf("header.CompanyDescriptiveDate=%s", header.CompanyDescriptiveDate)
	}
	if header.EffectiveEntryDate != "190920" {
		t.Errorf("header.EffectiveEntryDate=%s", header.EffectiveEntryDate)
	}
}

func TestFile__IATdatetimeParse(t *testing.T) {
	bs, err := ioutil.ReadFile(filepath.Join("test", "testdata", "iat-debit.json"))
	if err != nil {
		t.Fatal(err)
	}
	file, err := FileFromJSON(bs)
	if err != nil {
		t.Fatal(err)
	}

	if file.ID != "iat-datetime" {
		t.Fatalf("file.ID=%s", file.ID)
	}
	if len(file.IATBatches) != 1 {
		t.Errorf("got %d IAT batches", len(file.IATBatches))
	}
	if date := file.IATBatches[0].Header.EffectiveEntryDate; date != "190923" {
		t.Errorf("file.IATBatches[0].Header.EffectiveEntryDate=%s", date)
	}
}

func TestFile__datetimeParse(t *testing.T) {
	// from javascript: (new Date).toISOString()
	if ts, err := datetimeParse("2019-09-20T20:49:35.177Z"); err != nil {
		t.Error(err)
	} else {
		if v := ts.Format("060102"); v != "190920" {
			t.Errorf("got %s", v)
		}
	}

	// RFC3339 format
	if ts, err := datetimeParse("2019-09-23T09:50:52-07:00"); err != nil {
		t.Error(err)
	} else {
		if v := ts.Format("060102"); v != "190923" {
			t.Errorf("got %s", v)
		}
	}

	// other, expect zero time
	if ts, err := datetimeParse(""); !ts.IsZero() || err == nil {
		t.Errorf("ts=%v error=%v", ts, err)
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
	if !base.Match(err, NewErrRecordType(9)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileControlValidate validates PPD File Control
func TestFileControlValidate(t *testing.T) {
	file := mockFilePPD()

	file.Control.recordType = "22"
	err := file.Validate()
	if !base.Match(err, NewErrRecordType(9)) {
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

// TestBatchControlNil Batch Control Nil
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
	if _, err := FileFromJSON(bs); err == nil {
		t.Error("expected error")
	} else {
		if !strings.Contains(err.Error(), "problem reading File") {
			t.Fatal(err)
		}
	}

	// a file which fails .Validate()
	path = filepath.Join("test", "testdata", "ppd-invalid-EntryDetail-checkDigit.json")
	bs, err = ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := FileFromJSON(bs); err == nil {
		t.Error("expected error")
	} else {
		if !strings.Contains(err.Error(), "batch #1 (PPD) FieldError RDFIIdentification 1 does not match calculated check digit 4") {
			t.Error(err)
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

func TestFile_largeFileEntryHash(t *testing.T) {
	// To create a file
	fh := NewFileHeader()
	fh.ImmediateDestination = "231380104"
	fh.ImmediateOrigin = "121042882"
	fh.FileCreationDate = time.Now().Format("060102")
	fh.ImmediateDestinationName = "Citadel"
	fh.ImmediateOriginName = "Wells Fargo"
	file := NewFile()
	file.SetHeader(fh)

	// Create 2 Batches of SEC Code PPD

	bh := NewBatchHeader()
	bh.ServiceClassCode = MixedDebitsAndCredits
	bh.CompanyName = "Wells Fargo"
	bh.CompanyIdentification = "121042882"
	bh.StandardEntryClassCode = PPD
	bh.CompanyEntryDescription = "Trans. Description"
	bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1).Format("060102")
	bh.ODFIIdentification = "121042882"

	batch, _ := NewBatch(bh)

	// Create Entry
	entrySeq := 0
	for i := 0; i < 1250; i++ {
		entrySeq = entrySeq + 1

		entryEntrySeq := NewEntryDetail()
		entryEntrySeq.TransactionCode = CheckingCredit
		entryEntrySeq.SetRDFI("231380104")
		entryEntrySeq.DFIAccountNumber = "81967038518"
		entryEntrySeq.Amount = 100000
		entryEntrySeq.IndividualName = "Steven Tander"
		entryEntrySeq.SetTraceNumber(bh.ODFIIdentification, entrySeq)
		entryEntrySeq.IdentificationNumber = "#83738AB#"
		entryEntrySeq.Category = CategoryForward
		entryEntrySeq.AddendaRecordIndicator = 1

		// Add addenda record for an entry
		addendaEntrySeq := NewAddenda05()
		addendaEntrySeq.PaymentRelatedInformation = "bonus pay for amazing work on #OSS"
		entryEntrySeq.AddAddenda05(addendaEntrySeq)

		// Add entries
		batch.AddEntry(entryEntrySeq)

	}

	// Create the batch.
	if err := batch.Create(); err != nil {
		fmt.Printf("%T: %s", err, err)
	}

	// Add batch to the file
	file.AddBatch(batch)

	// Create the file
	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	// ensure we have a validated file structure
	if err := file.Validate(); err != nil {
		t.Errorf("Could not validate entire file: %v", err)
	}

	// Per BatchControl.Parse() and Batch.calculateEntryHash()
	// EntryHash is essentially the sum of all the RDFI routing numbers in the batch. If the sum exceeds 10 digits
	// (because you have lots of Entry Detail Records), lop off the most significant digits of the sum until there
	// are only 10.  This allows the entryHash to be 10 digits as per the ACH specification.

	testHash := 0
	for _, batch := range file.Batches {

		for _, entry := range batch.GetEntries() {

			entryRDFI, _ := strconv.Atoi(entry.RDFIIdentification)

			testHash = testHash + entryRDFI
		}

		if testHash != 28922512500 {
			t.Errorf("Expected '28922512500' Calculated Entry Hash: %d", testHash)
		}

		s := strconv.Itoa(testHash)
		ln := uint(len(s))
		if ln > 10 {
			s = s[ln-10:]
		}

		testHash, _ = strconv.Atoi(s)

		if testHash != 8922512500 {
			t.Errorf("Expected '8922512500' Calculated Entry Hash: %d", testHash)
		}
	}
}

func TestFile__RemoveBatch(t *testing.T) {
	file, err := readACHFilepath(filepath.Join("test", "testdata", "ppd-debit.ach"))
	if err != nil {
		t.Fatal(err)
	}
	if len(file.Batches) != 1 {
		t.Errorf("unexpected number of batches: %d", len(file.Batches))
	}

	// remove the batch and check
	file.RemoveBatch(file.Batches[0])
	if len(file.Batches) != 0 {
		t.Errorf("unexpected number of batches: %d", len(file.Batches))
	}

	// NOC Entries
	nocHeader := NewBatchHeader()
	nocHeader.ServiceClassCode = CreditsOnly
	nocHeader.StandardEntryClassCode = COR
	nocHeader.CompanyName = "Your Company, inc"
	nocHeader.CompanyIdentification = "121042882"
	nocHeader.CompanyEntryDescription = "Vendor Pay"
	nocHeader.ODFIIdentification = "121042882"
	noc := NewBatchCOR(nocHeader)
	nocED := mockCOREntryDetail()
	nocED.Addenda98 = mockAddenda98()
	nocED.Category = CategoryNOC
	nocED.AddendaRecordIndicator = 1
	noc.AddEntry(nocED)
	if err := noc.Create(); err != nil {
		t.Fatal(err)
	}
	file.AddBatch(noc)
	if len(file.NotificationOfChange) != 1 {
		t.Errorf("unexpected number of NOC batches: %d", len(file.NotificationOfChange))
	}
	file.RemoveBatch(noc)
	if len(file.NotificationOfChange) != 0 {
		t.Errorf("unexpected number of NOC batches: %d", len(file.NotificationOfChange))
	}

	// Returns
	file, err = readACHFilepath(filepath.Join("test", "testdata", "ppd-debit.ach"))
	if err != nil {
		t.Fatal(err)
	}
	if len(file.ReturnEntries) != 0 {
		t.Errorf("unexpected number of return entries: %d", len(file.ReturnEntries))
	}
	ppdHeader := mockBatchPPDHeader()
	ppd := NewBatchPPD(ppdHeader)
	ppdED := mockPPDEntryDetail()
	ppdED.Addenda99 = mockAddenda99()
	ppdED.Category = CategoryReturn
	ppd.AddEntry(ppdED)

	file.AddBatch(ppd)
	if len(file.ReturnEntries) != 1 {
		t.Errorf("unexpected number of return entries: %d", len(file.ReturnEntries))
	}
	file.RemoveBatch(ppd)
	if len(file.ReturnEntries) != 0 {
		t.Errorf("unexpected number of return entries: %d", len(file.ReturnEntries))
	}
}

func TestFile__SegmentFile(t *testing.T) {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("test", "testdata", "ppd-mixedDebitCredit.ach"))

	if err != nil {
		t.Fatal(err)
	}
	r := NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		t.Fatalf("Issue reading file: %+v \n", err)
	}

	// ensure we have a validated file structure
	if achFile.Validate(); err != nil {
		t.Fatalf("Could not validate entire read file: %v", err)
	}

	sfc := NewSegmentFileConfiguration()
	creditFile, debitFile, err := achFile.SegmentFile(sfc)

	if err != nil {
		t.Fatalf("Could not segment the file: %+v \n", err)
	}

	if err := creditFile.Validate(); err != nil {
		t.Fatalf("Credit file did not validate: %+v \n", err)
	}

	if err := debitFile.Validate(); err != nil {
		t.Fatalf("Debit File did not validate: %+v \n", err)
	}
}

func TestFile__SegmentFileCredit(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "ach-ppd-read", "ppd-credit.ach"))
	if err != nil {
		t.Fatal(err)
	}
	r := NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		t.Fatalf("Issue reading file: %+v \n", err)
	}

	// ensure we have a validated file structure
	if achFile.Validate(); err != nil {
		t.Fatalf("Could not validate entire read file: %v", err)
	}

	sfc := NewSegmentFileConfiguration()
	_, _, err = achFile.SegmentFile(sfc)

	if err != nil {
		if !base.Match(err, ErrFileNoBatches) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFile__SegmentFileDebitOnly(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "ach-ppd-read", "ppd-debit.ach"))
	if err != nil {
		t.Fatal(err)
	}
	r := NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		t.Fatalf("Issue reading file: %+v \n", err)
	}

	// ensure we have a validated file structure
	if achFile.Validate(); err != nil {
		t.Fatalf("Could not validate entire read file: %v", err)
	}

	sfc := NewSegmentFileConfiguration()
	_, _, err = achFile.SegmentFile(sfc)

	if err != nil {
		if !base.Match(err, ErrFileNoBatches) {
			t.Errorf("%T: %s", err, err)
		}
	}

}

func TestFile__SegmentADVFile(t *testing.T) {
	bs, err := ioutil.ReadFile(filepath.Join("test", "testdata", "adv-valid.json"))
	if err != nil {
		t.Fatal(err)
	}
	f, err := FileFromJSON(bs)
	if err != nil {
		t.Fatal(err)
	}

	creditFile, debitFile, err := f.SegmentFile(nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(debitFile.Batches) != 0 {
		t.Fatalf("debitFile.Batches=%#v", debitFile.Batches)
	}
	if !creditFile.IsADV() {
		t.Errorf("creditFile.IsADV=%v", creditFile.IsADV())
	}
	if len(creditFile.Batches) != 1 {
		t.Fatalf("credit: batches=%d", len(creditFile.Batches))
	}

	creditADVEntries := creditFile.Batches[0].GetADVEntries()
	if len(creditADVEntries) != 1 {
		t.Errorf("creditADVEntries=%d", len(creditADVEntries))
	}

	entry := creditADVEntries[0]
	if entry.ID != "adv-01" {
		t.Errorf("ADV entry: %#v", entry)
	}
	if entry.TransactionCode != CreditForDebitsOriginated {
		t.Errorf("ADV entry: %#v", entry)
	}
	if entry.AdviceRoutingNumber != "121042882" {
		t.Errorf("ADV entry: %#v", entry)
	}
	if entry.FileIdentification != "11131" {
		t.Errorf("ADV entry: %#v", entry)
	}
	if entry.ACHOperatorRoutingNumber != "01100001" {
		t.Errorf("ADV entry: %#v", entry)
	}
}

func TestFile__SegmentADVFileDebit(t *testing.T) {
	bs, err := ioutil.ReadFile(filepath.Join("test", "testdata", "adv-valid.json"))
	if err != nil {
		t.Fatal(err)
	}
	f, err := FileFromJSON(bs)
	if err != nil {
		t.Fatal(err)
	}

	// Force Batch Header and Control to ADV types
	bh := f.Batches[0].GetHeader()
	bh.ServiceClassCode = AutomatedAccountingAdvices
	f.Batches[0].SetHeader(bh)

	bc := f.Batches[0].GetControl()
	bc.ServiceClassCode = AutomatedAccountingAdvices
	f.Batches[0].SetControl(bc)

	// Force the ADVEntryDetail to Debit
	f.Batches[0].GetADVEntries()[0].TransactionCode = DebitForDebitsReceived

	creditFile, debitFile, err := f.SegmentFile(nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(creditFile.Batches) != 0 {
		t.Fatalf("creditFile.Batches=%#v", creditFile.Batches)
	}
	if !debitFile.IsADV() {
		t.Fatalf("debitFile.IsADV=%v", debitFile.IsADV())
	}

	debitADVEntries := debitFile.Batches[0].GetADVEntries()

	if len(debitADVEntries) != 1 {
		t.Errorf("debitADVEntries=%d", len(debitADVEntries))
	}

	entry := debitADVEntries[0]
	if entry.ID != "adv-01" {
		t.Errorf("ADV entry: %#v", entry)
	}
	if entry.TransactionCode != DebitForDebitsReceived {
		t.Errorf("ADV entry: %#v", entry)
	}
	if entry.AdviceRoutingNumber != "121042882" {
		t.Errorf("ADV entry: %#v", entry)
	}
	if entry.FileIdentification != "11131" {
		t.Errorf("ADV entry: %#v", entry)
	}
	if entry.ACHOperatorRoutingNumber != "01100001" {
		t.Errorf("ADV entry: %#v", entry)
	}
}

func TestSegmentFile_FileHeaderError(t *testing.T) {
	achFile := NewFile()

	sfc := NewSegmentFileConfiguration()
	_, _, err := achFile.SegmentFile(sfc)

	if err != nil {
		if !base.Match(err, ErrConstructor) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestFileSegmentFileBatchControlCreditAmount(t *testing.T) {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("test", "testdata", "ppd-mixedDebitCredit.ach"))
	if err != nil {
		t.Fatal(err)
	}
	r := NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		t.Fatalf("Issue reading file: %+v \n", err)
	}

	// ensure we have a validated file structure
	if achFile.Validate(); err != nil {
		t.Fatalf("Could not validate entire read file: %v", err)
	}

	sfc := NewSegmentFileConfiguration()
	creditFile, debitFile, err := achFile.SegmentFile(sfc)

	if err != nil {
		t.Fatalf("Could not segment the file: %+v \n", err)
	}

	if err := creditFile.Validate(); err != nil {
		t.Fatalf("Credit file did not validate: %+v \n", err)
	}

	if err := debitFile.Validate(); err != nil {
		t.Fatalf("Debit File did not validate: %+v \n", err)
	}

	if creditFile.Batches[0].GetControl().TotalCreditEntryDollarAmount != 200000000 {
		t.Errorf("expected %s received %v", "200000000", creditFile.Batches[0].GetControl().TotalCreditEntryDollarAmount)
	}
}

func TestFileSegmentFileBatchControlDebitAmount(t *testing.T) {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("test", "testdata", "ppd-mixedDebitCredit.ach"))
	if err != nil {
		t.Fatal(err)
	}
	r := NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		t.Fatalf("Issue reading file: %+v \n", err)
	}

	// ensure we have a validated file structure
	if achFile.Validate(); err != nil {
		t.Fatalf("Could not validate entire read file: %v", err)
	}

	sfc := NewSegmentFileConfiguration()
	creditFile, debitFile, err := achFile.SegmentFile(sfc)

	if err != nil {
		t.Fatalf("Could not segment the file: %+v \n", err)
	}

	if err := creditFile.Validate(); err != nil {
		t.Fatalf("Credit file did not validate: %+v \n", err)
	}

	if err := debitFile.Validate(); err != nil {
		t.Fatalf("Debit File did not validate: %+v \n", err)
	}

	if debitFile.Batches[0].GetControl().TotalDebitEntryDollarAmount != 200000000 {
		t.Errorf("expected %s received %v", "200000000", debitFile.Batches[0].GetControl().TotalDebitEntryDollarAmount)
	}
}

func TestFileSegmentFileCreditBatches(t *testing.T) {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("test", "testdata", "ppd-mixedDebitCredit.ach"))
	if err != nil {
		t.Fatal(err)
	}
	r := NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		t.Fatalf("Issue reading file: %+v \n", err)
	}

	// ensure we have a validated file structure
	if achFile.Validate(); err != nil {
		t.Fatalf("Could not validate entire read file: %v", err)
	}

	sfc := NewSegmentFileConfiguration()
	creditFile, debitFile, err := achFile.SegmentFile(sfc)

	if err != nil {
		t.Fatalf("Could not segment the file: %+v \n", err)
	}

	if err := creditFile.Validate(); err != nil {
		t.Fatalf("Credit file did not validate: %+v \n", err)
	}

	if err := debitFile.Validate(); err != nil {
		t.Fatalf("Debit File did not validate: %+v \n", err)
	}

	if len(creditFile.Batches) != 1 {
		t.Errorf("expected %s received %v", "1", len(creditFile.Batches))
	}
}

func TestFileSegmentFileDebitBatches(t *testing.T) {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("test", "testdata", "ppd-mixedDebitCredit.ach"))
	if err != nil {
		t.Fatal(err)
	}
	r := NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		t.Fatalf("Issue reading file: %+v \n", err)
	}

	// ensure we have a validated file structure
	if achFile.Validate(); err != nil {
		t.Fatalf("Could not validate entire read file: %v", err)
	}

	sfc := NewSegmentFileConfiguration()
	creditFile, debitFile, err := achFile.SegmentFile(sfc)

	if err != nil {
		t.Fatalf("Could not segment the file: %+v \n", err)
	}

	if err := creditFile.Validate(); err != nil {
		t.Fatalf("Credit file did not validate: %+v \n", err)
	}

	if err := debitFile.Validate(); err != nil {
		t.Fatalf("Debit File did not validate: %+v \n", err)
	}

	if len(debitFile.Batches) != 1 {
		t.Errorf("expected %s received %v", "1", len(debitFile.Batches))
	}
}

func TestSegmentFileCreditOnly(t *testing.T) {
	// write an ACH file into repository
	fd, err := os.Open(filepath.Join("test", "testdata", "ppd-valid.json"))
	if fd == nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()
	bs, _ := ioutil.ReadAll(fd)
	file, _ := FileFromJSON(bs)

	sfc := NewSegmentFileConfiguration()
	creditFile, debitFile, err := file.SegmentFile(sfc)

	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if len(creditFile.Batches) != 1 {
		t.Errorf("expected %s received %v", "1", len(creditFile.Batches))
	}

	if debitFile.ID != "" {
		t.Error("No Debit File")
	}
}

func TestSegmentFileDebitOnly(t *testing.T) {
	// write an ACH file into repository
	fd, err := os.Open(filepath.Join("test", "testdata", "ppd-valid-debit.json"))
	if fd == nil {
		t.Fatalf("empty ACH file: %v", err)
	}
	defer fd.Close()
	bs, _ := ioutil.ReadAll(fd)
	file, _ := FileFromJSON(bs)

	sfc := NewSegmentFileConfiguration()
	creditFile, debitFile, err := file.SegmentFile(sfc)

	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	if creditFile.ID != "" {
		t.Error("No Debit File")
	}

	if len(debitFile.Batches) != 1 {
		t.Errorf("expected %s received %v", "1", len(debitFile.Batches))
	}
}

func TestFileIATSegmentFileCreditOnly(t *testing.T) {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("test", "ach-iat-read", "iat-credit.ach"))

	if err != nil {
		t.Fatal(err)
	}
	r := NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		t.Fatalf("Issue reading file: %+v \n", err)
	}

	// ensure we have a validated file structure
	if achFile.Validate(); err != nil {
		t.Fatalf("Could not validate entire read file: %v", err)
	}

	sfc := NewSegmentFileConfiguration()
	creditFile, debitFile, err := achFile.SegmentFile(sfc)

	if err != nil {
		t.Fatalf("Could not segment the file: %+v \n", err)
	}

	if err := creditFile.Validate(); err != nil {
		t.Fatalf("Credit file did not validate: %+v \n", err)
	}

	if len(debitFile.IATBatches) > 0 {
		t.Fatalf("IATFile should not have IAT batches: %+v \n", err)
	}
}

func TestFileIATSegmentFileDebitOnly(t *testing.T) {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("test", "testdata", "iat-debit.ach"))

	if err != nil {
		t.Fatal(err)
	}
	r := NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		t.Fatalf("Issue reading file: %+v \n", err)
	}

	// ensure we have a validated file structure
	if achFile.Validate(); err != nil {
		t.Fatalf("Could not validate entire read file: %v", err)
	}

	sfc := NewSegmentFileConfiguration()
	creditFile, debitFile, err := achFile.SegmentFile(sfc)

	if err != nil {
		t.Fatalf("Could not segment the file: %+v \n", err)
	}

	if len(creditFile.IATBatches) > 0 {
		t.Fatalf("IATFile should not have IAT credit batches: %+v \n", err)
	}

	if err := debitFile.Validate(); err != nil {
		t.Fatalf("Debit file did not validate: %+v \n", err)
	}
}

// TestFileIAT__SegmentFile test segmenting and IAT File
func TestFileIAT__SegmentFile(t *testing.T) {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("test", "testdata", "iat-mixedDebitCredit.ach"))

	if err != nil {
		t.Fatal(err)
	}
	r := NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		t.Fatalf("Issue reading file: %+v \n", err)
	}

	// ensure we have a validated file structure
	if achFile.Validate(); err != nil {
		t.Fatalf("Could not validate entire read file: %v", err)
	}

	sfc := NewSegmentFileConfiguration()
	creditFile, debitFile, err := achFile.SegmentFile(sfc)

	if err != nil {
		t.Fatalf("Could not segment the file: %+v \n", err)
	}

	if err := creditFile.Validate(); err != nil {
		t.Fatalf("Credit file did not validate: %+v \n", err)
	}

	if err := debitFile.Validate(); err != nil {
		t.Fatalf("Debit File did not validate: %+v \n", err)
	}
}

// TestFile_FlattenFileOneBatchHeader
func TestFile_FlattenFileOneBatchHeader(t *testing.T) {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("test", "testdata", "flattenBatchesOneBatchHeader.ach"))

	if err != nil {
		t.Fatal(err)
	}
	r := NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		t.Fatalf("Issue reading file: %+v \n", err)
	}

	ff, err := achFile.FlattenBatches()
	if err != nil {
		t.Fatalf("Could not flatten the file: %+v \n", err)
	}
	if err := ff.Validate(); err != nil {
		t.Fatalf("Flattend file did not validate: %+v \n", err)
	}
}

// TestFileFlattenFileMultipleBatchHeaders
func TestFileFlattenFileMultipleBatchHeaders(t *testing.T) {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("test", "testdata", "flattenBatchesMultipleBatchHeaders.ach"))

	if err != nil {
		t.Fatal(err)
	}
	r := NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		t.Fatalf("Issue reading file: %+v \n", err)
	}

	ff, err := achFile.FlattenBatches()
	if err != nil {
		t.Fatalf("Could not flatten the file: %+v \n", err)
	}

	if err := ff.Validate(); err != nil {
		t.Fatalf("Flattend file did not validate: %+v \n", err)
	}
}

func TestFlattenFile_FileHeaderError(t *testing.T) {
	achFile := NewFile()

	_, err := achFile.FlattenBatches()

	if err != nil {
		if !base.Match(err, ErrConstructor) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

// TestFile_FlattenFileOneIATBatchHeader
func TestFile_FlattenFileOneIATBatchHeader(t *testing.T) {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("test", "testdata", "flattenIATBatchesOneBatchHeader.ach"))

	if err != nil {
		t.Fatal(err)
	}
	r := NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		t.Fatalf("Issue reading file: %+v \n", err)
	}

	ff, err := achFile.FlattenBatches()

	if err != nil {
		t.Fatalf("Could not flatten the file: %+v \n", err)
	}

	if err := ff.Validate(); err != nil {
		t.Fatalf("Flattend file did not validate: %+v \n", err)
	}
}

// TestFileFlattenFileMultipleIATBatchHeaders
func TestFileFlattenFileMultipleIATBatchHeaders(t *testing.T) {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("test", "testdata", "flattenIATBatchesMultipleBatchHeaders.ach"))

	if err != nil {
		t.Fatal(err)
	}
	r := NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		t.Fatalf("Issue reading file: %+v \n", err)
	}

	ff, err := achFile.FlattenBatches()

	if err != nil {
		t.Fatalf("Could not flatten the file: %+v \n", err)
	}

	if err := ff.Validate(); err != nil {
		t.Fatalf("Flattend file did not validate: %+v \n", err)
	}
}

// TestFile_FlattenFileOneADVBatchHeader
func TestFile_FlattenFileOneADVBatchHeader(t *testing.T) {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("test", "testdata", "flattenADVBatchesOneBatchHeader.ach"))

	if err != nil {
		t.Fatal(err)
	}
	r := NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		t.Fatalf("Issue reading file: %+v \n", err)
	}

	ff, err := achFile.FlattenBatches()

	if err != nil {
		t.Fatalf("Could not flatten the file: %+v \n", err)
	}

	if err := ff.Validate(); err != nil {
		t.Fatalf("Flattend file did not validate: %+v \n", err)
	}
}

// TestFileFlattenFileMultipleADVBatchHeaders
func TestFileFlattenFileMultipleADVTBatchHeaders(t *testing.T) {
	// open a file for reading. Any io.Reader Can be used
	f, err := os.Open(filepath.Join("test", "testdata", "flattenADVBatchesMultipleBatchHeaders.ach"))

	if err != nil {
		t.Fatal(err)
	}
	r := NewReader(f)
	achFile, err := r.Read()
	if err != nil {
		t.Fatalf("Issue reading file: %+v \n", err)
	}

	ff, err := achFile.FlattenBatches()

	if err != nil {
		t.Fatalf("Could not flatten the file: %+v \n", err)
	}

	if err := ff.Validate(); err != nil {
		t.Fatalf("Flattend file did not validate: %+v \n", err)
	}
}
