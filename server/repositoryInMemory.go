package server

// TODO: rename to InMemory and move into a repository directory when stable.
import (
	"sync"

	"github.com/moov-io/ach"
)

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

func (r *repositoryInMemory) StoreBatch(fileID string, batch ach.Batcher) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if _, ok := r.files[fileID]; !ok {
		return ErrNotFound
	}
	// TODO range batches to see if batch already exist and error if it does
	r.files[fileID].AddBatch(batch)
	return nil
}
