package ach

import (
	"fmt"
)

// Batcher abstract the different ACH batch types that can exist in a file.
// Each batch type is defined by SEC (Standard Entry Class) code in the Batch Header
// * SEC identifies the payment type (product) found within an ACH batch-using a 3-character code
// * The SEC Code pertains to all items within batch
//    * Determines format of the entry detail records
//    * Determines addenda records (required or optional PLUS one or up to 9,999 records)
//    * Determines rules to follow (return timeframes)
// 	  * Some SEC codes require specific data in predetermined fields within the ACH record
type Batcher interface {
	GetHeader() *BatchHeader
	SetHeader(*BatchHeader)
	GetControl() *BatchControl
	SetControl(*BatchControl)
	GetEntries() []*EntryDetail
	AddEntry(*EntryDetail)
	Create() error
	Validate() error
	// Category defines if a Forward or Return
	Category() string
}

// BatchError is an Error that describes batch validation issues
type BatchError struct {
	BatchNumber int
	FieldName   string
	Msg         string
}

func (e *BatchError) Error() string {
	return fmt.Sprintf("BatchNumber %d %s %s", e.BatchNumber, e.FieldName, e.Msg)
}

// Errors specific to parsing a Batch container
var (
	// generic messages
	msgBatchHeaderControlEquality     = "header %v is not equal to control %v"
	msgBatchCalculatedControlEquality = "calculated %v is out-of-balance with control %v"
	msgBatchAscending                 = "%v is less than last %v. Must be in ascending order"
	msgBatchFieldInclusion            = "%v is a required field "
	// specific messages for error
	msgBatchOriginatorDNE         = "%v is not “2” for DNE with entry transaction code of 23 or 33"
	msgBatchTraceNumberNotODFI    = "%v in header does not match entry trace number %v"
	msgBatchAddendaIndicator      = "is 0 but found addenda record(s)"
	msgBatchAddendaTraceNumber    = "%v does not match proceeding entry detail trace number %v"
	msgBatchEntries               = "must have Entry Record(s) to be built"
	msgBatchAddendaCount          = "%v addendum found where %v is allowed for batch type %v"
	msgBatchTransactionCodeCredit = "%v a credit is not allowed"
	msgBatchSECType               = "header SEC type code %v for batch type %v"
	msgBatchTypeCode              = "%v found in addenda and expecting %v for batch type %v"
	msgBatchForwardReturn         = "Forward and Return entries found in the same batch"
)
