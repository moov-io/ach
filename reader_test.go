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
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/moov-io/base"
	"github.com/stretchr/testify/require"
)

func readACHFilepath(path string) (*File, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	f, err := NewReader(fd).Read()
	return &f, err
}

func TestReadFile(t *testing.T) {
	file, err := ReadFile(filepath.Join("test", "testdata", "web-debit.ach"))
	if err != nil {
		t.Fatal(err)
	}
	if origin := file.Header.ImmediateOrigin; origin != "231380104" {
		t.Errorf("origin=%s", origin)
	}
}

func TestReadFiles(t *testing.T) {
	paths := []string{
		filepath.Join("test", "testdata", "return-WEB.ach"),
		filepath.Join("test", "testdata", "web-debit.ach"),
	}
	files, err := ReadFiles(paths)
	if err != nil {
		t.Fatal(err)
	}
	if n := len(files); n != 2 {
		t.Fatalf("read %d files", n)
	}
}

func TestReadPartial(t *testing.T) {
	file, err := ReadFile(filepath.Join("test", "testdata", "bh-ed-ad-bh-ed-ad-ed-ad.ach"))
	if err != nil {
		t.Log(err) // we expect errors -- display them
	}
	if len(file.Batches) != 1 {
		t.Fatalf("unexpected batches: %#v", file.Batches)
	}
	entries := file.Batches[0].GetEntries()
	if len(entries) != 3 {
		t.Error("unexpected entry details")
		for i := range entries {
			t.Errorf("  %#v", entries[i])
		}
		t.Fatal("")
	}

	// batch
	bh := file.Batches[0].GetHeader()
	if bh.ServiceClassCode != DebitsOnly {
		t.Errorf("ServiceClassCode=%d", bh.ServiceClassCode)
	}
	if bh.CompanyName != "Adam Shannon" {
		t.Errorf("CompanyName=%s", bh.CompanyName)
	}
	if bh.CompanyIdentification != "MOOVYYYYYY" {
		t.Errorf("CompanyIdentification=%s", bh.CompanyIdentification)
	}
	if bh.StandardEntryClassCode != "PPD" {
		t.Errorf("StandardEntryClassCode=%s", bh.StandardEntryClassCode)
	}

	// entries
	entry := entries[0]
	if entry.TransactionCode != CheckingReturnNOCDebit {
		t.Errorf("TransactionCode=%d", entry.TransactionCode)
	}
	if num := strings.TrimSpace(entry.DFIAccountNumber); num != "15XXXXXXXXXX1" {
		t.Errorf("DFIAccountNumber=%q", num)
	}
	if entry.Addenda99 != nil {
		if code := entry.Addenda99.ReturnCode; code != "R02" {
			t.Errorf("ReturnCode=%s", code)
		}
	} else {
		t.Errorf("nil Addenda99")
	}

	entry = entries[1]
	if entry.TransactionCode != CheckingReturnNOCCredit {
		t.Errorf("TransactionCode=%d", entry.TransactionCode)
	}
	if num := strings.TrimSpace(entry.DFIAccountNumber); num != "1XXXXXXXXXXX2" {
		t.Errorf("DFIAccountNumber=%q", num)
	}
	if entry.Addenda99 != nil {
		if code := entry.Addenda99.ReturnCode; code != "R03" {
			t.Errorf("ReturnCode=%s", code)
		}
	} else {
		t.Errorf("nil Addenda99")
	}
}

func TestReader__crashers(t *testing.T) {
	dir := filepath.Join("test", "testdata", "crashers")
	fds, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	var currentFile string

	// log current file when we panic
	defer func() {
		if v := recover(); v != nil {
			if _, ok := v.(error); ok {
				t.Errorf("panic from parsing %s", filepath.Join(dir, currentFile))
				panic(v) // throw original panic so testing package emits trace
			}
		}
	}()

	for i := range fds {
		currentFile = fds[i].Name()
		t.Logf("parsing %s to see if it crashes...", currentFile)

		f, err := os.Open(filepath.Join(dir, fds[i].Name()))
		if err != nil {
			t.Fatal(err)
		}
		NewReader(f).Read() // making sure we don't panic
	}
}

func TestParser__errors(t *testing.T) {
	var errorList base.ErrorList
	var buf bytes.Buffer

	// nil
	errorList.Print(&buf)
	if v := buf.String(); v != "<nil>" {
		t.Errorf("got %q", v)
	}
	if v := errorList.Error(); v != "<nil>" {
		t.Errorf("got %q", v)
	}
	buf.Reset()

	// empty
	errorList = make(base.ErrorList, 0)
	errorList.Print(&buf)
	if v := buf.String(); v != "<nil>" {
		t.Errorf("got %q", v)
	}
	if v := errorList.Error(); v != "<nil>" {
		t.Errorf("got %q", v)
	}
	buf.Reset()

	// one error
	errorList.Add(errors.New("testing"))
	errorList.Print(&buf)
	if v := buf.String(); v != "testing" {
		t.Errorf("got %q", v)
	}
	if v := errorList.Error(); v != "testing" {
		t.Errorf("got %q", v)
	}
	buf.Reset()

	// multiple errors
	errorList.Add(errors.New("other error"))
	errorList.Add(errors.New("last"))
	errorList.Print(&buf)
	if v := buf.String(); v != "testing\n  other error\n  last" {
		t.Errorf("got %q", v)
	}
	if v := errorList.Error(); v != "testing\n  other error\n  last" {
		t.Errorf("got %q", v)
	}
}

func TestParser__errorJSON(t *testing.T) {
	var errorList base.ErrorList
	errorList.Add(errors.New("first"))
	errorList.Add(errors.New("second"))
	errorList.Add(errors.New("third"))

	// marshal json
	bs, err := json.Marshal(errorList)
	if err != nil {
		t.Fatal(err)
	}
	v := string(bs)
	if v != "\"first\\n  second\\n  third\"" { // JSON strings are quoted
		t.Errorf("got %q", v)
	}
}

