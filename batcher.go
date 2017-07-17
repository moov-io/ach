package ach

// batcher abstract the different ACH batch types that can exist in a file.
// Each batch type is defined by SEC (Standard Entry Class) code in the Batch Header
// * SEC identifies the payment type (product) found within an ACH batch-using a 3-character code
// * The SEC Code pertains to all items within batch
//    * Determines format of the entry detail records
//    * Determines addenda records (required or optional PLUS one or up to 9,999 records)
//    * Determines rules to follow (return timeframes)
// * Some SEC codes require specific data in predetermined fields within the ACH record
type batcher interface {
	getHeader() BatchHeader
	getControl() BatchControl
	Validate() error
	// many other functions need to be added.
	// ValidateAll() error
	//
}
