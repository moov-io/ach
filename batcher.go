// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

// Batcher abstract the different ACH batch types that can exist in a file.
// Each batch type is defined by SEC (Standard Entry Class) code in the Batch Header
// * SEC identifies the payment type (product) found within an ACH batch-using a 3-character code
// * The SEC Code pertains to all items within batch
//    * Determines format of the entry detail records
//    * Determines addenda records (required or optional PLUS one or up to 9,999 records)
//    * Determines rules to follow (return time frames)
// 	  * Some SEC codes require specific data in predetermined fields within the ACH record
type Batcher interface {
	GetHeader() *BatchHeader
	SetHeader(*BatchHeader)
	GetControl() *BatchControl
	SetControl(*BatchControl)
	GetADVControl() *ADVBatchControl
	SetADVControl(*ADVBatchControl)
	GetEntries() []*EntryDetail
	AddEntry(*EntryDetail)
	GetADVEntries() []*ADVEntryDetail
	AddADVEntry(*ADVEntryDetail)
	Create() error
	Validate() error
	SetID(string)
	ID() string
	// Category defines if a Forward or Return
	Category() string
	Error(string, error, ...interface{}) error
	Equal(other Batcher) bool
}
