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

package server

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/moov-io/ach"
	"github.com/moov-io/base"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

// Service is a REST interface for interacting with ACH file structures
// TODO: Add ctx to function parameters to pass the client security token
type Service interface {
	// CreateFile creates a new ach file record and returns a resource ID
	CreateFile(f *ach.FileHeader) (string, error)
	// AddFile retrieves a file based on the File id
	GetFile(id string) (*ach.File, error)
	// GetFiles retrieves all files accessible from the client.
	GetFiles() []*ach.File
	// BuildFile tabulates file values according to the Nacha spec
	BuildFile(id string) (*ach.File, error)
	// DeleteFile takes a file resource ID and deletes it from the store
	DeleteFile(id string) error
	// GetFileContents creates a valid plaintext file in memory assuming it has a FileHeader and at least one Batch record.
	GetFileContents(id string) (io.Reader, error)
	// ValidateFile
	ValidateFile(id string, opts *ach.ValidateOpts) error
	// BalanceFile will apply a given offset record to the file
	BalanceFile(fileID string, off *ach.Offset) (*ach.File, error)
	// SegmentFileID segments an ach file
	SegmentFileID(id string, opts *ach.SegmentFileConfiguration) (*ach.File, *ach.File, error)
	// SegmentFile segments an ach file
	SegmentFile(file *ach.File, opts *ach.SegmentFileConfiguration) (*ach.File, *ach.File, error)
	// FlattenBatches will minimize the ach.Batch objects in a file by consolidating EntryDetails under distinct batch headers
	FlattenBatches(id string) (*ach.File, error)
	// CreateBatch creates a new batch within and ach file and returns its resource ID
	CreateBatch(fileID string, bh ach.Batcher) (string, error)
	// GetBatch retrieves a batch based oin the file id and batch id
	GetBatch(fileID string, batchID string) (ach.Batcher, error)
	// GetBatches retrieves all batches associated with the file id.
	GetBatches(fileID string) []ach.Batcher
	// DeleteBatch takes a fileID and BatchID and removes the batch from the file
	DeleteBatch(fileID string, batchID string) error
}

// service a concrete implementation of the service.
type service struct {
	store Repository
}

// NewService creates a new concrete service
func NewService(r Repository) Service {
	return &service{
		store: r,
	}
}

// CreateFile add a file to storage
// TODO(adam): the HTTP endpoint accepts malformed bodies (and missing data)
func (s *service) CreateFile(fh *ach.FileHeader) (string, error) {
	// create a new file
	f := ach.NewFile()
	f.SetHeader(*fh)
	// set resource id's
	if fh.ID == "" {
		id := base.ID()
		f.ID = id
		f.Header.ID = id
		f.Control.ID = id
	} else {
		f.ID = fh.ID
		f.Control.ID = fh.ID
	}

	if err := s.store.StoreFile(f); err != nil {
		return "", err
	}
	return f.ID, nil
}

// GetFile returns a files based on the supplied id
func (s *service) GetFile(id string) (*ach.File, error) {
	f, err := s.store.FindFile(id)
	if err != nil {
		return nil, ErrNotFound
	}
	return f, nil
}

func (s *service) GetFiles() []*ach.File {
	return s.store.FindAllFiles()
}

// BuildFile tabulates file values according to the Nacha spec
func (s *service) BuildFile(id string) (*ach.File, error) {
	file, err := s.GetFile(id)
	if err != nil {
		return nil, fmt.Errorf("build file: error reading file %s: %v", id, err)
	}
	err = file.Create()
	return file, err
}

func (s *service) DeleteFile(id string) error {
	return s.store.DeleteFile(id)
}

