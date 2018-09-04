package server

import (
	"sync"

	"github.com/moov-io/ach"
)

// Repository is the Service storage mechanism abstraction
type Repository interface {
	StoreFile(file *ach.File) error
	FindFile(id string) (*ach.File, error)
	FindAllFiles() []*ach.File
	DeleteFile(id string) error
	StoreBatch(fileID string, batch ach.Batcher) error
	FindBatch(fileID string, batchID string) (ach.Batcher, error)
	FindAllBatches(fileID string) []ach.Batcher
	DeleteBatch(fileID string, batchID string) error
}

type repositoryInMemory struct {
	mtx   sync.RWMutex
	files map[string]*ach.File
}

// NewRepositoryInMemory is an in memory ach storage repository for files
func NewRepositoryInMemory() Repository {
	f := map[string]*ach.File{}
	return &repositoryInMemory{
		files: f,
	}
}
func (r *repositoryInMemory) StoreFile(f *ach.File) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if _, ok := r.files[f.ID]; ok {
		return ErrAlreadyExists
	}
	r.files[f.ID] = f
	return nil
}

// FindFile retrieves a ach.File based on the supplied ID
func (r *repositoryInMemory) FindFile(id string) (*ach.File, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if val, ok := r.files[id]; ok {
		return val, nil
	}
	return nil, ErrNotFound
}

// FindAllFiles returns all files that have been saved in memory
func (r *repositoryInMemory) FindAllFiles() []*ach.File {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	files := make([]*ach.File, 0, len(r.files))
	for _, val := range r.files {
		files = append(files, val)
	}
	return files
}

func (r *repositoryInMemory) DeleteFile(id string) error {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	delete(r.files, id)
	return nil
}

// TODO(adam): was copying this causing issues?
func (r *repositoryInMemory) StoreBatch(fileID string, batch ach.Batcher) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	// Ensure the file does not already exist
	if _, ok := r.files[fileID]; !ok {
		return ErrNotFound
	}
	// ensure the batch does not already exist
	for _, val := range r.files[fileID].Batches {
		if val.ID() == batch.ID() {
			return ErrAlreadyExists
		}
	}
	// Add the batch to the file
	r.files[fileID].AddBatch(batch)
	return nil
}

// FindBatch retrieves a ach.Batcher based on the supplied ID
func (r *repositoryInMemory) FindBatch(fileID string, batchID string) (ach.Batcher, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	for _, val := range r.files[fileID].Batches {
		if val.ID() == batchID {
			return val, nil
		}
	}
	return nil, ErrNotFound
}

// FindAllBatches
func (r *repositoryInMemory) FindAllBatches(fileID string) []ach.Batcher {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	batches := make([]ach.Batcher, 0, len(r.files[fileID].Batches))
	for _, val := range r.files[fileID].Batches {
		batches = append(batches, val)
	}
	return batches
}

func (r *repositoryInMemory) DeleteBatch(fileID string, batchID string) error {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	for i := len(r.files[fileID].Batches) - 1; i >= 0; i-- {
		if r.files[fileID].Batches[i].ID() == batchID {
			r.files[fileID].Batches = append(r.files[fileID].Batches[:i], r.files[fileID].Batches[i+1:]...)
			//fmt.Println(r.files[fileID].Batches)
			return nil
		}
	}
	return ErrNotFound
}
