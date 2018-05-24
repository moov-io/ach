package server

import (
	"testing"

	"github.com/moov-io/ach"
)

func mockServiceInMemory() Service {
	repository := NewRepositoryInMemory()
	repository.StoreFile(&ach.File{ID: "98765"})
	return NewService(repository)
}

// CreateFile tests
func TestCreateFile(t *testing.T) {
	s := mockServiceInMemory()
	id, err := s.CreateFile(ach.File{ID: "12345"})
	if id != "12345" {
		t.Errorf("expected %s received %s w/ error %s", "12345", id, err)
	}
}
func TestCreateFileIDExists(t *testing.T) {
	s := mockServiceInMemory()
	id, err := s.CreateFile(ach.File{ID: "98765"})
	if err != ErrAlreadyExists {
		t.Errorf("expected %s received %s w/ error %s", "ErrAlreadyExists", id, err)
	}
}

func TestCreateFileNoID(t *testing.T) {
	s := mockServiceInMemory()
	id, err := s.CreateFile(*ach.NewFile())
	if len(id) < 3 {
		t.Errorf("expected %s received %s w/ error %s", "NextID", id, err)
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
