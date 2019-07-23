// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package server

import (
	"io/ioutil"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/moov-io/ach"
)

// test mocks are in mock_test.go

// CreateFile tests
func TestCreateFile(t *testing.T) {
	s := mockServiceInMemory()
	id, err := s.CreateFile(mockFileHeader())
	if err != nil {
		t.Fatal(err.Error())
	}
	if id != "12345" {
		t.Errorf("expected %s received %s w/ error %s", "12345", id, err)
	}
}
func TestCreateFileIDExists(t *testing.T) {
	s := mockServiceInMemory()
	h := ach.FileHeader{ID: "98765"}
	id, err := s.CreateFile(&h)
	if err != ErrAlreadyExists {
		t.Errorf("expected %s received %s w/ error %s", "ErrAlreadyExists", id, err)
	}
}

func TestCreateFileNoID(t *testing.T) {
	s := mockServiceInMemory()
	h := ach.NewFileHeader()
	id, err := s.CreateFile(&h)
	if len(id) < 3 {
		t.Errorf("expected %s received %s w/ error %s", "NextID", id, err)
	}
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Service.GetFile tests

func TestGetFile(t *testing.T) {
	s := mockServiceInMemory()
	f, err := s.GetFile("98765")
	if err != nil {
		t.Errorf("expected %s received %s w/ error %s", "98765", f.ID, err)
	}
}

func TestGetFileNotFound(t *testing.T) {
	s := mockServiceInMemory()
	f, err := s.GetFile("12345")
	if err != ErrNotFound {
		t.Errorf("expected %s received %s w/ error %s", "ErrNotFound", f.ID, err)
	}
}

// Service.GetFiles tests

func TestGetFiles(t *testing.T) {
	s := mockServiceInMemory()
	files := s.GetFiles()
	if len(files) != 1 {
		t.Errorf("expected %s received %v", "1", len(files))
	}
}

// Service.DeleteFile tests

func TestDeleteFile(t *testing.T) {
	s := mockServiceInMemory()
	err := s.DeleteFile("98765")
	if err != nil {
		t.Errorf("expected %s received %s", "nil", err)
	}
	_, err = s.GetFile("98765")
	if err != ErrNotFound {
		t.Errorf("expected %s received %s", "ErrNotFound", err)
	}
}

// Service.GetFileContents tests

func TestGetFileContents(t *testing.T) {
	s := mockServiceInMemory()
	id, err := s.CreateFile(mockFileHeader())
	if err != nil {
		t.Fatal(err.Error())
	}

	// make the file valid
	batch := mockBatchWEB()
	s.CreateBatch(id, batch)

	// build file
	r, err := s.GetFileContents(id)
	if err != nil {
		if !strings.Contains(err.Error(), "mandatory ") {
			t.Fatal(err.Error())
		}
	}
	if r != nil {
		bs, err := ioutil.ReadAll(r)
		if err != nil {
			t.Fatal(err.Error())
		}

		if len(bs) == 0 {
			t.Fatal("expected to read fil")
		}
	}
}

// Service.ValidateFile tests

func TestValidateFile(t *testing.T) {
	s := mockServiceInMemory()
	id, err := s.CreateFile(mockFileHeader())
	if err != nil {
		t.Fatal(err.Error())
	}
	if err := s.ValidateFile(id); err != nil {
		if !strings.Contains(err.Error(), "mandatory ") {
			t.Fatal(err.Error())
		}
	}
}

func TestValidateFileMissing(t *testing.T) {
	s := mockServiceInMemory()
	err := s.ValidateFile("missing")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestValidateFileBad(t *testing.T) {
	s := mockServiceInMemory()

	fId, _ := s.CreateFile(mockFileHeader())

	// setup batch
	bh := mockBatchHeaderWeb()
	bh.ID = "11111"
	b, _ := ach.NewBatch(bh)
	bId, e1 := s.CreateBatch(fId, b)
	batch, e2 := s.GetBatch(fId, bId)
	if batch == nil {
		t.Fatalf("couldn't get batch, e1=%v, e2=%v", e1, e2)
	}

	// setup file, add batch
	f, err := s.GetFile(fId)
	if f == nil {
		t.Fatalf("couldn't get file: %v", err)
	}
	if len(f.AddBatch(batch)) == 0 {
		t.Fatal("problem adding batch to file")
	}

	// validate
	if err := s.ValidateFile(fId); err == nil {
		t.Fatal("expected error")
	}
}

// Service.CreateBatch tests

// TestCreateBatch tests creating a new batch when file.ID exists and batch.id does not exist
func TestCreateBatch(t *testing.T) {
	s := mockServiceInMemory()
	bh := mockBatchHeaderWeb()
	bh.ID = "11111"
	b, _ := ach.NewBatch(bh)
	id, err := s.CreateBatch("98765", b)
	if err != nil {
		t.Fatal(err.Error())
	}
	if id != "11111" {
		t.Errorf("expected %s received %s w/ error %v", "11111", id, err)
	}
}

// TestCreateBatchIDExists Create a new batch with batch.id already present. Should fail.
func TestCreateBatchIDExists(t *testing.T) {
	s := mockServiceInMemory()
	b, _ := ach.NewBatch(mockBatchHeaderWeb())
	id, err := s.CreateBatch("98765", b)
	if err != ErrAlreadyExists {
		t.Errorf("expected %s received %s w/ error %v", "ErrAlreadyExists", id, err)
	}
}

// TestCreateBatchFileIDExits create a batch when the file.id does not exist. Should fail.
func TestCreateBatchFileIDExits(t *testing.T) {
	s := mockServiceInMemory()
	b, _ := ach.NewBatch(mockBatchHeaderWeb())
	id, err := s.CreateBatch("55555", b)
	if err != ErrNotFound {
		t.Errorf("expected %s received %s w/ error %v", "ErrNotFound", id, err)
	}
}

// TestCreateBatchIDBank create a new batch when the batch.id is nil but file.id is valid. Should generate batch.id and save.
func TestCreateBatchIDBlank(t *testing.T) {
	s := mockServiceInMemory()
	bh := mockBatchHeaderWeb()
	bh.ID = ""
	b, _ := ach.NewBatch(bh)
	id, err := s.CreateBatch("98765", b)
	if len(id) < 3 {
		t.Errorf("expected %s received %s w/ error %v", "NextID", id, err)
	}
	if err != nil {
		t.Fatal(err.Error())
	}
}

// Service.GetBatch

// TestGetBatch return a batch for the existing file.id and batch.id
func TestGetBatch(t *testing.T) {
	s := mockServiceInMemory()
	b, err := s.GetBatch("98765", "54321")
	if err != nil {
		t.Errorf("problem getting batch: %v", err)
	}
	if b.ID() != "54321" {
		t.Errorf("expected %s received %s w/ error %v", "54321", b.ID(), err)
	}
}

// TestGetBatchNotFound return a failure if the batch.id is not found
func TestGetBatchNotFound(t *testing.T) {
	s := mockServiceInMemory()
	b, err := s.GetBatch("98765", "55555")
	if err != ErrNotFound {
		t.Errorf("expected %s received %#v w/ error %v", "ErrNotFound", b, err)
	}
}

// Service.GetBatches

// TestGetBatches return a list of batches for the supplied file.id
func TestGetBatches(t *testing.T) {
	s := mockServiceInMemory()
	batches := s.GetBatches("98765")
	if len(batches) != 1 {
		t.Errorf("expected %s received %v", "1", len(batches))
	}
}

// Service.DeleteBatch

// TestDeleteBatch removes a batch with existing file and batch id.
func TestDeleteBatch(t *testing.T) {
	s := mockServiceInMemory()
	err := s.DeleteBatch("98765", "54321")
	if err != nil {
		t.Errorf("expected %s received error %v", "nil", err)
	}
}

// TestSegmentFile creates a Segmented File from an existing ACH File
func TestSegmentFile(t *testing.T) {
	s := mockServiceInMemory()

	fh := ach.NewFileHeader()
	fh.ID = "333339"
	fh.ImmediateDestination = "231380104"
	fh.ImmediateOrigin = "121042882"
	fh.FileCreationDate = time.Now().Format("060102")
	fh.FileCreationTime = time.Now().AddDate(0, 0, 1).Format("1504") // HHmm
	fh.ImmediateDestinationName = "Federal Reserve Bank"
	fh.ImmediateOriginName = "My Bank Name"

	bh := ach.NewBatchHeader()
	bh.ServiceClassCode = ach.MixedDebitsAndCredits
	bh.CompanyName = "Name on Account"
	bh.CompanyIdentification = fh.ImmediateOrigin
	bh.StandardEntryClassCode = ach.PPD
	bh.CompanyEntryDescription = "REG.SALARY"
	bh.EffectiveEntryDate = time.Now().AddDate(0, 0, 1).Format("060102")
	bh.ODFIIdentification = "121042882"
	bh.ID = "433333"
	b, _ := ach.NewBatch(bh)

	entryOne := ach.NewEntryDetail()
	entryOne.TransactionCode = ach.CheckingDebit
	entryOne.SetRDFI("231380104")
	entryOne.DFIAccountNumber = "123456789"
	entryOne.Amount = 200000000
	entryOne.SetTraceNumber(bh.ODFIIdentification, 1)
	entryOne.IndividualName = "Debit Account"

	entryTwo := ach.NewEntryDetail()
	entryTwo.TransactionCode = ach.CheckingCredit
	entryTwo.SetRDFI("231380104")
	entryTwo.DFIAccountNumber = "987654321"
	entryTwo.Amount = 100000000
	entryTwo.SetTraceNumber(bh.ODFIIdentification, 2)
	entryTwo.IndividualName = "Credit Account 1"

	entryThree := ach.NewEntryDetail()
	entryThree.TransactionCode = ach.CheckingCredit
	entryThree.SetRDFI("231380104")
	entryThree.DFIAccountNumber = "837098765"
	entryThree.Amount = 100000000
	entryThree.SetTraceNumber(bh.ODFIIdentification, 3)
	entryThree.IndividualName = "Credit Account 2"

	b.AddEntry(entryOne)
	b.AddEntry(entryTwo)
	b.AddEntry(entryThree)
	if err := b.Create(); err != nil {
		t.Fatalf("Unexpected error building batch: %s\n", err)
	}

	file := ach.NewFile()
	file.SetHeader(fh)
	file.AddBatch(b)
	if err := file.Create(); err != nil {
		log.Fatalf("Unexpected error building file: %s\n", err)
	}

	fileID, err := s.CreateFile(&fh)
	if err != nil {
		t.Fatal(err.Error())
	}

	batchID, err := s.CreateBatch("333339", b)
	if err != nil {
		t.Fatal(err.Error())
	}

	if batchID == "" {
		t.Fatal("No Batch ID")
	}

	f, err := s.GetFile(fileID)

	f.Create()

	if err != nil {
		t.Fatalf("expected %s received %s w/ error %s", "111333", f.ID, err)
	}

	creditFile, debitFile, err := s.SegmentFile(f)

	if err != nil {
		t.Fatalf("could not segment file w/ error %v", err)
	}

	if creditFile == nil {
		t.Fatal("No credit File")
	}

	if debitFile == nil {
		t.Fatal("No debit File")
	}

}

// ToDo: create helper function to use for additional tests?
