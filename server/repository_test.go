package server

import (
	"testing"

	"github.com/moov-io/ach"
)

func TestRepositoryFiles(t *testing.T) {
	r := NewRepositoryInMemory()

	if v := len(r.FindAllFiles()); v != 0 {
		t.Errorf("unexpected length: %d", v)
	}

	header := mockFileHeader()
	f := &ach.File{
		ID:     NextID(),
		Header: *header,
	}
	if err := r.StoreFile(f); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	found, err := r.FindFile(f.ID)
	if err != nil || found == nil {
		t.Errorf("found=%v, err=%v", found, err)
	}

	if v := len(r.FindAllFiles()); v != 1 {
		t.Errorf("unexpected length: %d", v)
	}

	if err := r.DeleteFile(f.ID); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestRepositoryBatches(t *testing.T) {
	r := NewRepositoryInMemory()

	// make sure our tests are setup
	if v := len(r.FindAllFiles()); v != 0 {
		t.Errorf("unexpected length: %d", v)
	}

	header := mockFileHeader()
	f := &ach.File{
		ID:     NextID(),
		Header: *header,
	}
	if err := r.StoreFile(f); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// batch tests
	if v := len(r.FindAllBatches(f.ID)); v != 0 {
		t.Errorf("unexpected length: %d", v)
	}

	batch := mockBatchWEB()
	b, err := r.FindBatch(f.ID, batch.ID())
	if err == nil || b != nil {
		t.Errorf("b=%v, err=%v", b, err)
	}

	if err := r.StoreBatch(f.ID, batch); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if v := len(r.FindAllBatches(f.ID)); v != 1 {
		t.Errorf("unexpected length: %d", v)
	}

	if err := r.DeleteBatch(f.ID, batch.ID()); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if v := len(r.FindAllBatches(f.ID)); v != 0 {
		t.Errorf("unexpected length: %d", v)
	}
}
