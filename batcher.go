// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

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
	msgBatchCalculatedControlEquality = "calculated %v is out-of-balance with batch control %v"
	msgBatchAscending                 = "%v is less than last %v. Must be in ascending order"
	// specific messages for error
	msgBatchCompanyEntryDescription = "Company entry description %v is not valid for batch type %v"
	msgBatchOriginatorDNE           = "%v is not “2” for DNE with entry transaction code of 23 or 33"
	msgBatchTraceNumberNotODFI      = "%v in header does not match entry trace number %v"
	msgBatchAddendaIndicator        = "is 0 but found addenda record(s)"
	msgBatchAddendaTraceNumber      = "%v does not match proceeding entry detail trace number %v"
	msgBatchEntries                 = "must have Entry Record(s) to be built"
	msgBatchAddendaCount            = "%v addendum found where %v is allowed for batch type %v"
	msgBatchRequiredAddendaCount    = "%v addendum found where %v is required for batch type %v"
	msgBatchTransactionCodeCredit   = "%v a credit is not allowed"
	msgBatchSECType                 = "header SEC type code %v for batch type %v"
	msgBatchServiceClassTranCode    = "%v is not valid for %v"
	msgBatchServiceClassCode        = "Service Class Code %v is not valid for batch type %v"
	msgBatchCategory                = "%v category found in batch with category %v"
	msgBatchAmount                  = "Amount must be less than %v for SEC code %v"
	msgBatchCheckSerialNumber       = "Check Serial Number is required for SEC code %v"
	msgBatchTransactionCode         = "%v is not allowed for batch type %v"
	msgBatchCardTransactionType     = "Card Transaction Type %v is invalid"
	msgBatchTransactionCodeAddenda  = "Addenda not allowed for transaction code %v for batch type %v"
	msgBatchAmountZero              = "%v must be zero for SEC code %v"
	msgBatchAmountNonZero           = "%v must be non-zero for SEC code %s"
	msgBatchAddenda                 = "%v not allowed for category %v for batch type %v"
	msgBatchADVCount                = "There can be a maximum of %v ADV Sequence Numbers (ADV Entry Detail Records)"
)
