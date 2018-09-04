package server

import (
	"errors"
	"strings"

	"github.com/moov-io/ach"

	uuid "github.com/gofrs/uuid/v3"
)

var (
	ErrNotFound      = errors.New("Not Found")
	ErrAlreadyExists = errors.New("Already Exists")
)

// Service is a REST interface for interacting with ACH file structures
// TODO: Add ctx to function parameters to pass the client security token
type Service interface {
	// CreateFile creates a new ach file record and returns a resource ID
	CreateFile(f ach.FileHeader) (string, error)
	// AddFile retrieves a file based on the File id
	GetFile(id string) (ach.File, error)
	// GetFiles retrieves all files accessible from the client.
	GetFiles() []ach.File
	// DeleteFile takes a file resource ID and deletes it from the store
	DeleteFile(id string) error
	// UpdateFile updates the changes properties of a matching File ID
	// UpdateFile(f ach.File) (string, error)

	// CreateBatch creates a new batch within and ach file and returns its resource ID
	CreateBatch(fileID string, bh ach.BatchHeader) (string, error)
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
func (s *service) CreateFile(fh ach.FileHeader) (string, error) {
	// create a new file
	f := ach.NewFile()
	f.SetHeader(fh)
	// set resource id's
	if fh.ID == "" {
		id := NextID()
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
func (s *service) GetFile(id string) (ach.File, error) {
	f, err := s.store.FindFile(id)
	if err != nil {
		return ach.File{}, ErrNotFound
	}
	return *f, nil
}

func (s *service) GetFiles() []ach.File {
	var result []ach.File
	for _, f := range s.store.FindAllFiles() {
		result = append(result, *f)
	}
	return result
}

func (s *service) DeleteFile(id string) error {
	return s.store.DeleteFile(id)
}

func (s *service) CreateBatch(fileID string, bh ach.BatchHeader) (string, error) {
	batch, err := ach.NewBatch(&bh)
	if err != nil {
		return bh.ID, err
	}
	if bh.ID == "" {
		id := NextID()
		batch.SetID(id)
		batch.GetHeader().ID = id
		batch.GetControl().ID = id
	} else {
		batch.SetID(bh.ID)
		batch.GetControl().ID = bh.ID
	}

	if err := s.store.StoreBatch(fileID, batch); err != nil {
		return "", err
	}
	return bh.ID, nil
}

func (s *service) GetBatch(fileID string, batchID string) (ach.Batcher, error) {
	b, err := s.store.FindBatch(fileID, batchID)
	if err != nil {
		return nil, ErrNotFound
	}
	return *b, nil
}

func (s *service) GetBatches(fileID string) []ach.Batcher {
	var result []ach.Batcher
	for _, b := range s.store.FindAllBatches(fileID) {
		result = append(result, *b)
	}
	return result
}

func (s *service) DeleteBatch(fileID string, batchID string) error {
	return s.store.DeleteBatch(fileID, batchID)
}

// Utility Functions

// NextID generates a new resource ID
func NextID() string {
	id, _ := uuid.NewV4()
	//return id.String()
	// make it shorter for testing URL
	return string(strings.Split(strings.ToUpper(id.String()), "-")[0])
}
