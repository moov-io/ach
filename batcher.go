package ach

// Batcher abstract the different ACH batch types that can exist in a file.
// Each batch type is defined by SEC (Standard Entry Class) code in the Batch Header
// * SEC identifies the payment type (product) found within an ACH batch-using a 3-character code
// * The SEC Code pertains to all items within batch
//    * Determines format of the entry detail records
//    * Determines addenda records (required or optional PLUS one or up to 9,999 records)
//    * Determines rules to follow (return timeframes)
// * Some SEC codes require specific data in predetermined fields within the ACH record
type Batcher interface {
	GetHeader() *BatchHeader
	SetHeader(*BatchHeader)
	GetControl() *BatchControl
	SetControl(*BatchControl)
	GetEntries() []*EntryDetail
	AddEntry(*EntryDetail) Batcher
	Build() error
	Validate() error
	ValidateAll() error
}
