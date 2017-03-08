package ach

import (
	"testing"
)

func TestPPDWrite(t *testing.T) {
	//fileHead := mockFileHeader()
	//batchHead := mockBatchHeader()
	//entryDetail := mockEntryDetail()
	//file := NewF
	file := NewFile().setHeader(mockFileHeader())
	batch := NewBatch().setHeader(mockBatchHeader()).addEntryDetail(mockEntryDetail())
	file.addBatch(batch)
	//.addBatch(NewBatch().setHeader(mockBatchHeader()))
	//fmt.Printf("file: %+v", file)
	//file.addBatch(batch)
	//file.
	//w := NewWriter()

}
