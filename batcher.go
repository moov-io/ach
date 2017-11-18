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

// BatchParam contains information about the company(Originator) and the type of detail records to follow.
// It is a subset of BatchHeader used for simplifying the client api build process.
type BatchParam struct {
	// ServiceClassCode a three digit code identifies:
	// 	- 200 mixed debits and credits
	// 	- 220 credits only
	// 	- 225 debits only
	ServiceClassCode string `json:"service_class_code"`
	// CompanyName is the legal company name making the transaction.
	CompanyName string `json:"company_name"`
	// CompanyIdentification is assigned by your bank to identify your company. Frequently the federal tax ID
	CompanyIdentification string `json:"company_identification"`
	// StandardEntryClass identifies the payment type (product) found within the batch using a 3-character code
	StandardEntryClass string `json:"standard_entry_class"`
	// CompanyEntryDescription describes the transaction. For example "PAYROLL"
	CompanyEntryDescription string `json:"company_entry_description"`
	// CompanyDescriptiveDate a date chosen to identify the transactions in YYMMDD format.
	CompanyDescriptiveDate string `json:"company_descriptive_date"`
	// Date transactions are to be posted to the receivers’ account in YYMMDD format.
	EffectiveEntryDate string `json:"effective_entry_date"`
	// ODFIIdentification originating ODFI's routing number without the last digit
	ODFIIdentification string `json:"ODFI_identification"`
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