// testPPDDebitRead validates reading a PPD debit
func testPPDDebitRead(t testing.TB) {
	f, err := os.Open(filepath.Join("test", "testdata", "ppd-debit.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	if _, err := r.Read(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err = r.File.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestPPDDebitRead tests validating reading a PPD debit
func TestPPDDebitRead(t *testing.T) {
	testPPDDebitRead(t)
}

// BenchmarkPPDDebitRead benchmarks validating reading a PPD debit
func BenchmarkPPDDebitRead(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testPPDDebitRead(b)
	}
}

// testWEBDebitRead validates reading a WEB debit
func testWEBDebitRead(t testing.TB) {
	f, err := os.Open(filepath.Join("test", "testdata", "web-debit.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err = r.File.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestWEBDebitRead tests validating reading a WEB debit
func TestWEBDebitRead(t *testing.T) {
	testWEBDebitRead(t)
}

// BenchmarkWEBDebitRead benchmarks validating reading a WEB debit
func BenchmarkWEBDebitRead(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testWEBDebitRead(b)
	}
}

// testPPDDebitFixedLengthRead validates reading a PPD debit fixed width length
func testPPDDebitFixedLengthRead(t testing.TB) {
	f, err := os.Open(filepath.Join("test", "testdata", "ppd-debit-fixedLength.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
}

// TestPPDDebitFixedLengthRead test validates reading a PPD debit fixed width length
func TestPPDDebitFixedLengthRead(t *testing.T) {
	testPPDDebitFixedLengthRead(t)
}

// BenchmarkPPDDebitFixedLengthRead benchmark validates reading a PPD debit fixed width length
func BenchmarkPPDDebitFixedLengthRead(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testPPDDebitFixedLengthRead(b)
	}
}

func TestPPDDebitFixedLengthRead__InvalidLength(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "ppd-debit-fixedLengthInvalid.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()

	r := NewReader(f)
	_, err = r.Read()
	if err == nil {
		t.Errorf("expected error")
	}

	wantErr := "1 extra character(s) in ACH file: must be 470 but found 471"
	if err != nil && !strings.Contains(err.Error(), wantErr) {
		t.Errorf("want: %v, got: %v", wantErr, err)
	}
}

// testRecordTypeUnknown validates record type unknown
func testRecordTypeUnknown(t testing.TB) {
	var line = "301 076401251 0764012510807291511A094101achdestname            companyname                    "
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if !base.Has(err, NewErrUnknownRecordType("3")) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestRecordTypeUnknown tests validating record type unknown
func TestRecordTypeUnknown(t *testing.T) {
	testRecordTypeUnknown(t)
}

// BenchmarkRecordTypeUnknown benchmarks validating record type unknown
func BenchmarkRecordTypeUnknown(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testRecordTypeUnknown(b)
	}
}

// testTwoFileHeaders validates one file header
func testTwoFileHeaders(t testing.TB) {
	var line = "101 076401251 0764012510807291511A094101achdestname            companyname                    "
	var twoHeaders = line + "\n" + line
	r := NewReader(strings.NewReader(twoHeaders))
	_, err := r.Read()

	if !base.Has(err, ErrFileHeader) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestTwoFileHeaders tests validating one file header
func TestTwoFileHeaders(t *testing.T) {
	testTwoFileHeaders(t)
}

// BenchmarkTwoFileHeaders benchmarks
func BenchmarkTwoFileHeaders(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testTwoFileHeaders(b)
	}
}

// testTwoFileControls validates one file control
func testTwoFileControls(t testing.TB) {
	var line = "9000001000001000000010005320001000000010500000000000000                                       "
	var twoControls = line + "\n" + line
	r := NewReader(strings.NewReader(twoControls))
	r.addCurrentBatch(NewBatchPPD(mockBatchPPDHeader()))
	bc := BatchControl{EntryAddendaCount: 1,
		TotalDebitEntryDollarAmount: 10500,
		EntryHash:                   5320001}
	r.currentBatch.SetControl(&bc)

	r.File.AddBatch(r.currentBatch)
	r.File.Control.EntryHash = 5320001

	_, err := r.Read()
	if !base.Has(err, ErrFileControl) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestTwoFileControls tests validating one file control
func TestTwoFileControls(t *testing.T) {
	testTwoFileControls(t)
}

// BenchmarkTwoFileControls benchmarks validating one file control
func BenchmarkTwoFileControls(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testTwoFileControls(b)
	}
}

// testFileLineEmpty verifies empty files fail to parse
func testFileLineEmpty(t testing.TB) {
	line := ""
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if !base.Has(err, ErrFileHeader) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileLineEmpty tests validating empty file fails to parse
func TestFileLineEmpty(t *testing.T) {
	testFileLineEmpty(t)
}

// BenchmarkFileLineEmpty benchmarks validating empty file fails to parse
func BenchmarkFileLineEmpty(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileLineEmpty(b)
	}
}

// testFileLineShort validates file line is short
func testFileLineShort(t testing.TB) {
	line := "1 line is only 70 characters ........................................!"
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()

	if !base.Has(err, NewRecordWrongLengthErr(70)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileLineShort tests validating file line is short
func TestFileLineShort(t *testing.T) {
	testFileLineShort(t)
}

// BenchmarkFileLineShort benchmarks validating file line is short
func BenchmarkFileLineShort(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileLineShort(b)
	}
}

// testFileLineLong validates file line is long
func testFileLineLong(t testing.TB) {
	var line = "1 line is 100 characters ..........................................................................!\n2 line is 94 characters ....................................................................!\n"
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()

	if !base.Has(err, NewRecordWrongLengthErr(100)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileLineLong tests validating file line is long
func TestFileLineLong(t *testing.T) {
	testFileLineLong(t)
}

// BenchmarkFileLineLong benchmarks validating file line is long
func BenchmarkFileLineLong(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileLineLong(b)
	}
}

// testFileFileHeaderErr validates error flows back from the parser
func testFileFileHeaderErr(t testing.TB) {
	fh := mockFileHeader()
	// fh.ImmediateOrigin = "0000000000"
	fh.ImmediateOrigin = ""
	r := NewReader(strings.NewReader(fh.String()))
	// necessary to have a file control not nil
	r.File.Control = mockFileControl()
	_, err := r.Read()
	// TODO: is this testing what we want to be testing?
	if !base.Has(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileFileHeaderErr tests validating error flows back from the parser
func TestFileFileHeaderErr(t *testing.T) {
	testFileFileHeaderErr(t)
}

// BenchmarkFileFileHeaderErr benchmarks validating error flows back from the parse
func BenchmarkFileFileHeaderErr(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileFileHeaderErr(b)
	}
}

// testFileBatchHeaderErr validates error flows back from the parser
func testFileBatchHeaderErr(t testing.TB) {
	bh := mockBatchHeader()
	//bh.ODFIIdentification = 0
	bh.ODFIIdentification = ""
	r := NewReader(strings.NewReader(bh.String()))
	_, err := r.Read()
	// TODO: is this testing what we want to be testing?
	if !base.Has(err, ErrFileHeader) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileBatchHeaderErr tests validating error flows back from the parser
func TestFileBatchHeaderErr(t *testing.T) {
	testFileBatchHeaderErr(t)
}

// BenchmarkFileBatchHeaderErr benchmarks validating error flows back from the parser
func BenchmarkFileBatchHeaderErr(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileBatchHeaderErr(b)
	}
}

// testFileBatchHeaderDuplicate validates when two batch headers exists in a current batch
func testFileBatchHeaderDuplicate(t testing.TB) {
	// create a new Batch header string
	bh := mockBatchPPDHeader()
	r := NewReader(strings.NewReader(bh.String()))
	// instantiate a batch header in the reader
	r.addCurrentBatch(NewBatchPPD(bh))
	// read should fail because it is parsing a second batch header and there can only be one.
	_, err := r.Read()
	if !base.Has(err, ErrFileBatchHeaderInsideBatch) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileBatchHeaderDuplicate tests validating when two batch headers exists in a current batch
func TestFileBatchHeaderDuplicate(t *testing.T) {
	testFileBatchHeaderDuplicate(t)
}

// BenchmarkFileBatchHeaderDuplicate benchmarks validating when two batch headers exists in a current batch
func BenchmarkFileBatchHeaderDuplicate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileBatchHeaderDuplicate(b)
	}
}

// testFileEntryDetailOutsideBatch validates entry detail does not exist outside of batch
func testFileEntryDetailOutsideBatch(t testing.TB) {
	ed := mockEntryDetail()
	r := NewReader(strings.NewReader(ed.String()))
	_, err := r.Read()
	if !base.Has(err, ErrFileEntryOutsideBatch) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileEntryDetailOutsideBatch tests validating entry detail does not exist outside of batch
func TestFileEntryDetailOutsideBatch(t *testing.T) {
	testFileEntryDetailOutsideBatch(t)
}

// BenchmarkFileEntryDetailOutsideBatch benchmarks validating entry detail does not exist outside of batch
func BenchmarkFileEntryDetailOutsideBatch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileEntryDetailOutsideBatch(b)
	}
}

// testFileEntryDetail validates error populates through the reader
func testFileEntryDetail(t testing.TB) {
	ed := mockEntryDetail()
	ed.TransactionCode = 0
	line := ed.String()
	r := NewReader(strings.NewReader(line))
	r.addCurrentBatch(NewBatchPPD(mockBatchPPDHeader()))
	r.currentBatch.SetHeader(mockBatchHeader())
	_, err := r.Read()
	if !base.Has(err, NewRecordWrongLengthErr(93)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileEntryDetail tests validating error populates through the reader
func TestFileEntryDetail(t *testing.T) {
	testFileEntryDetail(t)
}

// BenchmarkFileEntryDetail benchmarks validating error populates through the reader
func BenchmarkFileEntryDetail(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileEntryDetail(b)
	}
}

// testFileAddenda05 validates error for an invalid addenda05
func testFileAddenda05(t testing.TB) {
	bh := mockBatchHeader()
	ed := mockEntryDetail()
	addenda05 := mockAddenda05()
	addenda05.SequenceNumber = 0
	ed.AddAddenda05(addenda05)
	line := bh.String() + "\n" + ed.String() + "\n" + ed.Addenda05[0].String()
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if !base.Has(err, ErrBatchAddendaIndicator) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileAddenda05 tests validating error for an invalid addenda05
func TestFileAddenda05(t *testing.T) {
	testFileAddenda05(t)
}

// BenchmarkFileAddenda05 benchmarks validating error for an invalid addenda05
func BenchmarkFileAddenda05(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileAddenda05(b)
	}
}

// testFileAddenda02invalid validates error for an invalid addenda02
func testFileAddenda02invalid(t testing.TB) {
	bh := mockBatchPOSHeader()
	ed := mockPOSEntryDetail()
	addenda02 := mockAddenda02()
	addenda02.TransactionDate = "0000"
	ed.Addenda02 = addenda02
	line := bh.String() + "\n" + ed.String() + "\n" + ed.Addenda02.String()
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if !base.Has(err, ErrBatchAddendaIndicator) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileAddenda02invalid tests validating error for an invalid addenda02
func TestFileAddenda02invalid(t *testing.T) {
	testFileAddenda02invalid(t)
}

// BenchmarkFileAddenda02invalid benchmarks validating error for an invalid addenda02
func BenchmarkFileAddenda02invalid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileAddenda02invalid(b)
	}
}

// testFileAddenda02 validates a valid addenda02
func testFileAddenda02(t testing.TB) {
	bh := mockBatchPOSHeader()
	ed := mockPOSEntryDetail()
	addenda02 := mockAddenda02()
	ed.Addenda02 = addenda02
	line := bh.String() + "\n" + ed.String() + "\n" + ed.Addenda02.String()
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if !base.Has(err, ErrBatchAddendaIndicator) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileAddenda02invalid tests validating a valid addenda02
func TestFileAddenda02(t *testing.T) {
	testFileAddenda02(t)
}

// BenchmarkFileAddenda02 benchmarks validating a valid addenda02
func BenchmarkFileAddenda02(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileAddenda02(b)
	}
}

// testFileAddenda98 validates error for an invalid addenda98
func testFileAddenda98invalid(t testing.TB) {
	bh := mockBatchPPDHeader()
	ed := mockPPDEntryDetail()
	addenda98 := mockAddenda98()
	addenda98.TraceNumber = "0000001"
	addenda98.ChangeCode = "C50"
	addenda98.CorrectedData = "ACME One Corporation"
	ed.Category = CategoryNOC
	ed.Addenda98 = addenda98
	line := bh.String() + "\n" + ed.String() + "\n" + ed.Addenda98.String()
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if !base.Has(err, ErrBatchAddendaIndicator) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileAddenda98 tests validating error for an invalid addenda98
func TestFileAddenda98invalid(t *testing.T) {
	testFileAddenda98invalid(t)
}

// BenchmarkFileAddenda98 benchmarks validating error for an invalid addenda98
func BenchmarkFileAddenda98invalid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileAddenda98invalid(b)
	}
}

// testFileAddenda98 validates a valid addenda98
func testFileAddenda98(t testing.TB) {
	bh := mockBatchHeader()
	ed := mockEntryDetail()
	addenda98 := mockAddenda98()
	addenda98.TraceNumber = "0000001"
	addenda98.ChangeCode = "C10"
	addenda98.CorrectedData = "ACME One Corporation"
	ed.Category = CategoryNOC
	ed.Addenda98 = addenda98
	line := bh.String() + "\n" + ed.String() + "\n" + ed.Addenda98.String()
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if !base.Has(err, ErrBatchAddendaIndicator) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileAddenda98 tests validating a valid addenda98
func TestFileAddenda98(t *testing.T) {
	testFileAddenda98(t)
}

// BenchmarkFileAddenda98 benchmarks validating a valid addenda98
func BenchmarkFileAddenda98(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileAddenda98(b)
	}
}

// testFileAddenda99invalid validates error for an invalid addenda99
func testFileAddenda99invalid(t testing.TB) {
	bh := mockBatchPPDHeader()
	ed := mockPPDEntryDetail()
	addenda99 := mockAddenda99()
	addenda99.TraceNumber = "0000001"
	addenda99.ReturnCode = "100"
	ed.Category = CategoryReturn
	ed.Addenda99 = addenda99
	line := bh.String() + "\n" + ed.String() + "\n" + ed.Addenda99.String()
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if !base.Has(err, ErrBatchAddendaIndicator) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileAddenda99invalid tests validating error for an invalid addenda99
func TestFileAddenda99invalid(t *testing.T) {
	testFileAddenda99invalid(t)
}

// BenchmarkFileAddenda99invalid benchmarks validating error for an invalid addenda99
func BenchmarkFileAddenda99invalid(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileAddenda99invalid(b)
	}
}

// testFileAddenda99 validates a valid addenda99
func testFileAddenda99(t testing.TB) {
	bh := mockBatchHeader()
	ed := mockEntryDetail()
	addenda99 := mockAddenda99()
	addenda99.TraceNumber = "0000001"
	addenda99.ReturnCode = "R02"
	ed.Category = CategoryReturn
	ed.Addenda99 = addenda99
	line := bh.String() + "\n" + ed.String() + "\n" + ed.Addenda99.String()
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if !base.Has(err, ErrBatchAddendaIndicator) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileAddenda99 tests validating a valid addenda99
func TestFileAddenda99(t *testing.T) {
	testFileAddenda99(t)
}

// BenchmarkFileAddenda99 benchmarks validating a valid addenda99
func BenchmarkFileAddenda99(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileAddenda99(b)
	}
}

// testFileAddendaOutsideBatch validates error populates through the reader
func testFileAddendaOutsideBatch(t testing.TB) {
	ed := mockEntryDetail()
	addenda := mockAddenda05()
	line := ed.String() + "\n" + addenda.String()
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()

	// Note that the entry doesn't get counted since it is rejected due to being outside of a batch
	// So the parser considers the addenda to be outside of an entry since there are no valid entries
	if !base.Has(err, ErrFileAddendaOutsideEntry) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileAddendaOutsideBatch tests validating error populates through the reader
func TestFileAddendaOutsideBatch(t *testing.T) {
	testFileAddendaOutsideBatch(t)
}

// BenchmarkFileAddendaOutsideBatch benchmarks validating error populates through the reader
func BenchmarkFileAddendaOutsideBatch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileAddendaOutsideBatch(b)
	}
}

// testFileAddendaNoIndicator validates no addenda indicator
func testFileAddendaNoIndicator(t testing.TB) {
	bh := mockBatchHeader()
	ed := mockEntryDetail()
	addenda := mockAddenda05()
	line := bh.String() + "\n" + ed.String() + "\n" + addenda.String()
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if !base.Has(err, ErrBatchAddendaIndicator) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileAddendaNoIndicator tests validating no addenda indicator
func TestFileAddendaNoIndicator(t *testing.T) {
	testFileAddendaNoIndicator(t)
}

// BenchmarkFileAddendaNoIndicator benchmarks validating no addenda indicator
func BenchmarkFileAddendaNoIndicator(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileAddendaNoIndicator(b)
	}
}

// testFileFileControlErr validates a file control error
func testFileFileControlErr(t testing.TB) {
	fc := mockFileControl()
	fc.BatchCount = 0
	r := NewReader(strings.NewReader(fc.String()))
	_, err := r.Read()
	if !base.Has(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileFileControlErr tests validating a file control error
func TestFileFileControlErr(t *testing.T) {
	testFileFileControlErr(t)
}

// BenchmarkFileFileControlErr benchmarks validating a file control error
func BenchmarkFileFileControlErr(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileFileControlErr(b)
	}
}

// testFileBatchHeaderSEC validates batch header SEC
func testFileBatchHeaderSEC(t testing.TB) {
	bh := mockBatchHeader()
	bh.StandardEntryClassCode = "ABC"
	r := NewReader(strings.NewReader(bh.String()))
	_, err := r.Read()
	if !base.Has(err, ErrSECCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileBatchHeaderSEC tests validating batch header SEC
func TestFileBatchHeaderSEC(t *testing.T) {
	testFileBatchHeaderSEC(t)
}

// BenchmarkFileBatchHeaderSEC benchmarks validating batch header SEC
func BenchmarkFileBatchHeaderSEC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileBatchHeaderSEC(b)
	}
}

// testFileBatchControlNoCurrentBatch validates no current batch
func testFileBatchControlNoCurrentBatch(t testing.TB) {
	bc := mockBatchControl()
	r := NewReader(strings.NewReader(bc.String()))
	_, err := r.Read()
	if !base.Has(err, ErrFileBatchControlOutsideBatch) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileBatchControlNoCurrentBatch tests validating no current batch
func TestFileBatchControlNoCurrentBatch(t *testing.T) {
	testFileBatchControlNoCurrentBatch(t)
}

// BenchmarkFileBatchControlNoCurrentBatch benchmarks validating no current batch
func BenchmarkFileBatchControlNoCurrentBatch(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileBatchControlNoCurrentBatch(b)
	}
}

// testFileBatchControlValidate validates a batch control
func testFileBatchControlValidate(t testing.TB) {
	bh := mockBatchHeader()
	ed := mockEntryDetail()
	bc := mockBatchControl()
	bc.CompanyIdentification = "B1G C0MP®NY"
	line := bh.String() + "\n" + ed.String() + "\n" + bc.String()
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if !base.Has(err, NewErrBatchHeaderControlEquality(220, 200)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileBatchControlValidate tests validating a batch control
func TestFileBatchControlValidate(t *testing.T) {
	testFileBatchControlValidate(t)
}

// BenchmarkFileBatchControlValidate benchmarks validating a batch control
func BenchmarkFileBatchControlValidate(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileBatchControlValidate(b)
	}
}

// testFileAddBatchValidation validates a batch
func testFileAddBatchValidation(t testing.TB) {
	bh := mockBatchHeader()
	ed := mockEntryDetail()
	bc := mockBatchControl()
	line := bh.String() + "\n" + ed.String() + "\n" + bc.String()
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if !base.Has(err, NewErrBatchCalculatedControlEquality(1, 0)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileAddBatchValidation tests validating a batch
func TestFileAddBatchValidation(t *testing.T) {
	testFileAddBatchValidation(t)
}

// BenchmarkFileAddBatchValidation benchmarks validating a batch
func BenchmarkFileAddBatchValidation(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileAddBatchValidation(b)
	}
}

// testFileLongErr Batch Header Service Class is 000 which does not validate
func testFileLongErr(t testing.TB) {
	line := "101 076401251 0764012510807291511A094101achdestname            companyname                    5000companyname                         origid    PPDCHECKPAYMT000002080730   1076401250000001"
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()
	if !base.Has(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileLongErr tests Batch Header Service Class is 000 which does not validate
func TestFileLongErr(t *testing.T) {
	testFileLongErr(t)
}

// BenchmarkFileLongErr benchmarks Batch Header Service Class is 000 which does not validate
func BenchmarkFileLongErr(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileLongErr(b)
	}
}

// testFileAddendaOutsideEntry validates an addenda is within an entry detail
func testFileAddendaOutsideEntry(t testing.TB) {
	bh := mockBatchHeader()
	addenda := mockAddenda05()
	line := bh.String() + "\n" + addenda.String()
	r := NewReader(strings.NewReader(line))
	_, err := r.Read()

	if !base.Has(err, ErrFileAddendaOutsideEntry) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileAddendaOutsideEntry tests validating an addenda is within an entry detail
func TestFileAddendaOutsideEntry(t *testing.T) {
	testFileAddendaOutsideEntry(t)
}

// BenchmarkFileAddendaOutsideEntry benchmarks validating an addenda is within an entry detail
func BenchmarkFileAddendaOutsideEntry(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileAddendaOutsideEntry(b)
	}
}

// testFileFHImmediateOrigin validates file header immediate origin
func testFileFHImmediateOrigin(t testing.TB) {
	fh := mockFileHeader()
	fh.ImmediateDestination = ""
	r := NewReader(strings.NewReader(fh.String()))
	// necessary to have a file control not nil
	r.File.Control = mockFileControl()
	_, err := r.Read()
	if !base.Has(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestFileFHImmediateOrigin tests validating file header immediate origin
func TestFileFHImmediateOrigin(t *testing.T) {
	testFileFHImmediateOrigin(t)
}

// BenchmarkFileFHImmediateOrigin benchmarks validating file header immediate origin
func BenchmarkFileFHImmediateOrigin(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testFileFHImmediateOrigin(b)
	}
}

// testACHFileRead validates reading a file with PPD and IAT entries
func testACHFileRead(t testing.TB) {
	f, err := os.Open(filepath.Join("test", "testdata", "20110805A.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if err != nil {
		t.Errorf("%T: %s", err, err)
	}

	err2 := r.File.Validate()
	// TODO: is this supposed to have an error here?
	if !base.Match(err2, NewErrFileCalculatedControlEquality("BatchCount", 2, 5)) {
		t.Errorf("%T: %s", err2, err2)
	}
}

// TestACHFileRead tests validating reading a file with PPD and IAT entries
func TestACHFileRead(t *testing.T) {
	testACHFileRead(t)
}

// BenchmarkACHFileRead benchmarks validating reading a file with PPD and IAT entries
func BenchmarkACHFileRead(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testACHFileRead(b)
	}
}

// testACHFileRead2 validates reading a file with PPD and IAT entries
func testACHFileRead2(t testing.TB) {
	f, err := os.Open(filepath.Join("test", "testdata", "20110805A.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if err != nil {
		t.Errorf("%T: %s", err, err)
	}

	err2 := r.File.Validate()
	// TODO: is this supposed to have an error here?
	if !base.Match(err2, NewErrFileCalculatedControlEquality("BatchCount", 2, 5)) {
		t.Errorf("%T: %s", err2, err2)
	}
}

// TestACHFileRead2 tests validating reading a file with PPD and IAT entries
func TestACHFileRead2(t *testing.T) {
	testACHFileRead2(t)
}

// BenchmarkACHFileRead2 benchmarks validating reading a file with PPD and IAT entries
func BenchmarkACHFileRead2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testACHFileRead2(b)
	}
}

// testACHFileRead3 validates reading a file with IAT entries only
func testACHFileRead3(t testing.TB) {
	f, err := os.Open(filepath.Join("test", "testdata", "20180713-IAT.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()

	if file, err := NewReader(f).Read(); err != nil {
		t.Error(err)
	} else {
		if err := file.Validate(); err != nil {
			t.Error(err)
		}
	}
}

// TestACHFileRead3 tests validating reading a file with IAT entries that only
func TestACHFileRead3(t *testing.T) {
	testACHFileRead3(t)
}

// BenchmarkACHFileRead3 benchmarks validating reading a file with IAT entries only
func BenchmarkACHFileRead3(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testACHFileRead3(b)
	}
}

// testACHIATAddenda17 validates reading a file with IAT and Addenda17 entries
func testACHIATAddenda17(t testing.TB) {
	f, err := os.Open(filepath.Join("test", "testdata", "20180716-IAT-A17.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if err != nil {
		t.Errorf("%T: %s", err, err)
	}

	err2 := r.File.Validate()
	if err2 != nil {
		t.Errorf("%T: %s", err2, err2)
	}
}

// TestACHIATAddenda17 tests validating reading a file with IAT and Addenda17 entries that
func TestACHIATAddenda17(t *testing.T) {
	testACHIATAddenda17(t)
}

// BenchmarkACHIATAddenda17  benchmarks validating reading a file with IAT and Addenda17 entries
func BenchmarkACHIATAddenda17(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testACHIATAddenda17(b)
	}
}

// testACHIATAddenda1718 validates reading a file with IAT and Addenda17 and Addenda18 entries
func testACHIATAddenda1718(t testing.TB) {
	f, err := os.Open(filepath.Join("test", "testdata", "20180716-IAT-A17-A18.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if err != nil {
		t.Errorf("%T: %s", err, err)
	}

	err2 := r.File.Validate()
	if err2 != nil {
		t.Errorf("%T: %s", err2, err2)
	}
}

// TestACHIATAddenda1718 tests validating reading a file with IAT and Addenda17 and Addenda18 entries
func TestACHIATAddenda1718(t *testing.T) {
	testACHIATAddenda1718(t)
}

// BenchmarkACHIATAddenda17  benchmarks validating reading a file with IAT Addenda17 and Addenda18 entries
func BenchmarkACHIATAddenda1718(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testACHIATAddenda1718(b)
	}
}

// testACHFileIATBatchHeader validates error when reading an invalid IATBatchHeader
func testACHFileIATBatchHeader(t testing.TB) {
	f, err := os.Open(filepath.Join("test", "testdata", "iat-invalidBatchHeader.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrServiceClass) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestACHFileIATBatchHeader tests validating error when reading an invalid IATBatchHeader
func TestACHFileIATBatchHeader(t *testing.T) {
	testACHFileIATBatchHeader(t)
}

// BenchmarkACHFileIATBatchHeader benchmarks validating error when reading an invalid IATBatchHeader
func BenchmarkACHFileIATBatchHeader(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testACHFileIATBatchHeader(b)
	}
}

// testACHFileIATEntryDetail validates error when reading an invalid IATEntryDetail
func testACHFileIATEntryDetail(t testing.TB) {
	f, err := os.Open(filepath.Join("test", "testdata", "iat-invalidEntryDetail.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrTransactionCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestACHFileIATEntryDetail tests validating error when reading an invalid IATEntryDetail
func TestACHFileIATEntryDetail(t *testing.T) {
	testACHFileIATEntryDetail(t)
}

// BenchmarkACHFileIATEntryDetail benchmarks validating error when reading an invalid IATEntryDetail
func BenchmarkACHFileIATEntryDetail(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testACHFileIATEntryDetail(b)
	}
}

// TestIATAddendaRecordIndicator validates error when reading an invalid IATEntryDetail
func TestIATAddendaRecordIndicator(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "iat-invalidAddendaRecordIndicator.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrIATBatchAddendaIndicator) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestACHFileIATAddenda10 validates error when reading an invalid IATAddenda10
func TestACHFileIATAddenda10(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "iat-invalidAddenda10.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrTransactionTypeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestACHFileIATAddenda11 validates error when reading an invalid IATAddenda10
func TestACHFileIATAddenda11(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "iat-invalidAddenda11.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestACHFileIATAddenda12 validates error when reading an invalid IATAddenda10
func TestACHFileIATAddenda12(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "iat-invalidAddenda12.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestACHFileIATAddenda13 validates error when reading an invalid IATAddenda13
func TestACHFileIATAddenda13(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "Iat-invalidAddenda13.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestACHFileIATAddenda14 validates error when reading an invalid IATAddenda14
func TestACHFileIATAddenda14(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "iat-invalidAddenda14.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestACHFileIATAddenda15 validates error when reading an invalid IATAddenda15
func TestACHFileIATAddenda15(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "iat-invalidAddenda15.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestACHFileIATAddenda16 validates error when reading an invalid IATAddenda16
func TestACHFileIATAddenda16(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "iat-invalidAddenda16.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestACHFileIATAddenda17 validates error when reading an invalid IATAddenda17
func TestACHFileIATAddenda17(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "iat-invalidAddenda17.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestACHFileIATAddenda18 validates error when reading an invalid IATAddenda18
func TestACHFileIATAddenda18(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "iat-invalidAddenda18.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestACHFileIATAddenda98 validates error when reading an invalid IATAddenda98
func TestACHFileIATAddenda98(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "iat-invalidAddenda98.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrAddenda98ChangeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestACHFileIATAddenda99 validates error when reading an invalid IATAddenda99
func TestACHFileIATAddenda99(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "iat-invalidAddenda99.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrAddenda99ReturnCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestPOSInvalidReturnFile validates error when reading an invalid POS Return
func TestPOSInvalidReturnFile(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "pos-invalidReturnFile.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrAddenda99ReturnCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestWEBInvalidNOCFile validates error when reading an invalid WEB NOC
func TestWEBInvalidNOCFile(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "web-invalidNOCFile.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrAddenda98ChangeCode) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestPOSInvalidEntryDetail validates error when reading an invalid POS EntryDetail
func TestPOSInvalidEntryDetail(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "pos-invalidEntryDetail.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrNonAlphanumeric) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestADVInvalidBatchEntries validates error when reading an invalid ADV file with no entries in a batch
func TestADVInvalidBatchEntries(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "adv-invalidBatchEntries.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrFileAddendaOutsideEntry) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestADVNoFileControl validates error when reading an invalid ADV file with no FileControl
func TestADVNoFileControl(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "adv-noFileControl.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrFileControl) {
		t.Errorf("%T: %s", err, err)
	}
}

func TestADVCategoryReturn(t *testing.T) {
	fd, err := os.Open(filepath.Join("test", "testdata", "adv-return.json"))
	if err != nil {
		t.Error(err)
	}
	bs, err := io.ReadAll(fd)
	if err != nil {
		t.Error(err)
	}
	file, err := FileFromJSON(bs)
	if err != nil {
		t.Error(err)
	}

	if n := len(file.Batches); n != 1 {
		t.Errorf("got %d Batches", n)
	}
	if file.Batches[0].Category() != CategoryReturn {
		t.Errorf("Category=%s", file.Batches[0].Category())
	}
}

// testACHFileIATBC validates error when reading an invalid IAT Batch Control
func testACHFileIATBC(t testing.TB) {
	f, err := os.Open(filepath.Join("test", "testdata", "iat-invalidBatchControl.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, NewErrBatchHeaderControlEquality(23138010, 23100000)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestACHFileIATBC tests validating error when reading an invalid IAT Batch Control
func TestACHFileIATBC(t *testing.T) {
	testACHFileIATBC(t)
}

// BenchmarkACHFileIATBC benchmarks validating error when reading an invalid IAT Batch Control
func BenchmarkACHFileIATBC(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testACHFileIATBC(b)
	}
}

// testACHFileIATBH validates error when reading an invalid IAT Batch Header
func testACHFileIATBH(t testing.TB) {
	f, err := os.Open(filepath.Join("test", "testdata", "iat-batchHeaderErr.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrFileBatchHeaderInsideBatch) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestACHFileIATBH tests validating error when reading an invalid IAT Batch Header
func TestACHFileIATBH(t *testing.T) {
	testACHFileIATBH(t)
}

// BenchmarkACHFileIATBH benchmarks validating error when reading an invalid IAT Batch Header
func BenchmarkACHFileIATBH(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		testACHFileIATBH(b)
	}
}

// TestReturnACHFile test loading WEB return file
func TestReturnACHFile(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "return-WEB.ach"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	r := NewReader(f)
	data, err := r.Read()
	if err != nil {
		t.Fatal(err)
	}
	if err := data.Validate(); err != nil {
		t.Fatal(err)
	}
}

// TestReturnACHFile test loading PPD return file with a custom return code
func TestReturnACHFileCustomReasonCode(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "return-PPD-custom-reason-code.ach"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	r := NewReader(f)
	r.SetValidation(&ValidateOpts{
		CustomReturnCodes: true,
	})
	data, err := r.Read()
	if err != nil {
		t.Fatal(err)
	}
	if err := data.Validate(); err != nil {
		t.Fatal(err)
	}
}

// TestADVReturnError returns a Parse Error
func TestADVReturnError(t *testing.T) {
	file := NewFile().SetHeader(mockFileHeader())
	entry := mockADVEntryDetail()
	entry.Addenda99 = mockAddenda99()
	entry.Category = CategoryReturn
	advHeader := mockBatchADVHeader()
	batch := NewBatchADV(advHeader)
	batch.SetHeader(advHeader)
	batch.AddADVEntry(entry)
	if err := batch.Create(); err != nil {
		t.Fatal(err)
	}
	file.AddBatch(batch)

	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	b := &bytes.Buffer{}
	f := NewWriter(b)

	if err := f.Write(file); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	r := NewReader(strings.NewReader(b.String()))
	_, err := r.Read()
	if !base.Has(err, NewErrBatchCalculatedControlEquality(1, 2)) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestADVFileControl validates error when reading an invalid ADV File Control
func TestADVFileControl(t *testing.T) {
	f, err := os.Open(filepath.Join("test", "testdata", "adv-invalidFileControl.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrConstructor) {
		t.Errorf("%T: %s", err, err)
	}
}

// TestTwoFileADVControls validates one file control
func TestTwoFileADVControls(t *testing.T) {
	var line = "9000001000001000000010005320001000000010500000000000000                                       "
	var twoControls = line + "\n" + line
	r := NewReader(strings.NewReader(twoControls))
	r.addCurrentBatch(NewBatchADV(mockBatchADVHeader()))
	bc := ADVBatchControl{
		TotalDebitEntryDollarAmount: 10500,
		EntryHash:                   5320001}
	r.currentBatch.SetADVControl(&bc)

	r.File.AddBatch(r.currentBatch)
	r.File.ADVControl.EntryHash = 5320001

	_, err := r.Read()
	if !base.Has(err, ErrFileControl) {
		t.Errorf("%T: %s", err, err)
	}
}

// testACHFileTooLongErr checks that it errors on a file that is too long
func testACHFileTooLongErr(t testing.TB) {
	// To make testing this more manageable, we'll artificially cap the size of the file to 200 lines
	maxLines = 200

	f, err := os.Open(filepath.Join("test", "testdata", "20110729A.ach"))
	if err != nil {
		t.Errorf("%T: %s", err, err)
	}
	defer f.Close()
	r := NewReader(f)
	_, err = r.Read()

	if !base.Has(err, ErrFileTooLong) {
		t.Errorf("%T: %s", err, err)
	}

	// reset maxLines to its original value
	maxLines = 2 + 2000000 + 100000000 + 8
}

// TestCategoryAssignment ensures entry categories are set correctly
func TestCategoryAssignment(t *testing.T) {
	// Create a file containing entries of each category.
	file := NewFile().SetHeader(mockFileHeader())

	// "Forward"
	forwardEntry := mockEntryDetail()
	forwardEntry.Category = CategoryForward
	forwardEntry.DiscretionaryData = "01"
	forwardBatch := NewBatchWEB(mockBatchWEBHeader())
	forwardBatch.AddEntry(forwardEntry)
	if err := forwardBatch.Create(); err != nil {
		t.Fatal(err)
	}
	file.AddBatch(forwardBatch)

	// "Return"
	returnEntry := mockEntryDetail()
	returnEntry.Addenda99 = mockAddenda99()
	returnEntry.AddendaRecordIndicator = 1
	returnEntry.Category = CategoryReturn
	returnEntry.DiscretionaryData = "02"
	returnBatch := NewBatchWEB(mockBatchWEBHeader())
	returnBatch.AddEntry(returnEntry)
	if err := returnBatch.Create(); err != nil {
		t.Fatal(err)
	}
	file.AddBatch(returnBatch)

	// "DishonoredReturn"
	dishonoredReturnEntry := mockEntryDetail()
	dishonoredReturnEntry.Addenda99Dishonored = mockAddenda99Dishonored()
	dishonoredReturnEntry.AddendaRecordIndicator = 1
	dishonoredReturnEntry.Category = CategoryDishonoredReturn
	dishonoredReturnEntry.DiscretionaryData = "03"
	dishonoredReturnBatch := NewBatchWEB(mockBatchWEBHeader())
	dishonoredReturnBatch.AddEntry(dishonoredReturnEntry)
	if err := dishonoredReturnBatch.Create(); err != nil {
		t.Fatal(err)
	}
	file.AddBatch(dishonoredReturnBatch)

	// "DishonoredReturnContested"
	contestedDishonoredReturnEntry := mockEntryDetail()
	contestedDishonoredReturnEntry.Addenda99Contested = mockAddenda99Contested()
	contestedDishonoredReturnEntry.AddendaRecordIndicator = 1
	contestedDishonoredReturnEntry.Category = CategoryDishonoredReturnContested
	contestedDishonoredReturnEntry.DiscretionaryData = "04"
	contestedDishonoredReturnBatch := NewBatchWEB(mockBatchWEBHeader())
	contestedDishonoredReturnBatch.AddEntry(contestedDishonoredReturnEntry)
	if err := contestedDishonoredReturnBatch.Create(); err != nil {
		t.Fatal(err)
	}
	file.AddBatch(contestedDishonoredReturnBatch)

	// "NOC"
	nocEntry := mockCOREntryDetail()
	nocEntry.Addenda98 = mockAddenda98()
	nocEntry.AddendaRecordIndicator = 1
	nocEntry.Category = CategoryNOC
	nocEntry.DiscretionaryData = "05"
	nocBatch := NewBatchCOR(mockBatchCORHeader())
	nocBatch.AddEntry(nocEntry)
	if err := nocBatch.Create(); err != nil {
		t.Fatal(err)
	}
	file.AddBatch(nocBatch)

	if err := file.Create(); err != nil {
		t.Errorf("%T: %s", err, err)
	}
	if err := file.Validate(); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	b := &bytes.Buffer{}
	f := NewWriter(b)

	if err := f.Write(file); err != nil {
		t.Errorf("%T: %s", err, err)
	}

	r := NewReader(strings.NewReader(b.String()))
	readFile, err := r.Read()
	if err != nil {
		t.Fatal(err)
	}

	categoriesByDiscretionaryData := map[string]string{}
	for _, batch := range readFile.Batches {
		for _, entry := range batch.GetEntries() {
			categoriesByDiscretionaryData[entry.DiscretionaryData] = entry.Category
		}
	}

	require.Equal(t, CategoryForward, categoriesByDiscretionaryData["01"])
	require.Equal(t, CategoryReturn, categoriesByDiscretionaryData["02"])
	require.Equal(t, CategoryDishonoredReturn, categoriesByDiscretionaryData["03"])
	require.Equal(t, CategoryDishonoredReturnContested, categoriesByDiscretionaryData["04"])
	require.Equal(t, CategoryNOC, categoriesByDiscretionaryData["05"])
}

// TestACHFileTooLongErr checks that it errors on a file that is too long
func TestACHFileTooLongErr(t *testing.T) {
	testACHFileTooLongErr(t)
}

func TestReader_AddendaParse(t *testing.T) {
	var line = "702REFONEAREFTERM021000490614123456Target Store 0049          PHILADELPHIA   PA12104288000"
	r := NewReader(strings.NewReader(line))
	r.line = line

	if err := r.parseAddenda(); err != nil {
		if !base.Match(err, ErrFileAddendaOutsideBatch) {
			t.Errorf("%T: %s", err, err)
		}
	}
}

func TestReader__ShortLines(t *testing.T) {
	file, err := readACHFilepath(filepath.Join("test", "testdata", "short-line.ach"))
	if err != nil {
		t.Fatal(err)
	}
	if n := len(file.Batches); n != 1 {
		t.Errorf("got %d batches: %#v", n, file.Batches)
	}
}

func TestReader__LongLine(t *testing.T) {
	file, err := readACHFilepath(filepath.Join("test", "testdata", "long-line.ach"))
	if err != nil {
		t.Fatal(err)
	}
	if n := len(file.Batches); n != 1 {
		t.Errorf("got %d batches: %#v", n, file.Batches)
	}
}

func TestReader__morphing(t *testing.T) {
	out := trimSpacesFromLongLine(strings.Repeat("a", 94) + "    ")
	if len(out) != 94 {
		t.Errorf("out=%q (%d)", out, len(out))
	}

	out = trimSpacesFromLongLine(strings.Repeat("a", 94))
	if len(out) != 94 {
		t.Errorf("out=%q (%d)", out, len(out))
	}

	out, err := rightPadShortLine("")
	if err != nil {
		t.Error(err)
	}
	if len(out) != 94 {
		t.Errorf("out=%q (%d)", out, len(out))
	}

	out, err = rightPadShortLine("aaaaaaaaaa")
	if err != nil {
		t.Error(err)
	}
	if len(out) != 94 {
		t.Errorf("out=%q (%d)", out, len(out))
	}

	out, err = rightPadShortLine(strings.Repeat("a", 95))
	if err == nil {
		t.Error("expected error")
	}
	if len(out) != 95 {
		t.Errorf("out=%q (%d)", out, len(out))
	}
}

func TestReader__partial(t *testing.T) {
	file, err := readACHFilepath(filepath.Join("test", "testdata", "invalid-two-micro-deposits.ach"))
	if err == nil {
		t.Error("expected error")
	}
	if file == nil {
		t.Fatal("nil File")
	}

	// Under our current parser setup we append a lingering Batch to the resulting file
	// if an EntryDetail fails to parse or BatchControl record is missing. A partial File
	// is returend with any records that did parse. This test checks we don't silently break
	// this behavior.

	if n := file.Control.BatchCount; n != 2 {
		t.Errorf("got FileControl.BatchCount=%d", n)
	}
	if len(file.Batches) != 2 {
		t.Fatalf("got %d Batches", len(file.Batches))
	}
	entries := file.Batches[0].GetEntries()
	if len(entries) != 3 {
		t.Errorf("got %d entries", len(entries))
	}
	entries = file.Batches[1].GetEntries()
	if len(entries) != 3 {
		t.Errorf("got %d entries", len(entries))
	}
}

func TestJSONReader__IncludeFieldName(t *testing.T) {
	// Write a test so that we verify the field name is included in the error message so
	// it's easier to debug. Otherwise we get generic error messages that just include
	// "cannot unmarshal %T into %T"
	bs, err := os.ReadFile(filepath.Join("test", "testdata", "invalid-batchNumber.json"))
	if err != nil {
		t.Fatal(err)
	}

	if _, err := FileFromJSON(bs); err == nil {
		t.Fatal("expected error")
	} else {
		if !strings.Contains(err.Error(), `batchHeader.batchNumber: json: cannot unmarshal string into Go struct field BatchHeader.batchHeader.batchNumber of type int`) {
			t.Fatalf("unexpected error:\n  %v", err)
		}
	}
}
