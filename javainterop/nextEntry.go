package javainterop

import (
	"io"

	"github.com/moov-io/ach"
)

type EntryDetail ach.EntryDetail

type BatchHeader ach.BatchHeader

type Iterator ach.Iterator

type NextEntryResponse struct {
	NextBatchHeader BatchHeader
	NextEntryDetail EntryDetail
	NextEntryError  error
}

func (i *Iterator) NextEntry(r io.Reader) NextEntryResponse {
	batchHeader, entryFile, err := ach.NewIterator(r).NextEntry()

	return NextEntryResponse{
		NextBatchHeader: BatchHeader(*batchHeader),
		NextEntryDetail: EntryDetail(*entryFile),
		NextEntryError:  err,
	}
}