func (s *service) GetFileContents(id string) (io.Reader, error) {
	f, err := s.GetFile(id)
	if err != nil {
		return nil, fmt.Errorf("problem reading file %s: %v", id, err)
	}
	if err := f.Create(); err != nil {
		return nil, fmt.Errorf("problem creating file %s: %v", id, err)
	}

	var buf bytes.Buffer
	w := ach.NewWriter(&buf)
	if err := w.Write(f); err != nil {
		return nil, fmt.Errorf("problem writing plaintext file %s: %v", id, err)
	}
	if err := w.Flush(); err != nil {
		return nil, err
	}

	if buf.Len() == 0 {
		return nil, errors.New("empty ACH file contents")
	}

	return &buf, nil
}

func (s *service) ValidateFile(id string, opts *ach.ValidateOpts) error {
	f, err := s.GetFile(id)
	if err != nil {
		return fmt.Errorf("problem reading file %s: %v", id, err)
	}
	return f.ValidateWith(opts)
}

func (s *service) CreateBatch(fileID string, batch ach.Batcher) (string, error) {
	if batch == nil {
		return "", errors.New("no batch provided")
	}
	if batch.GetHeader().ID == "" {
		id := base.ID()
		batch.SetID(id)
		batch.GetHeader().ID = id
		batch.GetControl().ID = id
	} else {
		batch.SetID(batch.GetHeader().ID)
		batch.GetControl().ID = batch.GetHeader().ID
	}
	if err := s.store.StoreBatch(fileID, batch); err != nil {
		return "", err
	}
	return batch.ID(), nil
}

func (s *service) GetBatch(fileID string, batchID string) (ach.Batcher, error) {
	b, err := s.store.FindBatch(fileID, batchID)
	if err != nil {
		return nil, ErrNotFound
	}
	return b, nil
}

func (s *service) GetBatches(fileID string) []ach.Batcher {
	return s.store.FindAllBatches(fileID)
}

func (s *service) DeleteBatch(fileID string, batchID string) error {
	return s.store.DeleteBatch(fileID, batchID)
}

func (s *service) BalanceFile(fileID string, off *ach.Offset) (*ach.File, error) {
	f, err := s.GetFile(fileID)
	if err != nil {
		return nil, err
	}
	if err := f.Create(); err != nil {
		return nil, err
	}
	// Apply the Offset to each Batch and then re-create (to tabulate new EntryDetail records)
	for i := range f.Batches {
		f.Batches[i].WithOffset(off)
		if err := f.Batches[i].Create(); err != nil {
			return nil, err
		}
	}
	f.ID = base.ID() // overwrite the ID so it's new and unique
	if err := f.Create(); err != nil {
		return nil, err
	}
	// Save our new file
	if err := s.store.StoreFile(f); err != nil {
		return nil, err
	}
	return f, nil
}

// SegmentFileID takes an ACH FileID and segments the files into a credit ACH File and debit ACH File and adds to in memory storage.
func (s *service) SegmentFileID(fileID string, opts *ach.SegmentFileConfiguration) (*ach.File, *ach.File, error) {
	f, err := s.GetFile(fileID)
	if err != nil {
		return nil, nil, err
	}
	return s.SegmentFile(f, opts)
}

// SegmentFile takes an ACH File and segments the files into a credit ACH File and debit ACH File and adds to in memory storage.
func (s *service) SegmentFile(file *ach.File, opts *ach.SegmentFileConfiguration) (*ach.File, *ach.File, error) {
	// Build/tabulate file in the case it is malformed.
	if err := file.Create(); err != nil {
		return nil, nil, err
	}

	creditFile, debitFile, err := file.SegmentFile(opts)
	if err != nil {
		return nil, nil, err
	}
	return creditFile, debitFile, nil
}

// FlattenBatches consolidates batches that have the same BatchHeader
func (s *service) FlattenBatches(fileID string) (*ach.File, error) {
	f, err := s.GetFile(fileID)
	if err != nil {
		return nil, err
	}
	// File Create in the case a file is malformed.
	if err := f.Create(); err != nil {
		return nil, err
	}
	ff, err := f.FlattenBatches()
	if err != nil {
		return nil, err
	}
	return ff, err
}
