package server

import (
	"errors"
	"strings"

	"github.com/moov-io/ach"
	uuid "github.com/satori/go.uuid"
)

var (
	ErrNotFound      = errors.New("Not Found")
	ErrAlreadyExists = errors.New("Already Exists")
)

// Service is a REST interface for interacting with ACH file structures
// TODO: Add ctx to function parameters to pass the client security token
type Service interface {
	// CreateFile creates a new ach file record and returns a resource ID
	CreateFile(f ach.File) (string, error)
	// AddFile retrieves a file based on the File id
	GetFile(id string) (ach.File, error)
	// GetFiles retrieves all files accessible from the client.
	GetFiles() []ach.File
	// DeleteFile takes a file resource ID and deletes it from the repository
	DeleteFile(id string) error
	// UpdateFile updates the changes properties of a matching File ID
	// UpdateFile(f ach.File) (string, error)
}

// service a concrete implementation of the service.
type service struct {
	repository Repository
}

// NewService creates a new concrete service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// CreateFile add a file to storage
func (s *service) CreateFile(f ach.File) (string, error) {
	if f.ID == "" {
		f.ID = NextID()
	}
	if err := s.repository.StoreFile(&f); err != nil {
		return "", err
	}
	return f.ID, nil
}

// GetFile returns a files based on the supplied id
func (s *service) GetFile(id string) (ach.File, error) {
	f, err := s.repository.FindFile(id)
	if err != nil {
		return ach.File{}, ErrNotFound
	}
	return *f, nil
}

func (s *service) GetFiles() []ach.File {
	var result []ach.File
	for _, f := range s.repository.FindAllFiles() {
		result = append(result, *f)
	}
	return result
}

func (s *service) DeleteFile(id string) error {
	return s.repository.DeleteFile(id)
}

// Repository concrete implementations
// ********

// Repository is the Service storage mechanism abstraction
type Repository interface {
	StoreFile(file *ach.File) error
	FindFile(id string) (*ach.File, error)
	FindAllFiles() []*ach.File
	DeleteFile(id string) error
}

// Utility Functions
// *****

// NextID generates a new resource ID
func NextID() string {
	id, _ := uuid.NewV4()
	//return id.String()
	// make it shorter for testing URL
	return string(strings.Split(strings.ToUpper(id.String()), "-")[0])
}
