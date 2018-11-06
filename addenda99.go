// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package ach

import (
	"fmt"
	"strings"
	"time"
)

// When a Return Entry is prepared, the original Company/Batch Header Record, the original Entry Detail Record,
// and the Company/Batch Control Record are copied for return to the Originator.
//
// The Return Entry is a new Entry. These Entries must be assigned new batch and trace numbers, new identification numbers for the returning institution,
// appropriate transaction codes, etc., as required per format specifications.
//
// See Appendix Four: Return Entries in the NACHA Corporate

var (
	returnCodeDict = map[string]*returnCode{}

	// Error messages specific to Return Addenda
	msgAddenda99ReturnCode = "found is not a valid return code"
)

func init() {
	// populate the returnCode map with lookup values
	returnCodeDict = makeReturnCodeDict()
}

// Addenda99 utilized for Notification of Change Entry (COR) and Return types.
type Addenda99 struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// RecordType defines the type of record in the block. entryAddendaPos 7
	recordType string
	// TypeCode Addenda types code '99'
	TypeCode string `json:"typeCode"`
	// ReturnCode field contains a standard code used by an ACH Operator or RDFI to describe the reason for returning an Entry.
	// Must exist in returnCodeDict
	ReturnCode string `json:"returnCode"`
	// OriginalTrace This field contains the Trace Number as originally included on the forward Entry or Prenotification.
	// The RDFI must include the Original Entry Trace Number in the Addenda Record of an Entry being returned to an ODFI,
	// in the Addenda Record of an 98, within an Acknowledgment Entry, or with an RDFI request for a copy of an authorization.
	OriginalTrace int `json:"originalTrace"`
	// DateOfDeath The field date of death is to be supplied on Entries being returned for reason of death (return reason codes R14 and R15).
	DateOfDeath time.Time `json:"dateOfDeath"`
	// OriginalDFI field contains the Receiving DFI Identification (addenda.RDFIIdentification) as originally included on the forward Entry or Prenotification that the RDFI is returning or correcting.
	OriginalDFI string `json:"originalDFI"`
	// AddendaInformation
	AddendaInformation string `json:"addendaInformation,omitempty"`
	// TraceNumber matches the Entry Detail Trace Number of the entry being returned.
	TraceNumber int `json:"traceNumber,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for ACH to GoLang Converters
	converters
}

// returnCode holds a return Code, Reason/Title, and Description
//
// Table of return codes exists in Part 4.2 of the NACHA corporate rules and guidelines
type returnCode struct {
	Code, Reason, Description string
}

// NewAddenda99 returns a new Addenda99 with default values for none exported fields
func NewAddenda99() *Addenda99 {
	Addenda99 := &Addenda99{
		recordType: "7",
		TypeCode:   "99",
	}
	return Addenda99
}

// Parse takes the input record string and parses the Addenda99 values
func (Addenda99 *Addenda99) Parse(record string) {
	// 1-1 Always "7"
	Addenda99.recordType = "7"
	// 2-3 Defines the specific explanation and format for the addenda information contained in the same record
	Addenda99.TypeCode = record[1:3]
	// 4-6
	Addenda99.ReturnCode = record[3:6]
	// 7-21
	Addenda99.OriginalTrace = Addenda99.parseNumField(record[6:21])
	// 22-27, might be a date or blank
	Addenda99.DateOfDeath = Addenda99.parseSimpleDate(record[21:27])
	// 28-35
	Addenda99.OriginalDFI = Addenda99.parseStringField(record[27:35])
	// 36-79
	Addenda99.AddendaInformation = strings.TrimSpace(record[35:79])
	// 80-94
	Addenda99.TraceNumber = Addenda99.parseNumField(record[79:94])
}

// String writes the Addenda99 struct to a 94 character string
func (Addenda99 *Addenda99) String() string {
	var buf strings.Builder
	buf.Grow(94)
	buf.WriteString(Addenda99.recordType)
	buf.WriteString(Addenda99.TypeCode)
	buf.WriteString(Addenda99.ReturnCode)
	buf.WriteString(Addenda99.OriginalTraceField())
	buf.WriteString(Addenda99.DateOfDeathField())
	buf.WriteString(Addenda99.OriginalDFIField())
	buf.WriteString(Addenda99.AddendaInformationField())
	buf.WriteString(Addenda99.TraceNumberField())
	return buf.String()
}

// Validate verifies NACHA rules for Addenda99
func (Addenda99 *Addenda99) Validate() error {
	if Addenda99.recordType != "7" {
		msg := fmt.Sprintf(msgRecordType, 7)
		return &FieldError{FieldName: "recordType", Value: Addenda99.recordType, Msg: msg}
	}
	if Addenda99.TypeCode == "" {
		return &FieldError{
			FieldName: "TypeCode",
			Value:     Addenda99.TypeCode,
			Msg:       msgFieldInclusion + ", did you use NewAddenda99()?",
		}
	}
	if Addenda99.TypeCode != "99" {
		return &FieldError{FieldName: "TypeCode", Value: Addenda99.TypeCode, Msg: msgAddendaTypeCode}
	}

	_, ok := returnCodeDict[Addenda99.ReturnCode]
	if !ok {
		// Return Addenda requires a valid ReturnCode
		return &FieldError{FieldName: "ReturnCode", Value: Addenda99.ReturnCode, Msg: msgAddenda99ReturnCode}
	}
	return nil
}

// OriginalTraceField returns a zero padded OriginalTrace string
func (Addenda99 *Addenda99) OriginalTraceField() string {
	return Addenda99.numericField(Addenda99.OriginalTrace, 15)
}

// DateOfDeathField returns a space padded DateOfDeath string
func (Addenda99 *Addenda99) DateOfDeathField() string {
	// Return space padded 6 characters if it is a zero value of DateOfDeath
	if Addenda99.DateOfDeath.IsZero() {
		return Addenda99.alphaField("", 6)
	}
	// YYMMDD
	return Addenda99.formatSimpleDate(Addenda99.DateOfDeath)
}

// OriginalDFIField returns a zero padded OriginalDFI string
func (Addenda99 *Addenda99) OriginalDFIField() string {
	return Addenda99.stringField(Addenda99.OriginalDFI, 8)
}

//AddendaInformationField returns a space padded AddendaInformation string
func (Addenda99 *Addenda99) AddendaInformationField() string {
	return Addenda99.alphaField(Addenda99.AddendaInformation, 44)
}

//IATPaymentAmount sets original forward entry payment amount characters 1-10 of underlying AddendaInformation
func (Addenda99 *Addenda99) IATPaymentAmount(s string) {
	Addenda99.AddendaInformation = Addenda99.stringField(s, 10)
}

//IATAddendaInformation sets Addenda Information for IAT return items, characters 10-44 of
// underlying AddendaInformation
func (Addenda99 *Addenda99) IATAddendaInformation(s string) {
	Addenda99.AddendaInformation = Addenda99.AddendaInformation + Addenda99.alphaField(s, 34)
}

//IATPaymentAmountField returns original forward entry payment amount int, characters 1-10 of
// underlying AddendaInformation
func (Addenda99 *Addenda99) IATPaymentAmountField() int {
	return Addenda99.parseNumField(Addenda99.AddendaInformation[0:10])
}

//IATAddendaInformationField returns a space padded AddendaInformation string, characters 10-44 of
// underlying AddendaInformation
func (Addenda99 *Addenda99) IATAddendaInformationField() string {
	return Addenda99.alphaField(Addenda99.AddendaInformation[9:44], 34)
}

// TraceNumberField returns a zero padded traceNumber string
func (Addenda99 *Addenda99) TraceNumberField() string {
	return Addenda99.numericField(Addenda99.TraceNumber, 15)
}

func makeReturnCodeDict() map[string]*returnCode {
	dict := make(map[string]*returnCode)

	codes := []returnCode{
		// Return Reason Codes for RDFIs
		{"R01", "Insufficient Funds", "Available balance is not sufficient to cover the dollar value of the debit entry"},
		{"R02", "Account Closed", "Previously active account has been closed by customer or RDFI"},
		// R03 may not be used to return ARC, BOC or POP entries solely because they do not contain an Individual Name.
		{"R03", "No Account/Unable to Locate Account", "Account number structure is valid and passes editing process, but does not correspond to individual or is not an open account"},
		{"R04", "Invalid Account Number", "Account number structure not valid; entry may fail check digit validation or may contain an incorrect number of digits."},
		{"R05", "Improper Debit to Consumer Account", "A CCD, CTX, or CBR debit entry was transmitted to a Consumer Account of the Receiver and was not authorized by the Receiver"},
		{"R06", "Returned per ODFI's Request", "ODFI has requested RDFI to return the ACH entry (optional to RDFI - ODFI indemnifies RDFI)}"},
		// R07 Prohibited use for ARC, BOC, POP and RCK.
		{"R07", "Authorization Revoked by Customer", "Consumer, who previously authorized ACH payment, has revoked authorization from Originator (must be returned no later than 60 days from settlement date and customer must sign affidavit)"},
		{"R08", "Payment Stopped", "Receiver of a recurring debit transaction has stopped payment to a specific ACH debit. RDFI should verify the Receiver's intent when a request for stop payment is made to insure this is not intended to be a revocation of authorization"},
		{"R09", "Uncollected Funds", "Sufficient book or ledger balance exists to satisfy dollar value of the transaction, but the dollar value of transaction is in process of collection (i.e., uncollected checks) or cash reserve balance below dollar value of the debit entry."},
		{"R10", "Customer Advises Not Authorized", "Consumer has advised RDFI that Originator of transaction is not authorized to debit account (must be returned no later than 60 days from settlement date of original entry and customer must sign affidavit)."},
		{"R11", "Check Truncation Entry Returned", "Used when returning a check safekeeping entry; RDFI should use appropriate field in addenda record to specify reason for return (i.e., 'exceeds dollar limit,' 'stale date,' etc.)."},
		{"R12", "Branch Sold to Another DFI", "Financial institution receives entry destined for an account at a branch that has been sold to another financial institution."},
		{"R13", "RDFI not qualified to participate", "Financial institution does not receive commercial ACH entries"},
		{"R14", "Representative payee deceased or unable to continue in that capacity", "The representative payee authorized to accept entries on behalf of a beneficiary is either deceased or unable to continue in that capacity"},
		{"R15", "Beneficiary or bank account holder", "(Other than representative payee) deceased* - (1) the beneficiary entitled to payments is deceased or (2) the bank account holder other than a representative payee is deceased"},
		{"R16", "Bank account frozen", "Funds in bank account are unavailable due to action by RDFI or legal order"},
		{"R17", "File record edit criteria", "Fields rejected by RDFI processing (identified in return addenda)"},
		{"R18", "Improper effective entry date", "Entries have been presented prior to the first available processing window for the effective date."},
		{"R19", "Amount field error", "Improper formatting of the amount field"},
		{"R20", "Non-payment bank account", "Entry destined for non-payment bank account defined by reg."},
		{"R21", "Invalid company ID number", "The company ID information not valid (normally CIE entries)"},
		{"R22", "Invalid individual ID number", "Individual id used by receiver is incorrect (CIE entries)"},
		{"R23", "Credit entry refused by receiver", "Receiver returned entry because minimum or exact amount not remitted, bank account is subject to litigation, or payment represents an overpayment, originator is not known to receiver or receiver has not authorized this credit entry to this bank account"},
		{"R24", "Duplicate entry", "RDFI has received a duplicate entry"},
		{"R25", "Addenda error", "Improper formatting of the addenda record information"},
		{"R26", "Mandatory field error", "Improper information in one of the mandatory fields"},
		{"R27", "Trace number error", "Original entry trace number is not valid for return entry; or addenda trace numbers do not correspond with entry detail record"},
		{"R28", "Transit routing number check digit error", "Check digit for the transit routing number is incorrect"},
		{"R29", "Corporate customer advises not authorized", "RDFI has bee notified by corporate receiver that debit entry of originator is not authorized"},
		{"R30", "RDFI not participant in check truncation program", "Financial institution not participating in automated check safekeeping application"},
		{"R31", "Permissible return entry (CCD and CTX only)", "RDFI has been notified by the ODFI that it agrees to accept a CCD or CTX return entry"},
		{"R32", "RDFI non-settlement", "RDFI is not able to settle the entry"},
		{"R33", "Return of XCK entry", "RDFI determines at its sole discretion to return an XCK entry; an XCK return entry may be initiated by midnight of the sixtieth day following the settlement date if the XCK entry"},
		{"R34", "Limited participation RDFI", "RDFI participation has been limited by a federal or state supervisor"},
		{"R35", "Return of improper debit entry", "ACH debit not permitted for use with the CIE standard entry class code (except for reversals)"},
		{"R37", "Source Document Presented for Payment (Adjustment Entry)", "The source document to which an ARC, BOC or POP entry relateshas been presented for payment. RDFI must obtain a Written Statement and return the entry within 60 days following Settlement Date"},
		{"R38", "Stop Payment on Source Document (Adjustment Entry)", "A stop payment has been placed on the source document to which the ARC or BOC entry relates. RDFI must return no later than 60 days following Settlement Date. No Written Statement is required as the original stop payment form covers the return"},
		{"R39", "Improper Source Document", "The RDFI has determined the source document used for the ARC, BOC or POP entry to its Receiver’s account is improper."},
		{"R50", "State Law Affecting RCK Acceptance", "RDFI is located in a state that has not adopted Revised Article 4 of the UCC or the RDFI is located in a state that requires all canceled checks to be returned within the periodic statement"},
		{"R51", "Item Related to RCK Entry is Ineligible or RCK Entry is Improper", "The item to which the RCK entry relates was not eligible, Originator did not provide notice of the RCK policy, signature on the item was not genuine, the item has been altered or amount of the entry was not accurately obtained from the item. RDFI must obtain a Written Statement and return the entry within 60 days following Settlement Date"},
		{"R52", "Stop Payment on Item (Adjustment Entry)", "A stop payment has been placed on the item to which the RCK entry relates. RDFI must return no later than 60 days following Settlement Date. No Written Statement is required as the original stop payment form covers the return."},
		{"R53", "Item and RCK Entry Presented for Payment (Adjustment Entry)", "Both the RCK entry and check have been presented forpayment. RDFI must obtain a Written Statement and return the entry within 60 days following Settlement Date"},
		// More return codes will be added when more SEC types are added to the library.
	}
	// populate the map
	for _, code := range codes {
		dict[code.Code] = &code
	}
	return dict
}
