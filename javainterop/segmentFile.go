package javainterop

import (
	"github.com/moov-io/ach"
)

type File ach.File

type SegmentFileResponse struct {
	CreditEntryFile  *File
	DebitEntryFile   *File
	SegmentFileError error
}

func (f *File) SegmentFile(config *ach.SegmentFileConfiguration) SegmentFileResponse {
	creditFile, debitFile, err := ach.NewFile().SegmentFile(config)

	return SegmentFileResponse{
		CreditEntryFile:  (*File)(creditFile),
		DebitEntryFile:   (*File)(debitFile),
		SegmentFileError: err,
	}
}
