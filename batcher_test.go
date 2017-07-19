package ach

import (
	"testing"
)

func TestBatcher(t *testing.T) {
	batch := NewBatch()
	whatABatch(batch)

}

func whatABatch(batch Batcher) {
	//fmt.Printf("Batch Header: %v", batch.GetHeader().StandardEntryClassCode)
}
